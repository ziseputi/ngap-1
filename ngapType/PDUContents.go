package ngapType

type PDUSessionResourceSetupRequest struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceSetupRequestIEs
}

type ProtocolIEContainerPDUSessionResourceSetupRequestIEs struct {
	List []PDUSessionResourceSetupRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceSetupRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupRequestIEsValuePresentAMFUENGAPID                      int = 0
	PDUSessionResourceSetupRequestIEsValuePresentRANUENGAPID                      int = 1
	PDUSessionResourceSetupRequestIEsValuePresentRANPagingPriority                int = 2
	PDUSessionResourceSetupRequestIEsValuePresentNASPDU                           int = 3
	PDUSessionResourceSetupRequestIEsValuePresentPDUSessionResourceSetupListSUReq int = 4
	PDUSessionResourceSetupRequestIEsValuePresentUEAggregateMaximumBitRate        int = 5
)

type PDUSessionResourceSetupRequestIEsValue struct {
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
	List []PDUSessionResourceSetupResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceSetupResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupResponseIEsValuePresentAMFUENGAPID                              int = 0
	PDUSessionResourceSetupResponseIEsValuePresentRANUENGAPID                              int = 1
	PDUSessionResourceSetupResponseIEsValuePresentPDUSessionResourceSetupListSURes         int = 2
	PDUSessionResourceSetupResponseIEsValuePresentPDUSessionResourceFailedToSetupListSURes int = 3
	PDUSessionResourceSetupResponseIEsValuePresentCriticalityDiagnostics                   int = 4
)

type PDUSessionResourceSetupResponseIEsValue struct {
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
	List []PDUSessionResourceReleaseCommandIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceReleaseCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceReleaseCommandIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceReleaseCommandIEsValuePresentAMFUENGAPID                           int = 0
	PDUSessionResourceReleaseCommandIEsValuePresentRANUENGAPID                           int = 1
	PDUSessionResourceReleaseCommandIEsValuePresentRANPagingPriority                     int = 2
	PDUSessionResourceReleaseCommandIEsValuePresentNASPDU                                int = 3
	PDUSessionResourceReleaseCommandIEsValuePresentPDUSessionResourceToReleaseListRelCmd int = 4
)

type PDUSessionResourceReleaseCommandIEsValue struct {
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
	List []PDUSessionResourceReleaseResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceReleaseResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceReleaseResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceReleaseResponseIEsValuePresentAMFUENGAPID                          int = 0
	PDUSessionResourceReleaseResponseIEsValuePresentRANUENGAPID                          int = 1
	PDUSessionResourceReleaseResponseIEsValuePresentPDUSessionResourceReleasedListRelRes int = 2
	PDUSessionResourceReleaseResponseIEsValuePresentUserLocationInformation              int = 3
	PDUSessionResourceReleaseResponseIEsValuePresentCriticalityDiagnostics               int = 4
)

type PDUSessionResourceReleaseResponseIEsValue struct {
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
	List []PDUSessionResourceModifyRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceModifyRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceModifyRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyRequestIEsValuePresentAMFUENGAPID                        int = 0
	PDUSessionResourceModifyRequestIEsValuePresentRANUENGAPID                        int = 1
	PDUSessionResourceModifyRequestIEsValuePresentRANPagingPriority                  int = 2
	PDUSessionResourceModifyRequestIEsValuePresentPDUSessionResourceModifyListModReq int = 3
)

type PDUSessionResourceModifyRequestIEsValue struct {
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
	List []PDUSessionResourceModifyResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceModifyResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceModifyResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyResponseIEsValuePresentAMFUENGAPID                                int = 0
	PDUSessionResourceModifyResponseIEsValuePresentRANUENGAPID                                int = 1
	PDUSessionResourceModifyResponseIEsValuePresentPDUSessionResourceModifyListModRes         int = 2
	PDUSessionResourceModifyResponseIEsValuePresentPDUSessionResourceFailedToModifyListModRes int = 3
	PDUSessionResourceModifyResponseIEsValuePresentUserLocationInformation                    int = 4
	PDUSessionResourceModifyResponseIEsValuePresentCriticalityDiagnostics                     int = 5
)

type PDUSessionResourceModifyResponseIEsValue struct {
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
	List []PDUSessionResourceNotifyIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceNotifyIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceNotifyIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceNotifyIEsValuePresentAMFUENGAPID                       int = 0
	PDUSessionResourceNotifyIEsValuePresentRANUENGAPID                       int = 1
	PDUSessionResourceNotifyIEsValuePresentPDUSessionResourceNotifyList      int = 2
	PDUSessionResourceNotifyIEsValuePresentPDUSessionResourceReleasedListNot int = 3
	PDUSessionResourceNotifyIEsValuePresentUserLocationInformation           int = 4
)

type PDUSessionResourceNotifyIEsValue struct {
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
	List []PDUSessionResourceModifyIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceModifyIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceModifyIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyIndicationIEsValuePresentAMFUENGAPID                        int = 0
	PDUSessionResourceModifyIndicationIEsValuePresentRANUENGAPID                        int = 1
	PDUSessionResourceModifyIndicationIEsValuePresentPDUSessionResourceModifyListModInd int = 2
	PDUSessionResourceModifyIndicationIEsValuePresentUserLocationInformation            int = 3
)

type PDUSessionResourceModifyIndicationIEsValue struct {
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
	List []PDUSessionResourceModifyConfirmIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PDUSessionResourceModifyConfirmIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	Value        PDUSessionResourceModifyConfirmIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyConfirmIEsValuePresentAMFUENGAPID                                int = 0
	PDUSessionResourceModifyConfirmIEsValuePresentRANUENGAPID                                int = 1
	PDUSessionResourceModifyConfirmIEsValuePresentPDUSessionResourceModifyListModCfm         int = 2
	PDUSessionResourceModifyConfirmIEsValuePresentPDUSessionResourceFailedToModifyListModCfm int = 3
	PDUSessionResourceModifyConfirmIEsValuePresentCriticalityDiagnostics                     int = 4
)

type PDUSessionResourceModifyConfirmIEsValue struct {
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
	List []InitialContextSetupRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type InitialContextSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	Value        InitialContextSetupRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupRequestIEsValuePresentAMFUENGAPID                                 int = 0
	InitialContextSetupRequestIEsValuePresentRANUENGAPID                                 int = 1
	InitialContextSetupRequestIEsValuePresentOldAMF                                      int = 2
	InitialContextSetupRequestIEsValuePresentUEAggregateMaximumBitRate                   int = 3
	InitialContextSetupRequestIEsValuePresentCoreNetworkAssistanceInformationForInactive int = 4
	InitialContextSetupRequestIEsValuePresentGUAMI                                       int = 5
	InitialContextSetupRequestIEsValuePresentPDUSessionResourceSetupListCxtReq           int = 6
	InitialContextSetupRequestIEsValuePresentAllowedNSSAI                                int = 7
	InitialContextSetupRequestIEsValuePresentUESecurityCapabilities                      int = 8
	InitialContextSetupRequestIEsValuePresentSecurityKey                                 int = 9
	InitialContextSetupRequestIEsValuePresentTraceActivation                             int = 10
	InitialContextSetupRequestIEsValuePresentMobilityRestrictionList                     int = 11
	InitialContextSetupRequestIEsValuePresentUERadioCapability                           int = 12
	InitialContextSetupRequestIEsValuePresentIndexToRFSP                                 int = 13
	InitialContextSetupRequestIEsValuePresentMaskedIMEISV                                int = 14
	InitialContextSetupRequestIEsValuePresentNASPDU                                      int = 15
	InitialContextSetupRequestIEsValuePresentEmergencyFallbackIndicator                  int = 16
	InitialContextSetupRequestIEsValuePresentRRCInactiveTransitionReportRequest          int = 17
	InitialContextSetupRequestIEsValuePresentUERadioCapabilityForPaging                  int = 18
	InitialContextSetupRequestIEsValuePresentRedirectionVoiceFallback                    int = 19
	InitialContextSetupRequestIEsValuePresentLocationReportingRequestType                int = 20
	InitialContextSetupRequestIEsValuePresentCNAssistedRANTuning                         int = 21
	InitialContextSetupRequestIEsValuePresentSRVCCOperationPossible                      int = 22
)

