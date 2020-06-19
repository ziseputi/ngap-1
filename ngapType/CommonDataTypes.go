// Created By HaoDHH-245789 VHT2020

package ngapType

// ASN1 TYPE

// BIT STRING

// BitString is for an ASN.1 BIT STRING type, BitLength means the effective bits.
type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength uint64 // length in bits.
}

// OCTET STRING

// OctetString is for an ASN.1 OCTET STRING type
type OctetString []byte

// OBJECT IDENTIFIER

// ObjectIdentifier is for an ASN.1 OBJECT IDENTIFIER type
type ObjectIdentifier []byte

// ENUMERATED

// An Enumerated is represented as a plain uint64.
type Enumerated uint64

// INTEGER

// An Integer is represented as a plain int64.
type Integer int64

// PrintableString

// A PrintableString is represented as a plain string.
type PrintableString string

// All types, universal class

// | Name                                           | Value encodings                | Tag number
// |                                                |                                | Decimal | Hexadecimal
// | End-of-Content (EOC)                           | Primitive                      | 0       | 0
// | BOOLEAN                                        | Primitive                      | 1       | 1
// | INTEGER                                        | Primitive                      | 2       | 2
// | BIT STRING                                     | Both                           | 3       | 3
// | OCTET STRING                                   | Both                           | 4       | 4
// | NULL                                           | Primitive                      | 5       | 5
// | OBJECT IDENTIFIER                              | Primitive                      | 6       | 6
// | Object Descriptor                              | Both                           | 7       | 7
// | EXTERNAL                                       | Constructed                    | 8       | 8
// | REAL (float)                                   | Primitive                      | 9       | 9
// | ENUMERATED                                     | Primitive                      | 10      | A
// | EMBEDDED PDV                                   | Constructed                    | 11      | B
// | UTF8String                                     | Both                           | 12      | C
// | RELATIVE-OID                                   | Primitive                      | 13      | D
// | TIME                                           | Primitive                      | 14      | E
// | Reserved                                       |                                | 15      | F
// | SEQUENCE and SEQUENCE OF                       | Constructed                    | 16      | 10
// | SET and SET OF                                 | Constructed                    | 17      | 11
// | NumericString                                  | Both                           | 18      | 12
// | PrintableString                                | Both                           | 19      | 13
// | T61String                                      | Both                           | 20      | 14
// | VideotexString                                 | Both                           | 21      | 15
// | IA5String                                      | Both                           | 22      | 16
// | UTCTime                                        | Both                           | 23      | 17
// | GeneralizedTime                                | Both                           | 24      | 18
// | GraphicString                                  | Both                           | 25      | 19
// | VisibleString                                  | Both                           | 26      | 1A
// | GeneralString                                  | Both                           | 27      | 1B
// | UniversalString                                | Both                           | 28      | 1C
// | CHARACTER STRING                               | Both                           | 29      | 1D
// | BMPString                                      | Both                           | 30      | 1E
// | DATE                                           | Primitive                      | 31      | 1F
// | TIME-OF-DAY                                    | Primitive                      | 32      | 20
// | DATE-TIME                                      | Primitive                      | 33      | 21
// | DURATION                                       | Primitive                      | 34      | 22
// | OID-IRI                                        | Primitive                      | 35      | 23
// | RELATIVE-OID-IRI                               | Primitive                      | 36      | 24

// COMMONDATATYPE

const (
	CriticalityReject Enumerated = 0
	CriticalityIgnore Enumerated = 1
	CriticalityNotify Enumerated = 2
)

type Criticality struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	PresenceOptional    Enumerated = 0
	PresenceConditional Enumerated = 1
	PresenceMandatory   Enumerated = 2
)

type Presence struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	PrivateIEIDPresentLocal  int = 0
	PrivateIEIDPresentGlobal int = 1
)

type PrivateIEID struct {
	Present int
	Local  *Integer `vht:"valueMin:0,valueMax:65535"`
	Global *ObjectIdentifier
}

