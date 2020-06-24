// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CompletedCellsInEAI_EUTRA */
/* CompletedCellsInEAIEUTRAItem */
type CompletedCellsInEAIEUTRA struct {
	List []CompletedCellsInEAIEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
