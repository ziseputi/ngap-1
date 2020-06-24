// Created By HaoDHH-245789 VHT2020
package ngapType

type SNSSAI struct {
	SST          SST
	SD           *SD                                     `vht:"optional"`
	IEExtensions *ProtocolExtensionContainerSNSSAIExtIEs `vht:"optional"`
}
