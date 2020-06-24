// Created By HaoDHH-245789 VHT2020
package ngapType

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 *ExpectedActivityPeriod                                      `vht:"optional"`
	ExpectedIdlePeriod                     *ExpectedIdlePeriod                                          `vht:"optional"`
	SourceOfUEActivityBehaviourInformation *SourceOfUEActivityBehaviourInformation                      `vht:"optional"`
	IEExtensions                           *ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs `vht:"optional"`
}
