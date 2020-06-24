// Created By HaoDHH-245789 VHT2020
package ngapType

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList RecommendedRANNodeList
	IEExtensions           *ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs `vht:"optional"`
}
