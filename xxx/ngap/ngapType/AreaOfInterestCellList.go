// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct AreaOfInterestCellList */
/* AreaOfInterestCellItem */
type AreaOfInterestCellList struct {
	List []AreaOfInterestCellItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
