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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceSetupRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupRequestIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceSetupRequestIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceSetupRequestIEsTypeValuePresentRANPagingPriority int = 2
	PDUSessionResourceSetupRequestIEsTypeValuePresentNASPDU int = 3
	PDUSessionResourceSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListSUReq int = 4
	PDUSessionResourceSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate int = 5
)

type PDUSessionResourceSetupRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority *RANPagingPriority `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	NASPDU *NASPDU `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	PDUSessionResourceSetupListSUReq *PDUSessionResourceSetupListSUReq `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListSUReq"`
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
}

type PDUSessionResourceSetupResponse struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceSetupResponseIEs
}

type ProtocolIEContainerPDUSessionResourceSetupResponseIEs struct {
	List []PDUSessionResourceSetupResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceSetupResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupResponseIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceSetupResponseIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListSURes int = 2
	PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListSURes int = 3
	PDUSessionResourceSetupResponseIEsTypeValuePresentCriticalityDiagnostics int = 4
)

type PDUSessionResourceSetupResponseIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceSetupListSURes *PDUSessionResourceSetupListSURes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListSURes"`
	PDUSessionResourceFailedToSetupListSURes *PDUSessionResourceFailedToSetupListSURes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListSURes"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PDUSessionResourceReleaseCommand struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceReleaseCommandIEs
}

type ProtocolIEContainerPDUSessionResourceReleaseCommandIEs struct {
	List []PDUSessionResourceReleaseCommandIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceReleaseCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceReleaseCommandIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceReleaseCommandIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceReleaseCommandIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceReleaseCommandIEsTypeValuePresentRANPagingPriority int = 2
	PDUSessionResourceReleaseCommandIEsTypeValuePresentNASPDU int = 3
	PDUSessionResourceReleaseCommandIEsTypeValuePresentPDUSessionResourceToReleaseListRelCmd int = 4
)

type PDUSessionResourceReleaseCommandIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority *RANPagingPriority `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	NASPDU *NASPDU `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNASPDU"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceReleaseResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceReleaseResponseIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceReleaseResponseIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceReleaseResponseIEsTypeValuePresentPDUSessionResourceReleasedListRelRes int = 2
	PDUSessionResourceReleaseResponseIEsTypeValuePresentUserLocationInformation int = 3
	PDUSessionResourceReleaseResponseIEsTypeValuePresentCriticalityDiagnostics int = 4
)

type PDUSessionResourceReleaseResponseIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceReleasedListRelRes *PDUSessionResourceReleasedListRelRes `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListRelRes"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PDUSessionResourceModifyRequest struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyRequestIEs
}

type ProtocolIEContainerPDUSessionResourceModifyRequestIEs struct {
	List []PDUSessionResourceModifyRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceModifyRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyRequestIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceModifyRequestIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceModifyRequestIEsTypeValuePresentRANPagingPriority int = 2
	PDUSessionResourceModifyRequestIEsTypeValuePresentPDUSessionResourceModifyListModReq int = 3
)

type PDUSessionResourceModifyRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority *RANPagingPriority `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceModifyResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyResponseIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceModifyResponseIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceModifyListModRes int = 2
	PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceFailedToModifyListModRes int = 3
	PDUSessionResourceModifyResponseIEsTypeValuePresentUserLocationInformation int = 4
	PDUSessionResourceModifyResponseIEsTypeValuePresentCriticalityDiagnostics int = 5
)

type PDUSessionResourceModifyResponseIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceModifyListModRes *PDUSessionResourceModifyListModRes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModRes"`
	PDUSessionResourceFailedToModifyListModRes *PDUSessionResourceFailedToModifyListModRes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToModifyListModRes"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PDUSessionResourceNotify struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceNotifyIEs
}

type ProtocolIEContainerPDUSessionResourceNotifyIEs struct {
	List []PDUSessionResourceNotifyIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceNotifyIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceNotifyIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceNotifyIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceNotifyIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceNotifyList int = 2
	PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceReleasedListNot int = 3
	PDUSessionResourceNotifyIEsTypeValuePresentUserLocationInformation int = 4
)

type PDUSessionResourceNotifyIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceNotifyList *PDUSessionResourceNotifyList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceNotifyList"`
	PDUSessionResourceReleasedListNot *PDUSessionResourceReleasedListNot `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListNot"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type PDUSessionResourceModifyIndication struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyIndicationIEs
}

type ProtocolIEContainerPDUSessionResourceModifyIndicationIEs struct {
	List []PDUSessionResourceModifyIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceModifyIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyIndicationIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceModifyIndicationIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceModifyIndicationIEsTypeValuePresentPDUSessionResourceModifyListModInd int = 2
	PDUSessionResourceModifyIndicationIEsTypeValuePresentUserLocationInformation int = 3
)

type PDUSessionResourceModifyIndicationIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceModifyListModInd *PDUSessionResourceModifyListModInd `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModInd"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type PDUSessionResourceModifyConfirm struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyConfirmIEs
}

type ProtocolIEContainerPDUSessionResourceModifyConfirmIEs struct {
	List []PDUSessionResourceModifyConfirmIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PDUSessionResourceModifyConfirmIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PDUSessionResourceModifyConfirmIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyConfirmIEsTypeValuePresentAMFUENGAPID int = 0
	PDUSessionResourceModifyConfirmIEsTypeValuePresentRANUENGAPID int = 1
	PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceModifyListModCfm int = 2
	PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceFailedToModifyListModCfm int = 3
	PDUSessionResourceModifyConfirmIEsTypeValuePresentCriticalityDiagnostics int = 4
)

type PDUSessionResourceModifyConfirmIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceModifyListModCfm *PDUSessionResourceModifyListModCfm `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceModifyListModCfm"`
	PDUSessionResourceFailedToModifyListModCfm *PDUSessionResourceFailedToModifyListModCfm `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type InitialContextSetupRequest struct {
	ProtocolIEs ProtocolIEContainerInitialContextSetupRequestIEs
}

type ProtocolIEContainerInitialContextSetupRequestIEs struct {
	List []InitialContextSetupRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialContextSetupRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue InitialContextSetupRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupRequestIEsTypeValuePresentAMFUENGAPID int = 0
	InitialContextSetupRequestIEsTypeValuePresentRANUENGAPID int = 1
	InitialContextSetupRequestIEsTypeValuePresentOldAMF int = 2
	InitialContextSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate int = 3
	InitialContextSetupRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive int = 4
	InitialContextSetupRequestIEsTypeValuePresentGUAMI int = 5
	InitialContextSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListCxtReq int = 6
	InitialContextSetupRequestIEsTypeValuePresentAllowedNSSAI int = 7
	InitialContextSetupRequestIEsTypeValuePresentUESecurityCapabilities int = 8
	InitialContextSetupRequestIEsTypeValuePresentSecurityKey int = 9
	InitialContextSetupRequestIEsTypeValuePresentTraceActivation int = 10
	InitialContextSetupRequestIEsTypeValuePresentMobilityRestrictionList int = 11
	InitialContextSetupRequestIEsTypeValuePresentUERadioCapability int = 12
	InitialContextSetupRequestIEsTypeValuePresentIndexToRFSP int = 13
	InitialContextSetupRequestIEsTypeValuePresentMaskedIMEISV int = 14
	InitialContextSetupRequestIEsTypeValuePresentNASPDU int = 15
	InitialContextSetupRequestIEsTypeValuePresentEmergencyFallbackIndicator int = 16
	InitialContextSetupRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest int = 17
	InitialContextSetupRequestIEsTypeValuePresentUERadioCapabilityForPaging int = 18
	InitialContextSetupRequestIEsTypeValuePresentRedirectionVoiceFallback int = 19
	InitialContextSetupRequestIEsTypeValuePresentLocationReportingRequestType int = 20
	InitialContextSetupRequestIEsTypeValuePresentCNAssistedRANTuning int = 21
	InitialContextSetupRequestIEsTypeValuePresentSRVCCOperationPossible int = 22
)

type InitialContextSetupRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	OldAMF *AMFName `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDOldAMF"`
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate `vht:"Presence:PresenceConditional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	GUAMI *GUAMI `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGUAMI"`
	PDUSessionResourceSetupListCxtReq *PDUSessionResourceSetupListCxtReq `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListCxtReq"`
	AllowedNSSAI *AllowedNSSAI `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	UESecurityCapabilities *UESecurityCapabilities `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityKey *SecurityKey `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityKey"`
	TraceActivation *TraceActivation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceActivation"`
	MobilityRestrictionList *MobilityRestrictionList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMobilityRestrictionList"`
	UERadioCapability *UERadioCapability `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapability"`
	IndexToRFSP *IndexToRFSP `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDIndexToRFSP"`
	MaskedIMEISV *MaskedIMEISV `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMaskedIMEISV"`
	NASPDU *NASPDU `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNASPDU"`
	EmergencyFallbackIndicator *EmergencyFallbackIndicator `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDEmergencyFallbackIndicator"`
	RRCInactiveTransitionReportRequest *RRCInactiveTransitionReportRequest `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	UERadioCapabilityForPaging *UERadioCapabilityForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapabilityForPaging"`
	RedirectionVoiceFallback *RedirectionVoiceFallback `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRedirectionVoiceFallback"`
	LocationReportingRequestType *LocationReportingRequestType `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
	CNAssistedRANTuning *CNAssistedRANTuning `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible *SRVCCOperationPossible `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type InitialContextSetupResponse struct {
	ProtocolIEs ProtocolIEContainerInitialContextSetupResponseIEs
}

type ProtocolIEContainerInitialContextSetupResponseIEs struct {
	List []InitialContextSetupResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialContextSetupResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue InitialContextSetupResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupResponseIEsTypeValuePresentAMFUENGAPID int = 0
	InitialContextSetupResponseIEsTypeValuePresentRANUENGAPID int = 1
	InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListCxtRes int = 2
	InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtRes int = 3
	InitialContextSetupResponseIEsTypeValuePresentCriticalityDiagnostics int = 4
)

type InitialContextSetupResponseIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceSetupListCxtRes *PDUSessionResourceSetupListCxtRes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListCxtRes"`
	PDUSessionResourceFailedToSetupListCxtRes *PDUSessionResourceFailedToSetupListCxtRes `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type InitialContextSetupFailure struct {
	ProtocolIEs ProtocolIEContainerInitialContextSetupFailureIEs
}

type ProtocolIEContainerInitialContextSetupFailureIEs struct {
	List []InitialContextSetupFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialContextSetupFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue InitialContextSetupFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialContextSetupFailureIEsTypeValuePresentAMFUENGAPID int = 0
	InitialContextSetupFailureIEsTypeValuePresentRANUENGAPID int = 1
	InitialContextSetupFailureIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtFail int = 2
	InitialContextSetupFailureIEsTypeValuePresentCause int = 3
	InitialContextSetupFailureIEsTypeValuePresentCriticalityDiagnostics int = 4
)

type InitialContextSetupFailureIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceFailedToSetupListCxtFail *PDUSessionResourceFailedToSetupListCxtFail `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UEContextReleaseRequest struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseRequestIEs
}

type ProtocolIEContainerUEContextReleaseRequestIEs struct {
	List []UEContextReleaseRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextReleaseRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UEContextReleaseRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseRequestIEsTypeValuePresentAMFUENGAPID int = 0
	UEContextReleaseRequestIEsTypeValuePresentRANUENGAPID int = 1
	UEContextReleaseRequestIEsTypeValuePresentPDUSessionResourceListCxtRelReq int = 2
	UEContextReleaseRequestIEsTypeValuePresentCause int = 3
)

type UEContextReleaseRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceListCxtRelReq *PDUSessionResourceListCxtRelReq `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceListCxtRelReq"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type UEContextReleaseCommand struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseCommandIEs
}

type ProtocolIEContainerUEContextReleaseCommandIEs struct {
	List []UEContextReleaseCommandIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextReleaseCommandIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UEContextReleaseCommandIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseCommandIEsTypeValuePresentUENGAPIDs int = 0
	UEContextReleaseCommandIEsTypeValuePresentCause int = 1
)

type UEContextReleaseCommandIEsTypeValue struct {
	Present int
	UENGAPIDs *UENGAPIDs `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUENGAPIDs"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type UEContextReleaseComplete struct {
	ProtocolIEs ProtocolIEContainerUEContextReleaseCompleteIEs
}

type ProtocolIEContainerUEContextReleaseCompleteIEs struct {
	List []UEContextReleaseCompleteIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextReleaseCompleteIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UEContextReleaseCompleteIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextReleaseCompleteIEsTypeValuePresentAMFUENGAPID int = 0
	UEContextReleaseCompleteIEsTypeValuePresentRANUENGAPID int = 1
	UEContextReleaseCompleteIEsTypeValuePresentUserLocationInformation int = 2
	UEContextReleaseCompleteIEsTypeValuePresentInfoOnRecommendedCellsAndRANNodesForPaging int = 3
	UEContextReleaseCompleteIEsTypeValuePresentPDUSessionResourceListCxtRelCpl int = 4
	UEContextReleaseCompleteIEsTypeValuePresentCriticalityDiagnostics int = 5
)

type UEContextReleaseCompleteIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	InfoOnRecommendedCellsAndRANNodesForPaging *InfoOnRecommendedCellsAndRANNodesForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging"`
	PDUSessionResourceListCxtRelCpl *PDUSessionResourceListCxtRelCpl `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceListCxtRelCpl"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UEContextModificationRequest struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationRequestIEs
}

type ProtocolIEContainerUEContextModificationRequestIEs struct {
	List []UEContextModificationRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextModificationRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UEContextModificationRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationRequestIEsTypeValuePresentAMFUENGAPID int = 0
	UEContextModificationRequestIEsTypeValuePresentRANUENGAPID int = 1
	UEContextModificationRequestIEsTypeValuePresentRANPagingPriority int = 2
	UEContextModificationRequestIEsTypeValuePresentSecurityKey int = 3
	UEContextModificationRequestIEsTypeValuePresentIndexToRFSP int = 4
	UEContextModificationRequestIEsTypeValuePresentUEAggregateMaximumBitRate int = 5
	UEContextModificationRequestIEsTypeValuePresentUESecurityCapabilities int = 6
	UEContextModificationRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive int = 7
	UEContextModificationRequestIEsTypeValuePresentEmergencyFallbackIndicator int = 8
	UEContextModificationRequestIEsTypeValuePresentNewAMFUENGAPID int = 9
	UEContextModificationRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest int = 10
	UEContextModificationRequestIEsTypeValuePresentNewGUAMI int = 11
	UEContextModificationRequestIEsTypeValuePresentCNAssistedRANTuning int = 12
	UEContextModificationRequestIEsTypeValuePresentSRVCCOperationPossible int = 13
)

type UEContextModificationRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RANPagingPriority *RANPagingPriority `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	SecurityKey *SecurityKey `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityKey"`
	IndexToRFSP *IndexToRFSP `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDIndexToRFSP"`
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	UESecurityCapabilities *UESecurityCapabilities `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	EmergencyFallbackIndicator *EmergencyFallbackIndicator `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDEmergencyFallbackIndicator"`
	NewAMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewAMFUENGAPID"`
	RRCInactiveTransitionReportRequest *RRCInactiveTransitionReportRequest `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	NewGUAMI *GUAMI `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewGUAMI"`
	CNAssistedRANTuning *CNAssistedRANTuning `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible *SRVCCOperationPossible `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type UEContextModificationResponse struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationResponseIEs
}

