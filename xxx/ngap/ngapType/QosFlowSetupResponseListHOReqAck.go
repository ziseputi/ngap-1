// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowSetupResponseListHOReqAck */
/* QosFlowSetupResponseItemHOReqAck */
type QosFlowSetupResponseListHOReqAck struct {
	List []QosFlowSetupResponseItemHOReqAck `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
