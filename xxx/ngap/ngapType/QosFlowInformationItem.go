// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowInformationItem struct {
	QosFlowIdentifier QosFlowIdentifier
	DLForwarding      *DLForwarding                                           `vht:"optional"`
	IEExtensions      *ProtocolExtensionContainerQosFlowInformationItemExtIEs `vht:"optional"`
}