type InitialContextSetupRequestIEsValue struct {
	Present                                     int
	AMFUENGAPID                                 *AMFUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                                 *RANUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	OldAMF                                      *AMFName                                     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFName"`
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
	List []InitialContextSetupResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type InitialContextSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	Value        InitialContextSetupResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupResponseIEsValuePresentAMFUENGAPID                               int = 0
	InitialContextSetupResponseIEsValuePresentRANUENGAPID                               int = 1
	InitialContextSetupResponseIEsValuePresentPDUSessionResourceSetupListCxtRes         int = 2
	InitialContextSetupResponseIEsValuePresentPDUSessionResourceFailedToSetupListCxtRes int = 3
	InitialContextSetupResponseIEsValuePresentCriticalityDiagnostics                    int = 4
)

type InitialContextSetupResponseIEsValue struct {
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
	List []InitialContextSetupFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type InitialContextSetupFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	Value        InitialContextSetupFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupFailureIEsValuePresentAMFUENGAPID                                int = 0
	InitialContextSetupFailureIEsValuePresentRANUENGAPID                                int = 1
	InitialContextSetupFailureIEsValuePresentPDUSessionResourceFailedToSetupListCxtFail int = 2
	InitialContextSetupFailureIEsValuePresentCause                                      int = 3
	InitialContextSetupFailureIEsValuePresentCriticalityDiagnostics                     int = 4
)

type InitialContextSetupFailureIEsValue struct {
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
	List []UEContextReleaseRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UEContextReleaseRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                     `vht:"Reference:ProtocolIEID"`
	Value        UEContextReleaseRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseRequestIEsValuePresentAMFUENGAPID                     int = 0
	UEContextReleaseRequestIEsValuePresentRANUENGAPID                     int = 1
	UEContextReleaseRequestIEsValuePresentPDUSessionResourceListCxtRelReq int = 2
	UEContextReleaseRequestIEsValuePresentCause                           int = 3
)

type UEContextReleaseRequestIEsValue struct {
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
	List []UEContextReleaseCommandIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UEContextReleaseCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                     `vht:"Reference:ProtocolIEID"`
	Value        UEContextReleaseCommandIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseCommandIEsValuePresentUENGAPIDs int = 0
	UEContextReleaseCommandIEsValuePresentCause     int = 1
)

type UEContextReleaseCommandIEsValue struct {
	Present   int
	UENGAPIDs *UENGAPIDs `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUENGAPIDs"`
	Cause     *Cause     `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type UEContextReleaseComplete struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseCompleteIEs
}

type ProtocolIEContainerUEContextReleaseCompleteIEs struct {
	List []UEContextReleaseCompleteIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UEContextReleaseCompleteIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	Value        UEContextReleaseCompleteIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseCompleteIEsValuePresentAMFUENGAPID                                int = 0
	UEContextReleaseCompleteIEsValuePresentRANUENGAPID                                int = 1
	UEContextReleaseCompleteIEsValuePresentUserLocationInformation                    int = 2
	UEContextReleaseCompleteIEsValuePresentInfoOnRecommendedCellsAndRANNodesForPaging int = 3
	UEContextReleaseCompleteIEsValuePresentPDUSessionResourceListCxtRelCpl            int = 4
	UEContextReleaseCompleteIEsValuePresentCriticalityDiagnostics                     int = 5
)

type UEContextReleaseCompleteIEsValue struct {
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
	List []UEContextModificationRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UEContextModificationRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	Value        UEContextModificationRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationRequestIEsValuePresentAMFUENGAPID                                 int = 0
	UEContextModificationRequestIEsValuePresentRANUENGAPID                                 int = 1
	UEContextModificationRequestIEsValuePresentRANPagingPriority                           int = 2
	UEContextModificationRequestIEsValuePresentSecurityKey                                 int = 3
	UEContextModificationRequestIEsValuePresentIndexToRFSP                                 int = 4
	UEContextModificationRequestIEsValuePresentUEAggregateMaximumBitRate                   int = 5
	UEContextModificationRequestIEsValuePresentUESecurityCapabilities                      int = 6
	UEContextModificationRequestIEsValuePresentCoreNetworkAssistanceInformationForInactive int = 7
	UEContextModificationRequestIEsValuePresentEmergencyFallbackIndicator                  int = 8
	UEContextModificationRequestIEsValuePresentNewAMFUENGAPID                              int = 9
	UEContextModificationRequestIEsValuePresentRRCInactiveTransitionReportRequest          int = 10
	UEContextModificationRequestIEsValuePresentNewGUAMI                                    int = 11
	UEContextModificationRequestIEsValuePresentCNAssistedRANTuning                         int = 12
	UEContextModificationRequestIEsValuePresentSRVCCOperationPossible                      int = 13
)

type UEContextModificationRequestIEsValue struct {
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
	NewAMFUENGAPID                              *AMFUENGAPID                                 `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	NewGUAMI                                    *GUAMI                                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGUAMI"`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible                      *SRVCCOperationPossible                      `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type UEContextModificationResponse struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationResponseIEs
}

type ProtocolIEContainerUEContextModificationResponseIEs struct {
	List []UEContextModificationResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UEContextModificationResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                           `vht:"Reference:ProtocolIEID"`
	Value        UEContextModificationResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationResponseIEsValuePresentAMFUENGAPID             int = 0
	UEContextModificationResponseIEsValuePresentRANUENGAPID             int = 1
	UEContextModificationResponseIEsValuePresentRRCState                int = 2
	UEContextModificationResponseIEsValuePresentUserLocationInformation int = 3
	UEContextModificationResponseIEsValuePresentCriticalityDiagnostics  int = 4
)

type UEContextModificationResponseIEsValue struct {
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
	List []UEContextModificationFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UEContextModificationFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	Value        UEContextModificationFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationFailureIEsValuePresentAMFUENGAPID            int = 0
	UEContextModificationFailureIEsValuePresentRANUENGAPID            int = 1
	UEContextModificationFailureIEsValuePresentCause                  int = 2
	UEContextModificationFailureIEsValuePresentCriticalityDiagnostics int = 3
)

type UEContextModificationFailureIEsValue struct {
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
	List []RRCInactiveTransitionReportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type RRCInactiveTransitionReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	Value        RRCInactiveTransitionReportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	RRCInactiveTransitionReportIEsValuePresentAMFUENGAPID             int = 0
	RRCInactiveTransitionReportIEsValuePresentRANUENGAPID             int = 1
	RRCInactiveTransitionReportIEsValuePresentRRCState                int = 2
	RRCInactiveTransitionReportIEsValuePresentUserLocationInformation int = 3
)

