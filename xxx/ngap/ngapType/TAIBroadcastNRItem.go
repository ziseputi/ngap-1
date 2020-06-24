// Created By HaoDHH-245789 VHT2020
package ngapType

type TAIBroadcastNRItem struct {
	TAI                   TAI `vht:"valueExt"`
	CompletedCellsInTAINR CompletedCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAIBroadcastNRItemExtIEs `vht:"optional"`
}
