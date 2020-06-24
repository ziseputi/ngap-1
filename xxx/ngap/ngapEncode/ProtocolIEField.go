// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"
	"ngap/ngapType"
)

// ProtocolIEField 0 - AllowedNSSAI
func AllowedNSSAI (value ngapType.AllowedNSSAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 8 {
		return binData, bitEnd, errors.New("AllowedNSSAI: List: valueExt,sizeLB:1,sizeUB:8")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 3, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AllowedNSSAIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AllowedNSSAI: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 1 - AMFName
func AMFName (value ngapType.AMFName, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 150 {
		return binData, bitEnd, errors.New("AMFName: string: sizeExt,sizeLB:1,sizeUB:150")
	}
	binData, bitEnd = EncodeStringExt(value.Value, len(value.Value)-1, 9 , binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 2 - AMFOverloadResponse
func AMFOverloadResponse (value ngapType.OverloadResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return OverloadResponse(value, binData, bitEnd)
}

func OverloadResponse (value ngapType.OverloadResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.OverloadResponsePresentOverloadAction:
		if value.OverloadAction == nil {
			return binData, bitEnd, errors.New("OverloadResponse: OverloadAction: NIL")
		}
		binData, bitEnd, err = OverloadAction(*value.OverloadAction, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadResponse: " + err.Error())
		}
	case ngapType.OverloadResponsePresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("OverloadResponse: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerOverloadResponseExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadResponse: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("OverloadResponse: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 3 - AMFSetID
func AMFSetID (value ngapType.AMFSetID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 10 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("AMFSetID: BitString: sizeLB:10,sizeUB:10")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 4 - AMFTNLAssociationFailedToSetupList
func AMFTNLAssociationFailedToSetupList (value ngapType.TNLAssociationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return TNLAssociationList(value, binData, bitEnd)
}

func TNLAssociationList (value ngapType.TNLAssociationList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 32 {
		return binData, bitEnd, errors.New("TNLAssociationList: List: valueExt,sizeLB:1,sizeUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 5, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TNLAssociationItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TNLAssociationList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 5 - AMFTNLAssociationSetupList
func AMFTNLAssociationSetupList (value ngapType.AMFTNLAssociationSetupList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 32 {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupList: List: valueExt,sizeLB:1,sizeUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationSetupItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationSetupList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 6 - AMFTNLAssociationToAddList
func AMFTNLAssociationToAddList (value ngapType.AMFTNLAssociationToAddList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 32 {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddList: List: valueExt,sizeLB:1,sizeUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 5, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToAddItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToAddList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 7 - AMFTNLAssociationToRemoveList
func AMFTNLAssociationToRemoveList (value ngapType.AMFTNLAssociationToRemoveList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 32 {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveList: List: valueExt,sizeLB:1,sizeUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 5, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToRemoveItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 8 - AMFTNLAssociationToUpdateList
func AMFTNLAssociationToUpdateList (value ngapType.AMFTNLAssociationToUpdateList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 32 {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateList: List: valueExt,sizeLB:1,sizeUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 5, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToUpdateItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 9 - AMFTrafficLoadReductionIndication
func AMFTrafficLoadReductionIndication (value ngapType.TrafficLoadReductionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return TrafficLoadReductionIndication(value, binData, bitEnd)
}

func TrafficLoadReductionIndication (value ngapType.TrafficLoadReductionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 99 {
		return binData, bitEnd, errors.New("TrafficLoadReductionIndication: int64: valueLB:1,valueUB:99")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 10 - AMFUENGAPID
func AMFUENGAPID (value ngapType.AMFUENGAPID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1099511627775 {
		return binData, bitEnd, errors.New("AMFUENGAPID: int64: valueLB:0,valueUB:1099511627775")
	}
	binValue := ConvertInt64ToBytes(value.Value)
	binData, bitEnd = EncodeUint64(uint64(len(binValue))-1, 3, binData, bitEnd)
	binData, bitEnd = EncodeBytes(binValue, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 11 - AssistanceDataForPaging
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 12 - BroadcastCancelledAreaList
func BroadcastCancelledAreaList (value ngapType.BroadcastCancelledAreaList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.BroadcastCancelledAreaListPresentCellIDCancelledEUTRA:
		if value.CellIDCancelledEUTRA == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: CellIDCancelledEUTRA: NIL")
		}
		binData, bitEnd, err = CellIDCancelledEUTRA(*value.CellIDCancelledEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	case ngapType.BroadcastCancelledAreaListPresentTAICancelledEUTRA:
		if value.TAICancelledEUTRA == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: TAICancelledEUTRA: NIL")
		}
		binData, bitEnd, err = TAICancelledEUTRA(*value.TAICancelledEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	case ngapType.BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledEUTRA:
		if value.EmergencyAreaIDCancelledEUTRA == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: EmergencyAreaIDCancelledEUTRA: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDCancelledEUTRA(*value.EmergencyAreaIDCancelledEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	case ngapType.BroadcastCancelledAreaListPresentCellIDCancelledNR:
		if value.CellIDCancelledNR == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: CellIDCancelledNR: NIL")
		}
		binData, bitEnd, err = CellIDCancelledNR(*value.CellIDCancelledNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	case ngapType.BroadcastCancelledAreaListPresentTAICancelledNR:
		if value.TAICancelledNR == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: TAICancelledNR: NIL")
		}
		binData, bitEnd, err = TAICancelledNR(*value.TAICancelledNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	case ngapType.BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledNR:
		if value.EmergencyAreaIDCancelledNR == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: EmergencyAreaIDCancelledNR: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDCancelledNR(*value.EmergencyAreaIDCancelledNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	case ngapType.BroadcastCancelledAreaListPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCancelledAreaList: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("BroadcastCancelledAreaList: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 13 - BroadcastCompletedAreaList
func BroadcastCompletedAreaList (value ngapType.BroadcastCompletedAreaList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA:
		if value.CellIDBroadcastEUTRA == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: CellIDBroadcastEUTRA: NIL")
		}
		binData, bitEnd, err = CellIDBroadcastEUTRA(*value.CellIDBroadcastEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	case ngapType.BroadcastCompletedAreaListPresentTAIBroadcastEUTRA:
		if value.TAIBroadcastEUTRA == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: TAIBroadcastEUTRA: NIL")
		}
		binData, bitEnd, err = TAIBroadcastEUTRA(*value.TAIBroadcastEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	case ngapType.BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA:
		if value.EmergencyAreaIDBroadcastEUTRA == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: EmergencyAreaIDBroadcastEUTRA: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDBroadcastEUTRA(*value.EmergencyAreaIDBroadcastEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	case ngapType.BroadcastCompletedAreaListPresentCellIDBroadcastNR:
		if value.CellIDBroadcastNR == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: CellIDBroadcastNR: NIL")
		}
		binData, bitEnd, err = CellIDBroadcastNR(*value.CellIDBroadcastNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	case ngapType.BroadcastCompletedAreaListPresentTAIBroadcastNR:
		if value.TAIBroadcastNR == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: TAIBroadcastNR: NIL")
		}
		binData, bitEnd, err = TAIBroadcastNR(*value.TAIBroadcastNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	case ngapType.BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR:
		if value.EmergencyAreaIDBroadcastNR == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: EmergencyAreaIDBroadcastNR: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDBroadcastNR(*value.EmergencyAreaIDBroadcastNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	case ngapType.BroadcastCompletedAreaListPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastCompletedAreaList: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("BroadcastCompletedAreaList: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 14 - CancelAllWarningMessages
func CancelAllWarningMessages (value ngapType.CancelAllWarningMessages, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("CancelAllWarningMessages: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 15 - Cause
func Cause (value ngapType.Cause, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)
	switch value.Present {
	case ngapType.CausePresentRadioNetwork:
		if value.RadioNetwork == nil {
			return binData, bitEnd, errors.New("Cause: RadioNetwork: NIL")
		}
		binData, bitEnd, err = CauseRadioNetwork(*value.RadioNetwork, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Cause: " + err.Error())
		}
	case ngapType.CausePresentTransport:
		if value.Transport == nil {
			return binData, bitEnd, errors.New("Cause: Transport: NIL")
		}
		binData, bitEnd, err = CauseTransport(*value.Transport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Cause: " + err.Error())
		}
	case ngapType.CausePresentNas:
		if value.Nas == nil {
			return binData, bitEnd, errors.New("Cause: Nas: NIL")
		}
		binData, bitEnd, err = CauseNas(*value.Nas, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Cause: " + err.Error())
		}
	case ngapType.CausePresentProtocol:
		if value.Protocol == nil {
			return binData, bitEnd, errors.New("Cause: Protocol: NIL")
		}
		binData, bitEnd, err = CauseProtocol(*value.Protocol, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Cause: " + err.Error())
		}
	case ngapType.CausePresentMisc:
		if value.Misc == nil {
			return binData, bitEnd, errors.New("Cause: Misc: NIL")
		}
		binData, bitEnd, err = CauseMisc(*value.Misc, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Cause: " + err.Error())
		}
	case ngapType.CausePresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("Cause: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerCauseExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("Cause: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("Cause: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 16 - CellIDListForRestart
func CellIDListForRestart (value ngapType.CellIDListForRestart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.CellIDListForRestartPresentEUTRACGIListforRestart:
		if value.EUTRACGIListforRestart == nil {
			return binData, bitEnd, errors.New("CellIDListForRestart: EUTRACGIListforRestart: NIL")
		}
		binData, bitEnd, err = EUTRACGIList(*value.EUTRACGIListforRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDListForRestart: " + err.Error())
		}
	case ngapType.CellIDListForRestartPresentNRCGIListforRestart:
		if value.NRCGIListforRestart == nil {
			return binData, bitEnd, errors.New("CellIDListForRestart: NRCGIListforRestart: NIL")
		}
		binData, bitEnd, err = NRCGIList(*value.NRCGIListforRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDListForRestart: " + err.Error())
		}
	case ngapType.CellIDListForRestartPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("CellIDListForRestart: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerCellIDListForRestartExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellIDListForRestart: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("CellIDListForRestart: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 17 - ConcurrentWarningMessageInd
func ConcurrentWarningMessageInd (value ngapType.ConcurrentWarningMessageInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("ConcurrentWarningMessageInd: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 18 - CoreNetworkAssistanceInformation
func CoreNetworkAssistanceInformation (value ngapType.CoreNetworkAssistanceInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
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
	binData, bitEnd = EncodeUint64(option, 5, binData, bitEnd)
	binData, bitEnd, err = UEIdentityIndexValue(value.UEIdentityIndexValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
	}
	if value.UESpecificDRX != nil {
		binData, bitEnd, err = PagingDRX(*value.UESpecificDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
		}
	}
	binData, bitEnd, err = PeriodicRegistrationUpdateTimer(value.PeriodicRegistrationUpdateTimer, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
	}
	if value.MICOModeIndication != nil {
		binData, bitEnd, err = MICOModeIndication(*value.MICOModeIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
		}
	}
	binData, bitEnd, err = TAIListForInactive(value.TAIListForInactive, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
	}
	if value.ExpectedUEBehaviour != nil {
		binData, bitEnd, err = ExpectedUEBehaviour(*value.ExpectedUEBehaviour, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
		}
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCoreNetworkAssistanceInformationExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CoreNetworkAssistanceInformation: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 19 - CriticalityDiagnostics
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
	binData, bitEnd = EncodeUint64(uint64(option), 6, binData, bitEnd)
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

// ProtocolIEField 20 - DataCodingScheme
func DataCodingScheme (value ngapType.DataCodingScheme, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("DataCodingScheme: BitString: sizeLB:8,sizeUB:8")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 21 - DefaultPagingDRX
func DefaultPagingDRX (value ngapType.PagingDRX, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return PagingDRX(value, binData, bitEnd)
}

// ProtocolIEField 22 - DirectForwardingPathAvailability
func DirectForwardingPathAvailability (value ngapType.DirectForwardingPathAvailability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("DirectForwardingPathAvailability: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 1, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 23 - EmergencyAreaIDListForRestart
func EmergencyAreaIDListForRestart (value ngapType.EmergencyAreaIDListForRestart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("EmergencyAreaIDListForRestart: List: sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaID(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("EmergencyAreaIDListForRestart: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 24 - EmergencyFallbackIndicator
func EmergencyFallbackIndicator (value ngapType.EmergencyFallbackIndicator, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.EmergencyServiceTargetCN != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

// ProtocolIEField 25 - EUTRACGI
func EUTRACGI (value ngapType.EUTRACGI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 26 - FiveGSTMSI
func FiveGSTMSI (value ngapType.FiveGSTMSI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 27 - GlobalRANNodeID
func GlobalRANNodeID (value ngapType.GlobalRANNodeID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.GlobalRANNodeIDPresentGlobalGNBID:
		if value.GlobalGNBID == nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: GlobalGNBID: NIL")
		}
		binData, bitEnd, err = GlobalGNBID(*value.GlobalGNBID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: " + err.Error())
		}
	case ngapType.GlobalRANNodeIDPresentGlobalNgENBID:
		if value.GlobalNgENBID == nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: GlobalNgENBID: NIL")
		}
		binData, bitEnd, err = GlobalNgENBID(*value.GlobalNgENBID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: " + err.Error())
		}
	case ngapType.GlobalRANNodeIDPresentGlobalN3IWFID:
		if value.GlobalN3IWFID == nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: GlobalN3IWFID: NIL")
		}
		binData, bitEnd, err = GlobalN3IWFID(*value.GlobalN3IWFID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: " + err.Error())
		}
	case ngapType.GlobalRANNodeIDPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerGlobalRANNodeIDExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GlobalRANNodeID: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("GlobalRANNodeID: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 28 - GUAMI
func GUAMI (value ngapType.GUAMI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 29 - HandoverType
func HandoverType (value ngapType.HandoverType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("HandoverType: Enumerated: valueExt,valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 3, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 30 - IMSVoiceSupportIndicator
func IMSVoiceSupportIndicator (value ngapType.IMSVoiceSupportIndicator, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("IMSVoiceSupportIndicator: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 31 - IndexToRFSP
func IndexToRFSP (value ngapType.IndexToRFSP, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 256 {
		return binData, bitEnd, errors.New("IndexToRFSP: int64: valueExt,valueLB:1,valueUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value)-1, 16, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 32 - InfoOnRecommendedCellsAndRANNodesForPaging
func InfoOnRecommendedCellsAndRANNodesForPaging (value ngapType.InfoOnRecommendedCellsAndRANNodesForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 33 - LocationReportingRequestType
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 34 - MaskedIMEISV
func MaskedIMEISV (value ngapType.MaskedIMEISV, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 64 || len(value.Value.Bytes) != 8 {
		return binData, bitEnd, errors.New("MaskedIMEISV: BitString: sizeLB:64,sizeUB:64")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 35 - MessageIdentifier
func MessageIdentifier (value ngapType.MessageIdentifier, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("MessageIdentifier: BitString: sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 36 - MobilityRestrictionList
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
	binData, bitEnd = EncodeUint64(uint64(option), 6, binData, bitEnd)
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

// ProtocolIEField 37 - NASC
func NASC (value ngapType.NASPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return NASPDU(value, binData, bitEnd)
}

// ProtocolIEField 38 - NASPDU
func NASPDU (value ngapType.NASPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return append(binData, EncodeLengthValue(value.Value)...), 8, nil
}

// ProtocolIEField 39 - NASSecurityParametersFromNGRAN
func NASSecurityParametersFromNGRAN (value ngapType.NASSecurityParametersFromNGRAN, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 40 - NewAMFUENGAPID
func NewAMFUENGAPID (value ngapType.AMFUENGAPID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return AMFUENGAPID(value, binData, bitEnd)
}

// ProtocolIEField 41 - NewSecurityContextInd
func NewSecurityContextInd (value ngapType.NewSecurityContextInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("NewSecurityContextInd: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 1, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 42 - NGAPMessage
func NGAPMessage (value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	binData, bitEnd = EncodeOctetString(value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 43 - NGRANCGI
func NGRANCGI (value ngapType.NGRANCGI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.NGRANCGIPresentNRCGI:
		if value.NRCGI == nil {
			return binData, bitEnd, errors.New("NGRANCGI: NRCGI: NIL")
		}
		binData, bitEnd, err = NRCGI(*value.NRCGI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGRANCGI: " + err.Error())
		}
	case ngapType.NGRANCGIPresentEUTRACGI:
		if value.EUTRACGI == nil {
			return binData, bitEnd, errors.New("NGRANCGI: EUTRACGI: NIL")
		}
		binData, bitEnd, err = EUTRACGI(*value.EUTRACGI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGRANCGI: " + err.Error())
		}
	case ngapType.NGRANCGIPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("NGRANCGI: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerNGRANCGIExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGRANCGI: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGRANCGI: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 44 - NGRANTraceID
func NGRANTraceID (value ngapType.NGRANTraceID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 8 {
		return binData, bitEnd, errors.New("NGRANTraceID: OctetString: sizeLB:8,sizeUB:8")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 45 - NRCGI
func NRCGI (value ngapType.NRCGI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 46 - NRPPaPDU
func NRPPaPDU (value ngapType.NRPPaPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 47 - NumberOfBroadcastsRequested
func NumberOfBroadcastsRequested (value ngapType.NumberOfBroadcastsRequested, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 65535 {
		return binData, bitEnd, errors.New("NumberOfBroadcastsRequested: int64: valueLB:0,valueUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 48 - OldAMF
func OldAMF (value ngapType.AMFName, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return AMFName(value, binData, bitEnd)
}

// ProtocolIEField 49 - OverloadStartNSSAIList
func OverloadStartNSSAIList (value ngapType.OverloadStartNSSAIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 1024 {
		return binData, bitEnd, errors.New("OverloadStartNSSAIList: List: valueExt,sizeLB:1,sizeUB:1024")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = OverloadStartNSSAIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartNSSAIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 50 - PagingDRX
func PagingDRX (value ngapType.PagingDRX, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("PagingDRX: Enumerated: valueExt,valueLB:0,valueUB:3")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 3, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 51 - PagingOrigin
func PagingOrigin (value ngapType.PagingOrigin, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("PagingOrigin: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 52 - PagingPriority
func PagingPriority (value ngapType.PagingPriority, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 7 {
		return binData, bitEnd, errors.New("PagingPriority: Enumerated: valueExt,valueLB:0,valueUB:7")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 53 - PDUSessionResourceAdmittedList
func PDUSessionResourceAdmittedList (value ngapType.PDUSessionResourceAdmittedList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceAdmittedItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceAdmittedList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 54 - PDUSessionResourceFailedToModifyListModRes
func PDUSessionResourceFailedToModifyListModRes (value ngapType.PDUSessionResourceFailedToModifyListModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModRes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 55 - PDUSessionResourceFailedToSetupListCxtRes
func PDUSessionResourceFailedToSetupListCxtRes (value ngapType.PDUSessionResourceFailedToSetupListCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtRes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 56 - PDUSessionResourceFailedToSetupListHOAck
func PDUSessionResourceFailedToSetupListHOAck (value ngapType.PDUSessionResourceFailedToSetupListHOAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListHOAck: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemHOAck(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListHOAck: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 57 - PDUSessionResourceFailedToSetupListPSReq
func PDUSessionResourceFailedToSetupListPSReq (value ngapType.PDUSessionResourceFailedToSetupListPSReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListPSReq: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemPSReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListPSReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 58 - PDUSessionResourceFailedToSetupListSURes
func PDUSessionResourceFailedToSetupListSURes (value ngapType.PDUSessionResourceFailedToSetupListSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListSURes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemSURes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListSURes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 59 - PDUSessionResourceHandoverList
func PDUSessionResourceHandoverList (value ngapType.PDUSessionResourceHandoverList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceHandoverItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceHandoverList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 60 - PDUSessionResourceListCxtRelCpl
func PDUSessionResourceListCxtRelCpl (value ngapType.PDUSessionResourceListCxtRelCpl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelCpl: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemCxtRelCpl(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelCpl: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 61 - PDUSessionResourceListHORqd
func PDUSessionResourceListHORqd (value ngapType.PDUSessionResourceListHORqd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceListHORqd: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 8, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemHORqd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceListHORqd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 62 - PDUSessionResourceModifyListModCfm
func PDUSessionResourceModifyListModCfm (value ngapType.PDUSessionResourceModifyListModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModCfm: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModCfm(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModCfm: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 63 - PDUSessionResourceModifyListModInd
func PDUSessionResourceModifyListModInd (value ngapType.PDUSessionResourceModifyListModInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModInd: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModInd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModInd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 64 - PDUSessionResourceModifyListModReq
func PDUSessionResourceModifyListModReq (value ngapType.PDUSessionResourceModifyListModReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModReq: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 65 - PDUSessionResourceModifyListModRes
func PDUSessionResourceModifyListModRes (value ngapType.PDUSessionResourceModifyListModRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyListModRes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyListModRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 66 - PDUSessionResourceNotifyList
func PDUSessionResourceNotifyList (value ngapType.PDUSessionResourceNotifyList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceNotifyItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 67 - PDUSessionResourceReleasedListNot
func PDUSessionResourceReleasedListNot (value ngapType.PDUSessionResourceReleasedListNot, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListNot: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemNot(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListNot: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 68 - PDUSessionResourceReleasedListPSAck
func PDUSessionResourceReleasedListPSAck (value ngapType.PDUSessionResourceReleasedListPSAck, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSAck: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemPSAck(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSAck: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 69 - PDUSessionResourceReleasedListPSFail
func PDUSessionResourceReleasedListPSFail (value ngapType.PDUSessionResourceReleasedListPSFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSFail: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemPSFail(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListPSFail: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 70 - PDUSessionResourceReleasedListRelRes
func PDUSessionResourceReleasedListRelRes (value ngapType.PDUSessionResourceReleasedListRelRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedListRelRes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemRelRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleasedListRelRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 71 - PDUSessionResourceSetupListCxtReq
func PDUSessionResourceSetupListCxtReq (value ngapType.PDUSessionResourceSetupListCxtReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtReq: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 8, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemCxtReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 72 - PDUSessionResourceSetupListCxtRes
func PDUSessionResourceSetupListCxtRes (value ngapType.PDUSessionResourceSetupListCxtRes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtRes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemCxtRes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListCxtRes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 73 - PDUSessionResourceSetupListHOReq
func PDUSessionResourceSetupListHOReq (value ngapType.PDUSessionResourceSetupListHOReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListHOReq: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 8, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemHOReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListHOReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 74 - PDUSessionResourceSetupListSUReq
func PDUSessionResourceSetupListSUReq (value ngapType.PDUSessionResourceSetupListSUReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListSUReq: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemSUReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListSUReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 75 - PDUSessionResourceSetupListSURes
func PDUSessionResourceSetupListSURes (value ngapType.PDUSessionResourceSetupListSURes, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupListSURes: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemSURes(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupListSURes: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 76 - PDUSessionResourceToBeSwitchedDLList
func PDUSessionResourceToBeSwitchedDLList (value ngapType.PDUSessionResourceToBeSwitchedDLList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 77 - PDUSessionResourceSwitchedList
func PDUSessionResourceSwitchedList (value ngapType.PDUSessionResourceSwitchedList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSwitchedItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSwitchedList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 78 - PDUSessionResourceToReleaseListHOCmd
func PDUSessionResourceToReleaseListHOCmd (value ngapType.PDUSessionResourceToReleaseListHOCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListHOCmd: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToReleaseItemHOCmd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListHOCmd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 79 - PDUSessionResourceToReleaseListRelCmd
func PDUSessionResourceToReleaseListRelCmd (value ngapType.PDUSessionResourceToReleaseListRelCmd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListRelCmd: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToReleaseItemRelCmd(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceToReleaseListRelCmd: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 80 - PLMNSupportList
func PLMNSupportList (value ngapType.PLMNSupportList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 12 {
		return binData, bitEnd, errors.New("PLMNSupportList: List: valueExt,sizeLB:1,sizeUB:12")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PLMNSupportItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PLMNSupportList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 81 - PWSFailedCellIDList
func PWSFailedCellIDList (value ngapType.PWSFailedCellIDList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.PWSFailedCellIDListPresentEUTRACGIPWSFailedList:
		if value.EUTRACGIPWSFailedList == nil {
			return binData, bitEnd, errors.New("PWSFailedCellIDList: EUTRACGIPWSFailedList: NIL")
		}
		binData, bitEnd, err = EUTRACGIList(*value.EUTRACGIPWSFailedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailedCellIDList: " + err.Error())
		}
	case ngapType.PWSFailedCellIDListPresentNRCGIPWSFailedList:
		if value.NRCGIPWSFailedList == nil {
			return binData, bitEnd, errors.New("PWSFailedCellIDList: NRCGIPWSFailedList: NIL")
		}
		binData, bitEnd, err = NRCGIList(*value.NRCGIPWSFailedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailedCellIDList: " + err.Error())
		}
	case ngapType.PWSFailedCellIDListPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("PWSFailedCellIDList: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerPWSFailedCellIDListExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailedCellIDList: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSFailedCellIDList: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 82 - RANNodeName
func RANNodeName (value ngapType.RANNodeName, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 150 {
		return binData, bitEnd, errors.New("RANNodeName: string: sizeExt,sizeLB:1,sizeUB:150")
	}
	binData, bitEnd = EncodeUint64(uint64(len(value.Value)-1), 9, binData, bitEnd)
	binData, bitEnd = EncodeBytes([]byte(value.Value), binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 83 - RANPagingPriority
func RANPagingPriority (value ngapType.RANPagingPriority, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 256 {
		return binData, bitEnd, errors.New("RANPagingPriority: int64: valueLB:1,valueUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 84 - RANStatusTransferTransparentContainer
func RANStatusTransferTransparentContainer (value ngapType.RANStatusTransferTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 85 - RANUENGAPID
func RANUENGAPID (value ngapType.RANUENGAPID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4294967295 {
		return binData, bitEnd, errors.New("RANUENGAPID: int64: valueLB:0,valueUB:4294967295")
	}
	binValue := ConvertInt64ToBytes(value.Value)
	binData, bitEnd = EncodeUint64(uint64(len(binValue))-1, 2, binData, bitEnd)
	binData, bitEnd = EncodeBytes(binValue, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 86 - RelativeAMFCapacity
func RelativeAMFCapacity (value ngapType.RelativeAMFCapacity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("RelativeAMFCapacity: int64: valueLB:0,valueUB:255")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 8, binData, 8)
	return binData, bitEnd, nil
}

// ProtocolIEField 87 - RepetitionPeriod
func RepetitionPeriod (value ngapType.RepetitionPeriod, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 131071 {
		return binData, bitEnd, errors.New("RepetitionPeriod: int64: valueLB:0,valueUB:131071")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 88 - ResetType
func ResetType (value ngapType.ResetType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.ResetTypePresentNGInterface:
		if value.NGInterface == nil {
			return binData, bitEnd, errors.New("ResetType: NGInterface: NIL")
		}
		binData, bitEnd, err = ResetAll(*value.NGInterface, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ResetType: " + err.Error())
		}
	case ngapType.ResetTypePresentPartOfNGInterface:
		if value.PartOfNGInterface == nil {
			return binData, bitEnd, errors.New("ResetType: PartOfNGInterface: NIL")
		}
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionList(*value.PartOfNGInterface, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ResetType: " + err.Error())
		}
	case ngapType.ResetTypePresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("ResetType: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerResetTypeExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ResetType: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("ResetType: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 89 - RoutingID
func RoutingID (value ngapType.RoutingID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 90 - RRCEstablishmentCause
func RRCEstablishmentCause (value ngapType.RRCEstablishmentCause, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 9 {
		return binData, bitEnd, errors.New("RRCEstablishmentCause: Enumerated: valueExt,valueLB:0,valueUB:9")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 91 - RRCInactiveTransitionReportRequest
func RRCInactiveTransitionReportRequest (value ngapType.RRCInactiveTransitionReportRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportRequest: Enumerated: valueExt,valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 3, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 92 - RRCState
func RRCState (value ngapType.RRCState, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("RRCState: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 93 - SecurityContext
func SecurityContext (value ngapType.SecurityContext, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 94 - SecurityKey
func SecurityKey (value ngapType.SecurityKey, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 256 || len(value.Value.Bytes) != 32 {
		return binData, bitEnd, errors.New("SecurityKey: BitString: sizeLB:256,sizeUB:256")
	}
	binData, bitEnd = append(binData, value.Value.Bytes...), 8
	return binData, bitEnd, nil
}

// ProtocolIEField 95 - SerialNumber
func SerialNumber (value ngapType.SerialNumber, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("SerialNumber: BitString: sizeLB:16,sizeUB:16")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 96 - ServedGUAMIList

func ServedGUAMIList (value ngapType.ServedGUAMIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("ServedGUAMIList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 8, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ServedGUAMIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ServedGUAMIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 97 - SliceSupportList
func SliceSupportList (value ngapType.SliceSupportList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 1024 {
		return binData, bitEnd, errors.New("SliceSupportList: List: valueExt,sizeLB:1,sizeUB:1024")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 16, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SliceSupportItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SliceSupportList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 98 - SONConfigurationTransferDL
func SONConfigurationTransferDL (value ngapType.SONConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return SONConfigurationTransfer(value, binData, bitEnd)
}

func SONConfigurationTransfer (value ngapType.SONConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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
	binData, bitEnd, err = XnTNLConfigurationInfo(value.XnTNLConfigurationInfo, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerSONConfigurationTransferExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONConfigurationTransfer: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 99 - SONConfigurationTransferUL
func SONConfigurationTransferUL (value ngapType.SONConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return SONConfigurationTransfer(value, binData, bitEnd)
}

// ProtocolIEField 100 - SourceAMFUENGAPID
func SourceAMFUENGAPID (value ngapType.AMFUENGAPID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return AMFUENGAPID(value, binData, bitEnd)
}

// ProtocolIEField 101 - SourceToTargetTransparentContainer
func SourceToTargetTransparentContainer (value ngapType.SourceToTargetTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	binValue, bitEnd := EncodeOctetString(value.Value, binData, bitEnd)
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, nil
}

// ProtocolIEField 102 - SupportedTAList
func SupportedTAList (value ngapType.SupportedTAList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("SupportedTAList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 8, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SupportedTAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SupportedTAList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 103 - TAIListForPaging
func TAIListForPaging (value ngapType.TAIListForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("TAIListForPaging: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIListForPagingItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForPaging: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 104 - TAIListForRestart
func TAIListForRestart (value ngapType.TAIListForRestart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 2048 {
		return binData, bitEnd, errors.New("TAIListForRestart: List: valueExt,sizeLB:1,sizeUB:2048")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAI(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TAIListForRestart: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 105 - TargetID
func TargetID (value ngapType.TargetID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.TargetIDPresentTargetRANNodeID:
		if value.TargetRANNodeID == nil {
			return binData, bitEnd, errors.New("TargetID: TargetRANNodeID: NIL")
		}
		binData, bitEnd, err = TargetRANNodeID(*value.TargetRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetID: " + err.Error())
		}
	case ngapType.TargetIDPresentTargeteNBID:
		if value.TargeteNBID == nil {
			return binData, bitEnd, errors.New("TargetID: TargeteNBID: NIL")
		}
		binData, bitEnd, err = TargeteNBID(*value.TargeteNBID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetID: " + err.Error())
		}
	case ngapType.TargetIDPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("TargetID: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerTargetIDExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetID: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("TargetID: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 106 - TargetToSourceTransparentContainer
func TargetToSourceTransparentContainer (value ngapType.TargetToSourceTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 107 - TimeToWait
func TimeToWait (value ngapType.TimeToWait, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 5 {
		return binData, bitEnd, errors.New("TimeToWait: Enumerated: valueExt,valueLB:0,valueUB:5")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 108 - TraceActivation
func TraceActivation (value ngapType.TraceActivation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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
	binData, bitEnd = EncodeUint64(0, 1, binData, bitEnd)
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

// ProtocolIEField 109 - TraceCollectionEntityIPAddress
func TraceCollectionEntityIPAddress (value ngapType.TransportLayerAddress, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return TransportLayerAddress(value, binData, bitEnd)
}

// ProtocolIEField 110 - UEAggregateMaximumBitRate
func UEAggregateMaximumBitRate (value ngapType.UEAggregateMaximumBitRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 111 - UEAssociatedLogicalNGConnectionList
func UEAssociatedLogicalNGConnectionList (value ngapType.UEAssociatedLogicalNGConnectionList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65536 {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionList: List: valueExt,sizeLB:1,sizeUB:65536")
	}
	if numItem < 256 {
		binData, bitEnd = EncodeUint64(uint64(numItem), 8, binData, 8)
	} else {
		binData, bitEnd = EncodeUint64(uint64(numItem), 16, binData, 8)
	}
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 112 - UEContextRequest
func UEContextRequest (value ngapType.UEContextRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("UEContextRequest: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 113 - UENGAPIDs
func UENGAPIDs (value ngapType.UENGAPIDs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.UENGAPIDsPresentUENGAPIDPair:
		if value.UENGAPIDPair == nil {
			return binData, bitEnd, errors.New("UENGAPIDs: UENGAPIDPair: NIL")
		}
		binData, bitEnd, err = UENGAPIDPair(*value.UENGAPIDPair, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UENGAPIDs: " + err.Error())
		}
	case ngapType.UENGAPIDsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UENGAPIDs: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UENGAPIDs: " + err.Error())
		}
	case ngapType.UENGAPIDsPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("UENGAPIDs: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerUENGAPIDsExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UENGAPIDs: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UENGAPIDs: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 114 - UEPagingIdentity
func UEPagingIdentity (value ngapType.UEPagingIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.UEPagingIdentityPresentFiveGSTMSI:
		if value.FiveGSTMSI == nil {
			return binData, bitEnd, errors.New("UENGAPIDs: FiveGSTMSI: NIL")
		}
		binData, bitEnd, err = FiveGSTMSI(*value.FiveGSTMSI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEPagingIdentity: " + err.Error())
		}
	case ngapType.UEPagingIdentityPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("UENGAPIDs: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerUEPagingIdentityExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEPagingIdentity: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEPagingIdentity: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 115 - UEPresenceInAreaOfInterestList
func UEPresenceInAreaOfInterestList (value ngapType.UEPresenceInAreaOfInterestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 64 {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestList: List: valueExt,sizeLB:1,sizeUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEPresenceInAreaOfInterestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 116 - UERadioCapability
func UERadioCapability (value ngapType.UERadioCapability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return append(binData, EncodeLengthValue(value.Value)...), 8, nil
}

// ProtocolIEField 117 - UERadioCapabilityForPaging
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

// ProtocolIEField 118 - UESecurityCapabilities
func UESecurityCapabilities (value ngapType.UESecurityCapabilities, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 119 - UnavailableGUAMIList
func UnavailableGUAMIList (value ngapType.UnavailableGUAMIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("UnavailableGUAMIList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UnavailableGUAMIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnavailableGUAMIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 120 - UserLocationInformation
func UserLocationInformation (value ngapType.UserLocationInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.UserLocationInformationPresentUserLocationInformationEUTRA:
		if value.UserLocationInformationEUTRA == nil {
			return binData, bitEnd, errors.New("UserLocationInformation: UserLocationInformationEUTRA: NIL")
		}
		binData, bitEnd, err = UserLocationInformationEUTRA(*value.UserLocationInformationEUTRA, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformation: " + err.Error())
		}
	case ngapType.UserLocationInformationPresentUserLocationInformationNR:
		if value.UserLocationInformationNR == nil {
			return binData, bitEnd, errors.New("UserLocationInformation: UserLocationInformationNR: NIL")
		}
		binData, bitEnd, err = UserLocationInformationNR(*value.UserLocationInformationNR, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformation: " + err.Error())
		}
	case ngapType.UserLocationInformationPresentUserLocationInformationN3IWF:
		if value.UserLocationInformationN3IWF == nil {
			return binData, bitEnd, errors.New("UserLocationInformation: UserLocationInformationN3IWF: NIL")
		}
		binData, bitEnd, err = UserLocationInformationN3IWF(*value.UserLocationInformationN3IWF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformation: " + err.Error())
		}
	case ngapType.UserLocationInformationPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("UserLocationInformation: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerUserLocationInformationExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UserLocationInformation: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UserLocationInformation: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 121 - WarningAreaList
func WarningAreaList (value ngapType.WarningAreaList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.WarningAreaListPresentEUTRACGIListForWarning:
		if value.EUTRACGIListForWarning == nil {
			return binData, bitEnd, errors.New("WarningAreaList: EUTRACGIListForWarning: NIL")
		}
		binData, bitEnd, err = EUTRACGIListForWarning(*value.EUTRACGIListForWarning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WarningAreaList: " + err.Error())
		}
	case ngapType.WarningAreaListPresentNRCGIListForWarning:
		if value.NRCGIListForWarning == nil {
			return binData, bitEnd, errors.New("WarningAreaList: NRCGIListForWarning: NIL")
		}
		binData, bitEnd, err = NRCGIListForWarning(*value.NRCGIListForWarning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WarningAreaList: " + err.Error())
		}
	case ngapType.WarningAreaListPresentTAIListForWarning:
		if value.TAIListForWarning == nil {
			return binData, bitEnd, errors.New("WarningAreaList: TAIListForWarning: NIL")
		}
		binData, bitEnd, err = TAIListForWarning(*value.TAIListForWarning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WarningAreaList: " + err.Error())
		}
	case ngapType.WarningAreaListPresentEmergencyAreaIDList:
		if value.EmergencyAreaIDList == nil {
			return binData, bitEnd, errors.New("WarningAreaList: EmergencyAreaIDList: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDList(*value.EmergencyAreaIDList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WarningAreaList: " + err.Error())
		}
	case ngapType.WarningAreaListPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("WarningAreaList: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerWarningAreaListExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WarningAreaList: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("WarningAreaList: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 122 - WarningMessageContents
func WarningMessageContents (value ngapType.WarningMessageContents, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 9600 {
		return binData, bitEnd, errors.New("WarningMessageContents: OctetString: sizeLB:1,sizeUB:9600")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 123 - WarningSecurityInfo
func WarningSecurityInfo (value ngapType.WarningSecurityInfo, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 50 {
		return binData, bitEnd, errors.New("WarningSecurityInfo: OctetString: sizeLB:50,sizeUB:50")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 124 - WarningType
func WarningType (value ngapType.WarningType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("WarningType: OctetString: sizeLB:2,sizeUB:2")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 125 - AdditionalULNGUUPTNLInformation
func AdditionalULNGUUPTNLInformation (value ngapType.UPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return UPTransportLayerInformation(value, binData, bitEnd)
}

func UPTransportLayerInformation (value ngapType.UPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.UPTransportLayerInformationPresentGTPTunnel:
		if value.GTPTunnel == nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformation: GTPTunnel: NIL")
		}
		binData, bitEnd, err = GTPTunnel(*value.GTPTunnel, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformation: " + err.Error())
		}
	case ngapType.UPTransportLayerInformationPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformation: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerUPTransportLayerInformationExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UPTransportLayerInformation: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UPTransportLayerInformation: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 126 - DataForwardingNotPossible
func DataForwardingNotPossible (value ngapType.DataForwardingNotPossible, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("DataForwardingNotPossible: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 127 - DLNGUUPTNLInformation
func DLNGUUPTNLInformation (value ngapType.UPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return UPTransportLayerInformation(value, binData, bitEnd)
}

// ProtocolIEField 128 - NetworkInstance
func NetworkInstance (value ngapType.NetworkInstance, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 256 {
		return binData, bitEnd, errors.New("NetworkInstance: int64: valueExt,valueLB:1,valueUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 129 - PDUSessionAggregateMaximumBitRate
func PDUSessionAggregateMaximumBitRate (value ngapType.PDUSessionAggregateMaximumBitRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 130 - PDUSessionResourceFailedToModifyListModCfm
func PDUSessionResourceFailedToModifyListModCfm (value ngapType.PDUSessionResourceFailedToModifyListModCfm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModCfm: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModCfm(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyListModCfm: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 131 - PDUSessionResourceFailedToSetupListCxtFail
func PDUSessionResourceFailedToSetupListCxtFail (value ngapType.PDUSessionResourceFailedToSetupListCxtFail, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtFail: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtFail(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupListCxtFail: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 132 - PDUSessionResourceListCxtRelReq
func PDUSessionResourceListCxtRelReq (value ngapType.PDUSessionResourceListCxtRelReq, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelReq: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemCxtRelReq(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceListCxtRelReq: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 133 - PDUSessionType

func PDUSessionType (value ngapType.PDUSessionType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4 {
		return binData, bitEnd, errors.New("PDUSessionType: Enumerated: valueExt,valueLB:0,valueUB:4")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

// ProtocolIEField 134 - QosFlowAddOrModifyRequestList
func QosFlowAddOrModifyRequestList (value ngapType.QosFlowAddOrModifyRequestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 64 {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestList: List: valueExt,sizeLB:1,sizeUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowAddOrModifyRequestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 135 - QosFlowSetupRequestList
func QosFlowSetupRequestList (value ngapType.QosFlowSetupRequestList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 64 {
		return binData, bitEnd, errors.New("QosFlowSetupRequestList: List: valueExt,sizeLB:1,sizeUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowSetupRequestItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowSetupRequestList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 136 - QosFlowToReleaseList
func QosFlowToReleaseList (value ngapType.QosFlowList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return QosFlowList(value, binData, bitEnd)
}

func QosFlowList (value ngapType.QosFlowList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 64 {
		return binData, bitEnd, errors.New("QosFlowList: List: valueExt,sizeLB:1,sizeUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosFlowList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 137 - SecurityIndication
func SecurityIndication (value ngapType.SecurityIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.MaximumIntegrityProtectedDataRate != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
	binData, bitEnd, err = IntegrityProtectionIndication(value.IntegrityProtectionIndication, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndication: " + err.Error())
	}
	binData, bitEnd, err = ConfidentialityProtectionIndication(value.ConfidentialityProtectionIndication, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndication: " + err.Error())
	}
	if value.MaximumIntegrityProtectedDataRate != nil {
		binData, bitEnd, err = MaximumIntegrityProtectedDataRate(*value.MaximumIntegrityProtectedDataRate, binData, bitEnd)
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

// ProtocolIEField 138 - ULNGUUPTNLInformation
func ULNGUUPTNLInformation (value ngapType.UPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return UPTransportLayerInformation(value, binData, bitEnd)
}

// ProtocolIEField 139 - ULNGUUPTNLModifyList
func ULNGUUPTNLModifyList (value ngapType.ULNGUUPTNLModifyList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem != 4 {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyList: List: valueExt,sizeLB:4,sizeUB:4")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ULNGUUPTNLModifyItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ULNGUUPTNLModifyList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 140 - WarningAreaCoordinates
func WarningAreaCoordinates (value ngapType.WarningAreaCoordinates, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 1024 {
		return binData, bitEnd, errors.New("WarningAreaCoordinates: OctetString: sizeLB:1,sizeUB:1024")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

