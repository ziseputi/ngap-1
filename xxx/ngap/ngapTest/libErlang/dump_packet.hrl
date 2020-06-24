%% @author Tran Hoan <tvhoan88@gmail.com>
%% Created on 7:20:14 AM May 3, 2014
%% Purpose

%% maintainer

-ifndef(dump_packet).
-define(dump_packet,true).

-define(MAC1,<<0,0,16#5E,1,1,1>>).
-define(MAC2,<<0,0,16#5E,2,2,2>>).
-define(IP1,<<1,1,1,1>>).
-define(IP2,<<2,2,2,2>>).
-define(IPv4ADDR,<<1,1,1,1,2,2,2,2>>).
-define(SCCP1,<<16#43,1000:16/little,142>>).
-define(SCCP2,<<16#43,2000:16/little,142>>).
-define(SUDP,1111).
-define(DUDP,2222).

-record(ether_packet,{
	saddr = ?MAC1		::	binary(),			%% Source MAC
	daddr = ?MAC2		::  binary(),			%% Destination MAC
	type  = 16#0800		::	non_neg_integer(),	%% Type, default is IPv4 
	data				::	binary()			%% Layer 3 data
}).

-record(mtp2_packet,{
	bseqn = 61			::	0..127,
	bind  = 1			:: 	0..1,
	fseqn = 68			::	0..127,
	find  = 1			::	0..1,
	lind  = 63			::	0..63,
	data				::	binary()
}).

-record(ipv4_packet,{
	ds		= 0	::	byte(),				%% Differentiated Services
	id		= 1	::	non_neg_integer(),	%% Identification
	flags	= 0	::	byte(),				%% Flags
	frag_off= 0	::	non_neg_integer(),	%% Fragment offset
	ttl	= 255	::	byte(),				%% Time To Live
	proto = 132	::	byte(),				%% Protocol, default SCTP
	crc			::  non_neg_integer(),	%% CRC
	saddr = ?IP1::  binary(),			%% Source IP
	daddr = ?IP2::	binary(),			%% Destination IP
	data		::	binary()			%% Layer 4 data
}).

-record(sctp_data_packet,{
	sport	= 2905			:: 1..65535,		 %% Source Port
	dport	= 2905			:: 1..65535,		 %% Destionation Port
	vtag	= 1				:: non_neg_integer(),%% Verification Tag
	csum	= 16#ffffffff	:: non_neg_integer(),%% Check sum
	chunkType = 0			:: byte(),			 %% Chunk Type
	chunkFlag = 3			:: byte(),			 %% Chunk Flag
	tsn		= 16#ffffffff 	:: non_neg_integer(),%% TSN
	si		= 2				:: byte(),			 %% Stream Identifier
	seqn	= 16#ffff		:: non_neg_integer(),%% Stream Sequence Number
	proto	= 3				:: non_neg_integer(),%% Protocol Identifier
	data					:: binary()
}).

-record(m3ua_packet,{
	version		= 	1		:: byte(),
	reserved	=	0		:: byte(),
	class		= 	1		:: byte(),
	type		=	1		:: byte(),
	data					:: binary()
}).

-record(m3ua_user_packet,{
	opc	= 9416		:: non_neg_integer(),	%% M
	dpc	= 2805		:: non_neg_integer(),	%% M
	si	= 13		:: byte(),				%% M
	ni	= 3			:: byte(),				%% M
	mp	= 0			:: byte(),				%% M
	sls	= 1			:: byte(),				%% M
	data	:: binary()						%% M
}).

-record(sccp_udt_packet,{
	msgType	=	9	::	byte(),				%% M
	protoClass = 0	::	byte(),				%% M
	called_party	::	binary(),			%% M
	calling_party	::	binary(),			%% M
	data			::	binary()			%% M
}).

-record(sccp_cr_packet,{
	msgType	= 1				::	byte(),				%% M 	Message Type
	sourceRef = 1			::	non_neg_integer(),	%% M
	protoClass = 2			::	byte(),				%% M
	called_party = ?SCCP1	::	binary(),			%% M
	calling_party = ?SCCP2	::	binary(),			%% O
	data					::	binary()			%% O
}).

-record(sccp_dt1_packet,{
	msgType	= 6				::	byte(),				%% M	Message Type
	destRef = 1				::	non_neg_integer(),	%% M
	segmenting = 0			::	byte(),				%% M
	data					::	binary()			%% M
}).

-record(sccp_cc_packet,{
	msgType = 2				::	byte(),				%% M
	destRef = 1				::	non_neg_integer(),	%% M
	sourceRef = 1			::	non_neg_integer(),	%% M
	protoClass	= 2			::	byte(),				%% M
	data					::	binary()			%% O				
}).

-record(udp_packet,{
	sourcePort = ?SUDP		::  integer(),
	destPort   = ?DUDP		::  integer(),
	checkSum   = 0			::  integer(),
	data					::  binary()
}).

-record(cooked_capture,{
	type	   = 0			::  integer(),
	add_type   = 1			::	integer(),
	soucreIP				::	binary(),
	proto	   = 16#0800	::  integer(),
	data					::	binary()
}).

-define(sccp_EndOfOptional,0).
-define(sccp_DestinationLocalRef,1).
-define(sccp_SourceLocalRef,2).
-define(sccp_CalledPartyAddr,3).
-define(sccp_CallingPartyAddr,4).
-define(sccp_ProtocolClass,5).
-define(sccp_Data,15).
-define(sccp_Segmenting,6).

-endif.


