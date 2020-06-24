// Created By HaoDHH-245789 VHT2020
package ngapType

type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging RecommendedCellsForPaging                                          `vht:"valueExt"`
	IEExtensions              *ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs `vht:"optional"`
}
