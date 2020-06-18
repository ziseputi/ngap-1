package ngapType

const (
	NGAPPDUChoiceInitiatingMessage   int = 0
	NGAPPDUChoiceSuccessfulOutcome   int = 1
	NGAPPDUChoiceUnsuccessfulOutcome int = 2
	/* Extensions may appear below */
)

type NGAPPDU struct {
	Choice              int
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}

type InitiatingMessage struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality            `vht:"Reference:ProcedureCode"`
	Value         InitiatingMessageValue `vht:"Reference:ProcedureCode"`
}

type SuccessfulOutcome struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality            `vht:"Reference:ProcedureCode"`
	Value         SuccessfulOutcomeValue `vht:"Reference:ProcedureCode"`
}

type UnsuccessfulOutcome struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality              `vht:"Reference:ProcedureCode"`
	Value         UnsuccessfulOutcomeValue `vht:"Reference:ProcedureCode"`
}

const (
	InitiatingMessageValueChoiceAMFConfigurationUpdate                int = 0
	InitiatingMessageValueChoiceDeactivateTrace                       int = 1
	InitiatingMessageValueChoiceDownlinkNASTransport                  int = 2
	InitiatingMessageValueChoiceDownlinkNonUEAssociatedNRPPaTransport int = 3
	InitiatingMessageValueChoiceDownlinkRANConfigurationTransfer      int = 4
	InitiatingMessageValueChoiceDownlinkRANStatusTransfer             int = 5
	InitiatingMessageValueChoiceDownlinkUEAssociatedNRPPaTransport    int = 6
	InitiatingMessageValueChoiceErrorIndication                       int = 7
	InitiatingMessageValueChoiceHandoverCancel                        int = 8
	InitiatingMessageValueChoiceHandoverNotify                        int = 9
	InitiatingMessageValueChoiceHandoverRequired                      int = 10
	InitiatingMessageValueChoiceHandoverRequest                       int = 11
	InitiatingMessageValueChoiceInitialContextSetupRequest            int = 12
	InitiatingMessageValueChoiceInitialUEMessage                      int = 13
	InitiatingMessageValueChoiceLocationReport                        int = 14
	InitiatingMessageValueChoiceLocationReportingControl              int = 15
	InitiatingMessageValueChoiceLocationReportingFailureIndication    int = 16
	InitiatingMessageValueChoiceNASNonDeliveryIndication              int = 17
	InitiatingMessageValueChoiceNGReset                               int = 18
	InitiatingMessageValueChoiceNGSetupRequest                        int = 19
	InitiatingMessageValueChoiceOverloadStart                         int = 20
	InitiatingMessageValueChoiceOverloadStop                          int = 21
	InitiatingMessageValueChoicePaging                                int = 22
	InitiatingMessageValueChoicePathSwitchRequest                     int = 23
	InitiatingMessageValueChoicePDUSessionResourceModifyRequest       int = 24
	InitiatingMessageValueChoicePDUSessionResourceModifyIndication    int = 25
	InitiatingMessageValueChoicePDUSessionResourceNotify              int = 26
	InitiatingMessageValueChoicePDUSessionResourceReleaseCommand      int = 27
	InitiatingMessageValueChoicePDUSessionResourceSetupRequest        int = 28
	InitiatingMessageValueChoicePrivateMessage                        int = 29
	InitiatingMessageValueChoicePWSCancelRequest                      int = 30
	InitiatingMessageValueChoicePWSFailureIndication                  int = 31
	InitiatingMessageValueChoicePWSRestartIndication                  int = 32
	InitiatingMessageValueChoiceRANConfigurationUpdate                int = 33
	InitiatingMessageValueChoiceRerouteNASRequest                     int = 34
	InitiatingMessageValueChoiceRRCInactiveTransitionReport           int = 35
	InitiatingMessageValueChoiceSecondaryRATDataUsageReport           int = 36
	InitiatingMessageValueChoiceTraceFailureIndication                int = 37
	InitiatingMessageValueChoiceTraceStart                            int = 38
	InitiatingMessageValueChoiceUEContextModificationRequest          int = 39
	InitiatingMessageValueChoiceUEContextReleaseCommand               int = 40
	InitiatingMessageValueChoiceUEContextReleaseRequest               int = 41
	InitiatingMessageValueChoiceUERadioCapabilityCheckRequest         int = 42
	InitiatingMessageValueChoiceUERadioCapabilityInfoIndication       int = 43
	InitiatingMessageValueChoiceUETNLABindingReleaseRequest           int = 44
	InitiatingMessageValueChoiceUplinkNASTransport                    int = 45
	InitiatingMessageValueChoiceUplinkNonUEAssociatedNRPPaTransport   int = 46
	InitiatingMessageValueChoiceUplinkRANConfigurationTransfer        int = 47
	InitiatingMessageValueChoiceUplinkRANStatusTransfer               int = 48
	InitiatingMessageValueChoiceUplinkUEAssociatedNRPPaTransport      int = 49
	InitiatingMessageValueChoiceWriteReplaceWarningRequest            int = 50
	InitiatingMessageValueChoiceUplinkRIMInformationTransfer          int = 51
	InitiatingMessageValueChoiceDownlinkRIMInformationTransfer        int = 52
)

