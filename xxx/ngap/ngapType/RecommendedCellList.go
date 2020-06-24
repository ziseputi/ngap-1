// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct RecommendedCellList */
/* RecommendedCellItem */
type RecommendedCellList struct {
	List []RecommendedCellItem `vht:"valueExt,sizeLB:1,sizeUB:16"`
}
