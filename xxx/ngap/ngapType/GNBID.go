// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	GNBIDPresentNothing int = iota /* No components present */
	GNBIDPresentGNBID
	GNBIDPresentChoiceExtensions
)

type GNBID struct {
	Present          int
	GNBID            *BitString `vht:"sizeLB:22,sizeUB:32"`
	ChoiceExtensions *ProtocolIESingleContainerGNBIDExtIEs
}
