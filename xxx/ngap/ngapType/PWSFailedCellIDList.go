// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PWSFailedCellIDListPresentNothing int = iota /* No components present */
	PWSFailedCellIDListPresentEUTRACGIPWSFailedList
	PWSFailedCellIDListPresentNRCGIPWSFailedList
	PWSFailedCellIDListPresentChoiceExtensions
)

type PWSFailedCellIDList struct {
	Present               int
	EUTRACGIPWSFailedList *EUTRACGIList
	NRCGIPWSFailedList    *NRCGIList
	ChoiceExtensions      *ProtocolIESingleContainerPWSFailedCellIDListExtIEs
}
