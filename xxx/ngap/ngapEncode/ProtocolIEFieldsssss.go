// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"
	"vht5gc/lib/ngap/ngapType"
)

// ProtocolIEField 0 - AllowedNSSAI
func SNSSAIExtIEsExtensionValue (value ngapType.SNSSAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SNSSAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 1 - AMFName

// ProtocolIEField 2 - AMFOverloadResponse

// ProtocolIEField 3 - AMFSetID

// ProtocolIEField 4 - AMFTNLAssociationFailedToSetupList

// ProtocolIEField 5 - AMFTNLAssociationSetupList

// ProtocolIEField 6 - AMFTNLAssociationToAddList

// ProtocolIEField 7 - AMFTNLAssociationToRemoveList

// ProtocolIEField 8 - AMFTNLAssociationToUpdateList

// ProtocolIEField 9 - AMFTrafficLoadReductionIndication

// ProtocolIEField 10 - AMFUENGAPID

// ProtocolIEField 11 - AssistanceDataForPaging
func ProtocolExtensionContainerRecommendedCellItemExtIEs (value ngapType.ProtocolExtensionContainerRecommendedCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = RecommendedCellItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerRecommendedCellItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func RecommendedCellsForPagingExtIEsExtensionValue (value ngapType.RecommendedCellsForPagingExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedCellsForPagingExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RecommendedCellItemExtIEs (value ngapType.RecommendedCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = RecommendedCellItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func RecommendedCellItemExtIEsExtensionValue (value ngapType.RecommendedCellItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedCellItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 12 - BroadcastCancelledAreaList
func CellIDCancelledEUTRAItemExtIEsExtensionValue (value ngapType.CellIDCancelledEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDCancelledEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIExtIEs (value ngapType.TAIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = TAIExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TAIExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInTAIEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledEUTRAItemExtIEsExtensionValue (value ngapType.TAICancelledEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAICancelledEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInEAIEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue (value ngapType.EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDCancelledNRItemExtIEsExtensionValue (value ngapType.CellIDCancelledNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDCancelledNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs (value ngapType.ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInTAINRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAICancelledNRItemExtIEsExtensionValue (value ngapType.TAICancelledNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAICancelledNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs (value ngapType.ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CancelledCellsInEAINRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDCancelledNRItemExtIEsExtensionValue (value ngapType.EmergencyAreaIDCancelledNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDCancelledNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TAIExtIEsExtensionValue (value ngapType.TAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAIEUTRAItemExtIEs (value ngapType.CancelledCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInTAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAIEUTRAItemExtIEs (value ngapType.CancelledCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInEAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINRItemExtIEs (value ngapType.CancelledCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInTAINRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINRItemExtIEs (value ngapType.CancelledCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CancelledCellsInEAINRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAIEUTRAItemExtIEsExtensionValue (value ngapType.CancelledCellsInTAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInTAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAIEUTRAItemExtIEsExtensionValue (value ngapType.CancelledCellsInEAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInEAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInTAINRItemExtIEsExtensionValue (value ngapType.CancelledCellsInTAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInTAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CancelledCellsInEAINRItemExtIEsExtensionValue (value ngapType.CancelledCellsInEAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CancelledCellsInEAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 13 - BroadcastCompletedAreaList
func CellIDBroadcastEUTRAItemExtIEsExtensionValue (value ngapType.CellIDBroadcastEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDBroadcastEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInTAIEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastEUTRAItemExtIEsExtensionValue (value ngapType.TAIBroadcastEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIBroadcastEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs (value ngapType.ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInEAIEUTRAItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue (value ngapType.EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDBroadcastNRItemExtIEsExtensionValue (value ngapType.CellIDBroadcastNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDBroadcastNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs (value ngapType.ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInTAINRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func TAIBroadcastNRItemExtIEsExtensionValue (value ngapType.TAIBroadcastNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIBroadcastNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs (value ngapType.ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = CompletedCellsInEAINRItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue (value ngapType.EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRAItemExtIEs (value ngapType.CompletedCellsInTAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInTAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRAItemExtIEs (value ngapType.CompletedCellsInEAIEUTRAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInEAIEUTRAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINRItemExtIEs (value ngapType.CompletedCellsInTAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInTAINRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINRItemExtIEs (value ngapType.CompletedCellsInEAINRItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = CompletedCellsInEAINRItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAIEUTRAItemExtIEsExtensionValue (value ngapType.CompletedCellsInTAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInTAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAIEUTRAItemExtIEsExtensionValue (value ngapType.CompletedCellsInEAIEUTRAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInEAIEUTRAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInTAINRItemExtIEsExtensionValue (value ngapType.CompletedCellsInTAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInTAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CompletedCellsInEAINRItemExtIEsExtensionValue (value ngapType.CompletedCellsInEAINRItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CompletedCellsInEAINRItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 14 - CancelAllWarningMessages

// ProtocolIEField 15 - Cause

// ProtocolIEField 16 - CellIDListForRestart

// ProtocolIEField 17 - ConcurrentWarningMessageInd

// ProtocolIEField 18 - CoreNetworkAssistanceInformation
func TAIListForInactiveItemExtIEsExtensionValue (value ngapType.TAIListForInactiveItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("TAIListForInactiveItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ExpectedUEActivityBehaviourExtIEsExtensionValue (value ngapType.ExpectedUEActivityBehaviourExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExpectedUEActivityBehaviourExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ExpectedUEMovingTrajectoryItemExtIEs (value ngapType.ExpectedUEMovingTrajectoryItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = ExpectedUEMovingTrajectoryItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ExpectedUEMovingTrajectoryItemExtIEsExtensionValue (value ngapType.ExpectedUEMovingTrajectoryItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ExpectedUEMovingTrajectoryItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 19 - CriticalityDiagnostics
func CriticalityDiagnosticsIEItemExtIEsExtensionValue (value ngapType.CriticalityDiagnosticsIEItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CriticalityDiagnosticsIEItemExtIEsExtensionValue: Present: INVALID")
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

// ProtocolIEField 28 - GUAMI

// ProtocolIEField 29 - HandoverType

// ProtocolIEField 30 - IMSVoiceSupportIndicator

// ProtocolIEField 31 - IndexToRFSP

// ProtocolIEField 32 - InfoOnRecommendedCellsAndRANNodesForPaging
func ProtocolIESingleContainerAMFPagingTargetExtIEs (value ngapType.ProtocolIESingleContainerAMFPagingTargetExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	return binData, bitEnd, err
}

func RecommendedRANNodeItemExtIEs (value ngapType.RecommendedRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = RecommendedRANNodeItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func RecommendedRANNodeItemExtIEsExtensionValue (value ngapType.RecommendedRANNodeItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RecommendedRANNodeItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 33 - LocationReportingRequestType
func AreaOfInterestTAIItem (value ngapType.AreaOfInterestTAIItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func AreaOfInterestCellItem (value ngapType.AreaOfInterestCellItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func AreaOfInterestRANNodeItem (value ngapType.AreaOfInterestRANNodeItem, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
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

func AreaOfInterestExtIEs (value ngapType.AreaOfInterestExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AreaOfInterestExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AreaOfInterestItemExtIEsExtensionValue (value ngapType.AreaOfInterestItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs (value ngapType.ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestTAIItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestCellItemExtIEs (value ngapType.ProtocolExtensionContainerAreaOfInterestCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestCellItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestCellItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestCellItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs (value ngapType.ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = AreaOfInterestRANNodeItemExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AreaOfInterestExtIEsExtensionValue (value ngapType.AreaOfInterestExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestTAIItemExtIEs (value ngapType.AreaOfInterestTAIItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AreaOfInterestTAIItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AreaOfInterestCellItemExtIEs (value ngapType.AreaOfInterestCellItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AreaOfInterestCellItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AreaOfInterestRANNodeItemExtIEs (value ngapType.AreaOfInterestRANNodeItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AreaOfInterestRANNodeItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AreaOfInterestTAIItemExtIEsExtensionValue (value ngapType.AreaOfInterestTAIItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestTAIItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestCellItemExtIEsExtensionValue (value ngapType.AreaOfInterestCellItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestCellItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AreaOfInterestRANNodeItemExtIEsExtensionValue (value ngapType.AreaOfInterestRANNodeItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AreaOfInterestRANNodeItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 34 - MaskedIMEISV

// ProtocolIEField 35 - MessageIdentifier

// ProtocolIEField 36 - MobilityRestrictionList
func RATRestrictionsItemExtIEsExtensionValue (value ngapType.RATRestrictionsItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("RATRestrictionsItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ForbiddenAreaInformationItemExtIEsExtensionValue (value ngapType.ForbiddenAreaInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ForbiddenAreaInformationItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ServiceAreaInformationItemExtIEsExtensionValue (value ngapType.ServiceAreaInformationItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ServiceAreaInformationItemExtIEsExtensionValue: Present: INVALID")
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
func SliceOverloadItemExtIEs (value ngapType.SliceOverloadItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = SliceOverloadItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func SliceOverloadItemExtIEsExtensionValue (value ngapType.SliceOverloadItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SliceOverloadItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 50 - PagingDRX

// ProtocolIEField 51 - PagingOrigin

// ProtocolIEField 52 - PagingPriority

// ProtocolIEField 53 - PDUSessionResourceAdmittedList

// ProtocolIEField 54 - PDUSessionResourceFailedToModifyListModRes

// ProtocolIEField 55 - PDUSessionResourceFailedToSetupListCxtRes

// ProtocolIEField 56 - PDUSessionResourceFailedToSetupListHOAck

// ProtocolIEField 57 - PDUSessionResourceFailedToSetupListPSReq

// ProtocolIEField 58 - PDUSessionResourceFailedToSetupListSURes

// ProtocolIEField 59 - PDUSessionResourceHandoverList

// ProtocolIEField 60 - PDUSessionResourceListCxtRelCpl

// ProtocolIEField 61 - PDUSessionResourceListHORqd

// ProtocolIEField 62 - PDUSessionResourceModifyListModCfm

// ProtocolIEField 63 - PDUSessionResourceModifyListModInd

// ProtocolIEField 64 - PDUSessionResourceModifyListModReq

// ProtocolIEField 65 - PDUSessionResourceModifyListModRes

// ProtocolIEField 66 - PDUSessionResourceNotifyList

// ProtocolIEField 67 - PDUSessionResourceReleasedListNot

// ProtocolIEField 68 - PDUSessionResourceReleasedListPSAck

// ProtocolIEField 69 - PDUSessionResourceReleasedListPSFail

// ProtocolIEField 70 - PDUSessionResourceReleasedListRelRes

// ProtocolIEField 71 - PDUSessionResourceSetupListCxtReq

// ProtocolIEField 72 - PDUSessionResourceSetupListCxtRes

// ProtocolIEField 73 - PDUSessionResourceSetupListHOReq

// ProtocolIEField 74 - PDUSessionResourceSetupListSUReq

// ProtocolIEField 75 - PDUSessionResourceSetupListSURes

// ProtocolIEField 76 - PDUSessionResourceToBeSwitchedDLList

// ProtocolIEField 77 - PDUSessionResourceSwitchedList

// ProtocolIEField 78 - PDUSessionResourceToReleaseListHOCmd

// ProtocolIEField 79 - PDUSessionResourceToReleaseListRelCmd

// ProtocolIEField 80 - PLMNSupportList

// ProtocolIEField 81 - PWSFailedCellIDList

// ProtocolIEField 82 - RANNodeName

// ProtocolIEField 83 - RANPagingPriority

// ProtocolIEField 84 - RANStatusTransferTransparentContainer
func COUNTValueForPDCPSN12 (value ngapType.COUNTValueForPDCPSN12, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd = EncodeUint64(uint64(value.PDCPSN12), 0, binData, bitEnd)
	binData, bitEnd = EncodeUint64(uint64(value.HFNPDCPSN12), 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("COUNTValueForPDCPSN12: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusUL12ExtIEs (value ngapType.ProtocolExtensionContainerDRBStatusUL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL12ExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBStatusUL12ExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL12ExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
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
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd = EncodeUint64(uint64(value.PDCPSN18), 0, binData, bitEnd)
	binData, bitEnd = EncodeUint64(uint64(value.HFNPDCPSN18), 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("COUNTValueForPDCPSN18: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusUL18ExtIEs (value ngapType.ProtocolExtensionContainerDRBStatusUL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL18ExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBStatusUL18ExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusUL18ExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusDL12ExtIEs (value ngapType.ProtocolExtensionContainerDRBStatusDL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL12ExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBStatusDL12ExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL12ExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerDRBStatusDL18ExtIEs (value ngapType.ProtocolExtensionContainerDRBStatusDL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL18ExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = DRBStatusDL18ExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDRBStatusDL18ExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBsSubjectToStatusTransferItemExtIEsExtensionValue (value ngapType.DRBsSubjectToStatusTransferItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBsSubjectToStatusTransferItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs (value ngapType.ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = COUNTValueForPDCPSN12ExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusUL12ExtIEs (value ngapType.DRBStatusUL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = DRBStatusUL12ExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs (value ngapType.ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = COUNTValueForPDCPSN18ExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DRBStatusUL18ExtIEs (value ngapType.DRBStatusUL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = DRBStatusUL18ExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func DRBStatusDL12ExtIEs (value ngapType.DRBStatusDL12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = DRBStatusDL12ExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func DRBStatusDL18ExtIEs (value ngapType.DRBStatusDL18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = DRBStatusDL18ExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN12ExtIEs (value ngapType.COUNTValueForPDCPSN12ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = COUNTValueForPDCPSN12ExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func DRBStatusUL12ExtIEsExtensionValue (value ngapType.DRBStatusUL12ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusUL12ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN18ExtIEs (value ngapType.COUNTValueForPDCPSN18ExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	binData, bitEnd, err = COUNTValueForPDCPSN18ExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func DRBStatusUL18ExtIEsExtensionValue (value ngapType.DRBStatusUL18ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusUL18ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusDL12ExtIEsExtensionValue (value ngapType.DRBStatusDL12ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusDL12ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusDL18ExtIEsExtensionValue (value ngapType.DRBStatusDL18ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusDL18ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN12ExtIEsExtensionValue (value ngapType.COUNTValueForPDCPSN12ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN12ExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func COUNTValueForPDCPSN18ExtIEsExtensionValue (value ngapType.COUNTValueForPDCPSN18ExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("COUNTValueForPDCPSN18ExtIEsExtensionValue: Present: INVALID")
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

// ProtocolIEField 97 - SliceSupportList

// ProtocolIEField 98 - SONConfigurationTransferDL
func SONInformationReplyExtIEsExtensionValue (value ngapType.SONInformationReplyExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SONInformationReplyExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func XnExtTLAItemExtIEs (value ngapType.XnExtTLAItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = XnExtTLAItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func XnExtTLAItemExtIEsExtensionValue (value ngapType.XnExtTLAItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("XnExtTLAItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 99 - SONConfigurationTransferUL

// ProtocolIEField 100 - SourceAMFUENGAPID

// ProtocolIEField 101 - SourceToTargetTransparentContainer

// ProtocolIEField 102 - SupportedTAList
func BroadcastPLMNItemExtIEs (value ngapType.BroadcastPLMNItemExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEs: " + err.Error())
	}
	binData, bitEnd, err = BroadcastPLMNItemExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func BroadcastPLMNItemExtIEsExtensionValue (value ngapType.BroadcastPLMNItemExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("BroadcastPLMNItemExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 103 - TAIListForPaging

// ProtocolIEField 104 - TAIListForRestart

// ProtocolIEField 105 - TargetID
func EPSTAIExtIEsExtensionValue (value ngapType.EPSTAIExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("EPSTAIExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 106 - TargetToSourceTransparentContainer

// ProtocolIEField 107 - TimeToWait

// ProtocolIEField 108 - TraceActivation

// ProtocolIEField 109 - TraceCollectionEntityIPAddress

// ProtocolIEField 110 - UEAggregateMaximumBitRate

// ProtocolIEField 111 - UEAssociatedLogicalNGConnectionList

// ProtocolIEField 112 - UEContextRequest

// ProtocolIEField 113 - UENGAPIDs

// ProtocolIEField 114 - UEPagingIdentity

// ProtocolIEField 115 - UEPresenceInAreaOfInterestList

// ProtocolIEField 116 - UERadioCapability

// ProtocolIEField 117 - UERadioCapabilityForPaging

// ProtocolIEField 118 - UESecurityCapabilities

// ProtocolIEField 119 - UnavailableGUAMIList

// ProtocolIEField 120 - UserLocationInformation

// ProtocolIEField 121 - WarningAreaList

// ProtocolIEField 122 - WarningMessageContents

// ProtocolIEField 123 - WarningSecurityInfo

// ProtocolIEField 124 - WarningType

// ProtocolIEField 125 - AdditionalULNGUUPTNLInformation

// ProtocolIEField 126 - DataForwardingNotPossible

// ProtocolIEField 127 - DLNGUUPTNLInformation

// ProtocolIEField 128 - NetworkInstance

// ProtocolIEField 129 - PDUSessionAggregateMaximumBitRate

// ProtocolIEField 130 - PDUSessionResourceFailedToModifyListModCfm

// ProtocolIEField 131 - PDUSessionResourceFailedToSetupListCxtFail

// ProtocolIEField 132 - PDUSessionResourceListCxtRelReq

// ProtocolIEField 133 - PDUSessionType

// ProtocolIEField 134 - QosFlowAddOrModifyRequestList
func FiveQI (value ngapType.FiveQI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("FiveQI: int64: valueExt,valueLB:0,valueUB:255")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PriorityLevelQos (value ngapType.PriorityLevelQos, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 127 {
		return binData, bitEnd, errors.New("PriorityLevelQos: int64: valueExt,valueLB:1,valueUB:127")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func AveragingWindow (value ngapType.AveragingWindow, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4095 {
		return binData, bitEnd, errors.New("AveragingWindow: int64: valueExt,valueLB:0,valueUB:4095")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func MaximumDataBurstVolume (value ngapType.MaximumDataBurstVolume, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4095 {
		return binData, bitEnd, errors.New("MaximumDataBurstVolume: int64: valueExt,valueLB:0,valueUB:4095")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs (value ngapType.ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = NonDynamic5QIDescriptorExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func PacketDelayBudget (value ngapType.PacketDelayBudget, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1023 {
		return binData, bitEnd, errors.New("PacketDelayBudget: int64: valueExt,valueLB:0,valueUB:1023")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PacketErrorRate (value ngapType.PacketErrorRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	if value.IEExtensions != nil {
		option += 1<<0
	}
	binData, bitEnd = EncodeUint64(uint64(option), 2, binData, bitEnd)
	binData, bitEnd = EncodeUint64(uint64(value.PERScalar), 0, binData, bitEnd)
	binData, bitEnd = EncodeUint64(uint64(value.PERExponent), 0, binData, bitEnd)
	if value.IEExtensions != nil {
		binData, bitEnd, err = ProtocolExtensionContainerPacketErrorRateExtIEs(*value.IEExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PacketErrorRate: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func DelayCritical (value ngapType.DelayCritical, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("DelayCritical: Enumerated: valueExt,valueLB:0,valueUB:1")
	}
	binData, bitEnd = EncodeUint64(uint64(value.Value), 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ProtocolExtensionContainerDynamic5QIDescriptorExtIEs (value ngapType.ProtocolExtensionContainerDynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerDynamic5QIDescriptorExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = Dynamic5QIDescriptorExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerDynamic5QIDescriptorExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func AllocationAndRetentionPriorityExtIEs (value ngapType.AllocationAndRetentionPriorityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	binData, bitEnd, err = AllocationAndRetentionPriorityExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func GBRQosInformationExtIEs (value ngapType.GBRQosInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = GBRQosInformationExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GBRQosInformationExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func QosFlowLevelQosParametersExtIEsExtensionValue (value ngapType.QosFlowLevelQosParametersExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosFlowLevelQosParametersExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NonDynamic5QIDescriptorExtIEs (value ngapType.NonDynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = NonDynamic5QIDescriptorExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolExtensionContainerPacketErrorRateExtIEs (value ngapType.ProtocolExtensionContainerPacketErrorRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numItem := len(value.List)
	if numItem < 1 || numItem > 65535 {
		return binData, bitEnd, errors.New("ProtocolExtensionContainerPacketErrorRateExtIEs: List: sizeLB:1,sizeUB:65535")
	}
	binData, bitEnd = EncodeUint64(uint64(numItem-1), 0, binData, bitEnd)
	for i:=0; i<numItem; i++ {
		binData, bitEnd, err = PacketErrorRateExtIEs(value.List[i], binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolExtensionContainerPacketErrorRateExtIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
		}
	}
	return binData, bitEnd, err
}

func Dynamic5QIDescriptorExtIEs (value ngapType.Dynamic5QIDescriptorExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Dynamic5QIDescriptorExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func AllocationAndRetentionPriorityExtIEsExtensionValue (value ngapType.AllocationAndRetentionPriorityExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AllocationAndRetentionPriorityExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GBRQosInformationExtIEsExtensionValue (value ngapType.GBRQosInformationExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GBRQosInformationExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NonDynamic5QIDescriptorExtIEsExtensionValue (value ngapType.NonDynamic5QIDescriptorExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NonDynamic5QIDescriptorExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PacketErrorRateExtIEs (value ngapType.PacketErrorRateExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	var option uint64 = 0
	binData, bitEnd = EncodeUint64(uint64(option), 1, binData, bitEnd)
	binData, bitEnd, err = ProtocolExtensionID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PacketErrorRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PacketErrorRateExtIEs: " + err.Error())
	}
	binData, bitEnd, err = PacketErrorRateExtIEsExtensionValue(value.ExtensionValue, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PacketErrorRateExtIEs: " + err.Error())
	}
	return binData, bitEnd, err
}

func Dynamic5QIDescriptorExtIEsExtensionValue (value ngapType.Dynamic5QIDescriptorExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("Dynamic5QIDescriptorExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PacketErrorRateExtIEsExtensionValue (value ngapType.PacketErrorRateExtIEsExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 0, binData, bitEnd)
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PacketErrorRateExtIEsExtensionValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

// ProtocolIEField 135 - QosFlowSetupRequestList

// ProtocolIEField 136 - QosFlowToReleaseList

// ProtocolIEField 137 - SecurityIndication

// ProtocolIEField 138 - ULNGUUPTNLInformation

// ProtocolIEField 139 - ULNGUUPTNLModifyList

// ProtocolIEField 140 - WarningAreaCoordinates

