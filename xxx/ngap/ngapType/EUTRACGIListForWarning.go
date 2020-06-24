// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct EUTRA_CGIListForWarning */
/* EUTRACGI */
type EUTRACGIListForWarning struct {
	List []EUTRACGI `vht:"valueExt,sizeLB:1,sizeUB:65535"`
}
