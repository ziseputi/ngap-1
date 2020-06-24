// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceSetupListCxtRes */
/* PDUSessionResourceSetupItemCxtRes */
type PDUSessionResourceSetupListCxtRes struct {
	List []PDUSessionResourceSetupItemCxtRes `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
