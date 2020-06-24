// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CompletedCellsInTAI_NR */
/* CompletedCellsInTAINRItem */
type CompletedCellsInTAINR struct {
	List []CompletedCellsInTAINRItem `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
