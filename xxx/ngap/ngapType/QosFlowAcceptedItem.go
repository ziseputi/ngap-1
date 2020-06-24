// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowAcceptedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAcceptedItemExtIEs `vht:"optional"`
}
