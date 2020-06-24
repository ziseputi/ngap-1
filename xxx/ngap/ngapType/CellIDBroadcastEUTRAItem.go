// Created By HaoDHH-245789 VHT2020
package ngapType

type CellIDBroadcastEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                  `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs `vht:"optional"`
}
