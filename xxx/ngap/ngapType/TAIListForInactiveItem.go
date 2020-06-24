// Created By HaoDHH-245789 VHT2020
package ngapType

type TAIListForInactiveItem struct {
	TAI          TAI                                                     `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerTAIListForInactiveItemExtIEs `vht:"optional"`
}
