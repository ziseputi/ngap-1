// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CancelledCellsInTAI_EUTRA */
/* CancelledCellsInTAIEUTRAItem */
type CancelledCellsInTAIEUTRA struct {
	List []CancelledCellsInTAIEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
