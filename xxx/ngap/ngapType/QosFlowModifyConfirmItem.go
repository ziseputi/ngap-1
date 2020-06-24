// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs `vht:"optional"`
}
