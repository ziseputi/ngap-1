// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DRBStatusDLPresentNothing int = iota /* No components present */
	DRBStatusDLPresentDRBStatusDL12
	DRBStatusDLPresentDRBStatusDL18
	DRBStatusDLPresentChoiceExtensions
)

type DRBStatusDL struct {
	Present          int
	DRBStatusDL12    *DRBStatusDL12 `vht:"valueExt"`
	DRBStatusDL18    *DRBStatusDL18 `vht:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerDRBStatusDLExtIEs
}
