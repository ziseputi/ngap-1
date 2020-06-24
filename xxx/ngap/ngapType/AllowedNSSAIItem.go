// Created By HaoDHH-245789 VHT2020
package ngapType

type AllowedNSSAIItem struct {
	SNSSAI       SNSSAI                                            `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAllowedNSSAIItemExtIEs `vht:"optional"`
}
