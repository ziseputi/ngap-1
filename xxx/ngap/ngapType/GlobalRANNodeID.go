// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	GlobalRANNodeIDPresentNothing int = iota /* No components present */
	GlobalRANNodeIDPresentGlobalGNBID
	GlobalRANNodeIDPresentGlobalNgENBID
	GlobalRANNodeIDPresentGlobalN3IWFID
	GlobalRANNodeIDPresentChoiceExtensions
)

type GlobalRANNodeID struct {
	Present          int
	GlobalGNBID      *GlobalGNBID   `vht:"valueExt"`
	GlobalNgENBID    *GlobalNgENBID `vht:"valueExt"`
	GlobalN3IWFID    *GlobalN3IWFID `vht:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerGlobalRANNodeIDExtIEs
}
