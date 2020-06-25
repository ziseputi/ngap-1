// Created By HaoDHH-245789 VHT2020

package ngapType

type PDUSessionResourceSetupRequest struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceSetupRequestIEs
}

type ProtocolIEContainerPDUSessionResourceSetupRequestIEs struct {
	List []PDUSessionResourceSetupRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceSetupRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupRequestIEsPresentAMFUENGAPID                      int = 0
	PDUSessionResourceSetupRequestIEsPresentRANUENGAPID                      int = 1
	PDUSessionResourceSetupRequestIEsPresentRANPagingPriority                int = 2
	PDUSessionResourceSetupRequestIEsPresentNASPDU                           int = 3
	PDUSessionResourceSetupRequestIEsPresentPDUSessionResourceSetupListSUReq int = 4
	PDUSessionResourceSetupRequestIEsPresentUEAggregateMaximumBitRate        int = 5
)

type PDUSessionResourceSetupRequestIEsTypeValue struct {
	Present                          int
	AMFUENGAPID                      *AMFUENGAPID                      `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                      *RANUENGAPID                      `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority                *RANPagingPriority                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	NASPDU                           *NASPDU                           `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	PDUSessionResourceSetupListSUReq *PDUSessionResourceSetupListSUReq `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListSUReq"`
	UEAggregateMaximumBitRate        *UEAggregateMaximumBitRate        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
}

type PDUSessionResourceSetupResponse struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceSetupResponseIEs
}

type ProtocolIEContainerPDUSessionResourceSetupResponseIEs struct {
	List []PDUSessionResourceSetupResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                 `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceSetupResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupResponseIEsPresentAMFUENGAPID                              int = 0
	PDUSessionResourceSetupResponseIEsPresentRANUENGAPID                              int = 1
	PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceSetupListSURes         int = 2
	PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceFailedToSetupListSURes int = 3
	PDUSessionResourceSetupResponseIEsPresentCriticalityDiagnostics                   int = 4
)

type PDUSessionResourceSetupResponseIEsTypeValue struct {
	Present                                  int
	AMFUENGAPID                              *AMFUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                              *RANUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceSetupListSURes         *PDUSessionResourceSetupListSURes         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListSURes"`
	PDUSessionResourceFailedToSetupListSURes *PDUSessionResourceFailedToSetupListSURes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListSURes"`
	CriticalityDiagnostics                   *CriticalityDiagnostics                   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PDUSessionResourceReleaseCommand struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceReleaseCommandIEs
}

type ProtocolIEContainerPDUSessionResourceReleaseCommandIEs struct {
	List []PDUSessionResourceReleaseCommandIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceReleaseCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                  `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceReleaseCommandIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceReleaseCommandIEsPresentAMFUENGAPID                           int = 0
	PDUSessionResourceReleaseCommandIEsPresentRANUENGAPID                           int = 1
	PDUSessionResourceReleaseCommandIEsPresentRANPagingPriority                     int = 2
	PDUSessionResourceReleaseCommandIEsPresentNASPDU                                int = 3
	PDUSessionResourceReleaseCommandIEsPresentPDUSessionResourceToReleaseListRelCmd int = 4
)

type PDUSessionResourceReleaseCommandIEsTypeValue struct {
	Present                               int
	AMFUENGAPID                           *AMFUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                           *RANUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority                     *RANPagingPriority                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	NASPDU                                *NASPDU                                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNASPDU"`
	PDUSessionResourceToReleaseListRelCmd *PDUSessionResourceToReleaseListRelCmd `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceToReleaseListRelCmd"`
}

type PDUSessionResourceReleaseResponse struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceReleaseResponseIEs
}

type ProtocolIEContainerPDUSessionResourceReleaseResponseIEs struct {
	List []PDUSessionResourceReleaseResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceReleaseResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                   `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceReleaseResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceReleaseResponseIEsPresentAMFUENGAPID                          int = 0
	PDUSessionResourceReleaseResponseIEsPresentRANUENGAPID                          int = 1
	PDUSessionResourceReleaseResponseIEsPresentPDUSessionResourceReleasedListRelRes int = 2
	PDUSessionResourceReleaseResponseIEsPresentUserLocationInformation              int = 3
	PDUSessionResourceReleaseResponseIEsPresentCriticalityDiagnostics               int = 4
)

type PDUSessionResourceReleaseResponseIEsTypeValue struct {
	Present                              int
	AMFUENGAPID                          *AMFUENGAPID                          `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                          *RANUENGAPID                          `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceReleasedListRelRes *PDUSessionResourceReleasedListRelRes `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListRelRes"`
	UserLocationInformation              *UserLocationInformation              `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	CriticalityDiagnostics               *CriticalityDiagnostics               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PDUSessionResourceModifyRequest struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyRequestIEs
}

type ProtocolIEContainerPDUSessionResourceModifyRequestIEs struct {
	List []PDUSessionResourceModifyRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                 `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceModifyRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyRequestIEsPresentAMFUENGAPID                        int = 0
	PDUSessionResourceModifyRequestIEsPresentRANUENGAPID                        int = 1
	PDUSessionResourceModifyRequestIEsPresentRANPagingPriority                  int = 2
	PDUSessionResourceModifyRequestIEsPresentPDUSessionResourceModifyListModReq int = 3
)

type PDUSessionResourceModifyRequestIEsTypeValue struct {
	Present                            int
	AMFUENGAPID                        *AMFUENGAPID                        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                        *RANUENGAPID                        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority                  *RANPagingPriority                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	PDUSessionResourceModifyListModReq *PDUSessionResourceModifyListModReq `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModReq"`
}

type PDUSessionResourceModifyResponse struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyResponseIEs
}

type ProtocolIEContainerPDUSessionResourceModifyResponseIEs struct {
	List []PDUSessionResourceModifyResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                  `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceModifyResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyResponseIEsPresentAMFUENGAPID                                int = 0
	PDUSessionResourceModifyResponseIEsPresentRANUENGAPID                                int = 1
	PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceModifyListModRes         int = 2
	PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceFailedToModifyListModRes int = 3
	PDUSessionResourceModifyResponseIEsPresentUserLocationInformation                    int = 4
	PDUSessionResourceModifyResponseIEsPresentCriticalityDiagnostics                     int = 5
)

type PDUSessionResourceModifyResponseIEsTypeValue struct {
	Present                                    int
	AMFUENGAPID                                *AMFUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                *RANUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceModifyListModRes         *PDUSessionResourceModifyListModRes         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModRes"`
	PDUSessionResourceFailedToModifyListModRes *PDUSessionResourceFailedToModifyListModRes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToModifyListModRes"`
	UserLocationInformation                    *UserLocationInformation                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	CriticalityDiagnostics                     *CriticalityDiagnostics                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PDUSessionResourceNotify struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceNotifyIEs
}

type ProtocolIEContainerPDUSessionResourceNotifyIEs struct {
	List []PDUSessionResourceNotifyIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceNotifyIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceNotifyIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceNotifyIEsPresentAMFUENGAPID                       int = 0
	PDUSessionResourceNotifyIEsPresentRANUENGAPID                       int = 1
	PDUSessionResourceNotifyIEsPresentPDUSessionResourceNotifyList      int = 2
	PDUSessionResourceNotifyIEsPresentPDUSessionResourceReleasedListNot int = 3
	PDUSessionResourceNotifyIEsPresentUserLocationInformation           int = 4
)

type PDUSessionResourceNotifyIEsTypeValue struct {
	Present                           int
	AMFUENGAPID                       *AMFUENGAPID                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                       *RANUENGAPID                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceNotifyList      *PDUSessionResourceNotifyList      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceNotifyList"`
	PDUSessionResourceReleasedListNot *PDUSessionResourceReleasedListNot `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListNot"`
	UserLocationInformation           *UserLocationInformation           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type PDUSessionResourceModifyIndication struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyIndicationIEs
}

type ProtocolIEContainerPDUSessionResourceModifyIndicationIEs struct {
	List []PDUSessionResourceModifyIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                    `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceModifyIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyIndicationIEsPresentAMFUENGAPID                        int = 0
	PDUSessionResourceModifyIndicationIEsPresentRANUENGAPID                        int = 1
	PDUSessionResourceModifyIndicationIEsPresentPDUSessionResourceModifyListModInd int = 2
	PDUSessionResourceModifyIndicationIEsPresentUserLocationInformation            int = 3
)

type PDUSessionResourceModifyIndicationIEsTypeValue struct {
	Present                            int
	AMFUENGAPID                        *AMFUENGAPID                        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                        *RANUENGAPID                        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceModifyListModInd *PDUSessionResourceModifyListModInd `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModInd"`
	UserLocationInformation            *UserLocationInformation            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type PDUSessionResourceModifyConfirm struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyConfirmIEs
}

