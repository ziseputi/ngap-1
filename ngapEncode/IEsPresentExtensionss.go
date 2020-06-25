package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AMFPagingTargetExtIEsTypeValue(value ngapType.AMFPagingTargetExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("AMFPagingTargetExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func BroadcastCancelledAreaListExtIEsTypeValue(value ngapType.BroadcastCancelledAreaListExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("BroadcastCancelledAreaListExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func BroadcastCompletedAreaListExtIEsTypeValue(value ngapType.BroadcastCompletedAreaListExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("BroadcastCompletedAreaListExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CauseExtIEsTypeValue(value ngapType.CauseExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CauseExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellIDListForRestartExtIEsTypeValue(value ngapType.CellIDListForRestartExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("CellIDListForRestartExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CPTransportLayerInformationExtIEsTypeValue(value ngapType.CPTransportLayerInformationExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.CPTransportLayerInformationExtIEsPresentEndpointIPAddressAndPort:
		if value.EndpointIPAddressAndPort == nil {
			return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEsTypeValue: EndpointIPAddressAndPort: NIL")
		}
		binData, bitEnd, err = EndpointIPAddressAndPort(*value.EndpointIPAddressAndPort, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("CPTransportLayerInformationExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusDLExtIEsTypeValue(value ngapType.DRBStatusDLExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusDLExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DRBStatusULExtIEsTypeValue(value ngapType.DRBStatusULExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("DRBStatusULExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GlobalRANNodeIDExtIEsTypeValue(value ngapType.GlobalRANNodeIDExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GlobalRANNodeIDExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func GNBIDExtIEsTypeValue(value ngapType.GNBIDExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("GNBIDExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LastVisitedCellInformationExtIEsTypeValue(value ngapType.LastVisitedCellInformationExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("LastVisitedCellInformationExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func N3IWFIDExtIEsTypeValue(value ngapType.N3IWFIDExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("N3IWFIDExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NgENBIDExtIEsTypeValue(value ngapType.NgENBIDExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NgENBIDExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGRANCGIExtIEsTypeValue(value ngapType.NGRANCGIExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("NGRANCGIExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func OverloadResponseExtIEsTypeValue(value ngapType.OverloadResponseExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("OverloadResponseExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSFailedCellIDListExtIEsTypeValue(value ngapType.PWSFailedCellIDListExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PWSFailedCellIDListExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func QosCharacteristicsExtIEsTypeValue(value ngapType.QosCharacteristicsExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("QosCharacteristicsExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ResetTypeExtIEsTypeValue(value ngapType.ResetTypeExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("ResetTypeExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SONInformationExtIEsTypeValue(value ngapType.SONInformationExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("SONInformationExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TargetIDExtIEsTypeValue(value ngapType.TargetIDExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.TargetIDExtIEsPresentTargetRNCID:
		if value.TargetRNCID == nil {
			return binData, bitEnd, errors.New("TargetIDExtIEsTypeValue: TargetRNCID: NIL")
		}
		binData, bitEnd, err = TargetRNCID(*value.TargetRNCID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TargetIDExtIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("TargetIDExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEIdentityIndexValueExtIEsTypeValue(value ngapType.UEIdentityIndexValueExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEIdentityIndexValueExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UENGAPIDsExtIEsTypeValue(value ngapType.UENGAPIDsExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UENGAPIDsExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEPagingIdentityExtIEsTypeValue(value ngapType.UEPagingIdentityExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UEPagingIdentityExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UPTransportLayerInformationExtIEsTypeValue(value ngapType.UPTransportLayerInformationExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UPTransportLayerInformationExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UserLocationInformationExtIEsTypeValue(value ngapType.UserLocationInformationExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("UserLocationInformationExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func WarningAreaListExtIEsTypeValue(value ngapType.WarningAreaListExtIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("WarningAreaListExtIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}
