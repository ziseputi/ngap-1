// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DRBStatusULPresentNothing int = iota /* No components present */
	DRBStatusULPresentDRBStatusUL12
	DRBStatusULPresentDRBStatusUL18
	DRBStatusULPresentChoiceExtensions
)

type DRBStatusUL struct {
	Present          int
	DRBStatusUL12    *DRBStatusUL12 `vht:"valueExt"`
	DRBStatusUL18    *DRBStatusUL18 `vht:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerDRBStatusULExtIEs
}
