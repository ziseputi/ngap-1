// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceListCxtRelReq */
/* PDUSessionResourceItemCxtRelReq */
type PDUSessionResourceListCxtRelReq struct {
	List []PDUSessionResourceItemCxtRelReq `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
