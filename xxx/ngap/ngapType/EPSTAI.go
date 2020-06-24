// Created By HaoDHH-245789 VHT2020
package ngapType

type EPSTAI struct {
	PLMNIdentity PLMNIdentity
	EPSTAC       EPSTAC
	IEExtensions *ProtocolExtensionContainerEPSTAIExtIEs `vht:"optional"`
}
