// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct TAIBroadcastNR */
/* TAIBroadcastNRItem */
type TAIBroadcastNR struct {
	List []TAIBroadcastNRItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
