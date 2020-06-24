// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct AllowedTACs */
/* TAC */
type AllowedTACs struct {
	List []TAC `vht:"sizeLB:1,sizeUB:16"`
}
