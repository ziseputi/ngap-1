// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowList */
/* QosFlowItem */
type QosFlowList struct {
	List []QosFlowItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
