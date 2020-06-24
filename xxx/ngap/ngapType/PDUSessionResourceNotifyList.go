// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceNotifyList */
/* PDUSessionResourceNotifyItem */
type PDUSessionResourceNotifyList struct {
	List []PDUSessionResourceNotifyItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
