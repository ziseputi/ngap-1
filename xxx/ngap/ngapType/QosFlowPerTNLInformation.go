// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation `vht:"valueLB:0,valueUB:1"`
	AssociatedQosFlowList       AssociatedQosFlowList
	IEExtensions                *ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs `vht:"optional"`
}
