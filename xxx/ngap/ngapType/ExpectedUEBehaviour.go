// Created By HaoDHH-245789 VHT2020
package ngapType

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour *ExpectedUEActivityBehaviour                         `vht:"valueExt,optional"`
	ExpectedHOInterval          *ExpectedHOInterval                                  `vht:"optional"`
	ExpectedUEMobility          *ExpectedUEMobility                                  `vht:"optional"`
	ExpectedUEMovingTrajectory  *ExpectedUEMovingTrajectory                          `vht:"optional"`
	IEExtensions                *ProtocolExtensionContainerExpectedUEBehaviourExtIEs `vht:"optional"`
}
