// Created By HaoDHH-245789 VHT2020
package ngapType

type TAIListForPagingItem struct {
	TAI          TAI                                                   `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerTAIListForPagingItemExtIEs `vht:"optional"`
}
