// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToSetupListSURes */
/* PDUSessionResourceFailedToSetupItemSURes */
type PDUSessionResourceFailedToSetupListSURes struct {
	List []PDUSessionResourceFailedToSetupItemSURes `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
