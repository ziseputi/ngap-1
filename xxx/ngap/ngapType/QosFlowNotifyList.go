// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowNotifyList */
/* QosFlowNotifyItem */
type QosFlowNotifyList struct {
	List []QosFlowNotifyItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
