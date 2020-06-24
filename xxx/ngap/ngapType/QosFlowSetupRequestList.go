// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowSetupRequestList */
/* QosFlowSetupRequestItem */
type QosFlowSetupRequestList struct {
	List []QosFlowSetupRequestItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
