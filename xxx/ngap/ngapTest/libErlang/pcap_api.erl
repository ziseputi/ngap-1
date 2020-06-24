%% @author Tran Hoan <tvhoan88@gmail.com>
%% Created on 9:51:10 PM May 2, 2014
%% Purpose

%% maintainer

-module(pcap_api).

-include("dump_packet.hrl").

-export([new_pcapng/1,
		 new_pcap/1,
		 new_pcap/2,
		 append_pcap/5
		]).

-export([test/0]).

-record(pcap_header,{
	magic_number 	= 16#a1b2c3d4,
	version_major 	= 2,
	version_minor	= 4,
	thiszone		= 0,
	sigfigs			= 0,
	snaplen			= 16#ffff,
	network			= 1 
}).

-record(pcap_record,{
	ts_sec		:: non_neg_integer(),	%% timestamp seconds
	ts_usec		:: non_neg_integer(),	%% timestamp microseconds
	incl_len	:: non_neg_integer(),	%% number of octets of packet saved in file
	orig_len	:: non_neg_integer(),	%% actual length of packet
	data		:: binary()				%% Packet data
}).

new_pcapng(_) ->
	ok.

new_pcap(Name) ->
	new_pcap(Name,ether).

new_pcap(Name,mtp2) ->
	file:write_file(Name,encode(#pcap_header{network = 140}),[write]);
new_pcap(Name,_) ->
	file:write_file(Name,encode(#pcap_header{}),[write]).

append_pcap(Name,Ts,UTs,Packet,_Opts) when is_binary(Packet) ->
	R = #pcap_record{ts_sec = Ts,ts_usec = UTs,data = Packet},
	file:write_file(Name,encode(R),[append]);

append_pcap(Name,Ts,UTs,Packet,Opts) ->
	R = #pcap_record{ts_sec = Ts,ts_usec = UTs,data = encode_packet(Packet,Opts)},
	file:write_file(Name,encode(R),[append]).

encode(#pcap_header{magic_number = Mg,version_major = Vma,version_minor = Vmi,thiszone = Zone,
					sigfigs = Sig,snaplen = Snap,network = Network}) ->
	<<Mg:32,Vma:16,Vmi:16,Zone:32,Sig:32,Snap:32,Network:32>>;

encode(#pcap_record{ts_sec = Sec,ts_usec = Usec,data = Bin}) ->
	Len = byte_size(Bin),
	<<Sec:32,Usec:32,Len:32,Len:32,Bin/binary>>;

encode(X) ->
	erlang:error({not_implemented,X}).


encode_packet(Packet,_) when is_binary(Packet) ->
	Packet;

encode_packet(Packet,_) when is_record(Packet,mtp2_packet) ->
	dump_packet:encode(Packet);

encode_packet(Packet,_) when is_record(Packet,ether_packet) ->
	dump_packet:encode(Packet);

encode_packet(Packet,Opts) when is_record(Packet,ipv4_packet) ->
	IP = dump_packet:encode(Packet),
	encode_packet(#ether_packet{data = IP},Opts);

encode_packet(Packet,Opts) when is_record(Packet,sctp_data_packet) ->
	SCTP = dump_packet:encode(Packet),
	encode_packet(#ipv4_packet{data = SCTP},Opts);

encode_packet(Packet,_) when is_record(Packet,udp_packet) ->
	SCTP = dump_packet:encode(Packet),
	IP = dump_packet:encode(#ipv4_packet{ds = 0,
							   id = 23466,
							   flags = 4,
							   frag_off = 0,
							   ttl = 64,
							   proto = 17,
							   saddr = <<16#78,16#78,16#78,16#29>>,
							   daddr = <<16#78,16#78,16#78,16#14>>,
							   data = SCTP}),
	Ether = dump_packet:encode(#ether_packet{saddr = <<16#08,16#00,16#27,16#b4,16#bb,16#b7>>,
											 daddr = <<16#b0,16#83,16#fe,16#6a,16#66,16#c8>>,
											 type = 16#0800,data = IP}),
	Ether;

encode_packet(Packet,Opts) when is_record(Packet,m3ua_packet) ->
	M3UA = dump_packet:encode(Packet),
	encode_packet(#sctp_data_packet{data = M3UA},Opts);


encode_packet(Packet,Opts) when is_record(Packet,m3ua_user_packet) ->
	M3UA = dump_packet:encode(Packet),
	encode_packet(#m3ua_packet{data = M3UA},Opts);

encode_packet(Packet,Opts) when is_record(Packet,sccp_cr_packet);
						   is_record(Packet,sccp_cc_packet);
						   is_record(Packet,sccp_udt_packet);
						   is_record(Packet,sccp_dt1_packet)->
	SCCP = dump_packet:encode(Packet),
	encode_packet(#m3ua_user_packet{si = 3,data = SCCP,opc = get_opc(Opts), dpc = get_dpc(Opts)},Opts);

encode_packet(X,_) ->
	erlang:error({not_implemented,X}).

get_opc(Opts) ->
	case lists:keyfind(opc,1,Opts) of
		{opc,OPC} ->
			OPC;
		_ ->
			1000
	end.

get_dpc(Opts) ->
	case lists:keyfind(dpc,1,Opts) of
		{dpc,DPC} ->
			DPC;
		_ ->
			2000
	end.


test() ->
	new_pcap("/home/tran/tmp/edump.pcap",mtp2),
	append_pcap("/home/tran/tmp/edump.pcap",0,0,
				#mtp2_packet{data = hex_dump:parse_hex_str("85 E6 05 10 31 24 00 01 10 40 00 0A 00 02 09 07 83 10 60 09 06 50 0F 0A 08 84 13 48 69 23 53 12 07 31 02 00 00 03 1F 71 08 A8 4F 54 67 6F 73 43 FF 6D 13 A0 10 02 31 40 20 31 14 90 25 F3 F0 FF 0F 00 00 00 80 FF 20 54 00 30 80 81 13 30 31 32 30 31 33 30 34 30 32 31 33 34 31 30 39 35 32 33 A2 80 80 05 34 34 35 37 36 A1 80 80 05 56 49 45 54 54 A1 80 81 08 84 13 48 89 20 10 02 03 00 00 00 00 00 00 83 05 36 33 37 33 34 84 01 00 85 03 80 90 A3 A7 80 81 01 00 00 00 00 00 39 04 78 83 31 C0 00")},
				[]).