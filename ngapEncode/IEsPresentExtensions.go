package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AMFPagingTargetExtIEs(value ngapType.AMFPagingTargetExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("AMFPagingTargetExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFPagingTargetExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFPagingTargetExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFPagingTargetExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFPagingTargetExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func BroadcastCancelledAreaListExtIEs(value ngapType.BroadcastCancelledAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("BroadcastCancelledAreaListExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastCancelledAreaListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastCancelledAreaListExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = BroadcastCancelledAreaListExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastCancelledAreaListExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func BroadcastCompletedAreaListExtIEs(value ngapType.BroadcastCompletedAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("BroadcastCompletedAreaListExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastCompletedAreaListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastCompletedAreaListExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = BroadcastCompletedAreaListExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("BroadcastCompletedAreaListExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CauseExtIEs(value ngapType.CauseExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("CauseExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CauseExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CauseExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CauseExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CauseExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellIDListForRestartExtIEs(value ngapType.CellIDListForRestartExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("CellIDListForRestartExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDListForRestartExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDListForRestartExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellIDListForRestartExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellIDListForRestartExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CPTransportLayerInformationExtIEs(value ngapType.CPTransportLayerInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDEndpointIPAddressAndPort:
		if value.TypeValue.Present != ngapType.CPTransportLayerInformationExtIEsPresentEndpointIPAddressAndPort {
			return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CPTransportLayerInformationExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBStatusDLExtIEs(value ngapType.DRBStatusDLExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("DRBStatusDLExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDLExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDLExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBStatusDLExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusDLExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DRBStatusULExtIEs(value ngapType.DRBStatusULExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("DRBStatusULExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusULExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusULExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DRBStatusULExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DRBStatusULExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GlobalRANNodeIDExtIEs(value ngapType.GlobalRANNodeIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("GlobalRANNodeIDExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalRANNodeIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalRANNodeIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GlobalRANNodeIDExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GlobalRANNodeIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func GNBIDExtIEs(value ngapType.GNBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("GNBIDExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GNBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("GNBIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = GNBIDExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("GNBIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LastVisitedCellInformationExtIEs(value ngapType.LastVisitedCellInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("LastVisitedCellInformationExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LastVisitedCellInformationExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedCellInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func N3IWFIDExtIEs(value ngapType.N3IWFIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("N3IWFIDExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("N3IWFIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("N3IWFIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = N3IWFIDExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("N3IWFIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NgENBIDExtIEs(value ngapType.NgENBIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("NgENBIDExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NgENBIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NgENBIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NgENBIDExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NgENBIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGRANCGIExtIEs(value ngapType.NGRANCGIExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("NGRANCGIExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANCGIExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANCGIExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGRANCGIExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGRANCGIExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func OverloadResponseExtIEs(value ngapType.OverloadResponseExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("OverloadResponseExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadResponseExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadResponseExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = OverloadResponseExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadResponseExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PWSFailedCellIDListExtIEs(value ngapType.PWSFailedCellIDListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("PWSFailedCellIDListExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailedCellIDListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailedCellIDListExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSFailedCellIDListExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailedCellIDListExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func QosCharacteristicsExtIEs(value ngapType.QosCharacteristicsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("QosCharacteristicsExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosCharacteristicsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("QosCharacteristicsExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = QosCharacteristicsExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("QosCharacteristicsExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ResetTypeExtIEs(value ngapType.ResetTypeExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("ResetTypeExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ResetTypeExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ResetTypeExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ResetTypeExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ResetTypeExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SONInformationExtIEs(value ngapType.SONInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("SONInformationExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SONInformationExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SONInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TargetIDExtIEs(value ngapType.TargetIDExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDTargetRNCID:
		if value.TypeValue.Present != ngapType.TargetIDExtIEsPresentTargetRNCID {
			return binData, bitEnd, errors.New("TargetIDExtIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("TargetIDExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetIDExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetIDExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TargetIDExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TargetIDExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEIdentityIndexValueExtIEs(value ngapType.UEIdentityIndexValueExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("UEIdentityIndexValueExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEIdentityIndexValueExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEIdentityIndexValueExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEIdentityIndexValueExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEIdentityIndexValueExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UENGAPIDsExtIEs(value ngapType.UENGAPIDsExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("UENGAPIDsExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDsExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDsExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UENGAPIDsExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UENGAPIDsExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEPagingIdentityExtIEs(value ngapType.UEPagingIdentityExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("UEPagingIdentityExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPagingIdentityExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEPagingIdentityExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEPagingIdentityExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEPagingIdentityExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UPTransportLayerInformationExtIEs(value ngapType.UPTransportLayerInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("UPTransportLayerInformationExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UPTransportLayerInformationExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UPTransportLayerInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UserLocationInformationExtIEs(value ngapType.UserLocationInformationExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("UserLocationInformationExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UserLocationInformationExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UserLocationInformationExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func WarningAreaListExtIEs(value ngapType.WarningAreaListExtIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("WarningAreaListExtIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WarningAreaListExtIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WarningAreaListExtIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = WarningAreaListExtIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("WarningAreaListExtIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}
