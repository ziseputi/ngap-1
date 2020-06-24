// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct UEHistoryInformation */
/* LastVisitedCellItem */
type UEHistoryInformation struct {
	List []LastVisitedCellItem `vht:"valueExt,sizeLB:1,sizeUB:16"`
}
