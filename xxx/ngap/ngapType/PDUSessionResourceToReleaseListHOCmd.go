// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceToReleaseListHOCmd */
/* PDUSessionResourceToReleaseItemHOCmd */
type PDUSessionResourceToReleaseListHOCmd struct {
	List []PDUSessionResourceToReleaseItemHOCmd `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
