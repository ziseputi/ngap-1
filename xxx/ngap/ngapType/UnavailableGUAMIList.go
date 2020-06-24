// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct UnavailableGUAMIList */
/* UnavailableGUAMIItem */
type UnavailableGUAMIList struct {
	List []UnavailableGUAMIItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
