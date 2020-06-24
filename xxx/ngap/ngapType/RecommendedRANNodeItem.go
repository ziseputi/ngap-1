// Created By HaoDHH-245789 VHT2020
package ngapType

type RecommendedRANNodeItem struct {
	AMFPagingTarget AMFPagingTarget                                         `vht:"valueLB:0,valueUB:2"`
	IEExtensions    *ProtocolExtensionContainerRecommendedRANNodeItemExtIEs `vht:"optional"`
}
