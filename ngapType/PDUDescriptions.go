package ngapType

const (
	NGAPPDUPresentInitiatingMessage   int = 0
	NGAPPDUPresentSuccessfulOutcome   int = 1
	NGAPPDUPresentUnsuccessfulOutcome int = 2
	/* Extensions may appear below */
)

type NGAPPDU struct {
	Present             int
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
	InitiatingMessagePresentAMFConfigurationUpdate                int = 0
	InitiatingMessagePresentAMFStatusIndication                   int = 1
	InitiatingMessagePresentCellTrafficTrace                      int = 2
	InitiatingMessagePresentDeactivateTrace                       int = 3
	InitiatingMessagePresentDownlinkNASTransport                  int = 4
	InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport int = 5
	InitiatingMessagePresentDownlinkRANConfigurationTransfer      int = 6
	InitiatingMessagePresentDownlinkRANStatusTransfer             int = 7
	InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport    int = 8
	InitiatingMessagePresentErrorIndication                       int = 9
	InitiatingMessagePresentHandoverCancel                        int = 10
	InitiatingMessagePresentHandoverNotify                        int = 11
	InitiatingMessagePresentHandoverRequired                      int = 12
	InitiatingMessagePresentHandoverRequest                       int = 13
	InitiatingMessagePresentInitialContextSetupRequest            int = 14
	InitiatingMessagePresentInitialUEMessage                      int = 15
	InitiatingMessagePresentLocationReport                        int = 16
	InitiatingMessagePresentLocationReportingControl              int = 17
	InitiatingMessagePresentLocationReportingFailureIndication    int = 18
	InitiatingMessagePresentNASNonDeliveryIndication              int = 19
	InitiatingMessagePresentNGReset                               int = 20
	InitiatingMessagePresentNGSetupRequest                        int = 21
	InitiatingMessagePresentOverloadStart                         int = 22
	InitiatingMessagePresentOverloadStop                          int = 23
	InitiatingMessagePresentPaging                                int = 24
	InitiatingMessagePresentPathSwitchRequest                     int = 25
	InitiatingMessagePresentPDUSessionResourceModifyRequest       int = 26
	InitiatingMessagePresentPDUSessionResourceModifyIndication    int = 27
	InitiatingMessagePresentPDUSessionResourceNotify              int = 28
	InitiatingMessagePresentPDUSessionResourceReleaseCommand      int = 29
	InitiatingMessagePresentPDUSessionResourceSetupRequest        int = 30
	InitiatingMessagePresentPrivateMessage                        int = 31
	InitiatingMessagePresentPWSCancelRequest                      int = 32
	InitiatingMessagePresentPWSFailureIndication                  int = 33
	InitiatingMessagePresentPWSRestartIndication                  int = 34
	InitiatingMessagePresentRANConfigurationUpdate                int = 35
	InitiatingMessagePresentRerouteNASRequest                     int = 36
	InitiatingMessagePresentRRCInactiveTransitionReport           int = 37
	InitiatingMessagePresentSecondaryRATDataUsageReport           int = 38
	InitiatingMessagePresentTraceFailureIndication                int = 39
	InitiatingMessagePresentTraceStart                            int = 40
	InitiatingMessagePresentUEContextModificationRequest          int = 41
	InitiatingMessagePresentUEContextReleaseCommand               int = 42
	InitiatingMessagePresentUEContextReleaseRequest               int = 43
	InitiatingMessagePresentUERadioCapabilityCheckRequest         int = 44
	InitiatingMessagePresentUERadioCapabilityInfoIndication       int = 45
	InitiatingMessagePresentUETNLABindingReleaseRequest           int = 46
	InitiatingMessagePresentUplinkNASTransport                    int = 47
	InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport   int = 48
	InitiatingMessagePresentUplinkRANConfigurationTransfer        int = 49
	InitiatingMessagePresentUplinkRANStatusTransfer               int = 50
	InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport      int = 51
	InitiatingMessagePresentWriteReplaceWarningRequest            int = 52
	InitiatingMessagePresentUplinkRIMInformationTransfer          int = 53
	InitiatingMessagePresentDownlinkRIMInformationTransfer        int = 54
)

type InitiatingMessageValue struct {
	Present                               int
	AMFConfigurationUpdate                *AMFConfigurationUpdate                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeAMFConfigurationUpdate"`
	AMFStatusIndication                   *AMFStatusIndication                   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeAMFStatusIndication"`
	CellTrafficTrace                      *CellTrafficTrace                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProcedureCode:ProcedureCodeCellTrafficTrace"`
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
	SuccessfulOutcomePresentAMFConfigurationUpdateAcknowledge int = 0
	SuccessfulOutcomePresentHandoverCancelAcknowledge         int = 1
	SuccessfulOutcomePresentHandoverCommand                   int = 2
	SuccessfulOutcomePresentHandoverRequestAcknowledge        int = 3
	SuccessfulOutcomePresentInitialContextSetupResponse       int = 4
	SuccessfulOutcomePresentNGResetAcknowledge                int = 5
	SuccessfulOutcomePresentNGSetupResponse                   int = 6
	SuccessfulOutcomePresentPathSwitchRequestAcknowledge      int = 7
	SuccessfulOutcomePresentPDUSessionResourceModifyResponse  int = 8
	SuccessfulOutcomePresentPDUSessionResourceModifyConfirm   int = 9
	SuccessfulOutcomePresentPDUSessionResourceReleaseResponse int = 10
	SuccessfulOutcomePresentPDUSessionResourceSetupResponse   int = 11
	SuccessfulOutcomePresentPWSCancelResponse                 int = 12
	SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge int = 13
	SuccessfulOutcomePresentUEContextModificationResponse     int = 14
	SuccessfulOutcomePresentUEContextReleaseComplete          int = 15
	SuccessfulOutcomePresentUERadioCapabilityCheckResponse    int = 16
	SuccessfulOutcomePresentWriteReplaceWarningResponse       int = 17
)

type SuccessfulOutcomeValue struct {
	Present                           int
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
	UnsuccessfulOutcomePresentAMFConfigurationUpdateFailure int = 0
	UnsuccessfulOutcomePresentHandoverPreparationFailure    int = 1
	UnsuccessfulOutcomePresentHandoverFailure               int = 2
	UnsuccessfulOutcomePresentInitialContextSetupFailure    int = 3
	UnsuccessfulOutcomePresentNGSetupFailure                int = 4
	UnsuccessfulOutcomePresentPathSwitchRequestFailure      int = 5
	UnsuccessfulOutcomePresentRANConfigurationUpdateFailure int = 6
	UnsuccessfulOutcomePresentUEContextModificationFailure  int = 7
)

type UnsuccessfulOutcomeValue struct {
	Present                       int
	AMFConfigurationUpdateFailure *AMFConfigurationUpdateFailure `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeAMFConfigurationUpdate"`
	HandoverPreparationFailure    *HandoverPreparationFailure    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverPreparation"`
	HandoverFailure               *HandoverFailure               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverResourceAllocation"`
	InitialContextSetupFailure    *InitialContextSetupFailure    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeInitialContextSetup"`
	NGSetupFailure                *NGSetupFailure                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGSetup"`
	PathSwitchRequestFailure      *PathSwitchRequestFailure      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePathSwitchRequest"`
	RANConfigurationUpdateFailure *RANConfigurationUpdateFailure `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeRANConfigurationUpdate"`
	UEContextModificationFailure  *UEContextModificationFailure  `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextModification"`
}
