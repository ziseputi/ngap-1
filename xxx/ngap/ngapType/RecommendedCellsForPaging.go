// Created By HaoDHH-245789 VHT2020
package ngapType

type RecommendedCellsForPaging struct {
	RecommendedCellList RecommendedCellList
	IEExtensions        *ProtocolExtensionContainerRecommendedCellsForPagingExtIEs `vht:"optional"`
}
