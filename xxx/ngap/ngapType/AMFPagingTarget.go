// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	AMFPagingTargetPresentNothing int = iota /* No components present */
	AMFPagingTargetPresentGlobalRANNodeID
	AMFPagingTargetPresentTAI
	AMFPagingTargetPresentChoiceExtensions
)

type AMFPagingTarget struct {
	Present          int
	GlobalRANNodeID  *GlobalRANNodeID `vht:"valueLB:0,valueUB:3"`
	TAI              *TAI             `vht:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerAMFPagingTargetExtIEs
}
