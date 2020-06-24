// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters QosFlowLevelQosParameters                                `vht:"valueExt"`
	ERABID                    *ERABID                                                  `vht:"optional"`
	IEExtensions              *ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs `vht:"optional"`
}
