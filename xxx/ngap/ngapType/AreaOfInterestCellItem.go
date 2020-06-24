// Created By HaoDHH-245789 VHT2020
package ngapType

type AreaOfInterestCellItem struct {
	NGRANCGI     NGRANCGI                                                `vht:"valueLB:0,valueUB:2"`
	IEExtensions *ProtocolExtensionContainerAreaOfInterestCellItemExtIEs `vht:"optional"`
}
