// Created By HaoDHH-245789 VHT2020
package ngapType

type TAI struct {
	PLMNIdentity PLMNIdentity
	TAC          TAC
	IEExtensions *ProtocolExtensionContainerTAIExtIEs `vht:"optional"`
}
