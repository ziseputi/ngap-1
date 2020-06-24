// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceReleasedListNot */
/* PDUSessionResourceReleasedItemNot */
type PDUSessionResourceReleasedListNot struct {
	List []PDUSessionResourceReleasedItemNot `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
