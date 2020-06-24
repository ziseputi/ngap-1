// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceSwitchedList */
/* PDUSessionResourceSwitchedItem */
type PDUSessionResourceSwitchedList struct {
	List []PDUSessionResourceSwitchedItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
