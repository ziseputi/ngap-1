// Created By HaoDHH-245789 VHT2020
package ngapBuild

import "vht5gc/lib/ngap/ngapType"

//const ProtocolIEIDAllowedNSSAI int64 = 0
func AllowedNSSAI () (value ngapType.AllowedNSSAI) {
	item := AllowedNSSAIItem()
	value.List = append(value.List, item, item, item)
	return
}

//const ProtocolIEIDAMFName int64 = 1
func AMFName () (value ngapType.AMFName) {
	value.Value = "haodhh-ngap-amf-5gc"
	return
}

//const ProtocolIEIDAMFOverloadResponse int64 = 2
func AMFOverloadResponse () (value ngapType.OverloadResponse) {
	return OverloadResponse()
}

//const ProtocolIEIDAMFSetID int64 = 3
func AMFSetID () (value ngapType.AMFSetID) {
	value.Value.BitLength = 10
	value.Value.Bytes = []byte{171,255}
	return
}

//const ProtocolIEIDAMFTNLAssociationFailedToSetupList int64 = 4
func AMFTNLAssociationFailedToSetupList () (value ngapType.TNLAssociationList) {
	return TNLAssociationList()
}

//const ProtocolIEIDAMFTNLAssociationSetupList int64 = 5
func AMFTNLAssociationSetupList () (value ngapType.AMFTNLAssociationSetupList) {
	item := AMFTNLAssociationSetupItem()
	value.List = append(value.List, item, item, item)
	return
}

//const ProtocolIEIDAMFTNLAssociationToAddList int64 = 6
func AMFTNLAssociationToAddList () (value ngapType.AMFTNLAssociationToAddList) {
	item := AMFTNLAssociationToAddItem()
	value.List = append(value.List, item, item, item)
	return
}

//const ProtocolIEIDAMFTNLAssociationToRemoveList int64 = 7
func AMFTNLAssociationToRemoveList () (value ngapType.AMFTNLAssociationToRemoveList) {
	item := AMFTNLAssociationToRemoveItem()
	value.List = append(value.List, item, item, item)
	return
}

//const ProtocolIEIDAMFTNLAssociationToUpdateList int64 = 8
func AMFTNLAssociationToUpdateList () (value ngapType.AMFTNLAssociationToUpdateList) {
	item := AMFTNLAssociationToUpdateItem()
	value.List = append(value.List, item, item, item)
	return
}

//const ProtocolIEIDAMFTrafficLoadReductionIndication int64 = 9
func AMFTrafficLoadReductionIndication () (value ngapType.TrafficLoadReductionIndication) {
	value.Value = 49 // 1-99
	return
}

//const ProtocolIEIDAMFUENGAPID int64 = 10
func AMFUENGAPID () (value ngapType.AMFUENGAPID) {
	value.Value = 1099511627775 // 0-1099511627775
	return
}

// const ProtocolIEIDAssistanceDataForPaging int64 = 11
func AssistanceDataForPaging () (value ngapType.AssistanceDataForPaging) {
	value.AssistanceDataForRecommendedCells = new(ngapType.AssistanceDataForRecommendedCells)
	*value.AssistanceDataForRecommendedCells = AssistanceDataForRecommendedCells()
	value.PagingAttemptInformation = new(ngapType.PagingAttemptInformation)
	*value.PagingAttemptInformation = PagingAttemptInformation()
	return
}

// const ProtocolIEIDBroadcastCancelledAreaList int64 = 12
func BroadcastCancelledAreaList () (value ngapType.BroadcastCancelledAreaList) {
	value.Present = ngapType.BroadcastCancelledAreaListPresentCellIDCancelledEUTRA
	value.CellIDCancelledEUTRA = new(ngapType.CellIDCancelledEUTRA)
	*value.CellIDCancelledEUTRA = CellIDCancelledEUTRA()
	value.Present = ngapType.BroadcastCancelledAreaListPresentTAICancelledEUTRA
	value.TAICancelledEUTRA = new(ngapType.TAICancelledEUTRA)
	*value.TAICancelledEUTRA = TAICancelledEUTRA()
	value.Present = ngapType.BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledEUTRA
	value.EmergencyAreaIDCancelledEUTRA = new(ngapType.EmergencyAreaIDCancelledEUTRA)
	*value.EmergencyAreaIDCancelledEUTRA = EmergencyAreaIDCancelledEUTRA()
	value.Present = ngapType.BroadcastCancelledAreaListPresentCellIDCancelledNR
	value.CellIDCancelledNR = new(ngapType.CellIDCancelledNR)
	*value.CellIDCancelledNR = CellIDCancelledNR()
	value.Present = ngapType.BroadcastCancelledAreaListPresentTAICancelledNR
	value.TAICancelledNR = new(ngapType.TAICancelledNR)
	*value.TAICancelledNR = TAICancelledNR()
	value.Present = ngapType.BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledNR
	value.EmergencyAreaIDCancelledNR = new(ngapType.EmergencyAreaIDCancelledNR)
	*value.EmergencyAreaIDCancelledNR = EmergencyAreaIDCancelledNR()
	return
}

// const ProtocolIEIDBroadcastCompletedAreaList int64 = 13
func BroadcastCompletedAreaList () (value ngapType.BroadcastCompletedAreaList) {
	value.Present = ngapType.BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA
	value.CellIDBroadcastEUTRA = new(ngapType.CellIDBroadcastEUTRA)
	*value.CellIDBroadcastEUTRA = CellIDBroadcastEUTRA()
	value.Present = ngapType.BroadcastCompletedAreaListPresentTAIBroadcastEUTRA
	value.TAIBroadcastEUTRA = new(ngapType.TAIBroadcastEUTRA)
	*value.TAIBroadcastEUTRA = TAIBroadcastEUTRA()
	value.Present = ngapType.BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA
	value.EmergencyAreaIDBroadcastEUTRA = new(ngapType.EmergencyAreaIDBroadcastEUTRA)
	*value.EmergencyAreaIDBroadcastEUTRA = EmergencyAreaIDBroadcastEUTRA()
	value.Present = ngapType.BroadcastCompletedAreaListPresentCellIDBroadcastNR
	value.CellIDBroadcastNR = new(ngapType.CellIDBroadcastNR)
	*value.CellIDBroadcastNR = CellIDBroadcastNR()
	value.Present = ngapType.BroadcastCompletedAreaListPresentTAIBroadcastNR
	value.TAIBroadcastNR = new(ngapType.TAIBroadcastNR)
	*value.TAIBroadcastNR = TAIBroadcastNR()
	value.Present = ngapType.BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR
	value.EmergencyAreaIDBroadcastNR = new(ngapType.EmergencyAreaIDBroadcastNR)
	*value.EmergencyAreaIDBroadcastNR = EmergencyAreaIDBroadcastNR()
	return
}

// const ProtocolIEIDCancelAllWarningMessages int64 = 14
func CancelAllWarningMessages () (value ngapType.CancelAllWarningMessages) {
	value.Value = ngapType.CancelAllWarningMessagesPresentTrue
	return
}

//const ProtocolIEIDCause int64 = 15
func Cause () (value ngapType.Cause) {
	value.Present = ngapType.CausePresentRadioNetwork
	value.RadioNetwork = new(ngapType.CauseRadioNetwork)
	*value.RadioNetwork = CauseRadioNetwork()
	value.Present = ngapType.CausePresentTransport
	value.Transport = new(ngapType.CauseTransport)
	*value.Transport = CauseTransport()
	value.Present = ngapType.CausePresentNas
	value.Nas = new(ngapType.CauseNas)
	*value.Nas = CauseNas()
	value.Present = ngapType.CausePresentProtocol
	value.Protocol = new(ngapType.CauseProtocol)
	*value.Protocol = CauseProtocol()
	value.Present = ngapType.CausePresentMisc
	value.Misc = new(ngapType.CauseMisc)
	*value.Misc = CauseMisc()
	return
}

// const ProtocolIEIDCellIDListForRestart int64 = 16
func CellIDListForRestart () (value ngapType.CellIDListForRestart) {
	value.Present = ngapType.CellIDListForRestartPresentEUTRACGIListforRestart
	value.EUTRACGIListforRestart = new(ngapType.EUTRACGIList)
	*value.EUTRACGIListforRestart = EUTRACGIList()
	value.Present = ngapType.CellIDListForRestartPresentNRCGIListforRestart
	value.NRCGIListforRestart = new(ngapType.NRCGIList)
	*value.NRCGIListforRestart = NRCGIList()
	return
}

// const ProtocolIEIDConcurrentWarningMessageInd int64 = 17
func ConcurrentWarningMessageInd () (value ngapType.ConcurrentWarningMessageInd) {
	value.Value = ngapType.ConcurrentWarningMessageIndPresentTrue
	return
}

// const ProtocolIEIDCoreNetworkAssistanceInformation int64 = 18
func CoreNetworkAssistanceInformation () (value ngapType.CoreNetworkAssistanceInformation) {
	value.UEIdentityIndexValue = UEIdentityIndexValue()
	value.UESpecificDRX = new(ngapType.PagingDRX)
	*value.UESpecificDRX = PagingDRX()
	value.PeriodicRegistrationUpdateTimer = PeriodicRegistrationUpdateTimer()
	value.MICOModeIndication = new(ngapType.MICOModeIndication)
	*value.MICOModeIndication = MICOModeIndication()
	value.TAIListForInactive = TAIListForInactive()
	value.ExpectedUEBehaviour = new(ngapType.ExpectedUEBehaviour)
	*value.ExpectedUEBehaviour = ExpectedUEBehaviour()
	return
}

// const ProtocolIEIDCriticalityDiagnostics int64 = 19
func CriticalityDiagnostics () (value ngapType.CriticalityDiagnostics) {
	value.ProcedureCode = new(ngapType.ProcedureCode)
	*value.ProcedureCode = ProcedureCode()
	value.TriggeringMessage = new(ngapType.TriggeringMessage)
	*value.TriggeringMessage = TriggeringMessage()
	value.ProcedureCriticality = new(ngapType.Criticality)
	*value.ProcedureCriticality = Criticality()
	value.IEsCriticalityDiagnostics = new(ngapType.CriticalityDiagnosticsIEList)
	*value.IEsCriticalityDiagnostics = CriticalityDiagnosticsIEList()
	return
}

// const ProtocolIEIDDataCodingScheme int64 = 20
func DataCodingScheme () (value ngapType.DataCodingScheme) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{183}
	return
}

// const ProtocolIEIDDefaultPagingDRX int64 = 21
func DefaultPagingDRX () (value ngapType.PagingDRX) {
	return PagingDRX()
}

// const ProtocolIEIDDirectForwardingPathAvailability int64 = 22
func DirectForwardingPathAvailability () (value ngapType.DirectForwardingPathAvailability) {
	value.Value = ngapType.DirectForwardingPathAvailabilityPresentDirectPathAvailable
	return
}