type ProtocolIEContainerUEContextModificationResponseIEs struct {
	List []UEContextModificationResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextModificationResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UEContextModificationResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationResponseIEsTypeValuePresentAMFUENGAPID int = 0
	UEContextModificationResponseIEsTypeValuePresentRANUENGAPID int = 1
	UEContextModificationResponseIEsTypeValuePresentRRCState int = 2
	UEContextModificationResponseIEsTypeValuePresentUserLocationInformation int = 3
	UEContextModificationResponseIEsTypeValuePresentCriticalityDiagnostics int = 4
)

type UEContextModificationResponseIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RRCState *RRCState `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCState"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type UEContextModificationFailure struct {
	ProtocolIEs ProtocolIEContainerUEContextModificationFailureIEs
}

type ProtocolIEContainerUEContextModificationFailureIEs struct {
	List []UEContextModificationFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UEContextModificationFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UEContextModificationFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEContextModificationFailureIEsTypeValuePresentAMFUENGAPID int = 0
	UEContextModificationFailureIEsTypeValuePresentRANUENGAPID int = 1
	UEContextModificationFailureIEsTypeValuePresentCause int = 2
	UEContextModificationFailureIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type UEContextModificationFailureIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue RRCInactiveTransitionReportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RRCInactiveTransitionReportIEsTypeValuePresentAMFUENGAPID int = 0
	RRCInactiveTransitionReportIEsTypeValuePresentRANUENGAPID int = 1
	RRCInactiveTransitionReportIEsTypeValuePresentRRCState int = 2
	RRCInactiveTransitionReportIEsTypeValuePresentUserLocationInformation int = 3
)

type RRCInactiveTransitionReportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RRCState *RRCState `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCState"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverRequiredIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequiredIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverRequiredIEsTypeValuePresentRANUENGAPID int = 1
	HandoverRequiredIEsTypeValuePresentHandoverType int = 2
	HandoverRequiredIEsTypeValuePresentCause int = 3
	HandoverRequiredIEsTypeValuePresentTargetID int = 4
	HandoverRequiredIEsTypeValuePresentDirectForwardingPathAvailability int = 5
	HandoverRequiredIEsTypeValuePresentPDUSessionResourceListHORqd int = 6
	HandoverRequiredIEsTypeValuePresentSourceToTargetTransparentContainer int = 7
)

type HandoverRequiredIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	HandoverType *HandoverType `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TargetID *TargetID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetID"`
	DirectForwardingPathAvailability *DirectForwardingPathAvailability `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDirectForwardingPathAvailability"`
	PDUSessionResourceListHORqd *PDUSessionResourceListHORqd `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceListHORqd"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverCommandIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCommandIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverCommandIEsTypeValuePresentRANUENGAPID int = 1
	HandoverCommandIEsTypeValuePresentHandoverType int = 2
	HandoverCommandIEsTypeValuePresentNASSecurityParametersFromNGRAN int = 3
	HandoverCommandIEsTypeValuePresentPDUSessionResourceHandoverList int = 4
	HandoverCommandIEsTypeValuePresentPDUSessionResourceToReleaseListHOCmd int = 5
	HandoverCommandIEsTypeValuePresentTargetToSourceTransparentContainer int = 6
	HandoverCommandIEsTypeValuePresentCriticalityDiagnostics int = 7
)

type HandoverCommandIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	HandoverType *HandoverType `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	NASSecurityParametersFromNGRAN *NASSecurityParametersFromNGRAN `vht:"Presence:PresenceConditional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASSecurityParametersFromNGRAN"`
	PDUSessionResourceHandoverList *PDUSessionResourceHandoverList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceHandoverList"`
	PDUSessionResourceToReleaseListHOCmd *PDUSessionResourceToReleaseListHOCmd `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceToReleaseListHOCmd"`
	TargetToSourceTransparentContainer *TargetToSourceTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetToSourceTransparentContainer"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverPreparationFailure struct {
	ProtocolIEs ProtocolIEContainerHandoverPreparationFailureIEs
}

type ProtocolIEContainerHandoverPreparationFailureIEs struct {
	List []HandoverPreparationFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverPreparationFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverPreparationFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverPreparationFailureIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverPreparationFailureIEsTypeValuePresentRANUENGAPID int = 1
	HandoverPreparationFailureIEsTypeValuePresentCause int = 2
	HandoverPreparationFailureIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type HandoverPreparationFailureIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequestIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverRequestIEsTypeValuePresentHandoverType int = 1
	HandoverRequestIEsTypeValuePresentCause int = 2
	HandoverRequestIEsTypeValuePresentUEAggregateMaximumBitRate int = 3
	HandoverRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive int = 4
	HandoverRequestIEsTypeValuePresentUESecurityCapabilities int = 5
	HandoverRequestIEsTypeValuePresentSecurityContext int = 6
	HandoverRequestIEsTypeValuePresentNewSecurityContextInd int = 7
	HandoverRequestIEsTypeValuePresentNASC int = 8
	HandoverRequestIEsTypeValuePresentPDUSessionResourceSetupListHOReq int = 9
	HandoverRequestIEsTypeValuePresentAllowedNSSAI int = 10
	HandoverRequestIEsTypeValuePresentTraceActivation int = 11
	HandoverRequestIEsTypeValuePresentMaskedIMEISV int = 12
	HandoverRequestIEsTypeValuePresentSourceToTargetTransparentContainer int = 13
	HandoverRequestIEsTypeValuePresentMobilityRestrictionList int = 14
	HandoverRequestIEsTypeValuePresentLocationReportingRequestType int = 15
	HandoverRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest int = 16
	HandoverRequestIEsTypeValuePresentGUAMI int = 17
	HandoverRequestIEsTypeValuePresentRedirectionVoiceFallback int = 18
	HandoverRequestIEsTypeValuePresentCNAssistedRANTuning int = 19
	HandoverRequestIEsTypeValuePresentSRVCCOperationPossible int = 20
)

type HandoverRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	HandoverType *HandoverType `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDHandoverType"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	UESecurityCapabilities *UESecurityCapabilities `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityContext *SecurityContext `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityContext"`
	NewSecurityContextInd *NewSecurityContextInd `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewSecurityContextInd"`
	NASC *NASPDU `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASC"`
	PDUSessionResourceSetupListHOReq *PDUSessionResourceSetupListHOReq `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceSetupListHOReq"`
	AllowedNSSAI *AllowedNSSAI `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	TraceActivation *TraceActivation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTraceActivation"`
	MaskedIMEISV *MaskedIMEISV `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMaskedIMEISV"`
	SourceToTargetTransparentContainer *SourceToTargetTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSourceToTargetTransparentContainer"`
	MobilityRestrictionList *MobilityRestrictionList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMobilityRestrictionList"`
	LocationReportingRequestType *LocationReportingRequestType `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
	RRCInactiveTransitionReportRequest *RRCInactiveTransitionReportRequest `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	GUAMI *GUAMI `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGUAMI"`
	RedirectionVoiceFallback *RedirectionVoiceFallback `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRedirectionVoiceFallback"`
	CNAssistedRANTuning *CNAssistedRANTuning `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible *SRVCCOperationPossible `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type HandoverRequestAcknowledge struct {
	ProtocolIEs ProtocolIEContainerHandoverRequestAcknowledgeIEs
}

type ProtocolIEContainerHandoverRequestAcknowledgeIEs struct {
	List []HandoverRequestAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverRequestAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverRequestAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverRequestAcknowledgeIEsTypeValuePresentRANUENGAPID int = 1
	HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceAdmittedList int = 2
	HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceFailedToSetupListHOAck int = 3
	HandoverRequestAcknowledgeIEsTypeValuePresentTargetToSourceTransparentContainer int = 4
	HandoverRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics int = 5
)

type HandoverRequestAcknowledgeIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceAdmittedList *PDUSessionResourceAdmittedList `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceAdmittedList"`
	PDUSessionResourceFailedToSetupListHOAck *PDUSessionResourceFailedToSetupListHOAck `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck"`
	TargetToSourceTransparentContainer *TargetToSourceTransparentContainer `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetToSourceTransparentContainer"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverFailure struct {
	ProtocolIEs ProtocolIEContainerHandoverFailureIEs
}

type ProtocolIEContainerHandoverFailureIEs struct {
	List []HandoverFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverFailureIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverFailureIEsTypeValuePresentCause int = 1
	HandoverFailureIEsTypeValuePresentCriticalityDiagnostics int = 2
)

type HandoverFailureIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverNotifyIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverNotifyIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverNotifyIEsTypeValuePresentRANUENGAPID int = 1
	HandoverNotifyIEsTypeValuePresentUserLocationInformation int = 2
)

type HandoverNotifyIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PathSwitchRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestIEsTypeValuePresentRANUENGAPID int = 0
	PathSwitchRequestIEsTypeValuePresentSourceAMFUENGAPID int = 1
	PathSwitchRequestIEsTypeValuePresentUserLocationInformation int = 2
	PathSwitchRequestIEsTypeValuePresentUESecurityCapabilities int = 3
	PathSwitchRequestIEsTypeValuePresentPDUSessionResourceToBeSwitchedDLList int = 4
	PathSwitchRequestIEsTypeValuePresentPDUSessionResourceFailedToSetupListPSReq int = 5
)

type PathSwitchRequestIEsTypeValue struct {
	Present int
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	SourceAMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSourceAMFUENGAPID"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	UESecurityCapabilities *UESecurityCapabilities `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	PDUSessionResourceToBeSwitchedDLList *PDUSessionResourceToBeSwitchedDLList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionResourceToBeSwitchedDLList"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PathSwitchRequestAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID int = 0
	PathSwitchRequestAcknowledgeIEsTypeValuePresentRANUENGAPID int = 1
	PathSwitchRequestAcknowledgeIEsTypeValuePresentUESecurityCapabilities int = 2
	PathSwitchRequestAcknowledgeIEsTypeValuePresentSecurityContext int = 3
	PathSwitchRequestAcknowledgeIEsTypeValuePresentNewSecurityContextInd int = 4
	PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceSwitchedList int = 5
	PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceReleasedListPSAck int = 6
	PathSwitchRequestAcknowledgeIEsTypeValuePresentAllowedNSSAI int = 7
	PathSwitchRequestAcknowledgeIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive int = 8
	PathSwitchRequestAcknowledgeIEsTypeValuePresentRRCInactiveTransitionReportRequest int = 9
	PathSwitchRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics int = 10
	PathSwitchRequestAcknowledgeIEsTypeValuePresentRedirectionVoiceFallback int = 11
	PathSwitchRequestAcknowledgeIEsTypeValuePresentCNAssistedRANTuning int = 12
	PathSwitchRequestAcknowledgeIEsTypeValuePresentSRVCCOperationPossible int = 13
)

type PathSwitchRequestAcknowledgeIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UESecurityCapabilities *UESecurityCapabilities `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUESecurityCapabilities"`
	SecurityContext *SecurityContext `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityContext"`
	NewSecurityContextInd *NewSecurityContextInd `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNewSecurityContextInd"`
	PDUSessionResourceSwitchedList *PDUSessionResourceSwitchedList `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSwitchedList"`
	PDUSessionResourceReleasedListPSAck *PDUSessionResourceReleasedListPSAck `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListPSAck"`
	AllowedNSSAI *AllowedNSSAI `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCoreNetworkAssistanceInformationForInactive"`
	RRCInactiveTransitionReportRequest *RRCInactiveTransitionReportRequest `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCInactiveTransitionReportRequest"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
	RedirectionVoiceFallback *RedirectionVoiceFallback `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRedirectionVoiceFallback"`
	CNAssistedRANTuning *CNAssistedRANTuning `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCNAssistedRANTuning"`
	SRVCCOperationPossible *SRVCCOperationPossible `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type PathSwitchRequestFailure struct {
	ProtocolIEs ProtocolIEContainerPathSwitchRequestFailureIEs
}

type ProtocolIEContainerPathSwitchRequestFailureIEs struct {
	List []PathSwitchRequestFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PathSwitchRequestFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PathSwitchRequestFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PathSwitchRequestFailureIEsTypeValuePresentAMFUENGAPID int = 0
	PathSwitchRequestFailureIEsTypeValuePresentRANUENGAPID int = 1
	PathSwitchRequestFailureIEsTypeValuePresentPDUSessionResourceReleasedListPSFail int = 2
	PathSwitchRequestFailureIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type PathSwitchRequestFailureIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceReleasedListPSFail *PDUSessionResourceReleasedListPSFail `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceReleasedListPSFail"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type HandoverCancel struct {
	ProtocolIEs ProtocolIEContainerHandoverCancelIEs
}

type ProtocolIEContainerHandoverCancelIEs struct {
	List []HandoverCancelIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverCancelIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverCancelIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCancelIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverCancelIEsTypeValuePresentRANUENGAPID int = 1
	HandoverCancelIEsTypeValuePresentCause int = 2
)

type HandoverCancelIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type HandoverCancelAcknowledge struct {
	ProtocolIEs ProtocolIEContainerHandoverCancelAcknowledgeIEs
}

type ProtocolIEContainerHandoverCancelAcknowledgeIEs struct {
	List []HandoverCancelAcknowledgeIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type HandoverCancelAcknowledgeIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue HandoverCancelAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	HandoverCancelAcknowledgeIEsTypeValuePresentAMFUENGAPID int = 0
	HandoverCancelAcknowledgeIEsTypeValuePresentRANUENGAPID int = 1
	HandoverCancelAcknowledgeIEsTypeValuePresentCriticalityDiagnostics int = 2
)

type HandoverCancelAcknowledgeIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UplinkRANStatusTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID int = 0
	UplinkRANStatusTransferIEsTypeValuePresentRANUENGAPID int = 1
	UplinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer int = 2
)

type UplinkRANStatusTransferIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DownlinkRANStatusTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID int = 0
	DownlinkRANStatusTransferIEsTypeValuePresentRANUENGAPID int = 1
	DownlinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer int = 2
)

type DownlinkRANStatusTransferIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PagingIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PagingIEsTypeValuePresentUEPagingIdentity int = 0
	PagingIEsTypeValuePresentPagingDRX int = 1
	PagingIEsTypeValuePresentTAIListForPaging int = 2
	PagingIEsTypeValuePresentPagingPriority int = 3
	PagingIEsTypeValuePresentUERadioCapabilityForPaging int = 4
	PagingIEsTypeValuePresentPagingOrigin int = 5
	PagingIEsTypeValuePresentAssistanceDataForPaging int = 6
)

type PagingIEsTypeValue struct {
	Present int
	UEPagingIdentity *UEPagingIdentity `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEPagingIdentity"`
	PagingDRX *PagingDRX `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingDRX"`
	TAIListForPaging *TAIListForPaging `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTAIListForPaging"`
	PagingPriority *PagingPriority `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingPriority"`
	UERadioCapabilityForPaging *UERadioCapabilityForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapabilityForPaging"`
	PagingOrigin *PagingOrigin `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPagingOrigin"`
	AssistanceDataForPaging *AssistanceDataForPaging `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAssistanceDataForPaging"`
}

type InitialUEMessage struct {
	ProtocolIEs ProtocolIEContainerInitialUEMessageIEs
}

type ProtocolIEContainerInitialUEMessageIEs struct {
	List []InitialUEMessageIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type InitialUEMessageIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue InitialUEMessageIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	InitialUEMessageIEsTypeValuePresentRANUENGAPID int = 0
	InitialUEMessageIEsTypeValuePresentNASPDU int = 1
	InitialUEMessageIEsTypeValuePresentUserLocationInformation int = 2
	InitialUEMessageIEsTypeValuePresentRRCEstablishmentCause int = 3
	InitialUEMessageIEsTypeValuePresentFiveGSTMSI int = 4
	InitialUEMessageIEsTypeValuePresentAMFSetID int = 5
	InitialUEMessageIEsTypeValuePresentUEContextRequest int = 6
	InitialUEMessageIEsTypeValuePresentAllowedNSSAI int = 7
	InitialUEMessageIEsTypeValuePresentSourceToTargetAMFInformationReroute int = 8
)

type InitialUEMessageIEsTypeValue struct {
	Present int
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NASPDU *NASPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	RRCEstablishmentCause *RRCEstablishmentCause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRRCEstablishmentCause"`
	FiveGSTMSI *FiveGSTMSI `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDFiveGSTMSI"`
	AMFSetID *AMFSetID `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFSetID"`
	UEContextRequest *UEContextRequest `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEContextRequest"`
	AllowedNSSAI *AllowedNSSAI `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DownlinkNASTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkNASTransportIEsTypeValuePresentAMFUENGAPID int = 0
	DownlinkNASTransportIEsTypeValuePresentRANUENGAPID int = 1
	DownlinkNASTransportIEsTypeValuePresentOldAMF int = 2
	DownlinkNASTransportIEsTypeValuePresentRANPagingPriority int = 3
	DownlinkNASTransportIEsTypeValuePresentNASPDU int = 4
	DownlinkNASTransportIEsTypeValuePresentMobilityRestrictionList int = 5
	DownlinkNASTransportIEsTypeValuePresentIndexToRFSP int = 6
	DownlinkNASTransportIEsTypeValuePresentUEAggregateMaximumBitRate int = 7
	DownlinkNASTransportIEsTypeValuePresentAllowedNSSAI int = 8
	DownlinkNASTransportIEsTypeValuePresentSRVCCOperationPossible int = 9
)

type DownlinkNASTransportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	OldAMF *AMFName `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDOldAMF"`
	RANPagingPriority *RANPagingPriority `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANPagingPriority"`
	NASPDU *NASPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
	MobilityRestrictionList *MobilityRestrictionList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDMobilityRestrictionList"`
	IndexToRFSP *IndexToRFSP `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDIndexToRFSP"`
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAggregateMaximumBitRate"`
	AllowedNSSAI *AllowedNSSAI `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
	SRVCCOperationPossible *SRVCCOperationPossible `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSRVCCOperationPossible"`
}

type UplinkNASTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkNASTransportIEs
}

type ProtocolIEContainerUplinkNASTransportIEs struct {
	List []UplinkNASTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkNASTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UplinkNASTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkNASTransportIEsTypeValuePresentAMFUENGAPID int = 0
	UplinkNASTransportIEsTypeValuePresentRANUENGAPID int = 1
	UplinkNASTransportIEsTypeValuePresentNASPDU int = 2
	UplinkNASTransportIEsTypeValuePresentUserLocationInformation int = 3
)

type UplinkNASTransportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NASPDU *NASPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNASPDU"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue NASNonDeliveryIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NASNonDeliveryIndicationIEsTypeValuePresentAMFUENGAPID int = 0
	NASNonDeliveryIndicationIEsTypeValuePresentRANUENGAPID int = 1
	NASNonDeliveryIndicationIEsTypeValuePresentNASPDU int = 2
	NASNonDeliveryIndicationIEsTypeValuePresentCause int = 3
)

type NASNonDeliveryIndicationIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NASPDU *NASPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNASPDU"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type RerouteNASRequest struct {
	ProtocolIEs ProtocolIEContainerRerouteNASRequestIEs
}

type ProtocolIEContainerRerouteNASRequestIEs struct {
	List []RerouteNASRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type RerouteNASRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue RerouteNASRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RerouteNASRequestIEsTypeValuePresentRANUENGAPID int = 0
	RerouteNASRequestIEsTypeValuePresentAMFUENGAPID int = 1
	RerouteNASRequestIEsTypeValuePresentNGAPMessage int = 2
	RerouteNASRequestIEsTypeValuePresentAMFSetID int = 3
	RerouteNASRequestIEsTypeValuePresentAllowedNSSAI int = 4
	RerouteNASRequestIEsTypeValuePresentSourceToTargetAMFInformationReroute int = 5
)

type RerouteNASRequestIEsTypeValue struct {
	Present int
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	NGAPMessage *OctetString `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNGAPMessage"`
	AMFSetID *AMFSetID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFSetID"`
	AllowedNSSAI *AllowedNSSAI `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAllowedNSSAI"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue NGSetupRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupRequestIEsTypeValuePresentGlobalRANNodeID int = 0
	NGSetupRequestIEsTypeValuePresentRANNodeName int = 1
	NGSetupRequestIEsTypeValuePresentSupportedTAList int = 2
	NGSetupRequestIEsTypeValuePresentDefaultPagingDRX int = 3
	NGSetupRequestIEsTypeValuePresentUERetentionInformation int = 4
)

type NGSetupRequestIEsTypeValue struct {
	Present int
	GlobalRANNodeID *GlobalRANNodeID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	RANNodeName *RANNodeName `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANNodeName"`
	SupportedTAList *SupportedTAList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSupportedTAList"`
	DefaultPagingDRX *PagingDRX `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDefaultPagingDRX"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue NGSetupResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupResponseIEsTypeValuePresentAMFName int = 0
	NGSetupResponseIEsTypeValuePresentServedGUAMIList int = 1
	NGSetupResponseIEsTypeValuePresentRelativeAMFCapacity int = 2
	NGSetupResponseIEsTypeValuePresentPLMNSupportList int = 3
	NGSetupResponseIEsTypeValuePresentCriticalityDiagnostics int = 4
	NGSetupResponseIEsTypeValuePresentUERetentionInformation int = 5
)

type NGSetupResponseIEsTypeValue struct {
	Present int
	AMFName *AMFName `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFName"`
	ServedGUAMIList *ServedGUAMIList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDServedGUAMIList"`
	RelativeAMFCapacity *RelativeAMFCapacity `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRelativeAMFCapacity"`
	PLMNSupportList *PLMNSupportList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPLMNSupportList"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue NGSetupFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGSetupFailureIEsTypeValuePresentCause int = 0
	NGSetupFailureIEsTypeValuePresentTimeToWait int = 1
	NGSetupFailureIEsTypeValuePresentCriticalityDiagnostics int = 2
)

