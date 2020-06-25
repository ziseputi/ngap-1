package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AdditionalDLUPTNLInformationForHOItemExtIEs(value ngapType.AdditionalDLUPTNLInformationForHOItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AllocationAndRetentionPriorityExtIEs(value ngapType.AllocationAndRetentionPriorityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AllocationAndRetentionPriorityExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AllowedNSSAIItemExtIEs(value ngapType.AllowedNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AllowedNSSAIItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFTNLAssociationSetupItemExtIEs(value ngapType.AMFTNLAssociationSetupItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFTNLAssociationSetupItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFTNLAssociationToAddItemExtIEs(value ngapType.AMFTNLAssociationToAddItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFTNLAssociationToAddItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFTNLAssociationToRemoveItemExtIEs(value ngapType.AMFTNLAssociationToRemoveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDTNLAssociationTransportLayerAddressNGRAN:
		if value.ExtensionValue.Present != ngapType.AMFTNLAssociationToRemoveItemExtIEsPresentTNLAssociationTransportLayerAddressNGRAN {
			return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFTNLAssociationToRemoveItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFTNLAssociationToUpdateItemExtIEs(value ngapType.AMFTNLAssociationToUpdateItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFTNLAssociationToUpdateItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AreaOfInterestExtIEs(value ngapType.AreaOfInterestExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AreaOfInterestExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AreaOfInterestCellItemExtIEs(value ngapType.AreaOfInterestCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AreaOfInterestCellItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AreaOfInterestItemExtIEs(value ngapType.AreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AreaOfInterestItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AreaOfInterestRANNodeItemExtIEs(value ngapType.AreaOfInterestRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AreaOfInterestRANNodeItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AreaOfInterestTAIItemExtIEs(value ngapType.AreaOfInterestTAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AreaOfInterestTAIItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AssistanceDataForPagingExtIEs(value ngapType.AssistanceDataForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AssistanceDataForPagingExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AssistanceDataForRecommendedCellsExtIEs(value ngapType.AssistanceDataForRecommendedCellsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AssistanceDataForRecommendedCellsExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AssociatedQosFlowItemExtIEs(value ngapType.AssociatedQosFlowItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssociatedQosFlowItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssociatedQosFlowItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AssociatedQosFlowItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AssociatedQosFlowItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func BroadcastPLMNItemExtIEs(value ngapType.BroadcastPLMNItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = BroadcastPLMNItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CancelledCellsInEAIEUTRAItemExtIEs(value ngapType.CancelledCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CancelledCellsInEAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CancelledCellsInEAINRItemExtIEs(value ngapType.CancelledCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CancelledCellsInEAINRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CancelledCellsInTAIEUTRAItemExtIEs(value ngapType.CancelledCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CancelledCellsInTAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CancelledCellsInTAINRItemExtIEs(value ngapType.CancelledCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CancelledCellsInTAINRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellIDBroadcastEUTRAItemExtIEs(value ngapType.CellIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellIDBroadcastEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellIDBroadcastNRItemExtIEs(value ngapType.CellIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellIDBroadcastNRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellIDCancelledEUTRAItemExtIEs(value ngapType.CellIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellIDCancelledEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellIDCancelledNRItemExtIEs(value ngapType.CellIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellIDCancelledNRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellTypeExtIEs(value ngapType.CellTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTypeExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTypeExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellTypeExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellTypeExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CNAssistedRANTuningExtIEs(value ngapType.CNAssistedRANTuningExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CNAssistedRANTuningExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CNAssistedRANTuningExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CNAssistedRANTuningExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CNAssistedRANTuningExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CNTypeRestrictionsForEquivalentItemExtIEs(value ngapType.CNTypeRestrictionsForEquivalentItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CNTypeRestrictionsForEquivalentItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalentItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CompletedCellsInEAIEUTRAItemExtIEs(value ngapType.CompletedCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CompletedCellsInEAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CompletedCellsInEAINRItemExtIEs(value ngapType.CompletedCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CompletedCellsInEAINRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CompletedCellsInTAIEUTRAItemExtIEs(value ngapType.CompletedCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CompletedCellsInTAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CompletedCellsInTAINRItemExtIEs(value ngapType.CompletedCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CompletedCellsInTAINRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CoreNetworkAssistanceInformationForInactiveExtIEs(value ngapType.CoreNetworkAssistanceInformationForInactiveExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactiveExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactiveExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationForInactiveExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func COUNTValueForPDCPSN12ExtIEs(value ngapType.COUNTValueForPDCPSN12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = COUNTValueForPDCPSN12ExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func COUNTValueForPDCPSN18ExtIEs(value ngapType.COUNTValueForPDCPSN18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = COUNTValueForPDCPSN18ExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CriticalityDiagnosticsExtIEs(value ngapType.CriticalityDiagnosticsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CriticalityDiagnosticsExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CriticalityDiagnosticsIEItemExtIEs(value ngapType.CriticalityDiagnosticsIEItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CriticalityDiagnosticsIEItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DataForwardingResponseDRBItemExtIEs(value ngapType.DataForwardingResponseDRBItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseDRBItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseDRBItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DataForwardingResponseDRBItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseDRBItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DataForwardingResponseERABListItemExtIEs(value ngapType.DataForwardingResponseERABListItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseERABListItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseERABListItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DataForwardingResponseERABListItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DataForwardingResponseERABListItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBsSubjectToStatusTransferItemExtIEs(value ngapType.DRBsSubjectToStatusTransferItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDOldAssociatedQosFlowListULendmarkerexpected:
		if value.ExtensionValue.Present != ngapType.DRBsSubjectToStatusTransferItemExtIEsPresentOldAssociatedQosFlowListULendmarkerexpected {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBsSubjectToStatusTransferItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBStatusDL12ExtIEs(value ngapType.DRBStatusDL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBStatusDL12ExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBStatusDL18ExtIEs(value ngapType.DRBStatusDL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBStatusDL18ExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBStatusUL12ExtIEs(value ngapType.DRBStatusUL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBStatusUL12ExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBStatusUL18ExtIEs(value ngapType.DRBStatusUL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBStatusUL18ExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBsToQosFlowsMappingItemExtIEs(value ngapType.DRBsToQosFlowsMappingItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBsToQosFlowsMappingItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func Dynamic5QIDescriptorExtIEs(value ngapType.Dynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = Dynamic5QIDescriptorExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EmergencyAreaIDBroadcastEUTRAItemExtIEs(value ngapType.EmergencyAreaIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EmergencyAreaIDBroadcastNRItemExtIEs(value ngapType.EmergencyAreaIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EmergencyAreaIDCancelledEUTRAItemExtIEs(value ngapType.EmergencyAreaIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EmergencyAreaIDCancelledNRItemExtIEs(value ngapType.EmergencyAreaIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EmergencyAreaIDCancelledNRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EmergencyFallbackIndicatorExtIEs(value ngapType.EmergencyFallbackIndicatorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EmergencyFallbackIndicatorExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EndpointIPAddressAndPortExtIEs(value ngapType.EndpointIPAddressAndPortExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EndpointIPAddressAndPortExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EndpointIPAddressAndPortExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EndpointIPAddressAndPortExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EndpointIPAddressAndPortExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EPSTAIExtIEs(value ngapType.EPSTAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EPSTAIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ERABInformationItemExtIEs(value ngapType.ERABInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ERABInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ERABInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ERABInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ERABInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func EUTRACGIExtIEs(value ngapType.EUTRACGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = EUTRACGIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ExpectedUEActivityBehaviourExtIEs(value ngapType.ExpectedUEActivityBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ExpectedUEActivityBehaviourExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ExpectedUEBehaviourExtIEs(value ngapType.ExpectedUEBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ExpectedUEBehaviourExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ExpectedUEMovingTrajectoryItemExtIEs(value ngapType.ExpectedUEMovingTrajectoryItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ExpectedUEMovingTrajectoryItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ExtendedRATRestrictionInformationExtIEs(value ngapType.ExtendedRATRestrictionInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExtendedRATRestrictionInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExtendedRATRestrictionInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ExtendedRATRestrictionInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ExtendedRATRestrictionInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func FiveGSTMSIExtIEs(value ngapType.FiveGSTMSIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = FiveGSTMSIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ForbiddenAreaInformationItemExtIEs(value ngapType.ForbiddenAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ForbiddenAreaInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GBRQosInformationExtIEs(value ngapType.GBRQosInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GBRQosInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GlobalGNBIDExtIEs(value ngapType.GlobalGNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GlobalGNBIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GlobalN3IWFIDExtIEs(value ngapType.GlobalN3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GlobalN3IWFIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GlobalNgENBIDExtIEs(value ngapType.GlobalNgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GlobalNgENBIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GTPTunnelExtIEs(value ngapType.GTPTunnelExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnelExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnelExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GTPTunnelExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnelExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GUAMIExtIEs(value ngapType.GUAMIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GUAMIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverCommandTransferExtIEs(value ngapType.HandoverCommandTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDAdditionalDLForwardingUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.HandoverCommandTransferExtIEsPresentAdditionalDLForwardingUPTNLInformation {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDULForwardingUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.HandoverCommandTransferExtIEsPresentULForwardingUPTNLInformation {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDAdditionalULForwardingUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.HandoverCommandTransferExtIEsPresentAdditionalULForwardingUPTNLInformation {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDDataForwardingResponseERABList:
		if value.ExtensionValue.Present != ngapType.HandoverCommandTransferExtIEsPresentDataForwardingResponseERABList {
			return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCommandTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverPreparationUnsuccessfulTransferExtIEs(value ngapType.HandoverPreparationUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverPreparationUnsuccessfulTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverRequestAcknowledgeTransferExtIEs(value ngapType.HandoverRequestAcknowledgeTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDAdditionalDLUPTNLInformationForHOList:
		if value.ExtensionValue.Present != ngapType.HandoverRequestAcknowledgeTransferExtIEsPresentAdditionalDLUPTNLInformationForHOList {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDULForwardingUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.HandoverRequestAcknowledgeTransferExtIEsPresentULForwardingUPTNLInformation {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDAdditionalULForwardingUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.HandoverRequestAcknowledgeTransferExtIEsPresentAdditionalULForwardingUPTNLInformation {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDDataForwardingResponseERABList:
		if value.ExtensionValue.Present != ngapType.HandoverRequestAcknowledgeTransferExtIEsPresentDataForwardingResponseERABList {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequestAcknowledgeTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverRequiredTransferExtIEs(value ngapType.HandoverRequiredTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequiredTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverResourceAllocationUnsuccessfulTransferExtIEs(value ngapType.HandoverResourceAllocationUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverResourceAllocationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func InfoOnRecommendedCellsAndRANNodesForPagingExtIEs(value ngapType.InfoOnRecommendedCellsAndRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LAIExtIEs(value ngapType.LAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LAIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LAIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LAIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LastVisitedCellItemExtIEs(value ngapType.LastVisitedCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LastVisitedCellItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LastVisitedNGRANCellInformationExtIEs(value ngapType.LastVisitedNGRANCellInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LastVisitedNGRANCellInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedNGRANCellInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LocationReportingRequestTypeExtIEs(value ngapType.LocationReportingRequestTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDLocationReportingAdditionalInfo:
		if value.ExtensionValue.Present != ngapType.LocationReportingRequestTypeExtIEsPresentLocationReportingAdditionalInfo {
			return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportingRequestTypeExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func MobilityRestrictionListExtIEs(value ngapType.MobilityRestrictionListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDLastEUTRANPLMNIdentity:
		if value.ExtensionValue.Present != ngapType.MobilityRestrictionListExtIEsPresentLastEUTRANPLMNIdentity {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDCNTypeRestrictionsForServing:
		if value.ExtensionValue.Present != ngapType.MobilityRestrictionListExtIEsPresentCNTypeRestrictionsForServing {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDCNTypeRestrictionsForEquivalent:
		if value.ExtensionValue.Present != ngapType.MobilityRestrictionListExtIEsPresentCNTypeRestrictionsForEquivalent {
			return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = MobilityRestrictionListExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGRANTNLAssociationToRemoveItemExtIEs(value ngapType.NGRANTNLAssociationToRemoveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGRANTNLAssociationToRemoveItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NonDynamic5QIDescriptorExtIEs(value ngapType.NonDynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NonDynamic5QIDescriptorExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NRCGIExtIEs(value ngapType.NRCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NRCGIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func OverloadStartNSSAIItemExtIEs(value ngapType.OverloadStartNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = OverloadStartNSSAIItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PacketErrorRateExtIEs(value ngapType.PacketErrorRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PacketErrorRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PacketErrorRateExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PacketErrorRateExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PacketErrorRateExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PagingAttemptInformationExtIEs(value ngapType.PagingAttemptInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PagingAttemptInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestAcknowledgeTransferExtIEs(value ngapType.PathSwitchRequestAcknowledgeTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDAdditionalNGUUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.PathSwitchRequestAcknowledgeTransferExtIEsPresentAdditionalNGUUPTNLInformation {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestSetupFailedTransferExtIEs(value ngapType.PathSwitchRequestSetupFailedTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestSetupFailedTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestSetupFailedTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestSetupFailedTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestSetupFailedTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestTransferExtIEs(value ngapType.PathSwitchRequestTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDAdditionalDLQosFlowPerTNLInformation:
		if value.ExtensionValue.Present != ngapType.PathSwitchRequestTransferExtIEsPresentAdditionalDLQosFlowPerTNLInformation {
			return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestUnsuccessfulTransferExtIEs(value ngapType.PathSwitchRequestUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestUnsuccessfulTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestUnsuccessfulTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestUnsuccessfulTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionAggregateMaximumBitRateExtIEs(value ngapType.PDUSessionAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionAggregateMaximumBitRateExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceAdmittedItemExtIEs(value ngapType.PDUSessionResourceAdmittedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceAdmittedItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToModifyItemModCfmExtIEs(value ngapType.PDUSessionResourceFailedToModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToModifyItemModResExtIEs(value ngapType.PDUSessionResourceFailedToModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToSetupItemCxtFailExtIEs(value ngapType.PDUSessionResourceFailedToSetupItemCxtFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToSetupItemCxtResExtIEs(value ngapType.PDUSessionResourceFailedToSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToSetupItemHOAckExtIEs(value ngapType.PDUSessionResourceFailedToSetupItemHOAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToSetupItemPSReqExtIEs(value ngapType.PDUSessionResourceFailedToSetupItemPSReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceFailedToSetupItemSUResExtIEs(value ngapType.PDUSessionResourceFailedToSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceHandoverItemExtIEs(value ngapType.PDUSessionResourceHandoverItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceHandoverItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceInformationItemExtIEs(value ngapType.PDUSessionResourceInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceItemCxtRelCplExtIEs(value ngapType.PDUSessionResourceItemCxtRelCplExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDPDUSessionResourceReleaseResponseTransfer:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceItemCxtRelCplExtIEsPresentPDUSessionResourceReleaseResponseTransfer {
			return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceItemCxtRelCplExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceItemCxtRelReqExtIEs(value ngapType.PDUSessionResourceItemCxtRelReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceItemCxtRelReqExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceItemHORqdExtIEs(value ngapType.PDUSessionResourceItemHORqdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceItemHORqdExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyConfirmTransferExtIEs(value ngapType.PDUSessionResourceModifyConfirmTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyConfirmTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs(value ngapType.PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyRequestTransferIEs(value ngapType.PDUSessionResourceModifyRequestTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDPDUSessionAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentPDUSessionAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDULNGUUPTNLModifyList:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentULNGUUPTNLModifyList {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNetworkInstance:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentNetworkInstance {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDQosFlowAddOrModifyRequestList:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentQosFlowAddOrModifyRequestList {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDQosFlowToReleaseList:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentQosFlowToReleaseList {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAdditionalULNGUUPTNLInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentAdditionalULNGUUPTNLInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCommonNetworkInstance:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestTransferIEsTypeValuePresentCommonNetworkInstance {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyRequestTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyResponseTransferExtIEs(value ngapType.PDUSessionResourceModifyResponseTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDAdditionalNGUUPTNLInformation:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceModifyResponseTransferExtIEsPresentAdditionalNGUUPTNLInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyResponseTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyIndicationTransferExtIEs(value ngapType.PDUSessionResourceModifyIndicationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSecondaryRATUsageInformation:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceModifyIndicationTransferExtIEsPresentSecondaryRATUsageInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityResult:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceModifyIndicationTransferExtIEsPresentSecurityResult {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyItemModCfmExtIEs(value ngapType.PDUSessionResourceModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyItemModCfmExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyItemModIndExtIEs(value ngapType.PDUSessionResourceModifyItemModIndExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyItemModIndExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyItemModReqExtIEs(value ngapType.PDUSessionResourceModifyItemModReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSNSSAI:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceModifyItemModReqExtIEsPresentSNSSAI {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyItemModReqExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyItemModResExtIEs(value ngapType.PDUSessionResourceModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyItemModResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyUnsuccessfulTransferExtIEs(value ngapType.PDUSessionResourceModifyUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyUnsuccessfulTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceNotifyItemExtIEs(value ngapType.PDUSessionResourceNotifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceNotifyItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceNotifyReleasedTransferExtIEs(value ngapType.PDUSessionResourceNotifyReleasedTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSecondaryRATUsageInformation:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceNotifyReleasedTransferExtIEsPresentSecondaryRATUsageInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyReleasedTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceNotifyTransferExtIEs(value ngapType.PDUSessionResourceNotifyTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSecondaryRATUsageInformation:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceNotifyTransferExtIEsPresentSecondaryRATUsageInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceNotifyTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleaseCommandTransferExtIEs(value ngapType.PDUSessionResourceReleaseCommandTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleaseCommandTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleasedItemNotExtIEs(value ngapType.PDUSessionResourceReleasedItemNotExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleasedItemNotExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleasedItemPSAckExtIEs(value ngapType.PDUSessionResourceReleasedItemPSAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleasedItemPSFailExtIEs(value ngapType.PDUSessionResourceReleasedItemPSFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleasedItemRelResExtIEs(value ngapType.PDUSessionResourceReleasedItemRelResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleasedItemRelResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleaseResponseTransferExtIEs(value ngapType.PDUSessionResourceReleaseResponseTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSecondaryRATUsageInformation:
		if value.ExtensionValue.Present != ngapType.PDUSessionResourceReleaseResponseTransferExtIEsPresentSecondaryRATUsageInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSecondaryRATUsageItemExtIEs(value ngapType.PDUSessionResourceSecondaryRATUsageItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupItemCxtReqExtIEs(value ngapType.PDUSessionResourceSetupItemCxtReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupItemCxtResExtIEs(value ngapType.PDUSessionResourceSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupItemCxtResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupItemHOReqExtIEs(value ngapType.PDUSessionResourceSetupItemHOReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupItemHOReqExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupItemSUReqExtIEs(value ngapType.PDUSessionResourceSetupItemSUReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupItemSUReqExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupItemSUResExtIEs(value ngapType.PDUSessionResourceSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupItemSUResExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupRequestTransferIEs(value ngapType.PDUSessionResourceSetupRequestTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDPDUSessionAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentPDUSessionAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDULNGUUPTNLInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentULNGUUPTNLInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAdditionalULNGUUPTNLInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentAdditionalULNGUUPTNLInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDataForwardingNotPossible:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentDataForwardingNotPossible {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionType:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentPDUSessionType {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityIndication:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentSecurityIndication {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNetworkInstance:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentNetworkInstance {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDQosFlowSetupRequestList:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentQosFlowSetupRequestList {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCommonNetworkInstance:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentCommonNetworkInstance {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDirectForwardingPathAvailability:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestTransferIEsTypeValuePresentDirectForwardingPathAvailability {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupRequestTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupResponseTransferExtIEs(value ngapType.PDUSessionResourceSetupResponseTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupResponseTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupUnsuccessfulTransferExtIEs(value ngapType.PDUSessionResourceSetupUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupUnsuccessfulTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSwitchedItemExtIEs(value ngapType.PDUSessionResourceSwitchedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSwitchedItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceToBeSwitchedDLItemExtIEs(value ngapType.PDUSessionResourceToBeSwitchedDLItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceToReleaseItemHOCmdExtIEs(value ngapType.PDUSessionResourceToReleaseItemHOCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceToReleaseItemRelCmdExtIEs(value ngapType.PDUSessionResourceToReleaseItemRelCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionUsageReportExtIEs(value ngapType.PDUSessionUsageReportExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionUsageReportExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionUsageReportExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionUsageReportExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionUsageReportExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PLMNSupportItemExtIEs(value ngapType.PLMNSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PLMNSupportItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowAcceptedItemExtIEs(value ngapType.QosFlowAcceptedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAcceptedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAcceptedItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowAcceptedItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAcceptedItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowAddOrModifyRequestItemExtIEs(value ngapType.QosFlowAddOrModifyRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowAddOrModifyRequestItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowAddOrModifyResponseItemExtIEs(value ngapType.QosFlowAddOrModifyResponseItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowAddOrModifyResponseItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowInformationItemExtIEs(value ngapType.QosFlowInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDULForwarding:
		if value.ExtensionValue.Present != ngapType.QosFlowInformationItemExtIEsPresentULForwarding {
			return binData, bitEnd, errors.New("QosFlowInformationItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("QosFlowInformationItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowLevelQosParametersExtIEs(value ngapType.QosFlowLevelQosParametersExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDQosMonitoringRequest:
		if value.ExtensionValue.Present != ngapType.QosFlowLevelQosParametersExtIEsPresentQosMonitoringRequest {
			return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowLevelQosParametersExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowWithCauseItemExtIEs(value ngapType.QosFlowWithCauseItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowWithCauseItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowWithCauseItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowWithCauseItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowWithCauseItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowModifyConfirmItemExtIEs(value ngapType.QosFlowModifyConfirmItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowModifyConfirmItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowModifyConfirmItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowModifyConfirmItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowModifyConfirmItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowNotifyItemExtIEs(value ngapType.QosFlowNotifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowNotifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowNotifyItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowNotifyItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowNotifyItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowPerTNLInformationExtIEs(value ngapType.QosFlowPerTNLInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowPerTNLInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowPerTNLInformationItemExtIEs(value ngapType.QosFlowPerTNLInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowPerTNLInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowSetupRequestItemExtIEs(value ngapType.QosFlowSetupRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowSetupRequestItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowItemWithDataForwardingExtIEs(value ngapType.QosFlowItemWithDataForwardingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemWithDataForwardingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemWithDataForwardingExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowItemWithDataForwardingExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemWithDataForwardingExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosFlowToBeForwardedItemExtIEs(value ngapType.QosFlowToBeForwardedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowToBeForwardedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowToBeForwardedItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosFlowToBeForwardedItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowToBeForwardedItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QoSFlowsUsageReportItemExtIEs(value ngapType.QoSFlowsUsageReportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QoSFlowsUsageReportItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RANStatusTransferTransparentContainerExtIEs(value ngapType.RANStatusTransferTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANStatusTransferTransparentContainerExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RATRestrictionsItemExtIEs(value ngapType.RATRestrictionsItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDExtendedRATRestrictionInformation:
		if value.ExtensionValue.Present != ngapType.RATRestrictionsItemExtIEsPresentExtendedRATRestrictionInformation {
			return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RATRestrictionsItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RecommendedCellsForPagingExtIEs(value ngapType.RecommendedCellsForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RecommendedCellsForPagingExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RecommendedCellItemExtIEs(value ngapType.RecommendedCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RecommendedCellItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RecommendedRANNodesForPagingExtIEs(value ngapType.RecommendedRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RecommendedRANNodesForPagingExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RecommendedRANNodeItemExtIEs(value ngapType.RecommendedRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RecommendedRANNodeItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RIMInformationTransferExtIEs(value ngapType.RIMInformationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformationTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformationTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RIMInformationTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RIMInformationTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SecondaryRATUsageInformationExtIEs(value ngapType.SecondaryRATUsageInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATUsageInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATUsageInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SecondaryRATUsageInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATUsageInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SecondaryRATDataUsageReportTransferExtIEs(value ngapType.SecondaryRATDataUsageReportTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SecondaryRATDataUsageReportTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SecurityContextExtIEs(value ngapType.SecurityContextExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContextExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContextExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SecurityContextExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContextExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SecurityIndicationExtIEs(value ngapType.SecurityIndicationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDMaximumIntegrityProtectedDataRateDL:
		if value.ExtensionValue.Present != ngapType.SecurityIndicationExtIEsPresentMaximumIntegrityProtectedDataRateDL {
			return binData, bitEnd, errors.New("SecurityIndicationExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SecurityIndicationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SecurityResultExtIEs(value ngapType.SecurityResultExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityResultExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityResultExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SecurityResultExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityResultExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ServedGUAMIItemExtIEs(value ngapType.ServedGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDGUAMIType:
		if value.ExtensionValue.Present != ngapType.ServedGUAMIItemExtIEsPresentGUAMIType {
			return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ServedGUAMIItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ServiceAreaInformationItemExtIEs(value ngapType.ServiceAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ServiceAreaInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SliceOverloadItemExtIEs(value ngapType.SliceOverloadItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SliceOverloadItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SliceSupportItemExtIEs(value ngapType.SliceSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SliceSupportItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SNSSAIExtIEs(value ngapType.SNSSAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SNSSAIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SONConfigurationTransferExtIEs(value ngapType.SONConfigurationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SONConfigurationTransferExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SONInformationReplyExtIEs(value ngapType.SONInformationReplyExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationReplyExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationReplyExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SONInformationReplyExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationReplyExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs(value ngapType.SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSgNBUEX2APID:
		if value.ExtensionValue.Present != ngapType.SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsPresentSgNBUEX2APID {
			return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SourceRANNodeIDExtIEs(value ngapType.SourceRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SourceRANNodeIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SourceToTargetAMFInformationRerouteExtIEs(value ngapType.SourceToTargetAMFInformationRerouteExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceToTargetAMFInformationRerouteExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceToTargetAMFInformationRerouteExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SourceToTargetAMFInformationRerouteExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SourceToTargetAMFInformationRerouteExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SupportedTAItemExtIEs(value ngapType.SupportedTAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDRATInformation:
		if value.ExtensionValue.Present != ngapType.SupportedTAItemExtIEsPresentRATInformation {
			return binData, bitEnd, errors.New("SupportedTAItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SupportedTAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAIExtIEs(value ngapType.TAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAIExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAIBroadcastEUTRAItemExtIEs(value ngapType.TAIBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAIBroadcastEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAIBroadcastNRItemExtIEs(value ngapType.TAIBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAIBroadcastNRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAICancelledEUTRAItemExtIEs(value ngapType.TAICancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAICancelledEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAICancelledNRItemExtIEs(value ngapType.TAICancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAICancelledNRItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAIListForInactiveItemExtIEs(value ngapType.TAIListForInactiveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAIListForInactiveItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TAIListForPagingItemExtIEs(value ngapType.TAIListForPagingItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TAIListForPagingItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TargeteNBIDExtIEs(value ngapType.TargeteNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TargeteNBIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs(value ngapType.TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TargetRANNodeIDExtIEs(value ngapType.TargetRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TargetRANNodeIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TargetRNCIDExtIEs(value ngapType.TargetRNCIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRNCIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRNCIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TargetRNCIDExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRNCIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TNLAssociationItemExtIEs(value ngapType.TNLAssociationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TNLAssociationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TraceActivationExtIEs(value ngapType.TraceActivationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TraceActivationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEAggregateMaximumBitRateExtIEs(value ngapType.UEAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEAggregateMaximumBitRateExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEAssociatedLogicalNGConnectionItemExtIEs(value ngapType.UEAssociatedLogicalNGConnectionItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UENGAPIDPairExtIEs(value ngapType.UENGAPIDPairExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UENGAPIDPairExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEPresenceInAreaOfInterestItemExtIEs(value ngapType.UEPresenceInAreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEPresenceInAreaOfInterestItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UERadioCapabilityForPagingExtIEs(value ngapType.UERadioCapabilityForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityForPagingExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UESecurityCapabilitiesExtIEs(value ngapType.UESecurityCapabilitiesExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UESecurityCapabilitiesExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ULNGUUPTNLModifyItemExtIEs(value ngapType.ULNGUUPTNLModifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ULNGUUPTNLModifyItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UnavailableGUAMIItemExtIEs(value ngapType.UnavailableGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UnavailableGUAMIItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UPTransportLayerInformationItemExtIEs(value ngapType.UPTransportLayerInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UPTransportLayerInformationItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UPTransportLayerInformationPairItemExtIEs(value ngapType.UPTransportLayerInformationPairItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UPTransportLayerInformationPairItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UserLocationInformationEUTRAExtIEs(value ngapType.UserLocationInformationEUTRAExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDPSCellInformation:
		if value.ExtensionValue.Present != ngapType.UserLocationInformationEUTRAExtIEsPresentPSCellInformation {
			return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UserLocationInformationEUTRAExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UserLocationInformationN3IWFExtIEs(value ngapType.UserLocationInformationN3IWFExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UserLocationInformationN3IWFExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UserLocationInformationNRExtIEs(value ngapType.UserLocationInformationNRExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDPSCellInformation:
		if value.ExtensionValue.Present != ngapType.UserLocationInformationNRExtIEsPresentPSCellInformation {
			return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UserLocationInformationNRExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UserPlaneSecurityInformationExtIEs(value ngapType.UserPlaneSecurityInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserPlaneSecurityInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserPlaneSecurityInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UserPlaneSecurityInformationExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UserPlaneSecurityInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func VolumeTimedReportItemExtIEs(value ngapType.VolumeTimedReportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("VolumeTimedReportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("VolumeTimedReportItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = VolumeTimedReportItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("VolumeTimedReportItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func XnExtTLAItemExtIEs(value ngapType.XnExtTLAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolExtensionID.Value {
	case ngapType.ProtocolIEIDSCTPTLAs:
		if value.ExtensionValue.Present != ngapType.XnExtTLAItemExtIEsPresentSCTPTLAs {
			return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: ProtocolExtensionID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: ProtocolExtensionID: INVALID")
	}
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = XnExtTLAItemExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func XnTNLConfigurationInfoExtIEs(value ngapType.XnTNLConfigurationInfoExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = XnTNLConfigurationInfoExtIEsExtensionValue(value.ExtensionValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}
