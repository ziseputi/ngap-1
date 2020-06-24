// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters *QosFlowLevelQosParameters                                     `vht:"valueExt,optional"`
	ERABID                    *ERABID                                                        `vht:"optional"`
	IEExtensions              *ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs `vht:"optional"`
}
