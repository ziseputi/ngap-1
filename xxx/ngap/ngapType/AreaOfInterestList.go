// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct AreaOfInterestList */
/* AreaOfInterestItem */
type AreaOfInterestList struct {
	List []AreaOfInterestItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
