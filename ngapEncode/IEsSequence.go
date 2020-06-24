package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AdditionalDLUPTNLInformationForHOItem (value ngapType.AdditionalDLUPTNLInformationForHOItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.AdditionalDLForwardingUPTNLInformation != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.AdditionalDLNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItem: " + err.Error())
	}
	binData, bitEnd, err = QosFlowListWithDataForwarding(value.AdditionalQosFlowSetupResponseList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItem: " + err.Error())
	}
	if value.AdditionalDLForwardingUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.AdditionalDLForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllocationAndRetentionPriority (value ngapType.AllocationAndRetentionPriority, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PriorityLevelARP(value.PriorityLevelARP, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriority: " + err.Error())
	}
	binData, bitEnd, err = PreEmptionCapability(value.PreEmptionCapability, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriority: " + err.Error())
	}
	binData, bitEnd, err = PreEmptionVulnerability(value.PreEmptionVulnerability, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriority: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AllocationAndRetentionPriority: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllowedNSSAIItem (value ngapType.AllowedNSSAIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAllowedNSSAIItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AllowedNSSAIItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationSetupItem (value ngapType.AMFTNLAssociationSetupItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = CPTransportLayerInformation(value.AMFTNLAssociationAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationSetupItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToAddItem (value ngapType.AMFTNLAssociationToAddItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TNLAssociationUsage != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = CPTransportLayerInformation(value.AMFTNLAssociationAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItem: " + err.Error())
	}
	if value.TNLAssociationUsage != nil {
		binData, bitEnd, err = TNLAssociationUsage(*value.TNLAssociationUsage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToAddItem: " + err.Error())
		}
	}
	binData, bitEnd, err = TNLAddressWeightFactor(value.TNLAddressWeightFactor, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToAddItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToRemoveItem (value ngapType.AMFTNLAssociationToRemoveItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = CPTransportLayerInformation(value.AMFTNLAssociationAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToUpdateItem (value ngapType.AMFTNLAssociationToUpdateItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TNLAddressWeightFactor != nil {
		option += 1<<1
	}
	if value.TNLAssociationUsage != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = CPTransportLayerInformation(value.AMFTNLAssociationAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItem: " + err.Error())
	}
	if value.TNLAssociationUsage != nil {
		binData, bitEnd, err = TNLAssociationUsage(*value.TNLAssociationUsage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItem: " + err.Error())
		}
	}
	if value.TNLAddressWeightFactor != nil {
		binData, bitEnd, err = TNLAddressWeightFactor(*value.TNLAddressWeightFactor, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterest (value ngapType.AreaOfInterest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.AreaOfInterestRANNodeList != nil {
		option += 1<<1
	}
	if value.AreaOfInterestCellList != nil {
		option += 1<<2
	}
	if value.AreaOfInterestTAIList != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	if value.AreaOfInterestTAIList != nil {
		binData, bitEnd, err = AreaOfInterestTAIList(*value.AreaOfInterestTAIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterest: " + err.Error())
		}
	}
	if value.AreaOfInterestCellList != nil {
		binData, bitEnd, err = AreaOfInterestCellList(*value.AreaOfInterestCellList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterest: " + err.Error())
		}
	}
	if value.AreaOfInterestRANNodeList != nil {
		binData, bitEnd, err = AreaOfInterestRANNodeList(*value.AreaOfInterestRANNodeList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterest: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAreaOfInterestExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterest: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestCellItem (value ngapType.AreaOfInterestCellItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NGRANCGI(value.NGRANCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAreaOfInterestCellItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestCellItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestItem (value ngapType.AreaOfInterestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = AreaOfInterest(value.AreaOfInterest, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItem: " + err.Error())
	}
	binData, bitEnd, err = LocationReportingReferenceID(value.LocationReportingReferenceID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAreaOfInterestItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestRANNodeItem (value ngapType.AreaOfInterestRANNodeItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = GlobalRANNodeID(value.GlobalRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestRANNodeItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestTAIItem (value ngapType.AreaOfInterestTAIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestTAIItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AssistanceDataForPaging (value ngapType.AssistanceDataForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.PagingAttemptInformation != nil {
		option += 1<<1
	}
	if value.AssistanceDataForRecommendedCells != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.AssistanceDataForRecommendedCells != nil {
		binData, bitEnd, err = AssistanceDataForRecommendedCells(*value.AssistanceDataForRecommendedCells, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AssistanceDataForPaging: " + err.Error())
		}
	}
	if value.PagingAttemptInformation != nil {
		binData, bitEnd, err = PagingAttemptInformation(*value.PagingAttemptInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AssistanceDataForPaging: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAssistanceDataForPagingExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AssistanceDataForPaging: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AssistanceDataForRecommendedCells (value ngapType.AssistanceDataForRecommendedCells, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = RecommendedCellsForPaging(value.RecommendedCellsForPaging, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCells: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AssistanceDataForRecommendedCells: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AssociatedQosFlowItem (value ngapType.AssociatedQosFlowItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.QosFlowMappingIndication != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssociatedQosFlowItem: " + err.Error())
	}
	if *value.QosFlowMappingIndication < 0 || *value.QosFlowMappingIndication > 2 {
		return binData, bitEnd, errors.New("AssociatedQosFlowItem: QosFlowMappingIndication: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(*value.QosFlowMappingIndication, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerAssociatedQosFlowItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AssociatedQosFlowItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func BroadcastPLMNItem (value ngapType.BroadcastPLMNItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItem: " + err.Error())
	}
	binData, bitEnd, err = SliceSupportList(value.TAISliceSupportList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerBroadcastPLMNItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastPLMNItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAIEUTRAItem (value ngapType.CancelledCellsInEAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = NumberOfBroadcasts(value.NumberOfBroadcasts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINRItem (value ngapType.CancelledCellsInEAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItem: " + err.Error())
	}
	binData, bitEnd, err = NumberOfBroadcasts(value.NumberOfBroadcasts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInEAINRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAIEUTRAItem (value ngapType.CancelledCellsInTAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = NumberOfBroadcasts(value.NumberOfBroadcasts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINRItem (value ngapType.CancelledCellsInTAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItem: " + err.Error())
	}
	binData, bitEnd, err = NumberOfBroadcasts(value.NumberOfBroadcasts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInTAINRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDBroadcastEUTRAItem (value ngapType.CellIDBroadcastEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDBroadcastNRItem (value ngapType.CellIDBroadcastNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDBroadcastNRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDCancelledEUTRAItem (value ngapType.CellIDCancelledEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = NumberOfBroadcasts(value.NumberOfBroadcasts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDCancelledEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDCancelledNRItem (value ngapType.CellIDCancelledNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItem: " + err.Error())
	}
	binData, bitEnd, err = NumberOfBroadcasts(value.NumberOfBroadcasts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCellIDCancelledNRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDCancelledNRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellType (value ngapType.CellType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = CellSize(value.CellSize, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellType: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCellTypeExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellType: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CNAssistedRANTuning (value ngapType.CNAssistedRANTuning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ExpectedUEBehaviour != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	if value.ExpectedUEBehaviour != nil {
		binData, bitEnd, err = ExpectedUEBehaviour(*value.ExpectedUEBehaviour, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CNAssistedRANTuning: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCNAssistedRANTuningExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CNAssistedRANTuning: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CNTypeRestrictionsForEquivalentItem (value ngapType.CNTypeRestrictionsForEquivalentItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PlmnIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItem: " + err.Error())
	}
	if value.CnType < 0 || value.CnType > 2 {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItem: CnType: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.CnType, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRAItem (value ngapType.CompletedCellsInEAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINRItem (value ngapType.CompletedCellsInEAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInEAINRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRAItem (value ngapType.CompletedCellsInTAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINRItem (value ngapType.CompletedCellsInTAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInTAINRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CoreNetworkAssistanceInformationForInactive (value ngapType.CoreNetworkAssistanceInformationForInactive, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ExpectedUEBehaviour != nil {
		option += 1<<1
	}
	if value.MICOModeIndication != nil {
		option += 1<<2
	}
	if value.UESpecificDRX != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	binData, bitEnd, err = UEIdentityIndexValue(value.UEIdentityIndexValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
	}
	if value.UESpecificDRX != nil {
		binData, bitEnd, err = PagingDRX(*value.UESpecificDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
		}
	}
	binData, bitEnd, err = PeriodicRegistrationUpdateTimer(value.PeriodicRegistrationUpdateTimer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
	}
	if value.MICOModeIndication != nil {
		binData, bitEnd, err = MICOModeIndication(*value.MICOModeIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
		}
	}
	binData, bitEnd, err = TAIListForInactive(value.TAIListForInactive, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
	}
	if value.ExpectedUEBehaviour != nil {
		binData, bitEnd, err = ExpectedUEBehaviour(*value.ExpectedUEBehaviour, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactive: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN12 (value ngapType.COUNTValueForPDCPSN12, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if value.PDCPSN12 < 0 || value.PDCPSN12 > 4095 {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12: PDCPSN12: Integer: valueMin:0,valueMax:4095")
	}
	binData, bitEnd = EncodeInteger(value.PDCPSN12, 0, binData, bitEnd)
	if value.HFNPDCPSN12 < 0 || value.HFNPDCPSN12 > 1048575 {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12: HFNPDCPSN12: Integer: valueMin:0,valueMax:1048575")
	}
	binData, bitEnd = EncodeInteger(value.HFNPDCPSN12, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("COUNTValueForPDCPSN12: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN18 (value ngapType.COUNTValueForPDCPSN18, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if value.PDCPSN18 < 0 || value.PDCPSN18 > 262143 {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18: PDCPSN18: Integer: valueMin:0,valueMax:262143")
	}
	binData, bitEnd = EncodeInteger(value.PDCPSN18, 0, binData, bitEnd)
	if value.HFNPDCPSN18 < 0 || value.HFNPDCPSN18 > 16383 {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18: HFNPDCPSN18: Integer: valueMin:0,valueMax:16383")
	}
	binData, bitEnd = EncodeInteger(value.HFNPDCPSN18, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("COUNTValueForPDCPSN18: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CriticalityDiagnostics (value ngapType.CriticalityDiagnostics, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.IEsCriticalityDiagnostics != nil {
		option += 1<<1
	}
	if value.ProcedureCriticality != nil {
		option += 1<<2
	}
	if value.TriggeringMessage != nil {
		option += 1<<3
	}
	if value.ProcedureCode != nil {
		option += 1<<4
	}
	binData, bitEnd = EncodeOption(option, 6, binData, bitEnd)
	if value.ProcedureCode != nil {
		binData, bitEnd, err = ProcedureCode(*value.ProcedureCode, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnostics: " + err.Error())
		}
	}
	if value.TriggeringMessage != nil {
		binData, bitEnd, err = TriggeringMessage(*value.TriggeringMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnostics: " + err.Error())
		}
	}
	if value.ProcedureCriticality != nil {
		binData, bitEnd, err = Criticality(*value.ProcedureCriticality, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnostics: " + err.Error())
		}
	}
	if value.IEsCriticalityDiagnostics != nil {
		binData, bitEnd, err = CriticalityDiagnosticsIEList(*value.IEsCriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnostics: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCriticalityDiagnosticsExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnostics: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CriticalityDiagnosticsIEItem (value ngapType.CriticalityDiagnosticsIEItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Criticality(value.IECriticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItem: " + err.Error())
	}
	binData, bitEnd, err = ProtocolIEID(value.IEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItem: " + err.Error())
	}
	binData, bitEnd, err = TypeOfError(value.TypeOfError, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DataForwardingResponseDRBItem (value ngapType.DataForwardingResponseDRBItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ULForwardingUPTNLInformation != nil {
		option += 1<<1
	}
	if value.DLForwardingUPTNLInformation != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = DRBID(value.DRBID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseDRBItem: " + err.Error())
	}
	if value.DLForwardingUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.DLForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseDRBItem: " + err.Error())
		}
	}
	if value.ULForwardingUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.ULForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseDRBItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseDRBItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DataForwardingResponseERABListItem (value ngapType.DataForwardingResponseERABListItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DLForwarding != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = ERABID(value.ERABID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseERABListItem: " + err.Error())
	}
	if value.DLForwarding != nil {
		binData, bitEnd, err = DLForwarding(*value.DLForwarding, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseERABListItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseERABListItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBsSubjectToStatusTransferItem (value ngapType.DRBsSubjectToStatusTransferItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = DRBID(value.DRBID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItem: " + err.Error())
	}
	binData, bitEnd, err = DRBStatusUL(value.DRBStatusUL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItem: " + err.Error())
	}
	binData, bitEnd, err = DRBStatusDL(value.DRBStatusDL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusDL12 (value ngapType.DRBStatusDL12, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN12(value.DLCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusDL12ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL12: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusDL18 (value ngapType.DRBStatusDL18, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN18(value.DLCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusDL18ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL18: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusUL12 (value ngapType.DRBStatusUL12, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ReceiveStatusOfULPDCPSDUs != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN12(value.ULCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12: " + err.Error())
	}
	if value.ReceiveStatusOfULPDCPSDUs.BitLength < 1 || value.ReceiveStatusOfULPDCPSDUs.BitLength > 2048 || len(value.ReceiveStatusOfULPDCPSDUs.Bytes) != int((value.ReceiveStatusOfULPDCPSDUs.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("DRBStatusUL12: ReceiveStatusOfULPDCPSDUs: BitString: sizeMin:1,sizeMax:2048")
	}
	binData, bitEnd = EncodeBitStringExt(*value.ReceiveStatusOfULPDCPSDUs, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusUL12ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL12: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusUL18 (value ngapType.DRBStatusUL18, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ReceiveStatusOfULPDCPSDUs != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN18(value.ULCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18: " + err.Error())
	}
	if value.ReceiveStatusOfULPDCPSDUs.BitLength < 1 || value.ReceiveStatusOfULPDCPSDUs.BitLength > 131072 || len(value.ReceiveStatusOfULPDCPSDUs.Bytes) != int((value.ReceiveStatusOfULPDCPSDUs.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("DRBStatusUL18: ReceiveStatusOfULPDCPSDUs: BitString: sizeMin:1,sizeMax:131072")
	}
	binData, bitEnd = EncodeBitStringExt(*value.ReceiveStatusOfULPDCPSDUs, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusUL18ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL18: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBsToQosFlowsMappingItem (value ngapType.DRBsToQosFlowsMappingItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = DRBID(value.DRBID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItem: " + err.Error())
	}
	binData, bitEnd, err = AssociatedQosFlowList(value.AssociatedQosFlowList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func Dynamic5QIDescriptor (value ngapType.Dynamic5QIDescriptor, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.MaximumDataBurstVolume != nil {
		option += 1<<1
	}
	if value.AveragingWindow != nil {
		option += 1<<2
	}
	if value.DelayCritical != nil {
		option += 1<<3
	}
	if value.FiveQI != nil {
		option += 1<<4
	}
	binData, bitEnd = EncodeOption(option, 6, binData, bitEnd)
	binData, bitEnd, err = PriorityLevelQos(value.PriorityLevelQos, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
	}
	binData, bitEnd, err = PacketDelayBudget(value.PacketDelayBudget, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
	}
	binData, bitEnd, err = PacketErrorRate(value.PacketErrorRate, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
	}
	if value.FiveQI != nil {
		binData, bitEnd, err = FiveQI(*value.FiveQI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.DelayCritical != nil {
		binData, bitEnd, err = DelayCritical(*value.DelayCritical, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.AveragingWindow != nil {
		binData, bitEnd, err = AveragingWindow(*value.AveragingWindow, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.MaximumDataBurstVolume != nil {
		binData, bitEnd, err = MaximumDataBurstVolume(*value.MaximumDataBurstVolume, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDynamic5QIDescriptorExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Dynamic5QIDescriptor: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastEUTRAItem (value ngapType.EmergencyAreaIDBroadcastEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EmergencyAreaID(value.EmergencyAreaID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInEAIEUTRA(value.CompletedCellsInEAIEUTRA, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastNRItem (value ngapType.EmergencyAreaIDBroadcastNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EmergencyAreaID(value.EmergencyAreaID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItem: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInEAINR(value.CompletedCellsInEAINR, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledEUTRAItem (value ngapType.EmergencyAreaIDCancelledEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EmergencyAreaID(value.EmergencyAreaID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInEAIEUTRA(value.CancelledCellsInEAIEUTRA, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledNRItem (value ngapType.EmergencyAreaIDCancelledNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = EmergencyAreaID(value.EmergencyAreaID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItem: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInEAINR(value.CancelledCellsInEAINR, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyFallbackIndicator (value ngapType.EmergencyFallbackIndicator, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.EmergencyServiceTargetCN != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = EmergencyFallbackRequestIndicator(value.EmergencyFallbackRequestIndicator, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicator: " + err.Error())
	}
	if value.EmergencyServiceTargetCN != nil {
		binData, bitEnd, err = EmergencyServiceTargetCN(*value.EmergencyServiceTargetCN, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyFallbackIndicator: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyFallbackIndicator: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EndpointIPAddressAndPort (value ngapType.EndpointIPAddressAndPort, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeOption(option, 1, binData, bitEnd)
	binData, bitEnd, err = TransportLayerAddress(value.EndpointIPAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EndpointIPAddressAndPort: " + err.Error())
	}
	binData, bitEnd, err = PortNumber(value.PortNumber, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EndpointIPAddressAndPort: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EndpointIPAddressAndPort: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EPSTAI (value ngapType.EPSTAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAI: " + err.Error())
	}
	binData, bitEnd, err = EPSTAC(value.EPSTAC, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEPSTAIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EPSTAI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ERABInformationItem (value ngapType.ERABInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DLForwarding != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = ERABID(value.ERABID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ERABInformationItem: " + err.Error())
	}
	if value.DLForwarding != nil {
		binData, bitEnd, err = DLForwarding(*value.DLForwarding, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ERABInformationItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerERABInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ERABInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EUTRACGI (value ngapType.EUTRACGI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGI: " + err.Error())
	}
	binData, bitEnd, err = EUTRACellIdentity(value.EUTRACellIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerEUTRACGIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EUTRACGI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedUEActivityBehaviour (value ngapType.ExpectedUEActivityBehaviour, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.SourceOfUEActivityBehaviourInformation != nil {
		option += 1<<1
	}
	if value.ExpectedIdlePeriod != nil {
		option += 1<<2
	}
	if value.ExpectedActivityPeriod != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	if value.ExpectedActivityPeriod != nil {
		binData, bitEnd, err = ExpectedActivityPeriod(*value.ExpectedActivityPeriod, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEActivityBehaviour: " + err.Error())
		}
	}
	if value.ExpectedIdlePeriod != nil {
		binData, bitEnd, err = ExpectedIdlePeriod(*value.ExpectedIdlePeriod, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEActivityBehaviour: " + err.Error())
		}
	}
	if value.SourceOfUEActivityBehaviourInformation != nil {
		binData, bitEnd, err = SourceOfUEActivityBehaviourInformation(*value.SourceOfUEActivityBehaviourInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEActivityBehaviour: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEActivityBehaviour: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedUEBehaviour (value ngapType.ExpectedUEBehaviour, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ExpectedUEMovingTrajectory != nil {
		option += 1<<1
	}
	if value.ExpectedUEMobility != nil {
		option += 1<<2
	}
	if value.ExpectedHOInterval != nil {
		option += 1<<3
	}
	if value.ExpectedUEActivityBehaviour != nil {
		option += 1<<4
	}
	binData, bitEnd = EncodeOption(option, 6, binData, bitEnd)
	if value.ExpectedUEActivityBehaviour != nil {
		binData, bitEnd, err = ExpectedUEActivityBehaviour(*value.ExpectedUEActivityBehaviour, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEBehaviour: " + err.Error())
		}
	}
	if value.ExpectedHOInterval != nil {
		binData, bitEnd, err = ExpectedHOInterval(*value.ExpectedHOInterval, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEBehaviour: " + err.Error())
		}
	}
	if value.ExpectedUEMobility != nil {
		binData, bitEnd, err = ExpectedUEMobility(*value.ExpectedUEMobility, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEBehaviour: " + err.Error())
		}
	}
	if value.ExpectedUEMovingTrajectory != nil {
		binData, bitEnd, err = ExpectedUEMovingTrajectory(*value.ExpectedUEMovingTrajectory, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEBehaviour: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerExpectedUEBehaviourExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEBehaviour: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedUEMovingTrajectoryItem (value ngapType.ExpectedUEMovingTrajectoryItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TimeStayedInCell != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = NGRANCGI(value.NGRANCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItem: " + err.Error())
	}
	if *value.TimeStayedInCell < 0 || *value.TimeStayedInCell > 4095 {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItem: TimeStayedInCell: Integer: valueMin:0,valueMax:4095")
	}
	binData, bitEnd = EncodeInteger(*value.TimeStayedInCell, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExtendedRATRestrictionInformation (value ngapType.ExtendedRATRestrictionInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if value.PrimaryRATRestriction.BitLength != 8 || len(value.PrimaryRATRestriction.Bytes) != 1 {
		return binData, bitEnd, errors.New("ExtendedRATRestrictionInformation: PrimaryRATRestriction: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.PrimaryRATRestriction, binData, bitEnd)
	if value.SecondaryRATRestriction.BitLength != 8 || len(value.SecondaryRATRestriction.Bytes) != 1 {
		return binData, bitEnd, errors.New("ExtendedRATRestrictionInformation: SecondaryRATRestriction: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.SecondaryRATRestriction, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExtendedRATRestrictionInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func FiveGSTMSI (value ngapType.FiveGSTMSI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = AMFSetID(value.AMFSetID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSI: " + err.Error())
	}
	binData, bitEnd, err = AMFPointer(value.AMFPointer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSI: " + err.Error())
	}
	binData, bitEnd, err = FiveGTMSI(value.FiveGTMSI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerFiveGSTMSIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("FiveGSTMSI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ForbiddenAreaInformationItem (value ngapType.ForbiddenAreaInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItem: " + err.Error())
	}
	binData, bitEnd, err = ForbiddenTACs(value.ForbiddenTACs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ForbiddenAreaInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func GBRQosInformation (value ngapType.GBRQosInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.MaximumPacketLossRateUL != nil {
		option += 1<<1
	}
	if value.MaximumPacketLossRateDL != nil {
		option += 1<<2
	}
	if value.NotificationControl != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	binData, bitEnd, err = BitRate(value.MaximumFlowBitRateDL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
	}
	binData, bitEnd, err = BitRate(value.MaximumFlowBitRateUL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
	}
	binData, bitEnd, err = BitRate(value.GuaranteedFlowBitRateDL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
	}
	binData, bitEnd, err = BitRate(value.GuaranteedFlowBitRateUL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
	}
	if value.NotificationControl != nil {
		binData, bitEnd, err = NotificationControl(*value.NotificationControl, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
		}
	}
	if value.MaximumPacketLossRateDL != nil {
		binData, bitEnd, err = PacketLossRate(*value.MaximumPacketLossRateDL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
		}
	}
	if value.MaximumPacketLossRateUL != nil {
		binData, bitEnd, err = PacketLossRate(*value.MaximumPacketLossRateUL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerGBRQosInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GBRQosInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func GlobalGNBID (value ngapType.GlobalGNBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBID: " + err.Error())
	}
	binData, bitEnd, err = GNBID(value.GNBID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBID: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerGlobalGNBIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalGNBID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func GlobalN3IWFID (value ngapType.GlobalN3IWFID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFID: " + err.Error())
	}
	binData, bitEnd, err = N3IWFID(value.N3IWFID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFID: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerGlobalN3IWFIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalN3IWFID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func GlobalNgENBID (value ngapType.GlobalNgENBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBID: " + err.Error())
	}
	binData, bitEnd, err = NgENBID(value.NgENBID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBID: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerGlobalNgENBIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalNgENBID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func GTPTunnel (value ngapType.GTPTunnel, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TransportLayerAddress(value.TransportLayerAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnel: " + err.Error())
	}
	binData, bitEnd, err = GTPTEID(value.GTPTEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnel: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerGTPTunnelExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GTPTunnel: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func GUAMI (value ngapType.GUAMI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMI: " + err.Error())
	}
	binData, bitEnd, err = AMFRegionID(value.AMFRegionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMI: " + err.Error())
	}
	binData, bitEnd, err = AMFSetID(value.AMFSetID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMI: " + err.Error())
	}
	binData, bitEnd, err = AMFPointer(value.AMFPointer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerGUAMIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GUAMI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func HandoverCommandTransfer (value ngapType.HandoverCommandTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DataForwardingResponseDRBList != nil {
		option += 1<<1
	}
	if value.QosFlowToBeForwardedList != nil {
		option += 1<<2
	}
	if value.DLForwardingUPTNLInformation != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	if value.DLForwardingUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.DLForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransfer: " + err.Error())
		}
	}
	if value.QosFlowToBeForwardedList != nil {
		binData, bitEnd, err = QosFlowToBeForwardedList(*value.QosFlowToBeForwardedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransfer: " + err.Error())
		}
	}
	if value.DataForwardingResponseDRBList != nil {
		binData, bitEnd, err = DataForwardingResponseDRBList(*value.DataForwardingResponseDRBList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerHandoverCommandTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func HandoverPreparationUnsuccessfulTransfer (value ngapType.HandoverPreparationUnsuccessfulTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationUnsuccessfulTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationUnsuccessfulTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func HandoverRequestAcknowledgeTransfer (value ngapType.HandoverRequestAcknowledgeTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DataForwardingResponseDRBList != nil {
		option += 1<<1
	}
	if value.QosFlowFailedToSetupList != nil {
		option += 1<<2
	}
	if value.SecurityResult != nil {
		option += 1<<3
	}
	if value.DLForwardingUPTNLInformation != nil {
		option += 1<<4
	}
	binData, bitEnd = EncodeOption(option, 6, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.DLNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
	}
	if value.DLForwardingUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.DLForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	if value.SecurityResult != nil {
		binData, bitEnd, err = SecurityResult(*value.SecurityResult, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	binData, bitEnd, err = QosFlowListWithDataForwarding(value.QosFlowSetupResponseList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
	}
	if value.QosFlowFailedToSetupList != nil {
		binData, bitEnd, err = QosFlowListWithCause(*value.QosFlowFailedToSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	if value.DataForwardingResponseDRBList != nil {
		binData, bitEnd, err = DataForwardingResponseDRBList(*value.DataForwardingResponseDRBList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func HandoverRequiredTransfer (value ngapType.HandoverRequiredTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DirectForwardingPathAvailability != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	if value.DirectForwardingPathAvailability != nil {
		binData, bitEnd, err = DirectForwardingPathAvailability(*value.DirectForwardingPathAvailability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerHandoverRequiredTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func HandoverResourceAllocationUnsuccessfulTransfer (value ngapType.HandoverResourceAllocationUnsuccessfulTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.CriticalityDiagnostics != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransfer: " + err.Error())
	}
	if value.CriticalityDiagnostics != nil {
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func InfoOnRecommendedCellsAndRANNodesForPaging (value ngapType.InfoOnRecommendedCellsAndRANNodesForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = RecommendedCellsForPaging(value.RecommendedCellsForPaging, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPaging: " + err.Error())
	}
	binData, bitEnd, err = RecommendedRANNodesForPaging(value.RecommendRANNodesForPaging, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPaging: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPaging: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func LAI (value ngapType.LAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNidentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LAI: " + err.Error())
	}
	binData, bitEnd, err = LAC(value.LAC, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LAI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerLAIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LAI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func LastVisitedCellItem (value ngapType.LastVisitedCellItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = LastVisitedCellInformation(value.LastVisitedCellInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerLastVisitedCellItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedCellItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func LastVisitedNGRANCellInformation (value ngapType.LastVisitedNGRANCellInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.HOCauseValue != nil {
		option += 1<<1
	}
	if value.TimeUEStayedInCellEnhancedGranularity != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = NGRANCGI(value.GlobalCellID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformation: " + err.Error())
	}
	binData, bitEnd, err = CellType(value.CellType, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformation: " + err.Error())
	}
	binData, bitEnd, err = TimeUEStayedInCell(value.TimeUEStayedInCell, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformation: " + err.Error())
	}
	if value.TimeUEStayedInCellEnhancedGranularity != nil {
		binData, bitEnd, err = TimeUEStayedInCellEnhancedGranularity(*value.TimeUEStayedInCellEnhancedGranularity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedNGRANCellInformation: " + err.Error())
		}
	}
	if value.HOCauseValue != nil {
		binData, bitEnd, err = Cause(*value.HOCauseValue, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedNGRANCellInformation: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedNGRANCellInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func LocationReportingRequestType (value ngapType.LocationReportingRequestType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.LocationReportingReferenceIDToBeCancelled != nil {
		option += 1<<1
	}
	if value.AreaOfInterestList != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = EventType(value.EventType, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestType: " + err.Error())
	}
	binData, bitEnd, err = ReportArea(value.ReportArea, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestType: " + err.Error())
	}
	if value.AreaOfInterestList != nil {
		binData, bitEnd, err = AreaOfInterestList(*value.AreaOfInterestList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingRequestType: " + err.Error())
		}
	}
	if value.LocationReportingReferenceIDToBeCancelled != nil {
		binData, bitEnd, err = LocationReportingReferenceID(*value.LocationReportingReferenceIDToBeCancelled, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingRequestType: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerLocationReportingRequestTypeExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingRequestType: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func MobilityRestrictionList (value ngapType.MobilityRestrictionList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ServiceAreaInformation != nil {
		option += 1<<1
	}
	if value.ForbiddenAreaInformation != nil {
		option += 1<<2
	}
	if value.RATRestrictions != nil {
		option += 1<<3
	}
	if value.EquivalentPLMNs != nil {
		option += 1<<4
	}
	binData, bitEnd = EncodeOption(option, 6, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.ServingPLMN, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionList: " + err.Error())
	}
	if value.EquivalentPLMNs != nil {
		binData, bitEnd, err = EquivalentPLMNs(*value.EquivalentPLMNs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionList: " + err.Error())
		}
	}
	if value.RATRestrictions != nil {
		binData, bitEnd, err = RATRestrictions(*value.RATRestrictions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionList: " + err.Error())
		}
	}
	if value.ForbiddenAreaInformation != nil {
		binData, bitEnd, err = ForbiddenAreaInformation(*value.ForbiddenAreaInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionList: " + err.Error())
		}
	}
	if value.ServiceAreaInformation != nil {
		binData, bitEnd, err = ServiceAreaInformation(*value.ServiceAreaInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionList: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerMobilityRestrictionListExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionList: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NGRANTNLAssociationToRemoveItem (value ngapType.NGRANTNLAssociationToRemoveItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.TNLAssociationTransportLayerAddressAMF != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = CPTransportLayerInformation(value.TNLAssociationTransportLayerAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItem: " + err.Error())
	}
	if value.TNLAssociationTransportLayerAddressAMF != nil {
		binData, bitEnd, err = CPTransportLayerInformation(*value.TNLAssociationTransportLayerAddressAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NonDynamic5QIDescriptor (value ngapType.NonDynamic5QIDescriptor, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.MaximumDataBurstVolume != nil {
		option += 1<<1
	}
	if value.AveragingWindow != nil {
		option += 1<<2
	}
	if value.PriorityLevelQos != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	binData, bitEnd, err = FiveQI(value.FiveQI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptor: " + err.Error())
	}
	if value.PriorityLevelQos != nil {
		binData, bitEnd, err = PriorityLevelQos(*value.PriorityLevelQos, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NonDynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.AveragingWindow != nil {
		binData, bitEnd, err = AveragingWindow(*value.AveragingWindow, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NonDynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.MaximumDataBurstVolume != nil {
		binData, bitEnd, err = MaximumDataBurstVolume(*value.MaximumDataBurstVolume, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NonDynamic5QIDescriptor: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NonDynamic5QIDescriptor: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NRCGI (value ngapType.NRCGI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGI: " + err.Error())
	}
	binData, bitEnd, err = NRCellIdentity(value.NRCellIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerNRCGIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NRCGI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func OverloadStartNSSAIItem (value ngapType.OverloadStartNSSAIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.SliceTrafficLoadReductionIndication != nil {
		option += 1<<1
	}
	if value.SliceOverloadResponse != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = SliceOverloadList(value.SliceOverloadList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItem: " + err.Error())
	}
	if value.SliceOverloadResponse != nil {
		binData, bitEnd, err = OverloadResponse(*value.SliceOverloadResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartNSSAIItem: " + err.Error())
		}
	}
	if value.SliceTrafficLoadReductionIndication != nil {
		binData, bitEnd, err = TrafficLoadReductionIndication(*value.SliceTrafficLoadReductionIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartNSSAIItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartNSSAIItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PacketErrorRate (value ngapType.PacketErrorRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if value.PERScalar < 0 || value.PERScalar > 9 {
		return binData, bitEnd, errors.New("PacketErrorRate: PERScalar: Integer: valueMin:0,valueMax:9")
	}
	binData, bitEnd = EncodeIntegerExt(value.PERScalar, 0, binData, bitEnd)
	if value.PERExponent < 0 || value.PERExponent > 9 {
		return binData, bitEnd, errors.New("PacketErrorRate: PERExponent: Integer: valueMin:0,valueMax:9")
	}
	binData, bitEnd = EncodeIntegerExt(value.PERExponent, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPacketErrorRateExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PacketErrorRate: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PagingAttemptInformation (value ngapType.PagingAttemptInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NextPagingAreaScope != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = PagingAttemptCount(value.PagingAttemptCount, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformation: " + err.Error())
	}
	binData, bitEnd, err = IntendedNumberOfPagingAttempts(value.IntendedNumberOfPagingAttempts, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformation: " + err.Error())
	}
	if value.NextPagingAreaScope != nil {
		binData, bitEnd, err = NextPagingAreaScope(*value.NextPagingAreaScope, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingAttemptInformation: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPagingAttemptInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingAttemptInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PathSwitchRequestAcknowledgeTransfer (value ngapType.PathSwitchRequestAcknowledgeTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.SecurityIndication != nil {
		option += 1<<1
	}
	if value.ULNGUUPTNLInformation != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.ULNGUUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.ULNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	if value.SecurityIndication != nil {
		binData, bitEnd, err = SecurityIndication(*value.SecurityIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PathSwitchRequestSetupFailedTransfer (value ngapType.PathSwitchRequestSetupFailedTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestSetupFailedTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestSetupFailedTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PathSwitchRequestTransfer (value ngapType.PathSwitchRequestTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.UserPlaneSecurityInformation != nil {
		option += 1<<1
	}
	if value.DLNGUTNLInformationReused != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.DLNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestTransfer: " + err.Error())
	}
	if value.DLNGUTNLInformationReused != nil {
		binData, bitEnd, err = DLNGUTNLInformationReused(*value.DLNGUTNLInformationReused, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestTransfer: " + err.Error())
		}
	}
	if value.UserPlaneSecurityInformation != nil {
		binData, bitEnd, err = UserPlaneSecurityInformation(*value.UserPlaneSecurityInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestTransfer: " + err.Error())
		}
	}
	binData, bitEnd, err = QosFlowAcceptedList(value.QosFlowAcceptedList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPathSwitchRequestTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PathSwitchRequestUnsuccessfulTransfer (value ngapType.PathSwitchRequestUnsuccessfulTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestUnsuccessfulTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestUnsuccessfulTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionAggregateMaximumBitRate (value ngapType.PDUSessionAggregateMaximumBitRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = BitRate(value.PDUSessionAggregateMaximumBitRateDL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRate: " + err.Error())
	}
	binData, bitEnd, err = BitRate(value.PDUSessionAggregateMaximumBitRateUL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRate: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRate: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceAdmittedItem (value ngapType.PDUSessionResourceAdmittedItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItem: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.HandoverRequestAcknowledgeTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToModifyItemModCfm (value ngapType.PDUSessionResourceFailedToModifyItemModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfm: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceModifyIndicationUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfm: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfm: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToModifyItemModRes (value ngapType.PDUSessionResourceFailedToModifyItemModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModRes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceModifyUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModRes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemCxtFail (value ngapType.PDUSessionResourceFailedToSetupItemCxtFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFail: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFail: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFail: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemCxtRes (value ngapType.PDUSessionResourceFailedToSetupItemCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtRes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtRes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemHOAck (value ngapType.PDUSessionResourceFailedToSetupItemHOAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAck: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.HandoverResourceAllocationUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAck: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAck: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemPSReq (value ngapType.PDUSessionResourceFailedToSetupItemPSReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReq: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PathSwitchRequestSetupFailedTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReq: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemSURes (value ngapType.PDUSessionResourceFailedToSetupItemSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSURes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSURes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSURes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceHandoverItem (value ngapType.PDUSessionResourceHandoverItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItem: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.HandoverCommandTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceHandoverItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceInformationItem (value ngapType.PDUSessionResourceInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DRBsToQosFlowsMappingList != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceInformationItem: " + err.Error())
	}
	binData, bitEnd, err = QosFlowInformationList(value.QosFlowInformationList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceInformationItem: " + err.Error())
	}
	if value.DRBsToQosFlowsMappingList != nil {
		binData, bitEnd, err = DRBsToQosFlowsMappingList(*value.DRBsToQosFlowsMappingList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceInformationItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceItemCxtRelCpl (value ngapType.PDUSessionResourceItemCxtRelCpl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCpl: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCpl: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceItemCxtRelReq (value ngapType.PDUSessionResourceItemCxtRelReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReq: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceItemHORqd (value ngapType.PDUSessionResourceItemHORqd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqd: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.HandoverRequiredTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqd: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceItemHORqd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyConfirmTransfer (value ngapType.PDUSessionResourceModifyConfirmTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.QosFlowFailedToModifyList != nil {
		option += 1<<1
	}
	if value.AdditionalNGUUPTNLInformation != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = QosFlowModifyConfirmList(value.QosFlowModifyConfirmList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransfer: " + err.Error())
	}
	binData, bitEnd, err = UPTransportLayerInformation(value.ULNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransfer: " + err.Error())
	}
	if value.AdditionalNGUUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformationPairList(*value.AdditionalNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransfer: " + err.Error())
		}
	}
	if value.QosFlowFailedToModifyList != nil {
		binData, bitEnd, err = QosFlowListWithCause(*value.QosFlowFailedToModifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndicationUnsuccessfulTransfer (value ngapType.PDUSessionResourceModifyIndicationUnsuccessfulTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationUnsuccessfulTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationUnsuccessfulTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyRequestTransfer (value ngapType.PDUSessionResourceModifyRequestTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeOption(option, 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyResponseTransfer (value ngapType.PDUSessionResourceModifyResponseTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.QosFlowFailedToAddOrModifyList != nil {
		option += 1<<1
	}
	if value.AdditionalDLQosFlowPerTNLInformation != nil {
		option += 1<<2
	}
	if value.QosFlowAddOrModifyResponseList != nil {
		option += 1<<3
	}
	if value.ULNGUUPTNLInformation != nil {
		option += 1<<4
	}
	if value.DLNGUUPTNLInformation != nil {
		option += 1<<5
	}
	binData, bitEnd = EncodeOption(option, 7, binData, bitEnd)
	if value.DLNGUUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.DLNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransfer: " + err.Error())
		}
	}
	if value.ULNGUUPTNLInformation != nil {
		binData, bitEnd, err = UPTransportLayerInformation(*value.ULNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransfer: " + err.Error())
		}
	}
	if value.QosFlowAddOrModifyResponseList != nil {
		binData, bitEnd, err = QosFlowAddOrModifyResponseList(*value.QosFlowAddOrModifyResponseList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransfer: " + err.Error())
		}
	}
	if value.AdditionalDLQosFlowPerTNLInformation != nil {
		binData, bitEnd, err = QosFlowPerTNLInformationList(*value.AdditionalDLQosFlowPerTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransfer: " + err.Error())
		}
	}
	if value.QosFlowFailedToAddOrModifyList != nil {
		binData, bitEnd, err = QosFlowListWithCause(*value.QosFlowFailedToAddOrModifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndicationTransfer (value ngapType.PDUSessionResourceModifyIndicationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.AdditionalDLQosFlowPerTNLInformation != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = QosFlowPerTNLInformation(value.DLQosFlowPerTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransfer: " + err.Error())
	}
	if value.AdditionalDLQosFlowPerTNLInformation != nil {
		binData, bitEnd, err = QosFlowPerTNLInformationList(*value.AdditionalDLQosFlowPerTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModCfm (value ngapType.PDUSessionResourceModifyItemModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfm: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceModifyConfirmTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfm: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfm: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModInd (value ngapType.PDUSessionResourceModifyItemModInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModInd: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceModifyIndicationTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModInd: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModInd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModReq (value ngapType.PDUSessionResourceModifyItemModReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NASPDU != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReq: " + err.Error())
	}
	if value.NASPDU != nil {
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReq: " + err.Error())
		}
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceModifyRequestTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReq: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModRes (value ngapType.PDUSessionResourceModifyItemModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModRes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceModifyResponseTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModRes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyUnsuccessfulTransfer (value ngapType.PDUSessionResourceModifyUnsuccessfulTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.CriticalityDiagnostics != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransfer: " + err.Error())
	}
	if value.CriticalityDiagnostics != nil {
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyItem (value ngapType.PDUSessionResourceNotifyItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItem: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceNotifyTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyReleasedTransfer (value ngapType.PDUSessionResourceNotifyReleasedTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyTransfer (value ngapType.PDUSessionResourceNotifyTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.QosFlowReleasedList != nil {
		option += 1<<1
	}
	if value.QosFlowNotifyList != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.QosFlowNotifyList != nil {
		binData, bitEnd, err = QosFlowNotifyList(*value.QosFlowNotifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransfer: " + err.Error())
		}
	}
	if value.QosFlowReleasedList != nil {
		binData, bitEnd, err = QosFlowListWithCause(*value.QosFlowReleasedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseCommandTransfer (value ngapType.PDUSessionResourceReleaseCommandTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemNot (value ngapType.PDUSessionResourceReleasedItemNot, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNot: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceNotifyReleasedTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNot: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNot: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemPSAck (value ngapType.PDUSessionResourceReleasedItemPSAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAck: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PathSwitchRequestUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAck: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAck: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemPSFail (value ngapType.PDUSessionResourceReleasedItemPSFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFail: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PathSwitchRequestUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFail: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFail: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemRelRes (value ngapType.PDUSessionResourceReleasedItemRelRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelRes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceReleaseResponseTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelRes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseResponseTransfer (value ngapType.PDUSessionResourceReleaseResponseTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSecondaryRATUsageItem (value ngapType.PDUSessionResourceSecondaryRATUsageItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItem: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.SecondaryRATDataUsageReportTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemCxtReq (value ngapType.PDUSessionResourceSetupItemCxtReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NASPDU != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReq: " + err.Error())
	}
	if value.NASPDU != nil {
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReq: " + err.Error())
		}
	}
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReq: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupRequestTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReq: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemCxtRes (value ngapType.PDUSessionResourceSetupItemCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtRes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupResponseTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtRes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemHOReq (value ngapType.PDUSessionResourceSetupItemHOReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
	}
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.HandoverRequestTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemSUReq (value ngapType.PDUSessionResourceSetupItemSUReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.PDUSessionNASPDU != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReq: " + err.Error())
	}
	if value.PDUSessionNASPDU != nil {
		binData, bitEnd, err = NASPDU(*value.PDUSessionNASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReq: " + err.Error())
		}
	}
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReq: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupRequestTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReq: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemSURes (value ngapType.PDUSessionResourceSetupItemSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSURes: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceSetupResponseTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSURes: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSURes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupRequestTransfer (value ngapType.PDUSessionResourceSetupRequestTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeOption(option, 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupResponseTransfer (value ngapType.PDUSessionResourceSetupResponseTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.QosFlowFailedToSetupList != nil {
		option += 1<<1
	}
	if value.SecurityResult != nil {
		option += 1<<2
	}
	if value.AdditionalDLQosFlowPerTNLInformation != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	binData, bitEnd, err = QosFlowPerTNLInformation(value.DLQosFlowPerTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransfer: " + err.Error())
	}
	if value.AdditionalDLQosFlowPerTNLInformation != nil {
		binData, bitEnd, err = QosFlowPerTNLInformationList(*value.AdditionalDLQosFlowPerTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransfer: " + err.Error())
		}
	}
	if value.SecurityResult != nil {
		binData, bitEnd, err = SecurityResult(*value.SecurityResult, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransfer: " + err.Error())
		}
	}
	if value.QosFlowFailedToSetupList != nil {
		binData, bitEnd, err = QosFlowListWithCause(*value.QosFlowFailedToSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupUnsuccessfulTransfer (value ngapType.PDUSessionResourceSetupUnsuccessfulTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.CriticalityDiagnostics != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransfer: " + err.Error())
	}
	if value.CriticalityDiagnostics != nil {
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSwitchedItem (value ngapType.PDUSessionResourceSwitchedItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItem: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PathSwitchRequestAcknowledgeTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToBeSwitchedDLItem (value ngapType.PDUSessionResourceToBeSwitchedDLItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItem: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PathSwitchRequestTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToReleaseItemHOCmd (value ngapType.PDUSessionResourceToReleaseItemHOCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmd: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.HandoverPreparationUnsuccessfulTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmd: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToReleaseItemRelCmd (value ngapType.PDUSessionResourceToReleaseItemRelCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmd: " + err.Error())
	}
	binData, bitEnd, err = EncodeOctetStringContaining(value.PDUSessionResourceReleaseCommandTransfer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmd: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionUsageReport (value ngapType.PDUSessionUsageReport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if value.RATType < 0 || value.RATType > 4 {
		return binData, bitEnd, errors.New("PDUSessionUsageReport: RATType: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.RATType, 0, binData, bitEnd)
	binData, bitEnd, err = VolumeTimedReportList(value.PDUSessionTimedReportList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionUsageReport: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionUsageReportExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionUsageReport: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PLMNSupportItem (value ngapType.PLMNSupportItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItem: " + err.Error())
	}
	binData, bitEnd, err = SliceSupportList(value.SliceSupportList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPLMNSupportItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PLMNSupportItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAcceptedItem (value ngapType.QosFlowAcceptedItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAcceptedItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowAcceptedItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAcceptedItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyRequestItem (value ngapType.QosFlowAddOrModifyRequestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ERABID != nil {
		option += 1<<1
	}
	if value.QosFlowLevelQosParameters != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItem: " + err.Error())
	}
	if value.QosFlowLevelQosParameters != nil {
		binData, bitEnd, err = QosFlowLevelQosParameters(*value.QosFlowLevelQosParameters, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItem: " + err.Error())
		}
	}
	if value.ERABID != nil {
		binData, bitEnd, err = ERABID(*value.ERABID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyResponseItem (value ngapType.QosFlowAddOrModifyResponseItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowInformationItem (value ngapType.QosFlowInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DLForwarding != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowInformationItem: " + err.Error())
	}
	if value.DLForwarding != nil {
		binData, bitEnd, err = DLForwarding(*value.DLForwarding, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowInformationItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowLevelQosParameters (value ngapType.QosFlowLevelQosParameters, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.AdditionalQosFlowInformation != nil {
		option += 1<<1
	}
	if value.ReflectiveQosAttribute != nil {
		option += 1<<2
	}
	if value.GBRQosInformation != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	binData, bitEnd, err = QosCharacteristics(value.QosCharacteristics, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParameters: " + err.Error())
	}
	binData, bitEnd, err = AllocationAndRetentionPriority(value.AllocationAndRetentionPriority, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParameters: " + err.Error())
	}
	if value.GBRQosInformation != nil {
		binData, bitEnd, err = GBRQosInformation(*value.GBRQosInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowLevelQosParameters: " + err.Error())
		}
	}
	if value.ReflectiveQosAttribute != nil {
		binData, bitEnd, err = ReflectiveQosAttribute(*value.ReflectiveQosAttribute, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowLevelQosParameters: " + err.Error())
		}
	}
	if value.AdditionalQosFlowInformation != nil {
		binData, bitEnd, err = AdditionalQosFlowInformation(*value.AdditionalQosFlowInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowLevelQosParameters: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowLevelQosParameters: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowWithCauseItem (value ngapType.QosFlowWithCauseItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowWithCauseItem: " + err.Error())
	}
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowWithCauseItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowWithCauseItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowWithCauseItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowModifyConfirmItem (value ngapType.QosFlowModifyConfirmItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowModifyConfirmItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowModifyConfirmItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowNotifyItem (value ngapType.QosFlowNotifyItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowNotifyItem: " + err.Error())
	}
	binData, bitEnd, err = NotificationCause(value.NotificationCause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowNotifyItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowNotifyItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowNotifyItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowPerTNLInformation (value ngapType.QosFlowPerTNLInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.UPTransportLayerInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformation: " + err.Error())
	}
	binData, bitEnd, err = AssociatedQosFlowList(value.AssociatedQosFlowList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformation: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowPerTNLInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowPerTNLInformationItem (value ngapType.QosFlowPerTNLInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowPerTNLInformation(value.QosFlowPerTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowPerTNLInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowSetupRequestItem (value ngapType.QosFlowSetupRequestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ERABID != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItem: " + err.Error())
	}
	binData, bitEnd, err = QosFlowLevelQosParameters(value.QosFlowLevelQosParameters, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItem: " + err.Error())
	}
	if value.ERABID != nil {
		binData, bitEnd, err = ERABID(*value.ERABID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowSetupRequestItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowSetupRequestItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowItemWithDataForwarding (value ngapType.QosFlowItemWithDataForwarding, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.DataForwardingAccepted != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemWithDataForwarding: " + err.Error())
	}
	if value.DataForwardingAccepted != nil {
		binData, bitEnd, err = DataForwardingAccepted(*value.DataForwardingAccepted, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowItemWithDataForwarding: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowItemWithDataForwarding: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowToBeForwardedItem (value ngapType.QosFlowToBeForwardedItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowToBeForwardedItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowToBeForwardedItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QoSFlowsUsageReportItem (value ngapType.QoSFlowsUsageReportItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItem: " + err.Error())
	}
	if value.RATType < 0 || value.RATType > 4 {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItem: RATType: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.RATType, 0, binData, bitEnd)
	binData, bitEnd, err = VolumeTimedReportList(value.QoSFlowsTimedReportList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QoSFlowsUsageReportItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RANStatusTransferTransparentContainer (value ngapType.RANStatusTransferTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = DRBsSubjectToStatusTransferList(value.DRBsSubjectToStatusTransferList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANStatusTransferTransparentContainer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RATRestrictionsItem (value ngapType.RATRestrictionsItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItem: " + err.Error())
	}
	binData, bitEnd, err = RATRestrictionInformation(value.RATRestrictionInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRATRestrictionsItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RATRestrictionsItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedCellsForPaging (value ngapType.RecommendedCellsForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = RecommendedCellList(value.RecommendedCellList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPaging: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRecommendedCellsForPagingExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedCellsForPaging: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedCellItem (value ngapType.RecommendedCellItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TimeStayedInCell != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = NGRANCGI(value.NGRANCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItem: " + err.Error())
	}
	if *value.TimeStayedInCell < 0 || *value.TimeStayedInCell > 4095 {
		return binData, bitEnd, errors.New("RecommendedCellItem: TimeStayedInCell: Integer: valueMin:0,valueMax:4095")
	}
	binData, bitEnd = EncodeInteger(*value.TimeStayedInCell, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRecommendedCellItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedCellItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedRANNodesForPaging (value ngapType.RecommendedRANNodesForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = RecommendedRANNodeList(value.RecommendedRANNodeList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPaging: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedRANNodesForPaging: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedRANNodeItem (value ngapType.RecommendedRANNodeItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = AMFPagingTarget(value.AMFPagingTarget, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRecommendedRANNodeItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedRANNodeItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RIMInformationTransfer (value ngapType.RIMInformationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TargetRANNodeID(value.TargetRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformationTransfer: " + err.Error())
	}
	binData, bitEnd, err = SourceRANNodeID(value.SourceRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformationTransfer: " + err.Error())
	}
	binData, bitEnd, err = RIMInformation(value.RIMInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformationTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRIMInformationTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RIMInformationTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RIMInformation (value ngapType.RIMInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeOption(option, 1, binData, bitEnd)
	binData, bitEnd, err = GNBSetID(value.TargetgNBSetID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformation: " + err.Error())
	}
	if value.RIMRSDetection < 0 || value.RIMRSDetection > 2 {
		return binData, bitEnd, errors.New("RIMInformation: RIMRSDetection: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.RIMRSDetection, 0, binData, bitEnd)
	return binData, bitEnd, err
}

func SecondaryRATUsageInformation (value ngapType.SecondaryRATUsageInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.QosFlowsUsageReportList != nil {
		option += 1<<1
	}
	if value.PDUSessionUsageReport != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.PDUSessionUsageReport != nil {
		binData, bitEnd, err = PDUSessionUsageReport(*value.PDUSessionUsageReport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATUsageInformation: " + err.Error())
		}
	}
	if value.QosFlowsUsageReportList != nil {
		binData, bitEnd, err = QoSFlowsUsageReportList(*value.QosFlowsUsageReportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATUsageInformation: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATUsageInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SecondaryRATDataUsageReportTransfer (value ngapType.SecondaryRATDataUsageReportTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.SecondaryRATUsageInformation != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	if value.SecondaryRATUsageInformation != nil {
		binData, bitEnd, err = SecondaryRATUsageInformation(*value.SecondaryRATUsageInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SecurityContext (value ngapType.SecurityContext, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NextHopChainingCount(value.NextHopChainingCount, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContext: " + err.Error())
	}
	binData, bitEnd, err = SecurityKey(value.NextHopNH, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContext: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSecurityContextExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecurityContext: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SecurityIndication (value ngapType.SecurityIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.MaximumIntegrityProtectedDataRateUL != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = IntegrityProtectionIndication(value.IntegrityProtectionIndication, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndication: " + err.Error())
	}
	binData, bitEnd, err = ConfidentialityProtectionIndication(value.ConfidentialityProtectionIndication, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndication: " + err.Error())
	}
	if value.MaximumIntegrityProtectedDataRateUL != nil {
		binData, bitEnd, err = MaximumIntegrityProtectedDataRate(*value.MaximumIntegrityProtectedDataRateUL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecurityIndication: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSecurityIndicationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecurityIndication: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SecurityResult (value ngapType.SecurityResult, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = IntegrityProtectionResult(value.IntegrityProtectionResult, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityResult: " + err.Error())
	}
	binData, bitEnd, err = ConfidentialityProtectionResult(value.ConfidentialityProtectionResult, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityResult: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSecurityResultExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecurityResult: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ServedGUAMIItem (value ngapType.ServedGUAMIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.BackupAMFName != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = GUAMI(value.GUAMI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItem: " + err.Error())
	}
	if value.BackupAMFName != nil {
		binData, bitEnd, err = AMFName(*value.BackupAMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServedGUAMIItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerServedGUAMIItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServedGUAMIItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ServiceAreaInformationItem (value ngapType.ServiceAreaInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NotAllowedTACs != nil {
		option += 1<<1
	}
	if value.AllowedTACs != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItem: " + err.Error())
	}
	if value.AllowedTACs != nil {
		binData, bitEnd, err = AllowedTACs(*value.AllowedTACs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServiceAreaInformationItem: " + err.Error())
		}
	}
	if value.NotAllowedTACs != nil {
		binData, bitEnd, err = NotAllowedTACs(*value.NotAllowedTACs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServiceAreaInformationItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerServiceAreaInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServiceAreaInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SliceOverloadItem (value ngapType.SliceOverloadItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSliceOverloadItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SliceOverloadItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SliceSupportItem (value ngapType.SliceSupportItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSliceSupportItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SliceSupportItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SNSSAI (value ngapType.SNSSAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.SD != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = SST(value.SST, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAI: " + err.Error())
	}
	if value.SD != nil {
		binData, bitEnd, err = SD(*value.SD, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SNSSAI: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSNSSAIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SNSSAI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SONConfigurationTransfer (value ngapType.SONConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.XnTNLConfigurationInfo != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = TargetRANNodeID(value.TargetRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
	}
	binData, bitEnd, err = SourceRANNodeID(value.SourceRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
	}
	binData, bitEnd, err = SONInformation(value.SONInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
	}
	if value.XnTNLConfigurationInfo != nil {
		binData, bitEnd, err = XnTNLConfigurationInfo(*value.XnTNLConfigurationInfo, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSONConfigurationTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SONInformationReply (value ngapType.SONInformationReply, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.XnTNLConfigurationInfo != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	if value.XnTNLConfigurationInfo != nil {
		binData, bitEnd, err = XnTNLConfigurationInfo(*value.XnTNLConfigurationInfo, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformationReply: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSONInformationReplyExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformationReply: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SourceNGRANNodeToTargetNGRANNodeTransparentContainer (value ngapType.SourceNGRANNodeToTargetNGRANNodeTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.IndexToRFSP != nil {
		option += 1<<1
	}
	if value.ERABInformationList != nil {
		option += 1<<2
	}
	if value.PDUSessionResourceInformationList != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	binData, bitEnd, err = RRCContainer(value.RRCContainer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
	}
	if value.PDUSessionResourceInformationList != nil {
		binData, bitEnd, err = PDUSessionResourceInformationList(*value.PDUSessionResourceInformationList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
		}
	}
	if value.ERABInformationList != nil {
		binData, bitEnd, err = ERABInformationList(*value.ERABInformationList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
		}
	}
	binData, bitEnd, err = NGRANCGI(value.TargetCellID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
	}
	if value.IndexToRFSP != nil {
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
		}
	}
	binData, bitEnd, err = UEHistoryInformation(value.UEHistoryInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SourceRANNodeID (value ngapType.SourceRANNodeID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = GlobalRANNodeID(value.GlobalRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeID: " + err.Error())
	}
	binData, bitEnd, err = TAI(value.SelectedTAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeID: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSourceRANNodeIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceRANNodeID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SourceToTargetAMFInformationReroute (value ngapType.SourceToTargetAMFInformationReroute, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.RejectedNSSAIinTA != nil {
		option += 1<<1
	}
	if value.RejectedNSSAIinPLMN != nil {
		option += 1<<2
	}
	if value.ConfiguredNSSAI != nil {
		option += 1<<3
	}
	binData, bitEnd = EncodeOption(option, 5, binData, bitEnd)
	if value.ConfiguredNSSAI != nil {
		binData, bitEnd, err = ConfiguredNSSAI(*value.ConfiguredNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceToTargetAMFInformationReroute: " + err.Error())
		}
	}
	if value.RejectedNSSAIinPLMN != nil {
		binData, bitEnd, err = RejectedNSSAIinPLMN(*value.RejectedNSSAIinPLMN, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceToTargetAMFInformationReroute: " + err.Error())
		}
	}
	if value.RejectedNSSAIinTA != nil {
		binData, bitEnd, err = RejectedNSSAIinTA(*value.RejectedNSSAIinTA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceToTargetAMFInformationReroute: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceToTargetAMFInformationReroute: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SupportedTAItem (value ngapType.SupportedTAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAC(value.TAC, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItem: " + err.Error())
	}
	binData, bitEnd, err = BroadcastPLMNList(value.BroadcastPLMNList, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSupportedTAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SupportedTAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAI (value ngapType.TAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAI: " + err.Error())
	}
	binData, bitEnd, err = TAC(value.TAC, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAI: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAIExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAI: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastEUTRAItem (value ngapType.TAIBroadcastEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInTAIEUTRA(value.CompletedCellsInTAIEUTRA, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIBroadcastEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastNRItem (value ngapType.TAIBroadcastNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItem: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInTAINR(value.CompletedCellsInTAINR, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAIBroadcastNRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIBroadcastNRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledEUTRAItem (value ngapType.TAICancelledEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItem: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInTAIEUTRA(value.CancelledCellsInTAIEUTRA, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAICancelledEUTRAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledNRItem (value ngapType.TAICancelledNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItem: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInTAINR(value.CancelledCellsInTAINR, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAICancelledNRItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAICancelledNRItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForInactiveItem (value ngapType.TAIListForInactiveItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAIListForInactiveItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForInactiveItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForPagingItem (value ngapType.TAIListForPagingItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTAIListForPagingItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForPagingItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TargeteNBID (value ngapType.TargeteNBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = GlobalNgENBID(value.GlobalENBID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBID: " + err.Error())
	}
	binData, bitEnd, err = EPSTAI(value.SelectedEPSTAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBID: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTargeteNBIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargeteNBID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TargetNGRANNodeToSourceNGRANNodeTransparentContainer (value ngapType.TargetNGRANNodeToSourceNGRANNodeTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = RRCContainer(value.RRCContainer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetNGRANNodeToSourceNGRANNodeTransparentContainer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetNGRANNodeToSourceNGRANNodeTransparentContainer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TargetRANNodeID (value ngapType.TargetRANNodeID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = GlobalRANNodeID(value.GlobalRANNodeID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeID: " + err.Error())
	}
	binData, bitEnd, err = TAI(value.SelectedTAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeID: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTargetRANNodeIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetRANNodeID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TargetRNCID (value ngapType.TargetRNCID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ExtendedRNCID != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = LAI(value.LAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRNCID: " + err.Error())
	}
	binData, bitEnd, err = RNCID(value.RNCID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRNCID: " + err.Error())
	}
	if value.ExtendedRNCID != nil {
		binData, bitEnd, err = ExtendedRNCID(*value.ExtendedRNCID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetRNCID: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTargetRNCIDExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetRNCID: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TNLAssociationItem (value ngapType.TNLAssociationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = CPTransportLayerInformation(value.TNLAssociationAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItem: " + err.Error())
	}
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTNLAssociationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TNLAssociationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TraceActivation (value ngapType.TraceActivation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NGRANTraceID(value.NGRANTraceID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivation: " + err.Error())
	}
	binData, bitEnd, err = InterfacesToTrace(value.InterfacesToTrace, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivation: " + err.Error())
	}
	binData, bitEnd, err = TraceDepth(value.TraceDepth, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivation: " + err.Error())
	}
	binData, bitEnd, err = TransportLayerAddress(value.TraceCollectionEntityIPAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivation: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerTraceActivationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceActivation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UEAggregateMaximumBitRate (value ngapType.UEAggregateMaximumBitRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = BitRate(value.UEAggregateMaximumBitRateDL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRate: " + err.Error())
	}
	binData, bitEnd, err = BitRate(value.UEAggregateMaximumBitRateUL, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRate: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEAggregateMaximumBitRate: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UEAssociatedLogicalNGConnectionItem (value ngapType.UEAssociatedLogicalNGConnectionItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.RANUENGAPID != nil {
		option += 1<<1
	}
	if value.AMFUENGAPID != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.AMFUENGAPID != nil {
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItem: " + err.Error())
		}
	}
	if value.RANUENGAPID != nil {
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UENGAPIDPair (value ngapType.UENGAPIDPair, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = AMFUENGAPID(value.AMFUENGAPID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPair: " + err.Error())
	}
	binData, bitEnd, err = RANUENGAPID(value.RANUENGAPID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPair: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUENGAPIDPairExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UENGAPIDPair: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UEPresenceInAreaOfInterestItem (value ngapType.UEPresenceInAreaOfInterestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = LocationReportingReferenceID(value.LocationReportingReferenceID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItem: " + err.Error())
	}
	binData, bitEnd, err = UEPresence(value.UEPresence, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UERadioCapabilityForPaging (value ngapType.UERadioCapabilityForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.UERadioCapabilityForPagingOfEUTRA != nil {
		option += 1<<1
	}
	if value.UERadioCapabilityForPagingOfNR != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.UERadioCapabilityForPagingOfNR != nil {
		binData, bitEnd, err = UERadioCapabilityForPagingOfNR(*value.UERadioCapabilityForPagingOfNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityForPaging: " + err.Error())
		}
	}
	if value.UERadioCapabilityForPagingOfEUTRA != nil {
		binData, bitEnd, err = UERadioCapabilityForPagingOfEUTRA(*value.UERadioCapabilityForPagingOfEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityForPaging: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityForPaging: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UESecurityCapabilities (value ngapType.UESecurityCapabilities, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = NRencryptionAlgorithms(value.NRencryptionAlgorithms, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilities: " + err.Error())
	}
	binData, bitEnd, err = NRintegrityProtectionAlgorithms(value.NRintegrityProtectionAlgorithms, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilities: " + err.Error())
	}
	binData, bitEnd, err = EUTRAencryptionAlgorithms(value.EUTRAencryptionAlgorithms, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilities: " + err.Error())
	}
	binData, bitEnd, err = EUTRAintegrityProtectionAlgorithms(value.EUTRAintegrityProtectionAlgorithms, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilities: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUESecurityCapabilitiesExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UESecurityCapabilities: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ULNGUUPTNLModifyItem (value ngapType.ULNGUUPTNLModifyItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.ULNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItem: " + err.Error())
	}
	binData, bitEnd, err = UPTransportLayerInformation(value.DLNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ULNGUUPTNLModifyItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UnavailableGUAMIItem (value ngapType.UnavailableGUAMIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.BackupAMFName != nil {
		option += 1<<1
	}
	if value.TimerApproachForGUAMIRemoval != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	binData, bitEnd, err = GUAMI(value.GUAMI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItem: " + err.Error())
	}
	if value.TimerApproachForGUAMIRemoval != nil {
		binData, bitEnd, err = TimerApproachForGUAMIRemoval(*value.TimerApproachForGUAMIRemoval, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnavailableGUAMIItem: " + err.Error())
		}
	}
	if value.BackupAMFName != nil {
		binData, bitEnd, err = AMFName(*value.BackupAMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnavailableGUAMIItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUnavailableGUAMIItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnavailableGUAMIItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationItem (value ngapType.UPTransportLayerInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.NGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformationItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationPairItem (value ngapType.UPTransportLayerInformationPairItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = UPTransportLayerInformation(value.ULNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairItem: " + err.Error())
	}
	binData, bitEnd, err = UPTransportLayerInformation(value.DLNGUUPTNLInformation, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformationPairItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UserLocationInformationEUTRA (value ngapType.UserLocationInformationEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TimeStamp != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = EUTRACGI(value.EUTRACGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRA: " + err.Error())
	}
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRA: " + err.Error())
	}
	if value.TimeStamp != nil {
		binData, bitEnd, err = TimeStamp(*value.TimeStamp, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationEUTRA: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationEUTRA: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UserLocationInformationN3IWF (value ngapType.UserLocationInformationN3IWF, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = TransportLayerAddress(value.IPAddress, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWF: " + err.Error())
	}
	binData, bitEnd, err = PortNumber(value.PortNumber, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWF: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationN3IWF: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UserLocationInformationNR (value ngapType.UserLocationInformationNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TimeStamp != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = NRCGI(value.NRCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNR: " + err.Error())
	}
	binData, bitEnd, err = TAI(value.TAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNR: " + err.Error())
	}
	if value.TimeStamp != nil {
		binData, bitEnd, err = TimeStamp(*value.TimeStamp, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationNR: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUserLocationInformationNRExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationNR: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UserPlaneSecurityInformation (value ngapType.UserPlaneSecurityInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	binData, bitEnd, err = SecurityResult(value.SecurityResult, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserPlaneSecurityInformation: " + err.Error())
	}
	binData, bitEnd, err = SecurityIndication(value.SecurityIndication, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserPlaneSecurityInformation: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserPlaneSecurityInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func VolumeTimedReportItem (value ngapType.VolumeTimedReportItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeOption(option, 2, binData, bitEnd)
	if len(value.StartTimeStamp) != 4 {
		return binData, bitEnd, errors.New("VolumeTimedReportItem: StartTimeStamp: OctetString: sizeMin:4,sizeMax:4")
	}
	binData, bitEnd = EncodeOctetStringFix(value.StartTimeStamp, binData, bitEnd)
	if len(value.EndTimeStamp) != 4 {
		return binData, bitEnd, errors.New("VolumeTimedReportItem: EndTimeStamp: OctetString: sizeMin:4,sizeMax:4")
	}
	binData, bitEnd = EncodeOctetStringFix(value.EndTimeStamp, binData, bitEnd)
	if value.UsageCountUL < 0 || value.UsageCountUL > 18446744073709551615 {
		return binData, bitEnd, errors.New("VolumeTimedReportItem: UsageCountUL: Integer: valueMin:0,valueMax:18446744073709551615")
	}
	binData, bitEnd = EncodeInteger(value.UsageCountUL, 0, binData, bitEnd)
	if value.UsageCountDL < 0 || value.UsageCountDL > 18446744073709551615 {
		return binData, bitEnd, errors.New("VolumeTimedReportItem: UsageCountDL: Integer: valueMin:0,valueMax:18446744073709551615")
	}
	binData, bitEnd = EncodeInteger(value.UsageCountDL, 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerVolumeTimedReportItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("VolumeTimedReportItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnExtTLAItem (value ngapType.XnExtTLAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.GTPTLAs != nil {
		option += 1<<1
	}
	if value.IPsecTLA != nil {
		option += 1<<2
	}
	binData, bitEnd = EncodeOption(option, 4, binData, bitEnd)
	if value.IPsecTLA != nil {
		binData, bitEnd, err = TransportLayerAddress(*value.IPsecTLA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnExtTLAItem: " + err.Error())
		}
	}
	if value.GTPTLAs != nil {
		binData, bitEnd, err = XnGTPTLAs(*value.GTPTLAs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnExtTLAItem: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerXnExtTLAItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnExtTLAItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnTNLConfigurationInfo (value ngapType.XnTNLConfigurationInfo, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.XnExtendedTransportLayerAddresses != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeOption(option, 3, binData, bitEnd)
	binData, bitEnd, err = XnTLAs(value.XnTransportLayerAddresses, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfo: " + err.Error())
	}
	if value.XnExtendedTransportLayerAddresses != nil {
		binData, bitEnd, err = XnExtTLAs(*value.XnExtendedTransportLayerAddresses, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnTNLConfigurationInfo: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnTNLConfigurationInfo: " + err.Error())
		}
	}
	return binData, bitEnd, err
}
