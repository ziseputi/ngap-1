// Created By HaoDHH-245789 VHT2020
package ngapType

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  RecommendedCellsForPaging                                                   `vht:"valueExt"`
	RecommendRANNodesForPaging RecommendedRANNodesForPaging                                                `vht:"valueExt"`
	IEExtensions               *ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs `vht:"optional"`
}
