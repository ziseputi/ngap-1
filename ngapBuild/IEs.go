// Created By HaoDHH-245789 VHT2020
package ngapBuild

import "ngap/ngapType"

func AdditionalDLUPTNLInformationForHOList() (value ngapType.AdditionalDLUPTNLInformationForHOList) {
	item := AdditionalDLUPTNLInformationForHOItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item)
	return
}

func AdditionalDLUPTNLInformationForHOItem() (value ngapType.AdditionalDLUPTNLInformationForHOItem) {
	value.AdditionalDLNGUUPTNLInformation = UPTransportLayerInformation()
	value.AdditionalQosFlowSetupResponseList = QosFlowListWithDataForwarding()
	value.AdditionalDLForwardingUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.AdditionalDLForwardingUPTNLInformation = UPTransportLayerInformation()
	return
}

func AdditionalQosFlowInformation() (value ngapType.AdditionalQosFlowInformation) {
	value.Value = ngapType.AdditionalQosFlowInformationMoreLikely
	return
}

func AllocationAndRetentionPriority() (value ngapType.AllocationAndRetentionPriority) {
	value.PriorityLevelARP = PriorityLevelARP()
	value.PreEmptionCapability = PreEmptionCapability()
	value.PreEmptionVulnerability = PreEmptionVulnerability()
	return
}

func AllowedNSSAI() (value ngapType.AllowedNSSAI) {
	item := AllowedNSSAIItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item)
	return
}

func AllowedNSSAIItem() (value ngapType.AllowedNSSAIItem) {
	value.SNSSAI = SNSSAI()
	return
}

func AllowedTACs() (value ngapType.AllowedTACs) {
	item := TAC()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AMFName() (value ngapType.AMFName) {
	value.Value = "-haodhh-vht5gc-"
	value.Value = "-haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc-"
	return
}

func AMFPagingTarget() (value ngapType.AMFPagingTarget) {
	value.Present = ngapType.AMFPagingTargetPresentGlobalRANNodeID
	value.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*value.GlobalRANNodeID = GlobalRANNodeID()
	value.Present = ngapType.AMFPagingTargetPresentTAI
	value.TAI = new(ngapType.TAI)
	*value.TAI = TAI()
	return
}

func AMFPointer() (value ngapType.AMFPointer) {
	value.Value.BitLength = 6
	value.Value.Bytes = []byte{171}
	return
}

func AMFRegionID() (value ngapType.AMFRegionID) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{171}
	return
}

func AMFSetID() (value ngapType.AMFSetID) {
	value.Value.BitLength = 10
	value.Value.Bytes = []byte{171,255}
	return
}

