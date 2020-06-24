package ngapEncode

import (
	"errors"
	"strconv"

	"ngap/ngapType"
)

func AdditionalDLUPTNLInformationForHOList (value ngapType.AdditionalDLUPTNLInformationForHOList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofMultiConnectivityMinusOne {
		return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOList: List: valueExt,sizeLB:1,sizeUB:maxnoofMultiConnectivityMinusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofMultiConnectivityMinusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AdditionalDLUPTNLInformationForHOItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AdditionalDLUPTNLInformationForHOList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllowedNSSAI (value ngapType.AllowedNSSAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofAllowedSNSSAIs {
		return binData, bitEnd, errors.New("AllowedNSSAI: List: valueExt,sizeLB:1,sizeUB:maxnoofAllowedS-NSSAIs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofAllowedSNSSAIs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AllowedNSSAIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AllowedNSSAI: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllowedTACs (value ngapType.AllowedTACs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofAllowedAreas {
		return binData, bitEnd, errors.New("AllowedTACs: List: valueExt,sizeLB:1,sizeUB:maxnoofAllowedAreas")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofAllowedAreas, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAC(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AllowedTACs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationSetupList (value ngapType.AMFTNLAssociationSetupList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTNLAssociations {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupList: List: valueExt,sizeLB:1,sizeUB:maxnoofTNLAssociations")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTNLAssociations, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationSetupItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationSetupList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToAddList (value ngapType.AMFTNLAssociationToAddList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTNLAssociations {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddList: List: valueExt,sizeLB:1,sizeUB:maxnoofTNLAssociations")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTNLAssociations, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToAddItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToAddList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToRemoveList (value ngapType.AMFTNLAssociationToRemoveList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTNLAssociations {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveList: List: valueExt,sizeLB:1,sizeUB:maxnoofTNLAssociations")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTNLAssociations, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToRemoveItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AMFTNLAssociationToUpdateList (value ngapType.AMFTNLAssociationToUpdateList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTNLAssociations {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateList: List: valueExt,sizeLB:1,sizeUB:maxnoofTNLAssociations")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTNLAssociations, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToUpdateItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestCellList (value ngapType.AreaOfInterestCellList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinAoI {
		return binData, bitEnd, errors.New("AreaOfInterestCellList: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinAoI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinAoI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestCellItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestCellList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestList (value ngapType.AreaOfInterestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofAoI {
		return binData, bitEnd, errors.New("AreaOfInterestList: List: valueExt,sizeLB:1,sizeUB:maxnoofAoI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofAoI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestRANNodeList (value ngapType.AreaOfInterestRANNodeList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofRANNodeinAoI {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeList: List: valueExt,sizeLB:1,sizeUB:maxnoofRANNodeinAoI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofRANNodeinAoI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestRANNodeItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestRANNodeList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestTAIList (value ngapType.AreaOfInterestTAIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIinAoI {
		return binData, bitEnd, errors.New("AreaOfInterestTAIList: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIinAoI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIinAoI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestTAIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestTAIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AssociatedQosFlowList (value ngapType.AssociatedQosFlowList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("AssociatedQosFlowList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AssociatedQosFlowItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AssociatedQosFlowList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func BroadcastPLMNList (value ngapType.BroadcastPLMNList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofBPLMNs {
		return binData, bitEnd, errors.New("BroadcastPLMNList: List: valueExt,sizeLB:1,sizeUB:maxnoofBPLMNs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofBPLMNs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = BroadcastPLMNItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastPLMNList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAIEUTRA (value ngapType.CancelledCellsInEAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinEAI {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinEAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinEAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInEAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINR (value ngapType.CancelledCellsInEAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinEAI {
		return binData, bitEnd, errors.New("CancelledCellsInEAINR: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinEAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinEAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInEAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInEAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAIEUTRA (value ngapType.CancelledCellsInTAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinTAI {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinTAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinTAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInTAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINR (value ngapType.CancelledCellsInTAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinTAI {
		return binData, bitEnd, errors.New("CancelledCellsInTAINR: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinTAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinTAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInTAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInTAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDBroadcastEUTRA (value ngapType.CellIDBroadcastEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellIDforWarning {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofCellIDforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellIDforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDBroadcastEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDBroadcastEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDBroadcastNR (value ngapType.CellIDBroadcastNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellIDforWarning {
		return binData, bitEnd, errors.New("CellIDBroadcastNR: List: valueExt,sizeLB:1,sizeUB:maxnoofCellIDforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellIDforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDBroadcastNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDBroadcastNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDCancelledEUTRA (value ngapType.CellIDCancelledEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellIDforWarning {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofCellIDforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellIDforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDCancelledEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDCancelledEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDCancelledNR (value ngapType.CellIDCancelledNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellIDforWarning {
		return binData, bitEnd, errors.New("CellIDCancelledNR: List: valueExt,sizeLB:1,sizeUB:maxnoofCellIDforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellIDforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDCancelledNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDCancelledNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CNTypeRestrictionsForEquivalent (value ngapType.CNTypeRestrictionsForEquivalent, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEPLMNs {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalent: List: valueExt,sizeLB:1,sizeUB:maxnoofEPLMNs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEPLMNs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CNTypeRestrictionsForEquivalentItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CNTypeRestrictionsForEquivalent: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRA (value ngapType.CompletedCellsInEAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinEAI {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinEAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinEAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInEAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINR (value ngapType.CompletedCellsInEAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinEAI {
		return binData, bitEnd, errors.New("CompletedCellsInEAINR: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinEAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinEAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInEAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInEAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRA (value ngapType.CompletedCellsInTAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinTAI {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinTAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinTAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInTAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINR (value ngapType.CompletedCellsInTAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellinTAI {
		return binData, bitEnd, errors.New("CompletedCellsInTAINR: List: valueExt,sizeLB:1,sizeUB:maxnoofCellinTAI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellinTAI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInTAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInTAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CriticalityDiagnosticsIEList (value ngapType.CriticalityDiagnosticsIEList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofErrors {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEList: List: valueExt,sizeLB:1,sizeUB:maxnoofErrors")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofErrors, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CriticalityDiagnosticsIEItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnosticsIEList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DataForwardingResponseDRBList (value ngapType.DataForwardingResponseDRBList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofDRBs {
		return binData, bitEnd, errors.New("DataForwardingResponseDRBList: List: valueExt,sizeLB:1,sizeUB:maxnoofDRBs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofDRBs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DataForwardingResponseDRBItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseDRBList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DataForwardingResponseERABList (value ngapType.DataForwardingResponseERABList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofERABs {
		return binData, bitEnd, errors.New("DataForwardingResponseERABList: List: valueExt,sizeLB:1,sizeUB:maxnoofE-RABs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofERABs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DataForwardingResponseERABListItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DataForwardingResponseERABList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBsSubjectToStatusTransferList (value ngapType.DRBsSubjectToStatusTransferList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofDRBs {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferList: List: valueExt,sizeLB:1,sizeUB:maxnoofDRBs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofDRBs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBsSubjectToStatusTransferItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBsToQosFlowsMappingList (value ngapType.DRBsToQosFlowsMappingList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofDRBs {
		return binData, bitEnd, errors.New("DRBsToQosFlowsMappingList: List: valueExt,sizeLB:1,sizeUB:maxnoofDRBs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofDRBs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBsToQosFlowsMappingItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsToQosFlowsMappingList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastEUTRA (value ngapType.EmergencyAreaIDBroadcastEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEmergencyAreaID {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofEmergencyAreaID")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEmergencyAreaID, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDBroadcastEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastNR (value ngapType.EmergencyAreaIDBroadcastNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEmergencyAreaID {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNR: List: valueExt,sizeLB:1,sizeUB:maxnoofEmergencyAreaID")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEmergencyAreaID, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDBroadcastNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledEUTRA (value ngapType.EmergencyAreaIDCancelledEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEmergencyAreaID {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofEmergencyAreaID")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEmergencyAreaID, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDCancelledEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledNR (value ngapType.EmergencyAreaIDCancelledNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEmergencyAreaID {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNR: List: valueExt,sizeLB:1,sizeUB:maxnoofEmergencyAreaID")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEmergencyAreaID, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDCancelledNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDList (value ngapType.EmergencyAreaIDList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEmergencyAreaID {
		return binData, bitEnd, errors.New("EmergencyAreaIDList: List: valueExt,sizeLB:1,sizeUB:maxnoofEmergencyAreaID")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEmergencyAreaID, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaID(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDListForRestart (value ngapType.EmergencyAreaIDListForRestart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEAIforRestart {
		return binData, bitEnd, errors.New("EmergencyAreaIDListForRestart: List: valueExt,sizeLB:1,sizeUB:maxnoofEAIforRestart")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEAIforRestart, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaID(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDListForRestart: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EquivalentPLMNs (value ngapType.EquivalentPLMNs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEPLMNs {
		return binData, bitEnd, errors.New("EquivalentPLMNs: List: valueExt,sizeLB:1,sizeUB:maxnoofEPLMNs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEPLMNs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PLMNIdentity(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EquivalentPLMNs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ERABInformationList (value ngapType.ERABInformationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofERABs {
		return binData, bitEnd, errors.New("ERABInformationList: List: valueExt,sizeLB:1,sizeUB:maxnoofE-RABs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofERABs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ERABInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ERABInformationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EUTRACGIList (value ngapType.EUTRACGIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellsinngeNB {
		return binData, bitEnd, errors.New("EUTRACGIList: List: valueExt,sizeLB:1,sizeUB:maxnoofCellsinngeNB")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellsinngeNB, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EUTRACGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EUTRACGIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EUTRACGIListForWarning (value ngapType.EUTRACGIListForWarning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellIDforWarning {
		return binData, bitEnd, errors.New("EUTRACGIListForWarning: List: valueExt,sizeLB:1,sizeUB:maxnoofCellIDforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellIDforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EUTRACGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EUTRACGIListForWarning: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedUEMovingTrajectory (value ngapType.ExpectedUEMovingTrajectory, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellsUEMovingTrajectory {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectory: List: valueExt,sizeLB:1,sizeUB:maxnoofCellsUEMovingTrajectory")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellsUEMovingTrajectory, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ExpectedUEMovingTrajectoryItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEMovingTrajectory: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ForbiddenAreaInformation (value ngapType.ForbiddenAreaInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEPLMNsPlusOne {
		return binData, bitEnd, errors.New("ForbiddenAreaInformation: List: valueExt,sizeLB:1,sizeUB:maxnoofEPLMNsPlusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEPLMNsPlusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ForbiddenAreaInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ForbiddenAreaInformation: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ForbiddenTACs (value ngapType.ForbiddenTACs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofForbTACs {
		return binData, bitEnd, errors.New("ForbiddenTACs: List: valueExt,sizeLB:1,sizeUB:maxnoofForbTACs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofForbTACs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAC(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ForbiddenTACs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NGRANTNLAssociationToRemoveList (value ngapType.NGRANTNLAssociationToRemoveList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTNLAssociations {
		return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveList: List: valueExt,sizeLB:1,sizeUB:maxnoofTNLAssociations")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTNLAssociations, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NGRANTNLAssociationToRemoveItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGRANTNLAssociationToRemoveList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NotAllowedTACs (value ngapType.NotAllowedTACs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofAllowedAreas {
		return binData, bitEnd, errors.New("NotAllowedTACs: List: valueExt,sizeLB:1,sizeUB:maxnoofAllowedAreas")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofAllowedAreas, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAC(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NotAllowedTACs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NRCGIList (value ngapType.NRCGIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellsingNB {
		return binData, bitEnd, errors.New("NRCGIList: List: valueExt,sizeLB:1,sizeUB:maxnoofCellsingNB")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellsingNB, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NRCGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NRCGIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NRCGIListForWarning (value ngapType.NRCGIListForWarning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellIDforWarning {
		return binData, bitEnd, errors.New("NRCGIListForWarning: List: valueExt,sizeLB:1,sizeUB:maxnoofCellIDforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellIDforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NRCGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NRCGIListForWarning: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func OverloadStartNSSAIList (value ngapType.OverloadStartNSSAIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofSliceItems {
		return binData, bitEnd, errors.New("OverloadStartNSSAIList: List: valueExt,sizeLB:1,sizeUB:maxnoofSliceItems")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofSliceItems, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = OverloadStartNSSAIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartNSSAIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceAdmittedList (value ngapType.PDUSessionResourceAdmittedList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceAdmittedItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceAdmittedList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToModifyListModCfm (value ngapType.PDUSessionResourceFailedToModifyListModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModCfm: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModCfm(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModCfm: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToModifyListModRes (value ngapType.PDUSessionResourceFailedToModifyListModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModRes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupListCxtFail (value ngapType.PDUSessionResourceFailedToSetupListCxtFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtFail: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtFail(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtFail: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupListCxtRes (value ngapType.PDUSessionResourceFailedToSetupListCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtRes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupListHOAck (value ngapType.PDUSessionResourceFailedToSetupListHOAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListHOAck: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemHOAck(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListHOAck: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupListPSReq (value ngapType.PDUSessionResourceFailedToSetupListPSReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListPSReq: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemPSReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListPSReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceFailedToSetupListSURes (value ngapType.PDUSessionResourceFailedToSetupListSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListSURes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemSURes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListSURes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceHandoverList (value ngapType.PDUSessionResourceHandoverList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceHandoverItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceHandoverList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceInformationList (value ngapType.PDUSessionResourceInformationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceInformationList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceInformationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceListCxtRelCpl (value ngapType.PDUSessionResourceListCxtRelCpl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelCpl: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemCxtRelCpl(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelCpl: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceListCxtRelReq (value ngapType.PDUSessionResourceListCxtRelReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelReq: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemCxtRelReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceListHORqd (value ngapType.PDUSessionResourceListHORqd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceListHORqd: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemHORqd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceListHORqd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyListModCfm (value ngapType.PDUSessionResourceModifyListModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModCfm: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModCfm(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModCfm: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyListModInd (value ngapType.PDUSessionResourceModifyListModInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModInd: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModInd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModInd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyListModReq (value ngapType.PDUSessionResourceModifyListModReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModReq: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyListModRes (value ngapType.PDUSessionResourceModifyListModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModRes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyList (value ngapType.PDUSessionResourceNotifyList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceNotifyItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedListNot (value ngapType.PDUSessionResourceReleasedListNot, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListNot: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemNot(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListNot: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedListPSAck (value ngapType.PDUSessionResourceReleasedListPSAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSAck: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemPSAck(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSAck: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedListPSFail (value ngapType.PDUSessionResourceReleasedListPSFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSFail: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemPSFail(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSFail: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleasedListRelRes (value ngapType.PDUSessionResourceReleasedListRelRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListRelRes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemRelRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListRelRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSecondaryRATUsageList (value ngapType.PDUSessionResourceSecondaryRATUsageList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSecondaryRATUsageItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSecondaryRATUsageList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupListCxtReq (value ngapType.PDUSessionResourceSetupListCxtReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtReq: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemCxtReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupListCxtRes (value ngapType.PDUSessionResourceSetupListCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtRes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemCxtRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupListHOReq (value ngapType.PDUSessionResourceSetupListHOReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListHOReq: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemHOReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListHOReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupListSUReq (value ngapType.PDUSessionResourceSetupListSUReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListSUReq: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemSUReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListSUReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupListSURes (value ngapType.PDUSessionResourceSetupListSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListSURes: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemSURes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListSURes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSwitchedList (value ngapType.PDUSessionResourceSwitchedList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSwitchedItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSwitchedList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToBeSwitchedDLList (value ngapType.PDUSessionResourceToBeSwitchedDLList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLList: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToReleaseListHOCmd (value ngapType.PDUSessionResourceToReleaseListHOCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListHOCmd: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToReleaseItemHOCmd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListHOCmd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PDUSessionResourceToReleaseListRelCmd (value ngapType.PDUSessionResourceToReleaseListRelCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPDUSessions {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListRelCmd: List: valueExt,sizeLB:1,sizeUB:maxnoofPDUSessions")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPDUSessions, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToReleaseItemRelCmd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListRelCmd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PLMNSupportList (value ngapType.PLMNSupportList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofPLMNs {
		return binData, bitEnd, errors.New("PLMNSupportList: List: valueExt,sizeLB:1,sizeUB:maxnoofPLMNs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofPLMNs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PLMNSupportItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PLMNSupportList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAcceptedList (value ngapType.QosFlowAcceptedList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowAcceptedList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowAcceptedItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAcceptedList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyRequestList (value ngapType.QosFlowAddOrModifyRequestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowAddOrModifyRequestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyResponseList (value ngapType.QosFlowAddOrModifyResponseList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowAddOrModifyResponseItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyResponseList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowInformationList (value ngapType.QosFlowInformationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowInformationList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowInformationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowListWithCause (value ngapType.QosFlowListWithCause, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowListWithCause: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowWithCauseItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowListWithCause: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowModifyConfirmList (value ngapType.QosFlowModifyConfirmList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowModifyConfirmList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowModifyConfirmItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowModifyConfirmList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowNotifyList (value ngapType.QosFlowNotifyList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowNotifyList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowNotifyItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowNotifyList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowPerTNLInformationList (value ngapType.QosFlowPerTNLInformationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofMultiConnectivityMinusOne {
		return binData, bitEnd, errors.New("QosFlowPerTNLInformationList: List: valueExt,sizeLB:1,sizeUB:maxnoofMultiConnectivityMinusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofMultiConnectivityMinusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowPerTNLInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowPerTNLInformationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowSetupRequestList (value ngapType.QosFlowSetupRequestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowSetupRequestList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowSetupRequestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowSetupRequestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowListWithDataForwarding (value ngapType.QosFlowListWithDataForwarding, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowListWithDataForwarding: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowItemWithDataForwarding(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowListWithDataForwarding: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowToBeForwardedList (value ngapType.QosFlowToBeForwardedList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QosFlowToBeForwardedList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowToBeForwardedItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowToBeForwardedList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QoSFlowsUsageReportList (value ngapType.QoSFlowsUsageReportList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofQosFlows {
		return binData, bitEnd, errors.New("QoSFlowsUsageReportList: List: valueExt,sizeLB:1,sizeUB:maxnoofQosFlows")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofQosFlows, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QoSFlowsUsageReportItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QoSFlowsUsageReportList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RATRestrictions (value ngapType.RATRestrictions, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEPLMNsPlusOne {
		return binData, bitEnd, errors.New("RATRestrictions: List: valueExt,sizeLB:1,sizeUB:maxnoofEPLMNsPlusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEPLMNsPlusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RATRestrictionsItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RATRestrictions: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedCellList (value ngapType.RecommendedCellList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofRecommendedCells {
		return binData, bitEnd, errors.New("RecommendedCellList: List: valueExt,sizeLB:1,sizeUB:maxnoofRecommendedCells")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofRecommendedCells, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedCellItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedCellList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedRANNodeList (value ngapType.RecommendedRANNodeList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofRecommendedRANNodes {
		return binData, bitEnd, errors.New("RecommendedRANNodeList: List: valueExt,sizeLB:1,sizeUB:maxnoofRecommendedRANNodes")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofRecommendedRANNodes, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedRANNodeItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedRANNodeList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SCTPTLAs (value ngapType.SCTPTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofXnTLAs {
		return binData, bitEnd, errors.New("SCTPTLAs: List: valueExt,sizeLB:1,sizeUB:maxnoofXnTLAs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofXnTLAs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TransportLayerAddress(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SCTPTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ServedGUAMIList (value ngapType.ServedGUAMIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofServedGUAMIs {
		return binData, bitEnd, errors.New("ServedGUAMIList: List: valueExt,sizeLB:1,sizeUB:maxnoofServedGUAMIs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofServedGUAMIs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ServedGUAMIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServedGUAMIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ServiceAreaInformation (value ngapType.ServiceAreaInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofEPLMNsPlusOne {
		return binData, bitEnd, errors.New("ServiceAreaInformation: List: valueExt,sizeLB:1,sizeUB:maxnoofEPLMNsPlusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofEPLMNsPlusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ServiceAreaInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServiceAreaInformation: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SliceOverloadList (value ngapType.SliceOverloadList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofSliceItems {
		return binData, bitEnd, errors.New("SliceOverloadList: List: valueExt,sizeLB:1,sizeUB:maxnoofSliceItems")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofSliceItems, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SliceOverloadItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SliceOverloadList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SliceSupportList (value ngapType.SliceSupportList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofSliceItems {
		return binData, bitEnd, errors.New("SliceSupportList: List: valueExt,sizeLB:1,sizeUB:maxnoofSliceItems")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofSliceItems, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SliceSupportItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SliceSupportList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SupportedTAList (value ngapType.SupportedTAList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTACs {
		return binData, bitEnd, errors.New("SupportedTAList: List: valueExt,sizeLB:1,sizeUB:maxnoofTACs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTACs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SupportedTAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SupportedTAList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastEUTRA (value ngapType.TAIBroadcastEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforWarning {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIBroadcastEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIBroadcastEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastNR (value ngapType.TAIBroadcastNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforWarning {
		return binData, bitEnd, errors.New("TAIBroadcastNR: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIBroadcastNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIBroadcastNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledEUTRA (value ngapType.TAICancelledEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforWarning {
		return binData, bitEnd, errors.New("TAICancelledEUTRA: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAICancelledEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAICancelledEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledNR (value ngapType.TAICancelledNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforWarning {
		return binData, bitEnd, errors.New("TAICancelledNR: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAICancelledNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAICancelledNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForInactive (value ngapType.TAIListForInactive, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforInactive {
		return binData, bitEnd, errors.New("TAIListForInactive: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforInactive")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforInactive, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIListForInactiveItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForInactive: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForPaging (value ngapType.TAIListForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforPaging {
		return binData, bitEnd, errors.New("TAIListForPaging: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforPaging")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforPaging, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIListForPagingItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForPaging: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForRestart (value ngapType.TAIListForRestart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforRestart {
		return binData, bitEnd, errors.New("TAIListForRestart: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforRestart")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforRestart, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForRestart: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForWarning (value ngapType.TAIListForWarning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTAIforWarning {
		return binData, bitEnd, errors.New("TAIListForWarning: List: valueExt,sizeLB:1,sizeUB:maxnoofTAIforWarning")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTAIforWarning, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForWarning: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TNLAssociationList (value ngapType.TNLAssociationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTNLAssociations {
		return binData, bitEnd, errors.New("TNLAssociationList: List: valueExt,sizeLB:1,sizeUB:maxnoofTNLAssociations")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTNLAssociations, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TNLAssociationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TNLAssociationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UEAssociatedLogicalNGConnectionList (value ngapType.UEAssociatedLogicalNGConnectionList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofNGConnectionsToReset {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionList: List: valueExt,sizeLB:1,sizeUB:maxnoofNGConnectionsToReset")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofNGConnectionsToReset, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UEHistoryInformation (value ngapType.UEHistoryInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofCellsinUEHistoryInfo {
		return binData, bitEnd, errors.New("UEHistoryInformation: List: valueExt,sizeLB:1,sizeUB:maxnoofCellsinUEHistoryInfo")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofCellsinUEHistoryInfo, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = LastVisitedCellItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEHistoryInformation: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UEPresenceInAreaOfInterestList (value ngapType.UEPresenceInAreaOfInterestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofAoI {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestList: List: valueExt,sizeLB:1,sizeUB:maxnoofAoI")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofAoI, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEPresenceInAreaOfInterestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ULNGUUPTNLModifyList (value ngapType.ULNGUUPTNLModifyList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofMultiConnectivity {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyList: List: valueExt,sizeLB:1,sizeUB:maxnoofMultiConnectivity")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofMultiConnectivity, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ULNGUUPTNLModifyItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ULNGUUPTNLModifyList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UnavailableGUAMIList (value ngapType.UnavailableGUAMIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofServedGUAMIs {
		return binData, bitEnd, errors.New("UnavailableGUAMIList: List: valueExt,sizeLB:1,sizeUB:maxnoofServedGUAMIs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofServedGUAMIs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UnavailableGUAMIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnavailableGUAMIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationList (value ngapType.UPTransportLayerInformationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofMultiConnectivityMinusOne {
		return binData, bitEnd, errors.New("UPTransportLayerInformationList: List: valueExt,sizeLB:1,sizeUB:maxnoofMultiConnectivityMinusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofMultiConnectivityMinusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UPTransportLayerInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationPairList (value ngapType.UPTransportLayerInformationPairList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofMultiConnectivityMinusOne {
		return binData, bitEnd, errors.New("UPTransportLayerInformationPairList: List: valueExt,sizeLB:1,sizeUB:maxnoofMultiConnectivityMinusOne")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofMultiConnectivityMinusOne, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UPTransportLayerInformationPairItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformationPairList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func VolumeTimedReportList (value ngapType.VolumeTimedReportList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofTimePeriods {
		return binData, bitEnd, errors.New("VolumeTimedReportList: List: valueExt,sizeLB:1,sizeUB:maxnoofTimePeriods")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofTimePeriods, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = VolumeTimedReportItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("VolumeTimedReportList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnExtTLAs (value ngapType.XnExtTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofXnExtTLAs {
		return binData, bitEnd, errors.New("XnExtTLAs: List: valueExt,sizeLB:1,sizeUB:maxnoofXnExtTLAs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofXnExtTLAs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = XnExtTLAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnExtTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnGTPTLAs (value ngapType.XnGTPTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofXnGTPTLAs {
		return binData, bitEnd, errors.New("XnGTPTLAs: List: valueExt,sizeLB:1,sizeUB:maxnoofXnGTP-TLAs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofXnGTPTLAs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TransportLayerAddress(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnGTPTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnTLAs (value ngapType.XnTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > ngapType.MaxnoofXnTLAs {
		return binData, bitEnd, errors.New("XnTLAs: List: valueExt,sizeLB:1,sizeUB:maxnoofXnTLAs")
	}
	binData, bitEnd = EncodeNumberItemSequenceOf(numItem, 1, ngapType.MaxnoofXnTLAs, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TransportLayerAddress(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}
