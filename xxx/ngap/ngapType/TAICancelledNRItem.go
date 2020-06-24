// Created By HaoDHH-245789 VHT2020
package ngapType

type TAICancelledNRItem struct {
	TAI                   TAI `vht:"valueExt"`
	CancelledCellsInTAINR CancelledCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAICancelledNRItemExtIEs `vht:"optional"`
}
