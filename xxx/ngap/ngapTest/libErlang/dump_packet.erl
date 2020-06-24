%% @author Tran Hoan <tvhoan88@gmail.com>
%% Created on 9:52:24 PM May 2, 2014
%% Purpose

%% maintainer

-module(dump_packet).

-include("dump_packet.hrl").

-export([encode/1,
		 encode/2]).

encode(Raw) when is_binary(Raw) ->
	Raw;

encode(#ether_packet{saddr = MAC1,daddr = MAC2,type = Type,data = Bin}) ->
	<<MAC2/binary,MAC1/binary,Type:16,Bin/binary>>;

encode(#ipv4_packet{ds = DS,id = Id,flags = Flag,frag_off = Frag,ttl = TTL,proto = Proto,
					saddr = Sip,daddr = Dip,data = Bin}) ->
	IpLen = byte_size(Bin) + 20,
	CRC = ipv4_crc(<<16#45,DS,IpLen:16,Id:16,Flag:4,Frag:12,TTL,Proto,Sip/binary,Dip/binary>>),
	<<16#45,DS,IpLen:16,Id:16,Flag:4,Frag:12,TTL,Proto,CRC:16,Sip/binary,Dip/binary,Bin/binary>>;

encode(#sctp_data_packet{sport = Hport,dport = Pport,vtag = Vtag,csum = CheckSum,chunkType = ChunkType,
						 chunkFlag = ChunkFlag,tsn = TSN,si = SI,seqn = Seqn,proto = Proto,data = Bin}) ->
	SCTPLen = byte_size(Bin) + 16,
	<<Hport:16/little,Pport:16/little,Vtag:32,CheckSum:32,ChunkType,ChunkFlag,SCTPLen:16,
	  TSN:32,SI:16,Seqn:16,Proto:32,Bin/binary>>;

