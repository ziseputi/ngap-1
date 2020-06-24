package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AMFPagingTarget (value ngapType.AMFPagingTarget, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.AMFPagingTargetPresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: " + err.Error())
		}
	case ngapType.AMFPagingTargetPresentTAI:
		if value.TAI == nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: TAI: NIL")
		}
		binData, bitEnd, err = TAI(*value.TAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: " + err.Error())
		}
	case ngapType.AMFPagingTargetPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("AMFPagingTarget: ChoiceExtensions: NIL")
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

func BroadcastCancelledAreaList (value ngapType.BroadcastCancelledAreaList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)
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

func BroadcastCompletedAreaList (value ngapType.BroadcastCompletedAreaList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)
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

func CellIDListForRestart (value ngapType.CellIDListForRestart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
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

func CPTransportLayerInformation (value ngapType.CPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
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

func DRBStatusDL (value ngapType.DRBStatusDL, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.DRBStatusDLPresentDRBStatusDL12:
		if value.DRBStatusDL12 == nil {
			return binData, bitEnd, errors.New("DRBStatusDL: DRBStatusDL12: NIL")
		}
		binData, bitEnd, err = DRBStatusDL12(*value.DRBStatusDL12, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL: " + err.Error())
		}
	case ngapType.DRBStatusDLPresentDRBStatusDL18:
		if value.DRBStatusDL18 == nil {
			return binData, bitEnd, errors.New("DRBStatusDL: DRBStatusDL18: NIL")
		}
		binData, bitEnd, err = DRBStatusDL18(*value.DRBStatusDL18, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusDL: " + err.Error())
		}
	case ngapType.DRBStatusDLPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("DRBStatusDL: ChoiceExtensions: NIL")
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

func DRBStatusUL (value ngapType.DRBStatusUL, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.DRBStatusULPresentDRBStatusUL12:
		if value.DRBStatusUL12 == nil {
			return binData, bitEnd, errors.New("DRBStatusUL: DRBStatusUL12: NIL")
		}
		binData, bitEnd, err = DRBStatusUL12(*value.DRBStatusUL12, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL: " + err.Error())
		}
	case ngapType.DRBStatusULPresentDRBStatusUL18:
		if value.DRBStatusUL18 == nil {
			return binData, bitEnd, errors.New("DRBStatusUL: DRBStatusUL18: NIL")
		}
		binData, bitEnd, err = DRBStatusUL18(*value.DRBStatusUL18, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DRBStatusUL: " + err.Error())
		}
	case ngapType.DRBStatusULPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("DRBStatusUL: ChoiceExtensions: NIL")
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

func GNBID (value ngapType.GNBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
	switch value.Present {
	case ngapType.GNBIDPresentGNBID:
		if value.GNBID == nil {
			return binData, bitEnd, errors.New("GNBID: GNBID: NIL")
		}
		if value.GNBID.BitLength < 22 || value.GNBID.BitLength > 32 || len(value.GNBID.Bytes) != int((value.GNBID.BitLength-1)/8+1) {
			return binData, bitEnd, errors.New("GNBID: GNBID: BitString: sizeMin:22,sizeMax:32")
		}
		binData, bitEnd = EncodeBitStringExt(*value.GNBID, binData, bitEnd)
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

func LastVisitedCellInformation (value ngapType.LastVisitedCellInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)
	switch value.Present {
	case ngapType.LastVisitedCellInformationPresentNGRANCell:
		if value.NGRANCell == nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: NGRANCell: NIL")
		}
		binData, bitEnd, err = LastVisitedNGRANCellInformation(*value.NGRANCell, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: " + err.Error())
		}
	case ngapType.LastVisitedCellInformationPresentEUTRANCell:
		if value.EUTRANCell == nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: EUTRANCell: NIL")
		}
		binData, bitEnd, err = LastVisitedEUTRANCellInformation(*value.EUTRANCell, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: " + err.Error())
		}
	case ngapType.LastVisitedCellInformationPresentUTRANCell:
		if value.UTRANCell == nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: UTRANCell: NIL")
		}
		binData, bitEnd, err = LastVisitedUTRANCellInformation(*value.UTRANCell, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: " + err.Error())
		}
	case ngapType.LastVisitedCellInformationPresentGERANCell:
		if value.GERANCell == nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: GERANCell: NIL")
		}
		binData, bitEnd, err = LastVisitedGERANCellInformation(*value.GERANCell, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: " + err.Error())
		}
	case ngapType.LastVisitedCellInformationPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: ChoiceExtensions: NIL")
		}
		binData, bitEnd, err = ProtocolIESingleContainerLastVisitedCellInformationExtIEs(*value.ChoiceExtensions, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LastVisitedCellInformation: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LastVisitedCellInformation: Present: INVALID")
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
		if value.N3IWFID.BitLength != 16 || len(value.N3IWFID.Bytes) != 2 {
			return binData, bitEnd, errors.New("N3IWFID: N3IWFID: BitString: sizeMin:16,sizeMax:16")
		}
		binData, bitEnd = EncodeBitStringFix(*value.N3IWFID, binData, bitEnd)
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

func NgENBID (value ngapType.NgENBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.NgENBIDPresentMacroNgENBID:
		if value.MacroNgENBID == nil {
			return binData, bitEnd, errors.New("NgENBID: MacroNgENBID: NIL")
		}
		if value.MacroNgENBID.BitLength != 20 || len(value.MacroNgENBID.Bytes) != 3 {
			return binData, bitEnd, errors.New("NgENBID: MacroNgENBID: BitString: sizeMin:20,sizeMax:20")
		}
		binData, bitEnd = EncodeBitStringFix(*value.MacroNgENBID, binData, bitEnd)
	case ngapType.NgENBIDPresentShortMacroNgENBID:
		if value.ShortMacroNgENBID == nil {
			return binData, bitEnd, errors.New("NgENBID: ShortMacroNgENBID: NIL")
		}
		if value.ShortMacroNgENBID.BitLength != 18 || len(value.ShortMacroNgENBID.Bytes) != 3 {
			return binData, bitEnd, errors.New("NgENBID: ShortMacroNgENBID: BitString: sizeMin:18,sizeMax:18")
		}
		binData, bitEnd = EncodeBitStringFix(*value.ShortMacroNgENBID, binData, bitEnd)
	case ngapType.NgENBIDPresentLongMacroNgENBID:
		if value.LongMacroNgENBID == nil {
			return binData, bitEnd, errors.New("NgENBID: LongMacroNgENBID: NIL")
		}
		if value.LongMacroNgENBID.BitLength != 21 || len(value.LongMacroNgENBID.Bytes) != 3 {
			return binData, bitEnd, errors.New("NgENBID: LongMacroNgENBID: BitString: sizeMin:21,sizeMax:21")
		}
		binData, bitEnd = EncodeBitStringFix(*value.LongMacroNgENBID, binData, bitEnd)
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

func OverloadResponse (value ngapType.OverloadResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
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

func PWSFailedCellIDList (value ngapType.PWSFailedCellIDList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
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

func QosCharacteristics (value ngapType.QosCharacteristics, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.QosCharacteristicsPresentNonDynamic5QI:
		if value.NonDynamic5QI == nil {
			return binData, bitEnd, errors.New("QosCharacteristics: NonDynamic5QI: NIL")
		}
		binData, bitEnd, err = NonDynamic5QIDescriptor(*value.NonDynamic5QI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosCharacteristics: " + err.Error())
		}
	case ngapType.QosCharacteristicsPresentDynamic5QI:
		if value.Dynamic5QI == nil {
			return binData, bitEnd, errors.New("QosCharacteristics: Dynamic5QI: NIL")
		}
		binData, bitEnd, err = Dynamic5QIDescriptor(*value.Dynamic5QI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("QosCharacteristics: " + err.Error())
		}
	case ngapType.QosCharacteristicsPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("QosCharacteristics: ChoiceExtensions: NIL")
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

func SONInformation (value ngapType.SONInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
	switch value.Present {
	case ngapType.SONInformationPresentSONInformationRequest:
		if value.SONInformationRequest == nil {
			return binData, bitEnd, errors.New("SONInformation: SONInformationRequest: NIL")
		}
		binData, bitEnd, err = SONInformationRequest(*value.SONInformationRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformation: " + err.Error())
		}
	case ngapType.SONInformationPresentSONInformationReply:
		if value.SONInformationReply == nil {
			return binData, bitEnd, errors.New("SONInformation: SONInformationReply: NIL")
		}
		binData, bitEnd, err = SONInformationReply(*value.SONInformationReply, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SONInformation: " + err.Error())
		}
	case ngapType.SONInformationPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("SONInformation: ChoiceExtensions: NIL")
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

func UEIdentityIndexValue (value ngapType.UEIdentityIndexValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
	switch value.Present {
	case ngapType.UEIdentityIndexValuePresentIndexLength10:
		if value.IndexLength10 == nil {
			return binData, bitEnd, errors.New("UEIdentityIndexValue: IndexLength10: NIL")
		}
		if value.IndexLength10.BitLength != 10 || len(value.IndexLength10.Bytes) != 2 {
			return binData, bitEnd, errors.New("UEIdentityIndexValue: IndexLength10: BitString: sizeMin:10,sizeMax:10")
		}
		binData, bitEnd = EncodeBitStringFix(*value.IndexLength10, binData, bitEnd)
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

func UENGAPIDs (value ngapType.UENGAPIDs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 2, binData, bitEnd)
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

func UEPagingIdentity (value ngapType.UEPagingIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
	switch value.Present {
	case ngapType.UEPagingIdentityPresentFiveGSTMSI:
		if value.FiveGSTMSI == nil {
			return binData, bitEnd, errors.New("UEPagingIdentity: FiveGSTMSI: NIL")
		}
		binData, bitEnd, err = FiveGSTMSI(*value.FiveGSTMSI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEPagingIdentity: " + err.Error())
		}
	case ngapType.UEPagingIdentityPresentChoiceExtensions:
		if value.ChoiceExtensions == nil {
			return binData, bitEnd, errors.New("UEPagingIdentity: ChoiceExtensions: NIL")
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

func UPTransportLayerInformation (value ngapType.UPTransportLayerInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 1, binData, bitEnd)
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

func WarningAreaList (value ngapType.WarningAreaList, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)
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
