// Created By HaoDHH-245789 VHT2020
package ngapType

type CancelledCellsInTAIEUTRAItem struct {
	EUTRACGI           EUTRACGI `vht:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs `vht:"optional"`
}
