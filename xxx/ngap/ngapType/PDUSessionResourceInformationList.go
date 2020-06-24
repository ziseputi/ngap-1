// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceInformationList */
/* PDUSessionResourceInformationItem */
type PDUSessionResourceInformationList struct {
	List []PDUSessionResourceInformationItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
