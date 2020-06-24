// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToSetupListHOAck */
/* PDUSessionResourceFailedToSetupItemHOAck */
type PDUSessionResourceFailedToSetupListHOAck struct {
	List []PDUSessionResourceFailedToSetupItemHOAck `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
