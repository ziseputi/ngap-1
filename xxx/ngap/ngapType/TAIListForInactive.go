// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct TAIListForInactive */
/* TAIListForInactiveItem */
type TAIListForInactive struct {
	List []TAIListForInactiveItem `vht:"valueExt,sizeLB:1,sizeUB:16"`
}