// Created By HaoDHH-245789 VHT2020
package ngapType

type UnsuccessfulOutcome struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality
	Value         UnsuccessfulOutcomeValue `vht:"openType,referenceFieldName:ProcedureCode"`
}

const (
	UnsuccessfulOutcomePresentNothing int = iota /* No components present */
	UnsuccessfulOutcomePresentAMFConfigurationUpdateFailure
	UnsuccessfulOutcomePresentHandoverPreparationFailure
	UnsuccessfulOutcomePresentHandoverFailure
	UnsuccessfulOutcomePresentInitialContextSetupFailure
	UnsuccessfulOutcomePresentNGSetupFailure
	UnsuccessfulOutcomePresentPathSwitchRequestFailure
	UnsuccessfulOutcomePresentRANConfigurationUpdateFailure
	UnsuccessfulOutcomePresentUEContextModificationFailure
)

type UnsuccessfulOutcomeValue struct {
	Present                       int
	AMFConfigurationUpdateFailure *AMFConfigurationUpdateFailure `vht:"valueExt,referenceFieldValue:0"`
	HandoverPreparationFailure    *HandoverPreparationFailure    `vht:"valueExt,referenceFieldValue:12"`
	HandoverFailure               *HandoverFailure               `vht:"valueExt,referenceFieldValue:13"`
	InitialContextSetupFailure    *InitialContextSetupFailure    `vht:"valueExt,referenceFieldValue:14"`
	NGSetupFailure                *NGSetupFailure                `vht:"valueExt,referenceFieldValue:21"`
	PathSwitchRequestFailure      *PathSwitchRequestFailure      `vht:"valueExt,referenceFieldValue:25"`
	RANConfigurationUpdateFailure *RANConfigurationUpdateFailure `vht:"valueExt,referenceFieldValue:35"`
	UEContextModificationFailure  *UEContextModificationFailure  `vht:"valueExt,referenceFieldValue:40"`
}
