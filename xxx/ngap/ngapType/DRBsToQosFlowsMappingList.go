// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct DRBsToQosFlowsMappingList */
/* DRBsToQosFlowsMappingItem */
type DRBsToQosFlowsMappingList struct {
	List []DRBsToQosFlowsMappingItem `vht:"valueExt,sizeLB:1,sizeUB:32"`
}
