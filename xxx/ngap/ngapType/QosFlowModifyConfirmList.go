// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowModifyConfirmList */
/* QosFlowModifyConfirmItem */
type QosFlowModifyConfirmList struct {
	List []QosFlowModifyConfirmItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
