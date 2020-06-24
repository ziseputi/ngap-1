// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceListHORqd */
/* PDUSessionResourceItemHORqd */
type PDUSessionResourceListHORqd struct {
	List []PDUSessionResourceItemHORqd `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
