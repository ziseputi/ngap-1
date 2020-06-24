// Created By HaoDHH-245789 VHT2020
package ngapType

type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation                           `vht:"valueLB:0,valueUB:1"`
	DLNGUUPTNLInformation UPTransportLayerInformation                           `vht:"valueLB:0,valueUB:1"`
	IEExtensions          *ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs `vht:"optional"`
}
