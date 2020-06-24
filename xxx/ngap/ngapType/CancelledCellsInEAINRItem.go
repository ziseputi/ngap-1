// Created By HaoDHH-245789 VHT2020
package ngapType

type CancelledCellsInEAINRItem struct {
	NRCGI              NRCGI `vht:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs `vht:"optional"`
}
