// Created By HaoDHH-245789 VHT2020
package ngapType

type NonDynamic5QIDescriptor struct {
	FiveQI                 FiveQI
	PriorityLevelQos       *PriorityLevelQos                                        `vht:"optional"`
	AveragingWindow        *AveragingWindow                                         `vht:"optional"`
	MaximumDataBurstVolume *MaximumDataBurstVolume                                  `vht:"optional"`
	IEExtensions           *ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs `vht:"optional"`
}
