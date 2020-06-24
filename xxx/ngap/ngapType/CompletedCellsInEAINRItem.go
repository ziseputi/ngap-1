// Created By HaoDHH-245789 VHT2020
package ngapType

type CompletedCellsInEAINRItem struct {
	NRCGI        NRCGI                                                      `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs `vht:"optional"`
}