const (
	ProcedureCodeAMFConfigurationUpdate                Integer = 0
	ProcedureCodeAMFStatusIndication                   Integer = 1
	ProcedureCodeCellTrafficTrace                      Integer = 2
	ProcedureCodeDeactivateTrace                       Integer = 3
	ProcedureCodeDownlinkNASTransport                  Integer = 4
	ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport Integer = 5
	ProcedureCodeDownlinkRANConfigurationTransfer      Integer = 6
	ProcedureCodeDownlinkRANStatusTransfer             Integer = 7
	ProcedureCodeDownlinkUEAssociatedNRPPaTransport    Integer = 8
	ProcedureCodeErrorIndication                       Integer = 9
	ProcedureCodeHandoverCancel                        Integer = 10
	ProcedureCodeHandoverNotification                  Integer = 11
	ProcedureCodeHandoverPreparation                   Integer = 12
	ProcedureCodeHandoverResourceAllocation            Integer = 13
	ProcedureCodeInitialContextSetup                   Integer = 14
	ProcedureCodeInitialUEMessage                      Integer = 15
	ProcedureCodeLocationReportingControl              Integer = 16
	ProcedureCodeLocationReportingFailureIndication    Integer = 17
	ProcedureCodeLocationReport                        Integer = 18
	ProcedureCodeNASNonDeliveryIndication              Integer = 19
	ProcedureCodeNGReset                               Integer = 20
	ProcedureCodeNGSetup                               Integer = 21
	ProcedureCodeOverloadStart                         Integer = 22
	ProcedureCodeOverloadStop                          Integer = 23
	ProcedureCodePaging                                Integer = 24
	ProcedureCodePathSwitchRequest                     Integer = 25
	ProcedureCodePDUSessionResourceModify              Integer = 26
	ProcedureCodePDUSessionResourceModifyIndication    Integer = 27
	ProcedureCodePDUSessionResourceRelease             Integer = 28
	ProcedureCodePDUSessionResourceSetup               Integer = 29
	ProcedureCodePDUSessionResourceNotify              Integer = 30
	ProcedureCodePrivateMessage                        Integer = 31
	ProcedureCodePWSCancel                             Integer = 32
	ProcedureCodePWSFailureIndication                  Integer = 33
	ProcedureCodePWSRestartIndication                  Integer = 34
	ProcedureCodeRANConfigurationUpdate                Integer = 35
	ProcedureCodeRerouteNASRequest                     Integer = 36
	ProcedureCodeRRCInactiveTransitionReport           Integer = 37
	ProcedureCodeTraceFailureIndication                Integer = 38
	ProcedureCodeTraceStart                            Integer = 39
	ProcedureCodeUEContextModification                 Integer = 40
	ProcedureCodeUEContextRelease                      Integer = 41
	ProcedureCodeUEContextReleaseRequest               Integer = 42
	ProcedureCodeUERadioCapabilityCheck                Integer = 43
	ProcedureCodeUERadioCapabilityInfoIndication       Integer = 44
	ProcedureCodeUETNLABindingRelease                  Integer = 45
	ProcedureCodeUplinkNASTransport                    Integer = 46
	ProcedureCodeUplinkNonUEAssociatedNRPPaTransport   Integer = 47
	ProcedureCodeUplinkRANConfigurationTransfer        Integer = 48
	ProcedureCodeUplinkRANStatusTransfer               Integer = 49
	ProcedureCodeUplinkUEAssociatedNRPPaTransport      Integer = 50
	ProcedureCodeWriteReplaceWarning                   Integer = 51
	ProcedureCodeSecondaryRATDataUsageReport           Integer = 52
	ProcedureCodeUplinkRIMInformationTransfer          Integer = 53
	ProcedureCodeDownlinkRIMInformationTransfer        Integer = 54
)

type ProcedureCode struct {
	Value Integer `vht:"valueMin:0,valueMax:255"`
}

type ProtocolExtensionID struct {
	Value Integer `vht:"valueMin:0,valueMax:65535"`
}

