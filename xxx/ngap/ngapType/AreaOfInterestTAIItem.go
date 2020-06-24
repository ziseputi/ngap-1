// Created By HaoDHH-245789 VHT2020
package ngapType

type AreaOfInterestTAIItem struct {
	TAI          TAI                                                    `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs `vht:"optional"`
}
