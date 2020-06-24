// Created By HaoDHH-245789 VHT2020
package ngapType

type AreaOfInterestRANNodeItem struct {
	GlobalRANNodeID GlobalRANNodeID                                            `vht:"valueLB:0,valueUB:3"`
	IEExtensions    *ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs `vht:"optional"`
}
