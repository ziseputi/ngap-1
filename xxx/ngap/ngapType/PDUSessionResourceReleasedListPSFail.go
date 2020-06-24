// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceReleasedListPSFail */
/* PDUSessionResourceReleasedItemPSFail */
type PDUSessionResourceReleasedListPSFail struct {
	List []PDUSessionResourceReleasedItemPSFail `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
