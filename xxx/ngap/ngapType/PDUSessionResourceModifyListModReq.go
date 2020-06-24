// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceModifyListModReq */
/* PDUSessionResourceModifyItemModReq */
type PDUSessionResourceModifyListModReq struct {
	List []PDUSessionResourceModifyItemModReq `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
