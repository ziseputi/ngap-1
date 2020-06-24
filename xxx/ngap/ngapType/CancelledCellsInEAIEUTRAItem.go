// Created By HaoDHH-245789 VHT2020
package ngapType

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI `vht:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs `vht:"optional"`
}
