%% Author: Tran Hoan
%% Created: Jul 18, 2012
%% Description: hex string utility
-module(hex_dump).

%%
%% Include files
%%

%%
%% Exported Functions
%%
-export([dump/1,
		 dumpFd/2,
		 str2hex/1,
		 hex2str/1,
		 time2hex/0,
		 node2hex/0,
		 pid2hex/0,
		 parse_hex_str/1,
		 reformat_str/1,
		 hex_value/1,
		 compare/2]).

%%
%% API Functions
%%

dump(Term) when is_binary(Term) ->
	Str = dump_bin(Term," ",""),
	io:format("~s~n", [Str]).

dump_bin(<<>>,_Sep,Str) ->
	Str;
dump_bin(<<Oct:8>>,Sep,Str) ->
	Hex = io_lib:format("~2.16.0B", [Oct]),
	string:join([Str,Hex], Sep);
dump_bin(<<Oct:8,Rest/binary>>,Sep,"") ->
	Hex = io_lib:format("~2.16.0B", [Oct]),
	dump_bin(Rest,Sep,Hex);
dump_bin(<<Oct:8,Rest/binary>>,Sep,Str) ->
	Hex = io_lib:format("~2.16.0B", [Oct]),
	NewStr = string:join([Str,Hex], Sep),
	dump_bin(Rest, Sep,NewStr).

dumpFd(Fd,BufLen) ->
	case file:read(Fd,BufLen) of
		{ok,Data} ->
			Str = dump_bin(Data,"",""),
			io:format("~s",[Str]),
			dumpFd(Fd, BufLen);
		_ ->
			io:format("~n")
	end.

%% -------------------------------------------
%% time to hex
%% -------------------------------------------
time2hex() ->
	{MegaSec,Sec,MicroSec} = now(),
	Time = MegaSec * 1000 + Sec,
	lists:flatten(io_lib:format("~8.16.0B~8.16.0B", [Time,MicroSec])).

%% ------------------------------------------
%% Node to hex
%% ------------------------------------------
node2hex() ->
	Str = atom_to_list(node()),
	lists:flatten(str2hex(Str)).

%% -----------------------------------------
%% Pid to hex
%% -----------------------------------------
pid2hex() ->
	Str = pid_to_list(self()),
	StrLen = string:len(Str),
	Str1 = string:sub_string(Str,2,StrLen -1),
	lists:flatten(str2hex(Str1)).

%% --------------------------------------------
%% str2hex
%% --------------------------------------------
str2hex(Str) ->
	[io_lib:format("~2.16.0B", [X]) || X <- Str].

%% --------------------------------------------
%% hex2str
%% --------------------------------------------
hex2str(Oct) ->
  io_lib:format("~2.16.0B", [Oct]).
%% ---------------------------------------------
%% parse_hex_str
%% ---------------------------------------------
parse_hex_str(Str) ->
	Str1 = reformat_str(Str),
	Len = string:len(Str1),
	parse_hex_str(Str1, <<>>, 1, Len).

parse_hex_str(Str,Bin,Index,Len) when Index < Len ->
	NewIndex = Index +2,
	Index1 = Index + 1,
	Oct = string:sub_string(Str, Index, Index1),
	OctBin = hex_value(Oct),
	NewBin = <<Bin/binary,OctBin:8>>,
	parse_hex_str(Str, NewBin, NewIndex, Len);
parse_hex_str(_Str,Bin,Index,Len) when Index >= Len ->
	Bin.

%% ---------------------------------------------
%% reformat string
%% ---------------------------------------------

reformat_str(Str) ->
	Str1 = string:strip(Str),
	Str2 = string:to_upper(Str1),
	Fun = fun
			 (Elem, Acc) ->
				  if
					  Elem >= $0 andalso Elem =< $9 ->
						  string:join([Acc,[Elem]],"");
					  Elem >= $A andalso Elem =< $F ->
						  string:join([Acc,[Elem]],"");
					  true ->
						  Acc
				  end		  
		  end,
	lists:foldl(Fun,"",Str2).

%% ----------------------------
%% hex_value
%% ----------------------------
hex_value(Str) ->
	Fun = fun
			 (Elem, Acc) ->
				  if
					  Elem >= $0 andalso Elem =< $9 ->
						  (Acc bsl 4) + (Elem - $0);
					  Elem >= $A andalso Elem =< $F ->
						  (Acc bsl 4) + (Elem - $A + 10);
					  true ->
						  Acc
				  end
		  end,
	lists:foldl(Fun, 0, Str).

%% -----------------------------------
%% Compare two erlang binary
%% -----------------------------------
compare(Bin,Bin1) ->
	List = binary_to_list(Bin),
	List1 = binary_to_list(Bin1),
	S1 = length(List),
	S2 = length(List1),
	Len = erlang:min(S1,S2),
	compare(Len, 1,0,List,List1).

compare(Len,Idx,Error,List1,List2) when Idx =< Len ->
	E1 = lists:nth(Idx,List1),
	E2 = lists:nth(Idx,List2),
	NewIdx = Idx+1,
	if 
		E1 == E2 ->
			io:format("~2.16.0B   ok   ~2.16.0B~n",[E1,E2]),
			compare(Len, NewIdx,Error,List1,List2);
		true ->
			io:format("~2.16.0B  error ~2.16.0B~n",[E1,E2]),
			NewError = Error + 1,
			compare(Len, NewIdx,NewError,List1,List2)
	end;

compare(Len,Idx,Error,List1,List2) when Idx > Len ->
	Size1 = length(List1),
	Size2 = length(List2),
	io:format("Number of error ~p~n",[Error]),
	if 
		Size1 /= Size2 ->
			io:format("Error  size ~p vs ~p ~n",[Size1,Size2]);
		true ->
			ok
	end.