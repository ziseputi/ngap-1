// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct DRBsSubjectToStatusTransferList */
/* DRBsSubjectToStatusTransferItem */
type DRBsSubjectToStatusTransferList struct {
	List []DRBsSubjectToStatusTransferItem `vht:"valueExt,sizeLB:1,sizeUB:32"`
}
