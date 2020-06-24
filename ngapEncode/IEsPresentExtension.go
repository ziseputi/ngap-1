package ngapEncode

import (
	"errors"
	"strconv"

	"ngap/ngapType"
)

func ProtocolIESingleContainerAMFPagingTargetExtIEs(value ngapType.ProtocolIESingleContainerAMFPagingTargetExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerAMFPagingTargetExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFPagingTargetExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerAMFPagingTargetExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs(value ngapType.ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = BroadcastCancelledAreaListExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs(value ngapType.ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = BroadcastCompletedAreaListExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerCauseExtIEs(value ngapType.ProtocolIESingleContainerCauseExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerCauseExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CauseExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerCauseExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerCellIDListForRestartExtIEs(value ngapType.ProtocolIESingleContainerCellIDListForRestartExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerCellIDListForRestartExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellIDListForRestartExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerCellIDListForRestartExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerCPTransportLayerInformationExtIEs(value ngapType.ProtocolIESingleContainerCPTransportLayerInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerCPTransportLayerInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CPTransportLayerInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerCPTransportLayerInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerDRBStatusDLExtIEs(value ngapType.ProtocolIESingleContainerDRBStatusDLExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerDRBStatusDLExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBStatusDLExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerDRBStatusDLExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerDRBStatusULExtIEs(value ngapType.ProtocolIESingleContainerDRBStatusULExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerDRBStatusULExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DRBStatusULExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerDRBStatusULExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerGlobalRANNodeIDExtIEs(value ngapType.ProtocolIESingleContainerGlobalRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerGlobalRANNodeIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GlobalRANNodeIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerGlobalRANNodeIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerGNBIDExtIEs(value ngapType.ProtocolIESingleContainerGNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerGNBIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = GNBIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerGNBIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerLastVisitedCellInformationExtIEs(value ngapType.ProtocolIESingleContainerLastVisitedCellInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerLastVisitedCellInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LastVisitedCellInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerLastVisitedCellInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerN3IWFIDExtIEs(value ngapType.ProtocolIESingleContainerN3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerN3IWFIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = N3IWFIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerN3IWFIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerNgENBIDExtIEs(value ngapType.ProtocolIESingleContainerNgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerNgENBIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NgENBIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerNgENBIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerNGRANCGIExtIEs(value ngapType.ProtocolIESingleContainerNGRANCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerNGRANCGIExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGRANCGIExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerNGRANCGIExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerOverloadResponseExtIEs(value ngapType.ProtocolIESingleContainerOverloadResponseExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerOverloadResponseExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = OverloadResponseExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerOverloadResponseExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerPWSFailedCellIDListExtIEs(value ngapType.ProtocolIESingleContainerPWSFailedCellIDListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerPWSFailedCellIDListExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PWSFailedCellIDListExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerPWSFailedCellIDListExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerQosCharacteristicsExtIEs(value ngapType.ProtocolIESingleContainerQosCharacteristicsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerQosCharacteristicsExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = QosCharacteristicsExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerQosCharacteristicsExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerResetTypeExtIEs(value ngapType.ProtocolIESingleContainerResetTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerResetTypeExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ResetTypeExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerResetTypeExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerSONInformationExtIEs(value ngapType.ProtocolIESingleContainerSONInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerSONInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SONInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerSONInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerTargetIDExtIEs(value ngapType.ProtocolIESingleContainerTargetIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerTargetIDExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TargetIDExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerTargetIDExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerUEIdentityIndexValueExtIEs(value ngapType.ProtocolIESingleContainerUEIdentityIndexValueExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerUEIdentityIndexValueExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEIdentityIndexValueExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerUEIdentityIndexValueExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerUENGAPIDsExtIEs(value ngapType.ProtocolIESingleContainerUENGAPIDsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerUENGAPIDsExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UENGAPIDsExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerUENGAPIDsExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerUEPagingIdentityExtIEs(value ngapType.ProtocolIESingleContainerUEPagingIdentityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerUEPagingIdentityExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEPagingIdentityExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerUEPagingIdentityExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerUPTransportLayerInformationExtIEs(value ngapType.ProtocolIESingleContainerUPTransportLayerInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerUPTransportLayerInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UPTransportLayerInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerUPTransportLayerInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerUserLocationInformationExtIEs(value ngapType.ProtocolIESingleContainerUserLocationInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerUserLocationInformationExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UserLocationInformationExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerUserLocationInformationExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIESingleContainerWarningAreaListExtIEs(value ngapType.ProtocolIESingleContainerWarningAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIESingleContainerWarningAreaListExtIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = WarningAreaListExtIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIESingleContainerWarningAreaListExtIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}