type InitiatingMessageValue struct {
	Choice                                int
	AMFConfigurationUpdate                *AMFConfigurationUpdate                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeAMFConfigurationUpdate"`
	DeactivateTrace                       *DeactivateTrace                       `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDeactivateTrace"`
	DownlinkNASTransport                  *DownlinkNASTransport                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDownlinkNASTransport"`
	DownlinkNonUEAssociatedNRPPaTransport *DownlinkNonUEAssociatedNRPPaTransport `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport"`
	DownlinkRANConfigurationTransfer      *DownlinkRANConfigurationTransfer      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDownlinkRANConfigurationTransfer"`
	DownlinkRANStatusTransfer             *DownlinkRANStatusTransfer             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDownlinkRANStatusTransfer"`
	DownlinkUEAssociatedNRPPaTransport    *DownlinkUEAssociatedNRPPaTransport    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDownlinkUEAssociatedNRPPaTransport"`
	ErrorIndication                       *ErrorIndication                       `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeErrorIndication"`
	HandoverCancel                        *HandoverCancel                        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverCancel"`
	HandoverNotify                        *HandoverNotify                        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeHandoverNotification"`
	HandoverRequired                      *HandoverRequired                      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverPreparation"`
	HandoverRequest                       *HandoverRequest                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverResourceAllocation"`
	InitialContextSetupRequest            *InitialContextSetupRequest            `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeInitialContextSetup"`
	InitialUEMessage                      *InitialUEMessage                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeInitialUEMessage"`
	LocationReport                        *LocationReport                        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeLocationReport"`
	LocationReportingControl              *LocationReportingControl              `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeLocationReportingControl"`
	LocationReportingFailureIndication    *LocationReportingFailureIndication    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeLocationReportingFailureIndication"`
	NASNonDeliveryIndication              *NASNonDeliveryIndication              `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeNASNonDeliveryIndication"`
	NGReset                               *NGReset                               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGReset"`
	NGSetupRequest                        *NGSetupRequest                        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGSetup"`
	OverloadStart                         *OverloadStart                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeOverloadStart"`
	OverloadStop                          *OverloadStop                          `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeOverloadStop"`
	Paging                                *Paging                                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodePaging"`
	PathSwitchRequest                     *PathSwitchRequest                     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePathSwitchRequest"`
	PDUSessionResourceModifyRequest       *PDUSessionResourceModifyRequest       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceModify"`
	PDUSessionResourceModifyIndication    *PDUSessionResourceModifyIndication    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceModifyIndication"`
	PDUSessionResourceNotify              *PDUSessionResourceNotify              `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodePDUSessionResourceNotify"`
	PDUSessionResourceReleaseCommand      *PDUSessionResourceReleaseCommand      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceRelease"`
	PDUSessionResourceSetupRequest        *PDUSessionResourceSetupRequest        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceSetup"`
	PrivateMessage                        *PrivateMessage                        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodePrivateMessage"`
	PWSCancelRequest                      *PWSCancelRequest                      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePWSCancel"`
	PWSFailureIndication                  *PWSFailureIndication                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodePWSFailureIndication"`
	PWSRestartIndication                  *PWSRestartIndication                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodePWSRestartIndication"`
	RANConfigurationUpdate                *RANConfigurationUpdate                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeRANConfigurationUpdate"`
	RerouteNASRequest                     *RerouteNASRequest                     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeRerouteNASRequest"`
	RRCInactiveTransitionReport           *RRCInactiveTransitionReport           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeRRCInactiveTransitionReport"`
	SecondaryRATDataUsageReport           *SecondaryRATDataUsageReport           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeSecondaryRATDataUsageReport"`
	TraceFailureIndication                *TraceFailureIndication                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeTraceFailureIndication"`
	TraceStart                            *TraceStart                            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeTraceStart"`
	UEContextModificationRequest          *UEContextModificationRequest          `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextModification"`
	UEContextReleaseCommand               *UEContextReleaseCommand               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextRelease"`
	UEContextReleaseRequest               *UEContextReleaseRequest               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUEContextReleaseRequest"`
	UERadioCapabilityCheckRequest         *UERadioCapabilityCheckRequest         `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUERadioCapabilityCheck"`
	UERadioCapabilityInfoIndication       *UERadioCapabilityInfoIndication       `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUERadioCapabilityInfoIndication"`
	UETNLABindingReleaseRequest           *UETNLABindingReleaseRequest           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUETNLABindingRelease"`
	UplinkNASTransport                    *UplinkNASTransport                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUplinkNASTransport"`
	UplinkNonUEAssociatedNRPPaTransport   *UplinkNonUEAssociatedNRPPaTransport   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUplinkNonUEAssociatedNRPPaTransport"`
	UplinkRANConfigurationTransfer        *UplinkRANConfigurationTransfer        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUplinkRANConfigurationTransfer"`
	UplinkRANStatusTransfer               *UplinkRANStatusTransfer               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUplinkRANStatusTransfer"`
	UplinkUEAssociatedNRPPaTransport      *UplinkUEAssociatedNRPPaTransport      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUplinkUEAssociatedNRPPaTransport"`
	WriteReplaceWarningRequest            *WriteReplaceWarningRequest            `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeWriteReplaceWarning"`
	UplinkRIMInformationTransfer          *UplinkRIMInformationTransfer          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeUplinkRIMInformationTransfer"`
	DownlinkRIMInformationTransfer        *DownlinkRIMInformationTransfer        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeDownlinkRIMInformationTransfer"`
}

