// Created By HaoDHH-245789 VHT2020
package ngapType

type SuccessfulOutcome struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality
	Value         SuccessfulOutcomeValue `vht:"openType,referenceFieldName:ProcedureCode"`
}

const (
	SuccessfulOutcomePresentNothing int = iota /* No components present */
	SuccessfulOutcomePresentAMFConfigurationUpdateAcknowledge
	SuccessfulOutcomePresentHandoverCancelAcknowledge
	SuccessfulOutcomePresentHandoverCommand
	SuccessfulOutcomePresentHandoverRequestAcknowledge
	SuccessfulOutcomePresentInitialContextSetupResponse
	SuccessfulOutcomePresentNGResetAcknowledge
	SuccessfulOutcomePresentNGSetupResponse
	SuccessfulOutcomePresentPathSwitchRequestAcknowledge
	SuccessfulOutcomePresentPDUSessionResourceModifyResponse
	SuccessfulOutcomePresentPDUSessionResourceModifyConfirm
	SuccessfulOutcomePresentPDUSessionResourceReleaseResponse
	SuccessfulOutcomePresentPDUSessionResourceSetupResponse
	SuccessfulOutcomePresentPWSCancelResponse
	SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge
	SuccessfulOutcomePresentUEContextModificationResponse
	SuccessfulOutcomePresentUEContextReleaseComplete
	SuccessfulOutcomePresentUERadioCapabilityCheckResponse
	SuccessfulOutcomePresentWriteReplaceWarningResponse
)

type SuccessfulOutcomeValue struct {
	Present                           int
	AMFConfigurationUpdateAcknowledge *AMFConfigurationUpdateAcknowledge `vht:"valueExt,referenceFieldValue:0"`
	HandoverCancelAcknowledge         *HandoverCancelAcknowledge         `vht:"valueExt,referenceFieldValue:10"`
	HandoverCommand                   *HandoverCommand                   `vht:"valueExt,referenceFieldValue:12"`
	HandoverRequestAcknowledge        *HandoverRequestAcknowledge        `vht:"valueExt,referenceFieldValue:13"`
	InitialContextSetupResponse       *InitialContextSetupResponse       `vht:"valueExt,referenceFieldValue:14"`
	NGResetAcknowledge                *NGResetAcknowledge                `vht:"valueExt,referenceFieldValue:20"`
	NGSetupResponse                   *NGSetupResponse                   `vht:"valueExt,referenceFieldValue:21"`
	PathSwitchRequestAcknowledge      *PathSwitchRequestAcknowledge      `vht:"valueExt,referenceFieldValue:25"`
	PDUSessionResourceModifyResponse  *PDUSessionResourceModifyResponse  `vht:"valueExt,referenceFieldValue:26"`
	PDUSessionResourceModifyConfirm   *PDUSessionResourceModifyConfirm   `vht:"valueExt,referenceFieldValue:27"`
	PDUSessionResourceReleaseResponse *PDUSessionResourceReleaseResponse `vht:"valueExt,referenceFieldValue:28"`
	PDUSessionResourceSetupResponse   *PDUSessionResourceSetupResponse   `vht:"valueExt,referenceFieldValue:29"`
	PWSCancelResponse                 *PWSCancelResponse                 `vht:"valueExt,referenceFieldValue:32"`
	RANConfigurationUpdateAcknowledge *RANConfigurationUpdateAcknowledge `vht:"valueExt,referenceFieldValue:35"`
	UEContextModificationResponse     *UEContextModificationResponse     `vht:"valueExt,referenceFieldValue:40"`
	UEContextReleaseComplete          *UEContextReleaseComplete          `vht:"valueExt,referenceFieldValue:41"`
	UERadioCapabilityCheckResponse    *UERadioCapabilityCheckResponse    `vht:"valueExt,referenceFieldValue:43"`
	WriteReplaceWarningResponse       *WriteReplaceWarningResponse       `vht:"valueExt,referenceFieldValue:51"`
}