// const ProtocolIEIDEmergencyAreaIDListForRestart int64 = 23
func EmergencyAreaIDListForRestart () (value ngapType.EmergencyAreaIDListForRestart) {
	item := EmergencyAreaID()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDEmergencyFallbackIndicator int64 = 24
func EmergencyFallbackIndicator () (value ngapType.EmergencyFallbackIndicator) {
	value.EmergencyFallbackRequestIndicator = EmergencyFallbackRequestIndicator()
	value.EmergencyServiceTargetCN = new(ngapType.EmergencyServiceTargetCN)
	*value.EmergencyServiceTargetCN = EmergencyServiceTargetCN()
	return
}

// const ProtocolIEIDEUTRACGI int64 = 25
func EUTRACGI () (value ngapType.EUTRACGI) {
	value.PLMNIdentity = PLMNIdentity()
	value.EUTRACellIdentity = EUTRACellIdentity()
	return
}

// const ProtocolIEIDFiveGSTMSI int64 = 26
func FiveGSTMSI () (value ngapType.FiveGSTMSI) {
	value.AMFSetID = AMFSetID()
	value.AMFPointer = AMFPointer()
	value.FiveGTMSI = FiveGTMSI()
	return
}

// const ProtocolIEIDGlobalRANNodeID int64 = 27
func GlobalRANNodeID () (value ngapType.GlobalRANNodeID) {
	value.Present = ngapType.GlobalRANNodeIDPresentGlobalGNBID
	value.GlobalGNBID = new(ngapType.GlobalGNBID)
	*value.GlobalGNBID = GlobalGNBID()
	value.Present = ngapType.GlobalRANNodeIDPresentGlobalNgENBID
	value.GlobalNgENBID = new(ngapType.GlobalNgENBID)
	*value.GlobalNgENBID = GlobalNgENBID()
	value.Present = ngapType.GlobalRANNodeIDPresentGlobalN3IWFID
	value.GlobalN3IWFID = new(ngapType.GlobalN3IWFID)
	*value.GlobalN3IWFID = GlobalN3IWFID()
	return
}

// const ProtocolIEIDGUAMI int64 = 28
func GUAMI () (value ngapType.GUAMI) {
	value.PLMNIdentity = PLMNIdentity()
	value.AMFRegionID = AMFRegionID()
	value.AMFSetID = AMFSetID()
	value.AMFPointer = AMFPointer()
	return
}

// const ProtocolIEIDHandoverType int64 = 29
func HandoverType () (value ngapType.HandoverType) {
	value.Value = ngapType.HandoverTypePresentIntra5gs
	value.Value = ngapType.HandoverTypePresentFivegsToEps
	value.Value = ngapType.HandoverTypePresentEpsTo5gs
	return
}

// const ProtocolIEIDIMSVoiceSupportIndicator int64 = 30
func IMSVoiceSupportIndicator () (value ngapType.IMSVoiceSupportIndicator) {
	value.Value = ngapType.IMSVoiceSupportIndicatorPresentSupported
	value.Value = ngapType.IMSVoiceSupportIndicatorPresentNotSupported
	return
}

// const ProtocolIEIDIndexToRFSP int64 = 31
func IndexToRFSP () (value ngapType.IndexToRFSP) {
	value.Value = 256 // 1-256
	return
}

// const ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging int64 = 32
func InfoOnRecommendedCellsAndRANNodesForPaging () (value ngapType.InfoOnRecommendedCellsAndRANNodesForPaging) {
	value.RecommendedCellsForPaging = RecommendedCellsForPaging()
	value.RecommendRANNodesForPaging = RecommendedRANNodesForPaging()
	return
}

// const ProtocolIEIDLocationReportingRequestType int64 = 33
func LocationReportingRequestType () (value ngapType.LocationReportingRequestType) {
	value.EventType = EventType()
	value.ReportArea = ReportArea()
	value.AreaOfInterestList = new(ngapType.AreaOfInterestList)
	*value.AreaOfInterestList = AreaOfInterestList()
	value.LocationReportingReferenceIDToBeCancelled = new(ngapType.LocationReportingReferenceID)
	*value.LocationReportingReferenceIDToBeCancelled = LocationReportingReferenceID()
	return
}

// const ProtocolIEIDMaskedIMEISV int64 = 34
func MaskedIMEISV () (value ngapType.MaskedIMEISV) {
	value.Value.BitLength = 64
	value.Value.Bytes = []byte{171,11,22,33,44,55,66,255}
	return
}

// const ProtocolIEIDMessageIdentifier int64 = 35
func MessageIdentifier () (value ngapType.MessageIdentifier) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

// const ProtocolIEIDMobilityRestrictionList int64 = 36
func MobilityRestrictionList () (value ngapType.MobilityRestrictionList) {
	value.ServingPLMN = PLMNIdentity()
	value.EquivalentPLMNs = new(ngapType.EquivalentPLMNs)
	*value.EquivalentPLMNs = EquivalentPLMNs()
	value.RATRestrictions = new(ngapType.RATRestrictions)
	*value.RATRestrictions = RATRestrictions()
	value.ForbiddenAreaInformation = new(ngapType.ForbiddenAreaInformation)
	*value.ForbiddenAreaInformation = ForbiddenAreaInformation()
	value.ServiceAreaInformation = new(ngapType.ServiceAreaInformation)
	*value.ServiceAreaInformation = ServiceAreaInformation()
	return
}

// const ProtocolIEIDNASC int64 = 37
func NASC () (value ngapType.NASPDU) {
	return NASPDU()
}

// const ProtocolIEIDNASPDU int64 = 38
func NASPDU () (value ngapType.NASPDU) {
	value.Value = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

// const ProtocolIEIDNASSecurityParametersFromNGRAN int64 = 39
func NASSecurityParametersFromNGRAN () (value ngapType.NASSecurityParametersFromNGRAN) {
	value.Value = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

// const ProtocolIEIDNewAMFUENGAPID int64 = 40
func NewAMFUENGAPID () (value ngapType.AMFUENGAPID) {
	return AMFUENGAPID()
}

// const ProtocolIEIDNewSecurityContextInd int64 = 41
func NewSecurityContextInd () (value ngapType.NewSecurityContextInd) {
	value.Value = ngapType.NewSecurityContextIndPresentTrue
	return
}

// const ProtocolIEIDNGAPMessage int64 = 42
func NGAPMessage () (value []byte) {
	value = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

// const ProtocolIEIDNGRANCGI int64 = 43
func NGRANCGI () (value ngapType.NGRANCGI) {
	value.Present = ngapType.NGRANCGIPresentNRCGI
	value.NRCGI = new(ngapType.NRCGI)
	*value.NRCGI = NRCGI()
	value.Present = ngapType.NGRANCGIPresentEUTRACGI
	value.EUTRACGI = new(ngapType.EUTRACGI)
	*value.EUTRACGI = EUTRACGI()
	return
}

// const ProtocolIEIDNGRANTraceID int64 = 44
func NGRANTraceID () (value ngapType.NGRANTraceID) {
	value.Value = []byte{0x54,0xf2,0x40,0x22,0x33,0x44,0x55,0x66}
	return
}

// const ProtocolIEIDNRCGI int64 = 45
func NRCGI () (value ngapType.NRCGI) {
	value.PLMNIdentity = PLMNIdentity()
	value.NRCellIdentity = NRCellIdentity()
	return
}

// const ProtocolIEIDNRPPaPDU int64 = 46
func NRPPaPDU () (value ngapType.NRPPaPDU) {
	value.Value = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

// const ProtocolIEIDNumberOfBroadcastsRequested int64 = 47
func NumberOfBroadcastsRequested () (value ngapType.NumberOfBroadcastsRequested) {
	value.Value = 65535 // 0-65535
	return
}

// const ProtocolIEIDOldAMF int64 = 48
func OldAMF () (value ngapType.AMFName) {
	return AMFName()
}

// const ProtocolIEIDOverloadStartNSSAIList int64 = 49
func OverloadStartNSSAIList () (value ngapType.OverloadStartNSSAIList) {
	item := OverloadStartNSSAIItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPagingDRX int64 = 50
func PagingDRX () (value ngapType.PagingDRX) {
	value.Value = ngapType.PagingDRXPresentV32
	value.Value = ngapType.PagingDRXPresentV64
	value.Value = ngapType.PagingDRXPresentV128
	value.Value = ngapType.PagingDRXPresentV256
	return
}

// const ProtocolIEIDPagingOrigin int64 = 51
func PagingOrigin () (value ngapType.PagingOrigin) {
	value.Value = ngapType.PagingOriginPresentNon3gpp
	return
}

// const ProtocolIEIDPagingPriority int64 = 52
func PagingPriority () (value ngapType.PagingPriority) {
	value.Value = ngapType.PagingPriorityPresentPriolevel1
	value.Value = ngapType.PagingPriorityPresentPriolevel2
	value.Value = ngapType.PagingPriorityPresentPriolevel3
	value.Value = ngapType.PagingPriorityPresentPriolevel4
	value.Value = ngapType.PagingPriorityPresentPriolevel5
	value.Value = ngapType.PagingPriorityPresentPriolevel6
	value.Value = ngapType.PagingPriorityPresentPriolevel7
	value.Value = ngapType.PagingPriorityPresentPriolevel8
	return
}

// const ProtocolIEIDPDUSessionResourceAdmittedList int64 = 53
func PDUSessionResourceAdmittedList () (value ngapType.PDUSessionResourceAdmittedList) {
	item := PDUSessionResourceAdmittedItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToModifyListModRes int64 = 54
func PDUSessionResourceFailedToModifyListModRes () (value ngapType.PDUSessionResourceFailedToModifyListModRes) {
	item := PDUSessionResourceFailedToModifyItemModRes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes int64 = 55
func PDUSessionResourceFailedToSetupListCxtRes () (value ngapType.PDUSessionResourceFailedToSetupListCxtRes) {
	item := PDUSessionResourceFailedToSetupItemCxtRes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck int64 = 56
func PDUSessionResourceFailedToSetupListHOAck () (value ngapType.PDUSessionResourceFailedToSetupListHOAck) {
	item := PDUSessionResourceFailedToSetupItemHOAck()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq int64 = 57
func PDUSessionResourceFailedToSetupListPSReq () (value ngapType.PDUSessionResourceFailedToSetupListPSReq) {
	item := PDUSessionResourceFailedToSetupItemPSReq()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToSetupListSURes int64 = 58
func PDUSessionResourceFailedToSetupListSURes () (value ngapType.PDUSessionResourceFailedToSetupListSURes) {
	item := PDUSessionResourceFailedToSetupItemSURes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceHandoverList int64 = 59
func PDUSessionResourceHandoverList () (value ngapType.PDUSessionResourceHandoverList) {
	item := PDUSessionResourceHandoverItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceListCxtRelCpl int64 = 60
func PDUSessionResourceListCxtRelCpl () (value ngapType.PDUSessionResourceListCxtRelCpl) {
	item := PDUSessionResourceItemCxtRelCpl()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceListHORqd int64 = 61
func PDUSessionResourceListHORqd () (value ngapType.PDUSessionResourceListHORqd) {
	item := PDUSessionResourceItemHORqd()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceModifyListModCfm int64 = 62
func PDUSessionResourceModifyListModCfm () (value ngapType.PDUSessionResourceModifyListModCfm) {
	item := PDUSessionResourceModifyItemModCfm()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceModifyListModInd int64 = 63
func PDUSessionResourceModifyListModInd () (value ngapType.PDUSessionResourceModifyListModInd) {
	item := PDUSessionResourceModifyItemModInd()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceModifyListModReq int64 = 64
func PDUSessionResourceModifyListModReq () (value ngapType.PDUSessionResourceModifyListModReq) {
	item := PDUSessionResourceModifyItemModReq()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceModifyListModRes int64 = 65
func PDUSessionResourceModifyListModRes () (value ngapType.PDUSessionResourceModifyListModRes) {
	item := PDUSessionResourceModifyItemModRes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceNotifyList int64 = 66
func PDUSessionResourceNotifyList () (value ngapType.PDUSessionResourceNotifyList) {
	item := PDUSessionResourceNotifyItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceReleasedListNot int64 = 67
func PDUSessionResourceReleasedListNot () (value ngapType.PDUSessionResourceReleasedListNot) {
	item := PDUSessionResourceReleasedItemNot()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceReleasedListPSAck int64 = 68
func PDUSessionResourceReleasedListPSAck () (value ngapType.PDUSessionResourceReleasedListPSAck) {
	item := PDUSessionResourceReleasedItemPSAck()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceReleasedListPSFail int64 = 69
func PDUSessionResourceReleasedListPSFail () (value ngapType.PDUSessionResourceReleasedListPSFail) {
	item := PDUSessionResourceReleasedItemPSFail()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceReleasedListRelRes int64 = 70
func PDUSessionResourceReleasedListRelRes () (value ngapType.PDUSessionResourceReleasedListRelRes) {
	item := PDUSessionResourceReleasedItemRelRes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceSetupListCxtReq int64 = 71
func PDUSessionResourceSetupListCxtReq () (value ngapType.PDUSessionResourceSetupListCxtReq) {
	item := PDUSessionResourceSetupItemCxtReq()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceSetupListCxtRes int64 = 72
func PDUSessionResourceSetupListCxtRes () (value ngapType.PDUSessionResourceSetupListCxtRes) {
	item := PDUSessionResourceSetupItemCxtRes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceSetupListHOReq int64 = 73
func PDUSessionResourceSetupListHOReq () (value ngapType.PDUSessionResourceSetupListHOReq) {
	item := PDUSessionResourceSetupItemHOReq()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceSetupListSUReq int64 = 74
func PDUSessionResourceSetupListSUReq () (value ngapType.PDUSessionResourceSetupListSUReq) {
	item := PDUSessionResourceSetupItemSUReq()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceSetupListSURes int64 = 75
func PDUSessionResourceSetupListSURes () (value ngapType.PDUSessionResourceSetupListSURes) {
	item := PDUSessionResourceSetupItemSURes()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceToBeSwitchedDLList int64 = 76
func PDUSessionResourceToBeSwitchedDLList () (value ngapType.PDUSessionResourceToBeSwitchedDLList) {
	item := PDUSessionResourceToBeSwitchedDLItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceSwitchedList int64 = 77
func PDUSessionResourceSwitchedList () (value ngapType.PDUSessionResourceSwitchedList) {
	item := PDUSessionResourceSwitchedItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceToReleaseListHOCmd int64 = 78
func PDUSessionResourceToReleaseListHOCmd () (value ngapType.PDUSessionResourceToReleaseListHOCmd) {
	item := PDUSessionResourceToReleaseItemHOCmd()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceToReleaseListRelCmd int64 = 79
func PDUSessionResourceToReleaseListRelCmd () (value ngapType.PDUSessionResourceToReleaseListRelCmd) {
	item := PDUSessionResourceToReleaseItemRelCmd()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPLMNSupportList int64 = 80
func PLMNSupportList () (value ngapType.PLMNSupportList) {
	item := PLMNSupportItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPWSFailedCellIDList int64 = 81
func PWSFailedCellIDList () (value ngapType.PWSFailedCellIDList) {
	value.Present = ngapType.PWSFailedCellIDListPresentEUTRACGIPWSFailedList
	value.EUTRACGIPWSFailedList = new(ngapType.EUTRACGIList)
	*value.EUTRACGIPWSFailedList = EUTRACGIList()
	value.Present = ngapType.PWSFailedCellIDListPresentNRCGIPWSFailedList
	value.NRCGIPWSFailedList = new(ngapType.NRCGIList)
	*value.NRCGIPWSFailedList = NRCGIList()
	return
}

// const ProtocolIEIDRANNodeName int64 = 82
func RANNodeName () (value ngapType.RANNodeName) {
	value.Value = "5gc-haodhh-amf-ngap"
	return
}

// const ProtocolIEIDRANPagingPriority int64 = 83
func RANPagingPriority () (value ngapType.RANPagingPriority) {
	value.Value = 137 // 1-256
	return
}

// const ProtocolIEIDRANStatusTransferTransparentContainer int64 = 84
func RANStatusTransferTransparentContainer () (value ngapType.RANStatusTransferTransparentContainer) {
	value.DRBsSubjectToStatusTransferList = DRBsSubjectToStatusTransferList()
	return
}

// const ProtocolIEIDRANUENGAPID int64 = 85
func RANUENGAPID () (value ngapType.RANUENGAPID) {
	value.Value = 4294967295 // 0-4294967295
	return
}

// const ProtocolIEIDRelativeAMFCapacity int64 = 86
func RelativeAMFCapacity () (value ngapType.RelativeAMFCapacity) {
	value.Value = 137 // 0-255
	return
}

// const ProtocolIEIDRepetitionPeriod int64 = 87
func RepetitionPeriod () (value ngapType.RepetitionPeriod) {
	value.Value = 131071 // 0-131071
	return
}

// const ProtocolIEIDResetType int64 = 88
func ResetType () (value ngapType.ResetType) {
	value.Present = ngapType.ResetTypePresentNGInterface
	value.NGInterface = new(ngapType.ResetAll)
	*value.NGInterface = ResetAll()
	value.Present = ngapType.ResetTypePresentPartOfNGInterface
	value.PartOfNGInterface = new(ngapType.UEAssociatedLogicalNGConnectionList)
	*value.PartOfNGInterface = UEAssociatedLogicalNGConnectionList()
	return
}

// const ProtocolIEIDRoutingID int64 = 89
func RoutingID () (value ngapType.RoutingID) {
	value.Value = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

// const ProtocolIEIDRRCEstablishmentCause int64 = 90
func RRCEstablishmentCause () (value ngapType.RRCEstablishmentCause) {
	value.Value = ngapType.RRCEstablishmentCausePresentEmergency
	value.Value = ngapType.RRCEstablishmentCausePresentHighPriorityAccess
	value.Value = ngapType.RRCEstablishmentCausePresentMtAccess
	value.Value = ngapType.RRCEstablishmentCausePresentMoSignalling
	value.Value = ngapType.RRCEstablishmentCausePresentMoData
	value.Value = ngapType.RRCEstablishmentCausePresentMoVoiceCall
	value.Value = ngapType.RRCEstablishmentCausePresentMoVideoCall
	value.Value = ngapType.RRCEstablishmentCausePresentMoSMS
	value.Value = ngapType.RRCEstablishmentCausePresentMpsPriorityAccess
	value.Value = ngapType.RRCEstablishmentCausePresentMcsPriorityAccess
	return
}

// const ProtocolIEIDRRCInactiveTransitionReportRequest int64 = 91
func RRCInactiveTransitionReportRequest () (value ngapType.RRCInactiveTransitionReportRequest) {
	value.Value = ngapType.RRCInactiveTransitionReportRequestPresentSubsequentStateTransitionReport
	value.Value = ngapType.RRCInactiveTransitionReportRequestPresentSingleRrcConnectedStateReport
	value.Value = ngapType.RRCInactiveTransitionReportRequestPresentCancelReport
	return
}

// const ProtocolIEIDRRCState int64 = 92
func RRCState () (value ngapType.RRCState) {
	value.Value = ngapType.RRCStatePresentInactive
	value.Value = ngapType.RRCStatePresentConnected
	return
}

// const ProtocolIEIDSecurityContext int64 = 93
func SecurityContext () (value ngapType.SecurityContext) {
	value.NextHopChainingCount = NextHopChainingCount()
	value.NextHopNH = SecurityKey()
	return
}

// const ProtocolIEIDSecurityKey int64 = 94
func SecurityKey () (value ngapType.SecurityKey) {
	value.Value.BitLength = 256
	value.Value.Bytes = []byte{171,11,22,33,44,55,66,77,88,99,255,11,22,33,44,55,66,77,88,99,255,11,22,33,44,55,66,77,88,99,255,171}
	return
}

// const ProtocolIEIDSerialNumber int64 = 95
func SerialNumber () (value ngapType.SerialNumber) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

// const ProtocolIEIDServedGUAMIList int64 = 96
func ServedGUAMIList () (value ngapType.ServedGUAMIList) {
	item := ServedGUAMIItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDSliceSupportList int64 = 97
func SliceSupportList () (value ngapType.SliceSupportList) {
	item := SliceSupportItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDSONConfigurationTransferDL int64 = 98
func SONConfigurationTransferDL () (value ngapType.SONConfigurationTransfer) {
	return SONConfigurationTransfer()
}

// const ProtocolIEIDSONConfigurationTransferUL int64 = 99
func SONConfigurationTransferUL () (value ngapType.SONConfigurationTransfer) {
	return SONConfigurationTransfer()
}

// const ProtocolIEIDSourceAMFUENGAPID int64 = 100
func SourceAMFUENGAPID () (value ngapType.AMFUENGAPID) {
	return AMFUENGAPID()
}

// const ProtocolIEIDSourceToTargetTransparentContainer int64 = 101
func SourceToTargetTransparentContainer () (value ngapType.SourceToTargetTransparentContainer) {
	value.Value = []byte{171,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,171,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255,11,22,33,255}
	return
}

// const ProtocolIEIDSupportedTAList int64 = 102
func SupportedTAList () (value ngapType.SupportedTAList) {
	item := SupportedTAItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDTAIListForPaging int64 = 103
func TAIListForPaging () (value ngapType.TAIListForPaging) {
	item := TAIListForPagingItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDTAIListForRestart int64 = 104
func TAIListForRestart () (value ngapType.TAIListForRestart) {
	item := TAI()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDTargetID int64 = 105
func TargetID () (value ngapType.TargetID) {
	value.Present = ngapType.TargetIDPresentTargetRANNodeID
	value.TargetRANNodeID = new(ngapType.TargetRANNodeID)
	*value.TargetRANNodeID = TargetRANNodeID()
	value.Present = ngapType.TargetIDPresentTargeteNBID
	value.TargeteNBID = new(ngapType.TargeteNBID)
	*value.TargeteNBID = TargeteNBID()
	return
}

// const ProtocolIEIDTargetToSourceTransparentContainer int64 = 106
func TargetToSourceTransparentContainer () (value ngapType.TargetToSourceTransparentContainer) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

// const ProtocolIEIDTimeToWait int64 = 107
func TimeToWait () (value ngapType.TimeToWait) {
	value.Value = ngapType.TimeToWaitPresentV1s
	value.Value = ngapType.TimeToWaitPresentV2s
	value.Value = ngapType.TimeToWaitPresentV5s
	value.Value = ngapType.TimeToWaitPresentV10s
	value.Value = ngapType.TimeToWaitPresentV20s
	value.Value = ngapType.TimeToWaitPresentV60s
	return
}

// const ProtocolIEIDTraceActivation int64 = 108
func TraceActivation () (value ngapType.TraceActivation) {
	value.NGRANTraceID = NGRANTraceID()
	value.InterfacesToTrace = InterfacesToTrace()
	value.TraceDepth = TraceDepth()
	value.TraceCollectionEntityIPAddress = TransportLayerAddress()
	return
}

// const ProtocolIEIDTraceCollectionEntityIPAddress int64 = 109
func TraceCollectionEntityIPAddress () (value ngapType.TransportLayerAddress) {
	return TransportLayerAddress()
}

// const ProtocolIEIDUEAggregateMaximumBitRate int64 = 110
func UEAggregateMaximumBitRate () (value ngapType.UEAggregateMaximumBitRate) {
	value.UEAggregateMaximumBitRateDL = BitRate()
	value.UEAggregateMaximumBitRateUL = BitRate()
	return
}

// const ProtocolIEIDUEAssociatedLogicalNGConnectionList int64 = 111
func UEAssociatedLogicalNGConnectionList () (value ngapType.UEAssociatedLogicalNGConnectionList) {
	item := UEAssociatedLogicalNGConnectionItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDUEContextRequest int64 = 112
func UEContextRequest () (value ngapType.UEContextRequest) {
	value.Value = ngapType.UEContextRequestPresentRequested
	return
}

// const ProtocolIEIDUENGAPIDs int64 = 114
func UENGAPIDs () (value ngapType.UENGAPIDs) {
	value.Present = ngapType.UENGAPIDsPresentUENGAPIDPair
	value.UENGAPIDPair = new(ngapType.UENGAPIDPair)
	*value.UENGAPIDPair = UENGAPIDPair()
	value.Present = ngapType.UENGAPIDsPresentAMFUENGAPID
	value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*value.AMFUENGAPID = AMFUENGAPID()
	return
}

// const ProtocolIEIDUEPagingIdentity int64 = 115
func UEPagingIdentity () (value ngapType.UEPagingIdentity) {
	value.Present = ngapType.UEPagingIdentityPresentFiveGSTMSI
	value.FiveGSTMSI = new(ngapType.FiveGSTMSI)
	*value.FiveGSTMSI = FiveGSTMSI()
	return
}

// const ProtocolIEIDUEPresenceInAreaOfInterestList int64 = 116
func UEPresenceInAreaOfInterestList () (value ngapType.UEPresenceInAreaOfInterestList) {
	item := UEPresenceInAreaOfInterestItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDUERadioCapability int64 = 117
func UERadioCapability () (value ngapType.UERadioCapability) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

// const ProtocolIEIDUERadioCapabilityForPaging int64 = 118
func UERadioCapabilityForPaging () (value ngapType.UERadioCapabilityForPaging) {
	value.UERadioCapabilityForPagingOfNR = new(ngapType.UERadioCapabilityForPagingOfNR)
	*value.UERadioCapabilityForPagingOfNR = UERadioCapabilityForPagingOfNR()
	value.UERadioCapabilityForPagingOfEUTRA = new(ngapType.UERadioCapabilityForPagingOfEUTRA)
	*value.UERadioCapabilityForPagingOfEUTRA = UERadioCapabilityForPagingOfEUTRA()
	return
}

// const ProtocolIEIDUESecurityCapabilities int64 = 119
func UESecurityCapabilities () (value ngapType.UESecurityCapabilities) {
	value.NRencryptionAlgorithms = NRencryptionAlgorithms()
	value.NRintegrityProtectionAlgorithms = NRintegrityProtectionAlgorithms()
	value.EUTRAencryptionAlgorithms = EUTRAencryptionAlgorithms()
	value.EUTRAintegrityProtectionAlgorithms = EUTRAintegrityProtectionAlgorithms()
	return
}

// const ProtocolIEIDUnavailableGUAMIList int64 = 120
func UnavailableGUAMIList () (value ngapType.UnavailableGUAMIList) {
	item := UnavailableGUAMIItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDUserLocationInformation int64 = 121
func UserLocationInformation () (value ngapType.UserLocationInformation) {
	value.Present = ngapType.UserLocationInformationPresentUserLocationInformationEUTRA
	value.UserLocationInformationEUTRA = new(ngapType.UserLocationInformationEUTRA)
	*value.UserLocationInformationEUTRA = UserLocationInformationEUTRA()
	value.Present = ngapType.UserLocationInformationPresentUserLocationInformationNR
	value.UserLocationInformationNR = new(ngapType.UserLocationInformationNR)
	*value.UserLocationInformationNR = UserLocationInformationNR()
	value.Present = ngapType.UserLocationInformationPresentUserLocationInformationN3IWF
	value.UserLocationInformationN3IWF = new(ngapType.UserLocationInformationN3IWF)
	*value.UserLocationInformationN3IWF = UserLocationInformationN3IWF()
	return
}

// const ProtocolIEIDWarningAreaList int64 = 122
func WarningAreaList () (value ngapType.WarningAreaList) {
	value.Present = ngapType.WarningAreaListPresentEUTRACGIListForWarning
	value.EUTRACGIListForWarning = new(ngapType.EUTRACGIListForWarning)
	*value.EUTRACGIListForWarning = EUTRACGIListForWarning()
	value.Present = ngapType.WarningAreaListPresentNRCGIListForWarning
	value.NRCGIListForWarning = new(ngapType.NRCGIListForWarning)
	*value.NRCGIListForWarning = NRCGIListForWarning()
	value.Present = ngapType.WarningAreaListPresentTAIListForWarning
	value.TAIListForWarning = new(ngapType.TAIListForWarning)
	*value.TAIListForWarning = TAIListForWarning()
	value.Present = ngapType.WarningAreaListPresentEmergencyAreaIDList
	value.EmergencyAreaIDList = new(ngapType.EmergencyAreaIDList)
	*value.EmergencyAreaIDList = EmergencyAreaIDList()
	return
}

// const ProtocolIEIDWarningMessageContents int64 = 123
func WarningMessageContents () (value ngapType.WarningMessageContents) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,99,255} // 1-9600
	return
}

// const ProtocolIEIDWarningSecurityInfo int64 = 124
func WarningSecurityInfo () (value ngapType.WarningSecurityInfo) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,255,171,11,22,33,44,55,66,77,88,255,171,11,22,33,44,55,66,77,88,255,171,11,22,33,44,55,66,77,88,255,171,11,22,33,44,55,66,77,88,255}
	return
}

// const ProtocolIEIDWarningType int64 = 125
func WarningType () (value ngapType.WarningType) {
	value.Value = []byte{171,255}
	return
}

// const ProtocolIEIDAdditionalULNGUUPTNLInformation int64 = 126
func AdditionalULNGUUPTNLInformation () (value ngapType.UPTransportLayerInformation) {
	return UPTransportLayerInformation()
}

// const ProtocolIEIDDataForwardingNotPossible int64 = 127
func DataForwardingNotPossible () (value ngapType.DataForwardingNotPossible) {
	value.Value = ngapType.DataForwardingNotPossiblePresentDataForwardingNotPossible
	return
}

// const ProtocolIEIDDLNGUUPTNLInformation int64 = 128
func DLNGUUPTNLInformation () (value ngapType.UPTransportLayerInformation) {
	return UPTransportLayerInformation()
}

// const ProtocolIEIDNetworkInstance int64 = 129
func NetworkInstance () (value ngapType.NetworkInstance) {
	value.Value = 137 // 1-256
	return
}

// const ProtocolIEIDPDUSessionAggregateMaximumBitRate int64 = 130
func PDUSessionAggregateMaximumBitRate () (value ngapType.PDUSessionAggregateMaximumBitRate) {
	value.PDUSessionAggregateMaximumBitRateDL = BitRate()
	value.PDUSessionAggregateMaximumBitRateUL = BitRate()
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm int64 = 131
func PDUSessionResourceFailedToModifyListModCfm () (value ngapType.PDUSessionResourceFailedToModifyListModCfm) {
	item := PDUSessionResourceFailedToModifyItemModCfm()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail int64 = 132
func PDUSessionResourceFailedToSetupListCxtFail () (value ngapType.PDUSessionResourceFailedToSetupListCxtFail) {
	item := PDUSessionResourceFailedToSetupItemCxtFail()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionResourceListCxtRelReq int64 = 133
func PDUSessionResourceListCxtRelReq () (value ngapType.PDUSessionResourceListCxtRelReq) {
	item := PDUSessionResourceItemCxtRelReq()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDPDUSessionType int64 = 134
func PDUSessionType () (value ngapType.PDUSessionType) {
	value.Value = ngapType.PDUSessionTypePresentIpv4
	value.Value = ngapType.PDUSessionTypePresentIpv6
	value.Value = ngapType.PDUSessionTypePresentIpv4v6
	value.Value = ngapType.PDUSessionTypePresentEthernet
	value.Value = ngapType.PDUSessionTypePresentUnstructured
	return
}

// const ProtocolIEIDQosFlowAddOrModifyRequestList int64 = 135
func QosFlowAddOrModifyRequestList () (value ngapType.QosFlowAddOrModifyRequestList) {
	item := QosFlowAddOrModifyRequestItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDQosFlowSetupRequestList int64 = 136
func QosFlowSetupRequestList () (value ngapType.QosFlowSetupRequestList) {
	item := QosFlowSetupRequestItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDQosFlowToReleaseList int64 = 137
func QosFlowToReleaseList () (value ngapType.QosFlowList) {
	item := QosFlowItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDSecurityIndication int64 = 138
func SecurityIndication () (value ngapType.SecurityIndication) {
	value.IntegrityProtectionIndication = IntegrityProtectionIndication()
	value.ConfidentialityProtectionIndication = ConfidentialityProtectionIndication()
	value.MaximumIntegrityProtectedDataRate = new(ngapType.MaximumIntegrityProtectedDataRate)
	*value.MaximumIntegrityProtectedDataRate = MaximumIntegrityProtectedDataRate()
	return
}

// const ProtocolIEIDULNGUUPTNLInformation int64 = 139
func ULNGUUPTNLInformation () (value ngapType.UPTransportLayerInformation) {
	return UPTransportLayerInformation()
}

// const ProtocolIEIDULNGUUPTNLModifyList int64 = 140
func ULNGUUPTNLModifyList () (value ngapType.ULNGUUPTNLModifyList) {
	item := ULNGUUPTNLModifyItem()
	value.List = append(value.List, item, item, item)
	return
}

// const ProtocolIEIDWarningAreaCoordinates int64 = 141
func WarningAreaCoordinates () (value ngapType.WarningAreaCoordinates) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

////////////////////////////////////////////
/////// COMMON ////////////// COMMON ///////
////////////////////////////////////////////

func AllowedNSSAIItem () (value ngapType.AllowedNSSAIItem) {
	value.SNSSAI = SNSSAI()
	return
}

func SNSSAI () (value ngapType.SNSSAI) {
	value.SST = SST()
	value.SD = new(ngapType.SD)
	*value.SD = SD()
	return
}

func SST () (value ngapType.SST) {
	value.Value = []byte{171}
	return
}

func SD () (value ngapType.SD) {
	value.Value = []byte{183,137,255}
	return
}

func OverloadResponse () (value ngapType.OverloadResponse) {
	value.Present = ngapType.OverloadResponsePresentOverloadAction
	value.OverloadAction = new(ngapType.OverloadAction)
	*value.OverloadAction = OverloadAction()
	return
}

func OverloadAction () (value ngapType.OverloadAction) {
	value.Value = ngapType.OverloadActionPresentRejectNonEmergencyMoDt
	value.Value = ngapType.OverloadActionPresentRejectRrcCrSignalling
	value.Value = ngapType.OverloadActionPresentPermitEmergencySessionsAndMobileTerminatedServicesOnly
	value.Value = ngapType.OverloadActionPresentPermitHighPrioritySessionsAndMobileTerminatedServicesOnly
	return
}

func TNLAssociationList () (value ngapType.TNLAssociationList) {
	item := TNLAssociationItem()
	value.List = append(value.List, item, item, item)
	return
}

func TNLAssociationItem () (value ngapType.TNLAssociationItem) {
	value.TNLAssociationAddress = CPTransportLayerInformation()
	value.Cause = Cause()
	return
}

func CPTransportLayerInformation () (value ngapType.CPTransportLayerInformation) {
	value.Present = ngapType.CPTransportLayerInformationPresentEndpointIPAddress
	value.EndpointIPAddress = new(ngapType.TransportLayerAddress)
	*value.EndpointIPAddress = TransportLayerAddress()
	return
}

func TransportLayerAddress () (value ngapType.TransportLayerAddress) {
	value.Value.BitLength = 91 // 1-160
	value.Value.Bytes = []byte{171,00,11,22,33,44,55,66,77,88,99,224}
	return
}

func AMFTNLAssociationSetupItem () (value ngapType.AMFTNLAssociationSetupItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	return
}

func AMFTNLAssociationToAddItem () (value ngapType.AMFTNLAssociationToAddItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	value.TNLAssociationUsage = new(ngapType.TNLAssociationUsage)
	*value.TNLAssociationUsage = TNLAssociationUsage()
	value.TNLAddressWeightFactor = TNLAddressWeightFactor()
	return
}

func TNLAssociationUsage () (value ngapType.TNLAssociationUsage) {
	value.Value = ngapType.TNLAssociationUsagePresentUe
	value.Value = ngapType.TNLAssociationUsagePresentNonUe
	value.Value = ngapType.TNLAssociationUsagePresentBoth
	return
}

func TNLAddressWeightFactor () (value ngapType.TNLAddressWeightFactor) {
	value.Value = 137 // 0-255
	return
}

func AMFTNLAssociationToRemoveItem () (value ngapType.AMFTNLAssociationToRemoveItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	return
}

func AMFTNLAssociationToUpdateItem () (value ngapType.AMFTNLAssociationToUpdateItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	value.TNLAssociationUsage = new(ngapType.TNLAssociationUsage)
	*value.TNLAssociationUsage = TNLAssociationUsage()
	value.TNLAddressWeightFactor = new(ngapType.TNLAddressWeightFactor)
	*value.TNLAddressWeightFactor = TNLAddressWeightFactor()
	return
}

func AssistanceDataForRecommendedCells () (value ngapType.AssistanceDataForRecommendedCells) {
	value.RecommendedCellsForPaging = RecommendedCellsForPaging()
	return
}

func RecommendedCellsForPaging () (value ngapType.RecommendedCellsForPaging) {
	value.RecommendedCellList = RecommendedCellList()
	return
}

func RecommendedCellList () (value ngapType.RecommendedCellList) {
	item := RecommendedCellItem()
	value.List = append(value.List, item, item, item)
	return
}

func RecommendedCellItem () (value ngapType.RecommendedCellItem) {
	value.NGRANCGI = NGRANCGI()
	value.TimeStayedInCell = new(int64)
	*value.TimeStayedInCell = 4095 // 0-4095
	return
}

func PagingAttemptInformation () (value ngapType.PagingAttemptInformation) {
	value.PagingAttemptCount = PagingAttemptCount()
	value.IntendedNumberOfPagingAttempts = IntendedNumberOfPagingAttempts()
	value.NextPagingAreaScope = new(ngapType.NextPagingAreaScope)
	*value.NextPagingAreaScope = NextPagingAreaScope()
	return
}

func PagingAttemptCount () (value ngapType.PagingAttemptCount) {
	value.Value = 16 // 1-16
	return
}

func IntendedNumberOfPagingAttempts () (value ngapType.IntendedNumberOfPagingAttempts) {
	value.Value = 16 // 1-16
	return
}

func NextPagingAreaScope () (value ngapType.NextPagingAreaScope) {
	value.Value = ngapType.NextPagingAreaScopePresentSame
	value.Value = ngapType.NextPagingAreaScopePresentChanged
	return
}

func CellIDCancelledEUTRA () (value ngapType.CellIDCancelledEUTRA) {
	item := CellIDCancelledEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func CellIDCancelledEUTRAItem () (value ngapType.CellIDCancelledEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func NumberOfBroadcasts () (value ngapType.NumberOfBroadcasts) {
	value.Value = 65535 // 0-65535
	return
}

func TAICancelledEUTRA () (value ngapType.TAICancelledEUTRA) {
	item := TAICancelledEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func TAICancelledEUTRAItem () (value ngapType.TAICancelledEUTRAItem) {
	value.TAI = TAI()
	value.CancelledCellsInTAIEUTRA = CancelledCellsInTAIEUTRA()
	return
}

func TAI () (value ngapType.TAI) {
	value.PLMNIdentity = PLMNIdentity()
	value.TAC = TAC()
	return
}

func PLMNIdentity () (value ngapType.PLMNIdentity) {
	// 54f240
	value.Value = []byte{0x54,0xf2,0x40}
	return
}

func TAC () (value ngapType.TAC) {
	value.Value = []byte{171,137,255}
	return
}

func CancelledCellsInTAIEUTRA () (value ngapType.CancelledCellsInTAIEUTRA) {
	item := CancelledCellsInTAIEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func CancelledCellsInTAIEUTRAItem () (value ngapType.CancelledCellsInTAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func EmergencyAreaIDCancelledEUTRA () (value ngapType.EmergencyAreaIDCancelledEUTRA) {
	item := EmergencyAreaIDCancelledEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func EmergencyAreaIDCancelledEUTRAItem () (value ngapType.EmergencyAreaIDCancelledEUTRAItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CancelledCellsInEAIEUTRA = CancelledCellsInEAIEUTRA()
	return
}

func EmergencyAreaID () (value ngapType.EmergencyAreaID) {
	value.Value = []byte{171,137,255}
	return
}

func CancelledCellsInEAIEUTRA () (value ngapType.CancelledCellsInEAIEUTRA) {
	item := CancelledCellsInEAIEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func CancelledCellsInEAIEUTRAItem () (value ngapType.CancelledCellsInEAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CellIDCancelledNR () (value ngapType.CellIDCancelledNR) {
	item := CellIDCancelledNRItem()
	value.List = append(value.List, item, item, item)
	return
}

func CellIDCancelledNRItem () (value ngapType.CellIDCancelledNRItem) {
	value.NRCGI = NRCGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func TAICancelledNR () (value ngapType.TAICancelledNR) {
	item := TAICancelledNRItem()
	value.List = append(value.List, item, item, item)
	return
}

func TAICancelledNRItem () (value ngapType.TAICancelledNRItem) {
	value.TAI = TAI()
	value.CancelledCellsInTAINR = CancelledCellsInTAINR()
	return
}

func CancelledCellsInTAINR () (value ngapType.CancelledCellsInTAINR) {
	item := CancelledCellsInTAINRItem()
	value.List = append(value.List, item, item, item)
	return
}

func CancelledCellsInTAINRItem () (value ngapType.CancelledCellsInTAINRItem) {
	value.NRCGI = NRCGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func EmergencyAreaIDCancelledNR () (value ngapType.EmergencyAreaIDCancelledNR) {
	item := EmergencyAreaIDCancelledNRItem()
	value.List = append(value.List, item, item, item)
	return
}

func EmergencyAreaIDCancelledNRItem () (value ngapType.EmergencyAreaIDCancelledNRItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CancelledCellsInEAINR = CancelledCellsInEAINR()
	return
}

func CancelledCellsInEAINR () (value ngapType.CancelledCellsInEAINR) {
	item := CancelledCellsInEAINRItem()
	value.List = append(value.List, item, item, item)
	return
}

func CancelledCellsInEAINRItem () (value ngapType.CancelledCellsInEAINRItem) {
	value.NRCGI = NRCGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CellIDBroadcastEUTRA () (value ngapType.CellIDBroadcastEUTRA) {
	item := CellIDBroadcastEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func CellIDBroadcastEUTRAItem () (value ngapType.CellIDBroadcastEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	return
}

func TAIBroadcastEUTRA () (value ngapType.TAIBroadcastEUTRA) {
	item := TAIBroadcastEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func TAIBroadcastEUTRAItem () (value ngapType.TAIBroadcastEUTRAItem) {
	value.TAI = TAI()
	value.CompletedCellsInTAIEUTRA = CompletedCellsInTAIEUTRA()
	return
}

func CompletedCellsInTAIEUTRA () (value ngapType.CompletedCellsInTAIEUTRA) {
	item := CompletedCellsInTAIEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func CompletedCellsInTAIEUTRAItem () (value ngapType.CompletedCellsInTAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	return
}

func EmergencyAreaIDBroadcastEUTRA () (value ngapType.EmergencyAreaIDBroadcastEUTRA) {
	item := EmergencyAreaIDBroadcastEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func EmergencyAreaIDBroadcastEUTRAItem () (value ngapType.EmergencyAreaIDBroadcastEUTRAItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CompletedCellsInEAIEUTRA = CompletedCellsInEAIEUTRA()
	return
}

func CompletedCellsInEAIEUTRA () (value ngapType.CompletedCellsInEAIEUTRA) {
	item := CompletedCellsInEAIEUTRAItem()
	value.List = append(value.List, item, item, item)
	return
}

func CompletedCellsInEAIEUTRAItem () (value ngapType.CompletedCellsInEAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	return
}

func CellIDBroadcastNR () (value ngapType.CellIDBroadcastNR) {
	item := CellIDBroadcastNRItem()
	value.List = append(value.List, item, item, item)
	return
}

func CellIDBroadcastNRItem () (value ngapType.CellIDBroadcastNRItem) {
	value.NRCGI = NRCGI()
	return
}

func TAIBroadcastNR () (value ngapType.TAIBroadcastNR) {
	item := TAIBroadcastNRItem()
	value.List = append(value.List, item, item, item)
	return
}

func TAIBroadcastNRItem () (value ngapType.TAIBroadcastNRItem) {
	value.TAI = TAI()
	value.CompletedCellsInTAINR = CompletedCellsInTAINR()
	return
}

func CompletedCellsInTAINR () (value ngapType.CompletedCellsInTAINR) {
	item := CompletedCellsInTAINRItem()
	value.List = append(value.List, item, item, item)
	return
}

func CompletedCellsInTAINRItem () (value ngapType.CompletedCellsInTAINRItem) {
	value.NRCGI = NRCGI()
	return
}

func EmergencyAreaIDBroadcastNR () (value ngapType.EmergencyAreaIDBroadcastNR) {
	item := EmergencyAreaIDBroadcastNRItem()
	value.List = append(value.List, item, item, item)
	return
}

func EmergencyAreaIDBroadcastNRItem () (value ngapType.EmergencyAreaIDBroadcastNRItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CompletedCellsInEAINR = CompletedCellsInEAINR()
	return
}

func CompletedCellsInEAINR () (value ngapType.CompletedCellsInEAINR) {
	item := CompletedCellsInEAINRItem()
	value.List = append(value.List, item, item, item)
	return
}

func CompletedCellsInEAINRItem () (value ngapType.CompletedCellsInEAINRItem) {
	value.NRCGI = NRCGI()
	return
}

func CauseRadioNetwork () (value ngapType.CauseRadioNetwork) {
	value.Value = ngapType.CauseRadioNetworkPresentTxnrelocoverallExpiry
	value.Value = ngapType.CauseRadioNetworkPresentSuccessfulHandover
	value.Value = ngapType.CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason
	value.Value = ngapType.CauseRadioNetworkPresentReleaseDueTo5gcGeneratedReason
	value.Value = ngapType.CauseRadioNetworkPresentHandoverCancelled
	value.Value = ngapType.CauseRadioNetworkPresentPartialHandover
	value.Value = ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem
	value.Value = ngapType.CauseRadioNetworkPresentHoTargetNotAllowed
	value.Value = ngapType.CauseRadioNetworkPresentTngrelocoverallExpiry
	value.Value = ngapType.CauseRadioNetworkPresentTngrelocprepExpiry
	value.Value = ngapType.CauseRadioNetworkPresentCellNotAvailable
	value.Value = ngapType.CauseRadioNetworkPresentUnknownTargetID
	value.Value = ngapType.CauseRadioNetworkPresentNoRadioResourcesAvailableInTargetCell
	value.Value = ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID
	value.Value = ngapType.CauseRadioNetworkPresentInconsistentRemoteUENGAPID
	value.Value = ngapType.CauseRadioNetworkPresentHandoverDesirableForRadioReason
	value.Value = ngapType.CauseRadioNetworkPresentTimeCriticalHandover
	value.Value = ngapType.CauseRadioNetworkPresentResourceOptimisationHandover
	value.Value = ngapType.CauseRadioNetworkPresentReduceLoadInServingCell
	value.Value = ngapType.CauseRadioNetworkPresentUserInactivity
	value.Value = ngapType.CauseRadioNetworkPresentRadioConnectionWithUeLost
	value.Value = ngapType.CauseRadioNetworkPresentRadioResourcesNotAvailable
	value.Value = ngapType.CauseRadioNetworkPresentInvalidQosCombination
	value.Value = ngapType.CauseRadioNetworkPresentFailureInRadioInterfaceProcedure
	value.Value = ngapType.CauseRadioNetworkPresentInteractionWithOtherProcedure
	value.Value = ngapType.CauseRadioNetworkPresentUnknownPDUSessionID
	value.Value = ngapType.CauseRadioNetworkPresentUnkownQosFlowID
	value.Value = ngapType.CauseRadioNetworkPresentMultiplePDUSessionIDInstances
	value.Value = ngapType.CauseRadioNetworkPresentMultipleQosFlowIDInstances
	value.Value = ngapType.CauseRadioNetworkPresentEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported
	value.Value = ngapType.CauseRadioNetworkPresentNgIntraSystemHandoverTriggered
	value.Value = ngapType.CauseRadioNetworkPresentNgInterSystemHandoverTriggered
	value.Value = ngapType.CauseRadioNetworkPresentXnHandoverTriggered
	value.Value = ngapType.CauseRadioNetworkPresentNotSupported5QIValue
	value.Value = ngapType.CauseRadioNetworkPresentUeContextTransfer
	value.Value = ngapType.CauseRadioNetworkPresentImsVoiceEpsFallbackOrRatFallbackTriggered
	value.Value = ngapType.CauseRadioNetworkPresentUpIntegrityProtectionNotPossible
	value.Value = ngapType.CauseRadioNetworkPresentUpConfidentialityProtectionNotPossible
	value.Value = ngapType.CauseRadioNetworkPresentSliceNotSupported
	value.Value = ngapType.CauseRadioNetworkPresentUeInRrcInactiveStateNotReachable
	value.Value = ngapType.CauseRadioNetworkPresentRedirection
	value.Value = ngapType.CauseRadioNetworkPresentResourcesNotAvailableForTheSlice
	value.Value = ngapType.CauseRadioNetworkPresentUeMaxIntegrityProtectedDataRateReason
	value.Value = ngapType.CauseRadioNetworkPresentReleaseDueToCnDetectedMobility
	return
}

func CauseTransport () (value ngapType.CauseTransport) {
	value.Value = ngapType.CauseTransportPresentTransportResourceUnavailable
	value.Value = ngapType.CauseTransportPresentUnspecified
	return
}

func CauseNas () (value ngapType.CauseNas) {
	value.Value = ngapType.CauseNasPresentNormalRelease
	value.Value = ngapType.CauseNasPresentAuthenticationFailure
	value.Value = ngapType.CauseNasPresentDeregister
	value.Value = ngapType.CauseNasPresentUnspecified
	return
}

func CauseProtocol () (value ngapType.CauseProtocol) {
	value.Value = ngapType.CauseProtocolPresentTransferSyntaxError
	value.Value = ngapType. CauseProtocolPresentAbstractSyntaxErrorReject
	value.Value = ngapType. CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify
	value.Value = ngapType. CauseProtocolPresentMessageNotCompatibleWithReceiverState
	value.Value = ngapType. CauseProtocolPresentSemanticError
	value.Value = ngapType. CauseProtocolPresentAbstractSyntaxErrorFalselyConstructedMessage
	value.Value = ngapType. CauseProtocolPresentUnspecified
	return
}

func CauseMisc () (value ngapType.CauseMisc) {
	value.Value = ngapType.CauseMiscPresentControlProcessingOverload
	value.Value = ngapType.CauseMiscPresentNotEnoughUserPlaneProcessingResources
	value.Value = ngapType.CauseMiscPresentHardwareFailure
	value.Value = ngapType.CauseMiscPresentOmIntervention
	value.Value = ngapType.CauseMiscPresentUnknownPLMN
	value.Value = ngapType.CauseMiscPresentUnspecified
	return
}

func EUTRACGIList () (value ngapType.EUTRACGIList) {
	item := EUTRACGI()
	value.List = append(value.List, item, item, item)
	return
}

func NRCGIList () (value ngapType.NRCGIList) {
	item := NRCGI()
	value.List = append(value.List, item, item, item)
	return
}

func UEIdentityIndexValue () (value ngapType.UEIdentityIndexValue) {
	value.Present = ngapType.UEIdentityIndexValuePresentIndexLength10
	value.IndexLength10 = new(ngapType.BitString)
	value.IndexLength10.BitLength = 10
	value.IndexLength10.Bytes = []byte{171,255}
	return
}

func PeriodicRegistrationUpdateTimer () (value ngapType.PeriodicRegistrationUpdateTimer) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{183}
	return
}

func MICOModeIndication () (value ngapType.MICOModeIndication) {
	value.Value = ngapType.MICOModeIndicationPresentTrue
	return
}

func TAIListForInactive () (value ngapType.TAIListForInactive) {
	item := TAIListForInactiveItem()
	value.List = append(value.List, item, item, item)
	return
}

func TAIListForInactiveItem () (value ngapType.TAIListForInactiveItem) {
	value.TAI = TAI()
	return
}

func ExpectedUEBehaviour () (value ngapType.ExpectedUEBehaviour) {
	value.ExpectedUEActivityBehaviour = new(ngapType.ExpectedUEActivityBehaviour)
	*value.ExpectedUEActivityBehaviour = ExpectedUEActivityBehaviour()
	value.ExpectedHOInterval = new(ngapType.ExpectedHOInterval)
	*value.ExpectedHOInterval = ExpectedHOInterval()
	value.ExpectedUEMobility = new(ngapType.ExpectedUEMobility)
	*value.ExpectedUEMobility = ExpectedUEMobility()
	value.ExpectedUEMovingTrajectory = new(ngapType.ExpectedUEMovingTrajectory)
	*value.ExpectedUEMovingTrajectory = ExpectedUEMovingTrajectory()
	return
}

func ExpectedUEActivityBehaviour () (value ngapType.ExpectedUEActivityBehaviour) {
	value.ExpectedActivityPeriod = new(ngapType.ExpectedActivityPeriod)
	*value.ExpectedActivityPeriod = ExpectedActivityPeriod()
	value.ExpectedIdlePeriod = new(ngapType.ExpectedIdlePeriod)
	*value.ExpectedIdlePeriod = ExpectedIdlePeriod()
	value.SourceOfUEActivityBehaviourInformation = new(ngapType.SourceOfUEActivityBehaviourInformation)
	*value.SourceOfUEActivityBehaviourInformation = SourceOfUEActivityBehaviourInformation()
	return
}

func ExpectedActivityPeriod () (value ngapType.ExpectedActivityPeriod) {
	value.Value = 156
	return
}

func ExpectedIdlePeriod () (value ngapType.ExpectedIdlePeriod) {
	value.Value = 122
	return
}

func SourceOfUEActivityBehaviourInformation () (value ngapType.SourceOfUEActivityBehaviourInformation) {
	value.Value = ngapType.SourceOfUEActivityBehaviourInformationPresentSubscriptionInformation
	value.Value = ngapType.SourceOfUEActivityBehaviourInformationPresentStatistics
	return
}

func ExpectedHOInterval () (value ngapType.ExpectedHOInterval) {
	value.Value = ngapType.ExpectedHOIntervalPresentSec15
	value.Value = ngapType.ExpectedHOIntervalPresentSec30
	value.Value = ngapType.ExpectedHOIntervalPresentSec60
	value.Value = ngapType.ExpectedHOIntervalPresentSec90
	value.Value = ngapType.ExpectedHOIntervalPresentSec120
	value.Value = ngapType.ExpectedHOIntervalPresentSec180
	value.Value = ngapType.ExpectedHOIntervalPresentLongTime
	return
}

func ExpectedUEMobility () (value ngapType.ExpectedUEMobility) {
	value.Value = ngapType.ExpectedUEMobilityPresentStationary
	value.Value = ngapType.ExpectedUEMobilityPresentMobile
	return
}

func ExpectedUEMovingTrajectory () (value ngapType.ExpectedUEMovingTrajectory) {
	item := ExpectedUEMovingTrajectoryItem()
	value.List = append(value.List, item, item, item)
	return
}

func ExpectedUEMovingTrajectoryItem () (value ngapType.ExpectedUEMovingTrajectoryItem) {
	value.NGRANCGI = NGRANCGI()
	value.TimeStayedInCell = new(int64)
	*value.TimeStayedInCell = 4095 // 0-4095
	return
}

func ProcedureCode () (value ngapType.ProcedureCode) {
	value.Value = 137 // 0-255
	return
}

func TriggeringMessage () (value ngapType.TriggeringMessage) {
	value.Value = ngapType.TriggeringMessagePresentInitiatingMessage
	value.Value = ngapType.TriggeringMessagePresentSuccessfulOutcome
	value.Value = ngapType.TriggeringMessagePresentUnsuccessfullOutcome
	return
}

func Criticality () (value ngapType.Criticality) {
	value.Value = ngapType.CriticalityPresentReject
	value.Value = ngapType.CriticalityPresentIgnore
	value.Value = ngapType.CriticalityPresentNotify
	return
}

func CriticalityDiagnosticsIEList () (value ngapType.CriticalityDiagnosticsIEList) {
	item := CriticalityDiagnosticsIEItem()
	value.List = append(value.List, item, item, item)
	return
}

func CriticalityDiagnosticsIEItem () (value ngapType.CriticalityDiagnosticsIEItem) {
	value.IECriticality = Criticality()
	value.IEID = ProtocolIEID()
	value.TypeOfError = TypeOfError()
	return
}

func ProtocolIEID () (value ngapType.ProtocolIEID) {
	value.Value = 65535 // 0-65535
	return
}

func TypeOfError () (value ngapType.TypeOfError) {
	value.Value = ngapType.TypeOfErrorPresentNotUnderstood
	value.Value = ngapType.TypeOfErrorPresentMissing
	return
}

func EmergencyFallbackRequestIndicator () (value ngapType.EmergencyFallbackRequestIndicator) {
	value.Value = ngapType.EmergencyFallbackRequestIndicatorPresentEmergencyFallbackRequested
	return
}

func EmergencyServiceTargetCN () (value ngapType.EmergencyServiceTargetCN) {
	value.Value = ngapType.EmergencyServiceTargetCNPresentFiveGC
	value.Value = ngapType.EmergencyServiceTargetCNPresentEpc
	return
}

func EUTRACellIdentity () (value ngapType.EUTRACellIdentity) {
	value.Value.BitLength = 28
	value.Value.Bytes = []byte{171,11,22,240}
	return
}

func AMFPointer () (value ngapType.AMFPointer) {
	value.Value.BitLength = 6
	value.Value.Bytes = []byte{183}
	return
}

func FiveGTMSI () (value ngapType.FiveGTMSI) {
	value.Value = []byte{171,11,22,255}
	return
}

func GlobalGNBID () (value ngapType.GlobalGNBID) {
	value.PLMNIdentity = PLMNIdentity()
	value.GNBID = GNBID()
	return
}

func GNBID () (value ngapType.GNBID) {
	value.Present = ngapType.GNBIDPresentGNBID
	value.GNBID = new(ngapType.BitString)
	value.GNBID.BitLength = 29 // 22-32
	value.GNBID.Bytes = []byte{171,11,22,255}
	return
}

func GlobalNgENBID () (value ngapType.GlobalNgENBID) {
	value.PLMNIdentity = PLMNIdentity()
	value.NgENBID = NgENBID()
	return
}

func NgENBID () (value ngapType.NgENBID) {
	value.Present = ngapType.NgENBIDPresentMacroNgENBID
	value.MacroNgENBID = new(ngapType.BitString)
	value.MacroNgENBID.BitLength = 20
	value.MacroNgENBID.Bytes = []byte{171,11,255}
	value.Present = ngapType.NgENBIDPresentShortMacroNgENBID
	value.ShortMacroNgENBID = new(ngapType.BitString)
	value.ShortMacroNgENBID.BitLength = 18
	value.ShortMacroNgENBID.Bytes = []byte{171,11,255}
	value.Present = ngapType.NgENBIDPresentLongMacroNgENBID
	value.LongMacroNgENBID = new(ngapType.BitString)
	value.LongMacroNgENBID.BitLength = 21
	value.LongMacroNgENBID.Bytes = []byte{171,11,255}
	return
}

func GlobalN3IWFID () (value ngapType.GlobalN3IWFID) {
	value.PLMNIdentity = PLMNIdentity()
	value.N3IWFID = N3IWFID()
	return
}

func N3IWFID () (value ngapType.N3IWFID) {
	value.Present = ngapType.N3IWFIDPresentN3IWFID
	value.N3IWFID = new(ngapType.BitString)
	value.N3IWFID.BitLength = 16
	value.N3IWFID.Bytes = []byte{171,255}
	return
}

func AMFRegionID () (value ngapType.AMFRegionID) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{183}
	return
}

func RecommendedRANNodesForPaging () (value ngapType.RecommendedRANNodesForPaging) {
	value.RecommendedRANNodeList = RecommendedRANNodeList()
	return
}

func RecommendedRANNodeList () (value ngapType.RecommendedRANNodeList) {
	item := RecommendedRANNodeItem()
	value.List = append(value.List, item, item, item)
	return
}

func RecommendedRANNodeItem () (value ngapType.RecommendedRANNodeItem) {
	value.AMFPagingTarget = AMFPagingTarget()
	return
}

func AMFPagingTarget () (value ngapType.AMFPagingTarget) {
	value.Present = ngapType.AMFPagingTargetPresentGlobalRANNodeID
	value.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*value.GlobalRANNodeID = GlobalRANNodeID()
	value.Present = ngapType.AMFPagingTargetPresentTAI
	value.TAI = new(ngapType.TAI)
	*value.TAI = TAI()
	return
}

func EventType () (value ngapType.EventType) {
	value.Value = ngapType.EventTypePresentDirect
	value.Value = ngapType.EventTypePresentChangeOfServeCell
	value.Value = ngapType.EventTypePresentUePresenceInAreaOfInterest
	value.Value = ngapType.EventTypePresentStopChangeOfServeCell
	value.Value = ngapType.EventTypePresentStopUePresenceInAreaOfInterest
	value.Value = ngapType.EventTypePresentCancelLocationReportingForTheUe
	return
}

func ReportArea () (value ngapType.ReportArea) {
	value.Value = ngapType.ReportAreaPresentCell
	return
}

func AreaOfInterestList () (value ngapType.AreaOfInterestList) {
	item := AreaOfInterestItem()
	value.List = append(value.List, item, item, item)
	return
}

func AreaOfInterestItem () (value ngapType.AreaOfInterestItem) {
	value.AreaOfInterest = AreaOfInterest()
	value.LocationReportingReferenceID = LocationReportingReferenceID()
	return
}

func AreaOfInterest () (value ngapType.AreaOfInterest) {
	value.AreaOfInterestTAIList = new(ngapType.AreaOfInterestTAIList)
	*value.AreaOfInterestTAIList = AreaOfInterestTAIList()
	value.AreaOfInterestCellList = new(ngapType.AreaOfInterestCellList)
	*value.AreaOfInterestCellList = AreaOfInterestCellList()
	value.AreaOfInterestRANNodeList = new(ngapType.AreaOfInterestRANNodeList)
	*value.AreaOfInterestRANNodeList = AreaOfInterestRANNodeList()
	return
}

func AreaOfInterestTAIList () (value ngapType.AreaOfInterestTAIList) {
	item := AreaOfInterestTAIItem()
	value.List = append(value.List, item, item, item)
	return
}

func AreaOfInterestTAIItem () (value ngapType.AreaOfInterestTAIItem) {
	value.TAI = TAI()
	return
}

func AreaOfInterestCellList () (value ngapType.AreaOfInterestCellList) {
	item := AreaOfInterestCellItem()
	value.List = append(value.List, item, item, item)
	return
}

func AreaOfInterestCellItem () (value ngapType.AreaOfInterestCellItem) {
	value.NGRANCGI = NGRANCGI()
	return
}

func AreaOfInterestRANNodeList () (value ngapType.AreaOfInterestRANNodeList) {
	item := AreaOfInterestRANNodeItem()
	value.List = append(value.List, item, item, item)
	return
}

func AreaOfInterestRANNodeItem () (value ngapType.AreaOfInterestRANNodeItem) {
	value.GlobalRANNodeID = GlobalRANNodeID()
	return
}

func LocationReportingReferenceID () (value ngapType.LocationReportingReferenceID) {
	value.Value = 63 // 1-64
	return
}

func EquivalentPLMNs () (value ngapType.EquivalentPLMNs) {
	item := PLMNIdentity()
	value.List = append(value.List, item, item, item)
	return
}

func RATRestrictions () (value ngapType.RATRestrictions) {
	item := RATRestrictionsItem()
	value.List = append(value.List, item, item, item, item, item, item, item, item, item, item, item, item, item, item, item, item)
	return
}

func RATRestrictionsItem () (value ngapType.RATRestrictionsItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.RATRestrictionInformation = RATRestrictionInformation()
	return
}

func RATRestrictionInformation () (value ngapType.RATRestrictionInformation) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{183}
	return
}

func ForbiddenAreaInformation () (value ngapType.ForbiddenAreaInformation) {
	item := ForbiddenAreaInformationItem()
	value.List = append(value.List, item, item, item)
	return
}

func ForbiddenAreaInformationItem () (value ngapType.ForbiddenAreaInformationItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.ForbiddenTACs = ForbiddenTACs()
	return
}

func ForbiddenTACs () (value ngapType.ForbiddenTACs) {
	item := TAC()
	value.List = append(value.List, item, item, item)
	return
}

func ServiceAreaInformation () (value ngapType.ServiceAreaInformation) {
	item := ServiceAreaInformationItem()
	value.List = append(value.List, item, item, item)
	return
}

func ServiceAreaInformationItem () (value ngapType.ServiceAreaInformationItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.AllowedTACs = new(ngapType.AllowedTACs)
	*value.AllowedTACs = AllowedTACs()
	value.NotAllowedTACs = new(ngapType.NotAllowedTACs)
	*value.NotAllowedTACs = NotAllowedTACs()
	return
}

func AllowedTACs () (value ngapType.AllowedTACs) {
	item := TAC()
	value.List = append(value.List, item, item, item)
	return
}

func NotAllowedTACs () (value ngapType.NotAllowedTACs) {
	item := TAC()
	value.List = append(value.List, item, item, item)
	return
}

func NRCellIdentity () (value ngapType.NRCellIdentity) {
	value.Value.BitLength = 36
	value.Value.Bytes = []byte{171,11,22,33,240}
	return
}

func OverloadStartNSSAIItem () (value ngapType.OverloadStartNSSAIItem) {
	value.SliceOverloadList = SliceOverloadList()
	value.SliceOverloadResponse = new(ngapType.OverloadResponse)
	*value.SliceOverloadResponse = OverloadResponse()
	value.SliceTrafficLoadReductionIndication = new(ngapType.TrafficLoadReductionIndication)
	*value.SliceTrafficLoadReductionIndication = TrafficLoadReductionIndication()
	return
}

func SliceOverloadList () (value ngapType.SliceOverloadList) {
	item := SliceOverloadItem()
	value.List = append(value.List, item, item, item)
	return
}

func SliceOverloadItem () (value ngapType.SliceOverloadItem) {
	value.SNSSAI = SNSSAI()
	return
}

func TrafficLoadReductionIndication () (value ngapType.TrafficLoadReductionIndication) {
	value.Value = 99 // 1-99
	return
}

func PDUSessionResourceAdmittedItem () (value ngapType.PDUSessionResourceAdmittedItem) {
	value.PDUSessionID = PDUSessionID()
	value.HandoverRequestAcknowledgeTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionID () (value ngapType.PDUSessionID) {
	value.Value = 183
	return
}

func PDUSessionResourceFailedToModifyItemModRes () (value ngapType.PDUSessionResourceFailedToModifyItemModRes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceModifyUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceFailedToSetupItemCxtRes () (value ngapType.PDUSessionResourceFailedToSetupItemCxtRes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceSetupUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceFailedToSetupItemHOAck () (value ngapType.PDUSessionResourceFailedToSetupItemHOAck) {
	value.PDUSessionID = PDUSessionID()
	value.HandoverResourceAllocationUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceFailedToSetupItemPSReq () (value ngapType.PDUSessionResourceFailedToSetupItemPSReq) {
	value.PDUSessionID = PDUSessionID()
	value.PathSwitchRequestSetupFailedTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceFailedToSetupItemSURes () (value ngapType.PDUSessionResourceFailedToSetupItemSURes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceSetupUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceHandoverItem () (value ngapType.PDUSessionResourceHandoverItem) {
	value.PDUSessionID = PDUSessionID()
	value.HandoverCommandTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceItemCxtRelCpl () (value ngapType.PDUSessionResourceItemCxtRelCpl) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceItemHORqd () (value ngapType.PDUSessionResourceItemHORqd) {
	value.PDUSessionID = PDUSessionID()
	value.HandoverRequiredTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceModifyItemModCfm () (value ngapType.PDUSessionResourceModifyItemModCfm) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceModifyConfirmTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceModifyItemModInd () (value ngapType.PDUSessionResourceModifyItemModInd) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceModifyIndicationTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceModifyItemModReq () (value ngapType.PDUSessionResourceModifyItemModReq) {
	value.PDUSessionID = PDUSessionID()
	value.NASPDU = new(ngapType.NASPDU)
	*value.NASPDU = NASPDU()
	value.PDUSessionResourceModifyRequestTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceModifyItemModRes () (value ngapType.PDUSessionResourceModifyItemModRes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceModifyResponseTransfer = new(ngapType.OctetString)
	*value.PDUSessionResourceModifyResponseTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceNotifyItem () (value ngapType.PDUSessionResourceNotifyItem) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceNotifyTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceReleasedItemNot () (value ngapType.PDUSessionResourceReleasedItemNot) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceNotifyReleasedTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceReleasedItemPSAck () (value ngapType.PDUSessionResourceReleasedItemPSAck) {
	value.PDUSessionID = PDUSessionID()
	value.PathSwitchRequestUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceReleasedItemPSFail () (value ngapType.PDUSessionResourceReleasedItemPSFail) {
	value.PDUSessionID = PDUSessionID()
	value.PathSwitchRequestUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceReleasedItemRelRes () (value ngapType.PDUSessionResourceReleasedItemRelRes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceReleaseResponseTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceSetupItemCxtReq () (value ngapType.PDUSessionResourceSetupItemCxtReq) {
	value.PDUSessionID = PDUSessionID()
	value.NASPDU = new(ngapType.NASPDU)
	*value.NASPDU = NASPDU()
	value.SNSSAI = SNSSAI()
	value.PDUSessionResourceSetupRequestTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceSetupItemCxtRes () (value ngapType.PDUSessionResourceSetupItemCxtRes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceSetupResponseTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceSetupItemHOReq () (value ngapType.PDUSessionResourceSetupItemHOReq) {
	value.PDUSessionID = PDUSessionID()
	value.SNSSAI = SNSSAI()
	value.HandoverRequestTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceSetupItemSUReq () (value ngapType.PDUSessionResourceSetupItemSUReq) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionNASPDU = new(ngapType.NASPDU)
	*value.PDUSessionNASPDU = NASPDU()
	value.SNSSAI = SNSSAI()
	value.PDUSessionResourceSetupRequestTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceSetupItemSURes () (value ngapType.PDUSessionResourceSetupItemSURes) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceSetupResponseTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceToBeSwitchedDLItem () (value ngapType.PDUSessionResourceToBeSwitchedDLItem) {
	value.PDUSessionID = PDUSessionID()
	value.PathSwitchRequestTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceSwitchedItem () (value ngapType.PDUSessionResourceSwitchedItem) {
	value.PDUSessionID = PDUSessionID()
	value.PathSwitchRequestAcknowledgeTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceToReleaseItemHOCmd () (value ngapType.PDUSessionResourceToReleaseItemHOCmd) {
	value.PDUSessionID = PDUSessionID()
	value.HandoverPreparationUnsuccessfulTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PDUSessionResourceToReleaseItemRelCmd () (value ngapType.PDUSessionResourceToReleaseItemRelCmd) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceReleaseCommandTransfer = []byte{0, 3, 224, 172, 16, 108, 181, 0, 1, 134, 160, 0, 7}
	return
}

func PLMNSupportItem () (value ngapType.PLMNSupportItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.SliceSupportList = SliceSupportList()
	return
}

func DRBsSubjectToStatusTransferList () (value ngapType.DRBsSubjectToStatusTransferList) {
	item := DRBsSubjectToStatusTransferItem()
	value.List = append(value.List, item, item, item)
	return
}

func DRBsSubjectToStatusTransferItem () (value ngapType.DRBsSubjectToStatusTransferItem) {
	value.DRBID = DRBID()
	value.DRBStatusUL = DRBStatusUL()
	value.DRBStatusDL = DRBStatusDL()
	return
}

func DRBID () (value ngapType.DRBID) {
	value.Value = 29 // 1-32
	return
}

func DRBStatusUL () (value ngapType.DRBStatusUL) {
	value.Present = ngapType.DRBStatusULPresentDRBStatusUL12
	value.DRBStatusUL12 = new(ngapType.DRBStatusUL12)
	*value.DRBStatusUL12 = DRBStatusUL12()
	value.Present = ngapType.DRBStatusULPresentDRBStatusUL18
	value.DRBStatusUL18 = new(ngapType.DRBStatusUL18)
	*value.DRBStatusUL18 = DRBStatusUL18()
	return
}

func DRBStatusUL12 () (value ngapType.DRBStatusUL12) {
	value.ULCOUNTValue = COUNTValueForPDCPSN12()
	value.ReceiveStatusOfULPDCPSDUs = new(ngapType.BitString)
	value.ReceiveStatusOfULPDCPSDUs.BitLength = 513 // 1-2048
	value.ReceiveStatusOfULPDCPSDUs.Bytes = []byte{171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255}
	return
}

func COUNTValueForPDCPSN12 () (value ngapType.COUNTValueForPDCPSN12) {
	value.PDCPSN12 = 4095 // 0-4095
	value.HFNPDCPSN12 = 1048575 // 0-1048575
	return
}

func DRBStatusUL18 () (value ngapType.DRBStatusUL18) {
	value.ULCOUNTValue = COUNTValueForPDCPSN18()
	value.ReceiveStatusOfULPDCPSDUs = new(ngapType.BitString)
	value.ReceiveStatusOfULPDCPSDUs.BitLength = 513 // 1-131072
	value.ReceiveStatusOfULPDCPSDUs.Bytes = []byte{171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255,171,11,22,33,44,55,66,77,88,99,255}
	return
}

func COUNTValueForPDCPSN18 () (value ngapType.COUNTValueForPDCPSN18) {
	value.PDCPSN18 = 262143 // 0-262143
	value.HFNPDCPSN18 = 16383 // 0-16383
	return
}

func DRBStatusDL () (value ngapType.DRBStatusDL) {
	value.Present = ngapType.DRBStatusDLPresentDRBStatusDL12
	value.DRBStatusDL12 = new(ngapType.DRBStatusDL12)
	*value.DRBStatusDL12 = DRBStatusDL12()
	value.Present = ngapType.DRBStatusDLPresentDRBStatusDL18
	value.DRBStatusDL18 = new(ngapType.DRBStatusDL18)
	*value.DRBStatusDL18 = DRBStatusDL18()
	return
}

func DRBStatusDL12 () (value ngapType.DRBStatusDL12) {
	value.DLCOUNTValue = COUNTValueForPDCPSN12()
	return
}

func DRBStatusDL18 () (value ngapType.DRBStatusDL18) {
	value.DLCOUNTValue = COUNTValueForPDCPSN18()
	return
}

func ResetAll () (value ngapType.ResetAll) {
	value.Value = ngapType.ResetAllPresentResetAll
	return
}

func NextHopChainingCount () (value ngapType.NextHopChainingCount) {
	value.Value = 5 // 0-7
	return
}

func ServedGUAMIItem () (value ngapType.ServedGUAMIItem) {
	value.GUAMI = GUAMI()
	//value.BackupAMFName = new(ngapType.AMFName)
	//*value.BackupAMFName = AMFName()
	return
}

func SliceSupportItem () (value ngapType.SliceSupportItem) {
	value.SNSSAI = SNSSAI()
	return
}

func SONConfigurationTransfer () (value ngapType.SONConfigurationTransfer) {
	value.TargetRANNodeID = TargetRANNodeID()
	value.SourceRANNodeID = SourceRANNodeID()
	value.SONInformation = SONInformation()
	value.XnTNLConfigurationInfo = XnTNLConfigurationInfo()
	return
}

func TargetRANNodeID () (value ngapType.TargetRANNodeID) {
	value.GlobalRANNodeID = GlobalRANNodeID()
	value.SelectedTAI = TAI()
	return
}

func SourceRANNodeID () (value ngapType.SourceRANNodeID) {
	value.GlobalRANNodeID = GlobalRANNodeID()
	value.SelectedTAI = TAI()
	return
}

func SONInformation () (value ngapType.SONInformation) {
	value.Present = ngapType.SONInformationPresentSONInformationRequest
	value.SONInformationRequest = new(ngapType.SONInformationRequest)
	*value.SONInformationRequest = SONInformationRequest()
	value.Present = ngapType.SONInformationPresentSONInformationReply
	value.SONInformationReply = new(ngapType.SONInformationReply)
	*value.SONInformationReply = SONInformationReply()
	return
}

func SONInformationRequest () (value ngapType.SONInformationRequest) {
	value.Value = ngapType.SONInformationRequestPresentXnTNLConfigurationInfo
	return
}

func SONInformationReply () (value ngapType.SONInformationReply) {
	value.XnTNLConfigurationInfo = new(ngapType.XnTNLConfigurationInfo)
	*value.XnTNLConfigurationInfo = XnTNLConfigurationInfo()
	return
}

func XnTNLConfigurationInfo () (value ngapType.XnTNLConfigurationInfo) {
	value.XnTransportLayerAddresses = XnTLAs()
	value.XnExtendedTransportLayerAddresses = new(ngapType.XnExtTLAs)
	*value.XnExtendedTransportLayerAddresses = XnExtTLAs()
	return
}

func XnTLAs () (value ngapType.XnTLAs) {
	item := TransportLayerAddress()
	value.List = append(value.List, item, item, item)
	return
}

func XnExtTLAs () (value ngapType.XnExtTLAs) {
	item := XnExtTLAItem()
	value.List = append(value.List, item, item, item)
	return
}

func XnExtTLAItem () (value ngapType.XnExtTLAItem) {
	value.IPsecTLA = new(ngapType.TransportLayerAddress)
	*value.IPsecTLA = TransportLayerAddress()
	value.GTPTLAs = new(ngapType.XnGTPTLAs)
	*value.GTPTLAs = XnGTPTLAs()
	return
}

func XnGTPTLAs () (value ngapType.XnGTPTLAs) {
	item := TransportLayerAddress()
	value.List = append(value.List, item, item, item)
	return
}

func SupportedTAItem () (value ngapType.SupportedTAItem) {
	value.TAC = TAC()
	value.BroadcastPLMNList = BroadcastPLMNList()
	return
}

func BroadcastPLMNList () (value ngapType.BroadcastPLMNList) {
	item := BroadcastPLMNItem()
	value.List = append(value.List, item, item, item)
	return
}

func BroadcastPLMNItem () (value ngapType.BroadcastPLMNItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.TAISliceSupportList = SliceSupportList()
	return
}

func TAIListForPagingItem () (value ngapType.TAIListForPagingItem) {
	value.TAI = TAI()
	return
}

func TargeteNBID () (value ngapType.TargeteNBID) {
	value.GlobalENBID = GlobalNgENBID()
	value.SelectedEPSTAI = EPSTAI()
	return
}

func EPSTAI () (value ngapType.EPSTAI) {
	value.PLMNIdentity = PLMNIdentity()
	value.EPSTAC = EPSTAC()
	return
}

func EPSTAC () (value ngapType.EPSTAC) {
	value.Value = []byte{171,183}
	return
}

func InterfacesToTrace () (value ngapType.InterfacesToTrace) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{171}
	return
}

func TraceDepth () (value ngapType.TraceDepth) {
	value.Value = ngapType.TraceDepthPresentMinimum
	value.Value = ngapType.TraceDepthPresentMedium
	value.Value = ngapType.TraceDepthPresentMaximum
	value.Value = ngapType.TraceDepthPresentMinimumWithoutVendorSpecificExtension
	value.Value = ngapType.TraceDepthPresentMediumWithoutVendorSpecificExtension
	value.Value = ngapType.TraceDepthPresentMaximumWithoutVendorSpecificExtension
	return
}

func BitRate () (value ngapType.BitRate) {
	value.Value = 4000000000000 // 0-4000000000000
	return
}

func UEAssociatedLogicalNGConnectionItem () (value ngapType.UEAssociatedLogicalNGConnectionItem) {
	value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*value.AMFUENGAPID = AMFUENGAPID()
	value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*value.RANUENGAPID = RANUENGAPID()
	return
}

func UENGAPIDPair () (value ngapType.UENGAPIDPair) {
	value.AMFUENGAPID = AMFUENGAPID()
	value.RANUENGAPID = RANUENGAPID()
	return
}

func UEPresenceInAreaOfInterestItem () (value ngapType.UEPresenceInAreaOfInterestItem) {
	value.LocationReportingReferenceID = LocationReportingReferenceID()
	value.UEPresence = UEPresence()
	return
}

func UEPresence () (value ngapType.UEPresence) {
	value.Value = ngapType.UEPresencePresentIn
	value.Value = ngapType.UEPresencePresentOut
	value.Value = ngapType.UEPresencePresentUnknown
	return
}

func UERadioCapabilityForPagingOfNR () (value ngapType.UERadioCapabilityForPagingOfNR) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

func UERadioCapabilityForPagingOfEUTRA () (value ngapType.UERadioCapabilityForPagingOfEUTRA) {
	value.Value = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

func NRencryptionAlgorithms () (value ngapType.NRencryptionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func NRintegrityProtectionAlgorithms () (value ngapType.NRintegrityProtectionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func EUTRAencryptionAlgorithms () (value ngapType.EUTRAencryptionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func EUTRAintegrityProtectionAlgorithms () (value ngapType.EUTRAintegrityProtectionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func UnavailableGUAMIItem () (value ngapType.UnavailableGUAMIItem) {
	value.GUAMI = GUAMI()
	value.TimerApproachForGUAMIRemoval = new(ngapType.TimerApproachForGUAMIRemoval)
	*value.TimerApproachForGUAMIRemoval = TimerApproachForGUAMIRemoval()
	value.BackupAMFName = new(ngapType.AMFName)
	*value.BackupAMFName = AMFName()
	return
}

func TimerApproachForGUAMIRemoval () (value ngapType.TimerApproachForGUAMIRemoval) {
	value.Value = ngapType.TimerApproachForGUAMIRemovalPresentApplyTimer
	return
}

func UserLocationInformationEUTRA () (value ngapType.UserLocationInformationEUTRA) {
	value.EUTRACGI = EUTRACGI()
	value.TAI = TAI()
	value.TimeStamp = new(ngapType.TimeStamp)
	*value.TimeStamp = TimeStamp()
	return
}

func TimeStamp () (value ngapType.TimeStamp) {
	value.Value = []byte{171,11,22,255}
	return
}

func UserLocationInformationNR () (value ngapType.UserLocationInformationNR) {
	value.NRCGI = NRCGI()
	value.TAI = TAI()
	value.TimeStamp = new(ngapType.TimeStamp)
	*value.TimeStamp = TimeStamp()
	return
}

func UserLocationInformationN3IWF () (value ngapType.UserLocationInformationN3IWF) {
	value.IPAddress = TransportLayerAddress()
	value.PortNumber = PortNumber()
	return
}

func PortNumber () (value ngapType.PortNumber) {
	value.Value = []byte{171,255}
	return
}

func EUTRACGIListForWarning () (value ngapType.EUTRACGIListForWarning) {
	item := EUTRACGI()
	value.List = append(value.List, item, item, item)
	return
}

func NRCGIListForWarning () (value ngapType.NRCGIListForWarning) {
	item := NRCGI()
	value.List = append(value.List, item, item, item)
	return
}

func TAIListForWarning () (value ngapType.TAIListForWarning) {
	item := TAI()
	value.List = append(value.List, item, item, item)
	return
}

func EmergencyAreaIDList () (value ngapType.EmergencyAreaIDList) {
	item := EmergencyAreaID()
	value.List = append(value.List, item, item, item)
	return
}

func UPTransportLayerInformation () (value ngapType.UPTransportLayerInformation) {
	value.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	value.GTPTunnel = new(ngapType.GTPTunnel)
	*value.GTPTunnel = GTPTunnel()
	return
}

func GTPTunnel () (value ngapType.GTPTunnel) {
	value.TransportLayerAddress = TransportLayerAddress()
	value.GTPTEID = GTPTEID()
	return
}

func GTPTEID () (value ngapType.GTPTEID) {
	value.Value = []byte{171,11,22,255}
	return
}

func PDUSessionResourceFailedToModifyItemModCfm () (value ngapType.PDUSessionResourceFailedToModifyItemModCfm) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceModifyIndicationUnsuccessfulTransfer = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

func PDUSessionResourceFailedToSetupItemCxtFail () (value ngapType.PDUSessionResourceFailedToSetupItemCxtFail) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionResourceSetupUnsuccessfulTransfer = []byte{171,11,22,33,44,55,66,77,88,99,255}
	return
}

func PDUSessionResourceItemCxtRelReq () (value ngapType.PDUSessionResourceItemCxtRelReq) {
	value.PDUSessionID = PDUSessionID()
	return
}

func QosFlowAddOrModifyRequestItem () (value ngapType.QosFlowAddOrModifyRequestItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.QosFlowLevelQosParameters = new(ngapType.QosFlowLevelQosParameters)
	*value.QosFlowLevelQosParameters = QosFlowLevelQosParameters()
	value.ERABID = new(ngapType.ERABID)
	*value.ERABID = ERABID()
	return
}

func QosFlowIdentifier () (value ngapType.QosFlowIdentifier) {
	value.Value = 63 // 0-63
	return
}

func QosFlowLevelQosParameters () (value ngapType.QosFlowLevelQosParameters) {
	value.QosCharacteristics = QosCharacteristics()
	value.AllocationAndRetentionPriority = AllocationAndRetentionPriority()
	value.GBRQosInformation = new(ngapType.GBRQosInformation)
	*value.GBRQosInformation = GBRQosInformation()
	value.ReflectiveQosAttribute = new(ngapType.ReflectiveQosAttribute)
	*value.ReflectiveQosAttribute = ReflectiveQosAttribute()
	value.AdditionalQosFlowInformation = new(ngapType.AdditionalQosFlowInformation)
	*value.AdditionalQosFlowInformation = AdditionalQosFlowInformation()
	return
}

func QosCharacteristics () (value ngapType.QosCharacteristics) {
	value.Present = ngapType.QosCharacteristicsPresentNonDynamic5QI
	value.NonDynamic5QI = new(ngapType.NonDynamic5QIDescriptor)
	*value.NonDynamic5QI = NonDynamic5QIDescriptor()
	value.Present = ngapType.QosCharacteristicsPresentDynamic5QI
	value.Dynamic5QI = new(ngapType.Dynamic5QIDescriptor)
	*value.Dynamic5QI = Dynamic5QIDescriptor()
	return
}

func NonDynamic5QIDescriptor () (value ngapType.NonDynamic5QIDescriptor) {
	value.FiveQI = FiveQI()
	value.PriorityLevelQos = new(ngapType.PriorityLevelQos)
	*value.PriorityLevelQos = PriorityLevelQos()
	value.AveragingWindow = new(ngapType.AveragingWindow)
	*value.AveragingWindow = AveragingWindow()
	value.MaximumDataBurstVolume = new(ngapType.MaximumDataBurstVolume)
	*value.MaximumDataBurstVolume = MaximumDataBurstVolume()
	return
}

func FiveQI () (value ngapType.FiveQI) {
	value.Value = 137 // 0-255
	return
}

func PriorityLevelQos () (value ngapType.PriorityLevelQos) {
	value.Value = 127 // 1-127
	return
}

func AveragingWindow () (value ngapType.AveragingWindow) {
	value.Value = 4095 // 0-4095
	return
}

func MaximumDataBurstVolume () (value ngapType.MaximumDataBurstVolume) {
	value.Value = 4095 // 0-4095
	return
}

func Dynamic5QIDescriptor () (value ngapType.Dynamic5QIDescriptor) {
	value.PriorityLevelQos = PriorityLevelQos()
	value.PacketDelayBudget = PacketDelayBudget()
	value.PacketErrorRate = PacketErrorRate()
	value.FiveQI = new(ngapType.FiveQI)
	*value.FiveQI = FiveQI()
	value.DelayCritical = new(ngapType.DelayCritical)
	*value.DelayCritical = DelayCritical()
	value.AveragingWindow = new(ngapType.AveragingWindow)
	*value.AveragingWindow = AveragingWindow()
	value.MaximumDataBurstVolume = new(ngapType.MaximumDataBurstVolume)
	*value.MaximumDataBurstVolume = MaximumDataBurstVolume()
	return
}

func PacketDelayBudget () (value ngapType.PacketDelayBudget) {
	value.Value = 1023 // 0-1023
	return
}

func PacketErrorRate () (value ngapType.PacketErrorRate) {
	value.PERScalar = 9 // 0-9
	value.PERExponent = 9 // 0-9
	return
}

func DelayCritical () (value ngapType.DelayCritical) {
	value.Value = ngapType.DelayCriticalPresentDelayCritical
	value.Value = ngapType.DelayCriticalPresentNonDelayCritical
	return
}

func AllocationAndRetentionPriority () (value ngapType.AllocationAndRetentionPriority) {
	value.PriorityLevelARP = PriorityLevelARP()
	value.PreEmptionCapability = PreEmptionCapability()
	value.PreEmptionVulnerability = PreEmptionVulnerability()
	return
}

func PriorityLevelARP () (value ngapType.PriorityLevelARP) {
	value.Value = 15 // 1-15
	return
}

func PreEmptionCapability () (value ngapType.PreEmptionCapability) {
	value.Value = ngapType.PreEmptionCapabilityPresentShallNotTriggerPreEmption
	value.Value = ngapType.PreEmptionCapabilityPresentMayTriggerPreEmption
	return
}

func PreEmptionVulnerability () (value ngapType.PreEmptionVulnerability) {
	value.Value = ngapType.PreEmptionVulnerabilityPresentNotPreEmptable
	value.Value = ngapType.PreEmptionVulnerabilityPresentPreEmptable
	return
}

func GBRQosInformation () (value ngapType.GBRQosInformation) {
	value.MaximumFlowBitRateDL = BitRate()
	value.MaximumFlowBitRateUL = BitRate()
	value.GuaranteedFlowBitRateDL = BitRate()
	value.GuaranteedFlowBitRateUL = BitRate()
	value.NotificationControl = new(ngapType.NotificationControl)
	*value.NotificationControl = NotificationControl()
	value.MaximumPacketLossRateDL = new(ngapType.PacketLossRate)
	*value.MaximumPacketLossRateDL = PacketLossRate()
	value.MaximumPacketLossRateUL = new(ngapType.PacketLossRate)
	*value.MaximumPacketLossRateUL = PacketLossRate()
	return
}

func NotificationControl () (value ngapType.NotificationControl) {
	value.Value = ngapType.NotificationControlPresentNotificationRequested
	return
}

func PacketLossRate () (value ngapType.PacketLossRate) {
	value.Value = 999 // 0-1000
	return
}

func ReflectiveQosAttribute () (value ngapType.ReflectiveQosAttribute) {
	value.Value = ngapType.ReflectiveQosAttributePresentSubjectTo
	return
}

func AdditionalQosFlowInformation () (value ngapType.AdditionalQosFlowInformation) {
	value.Value = ngapType.AdditionalQosFlowInformationPresentMoreLikely
	return
}

func ERABID () (value ngapType.ERABID) {
	value.Value = 15 // 0-15
	return
}

func QosFlowSetupRequestItem () (value ngapType.QosFlowSetupRequestItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.QosFlowLevelQosParameters = QosFlowLevelQosParameters()
	value.ERABID = new(ngapType.ERABID)
	*value.ERABID = ERABID()
	return
}

func QosFlowItem () (value ngapType.QosFlowItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.Cause = Cause()
	return
}

func IntegrityProtectionIndication () (value ngapType.IntegrityProtectionIndication) {
	value.Value = ngapType.IntegrityProtectionIndicationPresentRequired
	value.Value = ngapType.IntegrityProtectionIndicationPresentPreferred
	value.Value = ngapType.IntegrityProtectionIndicationPresentNotNeeded
	return
}

func ConfidentialityProtectionIndication () (value ngapType.ConfidentialityProtectionIndication) {
	value.Value = ngapType.ConfidentialityProtectionIndicationPresentRequired
	value.Value = ngapType.ConfidentialityProtectionIndicationPresentPreferred
	value.Value = ngapType.ConfidentialityProtectionIndicationPresentNotNeeded
	return
}

func MaximumIntegrityProtectedDataRate () (value ngapType.MaximumIntegrityProtectedDataRate) {
	value.Value = ngapType.MaximumIntegrityProtectedDataRatePresentBitrate64kbs
	value.Value = ngapType.MaximumIntegrityProtectedDataRatePresentMaximumUERate
	return
}

func ULNGUUPTNLModifyItem () (value ngapType.ULNGUUPTNLModifyItem) {
	value.ULNGUUPTNLInformation = UPTransportLayerInformation()
	value.DLNGUUPTNLInformation = UPTransportLayerInformation()
	return
}