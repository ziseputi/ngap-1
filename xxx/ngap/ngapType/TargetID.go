// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	TargetIDPresentNothing int = iota /* No components present */
	TargetIDPresentTargetRANNodeID
	TargetIDPresentTargeteNBID
	TargetIDPresentChoiceExtensions
)

type TargetID struct {
	Present          int
	TargetRANNodeID  *TargetRANNodeID `vht:"valueExt"`
	TargeteNBID      *TargeteNBID     `vht:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerTargetIDExtIEs
}
