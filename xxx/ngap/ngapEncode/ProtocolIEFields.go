// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"
	"ngap/ngapType"
)

// ProtocolIEField 0 - AllowedNSSAI
func AllowedNSSAIItem (value ngapType.AllowedNSSAIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 1 - AMFName

// ProtocolIEField 2 - AMFOverloadResponse
func OverloadAction (value ngapType.OverloadAction, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("OverloadAction: Enumerated: valueExt,valueLB:0,valueUB:3")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolIESingleContainerOverloadResponseExtIEs (value ngapType.ProtocolIESingleContainerOverloadResponseExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 3 - AMFSetID

// ProtocolIEField 4 - AMFTNLAssociationFailedToSetupList
func TNLAssociationItem (value ngapType.TNLAssociationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 5 - AMFTNLAssociationSetupList
func AMFTNLAssociationSetupItem (value ngapType.AMFTNLAssociationSetupItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 6 - AMFTNLAssociationToAddList
func AMFTNLAssociationToAddItem (value ngapType.AMFTNLAssociationToAddItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TNLAssociationUsage != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

// ProtocolIEField 7 - AMFTNLAssociationToRemoveList
func AMFTNLAssociationToRemoveItem (value ngapType.AMFTNLAssociationToRemoveItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 8 - AMFTNLAssociationToUpdateList
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 9 - AMFTrafficLoadReductionIndication

// ProtocolIEField 10 - AMFUENGAPID

// ProtocolIEField 11 - AssistanceDataForPaging
func AssistanceDataForRecommendedCells (value ngapType.AssistanceDataForRecommendedCells, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func PagingAttemptInformation (value ngapType.PagingAttemptInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NextPagingAreaScope != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

func ProtocolExtensionContainerAssistanceDataForPagingExtIEs (value ngapType.ProtocolExtensionContainerAssistanceDataForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForPagingExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AssistanceDataForPagingExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForPagingExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 12 - BroadcastCancelledAreaList
func CellIDCancelledEUTRA (value ngapType.CellIDCancelledEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDCancelledEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDCancelledEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledEUTRA (value ngapType.TAICancelledEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("TAICancelledEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAICancelledEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAICancelledEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledEUTRA (value ngapType.EmergencyAreaIDCancelledEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDCancelledEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDCancelledNR (value ngapType.CellIDCancelledNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CellIDCancelledNR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDCancelledNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDCancelledNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledNR (value ngapType.TAICancelledNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("TAICancelledNR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAICancelledNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAICancelledNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledNR (value ngapType.EmergencyAreaIDCancelledNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDCancelledNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs (value ngapType.ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 13 - BroadcastCompletedAreaList
func CellIDBroadcastEUTRA (value ngapType.CellIDBroadcastEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDBroadcastEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDBroadcastEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastEUTRA (value ngapType.TAIBroadcastEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIBroadcastEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIBroadcastEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastEUTRA (value ngapType.EmergencyAreaIDBroadcastEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDBroadcastEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CellIDBroadcastNR (value ngapType.CellIDBroadcastNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CellIDBroadcastNR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDBroadcastNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDBroadcastNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastNR (value ngapType.TAIBroadcastNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("TAIBroadcastNR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIBroadcastNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIBroadcastNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastNR (value ngapType.EmergencyAreaIDBroadcastNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDBroadcastNRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs (value ngapType.ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 14 - CancelAllWarningMessages

// ProtocolIEField 15 - Cause
func CauseRadioNetwork (value ngapType.CauseRadioNetwork, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 44 {
		return binData, bitEnd, errors.New("CauseRadioNetwork: Enumerated: valueExt,valueLB:0,valueUB:44")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 7, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseTransport (value ngapType.CauseTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("CauseTransport: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseNas (value ngapType.CauseNas, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("CauseNas: Enumerated: valueExt,valueLB:0,valueUB:3")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseProtocol (value ngapType.CauseProtocol, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 6 {
		return binData, bitEnd, errors.New("CauseProtocol: Enumerated: valueExt,valueLB:0,valueUB:6")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseMisc (value ngapType.CauseMisc, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 5 {
		return binData, bitEnd, errors.New("CauseMisc: Enumerated: valueExt,valueLB:0,valueUB:5")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolIESingleContainerCauseExtIEs (value ngapType.ProtocolIESingleContainerCauseExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 16 - CellIDListForRestart
func EUTRACGIList (value ngapType.EUTRACGIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("EUTRACGIList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EUTRACGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EUTRACGIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NRCGIList (value ngapType.NRCGIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16384 {
		return binData, bitEnd, errors.New("NRCGIList: List: valueExt,sizeLB:1,sizeUB:16384")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NRCGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NRCGIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerCellIDListForRestartExtIEs (value ngapType.ProtocolIESingleContainerCellIDListForRestartExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 17 - ConcurrentWarningMessageInd

// ProtocolIEField 18 - CoreNetworkAssistanceInformation
func UEIdentityIndexValue (value ngapType.UEIdentityIndexValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
	switch value.Present {
	case ngapType.UEIdentityIndexValuePresentIndexLength10:
		if value.IndexLength10 == nil {
			return binData, bitEnd, errors.New("UEIdentityIndexValue: IndexLength10: NIL")
		}
		if value.IndexLength10.BitLength != 10 || len(value.IndexLength10.Bytes) != 2 {
			return binData, bitEnd, errors.New("UEIdentityIndexValue: IndexLength10: sizeLB:10,sizeUB:10")
		}
		binData, bitEnd = EncodeBitString(*value.IndexLength10, binData, bitEnd)
	case ngapType.UEIdentityIndexValuePresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("UEIdentityIndexValue: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerUEIdentityIndexValueExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEIdentityIndexValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEIdentityIndexValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PeriodicRegistrationUpdateTimer (value ngapType.PeriodicRegistrationUpdateTimer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("PeriodicRegistrationUpdateTimer: BitString: sizeLB:8,sizeUB:8")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func MICOModeIndication (value ngapType.MICOModeIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("MICOModeIndication: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 1, binData, bitEnd)
	return binData, bitEnd, nil
}

func TAIListForInactive (value ngapType.TAIListForInactive, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("TAIListForInactive: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIListForInactiveItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForInactive: ["+strconv.Itoa(i)+"]: " + err.Error())
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
	binData, bitEnd = EncodeUint64(uint64(option), 6, binData, bitEnd)
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

func ProtocolExtensionContainerCoreNetworkAssistanceInformationExtIEs (value ngapType.ProtocolExtensionContainerCoreNetworkAssistanceInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCoreNetworkAssistanceInformationExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CoreNetworkAssistanceInformationExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCoreNetworkAssistanceInformationExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 19 - CriticalityDiagnostics
func ProcedureCode (value ngapType.ProcedureCode, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("ProcedureCode: int64: valueLB:0,valueUB:255")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 8, binData, 8)
	return binData, bitEnd, nil
}

func TriggeringMessage (value ngapType.TriggeringMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("TriggeringMessage: Enumerated: valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func CriticalityDiagnosticsIEList (value ngapType.CriticalityDiagnosticsIEList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CriticalityDiagnosticsIEItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CriticalityDiagnosticsIEList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCriticalityDiagnosticsExtIEs (value ngapType.ProtocolExtensionContainerCriticalityDiagnosticsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CriticalityDiagnosticsExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 20 - DataCodingScheme

// ProtocolIEField 21 - DefaultPagingDRX

// ProtocolIEField 22 - DirectForwardingPathAvailability

// ProtocolIEField 23 - EmergencyAreaIDListForRestart

// ProtocolIEField 24 - EmergencyFallbackIndicator
func EmergencyFallbackRequestIndicator (value ngapType.EmergencyFallbackRequestIndicator, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("EmergencyFallbackRequestIndicator: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func EmergencyServiceTargetCN (value ngapType.EmergencyServiceTargetCN, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("EmergencyServiceTargetCN: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs (value ngapType.ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyFallbackIndicatorExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 25 - EUTRACGI
func EUTRACellIdentity (value ngapType.EUTRACellIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 28 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("EUTRACellIdentity: BitString: sizeLB:28,sizeUB:28")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerEUTRACGIExtIEs (value ngapType.ProtocolExtensionContainerEUTRACGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEUTRACGIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EUTRACGIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEUTRACGIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 26 - FiveGSTMSI
func AMFPointer (value ngapType.AMFPointer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 6 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("AMFPointer: BitString: sizeLB:6,sizeUB:6")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func FiveGTMSI (value ngapType.FiveGTMSI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 4 {
		return binData, bitEnd, errors.New("FiveGTMSI: OctetString: sizeLB:4,sizeUB:4")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerFiveGSTMSIExtIEs (value ngapType.ProtocolExtensionContainerFiveGSTMSIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerFiveGSTMSIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = FiveGSTMSIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerFiveGSTMSIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 27 - GlobalRANNodeID
func GlobalGNBID (value ngapType.GlobalGNBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func GlobalNgENBID (value ngapType.GlobalNgENBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func GlobalN3IWFID (value ngapType.GlobalN3IWFID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolIESingleContainerGlobalRANNodeIDExtIEs (value ngapType.ProtocolIESingleContainerGlobalRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 28 - GUAMI
func AMFRegionID (value ngapType.AMFRegionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("AMFRegionID: BitString: sizeLB:8,sizeUB:8")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerGUAMIExtIEs (value ngapType.ProtocolExtensionContainerGUAMIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGUAMIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = GUAMIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGUAMIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 29 - HandoverType

// ProtocolIEField 30 - IMSVoiceSupportIndicator

// ProtocolIEField 31 - IndexToRFSP

// ProtocolIEField 32 - InfoOnRecommendedCellsAndRANNodesForPaging
func RecommendedRANNodesForPaging (value ngapType.RecommendedRANNodesForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs (value ngapType.ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPagingExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 33 - LocationReportingRequestType
func EventType (value ngapType.EventType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 5 {
		return binData, bitEnd, errors.New("EventType: Enumerated: valueExt,valueLB:0,valueUB:5")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func ReportArea (value ngapType.ReportArea, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("ReportArea: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 1, binData, bitEnd)
	return binData, bitEnd, nil
}

func AreaOfInterestList (value ngapType.AreaOfInterestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 64 {
		return binData, bitEnd, errors.New("AreaOfInterestList: List: valueExt,sizeLB:1,sizeUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 6, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerLocationReportingRequestTypeExtIEs (value ngapType.ProtocolExtensionContainerLocationReportingRequestTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerLocationReportingRequestTypeExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = LocationReportingRequestTypeExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerLocationReportingRequestTypeExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 34 - MaskedIMEISV

// ProtocolIEField 35 - MessageIdentifier

// ProtocolIEField 36 - MobilityRestrictionList
func EquivalentPLMNs (value ngapType.EquivalentPLMNs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 15 {
		return binData, bitEnd, errors.New("EquivalentPLMNs: List: sizeLB:1,sizeUB:15")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PLMNIdentity(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EquivalentPLMNs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RATRestrictions (value ngapType.RATRestrictions, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem != 16 {
		return binData, bitEnd, errors.New("RATRestrictions: List: valueExt,sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RATRestrictionsItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RATRestrictions: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ForbiddenAreaInformation (value ngapType.ForbiddenAreaInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("ForbiddenAreaInformation: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ForbiddenAreaInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ForbiddenAreaInformation: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ServiceAreaInformation (value ngapType.ServiceAreaInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("ServiceAreaInformation: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ServiceAreaInformationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServiceAreaInformation: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerMobilityRestrictionListExtIEs (value ngapType.ProtocolExtensionContainerMobilityRestrictionListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerMobilityRestrictionListExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = MobilityRestrictionListExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerMobilityRestrictionListExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 37 - NASC

// ProtocolIEField 38 - NASPDU

// ProtocolIEField 39 - NASSecurityParametersFromNGRAN

// ProtocolIEField 40 - NewAMFUENGAPID

// ProtocolIEField 41 - NewSecurityContextInd

// ProtocolIEField 42 - NGAPMessage

// ProtocolIEField 43 - NGRANCGI
func ProtocolIESingleContainerNGRANCGIExtIEs (value ngapType.ProtocolIESingleContainerNGRANCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 44 - NGRANTraceID

// ProtocolIEField 45 - NRCGI
func NRCellIdentity (value ngapType.NRCellIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 36 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("NRCellIdentity: BitString: sizeLB:36,sizeUB:36")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerNRCGIExtIEs (value ngapType.ProtocolExtensionContainerNRCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerNRCGIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NRCGIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerNRCGIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 46 - NRPPaPDU

// ProtocolIEField 47 - NumberOfBroadcastsRequested

// ProtocolIEField 48 - OldAMF

// ProtocolIEField 49 - OverloadStartNSSAIList
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 50 - PagingDRX

// ProtocolIEField 51 - PagingOrigin

// ProtocolIEField 52 - PagingPriority

// ProtocolIEField 53 - PDUSessionResourceAdmittedList
func PDUSessionResourceAdmittedItem (value ngapType.PDUSessionResourceAdmittedItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItem: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.HandoverRequestAcknowledgeTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 54 - PDUSessionResourceFailedToModifyListModRes
func PDUSessionResourceFailedToModifyItemModRes (value ngapType.PDUSessionResourceFailedToModifyItemModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModRes: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceModifyUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 55 - PDUSessionResourceFailedToSetupListCxtRes
func PDUSessionResourceFailedToSetupItemCxtRes (value ngapType.PDUSessionResourceFailedToSetupItemCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtRes: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 56 - PDUSessionResourceFailedToSetupListHOAck
func PDUSessionResourceFailedToSetupItemHOAck (value ngapType.PDUSessionResourceFailedToSetupItemHOAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAck: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.HandoverResourceAllocationUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAck: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 57 - PDUSessionResourceFailedToSetupListPSReq
func PDUSessionResourceFailedToSetupItemPSReq (value ngapType.PDUSessionResourceFailedToSetupItemPSReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReq: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PathSwitchRequestSetupFailedTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 58 - PDUSessionResourceFailedToSetupListSURes
func PDUSessionResourceFailedToSetupItemSURes (value ngapType.PDUSessionResourceFailedToSetupItemSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSURes: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSURes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 59 - PDUSessionResourceHandoverList
func PDUSessionResourceHandoverItem (value ngapType.PDUSessionResourceHandoverItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItem: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.HandoverCommandTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceHandoverItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 60 - PDUSessionResourceListCxtRelCpl
func PDUSessionResourceItemCxtRelCpl (value ngapType.PDUSessionResourceItemCxtRelCpl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 61 - PDUSessionResourceListHORqd
func PDUSessionResourceItemHORqd (value ngapType.PDUSessionResourceItemHORqd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqd: " + err.Error())
	}
	binValue, bitEnd := EncodeOctetString(value.HandoverRequiredTransfer, binData, bitEnd)
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceItemHORqd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 62 - PDUSessionResourceModifyListModCfm
func PDUSessionResourceModifyItemModCfm (value ngapType.PDUSessionResourceModifyItemModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfm: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceModifyConfirmTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfm: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 63 - PDUSessionResourceModifyListModInd
func PDUSessionResourceModifyItemModInd (value ngapType.PDUSessionResourceModifyItemModInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModInd: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceModifyIndicationTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModInd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 64 - PDUSessionResourceModifyListModReq
func PDUSessionResourceModifyItemModReq (value ngapType.PDUSessionResourceModifyItemModReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NASPDU != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceModifyRequestTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 65 - PDUSessionResourceModifyListModRes
func PDUSessionResourceModifyItemModRes (value ngapType.PDUSessionResourceModifyItemModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.PDUSessionResourceModifyResponseTransfer != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModRes: " + err.Error())
	}
	if value.PDUSessionResourceModifyResponseTransfer != nil {
		binData, bitEnd = EncodeOctetString(*value.PDUSessionResourceModifyResponseTransfer, binData, bitEnd)
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 66 - PDUSessionResourceNotifyList
func PDUSessionResourceNotifyItem (value ngapType.PDUSessionResourceNotifyItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItem: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceNotifyTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 67 - PDUSessionResourceReleasedListNot
func PDUSessionResourceReleasedItemNot (value ngapType.PDUSessionResourceReleasedItemNot, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNot: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceNotifyReleasedTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNot: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 68 - PDUSessionResourceReleasedListPSAck
func PDUSessionResourceReleasedItemPSAck (value ngapType.PDUSessionResourceReleasedItemPSAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAck: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PathSwitchRequestUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAck: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 69 - PDUSessionResourceReleasedListPSFail
func PDUSessionResourceReleasedItemPSFail (value ngapType.PDUSessionResourceReleasedItemPSFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFail: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PathSwitchRequestUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFail: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 70 - PDUSessionResourceReleasedListRelRes
func PDUSessionResourceReleasedItemRelRes (value ngapType.PDUSessionResourceReleasedItemRelRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelRes: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceReleaseResponseTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 71 - PDUSessionResourceSetupListCxtReq
func PDUSessionResourceSetupItemCxtReq (value ngapType.PDUSessionResourceSetupItemCxtReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.NASPDU != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupRequestTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 72 - PDUSessionResourceSetupListCxtRes
func PDUSessionResourceSetupItemCxtRes (value ngapType.PDUSessionResourceSetupItemCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtRes: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupResponseTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtRes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 73 - PDUSessionResourceSetupListHOReq
func PDUSessionResourceSetupItemHOReq (value ngapType.PDUSessionResourceSetupItemHOReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
	}
	binData, bitEnd, err = SNSSAI(value.SNSSAI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(EncodeLengthValue(value.HandoverRequestTransfer), binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 74 - PDUSessionResourceSetupListSUReq
func PDUSessionResourceSetupItemSUReq (value ngapType.PDUSessionResourceSetupItemSUReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.PDUSessionNASPDU != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupRequestTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReq: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 75 - PDUSessionResourceSetupListSURes
func PDUSessionResourceSetupItemSURes (value ngapType.PDUSessionResourceSetupItemSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSURes: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupResponseTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSURes: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 76 - PDUSessionResourceToBeSwitchedDLList
func PDUSessionResourceToBeSwitchedDLItem (value ngapType.PDUSessionResourceToBeSwitchedDLItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItem: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PathSwitchRequestTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 77 - PDUSessionResourceSwitchedList
func PDUSessionResourceSwitchedItem (value ngapType.PDUSessionResourceSwitchedItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItem: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PathSwitchRequestAcknowledgeTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 78 - PDUSessionResourceToReleaseListHOCmd
func PDUSessionResourceToReleaseItemHOCmd (value ngapType.PDUSessionResourceToReleaseItemHOCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmd: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.HandoverPreparationUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 79 - PDUSessionResourceToReleaseListRelCmd
func PDUSessionResourceToReleaseItemRelCmd (value ngapType.PDUSessionResourceToReleaseItemRelCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmd: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceReleaseCommandTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmd: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 80 - PLMNSupportList
func PLMNSupportItem (value ngapType.PLMNSupportItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 81 - PWSFailedCellIDList
func ProtocolIESingleContainerPWSFailedCellIDListExtIEs (value ngapType.ProtocolIESingleContainerPWSFailedCellIDListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 82 - RANNodeName

// ProtocolIEField 83 - RANPagingPriority

// ProtocolIEField 84 - RANStatusTransferTransparentContainer
func DRBsSubjectToStatusTransferList (value ngapType.DRBsSubjectToStatusTransferList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 32 {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferList: List: valueExt,sizeLB:1,sizeUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBsSubjectToStatusTransferItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs (value ngapType.ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RANStatusTransferTransparentContainerExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 85 - RANUENGAPID

// ProtocolIEField 86 - RelativeAMFCapacity

// ProtocolIEField 87 - RepetitionPeriod

// ProtocolIEField 88 - ResetType
func ResetAll (value ngapType.ResetAll, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("ResetAll: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 1, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolIESingleContainerResetTypeExtIEs (value ngapType.ProtocolIESingleContainerResetTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 89 - RoutingID

// ProtocolIEField 90 - RRCEstablishmentCause

// ProtocolIEField 91 - RRCInactiveTransitionReportRequest

// ProtocolIEField 92 - RRCState

// ProtocolIEField 93 - SecurityContext
func NextHopChainingCount (value ngapType.NextHopChainingCount, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 7 {
		return binData, bitEnd, errors.New("NextHopChainingCount: int64: valueLB:0,valueUB:7")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerSecurityContextExtIEs (value ngapType.ProtocolExtensionContainerSecurityContextExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityContextExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SecurityContextExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityContextExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 94 - SecurityKey

// ProtocolIEField 95 - SerialNumber

// ProtocolIEField 96 - ServedGUAMIList
func ServedGUAMIItem (value ngapType.ServedGUAMIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.BackupAMFName != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

// ProtocolIEField 97 - SliceSupportList
func SliceSupportItem (value ngapType.SliceSupportItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 98 - SONConfigurationTransferDL
func TargetRANNodeID (value ngapType.TargetRANNodeID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func SourceRANNodeID (value ngapType.SourceRANNodeID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func SONInformation (value ngapType.SONInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.SONInformationPresentSONInformationRequest:
		if value.SONInformationRequest == nil {
			return binData, bitEnd, errors.New("Cause: SONInformationRequest: NIL")
		}
		binData, bitEnd, err = SONInformationRequest(*value.SONInformationRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformation: " + err.Error())
		}
	case ngapType.SONInformationPresentSONInformationReply:
		if value.SONInformationReply == nil {
			return binData, bitEnd, errors.New("Cause: SONInformationReply: NIL")
		}
		binData, bitEnd, err = SONInformationReply(*value.SONInformationReply, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformation: " + err.Error())
		}
	case ngapType.SONInformationPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("Cause: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerSONInformationExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformation: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("SONInformation: Present: INVALID")
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
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

func ProtocolExtensionContainerSONConfigurationTransferExtIEs (value ngapType.ProtocolExtensionContainerSONConfigurationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSONConfigurationTransferExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SONConfigurationTransferExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSONConfigurationTransferExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 99 - SONConfigurationTransferUL

// ProtocolIEField 100 - SourceAMFUENGAPID

// ProtocolIEField 101 - SourceToTargetTransparentContainer

// ProtocolIEField 102 - SupportedTAList
func SupportedTAItem (value ngapType.SupportedTAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 103 - TAIListForPaging
func TAIListForPagingItem (value ngapType.TAIListForPagingItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 104 - TAIListForRestart

// ProtocolIEField 105 - TargetID
func TargeteNBID (value ngapType.TargeteNBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolIESingleContainerTargetIDExtIEs (value ngapType.ProtocolIESingleContainerTargetIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 106 - TargetToSourceTransparentContainer

// ProtocolIEField 107 - TimeToWait

// ProtocolIEField 108 - TraceActivation
func InterfacesToTrace (value ngapType.InterfacesToTrace, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("InterfacesToTrace: BitString: sizeLB:8,sizeUB:8")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func TraceDepth (value ngapType.TraceDepth, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 5 {
		return binData, bitEnd, errors.New("TraceDepth: Enumerated: valueExt,valueLB:0,valueUB:5")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerTraceActivationExtIEs (value ngapType.ProtocolExtensionContainerTraceActivationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTraceActivationExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TraceActivationExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTraceActivationExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 109 - TraceCollectionEntityIPAddress

// ProtocolIEField 110 - UEAggregateMaximumBitRate
func BitRate (value ngapType.BitRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4000000000000 {
		return binData, bitEnd, errors.New("BitRate: int64: valueExt,valueLB:0,valueUB:4000000000000")
	}
	binValue := ConvertInt64ToBytes(value.Value)
	binData, bitEnd = EncodeUint64(uint64(len(binValue))-1, 4, binData, bitEnd)
	binData, bitEnd = EncodeBytes(binValue, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs (value ngapType.ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEAggregateMaximumBitRateExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 111 - UEAssociatedLogicalNGConnectionList
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 112 - UEContextRequest

// ProtocolIEField 113 - UENGAPIDs
func UENGAPIDPair (value ngapType.UENGAPIDPair, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolIESingleContainerUENGAPIDsExtIEs (value ngapType.ProtocolIESingleContainerUENGAPIDsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 114 - UEPagingIdentity
func ProtocolIESingleContainerUEPagingIdentityExtIEs (value ngapType.ProtocolIESingleContainerUEPagingIdentityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 115 - UEPresenceInAreaOfInterestList
func UEPresenceInAreaOfInterestItem (value ngapType.UEPresenceInAreaOfInterestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 116 - UERadioCapability

// ProtocolIEField 117 - UERadioCapabilityForPaging
func UERadioCapabilityForPagingOfNR (value ngapType.UERadioCapabilityForPagingOfNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return append(binData, EncodeLengthValue(value.Value)...), 8, nil
}

func UERadioCapabilityForPagingOfEUTRA (value ngapType.UERadioCapabilityForPagingOfEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return append(binData, EncodeLengthValue(value.Value)...), 8, nil
}

func ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs (value ngapType.ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UERadioCapabilityForPagingExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 118 - UESecurityCapabilities
func NRencryptionAlgorithms (value ngapType.NRencryptionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("NRencryptionAlgorithms: BitString: sizeExt,sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(0, 1, binData, bitEnd)
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func NRintegrityProtectionAlgorithms (value ngapType.NRintegrityProtectionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("NRintegrityProtectionAlgorithms: BitString: sizeExt,sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(0, 1, binData, bitEnd)
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func EUTRAencryptionAlgorithms (value ngapType.EUTRAencryptionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("EUTRAencryptionAlgorithms: BitString: sizeExt,sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(0, 1, binData, bitEnd)
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func EUTRAintegrityProtectionAlgorithms (value ngapType.EUTRAintegrityProtectionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("EUTRAintegrityProtectionAlgorithms: BitString: sizeExt,sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(0, 1, binData, bitEnd)
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerUESecurityCapabilitiesExtIEs (value ngapType.ProtocolExtensionContainerUESecurityCapabilitiesExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUESecurityCapabilitiesExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UESecurityCapabilitiesExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUESecurityCapabilitiesExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 119 - UnavailableGUAMIList
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 120 - UserLocationInformation
func UserLocationInformationEUTRA (value ngapType.UserLocationInformationEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TimeStamp != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

func UserLocationInformationNR (value ngapType.UserLocationInformationNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.TimeStamp != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

func UserLocationInformationN3IWF (value ngapType.UserLocationInformationN3IWF, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolIESingleContainerUserLocationInformationExtIEs (value ngapType.ProtocolIESingleContainerUserLocationInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 121 - WarningAreaList
func EUTRACGIListForWarning (value ngapType.EUTRACGIListForWarning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("EUTRACGIListForWarning: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EUTRACGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EUTRACGIListForWarning: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NRCGIListForWarning (value ngapType.NRCGIListForWarning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("NRCGIListForWarning: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NRCGI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NRCGIListForWarning: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIListForWarning (value ngapType.TAIListForWarning, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("TAIListForWarning: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForWarning: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDList (value ngapType.EmergencyAreaIDList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("EmergencyAreaIDList: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaID(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerWarningAreaListExtIEs (value ngapType.ProtocolIESingleContainerWarningAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 122 - WarningMessageContents

// ProtocolIEField 123 - WarningSecurityInfo

// ProtocolIEField 124 - WarningType

// ProtocolIEField 125 - AdditionalULNGUUPTNLInformation
func GTPTunnel (value ngapType.GTPTunnel, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolIESingleContainerUPTransportLayerInformationExtIEs (value ngapType.ProtocolIESingleContainerUPTransportLayerInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

// ProtocolIEField 126 - DataForwardingNotPossible

// ProtocolIEField 127 - DLNGUUPTNLInformation

// ProtocolIEField 128 - NetworkInstance

// ProtocolIEField 129 - PDUSessionAggregateMaximumBitRate
func ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionAggregateMaximumBitRateExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 130 - PDUSessionResourceFailedToModifyListModCfm
func PDUSessionResourceFailedToModifyItemModCfm (value ngapType.PDUSessionResourceFailedToModifyItemModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfm: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceModifyIndicationUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfm: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 131 - PDUSessionResourceFailedToSetupListCxtFail
func PDUSessionResourceFailedToSetupItemCxtFail (value ngapType.PDUSessionResourceFailedToSetupItemCxtFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PDUSessionID(value.PDUSessionID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFail: " + err.Error())
	}
	binData, bitEnd = EncodeOctetString(value.PDUSessionResourceSetupUnsuccessfulTransfer, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFail: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 132 - PDUSessionResourceListCxtRelReq
func PDUSessionResourceItemCxtRelReq (value ngapType.PDUSessionResourceItemCxtRelReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 133 - PDUSessionType

// ProtocolIEField 134 - QosFlowAddOrModifyRequestList
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 135 - QosFlowSetupRequestList
func QosFlowSetupRequestItem (value ngapType.QosFlowSetupRequestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.ERABID != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

// ProtocolIEField 136 - QosFlowToReleaseList
func QosFlowItem (value ngapType.QosFlowItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = QosFlowIdentifier(value.QosFlowIdentifier, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItem: " + err.Error())
	}
	binData, bitEnd, err = Cause(value.Cause, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItem: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerQosFlowItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 137 - SecurityIndication
func IntegrityProtectionIndication (value ngapType.IntegrityProtectionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("IntegrityProtectionIndication: Enumerated: valueExt,valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ConfidentialityProtectionIndication (value ngapType.ConfidentialityProtectionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("ConfidentialityProtectionIndication: Enumerated: valueExt,valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func MaximumIntegrityProtectedDataRate (value ngapType.MaximumIntegrityProtectedDataRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("MaximumIntegrityProtectedDataRate: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerSecurityIndicationExtIEs (value ngapType.ProtocolExtensionContainerSecurityIndicationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityIndicationExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SecurityIndicationExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSecurityIndicationExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 138 - ULNGUUPTNLInformation

// ProtocolIEField 139 - ULNGUUPTNLModifyList
func ULNGUUPTNLModifyItem (value ngapType.ULNGUUPTNLModifyItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 140 - WarningAreaCoordinates

