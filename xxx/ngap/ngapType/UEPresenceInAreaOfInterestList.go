// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct UEPresenceInAreaOfInterestList */
/* UEPresenceInAreaOfInterestItem */
type UEPresenceInAreaOfInterestList struct {
	List []UEPresenceInAreaOfInterestItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
