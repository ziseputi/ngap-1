// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowAddOrModifyResponseList */
/* QosFlowAddOrModifyResponseItem */
type QosFlowAddOrModifyResponseList struct {
	List []QosFlowAddOrModifyResponseItem `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
