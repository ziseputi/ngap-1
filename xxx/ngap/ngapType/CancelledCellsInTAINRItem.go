// Created By HaoDHH-245789 VHT2020
package ngapType

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI `vht:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs `vht:"optional"`
}
