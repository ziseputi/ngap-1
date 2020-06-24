// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceModifyListModRes */
/* PDUSessionResourceModifyItemModRes */
type PDUSessionResourceModifyListModRes struct {
	List []PDUSessionResourceModifyItemModRes `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
