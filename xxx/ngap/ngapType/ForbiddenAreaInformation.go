// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct ForbiddenAreaInformation */
/* ForbiddenAreaInformationItem */
type ForbiddenAreaInformation struct {
	List []ForbiddenAreaInformationItem `vht:"valueExt,sizeLB:1,sizeUB:16"`
}
