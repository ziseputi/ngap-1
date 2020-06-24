// Created By HaoDHH-245789 VHT2020
package ngapType

type TNLMappingItem struct {
	DLNGUUPTNLInformation UPTransportLayerInformation                     `vht:"valueLB:0,valueUB:1"`
	ULNGUUPTNLInformation UPTransportLayerInformation                     `vht:"valueLB:0,valueUB:1"`
	IEExtensions          *ProtocolExtensionContainerTNLMappingItemExtIEs `vht:"optional"`
}
