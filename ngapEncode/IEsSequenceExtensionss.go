package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValue(value ngapType.AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AllocationAndRetentionPriorityExtIEsExtensionValue(value ngapType.AllocationAndRetentionPriorityExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AllowedNSSAIItemExtIEsExtensionValue(value ngapType.AllowedNSSAIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationSetupItemExtIEsExtensionValue(value ngapType.AMFTNLAssociationSetupItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToAddItemExtIEsExtensionValue(value ngapType.AMFTNLAssociationToAddItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToRemoveItemExtIEsExtensionValue(value ngapType.AMFTNLAssociationToRemoveItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFTNLAssociationToRemoveItemExtIEsExtensionValuePresentTNLAssociationTransportLayerAddressNGRAN:
		if value.TNLAssociationTransportLayerAddressNGRAN == nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEsExtensionValue: TNLAssociationTransportLayerAddressNGRAN: NIL")
		}
		binData, bitEnd, err = CPTransportLayerInformation(*value.TNLAssociationTransportLayerAddressNGRAN, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToUpdateItemExtIEsExtensionValue(value ngapType.AMFTNLAssociationToUpdateItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestExtIEsExtensionValue(value ngapType.AreaOfInterestExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestCellItemExtIEsExtensionValue(value ngapType.AreaOfInterestCellItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestItemExtIEsExtensionValue(value ngapType.AreaOfInterestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestRANNodeItemExtIEsExtensionValue(value ngapType.AreaOfInterestRANNodeItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestTAIItemExtIEsExtensionValue(value ngapType.AreaOfInterestTAIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AssistanceDataForPagingExtIEsExtensionValue(value ngapType.AssistanceDataForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AssistanceDataForRecommendedCellsExtIEsExtensionValue(value ngapType.AssistanceDataForRecommendedCellsExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AssociatedQosFlowItemExtIEsExtensionValue(value ngapType.AssociatedQosFlowItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AssociatedQosFlowItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func BroadcastPLMNItemExtIEsExtensionValue(value ngapType.BroadcastPLMNItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAIEUTRAItemExtIEsExtensionValue(value ngapType.CancelledCellsInEAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINRItemExtIEsExtensionValue(value ngapType.CancelledCellsInEAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAIEUTRAItemExtIEsExtensionValue(value ngapType.CancelledCellsInTAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINRItemExtIEsExtensionValue(value ngapType.CancelledCellsInTAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDBroadcastEUTRAItemExtIEsExtensionValue(value ngapType.CellIDBroadcastEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDBroadcastNRItemExtIEsExtensionValue(value ngapType.CellIDBroadcastNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDCancelledEUTRAItemExtIEsExtensionValue(value ngapType.CellIDCancelledEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDCancelledNRItemExtIEsExtensionValue(value ngapType.CellIDCancelledNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellTypeExtIEsExtensionValue(value ngapType.CellTypeExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellTypeExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CNAssistedRANTuningExtIEsExtensionValue(value ngapType.CNAssistedRANTuningExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CNAssistedRANTuningExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CNTypeRestrictionsForEquivalentItemExtIEsExtensionValue(value ngapType.CNTypeRestrictionsForEquivalentItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRAItemExtIEsExtensionValue(value ngapType.CompletedCellsInEAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINRItemExtIEsExtensionValue(value ngapType.CompletedCellsInEAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRAItemExtIEsExtensionValue(value ngapType.CompletedCellsInTAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINRItemExtIEsExtensionValue(value ngapType.CompletedCellsInTAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValue(value ngapType.CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN12ExtIEsExtensionValue(value ngapType.COUNTValueForPDCPSN12ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN18ExtIEsExtensionValue(value ngapType.COUNTValueForPDCPSN18ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CriticalityDiagnosticsExtIEsExtensionValue(value ngapType.CriticalityDiagnosticsExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CriticalityDiagnosticsIEItemExtIEsExtensionValue(value ngapType.CriticalityDiagnosticsIEItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DataForwardingResponseDRBItemExtIEsExtensionValue(value ngapType.DataForwardingResponseDRBItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DataForwardingResponseDRBItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DataForwardingResponseERABListItemExtIEsExtensionValue(value ngapType.DataForwardingResponseERABListItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DataForwardingResponseERABListItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBsSubjectToStatusTransferItemExtIEsExtensionValue(value ngapType.DRBsSubjectToStatusTransferItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DRBsSubjectToStatusTransferItemExtIEsExtensionValuePresentOldAssociatedQosFlowListULendmarkerexpected:
		if value.OldAssociatedQosFlowListULendmarkerexpected == nil {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEsExtensionValue: OldAssociatedQosFlowListULendmarkerexpected: NIL")
		}
		binData, bitEnd, err = AssociatedQosFlowList(*value.OldAssociatedQosFlowListULendmarkerexpected, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusDL12ExtIEsExtensionValue(value ngapType.DRBStatusDL12ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusDL18ExtIEsExtensionValue(value ngapType.DRBStatusDL18ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusUL12ExtIEsExtensionValue(value ngapType.DRBStatusUL12ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusUL18ExtIEsExtensionValue(value ngapType.DRBStatusUL18ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBsToQosFlowsMappingItemExtIEsExtensionValue(value ngapType.DRBsToQosFlowsMappingItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func Dynamic5QIDescriptorExtIEsExtensionValue(value ngapType.Dynamic5QIDescriptorExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue(value ngapType.EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue(value ngapType.EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue(value ngapType.EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledNRItemExtIEsExtensionValue(value ngapType.EmergencyAreaIDCancelledNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EmergencyFallbackIndicatorExtIEsExtensionValue(value ngapType.EmergencyFallbackIndicatorExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EndpointIPAddressAndPortExtIEsExtensionValue(value ngapType.EndpointIPAddressAndPortExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EndpointIPAddressAndPortExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EPSTAIExtIEsExtensionValue(value ngapType.EPSTAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EPSTAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ERABInformationItemExtIEsExtensionValue(value ngapType.ERABInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ERABInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func EUTRACGIExtIEsExtensionValue(value ngapType.EUTRACGIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EUTRACGIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ExpectedUEActivityBehaviourExtIEsExtensionValue(value ngapType.ExpectedUEActivityBehaviourExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ExpectedUEBehaviourExtIEsExtensionValue(value ngapType.ExpectedUEBehaviourExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ExpectedUEMovingTrajectoryItemExtIEsExtensionValue(value ngapType.ExpectedUEMovingTrajectoryItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ExtendedRATRestrictionInformationExtIEsExtensionValue(value ngapType.ExtendedRATRestrictionInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExtendedRATRestrictionInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func FiveGSTMSIExtIEsExtensionValue(value ngapType.FiveGSTMSIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ForbiddenAreaInformationItemExtIEsExtensionValue(value ngapType.ForbiddenAreaInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GBRQosInformationExtIEsExtensionValue(value ngapType.GBRQosInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GBRQosInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GlobalGNBIDExtIEsExtensionValue(value ngapType.GlobalGNBIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GlobalN3IWFIDExtIEsExtensionValue(value ngapType.GlobalN3IWFIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GlobalNgENBIDExtIEsExtensionValue(value ngapType.GlobalNgENBIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GTPTunnelExtIEsExtensionValue(value ngapType.GTPTunnelExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GTPTunnelExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GUAMIExtIEsExtensionValue(value ngapType.GUAMIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GUAMIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCommandTransferExtIEsExtensionValue(value ngapType.HandoverCommandTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCommandTransferExtIEsExtensionValuePresentAdditionalDLForwardingUPTNLInformation:
		if value.AdditionalDLForwardingUPTNLInformation == nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: AdditionalDLForwardingUPTNLInformation: NIL")
		}
		binData, bitEnd, err = QosFlowPerTNLInformationList(*value.AdditionalDLForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.HandoverCommandTransferExtIEsExtensionValuePresentULForwardingUPTNLInformation:
		if value.ULForwardingUPTNLInformation == nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: ULForwardingUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformation(*value.ULForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.HandoverCommandTransferExtIEsExtensionValuePresentAdditionalULForwardingUPTNLInformation:
		if value.AdditionalULForwardingUPTNLInformation == nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: AdditionalULForwardingUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformationList(*value.AdditionalULForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.HandoverCommandTransferExtIEsExtensionValuePresentDataForwardingResponseERABList:
		if value.DataForwardingResponseERABList == nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: DataForwardingResponseERABList: NIL")
		}
		binData, bitEnd, err = DataForwardingResponseERABList(*value.DataForwardingResponseERABList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCommandTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverPreparationUnsuccessfulTransferExtIEsExtensionValue(value ngapType.HandoverPreparationUnsuccessfulTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("HandoverPreparationUnsuccessfulTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequestAcknowledgeTransferExtIEsExtensionValue(value ngapType.HandoverRequestAcknowledgeTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequestAcknowledgeTransferExtIEsExtensionValuePresentAdditionalDLUPTNLInformationForHOList:
		if value.AdditionalDLUPTNLInformationForHOList == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: AdditionalDLUPTNLInformationForHOList: NIL")
		}
		binData, bitEnd, err = AdditionalDLUPTNLInformationForHOList(*value.AdditionalDLUPTNLInformationForHOList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeTransferExtIEsExtensionValuePresentULForwardingUPTNLInformation:
		if value.ULForwardingUPTNLInformation == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: ULForwardingUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformation(*value.ULForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeTransferExtIEsExtensionValuePresentAdditionalULForwardingUPTNLInformation:
		if value.AdditionalULForwardingUPTNLInformation == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: AdditionalULForwardingUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformationList(*value.AdditionalULForwardingUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeTransferExtIEsExtensionValuePresentDataForwardingResponseERABList:
		if value.DataForwardingResponseERABList == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: DataForwardingResponseERABList: NIL")
		}
		binData, bitEnd, err = DataForwardingResponseERABList(*value.DataForwardingResponseERABList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequiredTransferExtIEsExtensionValue(value ngapType.HandoverRequiredTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("HandoverRequiredTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValue(value ngapType.HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue(value ngapType.InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LAIExtIEsExtensionValue(value ngapType.LAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("LAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LastVisitedCellItemExtIEsExtensionValue(value ngapType.LastVisitedCellItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("LastVisitedCellItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LastVisitedNGRANCellInformationExtIEsExtensionValue(value ngapType.LastVisitedNGRANCellInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReportingRequestTypeExtIEsExtensionValue(value ngapType.LocationReportingRequestTypeExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportingRequestTypeExtIEsExtensionValuePresentLocationReportingAdditionalInfo:
		if value.LocationReportingAdditionalInfo == nil {
			return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEsExtensionValue: LocationReportingAdditionalInfo: NIL")
		}
		binData, bitEnd, err = LocationReportingAdditionalInfo(*value.LocationReportingAdditionalInfo, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func MobilityRestrictionListExtIEsExtensionValue(value ngapType.MobilityRestrictionListExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.MobilityRestrictionListExtIEsExtensionValuePresentLastEUTRANPLMNIdentity:
		if value.LastEUTRANPLMNIdentity == nil {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: LastEUTRANPLMNIdentity: NIL")
		}
		binData, bitEnd, err = PLMNIdentity(*value.LastEUTRANPLMNIdentity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.MobilityRestrictionListExtIEsExtensionValuePresentCNTypeRestrictionsForServing:
		if value.CNTypeRestrictionsForServing == nil {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: CNTypeRestrictionsForServing: NIL")
		}
		binData, bitEnd, err = CNTypeRestrictionsForServing(*value.CNTypeRestrictionsForServing, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.MobilityRestrictionListExtIEsExtensionValuePresentCNTypeRestrictionsForEquivalent:
		if value.CNTypeRestrictionsForEquivalent == nil {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: CNTypeRestrictionsForEquivalent: NIL")
		}
		binData, bitEnd, err = CNTypeRestrictionsForEquivalent(*value.CNTypeRestrictionsForEquivalent, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGRANTNLAssociationToRemoveItemExtIEsExtensionValue(value ngapType.NGRANTNLAssociationToRemoveItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NonDynamic5QIDescriptorExtIEsExtensionValue(value ngapType.NonDynamic5QIDescriptorExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NRCGIExtIEsExtensionValue(value ngapType.NRCGIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NRCGIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func OverloadStartNSSAIItemExtIEsExtensionValue(value ngapType.OverloadStartNSSAIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PacketErrorRateExtIEsExtensionValue(value ngapType.PacketErrorRateExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PacketErrorRateExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PagingAttemptInformationExtIEsExtensionValue(value ngapType.PagingAttemptInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue(value ngapType.PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestAcknowledgeTransferExtIEsExtensionValuePresentAdditionalNGUUPTNLInformation:
		if value.AdditionalNGUUPTNLInformation == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue: AdditionalNGUUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformationPairList(*value.AdditionalNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestSetupFailedTransferExtIEsExtensionValue(value ngapType.PathSwitchRequestSetupFailedTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestSetupFailedTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestTransferExtIEsExtensionValue(value ngapType.PathSwitchRequestTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestTransferExtIEsExtensionValuePresentAdditionalDLQosFlowPerTNLInformation:
		if value.AdditionalDLQosFlowPerTNLInformation == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEsExtensionValue: AdditionalDLQosFlowPerTNLInformation: NIL")
		}
		binData, bitEnd, err = QosFlowPerTNLInformationList(*value.AdditionalDLQosFlowPerTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValue(value ngapType.PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionAggregateMaximumBitRateExtIEsExtensionValue(value ngapType.PDUSessionAggregateMaximumBitRateExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceAdmittedItemExtIEsExtensionValue(value ngapType.PDUSessionResourceAdmittedItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue(value ngapType.PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceHandoverItemExtIEsExtensionValue(value ngapType.PDUSessionResourceHandoverItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceInformationItemExtIEsExtensionValue(value ngapType.PDUSessionResourceInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceItemCxtRelCplExtIEsExtensionValue(value ngapType.PDUSessionResourceItemCxtRelCplExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceItemCxtRelCplExtIEsExtensionValuePresentPDUSessionResourceReleaseResponseTransfer:
		if value.PDUSessionResourceReleaseResponseTransfer == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEsExtensionValue: PDUSessionResourceReleaseResponseTransfer: NIL")
		}
		binData, bitEnd, err = EncodeOctetStringNo(*value.PDUSessionResourceReleaseResponseTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceItemCxtRelReqExtIEsExtensionValue(value ngapType.PDUSessionResourceItemCxtRelReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceItemHORqdExtIEsExtensionValue(value ngapType.PDUSessionResourceItemHORqdExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyConfirmTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyConfirmTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyRequestTransferIEsTypeValue(value ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentPDUSessionAggregateMaximumBitRate:
		if value.PDUSessionAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: PDUSessionAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = PDUSessionAggregateMaximumBitRate(*value.PDUSessionAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentULNGUUPTNLModifyList:
		if value.ULNGUUPTNLModifyList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: ULNGUUPTNLModifyList: NIL")
		}
		binData, bitEnd, err = ULNGUUPTNLModifyList(*value.ULNGUUPTNLModifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentNetworkInstance:
		if value.NetworkInstance == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: NetworkInstance: NIL")
		}
		binData, bitEnd, err = NetworkInstance(*value.NetworkInstance, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentQosFlowAddOrModifyRequestList:
		if value.QosFlowAddOrModifyRequestList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: QosFlowAddOrModifyRequestList: NIL")
		}
		binData, bitEnd, err = QosFlowAddOrModifyRequestList(*value.QosFlowAddOrModifyRequestList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentQosFlowToReleaseList:
		if value.QosFlowToReleaseList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: QosFlowToReleaseList: NIL")
		}
		binData, bitEnd, err = QosFlowListWithCause(*value.QosFlowToReleaseList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentAdditionalULNGUUPTNLInformation:
		if value.AdditionalULNGUUPTNLInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: AdditionalULNGUUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformationList(*value.AdditionalULNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentCommonNetworkInstance:
		if value.CommonNetworkInstance == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: CommonNetworkInstance: NIL")
		}
		binData, bitEnd, err = CommonNetworkInstance(*value.CommonNetworkInstance, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyResponseTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyResponseTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyResponseTransferExtIEsExtensionValuePresentAdditionalNGUUPTNLInformation:
		if value.AdditionalNGUUPTNLInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEsExtensionValue: AdditionalNGUUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformationPairList(*value.AdditionalNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyIndicationTransferExtIEsExtensionValuePresentSecondaryRATUsageInformation:
		if value.SecondaryRATUsageInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue: SecondaryRATUsageInformation: NIL")
		}
		binData, bitEnd, err = SecondaryRATUsageInformation(*value.SecondaryRATUsageInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationTransferExtIEsExtensionValuePresentSecurityResult:
		if value.SecurityResult == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue: SecurityResult: NIL")
		}
		binData, bitEnd, err = SecurityResult(*value.SecurityResult, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModCfmExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyItemModCfmExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModIndExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyItemModIndExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModReqExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyItemModReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyItemModReqExtIEsExtensionValuePresentSNSSAI:
		if value.SNSSAI == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEsExtensionValue: SNSSAI: NIL")
		}
		binData, bitEnd, err = SNSSAI(*value.SNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyItemModResExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyItemModResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyItemExtIEsExtensionValue(value ngapType.PDUSessionResourceNotifyItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValuePresentSecondaryRATUsageInformation:
		if value.SecondaryRATUsageInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue: SecondaryRATUsageInformation: NIL")
		}
		binData, bitEnd, err = SecondaryRATUsageInformation(*value.SecondaryRATUsageInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceNotifyTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceNotifyTransferExtIEsExtensionValuePresentSecondaryRATUsageInformation:
		if value.SecondaryRATUsageInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEsExtensionValue: SecondaryRATUsageInformation: NIL")
		}
		binData, bitEnd, err = SecondaryRATUsageInformation(*value.SecondaryRATUsageInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseCommandTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceReleaseCommandTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemNotExtIEsExtensionValue(value ngapType.PDUSessionResourceReleasedItemNotExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue(value ngapType.PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue(value ngapType.PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedItemRelResExtIEsExtensionValue(value ngapType.PDUSessionResourceReleasedItemRelResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceReleaseResponseTransferExtIEsExtensionValuePresentSecondaryRATUsageInformation:
		if value.SecondaryRATUsageInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue: SecondaryRATUsageInformation: NIL")
		}
		binData, bitEnd, err = SecondaryRATUsageInformation(*value.SecondaryRATUsageInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValue(value ngapType.PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemCxtResExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupItemCxtResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemHOReqExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupItemHOReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemSUReqExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupItemSUReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupItemSUResExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupItemSUResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupRequestTransferIEsTypeValue(value ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentPDUSessionAggregateMaximumBitRate:
		if value.PDUSessionAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: PDUSessionAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = PDUSessionAggregateMaximumBitRate(*value.PDUSessionAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentULNGUUPTNLInformation:
		if value.ULNGUUPTNLInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: ULNGUUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformation(*value.ULNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentAdditionalULNGUUPTNLInformation:
		if value.AdditionalULNGUUPTNLInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: AdditionalULNGUUPTNLInformation: NIL")
		}
		binData, bitEnd, err = UPTransportLayerInformationList(*value.AdditionalULNGUUPTNLInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentDataForwardingNotPossible:
		if value.DataForwardingNotPossible == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: DataForwardingNotPossible: NIL")
		}
		binData, bitEnd, err = DataForwardingNotPossible(*value.DataForwardingNotPossible, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentPDUSessionType:
		if value.PDUSessionType == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: PDUSessionType: NIL")
		}
		binData, bitEnd, err = PDUSessionType(*value.PDUSessionType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentSecurityIndication:
		if value.SecurityIndication == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: SecurityIndication: NIL")
		}
		binData, bitEnd, err = SecurityIndication(*value.SecurityIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentNetworkInstance:
		if value.NetworkInstance == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: NetworkInstance: NIL")
		}
		binData, bitEnd, err = NetworkInstance(*value.NetworkInstance, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentQosFlowSetupRequestList:
		if value.QosFlowSetupRequestList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: QosFlowSetupRequestList: NIL")
		}
		binData, bitEnd, err = QosFlowSetupRequestList(*value.QosFlowSetupRequestList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentCommonNetworkInstance:
		if value.CommonNetworkInstance == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: CommonNetworkInstance: NIL")
		}
		binData, bitEnd, err = CommonNetworkInstance(*value.CommonNetworkInstance, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentDirectForwardingPathAvailability:
		if value.DirectForwardingPathAvailability == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: DirectForwardingPathAvailability: NIL")
		}
		binData, bitEnd, err = DirectForwardingPathAvailability(*value.DirectForwardingPathAvailability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupResponseTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupResponseTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValue(value ngapType.PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSwitchedItemExtIEsExtensionValue(value ngapType.PDUSessionResourceSwitchedItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue(value ngapType.PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue(value ngapType.PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue(value ngapType.PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionUsageReportExtIEsExtensionValue(value ngapType.PDUSessionUsageReportExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionUsageReportExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PLMNSupportItemExtIEsExtensionValue(value ngapType.PLMNSupportItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowAcceptedItemExtIEsExtensionValue(value ngapType.QosFlowAcceptedItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowAcceptedItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyRequestItemExtIEsExtensionValue(value ngapType.QosFlowAddOrModifyRequestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyResponseItemExtIEsExtensionValue(value ngapType.QosFlowAddOrModifyResponseItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowInformationItemExtIEsExtensionValue(value ngapType.QosFlowInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.QosFlowInformationItemExtIEsExtensionValuePresentULForwarding:
		if value.ULForwarding == nil {
			return binData, bitEnd, errors.New("QosFlowInformationItemExtIEsExtensionValue: ULForwarding: NIL")
		}
		binData, bitEnd, err = ULForwarding(*value.ULForwarding, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowInformationItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("QosFlowInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowLevelQosParametersExtIEsExtensionValue(value ngapType.QosFlowLevelQosParametersExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.QosFlowLevelQosParametersExtIEsExtensionValuePresentQosMonitoringRequest:
		if value.QosMonitoringRequest == nil {
			return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEsExtensionValue: QosMonitoringRequest: NIL")
		}
		binData, bitEnd, err = QosMonitoringRequest(*value.QosMonitoringRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowWithCauseItemExtIEsExtensionValue(value ngapType.QosFlowWithCauseItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowWithCauseItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowModifyConfirmItemExtIEsExtensionValue(value ngapType.QosFlowModifyConfirmItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowModifyConfirmItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowNotifyItemExtIEsExtensionValue(value ngapType.QosFlowNotifyItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowNotifyItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowPerTNLInformationExtIEsExtensionValue(value ngapType.QosFlowPerTNLInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowPerTNLInformationItemExtIEsExtensionValue(value ngapType.QosFlowPerTNLInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowSetupRequestItemExtIEsExtensionValue(value ngapType.QosFlowSetupRequestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowItemWithDataForwardingExtIEsExtensionValue(value ngapType.QosFlowItemWithDataForwardingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowItemWithDataForwardingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosFlowToBeForwardedItemExtIEsExtensionValue(value ngapType.QosFlowToBeForwardedItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowToBeForwardedItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QoSFlowsUsageReportItemExtIEsExtensionValue(value ngapType.QoSFlowsUsageReportItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANStatusTransferTransparentContainerExtIEsExtensionValue(value ngapType.RANStatusTransferTransparentContainerExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RATRestrictionsItemExtIEsExtensionValue(value ngapType.RATRestrictionsItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RATRestrictionsItemExtIEsExtensionValuePresentExtendedRATRestrictionInformation:
		if value.ExtendedRATRestrictionInformation == nil {
			return binData, bitEnd, errors.New("RATRestrictionsItemExtIEsExtensionValue: ExtendedRATRestrictionInformation: NIL")
		}
		binData, bitEnd, err = ExtendedRATRestrictionInformation(*value.ExtendedRATRestrictionInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RATRestrictionsItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RecommendedCellsForPagingExtIEsExtensionValue(value ngapType.RecommendedCellsForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RecommendedCellItemExtIEsExtensionValue(value ngapType.RecommendedCellItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RecommendedRANNodesForPagingExtIEsExtensionValue(value ngapType.RecommendedRANNodesForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RecommendedRANNodeItemExtIEsExtensionValue(value ngapType.RecommendedRANNodeItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RIMInformationTransferExtIEsExtensionValue(value ngapType.RIMInformationTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RIMInformationTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SecondaryRATUsageInformationExtIEsExtensionValue(value ngapType.SecondaryRATUsageInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SecondaryRATUsageInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SecondaryRATDataUsageReportTransferExtIEsExtensionValue(value ngapType.SecondaryRATDataUsageReportTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SecurityContextExtIEsExtensionValue(value ngapType.SecurityContextExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SecurityContextExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SecurityIndicationExtIEsExtensionValue(value ngapType.SecurityIndicationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.SecurityIndicationExtIEsExtensionValuePresentMaximumIntegrityProtectedDataRateDL:
		if value.MaximumIntegrityProtectedDataRateDL == nil {
			return binData, bitEnd, errors.New("SecurityIndicationExtIEsExtensionValue: MaximumIntegrityProtectedDataRateDL: NIL")
		}
		binData, bitEnd, err = MaximumIntegrityProtectedDataRate(*value.MaximumIntegrityProtectedDataRateDL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecurityIndicationExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("SecurityIndicationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SecurityResultExtIEsExtensionValue(value ngapType.SecurityResultExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SecurityResultExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ServedGUAMIItemExtIEsExtensionValue(value ngapType.ServedGUAMIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.ServedGUAMIItemExtIEsExtensionValuePresentGUAMIType:
		if value.GUAMIType == nil {
			return binData, bitEnd, errors.New("ServedGUAMIItemExtIEsExtensionValue: GUAMIType: NIL")
		}
		binData, bitEnd, err = GUAMIType(*value.GUAMIType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServedGUAMIItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ServiceAreaInformationItemExtIEsExtensionValue(value ngapType.ServiceAreaInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SliceOverloadItemExtIEsExtensionValue(value ngapType.SliceOverloadItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SliceSupportItemExtIEsExtensionValue(value ngapType.SliceSupportItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SliceSupportItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SNSSAIExtIEsExtensionValue(value ngapType.SNSSAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SNSSAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SONConfigurationTransferExtIEsExtensionValue(value ngapType.SONConfigurationTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SONInformationReplyExtIEsExtensionValue(value ngapType.SONInformationReplyExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SONInformationReplyExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue(value ngapType.SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValuePresentSgNBUEX2APID:
		if value.SgNBUEX2APID == nil {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue: SgNBUEX2APID: NIL")
		}
		binData, bitEnd, err = SgNBUEX2APID(*value.SgNBUEX2APID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SourceRANNodeIDExtIEsExtensionValue(value ngapType.SourceRANNodeIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SourceToTargetAMFInformationRerouteExtIEsExtensionValue(value ngapType.SourceToTargetAMFInformationRerouteExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SourceToTargetAMFInformationRerouteExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SupportedTAItemExtIEsExtensionValue(value ngapType.SupportedTAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.SupportedTAItemExtIEsExtensionValuePresentRATInformation:
		if value.RATInformation == nil {
			return binData, bitEnd, errors.New("SupportedTAItemExtIEsExtensionValue: RATInformation: NIL")
		}
		binData, bitEnd, err = RATInformation(*value.RATInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SupportedTAItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("SupportedTAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIExtIEsExtensionValue(value ngapType.TAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIBroadcastEUTRAItemExtIEsExtensionValue(value ngapType.TAIBroadcastEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIBroadcastNRItemExtIEsExtensionValue(value ngapType.TAIBroadcastNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAICancelledEUTRAItemExtIEsExtensionValue(value ngapType.TAICancelledEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAICancelledNRItemExtIEsExtensionValue(value ngapType.TAICancelledNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIListForInactiveItemExtIEsExtensionValue(value ngapType.TAIListForInactiveItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIListForPagingItemExtIEsExtensionValue(value ngapType.TAIListForPagingItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TargeteNBIDExtIEsExtensionValue(value ngapType.TargeteNBIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TargeteNBIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValue(value ngapType.TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TargetRANNodeIDExtIEsExtensionValue(value ngapType.TargetRANNodeIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TargetRNCIDExtIEsExtensionValue(value ngapType.TargetRNCIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TargetRNCIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TNLAssociationItemExtIEsExtensionValue(value ngapType.TNLAssociationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TraceActivationExtIEsExtensionValue(value ngapType.TraceActivationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TraceActivationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEAggregateMaximumBitRateExtIEsExtensionValue(value ngapType.UEAggregateMaximumBitRateExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue(value ngapType.UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UENGAPIDPairExtIEsExtensionValue(value ngapType.UENGAPIDPairExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEPresenceInAreaOfInterestItemExtIEsExtensionValue(value ngapType.UEPresenceInAreaOfInterestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityForPagingExtIEsExtensionValue(value ngapType.UERadioCapabilityForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UESecurityCapabilitiesExtIEsExtensionValue(value ngapType.UESecurityCapabilitiesExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ULNGUUPTNLModifyItemExtIEsExtensionValue(value ngapType.ULNGUUPTNLModifyItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UnavailableGUAMIItemExtIEsExtensionValue(value ngapType.UnavailableGUAMIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationItemExtIEsExtensionValue(value ngapType.UPTransportLayerInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UPTransportLayerInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationPairItemExtIEsExtensionValue(value ngapType.UPTransportLayerInformationPairItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserLocationInformationEUTRAExtIEsExtensionValue(value ngapType.UserLocationInformationEUTRAExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UserLocationInformationEUTRAExtIEsExtensionValuePresentPSCellInformation:
		if value.PSCellInformation == nil {
			return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEsExtensionValue: PSCellInformation: NIL")
		}
		binData, bitEnd, err = NGRANCGI(*value.PSCellInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserLocationInformationN3IWFExtIEsExtensionValue(value ngapType.UserLocationInformationN3IWFExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserLocationInformationNRExtIEsExtensionValue(value ngapType.UserLocationInformationNRExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UserLocationInformationNRExtIEsExtensionValuePresentPSCellInformation:
		if value.PSCellInformation == nil {
			return binData, bitEnd, errors.New("UserLocationInformationNRExtIEsExtensionValue: PSCellInformation: NIL")
		}
		binData, bitEnd, err = NGRANCGI(*value.PSCellInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformationNRExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserPlaneSecurityInformationExtIEsExtensionValue(value ngapType.UserPlaneSecurityInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UserPlaneSecurityInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func VolumeTimedReportItemExtIEsExtensionValue(value ngapType.VolumeTimedReportItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("VolumeTimedReportItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func XnExtTLAItemExtIEsExtensionValue(value ngapType.XnExtTLAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.XnExtTLAItemExtIEsExtensionValuePresentSCTPTLAs:
		if value.SCTPTLAs == nil {
			return binData, bitEnd, errors.New("XnExtTLAItemExtIEsExtensionValue: SCTPTLAs: NIL")
		}
		binData, bitEnd, err = SCTPTLAs(*value.SCTPTLAs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnExtTLAItemExtIEsExtensionValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func XnTNLConfigurationInfoExtIEsExtensionValue(value ngapType.XnTNLConfigurationInfoExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}