func AMFTNLAssociationSetupList() (value ngapType.AMFTNLAssociationSetupList) {
	item := AMFTNLAssociationSetupItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AMFTNLAssociationSetupItem() (value ngapType.AMFTNLAssociationSetupItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	return
}

func AMFTNLAssociationToAddList() (value ngapType.AMFTNLAssociationToAddList) {
	item := AMFTNLAssociationToAddItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AMFTNLAssociationToAddItem() (value ngapType.AMFTNLAssociationToAddItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	value.TNLAssociationUsage = new(ngapType.TNLAssociationUsage)
	*value.TNLAssociationUsage = TNLAssociationUsage()
	value.TNLAddressWeightFactor = TNLAddressWeightFactor()
	return
}

func AMFTNLAssociationToRemoveList() (value ngapType.AMFTNLAssociationToRemoveList) {
	item := AMFTNLAssociationToRemoveItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AMFTNLAssociationToRemoveItem() (value ngapType.AMFTNLAssociationToRemoveItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	return
}

func AMFTNLAssociationToUpdateList() (value ngapType.AMFTNLAssociationToUpdateList) {
	item := AMFTNLAssociationToUpdateItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AMFTNLAssociationToUpdateItem() (value ngapType.AMFTNLAssociationToUpdateItem) {
	value.AMFTNLAssociationAddress = CPTransportLayerInformation()
	value.TNLAssociationUsage = new(ngapType.TNLAssociationUsage)
	*value.TNLAssociationUsage = TNLAssociationUsage()
	value.TNLAddressWeightFactor = new(ngapType.TNLAddressWeightFactor)
	*value.TNLAddressWeightFactor = TNLAddressWeightFactor()
	return
}

func AMFUENGAPID() (value ngapType.AMFUENGAPID) {
	value.Value = 0
	value.Value = 1099511627775
	return
}

func AreaOfInterest() (value ngapType.AreaOfInterest) {
	value.AreaOfInterestTAIList = new(ngapType.AreaOfInterestTAIList)
	*value.AreaOfInterestTAIList = AreaOfInterestTAIList()
	value.AreaOfInterestCellList = new(ngapType.AreaOfInterestCellList)
	*value.AreaOfInterestCellList = AreaOfInterestCellList()
	value.AreaOfInterestRANNodeList = new(ngapType.AreaOfInterestRANNodeList)
	*value.AreaOfInterestRANNodeList = AreaOfInterestRANNodeList()
	return
}

func AreaOfInterestCellList() (value ngapType.AreaOfInterestCellList) {
	item := AreaOfInterestCellItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AreaOfInterestCellItem() (value ngapType.AreaOfInterestCellItem) {
	value.NGRANCGI = NGRANCGI()
	return
}

func AreaOfInterestList() (value ngapType.AreaOfInterestList) {
	item := AreaOfInterestItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AreaOfInterestItem() (value ngapType.AreaOfInterestItem) {
	value.AreaOfInterest = AreaOfInterest()
	value.LocationReportingReferenceID = LocationReportingReferenceID()
	return
}

func AreaOfInterestRANNodeList() (value ngapType.AreaOfInterestRANNodeList) {
	item := AreaOfInterestRANNodeItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AreaOfInterestRANNodeItem() (value ngapType.AreaOfInterestRANNodeItem) {
	value.GlobalRANNodeID = GlobalRANNodeID()
	return
}

func AreaOfInterestTAIList() (value ngapType.AreaOfInterestTAIList) {
	item := AreaOfInterestTAIItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AreaOfInterestTAIItem() (value ngapType.AreaOfInterestTAIItem) {
	value.TAI = TAI()
	return
}

func AssistanceDataForPaging() (value ngapType.AssistanceDataForPaging) {
	value.AssistanceDataForRecommendedCells = new(ngapType.AssistanceDataForRecommendedCells)
	*value.AssistanceDataForRecommendedCells = AssistanceDataForRecommendedCells()
	value.PagingAttemptInformation = new(ngapType.PagingAttemptInformation)
	*value.PagingAttemptInformation = PagingAttemptInformation()
	return
}

func AssistanceDataForRecommendedCells() (value ngapType.AssistanceDataForRecommendedCells) {
	value.RecommendedCellsForPaging = RecommendedCellsForPaging()
	return
}

func AssociatedQosFlowList() (value ngapType.AssociatedQosFlowList) {
	item := AssociatedQosFlowItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func AssociatedQosFlowItem() (value ngapType.AssociatedQosFlowItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	return
}

func AveragingWindow() (value ngapType.AveragingWindow) {
	value.Value = 0
	value.Value = 4095
	return
}

func BitRate() (value ngapType.BitRate) {
	value.Value = 0
	value.Value = 4000000000000
	return
}

func BroadcastCancelledAreaList() (value ngapType.BroadcastCancelledAreaList) {
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

func BroadcastCompletedAreaList() (value ngapType.BroadcastCompletedAreaList) {
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

func BroadcastPLMNList() (value ngapType.BroadcastPLMNList) {
	item := BroadcastPLMNItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func BroadcastPLMNItem() (value ngapType.BroadcastPLMNItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.TAISliceSupportList = SliceSupportList()
	return
}

func CancelAllWarningMessages() (value ngapType.CancelAllWarningMessages) {
	value.Value = ngapType.CancelAllWarningMessagesTrue
	return
}

func CancelledCellsInEAIEUTRA() (value ngapType.CancelledCellsInEAIEUTRA) {
	item := CancelledCellsInEAIEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CancelledCellsInEAIEUTRAItem() (value ngapType.CancelledCellsInEAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CancelledCellsInEAINR() (value ngapType.CancelledCellsInEAINR) {
	item := CancelledCellsInEAINRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CancelledCellsInEAINRItem() (value ngapType.CancelledCellsInEAINRItem) {
	value.NRCGI = NRCGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CancelledCellsInTAIEUTRA() (value ngapType.CancelledCellsInTAIEUTRA) {
	item := CancelledCellsInTAIEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CancelledCellsInTAIEUTRAItem() (value ngapType.CancelledCellsInTAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CancelledCellsInTAINR() (value ngapType.CancelledCellsInTAINR) {
	item := CancelledCellsInTAINRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CancelledCellsInTAINRItem() (value ngapType.CancelledCellsInTAINRItem) {
	value.NRCGI = NRCGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func Cause() (value ngapType.Cause) {
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

func CauseMisc() (value ngapType.CauseMisc) {
	value.Value = ngapType.CauseMiscControlProcessingOverload
	value.Value = ngapType.CauseMiscNotEnoughUserPlaneProcessingResources
	value.Value = ngapType.CauseMiscHardwareFailure
	value.Value = ngapType.CauseMiscOmIntervention
	value.Value = ngapType.CauseMiscUnknownPLMN
	value.Value = ngapType.CauseMiscUnspecified
	return
}

func CauseNas() (value ngapType.CauseNas) {
	value.Value = ngapType.CauseNasNormalRelease
	value.Value = ngapType.CauseNasAuthenticationFailure
	value.Value = ngapType.CauseNasDeregister
	value.Value = ngapType.CauseNasUnspecified
	return
}

func CauseProtocol() (value ngapType.CauseProtocol) {
	value.Value = ngapType.CauseProtocolTransferSyntaxError
	value.Value = ngapType.CauseProtocolAbstractSyntaxErrorReject
	value.Value = ngapType.CauseProtocolAbstractSyntaxErrorIgnoreAndNotify
	value.Value = ngapType.CauseProtocolMessageNotCompatibleWithReceiverState
	value.Value = ngapType.CauseProtocolSemanticError
	value.Value = ngapType.CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage
	value.Value = ngapType.CauseProtocolUnspecified
	return
}

func CauseRadioNetwork() (value ngapType.CauseRadioNetwork) {
	value.Value = ngapType.CauseRadioNetworkUnspecified
	value.Value = ngapType.CauseRadioNetworkTxnrelocoverallExpiry
	value.Value = ngapType.CauseRadioNetworkSuccessfulHandover
	value.Value = ngapType.CauseRadioNetworkReleaseDueToNgranGeneratedReason
	value.Value = ngapType.CauseRadioNetworkReleaseDueTo5gcGeneratedReason
	value.Value = ngapType.CauseRadioNetworkHandoverCancelled
	value.Value = ngapType.CauseRadioNetworkPartialHandover
	value.Value = ngapType.CauseRadioNetworkHoFailureInTarget5GCNgranNodeOrTargetSystem
	value.Value = ngapType.CauseRadioNetworkHoTargetNotAllowed
	value.Value = ngapType.CauseRadioNetworkTngrelocoverallExpiry
	value.Value = ngapType.CauseRadioNetworkTngrelocprepExpiry
	value.Value = ngapType.CauseRadioNetworkCellNotAvailable
	value.Value = ngapType.CauseRadioNetworkUnknownTargetID
	value.Value = ngapType.CauseRadioNetworkNoRadioResourcesAvailableInTargetCell
	value.Value = ngapType.CauseRadioNetworkUnknownLocalUENGAPID
	value.Value = ngapType.CauseRadioNetworkInconsistentRemoteUENGAPID
	value.Value = ngapType.CauseRadioNetworkHandoverDesirableForRadioReason
	value.Value = ngapType.CauseRadioNetworkTimeCriticalHandover
	value.Value = ngapType.CauseRadioNetworkResourceOptimisationHandover
	value.Value = ngapType.CauseRadioNetworkReduceLoadInServingCell
	value.Value = ngapType.CauseRadioNetworkUserInactivity
	value.Value = ngapType.CauseRadioNetworkRadioConnectionWithUeLost
	value.Value = ngapType.CauseRadioNetworkRadioResourcesNotAvailable
	value.Value = ngapType.CauseRadioNetworkInvalidQosCombination
	value.Value = ngapType.CauseRadioNetworkFailureInRadioInterfaceProcedure
	value.Value = ngapType.CauseRadioNetworkInteractionWithOtherProcedure
	value.Value = ngapType.CauseRadioNetworkUnknownPDUSessionID
	value.Value = ngapType.CauseRadioNetworkUnkownQosFlowID
	value.Value = ngapType.CauseRadioNetworkMultiplePDUSessionIDInstances
	value.Value = ngapType.CauseRadioNetworkMultipleQosFlowIDInstances
	value.Value = ngapType.CauseRadioNetworkEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported
	value.Value = ngapType.CauseRadioNetworkNgIntraSystemHandoverTriggered
	value.Value = ngapType.CauseRadioNetworkNgInterSystemHandoverTriggered
	value.Value = ngapType.CauseRadioNetworkXnHandoverTriggered
	value.Value = ngapType.CauseRadioNetworkNotSupported5QIValue
	value.Value = ngapType.CauseRadioNetworkUeContextTransfer
	value.Value = ngapType.CauseRadioNetworkImsVoiceEpsFallbackOrRatFallbackTriggered
	value.Value = ngapType.CauseRadioNetworkUpIntegrityProtectionNotPossible
	value.Value = ngapType.CauseRadioNetworkUpConfidentialityProtectionNotPossible
	value.Value = ngapType.CauseRadioNetworkSliceNotSupported
	value.Value = ngapType.CauseRadioNetworkUeInRrcInactiveStateNotReachable
	value.Value = ngapType.CauseRadioNetworkRedirection
	value.Value = ngapType.CauseRadioNetworkResourcesNotAvailableForTheSlice
	value.Value = ngapType.CauseRadioNetworkUeMaxIntegrityProtectedDataRateReason
	value.Value = ngapType.CauseRadioNetworkReleaseDueToCnDetectedMobility
	value.Value = ngapType.CauseRadioNetworkN26InterfaceNotAvailable
	value.Value = ngapType.CauseRadioNetworkReleaseDueToPreEmption
	value.Value = ngapType.CauseRadioNetworkMultipleLocationReportingReferenceIDInstances
	return
}

func CauseTransport() (value ngapType.CauseTransport) {
	value.Value = ngapType.CauseTransportTransportResourceUnavailable
	value.Value = ngapType.CauseTransportUnspecified
	return
}

func CellIDBroadcastEUTRA() (value ngapType.CellIDBroadcastEUTRA) {
	item := CellIDBroadcastEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CellIDBroadcastEUTRAItem() (value ngapType.CellIDBroadcastEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	return
}

func CellIDBroadcastNR() (value ngapType.CellIDBroadcastNR) {
	item := CellIDBroadcastNRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CellIDBroadcastNRItem() (value ngapType.CellIDBroadcastNRItem) {
	value.NRCGI = NRCGI()
	return
}

func CellIDCancelledEUTRA() (value ngapType.CellIDCancelledEUTRA) {
	item := CellIDCancelledEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CellIDCancelledEUTRAItem() (value ngapType.CellIDCancelledEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CellIDCancelledNR() (value ngapType.CellIDCancelledNR) {
	item := CellIDCancelledNRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CellIDCancelledNRItem() (value ngapType.CellIDCancelledNRItem) {
	value.NRCGI = NRCGI()
	value.NumberOfBroadcasts = NumberOfBroadcasts()
	return
}

func CellIDListForRestart() (value ngapType.CellIDListForRestart) {
	value.Present = ngapType.CellIDListForRestartPresentEUTRACGIListforRestart
	value.EUTRACGIListforRestart = new(ngapType.EUTRACGIList)
	*value.EUTRACGIListforRestart = EUTRACGIList()
	value.Present = ngapType.CellIDListForRestartPresentNRCGIListforRestart
	value.NRCGIListforRestart = new(ngapType.NRCGIList)
	*value.NRCGIListforRestart = NRCGIList()
	return
}

func CellSize() (value ngapType.CellSize) {
	value.Value = ngapType.CellSizeVerysmall
	value.Value = ngapType.CellSizeSmall
	value.Value = ngapType.CellSizeMedium
	value.Value = ngapType.CellSizeLarge
	return
}

func CellType() (value ngapType.CellType) {
	value.CellSize = CellSize()
	return
}

func CNAssistedRANTuning() (value ngapType.CNAssistedRANTuning) {
	value.ExpectedUEBehaviour = new(ngapType.ExpectedUEBehaviour)
	*value.ExpectedUEBehaviour = ExpectedUEBehaviour()
	return
}

func CNTypeRestrictionsForEquivalent() (value ngapType.CNTypeRestrictionsForEquivalent) {
	item := CNTypeRestrictionsForEquivalentItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CNTypeRestrictionsForEquivalentItem() (value ngapType.CNTypeRestrictionsForEquivalentItem) {
	value.PlmnIdentity = PLMNIdentity()
	return
}

func CNTypeRestrictionsForServing() (value ngapType.CNTypeRestrictionsForServing) {
	value.Value = ngapType.CNTypeRestrictionsForServingEpcForbidden
	return
}

func CommonNetworkInstance() (value ngapType.CommonNetworkInstance) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func CompletedCellsInEAIEUTRA() (value ngapType.CompletedCellsInEAIEUTRA) {
	item := CompletedCellsInEAIEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CompletedCellsInEAIEUTRAItem() (value ngapType.CompletedCellsInEAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	return
}

func CompletedCellsInEAINR() (value ngapType.CompletedCellsInEAINR) {
	item := CompletedCellsInEAINRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CompletedCellsInEAINRItem() (value ngapType.CompletedCellsInEAINRItem) {
	value.NRCGI = NRCGI()
	return
}

func CompletedCellsInTAIEUTRA() (value ngapType.CompletedCellsInTAIEUTRA) {
	item := CompletedCellsInTAIEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CompletedCellsInTAIEUTRAItem() (value ngapType.CompletedCellsInTAIEUTRAItem) {
	value.EUTRACGI = EUTRACGI()
	return
}

func CompletedCellsInTAINR() (value ngapType.CompletedCellsInTAINR) {
	item := CompletedCellsInTAINRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CompletedCellsInTAINRItem() (value ngapType.CompletedCellsInTAINRItem) {
	value.NRCGI = NRCGI()
	return
}

func ConcurrentWarningMessageInd() (value ngapType.ConcurrentWarningMessageInd) {
	value.Value = ngapType.ConcurrentWarningMessageIndTrue
	return
}

func ConfidentialityProtectionIndication() (value ngapType.ConfidentialityProtectionIndication) {
	value.Value = ngapType.ConfidentialityProtectionIndicationRequired
	value.Value = ngapType.ConfidentialityProtectionIndicationPreferred
	value.Value = ngapType.ConfidentialityProtectionIndicationNotNeeded
	return
}

func ConfidentialityProtectionResult() (value ngapType.ConfidentialityProtectionResult) {
	value.Value = ngapType.ConfidentialityProtectionResultPerformed
	value.Value = ngapType.ConfidentialityProtectionResultNotPerformed
	return
}

func CoreNetworkAssistanceInformationForInactive() (value ngapType.CoreNetworkAssistanceInformationForInactive) {
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

func COUNTValueForPDCPSN12() (value ngapType.COUNTValueForPDCPSN12) {
	return
}

func COUNTValueForPDCPSN18() (value ngapType.COUNTValueForPDCPSN18) {
	return
}

func CPTransportLayerInformation() (value ngapType.CPTransportLayerInformation) {
	value.Present = ngapType.CPTransportLayerInformationPresentEndpointIPAddress
	value.EndpointIPAddress = new(ngapType.TransportLayerAddress)
	*value.EndpointIPAddress = TransportLayerAddress()
	return
}

func CriticalityDiagnostics() (value ngapType.CriticalityDiagnostics) {
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

func CriticalityDiagnosticsIEList() (value ngapType.CriticalityDiagnosticsIEList) {
	item := CriticalityDiagnosticsIEItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func CriticalityDiagnosticsIEItem() (value ngapType.CriticalityDiagnosticsIEItem) {
	value.IECriticality = Criticality()
	value.IEID = ProtocolIEID()
	value.TypeOfError = TypeOfError()
	return
}

func DataCodingScheme() (value ngapType.DataCodingScheme) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{171}
	return
}

func DataForwardingAccepted() (value ngapType.DataForwardingAccepted) {
	value.Value = ngapType.DataForwardingAcceptedDataForwardingAccepted
	return
}

func DataForwardingNotPossible() (value ngapType.DataForwardingNotPossible) {
	value.Value = ngapType.DataForwardingNotPossibleDataForwardingNotPossible
	return
}

func DataForwardingResponseDRBList() (value ngapType.DataForwardingResponseDRBList) {
	item := DataForwardingResponseDRBItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func DataForwardingResponseDRBItem() (value ngapType.DataForwardingResponseDRBItem) {
	value.DRBID = DRBID()
	value.DLForwardingUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.DLForwardingUPTNLInformation = UPTransportLayerInformation()
	value.ULForwardingUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.ULForwardingUPTNLInformation = UPTransportLayerInformation()
	return
}

func DataForwardingResponseERABList() (value ngapType.DataForwardingResponseERABList) {
	item := DataForwardingResponseERABListItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func DataForwardingResponseERABListItem() (value ngapType.DataForwardingResponseERABListItem) {
	value.ERABID = ERABID()
	value.DLForwarding = new(ngapType.DLForwarding)
	*value.DLForwarding = DLForwarding()
	return
}

func DelayCritical() (value ngapType.DelayCritical) {
	value.Value = ngapType.DelayCriticalDelayCritical
	value.Value = ngapType.DelayCriticalNonDelayCritical
	return
}

func DLForwarding() (value ngapType.DLForwarding) {
	value.Value = ngapType.DLForwardingDlForwardingProposed
	return
}

func DLNGUTNLInformationReused() (value ngapType.DLNGUTNLInformationReused) {
	value.Value = ngapType.DLNGUTNLInformationReusedTrue
	return
}

func DirectForwardingPathAvailability() (value ngapType.DirectForwardingPathAvailability) {
	value.Value = ngapType.DirectForwardingPathAvailabilityDirectPathAvailable
	return
}

func DRBID() (value ngapType.DRBID) {
	value.Value = 1
	value.Value = 32
	return
}

func DRBsSubjectToStatusTransferList() (value ngapType.DRBsSubjectToStatusTransferList) {
	item := DRBsSubjectToStatusTransferItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func DRBsSubjectToStatusTransferItem() (value ngapType.DRBsSubjectToStatusTransferItem) {
	value.DRBID = DRBID()
	value.DRBStatusUL = DRBStatusUL()
	value.DRBStatusDL = DRBStatusDL()
	return
}

func DRBStatusDL() (value ngapType.DRBStatusDL) {
	value.Present = ngapType.DRBStatusDLPresentDRBStatusDL12
	value.DRBStatusDL12 = new(ngapType.DRBStatusDL12)
	*value.DRBStatusDL12 = DRBStatusDL12()
	value.Present = ngapType.DRBStatusDLPresentDRBStatusDL18
	value.DRBStatusDL18 = new(ngapType.DRBStatusDL18)
	*value.DRBStatusDL18 = DRBStatusDL18()
	return
}

func DRBStatusDL12() (value ngapType.DRBStatusDL12) {
	value.DLCOUNTValue = COUNTValueForPDCPSN12()
	return
}

func DRBStatusDL18() (value ngapType.DRBStatusDL18) {
	value.DLCOUNTValue = COUNTValueForPDCPSN18()
	return
}

func DRBStatusUL() (value ngapType.DRBStatusUL) {
	value.Present = ngapType.DRBStatusULPresentDRBStatusUL12
	value.DRBStatusUL12 = new(ngapType.DRBStatusUL12)
	*value.DRBStatusUL12 = DRBStatusUL12()
	value.Present = ngapType.DRBStatusULPresentDRBStatusUL18
	value.DRBStatusUL18 = new(ngapType.DRBStatusUL18)
	*value.DRBStatusUL18 = DRBStatusUL18()
	return
}

func DRBStatusUL12() (value ngapType.DRBStatusUL12) {
	value.ULCOUNTValue = COUNTValueForPDCPSN12()
	return
}

func DRBStatusUL18() (value ngapType.DRBStatusUL18) {
	value.ULCOUNTValue = COUNTValueForPDCPSN18()
	return
}

func DRBsToQosFlowsMappingList() (value ngapType.DRBsToQosFlowsMappingList) {
	item := DRBsToQosFlowsMappingItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func DRBsToQosFlowsMappingItem() (value ngapType.DRBsToQosFlowsMappingItem) {
	value.DRBID = DRBID()
	value.AssociatedQosFlowList = AssociatedQosFlowList()
	return
}

func Dynamic5QIDescriptor() (value ngapType.Dynamic5QIDescriptor) {
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

func EmergencyAreaID() (value ngapType.EmergencyAreaID) {
	value.Value = []byte{171,1,2}
	return
}

func EmergencyAreaIDBroadcastEUTRA() (value ngapType.EmergencyAreaIDBroadcastEUTRA) {
	item := EmergencyAreaIDBroadcastEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EmergencyAreaIDBroadcastEUTRAItem() (value ngapType.EmergencyAreaIDBroadcastEUTRAItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CompletedCellsInEAIEUTRA = CompletedCellsInEAIEUTRA()
	return
}

func EmergencyAreaIDBroadcastNR() (value ngapType.EmergencyAreaIDBroadcastNR) {
	item := EmergencyAreaIDBroadcastNRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EmergencyAreaIDBroadcastNRItem() (value ngapType.EmergencyAreaIDBroadcastNRItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CompletedCellsInEAINR = CompletedCellsInEAINR()
	return
}

func EmergencyAreaIDCancelledEUTRA() (value ngapType.EmergencyAreaIDCancelledEUTRA) {
	item := EmergencyAreaIDCancelledEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EmergencyAreaIDCancelledEUTRAItem() (value ngapType.EmergencyAreaIDCancelledEUTRAItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CancelledCellsInEAIEUTRA = CancelledCellsInEAIEUTRA()
	return
}

func EmergencyAreaIDCancelledNR() (value ngapType.EmergencyAreaIDCancelledNR) {
	item := EmergencyAreaIDCancelledNRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EmergencyAreaIDCancelledNRItem() (value ngapType.EmergencyAreaIDCancelledNRItem) {
	value.EmergencyAreaID = EmergencyAreaID()
	value.CancelledCellsInEAINR = CancelledCellsInEAINR()
	return
}

func EmergencyAreaIDList() (value ngapType.EmergencyAreaIDList) {
	item := EmergencyAreaID()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EmergencyAreaIDListForRestart() (value ngapType.EmergencyAreaIDListForRestart) {
	item := EmergencyAreaID()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EmergencyFallbackIndicator() (value ngapType.EmergencyFallbackIndicator) {
	value.EmergencyFallbackRequestIndicator = EmergencyFallbackRequestIndicator()
	value.EmergencyServiceTargetCN = new(ngapType.EmergencyServiceTargetCN)
	*value.EmergencyServiceTargetCN = EmergencyServiceTargetCN()
	return
}

func EmergencyFallbackRequestIndicator() (value ngapType.EmergencyFallbackRequestIndicator) {
	value.Value = ngapType.EmergencyFallbackRequestIndicatorEmergencyFallbackRequested
	return
}

func EmergencyServiceTargetCN() (value ngapType.EmergencyServiceTargetCN) {
	value.Value = ngapType.EmergencyServiceTargetCNFiveGC
	value.Value = ngapType.EmergencyServiceTargetCNEpc
	return
}

func ENDCSONConfigurationTransfer() (value ngapType.ENDCSONConfigurationTransfer) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func EndpointIPAddressAndPort() (value ngapType.EndpointIPAddressAndPort) {
	value.EndpointIPAddress = TransportLayerAddress()
	value.PortNumber = PortNumber()
	return
}

func EquivalentPLMNs() (value ngapType.EquivalentPLMNs) {
	item := PLMNIdentity()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EPSTAC() (value ngapType.EPSTAC) {
	value.Value = []byte{171,1}
	return
}

func EPSTAI() (value ngapType.EPSTAI) {
	value.PLMNIdentity = PLMNIdentity()
	value.EPSTAC = EPSTAC()
	return
}

func ERABID() (value ngapType.ERABID) {
	value.Value = 0
	value.Value = 15
	return
}

func ERABInformationList() (value ngapType.ERABInformationList) {
	item := ERABInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func ERABInformationItem() (value ngapType.ERABInformationItem) {
	value.ERABID = ERABID()
	value.DLForwarding = new(ngapType.DLForwarding)
	*value.DLForwarding = DLForwarding()
	return
}

func EUTRACellIdentity() (value ngapType.EUTRACellIdentity) {
	value.Value.BitLength = 28
	value.Value.Bytes = []byte{171,101,102,255}
	return
}

func EUTRACGI() (value ngapType.EUTRACGI) {
	value.PLMNIdentity = PLMNIdentity()
	value.EUTRACellIdentity = EUTRACellIdentity()
	return
}

func EUTRACGIList() (value ngapType.EUTRACGIList) {
	item := EUTRACGI()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EUTRACGIListForWarning() (value ngapType.EUTRACGIListForWarning) {
	item := EUTRACGI()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func EUTRAencryptionAlgorithms() (value ngapType.EUTRAencryptionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func EUTRAintegrityProtectionAlgorithms() (value ngapType.EUTRAintegrityProtectionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func EventType() (value ngapType.EventType) {
	value.Value = ngapType.EventTypeDirect
	value.Value = ngapType.EventTypeChangeOfServeCell
	value.Value = ngapType.EventTypeUePresenceInAreaOfInterest
	value.Value = ngapType.EventTypeStopChangeOfServeCell
	value.Value = ngapType.EventTypeStopUePresenceInAreaOfInterest
	value.Value = ngapType.EventTypeCancelLocationReportingForTheUe
	return
}

func ExpectedActivityPeriod() (value ngapType.ExpectedActivityPeriod) {
	value.Value = 1
	value.Value = 30|40|50|60|80|100|120|150|180|181
	return
}

func ExpectedHOInterval() (value ngapType.ExpectedHOInterval) {
	value.Value = ngapType.ExpectedHOIntervalSec15
	value.Value = ngapType.ExpectedHOIntervalSec30
	value.Value = ngapType.ExpectedHOIntervalSec60
	value.Value = ngapType.ExpectedHOIntervalSec90
	value.Value = ngapType.ExpectedHOIntervalSec120
	value.Value = ngapType.ExpectedHOIntervalSec180
	value.Value = ngapType.ExpectedHOIntervalLongTime
	return
}

func ExpectedIdlePeriod() (value ngapType.ExpectedIdlePeriod) {
	value.Value = 1
	value.Value = 30|40|50|60|80|100|120|150|180|181
	return
}

func ExpectedUEActivityBehaviour() (value ngapType.ExpectedUEActivityBehaviour) {
	value.ExpectedActivityPeriod = new(ngapType.ExpectedActivityPeriod)
	*value.ExpectedActivityPeriod = ExpectedActivityPeriod()
	value.ExpectedIdlePeriod = new(ngapType.ExpectedIdlePeriod)
	*value.ExpectedIdlePeriod = ExpectedIdlePeriod()
	value.SourceOfUEActivityBehaviourInformation = new(ngapType.SourceOfUEActivityBehaviourInformation)
	*value.SourceOfUEActivityBehaviourInformation = SourceOfUEActivityBehaviourInformation()
	return
}

func ExpectedUEBehaviour() (value ngapType.ExpectedUEBehaviour) {
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

func ExpectedUEMobility() (value ngapType.ExpectedUEMobility) {
	value.Value = ngapType.ExpectedUEMobilityStationary
	value.Value = ngapType.ExpectedUEMobilityMobile
	return
}

func ExpectedUEMovingTrajectory() (value ngapType.ExpectedUEMovingTrajectory) {
	item := ExpectedUEMovingTrajectoryItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func ExpectedUEMovingTrajectoryItem() (value ngapType.ExpectedUEMovingTrajectoryItem) {
	value.NGRANCGI = NGRANCGI()
	return
}

func ExtendedRATRestrictionInformation() (value ngapType.ExtendedRATRestrictionInformation) {
	return
}

func ExtendedRNCID() (value ngapType.ExtendedRNCID) {
	value.Value = 4096
	value.Value = 65535
	return
}

func FiveGSTMSI() (value ngapType.FiveGSTMSI) {
	value.AMFSetID = AMFSetID()
	value.AMFPointer = AMFPointer()
	value.FiveGTMSI = FiveGTMSI()
	return
}

func FiveGTMSI() (value ngapType.FiveGTMSI) {
	value.Value = []byte{171,1,2,3}
	return
}

func FiveQI() (value ngapType.FiveQI) {
	value.Value = 0
	value.Value = 255
	return
}

func ForbiddenAreaInformation() (value ngapType.ForbiddenAreaInformation) {
	item := ForbiddenAreaInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func ForbiddenAreaInformationItem() (value ngapType.ForbiddenAreaInformationItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.ForbiddenTACs = ForbiddenTACs()
	return
}

func ForbiddenTACs() (value ngapType.ForbiddenTACs) {
	item := TAC()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func GBRQosInformation() (value ngapType.GBRQosInformation) {
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

func GlobalGNBID() (value ngapType.GlobalGNBID) {
	value.PLMNIdentity = PLMNIdentity()
	value.GNBID = GNBID()
	return
}

func GlobalN3IWFID() (value ngapType.GlobalN3IWFID) {
	value.PLMNIdentity = PLMNIdentity()
	value.N3IWFID = N3IWFID()
	return
}

func GlobalNgENBID() (value ngapType.GlobalNgENBID) {
	value.PLMNIdentity = PLMNIdentity()
	value.NgENBID = NgENBID()
	return
}

func GlobalRANNodeID() (value ngapType.GlobalRANNodeID) {
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

func GNBID() (value ngapType.GNBID) {
	value.Present = ngapType.GNBIDPresentGNBID
	value.GNBID = new(ngapType.BitString)
	value.GNBID.BitLength = 22
	value.GNBID.Bytes = []byte{171,219,255}
	value.GNBID.BitLength = 32
	value.GNBID.Bytes = []byte{171,219,217,255}
	return
}

func GTPTEID() (value ngapType.GTPTEID) {
	value.Value = []byte{171,1,2,3}
	return
}

func GTPTunnel() (value ngapType.GTPTunnel) {
	value.TransportLayerAddress = TransportLayerAddress()
	value.GTPTEID = GTPTEID()
	return
}

func GUAMI() (value ngapType.GUAMI) {
	value.PLMNIdentity = PLMNIdentity()
	value.AMFRegionID = AMFRegionID()
	value.AMFSetID = AMFSetID()
	value.AMFPointer = AMFPointer()
	return
}

func GUAMIType() (value ngapType.GUAMIType) {
	value.Value = ngapType.GUAMITypeNative
	value.Value = ngapType.GUAMITypeMapped
	return
}

func HandoverCommandTransfer() (value ngapType.HandoverCommandTransfer) {
	value.DLForwardingUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.DLForwardingUPTNLInformation = UPTransportLayerInformation()
	value.QosFlowToBeForwardedList = new(ngapType.QosFlowToBeForwardedList)
	*value.QosFlowToBeForwardedList = QosFlowToBeForwardedList()
	value.DataForwardingResponseDRBList = new(ngapType.DataForwardingResponseDRBList)
	*value.DataForwardingResponseDRBList = DataForwardingResponseDRBList()
	return
}

func HandoverFlag() (value ngapType.HandoverFlag) {
	value.Value = ngapType.HandoverFlagHandoverPreparation
	return
}

func HandoverPreparationUnsuccessfulTransfer() (value ngapType.HandoverPreparationUnsuccessfulTransfer) {
	value.Cause = Cause()
	return
}

func HandoverRequestAcknowledgeTransfer() (value ngapType.HandoverRequestAcknowledgeTransfer) {
	value.DLNGUUPTNLInformation = UPTransportLayerInformation()
	value.DLForwardingUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.DLForwardingUPTNLInformation = UPTransportLayerInformation()
	value.SecurityResult = new(ngapType.SecurityResult)
	*value.SecurityResult = SecurityResult()
	value.QosFlowSetupResponseList = QosFlowListWithDataForwarding()
	value.QosFlowFailedToSetupList = new(ngapType.QosFlowListWithCause)
	*value.QosFlowFailedToSetupList = QosFlowListWithCause()
	value.DataForwardingResponseDRBList = new(ngapType.DataForwardingResponseDRBList)
	*value.DataForwardingResponseDRBList = DataForwardingResponseDRBList()
	return
}

func HandoverRequiredTransfer() (value ngapType.HandoverRequiredTransfer) {
	value.DirectForwardingPathAvailability = new(ngapType.DirectForwardingPathAvailability)
	*value.DirectForwardingPathAvailability = DirectForwardingPathAvailability()
	return
}

func HandoverResourceAllocationUnsuccessfulTransfer() (value ngapType.HandoverResourceAllocationUnsuccessfulTransfer) {
	value.Cause = Cause()
	value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*value.CriticalityDiagnostics = CriticalityDiagnostics()
	return
}

func HandoverType() (value ngapType.HandoverType) {
	value.Value = ngapType.HandoverTypeIntra5gs
	value.Value = ngapType.HandoverTypeFivegsToEps
	value.Value = ngapType.HandoverTypeEpsTo5gs
	value.Value = ngapType.HandoverTypeFivegsToUtran
	return
}

func IMSVoiceSupportIndicator() (value ngapType.IMSVoiceSupportIndicator) {
	value.Value = ngapType.IMSVoiceSupportIndicatorSupported
	value.Value = ngapType.IMSVoiceSupportIndicatorNotSupported
	return
}

func IndexToRFSP() (value ngapType.IndexToRFSP) {
	value.Value = 1
	value.Value = 256
	return
}

func InfoOnRecommendedCellsAndRANNodesForPaging() (value ngapType.InfoOnRecommendedCellsAndRANNodesForPaging) {
	value.RecommendedCellsForPaging = RecommendedCellsForPaging()
	value.RecommendRANNodesForPaging = RecommendedRANNodesForPaging()
	return
}

func IntegrityProtectionIndication() (value ngapType.IntegrityProtectionIndication) {
	value.Value = ngapType.IntegrityProtectionIndicationRequired
	value.Value = ngapType.IntegrityProtectionIndicationPreferred
	value.Value = ngapType.IntegrityProtectionIndicationNotNeeded
	return
}

func IntegrityProtectionResult() (value ngapType.IntegrityProtectionResult) {
	value.Value = ngapType.IntegrityProtectionResultPerformed
	value.Value = ngapType.IntegrityProtectionResultNotPerformed
	return
}

func IntendedNumberOfPagingAttempts() (value ngapType.IntendedNumberOfPagingAttempts) {
	value.Value = 1
	value.Value = 16
	return
}

func InterfacesToTrace() (value ngapType.InterfacesToTrace) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{171}
	return
}

func LAC() (value ngapType.LAC) {
	value.Value = []byte{171,1}
	return
}

func LAI() (value ngapType.LAI) {
	value.PLMNidentity = PLMNIdentity()
	value.LAC = LAC()
	return
}

func LastVisitedCellInformation() (value ngapType.LastVisitedCellInformation) {
	value.Present = ngapType.LastVisitedCellInformationPresentNGRANCell
	value.NGRANCell = new(ngapType.LastVisitedNGRANCellInformation)
	*value.NGRANCell = LastVisitedNGRANCellInformation()
	value.Present = ngapType.LastVisitedCellInformationPresentEUTRANCell
	value.EUTRANCell = new(ngapType.LastVisitedEUTRANCellInformation)
	*value.EUTRANCell = LastVisitedEUTRANCellInformation()
	value.Present = ngapType.LastVisitedCellInformationPresentUTRANCell
	value.UTRANCell = new(ngapType.LastVisitedUTRANCellInformation)
	*value.UTRANCell = LastVisitedUTRANCellInformation()
	value.Present = ngapType.LastVisitedCellInformationPresentGERANCell
	value.GERANCell = new(ngapType.LastVisitedGERANCellInformation)
	*value.GERANCell = LastVisitedGERANCellInformation()
	return
}

func LastVisitedCellItem() (value ngapType.LastVisitedCellItem) {
	value.LastVisitedCellInformation = LastVisitedCellInformation()
	return
}

func LastVisitedEUTRANCellInformation() (value ngapType.LastVisitedEUTRANCellInformation) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func LastVisitedGERANCellInformation() (value ngapType.LastVisitedGERANCellInformation) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func LastVisitedNGRANCellInformation() (value ngapType.LastVisitedNGRANCellInformation) {
	value.GlobalCellID = NGRANCGI()
	value.CellType = CellType()
	value.TimeUEStayedInCell = TimeUEStayedInCell()
	value.TimeUEStayedInCellEnhancedGranularity = new(ngapType.TimeUEStayedInCellEnhancedGranularity)
	*value.TimeUEStayedInCellEnhancedGranularity = TimeUEStayedInCellEnhancedGranularity()
	value.HOCauseValue = new(ngapType.Cause)
	*value.HOCauseValue = Cause()
	return
}

func LastVisitedUTRANCellInformation() (value ngapType.LastVisitedUTRANCellInformation) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func LocationReportingAdditionalInfo() (value ngapType.LocationReportingAdditionalInfo) {
	value.Value = ngapType.LocationReportingAdditionalInfoIncludePSCell
	return
}

func LocationReportingReferenceID() (value ngapType.LocationReportingReferenceID) {
	value.Value = 1
	value.Value = 64
	return
}

func LocationReportingRequestType() (value ngapType.LocationReportingRequestType) {
	value.EventType = EventType()
	value.ReportArea = ReportArea()
	value.AreaOfInterestList = new(ngapType.AreaOfInterestList)
	*value.AreaOfInterestList = AreaOfInterestList()
	value.LocationReportingReferenceIDToBeCancelled = new(ngapType.LocationReportingReferenceID)
	*value.LocationReportingReferenceIDToBeCancelled = LocationReportingReferenceID()
	return
}

func MaskedIMEISV() (value ngapType.MaskedIMEISV) {
	value.Value.BitLength = 64
	value.Value.Bytes = []byte{171,101,102,103,104,105,106,255}
	return
}

func MaximumDataBurstVolume() (value ngapType.MaximumDataBurstVolume) {
	value.Value = 0
	value.Value = 2000000
	return
}

func MessageIdentifier() (value ngapType.MessageIdentifier) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func MaximumIntegrityProtectedDataRate() (value ngapType.MaximumIntegrityProtectedDataRate) {
	value.Value = ngapType.MaximumIntegrityProtectedDataRateBitrate64kbs
	value.Value = ngapType.MaximumIntegrityProtectedDataRateMaximumUERate
	return
}

func MICOModeIndication() (value ngapType.MICOModeIndication) {
	value.Value = ngapType.MICOModeIndicationTrue
	return
}

func MobilityRestrictionList() (value ngapType.MobilityRestrictionList) {
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

func N3IWFID() (value ngapType.N3IWFID) {
	value.Present = ngapType.N3IWFIDPresentN3IWFID
	value.N3IWFID = new(ngapType.BitString)
	value.N3IWFID.BitLength = 16
	value.N3IWFID.Bytes = []byte{171,219}
	return
}

func NASPDU() (value ngapType.NASPDU) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func NASSecurityParametersFromNGRAN() (value ngapType.NASSecurityParametersFromNGRAN) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func NetworkInstance() (value ngapType.NetworkInstance) {
	value.Value = 1
	value.Value = 256
	return
}

func NewSecurityContextInd() (value ngapType.NewSecurityContextInd) {
	value.Value = ngapType.NewSecurityContextIndTrue
	return
}

func NextHopChainingCount() (value ngapType.NextHopChainingCount) {
	value.Value = 0
	value.Value = 7
	return
}

func NextPagingAreaScope() (value ngapType.NextPagingAreaScope) {
	value.Value = ngapType.NextPagingAreaScopeSame
	value.Value = ngapType.NextPagingAreaScopeChanged
	return
}

func NgENBID() (value ngapType.NgENBID) {
	value.Present = ngapType.NgENBIDPresentMacroNgENBID
	value.MacroNgENBID = new(ngapType.BitString)
	value.MacroNgENBID.BitLength = 20
	value.MacroNgENBID.Bytes = []byte{171,219,255}
	value.Present = ngapType.NgENBIDPresentShortMacroNgENBID
	value.ShortMacroNgENBID = new(ngapType.BitString)
	value.ShortMacroNgENBID.BitLength = 18
	value.ShortMacroNgENBID.Bytes = []byte{171,219,255}
	value.Present = ngapType.NgENBIDPresentLongMacroNgENBID
	value.LongMacroNgENBID = new(ngapType.BitString)
	value.LongMacroNgENBID.BitLength = 21
	value.LongMacroNgENBID.Bytes = []byte{171,219,255}
	return
}

func NGRANCGI() (value ngapType.NGRANCGI) {
	value.Present = ngapType.NGRANCGIPresentNRCGI
	value.NRCGI = new(ngapType.NRCGI)
	*value.NRCGI = NRCGI()
	value.Present = ngapType.NGRANCGIPresentEUTRACGI
	value.EUTRACGI = new(ngapType.EUTRACGI)
	*value.EUTRACGI = EUTRACGI()
	return
}

func NGRANTNLAssociationToRemoveList() (value ngapType.NGRANTNLAssociationToRemoveList) {
	item := NGRANTNLAssociationToRemoveItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func NGRANTNLAssociationToRemoveItem() (value ngapType.NGRANTNLAssociationToRemoveItem) {
	value.TNLAssociationTransportLayerAddress = CPTransportLayerInformation()
	value.TNLAssociationTransportLayerAddressAMF = new(ngapType.CPTransportLayerInformation)
	*value.TNLAssociationTransportLayerAddressAMF = CPTransportLayerInformation()
	return
}

func NGRANTraceID() (value ngapType.NGRANTraceID) {
	value.Value = []byte{171,1,2,3,4,5,6,7}
	return
}

func NonDynamic5QIDescriptor() (value ngapType.NonDynamic5QIDescriptor) {
	value.FiveQI = FiveQI()
	value.PriorityLevelQos = new(ngapType.PriorityLevelQos)
	*value.PriorityLevelQos = PriorityLevelQos()
	value.AveragingWindow = new(ngapType.AveragingWindow)
	*value.AveragingWindow = AveragingWindow()
	value.MaximumDataBurstVolume = new(ngapType.MaximumDataBurstVolume)
	*value.MaximumDataBurstVolume = MaximumDataBurstVolume()
	return
}

func NotAllowedTACs() (value ngapType.NotAllowedTACs) {
	item := TAC()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func NotificationCause() (value ngapType.NotificationCause) {
	value.Value = ngapType.NotificationCauseFulfilled
	value.Value = ngapType.NotificationCauseNotFulfilled
	return
}

func NotificationControl() (value ngapType.NotificationControl) {
	value.Value = ngapType.NotificationControlNotificationRequested
	return
}

func NRCellIdentity() (value ngapType.NRCellIdentity) {
	value.Value.BitLength = 36
	value.Value.Bytes = []byte{171,101,102,103,255}
	return
}

func NRCGI() (value ngapType.NRCGI) {
	value.PLMNIdentity = PLMNIdentity()
	value.NRCellIdentity = NRCellIdentity()
	return
}

func NRCGIList() (value ngapType.NRCGIList) {
	item := NRCGI()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func NRCGIListForWarning() (value ngapType.NRCGIListForWarning) {
	item := NRCGI()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func NRencryptionAlgorithms() (value ngapType.NRencryptionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func NRintegrityProtectionAlgorithms() (value ngapType.NRintegrityProtectionAlgorithms) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func NRPPaPDU() (value ngapType.NRPPaPDU) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func NumberOfBroadcasts() (value ngapType.NumberOfBroadcasts) {
	value.Value = 0
	value.Value = 65535
	return
}

func NumberOfBroadcastsRequested() (value ngapType.NumberOfBroadcastsRequested) {
	value.Value = 0
	value.Value = 65535
	return
}

func OverloadAction() (value ngapType.OverloadAction) {
	value.Value = ngapType.OverloadActionRejectNonEmergencyMoDt
	value.Value = ngapType.OverloadActionRejectRrcCrSignalling
	value.Value = ngapType.OverloadActionPermitEmergencySessionsAndMobileTerminatedServicesOnly
	value.Value = ngapType.OverloadActionPermitHighPrioritySessionsAndMobileTerminatedServicesOnly
	return
}

func OverloadResponse() (value ngapType.OverloadResponse) {
	value.Present = ngapType.OverloadResponsePresentOverloadAction
	value.OverloadAction = new(ngapType.OverloadAction)
	*value.OverloadAction = OverloadAction()
	return
}

func OverloadStartNSSAIList() (value ngapType.OverloadStartNSSAIList) {
	item := OverloadStartNSSAIItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func OverloadStartNSSAIItem() (value ngapType.OverloadStartNSSAIItem) {
	value.SliceOverloadList = SliceOverloadList()
	value.SliceOverloadResponse = new(ngapType.OverloadResponse)
	*value.SliceOverloadResponse = OverloadResponse()
	value.SliceTrafficLoadReductionIndication = new(ngapType.TrafficLoadReductionIndication)
	*value.SliceTrafficLoadReductionIndication = TrafficLoadReductionIndication()
	return
}

func PacketDelayBudget() (value ngapType.PacketDelayBudget) {
	value.Value = 0
	value.Value = 1023
	return
}

func PacketErrorRate() (value ngapType.PacketErrorRate) {
	return
}

func PacketLossRate() (value ngapType.PacketLossRate) {
	value.Value = 0
	value.Value = 1000
	return
}

func PagingAttemptInformation() (value ngapType.PagingAttemptInformation) {
	value.PagingAttemptCount = PagingAttemptCount()
	value.IntendedNumberOfPagingAttempts = IntendedNumberOfPagingAttempts()
	value.NextPagingAreaScope = new(ngapType.NextPagingAreaScope)
	*value.NextPagingAreaScope = NextPagingAreaScope()
	return
}

func PagingAttemptCount() (value ngapType.PagingAttemptCount) {
	value.Value = 1
	value.Value = 16
	return
}

func PagingDRX() (value ngapType.PagingDRX) {
	value.Value = ngapType.PagingDRXV32
	value.Value = ngapType.PagingDRXV64
	value.Value = ngapType.PagingDRXV128
	value.Value = ngapType.PagingDRXV256
	return
}

func PagingOrigin() (value ngapType.PagingOrigin) {
	value.Value = ngapType.PagingOriginNon3gpp
	return
}

func PagingPriority() (value ngapType.PagingPriority) {
	value.Value = ngapType.PagingPriorityPriolevel1
	value.Value = ngapType.PagingPriorityPriolevel2
	value.Value = ngapType.PagingPriorityPriolevel3
	value.Value = ngapType.PagingPriorityPriolevel4
	value.Value = ngapType.PagingPriorityPriolevel5
	value.Value = ngapType.PagingPriorityPriolevel6
	value.Value = ngapType.PagingPriorityPriolevel7
	value.Value = ngapType.PagingPriorityPriolevel8
	return
}

func PathSwitchRequestAcknowledgeTransfer() (value ngapType.PathSwitchRequestAcknowledgeTransfer) {
	value.ULNGUUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.ULNGUUPTNLInformation = UPTransportLayerInformation()
	value.SecurityIndication = new(ngapType.SecurityIndication)
	*value.SecurityIndication = SecurityIndication()
	return
}

func PathSwitchRequestSetupFailedTransfer() (value ngapType.PathSwitchRequestSetupFailedTransfer) {
	value.Cause = Cause()
	return
}

func PathSwitchRequestTransfer() (value ngapType.PathSwitchRequestTransfer) {
	value.DLNGUUPTNLInformation = UPTransportLayerInformation()
	value.DLNGUTNLInformationReused = new(ngapType.DLNGUTNLInformationReused)
	*value.DLNGUTNLInformationReused = DLNGUTNLInformationReused()
	value.UserPlaneSecurityInformation = new(ngapType.UserPlaneSecurityInformation)
	*value.UserPlaneSecurityInformation = UserPlaneSecurityInformation()
	value.QosFlowAcceptedList = QosFlowAcceptedList()
	return
}

func PathSwitchRequestUnsuccessfulTransfer() (value ngapType.PathSwitchRequestUnsuccessfulTransfer) {
	value.Cause = Cause()
	return
}

func PDUSessionAggregateMaximumBitRate() (value ngapType.PDUSessionAggregateMaximumBitRate) {
	value.PDUSessionAggregateMaximumBitRateDL = BitRate()
	value.PDUSessionAggregateMaximumBitRateUL = BitRate()
	return
}

func PDUSessionID() (value ngapType.PDUSessionID) {
	value.Value = 0
	value.Value = 255
	return
}

func PDUSessionResourceAdmittedList() (value ngapType.PDUSessionResourceAdmittedList) {
	item := PDUSessionResourceAdmittedItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceAdmittedItem() (value ngapType.PDUSessionResourceAdmittedItem) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToModifyListModCfm() (value ngapType.PDUSessionResourceFailedToModifyListModCfm) {
	item := PDUSessionResourceFailedToModifyItemModCfm()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToModifyItemModCfm() (value ngapType.PDUSessionResourceFailedToModifyItemModCfm) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToModifyListModRes() (value ngapType.PDUSessionResourceFailedToModifyListModRes) {
	item := PDUSessionResourceFailedToModifyItemModRes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToModifyItemModRes() (value ngapType.PDUSessionResourceFailedToModifyItemModRes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToSetupListCxtFail() (value ngapType.PDUSessionResourceFailedToSetupListCxtFail) {
	item := PDUSessionResourceFailedToSetupItemCxtFail()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToSetupItemCxtFail() (value ngapType.PDUSessionResourceFailedToSetupItemCxtFail) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToSetupListCxtRes() (value ngapType.PDUSessionResourceFailedToSetupListCxtRes) {
	item := PDUSessionResourceFailedToSetupItemCxtRes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToSetupItemCxtRes() (value ngapType.PDUSessionResourceFailedToSetupItemCxtRes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToSetupListHOAck() (value ngapType.PDUSessionResourceFailedToSetupListHOAck) {
	item := PDUSessionResourceFailedToSetupItemHOAck()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToSetupItemHOAck() (value ngapType.PDUSessionResourceFailedToSetupItemHOAck) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToSetupListPSReq() (value ngapType.PDUSessionResourceFailedToSetupListPSReq) {
	item := PDUSessionResourceFailedToSetupItemPSReq()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToSetupItemPSReq() (value ngapType.PDUSessionResourceFailedToSetupItemPSReq) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceFailedToSetupListSURes() (value ngapType.PDUSessionResourceFailedToSetupListSURes) {
	item := PDUSessionResourceFailedToSetupItemSURes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceFailedToSetupItemSURes() (value ngapType.PDUSessionResourceFailedToSetupItemSURes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceHandoverList() (value ngapType.PDUSessionResourceHandoverList) {
	item := PDUSessionResourceHandoverItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceHandoverItem() (value ngapType.PDUSessionResourceHandoverItem) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceInformationList() (value ngapType.PDUSessionResourceInformationList) {
	item := PDUSessionResourceInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceInformationItem() (value ngapType.PDUSessionResourceInformationItem) {
	value.PDUSessionID = PDUSessionID()
	value.QosFlowInformationList = QosFlowInformationList()
	value.DRBsToQosFlowsMappingList = new(ngapType.DRBsToQosFlowsMappingList)
	*value.DRBsToQosFlowsMappingList = DRBsToQosFlowsMappingList()
	return
}

func PDUSessionResourceListCxtRelCpl() (value ngapType.PDUSessionResourceListCxtRelCpl) {
	item := PDUSessionResourceItemCxtRelCpl()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceItemCxtRelCpl() (value ngapType.PDUSessionResourceItemCxtRelCpl) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceListCxtRelReq() (value ngapType.PDUSessionResourceListCxtRelReq) {
	item := PDUSessionResourceItemCxtRelReq()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceItemCxtRelReq() (value ngapType.PDUSessionResourceItemCxtRelReq) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceListHORqd() (value ngapType.PDUSessionResourceListHORqd) {
	item := PDUSessionResourceItemHORqd()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceItemHORqd() (value ngapType.PDUSessionResourceItemHORqd) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceModifyConfirmTransfer() (value ngapType.PDUSessionResourceModifyConfirmTransfer) {
	value.QosFlowModifyConfirmList = QosFlowModifyConfirmList()
	value.ULNGUUPTNLInformation = UPTransportLayerInformation()
	value.AdditionalNGUUPTNLInformation = new(ngapType.UPTransportLayerInformationPairList)
	*value.AdditionalNGUUPTNLInformation = UPTransportLayerInformationPairList()
	value.QosFlowFailedToModifyList = new(ngapType.QosFlowListWithCause)
	*value.QosFlowFailedToModifyList = QosFlowListWithCause()
	return
}

func PDUSessionResourceModifyIndicationUnsuccessfulTransfer() (value ngapType.PDUSessionResourceModifyIndicationUnsuccessfulTransfer) {
	value.Cause = Cause()
	return
}

func PDUSessionResourceModifyRequestTransfer() (value ngapType.PDUSessionResourceModifyRequestTransfer) {
	return
}

func PDUSessionResourceModifyResponseTransfer() (value ngapType.PDUSessionResourceModifyResponseTransfer) {
	value.DLNGUUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.DLNGUUPTNLInformation = UPTransportLayerInformation()
	value.ULNGUUPTNLInformation = new(ngapType.UPTransportLayerInformation)
	*value.ULNGUUPTNLInformation = UPTransportLayerInformation()
	value.QosFlowAddOrModifyResponseList = new(ngapType.QosFlowAddOrModifyResponseList)
	*value.QosFlowAddOrModifyResponseList = QosFlowAddOrModifyResponseList()
	value.AdditionalDLQosFlowPerTNLInformation = new(ngapType.QosFlowPerTNLInformationList)
	*value.AdditionalDLQosFlowPerTNLInformation = QosFlowPerTNLInformationList()
	value.QosFlowFailedToAddOrModifyList = new(ngapType.QosFlowListWithCause)
	*value.QosFlowFailedToAddOrModifyList = QosFlowListWithCause()
	return
}

func PDUSessionResourceModifyIndicationTransfer() (value ngapType.PDUSessionResourceModifyIndicationTransfer) {
	value.DLQosFlowPerTNLInformation = QosFlowPerTNLInformation()
	value.AdditionalDLQosFlowPerTNLInformation = new(ngapType.QosFlowPerTNLInformationList)
	*value.AdditionalDLQosFlowPerTNLInformation = QosFlowPerTNLInformationList()
	return
}

func PDUSessionResourceModifyListModCfm() (value ngapType.PDUSessionResourceModifyListModCfm) {
	item := PDUSessionResourceModifyItemModCfm()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceModifyItemModCfm() (value ngapType.PDUSessionResourceModifyItemModCfm) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceModifyListModInd() (value ngapType.PDUSessionResourceModifyListModInd) {
	item := PDUSessionResourceModifyItemModInd()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceModifyItemModInd() (value ngapType.PDUSessionResourceModifyItemModInd) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceModifyListModReq() (value ngapType.PDUSessionResourceModifyListModReq) {
	item := PDUSessionResourceModifyItemModReq()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceModifyItemModReq() (value ngapType.PDUSessionResourceModifyItemModReq) {
	value.PDUSessionID = PDUSessionID()
	value.NASPDU = new(ngapType.NASPDU)
	*value.NASPDU = NASPDU()
	return
}

func PDUSessionResourceModifyListModRes() (value ngapType.PDUSessionResourceModifyListModRes) {
	item := PDUSessionResourceModifyItemModRes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceModifyItemModRes() (value ngapType.PDUSessionResourceModifyItemModRes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceModifyUnsuccessfulTransfer() (value ngapType.PDUSessionResourceModifyUnsuccessfulTransfer) {
	value.Cause = Cause()
	value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*value.CriticalityDiagnostics = CriticalityDiagnostics()
	return
}

func PDUSessionResourceNotifyList() (value ngapType.PDUSessionResourceNotifyList) {
	item := PDUSessionResourceNotifyItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceNotifyItem() (value ngapType.PDUSessionResourceNotifyItem) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceNotifyReleasedTransfer() (value ngapType.PDUSessionResourceNotifyReleasedTransfer) {
	value.Cause = Cause()
	return
}

func PDUSessionResourceNotifyTransfer() (value ngapType.PDUSessionResourceNotifyTransfer) {
	value.QosFlowNotifyList = new(ngapType.QosFlowNotifyList)
	*value.QosFlowNotifyList = QosFlowNotifyList()
	value.QosFlowReleasedList = new(ngapType.QosFlowListWithCause)
	*value.QosFlowReleasedList = QosFlowListWithCause()
	return
}

func PDUSessionResourceReleaseCommandTransfer() (value ngapType.PDUSessionResourceReleaseCommandTransfer) {
	value.Cause = Cause()
	return
}

func PDUSessionResourceReleasedListNot() (value ngapType.PDUSessionResourceReleasedListNot) {
	item := PDUSessionResourceReleasedItemNot()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceReleasedItemNot() (value ngapType.PDUSessionResourceReleasedItemNot) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceReleasedListPSAck() (value ngapType.PDUSessionResourceReleasedListPSAck) {
	item := PDUSessionResourceReleasedItemPSAck()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceReleasedItemPSAck() (value ngapType.PDUSessionResourceReleasedItemPSAck) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceReleasedListPSFail() (value ngapType.PDUSessionResourceReleasedListPSFail) {
	item := PDUSessionResourceReleasedItemPSFail()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceReleasedItemPSFail() (value ngapType.PDUSessionResourceReleasedItemPSFail) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceReleasedListRelRes() (value ngapType.PDUSessionResourceReleasedListRelRes) {
	item := PDUSessionResourceReleasedItemRelRes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceReleasedItemRelRes() (value ngapType.PDUSessionResourceReleasedItemRelRes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceReleaseResponseTransfer() (value ngapType.PDUSessionResourceReleaseResponseTransfer) {
	return
}

func PDUSessionResourceSecondaryRATUsageList() (value ngapType.PDUSessionResourceSecondaryRATUsageList) {
	item := PDUSessionResourceSecondaryRATUsageItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSecondaryRATUsageItem() (value ngapType.PDUSessionResourceSecondaryRATUsageItem) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceSetupListCxtReq() (value ngapType.PDUSessionResourceSetupListCxtReq) {
	item := PDUSessionResourceSetupItemCxtReq()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSetupItemCxtReq() (value ngapType.PDUSessionResourceSetupItemCxtReq) {
	value.PDUSessionID = PDUSessionID()
	value.NASPDU = new(ngapType.NASPDU)
	*value.NASPDU = NASPDU()
	value.SNSSAI = SNSSAI()
	return
}

func PDUSessionResourceSetupListCxtRes() (value ngapType.PDUSessionResourceSetupListCxtRes) {
	item := PDUSessionResourceSetupItemCxtRes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSetupItemCxtRes() (value ngapType.PDUSessionResourceSetupItemCxtRes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceSetupListHOReq() (value ngapType.PDUSessionResourceSetupListHOReq) {
	item := PDUSessionResourceSetupItemHOReq()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSetupItemHOReq() (value ngapType.PDUSessionResourceSetupItemHOReq) {
	value.PDUSessionID = PDUSessionID()
	value.SNSSAI = SNSSAI()
	return
}

func PDUSessionResourceSetupListSUReq() (value ngapType.PDUSessionResourceSetupListSUReq) {
	item := PDUSessionResourceSetupItemSUReq()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSetupItemSUReq() (value ngapType.PDUSessionResourceSetupItemSUReq) {
	value.PDUSessionID = PDUSessionID()
	value.PDUSessionNASPDU = new(ngapType.NASPDU)
	*value.PDUSessionNASPDU = NASPDU()
	value.SNSSAI = SNSSAI()
	return
}

func PDUSessionResourceSetupListSURes() (value ngapType.PDUSessionResourceSetupListSURes) {
	item := PDUSessionResourceSetupItemSURes()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSetupItemSURes() (value ngapType.PDUSessionResourceSetupItemSURes) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceSetupRequestTransfer() (value ngapType.PDUSessionResourceSetupRequestTransfer) {
	return
}

func PDUSessionResourceSetupResponseTransfer() (value ngapType.PDUSessionResourceSetupResponseTransfer) {
	value.DLQosFlowPerTNLInformation = QosFlowPerTNLInformation()
	value.AdditionalDLQosFlowPerTNLInformation = new(ngapType.QosFlowPerTNLInformationList)
	*value.AdditionalDLQosFlowPerTNLInformation = QosFlowPerTNLInformationList()
	value.SecurityResult = new(ngapType.SecurityResult)
	*value.SecurityResult = SecurityResult()
	value.QosFlowFailedToSetupList = new(ngapType.QosFlowListWithCause)
	*value.QosFlowFailedToSetupList = QosFlowListWithCause()
	return
}

func PDUSessionResourceSetupUnsuccessfulTransfer() (value ngapType.PDUSessionResourceSetupUnsuccessfulTransfer) {
	value.Cause = Cause()
	value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*value.CriticalityDiagnostics = CriticalityDiagnostics()
	return
}

func PDUSessionResourceSwitchedList() (value ngapType.PDUSessionResourceSwitchedList) {
	item := PDUSessionResourceSwitchedItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceSwitchedItem() (value ngapType.PDUSessionResourceSwitchedItem) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceToBeSwitchedDLList() (value ngapType.PDUSessionResourceToBeSwitchedDLList) {
	item := PDUSessionResourceToBeSwitchedDLItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceToBeSwitchedDLItem() (value ngapType.PDUSessionResourceToBeSwitchedDLItem) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceToReleaseListHOCmd() (value ngapType.PDUSessionResourceToReleaseListHOCmd) {
	item := PDUSessionResourceToReleaseItemHOCmd()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceToReleaseItemHOCmd() (value ngapType.PDUSessionResourceToReleaseItemHOCmd) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionResourceToReleaseListRelCmd() (value ngapType.PDUSessionResourceToReleaseListRelCmd) {
	item := PDUSessionResourceToReleaseItemRelCmd()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PDUSessionResourceToReleaseItemRelCmd() (value ngapType.PDUSessionResourceToReleaseItemRelCmd) {
	value.PDUSessionID = PDUSessionID()
	return
}

func PDUSessionType() (value ngapType.PDUSessionType) {
	value.Value = ngapType.PDUSessionTypeIpv4
	value.Value = ngapType.PDUSessionTypeIpv6
	value.Value = ngapType.PDUSessionTypeIpv4v6
	value.Value = ngapType.PDUSessionTypeEthernet
	value.Value = ngapType.PDUSessionTypeUnstructured
	return
}

func PDUSessionUsageReport() (value ngapType.PDUSessionUsageReport) {
	value.PDUSessionTimedReportList = VolumeTimedReportList()
	return
}

func PeriodicRegistrationUpdateTimer() (value ngapType.PeriodicRegistrationUpdateTimer) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{171}
	return
}

func PLMNIdentity() (value ngapType.PLMNIdentity) {
	value.Value = []byte{171,1,2}
	return
}

func PLMNSupportList() (value ngapType.PLMNSupportList) {
	item := PLMNSupportItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func PLMNSupportItem() (value ngapType.PLMNSupportItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.SliceSupportList = SliceSupportList()
	return
}

func PortNumber() (value ngapType.PortNumber) {
	value.Value = []byte{171,1}
	return
}

func PreEmptionCapability() (value ngapType.PreEmptionCapability) {
	value.Value = ngapType.PreEmptionCapabilityShallNotTriggerPreEmption
	value.Value = ngapType.PreEmptionCapabilityMayTriggerPreEmption
	return
}

func PreEmptionVulnerability() (value ngapType.PreEmptionVulnerability) {
	value.Value = ngapType.PreEmptionVulnerabilityNotPreEmptable
	value.Value = ngapType.PreEmptionVulnerabilityPreEmptable
	return
}

func PriorityLevelARP() (value ngapType.PriorityLevelARP) {
	value.Value = 1
	value.Value = 15
	return
}

func PriorityLevelQos() (value ngapType.PriorityLevelQos) {
	value.Value = 1
	value.Value = 127
	return
}

func PWSFailedCellIDList() (value ngapType.PWSFailedCellIDList) {
	value.Present = ngapType.PWSFailedCellIDListPresentEUTRACGIPWSFailedList
	value.EUTRACGIPWSFailedList = new(ngapType.EUTRACGIList)
	*value.EUTRACGIPWSFailedList = EUTRACGIList()
	value.Present = ngapType.PWSFailedCellIDListPresentNRCGIPWSFailedList
	value.NRCGIPWSFailedList = new(ngapType.NRCGIList)
	*value.NRCGIPWSFailedList = NRCGIList()
	return
}

func QosCharacteristics() (value ngapType.QosCharacteristics) {
	value.Present = ngapType.QosCharacteristicsPresentNonDynamic5QI
	value.NonDynamic5QI = new(ngapType.NonDynamic5QIDescriptor)
	*value.NonDynamic5QI = NonDynamic5QIDescriptor()
	value.Present = ngapType.QosCharacteristicsPresentDynamic5QI
	value.Dynamic5QI = new(ngapType.Dynamic5QIDescriptor)
	*value.Dynamic5QI = Dynamic5QIDescriptor()
	return
}

func QosFlowAcceptedList() (value ngapType.QosFlowAcceptedList) {
	item := QosFlowAcceptedItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowAcceptedItem() (value ngapType.QosFlowAcceptedItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	return
}

func QosFlowAddOrModifyRequestList() (value ngapType.QosFlowAddOrModifyRequestList) {
	item := QosFlowAddOrModifyRequestItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowAddOrModifyRequestItem() (value ngapType.QosFlowAddOrModifyRequestItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.QosFlowLevelQosParameters = new(ngapType.QosFlowLevelQosParameters)
	*value.QosFlowLevelQosParameters = QosFlowLevelQosParameters()
	value.ERABID = new(ngapType.ERABID)
	*value.ERABID = ERABID()
	return
}

func QosFlowAddOrModifyResponseList() (value ngapType.QosFlowAddOrModifyResponseList) {
	item := QosFlowAddOrModifyResponseItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowAddOrModifyResponseItem() (value ngapType.QosFlowAddOrModifyResponseItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	return
}

func QosFlowIdentifier() (value ngapType.QosFlowIdentifier) {
	value.Value = 0
	value.Value = 63
	return
}

func QosFlowInformationList() (value ngapType.QosFlowInformationList) {
	item := QosFlowInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowInformationItem() (value ngapType.QosFlowInformationItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.DLForwarding = new(ngapType.DLForwarding)
	*value.DLForwarding = DLForwarding()
	return
}

func QosFlowLevelQosParameters() (value ngapType.QosFlowLevelQosParameters) {
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

func QosMonitoringRequest() (value ngapType.QosMonitoringRequest) {
	value.Value = ngapType.QosMonitoringRequestUl
	value.Value = ngapType.QosMonitoringRequestDl
	value.Value = ngapType.QosMonitoringRequestBoth
	return
}

func QosFlowListWithCause() (value ngapType.QosFlowListWithCause) {
	item := QosFlowWithCauseItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowWithCauseItem() (value ngapType.QosFlowWithCauseItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.Cause = Cause()
	return
}

func QosFlowModifyConfirmList() (value ngapType.QosFlowModifyConfirmList) {
	item := QosFlowModifyConfirmItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowModifyConfirmItem() (value ngapType.QosFlowModifyConfirmItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	return
}

func QosFlowNotifyList() (value ngapType.QosFlowNotifyList) {
	item := QosFlowNotifyItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowNotifyItem() (value ngapType.QosFlowNotifyItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.NotificationCause = NotificationCause()
	return
}

func QosFlowPerTNLInformation() (value ngapType.QosFlowPerTNLInformation) {
	value.UPTransportLayerInformation = UPTransportLayerInformation()
	value.AssociatedQosFlowList = AssociatedQosFlowList()
	return
}

func QosFlowPerTNLInformationList() (value ngapType.QosFlowPerTNLInformationList) {
	item := QosFlowPerTNLInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item)
	return
}

func QosFlowPerTNLInformationItem() (value ngapType.QosFlowPerTNLInformationItem) {
	value.QosFlowPerTNLInformation = QosFlowPerTNLInformation()
	return
}

func QosFlowSetupRequestList() (value ngapType.QosFlowSetupRequestList) {
	item := QosFlowSetupRequestItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowSetupRequestItem() (value ngapType.QosFlowSetupRequestItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.QosFlowLevelQosParameters = QosFlowLevelQosParameters()
	value.ERABID = new(ngapType.ERABID)
	*value.ERABID = ERABID()
	return
}

func QosFlowListWithDataForwarding() (value ngapType.QosFlowListWithDataForwarding) {
	item := QosFlowItemWithDataForwarding()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowItemWithDataForwarding() (value ngapType.QosFlowItemWithDataForwarding) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.DataForwardingAccepted = new(ngapType.DataForwardingAccepted)
	*value.DataForwardingAccepted = DataForwardingAccepted()
	return
}

func QosFlowToBeForwardedList() (value ngapType.QosFlowToBeForwardedList) {
	item := QosFlowToBeForwardedItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QosFlowToBeForwardedItem() (value ngapType.QosFlowToBeForwardedItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	return
}

func QoSFlowsUsageReportList() (value ngapType.QoSFlowsUsageReportList) {
	item := QoSFlowsUsageReportItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func QoSFlowsUsageReportItem() (value ngapType.QoSFlowsUsageReportItem) {
	value.QosFlowIdentifier = QosFlowIdentifier()
	value.QoSFlowsTimedReportList = VolumeTimedReportList()
	return
}

func RANNodeName() (value ngapType.RANNodeName) {
	value.Value = "-haodhh-vht5gc-"
	value.Value = "-haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc-"
	return
}

func RANPagingPriority() (value ngapType.RANPagingPriority) {
	value.Value = 1
	value.Value = 256
	return
}

func RANStatusTransferTransparentContainer() (value ngapType.RANStatusTransferTransparentContainer) {
	value.DRBsSubjectToStatusTransferList = DRBsSubjectToStatusTransferList()
	return
}

func RANUENGAPID() (value ngapType.RANUENGAPID) {
	value.Value = 0
	value.Value = 4294967295
	return
}

func RATInformation() (value ngapType.RATInformation) {
	value.Value = ngapType.RATInformationUnlicensed
	return
}

func RATRestrictions() (value ngapType.RATRestrictions) {
	item := RATRestrictionsItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func RATRestrictionsItem() (value ngapType.RATRestrictionsItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.RATRestrictionInformation = RATRestrictionInformation()
	return
}

func RATRestrictionInformation() (value ngapType.RATRestrictionInformation) {
	value.Value.BitLength = 8
	value.Value.Bytes = []byte{171}
	return
}

func RecommendedCellsForPaging() (value ngapType.RecommendedCellsForPaging) {
	value.RecommendedCellList = RecommendedCellList()
	return
}

func RecommendedCellList() (value ngapType.RecommendedCellList) {
	item := RecommendedCellItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func RecommendedCellItem() (value ngapType.RecommendedCellItem) {
	value.NGRANCGI = NGRANCGI()
	return
}

func RecommendedRANNodesForPaging() (value ngapType.RecommendedRANNodesForPaging) {
	value.RecommendedRANNodeList = RecommendedRANNodeList()
	return
}

func RecommendedRANNodeList() (value ngapType.RecommendedRANNodeList) {
	item := RecommendedRANNodeItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func RecommendedRANNodeItem() (value ngapType.RecommendedRANNodeItem) {
	value.AMFPagingTarget = AMFPagingTarget()
	return
}

func RedirectionVoiceFallback() (value ngapType.RedirectionVoiceFallback) {
	value.Value = ngapType.RedirectionVoiceFallbackPossible
	value.Value = ngapType.RedirectionVoiceFallbackNotPossible
	return
}

func ReflectiveQosAttribute() (value ngapType.ReflectiveQosAttribute) {
	value.Value = ngapType.ReflectiveQosAttributeSubjectTo
	return
}

func RelativeAMFCapacity() (value ngapType.RelativeAMFCapacity) {
	value.Value = 0
	value.Value = 255
	return
}

func ReportArea() (value ngapType.ReportArea) {
	value.Value = ngapType.ReportAreaCell
	return
}

func RepetitionPeriod() (value ngapType.RepetitionPeriod) {
	value.Value = 0
	value.Value = 131071
	return
}

func ResetAll() (value ngapType.ResetAll) {
	value.Value = ngapType.ResetAllResetAll
	return
}

func ResetType() (value ngapType.ResetType) {
	value.Present = ngapType.ResetTypePresentNGInterface
	value.NGInterface = new(ngapType.ResetAll)
	*value.NGInterface = ResetAll()
	value.Present = ngapType.ResetTypePresentPartOfNGInterface
	value.PartOfNGInterface = new(ngapType.UEAssociatedLogicalNGConnectionList)
	*value.PartOfNGInterface = UEAssociatedLogicalNGConnectionList()
	return
}

func RNCID() (value ngapType.RNCID) {
	value.Value = 0
	value.Value = 4095
	return
}

func RoutingID() (value ngapType.RoutingID) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func RRCContainer() (value ngapType.RRCContainer) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func RRCEstablishmentCause() (value ngapType.RRCEstablishmentCause) {
	value.Value = ngapType.RRCEstablishmentCauseEmergency
	value.Value = ngapType.RRCEstablishmentCauseHighPriorityAccess
	value.Value = ngapType.RRCEstablishmentCauseMtAccess
	value.Value = ngapType.RRCEstablishmentCauseMoSignalling
	value.Value = ngapType.RRCEstablishmentCauseMoData
	value.Value = ngapType.RRCEstablishmentCauseMoVoiceCall
	value.Value = ngapType.RRCEstablishmentCauseMoVideoCall
	value.Value = ngapType.RRCEstablishmentCauseMoSMS
	value.Value = ngapType.RRCEstablishmentCauseMpsPriorityAccess
	value.Value = ngapType.RRCEstablishmentCauseMcsPriorityAccess
	value.Value = ngapType.RRCEstablishmentCauseNotAvailable
	return
}

func RRCInactiveTransitionReportRequest() (value ngapType.RRCInactiveTransitionReportRequest) {
	value.Value = ngapType.RRCInactiveTransitionReportRequestSubsequentStateTransitionReport
	value.Value = ngapType.RRCInactiveTransitionReportRequestSingleRrcConnectedStateReport
	value.Value = ngapType.RRCInactiveTransitionReportRequestCancelReport
	return
}

func RRCState() (value ngapType.RRCState) {
	value.Value = ngapType.RRCStateInactive
	value.Value = ngapType.RRCStateConnected
	return
}

func RIMInformationTransfer() (value ngapType.RIMInformationTransfer) {
	value.TargetRANNodeID = TargetRANNodeID()
	value.SourceRANNodeID = SourceRANNodeID()
	value.RIMInformation = RIMInformation()
	return
}

func RIMInformation() (value ngapType.RIMInformation) {
	value.TargetgNBSetID = GNBSetID()
	return
}

func GNBSetID() (value ngapType.GNBSetID) {
	value.Value.BitLength = 22
	value.Value.Bytes = []byte{171,101,255}
	return
}

func SCTPTLAs() (value ngapType.SCTPTLAs) {
	item := TransportLayerAddress()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item)
	return
}

func SD() (value ngapType.SD) {
	value.Value = []byte{171,1,2}
	return
}

func SecondaryRATUsageInformation() (value ngapType.SecondaryRATUsageInformation) {
	value.PDUSessionUsageReport = new(ngapType.PDUSessionUsageReport)
	*value.PDUSessionUsageReport = PDUSessionUsageReport()
	value.QosFlowsUsageReportList = new(ngapType.QoSFlowsUsageReportList)
	*value.QosFlowsUsageReportList = QoSFlowsUsageReportList()
	return
}

func SecondaryRATDataUsageReportTransfer() (value ngapType.SecondaryRATDataUsageReportTransfer) {
	value.SecondaryRATUsageInformation = new(ngapType.SecondaryRATUsageInformation)
	*value.SecondaryRATUsageInformation = SecondaryRATUsageInformation()
	return
}

func SecurityContext() (value ngapType.SecurityContext) {
	value.NextHopChainingCount = NextHopChainingCount()
	value.NextHopNH = SecurityKey()
	return
}

func SecurityIndication() (value ngapType.SecurityIndication) {
	value.IntegrityProtectionIndication = IntegrityProtectionIndication()
	value.ConfidentialityProtectionIndication = ConfidentialityProtectionIndication()
	value.MaximumIntegrityProtectedDataRateUL = new(ngapType.MaximumIntegrityProtectedDataRate)
	*value.MaximumIntegrityProtectedDataRateUL = MaximumIntegrityProtectedDataRate()
	return
}

func SecurityKey() (value ngapType.SecurityKey) {
	value.Value.BitLength = 256
	value.Value.Bytes = []byte{171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255}
	return
}

func SecurityResult() (value ngapType.SecurityResult) {
	value.IntegrityProtectionResult = IntegrityProtectionResult()
	value.ConfidentialityProtectionResult = ConfidentialityProtectionResult()
	return
}

func SerialNumber() (value ngapType.SerialNumber) {
	value.Value.BitLength = 16
	value.Value.Bytes = []byte{171,255}
	return
}

func ServedGUAMIList() (value ngapType.ServedGUAMIList) {
	item := ServedGUAMIItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func ServedGUAMIItem() (value ngapType.ServedGUAMIItem) {
	value.GUAMI = GUAMI()
	value.BackupAMFName = new(ngapType.AMFName)
	*value.BackupAMFName = AMFName()
	return
}

func ServiceAreaInformation() (value ngapType.ServiceAreaInformation) {
	item := ServiceAreaInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func ServiceAreaInformationItem() (value ngapType.ServiceAreaInformationItem) {
	value.PLMNIdentity = PLMNIdentity()
	value.AllowedTACs = new(ngapType.AllowedTACs)
	*value.AllowedTACs = AllowedTACs()
	value.NotAllowedTACs = new(ngapType.NotAllowedTACs)
	*value.NotAllowedTACs = NotAllowedTACs()
	return
}

func SgNBUEX2APID() (value ngapType.SgNBUEX2APID) {
	value.Value = 0
	value.Value = 4294967295
	return
}

func SliceOverloadList() (value ngapType.SliceOverloadList) {
	item := SliceOverloadItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func SliceOverloadItem() (value ngapType.SliceOverloadItem) {
	value.SNSSAI = SNSSAI()
	return
}

func SliceSupportList() (value ngapType.SliceSupportList) {
	item := SliceSupportItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func SliceSupportItem() (value ngapType.SliceSupportItem) {
	value.SNSSAI = SNSSAI()
	return
}

func SNSSAI() (value ngapType.SNSSAI) {
	value.SST = SST()
	value.SD = new(ngapType.SD)
	*value.SD = SD()
	return
}

func SONConfigurationTransfer() (value ngapType.SONConfigurationTransfer) {
	value.TargetRANNodeID = TargetRANNodeID()
	value.SourceRANNodeID = SourceRANNodeID()
	value.SONInformation = SONInformation()
	value.XnTNLConfigurationInfo = new(ngapType.XnTNLConfigurationInfo)
	*value.XnTNLConfigurationInfo = XnTNLConfigurationInfo()
	return
}

func SONInformation() (value ngapType.SONInformation) {
	value.Present = ngapType.SONInformationPresentSONInformationRequest
	value.SONInformationRequest = new(ngapType.SONInformationRequest)
	*value.SONInformationRequest = SONInformationRequest()
	value.Present = ngapType.SONInformationPresentSONInformationReply
	value.SONInformationReply = new(ngapType.SONInformationReply)
	*value.SONInformationReply = SONInformationReply()
	return
}

func SONInformationReply() (value ngapType.SONInformationReply) {
	value.XnTNLConfigurationInfo = new(ngapType.XnTNLConfigurationInfo)
	*value.XnTNLConfigurationInfo = XnTNLConfigurationInfo()
	return
}

func SONInformationRequest() (value ngapType.SONInformationRequest) {
	value.Value = ngapType.SONInformationRequestXnTNLConfigurationInfo
	return
}

func SourceNGRANNodeToTargetNGRANNodeTransparentContainer() (value ngapType.SourceNGRANNodeToTargetNGRANNodeTransparentContainer) {
	value.RRCContainer = RRCContainer()
	value.PDUSessionResourceInformationList = new(ngapType.PDUSessionResourceInformationList)
	*value.PDUSessionResourceInformationList = PDUSessionResourceInformationList()
	value.ERABInformationList = new(ngapType.ERABInformationList)
	*value.ERABInformationList = ERABInformationList()
	value.TargetCellID = NGRANCGI()
	value.IndexToRFSP = new(ngapType.IndexToRFSP)
	*value.IndexToRFSP = IndexToRFSP()
	value.UEHistoryInformation = UEHistoryInformation()
	return
}

func SourceOfUEActivityBehaviourInformation() (value ngapType.SourceOfUEActivityBehaviourInformation) {
	value.Value = ngapType.SourceOfUEActivityBehaviourInformationSubscriptionInformation
	value.Value = ngapType.SourceOfUEActivityBehaviourInformationStatistics
	return
}

func SourceRANNodeID() (value ngapType.SourceRANNodeID) {
	value.GlobalRANNodeID = GlobalRANNodeID()
	value.SelectedTAI = TAI()
	return
}

func SourceToTargetTransparentContainer() (value ngapType.SourceToTargetTransparentContainer) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func SourceToTargetAMFInformationReroute() (value ngapType.SourceToTargetAMFInformationReroute) {
	value.ConfiguredNSSAI = new(ngapType.ConfiguredNSSAI)
	*value.ConfiguredNSSAI = ConfiguredNSSAI()
	value.RejectedNSSAIinPLMN = new(ngapType.RejectedNSSAIinPLMN)
	*value.RejectedNSSAIinPLMN = RejectedNSSAIinPLMN()
	value.RejectedNSSAIinTA = new(ngapType.RejectedNSSAIinTA)
	*value.RejectedNSSAIinTA = RejectedNSSAIinTA()
	return
}

func SRVCCOperationPossible() (value ngapType.SRVCCOperationPossible) {
	value.Value = ngapType.SRVCCOperationPossiblePossible
	return
}

func ConfiguredNSSAI() (value ngapType.ConfiguredNSSAI) {
	value.Value = []byte{171,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127}
	return
}

func RejectedNSSAIinPLMN() (value ngapType.RejectedNSSAIinPLMN) {
	value.Value = []byte{171,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31}
	return
}

func RejectedNSSAIinTA() (value ngapType.RejectedNSSAIinTA) {
	value.Value = []byte{171,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31}
	return
}

func SST() (value ngapType.SST) {
	value.Value = []byte{171}
	return
}

func SupportedTAList() (value ngapType.SupportedTAList) {
	item := SupportedTAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func SupportedTAItem() (value ngapType.SupportedTAItem) {
	value.TAC = TAC()
	value.BroadcastPLMNList = BroadcastPLMNList()
	return
}

func TAC() (value ngapType.TAC) {
	value.Value = []byte{171,1,2}
	return
}

func TAI() (value ngapType.TAI) {
	value.PLMNIdentity = PLMNIdentity()
	value.TAC = TAC()
	return
}

func TAIBroadcastEUTRA() (value ngapType.TAIBroadcastEUTRA) {
	item := TAIBroadcastEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAIBroadcastEUTRAItem() (value ngapType.TAIBroadcastEUTRAItem) {
	value.TAI = TAI()
	value.CompletedCellsInTAIEUTRA = CompletedCellsInTAIEUTRA()
	return
}

func TAIBroadcastNR() (value ngapType.TAIBroadcastNR) {
	item := TAIBroadcastNRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAIBroadcastNRItem() (value ngapType.TAIBroadcastNRItem) {
	value.TAI = TAI()
	value.CompletedCellsInTAINR = CompletedCellsInTAINR()
	return
}

func TAICancelledEUTRA() (value ngapType.TAICancelledEUTRA) {
	item := TAICancelledEUTRAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAICancelledEUTRAItem() (value ngapType.TAICancelledEUTRAItem) {
	value.TAI = TAI()
	value.CancelledCellsInTAIEUTRA = CancelledCellsInTAIEUTRA()
	return
}

func TAICancelledNR() (value ngapType.TAICancelledNR) {
	item := TAICancelledNRItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAICancelledNRItem() (value ngapType.TAICancelledNRItem) {
	value.TAI = TAI()
	value.CancelledCellsInTAINR = CancelledCellsInTAINR()
	return
}

func TAIListForInactive() (value ngapType.TAIListForInactive) {
	item := TAIListForInactiveItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAIListForInactiveItem() (value ngapType.TAIListForInactiveItem) {
	value.TAI = TAI()
	return
}

func TAIListForPaging() (value ngapType.TAIListForPaging) {
	item := TAIListForPagingItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAIListForPagingItem() (value ngapType.TAIListForPagingItem) {
	value.TAI = TAI()
	return
}

func TAIListForRestart() (value ngapType.TAIListForRestart) {
	item := TAI()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TAIListForWarning() (value ngapType.TAIListForWarning) {
	item := TAI()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TargeteNBID() (value ngapType.TargeteNBID) {
	value.GlobalENBID = GlobalNgENBID()
	value.SelectedEPSTAI = EPSTAI()
	return
}

func TargetID() (value ngapType.TargetID) {
	value.Present = ngapType.TargetIDPresentTargetRANNodeID
	value.TargetRANNodeID = new(ngapType.TargetRANNodeID)
	*value.TargetRANNodeID = TargetRANNodeID()
	value.Present = ngapType.TargetIDPresentTargeteNBID
	value.TargeteNBID = new(ngapType.TargeteNBID)
	*value.TargeteNBID = TargeteNBID()
	return
}

func TargetNGRANNodeToSourceNGRANNodeTransparentContainer() (value ngapType.TargetNGRANNodeToSourceNGRANNodeTransparentContainer) {
	value.RRCContainer = RRCContainer()
	return
}

func TargetRANNodeID() (value ngapType.TargetRANNodeID) {
	value.GlobalRANNodeID = GlobalRANNodeID()
	value.SelectedTAI = TAI()
	return
}

func TargetRNCID() (value ngapType.TargetRNCID) {
	value.LAI = LAI()
	value.RNCID = RNCID()
	value.ExtendedRNCID = new(ngapType.ExtendedRNCID)
	*value.ExtendedRNCID = ExtendedRNCID()
	return
}

func TargetToSourceTransparentContainer() (value ngapType.TargetToSourceTransparentContainer) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func TimerApproachForGUAMIRemoval() (value ngapType.TimerApproachForGUAMIRemoval) {
	value.Value = ngapType.TimerApproachForGUAMIRemovalApplyTimer
	return
}

func TimeStamp() (value ngapType.TimeStamp) {
	value.Value = []byte{171,1,2,3}
	return
}

func TimeToWait() (value ngapType.TimeToWait) {
	value.Value = ngapType.TimeToWaitV1s
	value.Value = ngapType.TimeToWaitV2s
	value.Value = ngapType.TimeToWaitV5s
	value.Value = ngapType.TimeToWaitV10s
	value.Value = ngapType.TimeToWaitV20s
	value.Value = ngapType.TimeToWaitV60s
	return
}

func TimeUEStayedInCell() (value ngapType.TimeUEStayedInCell) {
	value.Value = 0
	value.Value = 4095
	return
}

func TimeUEStayedInCellEnhancedGranularity() (value ngapType.TimeUEStayedInCellEnhancedGranularity) {
	value.Value = 0
	value.Value = 40950
	return
}

func TNLAddressWeightFactor() (value ngapType.TNLAddressWeightFactor) {
	value.Value = 0
	value.Value = 255
	return
}

func TNLAssociationList() (value ngapType.TNLAssociationList) {
	item := TNLAssociationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func TNLAssociationItem() (value ngapType.TNLAssociationItem) {
	value.TNLAssociationAddress = CPTransportLayerInformation()
	value.Cause = Cause()
	return
}

func TNLAssociationUsage() (value ngapType.TNLAssociationUsage) {
	value.Value = ngapType.TNLAssociationUsageUe
	value.Value = ngapType.TNLAssociationUsageNonUe
	value.Value = ngapType.TNLAssociationUsageBoth
	return
}

func TraceActivation() (value ngapType.TraceActivation) {
	value.NGRANTraceID = NGRANTraceID()
	value.InterfacesToTrace = InterfacesToTrace()
	value.TraceDepth = TraceDepth()
	value.TraceCollectionEntityIPAddress = TransportLayerAddress()
	return
}

func TraceDepth() (value ngapType.TraceDepth) {
	value.Value = ngapType.TraceDepthMinimum
	value.Value = ngapType.TraceDepthMedium
	value.Value = ngapType.TraceDepthMaximum
	value.Value = ngapType.TraceDepthMinimumWithoutVendorSpecificExtension
	value.Value = ngapType.TraceDepthMediumWithoutVendorSpecificExtension
	value.Value = ngapType.TraceDepthMaximumWithoutVendorSpecificExtension
	return
}

func TrafficLoadReductionIndication() (value ngapType.TrafficLoadReductionIndication) {
	value.Value = 1
	value.Value = 99
	return
}

func TransportLayerAddress() (value ngapType.TransportLayerAddress) {
	value.Value.BitLength = 1
	value.Value.Bytes = []byte{171}
	value.Value.BitLength = 160
	value.Value.Bytes = []byte{171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255}
	return
}

func TypeOfError() (value ngapType.TypeOfError) {
	value.Value = ngapType.TypeOfErrorNotUnderstood
	value.Value = ngapType.TypeOfErrorMissing
	return
}

func UEAggregateMaximumBitRate() (value ngapType.UEAggregateMaximumBitRate) {
	value.UEAggregateMaximumBitRateDL = BitRate()
	value.UEAggregateMaximumBitRateUL = BitRate()
	return
}

func UEAssociatedLogicalNGConnectionList() (value ngapType.UEAssociatedLogicalNGConnectionList) {
	item := UEAssociatedLogicalNGConnectionItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func UEAssociatedLogicalNGConnectionItem() (value ngapType.UEAssociatedLogicalNGConnectionItem) {
	value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*value.AMFUENGAPID = AMFUENGAPID()
	value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*value.RANUENGAPID = RANUENGAPID()
	return
}

func UEContextRequest() (value ngapType.UEContextRequest) {
	value.Value = ngapType.UEContextRequestRequested
	return
}

func UEHistoryInformation() (value ngapType.UEHistoryInformation) {
	item := LastVisitedCellItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func UEIdentityIndexValue() (value ngapType.UEIdentityIndexValue) {
	value.Present = ngapType.UEIdentityIndexValuePresentIndexLength10
	value.IndexLength10 = new(ngapType.BitString)
	value.IndexLength10.BitLength = 10
	value.IndexLength10.Bytes = []byte{171,219}
	return
}

func UENGAPIDs() (value ngapType.UENGAPIDs) {
	value.Present = ngapType.UENGAPIDsPresentUENGAPIDPair
	value.UENGAPIDPair = new(ngapType.UENGAPIDPair)
	*value.UENGAPIDPair = UENGAPIDPair()
	value.Present = ngapType.UENGAPIDsPresentAMFUENGAPID
	value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*value.AMFUENGAPID = AMFUENGAPID()
	return
}

func UENGAPIDPair() (value ngapType.UENGAPIDPair) {
	value.AMFUENGAPID = AMFUENGAPID()
	value.RANUENGAPID = RANUENGAPID()
	return
}

func UEPagingIdentity() (value ngapType.UEPagingIdentity) {
	value.Present = ngapType.UEPagingIdentityPresentFiveGSTMSI
	value.FiveGSTMSI = new(ngapType.FiveGSTMSI)
	*value.FiveGSTMSI = FiveGSTMSI()
	return
}

func UEPresence() (value ngapType.UEPresence) {
	value.Value = ngapType.UEPresenceIn
	value.Value = ngapType.UEPresenceOut
	value.Value = ngapType.UEPresenceUnknown
	return
}

func UEPresenceInAreaOfInterestList() (value ngapType.UEPresenceInAreaOfInterestList) {
	item := UEPresenceInAreaOfInterestItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func UEPresenceInAreaOfInterestItem() (value ngapType.UEPresenceInAreaOfInterestItem) {
	value.LocationReportingReferenceID = LocationReportingReferenceID()
	value.UEPresence = UEPresence()
	return
}

func UERadioCapability() (value ngapType.UERadioCapability) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func UERadioCapabilityForPaging() (value ngapType.UERadioCapabilityForPaging) {
	value.UERadioCapabilityForPagingOfNR = new(ngapType.UERadioCapabilityForPagingOfNR)
	*value.UERadioCapabilityForPagingOfNR = UERadioCapabilityForPagingOfNR()
	value.UERadioCapabilityForPagingOfEUTRA = new(ngapType.UERadioCapabilityForPagingOfEUTRA)
	*value.UERadioCapabilityForPagingOfEUTRA = UERadioCapabilityForPagingOfEUTRA()
	return
}

func UERadioCapabilityForPagingOfNR() (value ngapType.UERadioCapabilityForPagingOfNR) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func UERadioCapabilityForPagingOfEUTRA() (value ngapType.UERadioCapabilityForPagingOfEUTRA) {
	value.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}
	value.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	return
}

func UERetentionInformation() (value ngapType.UERetentionInformation) {
	value.Value = ngapType.UERetentionInformationUesRetained
	return
}

func UESecurityCapabilities() (value ngapType.UESecurityCapabilities) {
	value.NRencryptionAlgorithms = NRencryptionAlgorithms()
	value.NRintegrityProtectionAlgorithms = NRintegrityProtectionAlgorithms()
	value.EUTRAencryptionAlgorithms = EUTRAencryptionAlgorithms()
	value.EUTRAintegrityProtectionAlgorithms = EUTRAintegrityProtectionAlgorithms()
	return
}

func ULNGUUPTNLModifyList() (value ngapType.ULNGUUPTNLModifyList) {
	item := ULNGUUPTNLModifyItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item)
	return
}

func ULNGUUPTNLModifyItem() (value ngapType.ULNGUUPTNLModifyItem) {
	value.ULNGUUPTNLInformation = UPTransportLayerInformation()
	value.DLNGUUPTNLInformation = UPTransportLayerInformation()
	return
}

func UnavailableGUAMIList() (value ngapType.UnavailableGUAMIList) {
	item := UnavailableGUAMIItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func UnavailableGUAMIItem() (value ngapType.UnavailableGUAMIItem) {
	value.GUAMI = GUAMI()
	value.TimerApproachForGUAMIRemoval = new(ngapType.TimerApproachForGUAMIRemoval)
	*value.TimerApproachForGUAMIRemoval = TimerApproachForGUAMIRemoval()
	value.BackupAMFName = new(ngapType.AMFName)
	*value.BackupAMFName = AMFName()
	return
}

func ULForwarding() (value ngapType.ULForwarding) {
	value.Value = ngapType.ULForwardingUlForwardingProposed
	return
}

func UPTransportLayerInformation() (value ngapType.UPTransportLayerInformation) {
	value.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	value.GTPTunnel = new(ngapType.GTPTunnel)
	*value.GTPTunnel = GTPTunnel()
	return
}

func UPTransportLayerInformationList() (value ngapType.UPTransportLayerInformationList) {
	item := UPTransportLayerInformationItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item)
	return
}

func UPTransportLayerInformationItem() (value ngapType.UPTransportLayerInformationItem) {
	value.NGUUPTNLInformation = UPTransportLayerInformation()
	return
}

func UPTransportLayerInformationPairList() (value ngapType.UPTransportLayerInformationPairList) {
	item := UPTransportLayerInformationPairItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item)
	return
}

func UPTransportLayerInformationPairItem() (value ngapType.UPTransportLayerInformationPairItem) {
	value.ULNGUUPTNLInformation = UPTransportLayerInformation()
	value.DLNGUUPTNLInformation = UPTransportLayerInformation()
	return
}

func UserLocationInformation() (value ngapType.UserLocationInformation) {
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

func UserLocationInformationEUTRA() (value ngapType.UserLocationInformationEUTRA) {
	value.EUTRACGI = EUTRACGI()
	value.TAI = TAI()
	value.TimeStamp = new(ngapType.TimeStamp)
	*value.TimeStamp = TimeStamp()
	return
}

func UserLocationInformationN3IWF() (value ngapType.UserLocationInformationN3IWF) {
	value.IPAddress = TransportLayerAddress()
	value.PortNumber = PortNumber()
	return
}

func UserLocationInformationNR() (value ngapType.UserLocationInformationNR) {
	value.NRCGI = NRCGI()
	value.TAI = TAI()
	value.TimeStamp = new(ngapType.TimeStamp)
	*value.TimeStamp = TimeStamp()
	return
}

func UserPlaneSecurityInformation() (value ngapType.UserPlaneSecurityInformation) {
	value.SecurityResult = SecurityResult()
	value.SecurityIndication = SecurityIndication()
	return
}

func VolumeTimedReportList() (value ngapType.VolumeTimedReportList) {
	item := VolumeTimedReportItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item)
	return
}

func VolumeTimedReportItem() (value ngapType.VolumeTimedReportItem) {
	return
}

func WarningAreaCoordinates() (value ngapType.WarningAreaCoordinates) {
	value.Value = []byte{171}
	value.Value = []byte{219,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255}
	return
}

func WarningAreaList() (value ngapType.WarningAreaList) {
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

func WarningMessageContents() (value ngapType.WarningMessageContents) {
	value.Value = []byte{171}
	value.Value = []byte{219,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,
		128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,
		192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,
		0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,
		64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127}
	return
}

func WarningSecurityInfo() (value ngapType.WarningSecurityInfo) {
	value.Value = []byte{171,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49}
	return
}

func WarningType() (value ngapType.WarningType) {
	value.Value = []byte{171,1}
	return
}

func XnExtTLAs() (value ngapType.XnExtTLAs) {
	item := XnExtTLAItem()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func XnExtTLAItem() (value ngapType.XnExtTLAItem) {
	value.IPsecTLA = new(ngapType.TransportLayerAddress)
	*value.IPsecTLA = TransportLayerAddress()
	value.GTPTLAs = new(ngapType.XnGTPTLAs)
	*value.GTPTLAs = XnGTPTLAs()
	return
}

func XnGTPTLAs() (value ngapType.XnGTPTLAs) {
	item := TransportLayerAddress()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item,item,item,item,item,item,item,item,item,item,item,item,item,item,item)
	return
}

func XnTLAs() (value ngapType.XnTLAs) {
	item := TransportLayerAddress()
	value.List = append(value.List,item)
	value.List = append(value.List,
		item,item)
	return
}

func XnTNLConfigurationInfo() (value ngapType.XnTNLConfigurationInfo) {
	value.XnTransportLayerAddresses = XnTLAs()
	value.XnExtendedTransportLayerAddresses = new(ngapType.XnExtTLAs)
	*value.XnExtendedTransportLayerAddresses = XnExtTLAs()
	return
}
