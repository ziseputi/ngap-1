// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct TNLInformationList */
/* TNLInformationItem */
type TNLInformationList struct {
	List []TNLInformationItem `vht:"valueExt,sizeLB:1,sizeUB:4"`
}
