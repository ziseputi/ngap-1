// Created By HaoDHH-245789 VHT2020
package ngapType

type CellIDCancelledNRItem struct {
	NRCGI              NRCGI `vht:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledNRItemExtIEs `vht:"optional"`
}