type RRCInactiveTransitionReportIEsValue struct {
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
	List []HandoverRequiredIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverRequiredIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	Value        HandoverRequiredIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequiredIEsValuePresentAMFUENGAPID                        int = 0
	HandoverRequiredIEsValuePresentRANUENGAPID                        int = 1
	HandoverRequiredIEsValuePresentHandoverType                       int = 2
	HandoverRequiredIEsValuePresentCause                              int = 3
	HandoverRequiredIEsValuePresentTargetID                           int = 4
	HandoverRequiredIEsValuePresentDirectForwardingPathAvailability   int = 5
	HandoverRequiredIEsValuePresentPDUSessionResourceListHORqd        int = 6
	HandoverRequiredIEsValuePresentSourceToTargetTransparentContainer int = 7
)

type HandoverRequiredIEsValue struct {
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
	List []HandoverCommandIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	Value        HandoverCommandIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCommandIEsValuePresentAMFUENGAPID                          int = 0
	HandoverCommandIEsValuePresentRANUENGAPID                          int = 1
	HandoverCommandIEsValuePresentHandoverType                         int = 2
	HandoverCommandIEsValuePresentNASSecurityParametersFromNGRAN       int = 3
	HandoverCommandIEsValuePresentPDUSessionResourceHandoverList       int = 4
	HandoverCommandIEsValuePresentPDUSessionResourceToReleaseListHOCmd int = 5
	HandoverCommandIEsValuePresentTargetToSourceTransparentContainer   int = 6
	HandoverCommandIEsValuePresentCriticalityDiagnostics               int = 7
)

type HandoverCommandIEsValue struct {
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
	List []HandoverPreparationFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverPreparationFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	Value        HandoverPreparationFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverPreparationFailureIEsValuePresentAMFUENGAPID            int = 0
	HandoverPreparationFailureIEsValuePresentRANUENGAPID            int = 1
	HandoverPreparationFailureIEsValuePresentCause                  int = 2
	HandoverPreparationFailureIEsValuePresentCriticalityDiagnostics int = 3
)

type HandoverPreparationFailureIEsValue struct {
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
	List []HandoverRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	Value        HandoverRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequestIEsValuePresentAMFUENGAPID                                 int = 0
	HandoverRequestIEsValuePresentHandoverType                                int = 1
	HandoverRequestIEsValuePresentCause                                       int = 2
	HandoverRequestIEsValuePresentUEAggregateMaximumBitRate                   int = 3
	HandoverRequestIEsValuePresentCoreNetworkAssistanceInformationForInactive int = 4
	HandoverRequestIEsValuePresentUESecurityCapabilities                      int = 5
	HandoverRequestIEsValuePresentSecurityContext                             int = 6
	HandoverRequestIEsValuePresentNewSecurityContextInd                       int = 7
	HandoverRequestIEsValuePresentNASC                                        int = 8
	HandoverRequestIEsValuePresentPDUSessionResourceSetupListHOReq            int = 9
	HandoverRequestIEsValuePresentAllowedNSSAI                                int = 10
	HandoverRequestIEsValuePresentTraceActivation                             int = 11
	HandoverRequestIEsValuePresentMaskedIMEISV                                int = 12
	HandoverRequestIEsValuePresentSourceToTargetTransparentContainer          int = 13
	HandoverRequestIEsValuePresentMobilityRestrictionList                     int = 14
	HandoverRequestIEsValuePresentLocationReportingRequestType                int = 15
	HandoverRequestIEsValuePresentRRCInactiveTransitionReportRequest          int = 16
	HandoverRequestIEsValuePresentGUAMI                                       int = 17
	HandoverRequestIEsValuePresentRedirectionVoiceFallback                    int = 18
	HandoverRequestIEsValuePresentCNAssistedRANTuning                         int = 19
	HandoverRequestIEsValuePresentSRVCCOperationPossible                      int = 20
)

