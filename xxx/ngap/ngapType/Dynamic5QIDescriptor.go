// Created By HaoDHH-245789 VHT2020
package ngapType

type Dynamic5QIDescriptor struct {
	PriorityLevelQos       PriorityLevelQos
	PacketDelayBudget      PacketDelayBudget
	PacketErrorRate        PacketErrorRate                                       `vht:"valueExt"`
	FiveQI                 *FiveQI                                               `vht:"optional"`
	DelayCritical          *DelayCritical                                        `vht:"optional"`
	AveragingWindow        *AveragingWindow                                      `vht:"optional"`
	MaximumDataBurstVolume *MaximumDataBurstVolume                               `vht:"optional"`
	IEExtensions           *ProtocolExtensionContainerDynamic5QIDescriptorExtIEs `vht:"optional"`
}
