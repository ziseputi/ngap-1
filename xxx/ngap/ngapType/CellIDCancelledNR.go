// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CellIDCancelledNR */
/* CellIDCancelledNRItem */
type CellIDCancelledNR struct {
	List []CellIDCancelledNRItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}