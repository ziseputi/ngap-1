// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceAdmittedList */
/* PDUSessionResourceAdmittedItem */
type PDUSessionResourceAdmittedList struct {
	List []PDUSessionResourceAdmittedItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
