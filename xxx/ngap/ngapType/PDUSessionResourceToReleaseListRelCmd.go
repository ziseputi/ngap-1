// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceToReleaseListRelCmd */
/* PDUSessionResourceToReleaseItemRelCmd */
type PDUSessionResourceToReleaseListRelCmd struct {
	List []PDUSessionResourceToReleaseItemRelCmd `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
