// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct ForbiddenTACs */
/* TAC */
type ForbiddenTACs struct {
	List []TAC `vht:"sizeLB:1,sizeUB:4096"`
}
