// Created By HaoDHH-245789 VHT2020
package ngapType

type CellIDCancelledEUTRAItem struct {
	EUTRACGI           EUTRACGI `vht:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs `vht:"optional"`
}