type NGSetupFailureIEsTypeValue struct {
	Present int
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait *TimeToWait `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue RANConfigurationUpdateIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateIEsTypeValuePresentRANNodeName int = 0
	RANConfigurationUpdateIEsTypeValuePresentSupportedTAList int = 1
	RANConfigurationUpdateIEsTypeValuePresentDefaultPagingDRX int = 2
	RANConfigurationUpdateIEsTypeValuePresentGlobalRANNodeID int = 3
	RANConfigurationUpdateIEsTypeValuePresentNGRANTNLAssociationToRemoveList int = 4
)

type RANConfigurationUpdateIEsTypeValue struct {
	Present int
	RANNodeName *RANNodeName `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANNodeName"`
	SupportedTAList *SupportedTAList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSupportedTAList"`
	DefaultPagingDRX *PagingDRX `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDefaultPagingDRX"`
	GlobalRANNodeID *GlobalRANNodeID `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue RANConfigurationUpdateAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics int = 0
)

type RANConfigurationUpdateAcknowledgeIEsTypeValue struct {
	Present int
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue RANConfigurationUpdateFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	RANConfigurationUpdateFailureIEsTypeValuePresentCause int = 0
	RANConfigurationUpdateFailureIEsTypeValuePresentTimeToWait int = 1
	RANConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics int = 2
)

type RANConfigurationUpdateFailureIEsTypeValue struct {
	Present int
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait *TimeToWait `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue AMFConfigurationUpdateIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateIEsTypeValuePresentAMFName int = 0
	AMFConfigurationUpdateIEsTypeValuePresentServedGUAMIList int = 1
	AMFConfigurationUpdateIEsTypeValuePresentRelativeAMFCapacity int = 2
	AMFConfigurationUpdateIEsTypeValuePresentPLMNSupportList int = 3
	AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToAddList int = 4
	AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToRemoveList int = 5
	AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToUpdateList int = 6
)

type AMFConfigurationUpdateIEsTypeValue struct {
	Present int
	AMFName *AMFName `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFName"`
	ServedGUAMIList *ServedGUAMIList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDServedGUAMIList"`
	RelativeAMFCapacity *RelativeAMFCapacity `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRelativeAMFCapacity"`
	PLMNSupportList *PLMNSupportList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPLMNSupportList"`
	AMFTNLAssociationToAddList *AMFTNLAssociationToAddList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationToAddList"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue AMFConfigurationUpdateAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationSetupList int = 0
	AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationFailedToSetupList int = 1
	AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics int = 2
)

type AMFConfigurationUpdateAcknowledgeIEsTypeValue struct {
	Present int
	AMFTNLAssociationSetupList *AMFTNLAssociationSetupList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationSetupList"`
	AMFTNLAssociationFailedToSetupList *TNLAssociationList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTNLAssociationFailedToSetupList"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type AMFConfigurationUpdateFailure struct {
	ProtocolIEs ProtocolIEContainerAMFConfigurationUpdateFailureIEs
}

type ProtocolIEContainerAMFConfigurationUpdateFailureIEs struct {
	List []AMFConfigurationUpdateFailureIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type AMFConfigurationUpdateFailureIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue AMFConfigurationUpdateFailureIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFConfigurationUpdateFailureIEsTypeValuePresentCause int = 0
	AMFConfigurationUpdateFailureIEsTypeValuePresentTimeToWait int = 1
	AMFConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics int = 2
)

type AMFConfigurationUpdateFailureIEsTypeValue struct {
	Present int
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
	TimeToWait *TimeToWait `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDTimeToWait"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue AMFStatusIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFStatusIndicationIEsTypeValuePresentUnavailableGUAMIList int = 0
)

type AMFStatusIndicationIEsTypeValue struct {
	Present int
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue NGResetIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGResetIEsTypeValuePresentCause int = 0
	NGResetIEsTypeValuePresentResetType int = 1
)

type NGResetIEsTypeValue struct {
	Present int
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue NGResetAcknowledgeIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGResetAcknowledgeIEsTypeValuePresentUEAssociatedLogicalNGConnectionList int = 0
	NGResetAcknowledgeIEsTypeValuePresentCriticalityDiagnostics int = 1
)

type NGResetAcknowledgeIEsTypeValue struct {
	Present int
	UEAssociatedLogicalNGConnectionList *UEAssociatedLogicalNGConnectionList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEAssociatedLogicalNGConnectionList"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type ErrorIndication struct {
	ProtocolIEs ProtocolIEContainerErrorIndicationIEs
}

type ProtocolIEContainerErrorIndicationIEs struct {
	List []ErrorIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type ErrorIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue ErrorIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	ErrorIndicationIEsTypeValuePresentAMFUENGAPID int = 0
	ErrorIndicationIEsTypeValuePresentRANUENGAPID int = 1
	ErrorIndicationIEsTypeValuePresentCause int = 2
	ErrorIndicationIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type ErrorIndicationIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause *Cause `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue OverloadStartIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	OverloadStartIEsTypeValuePresentAMFOverloadResponse int = 0
	OverloadStartIEsTypeValuePresentAMFTrafficLoadReductionIndication int = 1
	OverloadStartIEsTypeValuePresentOverloadStartNSSAIList int = 2
)

type OverloadStartIEsTypeValue struct {
	Present int
	AMFOverloadResponse *OverloadResponse `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFOverloadResponse"`
	AMFTrafficLoadReductionIndication *TrafficLoadReductionIndication `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFTrafficLoadReductionIndication"`
	OverloadStartNSSAIList *OverloadStartNSSAIList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDOverloadStartNSSAIList"`
}

type OverloadStop struct {
	ProtocolIEs ProtocolIEContainerOverloadStopIEs
}

type ProtocolIEContainerOverloadStopIEs struct {
	List []OverloadStopIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type OverloadStopIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue OverloadStopIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	OverloadStopIEsTypeValuePresentNothing int = 0
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UplinkRANConfigurationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferUL int = 0
	UplinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferUL int = 1
)

type UplinkRANConfigurationTransferIEsTypeValue struct {
	Present int
	SONConfigurationTransferUL *SONConfigurationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSONConfigurationTransferUL"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DownlinkRANConfigurationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferDL int = 0
	DownlinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferDL int = 1
)