const (
	ProtocolIEIDAllowedNSSAI                                Integer = 0
	ProtocolIEIDAMFName                                     Integer = 1
	ProtocolIEIDAMFOverloadResponse                         Integer = 2
	ProtocolIEIDAMFSetID                                    Integer = 3
	ProtocolIEIDAMFTNLAssociationFailedToSetupList          Integer = 4
	ProtocolIEIDAMFTNLAssociationSetupList                  Integer = 5
	ProtocolIEIDAMFTNLAssociationToAddList                  Integer = 6
	ProtocolIEIDAMFTNLAssociationToRemoveList               Integer = 7
	ProtocolIEIDAMFTNLAssociationToUpdateList               Integer = 8
	ProtocolIEIDAMFTrafficLoadReductionIndication           Integer = 9
	ProtocolIEIDAMFUENGAPID                                 Integer = 10
	ProtocolIEIDAssistanceDataForPaging                     Integer = 11
	ProtocolIEIDBroadcastCancelledAreaList                  Integer = 12
	ProtocolIEIDBroadcastCompletedAreaList                  Integer = 13
	ProtocolIEIDCancelAllWarningMessages                    Integer = 14
	ProtocolIEIDCause                                       Integer = 15
	ProtocolIEIDCellIDListForRestart                        Integer = 16
	ProtocolIEIDConcurrentWarningMessageInd                 Integer = 17
	ProtocolIEIDCoreNetworkAssistanceInformationForInactive Integer = 18
	ProtocolIEIDCriticalityDiagnostics                      Integer = 19
	ProtocolIEIDDataCodingScheme                            Integer = 20
	ProtocolIEIDDefaultPagingDRX                            Integer = 21
	ProtocolIEIDDirectForwardingPathAvailability            Integer = 22
	ProtocolIEIDEmergencyAreaIDListForRestart               Integer = 23
	ProtocolIEIDEmergencyFallbackIndicator                  Integer = 24
	ProtocolIEIDEUTRACGI                                    Integer = 25
	ProtocolIEIDFiveGSTMSI                                  Integer = 26
	ProtocolIEIDGlobalRANNodeID                             Integer = 27
	ProtocolIEIDGUAMI                                       Integer = 28
	ProtocolIEIDHandoverType                                Integer = 29
	ProtocolIEIDIMSVoiceSupportIndicator                    Integer = 30
	ProtocolIEIDIndexToRFSP                                 Integer = 31
	ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging  Integer = 32
	ProtocolIEIDLocationReportingRequestType                Integer = 33
	ProtocolIEIDMaskedIMEISV                                Integer = 34
	ProtocolIEIDMessageIdentifier                           Integer = 35
	ProtocolIEIDMobilityRestrictionList                     Integer = 36
	ProtocolIEIDNASC                                        Integer = 37
	ProtocolIEIDNASPDU                                      Integer = 38
	ProtocolIEIDNASSecurityParametersFromNGRAN              Integer = 39
	ProtocolIEIDNewAMFUENGAPID                              Integer = 40
	ProtocolIEIDNewSecurityContextInd                       Integer = 41
	ProtocolIEIDNGAPMessage                                 Integer = 42
	ProtocolIEIDNGRANCGI                                    Integer = 43
	ProtocolIEIDNGRANTraceID                                Integer = 44
	ProtocolIEIDNRCGI                                       Integer = 45
	ProtocolIEIDNRPPaPDU                                    Integer = 46
	ProtocolIEIDNumberOfBroadcastsRequested                 Integer = 47
	ProtocolIEIDOldAMF                                      Integer = 48
	ProtocolIEIDOverloadStartNSSAIList                      Integer = 49
	ProtocolIEIDPagingDRX                                   Integer = 50
	ProtocolIEIDPagingOrigin                                Integer = 51
	ProtocolIEIDPagingPriority                              Integer = 52
	ProtocolIEIDPDUSessionResourceAdmittedList              Integer = 53
	ProtocolIEIDPDUSessionResourceFailedToModifyListModRes  Integer = 54
	ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes   Integer = 55
	ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck    Integer = 56
	ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq    Integer = 57
	ProtocolIEIDPDUSessionResourceFailedToSetupListSURes    Integer = 58
	ProtocolIEIDPDUSessionResourceHandoverList              Integer = 59
	ProtocolIEIDPDUSessionResourceListCxtRelCpl             Integer = 60
	ProtocolIEIDPDUSessionResourceListHORqd                 Integer = 61
	ProtocolIEIDPDUSessionResourceModifyListModCfm          Integer = 62
	ProtocolIEIDPDUSessionResourceModifyListModInd          Integer = 63
	ProtocolIEIDPDUSessionResourceModifyListModReq          Integer = 64
	ProtocolIEIDPDUSessionResourceModifyListModRes          Integer = 65
	ProtocolIEIDPDUSessionResourceNotifyList                Integer = 66
	ProtocolIEIDPDUSessionResourceReleasedListNot           Integer = 67
	ProtocolIEIDPDUSessionResourceReleasedListPSAck         Integer = 68
	ProtocolIEIDPDUSessionResourceReleasedListPSFail        Integer = 69
	ProtocolIEIDPDUSessionResourceReleasedListRelRes        Integer = 70
	ProtocolIEIDPDUSessionResourceSetupListCxtReq           Integer = 71
	ProtocolIEIDPDUSessionResourceSetupListCxtRes           Integer = 72
	ProtocolIEIDPDUSessionResourceSetupListHOReq            Integer = 73
	ProtocolIEIDPDUSessionResourceSetupListSUReq            Integer = 74
	ProtocolIEIDPDUSessionResourceSetupListSURes            Integer = 75
	ProtocolIEIDPDUSessionResourceToBeSwitchedDLList        Integer = 76
	ProtocolIEIDPDUSessionResourceSwitchedList              Integer = 77
	ProtocolIEIDPDUSessionResourceToReleaseListHOCmd        Integer = 78
	ProtocolIEIDPDUSessionResourceToReleaseListRelCmd       Integer = 79
	ProtocolIEIDPLMNSupportList                             Integer = 80
	ProtocolIEIDPWSFailedCellIDList                         Integer = 81
	ProtocolIEIDRANNodeName                                 Integer = 82
	ProtocolIEIDRANPagingPriority                           Integer = 83
	ProtocolIEIDRANStatusTransferTransparentContainer       Integer = 84
	ProtocolIEIDRANUENGAPID                                 Integer = 85
	ProtocolIEIDRelativeAMFCapacity                         Integer = 86
	ProtocolIEIDRepetitionPeriod                            Integer = 87
	ProtocolIEIDResetType                                   Integer = 88
	ProtocolIEIDRoutingID                                   Integer = 89
	ProtocolIEIDRRCEstablishmentCause                       Integer = 90
	ProtocolIEIDRRCInactiveTransitionReportRequest          Integer = 91
	ProtocolIEIDRRCState                                    Integer = 92
	ProtocolIEIDSecurityContext                             Integer = 93
	ProtocolIEIDSecurityKey                                 Integer = 94
	ProtocolIEIDSerialNumber                                Integer = 95
	ProtocolIEIDServedGUAMIList                             Integer = 96
	ProtocolIEIDSliceSupportList                            Integer = 97
	ProtocolIEIDSONConfigurationTransferDL                  Integer = 98
	ProtocolIEIDSONConfigurationTransferUL                  Integer = 99
	ProtocolIEIDSourceAMFUENGAPID                           Integer = 100
	ProtocolIEIDSourceToTargetTransparentContainer          Integer = 101
	ProtocolIEIDSupportedTAList                             Integer = 102
	ProtocolIEIDTAIListForPaging                            Integer = 103
	ProtocolIEIDTAIListForRestart                           Integer = 104
	ProtocolIEIDTargetID                                    Integer = 105
	ProtocolIEIDTargetToSourceTransparentContainer          Integer = 106
	ProtocolIEIDTimeToWait                                  Integer = 107
	ProtocolIEIDTraceActivation                             Integer = 108
	ProtocolIEIDTraceCollectionEntityIPAddress              Integer = 109
	ProtocolIEIDUEAggregateMaximumBitRate                   Integer = 110
	ProtocolIEIDUEAssociatedLogicalNGConnectionList         Integer = 111
	ProtocolIEIDUEContextRequest                            Integer = 112
	ProtocolIEIDUENGAPIDs                                   Integer = 114
	ProtocolIEIDUEPagingIdentity                            Integer = 115
	ProtocolIEIDUEPresenceInAreaOfInterestList              Integer = 116
	ProtocolIEIDUERadioCapability                           Integer = 117
	ProtocolIEIDUERadioCapabilityForPaging                  Integer = 118
	ProtocolIEIDUESecurityCapabilities                      Integer = 119
	ProtocolIEIDUnavailableGUAMIList                        Integer = 120
	ProtocolIEIDUserLocationInformation                     Integer = 121
	ProtocolIEIDWarningAreaList                             Integer = 122
	ProtocolIEIDWarningMessageContents                      Integer = 123
	ProtocolIEIDWarningSecurityInfo                         Integer = 124
	ProtocolIEIDWarningType                                 Integer = 125
	ProtocolIEIDAdditionalULNGUUPTNLInformation             Integer = 126
	ProtocolIEIDDataForwardingNotPossible                   Integer = 127
	ProtocolIEIDDLNGUUPTNLInformation                       Integer = 128
	ProtocolIEIDNetworkInstance                             Integer = 129
	ProtocolIEIDPDUSessionAggregateMaximumBitRate           Integer = 130
	ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm  Integer = 131
	ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail  Integer = 132
	ProtocolIEIDPDUSessionResourceListCxtRelReq             Integer = 133
	ProtocolIEIDPDUSessionType                              Integer = 134
	ProtocolIEIDQosFlowAddOrModifyRequestList               Integer = 135
	ProtocolIEIDQosFlowSetupRequestList                     Integer = 136
	ProtocolIEIDQosFlowToReleaseList                        Integer = 137
	ProtocolIEIDSecurityIndication                          Integer = 138
	ProtocolIEIDULNGUUPTNLInformation                       Integer = 139
	ProtocolIEIDULNGUUPTNLModifyList                        Integer = 140
	ProtocolIEIDWarningAreaCoordinates                      Integer = 141
	ProtocolIEIDPDUSessionResourceSecondaryRATUsageList     Integer = 142
	ProtocolIEIDHandoverFlag                                Integer = 143
	ProtocolIEIDSecondaryRATUsageInformation                Integer = 144
	ProtocolIEIDPDUSessionResourceReleaseResponseTransfer   Integer = 145
	ProtocolIEIDRedirectionVoiceFallback                    Integer = 146
	ProtocolIEIDUERetentionInformation                      Integer = 147
	ProtocolIEIDSNSSAI                                      Integer = 148
	ProtocolIEIDPSCellInformation                           Integer = 149
	ProtocolIEIDLastEUTRANPLMNIdentity                      Integer = 150
	ProtocolIEIDMaximumIntegrityProtectedDataRateDL         Integer = 151
	ProtocolIEIDAdditionalDLForwardingUPTNLInformation      Integer = 152
	ProtocolIEIDAdditionalDLUPTNLInformationForHOList       Integer = 153
	ProtocolIEIDAdditionalNGUUPTNLInformation               Integer = 154
	ProtocolIEIDAdditionalDLQosFlowPerTNLInformation        Integer = 155
	ProtocolIEIDSecurityResult                              Integer = 156
	ProtocolIEIDENDCSONConfigurationTransferDL              Integer = 157
	ProtocolIEIDENDCSONConfigurationTransferUL              Integer = 158
	ProtocolIEIDOldAssociatedQosFlowListULendmarkerexpected Integer = 159
	ProtocolIEIDCNTypeRestrictionsForEquivalent             Integer = 160
	ProtocolIEIDCNTypeRestrictionsForServing                Integer = 161
	ProtocolIEIDNewGUAMI                                    Integer = 162
	ProtocolIEIDULForwarding                                Integer = 163
	ProtocolIEIDULForwardingUPTNLInformation                Integer = 164
	ProtocolIEIDCNAssistedRANTuning                         Integer = 165
	ProtocolIEIDCommonNetworkInstance                       Integer = 166
	ProtocolIEIDNGRANTNLAssociationToRemoveList             Integer = 167
	ProtocolIEIDTNLAssociationTransportLayerAddressNGRAN    Integer = 168
	ProtocolIEIDEndpointIPAddressAndPort                    Integer = 169
	ProtocolIEIDLocationReportingAdditionalInfo             Integer = 170
	ProtocolIEIDSourceToTargetAMFInformationReroute         Integer = 171
	ProtocolIEIDAdditionalULForwardingUPTNLInformation      Integer = 172
	ProtocolIEIDSCTPTLAs                                    Integer = 173
	ProtocolIEIDDataForwardingResponseERABList              Integer = 174
	ProtocolIEIDRIMInformationTransfer                      Integer = 175
	ProtocolIEIDGUAMIType                                   Integer = 176
	ProtocolIEIDSRVCCOperationPossible                      Integer = 177
	ProtocolIEIDTargetRNCID                                 Integer = 178
	ProtocolIEIDRATInformation                              Integer = 179
	ProtocolIEIDExtendedRATRestrictionInformation           Integer = 180
	ProtocolIEIDQosMonitoringRequest                        Integer = 181
	ProtocolIEIDSgNBUEX2APID                                Integer = 182
)

