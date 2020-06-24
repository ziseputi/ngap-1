// Created By HaoDHH-245789 VHT2020
package ngapType

type CompletedCellsInTAINRItem struct {
	NRCGI        NRCGI                                                      `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs `vht:"optional"`
}
