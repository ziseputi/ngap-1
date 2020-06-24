package ngapEncode

import (
	"errors"
	"strconv"

	"ngap/ngapType"
)

func ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs(value ngapType.ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AdditionalDLUPTNLInformationForHOItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs(value ngapType.ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AllocationAndRetentionPriorityExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAllowedNSSAIItemExtIEs(value ngapType.ProtocolExtensionContainerAllowedNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAllowedNSSAIItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AllowedNSSAIItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAllowedNSSAIItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs(value ngapType.ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFTNLAssociationSetupItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs(value ngapType.ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFTNLAssociationToAddItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs(value ngapType.ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFTNLAssociationToRemoveItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs(value ngapType.ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFTNLAssociationToUpdateItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestExtIEs(value ngapType.ProtocolExtensionContainerAreaOfInterestExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AreaOfInterestExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestCellItemExtIEs(value ngapType.ProtocolExtensionContainerAreaOfInterestCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestCellItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AreaOfInterestCellItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestCellItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestItemExtIEs(value ngapType.ProtocolExtensionContainerAreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AreaOfInterestItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs(value ngapType.ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AreaOfInterestRANNodeItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs(value ngapType.ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AreaOfInterestTAIItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAssistanceDataForPagingExtIEs(value ngapType.ProtocolExtensionContainerAssistanceDataForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForPagingExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AssistanceDataForPagingExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForPagingExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs(value ngapType.ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AssistanceDataForRecommendedCellsExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAssociatedQosFlowItemExtIEs(value ngapType.ProtocolExtensionContainerAssociatedQosFlowItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAssociatedQosFlowItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AssociatedQosFlowItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAssociatedQosFlowItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerBroadcastPLMNItemExtIEs(value ngapType.ProtocolExtensionContainerBroadcastPLMNItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerBroadcastPLMNItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = BroadcastPLMNItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerBroadcastPLMNItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CancelledCellsInEAIEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs(value ngapType.ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CancelledCellsInEAINRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CancelledCellsInTAIEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs(value ngapType.ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CancelledCellsInTAINRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellIDBroadcastEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs(value ngapType.ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellIDBroadcastNRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellIDCancelledEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellIDCancelledNRItemExtIEs(value ngapType.ProtocolExtensionContainerCellIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledNRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellIDCancelledNRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledNRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellTypeExtIEs(value ngapType.ProtocolExtensionContainerCellTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellTypeExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellTypeExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellTypeExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCNAssistedRANTuningExtIEs(value ngapType.ProtocolExtensionContainerCNAssistedRANTuningExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCNAssistedRANTuningExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CNAssistedRANTuningExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCNAssistedRANTuningExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs(value ngapType.ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CNTypeRestrictionsForEquivalentItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CompletedCellsInEAIEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs(value ngapType.ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CompletedCellsInEAINRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CompletedCellsInTAIEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs(value ngapType.ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CompletedCellsInTAINRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs(value ngapType.ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CoreNetworkAssistanceInformationForInactiveExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs(value ngapType.ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = COUNTValueForPDCPSN12ExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs(value ngapType.ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = COUNTValueForPDCPSN18ExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCriticalityDiagnosticsExtIEs(value ngapType.ProtocolExtensionContainerCriticalityDiagnosticsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CriticalityDiagnosticsExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs(value ngapType.ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CriticalityDiagnosticsIEItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs(value ngapType.ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DataForwardingResponseDRBItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs(value ngapType.ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DataForwardingResponseERABListItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs(value ngapType.ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBsSubjectToStatusTransferItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusDL12ExtIEs(value ngapType.ProtocolExtensionContainerDRBStatusDL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL12ExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBStatusDL12ExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL12ExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusDL18ExtIEs(value ngapType.ProtocolExtensionContainerDRBStatusDL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL18ExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBStatusDL18ExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL18ExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusUL12ExtIEs(value ngapType.ProtocolExtensionContainerDRBStatusUL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL12ExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBStatusUL12ExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL12ExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusUL18ExtIEs(value ngapType.ProtocolExtensionContainerDRBStatusUL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL18ExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBStatusUL18ExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL18ExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs(value ngapType.ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBsToQosFlowsMappingItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDynamic5QIDescriptorExtIEs(value ngapType.ProtocolExtensionContainerDynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = Dynamic5QIDescriptorExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDynamic5QIDescriptorExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EmergencyAreaIDBroadcastEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs(value ngapType.ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EmergencyAreaIDBroadcastNRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EmergencyAreaIDCancelledEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs(value ngapType.ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EmergencyAreaIDCancelledNRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs(value ngapType.ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EmergencyFallbackIndicatorExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs(value ngapType.ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EndpointIPAddressAndPortExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEPSTAIExtIEs(value ngapType.ProtocolExtensionContainerEPSTAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEPSTAIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EPSTAIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEPSTAIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerERABInformationItemExtIEs(value ngapType.ProtocolExtensionContainerERABInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerERABInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ERABInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerERABInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEUTRACGIExtIEs(value ngapType.ProtocolExtensionContainerEUTRACGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEUTRACGIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = EUTRACGIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEUTRACGIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs(value ngapType.ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ExpectedUEActivityBehaviourExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerExpectedUEBehaviourExtIEs(value ngapType.ProtocolExtensionContainerExpectedUEBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEBehaviourExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ExpectedUEBehaviourExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEBehaviourExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs(value ngapType.ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ExpectedUEMovingTrajectoryItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs(value ngapType.ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ExtendedRATRestrictionInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerFiveGSTMSIExtIEs(value ngapType.ProtocolExtensionContainerFiveGSTMSIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerFiveGSTMSIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = FiveGSTMSIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerFiveGSTMSIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs(value ngapType.ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ForbiddenAreaInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGBRQosInformationExtIEs(value ngapType.ProtocolExtensionContainerGBRQosInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGBRQosInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GBRQosInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGBRQosInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGlobalGNBIDExtIEs(value ngapType.ProtocolExtensionContainerGlobalGNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalGNBIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GlobalGNBIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalGNBIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGlobalN3IWFIDExtIEs(value ngapType.ProtocolExtensionContainerGlobalN3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalN3IWFIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GlobalN3IWFIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalN3IWFIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGlobalNgENBIDExtIEs(value ngapType.ProtocolExtensionContainerGlobalNgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalNgENBIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GlobalNgENBIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalNgENBIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGTPTunnelExtIEs(value ngapType.ProtocolExtensionContainerGTPTunnelExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGTPTunnelExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GTPTunnelExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGTPTunnelExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGUAMIExtIEs(value ngapType.ProtocolExtensionContainerGUAMIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGUAMIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GUAMIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGUAMIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerHandoverCommandTransferExtIEs(value ngapType.ProtocolExtensionContainerHandoverCommandTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverCommandTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverCommandTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverCommandTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs(value ngapType.ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverPreparationUnsuccessfulTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs(value ngapType.ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverRequestAcknowledgeTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerHandoverRequiredTransferExtIEs(value ngapType.ProtocolExtensionContainerHandoverRequiredTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverRequiredTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverRequiredTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverRequiredTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs(value ngapType.ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverResourceAllocationUnsuccessfulTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs(value ngapType.ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPagingExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerLAIExtIEs(value ngapType.ProtocolExtensionContainerLAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerLAIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LAIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerLAIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerLastVisitedCellItemExtIEs(value ngapType.ProtocolExtensionContainerLastVisitedCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerLastVisitedCellItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LastVisitedCellItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerLastVisitedCellItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs(value ngapType.ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LastVisitedNGRANCellInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerLocationReportingRequestTypeExtIEs(value ngapType.ProtocolExtensionContainerLocationReportingRequestTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerLocationReportingRequestTypeExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LocationReportingRequestTypeExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerLocationReportingRequestTypeExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerMobilityRestrictionListExtIEs(value ngapType.ProtocolExtensionContainerMobilityRestrictionListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerMobilityRestrictionListExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = MobilityRestrictionListExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerMobilityRestrictionListExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs(value ngapType.ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGRANTNLAssociationToRemoveItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs(value ngapType.ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NonDynamic5QIDescriptorExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerNRCGIExtIEs(value ngapType.ProtocolExtensionContainerNRCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerNRCGIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NRCGIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerNRCGIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs(value ngapType.ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = OverloadStartNSSAIItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPacketErrorRateExtIEs(value ngapType.ProtocolExtensionContainerPacketErrorRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPacketErrorRateExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PacketErrorRateExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPacketErrorRateExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPagingAttemptInformationExtIEs(value ngapType.ProtocolExtensionContainerPagingAttemptInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPagingAttemptInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PagingAttemptInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPagingAttemptInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs(value ngapType.ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestAcknowledgeTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs(value ngapType.ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestSetupFailedTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPathSwitchRequestTransferExtIEs(value ngapType.ProtocolExtensionContainerPathSwitchRequestTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs(value ngapType.ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestUnsuccessfulTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionAggregateMaximumBitRateExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceAdmittedItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToModifyItemModCfmExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToModifyItemModResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtFailExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToSetupItemHOAckExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToSetupItemPSReqExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceFailedToSetupItemSUResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceHandoverItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceItemCxtRelCplExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceItemCxtRelReqExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceItemHORqdExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyConfirmTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyRequestTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyResponseTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyIndicationTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyItemModCfmExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyItemModIndExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyItemModReqExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyItemModResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyUnsuccessfulTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceNotifyItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceNotifyReleasedTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceNotifyTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleaseCommandTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleasedItemNotExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleasedItemPSAckExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleasedItemPSFailExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleasedItemRelResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleaseResponseTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSecondaryRATUsageItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupItemCxtReqExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupItemCxtResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupItemHOReqExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupItemSUReqExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupItemSUResExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs(value ngapType.ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupRequestTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupResponseTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupUnsuccessfulTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSwitchedItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceToBeSwitchedDLItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceToReleaseItemHOCmdExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceToReleaseItemRelCmdExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPDUSessionUsageReportExtIEs(value ngapType.ProtocolExtensionContainerPDUSessionUsageReportExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionUsageReportExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionUsageReportExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionUsageReportExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPLMNSupportItemExtIEs(value ngapType.ProtocolExtensionContainerPLMNSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPLMNSupportItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PLMNSupportItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPLMNSupportItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowAcceptedItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowAcceptedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAcceptedItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowAcceptedItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAcceptedItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowAddOrModifyRequestItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowAddOrModifyResponseItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowInformationItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs(value ngapType.ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowLevelQosParametersExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowWithCauseItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowWithCauseItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowWithCauseItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowWithCauseItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowWithCauseItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowModifyConfirmItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowNotifyItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowNotifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowNotifyItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowNotifyItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowNotifyItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs(value ngapType.ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowPerTNLInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowPerTNLInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowSetupRequestItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs(value ngapType.ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowItemWithDataForwardingExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs(value ngapType.ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosFlowToBeForwardedItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs(value ngapType.ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QoSFlowsUsageReportItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs(value ngapType.ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RANStatusTransferTransparentContainerExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRATRestrictionsItemExtIEs(value ngapType.ProtocolExtensionContainerRATRestrictionsItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRATRestrictionsItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RATRestrictionsItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRATRestrictionsItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedCellsForPagingExtIEs(value ngapType.ProtocolExtensionContainerRecommendedCellsForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellsForPagingExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RecommendedCellsForPagingExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellsForPagingExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedCellItemExtIEs(value ngapType.ProtocolExtensionContainerRecommendedCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RecommendedCellItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs(value ngapType.ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RecommendedRANNodesForPagingExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedRANNodeItemExtIEs(value ngapType.ProtocolExtensionContainerRecommendedRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodeItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RecommendedRANNodeItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodeItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRIMInformationTransferExtIEs(value ngapType.ProtocolExtensionContainerRIMInformationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRIMInformationTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RIMInformationTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRIMInformationTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs(value ngapType.ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SecondaryRATUsageInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs(value ngapType.ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SecondaryRATDataUsageReportTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSecurityContextExtIEs(value ngapType.ProtocolExtensionContainerSecurityContextExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityContextExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SecurityContextExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityContextExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSecurityIndicationExtIEs(value ngapType.ProtocolExtensionContainerSecurityIndicationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityIndicationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SecurityIndicationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityIndicationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSecurityResultExtIEs(value ngapType.ProtocolExtensionContainerSecurityResultExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityResultExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SecurityResultExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityResultExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerServedGUAMIItemExtIEs(value ngapType.ProtocolExtensionContainerServedGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerServedGUAMIItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ServedGUAMIItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerServedGUAMIItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerServiceAreaInformationItemExtIEs(value ngapType.ProtocolExtensionContainerServiceAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerServiceAreaInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ServiceAreaInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerServiceAreaInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSliceOverloadItemExtIEs(value ngapType.ProtocolExtensionContainerSliceOverloadItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceOverloadItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SliceOverloadItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceOverloadItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSliceSupportItemExtIEs(value ngapType.ProtocolExtensionContainerSliceSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceSupportItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SliceSupportItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceSupportItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSNSSAIExtIEs(value ngapType.ProtocolExtensionContainerSNSSAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSNSSAIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SNSSAIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSNSSAIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSONConfigurationTransferExtIEs(value ngapType.ProtocolExtensionContainerSONConfigurationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSONConfigurationTransferExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SONConfigurationTransferExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSONConfigurationTransferExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSONInformationReplyExtIEs(value ngapType.ProtocolExtensionContainerSONInformationReplyExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSONInformationReplyExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SONInformationReplyExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSONInformationReplyExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs(value ngapType.ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSourceRANNodeIDExtIEs(value ngapType.ProtocolExtensionContainerSourceRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceRANNodeIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SourceRANNodeIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceRANNodeIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs(value ngapType.ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SourceToTargetAMFInformationRerouteExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSupportedTAItemExtIEs(value ngapType.ProtocolExtensionContainerSupportedTAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSupportedTAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SupportedTAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSupportedTAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIExtIEs(value ngapType.ProtocolExtensionContainerTAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAIBroadcastEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIBroadcastNRItemExtIEs(value ngapType.ProtocolExtensionContainerTAIBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastNRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAIBroadcastNRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastNRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs(value ngapType.ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAICancelledEUTRAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAICancelledNRItemExtIEs(value ngapType.ProtocolExtensionContainerTAICancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledNRItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAICancelledNRItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledNRItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIListForInactiveItemExtIEs(value ngapType.ProtocolExtensionContainerTAIListForInactiveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForInactiveItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAIListForInactiveItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForInactiveItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIListForPagingItemExtIEs(value ngapType.ProtocolExtensionContainerTAIListForPagingItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForPagingItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TAIListForPagingItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForPagingItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTargeteNBIDExtIEs(value ngapType.ProtocolExtensionContainerTargeteNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTargeteNBIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TargeteNBIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTargeteNBIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs(value ngapType.ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTargetRANNodeIDExtIEs(value ngapType.ProtocolExtensionContainerTargetRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetRANNodeIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TargetRANNodeIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetRANNodeIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTargetRNCIDExtIEs(value ngapType.ProtocolExtensionContainerTargetRNCIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetRNCIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TargetRNCIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetRNCIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTNLAssociationItemExtIEs(value ngapType.ProtocolExtensionContainerTNLAssociationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTNLAssociationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TNLAssociationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTNLAssociationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTraceActivationExtIEs(value ngapType.ProtocolExtensionContainerTraceActivationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTraceActivationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TraceActivationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTraceActivationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs(value ngapType.ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEAggregateMaximumBitRateExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs(value ngapType.ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEAssociatedLogicalNGConnectionItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUENGAPIDPairExtIEs(value ngapType.ProtocolExtensionContainerUENGAPIDPairExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUENGAPIDPairExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UENGAPIDPairExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUENGAPIDPairExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs(value ngapType.ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEPresenceInAreaOfInterestItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs(value ngapType.ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UERadioCapabilityForPagingExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUESecurityCapabilitiesExtIEs(value ngapType.ProtocolExtensionContainerUESecurityCapabilitiesExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUESecurityCapabilitiesExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UESecurityCapabilitiesExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUESecurityCapabilitiesExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs(value ngapType.ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ULNGUUPTNLModifyItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUnavailableGUAMIItemExtIEs(value ngapType.ProtocolExtensionContainerUnavailableGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUnavailableGUAMIItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UnavailableGUAMIItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUnavailableGUAMIItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs(value ngapType.ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UPTransportLayerInformationItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs(value ngapType.ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UPTransportLayerInformationPairItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs(value ngapType.ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UserLocationInformationEUTRAExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs(value ngapType.ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UserLocationInformationN3IWFExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUserLocationInformationNRExtIEs(value ngapType.ProtocolExtensionContainerUserLocationInformationNRExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationNRExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UserLocationInformationNRExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationNRExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs(value ngapType.ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UserPlaneSecurityInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerVolumeTimedReportItemExtIEs(value ngapType.ProtocolExtensionContainerVolumeTimedReportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerVolumeTimedReportItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = VolumeTimedReportItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerVolumeTimedReportItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerXnExtTLAItemExtIEs(value ngapType.ProtocolExtensionContainerXnExtTLAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerXnExtTLAItemExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = XnExtTLAItemExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerXnExtTLAItemExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs(value ngapType.ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = XnTNLConfigurationInfoExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}
