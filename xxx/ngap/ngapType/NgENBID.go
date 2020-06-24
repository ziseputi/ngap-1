// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	NgENBIDPresentNothing int = iota /* No components present */
	NgENBIDPresentMacroNgENBID
	NgENBIDPresentShortMacroNgENBID
	NgENBIDPresentLongMacroNgENBID
	NgENBIDPresentChoiceExtensions
)

type NgENBID struct {
	Present           int
	MacroNgENBID      *BitString `vht:"sizeLB:20,sizeUB:20"`
	ShortMacroNgENBID *BitString `vht:"sizeLB:18,sizeUB:18"`
	LongMacroNgENBID  *BitString `vht:"sizeLB:21,sizeUB:21"`
	ChoiceExtensions  *ProtocolIESingleContainerNgENBIDExtIEs
}
