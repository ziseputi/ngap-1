// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct BroadcastPLMNList */
/* BroadcastPLMNItem */
type BroadcastPLMNList struct {
	List []BroadcastPLMNItem `vht:"valueExt,sizeLB:1,sizeUB:12"`
}