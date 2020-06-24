// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CellIDCancelledEUTRA */
/* CellIDCancelledEUTRAItem */
type CellIDCancelledEUTRA struct {
	List []CellIDCancelledEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
