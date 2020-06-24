// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"
	"vht5gc/lib/ngap/ngapType"
)

// ProtocolIEField 0 - AllowedNSSAI
func SNSSAI (value ngapType.SNSSAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	if value.SD != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

func ProtocolExtensionContainerAllowedNSSAIItemExtIEs (value ngapType.ProtocolExtensionContainerAllowedNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAllowedNSSAIItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AllowedNSSAIItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAllowedNSSAIItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 1 - AMFName

// ProtocolIEField 2 - AMFOverloadResponse

// ProtocolIEField 3 - AMFSetID

// ProtocolIEField 4 - AMFTNLAssociationFailedToSetupList
func CPTransportLayerInformation (value ngapType.CPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.CPTransportLayerInformationPresentEndpointIPAddress:
		if value.EndpointIPAddress == nil {
			return binData, bitEnd, errors.New("CPTransportLayerInformation: EndpointIPAddress: NIL")
		}
		binData, bitEnd, err = TransportLayerAddress(*value.EndpointIPAddress, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CPTransportLayerInformation: " + err.Error())
		}
	case ngapType.CPTransportLayerInformationPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("CPTransportLayerInformation: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerCPTransportLayerInformationExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CPTransportLayerInformation: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("CPTransportLayerInformation: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerTNLAssociationItemExtIEs (value ngapType.ProtocolExtensionContainerTNLAssociationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTNLAssociationItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TNLAssociationItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTNLAssociationItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 5 - AMFTNLAssociationSetupList
func ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs (value ngapType.ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationSetupItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 6 - AMFTNLAssociationToAddList
func TNLAssociationUsage (value ngapType.TNLAssociationUsage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("TNLAssociationUsage: Enumerated: valueExt,valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func TNLAddressWeightFactor (value ngapType.TNLAddressWeightFactor, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("TNLAddressWeightFactor: int64: valueLB:0,valueUB:255")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 8, binData, 8)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs (value ngapType.ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToAddItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 7 - AMFTNLAssociationToRemoveList
func ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs (value ngapType.ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToRemoveItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 8 - AMFTNLAssociationToUpdateList
func ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs (value ngapType.ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AMFTNLAssociationToUpdateItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 9 - AMFTrafficLoadReductionIndication

// ProtocolIEField 10 - AMFUENGAPID

// ProtocolIEField 11 - AssistanceDataForPaging
func RecommendedCellsForPaging (value ngapType.RecommendedCellsForPaging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs (value ngapType.ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AssistanceDataForRecommendedCellsExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PagingAttemptCount (value ngapType.PagingAttemptCount, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 16 {
		return binData, bitEnd, errors.New("PagingAttemptCount: int64: valueExt,valueLB:1,valueUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func IntendedNumberOfPagingAttempts (value ngapType.IntendedNumberOfPagingAttempts, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 16 {
		return binData, bitEnd, errors.New("IntendedNumberOfPagingAttempts: int64: valueExt,valueLB:1,valueUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func NextPagingAreaScope (value ngapType.NextPagingAreaScope, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("NextPagingAreaScope: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerPagingAttemptInformationExtIEs (value ngapType.ProtocolExtensionContainerPagingAttemptInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPagingAttemptInformationExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PagingAttemptInformationExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPagingAttemptInformationExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AssistanceDataForPagingExtIEs (value ngapType.AssistanceDataForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AssistanceDataForPagingExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AssistanceDataForPagingExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 12 - BroadcastCancelledAreaList
func CellIDCancelledEUTRAItem (value ngapType.CellIDCancelledEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAICancelledEUTRAItem (value ngapType.TAICancelledEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDCancelledEUTRAItem (value ngapType.EmergencyAreaIDCancelledEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func CellIDCancelledNRItem (value ngapType.CellIDCancelledNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAICancelledNRItem (value ngapType.TAICancelledNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDCancelledNRItem (value ngapType.EmergencyAreaIDCancelledNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 13 - BroadcastCompletedAreaList
func CellIDBroadcastEUTRAItem (value ngapType.CellIDBroadcastEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAIBroadcastEUTRAItem (value ngapType.TAIBroadcastEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDBroadcastEUTRAItem (value ngapType.EmergencyAreaIDBroadcastEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func CellIDBroadcastNRItem (value ngapType.CellIDBroadcastNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAIBroadcastNRItem (value ngapType.TAIBroadcastNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDBroadcastNRItem (value ngapType.EmergencyAreaIDBroadcastNRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

// ProtocolIEField 14 - CancelAllWarningMessages

// ProtocolIEField 15 - Cause

// ProtocolIEField 16 - CellIDListForRestart

// ProtocolIEField 17 - ConcurrentWarningMessageInd

// ProtocolIEField 18 - CoreNetworkAssistanceInformation
func ProtocolIESingleContainerUEIdentityIndexValueExtIEs (value ngapType.ProtocolIESingleContainerUEIdentityIndexValueExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func TAIListForInactiveItem (value ngapType.TAIListForInactiveItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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
	binData, bitEnd = EncodeUint64(uint64(option), 5, binData, bitEnd)
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

func ExpectedHOInterval (value ngapType.ExpectedHOInterval, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 6 {
		return binData, bitEnd, errors.New("ExpectedHOInterval: Enumerated: valueExt,valueLB:0,valueUB:6")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedUEMobility (value ngapType.ExpectedUEMobility, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("ExpectedUEMobility: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedUEMovingTrajectory (value ngapType.ExpectedUEMovingTrajectory, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectory: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ExpectedUEMovingTrajectoryItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ExpectedUEMovingTrajectory: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerExpectedUEBehaviourExtIEs (value ngapType.ProtocolExtensionContainerExpectedUEBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEBehaviourExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ExpectedUEBehaviourExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEBehaviourExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func CoreNetworkAssistanceInformationExtIEs (value ngapType.CoreNetworkAssistanceInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CoreNetworkAssistanceInformationExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CoreNetworkAssistanceInformationExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 19 - CriticalityDiagnostics
func CriticalityDiagnosticsIEItem (value ngapType.CriticalityDiagnosticsIEItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func CriticalityDiagnosticsExtIEs (value ngapType.CriticalityDiagnosticsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CriticalityDiagnosticsExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 20 - DataCodingScheme

// ProtocolIEField 21 - DefaultPagingDRX

// ProtocolIEField 22 - DirectForwardingPathAvailability

// ProtocolIEField 23 - EmergencyAreaIDListForRestart

// ProtocolIEField 24 - EmergencyFallbackIndicator
func EmergencyFallbackIndicatorExtIEs (value ngapType.EmergencyFallbackIndicatorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EmergencyFallbackIndicatorExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyFallbackIndicatorExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 25 - EUTRACGI
func EUTRACGIExtIEs (value ngapType.EUTRACGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EUTRACGIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EUTRACGIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 26 - FiveGSTMSI
func FiveGSTMSIExtIEs (value ngapType.FiveGSTMSIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = FiveGSTMSIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("FiveGSTMSIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 27 - GlobalRANNodeID
func GNBID (value ngapType.GNBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.GNBIDPresentGNBID:
		if value.GNBID == nil {
			return binData, bitEnd, errors.New("GNBID: GNBID: NIL")
		}
		if value.GNBID.BitLength < 22 || value.GNBID.BitLength > 32 || len(value.GNBID.Bytes) != int((value.GNBID.BitLength-1)/8+1) {
			return binData, bitEnd, errors.New("GNBID: GNBID: sizeLB:22,sizeUB:32")
		}
		binData, bitEnd = EncodeUint64(value.GNBID.BitLength, 5, binData, bitEnd)
		binData, bitEnd = EncodeBitString(*value.GNBID, binData, 8)
	case ngapType.GNBIDPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("GNBID: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerGNBIDExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("GNBID: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("GNBID: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGlobalGNBIDExtIEs (value ngapType.ProtocolExtensionContainerGlobalGNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalGNBIDExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = GlobalGNBIDExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalGNBIDExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NgENBID (value ngapType.NgENBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.NgENBIDPresentMacroNgENBID:
		if value.MacroNgENBID == nil {
			return binData, bitEnd, errors.New("NgENBID: MacroNgENBID: NIL")
		}
		if value.MacroNgENBID.BitLength != 20 || len(value.MacroNgENBID.Bytes) != 3 {
			return binData, bitEnd, errors.New("NgENBID: MacroNgENBID: sizeLB:20,sizeUB:20")
		}
		binData, bitEnd = EncodeBitString(*value.MacroNgENBID, binData, 8)
	case ngapType.NgENBIDPresentShortMacroNgENBID:
		if value.ShortMacroNgENBID == nil {
			return binData, bitEnd, errors.New("NgENBID: ShortMacroNgENBID: NIL")
		}
		if value.ShortMacroNgENBID.BitLength != 18 || len(value.ShortMacroNgENBID.Bytes) != 3 {
			return binData, bitEnd, errors.New("NgENBID: ShortMacroNgENBID: sizeLB:18,sizeUB:18")
		}
		binData, bitEnd = EncodeBitString(*value.ShortMacroNgENBID, binData, 8)
	case ngapType.NgENBIDPresentLongMacroNgENBID:
		if value.LongMacroNgENBID == nil {
			return binData, bitEnd, errors.New("NgENBID: LongMacroNgENBID: NIL")
		}
		if value.LongMacroNgENBID.BitLength != 21 || len(value.LongMacroNgENBID.Bytes) != 3 {
			return binData, bitEnd, errors.New("NgENBID: LongMacroNgENBID: sizeLB:21,sizeUB:21")
		}
		binData, bitEnd = EncodeBitString(*value.LongMacroNgENBID, binData, 8)
	case ngapType.NgENBIDPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("NgENBID: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerNgENBIDExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NgENBID: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NgENBID: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGlobalNgENBIDExtIEs (value ngapType.ProtocolExtensionContainerGlobalNgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalNgENBIDExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = GlobalNgENBIDExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalNgENBIDExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func N3IWFID (value ngapType.N3IWFID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
	switch value.Present {
	case ngapType.N3IWFIDPresentN3IWFID:
		if value.N3IWFID == nil {
			return binData, bitEnd, errors.New("N3IWFID: N3IWFID: NIL")
		}
		if value.N3IWFID.BitLength != 16 {
			return binData, bitEnd, errors.New("N3IWFID: N3IWFID: sizeLB:16,sizeUB:16")
		}
		binData, bitEnd = EncodeBitString(*value.N3IWFID, binData, bitEnd)
	case ngapType.N3IWFIDPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("N3IWFID: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerN3IWFIDExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("N3IWFID: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("N3IWFID: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerGlobalN3IWFIDExtIEs (value ngapType.ProtocolExtensionContainerGlobalN3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalN3IWFIDExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = GlobalN3IWFIDExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGlobalN3IWFIDExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 28 - GUAMI
func GUAMIExtIEs (value ngapType.GUAMIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = GUAMIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GUAMIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 29 - HandoverType

// ProtocolIEField 30 - IMSVoiceSupportIndicator

// ProtocolIEField 31 - IndexToRFSP

// ProtocolIEField 32 - InfoOnRecommendedCellsAndRANNodesForPaging
func RecommendedRANNodeList (value ngapType.RecommendedRANNodeList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("RecommendedRANNodeList: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedRANNodeItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedRANNodeList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs (value ngapType.ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedRANNodesForPagingExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func InfoOnRecommendedCellsAndRANNodesForPagingExtIEs (value ngapType.InfoOnRecommendedCellsAndRANNodesForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InfoOnRecommendedCellsAndRANNodesForPagingExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 33 - LocationReportingRequestType
func AreaOfInterestItem (value ngapType.AreaOfInterestItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func LocationReportingRequestTypeExtIEs (value ngapType.LocationReportingRequestTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: " + err.Error())
	}
	binData, bitEnd, err = LocationReportingRequestTypeExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingRequestTypeExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 34 - MaskedIMEISV

// ProtocolIEField 35 - MessageIdentifier

// ProtocolIEField 36 - MobilityRestrictionList
func RATRestrictionsItem (value ngapType.RATRestrictionsItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = PLMNIdentity(value.PLMNIdentity, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItem: " + err.Error())
	}
	binData, bitEnd = EncodeUint64(0, 1, binData, bitEnd)
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

func ForbiddenAreaInformationItem (value ngapType.ForbiddenAreaInformationItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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
	binData, bitEnd = EncodeUint64(uint64(option), 4, binData, bitEnd)
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

func MobilityRestrictionListExtIEs (value ngapType.MobilityRestrictionListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = MobilityRestrictionListExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("MobilityRestrictionListExtIEs: " + err.Error())
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
func NRCGIExtIEs (value ngapType.NRCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = NRCGIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRCGIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 46 - NRPPaPDU

// ProtocolIEField 47 - NumberOfBroadcastsRequested

// ProtocolIEField 48 - OldAMF

// ProtocolIEField 49 - OverloadStartNSSAIList
func SliceOverloadList (value ngapType.SliceOverloadList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 1024 {
		return binData, bitEnd, errors.New("SliceOverloadList: List: valueExt,sizeLB:1,sizeUB:1024")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SliceOverloadItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SliceOverloadList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs (value ngapType.ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = OverloadStartNSSAIItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 50 - PagingDRX

// ProtocolIEField 51 - PagingOrigin

// ProtocolIEField 52 - PagingPriority

// ProtocolIEField 53 - PDUSessionResourceAdmittedList
func PDUSessionID (value ngapType.PDUSessionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("PDUSessionID: int64: valueLB:0,valueUB:255")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 8, binData, 8)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceAdmittedItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 54 - PDUSessionResourceFailedToModifyListModRes
func ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 55 - PDUSessionResourceFailedToSetupListCxtRes
func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 56 - PDUSessionResourceFailedToSetupListHOAck
func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemHOAckExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 57 - PDUSessionResourceFailedToSetupListPSReq
func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemPSReqExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 58 - PDUSessionResourceFailedToSetupListSURes
func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemSUResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 59 - PDUSessionResourceHandoverList
func ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceHandoverItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 60 - PDUSessionResourceListCxtRelCpl
func ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemCxtRelCplExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 61 - PDUSessionResourceListHORqd
func ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemHORqdExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 62 - PDUSessionResourceModifyListModCfm
func ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModCfmExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 63 - PDUSessionResourceModifyListModInd
func ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModIndExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 64 - PDUSessionResourceModifyListModReq
func ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModReqExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 65 - PDUSessionResourceModifyListModRes
func ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceModifyItemModResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 66 - PDUSessionResourceNotifyList
func ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceNotifyItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 67 - PDUSessionResourceReleasedListNot
func ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemNotExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 68 - PDUSessionResourceReleasedListPSAck
func ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemPSAckExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 69 - PDUSessionResourceReleasedListPSFail
func ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemPSFailExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 70 - PDUSessionResourceReleasedListRelRes
func ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceReleasedItemRelResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 71 - PDUSessionResourceSetupListCxtReq
func ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemCxtReqExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 72 - PDUSessionResourceSetupListCxtRes
func ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemCxtResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 73 - PDUSessionResourceSetupListHOReq
func ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemHOReqExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 74 - PDUSessionResourceSetupListSUReq
func ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemSUReqExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 75 - PDUSessionResourceSetupListSURes
func ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSetupItemSUResExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 76 - PDUSessionResourceToBeSwitchedDLList
func ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 77 - PDUSessionResourceSwitchedList
func ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceSwitchedItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 78 - PDUSessionResourceToReleaseListHOCmd
func ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToReleaseItemHOCmdExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 79 - PDUSessionResourceToReleaseListRelCmd
func ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceToReleaseItemRelCmdExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 80 - PLMNSupportList
func ProtocolExtensionContainerPLMNSupportItemExtIEs (value ngapType.ProtocolExtensionContainerPLMNSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPLMNSupportItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PLMNSupportItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPLMNSupportItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 81 - PWSFailedCellIDList

// ProtocolIEField 82 - RANNodeName

// ProtocolIEField 83 - RANPagingPriority

// ProtocolIEField 84 - RANStatusTransferTransparentContainer
func DRBsSubjectToStatusTransferItem (value ngapType.DRBsSubjectToStatusTransferItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtension != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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
	if value.IEExtension != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs(*value.IEExtension, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RANStatusTransferTransparentContainerExtIEs (value ngapType.RANStatusTransferTransparentContainerExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEs: " + err.Error())
	}
	binData, bitEnd, err = RANStatusTransferTransparentContainerExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANStatusTransferTransparentContainerExtIEs: " + err.Error())
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
func SecurityContextExtIEs (value ngapType.SecurityContextExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContextExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContextExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SecurityContextExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityContextExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 94 - SecurityKey

// ProtocolIEField 95 - SerialNumber

// ProtocolIEField 96 - ServedGUAMIList
func ProtocolExtensionContainerServedGUAMIItemExtIEs (value ngapType.ProtocolExtensionContainerServedGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerServedGUAMIItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ServedGUAMIItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerServedGUAMIItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 97 - SliceSupportList
func ProtocolExtensionContainerSliceSupportItemExtIEs (value ngapType.ProtocolExtensionContainerSliceSupportItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceSupportItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SliceSupportItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceSupportItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 98 - SONConfigurationTransferDL
func ProtocolExtensionContainerTargetRANNodeIDExtIEs (value ngapType.ProtocolExtensionContainerTargetRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetRANNodeIDExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TargetRANNodeIDExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTargetRANNodeIDExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSourceRANNodeIDExtIEs (value ngapType.ProtocolExtensionContainerSourceRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceRANNodeIDExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SourceRANNodeIDExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSourceRANNodeIDExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SONInformationRequest (value ngapType.SONInformationRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("SONInformationRequest: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
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
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
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

func ProtocolIESingleContainerSONInformationExtIEs (value ngapType.ProtocolIESingleContainerSONInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func XnTLAs (value ngapType.XnTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("XnTLAs: List: sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TransportLayerAddress(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnExtTLAs (value ngapType.XnExtTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 2 {
		return binData, bitEnd, errors.New("XnExtTLAs: List: valueExt,sizeLB:1,sizeUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = XnExtTLAItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnExtTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs (value ngapType.ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = XnTNLConfigurationInfoExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SONConfigurationTransferExtIEs (value ngapType.SONConfigurationTransferExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SONConfigurationTransferExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONConfigurationTransferExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 99 - SONConfigurationTransferUL

// ProtocolIEField 100 - SourceAMFUENGAPID

// ProtocolIEField 101 - SourceToTargetTransparentContainer

// ProtocolIEField 102 - SupportedTAList
func BroadcastPLMNList (value ngapType.BroadcastPLMNList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 12 {
		return binData, bitEnd, errors.New("BroadcastPLMNList: List: valueExt,sizeLB:1,sizeUB:12")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = BroadcastPLMNItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("BroadcastPLMNList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerSupportedTAItemExtIEs (value ngapType.ProtocolExtensionContainerSupportedTAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSupportedTAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SupportedTAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSupportedTAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 103 - TAIListForPaging
func ProtocolExtensionContainerTAIListForPagingItemExtIEs (value ngapType.ProtocolExtensionContainerTAIListForPagingItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForPagingItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIListForPagingItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIListForPagingItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 104 - TAIListForRestart

// ProtocolIEField 105 - TargetID
func EPSTAI (value ngapType.EPSTAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func ProtocolExtensionContainerTargeteNBIDExtIEs (value ngapType.ProtocolExtensionContainerTargeteNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTargeteNBIDExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TargeteNBIDExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTargeteNBIDExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 106 - TargetToSourceTransparentContainer

// ProtocolIEField 107 - TimeToWait

// ProtocolIEField 108 - TraceActivation
func TraceActivationExtIEs (value ngapType.TraceActivationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TraceActivationExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceActivationExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 109 - TraceCollectionEntityIPAddress

// ProtocolIEField 110 - UEAggregateMaximumBitRate
func UEAggregateMaximumBitRateExtIEs (value ngapType.UEAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UEAggregateMaximumBitRateExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 111 - UEAssociatedLogicalNGConnectionList
func ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs (value ngapType.ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 112 - UEContextRequest

// ProtocolIEField 113 - UENGAPIDs
func ProtocolExtensionContainerUENGAPIDPairExtIEs (value ngapType.ProtocolExtensionContainerUENGAPIDPairExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUENGAPIDPairExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UENGAPIDPairExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUENGAPIDPairExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 114 - UEPagingIdentity

// ProtocolIEField 115 - UEPresenceInAreaOfInterestList
func UEPresence (value ngapType.UEPresence, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("UEPresence: Enumerated: valueExt,valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs (value ngapType.ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UEPresenceInAreaOfInterestItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 116 - UERadioCapability

// ProtocolIEField 117 - UERadioCapabilityForPaging
func UERadioCapabilityForPagingExtIEs (value ngapType.UERadioCapabilityForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UERadioCapabilityForPagingExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 118 - UESecurityCapabilities
func UESecurityCapabilitiesExtIEs (value ngapType.UESecurityCapabilitiesExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEs: " + err.Error())
	}
	binData, bitEnd, err = UESecurityCapabilitiesExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UESecurityCapabilitiesExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 119 - UnavailableGUAMIList
func TimerApproachForGUAMIRemoval (value ngapType.TimerApproachForGUAMIRemoval, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("TimerApproachForGUAMIRemoval: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerUnavailableGUAMIItemExtIEs (value ngapType.ProtocolExtensionContainerUnavailableGUAMIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUnavailableGUAMIItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UnavailableGUAMIItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUnavailableGUAMIItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 120 - UserLocationInformation
func TimeStamp (value ngapType.TimeStamp, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 4 {
		return binData, bitEnd, errors.New("TimeStamp: OctetString: sizeLB:4,sizeUB:4")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs (value ngapType.ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UserLocationInformationEUTRAExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerUserLocationInformationNRExtIEs (value ngapType.ProtocolExtensionContainerUserLocationInformationNRExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationNRExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UserLocationInformationNRExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationNRExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PortNumber (value ngapType.PortNumber, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("PortNumber: OctetString: sizeLB:2,sizeUB:2")
	}
	binData, bitEnd = EncodeOctetStringEx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs (value ngapType.ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = UserLocationInformationN3IWFExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 121 - WarningAreaList

// ProtocolIEField 122 - WarningMessageContents

// ProtocolIEField 123 - WarningSecurityInfo

// ProtocolIEField 124 - WarningType

// ProtocolIEField 125 - AdditionalULNGUUPTNLInformation
func GTPTEID (value ngapType.GTPTEID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 4 {
		return binData, bitEnd, errors.New("GTPTEID: OctetString: sizeLB:4,sizeUB:4")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerGTPTunnelExtIEs (value ngapType.ProtocolExtensionContainerGTPTunnelExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGTPTunnelExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = GTPTunnelExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGTPTunnelExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 126 - DataForwardingNotPossible

// ProtocolIEField 127 - DLNGUUPTNLInformation

// ProtocolIEField 128 - NetworkInstance

// ProtocolIEField 129 - PDUSessionAggregateMaximumBitRate
func PDUSessionAggregateMaximumBitRateExtIEs (value ngapType.PDUSessionAggregateMaximumBitRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PDUSessionAggregateMaximumBitRateExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionAggregateMaximumBitRateExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 130 - PDUSessionResourceFailedToModifyListModCfm
func ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToModifyItemModCfmExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 131 - PDUSessionResourceFailedToSetupListCxtFail
func ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceFailedToSetupItemCxtFailExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 132 - PDUSessionResourceListCxtRelReq
func ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs (value ngapType.ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PDUSessionResourceItemCxtRelReqExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 133 - PDUSessionType

// ProtocolIEField 134 - QosFlowAddOrModifyRequestList
func QosFlowIdentifier (value ngapType.QosFlowIdentifier, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 63 {
		return binData, bitEnd, errors.New("QosFlowIdentifier: int64: valueExt,valueLB:0,valueUB:63")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
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
	binData, bitEnd = EncodeUint64(uint64(option), 5, binData, bitEnd)
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

func ERABID (value ngapType.ERABID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 15 {
		return binData, bitEnd, errors.New("ERABID: int64: valueExt,valueLB:0,valueUB:15")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs (value ngapType.ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowAddOrModifyRequestItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 135 - QosFlowSetupRequestList
func ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs (value ngapType.ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowSetupRequestItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 136 - QosFlowToReleaseList
func ProtocolExtensionContainerQosFlowItemExtIEs (value ngapType.ProtocolExtensionContainerQosFlowItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = QosFlowItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerQosFlowItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 137 - SecurityIndication
func SecurityIndicationExtIEs (value ngapType.SecurityIndicationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SecurityIndicationExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecurityIndicationExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 138 - ULNGUUPTNLInformation

// ProtocolIEField 139 - ULNGUUPTNLModifyList
func ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs (value ngapType.ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ULNGUUPTNLModifyItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

// ProtocolIEField 140 - WarningAreaCoordinates

