// Created By HaoDHH-245789 VHT2020
package ngapType

type SingleTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation                           `vht:"valueLB:0,valueUB:1"`
	IEExtensions                *ProtocolExtensionContainerSingleTNLInformationExtIEs `vht:"optional"`
}
