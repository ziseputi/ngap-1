// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"
	"vht5gc/lib/ngap/ngapType"
)

// ProtocolIEField 0 - AllowedNSSAI
func SNSSAIExtIEs (value ngapType.SNSSAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SNSSAIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SNSSAIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolExtensionID (value ngapType.ProtocolExtensionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionID: int64: valueLB:0,valueUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func Criticality (value ngapType.Criticality, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("Criticality: Enumerated: valueLB:0,valueUB:2")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func AllowedNSSAIItemExtIEsExtensionValue (value ngapType.AllowedNSSAIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AllowedNSSAIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 1 - AMFName

// ProtocolIEField 2 - AMFOverloadResponse

// ProtocolIEField 3 - AMFSetID

// ProtocolIEField 4 - AMFTNLAssociationFailedToSetupList
func TNLAssociationItemExtIEsExtensionValue (value ngapType.TNLAssociationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TNLAssociationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 5 - AMFTNLAssociationSetupList
func AMFTNLAssociationSetupItemExtIEsExtensionValue (value ngapType.AMFTNLAssociationSetupItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationSetupItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 6 - AMFTNLAssociationToAddList
func AMFTNLAssociationToAddItemExtIEsExtensionValue (value ngapType.AMFTNLAssociationToAddItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToAddItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 7 - AMFTNLAssociationToRemoveList
func AMFTNLAssociationToRemoveItemExtIEsExtensionValue (value ngapType.AMFTNLAssociationToRemoveItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToRemoveItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 8 - AMFTNLAssociationToUpdateList
func AMFTNLAssociationToUpdateItemExtIEsExtensionValue (value ngapType.AMFTNLAssociationToUpdateItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFTNLAssociationToUpdateItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 9 - AMFTrafficLoadReductionIndication

// ProtocolIEField 10 - AMFUENGAPID

// ProtocolIEField 11 - AssistanceDataForPaging
func RecommendedCellItem (value ngapType.RecommendedCellItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
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
		return binData, bitEnd, errors.New("RecommendedCellItem: " + err.Error())
	}
	if value.TimeStayedInCell != nil {
		binData, bitEnd = EncodeUint64(uint64(*value.TimeStayedInCell), 0, binData, bitEnd)
	}
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerRecommendedCellItemExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RecommendedCellItem: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedCellsForPagingExtIEs (value ngapType.RecommendedCellsForPagingExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEs: " + err.Error())
	}
	binData, bitEnd, err = RecommendedCellsForPagingExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AssistanceDataForRecommendedCellsExtIEsExtensionValue (value ngapType.AssistanceDataForRecommendedCellsExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AssistanceDataForRecommendedCellsExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PagingAttemptInformationExtIEsExtensionValue (value ngapType.PagingAttemptInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PagingAttemptInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}


// ProtocolIEField 12 - BroadcastCancelledAreaList
func CellIDCancelledEUTRAItemExtIEs (value ngapType.CellIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CellIDCancelledEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func PLMNIdentity (value ngapType.PLMNIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("PLMNIdentity: OctetString: sizeLB:3,sizeUB:3")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func TAC (value ngapType.TAC, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("TAC: OctetString: sizeLB:3,sizeUB:3")
	}
	binData, bitEnd = EncodeOctetString(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerTAIExtIEs (value ngapType.ProtocolExtensionContainerTAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TAIExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerTAIExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
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
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAICancelledEUTRAItemExtIEs (value ngapType.TAICancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAICancelledEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAIEUTRAItem (value ngapType.CancelledCellsInEAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDCancelledEUTRAItemExtIEs (value ngapType.EmergencyAreaIDCancelledEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CellIDCancelledNRItemExtIEs (value ngapType.CellIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CellIDCancelledNRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINRItem (value ngapType.CancelledCellsInTAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAICancelledNRItemExtIEs (value ngapType.TAICancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAICancelledNRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINRItem (value ngapType.CancelledCellsInEAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDCancelledNRItemExtIEs (value ngapType.EmergencyAreaIDCancelledNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EmergencyAreaIDCancelledNRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 13 - BroadcastCompletedAreaList
func CellIDBroadcastEUTRAItemExtIEs (value ngapType.CellIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CellIDBroadcastEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRAItem (value ngapType.CompletedCellsInTAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAIBroadcastEUTRAItemExtIEs (value ngapType.TAIBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAIBroadcastEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRAItem (value ngapType.CompletedCellsInEAIEUTRAItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDBroadcastEUTRAItemExtIEs (value ngapType.EmergencyAreaIDBroadcastEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CellIDBroadcastNRItemExtIEs (value ngapType.CellIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CellIDBroadcastNRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINRItem (value ngapType.CompletedCellsInTAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func TAIBroadcastNRItemExtIEs (value ngapType.TAIBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAIBroadcastNRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINRItem (value ngapType.CompletedCellsInEAINRItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func EmergencyAreaIDBroadcastNRItemExtIEs (value ngapType.EmergencyAreaIDBroadcastNRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

// ProtocolIEField 14 - CancelAllWarningMessages

// ProtocolIEField 15 - Cause

// ProtocolIEField 16 - CellIDListForRestart

// ProtocolIEField 17 - ConcurrentWarningMessageInd

// ProtocolIEField 18 - CoreNetworkAssistanceInformation
func TAIListForInactiveItemExtIEs (value ngapType.TAIListForInactiveItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAIListForInactiveItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ExpectedUEActivityBehaviourExtIEs (value ngapType.ExpectedUEActivityBehaviourExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ExpectedUEActivityBehaviourExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs (value ngapType.ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = ExpectedUEMovingTrajectoryItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ExpectedUEBehaviourExtIEsExtensionValue (value ngapType.ExpectedUEBehaviourExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExpectedUEBehaviourExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 19 - CriticalityDiagnostics
func CriticalityDiagnosticsIEItemExtIEs (value ngapType.CriticalityDiagnosticsIEItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CriticalityDiagnosticsIEItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}


// ProtocolIEField 20 - DataCodingScheme

// ProtocolIEField 21 - DefaultPagingDRX

// ProtocolIEField 22 - DirectForwardingPathAvailability

// ProtocolIEField 23 - EmergencyAreaIDListForRestart

// ProtocolIEField 24 - EmergencyFallbackIndicator

// ProtocolIEField 25 - EUTRACGI

// ProtocolIEField 26 - FiveGSTMSI

// ProtocolIEField 27 - GlobalRANNodeID
func GlobalGNBIDExtIEsExtensionValue (value ngapType.GlobalGNBIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalGNBIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GlobalNgENBIDExtIEsExtensionValue (value ngapType.GlobalNgENBIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalNgENBIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GlobalN3IWFIDExtIEsExtensionValue (value ngapType.GlobalN3IWFIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalN3IWFIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 28 - GUAMI

// ProtocolIEField 29 - HandoverType

// ProtocolIEField 30 - IMSVoiceSupportIndicator

// ProtocolIEField 31 - IndexToRFSP

// ProtocolIEField 32 - InfoOnRecommendedCellsAndRANNodesForPaging
func AMFPagingTarget (value ngapType.AMFPagingTarget, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	case ngapType.AMFPagingTargetPresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("Cause: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: " + err.Error())
		}
	case ngapType.AMFPagingTargetPresentTAI:
		if value.TAI == nil {
			return binData, bitEnd, errors.New("Cause: TAI: NIL")
		}
		binData, bitEnd, err = TAI(*value.TAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: " + err.Error())
		}
	case ngapType.AMFPagingTargetPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("Cause: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerAMFPagingTargetExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFPagingTarget: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerRecommendedRANNodeItemExtIEs (value ngapType.ProtocolExtensionContainerRecommendedRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodeItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedRANNodeItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedRANNodeItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedRANNodesForPagingExtIEsExtensionValue (value ngapType.RecommendedRANNodesForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedRANNodesForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 33 - LocationReportingRequestType
func AreaOfInterestTAIList (value ngapType.AreaOfInterestTAIList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("AreaOfInterestTAIList: List: valueExt,sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 4, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestTAIItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestTAIList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestCellList (value ngapType.AreaOfInterestCellList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 256 {
		return binData, bitEnd, errors.New("AreaOfInterestCellList: List: valueExt,sizeLB:1,sizeUB:256")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 8, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestCellItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestCellList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestRANNodeList (value ngapType.AreaOfInterestRANNodeList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 64 {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeList: List: valueExt,sizeLB:1,sizeUB:64")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 6, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestRANNodeItem(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AreaOfInterestRANNodeList: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestExtIEs (value ngapType.ProtocolExtensionContainerAreaOfInterestExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestItemExtIEs (value ngapType.AreaOfInterestItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AreaOfInterestItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}


// ProtocolIEField 34 - MaskedIMEISV

// ProtocolIEField 35 - MessageIdentifier

// ProtocolIEField 36 - MobilityRestrictionList
func RATRestrictionsItemExtIEs (value ngapType.RATRestrictionsItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = RATRestrictionsItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ForbiddenAreaInformationItemExtIEs (value ngapType.ForbiddenAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ForbiddenAreaInformationItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ServiceAreaInformationItemExtIEs (value ngapType.ServiceAreaInformationItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ServiceAreaInformationItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEs: " + err.Error())
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

// ProtocolIEField 46 - NRPPaPDU

// ProtocolIEField 47 - NumberOfBroadcastsRequested

// ProtocolIEField 48 - OldAMF

// ProtocolIEField 49 - OverloadStartNSSAIList
func ProtocolExtensionContainerSliceOverloadItemExtIEs (value ngapType.ProtocolExtensionContainerSliceOverloadItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceOverloadItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = SliceOverloadItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerSliceOverloadItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func OverloadStartNSSAIItemExtIEsExtensionValue (value ngapType.OverloadStartNSSAIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("OverloadStartNSSAIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 50 - PagingDRX

// ProtocolIEField 51 - PagingOrigin

// ProtocolIEField 52 - PagingPriority

// ProtocolIEField 53 - PDUSessionResourceAdmittedList
func PDUSessionResourceAdmittedItemExtIEsExtensionValue (value ngapType.PDUSessionResourceAdmittedItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceAdmittedItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 54 - PDUSessionResourceFailedToModifyListModRes
func PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 55 - PDUSessionResourceFailedToSetupListCxtRes
func PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 56 - PDUSessionResourceFailedToSetupListHOAck
func PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 57 - PDUSessionResourceFailedToSetupListPSReq
func PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 58 - PDUSessionResourceFailedToSetupListSURes
func PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 59 - PDUSessionResourceHandoverList
func PDUSessionResourceHandoverItemExtIEsExtensionValue (value ngapType.PDUSessionResourceHandoverItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceHandoverItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 60 - PDUSessionResourceListCxtRelCpl
func PDUSessionResourceItemCxtRelCplExtIEsExtensionValue (value ngapType.PDUSessionResourceItemCxtRelCplExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelCplExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 61 - PDUSessionResourceListHORqd
func PDUSessionResourceItemHORqdExtIEsExtensionValue (value ngapType.PDUSessionResourceItemHORqdExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemHORqdExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 62 - PDUSessionResourceModifyListModCfm
func PDUSessionResourceModifyItemModCfmExtIEsExtensionValue (value ngapType.PDUSessionResourceModifyItemModCfmExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModCfmExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 63 - PDUSessionResourceModifyListModInd
func PDUSessionResourceModifyItemModIndExtIEsExtensionValue (value ngapType.PDUSessionResourceModifyItemModIndExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModIndExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 64 - PDUSessionResourceModifyListModReq
func PDUSessionResourceModifyItemModReqExtIEsExtensionValue (value ngapType.PDUSessionResourceModifyItemModReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 65 - PDUSessionResourceModifyListModRes
func PDUSessionResourceModifyItemModResExtIEsExtensionValue (value ngapType.PDUSessionResourceModifyItemModResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyItemModResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 66 - PDUSessionResourceNotifyList
func PDUSessionResourceNotifyItemExtIEsExtensionValue (value ngapType.PDUSessionResourceNotifyItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 67 - PDUSessionResourceReleasedListNot
func PDUSessionResourceReleasedItemNotExtIEsExtensionValue (value ngapType.PDUSessionResourceReleasedItemNotExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemNotExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 68 - PDUSessionResourceReleasedListPSAck
func PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue (value ngapType.PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 69 - PDUSessionResourceReleasedListPSFail
func PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue (value ngapType.PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 70 - PDUSessionResourceReleasedListRelRes
func PDUSessionResourceReleasedItemRelResExtIEsExtensionValue (value ngapType.PDUSessionResourceReleasedItemRelResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleasedItemRelResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 71 - PDUSessionResourceSetupListCxtReq
func PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue (value ngapType.PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 72 - PDUSessionResourceSetupListCxtRes
func PDUSessionResourceSetupItemCxtResExtIEsExtensionValue (value ngapType.PDUSessionResourceSetupItemCxtResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemCxtResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 73 - PDUSessionResourceSetupListHOReq
func PDUSessionResourceSetupItemHOReqExtIEsExtensionValue (value ngapType.PDUSessionResourceSetupItemHOReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemHOReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 74 - PDUSessionResourceSetupListSUReq
func PDUSessionResourceSetupItemSUReqExtIEsExtensionValue (value ngapType.PDUSessionResourceSetupItemSUReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 75 - PDUSessionResourceSetupListSURes
func PDUSessionResourceSetupItemSUResExtIEsExtensionValue (value ngapType.PDUSessionResourceSetupItemSUResExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupItemSUResExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 76 - PDUSessionResourceToBeSwitchedDLList
func PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue (value ngapType.PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 77 - PDUSessionResourceSwitchedList
func PDUSessionResourceSwitchedItemExtIEsExtensionValue (value ngapType.PDUSessionResourceSwitchedItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSwitchedItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 78 - PDUSessionResourceToReleaseListHOCmd
func PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue (value ngapType.PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 79 - PDUSessionResourceToReleaseListRelCmd
func PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue (value ngapType.PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 80 - PLMNSupportList
func PLMNSupportItemExtIEsExtensionValue (value ngapType.PLMNSupportItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PLMNSupportItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 81 - PWSFailedCellIDList

// ProtocolIEField 82 - RANNodeName

// ProtocolIEField 83 - RANPagingPriority

// ProtocolIEField 84 - RANStatusTransferTransparentContainer
func DRBStatusUL12 (value ngapType.DRBStatusUL12, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtension != nil {
		option += 1<<0
	}
	if value.ReceiveStatusOfULPDCPSDUs != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN12(value.ULCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12: " + err.Error())
	}
	if value.ReceiveStatusOfULPDCPSDUs != nil {
		binData, bitEnd = EncodeBitString(*value.ReceiveStatusOfULPDCPSDUs, binData, bitEnd)
	}
	if value.IEExtension != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusUL12ExtIEs(*value.IEExtension, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL12: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusUL18 (value ngapType.DRBStatusUL18, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtension != nil {
		option += 1<<0
	}
	if value.ReceiveStatusOfULPDCPSDUs != nil {
		option += 1<<1
	}
	binData, bitEnd = EncodeUint64(uint64(option), 3, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN18(value.ULCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18: " + err.Error())
	}
	if value.ReceiveStatusOfULPDCPSDUs != nil {
		binData, bitEnd = EncodeBitString(*value.ReceiveStatusOfULPDCPSDUs, binData, bitEnd)
	}
	if value.IEExtension != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusUL18ExtIEs(*value.IEExtension, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL18: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerDRBStatusULExtIEs (value ngapType.ProtocolIESingleContainerDRBStatusULExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func DRBStatusDL12 (value ngapType.DRBStatusDL12, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtension != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN12(value.DLCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12: " + err.Error())
	}
	if value.IEExtension != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusDL12ExtIEs(*value.IEExtension, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL12: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusDL18 (value ngapType.DRBStatusDL18, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtension != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd, err = COUNTValueForPDCPSN18(value.DLCOUNTValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18: " + err.Error())
	}
	if value.IEExtension != nil {
		binData, bitEnd, err = ProtocolExtensionContainerDRBStatusDL18ExtIEs(*value.IEExtension, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL18: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerDRBStatusDLExtIEs (value ngapType.ProtocolIESingleContainerDRBStatusDLExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func DRBsSubjectToStatusTransferItemExtIEs (value ngapType.DRBsSubjectToStatusTransferItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = DRBsSubjectToStatusTransferItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEs: " + err.Error())
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

// ProtocolIEField 94 - SecurityKey

// ProtocolIEField 95 - SerialNumber

// ProtocolIEField 96 - ServedGUAMIList
func ServedGUAMIItemExtIEsExtensionValue (value ngapType.ServedGUAMIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ServedGUAMIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 97 - SliceSupportList
func SliceSupportItemExtIEsExtensionValue (value ngapType.SliceSupportItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SliceSupportItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 98 - SONConfigurationTransferDL
func TargetRANNodeIDExtIEsExtensionValue (value ngapType.TargetRANNodeIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TargetRANNodeIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SourceRANNodeIDExtIEsExtensionValue (value ngapType.SourceRANNodeIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SourceRANNodeIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SONInformationReplyExtIEs (value ngapType.SONInformationReplyExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationReplyExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationReplyExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SONInformationReplyExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationReplyExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func XnGTPTLAs (value ngapType.XnGTPTLAs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 16 {
		return binData, bitEnd, errors.New("XnGTPTLAs: List: sizeLB:1,sizeUB:16")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = TransportLayerAddress(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("XnGTPTLAs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerXnExtTLAItemExtIEs (value ngapType.ProtocolExtensionContainerXnExtTLAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerXnExtTLAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = XnExtTLAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerXnExtTLAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func XnTNLConfigurationInfoExtIEsExtensionValue (value ngapType.XnTNLConfigurationInfoExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("XnTNLConfigurationInfoExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 99 - SONConfigurationTransferUL

// ProtocolIEField 100 - SourceAMFUENGAPID

// ProtocolIEField 101 - SourceToTargetTransparentContainer

// ProtocolIEField 102 - SupportedTAList
func ProtocolExtensionContainerBroadcastPLMNItemExtIEs (value ngapType.ProtocolExtensionContainerBroadcastPLMNItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerBroadcastPLMNItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = BroadcastPLMNItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerBroadcastPLMNItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func SupportedTAItemExtIEsExtensionValue (value ngapType.SupportedTAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SupportedTAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 103 - TAIListForPaging
func TAIListForPagingItemExtIEsExtensionValue (value ngapType.TAIListForPagingItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIListForPagingItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 104 - TAIListForRestart

// ProtocolIEField 105 - TargetID
func EPSTAIExtIEs (value ngapType.EPSTAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = EPSTAIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("EPSTAIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func TargeteNBIDExtIEsExtensionValue (value ngapType.TargeteNBIDExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TargeteNBIDExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 106 - TargetToSourceTransparentContainer

// ProtocolIEField 107 - TimeToWait

// ProtocolIEField 108 - TraceActivation

// ProtocolIEField 109 - TraceCollectionEntityIPAddress

// ProtocolIEField 110 - UEAggregateMaximumBitRate

// ProtocolIEField 111 - UEAssociatedLogicalNGConnectionList
func UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue (value ngapType.UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 112 - UEContextRequest

// ProtocolIEField 113 - UENGAPIDs
func UENGAPIDPairExtIEsExtensionValue (value ngapType.UENGAPIDPairExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UENGAPIDPairExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 114 - UEPagingIdentity

// ProtocolIEField 115 - UEPresenceInAreaOfInterestList
func UEPresenceInAreaOfInterestItemExtIEsExtensionValue (value ngapType.UEPresenceInAreaOfInterestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEPresenceInAreaOfInterestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 116 - UERadioCapability

// ProtocolIEField 117 - UERadioCapabilityForPaging

// ProtocolIEField 118 - UESecurityCapabilities

// ProtocolIEField 119 - UnavailableGUAMIList
func UnavailableGUAMIItemExtIEsExtensionValue (value ngapType.UnavailableGUAMIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UnavailableGUAMIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 120 - UserLocationInformation
func UserLocationInformationEUTRAExtIEsExtensionValue (value ngapType.UserLocationInformationEUTRAExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UserLocationInformationEUTRAExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserLocationInformationNRExtIEsExtensionValue (value ngapType.UserLocationInformationNRExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UserLocationInformationNRExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserLocationInformationN3IWFExtIEsExtensionValue (value ngapType.UserLocationInformationN3IWFExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UserLocationInformationN3IWFExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 121 - WarningAreaList

// ProtocolIEField 122 - WarningMessageContents

// ProtocolIEField 123 - WarningSecurityInfo

// ProtocolIEField 124 - WarningType

// ProtocolIEField 125 - AdditionalULNGUUPTNLInformation
func GTPTunnelExtIEsExtensionValue (value ngapType.GTPTunnelExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GTPTunnelExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 126 - DataForwardingNotPossible

// ProtocolIEField 127 - DLNGUUPTNLInformation

// ProtocolIEField 128 - NetworkInstance

// ProtocolIEField 129 - PDUSessionAggregateMaximumBitRate

// ProtocolIEField 130 - PDUSessionResourceFailedToModifyListModCfm
func PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 131 - PDUSessionResourceFailedToSetupListCxtFail
func PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue (value ngapType.PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 132 - PDUSessionResourceListCxtRelReq
func PDUSessionResourceItemCxtRelReqExtIEsExtensionValue (value ngapType.PDUSessionResourceItemCxtRelReqExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceItemCxtRelReqExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 133 - PDUSessionType

// ProtocolIEField 134 - QosFlowAddOrModifyRequestList
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
	binData, bitEnd = EncodeUint64(uint64(option), 5, binData, bitEnd)
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
	binData, bitEnd = EncodeUint64(uint64(option), 6, binData, bitEnd)
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

func ProtocolIESingleContainerQosCharacteristicsExtIEs (value ngapType.ProtocolIESingleContainerQosCharacteristicsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func PriorityLevelARP (value ngapType.PriorityLevelARP, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 15 {
		return binData, bitEnd, errors.New("PriorityLevelARP: int64: valueLB:1,valueUB:15")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PreEmptionCapability (value ngapType.PreEmptionCapability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("PreEmptionCapability: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PreEmptionVulnerability (value ngapType.PreEmptionVulnerability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("PreEmptionVulnerability: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs (value ngapType.ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AllocationAndRetentionPriorityExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func NotificationControl (value ngapType.NotificationControl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value != 0 {
		return binData, bitEnd, errors.New("NotificationControl: Enumerated: valueExt,valueLB:0,valueUB:0")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PacketLossRate (value ngapType.PacketLossRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1000 {
		return binData, bitEnd, errors.New("PacketLossRate: int64: valueExt,valueLB:0,valueUB:1000")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerGBRQosInformationExtIEs (value ngapType.ProtocolExtensionContainerGBRQosInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerGBRQosInformationExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = GBRQosInformationExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerGBRQosInformationExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func QosFlowLevelQosParametersExtIEs (value ngapType.QosFlowLevelQosParametersExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	binData, bitEnd, err = QosFlowLevelQosParametersExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func QosFlowAddOrModifyRequestItemExtIEsExtensionValue (value ngapType.QosFlowAddOrModifyRequestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowAddOrModifyRequestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 135 - QosFlowSetupRequestList
func QosFlowSetupRequestItemExtIEsExtensionValue (value ngapType.QosFlowSetupRequestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowSetupRequestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 136 - QosFlowToReleaseList
func QosFlowItemExtIEsExtensionValue (value ngapType.QosFlowItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 137 - SecurityIndication

// ProtocolIEField 138 - ULNGUUPTNLInformation

// ProtocolIEField 139 - ULNGUUPTNLModifyList
func ULNGUUPTNLModifyItemExtIEsExtensionValue (value ngapType.ULNGUUPTNLModifyItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ULNGUUPTNLModifyItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 140 - WarningAreaCoordinates

