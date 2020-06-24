// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CompletedCellsInTAI_EUTRA */
/* CompletedCellsInTAIEUTRAItem */
type CompletedCellsInTAIEUTRA struct {
	List []CompletedCellsInTAIEUTRAItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
