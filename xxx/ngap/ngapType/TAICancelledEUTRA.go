// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct TAICancelledEUTRA */
/* TAICancelledEUTRAItem */
type TAICancelledEUTRA struct {
	List []TAICancelledEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
