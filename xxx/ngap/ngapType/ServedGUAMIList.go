// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct ServedGUAMIList */
/* ServedGUAMIItem */
type ServedGUAMIList struct {
	List []ServedGUAMIItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
