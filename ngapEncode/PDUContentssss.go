package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func PDUSessionResourceSetupRequestIEsTypeValue(value ngapType.PDUSessionResourceSetupRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceSetupRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsPresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsPresentPDUSessionResourceSetupListSUReq:
		if value.PDUSessionResourceSetupListSUReq == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: PDUSessionResourceSetupListSUReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListSUReq(*value.PDUSessionResourceSetupListSUReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsPresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupResponseIEsTypeValue(value ngapType.PDUSessionResourceSetupResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceSetupResponseIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceSetupListSURes:
		if value.PDUSessionResourceSetupListSURes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: PDUSessionResourceSetupListSURes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListSURes(*value.PDUSessionResourceSetupListSURes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceFailedToSetupListSURes:
		if value.PDUSessionResourceFailedToSetupListSURes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: PDUSessionResourceFailedToSetupListSURes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListSURes(*value.PDUSessionResourceFailedToSetupListSURes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseCommandIEsTypeValue(value ngapType.PDUSessionResourceReleaseCommandIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceReleaseCommandIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsPresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsPresentPDUSessionResourceToReleaseListRelCmd:
		if value.PDUSessionResourceToReleaseListRelCmd == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: PDUSessionResourceToReleaseListRelCmd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToReleaseListRelCmd(*value.PDUSessionResourceToReleaseListRelCmd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseResponseIEsTypeValue(value ngapType.PDUSessionResourceReleaseResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceReleaseResponseIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsPresentPDUSessionResourceReleasedListRelRes:
		if value.PDUSessionResourceReleasedListRelRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: PDUSessionResourceReleasedListRelRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListRelRes(*value.PDUSessionResourceReleasedListRelRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyRequestIEsTypeValue(value ngapType.PDUSessionResourceModifyRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsPresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsPresentPDUSessionResourceModifyListModReq:
		if value.PDUSessionResourceModifyListModReq == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: PDUSessionResourceModifyListModReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModReq(*value.PDUSessionResourceModifyListModReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyResponseIEsTypeValue(value ngapType.PDUSessionResourceModifyResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyResponseIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceModifyListModRes:
		if value.PDUSessionResourceModifyListModRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: PDUSessionResourceModifyListModRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModRes(*value.PDUSessionResourceModifyListModRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceFailedToModifyListModRes:
		if value.PDUSessionResourceFailedToModifyListModRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: PDUSessionResourceFailedToModifyListModRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToModifyListModRes(*value.PDUSessionResourceFailedToModifyListModRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyIEsTypeValue(value ngapType.PDUSessionResourceNotifyIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceNotifyIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsPresentPDUSessionResourceNotifyList:
		if value.PDUSessionResourceNotifyList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: PDUSessionResourceNotifyList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceNotifyList(*value.PDUSessionResourceNotifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsPresentPDUSessionResourceReleasedListNot:
		if value.PDUSessionResourceReleasedListNot == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: PDUSessionResourceReleasedListNot: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListNot(*value.PDUSessionResourceReleasedListNot, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndicationIEsTypeValue(value ngapType.PDUSessionResourceModifyIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyIndicationIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsPresentPDUSessionResourceModifyListModInd:
		if value.PDUSessionResourceModifyListModInd == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: PDUSessionResourceModifyListModInd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModInd(*value.PDUSessionResourceModifyListModInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyConfirmIEsTypeValue(value ngapType.PDUSessionResourceModifyConfirmIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyConfirmIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceModifyListModCfm:
		if value.PDUSessionResourceModifyListModCfm == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: PDUSessionResourceModifyListModCfm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModCfm(*value.PDUSessionResourceModifyListModCfm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceFailedToModifyListModCfm:
		if value.PDUSessionResourceFailedToModifyListModCfm == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: PDUSessionResourceFailedToModifyListModCfm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToModifyListModCfm(*value.PDUSessionResourceFailedToModifyListModCfm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialContextSetupRequestIEsTypeValue(value ngapType.InitialContextSetupRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialContextSetupRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentOldAMF:
		if value.OldAMF == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: OldAMF: NIL")
		}
		binData, bitEnd, err = AMFName(*value.OldAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentGUAMI:
		if value.GUAMI == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: GUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.GUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentPDUSessionResourceSetupListCxtReq:
		if value.PDUSessionResourceSetupListCxtReq == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: PDUSessionResourceSetupListCxtReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListCxtReq(*value.PDUSessionResourceSetupListCxtReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentSecurityKey:
		if value.SecurityKey == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: SecurityKey: NIL")
		}
		binData, bitEnd, err = SecurityKey(*value.SecurityKey, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentMaskedIMEISV:
		if value.MaskedIMEISV == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: MaskedIMEISV: NIL")
		}
		binData, bitEnd, err = MaskedIMEISV(*value.MaskedIMEISV, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentEmergencyFallbackIndicator:
		if value.EmergencyFallbackIndicator == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: EmergencyFallbackIndicator: NIL")
		}
		binData, bitEnd, err = EmergencyFallbackIndicator(*value.EmergencyFallbackIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentRedirectionVoiceFallback:
		if value.RedirectionVoiceFallback == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: RedirectionVoiceFallback: NIL")
		}
		binData, bitEnd, err = RedirectionVoiceFallback(*value.RedirectionVoiceFallback, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsPresentSRVCCOperationPossible:
		if value.SRVCCOperationPossible == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: SRVCCOperationPossible: NIL")
		}
		binData, bitEnd, err = SRVCCOperationPossible(*value.SRVCCOperationPossible, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialContextSetupResponseIEsTypeValue(value ngapType.InitialContextSetupResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialContextSetupResponseIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsPresentPDUSessionResourceSetupListCxtRes:
		if value.PDUSessionResourceSetupListCxtRes == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: PDUSessionResourceSetupListCxtRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListCxtRes(*value.PDUSessionResourceSetupListCxtRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsPresentPDUSessionResourceFailedToSetupListCxtRes:
		if value.PDUSessionResourceFailedToSetupListCxtRes == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: PDUSessionResourceFailedToSetupListCxtRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListCxtRes(*value.PDUSessionResourceFailedToSetupListCxtRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialContextSetupFailureIEsTypeValue(value ngapType.InitialContextSetupFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialContextSetupFailureIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsPresentPDUSessionResourceFailedToSetupListCxtFail:
		if value.PDUSessionResourceFailedToSetupListCxtFail == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: PDUSessionResourceFailedToSetupListCxtFail: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListCxtFail(*value.PDUSessionResourceFailedToSetupListCxtFail, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextReleaseRequestIEsTypeValue(value ngapType.UEContextReleaseRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextReleaseRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsPresentPDUSessionResourceListCxtRelReq:
		if value.PDUSessionResourceListCxtRelReq == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: PDUSessionResourceListCxtRelReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListCxtRelReq(*value.PDUSessionResourceListCxtRelReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextReleaseCommandIEsTypeValue(value ngapType.UEContextReleaseCommandIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextReleaseCommandIEsPresentUENGAPIDs:
		if value.UENGAPIDs == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: UENGAPIDs: NIL")
		}
		binData, bitEnd, err = UENGAPIDs(*value.UENGAPIDs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCommandIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextReleaseCompleteIEsTypeValue(value ngapType.UEContextReleaseCompleteIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextReleaseCompleteIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsPresentInfoOnRecommendedCellsAndRANNodesForPaging:
		if value.InfoOnRecommendedCellsAndRANNodesForPaging == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: InfoOnRecommendedCellsAndRANNodesForPaging: NIL")
		}
		binData, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPaging(*value.InfoOnRecommendedCellsAndRANNodesForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsPresentPDUSessionResourceListCxtRelCpl:
		if value.PDUSessionResourceListCxtRelCpl == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: PDUSessionResourceListCxtRelCpl: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListCxtRelCpl(*value.PDUSessionResourceListCxtRelCpl, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextModificationRequestIEsTypeValue(value ngapType.UEContextModificationRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextModificationRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentSecurityKey:
		if value.SecurityKey == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: SecurityKey: NIL")
		}
		binData, bitEnd, err = SecurityKey(*value.SecurityKey, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentEmergencyFallbackIndicator:
		if value.EmergencyFallbackIndicator == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: EmergencyFallbackIndicator: NIL")
		}
		binData, bitEnd, err = EmergencyFallbackIndicator(*value.EmergencyFallbackIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentNewAMFUENGAPID:
		if value.NewAMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: NewAMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.NewAMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentNewGUAMI:
		if value.NewGUAMI == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: NewGUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.NewGUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsPresentSRVCCOperationPossible:
		if value.SRVCCOperationPossible == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: SRVCCOperationPossible: NIL")
		}
		binData, bitEnd, err = SRVCCOperationPossible(*value.SRVCCOperationPossible, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextModificationResponseIEsTypeValue(value ngapType.UEContextModificationResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextModificationResponseIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsPresentRRCState:
		if value.RRCState == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: RRCState: NIL")
		}
		binData, bitEnd, err = RRCState(*value.RRCState, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextModificationFailureIEsTypeValue(value ngapType.UEContextModificationFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextModificationFailureIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RRCInactiveTransitionReportIEsTypeValue(value ngapType.RRCInactiveTransitionReportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RRCInactiveTransitionReportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsPresentRRCState:
		if value.RRCState == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: RRCState: NIL")
		}
		binData, bitEnd, err = RRCState(*value.RRCState, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequiredIEsTypeValue(value ngapType.HandoverRequiredIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequiredIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentTargetID:
		if value.TargetID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: TargetID: NIL")
		}
		binData, bitEnd, err = TargetID(*value.TargetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentDirectForwardingPathAvailability:
		if value.DirectForwardingPathAvailability == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: DirectForwardingPathAvailability: NIL")
		}
		binData, bitEnd, err = DirectForwardingPathAvailability(*value.DirectForwardingPathAvailability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentPDUSessionResourceListHORqd:
		if value.PDUSessionResourceListHORqd == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: PDUSessionResourceListHORqd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListHORqd(*value.PDUSessionResourceListHORqd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsPresentSourceToTargetTransparentContainer:
		if value.SourceToTargetTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: SourceToTargetTransparentContainer: NIL")
		}
		binData, bitEnd, err = SourceToTargetTransparentContainer(*value.SourceToTargetTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCommandIEsTypeValue(value ngapType.HandoverCommandIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCommandIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentNASSecurityParametersFromNGRAN:
		if value.NASSecurityParametersFromNGRAN == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: NASSecurityParametersFromNGRAN: NIL")
		}
		binData, bitEnd, err = NASSecurityParametersFromNGRAN(*value.NASSecurityParametersFromNGRAN, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentPDUSessionResourceHandoverList:
		if value.PDUSessionResourceHandoverList == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: PDUSessionResourceHandoverList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceHandoverList(*value.PDUSessionResourceHandoverList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentPDUSessionResourceToReleaseListHOCmd:
		if value.PDUSessionResourceToReleaseListHOCmd == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: PDUSessionResourceToReleaseListHOCmd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToReleaseListHOCmd(*value.PDUSessionResourceToReleaseListHOCmd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentTargetToSourceTransparentContainer:
		if value.TargetToSourceTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: TargetToSourceTransparentContainer: NIL")
		}
		binData, bitEnd, err = TargetToSourceTransparentContainer(*value.TargetToSourceTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverPreparationFailureIEsTypeValue(value ngapType.HandoverPreparationFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverPreparationFailureIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequestIEsTypeValue(value ngapType.HandoverRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentSecurityContext:
		if value.SecurityContext == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: SecurityContext: NIL")
		}
		binData, bitEnd, err = SecurityContext(*value.SecurityContext, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentNewSecurityContextInd:
		if value.NewSecurityContextInd == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: NewSecurityContextInd: NIL")
		}
		binData, bitEnd, err = NewSecurityContextInd(*value.NewSecurityContextInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentNASC:
		if value.NASC == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: NASC: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASC, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentPDUSessionResourceSetupListHOReq:
		if value.PDUSessionResourceSetupListHOReq == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: PDUSessionResourceSetupListHOReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListHOReq(*value.PDUSessionResourceSetupListHOReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentMaskedIMEISV:
		if value.MaskedIMEISV == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: MaskedIMEISV: NIL")
		}
		binData, bitEnd, err = MaskedIMEISV(*value.MaskedIMEISV, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentSourceToTargetTransparentContainer:
		if value.SourceToTargetTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: SourceToTargetTransparentContainer: NIL")
		}
		binData, bitEnd, err = SourceToTargetTransparentContainer(*value.SourceToTargetTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentGUAMI:
		if value.GUAMI == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: GUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.GUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentRedirectionVoiceFallback:
		if value.RedirectionVoiceFallback == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: RedirectionVoiceFallback: NIL")
		}
		binData, bitEnd, err = RedirectionVoiceFallback(*value.RedirectionVoiceFallback, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsPresentSRVCCOperationPossible:
		if value.SRVCCOperationPossible == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: SRVCCOperationPossible: NIL")
		}
		binData, bitEnd, err = SRVCCOperationPossible(*value.SRVCCOperationPossible, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequestAcknowledgeIEsTypeValue(value ngapType.HandoverRequestAcknowledgeIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequestAcknowledgeIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsPresentPDUSessionResourceAdmittedList:
		if value.PDUSessionResourceAdmittedList == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: PDUSessionResourceAdmittedList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceAdmittedList(*value.PDUSessionResourceAdmittedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsPresentPDUSessionResourceFailedToSetupListHOAck:
		if value.PDUSessionResourceFailedToSetupListHOAck == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: PDUSessionResourceFailedToSetupListHOAck: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListHOAck(*value.PDUSessionResourceFailedToSetupListHOAck, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsPresentTargetToSourceTransparentContainer:
		if value.TargetToSourceTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: TargetToSourceTransparentContainer: NIL")
		}
		binData, bitEnd, err = TargetToSourceTransparentContainer(*value.TargetToSourceTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverFailureIEsTypeValue(value ngapType.HandoverFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverFailureIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverNotifyIEsTypeValue(value ngapType.HandoverNotifyIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverNotifyIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverNotifyIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverNotifyIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestIEsTypeValue(value ngapType.PathSwitchRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsPresentSourceAMFUENGAPID:
		if value.SourceAMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: SourceAMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.SourceAMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsPresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsPresentPDUSessionResourceToBeSwitchedDLList:
		if value.PDUSessionResourceToBeSwitchedDLList == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: PDUSessionResourceToBeSwitchedDLList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLList(*value.PDUSessionResourceToBeSwitchedDLList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsPresentPDUSessionResourceFailedToSetupListPSReq:
		if value.PDUSessionResourceFailedToSetupListPSReq == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: PDUSessionResourceFailedToSetupListPSReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListPSReq(*value.PDUSessionResourceFailedToSetupListPSReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestAcknowledgeIEsTypeValue(value ngapType.PathSwitchRequestAcknowledgeIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentSecurityContext:
		if value.SecurityContext == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: SecurityContext: NIL")
		}
		binData, bitEnd, err = SecurityContext(*value.SecurityContext, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentNewSecurityContextInd:
		if value.NewSecurityContextInd == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: NewSecurityContextInd: NIL")
		}
		binData, bitEnd, err = NewSecurityContextInd(*value.NewSecurityContextInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceSwitchedList:
		if value.PDUSessionResourceSwitchedList == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: PDUSessionResourceSwitchedList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSwitchedList(*value.PDUSessionResourceSwitchedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceReleasedListPSAck:
		if value.PDUSessionResourceReleasedListPSAck == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: PDUSessionResourceReleasedListPSAck: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListPSAck(*value.PDUSessionResourceReleasedListPSAck, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentRedirectionVoiceFallback:
		if value.RedirectionVoiceFallback == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: RedirectionVoiceFallback: NIL")
		}
		binData, bitEnd, err = RedirectionVoiceFallback(*value.RedirectionVoiceFallback, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsPresentSRVCCOperationPossible:
		if value.SRVCCOperationPossible == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: SRVCCOperationPossible: NIL")
		}
		binData, bitEnd, err = SRVCCOperationPossible(*value.SRVCCOperationPossible, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestFailureIEsTypeValue(value ngapType.PathSwitchRequestFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestFailureIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsPresentPDUSessionResourceReleasedListPSFail:
		if value.PDUSessionResourceReleasedListPSFail == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: PDUSessionResourceReleasedListPSFail: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListPSFail(*value.PDUSessionResourceReleasedListPSFail, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCancelIEsTypeValue(value ngapType.HandoverCancelIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCancelIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCancelAcknowledgeIEsTypeValue(value ngapType.HandoverCancelAcknowledgeIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCancelAcknowledgeIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelAcknowledgeIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkRANStatusTransferIEsTypeValue(value ngapType.UplinkRANStatusTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkRANStatusTransferIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkRANStatusTransferIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer:
		if value.RANStatusTransferTransparentContainer == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: RANStatusTransferTransparentContainer: NIL")
		}
		binData, bitEnd, err = RANStatusTransferTransparentContainer(*value.RANStatusTransferTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkRANStatusTransferIEsTypeValue(value ngapType.DownlinkRANStatusTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkRANStatusTransferIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkRANStatusTransferIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer:
		if value.RANStatusTransferTransparentContainer == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: RANStatusTransferTransparentContainer: NIL")
		}
		binData, bitEnd, err = RANStatusTransferTransparentContainer(*value.RANStatusTransferTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PagingIEsTypeValue(value ngapType.PagingIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PagingIEsPresentUEPagingIdentity:
		if value.UEPagingIdentity == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: UEPagingIdentity: NIL")
		}
		binData, bitEnd, err = UEPagingIdentity(*value.UEPagingIdentity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsPresentPagingDRX:
		if value.PagingDRX == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: PagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.PagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsPresentTAIListForPaging:
		if value.TAIListForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: TAIListForPaging: NIL")
		}
		binData, bitEnd, err = TAIListForPaging(*value.TAIListForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsPresentPagingPriority:
		if value.PagingPriority == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: PagingPriority: NIL")
		}
		binData, bitEnd, err = PagingPriority(*value.PagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsPresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsPresentPagingOrigin:
		if value.PagingOrigin == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: PagingOrigin: NIL")
		}
		binData, bitEnd, err = PagingOrigin(*value.PagingOrigin, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsPresentAssistanceDataForPaging:
		if value.AssistanceDataForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: AssistanceDataForPaging: NIL")
		}
		binData, bitEnd, err = AssistanceDataForPaging(*value.AssistanceDataForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PagingIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialUEMessageIEsTypeValue(value ngapType.InitialUEMessageIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialUEMessageIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentRRCEstablishmentCause:
		if value.RRCEstablishmentCause == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: RRCEstablishmentCause: NIL")
		}
		binData, bitEnd, err = RRCEstablishmentCause(*value.RRCEstablishmentCause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentFiveGSTMSI:
		if value.FiveGSTMSI == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: FiveGSTMSI: NIL")
		}
		binData, bitEnd, err = FiveGSTMSI(*value.FiveGSTMSI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentAMFSetID:
		if value.AMFSetID == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: AMFSetID: NIL")
		}
		binData, bitEnd, err = AMFSetID(*value.AMFSetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentUEContextRequest:
		if value.UEContextRequest == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: UEContextRequest: NIL")
		}
		binData, bitEnd, err = UEContextRequest(*value.UEContextRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsPresentSourceToTargetAMFInformationReroute:
		if value.SourceToTargetAMFInformationReroute == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: SourceToTargetAMFInformationReroute: NIL")
		}
		binData, bitEnd, err = SourceToTargetAMFInformationReroute(*value.SourceToTargetAMFInformationReroute, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkNASTransportIEsTypeValue(value ngapType.DownlinkNASTransportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkNASTransportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentOldAMF:
		if value.OldAMF == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: OldAMF: NIL")
		}
		binData, bitEnd, err = AMFName(*value.OldAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsPresentSRVCCOperationPossible:
		if value.SRVCCOperationPossible == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: SRVCCOperationPossible: NIL")
		}
		binData, bitEnd, err = SRVCCOperationPossible(*value.SRVCCOperationPossible, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkNASTransportIEsTypeValue(value ngapType.UplinkNASTransportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkNASTransportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NASNonDeliveryIndicationIEsTypeValue(value ngapType.NASNonDeliveryIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NASNonDeliveryIndicationIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsPresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RerouteNASRequestIEsTypeValue(value ngapType.RerouteNASRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RerouteNASRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsPresentNGAPMessage:
		if value.NGAPMessage == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: NGAPMessage: NIL")
		}
		binData, bitEnd, err = NGAPMessage(*value.NGAPMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsPresentAMFSetID:
		if value.AMFSetID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: AMFSetID: NIL")
		}
		binData, bitEnd, err = AMFSetID(*value.AMFSetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsPresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsPresentSourceToTargetAMFInformationReroute:
		if value.SourceToTargetAMFInformationReroute == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: SourceToTargetAMFInformationReroute: NIL")
		}
		binData, bitEnd, err = SourceToTargetAMFInformationReroute(*value.SourceToTargetAMFInformationReroute, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGSetupRequestIEsTypeValue(value ngapType.NGSetupRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGSetupRequestIEsPresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsPresentRANNodeName:
		if value.RANNodeName == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: RANNodeName: NIL")
		}
		binData, bitEnd, err = RANNodeName(*value.RANNodeName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsPresentSupportedTAList:
		if value.SupportedTAList == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: SupportedTAList: NIL")
		}
		binData, bitEnd, err = SupportedTAList(*value.SupportedTAList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsPresentDefaultPagingDRX:
		if value.DefaultPagingDRX == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: DefaultPagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.DefaultPagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsPresentUERetentionInformation:
		if value.UERetentionInformation == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: UERetentionInformation: NIL")
		}
		binData, bitEnd, err = UERetentionInformation(*value.UERetentionInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGSetupResponseIEsTypeValue(value ngapType.NGSetupResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGSetupResponseIEsPresentAMFName:
		if value.AMFName == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: AMFName: NIL")
		}
		binData, bitEnd, err = AMFName(*value.AMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsPresentServedGUAMIList:
		if value.ServedGUAMIList == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: ServedGUAMIList: NIL")
		}
		binData, bitEnd, err = ServedGUAMIList(*value.ServedGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsPresentRelativeAMFCapacity:
		if value.RelativeAMFCapacity == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: RelativeAMFCapacity: NIL")
		}
		binData, bitEnd, err = RelativeAMFCapacity(*value.RelativeAMFCapacity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsPresentPLMNSupportList:
		if value.PLMNSupportList == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: PLMNSupportList: NIL")
		}
		binData, bitEnd, err = PLMNSupportList(*value.PLMNSupportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsPresentUERetentionInformation:
		if value.UERetentionInformation == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: UERetentionInformation: NIL")
		}
		binData, bitEnd, err = UERetentionInformation(*value.UERetentionInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGSetupFailureIEsTypeValue(value ngapType.NGSetupFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGSetupFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupFailureIEsPresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANConfigurationUpdateIEsTypeValue(value ngapType.RANConfigurationUpdateIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RANConfigurationUpdateIEsPresentRANNodeName:
		if value.RANNodeName == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: RANNodeName: NIL")
		}
		binData, bitEnd, err = RANNodeName(*value.RANNodeName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsPresentSupportedTAList:
		if value.SupportedTAList == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: SupportedTAList: NIL")
		}
		binData, bitEnd, err = SupportedTAList(*value.SupportedTAList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsPresentDefaultPagingDRX:
		if value.DefaultPagingDRX == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: DefaultPagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.DefaultPagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsPresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsPresentNGRANTNLAssociationToRemoveList:
		if value.NGRANTNLAssociationToRemoveList == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: NGRANTNLAssociationToRemoveList: NIL")
		}
		binData, bitEnd, err = NGRANTNLAssociationToRemoveList(*value.NGRANTNLAssociationToRemoveList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANConfigurationUpdateAcknowledgeIEsTypeValue(value ngapType.RANConfigurationUpdateAcknowledgeIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RANConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANConfigurationUpdateFailureIEsTypeValue(value ngapType.RANConfigurationUpdateFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RANConfigurationUpdateFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateFailureIEsPresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFConfigurationUpdateIEsTypeValue(value ngapType.AMFConfigurationUpdateIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFConfigurationUpdateIEsPresentAMFName:
		if value.AMFName == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFName: NIL")
		}
		binData, bitEnd, err = AMFName(*value.AMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsPresentServedGUAMIList:
		if value.ServedGUAMIList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: ServedGUAMIList: NIL")
		}
		binData, bitEnd, err = ServedGUAMIList(*value.ServedGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsPresentRelativeAMFCapacity:
		if value.RelativeAMFCapacity == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: RelativeAMFCapacity: NIL")
		}
		binData, bitEnd, err = RelativeAMFCapacity(*value.RelativeAMFCapacity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsPresentPLMNSupportList:
		if value.PLMNSupportList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: PLMNSupportList: NIL")
		}
		binData, bitEnd, err = PLMNSupportList(*value.PLMNSupportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToAddList:
		if value.AMFTNLAssociationToAddList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFTNLAssociationToAddList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToAddList(*value.AMFTNLAssociationToAddList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToRemoveList:
		if value.AMFTNLAssociationToRemoveList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFTNLAssociationToRemoveList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToRemoveList(*value.AMFTNLAssociationToRemoveList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToUpdateList:
		if value.AMFTNLAssociationToUpdateList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFTNLAssociationToUpdateList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToUpdateList(*value.AMFTNLAssociationToUpdateList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFConfigurationUpdateAcknowledgeIEsTypeValue(value ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationSetupList:
		if value.AMFTNLAssociationSetupList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: AMFTNLAssociationSetupList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationSetupList(*value.AMFTNLAssociationSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationFailedToSetupList:
		if value.AMFTNLAssociationFailedToSetupList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: AMFTNLAssociationFailedToSetupList: NIL")
		}
		binData, bitEnd, err = TNLAssociationList(*value.AMFTNLAssociationFailedToSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFConfigurationUpdateFailureIEsTypeValue(value ngapType.AMFConfigurationUpdateFailureIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFConfigurationUpdateFailureIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateFailureIEsPresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateFailureIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFStatusIndicationIEsTypeValue(value ngapType.AMFStatusIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFStatusIndicationIEsPresentUnavailableGUAMIList:
		if value.UnavailableGUAMIList == nil {
			return binData, bitEnd, errors.New("AMFStatusIndicationIEsTypeValue: UnavailableGUAMIList: NIL")
		}
		binData, bitEnd, err = UnavailableGUAMIList(*value.UnavailableGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFStatusIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFStatusIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGResetIEsTypeValue(value ngapType.NGResetIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGResetIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NGResetIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetIEsTypeValue: " + err.Error())
		}
	case ngapType.NGResetIEsPresentResetType:
		if value.ResetType == nil {
			return binData, bitEnd, errors.New("NGResetIEsTypeValue: ResetType: NIL")
		}
		binData, bitEnd, err = ResetType(*value.ResetType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGResetIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGResetAcknowledgeIEsTypeValue(value ngapType.NGResetAcknowledgeIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGResetAcknowledgeIEsPresentUEAssociatedLogicalNGConnectionList:
		if value.UEAssociatedLogicalNGConnectionList == nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: UEAssociatedLogicalNGConnectionList: NIL")
		}
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionList(*value.UEAssociatedLogicalNGConnectionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.NGResetAcknowledgeIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ErrorIndicationIEsTypeValue(value ngapType.ErrorIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.ErrorIndicationIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func OverloadStartIEsTypeValue(value ngapType.OverloadStartIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.OverloadStartIEsPresentAMFOverloadResponse:
		if value.AMFOverloadResponse == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: AMFOverloadResponse: NIL")
		}
		binData, bitEnd, err = OverloadResponse(*value.AMFOverloadResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: " + err.Error())
		}
	case ngapType.OverloadStartIEsPresentAMFTrafficLoadReductionIndication:
		if value.AMFTrafficLoadReductionIndication == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: AMFTrafficLoadReductionIndication: NIL")
		}
		binData, bitEnd, err = TrafficLoadReductionIndication(*value.AMFTrafficLoadReductionIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: " + err.Error())
		}
	case ngapType.OverloadStartIEsPresentOverloadStartNSSAIList:
		if value.OverloadStartNSSAIList == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: OverloadStartNSSAIList: NIL")
		}
		binData, bitEnd, err = OverloadStartNSSAIList(*value.OverloadStartNSSAIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func OverloadStopIEsTypeValue(value ngapType.OverloadStopIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("OverloadStopIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkRANConfigurationTransferIEsTypeValue(value ngapType.UplinkRANConfigurationTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkRANConfigurationTransferIEsPresentSONConfigurationTransferUL:
		if value.SONConfigurationTransferUL == nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: SONConfigurationTransferUL: NIL")
		}
		binData, bitEnd, err = SONConfigurationTransfer(*value.SONConfigurationTransferUL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkRANConfigurationTransferIEsPresentENDCSONConfigurationTransferUL:
		if value.ENDCSONConfigurationTransferUL == nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: ENDCSONConfigurationTransferUL: NIL")
		}
		binData, bitEnd, err = ENDCSONConfigurationTransfer(*value.ENDCSONConfigurationTransferUL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkRANConfigurationTransferIEsTypeValue(value ngapType.DownlinkRANConfigurationTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkRANConfigurationTransferIEsPresentSONConfigurationTransferDL:
		if value.SONConfigurationTransferDL == nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: SONConfigurationTransferDL: NIL")
		}
		binData, bitEnd, err = SONConfigurationTransfer(*value.SONConfigurationTransferDL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkRANConfigurationTransferIEsPresentENDCSONConfigurationTransferDL:
		if value.ENDCSONConfigurationTransferDL == nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: ENDCSONConfigurationTransferDL: NIL")
		}
		binData, bitEnd, err = ENDCSONConfigurationTransfer(*value.ENDCSONConfigurationTransferDL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func WriteReplaceWarningRequestIEsTypeValue(value ngapType.WriteReplaceWarningRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.WriteReplaceWarningRequestIEsPresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentWarningAreaList:
		if value.WarningAreaList == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningAreaList: NIL")
		}
		binData, bitEnd, err = WarningAreaList(*value.WarningAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentRepetitionPeriod:
		if value.RepetitionPeriod == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: RepetitionPeriod: NIL")
		}
		binData, bitEnd, err = RepetitionPeriod(*value.RepetitionPeriod, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentNumberOfBroadcastsRequested:
		if value.NumberOfBroadcastsRequested == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: NumberOfBroadcastsRequested: NIL")
		}
		binData, bitEnd, err = NumberOfBroadcastsRequested(*value.NumberOfBroadcastsRequested, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentWarningType:
		if value.WarningType == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningType: NIL")
		}
		binData, bitEnd, err = WarningType(*value.WarningType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentWarningSecurityInfo:
		if value.WarningSecurityInfo == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningSecurityInfo: NIL")
		}
		binData, bitEnd, err = WarningSecurityInfo(*value.WarningSecurityInfo, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentDataCodingScheme:
		if value.DataCodingScheme == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: DataCodingScheme: NIL")
		}
		binData, bitEnd, err = DataCodingScheme(*value.DataCodingScheme, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentWarningMessageContents:
		if value.WarningMessageContents == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningMessageContents: NIL")
		}
		binData, bitEnd, err = WarningMessageContents(*value.WarningMessageContents, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentConcurrentWarningMessageInd:
		if value.ConcurrentWarningMessageInd == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: ConcurrentWarningMessageInd: NIL")
		}
		binData, bitEnd, err = ConcurrentWarningMessageInd(*value.ConcurrentWarningMessageInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsPresentWarningAreaCoordinates:
		if value.WarningAreaCoordinates == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningAreaCoordinates: NIL")
		}
		binData, bitEnd, err = WarningAreaCoordinates(*value.WarningAreaCoordinates, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func WriteReplaceWarningResponseIEsTypeValue(value ngapType.WriteReplaceWarningResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.WriteReplaceWarningResponseIEsPresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsPresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsPresentBroadcastCompletedAreaList:
		if value.BroadcastCompletedAreaList == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: BroadcastCompletedAreaList: NIL")
		}
		binData, bitEnd, err = BroadcastCompletedAreaList(*value.BroadcastCompletedAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSCancelRequestIEsTypeValue(value ngapType.PWSCancelRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSCancelRequestIEsPresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsPresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsPresentWarningAreaList:
		if value.WarningAreaList == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: WarningAreaList: NIL")
		}
		binData, bitEnd, err = WarningAreaList(*value.WarningAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsPresentCancelAllWarningMessages:
		if value.CancelAllWarningMessages == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: CancelAllWarningMessages: NIL")
		}
		binData, bitEnd, err = CancelAllWarningMessages(*value.CancelAllWarningMessages, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSCancelResponseIEsTypeValue(value ngapType.PWSCancelResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSCancelResponseIEsPresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsPresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsPresentBroadcastCancelledAreaList:
		if value.BroadcastCancelledAreaList == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: BroadcastCancelledAreaList: NIL")
		}
		binData, bitEnd, err = BroadcastCancelledAreaList(*value.BroadcastCancelledAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSRestartIndicationIEsTypeValue(value ngapType.PWSRestartIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSRestartIndicationIEsPresentCellIDListForRestart:
		if value.CellIDListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: CellIDListForRestart: NIL")
		}
		binData, bitEnd, err = CellIDListForRestart(*value.CellIDListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsPresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsPresentTAIListForRestart:
		if value.TAIListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: TAIListForRestart: NIL")
		}
		binData, bitEnd, err = TAIListForRestart(*value.TAIListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsPresentEmergencyAreaIDListForRestart:
		if value.EmergencyAreaIDListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: EmergencyAreaIDListForRestart: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDListForRestart(*value.EmergencyAreaIDListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSFailureIndicationIEsTypeValue(value ngapType.PWSFailureIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSFailureIndicationIEsPresentPWSFailedCellIDList:
		if value.PWSFailedCellIDList == nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: PWSFailedCellIDList: NIL")
		}
		binData, bitEnd, err = PWSFailedCellIDList(*value.PWSFailedCellIDList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSFailureIndicationIEsPresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkUEAssociatedNRPPaTransportIEsTypeValue(value ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkUEAssociatedNRPPaTransportIEsTypeValue(value ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue(value ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkNonUEAssociatedNRPPaTransportIEsTypeValue(value ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TraceStartIEsTypeValue(value ngapType.TraceStartIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.TraceStartIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceStartIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceStartIEsPresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("TraceStartIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TraceFailureIndicationIEsTypeValue(value ngapType.TraceFailureIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.TraceFailureIndicationIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsPresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DeactivateTraceIEsTypeValue(value ngapType.DeactivateTraceIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DeactivateTraceIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.DeactivateTraceIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.DeactivateTraceIEsPresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellTrafficTraceIEsTypeValue(value ngapType.CellTrafficTraceIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.CellTrafficTraceIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsPresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsPresentNGRANCGI:
		if value.NGRANCGI == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: NGRANCGI: NIL")
		}
		binData, bitEnd, err = NGRANCGI(*value.NGRANCGI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsPresentTraceCollectionEntityIPAddress:
		if value.TraceCollectionEntityIPAddress == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: TraceCollectionEntityIPAddress: NIL")
		}
		binData, bitEnd, err = TransportLayerAddress(*value.TraceCollectionEntityIPAddress, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReportingControlIEsTypeValue(value ngapType.LocationReportingControlIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportingControlIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingControlIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingControlIEsPresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReportingFailureIndicationIEsTypeValue(value ngapType.LocationReportingFailureIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportingFailureIndicationIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingFailureIndicationIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingFailureIndicationIEsPresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReportIEsTypeValue(value ngapType.LocationReportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsPresentUEPresenceInAreaOfInterestList:
		if value.UEPresenceInAreaOfInterestList == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: UEPresenceInAreaOfInterestList: NIL")
		}
		binData, bitEnd, err = UEPresenceInAreaOfInterestList(*value.UEPresenceInAreaOfInterestList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsPresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UETNLABindingReleaseRequestIEsTypeValue(value ngapType.UETNLABindingReleaseRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UETNLABindingReleaseRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UETNLABindingReleaseRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityInfoIndicationIEsTypeValue(value ngapType.UERadioCapabilityInfoIndicationIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UERadioCapabilityInfoIndicationIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsPresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsPresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityCheckRequestIEsTypeValue(value ngapType.UERadioCapabilityCheckRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UERadioCapabilityCheckRequestIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckRequestIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckRequestIEsPresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityCheckResponseIEsTypeValue(value ngapType.UERadioCapabilityCheckResponseIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UERadioCapabilityCheckResponseIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsPresentIMSVoiceSupportIndicator:
		if value.IMSVoiceSupportIndicator == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: IMSVoiceSupportIndicator: NIL")
		}
		binData, bitEnd, err = IMSVoiceSupportIndicator(*value.IMSVoiceSupportIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsPresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PrivateMessageIEsTypeValue(value ngapType.PrivateMessageIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	default:
		return binData, bitEnd, errors.New("PrivateMessageIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SecondaryRATDataUsageReportIEsTypeValue(value ngapType.SecondaryRATDataUsageReportIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.SecondaryRATDataUsageReportIEsPresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsPresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsPresentPDUSessionResourceSecondaryRATUsageList:
		if value.PDUSessionResourceSecondaryRATUsageList == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: PDUSessionResourceSecondaryRATUsageList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSecondaryRATUsageList(*value.PDUSessionResourceSecondaryRATUsageList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsPresentHandoverFlag:
		if value.HandoverFlag == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: HandoverFlag: NIL")
		}
		binData, bitEnd, err = HandoverFlag(*value.HandoverFlag, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsPresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkRIMInformationTransferIEsTypeValue(value ngapType.UplinkRIMInformationTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkRIMInformationTransferIEsPresentRIMInformationTransfer:
		if value.RIMInformationTransfer == nil {
			return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEsTypeValue: RIMInformationTransfer: NIL")
		}
		binData, bitEnd, err = RIMInformationTransfer(*value.RIMInformationTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkRIMInformationTransferIEsTypeValue(value ngapType.DownlinkRIMInformationTransferIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkRIMInformationTransferIEsPresentRIMInformationTransfer:
		if value.RIMInformationTransfer == nil {
			return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEsTypeValue: RIMInformationTransfer: NIL")
		}
		binData, bitEnd, err = RIMInformationTransfer(*value.RIMInformationTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEsTypeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEsTypeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}
