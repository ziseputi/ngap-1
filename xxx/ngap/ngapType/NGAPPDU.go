// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	NGAPPDUPresentNothing int = iota /* No components present */
	NGAPPDUPresentInitiatingMessage
	NGAPPDUPresentSuccessfulOutcome
	NGAPPDUPresentUnsuccessfulOutcome
	/* Extensions may appear below */
)

type NGAPPDU struct {
	Present             int
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}
