// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceHandoverList */
/* PDUSessionResourceHandoverItem */
type PDUSessionResourceHandoverList struct {
	List []PDUSessionResourceHandoverItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
