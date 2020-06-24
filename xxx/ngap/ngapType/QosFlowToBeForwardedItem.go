// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs `vht:"optional"`
}
