// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"
	"vht5gc/lib/ngap/ngapType"
)

// ProtocolIEField 0 - AllowedNSSAI
func SST (value ngapType.SST, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 1 {
		return binData, bitEnd, errors.New("SST: OctetString: sizeLB:1,sizeUB:1")
	}
	binData, bitEnd = EncodeOctetStringEx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func SD (value ngapType.SD, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("SD: OctetString: sizeLB:3,sizeUB:3")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerSNSSAIExtIEs (value ngapType.ProtocolExtensionContainerSNSSAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSNSSAIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SNSSAIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSNSSAIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllowedNSSAIItemExtIEs (value ngapType.AllowedNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AllowedNSSAIItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 1 - AMFName

// ProtocolIEField 2 - AMFOverloadResponse

// ProtocolIEField 3 - AMFSetID

// ProtocolIEField 4 - AMFTNLAssociationFailedToSetupList
func TransportLayerAddress (value ngapType.TransportLayerAddress, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength < 1 || value.Value.BitLength > 160 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("TransportLayerAddress: BitString: sizeExt,sizeLB:1,sizeUB:160")
	}
	binData, bitEnd = EncodeBitStringExt(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolIESingleContainerCPTransportLayerInformationExtIEs (value ngapType.ProtocolIESingleContainerCPTransportLayerInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func TNLAssociationItemExtIEs (value ngapType.TNLAssociationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TNLAssociationItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 5 - AMFTNLAssociationSetupList
func AMFTNLAssociationSetupItemExtIEs (value ngapType.AMFTNLAssociationSetupItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AMFTNLAssociationSetupItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 6 - AMFTNLAssociationToAddList
func AMFTNLAssociationToAddItemExtIEs (value ngapType.AMFTNLAssociationToAddItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AMFTNLAssociationToAddItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 7 - AMFTNLAssociationToRemoveList
func AMFTNLAssociationToRemoveItemExtIEs (value ngapType.AMFTNLAssociationToRemoveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AMFTNLAssociationToRemoveItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 8 - AMFTNLAssociationToUpdateList
func AMFTNLAssociationToUpdateItemExtIEs (value ngapType.AMFTNLAssociationToUpdateItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AMFTNLAssociationToUpdateItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 9 - AMFTrafficLoadReductionIndication

// ProtocolIEField 10 - AMFUENGAPID

// ProtocolIEField 11 - AssistanceDataForPaging
func RecommendedCellList (value ngapType.RecommendedCellList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("RecommendedCellList: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedCellItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedCellList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedCellsForPagingExtIEs (value ngapType.ProtocolExtensionContainerRecommendedCellsForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellsForPagingExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedCellsForPagingExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellsForPagingExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AssistanceDataForRecommendedCellsExtIEs (value ngapType.AssistanceDataForRecommendedCellsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AssistanceDataForRecommendedCellsExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func PagingAttemptInformationExtIEs (value ngapType.PagingAttemptInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PagingAttemptInformationExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AssistanceDataForPagingExtIEsExtensionValue (value ngapType.AssistanceDataForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 12 - BroadcastCancelledAreaList
func NumberOfBroadcasts (value ngapType.NumberOfBroadcasts, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 65535 {
		return binData, bitEnd, errors.New("NumberOfBroadcasts: int64: valueLB:0,valueUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDCancelledEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
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
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func CancelledCellsInTAIEUTRA (value ngapType.CancelledCellsInTAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInTAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAICancelledEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaID (value ngapType.EmergencyAreaID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("EmergencyAreaID: OctetString: sizeLB:3,sizeUB:3")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func CancelledCellsInEAIEUTRA (value ngapType.CancelledCellsInEAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInEAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDCancelledEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellIDCancelledNRItemExtIEs (value ngapType.ProtocolExtensionContainerCellIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledNRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDCancelledNRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDCancelledNRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINR (value ngapType.CancelledCellsInTAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CancelledCellsInTAINR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInTAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInTAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAICancelledNRItemExtIEs (value ngapType.ProtocolExtensionContainerTAICancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledNRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAICancelledNRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAICancelledNRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINR (value ngapType.CancelledCellsInEAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CancelledCellsInEAINR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInEAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CancelledCellsInEAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs (value ngapType.ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDCancelledNRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 13 - BroadcastCompletedAreaList
func ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDBroadcastEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRA (value ngapType.CompletedCellsInTAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInTAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIBroadcastEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRA (value ngapType.CompletedCellsInEAIEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRA: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInEAIEUTRAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRA: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDBroadcastEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs (value ngapType.ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CellIDBroadcastNRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINR (value ngapType.CompletedCellsInTAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CompletedCellsInTAINR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInTAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInTAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTAIBroadcastNRItemExtIEs (value ngapType.ProtocolExtensionContainerTAIBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastNRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIBroadcastNRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIBroadcastNRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINR (value ngapType.CompletedCellsInEAINR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("CompletedCellsInEAINR: List: valueExt,sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInEAINRItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CompletedCellsInEAINR: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs (value ngapType.ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EmergencyAreaIDBroadcastNRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 14 - CancelAllWarningMessages

// ProtocolIEField 15 - Cause

// ProtocolIEField 16 - CellIDListForRestart

// ProtocolIEField 17 - ConcurrentWarningMessageInd

// ProtocolIEField 18 - CoreNetworkAssistanceInformation
func ProtocolExtensionContainerTAIListForInactiveItemExtIEs (value ngapType.ProtocolExtensionContainerTAIListForInactiveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForInactiveItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIListForInactiveItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForInactiveItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedActivityPeriod (value ngapType.ExpectedActivityPeriod, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 181 {
		return binData, bitEnd, errors.New("ExpectedActivityPeriod: valueLB:1,valueUB:181")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value)-1, 9, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedIdlePeriod (value ngapType.ExpectedIdlePeriod, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 181 {
		return binData, bitEnd, errors.New("ExpectedActivityPeriod: valueLB:1,valueUB:181")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value)-1, 9, binData, bitEnd)
	return binData, bitEnd, nil
}

func SourceOfUEActivityBehaviourInformation (value ngapType.SourceOfUEActivityBehaviourInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("SourceOfUEActivityBehaviourInformation: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs (value ngapType.ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 24, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ExpectedUEActivityBehaviourExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
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
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
	binData, bitEnd, err = NGRANCGI(value.NGRANCGI, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItem: " + err.Error())
	}
	if value.TimeStayedInCell != nil {
		if *value.TimeStayedInCell < 0 || *value.TimeStayedInCell > 4095 {
			return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItem: TimeStayedInCell: valueLB:0,valueUB:4095")
		}
		binData, bitEnd = EncodeUint64(uint64(*value.TimeStayedInCell), 16, binData, 8)
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedUEBehaviourExtIEs (value ngapType.ExpectedUEBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ExpectedUEBehaviourExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CoreNetworkAssistanceInformationExtIEsExtensionValue (value ngapType.CoreNetworkAssistanceInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 19 - CriticalityDiagnostics
func ProtocolIEID (value ngapType.ProtocolIEID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 65535 {
		return binData, bitEnd, errors.New("ProtocolIEID: int64: valueLB:0,valueUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 16, binData, 8)
	return binData, bitEnd, nil
}

func TypeOfError (value ngapType.TypeOfError, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("TypeOfError: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs (value ngapType.ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CriticalityDiagnosticsIEItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CriticalityDiagnosticsExtIEsExtensionValue (value ngapType.CriticalityDiagnosticsExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 20 - DataCodingScheme

// ProtocolIEField 21 - DefaultPagingDRX

// ProtocolIEField 22 - DirectForwardingPathAvailability

// ProtocolIEField 23 - EmergencyAreaIDListForRestart

// ProtocolIEField 24 - EmergencyFallbackIndicator
func EmergencyFallbackIndicatorExtIEsExtensionValue (value ngapType.EmergencyFallbackIndicatorExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 25 - EUTRACGI
func EUTRACGIExtIEsExtensionValue (value ngapType.EUTRACGIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EUTRACGIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 26 - FiveGSTMSI
func FiveGSTMSIExtIEsExtensionValue (value ngapType.FiveGSTMSIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 27 - GlobalRANNodeID
func ProtocolIESingleContainerGNBIDExtIEs (value ngapType.ProtocolIESingleContainerGNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func GlobalGNBIDExtIEs (value ngapType.GlobalGNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = GlobalGNBIDExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerNgENBIDExtIEs (value ngapType.ProtocolIESingleContainerNgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func GlobalNgENBIDExtIEs (value ngapType.GlobalNgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = GlobalNgENBIDExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerN3IWFIDExtIEs (value ngapType.ProtocolIESingleContainerN3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func GlobalN3IWFIDExtIEs (value ngapType.GlobalN3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = GlobalN3IWFIDExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 28 - GUAMI
func GUAMIExtIEsExtensionValue (value ngapType.GUAMIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GUAMIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 29 - HandoverType

// ProtocolIEField 30 - IMSVoiceSupportIndicator

// ProtocolIEField 31 - IndexToRFSP

// ProtocolIEField 32 - InfoOnRecommendedCellsAndRANNodesForPaging
func RecommendedRANNodeItem (value ngapType.RecommendedRANNodeItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func RecommendedRANNodesForPagingExtIEs (value ngapType.RecommendedRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = RecommendedRANNodesForPagingExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue (value ngapType.InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 33 - LocationReportingRequestType
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
	binData, bitEnd = EncodeUint64(uint64(option), 5, binData, bitEnd)
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

func LocationReportingReferenceID (value ngapType.LocationReportingReferenceID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 64 {
		return binData, bitEnd, errors.New("LocationReportingReferenceID: int64: valueExt,valueLB:1,valueUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value)-1, 7, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerAreaOfInterestItemExtIEs (value ngapType.ProtocolExtensionContainerAreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func LocationReportingRequestTypeExtIEsExtensionValue (value ngapType.LocationReportingRequestTypeExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 34 - MaskedIMEISV

// ProtocolIEField 35 - MessageIdentifier

// ProtocolIEField 36 - MobilityRestrictionList
func RATRestrictionInformation (value ngapType.RATRestrictionInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("RATRestrictionInformation: BitString: sizeExt,sizeLB:8,sizeUB:8")
	}
	binData, bitEnd = EncodeBitString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerRATRestrictionsItemExtIEs (value ngapType.ProtocolExtensionContainerRATRestrictionsItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRATRestrictionsItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RATRestrictionsItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRATRestrictionsItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ForbiddenTACs (value ngapType.ForbiddenTACs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 4096 {
		return binData, bitEnd, errors.New("ForbiddenTACs: List: sizeLB:1,sizeUB:4096")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 16, binData, 8)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAC(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ForbiddenTACs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs (value ngapType.ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ForbiddenAreaInformationItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllowedTACs (value ngapType.AllowedTACs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("AllowedTACs: List: sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAC(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AllowedTACs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NotAllowedTACs (value ngapType.NotAllowedTACs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("NotAllowedTACs: List: sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAC(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NotAllowedTACs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerServiceAreaInformationItemExtIEs (value ngapType.ProtocolExtensionContainerServiceAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerServiceAreaInformationItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ServiceAreaInformationItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerServiceAreaInformationItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func MobilityRestrictionListExtIEsExtensionValue (value ngapType.MobilityRestrictionListExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEsExtensionValue: Present: INVALID")
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

// ProtocolIEField 44 - NGRANTraceID

// ProtocolIEField 45 - NRCGI
func NRCGIExtIEsExtensionValue (value ngapType.NRCGIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NRCGIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 46 - NRPPaPDU

// ProtocolIEField 47 - NumberOfBroadcastsRequested

// ProtocolIEField 48 - OldAMF

// ProtocolIEField 49 - OverloadStartNSSAIList
func SliceOverloadItem (value ngapType.SliceOverloadItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func OverloadStartNSSAIItemExtIEs (value ngapType.OverloadStartNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = OverloadStartNSSAIItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 50 - PagingDRX

// ProtocolIEField 51 - PagingOrigin

// ProtocolIEField 52 - PagingPriority

// ProtocolIEField 53 - PDUSessionResourceAdmittedList
func PDUSessionResourceAdmittedItemExtIEs (value ngapType.PDUSessionResourceAdmittedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceAdmittedItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 54 - PDUSessionResourceFailedToModifyListModRes
func PDUSessionResourceFailedToModifyItemModResExtIEs (value ngapType.PDUSessionResourceFailedToModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 55 - PDUSessionResourceFailedToSetupListCxtRes
func PDUSessionResourceFailedToSetupItemCxtResExtIEs (value ngapType.PDUSessionResourceFailedToSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 56 - PDUSessionResourceFailedToSetupListHOAck
func PDUSessionResourceFailedToSetupItemHOAckExtIEs (value ngapType.PDUSessionResourceFailedToSetupItemHOAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 57 - PDUSessionResourceFailedToSetupListPSReq
func PDUSessionResourceFailedToSetupItemPSReqExtIEs (value ngapType.PDUSessionResourceFailedToSetupItemPSReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 58 - PDUSessionResourceFailedToSetupListSURes
func PDUSessionResourceFailedToSetupItemSUResExtIEs (value ngapType.PDUSessionResourceFailedToSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 59 - PDUSessionResourceHandoverList
func PDUSessionResourceHandoverItemExtIEs (value ngapType.PDUSessionResourceHandoverItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceHandoverItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 60 - PDUSessionResourceListCxtRelCpl
func PDUSessionResourceItemCxtRelCplExtIEs (value ngapType.PDUSessionResourceItemCxtRelCplExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceItemCxtRelCplExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 61 - PDUSessionResourceListHORqd
func PDUSessionResourceItemHORqdExtIEs (value ngapType.PDUSessionResourceItemHORqdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceItemHORqdExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 62 - PDUSessionResourceModifyListModCfm
func PDUSessionResourceModifyItemModCfmExtIEs (value ngapType.PDUSessionResourceModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceModifyItemModCfmExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 63 - PDUSessionResourceModifyListModInd
func PDUSessionResourceModifyItemModIndExtIEs (value ngapType.PDUSessionResourceModifyItemModIndExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceModifyItemModIndExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 64 - PDUSessionResourceModifyListModReq
func PDUSessionResourceModifyItemModReqExtIEs (value ngapType.PDUSessionResourceModifyItemModReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceModifyItemModReqExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 65 - PDUSessionResourceModifyListModRes
func PDUSessionResourceModifyItemModResExtIEs (value ngapType.PDUSessionResourceModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceModifyItemModResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 66 - PDUSessionResourceNotifyList
func PDUSessionResourceNotifyItemExtIEs (value ngapType.PDUSessionResourceNotifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceNotifyItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 67 - PDUSessionResourceReleasedListNot
func PDUSessionResourceReleasedItemNotExtIEs (value ngapType.PDUSessionResourceReleasedItemNotExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceReleasedItemNotExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 68 - PDUSessionResourceReleasedListPSAck
func PDUSessionResourceReleasedItemPSAckExtIEs (value ngapType.PDUSessionResourceReleasedItemPSAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 69 - PDUSessionResourceReleasedListPSFail
func PDUSessionResourceReleasedItemPSFailExtIEs (value ngapType.PDUSessionResourceReleasedItemPSFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 70 - PDUSessionResourceReleasedListRelRes
func PDUSessionResourceReleasedItemRelResExtIEs (value ngapType.PDUSessionResourceReleasedItemRelResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceReleasedItemRelResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 71 - PDUSessionResourceSetupListCxtReq
func PDUSessionResourceSetupItemCxtReqExtIEs (value ngapType.PDUSessionResourceSetupItemCxtReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 72 - PDUSessionResourceSetupListCxtRes
func PDUSessionResourceSetupItemCxtResExtIEs (value ngapType.PDUSessionResourceSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceSetupItemCxtResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 73 - PDUSessionResourceSetupListHOReq
func PDUSessionResourceSetupItemHOReqExtIEs (value ngapType.PDUSessionResourceSetupItemHOReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceSetupItemHOReqExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 74 - PDUSessionResourceSetupListSUReq
func PDUSessionResourceSetupItemSUReqExtIEs (value ngapType.PDUSessionResourceSetupItemSUReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceSetupItemSUReqExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 75 - PDUSessionResourceSetupListSURes
func PDUSessionResourceSetupItemSUResExtIEs (value ngapType.PDUSessionResourceSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceSetupItemSUResExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 76 - PDUSessionResourceToBeSwitchedDLList
func PDUSessionResourceToBeSwitchedDLItemExtIEs (value ngapType.PDUSessionResourceToBeSwitchedDLItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 77 - PDUSessionResourceSwitchedList
func PDUSessionResourceSwitchedItemExtIEs (value ngapType.PDUSessionResourceSwitchedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceSwitchedItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 78 - PDUSessionResourceToReleaseListHOCmd
func PDUSessionResourceToReleaseItemHOCmdExtIEs (value ngapType.PDUSessionResourceToReleaseItemHOCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 79 - PDUSessionResourceToReleaseListRelCmd
func PDUSessionResourceToReleaseItemRelCmdExtIEs (value ngapType.PDUSessionResourceToReleaseItemRelCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 80 - PLMNSupportList
func PLMNSupportItemExtIEs (value ngapType.PLMNSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PLMNSupportItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 81 - PWSFailedCellIDList

// ProtocolIEField 82 - RANNodeName

// ProtocolIEField 83 - RANPagingPriority

// ProtocolIEField 84 - RANStatusTransferTransparentContainer
func DRBID (value ngapType.DRBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 32 {
		return binData, bitEnd, errors.New("DRBID: int64: valueExt,valueLB:1,valueUB:32")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func DRBStatusUL (value ngapType.DRBStatusUL, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.DRBStatusULPresentDRBStatusUL12:
		if value.DRBStatusUL12 == nil {
			return binData, bitEnd, errors.New("Cause: DRBStatusUL12: NIL")
		}
		binData, bitEnd, err = DRBStatusUL12(*value.DRBStatusUL12, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL: " + err.Error())
		}
	case ngapType.DRBStatusULPresentDRBStatusUL18:
		if value.DRBStatusUL18 == nil {
			return binData, bitEnd, errors.New("Cause: DRBStatusUL18: NIL")
		}
		binData, bitEnd, err = DRBStatusUL18(*value.DRBStatusUL18, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL: " + err.Error())
		}
	case ngapType.DRBStatusULPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("Cause: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerDRBStatusULExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DRBStatusUL: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusDL (value ngapType.DRBStatusDL, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.DRBStatusDLPresentDRBStatusDL12:
		if value.DRBStatusDL12 == nil {
			return binData, bitEnd, errors.New("Cause: DRBStatusDL12: NIL")
		}
		binData, bitEnd, err = DRBStatusDL12(*value.DRBStatusDL12, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL: " + err.Error())
		}
	case ngapType.DRBStatusDLPresentDRBStatusDL18:
		if value.DRBStatusDL18 == nil {
			return binData, bitEnd, errors.New("Cause: DRBStatusDL18: NIL")
		}
		binData, bitEnd, err = DRBStatusDL18(*value.DRBStatusDL18, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL: " + err.Error())
		}
	case ngapType.DRBStatusDLPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("Cause: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerDRBStatusDLExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DRBStatusDL: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs (value ngapType.ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBsSubjectToStatusTransferItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RANStatusTransferTransparentContainerExtIEsExtensionValue (value ngapType.RANStatusTransferTransparentContainerExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 85 - RANUENGAPID

// ProtocolIEField 86 - RelativeAMFCapacity

// ProtocolIEField 87 - RepetitionPeriod

// ProtocolIEField 88 - ResetType

// ProtocolIEField 89 - RoutingID

// ProtocolIEField 90 - RRCEstablishmentCause

// ProtocolIEField 91 - RRCInactiveTransitionReportRequest

// ProtocolIEField 92 - RRCState

// ProtocolIEField 93 - SecurityContext
func SecurityContextExtIEsExtensionValue (value ngapType.SecurityContextExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SecurityContextExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 94 - SecurityKey

// ProtocolIEField 95 - SerialNumber

// ProtocolIEField 96 - ServedGUAMIList
func ServedGUAMIItemExtIEs (value ngapType.ServedGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ServedGUAMIItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 97 - SliceSupportList
func SliceSupportItemExtIEs (value ngapType.SliceSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SliceSupportItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceSupportItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 98 - SONConfigurationTransferDL
func TargetRANNodeIDExtIEs (value ngapType.TargetRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TargetRANNodeIDExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func SourceRANNodeIDExtIEs (value ngapType.SourceRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SourceRANNodeIDExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSONInformationReplyExtIEs (value ngapType.ProtocolExtensionContainerSONInformationReplyExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSONInformationReplyExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SONInformationReplyExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSONInformationReplyExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

func XnTNLConfigurationInfoExtIEs (value ngapType.XnTNLConfigurationInfoExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	binData, bitEnd, err = XnTNLConfigurationInfoExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func SONConfigurationTransferExtIEsExtensionValue (value ngapType.SONConfigurationTransferExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 99 - SONConfigurationTransferUL

// ProtocolIEField 100 - SourceAMFUENGAPID

// ProtocolIEField 101 - SourceToTargetTransparentContainer

// ProtocolIEField 102 - SupportedTAList
func BroadcastPLMNItem (value ngapType.BroadcastPLMNItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func SupportedTAItemExtIEs (value ngapType.SupportedTAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SupportedTAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SupportedTAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 103 - TAIListForPaging
func TAIListForPagingItemExtIEs (value ngapType.TAIListForPagingItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAIListForPagingItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 104 - TAIListForRestart

// ProtocolIEField 105 - TargetID
func EPSTAC (value ngapType.EPSTAC, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("EPSTAC: OctetString: sizeLB:2,sizeUB:2")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerEPSTAIExtIEs (value ngapType.ProtocolExtensionContainerEPSTAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerEPSTAIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = EPSTAIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerEPSTAIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TargeteNBIDExtIEs (value ngapType.TargeteNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TargeteNBIDExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargeteNBIDExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 106 - TargetToSourceTransparentContainer

// ProtocolIEField 107 - TimeToWait

// ProtocolIEField 108 - TraceActivation
func TraceActivationExtIEsExtensionValue (value ngapType.TraceActivationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TraceActivationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 109 - TraceCollectionEntityIPAddress

// ProtocolIEField 110 - UEAggregateMaximumBitRate
func UEAggregateMaximumBitRateExtIEsExtensionValue (value ngapType.UEAggregateMaximumBitRateExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 111 - UEAssociatedLogicalNGConnectionList
func UEAssociatedLogicalNGConnectionItemExtIEs (value ngapType.UEAssociatedLogicalNGConnectionItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 112 - UEContextRequest

// ProtocolIEField 113 - UENGAPIDs
func UENGAPIDPairExtIEs (value ngapType.UENGAPIDPairExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UENGAPIDPairExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 114 - UEPagingIdentity

// ProtocolIEField 115 - UEPresenceInAreaOfInterestList
func UEPresenceInAreaOfInterestItemExtIEs (value ngapType.UEPresenceInAreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UEPresenceInAreaOfInterestItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 116 - UERadioCapability

// ProtocolIEField 117 - UERadioCapabilityForPaging
func UERadioCapabilityForPagingExtIEsExtensionValue (value ngapType.UERadioCapabilityForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 118 - UESecurityCapabilities
func UESecurityCapabilitiesExtIEsExtensionValue (value ngapType.UESecurityCapabilitiesExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 119 - UnavailableGUAMIList
func UnavailableGUAMIItemExtIEs (value ngapType.UnavailableGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UnavailableGUAMIItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 120 - UserLocationInformation
func UserLocationInformationEUTRAExtIEs (value ngapType.UserLocationInformationEUTRAExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UserLocationInformationEUTRAExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func UserLocationInformationNRExtIEs (value ngapType.UserLocationInformationNRExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UserLocationInformationNRExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func UserLocationInformationN3IWFExtIEs (value ngapType.UserLocationInformationN3IWFExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UserLocationInformationN3IWFExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 121 - WarningAreaList

// ProtocolIEField 122 - WarningMessageContents

// ProtocolIEField 123 - WarningSecurityInfo

// ProtocolIEField 124 - WarningType

// ProtocolIEField 125 - AdditionalULNGUUPTNLInformation
func GTPTunnelExtIEs (value ngapType.GTPTunnelExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnelExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnelExtIEs: " + err.Error())
	}
	binData, bitEnd, err = GTPTunnelExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GTPTunnelExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 126 - DataForwardingNotPossible

// ProtocolIEField 127 - DLNGUUPTNLInformation

// ProtocolIEField 128 - NetworkInstance

// ProtocolIEField 129 - PDUSessionAggregateMaximumBitRate
func PDUSessionAggregateMaximumBitRateExtIEsExtensionValue (value ngapType.PDUSessionAggregateMaximumBitRateExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 130 - PDUSessionResourceFailedToModifyListModCfm
func PDUSessionResourceFailedToModifyItemModCfmExtIEs (value ngapType.PDUSessionResourceFailedToModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 131 - PDUSessionResourceFailedToSetupListCxtFail
func PDUSessionResourceFailedToSetupItemCxtFailExtIEs (value ngapType.PDUSessionResourceFailedToSetupItemCxtFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 132 - PDUSessionResourceListCxtRelReq
func PDUSessionResourceItemCxtRelReqExtIEs (value ngapType.PDUSessionResourceItemCxtRelReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionResourceItemCxtRelReqExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 133 - PDUSessionType

// ProtocolIEField 134 - QosFlowAddOrModifyRequestList
func QosCharacteristics (value ngapType.QosCharacteristics, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.QosCharacteristicsPresentNonDynamic5QI:
		if value.NonDynamic5QI == nil {
			return binData, bitEnd, errors.New("Cause: NonDynamic5QI: NIL")
		}
		binData, bitEnd, err = NonDynamic5QIDescriptor(*value.NonDynamic5QI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosCharacteristics: " + err.Error())
		}
	case ngapType.QosCharacteristicsPresentDynamic5QI:
		if value.Dynamic5QI == nil {
			return binData, bitEnd, errors.New("Cause: Dynamic5QI: NIL")
		}
		binData, bitEnd, err = Dynamic5QIDescriptor(*value.Dynamic5QI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosCharacteristics: " + err.Error())
		}
	case ngapType.QosCharacteristicsPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("Cause: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerQosCharacteristicsExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosCharacteristics: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("QosCharacteristics: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AllocationAndRetentionPriority (value ngapType.AllocationAndRetentionPriority, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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
	binData, bitEnd = EncodeUint64(uint64(option), 5, binData, bitEnd)
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

func ReflectiveQosAttribute (value ngapType.ReflectiveQosAttribute, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("ReflectiveQosAttribute: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func AdditionalQosFlowInformation (value ngapType.AdditionalQosFlowInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("AdditionalQosFlowInformation: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs (value ngapType.ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowLevelQosParametersExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyRequestItemExtIEs (value ngapType.QosFlowAddOrModifyRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = QosFlowAddOrModifyRequestItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 135 - QosFlowSetupRequestList
func QosFlowSetupRequestItemExtIEs (value ngapType.QosFlowSetupRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = QosFlowSetupRequestItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 136 - QosFlowToReleaseList
func QosFlowItemExtIEs (value ngapType.QosFlowItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = QosFlowItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 137 - SecurityIndication
func SecurityIndicationExtIEsExtensionValue (value ngapType.SecurityIndicationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SecurityIndicationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 138 - ULNGUUPTNLInformation

// ProtocolIEField 139 - ULNGUUPTNLModifyList
func ULNGUUPTNLModifyItemExtIEs (value ngapType.ULNGUUPTNLModifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ULNGUUPTNLModifyItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 140 - WarningAreaCoordinates

