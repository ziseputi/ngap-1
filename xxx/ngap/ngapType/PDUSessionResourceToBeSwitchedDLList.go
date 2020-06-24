// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceToBeSwitchedDLList */
/* PDUSessionResourceToBeSwitchedDLItem */
type PDUSessionResourceToBeSwitchedDLList struct {
	List []PDUSessionResourceToBeSwitchedDLItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
