// Created By HaoDHH-245789 VHT2020
package ngapType

type CellIDBroadcastNRItem struct {
	NRCGI        NRCGI                                                  `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs `vht:"optional"`
}