type ProtocolIEID struct {
	Value Integer `vht:"valueMin:0,valueMax:65535"`
}

const (
	TriggeringMessageInitiatingMessage   Enumerated = 0
	TriggeringMessageSuccessfulOutcome   Enumerated = 1
	TriggeringMessageUnsuccessfulOutcome Enumerated = 2
)

type TriggeringMessage struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

/*
	Extension constants
*/
const (
	MaxnoofAllowedAreas              Integer = 16
	MaxnoofAllowedSNSSAIs            Integer = 8
	MaxnoofBPLMNs                    Integer = 12
	MaxnoofCellIDforWarning          Integer = 65535
	MaxnoofCellinAoI                 Integer = 256
	MaxnoofCellinEAI                 Integer = 65535
	MaxnoofCellinTAI                 Integer = 65535
	MaxnoofCellsingNB                Integer = 16384
	MaxnoofCellsinngeNB              Integer = 256
	MaxnoofCellsinUEHistoryInfo      Integer = 16
	MaxnoofCellsUEMovingTrajectory   Integer = 16
	MaxnoofDRBs                      Integer = 32
	MaxnoofEmergencyAreaID           Integer = 65535
	MaxnoofEAIforRestart             Integer = 256
	MaxnoofEPLMNs                    Integer = 15
	MaxnoofEPLMNsPlusOne             Integer = 16
	MaxnoofERABs                     Integer = 256
	MaxnoofErrors                    Integer = 256
	MaxnoofForbTACs                  Integer = 4096
	MaxnoofMultiConnectivity         Integer = 4
	MaxnoofMultiConnectivityMinusOne Integer = 3
	MaxnoofNGConnectionsToReset      Integer = 65536
	MaxnoofPDUSessions               Integer = 256
	MaxnoofPLMNs                     Integer = 12
	MaxnoofQosFlows                  Integer = 64
	MaxnoofRANNodeinAoI              Integer = 64
	MaxnoofRecommendedCells          Integer = 16
	MaxnoofRecommendedRANNodes       Integer = 16
	MaxnoofAoI                       Integer = 64
	MaxnoofServedGUAMIs              Integer = 256
	MaxnoofSliceItems                Integer = 1024
	MaxnoofTACs                      Integer = 256
	MaxnoofTAIforInactive            Integer = 16
	MaxnoofTAIforPaging              Integer = 16
	MaxnoofTAIforRestart             Integer = 2048
	MaxnoofTAIforWarning             Integer = 65535
	MaxnoofTAIinAoI                  Integer = 16
	MaxnoofTimePeriods               Integer = 2
	MaxnoofTNLAssociations           Integer = 32
	MaxnoofXnExtTLAs                 Integer = 16
	MaxnoofXnGTPTLAs                 Integer = 16
	MaxnoofXnTLAs                    Integer = 2
)

const (
	MaxPrivateIEs         Integer = 65535
	MaxProtocolExtensions Integer = 65535
	MaxProtocolIEs        Integer = 65535
)
