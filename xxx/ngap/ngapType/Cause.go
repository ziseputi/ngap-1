// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CausePresentNothing int = iota /* No components present */
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentNas
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtensions
)

type Cause struct {
	Present          int
	RadioNetwork     *CauseRadioNetwork
	Transport        *CauseTransport
	Nas              *CauseNas
	Protocol         *CauseProtocol
	Misc             *CauseMisc
	ChoiceExtensions *ProtocolIESingleContainerCauseExtIEs
}
