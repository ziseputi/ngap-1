// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct TAIBroadcastEUTRA */
/* TAIBroadcastEUTRAItem */
type TAIBroadcastEUTRA struct {
	List []TAIBroadcastEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
