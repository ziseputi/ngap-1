// Created By HaoDHH-245789 VHT2020
package ngapType

type QosFlowLevelQosParameters struct {
	QosCharacteristics             QosCharacteristics                                         `vht:"valueLB:0,valueUB:2"`
	AllocationAndRetentionPriority AllocationAndRetentionPriority                             `vht:"valueExt"`
	GBRQosInformation              *GBRQosInformation                                         `vht:"valueExt,optional"`
	ReflectiveQosAttribute         *ReflectiveQosAttribute                                    `vht:"optional"`
	AdditionalQosFlowInformation   *AdditionalQosFlowInformation                              `vht:"optional"`
	IEExtensions                   *ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs `vht:"optional"`
}
