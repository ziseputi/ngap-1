// Created By HaoDHH-245789 VHT2020

package ngapType

const (
	NGAPPDUPresentInitiatingMessage   int = 0
	NGAPPDUPresentSuccessfulOutcome   int = 1
	NGAPPDUPresentUnsuccessfulOutcome int = 2
	/* Extensions may appear below */
)

type NGAPPDU struct {
	Present              int
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
	InitiatingMessageValuePresentAMFConfigurationUpdate                int = 0
	InitiatingMessageValuePresentDeactivateTrace                       int = 1
	InitiatingMessageValuePresentDownlinkNASTransport                  int = 2
	InitiatingMessageValuePresentDownlinkNonUEAssociatedNRPPaTransport int = 3
	InitiatingMessageValuePresentDownlinkRANConfigurationTransfer      int = 4
	InitiatingMessageValuePresentDownlinkRANStatusTransfer             int = 5
	InitiatingMessageValuePresentDownlinkUEAssociatedNRPPaTransport    int = 6
	InitiatingMessageValuePresentErrorIndication                       int = 7
	InitiatingMessageValuePresentHandoverCancel                        int = 8
	InitiatingMessageValuePresentHandoverNotify                        int = 9
	InitiatingMessageValuePresentHandoverRequired                      int = 10
	InitiatingMessageValuePresentHandoverRequest                       int = 11
	InitiatingMessageValuePresentInitialContextSetupRequest            int = 12
	InitiatingMessageValuePresentInitialUEMessage                      int = 13
	InitiatingMessageValuePresentLocationReport                        int = 14
	InitiatingMessageValuePresentLocationReportingControl              int = 15
	InitiatingMessageValuePresentLocationReportingFailureIndication    int = 16
	InitiatingMessageValuePresentNASNonDeliveryIndication              int = 17
	InitiatingMessageValuePresentNGReset                               int = 18
	InitiatingMessageValuePresentNGSetupRequest                        int = 19
	InitiatingMessageValuePresentOverloadStart                         int = 20
	InitiatingMessageValuePresentOverloadStop                          int = 21
	InitiatingMessageValuePresentPaging                                int = 22
	InitiatingMessageValuePresentPathSwitchRequest                     int = 23
	InitiatingMessageValuePresentPDUSessionResourceModifyRequest       int = 24
	InitiatingMessageValuePresentPDUSessionResourceModifyIndication    int = 25
	InitiatingMessageValuePresentPDUSessionResourceNotify              int = 26
	InitiatingMessageValuePresentPDUSessionResourceReleaseCommand      int = 27
	InitiatingMessageValuePresentPDUSessionResourceSetupRequest        int = 28
	InitiatingMessageValuePresentPrivateMessage                        int = 29
	InitiatingMessageValuePresentPWSCancelRequest                      int = 30
	InitiatingMessageValuePresentPWSFailureIndication                  int = 31
	InitiatingMessageValuePresentPWSRestartIndication                  int = 32
	InitiatingMessageValuePresentRANConfigurationUpdate                int = 33
	InitiatingMessageValuePresentRerouteNASRequest                     int = 34
	InitiatingMessageValuePresentRRCInactiveTransitionReport           int = 35
	InitiatingMessageValuePresentSecondaryRATDataUsageReport           int = 36
	InitiatingMessageValuePresentTraceFailureIndication                int = 37
	InitiatingMessageValuePresentTraceStart                            int = 38
	InitiatingMessageValuePresentUEContextModificationRequest          int = 39
	InitiatingMessageValuePresentUEContextReleaseCommand               int = 40
	InitiatingMessageValuePresentUEContextReleaseRequest               int = 41
	InitiatingMessageValuePresentUERadioCapabilityCheckRequest         int = 42
	InitiatingMessageValuePresentUERadioCapabilityInfoIndication       int = 43
	InitiatingMessageValuePresentUETNLABindingReleaseRequest           int = 44
	InitiatingMessageValuePresentUplinkNASTransport                    int = 45
	InitiatingMessageValuePresentUplinkNonUEAssociatedNRPPaTransport   int = 46
	InitiatingMessageValuePresentUplinkRANConfigurationTransfer        int = 47
	InitiatingMessageValuePresentUplinkRANStatusTransfer               int = 48
	InitiatingMessageValuePresentUplinkUEAssociatedNRPPaTransport      int = 49
	InitiatingMessageValuePresentWriteReplaceWarningRequest            int = 50
	InitiatingMessageValuePresentUplinkRIMInformationTransfer          int = 51
	InitiatingMessageValuePresentDownlinkRIMInformationTransfer        int = 52
)

