// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToSetupListCxtRes */
/* PDUSessionResourceFailedToSetupItemCxtRes */
type PDUSessionResourceFailedToSetupListCxtRes struct {
	List []PDUSessionResourceFailedToSetupItemCxtRes `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
