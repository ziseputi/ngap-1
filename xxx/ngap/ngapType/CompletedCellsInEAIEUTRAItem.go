// Created By HaoDHH-245789 VHT2020
package ngapType

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs `vht:"optional"`
}
