// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowAddOrModifyRequestList */
/* QosFlowAddOrModifyRequestItem */
type QosFlowAddOrModifyRequestList struct {
	List []QosFlowAddOrModifyRequestItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