type HandoverRequestIEsValue struct {
	Present                                     int
	AMFUENGAPID                                 *AMFUENGAPID                                 `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	HandoverType                                *HandoverType                                `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	Cause                                       *Cause                                       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	UESecurityCapabilities                      *UESecurityCapabilities                      `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityContext                             *SecurityContext                             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityContext"`
	NewSecurityContextInd                       *NewSecurityContextInd                       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewSecurityContextInd"`
	NASC                                        *NASPDU                                      `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
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
	List []HandoverRequestAcknowledgeIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverRequestAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	Value        HandoverRequestAcknowledgeIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequestAcknowledgeIEsValuePresentAMFUENGAPID                              int = 0
	HandoverRequestAcknowledgeIEsValuePresentRANUENGAPID                              int = 1
	HandoverRequestAcknowledgeIEsValuePresentPDUSessionResourceAdmittedList           int = 2
	HandoverRequestAcknowledgeIEsValuePresentPDUSessionResourceFailedToSetupListHOAck int = 3
	HandoverRequestAcknowledgeIEsValuePresentTargetToSourceTransparentContainer       int = 4
	HandoverRequestAcknowledgeIEsValuePresentCriticalityDiagnostics                   int = 5
)

type HandoverRequestAcknowledgeIEsValue struct {
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
	List []HandoverFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	Value        HandoverFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverFailureIEsValuePresentAMFUENGAPID            int = 0
	HandoverFailureIEsValuePresentCause                  int = 1
	HandoverFailureIEsValuePresentCriticalityDiagnostics int = 2
)

type HandoverFailureIEsValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverNotify struct {
	ProtocolIEs ProtocolIEContainerHandoverNotifyIEs
}

type ProtocolIEContainerHandoverNotifyIEs struct {
	List []HandoverNotifyIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverNotifyIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	Value        HandoverNotifyIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverNotifyIEsValuePresentAMFUENGAPID             int = 0
	HandoverNotifyIEsValuePresentRANUENGAPID             int = 1
	HandoverNotifyIEsValuePresentUserLocationInformation int = 2
)

type HandoverNotifyIEsValue struct {
	Present                 int
	AMFUENGAPID             *AMFUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID             *RANUENGAPID             `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type PathSwitchRequest struct {
	ProtocolIEs ProtocolIEContainerPathSwitchRequestIEs
}

type ProtocolIEContainerPathSwitchRequestIEs struct {
	List []PathSwitchRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PathSwitchRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality               `vht:"Reference:ProtocolIEID"`
	Value        PathSwitchRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestIEsValuePresentRANUENGAPID                              int = 0
	PathSwitchRequestIEsValuePresentSourceAMFUENGAPID                        int = 1
	PathSwitchRequestIEsValuePresentUserLocationInformation                  int = 2
	PathSwitchRequestIEsValuePresentUESecurityCapabilities                   int = 3
	PathSwitchRequestIEsValuePresentPDUSessionResourceToBeSwitchedDLList     int = 4
	PathSwitchRequestIEsValuePresentPDUSessionResourceFailedToSetupListPSReq int = 5
)

type PathSwitchRequestIEsValue struct {
	Present                                  int
	RANUENGAPID                              *RANUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	SourceAMFUENGAPID                        *AMFUENGAPID                              `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	UserLocationInformation                  *UserLocationInformation                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	UESecurityCapabilities                   *UESecurityCapabilities                   `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	PDUSessionResourceToBeSwitchedDLList     *PDUSessionResourceToBeSwitchedDLList     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceToBeSwitchedDLList"`
	PDUSessionResourceFailedToSetupListPSReq *PDUSessionResourceFailedToSetupListPSReq `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq"`
}

type PathSwitchRequestAcknowledge struct {
	ProtocolIEs ProtocolIEContainerPathSwitchRequestAcknowledgeIEs
}

type ProtocolIEContainerPathSwitchRequestAcknowledgeIEs struct {
	List []PathSwitchRequestAcknowledgeIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PathSwitchRequestAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	Value        PathSwitchRequestAcknowledgeIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestAcknowledgeIEsValuePresentAMFUENGAPID                                 int = 0
	PathSwitchRequestAcknowledgeIEsValuePresentRANUENGAPID                                 int = 1
	PathSwitchRequestAcknowledgeIEsValuePresentUESecurityCapabilities                      int = 2
	PathSwitchRequestAcknowledgeIEsValuePresentSecurityContext                             int = 3
	PathSwitchRequestAcknowledgeIEsValuePresentNewSecurityContextInd                       int = 4
	PathSwitchRequestAcknowledgeIEsValuePresentPDUSessionResourceSwitchedList              int = 5
	PathSwitchRequestAcknowledgeIEsValuePresentPDUSessionResourceReleasedListPSAck         int = 6
	PathSwitchRequestAcknowledgeIEsValuePresentAllowedNSSAI                                int = 7
	PathSwitchRequestAcknowledgeIEsValuePresentCoreNetworkAssistanceInformationForInactive int = 8
	PathSwitchRequestAcknowledgeIEsValuePresentRRCInactiveTransitionReportRequest          int = 9
	PathSwitchRequestAcknowledgeIEsValuePresentCriticalityDiagnostics                      int = 10
	PathSwitchRequestAcknowledgeIEsValuePresentRedirectionVoiceFallback                    int = 11
	PathSwitchRequestAcknowledgeIEsValuePresentCNAssistedRANTuning                         int = 12
	PathSwitchRequestAcknowledgeIEsValuePresentSRVCCOperationPossible                      int = 13
)

type PathSwitchRequestAcknowledgeIEsValue struct {
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
	List []PathSwitchRequestFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PathSwitchRequestFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	Value        PathSwitchRequestFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestFailureIEsValuePresentAMFUENGAPID                          int = 0
	PathSwitchRequestFailureIEsValuePresentRANUENGAPID                          int = 1
	PathSwitchRequestFailureIEsValuePresentPDUSessionResourceReleasedListPSFail int = 2
	PathSwitchRequestFailureIEsValuePresentCriticalityDiagnostics               int = 3
)

type PathSwitchRequestFailureIEsValue struct {
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
	List []HandoverCancelIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverCancelIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	Value        HandoverCancelIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCancelIEsValuePresentAMFUENGAPID int = 0
	HandoverCancelIEsValuePresentRANUENGAPID int = 1
	HandoverCancelIEsValuePresentCause       int = 2
)

type HandoverCancelIEsValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause       *Cause       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type HandoverCancelAcknowledge struct {
	ProtocolIEs ProtocolIEContainerHandoverCancelAcknowledgeIEs
}

type ProtocolIEContainerHandoverCancelAcknowledgeIEs struct {
	List []HandoverCancelAcknowledgeIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type HandoverCancelAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                       `vht:"Reference:ProtocolIEID"`
	Value        HandoverCancelAcknowledgeIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCancelAcknowledgeIEsValuePresentAMFUENGAPID            int = 0
	HandoverCancelAcknowledgeIEsValuePresentRANUENGAPID            int = 1
	HandoverCancelAcknowledgeIEsValuePresentCriticalityDiagnostics int = 2
)

type HandoverCancelAcknowledgeIEsValue struct {
	Present                int
	AMFUENGAPID            *AMFUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID            *RANUENGAPID            `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UplinkRANStatusTransfer struct {
	ProtocolIEs ProtocolIEContainerUplinkRANStatusTransferIEs
}

type ProtocolIEContainerUplinkRANStatusTransferIEs struct {
	List []UplinkRANStatusTransferIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UplinkRANStatusTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                     `vht:"Reference:ProtocolIEID"`
	Value        UplinkRANStatusTransferIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRANStatusTransferIEsValuePresentAMFUENGAPID                           int = 0
	UplinkRANStatusTransferIEsValuePresentRANUENGAPID                           int = 1
	UplinkRANStatusTransferIEsValuePresentRANStatusTransferTransparentContainer int = 2
)

type UplinkRANStatusTransferIEsValue struct {
	Present                               int
	AMFUENGAPID                           *AMFUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                           *RANUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANStatusTransferTransparentContainer *RANStatusTransferTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANStatusTransferTransparentContainer"`
}

type DownlinkRANStatusTransfer struct {
	ProtocolIEs ProtocolIEContainerDownlinkRANStatusTransferIEs
}

type ProtocolIEContainerDownlinkRANStatusTransferIEs struct {
	List []DownlinkRANStatusTransferIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DownlinkRANStatusTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                       `vht:"Reference:ProtocolIEID"`
	Value        DownlinkRANStatusTransferIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRANStatusTransferIEsValuePresentAMFUENGAPID                           int = 0
	DownlinkRANStatusTransferIEsValuePresentRANUENGAPID                           int = 1
	DownlinkRANStatusTransferIEsValuePresentRANStatusTransferTransparentContainer int = 2
)

type DownlinkRANStatusTransferIEsValue struct {
	Present                               int
	AMFUENGAPID                           *AMFUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                           *RANUENGAPID                           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANStatusTransferTransparentContainer *RANStatusTransferTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANStatusTransferTransparentContainer"`
}

type Paging struct {
	ProtocolIEs ProtocolIEContainerPagingIEs
}

type ProtocolIEContainerPagingIEs struct {
	List []PagingIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PagingIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality    `vht:"Reference:ProtocolIEID"`
	Value        PagingIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PagingIEsValuePresentUEPagingIdentity           int = 0
	PagingIEsValuePresentPagingDRX                  int = 1
	PagingIEsValuePresentTAIListForPaging           int = 2
	PagingIEsValuePresentPagingPriority             int = 3
	PagingIEsValuePresentUERadioCapabilityForPaging int = 4
	PagingIEsValuePresentPagingOrigin               int = 5
	PagingIEsValuePresentAssistanceDataForPaging    int = 6
)

type PagingIEsValue struct {
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
	List []InitialUEMessageIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type InitialUEMessageIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	Value        InitialUEMessageIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialUEMessageIEsValuePresentRANUENGAPID                         int = 0
	InitialUEMessageIEsValuePresentNASPDU                              int = 1
	InitialUEMessageIEsValuePresentUserLocationInformation             int = 2
	InitialUEMessageIEsValuePresentRRCEstablishmentCause               int = 3
	InitialUEMessageIEsValuePresentFiveGSTMSI                          int = 4
	InitialUEMessageIEsValuePresentAMFSetID                            int = 5
	InitialUEMessageIEsValuePresentUEContextRequest                    int = 6
	InitialUEMessageIEsValuePresentAllowedNSSAI                        int = 7
	InitialUEMessageIEsValuePresentSourceToTargetAMFInformationReroute int = 8
)

type InitialUEMessageIEsValue struct {
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
	List []DownlinkNASTransportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DownlinkNASTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	Value        DownlinkNASTransportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkNASTransportIEsValuePresentAMFUENGAPID               int = 0
	DownlinkNASTransportIEsValuePresentRANUENGAPID               int = 1
	DownlinkNASTransportIEsValuePresentOldAMF                    int = 2
	DownlinkNASTransportIEsValuePresentRANPagingPriority         int = 3
	DownlinkNASTransportIEsValuePresentNASPDU                    int = 4
	DownlinkNASTransportIEsValuePresentMobilityRestrictionList   int = 5
	DownlinkNASTransportIEsValuePresentIndexToRFSP               int = 6
	DownlinkNASTransportIEsValuePresentUEAggregateMaximumBitRate int = 7
	DownlinkNASTransportIEsValuePresentAllowedNSSAI              int = 8
	DownlinkNASTransportIEsValuePresentSRVCCOperationPossible    int = 9
)

type DownlinkNASTransportIEsValue struct {
	Present                   int
	AMFUENGAPID               *AMFUENGAPID               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID               *RANUENGAPID               `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	OldAMF                    *AMFName                   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFName"`
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
	List []UplinkNASTransportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UplinkNASTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	Value        UplinkNASTransportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkNASTransportIEsValuePresentAMFUENGAPID             int = 0
	UplinkNASTransportIEsValuePresentRANUENGAPID             int = 1
	UplinkNASTransportIEsValuePresentNASPDU                  int = 2
	UplinkNASTransportIEsValuePresentUserLocationInformation int = 3
)

type UplinkNASTransportIEsValue struct {
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
	List []NASNonDeliveryIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type NASNonDeliveryIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	Value        NASNonDeliveryIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	NASNonDeliveryIndicationIEsValuePresentAMFUENGAPID int = 0
	NASNonDeliveryIndicationIEsValuePresentRANUENGAPID int = 1
	NASNonDeliveryIndicationIEsValuePresentNASPDU      int = 2
	NASNonDeliveryIndicationIEsValuePresentCause       int = 3
)

type NASNonDeliveryIndicationIEsValue struct {
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
	List []RerouteNASRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type RerouteNASRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality               `vht:"Reference:ProtocolIEID"`
	Value        RerouteNASRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	RerouteNASRequestIEsValuePresentRANUENGAPID                         int = 0
	RerouteNASRequestIEsValuePresentAMFUENGAPID                         int = 1
	RerouteNASRequestIEsValuePresentNGAPMessage                         int = 2
	RerouteNASRequestIEsValuePresentAMFSetID                            int = 3
	RerouteNASRequestIEsValuePresentAllowedNSSAI                        int = 4
	RerouteNASRequestIEsValuePresentSourceToTargetAMFInformationReroute int = 5
)

type RerouteNASRequestIEsValue struct {
	Present                             int
	RANUENGAPID                         *RANUENGAPID                         `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	AMFUENGAPID                         *AMFUENGAPID                         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	NGAPMessage                         *OCTET                               `vht:"Presence:PresencePRESENCE,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDOCTET"`
	AMFSetID                            *AMFSetID                            `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFSetID"`
	AllowedNSSAI                        *AllowedNSSAI                        `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSourceToTargetAMFInformationReroute"`
}

type NGSetupRequest struct {
	ProtocolIEs ProtocolIEContainerNGSetupRequestIEs
}

type ProtocolIEContainerNGSetupRequestIEs struct {
	List []NGSetupRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type NGSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	Value        NGSetupRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupRequestIEsValuePresentGlobalRANNodeID        int = 0
	NGSetupRequestIEsValuePresentRANNodeName            int = 1
	NGSetupRequestIEsValuePresentSupportedTAList        int = 2
	NGSetupRequestIEsValuePresentDefaultPagingDRX       int = 3
	NGSetupRequestIEsValuePresentUERetentionInformation int = 4
)

type NGSetupRequestIEsValue struct {
	Present                int
	GlobalRANNodeID        *GlobalRANNodeID        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	RANNodeName            *RANNodeName            `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANNodeName"`
	SupportedTAList        *SupportedTAList        `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSupportedTAList"`
	DefaultPagingDRX       *PagingDRX              `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingDRX"`
	UERetentionInformation *UERetentionInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERetentionInformation"`
}

type NGSetupResponse struct {
	ProtocolIEs ProtocolIEContainerNGSetupResponseIEs
}

type ProtocolIEContainerNGSetupResponseIEs struct {
	List []NGSetupResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type NGSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	Value        NGSetupResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupResponseIEsValuePresentAMFName                int = 0
	NGSetupResponseIEsValuePresentServedGUAMIList        int = 1
	NGSetupResponseIEsValuePresentRelativeAMFCapacity    int = 2
	NGSetupResponseIEsValuePresentPLMNSupportList        int = 3
	NGSetupResponseIEsValuePresentCriticalityDiagnostics int = 4
	NGSetupResponseIEsValuePresentUERetentionInformation int = 5
)

type NGSetupResponseIEsValue struct {
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
	List []NGSetupFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type NGSetupFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	Value        NGSetupFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupFailureIEsValuePresentCause                  int = 0
	NGSetupFailureIEsValuePresentTimeToWait             int = 1
	NGSetupFailureIEsValuePresentCriticalityDiagnostics int = 2
)

type NGSetupFailureIEsValue struct {
	Present                int
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait             *TimeToWait             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type RANConfigurationUpdate struct {
	ProtocolIEs ProtocolIEContainerRANConfigurationUpdateIEs
}

type ProtocolIEContainerRANConfigurationUpdateIEs struct {
	List []RANConfigurationUpdateIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type RANConfigurationUpdateIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	Value        RANConfigurationUpdateIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateIEsValuePresentRANNodeName                     int = 0
	RANConfigurationUpdateIEsValuePresentSupportedTAList                 int = 1
	RANConfigurationUpdateIEsValuePresentDefaultPagingDRX                int = 2
	RANConfigurationUpdateIEsValuePresentGlobalRANNodeID                 int = 3
	RANConfigurationUpdateIEsValuePresentNGRANTNLAssociationToRemoveList int = 4
)

type RANConfigurationUpdateIEsValue struct {
	Present                         int
	RANNodeName                     *RANNodeName                     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANNodeName"`
	SupportedTAList                 *SupportedTAList                 `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSupportedTAList"`
	DefaultPagingDRX                *PagingDRX                       `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingDRX"`
	GlobalRANNodeID                 *GlobalRANNodeID                 `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	NGRANTNLAssociationToRemoveList *NGRANTNLAssociationToRemoveList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNGRANTNLAssociationToRemoveList"`
}

type RANConfigurationUpdateAcknowledge struct {
	ProtocolIEs ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs
}

type ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs struct {
	List []RANConfigurationUpdateAcknowledgeIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type RANConfigurationUpdateAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	Value        RANConfigurationUpdateAcknowledgeIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateAcknowledgeIEsValuePresentCriticalityDiagnostics int = 0
)

type RANConfigurationUpdateAcknowledgeIEsValue struct {
	Present                int
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type RANConfigurationUpdateFailure struct {
	ProtocolIEs ProtocolIEContainerRANConfigurationUpdateFailureIEs
}

type ProtocolIEContainerRANConfigurationUpdateFailureIEs struct {
	List []RANConfigurationUpdateFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type RANConfigurationUpdateFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                           `vht:"Reference:ProtocolIEID"`
	Value        RANConfigurationUpdateFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateFailureIEsValuePresentCause                  int = 0
	RANConfigurationUpdateFailureIEsValuePresentTimeToWait             int = 1
	RANConfigurationUpdateFailureIEsValuePresentCriticalityDiagnostics int = 2
)

type RANConfigurationUpdateFailureIEsValue struct {
	Present                int
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait             *TimeToWait             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFConfigurationUpdate struct {
	ProtocolIEs ProtocolIEContainerAMFConfigurationUpdateIEs
}

type ProtocolIEContainerAMFConfigurationUpdateIEs struct {
	List []AMFConfigurationUpdateIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type AMFConfigurationUpdateIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	Value        AMFConfigurationUpdateIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateIEsValuePresentAMFName                       int = 0
	AMFConfigurationUpdateIEsValuePresentServedGUAMIList               int = 1
	AMFConfigurationUpdateIEsValuePresentRelativeAMFCapacity           int = 2
	AMFConfigurationUpdateIEsValuePresentPLMNSupportList               int = 3
	AMFConfigurationUpdateIEsValuePresentAMFTNLAssociationToAddList    int = 4
	AMFConfigurationUpdateIEsValuePresentAMFTNLAssociationToRemoveList int = 5
	AMFConfigurationUpdateIEsValuePresentAMFTNLAssociationToUpdateList int = 6
)

type AMFConfigurationUpdateIEsValue struct {
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
	List []AMFConfigurationUpdateAcknowledgeIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type AMFConfigurationUpdateAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	Value        AMFConfigurationUpdateAcknowledgeIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateAcknowledgeIEsValuePresentAMFTNLAssociationSetupList         int = 0
	AMFConfigurationUpdateAcknowledgeIEsValuePresentAMFTNLAssociationFailedToSetupList int = 1
	AMFConfigurationUpdateAcknowledgeIEsValuePresentCriticalityDiagnostics             int = 2
)

type AMFConfigurationUpdateAcknowledgeIEsValue struct {
	Present                            int
	AMFTNLAssociationSetupList         *AMFTNLAssociationSetupList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationSetupList"`
	AMFTNLAssociationFailedToSetupList *TNLAssociationList         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTNLAssociationList"`
	CriticalityDiagnostics             *CriticalityDiagnostics     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFConfigurationUpdateFailure struct {
	ProtocolIEs ProtocolIEContainerAMFConfigurationUpdateFailureIEs
}

type ProtocolIEContainerAMFConfigurationUpdateFailureIEs struct {
	List []AMFConfigurationUpdateFailureIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type AMFConfigurationUpdateFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                           `vht:"Reference:ProtocolIEID"`
	Value        AMFConfigurationUpdateFailureIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateFailureIEsValuePresentCause                  int = 0
	AMFConfigurationUpdateFailureIEsValuePresentTimeToWait             int = 1
	AMFConfigurationUpdateFailureIEsValuePresentCriticalityDiagnostics int = 2
)

type AMFConfigurationUpdateFailureIEsValue struct {
	Present                int
	Cause                  *Cause                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait             *TimeToWait             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFStatusIndication struct {
	ProtocolIEs ProtocolIEContainerAMFStatusIndicationIEs
}

type ProtocolIEContainerAMFStatusIndicationIEs struct {
	List []AMFStatusIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type AMFStatusIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                 `vht:"Reference:ProtocolIEID"`
	Value        AMFStatusIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFStatusIndicationIEsValuePresentUnavailableGUAMIList int = 0
)

type AMFStatusIndicationIEsValue struct {
	Present              int
	UnavailableGUAMIList *UnavailableGUAMIList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUnavailableGUAMIList"`
}

type NGReset struct {
	ProtocolIEs ProtocolIEContainerNGResetIEs
}

type ProtocolIEContainerNGResetIEs struct {
	List []NGResetIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type NGResetIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality     `vht:"Reference:ProtocolIEID"`
	Value        NGResetIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGResetIEsValuePresentCause     int = 0
	NGResetIEsValuePresentResetType int = 1
)

type NGResetIEsValue struct {
	Present   int
	Cause     *Cause     `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	ResetType *ResetType `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDResetType"`
}

type NGResetAcknowledge struct {
	ProtocolIEs ProtocolIEContainerNGResetAcknowledgeIEs
}

type ProtocolIEContainerNGResetAcknowledgeIEs struct {
	List []NGResetAcknowledgeIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type NGResetAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	Value        NGResetAcknowledgeIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGResetAcknowledgeIEsValuePresentUEAssociatedLogicalNGConnectionList int = 0
	NGResetAcknowledgeIEsValuePresentCriticalityDiagnostics              int = 1
)

type NGResetAcknowledgeIEsValue struct {
	Present                             int
	UEAssociatedLogicalNGConnectionList *UEAssociatedLogicalNGConnectionList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAssociatedLogicalNGConnectionList"`
	CriticalityDiagnostics              *CriticalityDiagnostics              `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type ErrorIndication struct {
	ProtocolIEs ProtocolIEContainerErrorIndicationIEs
}

type ProtocolIEContainerErrorIndicationIEs struct {
	List []ErrorIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type ErrorIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	Value        ErrorIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	ErrorIndicationIEsValuePresentAMFUENGAPID            int = 0
	ErrorIndicationIEsValuePresentRANUENGAPID            int = 1
	ErrorIndicationIEsValuePresentCause                  int = 2
	ErrorIndicationIEsValuePresentCriticalityDiagnostics int = 3
)

type ErrorIndicationIEsValue struct {
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
	List []OverloadStartIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type OverloadStartIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality           `vht:"Reference:ProtocolIEID"`
	Value        OverloadStartIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	OverloadStartIEsValuePresentAMFOverloadResponse               int = 0
	OverloadStartIEsValuePresentAMFTrafficLoadReductionIndication int = 1
	OverloadStartIEsValuePresentOverloadStartNSSAIList            int = 2
)

type OverloadStartIEsValue struct {
	Present                           int
	AMFOverloadResponse               *OverloadResponse               `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDOverloadResponse"`
	AMFTrafficLoadReductionIndication *TrafficLoadReductionIndication `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTrafficLoadReductionIndication"`
	OverloadStartNSSAIList            *OverloadStartNSSAIList         `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDOverloadStartNSSAIList"`
}

type OverloadStop struct {
	ProtocolIEs ProtocolIEContainerOverloadStopIEs
}

type ProtocolIEContainerOverloadStopIEs struct {
	List []OverloadStopIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type OverloadStopIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality          `vht:"Reference:ProtocolIEID"`
	Value        OverloadStopIEsValue `vht:"Reference:ProtocolIEID"`
}

const ()

type OverloadStopIEsValue struct {
	Present int
}

type UplinkRANConfigurationTransfer struct {
	ProtocolIEs ProtocolIEContainerUplinkRANConfigurationTransferIEs
}

type ProtocolIEContainerUplinkRANConfigurationTransferIEs struct {
	List []UplinkRANConfigurationTransferIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UplinkRANConfigurationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	Value        UplinkRANConfigurationTransferIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRANConfigurationTransferIEsValuePresentSONConfigurationTransferUL     int = 0
	UplinkRANConfigurationTransferIEsValuePresentENDCSONConfigurationTransferUL int = 1
)

type UplinkRANConfigurationTransferIEsValue struct {
	Present                        int
	SONConfigurationTransferUL     *SONConfigurationTransfer     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSONConfigurationTransfer"`
	ENDCSONConfigurationTransferUL *ENDCSONConfigurationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDENDCSONConfigurationTransfer"`
}

type DownlinkRANConfigurationTransfer struct {
	ProtocolIEs ProtocolIEContainerDownlinkRANConfigurationTransferIEs
}

type ProtocolIEContainerDownlinkRANConfigurationTransferIEs struct {
	List []DownlinkRANConfigurationTransferIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DownlinkRANConfigurationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	Value        DownlinkRANConfigurationTransferIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRANConfigurationTransferIEsValuePresentSONConfigurationTransferDL     int = 0
	DownlinkRANConfigurationTransferIEsValuePresentENDCSONConfigurationTransferDL int = 1
)

type DownlinkRANConfigurationTransferIEsValue struct {
	Present                        int
	SONConfigurationTransferDL     *SONConfigurationTransfer     `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSONConfigurationTransfer"`
	ENDCSONConfigurationTransferDL *ENDCSONConfigurationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDENDCSONConfigurationTransfer"`
}

type WriteReplaceWarningRequest struct {
	ProtocolIEs ProtocolIEContainerWriteReplaceWarningRequestIEs
}

type ProtocolIEContainerWriteReplaceWarningRequestIEs struct {
	List []WriteReplaceWarningRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type WriteReplaceWarningRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	Value        WriteReplaceWarningRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	WriteReplaceWarningRequestIEsValuePresentMessageIdentifier           int = 0
	WriteReplaceWarningRequestIEsValuePresentSerialNumber                int = 1
	WriteReplaceWarningRequestIEsValuePresentWarningAreaList             int = 2
	WriteReplaceWarningRequestIEsValuePresentRepetitionPeriod            int = 3
	WriteReplaceWarningRequestIEsValuePresentNumberOfBroadcastsRequested int = 4
	WriteReplaceWarningRequestIEsValuePresentWarningType                 int = 5
	WriteReplaceWarningRequestIEsValuePresentWarningSecurityInfo         int = 6
	WriteReplaceWarningRequestIEsValuePresentDataCodingScheme            int = 7
	WriteReplaceWarningRequestIEsValuePresentWarningMessageContents      int = 8
	WriteReplaceWarningRequestIEsValuePresentConcurrentWarningMessageInd int = 9
	WriteReplaceWarningRequestIEsValuePresentWarningAreaCoordinates      int = 10
)

type WriteReplaceWarningRequestIEsValue struct {
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
	List []WriteReplaceWarningResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type WriteReplaceWarningResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	Value        WriteReplaceWarningResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	WriteReplaceWarningResponseIEsValuePresentMessageIdentifier          int = 0
	WriteReplaceWarningResponseIEsValuePresentSerialNumber               int = 1
	WriteReplaceWarningResponseIEsValuePresentBroadcastCompletedAreaList int = 2
	WriteReplaceWarningResponseIEsValuePresentCriticalityDiagnostics     int = 3
)

type WriteReplaceWarningResponseIEsValue struct {
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
	List []PWSCancelRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PWSCancelRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	Value        PWSCancelRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSCancelRequestIEsValuePresentMessageIdentifier        int = 0
	PWSCancelRequestIEsValuePresentSerialNumber             int = 1
	PWSCancelRequestIEsValuePresentWarningAreaList          int = 2
	PWSCancelRequestIEsValuePresentCancelAllWarningMessages int = 3
)

type PWSCancelRequestIEsValue struct {
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
	List []PWSCancelResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PWSCancelResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality               `vht:"Reference:ProtocolIEID"`
	Value        PWSCancelResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSCancelResponseIEsValuePresentMessageIdentifier          int = 0
	PWSCancelResponseIEsValuePresentSerialNumber               int = 1
	PWSCancelResponseIEsValuePresentBroadcastCancelledAreaList int = 2
	PWSCancelResponseIEsValuePresentCriticalityDiagnostics     int = 3
)

type PWSCancelResponseIEsValue struct {
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
	List []PWSRestartIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PWSRestartIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	Value        PWSRestartIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSRestartIndicationIEsValuePresentCellIDListForRestart          int = 0
	PWSRestartIndicationIEsValuePresentGlobalRANNodeID               int = 1
	PWSRestartIndicationIEsValuePresentTAIListForRestart             int = 2
	PWSRestartIndicationIEsValuePresentEmergencyAreaIDListForRestart int = 3
)

type PWSRestartIndicationIEsValue struct {
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
	List []PWSFailureIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PWSFailureIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                  `vht:"Reference:ProtocolIEID"`
	Value        PWSFailureIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSFailureIndicationIEsValuePresentPWSFailedCellIDList int = 0
	PWSFailureIndicationIEsValuePresentGlobalRANNodeID     int = 1
)

type PWSFailureIndicationIEsValue struct {
	Present             int
	PWSFailedCellIDList *PWSFailedCellIDList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPWSFailedCellIDList"`
	GlobalRANNodeID     *GlobalRANNodeID     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
}

type DownlinkUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkUEAssociatedNRPPaTransportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DownlinkUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	Value        DownlinkUEAssociatedNRPPaTransportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkUEAssociatedNRPPaTransportIEsValuePresentAMFUENGAPID int = 0
	DownlinkUEAssociatedNRPPaTransportIEsValuePresentRANUENGAPID int = 1
	DownlinkUEAssociatedNRPPaTransportIEsValuePresentRoutingID   int = 2
	DownlinkUEAssociatedNRPPaTransportIEsValuePresentNRPPaPDU    int = 3
)

type DownlinkUEAssociatedNRPPaTransportIEsValue struct {
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
	List []UplinkUEAssociatedNRPPaTransportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UplinkUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                              `vht:"Reference:ProtocolIEID"`
	Value        UplinkUEAssociatedNRPPaTransportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkUEAssociatedNRPPaTransportIEsValuePresentAMFUENGAPID int = 0
	UplinkUEAssociatedNRPPaTransportIEsValuePresentRANUENGAPID int = 1
	UplinkUEAssociatedNRPPaTransportIEsValuePresentRoutingID   int = 2
	UplinkUEAssociatedNRPPaTransportIEsValuePresentNRPPaPDU    int = 3
)

type UplinkUEAssociatedNRPPaTransportIEsValue struct {
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
	List []DownlinkNonUEAssociatedNRPPaTransportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DownlinkNonUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                   `vht:"Reference:ProtocolIEID"`
	Value        DownlinkNonUEAssociatedNRPPaTransportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkNonUEAssociatedNRPPaTransportIEsValuePresentRoutingID int = 0
	DownlinkNonUEAssociatedNRPPaTransportIEsValuePresentNRPPaPDU  int = 1
)

type DownlinkNonUEAssociatedNRPPaTransportIEsValue struct {
	Present   int
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU  *NRPPaPDU  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type UplinkNonUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs struct {
	List []UplinkNonUEAssociatedNRPPaTransportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UplinkNonUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                 `vht:"Reference:ProtocolIEID"`
	Value        UplinkNonUEAssociatedNRPPaTransportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkNonUEAssociatedNRPPaTransportIEsValuePresentRoutingID int = 0
	UplinkNonUEAssociatedNRPPaTransportIEsValuePresentNRPPaPDU  int = 1
)

type UplinkNonUEAssociatedNRPPaTransportIEsValue struct {
	Present   int
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU  *NRPPaPDU  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type TraceStart struct {
	ProtocolIEs ProtocolIEContainerTraceStartIEs
}

type ProtocolIEContainerTraceStartIEs struct {
	List []TraceStartIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type TraceStartIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality        `vht:"Reference:ProtocolIEID"`
	Value        TraceStartIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	TraceStartIEsValuePresentAMFUENGAPID     int = 0
	TraceStartIEsValuePresentRANUENGAPID     int = 1
	TraceStartIEsValuePresentTraceActivation int = 2
)

type TraceStartIEsValue struct {
	Present         int
	AMFUENGAPID     *AMFUENGAPID     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID     *RANUENGAPID     `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	TraceActivation *TraceActivation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceActivation"`
}

type TraceFailureIndication struct {
	ProtocolIEs ProtocolIEContainerTraceFailureIndicationIEs
}

type ProtocolIEContainerTraceFailureIndicationIEs struct {
	List []TraceFailureIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type TraceFailureIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	Value        TraceFailureIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	TraceFailureIndicationIEsValuePresentAMFUENGAPID  int = 0
	TraceFailureIndicationIEsValuePresentRANUENGAPID  int = 1
	TraceFailureIndicationIEsValuePresentNGRANTraceID int = 2
	TraceFailureIndicationIEsValuePresentCause        int = 3
)

type TraceFailureIndicationIEsValue struct {
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
	List []DeactivateTraceIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DeactivateTraceIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	Value        DeactivateTraceIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DeactivateTraceIEsValuePresentAMFUENGAPID  int = 0
	DeactivateTraceIEsValuePresentRANUENGAPID  int = 1
	DeactivateTraceIEsValuePresentNGRANTraceID int = 2
)

type DeactivateTraceIEsValue struct {
	Present      int
	AMFUENGAPID  *AMFUENGAPID  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID  *RANUENGAPID  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NGRANTraceID *NGRANTraceID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANTraceID"`
}

type CellTrafficTrace struct {
	ProtocolIEs ProtocolIEContainerCellTrafficTraceIEs
}

type ProtocolIEContainerCellTrafficTraceIEs struct {
	List []CellTrafficTraceIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type CellTrafficTraceIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	Value        CellTrafficTraceIEsValue `vht:"Reference:ProtocolIEID"`
}

const ()

type CellTrafficTraceIEsValue struct {
	Present int
}

type LocationReportingControl struct {
	ProtocolIEs ProtocolIEContainerLocationReportingControlIEs
}

type ProtocolIEContainerLocationReportingControlIEs struct {
	List []LocationReportingControlIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type LocationReportingControlIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                      `vht:"Reference:ProtocolIEID"`
	Value        LocationReportingControlIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportingControlIEsValuePresentAMFUENGAPID                  int = 0
	LocationReportingControlIEsValuePresentRANUENGAPID                  int = 1
	LocationReportingControlIEsValuePresentLocationReportingRequestType int = 2
)

type LocationReportingControlIEsValue struct {
	Present                      int
	AMFUENGAPID                  *AMFUENGAPID                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID                  *RANUENGAPID                  `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	LocationReportingRequestType *LocationReportingRequestType `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
}

type LocationReportingFailureIndication struct {
	ProtocolIEs ProtocolIEContainerLocationReportingFailureIndicationIEs
}

type ProtocolIEContainerLocationReportingFailureIndicationIEs struct {
	List []LocationReportingFailureIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type LocationReportingFailureIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	Value        LocationReportingFailureIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportingFailureIndicationIEsValuePresentAMFUENGAPID int = 0
	LocationReportingFailureIndicationIEsValuePresentRANUENGAPID int = 1
	LocationReportingFailureIndicationIEsValuePresentCause       int = 2
)

type LocationReportingFailureIndicationIEsValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause       *Cause       `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type LocationReport struct {
	ProtocolIEs ProtocolIEContainerLocationReportIEs
}

type ProtocolIEContainerLocationReportIEs struct {
	List []LocationReportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type LocationReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	Value        LocationReportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportIEsValuePresentAMFUENGAPID                    int = 0
	LocationReportIEsValuePresentRANUENGAPID                    int = 1
	LocationReportIEsValuePresentUserLocationInformation        int = 2
	LocationReportIEsValuePresentUEPresenceInAreaOfInterestList int = 3
	LocationReportIEsValuePresentLocationReportingRequestType   int = 4
)

type LocationReportIEsValue struct {
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
	List []UETNLABindingReleaseRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UETNLABindingReleaseRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	Value        UETNLABindingReleaseRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UETNLABindingReleaseRequestIEsValuePresentAMFUENGAPID int = 0
	UETNLABindingReleaseRequestIEsValuePresentRANUENGAPID int = 1
)

type UETNLABindingReleaseRequestIEsValue struct {
	Present     int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
}

type UERadioCapabilityInfoIndication struct {
	ProtocolIEs ProtocolIEContainerUERadioCapabilityInfoIndicationIEs
}

type ProtocolIEContainerUERadioCapabilityInfoIndicationIEs struct {
	List []UERadioCapabilityInfoIndicationIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UERadioCapabilityInfoIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                             `vht:"Reference:ProtocolIEID"`
	Value        UERadioCapabilityInfoIndicationIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityInfoIndicationIEsValuePresentAMFUENGAPID                int = 0
	UERadioCapabilityInfoIndicationIEsValuePresentRANUENGAPID                int = 1
	UERadioCapabilityInfoIndicationIEsValuePresentUERadioCapability          int = 2
	UERadioCapabilityInfoIndicationIEsValuePresentUERadioCapabilityForPaging int = 3
)

type UERadioCapabilityInfoIndicationIEsValue struct {
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
	List []UERadioCapabilityCheckRequestIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UERadioCapabilityCheckRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                           `vht:"Reference:ProtocolIEID"`
	Value        UERadioCapabilityCheckRequestIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityCheckRequestIEsValuePresentAMFUENGAPID       int = 0
	UERadioCapabilityCheckRequestIEsValuePresentRANUENGAPID       int = 1
	UERadioCapabilityCheckRequestIEsValuePresentUERadioCapability int = 2
)

type UERadioCapabilityCheckRequestIEsValue struct {
	Present           int
	AMFUENGAPID       *AMFUENGAPID       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID       *RANUENGAPID       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UERadioCapability *UERadioCapability `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapability"`
}

type UERadioCapabilityCheckResponse struct {
	ProtocolIEs ProtocolIEContainerUERadioCapabilityCheckResponseIEs
}

type ProtocolIEContainerUERadioCapabilityCheckResponseIEs struct {
	List []UERadioCapabilityCheckResponseIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UERadioCapabilityCheckResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	Value        UERadioCapabilityCheckResponseIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityCheckResponseIEsValuePresentAMFUENGAPID              int = 0
	UERadioCapabilityCheckResponseIEsValuePresentRANUENGAPID              int = 1
	UERadioCapabilityCheckResponseIEsValuePresentIMSVoiceSupportIndicator int = 2
	UERadioCapabilityCheckResponseIEsValuePresentCriticalityDiagnostics   int = 3
)

type UERadioCapabilityCheckResponseIEsValue struct {
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
	List []PrivateMessageIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type PrivateMessageIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	Value        PrivateMessageIEsValue `vht:"Reference:ProtocolIEID"`
}

const ()

type PrivateMessageIEsValue struct {
	Present int
}

type SecondaryRATDataUsageReport struct {
	ProtocolIEs ProtocolIEContainerSecondaryRATDataUsageReportIEs
}

type ProtocolIEContainerSecondaryRATDataUsageReportIEs struct {
	List []SecondaryRATDataUsageReportIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type SecondaryRATDataUsageReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	Value        SecondaryRATDataUsageReportIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	SecondaryRATDataUsageReportIEsValuePresentAMFUENGAPID                             int = 0
	SecondaryRATDataUsageReportIEsValuePresentRANUENGAPID                             int = 1
	SecondaryRATDataUsageReportIEsValuePresentPDUSessionResourceSecondaryRATUsageList int = 2
	SecondaryRATDataUsageReportIEsValuePresentHandoverFlag                            int = 3
	SecondaryRATDataUsageReportIEsValuePresentUserLocationInformation                 int = 4
)

type SecondaryRATDataUsageReportIEsValue struct {
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
	List []UplinkRIMInformationTransferIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type UplinkRIMInformationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                          `vht:"Reference:ProtocolIEID"`
	Value        UplinkRIMInformationTransferIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRIMInformationTransferIEsValuePresentRIMInformationTransfer int = 0
)

type UplinkRIMInformationTransferIEsValue struct {
	Present                int
	RIMInformationTransfer *RIMInformationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRIMInformationTransfer"`
}

type DownlinkRIMInformationTransfer struct {
	ProtocolIEs ProtocolIEContainerDownlinkRIMInformationTransferIEs
}

type ProtocolIEContainerDownlinkRIMInformationTransferIEs struct {
	List []DownlinkRIMInformationTransferIEs `vht:"sizeMin:0,sizeMax:65535"`
}

type DownlinkRIMInformationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	Value        DownlinkRIMInformationTransferIEsValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRIMInformationTransferIEsValuePresentRIMInformationTransfer int = 0
)

type DownlinkRIMInformationTransferIEsValue struct {
	Present                int
	RIMInformationTransfer *RIMInformationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRIMInformationTransfer"`
}