const (
	SuccessfulOutcomeValueChoiceAMFConfigurationUpdateAcknowledge int = 0
	SuccessfulOutcomeValueChoiceHandoverCancelAcknowledge         int = 1
	SuccessfulOutcomeValueChoiceHandoverCommand                   int = 2
	SuccessfulOutcomeValueChoiceHandoverRequestAcknowledge        int = 3
	SuccessfulOutcomeValueChoiceInitialContextSetupResponse       int = 4
	SuccessfulOutcomeValueChoiceNGResetAcknowledge                int = 5
	SuccessfulOutcomeValueChoiceNGSetupResponse                   int = 6
	SuccessfulOutcomeValueChoicePathSwitchRequestAcknowledge      int = 7
	SuccessfulOutcomeValueChoicePDUSessionResourceModifyResponse  int = 8
	SuccessfulOutcomeValueChoicePDUSessionResourceModifyConfirm   int = 9
	SuccessfulOutcomeValueChoicePDUSessionResourceReleaseResponse int = 10
	SuccessfulOutcomeValueChoicePDUSessionResourceSetupResponse   int = 11
	SuccessfulOutcomeValueChoicePWSCancelResponse                 int = 12
	SuccessfulOutcomeValueChoiceRANConfigurationUpdateAcknowledge int = 13
	SuccessfulOutcomeValueChoiceUEContextModificationResponse     int = 14
	SuccessfulOutcomeValueChoiceUEContextReleaseComplete          int = 15
	SuccessfulOutcomeValueChoiceUERadioCapabilityCheckResponse    int = 16
	SuccessfulOutcomeValueChoiceWriteReplaceWarningResponse       int = 17
)

