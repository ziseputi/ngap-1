// Created By HaoDHH-245789 VHT2020
package ngapType

type TargetRANNodeID struct {
	GlobalRANNodeID GlobalRANNodeID                                  `vht:"valueLB:0,valueUB:3"`
	SelectedTAI     TAI                                              `vht:"valueExt"`
	IEExtensions    *ProtocolExtensionContainerTargetRANNodeIDExtIEs `vht:"optional"`
}
