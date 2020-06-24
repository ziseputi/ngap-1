// Created By HaoDHH-245789 VHT2020
package ngapType

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        QosFlowIdentifier
	QosFlowMappingIndication *Enumerated                                       `vht:"optional"`
	IEExtensions             *ProtocolExtensionContainerAssociatedQosFlowItemExtIEs `vht:"optional"`
}
