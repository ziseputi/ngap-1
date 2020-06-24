// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToModifyListModRes */
/* PDUSessionResourceFailedToModifyItemModRes */
type PDUSessionResourceFailedToModifyListModRes struct {
	List []PDUSessionResourceFailedToModifyItemModRes `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