encode(#mtp2_packet{bseqn = BSeq,bind = BInd,fseqn = FSeq,find = FInd,lind = LInd,data = Bin}) ->
	<<BInd:1,BSeq:7,FInd:1,FSeq:7,0:2,LInd:6,Bin/binary>>;

encode(#m3ua_packet{version = Vsn,reserved = Res,class = Class,type = Type,data = Bin}) ->
	Len = byte_size(Bin) + 8,
	case Len rem 4 of
		0 ->
			<<Vsn,Res,Class,Type,Len:32,Bin/binary>>;
		1 ->
			Len1 = Len + 3,
			<<Vsn,Res,Class,Type,Len1:32,Bin/binary,0,0,0>>;
		2 ->
			Len1 = Len + 2,
			<<Vsn,Res,Class,Type,Len1:32,Bin/binary,0,0>>;
		_ ->
			Len1 = Len + 1,
			<<Vsn,Res,Class,Type,Len1:32,Bin/binary,0>>
	end;

encode(#m3ua_user_packet{opc = Opc,dpc = Dpc,si = SI,ni = NI,mp = MP,sls = SLS, data = Bin}) ->
	Len = byte_size(Bin) + 16,
	<<16#0210:16,Len:16,Opc:32,Dpc:32,SI,NI,MP,SLS,Bin/binary>>;

encode(#sccp_cr_packet{msgType = Type,sourceRef = Ref,protoClass = Class,called_party = Called,
							  calling_party = Calling,data = Bin}) ->
	case lists:foldl(fun({C,P,V},{FM,VM,O,PList}) ->
							 encode_sccp_param(C,P,V,FM,VM,O,PList)
					 end,{<<>>,<<>>,<<>>,[]},
					 [{message_type,mandatory,Type},
					  {?sccp_SourceLocalRef,mandatory,Ref},
					  {?sccp_ProtocolClass,mandatory,Class},
					  {?sccp_CalledPartyAddr,mandatory,Called},
					  {?sccp_CallingPartyAddr,optional,Calling},
					  {?sccp_Data,optional,Bin}]) of
		{FM,VM,<<>>,PList} ->
			PBin = iolist_to_binary(lists:reverse(PList)),
			<<FM/binary,PBin/binary,VM/binary>>;
		{FM,VM,O,PList} ->
			P1 = [X + 1 || X <- PList],
			PO = byte_size(VM) + 1,
			PBin = iolist_to_binary(lists:reverse([PO|P1])),
			<<FM/binary,PBin/binary,VM/binary,O/binary,?sccp_EndOfOptional>>;
		Exp ->
			erlang:error({encode,Exp})
	end;

encode(#sccp_udt_packet{msgType = Type,protoClass = Class,called_party = Called,
						calling_party = Calling,data = Bin}) ->
	case lists:foldl(fun({C,P,V},{FM,VM,O,PList}) ->
							 encode_sccp_param(C,P,V,FM,VM,O,PList)
					 end,{<<>>,<<>>,<<>>,[]},
					 [{message_type,mandatory,Type},
					  {?sccp_ProtocolClass,mandatory,Class},
					  {?sccp_CalledPartyAddr,mandatory,Called},
					  {?sccp_CallingPartyAddr,mandatory,Calling},
					  {?sccp_Data,mandatory,Bin}]) of
		{FM,VM,<<>>,PList} ->
			PBin = iolist_to_binary(lists:reverse(PList)),
			<<FM/binary,PBin/binary,VM/binary>>;
		{FM,VM,O,PList} ->
			P1 = [X + 1 || X <- PList],
			PO = byte_size(VM) + 1,
			PBin = iolist_to_binary(lists:reverse([PO|P1])),
			<<FM/binary,PBin/binary,VM/binary,O/binary,?sccp_EndOfOptional>>;
		Exp ->
			erlang:error({encode,Exp})
	end;

encode(#udp_packet{sourcePort = SPort,destPort = DPort,checkSum = CS,data = Data}) ->
	Len = byte_size(Data) + 8,
	<<SPort:16,DPort:16,Len:16,CS:16,Data/binary>>;

encode(#sccp_cc_packet{msgType = Type,destRef = DRef,sourceRef = SRef,protoClass = Class,data = Bin}) ->
	case lists:foldl(fun({C,P,V},{FM,VM,O,PList}) ->
							 encode_sccp_param(C,P,V,FM,VM,O,PList)
					 end,{<<>>,<<>>,<<>>,[]},
					 [{message_type,mandatory,Type},
					  {?sccp_DestinationLocalRef,mandatory,DRef},
					  {?sccp_SourceLocalRef,mandatory,SRef},
					  {?sccp_ProtocolClass,mandatory,Class},
					  {?sccp_Data,optional,Bin}]) of
		{FM,VM,<<>>,PList} ->
			PBin = iolist_to_binary(lists:reverse(PList)),
			<<FM/binary,PBin/binary,VM/binary,?sccp_EndOfOptional>>;
		{FM,VM,O,PList} ->
			P1 = [X + 1 || X <- PList],
			PO = byte_size(VM) + 1,
			PBin = iolist_to_binary(lists:reverse([PO|P1])),
			<<FM/binary,PBin/binary,VM/binary,O/binary,?sccp_EndOfOptional>>;
		Exp ->
			erlang:error({encode,Exp})
	end;

encode(#sccp_dt1_packet{msgType = Type,destRef = Ref,segmenting = Seq,data = Bin}) ->
	case lists:foldl(fun({C,P,V},{FM,VM,O,PList}) ->
							 encode_sccp_param(C,P,V,FM,VM,O,PList)
					 end,{<<>>,<<>>,<<>>,[]},
					 [{message_type,mandatory,Type},
					  {?sccp_DestinationLocalRef,mandatory,Ref},
					  {?sccp_Segmenting,mandatory,Seq},
					  {?sccp_Data,mandatory,Bin}]) of
		{FM,VM,<<>>,PList} ->
			PBin = iolist_to_binary(lists:reverse(PList)),
			<<FM/binary,PBin/binary,VM/binary>>;
		{FM,VM,O,PList} ->
			P1 = [X + 1 || X <- PList],
			PO = byte_size(VM) + 1,
			PBin = iolist_to_binary(lists:reverse([PO|P1])),
			<<FM/binary,PBin/binary,VM/binary,O/binary,?sccp_EndOfOptional>>;
		Exp ->
			erlang:error({encode,Exp})
	end;

encode(X) ->
	erlang:error({not_implemented,X}).


encode_sccp_param(message_type,mandatory,V,FM,VM,O,P) ->
	{<<FM/binary,V>>,VM,O,P};

encode_sccp_param(Code,mandatory,V,FM,VM,O,P) when Code == ?sccp_ProtocolClass;
												   Code == ?sccp_Segmenting ->
	{<<FM/binary,V>>,VM,O,P};

encode_sccp_param(Code,mandatory,V,FM,VM,O,P) when (Code == ?sccp_SourceLocalRef orelse
													Code == ?sccp_DestinationLocalRef) andalso
													is_integer(V) ->
	{<<FM/binary,V:24/little>>,VM,O,P};

encode_sccp_param(Code,mandatory,V,FM,VM,O,P) when (Code == ?sccp_CalledPartyAddr orelse
												Code == ?sccp_Data orelse
												Code == ?sccp_CallingPartyAddr) andalso
												is_binary(V) ->
	Len = byte_size(V),
	P1 = [X + 1 || X <- P],
	Pn = byte_size(VM) + 1,
	{FM,<<VM/binary,Len,V/binary>>,O,[Pn|P1]};

encode_sccp_param(Code,optional,V,FM,VM,O,P) when (Code == ?sccp_SourceLocalRef orelse
													Code == ?sccp_DestinationLocalRef) andalso
												is_integer(V) ->
	{FM,VM,<<O/binary,Code,3,V:24/little>>,P};

encode_sccp_param(Code,optional,V,FM,VM,O,P) when (Code == ?sccp_CalledPartyAddr orelse
												Code == ?sccp_Data orelse
												Code == ?sccp_CallingPartyAddr) andalso
												is_binary(V) ->
	Len = byte_size(V),
	{FM,VM,<<O/binary,Code,Len,V/binary>>,P};

encode_sccp_param(_,optional,nil,FM,VM,O,P) ->
	{FM,VM,O,P};

encode_sccp_param(_,optional,undefined,FM,VM,O,P) ->
	{FM,VM,O,P};

encode_sccp_param(X,_,_,_,_,_,_) ->
	erlang:error({not_implemented,X}).

ipv4_crc(<<N1:16,N2:16,N3:16,N4:16,N5:16,N6:16,N7:16,N8:16,N9:16>>) ->
	Temp = N1 + N2 + N3 + N4 + N5 + N6 + N7 + N8 + N9,
	T0 = ((Temp bsr 16) band 15),
	T1 = ((Temp bsr 12) band 15),
	T2 = ((Temp bsr 8) band 15),
	T3 = ((Temp bsr 4) band 15),
	T4 = (Temp band 15),
	
	Temp1 = ((T1 bsl 12) bor (T2 bsl 8) bor (T3 bsl 4) bor T4) + T0,
	T11 = ((bnot((Temp1 bsr 12) band 15)) band 15),
	T21 = ((bnot((Temp1 bsr 8) band 15)) band 15),
	T31 = ((bnot((Temp1 bsr 4) band 15)) band 15),
	T41 = ((bnot((Temp1) band 15)) band 15),

	((T11 bsl 12) bor (T21 bsl 8) bor (T31 bsl 4) bor T41).

encode(#ipv4_packet{ds = DS,id = Id,flags = Flag,frag_off = Frag,ttl = TTL,proto = Proto,
					saddr = Sip,daddr = Dip,data = Bin,crc = CRC},padding) ->
	IpLen = byte_size(Bin) + 20,
	<<16#45,DS,IpLen:16,Id:16,Flag:4,Frag:12,TTL,Proto,CRC:16,Sip/binary,Dip/binary,Bin/binary>>;
encode(#sctp_data_packet{sport = Hport,dport = Pport,vtag = Vtag,csum = CheckSum,chunkType = ChunkType,
						 chunkFlag = ChunkFlag,tsn = TSN,si = SI,seqn = Seqn,proto = Proto,data = Bin},bin) ->
	SCTPLen = byte_size(Bin) + 16,
	<<Pport/binary,Hport/binary,Vtag:32,CheckSum:32,ChunkType,ChunkFlag,SCTPLen:16,
	  TSN:32,SI:16,Seqn:16,Proto:32,Bin/binary>>;
encode(#cooked_capture{proto = Proto,
					   soucreIP = IP,
					   add_type = ADD,
					   type = Type,
					   data = Data},_) ->
	Len = byte_size(IP),
	<<Type:16,ADD:16,Len:16,IP/binary,Proto:16,Data/binary>>;
encode(_,_) -> <<>>.