type SuccessfulOutcomeValue struct {
	Choice                            int
	AMFConfigurationUpdateAcknowledge *AMFConfigurationUpdateAcknowledge `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeAMFConfigurationUpdate"`
	HandoverCancelAcknowledge         *HandoverCancelAcknowledge         `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverCancel"`
	HandoverCommand                   *HandoverCommand                   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverPreparation"`
	HandoverRequestAcknowledge        *HandoverRequestAcknowledge        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverResourceAllocation"`
	InitialContextSetupResponse       *InitialContextSetupResponse       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeInitialContextSetup"`
	NGResetAcknowledge                *NGResetAcknowledge                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGReset"`
	NGSetupResponse                   *NGSetupResponse                   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGSetup"`
	PathSwitchRequestAcknowledge      *PathSwitchRequestAcknowledge      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePathSwitchRequest"`
	PDUSessionResourceModifyResponse  *PDUSessionResourceModifyResponse  `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceModify"`
	PDUSessionResourceModifyConfirm   *PDUSessionResourceModifyConfirm   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceModifyIndication"`
	PDUSessionResourceReleaseResponse *PDUSessionResourceReleaseResponse `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceRelease"`
	PDUSessionResourceSetupResponse   *PDUSessionResourceSetupResponse   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePDUSessionResourceSetup"`
	PWSCancelResponse                 *PWSCancelResponse                 `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePWSCancel"`
	RANConfigurationUpdateAcknowledge *RANConfigurationUpdateAcknowledge `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeRANConfigurationUpdate"`
	UEContextModificationResponse     *UEContextModificationResponse     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextModification"`
	UEContextReleaseComplete          *UEContextReleaseComplete          `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextRelease"`
	UERadioCapabilityCheckResponse    *UERadioCapabilityCheckResponse    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUERadioCapabilityCheck"`
	WriteReplaceWarningResponse       *WriteReplaceWarningResponse       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeWriteReplaceWarning"`
}

const (
	UnsuccessfulOutcomeValueChoiceAMFConfigurationUpdateFailure int = 0
	UnsuccessfulOutcomeValueChoiceHandoverPreparationFailure    int = 1
	UnsuccessfulOutcomeValueChoiceHandoverFailure               int = 2
	UnsuccessfulOutcomeValueChoiceInitialContextSetupFailure    int = 3
	UnsuccessfulOutcomeValueChoiceNGSetupFailure                int = 4
	UnsuccessfulOutcomeValueChoicePathSwitchRequestFailure      int = 5
	UnsuccessfulOutcomeValueChoiceRANConfigurationUpdateFailure int = 6
	UnsuccessfulOutcomeValueChoiceUEContextModificationFailure  int = 7
)

type UnsuccessfulOutcomeValue struct {
	Choice                        int
	AMFConfigurationUpdateFailure *AMFConfigurationUpdateFailure `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeAMFConfigurationUpdate"`
	HandoverPreparationFailure    *HandoverPreparationFailure    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverPreparation"`
	HandoverFailure               *HandoverFailure               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverResourceAllocation"`
	InitialContextSetupFailure    *InitialContextSetupFailure    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeInitialContextSetup"`
	NGSetupFailure                *NGSetupFailure                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGSetup"`
	PathSwitchRequestFailure      *PathSwitchRequestFailure      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePathSwitchRequest"`
	RANConfigurationUpdateFailure *RANConfigurationUpdateFailure `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeRANConfigurationUpdate"`
	UEContextModificationFailure  *UEContextModificationFailure  `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextModification"`
}
