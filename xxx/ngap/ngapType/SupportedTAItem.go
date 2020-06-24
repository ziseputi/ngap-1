// Created By HaoDHH-245789 VHT2020
package ngapType

type SupportedTAItem struct {
	TAC               TAC
	BroadcastPLMNList BroadcastPLMNList
	IEExtensions      *ProtocolExtensionContainerSupportedTAItemExtIEs `vht:"optional"`
}
