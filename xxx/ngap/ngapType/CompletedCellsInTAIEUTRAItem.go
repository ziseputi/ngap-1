// Created By HaoDHH-245789 VHT2020
package ngapType

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs `vht:"optional"`
}
