// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CancelledCellsInEAI_EUTRA */
/* CancelledCellsInEAIEUTRAItem */
type CancelledCellsInEAIEUTRA struct {
	List []CancelledCellsInEAIEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