type DownlinkRANConfigurationTransferIEsTypeValue struct {
	Present int
	SONConfigurationTransferDL *SONConfigurationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDSONConfigurationTransferDL"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue WriteReplaceWarningRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	WriteReplaceWarningRequestIEsTypeValuePresentMessageIdentifier int = 0
	WriteReplaceWarningRequestIEsTypeValuePresentSerialNumber int = 1
	WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaList int = 2
	WriteReplaceWarningRequestIEsTypeValuePresentRepetitionPeriod int = 3
	WriteReplaceWarningRequestIEsTypeValuePresentNumberOfBroadcastsRequested int = 4
	WriteReplaceWarningRequestIEsTypeValuePresentWarningType int = 5
	WriteReplaceWarningRequestIEsTypeValuePresentWarningSecurityInfo int = 6
	WriteReplaceWarningRequestIEsTypeValuePresentDataCodingScheme int = 7
	WriteReplaceWarningRequestIEsTypeValuePresentWarningMessageContents int = 8
	WriteReplaceWarningRequestIEsTypeValuePresentConcurrentWarningMessageInd int = 9
	WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaCoordinates int = 10
)

type WriteReplaceWarningRequestIEsTypeValue struct {
	Present int
	MessageIdentifier *MessageIdentifier `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber *SerialNumber `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	WarningAreaList *WarningAreaList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningAreaList"`
	RepetitionPeriod *RepetitionPeriod `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRepetitionPeriod"`
	NumberOfBroadcastsRequested *NumberOfBroadcastsRequested `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNumberOfBroadcastsRequested"`
	WarningType *WarningType `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningType"`
	WarningSecurityInfo *WarningSecurityInfo `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningSecurityInfo"`
	DataCodingScheme *DataCodingScheme `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDataCodingScheme"`
	WarningMessageContents *WarningMessageContents `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningMessageContents"`
	ConcurrentWarningMessageInd *ConcurrentWarningMessageInd `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDConcurrentWarningMessageInd"`
	WarningAreaCoordinates *WarningAreaCoordinates `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningAreaCoordinates"`
}

type WriteReplaceWarningResponse struct {
	ProtocolIEs ProtocolIEContainerWriteReplaceWarningResponseIEs
}

type ProtocolIEContainerWriteReplaceWarningResponseIEs struct {
	List []WriteReplaceWarningResponseIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type WriteReplaceWarningResponseIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue WriteReplaceWarningResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	WriteReplaceWarningResponseIEsTypeValuePresentMessageIdentifier int = 0
	WriteReplaceWarningResponseIEsTypeValuePresentSerialNumber int = 1
	WriteReplaceWarningResponseIEsTypeValuePresentBroadcastCompletedAreaList int = 2
	WriteReplaceWarningResponseIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type WriteReplaceWarningResponseIEsTypeValue struct {
	Present int
	MessageIdentifier *MessageIdentifier `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber *SerialNumber `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	BroadcastCompletedAreaList *BroadcastCompletedAreaList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDBroadcastCompletedAreaList"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PWSCancelRequest struct {
	ProtocolIEs ProtocolIEContainerPWSCancelRequestIEs
}

type ProtocolIEContainerPWSCancelRequestIEs struct {
	List []PWSCancelRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PWSCancelRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PWSCancelRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSCancelRequestIEsTypeValuePresentMessageIdentifier int = 0
	PWSCancelRequestIEsTypeValuePresentSerialNumber int = 1
	PWSCancelRequestIEsTypeValuePresentWarningAreaList int = 2
	PWSCancelRequestIEsTypeValuePresentCancelAllWarningMessages int = 3
)

type PWSCancelRequestIEsTypeValue struct {
	Present int
	MessageIdentifier *MessageIdentifier `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber *SerialNumber `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	WarningAreaList *WarningAreaList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDWarningAreaList"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PWSCancelResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSCancelResponseIEsTypeValuePresentMessageIdentifier int = 0
	PWSCancelResponseIEsTypeValuePresentSerialNumber int = 1
	PWSCancelResponseIEsTypeValuePresentBroadcastCancelledAreaList int = 2
	PWSCancelResponseIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type PWSCancelResponseIEsTypeValue struct {
	Present int
	MessageIdentifier *MessageIdentifier `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDMessageIdentifier"`
	SerialNumber *SerialNumber `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSerialNumber"`
	BroadcastCancelledAreaList *BroadcastCancelledAreaList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDBroadcastCancelledAreaList"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PWSRestartIndication struct {
	ProtocolIEs ProtocolIEContainerPWSRestartIndicationIEs
}

type ProtocolIEContainerPWSRestartIndicationIEs struct {
	List []PWSRestartIndicationIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type PWSRestartIndicationIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PWSRestartIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSRestartIndicationIEsTypeValuePresentCellIDListForRestart int = 0
	PWSRestartIndicationIEsTypeValuePresentGlobalRANNodeID int = 1
	PWSRestartIndicationIEsTypeValuePresentTAIListForRestart int = 2
	PWSRestartIndicationIEsTypeValuePresentEmergencyAreaIDListForRestart int = 3
)

type PWSRestartIndicationIEsTypeValue struct {
	Present int
	CellIDListForRestart *CellIDListForRestart `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDCellIDListForRestart"`
	GlobalRANNodeID *GlobalRANNodeID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
	TAIListForRestart *TAIListForRestart `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTAIListForRestart"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PWSFailureIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSFailureIndicationIEsTypeValuePresentPWSFailedCellIDList int = 0
	PWSFailureIndicationIEsTypeValuePresentGlobalRANNodeID int = 1
)

type PWSFailureIndicationIEsTypeValue struct {
	Present int
	PWSFailedCellIDList *PWSFailedCellIDList `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPWSFailedCellIDList"`
	GlobalRANNodeID *GlobalRANNodeID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDGlobalRANNodeID"`
}

type DownlinkUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DownlinkUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID int = 0
	DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID int = 1
	DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID int = 2
	DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU int = 3
)

type DownlinkUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU *NRPPaPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type UplinkUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs struct {
	List []UplinkUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UplinkUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID int = 0
	UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID int = 1
	UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID int = 2
	UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU int = 3
)

type UplinkUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU *NRPPaPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type DownlinkNonUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs struct {
	List []DownlinkNonUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DownlinkNonUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID int = 0
	DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU int = 1
)

type DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present int
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU *NRPPaPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type UplinkNonUEAssociatedNRPPaTransport struct {
	ProtocolIEs ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs
}

type ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs struct {
	List []UplinkNonUEAssociatedNRPPaTransportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkNonUEAssociatedNRPPaTransportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UplinkNonUEAssociatedNRPPaTransportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID int = 0
	UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU int = 1
)

type UplinkNonUEAssociatedNRPPaTransportIEsTypeValue struct {
	Present int
	RoutingID *RoutingID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRoutingID"`
	NRPPaPDU *NRPPaPDU `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNRPPaPDU"`
}

type TraceStart struct {
	ProtocolIEs ProtocolIEContainerTraceStartIEs
}

type ProtocolIEContainerTraceStartIEs struct {
	List []TraceStartIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type TraceStartIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue TraceStartIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	TraceStartIEsTypeValuePresentAMFUENGAPID int = 0
	TraceStartIEsTypeValuePresentRANUENGAPID int = 1
	TraceStartIEsTypeValuePresentTraceActivation int = 2
)

type TraceStartIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue TraceFailureIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	TraceFailureIndicationIEsTypeValuePresentAMFUENGAPID int = 0
	TraceFailureIndicationIEsTypeValuePresentRANUENGAPID int = 1
	TraceFailureIndicationIEsTypeValuePresentNGRANTraceID int = 2
	TraceFailureIndicationIEsTypeValuePresentCause int = 3
)

type TraceFailureIndicationIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NGRANTraceID *NGRANTraceID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANTraceID"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type DeactivateTrace struct {
	ProtocolIEs ProtocolIEContainerDeactivateTraceIEs
}

type ProtocolIEContainerDeactivateTraceIEs struct {
	List []DeactivateTraceIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type DeactivateTraceIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DeactivateTraceIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DeactivateTraceIEsTypeValuePresentAMFUENGAPID int = 0
	DeactivateTraceIEsTypeValuePresentRANUENGAPID int = 1
	DeactivateTraceIEsTypeValuePresentNGRANTraceID int = 2
)

type DeactivateTraceIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue CellTrafficTraceIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	CellTrafficTraceIEsTypeValuePresentAMFUENGAPID int = 0
	CellTrafficTraceIEsTypeValuePresentRANUENGAPID int = 1
	CellTrafficTraceIEsTypeValuePresentNGRANTraceID int = 2
	CellTrafficTraceIEsTypeValuePresentNGRANCGI int = 3
	CellTrafficTraceIEsTypeValuePresentTraceCollectionEntityIPAddress int = 4
)

type CellTrafficTraceIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	NGRANTraceID *NGRANTraceID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANTraceID"`
	NGRANCGI *NGRANCGI `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDNGRANCGI"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue LocationReportingControlIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportingControlIEsTypeValuePresentAMFUENGAPID int = 0
	LocationReportingControlIEsTypeValuePresentRANUENGAPID int = 1
	LocationReportingControlIEsTypeValuePresentLocationReportingRequestType int = 2
)

type LocationReportingControlIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue LocationReportingFailureIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportingFailureIndicationIEsTypeValuePresentAMFUENGAPID int = 0
	LocationReportingFailureIndicationIEsTypeValuePresentRANUENGAPID int = 1
	LocationReportingFailureIndicationIEsTypeValuePresentCause int = 2
)

type LocationReportingFailureIndicationIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	Cause *Cause `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCause"`
}

type LocationReport struct {
	ProtocolIEs ProtocolIEContainerLocationReportIEs
}

type ProtocolIEContainerLocationReportIEs struct {
	List []LocationReportIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type LocationReportIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue LocationReportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LocationReportIEsTypeValuePresentAMFUENGAPID int = 0
	LocationReportIEsTypeValuePresentRANUENGAPID int = 1
	LocationReportIEsTypeValuePresentUserLocationInformation int = 2
	LocationReportIEsTypeValuePresentUEPresenceInAreaOfInterestList int = 3
	LocationReportIEsTypeValuePresentLocationReportingRequestType int = 4
)

type LocationReportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
	UEPresenceInAreaOfInterestList *UEPresenceInAreaOfInterestList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUEPresenceInAreaOfInterestList"`
	LocationReportingRequestType *LocationReportingRequestType `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDLocationReportingRequestType"`
}

type UETNLABindingReleaseRequest struct {
	ProtocolIEs ProtocolIEContainerUETNLABindingReleaseRequestIEs
}

type ProtocolIEContainerUETNLABindingReleaseRequestIEs struct {
	List []UETNLABindingReleaseRequestIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UETNLABindingReleaseRequestIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UETNLABindingReleaseRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UETNLABindingReleaseRequestIEsTypeValuePresentAMFUENGAPID int = 0
	UETNLABindingReleaseRequestIEsTypeValuePresentRANUENGAPID int = 1
)

type UETNLABindingReleaseRequestIEsTypeValue struct {
	Present int
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UERadioCapabilityInfoIndicationIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityInfoIndicationIEsTypeValuePresentAMFUENGAPID int = 0
	UERadioCapabilityInfoIndicationIEsTypeValuePresentRANUENGAPID int = 1
	UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapability int = 2
	UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapabilityForPaging int = 3
)

type UERadioCapabilityInfoIndicationIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	UERadioCapability *UERadioCapability `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUERadioCapability"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UERadioCapabilityCheckRequestIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityCheckRequestIEsTypeValuePresentAMFUENGAPID int = 0
	UERadioCapabilityCheckRequestIEsTypeValuePresentRANUENGAPID int = 1
	UERadioCapabilityCheckRequestIEsTypeValuePresentUERadioCapability int = 2
)

type UERadioCapabilityCheckRequestIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UERadioCapabilityCheckResponseIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UERadioCapabilityCheckResponseIEsTypeValuePresentAMFUENGAPID int = 0
	UERadioCapabilityCheckResponseIEsTypeValuePresentRANUENGAPID int = 1
	UERadioCapabilityCheckResponseIEsTypeValuePresentIMSVoiceSupportIndicator int = 2
	UERadioCapabilityCheckResponseIEsTypeValuePresentCriticalityDiagnostics int = 3
)

type UERadioCapabilityCheckResponseIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	IMSVoiceSupportIndicator *IMSVoiceSupportIndicator `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDIMSVoiceSupportIndicator"`
	CriticalityDiagnostics *CriticalityDiagnostics `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCriticalityDiagnostics"`
}

type PrivateMessage struct {
	PrivateIEs PrivateIEContainerPrivateMessageIEs
}

type PrivateIEContainerPrivateMessageIEs struct {
	List []PrivateMessageIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type PrivateMessageIEs struct {
	PrivateIEID PrivateIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue PrivateMessageIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PrivateMessageIEsTypeValuePresentNothing int = 0
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue SecondaryRATDataUsageReportIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	SecondaryRATDataUsageReportIEsTypeValuePresentAMFUENGAPID int = 0
	SecondaryRATDataUsageReportIEsTypeValuePresentRANUENGAPID int = 1
	SecondaryRATDataUsageReportIEsTypeValuePresentPDUSessionResourceSecondaryRATUsageList int = 2
	SecondaryRATDataUsageReportIEsTypeValuePresentHandoverFlag int = 3
	SecondaryRATDataUsageReportIEsTypeValuePresentUserLocationInformation int = 4
)

type SecondaryRATDataUsageReportIEsTypeValue struct {
	Present int
	AMFUENGAPID *AMFUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDAMFUENGAPID"`
	RANUENGAPID *RANUENGAPID `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRANUENGAPID"`
	PDUSessionResourceSecondaryRATUsageList *PDUSessionResourceSecondaryRATUsageList `vht:"Presence:PresenceMandatory,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDPDUSessionResourceSecondaryRATUsageList"`
	HandoverFlag *HandoverFlag `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDHandoverFlag"`
	UserLocationInformation *UserLocationInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDUserLocationInformation"`
}

type UplinkRIMInformationTransfer struct {
	ProtocolIEs ProtocolIEContainerUplinkRIMInformationTransferIEs
}

type ProtocolIEContainerUplinkRIMInformationTransferIEs struct {
	List []UplinkRIMInformationTransferIEs `vht:"valueMin:0,valueMax:MaxPrivateIEs"`
}

type UplinkRIMInformationTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue UplinkRIMInformationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UplinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer int = 0
)

type UplinkRIMInformationTransferIEsTypeValue struct {
	Present int
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
	Criticality Criticality `vht:"Reference:ProtocolIEID"`
	TypeValue DownlinkRIMInformationTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DownlinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer int = 0
)

type DownlinkRIMInformationTransferIEsTypeValue struct {
	Present int
	RIMInformationTransfer *RIMInformationTransfer `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDRIMInformationTransfer"`
}