type InitiatingMessageValue struct {
	Present                                int
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
	SuccessfulOutcomeValuePresentAMFConfigurationUpdateAcknowledge int = 0
	SuccessfulOutcomeValuePresentHandoverCancelAcknowledge         int = 1
	SuccessfulOutcomeValuePresentHandoverCommand                   int = 2
	SuccessfulOutcomeValuePresentHandoverRequestAcknowledge        int = 3
	SuccessfulOutcomeValuePresentInitialContextSetupResponse       int = 4
	SuccessfulOutcomeValuePresentNGResetAcknowledge                int = 5
	SuccessfulOutcomeValuePresentNGSetupResponse                   int = 6
	SuccessfulOutcomeValuePresentPathSwitchRequestAcknowledge      int = 7
	SuccessfulOutcomeValuePresentPDUSessionResourceModifyResponse  int = 8
	SuccessfulOutcomeValuePresentPDUSessionResourceModifyConfirm   int = 9
	SuccessfulOutcomeValuePresentPDUSessionResourceReleaseResponse int = 10
	SuccessfulOutcomeValuePresentPDUSessionResourceSetupResponse   int = 11
	SuccessfulOutcomeValuePresentPWSCancelResponse                 int = 12
	SuccessfulOutcomeValuePresentRANConfigurationUpdateAcknowledge int = 13
	SuccessfulOutcomeValuePresentUEContextModificationResponse     int = 14
	SuccessfulOutcomeValuePresentUEContextReleaseComplete          int = 15
	SuccessfulOutcomeValuePresentUERadioCapabilityCheckResponse    int = 16
	SuccessfulOutcomeValuePresentWriteReplaceWarningResponse       int = 17
)

type SuccessfulOutcomeValue struct {
	Present                            int
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
	UnsuccessfulOutcomeValuePresentAMFConfigurationUpdateFailure int = 0
	UnsuccessfulOutcomeValuePresentHandoverPreparationFailure    int = 1
	UnsuccessfulOutcomeValuePresentHandoverFailure               int = 2
	UnsuccessfulOutcomeValuePresentInitialContextSetupFailure    int = 3
	UnsuccessfulOutcomeValuePresentNGSetupFailure                int = 4
	UnsuccessfulOutcomeValuePresentPathSwitchRequestFailure      int = 5
	UnsuccessfulOutcomeValuePresentRANConfigurationUpdateFailure int = 6
	UnsuccessfulOutcomeValuePresentUEContextModificationFailure  int = 7
)

type UnsuccessfulOutcomeValue struct {
	Present                        int
	AMFConfigurationUpdateFailure *AMFConfigurationUpdateFailure `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeAMFConfigurationUpdate"`
	HandoverPreparationFailure    *HandoverPreparationFailure    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverPreparation"`
	HandoverFailure               *HandoverFailure               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeHandoverResourceAllocation"`
	InitialContextSetupFailure    *InitialContextSetupFailure    `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeInitialContextSetup"`
	NGSetupFailure                *NGSetupFailure                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeNGSetup"`
	PathSwitchRequestFailure      *PathSwitchRequestFailure      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodePathSwitchRequest"`
	RANConfigurationUpdateFailure *RANConfigurationUpdateFailure `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeRANConfigurationUpdate"`
	UEContextModificationFailure  *UEContextModificationFailure  `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProcedureCode:ProcedureCodeUEContextModification"`
}
