// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceReleasedListRelRes */
/* PDUSessionResourceReleasedItemRelRes */
type PDUSessionResourceReleasedListRelRes struct {
	List []PDUSessionResourceReleasedItemRelRes `vht:"valueExt,sizeLB:1,sizeUB:256"`
}