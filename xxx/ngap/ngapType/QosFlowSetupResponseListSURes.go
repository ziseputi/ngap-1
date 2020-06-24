// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct QosFlowSetupResponseListSURes */
/* QosFlowSetupResponseItemSURes */
type QosFlowSetupResponseListSURes struct {
	List []QosFlowSetupResponseItemSURes `vht:"valueExt,sizeLB:1,sizeUB:64"`
}