type ProtocolIEContainerPDUSessionResourceModifyConfirmIEs struct {
	List []PDUSessionResourceModifyConfirmIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyConfirmIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                 `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceModifyConfirmIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyConfirmIEsPresentAMFUENGAPID                                int = 0
	PDUSessionResourceModifyConfirmIEsPresentRANUENGAPID                                int = 1
	PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceModifyListModCfm         int = 2
	PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceFailedToModifyListModCfm int = 3
	PDUSessionResourceModifyConfirmIEsPresentCriticalityDiagnostics                     int = 4
)

type PDUSessionResourceModifyConfirmIEsTypeValue struct {
	Present                                    int
	AMFUENGAPID                                *AMFUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                *RANUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceModifyListModCfm         *PDUSessionResourceModifyListModCfm         `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModCfm"`
	PDUSessionResourceFailedToModifyListModCfm *PDUSessionResourceFailedToModifyListModCfm `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm"`
	CriticalityDiagnostics                     *CriticalityDiagnostics                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type InitialContextSetupRequest struct {
	ProtocolIEs ProtocolIEContainerInitialContextSetupRequestIEs
}

type ProtocolIEContainerInitialContextSetupRequestIEs struct {
	List []InitialContextSetupRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialContextSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	TypeValue    InitialContextSetupRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupRequestIEsPresentAMFUENGAPID                                 int = 0
	InitialContextSetupRequestIEsPresentRANUENGAPID                                 int = 1
	InitialContextSetupRequestIEsPresentOldAMF                                      int = 2
	InitialContextSetupRequestIEsPresentUEAggregateMaximumBitRate                   int = 3
	InitialContextSetupRequestIEsPresentCoreNetworkAssistanceInformationForInactive int = 4
	InitialContextSetupRequestIEsPresentGUAMI                                       int = 5
	InitialContextSetupRequestIEsPresentPDUSessionResourceSetupListCxtReq           int = 6
	InitialContextSetupRequestIEsPresentAllowedNSSAI                                int = 7
	InitialContextSetupRequestIEsPresentUESecurityCapabilities                      int = 8
	InitialContextSetupRequestIEsPresentSecurityKey                                 int = 9
	InitialContextSetupRequestIEsPresentTraceActivation                             int = 10
	InitialContextSetupRequestIEsPresentMobilityRestrictionList                     int = 11
	InitialContextSetupRequestIEsPresentUERadioCapability                           int = 12
	InitialContextSetupRequestIEsPresentIndexToRFSP                                 int = 13
	InitialContextSetupRequestIEsPresentMaskedIMEISV                                int = 14
	InitialContextSetupRequestIEsPresentNASPDU                                      int = 15
	InitialContextSetupRequestIEsPresentEmergencyFallbackIndicator                  int = 16
	InitialContextSetupRequestIEsPresentRRCInactiveTransitionReportRequest          int = 17
	InitialContextSetupRequestIEsPresentUERadioCapabilityForPaging                  int = 18
	InitialContextSetupRequestIEsPresentRedirectionVoiceFallback                    int = 19
	InitialContextSetupRequestIEsPresentLocationReportingRequestType                int = 20
	InitialContextSetupRequestIEsPresentCNAssistedRANTuning                         int = 21
	InitialContextSetupRequestIEsPresentSRVCCOperationPossible                      int = 22
)

type InitialContextSetupRequestIEsTypeValue struct {
	Present                                     int
	AMFUENGAPID                                 *AMFUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                 *RANUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	OldAMF                                      *AMFName                                     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDOldAMF"`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `vht:"Presence:PresenceConditional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	GUAMI                                       *GUAMI                                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGUAMI"`
	PDUSessionResourceSetupListCxtReq           *PDUSessionResourceSetupListCxtReq           `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListCxtReq"`
	AllowedNSSAI                                *AllowedNSSAI                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	UESecurityCapabilities                      *UESecurityCapabilities                      `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityKey                                 *SecurityKey                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityKey"`
	TraceActivation                             *TraceActivation                             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceActivation"`
	MobilityRestrictionList                     *MobilityRestrictionList                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMobilityRestrictionList"`
	UERadioCapability                           *UERadioCapability                           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapability"`
	IndexToRFSP                                 *IndexToRFSP                                 `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDIndexToRFSP"`
	MaskedIMEISV                                *MaskedIMEISV                                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMaskedIMEISV"`
	NASPDU                                      *NASPDU                                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNASPDU"`
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator                  `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDEmergencyFallbackIndicator"`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	UERadioCapabilityForPaging                  *UERadioCapabilityForPaging                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapabilityForPaging"`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRedirectionVoiceFallback"`
	LocationReportingRequestType                *LocationReportingRequestType                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible                      *SRVCCOperationPossible                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type InitialContextSetupResponse struct {
	ProtocolIEs ProtocolIEContainerInitialContextSetupResponseIEs
}

type ProtocolIEContainerInitialContextSetupResponseIEs struct {
	List []InitialContextSetupResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialContextSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	TypeValue    InitialContextSetupResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupResponseIEsPresentAMFUENGAPID                               int = 0
	InitialContextSetupResponseIEsPresentRANUENGAPID                               int = 1
	InitialContextSetupResponseIEsPresentPDUSessionResourceSetupListCxtRes         int = 2
	InitialContextSetupResponseIEsPresentPDUSessionResourceFailedToSetupListCxtRes int = 3
	InitialContextSetupResponseIEsPresentCriticalityDiagnostics                    int = 4
)

type InitialContextSetupResponseIEsTypeValue struct {
	Present                                   int
	AMFUENGAPID                               *AMFUENGAPID                               `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                               *RANUENGAPID                               `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceSetupListCxtRes         *PDUSessionResourceSetupListCxtRes         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListCxtRes"`
	PDUSessionResourceFailedToSetupListCxtRes *PDUSessionResourceFailedToSetupListCxtRes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes"`
	CriticalityDiagnostics                    *CriticalityDiagnostics                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type InitialContextSetupFailure struct {
	ProtocolIEs ProtocolIEContainerInitialContextSetupFailureIEs
}

type ProtocolIEContainerInitialContextSetupFailureIEs struct {
	List []InitialContextSetupFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialContextSetupFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	TypeValue    InitialContextSetupFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupFailureIEsPresentAMFUENGAPID                                int = 0
	InitialContextSetupFailureIEsPresentRANUENGAPID                                int = 1
	InitialContextSetupFailureIEsPresentPDUSessionResourceFailedToSetupListCxtFail int = 2
	InitialContextSetupFailureIEsPresentCause                                      int = 3
	InitialContextSetupFailureIEsPresentCriticalityDiagnostics                     int = 4
)

type InitialContextSetupFailureIEsTypeValue struct {
	Present                                    int
	AMFUENGAPID                                *AMFUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                *RANUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceFailedToSetupListCxtFail *PDUSessionResourceFailedToSetupListCxtFail `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail"`
	Cause                                      *Cause                                      `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics                     *CriticalityDiagnostics                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UEContextReleaseRequest struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseRequestIEs
}

type ProtocolIEContainerUEContextReleaseRequestIEs struct {
	List []UEContextReleaseRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextReleaseRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	TypeValue    UEContextReleaseRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseRequestIEsPresentAMFUENGAPID                     int = 0
	UEContextReleaseRequestIEsPresentRANUENGAPID                     int = 1
	UEContextReleaseRequestIEsPresentPDUSessionResourceListCxtRelReq int = 2
	UEContextReleaseRequestIEsPresentCause                           int = 3
)

type UEContextReleaseRequestIEsTypeValue struct {
	Present                         int
	AMFUENGAPID                     *AMFUENGAPID                     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                     *RANUENGAPID                     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceListCxtRelReq *PDUSessionResourceListCxtRelReq `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceListCxtRelReq"`
	Cause                           *Cause                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type UEContextReleaseCommand struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseCommandIEs
}

type ProtocolIEContainerUEContextReleaseCommandIEs struct {
	List []UEContextReleaseCommandIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextReleaseCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	TypeValue    UEContextReleaseCommandIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseCommandIEsPresentUENGAPIDs int = 0
	UEContextReleaseCommandIEsPresentCause     int = 1
)

type UEContextReleaseCommandIEsTypeValue struct {
	Present   int
	UENGAPIDs *UENGAPIDs `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUENGAPIDs"`
	Cause     *Cause     `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type UEContextReleaseComplete struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseCompleteIEs
}

type ProtocolIEContainerUEContextReleaseCompleteIEs struct {
	List []UEContextReleaseCompleteIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextReleaseCompleteIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	TypeValue    UEContextReleaseCompleteIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseCompleteIEsPresentAMFUENGAPID                                int = 0
	UEContextReleaseCompleteIEsPresentRANUENGAPID                                int = 1
	UEContextReleaseCompleteIEsPresentUserLocationInformation                    int = 2
	UEContextReleaseCompleteIEsPresentInfoOnRecommendedCellsAndRANNodesForPaging int = 3
	UEContextReleaseCompleteIEsPresentPDUSessionResourceListCxtRelCpl            int = 4
	UEContextReleaseCompleteIEsPresentCriticalityDiagnostics                     int = 5
)

type UEContextReleaseCompleteIEsTypeValue struct {
	Present                                    int
	AMFUENGAPID                                *AMFUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                *RANUENGAPID                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UserLocationInformation                    *UserLocationInformation                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	InfoOnRecommendedCellsAndRANNodesForPaging *InfoOnRecommendedCellsAndRANNodesForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging"`
	PDUSessionResourceListCxtRelCpl            *PDUSessionResourceListCxtRelCpl            `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceListCxtRelCpl"`
	CriticalityDiagnostics                     *CriticalityDiagnostics                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UEContextModificationRequest struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationRequestIEs
}

type ProtocolIEContainerUEContextModificationRequestIEs struct {
	List []UEContextModificationRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextModificationRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	TypeValue    UEContextModificationRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationRequestIEsPresentAMFUENGAPID                                 int = 0
	UEContextModificationRequestIEsPresentRANUENGAPID                                 int = 1
	UEContextModificationRequestIEsPresentRANPagingPriority                           int = 2
	UEContextModificationRequestIEsPresentSecurityKey                                 int = 3
	UEContextModificationRequestIEsPresentIndexToRFSP                                 int = 4
	UEContextModificationRequestIEsPresentUEAggregateMaximumBitRate                   int = 5
	UEContextModificationRequestIEsPresentUESecurityCapabilities                      int = 6
	UEContextModificationRequestIEsPresentCoreNetworkAssistanceInformationForInactive int = 7
	UEContextModificationRequestIEsPresentEmergencyFallbackIndicator                  int = 8
	UEContextModificationRequestIEsPresentNewAMFUENGAPID                              int = 9
	UEContextModificationRequestIEsPresentRRCInactiveTransitionReportRequest          int = 10
	UEContextModificationRequestIEsPresentNewGUAMI                                    int = 11
	UEContextModificationRequestIEsPresentCNAssistedRANTuning                         int = 12
	UEContextModificationRequestIEsPresentSRVCCOperationPossible                      int = 13
)

type UEContextModificationRequestIEsTypeValue struct {
	Present                                     int
	AMFUENGAPID                                 *AMFUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                 *RANUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority                           *RANPagingPriority                           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	SecurityKey                                 *SecurityKey                                 `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityKey"`
	IndexToRFSP                                 *IndexToRFSP                                 `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDIndexToRFSP"`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	UESecurityCapabilities                      *UESecurityCapabilities                      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator                  `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDEmergencyFallbackIndicator"`
	NewAMFUENGAPID                              *AMFUENGAPID                                 `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewAMFUENGAPID"`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	NewGUAMI                                    *GUAMI                                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewGUAMI"`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible                      *SRVCCOperationPossible                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type UEContextModificationResponse struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationResponseIEs
}

type ProtocolIEContainerUEContextModificationResponseIEs struct {
	List []UEContextModificationResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextModificationResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    UEContextModificationResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationResponseIEsPresentAMFUENGAPID             int = 0
	UEContextModificationResponseIEsPresentRANUENGAPID             int = 1
	UEContextModificationResponseIEsPresentRRCState                int = 2
	UEContextModificationResponseIEsPresentUserLocationInformation int = 3
	UEContextModificationResponseIEsPresentCriticalityDiagnostics  int = 4
)

type UEContextModificationResponseIEsTypeValue struct {
	Present                 int
	AMFUENGAPID             *AMFUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID             *RANUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RRCState                *RRCState                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCState"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	CriticalityDiagnostics  *CriticalityDiagnostics  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UEContextModificationFailure struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationFailureIEs
}

type ProtocolIEContainerUEContextModificationFailureIEs struct {
	List []UEContextModificationFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextModificationFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	TypeValue    UEContextModificationFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationFailureIEsPresentAMFUENGAPID            int = 0
	UEContextModificationFailureIEsPresentRANUENGAPID            int = 1
	UEContextModificationFailureIEsPresentCause                  int = 2
	UEContextModificationFailureIEsPresentCriticalityDiagnostics int = 3
)

type UEContextModificationFailureIEsTypeValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID            *RANUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type RRCInactiveTransitionReport struct {
	ProtocolIEs ProtocolIEContainerRRCInactiveTransitionReportIEs
}

type ProtocolIEContainerRRCInactiveTransitionReportIEs struct {
	List []RRCInactiveTransitionReportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type RRCInactiveTransitionReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	TypeValue    RRCInactiveTransitionReportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RRCInactiveTransitionReportIEsPresentAMFUENGAPID             int = 0
	RRCInactiveTransitionReportIEsPresentRANUENGAPID             int = 1
	RRCInactiveTransitionReportIEsPresentRRCState                int = 2
	RRCInactiveTransitionReportIEsPresentUserLocationInformation int = 3
)

type RRCInactiveTransitionReportIEsTypeValue struct {
	Present                 int
	AMFUENGAPID             *AMFUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID             *RANUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RRCState                *RRCState                `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCState"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type HandoverRequired struct {
	ProtocolIEs ProtocolIEContainerHandoverRequiredIEs
}

type ProtocolIEContainerHandoverRequiredIEs struct {
	List []HandoverRequiredIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverRequiredIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverRequiredIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequiredIEsPresentAMFUENGAPID                        int = 0
	HandoverRequiredIEsPresentRANUENGAPID                        int = 1
	HandoverRequiredIEsPresentHandoverType                       int = 2
	HandoverRequiredIEsPresentCause                              int = 3
	HandoverRequiredIEsPresentTargetID                           int = 4
	HandoverRequiredIEsPresentDirectForwardingPathAvailability   int = 5
	HandoverRequiredIEsPresentPDUSessionResourceListHORqd        int = 6
	HandoverRequiredIEsPresentSourceToTargetTransparentContainer int = 7
)

type HandoverRequiredIEsTypeValue struct {
	Present                            int
	AMFUENGAPID                        *AMFUENGAPID                        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                        *RANUENGAPID                        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	HandoverType                       *HandoverType                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	Cause                              *Cause                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TargetID                           *TargetID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetID"`
	DirectForwardingPathAvailability   *DirectForwardingPathAvailability   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDirectForwardingPathAvailability"`
	PDUSessionResourceListHORqd        *PDUSessionResourceListHORqd        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceListHORqd"`
	SourceToTargetTransparentContainer *SourceToTargetTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSourceToTargetTransparentContainer"`
}

type HandoverCommand struct {
	ProtocolIEs ProtocolIEContainerHandoverCommandIEs
}

type ProtocolIEContainerHandoverCommandIEs struct {
	List []HandoverCommandIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverCommandIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCommandIEsPresentAMFUENGAPID                          int = 0
	HandoverCommandIEsPresentRANUENGAPID                          int = 1
	HandoverCommandIEsPresentHandoverType                         int = 2
	HandoverCommandIEsPresentNASSecurityParametersFromNGRAN       int = 3
	HandoverCommandIEsPresentPDUSessionResourceHandoverList       int = 4
	HandoverCommandIEsPresentPDUSessionResourceToReleaseListHOCmd int = 5
	HandoverCommandIEsPresentTargetToSourceTransparentContainer   int = 6
	HandoverCommandIEsPresentCriticalityDiagnostics               int = 7
)

type HandoverCommandIEsTypeValue struct {
	Present                              int
	AMFUENGAPID                          *AMFUENGAPID                          `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                          *RANUENGAPID                          `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	HandoverType                         *HandoverType                         `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	NASSecurityParametersFromNGRAN       *NASSecurityParametersFromNGRAN       `vht:"Presence:PresenceConditional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASSecurityParametersFromNGRAN"`
	PDUSessionResourceHandoverList       *PDUSessionResourceHandoverList       `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceHandoverList"`
	PDUSessionResourceToReleaseListHOCmd *PDUSessionResourceToReleaseListHOCmd `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceToReleaseListHOCmd"`
	TargetToSourceTransparentContainer   *TargetToSourceTransparentContainer   `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetToSourceTransparentContainer"`
	CriticalityDiagnostics               *CriticalityDiagnostics               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverPreparationFailure struct {
	ProtocolIEs ProtocolIEContainerHandoverPreparationFailureIEs
}

type ProtocolIEContainerHandoverPreparationFailureIEs struct {
	List []HandoverPreparationFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverPreparationFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverPreparationFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverPreparationFailureIEsPresentAMFUENGAPID            int = 0
	HandoverPreparationFailureIEsPresentRANUENGAPID            int = 1
	HandoverPreparationFailureIEsPresentCause                  int = 2
	HandoverPreparationFailureIEsPresentCriticalityDiagnostics int = 3
)

type HandoverPreparationFailureIEsTypeValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID            *RANUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverRequest struct {
	ProtocolIEs ProtocolIEContainerHandoverRequestIEs
}

type ProtocolIEContainerHandoverRequestIEs struct {
	List []HandoverRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequestIEsPresentAMFUENGAPID                                 int = 0
	HandoverRequestIEsPresentHandoverType                                int = 1
	HandoverRequestIEsPresentCause                                       int = 2
	HandoverRequestIEsPresentUEAggregateMaximumBitRate                   int = 3
	HandoverRequestIEsPresentCoreNetworkAssistanceInformationForInactive int = 4
	HandoverRequestIEsPresentUESecurityCapabilities                      int = 5
	HandoverRequestIEsPresentSecurityContext                             int = 6
	HandoverRequestIEsPresentNewSecurityContextInd                       int = 7
	HandoverRequestIEsPresentNASC                                        int = 8
	HandoverRequestIEsPresentPDUSessionResourceSetupListHOReq            int = 9
	HandoverRequestIEsPresentAllowedNSSAI                                int = 10
	HandoverRequestIEsPresentTraceActivation                             int = 11
	HandoverRequestIEsPresentMaskedIMEISV                                int = 12
	HandoverRequestIEsPresentSourceToTargetTransparentContainer          int = 13
	HandoverRequestIEsPresentMobilityRestrictionList                     int = 14
	HandoverRequestIEsPresentLocationReportingRequestType                int = 15
	HandoverRequestIEsPresentRRCInactiveTransitionReportRequest          int = 16
	HandoverRequestIEsPresentGUAMI                                       int = 17
	HandoverRequestIEsPresentRedirectionVoiceFallback                    int = 18
	HandoverRequestIEsPresentCNAssistedRANTuning                         int = 19
	HandoverRequestIEsPresentSRVCCOperationPossible                      int = 20
)

type HandoverRequestIEsTypeValue struct {
	Present                                     int
	AMFUENGAPID                                 *AMFUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	HandoverType                                *HandoverType                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	Cause                                       *Cause                                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	UESecurityCapabilities                      *UESecurityCapabilities                      `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityContext                             *SecurityContext                             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityContext"`
	NewSecurityContextInd                       *NewSecurityContextInd                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewSecurityContextInd"`
	NASC                                        *NASPDU                                      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASC"`
	PDUSessionResourceSetupListHOReq            *PDUSessionResourceSetupListHOReq            `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListHOReq"`
	AllowedNSSAI                                *AllowedNSSAI                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	TraceActivation                             *TraceActivation                             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceActivation"`
	MaskedIMEISV                                *MaskedIMEISV                                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMaskedIMEISV"`
	SourceToTargetTransparentContainer          *SourceToTargetTransparentContainer          `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSourceToTargetTransparentContainer"`
	MobilityRestrictionList                     *MobilityRestrictionList                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMobilityRestrictionList"`
	LocationReportingRequestType                *LocationReportingRequestType                `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	GUAMI                                       *GUAMI                                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGUAMI"`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRedirectionVoiceFallback"`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible                      *SRVCCOperationPossible                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type HandoverRequestAcknowledge struct {
	ProtocolIEs ProtocolIEContainerHandoverRequestAcknowledgeIEs
}

type ProtocolIEContainerHandoverRequestAcknowledgeIEs struct {
	List []HandoverRequestAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverRequestAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverRequestAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequestAcknowledgeIEsPresentAMFUENGAPID                              int = 0
	HandoverRequestAcknowledgeIEsPresentRANUENGAPID                              int = 1
	HandoverRequestAcknowledgeIEsPresentPDUSessionResourceAdmittedList           int = 2
	HandoverRequestAcknowledgeIEsPresentPDUSessionResourceFailedToSetupListHOAck int = 3
	HandoverRequestAcknowledgeIEsPresentTargetToSourceTransparentContainer       int = 4
	HandoverRequestAcknowledgeIEsPresentCriticalityDiagnostics                   int = 5
)

type HandoverRequestAcknowledgeIEsTypeValue struct {
	Present                                  int
	AMFUENGAPID                              *AMFUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                              *RANUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceAdmittedList           *PDUSessionResourceAdmittedList           `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceAdmittedList"`
	PDUSessionResourceFailedToSetupListHOAck *PDUSessionResourceFailedToSetupListHOAck `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck"`
	TargetToSourceTransparentContainer       *TargetToSourceTransparentContainer       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetToSourceTransparentContainer"`
	CriticalityDiagnostics                   *CriticalityDiagnostics                   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverFailure struct {
	ProtocolIEs ProtocolIEContainerHandoverFailureIEs
}

type ProtocolIEContainerHandoverFailureIEs struct {
	List []HandoverFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverFailureIEsPresentAMFUENGAPID            int = 0
	HandoverFailureIEsPresentCause                  int = 1
	HandoverFailureIEsPresentCriticalityDiagnostics int = 2
)

type HandoverFailureIEsTypeValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverNotify struct {
	ProtocolIEs ProtocolIEContainerHandoverNotifyIEs
}

type ProtocolIEContainerHandoverNotifyIEs struct {
	List []HandoverNotifyIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverNotifyIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverNotifyIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverNotifyIEsPresentAMFUENGAPID             int = 0
	HandoverNotifyIEsPresentRANUENGAPID             int = 1
	HandoverNotifyIEsPresentUserLocationInformation int = 2
)

type HandoverNotifyIEsTypeValue struct {
	Present                 int
	AMFUENGAPID             *AMFUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID             *RANUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type PathSwitchRequest struct {
	ProtocolIEs ProtocolIEContainerPathSwitchRequestIEs
}

type ProtocolIEContainerPathSwitchRequestIEs struct {
	List []PathSwitchRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PathSwitchRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                   `vht:"Reference:ProtocolIEID"`
	TypeValue    PathSwitchRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestIEsPresentRANUENGAPID                              int = 0
	PathSwitchRequestIEsPresentSourceAMFUENGAPID                        int = 1
	PathSwitchRequestIEsPresentUserLocationInformation                  int = 2
	PathSwitchRequestIEsPresentUESecurityCapabilities                   int = 3
	PathSwitchRequestIEsPresentPDUSessionResourceToBeSwitchedDLList     int = 4
	PathSwitchRequestIEsPresentPDUSessionResourceFailedToSetupListPSReq int = 5
)

type PathSwitchRequestIEsTypeValue struct {
	Present                                  int
	RANUENGAPID                              *RANUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	SourceAMFUENGAPID                        *AMFUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSourceAMFUENGAPID"`
	UserLocationInformation                  *UserLocationInformation                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	UESecurityCapabilities                   *UESecurityCapabilities                   `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	PDUSessionResourceToBeSwitchedDLList     *PDUSessionResourceToBeSwitchedDLList     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceToBeSwitchedDLList"`
	PDUSessionResourceFailedToSetupListPSReq *PDUSessionResourceFailedToSetupListPSReq `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq"`
}

type PathSwitchRequestAcknowledge struct {
	ProtocolIEs ProtocolIEContainerPathSwitchRequestAcknowledgeIEs
}

type ProtocolIEContainerPathSwitchRequestAcknowledgeIEs struct {
	List []PathSwitchRequestAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PathSwitchRequestAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	TypeValue    PathSwitchRequestAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestAcknowledgeIEsPresentAMFUENGAPID                                 int = 0
	PathSwitchRequestAcknowledgeIEsPresentRANUENGAPID                                 int = 1
	PathSwitchRequestAcknowledgeIEsPresentUESecurityCapabilities                      int = 2
	PathSwitchRequestAcknowledgeIEsPresentSecurityContext                             int = 3
	PathSwitchRequestAcknowledgeIEsPresentNewSecurityContextInd                       int = 4
	PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceSwitchedList              int = 5
	PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceReleasedListPSAck         int = 6
	PathSwitchRequestAcknowledgeIEsPresentAllowedNSSAI                                int = 7
	PathSwitchRequestAcknowledgeIEsPresentCoreNetworkAssistanceInformationForInactive int = 8
	PathSwitchRequestAcknowledgeIEsPresentRRCInactiveTransitionReportRequest          int = 9
	PathSwitchRequestAcknowledgeIEsPresentCriticalityDiagnostics                      int = 10
	PathSwitchRequestAcknowledgeIEsPresentRedirectionVoiceFallback                    int = 11
	PathSwitchRequestAcknowledgeIEsPresentCNAssistedRANTuning                         int = 12
	PathSwitchRequestAcknowledgeIEsPresentSRVCCOperationPossible                      int = 13
)

type PathSwitchRequestAcknowledgeIEsTypeValue struct {
	Present                                     int
	AMFUENGAPID                                 *AMFUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                 *RANUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UESecurityCapabilities                      *UESecurityCapabilities                      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityContext                             *SecurityContext                             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityContext"`
	NewSecurityContextInd                       *NewSecurityContextInd                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewSecurityContextInd"`
	PDUSessionResourceSwitchedList              *PDUSessionResourceSwitchedList              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSwitchedList"`
	PDUSessionResourceReleasedListPSAck         *PDUSessionResourceReleasedListPSAck         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListPSAck"`
	AllowedNSSAI                                *AllowedNSSAI                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	CriticalityDiagnostics                      *CriticalityDiagnostics                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRedirectionVoiceFallback"`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible                      *SRVCCOperationPossible                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type PathSwitchRequestFailure struct {
	ProtocolIEs ProtocolIEContainerPathSwitchRequestFailureIEs
}

type ProtocolIEContainerPathSwitchRequestFailureIEs struct {
	List []PathSwitchRequestFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PathSwitchRequestFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	TypeValue    PathSwitchRequestFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestFailureIEsPresentAMFUENGAPID                          int = 0
	PathSwitchRequestFailureIEsPresentRANUENGAPID                          int = 1
	PathSwitchRequestFailureIEsPresentPDUSessionResourceReleasedListPSFail int = 2
	PathSwitchRequestFailureIEsPresentCriticalityDiagnostics               int = 3
)

type PathSwitchRequestFailureIEsTypeValue struct {
	Present                              int
	AMFUENGAPID                          *AMFUENGAPID                          `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                          *RANUENGAPID                          `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceReleasedListPSFail *PDUSessionResourceReleasedListPSFail `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListPSFail"`
	CriticalityDiagnostics               *CriticalityDiagnostics               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverCancel struct {
	ProtocolIEs ProtocolIEContainerHandoverCancelIEs
}

type ProtocolIEContainerHandoverCancelIEs struct {
	List []HandoverCancelIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverCancelIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverCancelIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCancelIEsPresentAMFUENGAPID int = 0
	HandoverCancelIEsPresentRANUENGAPID int = 1
	HandoverCancelIEsPresentCause       int = 2
)

type HandoverCancelIEsTypeValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause       *Cause       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type HandoverCancelAcknowledge struct {
	ProtocolIEs ProtocolIEContainerHandoverCancelAcknowledgeIEs
}

type ProtocolIEContainerHandoverCancelAcknowledgeIEs struct {
	List []HandoverCancelAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverCancelAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                           `vht:"Reference:ProtocolIEID"`
	TypeValue    HandoverCancelAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCancelAcknowledgeIEsPresentAMFUENGAPID            int = 0
	HandoverCancelAcknowledgeIEsPresentRANUENGAPID            int = 1
	HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics int = 2
)

type HandoverCancelAcknowledgeIEsTypeValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID            *RANUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UplinkRANStatusTransfer struct {
	ProtocolIEs ProtocolIEContainerUplinkRANStatusTransferIEs
}

type ProtocolIEContainerUplinkRANStatusTransferIEs struct {
	List []UplinkRANStatusTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkRANStatusTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	TypeValue    UplinkRANStatusTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRANStatusTransferIEsPresentAMFUENGAPID                           int = 0
	UplinkRANStatusTransferIEsPresentRANUENGAPID                           int = 1
	UplinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer int = 2
)

type UplinkRANStatusTransferIEsTypeValue struct {
	Present                               int
	AMFUENGAPID                           *AMFUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                           *RANUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANStatusTransferTransparentContainer *RANStatusTransferTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANStatusTransferTransparentContainer"`
}

type DownlinkRANStatusTransfer struct {
	ProtocolIEs ProtocolIEContainerDownlinkRANStatusTransferIEs
}

type ProtocolIEContainerDownlinkRANStatusTransferIEs struct {
	List []DownlinkRANStatusTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkRANStatusTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                           `vht:"Reference:ProtocolIEID"`
	TypeValue    DownlinkRANStatusTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRANStatusTransferIEsPresentAMFUENGAPID                           int = 0
	DownlinkRANStatusTransferIEsPresentRANUENGAPID                           int = 1
	DownlinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer int = 2
)

type DownlinkRANStatusTransferIEsTypeValue struct {
	Present                               int
	AMFUENGAPID                           *AMFUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                           *RANUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANStatusTransferTransparentContainer *RANStatusTransferTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANStatusTransferTransparentContainer"`
}

type Paging struct {
	ProtocolIEs ProtocolIEContainerPagingIEs
}

type ProtocolIEContainerPagingIEs struct {
	List []PagingIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PagingIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality        `vht:"Reference:ProtocolIEID"`
	TypeValue    PagingIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PagingIEsPresentUEPagingIdentity           int = 0
	PagingIEsPresentPagingDRX                  int = 1
	PagingIEsPresentTAIListForPaging           int = 2
	PagingIEsPresentPagingPriority             int = 3
	PagingIEsPresentUERadioCapabilityForPaging int = 4
	PagingIEsPresentPagingOrigin               int = 5
	PagingIEsPresentAssistanceDataForPaging    int = 6
)

type PagingIEsTypeValue struct {
	Present                    int
	UEPagingIdentity           *UEPagingIdentity           `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEPagingIdentity"`
	PagingDRX                  *PagingDRX                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingDRX"`
	TAIListForPaging           *TAIListForPaging           `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTAIListForPaging"`
	PagingPriority             *PagingPriority             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingPriority"`
	UERadioCapabilityForPaging *UERadioCapabilityForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapabilityForPaging"`
	PagingOrigin               *PagingOrigin               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingOrigin"`
	AssistanceDataForPaging    *AssistanceDataForPaging    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAssistanceDataForPaging"`
}

type InitialUEMessage struct {
	ProtocolIEs ProtocolIEContainerInitialUEMessageIEs
}

type ProtocolIEContainerInitialUEMessageIEs struct {
	List []InitialUEMessageIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialUEMessageIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	TypeValue    InitialUEMessageIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialUEMessageIEsPresentRANUENGAPID                         int = 0
	InitialUEMessageIEsPresentNASPDU                              int = 1
	InitialUEMessageIEsPresentUserLocationInformation             int = 2
	InitialUEMessageIEsPresentRRCEstablishmentCause               int = 3
	InitialUEMessageIEsPresentFiveGSTMSI                          int = 4
	InitialUEMessageIEsPresentAMFSetID                            int = 5
	InitialUEMessageIEsPresentUEContextRequest                    int = 6
	InitialUEMessageIEsPresentAllowedNSSAI                        int = 7
	InitialUEMessageIEsPresentSourceToTargetAMFInformationReroute int = 8
)

type InitialUEMessageIEsTypeValue struct {
	Present                             int
	RANUENGAPID                         *RANUENGAPID                         `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NASPDU                              *NASPDU                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	UserLocationInformation             *UserLocationInformation             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	RRCEstablishmentCause               *RRCEstablishmentCause               `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCEstablishmentCause"`
	FiveGSTMSI                          *FiveGSTMSI                          `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDFiveGSTMSI"`
	AMFSetID                            *AMFSetID                            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFSetID"`
	UEContextRequest                    *UEContextRequest                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEContextRequest"`
	AllowedNSSAI                        *AllowedNSSAI                        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSourceToTargetAMFInformationReroute"`
}

type DownlinkNASTransport struct {
	ProtocolIEs ProtocolIEContainerDownlinkNASTransportIEs
}

type ProtocolIEContainerDownlinkNASTransportIEs struct {
	List []DownlinkNASTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkNASTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	TypeValue    DownlinkNASTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkNASTransportIEsPresentAMFUENGAPID               int = 0
	DownlinkNASTransportIEsPresentRANUENGAPID               int = 1
	DownlinkNASTransportIEsPresentOldAMF                    int = 2
	DownlinkNASTransportIEsPresentRANPagingPriority         int = 3
	DownlinkNASTransportIEsPresentNASPDU                    int = 4
	DownlinkNASTransportIEsPresentMobilityRestrictionList   int = 5
	DownlinkNASTransportIEsPresentIndexToRFSP               int = 6
	DownlinkNASTransportIEsPresentUEAggregateMaximumBitRate int = 7
	DownlinkNASTransportIEsPresentAllowedNSSAI              int = 8
	DownlinkNASTransportIEsPresentSRVCCOperationPossible    int = 9
)

type DownlinkNASTransportIEsTypeValue struct {
	Present                   int
	AMFUENGAPID               *AMFUENGAPID               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID               *RANUENGAPID               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	OldAMF                    *AMFName                   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDOldAMF"`
	RANPagingPriority         *RANPagingPriority         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	NASPDU                    *NASPDU                    `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	MobilityRestrictionList   *MobilityRestrictionList   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMobilityRestrictionList"`
	IndexToRFSP               *IndexToRFSP               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDIndexToRFSP"`
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	AllowedNSSAI              *AllowedNSSAI              `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	SRVCCOperationPossible    *SRVCCOperationPossible    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type UplinkNASTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkNASTransportIEs
}

type ProtocolIEContainerUplinkNASTransportIEs struct {
	List []UplinkNASTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkNASTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	TypeValue    UplinkNASTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkNASTransportIEsPresentAMFUENGAPID             int = 0
	UplinkNASTransportIEsPresentRANUENGAPID             int = 1
	UplinkNASTransportIEsPresentNASPDU                  int = 2
	UplinkNASTransportIEsPresentUserLocationInformation int = 3
)

type UplinkNASTransportIEsTypeValue struct {
	Present                 int
	AMFUENGAPID             *AMFUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID             *RANUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NASPDU                  *NASPDU                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type NASNonDeliveryIndication struct {
	ProtocolIEs ProtocolIEContainerNASNonDeliveryIndicationIEs
}

type ProtocolIEContainerNASNonDeliveryIndicationIEs struct {
	List []NASNonDeliveryIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type NASNonDeliveryIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	TypeValue    NASNonDeliveryIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NASNonDeliveryIndicationIEsPresentAMFUENGAPID int = 0
	NASNonDeliveryIndicationIEsPresentRANUENGAPID int = 1
	NASNonDeliveryIndicationIEsPresentNASPDU      int = 2
	NASNonDeliveryIndicationIEsPresentCause       int = 3
)

type NASNonDeliveryIndicationIEsTypeValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NASPDU      *NASPDU      `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNASPDU"`
	Cause       *Cause       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type RerouteNASRequest struct {
	ProtocolIEs ProtocolIEContainerRerouteNASRequestIEs
}

type ProtocolIEContainerRerouteNASRequestIEs struct {
	List []RerouteNASRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type RerouteNASRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                   `vht:"Reference:ProtocolIEID"`
	TypeValue    RerouteNASRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RerouteNASRequestIEsPresentRANUENGAPID                         int = 0
	RerouteNASRequestIEsPresentAMFUENGAPID                         int = 1
	RerouteNASRequestIEsPresentNGAPMessage                         int = 2
	RerouteNASRequestIEsPresentAMFSetID                            int = 3
	RerouteNASRequestIEsPresentAllowedNSSAI                        int = 4
	RerouteNASRequestIEsPresentSourceToTargetAMFInformationReroute int = 5
)

type RerouteNASRequestIEsTypeValue struct {
	Present                             int
	RANUENGAPID                         *RANUENGAPID                         `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	AMFUENGAPID                         *AMFUENGAPID                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	NGAPMessage                         *OctetString                         `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNGAPMessage"`
	AMFSetID                            *AMFSetID                            `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFSetID"`
	AllowedNSSAI                        *AllowedNSSAI                        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSourceToTargetAMFInformationReroute"`
}

type NGSetupRequest struct {
	ProtocolIEs ProtocolIEContainerNGSetupRequestIEs
}

type ProtocolIEContainerNGSetupRequestIEs struct {
	List []NGSetupRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type NGSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    NGSetupRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupRequestIEsPresentGlobalRANNodeID        int = 0
	NGSetupRequestIEsPresentRANNodeName            int = 1
	NGSetupRequestIEsPresentSupportedTAList        int = 2
	NGSetupRequestIEsPresentDefaultPagingDRX       int = 3
	NGSetupRequestIEsPresentUERetentionInformation int = 4
)

type NGSetupRequestIEsTypeValue struct {
	Present                int
	GlobalRANNodeID        *GlobalRANNodeID        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	RANNodeName            *RANNodeName            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANNodeName"`
	SupportedTAList        *SupportedTAList        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSupportedTAList"`
	DefaultPagingDRX       *PagingDRX              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDefaultPagingDRX"`
	UERetentionInformation *UERetentionInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERetentionInformation"`
}

type NGSetupResponse struct {
	ProtocolIEs ProtocolIEContainerNGSetupResponseIEs
}

type ProtocolIEContainerNGSetupResponseIEs struct {
	List []NGSetupResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type NGSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	TypeValue    NGSetupResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupResponseIEsPresentAMFName                int = 0
	NGSetupResponseIEsPresentServedGUAMIList        int = 1
	NGSetupResponseIEsPresentRelativeAMFCapacity    int = 2
	NGSetupResponseIEsPresentPLMNSupportList        int = 3
	NGSetupResponseIEsPresentCriticalityDiagnostics int = 4
	NGSetupResponseIEsPresentUERetentionInformation int = 5
)

type NGSetupResponseIEsTypeValue struct {
	Present                int
	AMFName                *AMFName                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFName"`
	ServedGUAMIList        *ServedGUAMIList        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDServedGUAMIList"`
	RelativeAMFCapacity    *RelativeAMFCapacity    `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRelativeAMFCapacity"`
	PLMNSupportList        *PLMNSupportList        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPLMNSupportList"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
	UERetentionInformation *UERetentionInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERetentionInformation"`
}

type NGSetupFailure struct {
	ProtocolIEs ProtocolIEContainerNGSetupFailureIEs
}

type ProtocolIEContainerNGSetupFailureIEs struct {
	List []NGSetupFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type NGSetupFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    NGSetupFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupFailureIEsPresentCause                  int = 0
	NGSetupFailureIEsPresentTimeToWait             int = 1
	NGSetupFailureIEsPresentCriticalityDiagnostics int = 2
)

type NGSetupFailureIEsTypeValue struct {
	Present                int
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait             *TimeToWait             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type RANConfigurationUpdate struct {
	ProtocolIEs ProtocolIEContainerRANConfigurationUpdateIEs
}

type ProtocolIEContainerRANConfigurationUpdateIEs struct {
	List []RANConfigurationUpdateIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type RANConfigurationUpdateIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	TypeValue    RANConfigurationUpdateIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateIEsPresentRANNodeName                     int = 0
	RANConfigurationUpdateIEsPresentSupportedTAList                 int = 1
	RANConfigurationUpdateIEsPresentDefaultPagingDRX                int = 2
	RANConfigurationUpdateIEsPresentGlobalRANNodeID                 int = 3
	RANConfigurationUpdateIEsPresentNGRANTNLAssociationToRemoveList int = 4
)

type RANConfigurationUpdateIEsTypeValue struct {
	Present                         int
	RANNodeName                     *RANNodeName                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANNodeName"`
	SupportedTAList                 *SupportedTAList                 `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSupportedTAList"`
	DefaultPagingDRX                *PagingDRX                       `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDefaultPagingDRX"`
	GlobalRANNodeID                 *GlobalRANNodeID                 `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	NGRANTNLAssociationToRemoveList *NGRANTNLAssociationToRemoveList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNGRANTNLAssociationToRemoveList"`
}

type RANConfigurationUpdateAcknowledge struct {
	ProtocolIEs ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs
}

type ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs struct {
	List []RANConfigurationUpdateAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type RANConfigurationUpdateAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                   `vht:"Reference:ProtocolIEID"`
	TypeValue    RANConfigurationUpdateAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics int = 0
)

type RANConfigurationUpdateAcknowledgeIEsTypeValue struct {
	Present                int
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type RANConfigurationUpdateFailure struct {
	ProtocolIEs ProtocolIEContainerRANConfigurationUpdateFailureIEs
}

type ProtocolIEContainerRANConfigurationUpdateFailureIEs struct {
	List []RANConfigurationUpdateFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type RANConfigurationUpdateFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    RANConfigurationUpdateFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateFailureIEsPresentCause                  int = 0
	RANConfigurationUpdateFailureIEsPresentTimeToWait             int = 1
	RANConfigurationUpdateFailureIEsPresentCriticalityDiagnostics int = 2
)

type RANConfigurationUpdateFailureIEsTypeValue struct {
	Present                int
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait             *TimeToWait             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFConfigurationUpdate struct {
	ProtocolIEs ProtocolIEContainerAMFConfigurationUpdateIEs
}

type ProtocolIEContainerAMFConfigurationUpdateIEs struct {
	List []AMFConfigurationUpdateIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type AMFConfigurationUpdateIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	TypeValue    AMFConfigurationUpdateIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateIEsPresentAMFName                       int = 0
	AMFConfigurationUpdateIEsPresentServedGUAMIList               int = 1
	AMFConfigurationUpdateIEsPresentRelativeAMFCapacity           int = 2
	AMFConfigurationUpdateIEsPresentPLMNSupportList               int = 3
	AMFConfigurationUpdateIEsPresentAMFTNLAssociationToAddList    int = 4
	AMFConfigurationUpdateIEsPresentAMFTNLAssociationToRemoveList int = 5
	AMFConfigurationUpdateIEsPresentAMFTNLAssociationToUpdateList int = 6
)

type AMFConfigurationUpdateIEsTypeValue struct {
	Present                       int
	AMFName                       *AMFName                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFName"`
	ServedGUAMIList               *ServedGUAMIList               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDServedGUAMIList"`
	RelativeAMFCapacity           *RelativeAMFCapacity           `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRelativeAMFCapacity"`
	PLMNSupportList               *PLMNSupportList               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPLMNSupportList"`
	AMFTNLAssociationToAddList    *AMFTNLAssociationToAddList    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationToAddList"`
	AMFTNLAssociationToRemoveList *AMFTNLAssociationToRemoveList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationToRemoveList"`
	AMFTNLAssociationToUpdateList *AMFTNLAssociationToUpdateList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationToUpdateList"`
}

type AMFConfigurationUpdateAcknowledge struct {
	ProtocolIEs ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs
}

type ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs struct {
	List []AMFConfigurationUpdateAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type AMFConfigurationUpdateAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                   `vht:"Reference:ProtocolIEID"`
	TypeValue    AMFConfigurationUpdateAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationSetupList         int = 0
	AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationFailedToSetupList int = 1
	AMFConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics             int = 2
)

type AMFConfigurationUpdateAcknowledgeIEsTypeValue struct {
	Present                            int
	AMFTNLAssociationSetupList         *AMFTNLAssociationSetupList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationSetupList"`
	AMFTNLAssociationFailedToSetupList *TNLAssociationList         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationFailedToSetupList"`
	CriticalityDiagnostics             *CriticalityDiagnostics     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFConfigurationUpdateFailure struct {
	ProtocolIEs ProtocolIEContainerAMFConfigurationUpdateFailureIEs
}

type ProtocolIEContainerAMFConfigurationUpdateFailureIEs struct {
	List []AMFConfigurationUpdateFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type AMFConfigurationUpdateFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    AMFConfigurationUpdateFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateFailureIEsPresentCause                  int = 0
	AMFConfigurationUpdateFailureIEsPresentTimeToWait             int = 1
	AMFConfigurationUpdateFailureIEsPresentCriticalityDiagnostics int = 2
)

type AMFConfigurationUpdateFailureIEsTypeValue struct {
	Present                int
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait             *TimeToWait             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFStatusIndication struct {
	ProtocolIEs ProtocolIEContainerAMFStatusIndicationIEs
}

type ProtocolIEContainerAMFStatusIndicationIEs struct {
	List []AMFStatusIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type AMFStatusIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                     `vht:"Reference:ProtocolIEID"`
	TypeValue    AMFStatusIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFStatusIndicationIEsPresentUnavailableGUAMIList int = 0
)

type AMFStatusIndicationIEsTypeValue struct {
	Present              int
	UnavailableGUAMIList *UnavailableGUAMIList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUnavailableGUAMIList"`
}

type NGReset struct {
	ProtocolIEs ProtocolIEContainerNGResetIEs
}

type ProtocolIEContainerNGResetIEs struct {
	List []NGResetIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type NGResetIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality         `vht:"Reference:ProtocolIEID"`
	TypeValue    NGResetIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGResetIEsPresentCause     int = 0
	NGResetIEsPresentResetType int = 1
)

type NGResetIEsTypeValue struct {
	Present   int
	Cause     *Cause     `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	ResetType *ResetType `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDResetType"`
}

type NGResetAcknowledge struct {
	ProtocolIEs ProtocolIEContainerNGResetAcknowledgeIEs
}

type ProtocolIEContainerNGResetAcknowledgeIEs struct {
	List []NGResetAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type NGResetAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	TypeValue    NGResetAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGResetAcknowledgeIEsPresentUEAssociatedLogicalNGConnectionList int = 0
	NGResetAcknowledgeIEsPresentCriticalityDiagnostics              int = 1
)

type NGResetAcknowledgeIEsTypeValue struct {
	Present                             int
	UEAssociatedLogicalNGConnectionList *UEAssociatedLogicalNGConnectionList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAssociatedLogicalNGConnectionList"`
	CriticalityDiagnostics              *CriticalityDiagnostics              `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type ErrorIndication struct {
	ProtocolIEs ProtocolIEContainerErrorIndicationIEs
}

type ProtocolIEContainerErrorIndicationIEs struct {
	List []ErrorIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type ErrorIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	TypeValue    ErrorIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	ErrorIndicationIEsPresentAMFUENGAPID            int = 0
	ErrorIndicationIEsPresentRANUENGAPID            int = 1
	ErrorIndicationIEsPresentCause                  int = 2
	ErrorIndicationIEsPresentCriticalityDiagnostics int = 3
)

type ErrorIndicationIEsTypeValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID            *RANUENGAPID            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause                  *Cause                  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type OverloadStart struct {
	ProtocolIEs ProtocolIEContainerOverloadStartIEs
}

type ProtocolIEContainerOverloadStartIEs struct {
	List []OverloadStartIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type OverloadStartIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality               `vht:"Reference:ProtocolIEID"`
	TypeValue    OverloadStartIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	OverloadStartIEsPresentAMFOverloadResponse               int = 0
	OverloadStartIEsPresentAMFTrafficLoadReductionIndication int = 1
	OverloadStartIEsPresentOverloadStartNSSAIList            int = 2
)

type OverloadStartIEsTypeValue struct {
	Present                           int
	AMFOverloadResponse               *OverloadResponse               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFOverloadResponse"`
	AMFTrafficLoadReductionIndication *TrafficLoadReductionIndication `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTrafficLoadReductionIndication"`
	OverloadStartNSSAIList            *OverloadStartNSSAIList         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDOverloadStartNSSAIList"`
}

type OverloadStop struct {
	ProtocolIEs ProtocolIEContainerOverloadStopIEs
}

type ProtocolIEContainerOverloadStopIEs struct {
	List []OverloadStopIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type OverloadStopIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	TypeValue    OverloadStopIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	OverloadStopIEsPresentNothing int = 0
)

type OverloadStopIEsTypeValue struct {
	Present int
}

type UplinkRANConfigurationTransfer struct {
	ProtocolIEs ProtocolIEContainerUplinkRANConfigurationTransferIEs
}

type ProtocolIEContainerUplinkRANConfigurationTransferIEs struct {
	List []UplinkRANConfigurationTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkRANConfigurationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	TypeValue    UplinkRANConfigurationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRANConfigurationTransferIEsPresentSONConfigurationTransferUL     int = 0
	UplinkRANConfigurationTransferIEsPresentENDCSONConfigurationTransferUL int = 1
)

type UplinkRANConfigurationTransferIEsTypeValue struct {
	Present                        int
	SONConfigurationTransferUL     *SONConfigurationTransfer     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSONConfigurationTransferUL"`
	ENDCSONConfigurationTransferUL *ENDCSONConfigurationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDENDCSONConfigurationTransferUL"`
}

type DownlinkRANConfigurationTransfer struct {
	ProtocolIEs ProtocolIEContainerDownlinkRANConfigurationTransferIEs
}

type ProtocolIEContainerDownlinkRANConfigurationTransferIEs struct {
	List []DownlinkRANConfigurationTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkRANConfigurationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                  `vht:"Reference:ProtocolIEID"`
	TypeValue    DownlinkRANConfigurationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRANConfigurationTransferIEsPresentSONConfigurationTransferDL     int = 0
	DownlinkRANConfigurationTransferIEsPresentENDCSONConfigurationTransferDL int = 1
)

type DownlinkRANConfigurationTransferIEsTypeValue struct {
	Present                        int
	SONConfigurationTransferDL     *SONConfigurationTransfer     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSONConfigurationTransferDL"`
	ENDCSONConfigurationTransferDL *ENDCSONConfigurationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDENDCSONConfigurationTransferDL"`
}

type WriteReplaceWarningRequest struct {
	ProtocolIEs ProtocolIEContainerWriteReplaceWarningRequestIEs
}

type ProtocolIEContainerWriteReplaceWarningRequestIEs struct {
	List []WriteReplaceWarningRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type WriteReplaceWarningRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	TypeValue    WriteReplaceWarningRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	WriteReplaceWarningRequestIEsPresentMessageIdentifier           int = 0
	WriteReplaceWarningRequestIEsPresentSerialNumber                int = 1
	WriteReplaceWarningRequestIEsPresentWarningAreaList             int = 2
	WriteReplaceWarningRequestIEsPresentRepetitionPeriod            int = 3
	WriteReplaceWarningRequestIEsPresentNumberOfBroadcastsRequested int = 4
	WriteReplaceWarningRequestIEsPresentWarningType                 int = 5
	WriteReplaceWarningRequestIEsPresentWarningSecurityInfo         int = 6
	WriteReplaceWarningRequestIEsPresentDataCodingScheme            int = 7
	WriteReplaceWarningRequestIEsPresentWarningMessageContents      int = 8
	WriteReplaceWarningRequestIEsPresentConcurrentWarningMessageInd int = 9
	WriteReplaceWarningRequestIEsPresentWarningAreaCoordinates      int = 10
)

type WriteReplaceWarningRequestIEsTypeValue struct {
	Present                     int
	MessageIdentifier           *MessageIdentifier           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber                *SerialNumber                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	WarningAreaList             *WarningAreaList             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningAreaList"`
	RepetitionPeriod            *RepetitionPeriod            `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRepetitionPeriod"`
	NumberOfBroadcastsRequested *NumberOfBroadcastsRequested `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNumberOfBroadcastsRequested"`
	WarningType                 *WarningType                 `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningType"`
	WarningSecurityInfo         *WarningSecurityInfo         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningSecurityInfo"`
	DataCodingScheme            *DataCodingScheme            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDataCodingScheme"`
	WarningMessageContents      *WarningMessageContents      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningMessageContents"`
	ConcurrentWarningMessageInd *ConcurrentWarningMessageInd `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDConcurrentWarningMessageInd"`
	WarningAreaCoordinates      *WarningAreaCoordinates      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningAreaCoordinates"`
}

type WriteReplaceWarningResponse struct {
	ProtocolIEs ProtocolIEContainerWriteReplaceWarningResponseIEs
}

type ProtocolIEContainerWriteReplaceWarningResponseIEs struct {
	List []WriteReplaceWarningResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type WriteReplaceWarningResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	TypeValue    WriteReplaceWarningResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	WriteReplaceWarningResponseIEsPresentMessageIdentifier          int = 0
	WriteReplaceWarningResponseIEsPresentSerialNumber               int = 1
	WriteReplaceWarningResponseIEsPresentBroadcastCompletedAreaList int = 2
	WriteReplaceWarningResponseIEsPresentCriticalityDiagnostics     int = 3
)

type WriteReplaceWarningResponseIEsTypeValue struct {
	Present                    int
	MessageIdentifier          *MessageIdentifier          `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber               *SerialNumber               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	BroadcastCompletedAreaList *BroadcastCompletedAreaList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDBroadcastCompletedAreaList"`
	CriticalityDiagnostics     *CriticalityDiagnostics     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PWSCancelRequest struct {
	ProtocolIEs ProtocolIEContainerPWSCancelRequestIEs
}

type ProtocolIEContainerPWSCancelRequestIEs struct {
	List []PWSCancelRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PWSCancelRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	TypeValue    PWSCancelRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSCancelRequestIEsPresentMessageIdentifier        int = 0
	PWSCancelRequestIEsPresentSerialNumber             int = 1
	PWSCancelRequestIEsPresentWarningAreaList          int = 2
	PWSCancelRequestIEsPresentCancelAllWarningMessages int = 3
)

type PWSCancelRequestIEsTypeValue struct {
	Present                  int
	MessageIdentifier        *MessageIdentifier        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber             *SerialNumber             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	WarningAreaList          *WarningAreaList          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningAreaList"`
	CancelAllWarningMessages *CancelAllWarningMessages `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDCancelAllWarningMessages"`
}

type PWSCancelResponse struct {
	ProtocolIEs ProtocolIEContainerPWSCancelResponseIEs
}

type ProtocolIEContainerPWSCancelResponseIEs struct {
	List []PWSCancelResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PWSCancelResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                   `vht:"Reference:ProtocolIEID"`
	TypeValue    PWSCancelResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSCancelResponseIEsPresentMessageIdentifier          int = 0
	PWSCancelResponseIEsPresentSerialNumber               int = 1
	PWSCancelResponseIEsPresentBroadcastCancelledAreaList int = 2
	PWSCancelResponseIEsPresentCriticalityDiagnostics     int = 3
)

type PWSCancelResponseIEsTypeValue struct {
	Present                    int
	MessageIdentifier          *MessageIdentifier          `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber               *SerialNumber               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	BroadcastCancelledAreaList *BroadcastCancelledAreaList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDBroadcastCancelledAreaList"`
	CriticalityDiagnostics     *CriticalityDiagnostics     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PWSRestartIndication struct {
	ProtocolIEs ProtocolIEContainerPWSRestartIndicationIEs
}

type ProtocolIEContainerPWSRestartIndicationIEs struct {
	List []PWSRestartIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PWSRestartIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	TypeValue    PWSRestartIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSRestartIndicationIEsPresentCellIDListForRestart          int = 0
	PWSRestartIndicationIEsPresentGlobalRANNodeID               int = 1
	PWSRestartIndicationIEsPresentTAIListForRestart             int = 2
	PWSRestartIndicationIEsPresentEmergencyAreaIDListForRestart int = 3
)

type PWSRestartIndicationIEsTypeValue struct {
	Present                       int
	CellIDListForRestart          *CellIDListForRestart          `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDCellIDListForRestart"`
	GlobalRANNodeID               *GlobalRANNodeID               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	TAIListForRestart             *TAIListForRestart             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTAIListForRestart"`
	EmergencyAreaIDListForRestart *EmergencyAreaIDListForRestart `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDEmergencyAreaIDListForRestart"`
}

type PWSFailureIndication struct {
	ProtocolIEs ProtocolIEContainerPWSFailureIndicationIEs
}

type ProtocolIEContainerPWSFailureIndicationIEs struct {
	List []PWSFailureIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PWSFailureIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	TypeValue    PWSFailureIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSFailureIndicationIEsPresentPWSFailedCellIDList int = 0
	PWSFailureIndicationIEsPresentGlobalRANNodeID     int = 1
)

type PWSFailureIndicationIEsTypeValue struct {
	Present             int
	PWSFailedCellIDList *PWSFailedCellIDList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPWSFailedCellIDList"`
	GlobalRANNodeID     *GlobalRANNodeID     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
}

type DownlinkUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                    `vht:"Reference:ProtocolIEID"`
	TypeValue    DownlinkUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID int = 0
	DownlinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID int = 1
	DownlinkUEAssociatedNRPPaTransportIEsPresentRoutingID   int = 2
	DownlinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU    int = 3
)

type DownlinkUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RoutingID   *RoutingID   `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU    *NRPPaPDU    `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type UplinkUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs struct {
	List []UplinkUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                  `vht:"Reference:ProtocolIEID"`
	TypeValue    UplinkUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID int = 0
	UplinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID int = 1
	UplinkUEAssociatedNRPPaTransportIEsPresentRoutingID   int = 2
	UplinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU    int = 3
)

type UplinkUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RoutingID   *RoutingID   `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU    *NRPPaPDU    `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type DownlinkNonUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkNonUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkNonUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                       `vht:"Reference:ProtocolIEID"`
	TypeValue    DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID int = 0
	DownlinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU  int = 1
)

type DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present   int
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU  *NRPPaPDU  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type UplinkNonUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs struct {
	List []UplinkNonUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkNonUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                     `vht:"Reference:ProtocolIEID"`
	TypeValue    UplinkNonUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID int = 0
	UplinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU  int = 1
)

type UplinkNonUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present   int
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU  *NRPPaPDU  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type TraceStart struct {
	ProtocolIEs ProtocolIEContainerTraceStartIEs
}

type ProtocolIEContainerTraceStartIEs struct {
	List []TraceStartIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type TraceStartIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	TypeValue    TraceStartIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	TraceStartIEsPresentAMFUENGAPID     int = 0
	TraceStartIEsPresentRANUENGAPID     int = 1
	TraceStartIEsPresentTraceActivation int = 2
)

type TraceStartIEsTypeValue struct {
	Present         int
	AMFUENGAPID     *AMFUENGAPID     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID     *RANUENGAPID     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	TraceActivation *TraceActivation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceActivation"`
}

type TraceFailureIndication struct {
	ProtocolIEs ProtocolIEContainerTraceFailureIndicationIEs
}

type ProtocolIEContainerTraceFailureIndicationIEs struct {
	List []TraceFailureIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type TraceFailureIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	TypeValue    TraceFailureIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	TraceFailureIndicationIEsPresentAMFUENGAPID  int = 0
	TraceFailureIndicationIEsPresentRANUENGAPID  int = 1
	TraceFailureIndicationIEsPresentNGRANTraceID int = 2
	TraceFailureIndicationIEsPresentCause        int = 3
)

type TraceFailureIndicationIEsTypeValue struct {
	Present      int
	AMFUENGAPID  *AMFUENGAPID  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID  *RANUENGAPID  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NGRANTraceID *NGRANTraceID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANTraceID"`
	Cause        *Cause        `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type DeactivateTrace struct {
	ProtocolIEs ProtocolIEContainerDeactivateTraceIEs
}

type ProtocolIEContainerDeactivateTraceIEs struct {
	List []DeactivateTraceIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DeactivateTraceIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	TypeValue    DeactivateTraceIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DeactivateTraceIEsPresentAMFUENGAPID  int = 0
	DeactivateTraceIEsPresentRANUENGAPID  int = 1
	DeactivateTraceIEsPresentNGRANTraceID int = 2
)

type DeactivateTraceIEsTypeValue struct {
	Present      int
	AMFUENGAPID  *AMFUENGAPID  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID  *RANUENGAPID  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NGRANTraceID *NGRANTraceID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANTraceID"`
}

type CellTrafficTrace struct {
	ProtocolIEs ProtocolIEContainerCellTrafficTraceIEs
}

type ProtocolIEContainerCellTrafficTraceIEs struct {
	List []CellTrafficTraceIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type CellTrafficTraceIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	TypeValue    CellTrafficTraceIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	CellTrafficTraceIEsPresentAMFUENGAPID                    int = 0
	CellTrafficTraceIEsPresentRANUENGAPID                    int = 1
	CellTrafficTraceIEsPresentNGRANTraceID                   int = 2
	CellTrafficTraceIEsPresentNGRANCGI                       int = 3
	CellTrafficTraceIEsPresentTraceCollectionEntityIPAddress int = 4
)

type CellTrafficTraceIEsTypeValue struct {
	Present                        int
	AMFUENGAPID                    *AMFUENGAPID           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                    *RANUENGAPID           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NGRANTraceID                   *NGRANTraceID          `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANTraceID"`
	NGRANCGI                       *NGRANCGI              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANCGI"`
	TraceCollectionEntityIPAddress *TransportLayerAddress `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceCollectionEntityIPAddress"`
}

type LocationReportingControl struct {
	ProtocolIEs ProtocolIEContainerLocationReportingControlIEs
}

type ProtocolIEContainerLocationReportingControlIEs struct {
	List []LocationReportingControlIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type LocationReportingControlIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	TypeValue    LocationReportingControlIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportingControlIEsPresentAMFUENGAPID                  int = 0
	LocationReportingControlIEsPresentRANUENGAPID                  int = 1
	LocationReportingControlIEsPresentLocationReportingRequestType int = 2
)

type LocationReportingControlIEsTypeValue struct {
	Present                      int
	AMFUENGAPID                  *AMFUENGAPID                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                  *RANUENGAPID                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	LocationReportingRequestType *LocationReportingRequestType `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
}

type LocationReportingFailureIndication struct {
	ProtocolIEs ProtocolIEContainerLocationReportingFailureIndicationIEs
}

type ProtocolIEContainerLocationReportingFailureIndicationIEs struct {
	List []LocationReportingFailureIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type LocationReportingFailureIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                    `vht:"Reference:ProtocolIEID"`
	TypeValue    LocationReportingFailureIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportingFailureIndicationIEsPresentAMFUENGAPID int = 0
	LocationReportingFailureIndicationIEsPresentRANUENGAPID int = 1
	LocationReportingFailureIndicationIEsPresentCause       int = 2
)

type LocationReportingFailureIndicationIEsTypeValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause       *Cause       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type LocationReport struct {
	ProtocolIEs ProtocolIEContainerLocationReportIEs
}

type ProtocolIEContainerLocationReportIEs struct {
	List []LocationReportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type LocationReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    LocationReportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportIEsPresentAMFUENGAPID                    int = 0
	LocationReportIEsPresentRANUENGAPID                    int = 1
	LocationReportIEsPresentUserLocationInformation        int = 2
	LocationReportIEsPresentUEPresenceInAreaOfInterestList int = 3
	LocationReportIEsPresentLocationReportingRequestType   int = 4
)

type LocationReportIEsTypeValue struct {
	Present                        int
	AMFUENGAPID                    *AMFUENGAPID                    `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                    *RANUENGAPID                    `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UserLocationInformation        *UserLocationInformation        `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	UEPresenceInAreaOfInterestList *UEPresenceInAreaOfInterestList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEPresenceInAreaOfInterestList"`
	LocationReportingRequestType   *LocationReportingRequestType   `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
}

type UETNLABindingReleaseRequest struct {
	ProtocolIEs ProtocolIEContainerUETNLABindingReleaseRequestIEs
}

type ProtocolIEContainerUETNLABindingReleaseRequestIEs struct {
	List []UETNLABindingReleaseRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UETNLABindingReleaseRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	TypeValue    UETNLABindingReleaseRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UETNLABindingReleaseRequestIEsPresentAMFUENGAPID int = 0
	UETNLABindingReleaseRequestIEsPresentRANUENGAPID int = 1
)

type UETNLABindingReleaseRequestIEsTypeValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
}

type UERadioCapabilityInfoIndication struct {
	ProtocolIEs ProtocolIEContainerUERadioCapabilityInfoIndicationIEs
}

type ProtocolIEContainerUERadioCapabilityInfoIndicationIEs struct {
	List []UERadioCapabilityInfoIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UERadioCapabilityInfoIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                 `vht:"Reference:ProtocolIEID"`
	TypeValue    UERadioCapabilityInfoIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityInfoIndicationIEsPresentAMFUENGAPID                int = 0
	UERadioCapabilityInfoIndicationIEsPresentRANUENGAPID                int = 1
	UERadioCapabilityInfoIndicationIEsPresentUERadioCapability          int = 2
	UERadioCapabilityInfoIndicationIEsPresentUERadioCapabilityForPaging int = 3
)

type UERadioCapabilityInfoIndicationIEsTypeValue struct {
	Present                    int
	AMFUENGAPID                *AMFUENGAPID                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                *RANUENGAPID                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UERadioCapability          *UERadioCapability          `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapability"`
	UERadioCapabilityForPaging *UERadioCapabilityForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapabilityForPaging"`
}

type UERadioCapabilityCheckRequest struct {
	ProtocolIEs ProtocolIEContainerUERadioCapabilityCheckRequestIEs
}

type ProtocolIEContainerUERadioCapabilityCheckRequestIEs struct {
	List []UERadioCapabilityCheckRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UERadioCapabilityCheckRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    UERadioCapabilityCheckRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityCheckRequestIEsPresentAMFUENGAPID       int = 0
	UERadioCapabilityCheckRequestIEsPresentRANUENGAPID       int = 1
	UERadioCapabilityCheckRequestIEsPresentUERadioCapability int = 2
)

type UERadioCapabilityCheckRequestIEsTypeValue struct {
	Present           int
	AMFUENGAPID       *AMFUENGAPID       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID       *RANUENGAPID       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UERadioCapability *UERadioCapability `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapability"`
}

type UERadioCapabilityCheckResponse struct {
	ProtocolIEs ProtocolIEContainerUERadioCapabilityCheckResponseIEs
}

type ProtocolIEContainerUERadioCapabilityCheckResponseIEs struct {
	List []UERadioCapabilityCheckResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UERadioCapabilityCheckResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	TypeValue    UERadioCapabilityCheckResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityCheckResponseIEsPresentAMFUENGAPID              int = 0
	UERadioCapabilityCheckResponseIEsPresentRANUENGAPID              int = 1
	UERadioCapabilityCheckResponseIEsPresentIMSVoiceSupportIndicator int = 2
	UERadioCapabilityCheckResponseIEsPresentCriticalityDiagnostics   int = 3
)

type UERadioCapabilityCheckResponseIEsTypeValue struct {
	Present                  int
	AMFUENGAPID              *AMFUENGAPID              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID              *RANUENGAPID              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	IMSVoiceSupportIndicator *IMSVoiceSupportIndicator `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDIMSVoiceSupportIndicator"`
	CriticalityDiagnostics   *CriticalityDiagnostics   `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PrivateMessage struct {
	PrivateIEs PrivateIEContainerPrivateMessageIEs
}

type PrivateIEContainerPrivateMessageIEs struct {
	List []PrivateMessageIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type PrivateMessageIEs struct {
	PrivateIEID PrivateIEID
	Criticality Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue   PrivateMessageIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PrivateMessageIEsPresentNothing int = 0
)

type PrivateMessageIEsTypeValue struct {
	Present int
}

type SecondaryRATDataUsageReport struct {
	ProtocolIEs ProtocolIEContainerSecondaryRATDataUsageReportIEs
}

type ProtocolIEContainerSecondaryRATDataUsageReportIEs struct {
	List []SecondaryRATDataUsageReportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type SecondaryRATDataUsageReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	TypeValue    SecondaryRATDataUsageReportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	SecondaryRATDataUsageReportIEsPresentAMFUENGAPID                             int = 0
	SecondaryRATDataUsageReportIEsPresentRANUENGAPID                             int = 1
	SecondaryRATDataUsageReportIEsPresentPDUSessionResourceSecondaryRATUsageList int = 2
	SecondaryRATDataUsageReportIEsPresentHandoverFlag                            int = 3
	SecondaryRATDataUsageReportIEsPresentUserLocationInformation                 int = 4
)

type SecondaryRATDataUsageReportIEsTypeValue struct {
	Present                                 int
	AMFUENGAPID                             *AMFUENGAPID                             `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                             *RANUENGAPID                             `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceSecondaryRATUsageList *PDUSessionResourceSecondaryRATUsageList `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSecondaryRATUsageList"`
	HandoverFlag                            *HandoverFlag                            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDHandoverFlag"`
	UserLocationInformation                 *UserLocationInformation                 `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type UplinkRIMInformationTransfer struct {
	ProtocolIEs ProtocolIEContainerUplinkRIMInformationTransferIEs
}

type ProtocolIEContainerUplinkRIMInformationTransferIEs struct {
	List []UplinkRIMInformationTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkRIMInformationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	TypeValue    UplinkRIMInformationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRIMInformationTransferIEsPresentRIMInformationTransfer int = 0
)

type UplinkRIMInformationTransferIEsTypeValue struct {
	Present                int
	RIMInformationTransfer *RIMInformationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRIMInformationTransfer"`
}

type DownlinkRIMInformationTransfer struct {
	ProtocolIEs ProtocolIEContainerDownlinkRIMInformationTransferIEs
}

type ProtocolIEContainerDownlinkRIMInformationTransferIEs struct {
	List []DownlinkRIMInformationTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkRIMInformationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	TypeValue    DownlinkRIMInformationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRIMInformationTransferIEsPresentRIMInformationTransfer int = 0
)

type DownlinkRIMInformationTransferIEsTypeValue struct {
	Present                int
	RIMInformationTransfer *RIMInformationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRIMInformationTransfer"`
}
