// Created By HaoDHH-245789 VHT2020
package ngapType

type RecommendedCellItem struct {
	NGRANCGI         NGRANCGI                                             `vht:"valueLB:0,valueUB:2"`
	TimeStayedInCell *int64                                               `vht:"valueLB:0,valueUB:4095,optional"`
	IEExtensions     *ProtocolExtensionContainerRecommendedCellItemExtIEs `vht:"optional"`
}
