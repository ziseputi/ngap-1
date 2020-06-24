package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func PDUSessionResourceSetupRequestIEsTypeValue(value ngapType.PDUSessionResourceSetupRequestIEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListSUReq:
		if value.PDUSessionResourceSetupListSUReq == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: PDUSessionResourceSetupListSUReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListSUReq(*value.PDUSessionResourceSetupListSUReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
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
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListSURes:
		if value.PDUSessionResourceSetupListSURes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: PDUSessionResourceSetupListSURes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListSURes(*value.PDUSessionResourceSetupListSURes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListSURes:
		if value.PDUSessionResourceFailedToSetupListSURes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: PDUSessionResourceFailedToSetupListSURes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListSURes(*value.PDUSessionResourceFailedToSetupListSURes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentPDUSessionResourceToReleaseListRelCmd:
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
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentPDUSessionResourceReleasedListRelRes:
		if value.PDUSessionResourceReleasedListRelRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: PDUSessionResourceReleasedListRelRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListRelRes(*value.PDUSessionResourceReleasedListRelRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentPDUSessionResourceModifyListModReq:
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
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceModifyListModRes:
		if value.PDUSessionResourceModifyListModRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: PDUSessionResourceModifyListModRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModRes(*value.PDUSessionResourceModifyListModRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceFailedToModifyListModRes:
		if value.PDUSessionResourceFailedToModifyListModRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: PDUSessionResourceFailedToModifyListModRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToModifyListModRes(*value.PDUSessionResourceFailedToModifyListModRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceNotifyList:
		if value.PDUSessionResourceNotifyList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: PDUSessionResourceNotifyList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceNotifyList(*value.PDUSessionResourceNotifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceReleasedListNot:
		if value.PDUSessionResourceReleasedListNot == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: PDUSessionResourceReleasedListNot: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListNot(*value.PDUSessionResourceReleasedListNot, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentUserLocationInformation:
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
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentPDUSessionResourceModifyListModInd:
		if value.PDUSessionResourceModifyListModInd == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: PDUSessionResourceModifyListModInd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModInd(*value.PDUSessionResourceModifyListModInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentUserLocationInformation:
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
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceModifyListModCfm:
		if value.PDUSessionResourceModifyListModCfm == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: PDUSessionResourceModifyListModCfm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModCfm(*value.PDUSessionResourceModifyListModCfm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceFailedToModifyListModCfm:
		if value.PDUSessionResourceFailedToModifyListModCfm == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: PDUSessionResourceFailedToModifyListModCfm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToModifyListModCfm(*value.PDUSessionResourceFailedToModifyListModCfm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsTypeValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentOldAMF:
		if value.OldAMF == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: OldAMF: NIL")
		}
		binData, bitEnd, err = AMFName(*value.OldAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentGUAMI:
		if value.GUAMI == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: GUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.GUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListCxtReq:
		if value.PDUSessionResourceSetupListCxtReq == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: PDUSessionResourceSetupListCxtReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListCxtReq(*value.PDUSessionResourceSetupListCxtReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentSecurityKey:
		if value.SecurityKey == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: SecurityKey: NIL")
		}
		binData, bitEnd, err = SecurityKey(*value.SecurityKey, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentMaskedIMEISV:
		if value.MaskedIMEISV == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: MaskedIMEISV: NIL")
		}
		binData, bitEnd, err = MaskedIMEISV(*value.MaskedIMEISV, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentEmergencyFallbackIndicator:
		if value.EmergencyFallbackIndicator == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: EmergencyFallbackIndicator: NIL")
		}
		binData, bitEnd, err = EmergencyFallbackIndicator(*value.EmergencyFallbackIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentRedirectionVoiceFallback:
		if value.RedirectionVoiceFallback == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: RedirectionVoiceFallback: NIL")
		}
		binData, bitEnd, err = RedirectionVoiceFallback(*value.RedirectionVoiceFallback, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentSRVCCOperationPossible:
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
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListCxtRes:
		if value.PDUSessionResourceSetupListCxtRes == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: PDUSessionResourceSetupListCxtRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListCxtRes(*value.PDUSessionResourceSetupListCxtRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtRes:
		if value.PDUSessionResourceFailedToSetupListCxtRes == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: PDUSessionResourceFailedToSetupListCxtRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListCxtRes(*value.PDUSessionResourceFailedToSetupListCxtRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtFail:
		if value.PDUSessionResourceFailedToSetupListCxtFail == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: PDUSessionResourceFailedToSetupListCxtFail: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListCxtFail(*value.PDUSessionResourceFailedToSetupListCxtFail, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentPDUSessionResourceListCxtRelReq:
		if value.PDUSessionResourceListCxtRelReq == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: PDUSessionResourceListCxtRelReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListCxtRelReq(*value.PDUSessionResourceListCxtRelReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentCause:
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
	case ngapType.UEContextReleaseCommandIEsTypeValuePresentUENGAPIDs:
		if value.UENGAPIDs == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: UENGAPIDs: NIL")
		}
		binData, bitEnd, err = UENGAPIDs(*value.UENGAPIDs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCommandIEsTypeValuePresentCause:
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
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentInfoOnRecommendedCellsAndRANNodesForPaging:
		if value.InfoOnRecommendedCellsAndRANNodesForPaging == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: InfoOnRecommendedCellsAndRANNodesForPaging: NIL")
		}
		binData, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPaging(*value.InfoOnRecommendedCellsAndRANNodesForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentPDUSessionResourceListCxtRelCpl:
		if value.PDUSessionResourceListCxtRelCpl == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: PDUSessionResourceListCxtRelCpl: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListCxtRelCpl(*value.PDUSessionResourceListCxtRelCpl, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.UEContextModificationRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentSecurityKey:
		if value.SecurityKey == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: SecurityKey: NIL")
		}
		binData, bitEnd, err = SecurityKey(*value.SecurityKey, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentEmergencyFallbackIndicator:
		if value.EmergencyFallbackIndicator == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: EmergencyFallbackIndicator: NIL")
		}
		binData, bitEnd, err = EmergencyFallbackIndicator(*value.EmergencyFallbackIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentNewAMFUENGAPID:
		if value.NewAMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: NewAMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.NewAMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentNewGUAMI:
		if value.NewGUAMI == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: NewGUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.NewGUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentSRVCCOperationPossible:
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
	case ngapType.UEContextModificationResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentRRCState:
		if value.RRCState == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: RRCState: NIL")
		}
		binData, bitEnd, err = RRCState(*value.RRCState, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.UEContextModificationFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRRCState:
		if value.RRCState == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: RRCState: NIL")
		}
		binData, bitEnd, err = RRCState(*value.RRCState, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsTypeValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentUserLocationInformation:
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
	case ngapType.HandoverRequiredIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentTargetID:
		if value.TargetID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: TargetID: NIL")
		}
		binData, bitEnd, err = TargetID(*value.TargetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentDirectForwardingPathAvailability:
		if value.DirectForwardingPathAvailability == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: DirectForwardingPathAvailability: NIL")
		}
		binData, bitEnd, err = DirectForwardingPathAvailability(*value.DirectForwardingPathAvailability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentPDUSessionResourceListHORqd:
		if value.PDUSessionResourceListHORqd == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: PDUSessionResourceListHORqd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListHORqd(*value.PDUSessionResourceListHORqd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentSourceToTargetTransparentContainer:
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
	case ngapType.HandoverCommandIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentNASSecurityParametersFromNGRAN:
		if value.NASSecurityParametersFromNGRAN == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: NASSecurityParametersFromNGRAN: NIL")
		}
		binData, bitEnd, err = NASSecurityParametersFromNGRAN(*value.NASSecurityParametersFromNGRAN, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceHandoverList:
		if value.PDUSessionResourceHandoverList == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: PDUSessionResourceHandoverList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceHandoverList(*value.PDUSessionResourceHandoverList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceToReleaseListHOCmd:
		if value.PDUSessionResourceToReleaseListHOCmd == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: PDUSessionResourceToReleaseListHOCmd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToReleaseListHOCmd(*value.PDUSessionResourceToReleaseListHOCmd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentTargetToSourceTransparentContainer:
		if value.TargetToSourceTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: TargetToSourceTransparentContainer: NIL")
		}
		binData, bitEnd, err = TargetToSourceTransparentContainer(*value.TargetToSourceTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.HandoverRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentSecurityContext:
		if value.SecurityContext == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: SecurityContext: NIL")
		}
		binData, bitEnd, err = SecurityContext(*value.SecurityContext, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentNewSecurityContextInd:
		if value.NewSecurityContextInd == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: NewSecurityContextInd: NIL")
		}
		binData, bitEnd, err = NewSecurityContextInd(*value.NewSecurityContextInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentNASC:
		if value.NASC == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: NASC: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASC, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentPDUSessionResourceSetupListHOReq:
		if value.PDUSessionResourceSetupListHOReq == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: PDUSessionResourceSetupListHOReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListHOReq(*value.PDUSessionResourceSetupListHOReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentMaskedIMEISV:
		if value.MaskedIMEISV == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: MaskedIMEISV: NIL")
		}
		binData, bitEnd, err = MaskedIMEISV(*value.MaskedIMEISV, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentSourceToTargetTransparentContainer:
		if value.SourceToTargetTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: SourceToTargetTransparentContainer: NIL")
		}
		binData, bitEnd, err = SourceToTargetTransparentContainer(*value.SourceToTargetTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentGUAMI:
		if value.GUAMI == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: GUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.GUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentRedirectionVoiceFallback:
		if value.RedirectionVoiceFallback == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: RedirectionVoiceFallback: NIL")
		}
		binData, bitEnd, err = RedirectionVoiceFallback(*value.RedirectionVoiceFallback, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentSRVCCOperationPossible:
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
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceAdmittedList:
		if value.PDUSessionResourceAdmittedList == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: PDUSessionResourceAdmittedList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceAdmittedList(*value.PDUSessionResourceAdmittedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceFailedToSetupListHOAck:
		if value.PDUSessionResourceFailedToSetupListHOAck == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: PDUSessionResourceFailedToSetupListHOAck: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListHOAck(*value.PDUSessionResourceFailedToSetupListHOAck, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentTargetToSourceTransparentContainer:
		if value.TargetToSourceTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: TargetToSourceTransparentContainer: NIL")
		}
		binData, bitEnd, err = TargetToSourceTransparentContainer(*value.TargetToSourceTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.HandoverFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.HandoverNotifyIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverNotifyIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverNotifyIEsTypeValuePresentUserLocationInformation:
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
	case ngapType.PathSwitchRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentSourceAMFUENGAPID:
		if value.SourceAMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: SourceAMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.SourceAMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceToBeSwitchedDLList:
		if value.PDUSessionResourceToBeSwitchedDLList == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: PDUSessionResourceToBeSwitchedDLList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLList(*value.PDUSessionResourceToBeSwitchedDLList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceFailedToSetupListPSReq:
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
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSecurityContext:
		if value.SecurityContext == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: SecurityContext: NIL")
		}
		binData, bitEnd, err = SecurityContext(*value.SecurityContext, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentNewSecurityContextInd:
		if value.NewSecurityContextInd == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: NewSecurityContextInd: NIL")
		}
		binData, bitEnd, err = NewSecurityContextInd(*value.NewSecurityContextInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceSwitchedList:
		if value.PDUSessionResourceSwitchedList == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: PDUSessionResourceSwitchedList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSwitchedList(*value.PDUSessionResourceSwitchedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceReleasedListPSAck:
		if value.PDUSessionResourceReleasedListPSAck == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: PDUSessionResourceReleasedListPSAck: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListPSAck(*value.PDUSessionResourceReleasedListPSAck, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive:
		if value.CoreNetworkAssistanceInformationForInactive == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: CoreNetworkAssistanceInformationForInactive: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformationForInactive(*value.CoreNetworkAssistanceInformationForInactive, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRedirectionVoiceFallback:
		if value.RedirectionVoiceFallback == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: RedirectionVoiceFallback: NIL")
		}
		binData, bitEnd, err = RedirectionVoiceFallback(*value.RedirectionVoiceFallback, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCNAssistedRANTuning:
		if value.CNAssistedRANTuning == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: CNAssistedRANTuning: NIL")
		}
		binData, bitEnd, err = CNAssistedRANTuning(*value.CNAssistedRANTuning, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSRVCCOperationPossible:
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
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentPDUSessionResourceReleasedListPSFail:
		if value.PDUSessionResourceReleasedListPSFail == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: PDUSessionResourceReleasedListPSFail: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListPSFail(*value.PDUSessionResourceReleasedListPSFail, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.HandoverCancelIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelIEsTypeValuePresentCause:
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
	case ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.UplinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer:
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
	case ngapType.DownlinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer:
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
	case ngapType.PagingIEsTypeValuePresentUEPagingIdentity:
		if value.UEPagingIdentity == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: UEPagingIdentity: NIL")
		}
		binData, bitEnd, err = UEPagingIdentity(*value.UEPagingIdentity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentPagingDRX:
		if value.PagingDRX == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: PagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.PagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentTAIListForPaging:
		if value.TAIListForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: TAIListForPaging: NIL")
		}
		binData, bitEnd, err = TAIListForPaging(*value.TAIListForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentPagingPriority:
		if value.PagingPriority == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: PagingPriority: NIL")
		}
		binData, bitEnd, err = PagingPriority(*value.PagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentPagingOrigin:
		if value.PagingOrigin == nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: PagingOrigin: NIL")
		}
		binData, bitEnd, err = PagingOrigin(*value.PagingOrigin, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsTypeValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentAssistanceDataForPaging:
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
	case ngapType.InitialUEMessageIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentRRCEstablishmentCause:
		if value.RRCEstablishmentCause == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: RRCEstablishmentCause: NIL")
		}
		binData, bitEnd, err = RRCEstablishmentCause(*value.RRCEstablishmentCause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentFiveGSTMSI:
		if value.FiveGSTMSI == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: FiveGSTMSI: NIL")
		}
		binData, bitEnd, err = FiveGSTMSI(*value.FiveGSTMSI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentAMFSetID:
		if value.AMFSetID == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: AMFSetID: NIL")
		}
		binData, bitEnd, err = AMFSetID(*value.AMFSetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentUEContextRequest:
		if value.UEContextRequest == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: UEContextRequest: NIL")
		}
		binData, bitEnd, err = UEContextRequest(*value.UEContextRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsTypeValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentSourceToTargetAMFInformationReroute:
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
	case ngapType.DownlinkNASTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentOldAMF:
		if value.OldAMF == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: OldAMF: NIL")
		}
		binData, bitEnd, err = AMFName(*value.OldAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentSRVCCOperationPossible:
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
	case ngapType.UplinkNASTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsTypeValuePresentUserLocationInformation:
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
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentCause:
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
	case ngapType.RerouteNASRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentNGAPMessage:
		if value.NGAPMessage == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: NGAPMessage: NIL")
		}
		binData, bitEnd, err = NGAPMessage(*value.NGAPMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentAMFSetID:
		if value.AMFSetID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: AMFSetID: NIL")
		}
		binData, bitEnd, err = AMFSetID(*value.AMFSetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentSourceToTargetAMFInformationReroute:
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
	case ngapType.NGSetupRequestIEsTypeValuePresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentRANNodeName:
		if value.RANNodeName == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: RANNodeName: NIL")
		}
		binData, bitEnd, err = RANNodeName(*value.RANNodeName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentSupportedTAList:
		if value.SupportedTAList == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: SupportedTAList: NIL")
		}
		binData, bitEnd, err = SupportedTAList(*value.SupportedTAList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentDefaultPagingDRX:
		if value.DefaultPagingDRX == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: DefaultPagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.DefaultPagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentUERetentionInformation:
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
	case ngapType.NGSetupResponseIEsTypeValuePresentAMFName:
		if value.AMFName == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: AMFName: NIL")
		}
		binData, bitEnd, err = AMFName(*value.AMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentServedGUAMIList:
		if value.ServedGUAMIList == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: ServedGUAMIList: NIL")
		}
		binData, bitEnd, err = ServedGUAMIList(*value.ServedGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentRelativeAMFCapacity:
		if value.RelativeAMFCapacity == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: RelativeAMFCapacity: NIL")
		}
		binData, bitEnd, err = RelativeAMFCapacity(*value.RelativeAMFCapacity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentPLMNSupportList:
		if value.PLMNSupportList == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: PLMNSupportList: NIL")
		}
		binData, bitEnd, err = PLMNSupportList(*value.PLMNSupportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentUERetentionInformation:
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
	case ngapType.NGSetupFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupFailureIEsTypeValuePresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.NGSetupFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentRANNodeName:
		if value.RANNodeName == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: RANNodeName: NIL")
		}
		binData, bitEnd, err = RANNodeName(*value.RANNodeName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentSupportedTAList:
		if value.SupportedTAList == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: SupportedTAList: NIL")
		}
		binData, bitEnd, err = SupportedTAList(*value.SupportedTAList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentDefaultPagingDRX:
		if value.DefaultPagingDRX == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: DefaultPagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.DefaultPagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentNGRANTNLAssociationToRemoveList:
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
	case ngapType.RANConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFName:
		if value.AMFName == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFName: NIL")
		}
		binData, bitEnd, err = AMFName(*value.AMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentServedGUAMIList:
		if value.ServedGUAMIList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: ServedGUAMIList: NIL")
		}
		binData, bitEnd, err = ServedGUAMIList(*value.ServedGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentRelativeAMFCapacity:
		if value.RelativeAMFCapacity == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: RelativeAMFCapacity: NIL")
		}
		binData, bitEnd, err = RelativeAMFCapacity(*value.RelativeAMFCapacity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentPLMNSupportList:
		if value.PLMNSupportList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: PLMNSupportList: NIL")
		}
		binData, bitEnd, err = PLMNSupportList(*value.PLMNSupportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToAddList:
		if value.AMFTNLAssociationToAddList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFTNLAssociationToAddList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToAddList(*value.AMFTNLAssociationToAddList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToRemoveList:
		if value.AMFTNLAssociationToRemoveList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: AMFTNLAssociationToRemoveList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToRemoveList(*value.AMFTNLAssociationToRemoveList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToUpdateList:
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
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationSetupList:
		if value.AMFTNLAssociationSetupList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: AMFTNLAssociationSetupList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationSetupList(*value.AMFTNLAssociationSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationFailedToSetupList:
		if value.AMFTNLAssociationFailedToSetupList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: AMFTNLAssociationFailedToSetupList: NIL")
		}
		binData, bitEnd, err = TNLAssociationList(*value.AMFTNLAssociationFailedToSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsTypeValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.AMFStatusIndicationIEsTypeValuePresentUnavailableGUAMIList:
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
	case ngapType.NGResetIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NGResetIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetIEsTypeValue: " + err.Error())
		}
	case ngapType.NGResetIEsTypeValuePresentResetType:
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
	case ngapType.NGResetAcknowledgeIEsTypeValuePresentUEAssociatedLogicalNGConnectionList:
		if value.UEAssociatedLogicalNGConnectionList == nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: UEAssociatedLogicalNGConnectionList: NIL")
		}
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionList(*value.UEAssociatedLogicalNGConnectionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsTypeValue: " + err.Error())
		}
	case ngapType.NGResetAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.ErrorIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.OverloadStartIEsTypeValuePresentAMFOverloadResponse:
		if value.AMFOverloadResponse == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: AMFOverloadResponse: NIL")
		}
		binData, bitEnd, err = OverloadResponse(*value.AMFOverloadResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: " + err.Error())
		}
	case ngapType.OverloadStartIEsTypeValuePresentAMFTrafficLoadReductionIndication:
		if value.AMFTrafficLoadReductionIndication == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: AMFTrafficLoadReductionIndication: NIL")
		}
		binData, bitEnd, err = TrafficLoadReductionIndication(*value.AMFTrafficLoadReductionIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsTypeValue: " + err.Error())
		}
	case ngapType.OverloadStartIEsTypeValuePresentOverloadStartNSSAIList:
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
	case ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferUL:
		if value.SONConfigurationTransferUL == nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: SONConfigurationTransferUL: NIL")
		}
		binData, bitEnd, err = SONConfigurationTransfer(*value.SONConfigurationTransferUL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferUL:
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
	case ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferDL:
		if value.SONConfigurationTransferDL == nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: SONConfigurationTransferDL: NIL")
		}
		binData, bitEnd, err = SONConfigurationTransfer(*value.SONConfigurationTransferDL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferDL:
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
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaList:
		if value.WarningAreaList == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningAreaList: NIL")
		}
		binData, bitEnd, err = WarningAreaList(*value.WarningAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentRepetitionPeriod:
		if value.RepetitionPeriod == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: RepetitionPeriod: NIL")
		}
		binData, bitEnd, err = RepetitionPeriod(*value.RepetitionPeriod, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentNumberOfBroadcastsRequested:
		if value.NumberOfBroadcastsRequested == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: NumberOfBroadcastsRequested: NIL")
		}
		binData, bitEnd, err = NumberOfBroadcastsRequested(*value.NumberOfBroadcastsRequested, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningType:
		if value.WarningType == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningType: NIL")
		}
		binData, bitEnd, err = WarningType(*value.WarningType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningSecurityInfo:
		if value.WarningSecurityInfo == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningSecurityInfo: NIL")
		}
		binData, bitEnd, err = WarningSecurityInfo(*value.WarningSecurityInfo, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentDataCodingScheme:
		if value.DataCodingScheme == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: DataCodingScheme: NIL")
		}
		binData, bitEnd, err = DataCodingScheme(*value.DataCodingScheme, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningMessageContents:
		if value.WarningMessageContents == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: WarningMessageContents: NIL")
		}
		binData, bitEnd, err = WarningMessageContents(*value.WarningMessageContents, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentConcurrentWarningMessageInd:
		if value.ConcurrentWarningMessageInd == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: ConcurrentWarningMessageInd: NIL")
		}
		binData, bitEnd, err = ConcurrentWarningMessageInd(*value.ConcurrentWarningMessageInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaCoordinates:
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
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentBroadcastCompletedAreaList:
		if value.BroadcastCompletedAreaList == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: BroadcastCompletedAreaList: NIL")
		}
		binData, bitEnd, err = BroadcastCompletedAreaList(*value.BroadcastCompletedAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.PWSCancelRequestIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsTypeValuePresentWarningAreaList:
		if value.WarningAreaList == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: WarningAreaList: NIL")
		}
		binData, bitEnd, err = WarningAreaList(*value.WarningAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsTypeValuePresentCancelAllWarningMessages:
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
	case ngapType.PWSCancelResponseIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsTypeValuePresentBroadcastCancelledAreaList:
		if value.BroadcastCancelledAreaList == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: BroadcastCancelledAreaList: NIL")
		}
		binData, bitEnd, err = BroadcastCancelledAreaList(*value.BroadcastCancelledAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.PWSRestartIndicationIEsTypeValuePresentCellIDListForRestart:
		if value.CellIDListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: CellIDListForRestart: NIL")
		}
		binData, bitEnd, err = CellIDListForRestart(*value.CellIDListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsTypeValuePresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsTypeValuePresentTAIListForRestart:
		if value.TAIListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: TAIListForRestart: NIL")
		}
		binData, bitEnd, err = TAIListForRestart(*value.TAIListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsTypeValuePresentEmergencyAreaIDListForRestart:
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
	case ngapType.PWSFailureIndicationIEsTypeValuePresentPWSFailedCellIDList:
		if value.PWSFailedCellIDList == nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: PWSFailedCellIDList: NIL")
		}
		binData, bitEnd, err = PWSFailedCellIDList(*value.PWSFailedCellIDList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.PWSFailureIndicationIEsTypeValuePresentGlobalRANNodeID:
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
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
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
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
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
	case ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
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
	case ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsTypeValue: " + err.Error())
		}
	case ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
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
	case ngapType.TraceStartIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceStartIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceStartIEsTypeValuePresentTraceActivation:
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
	case ngapType.TraceFailureIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsTypeValuePresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsTypeValuePresentCause:
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
	case ngapType.DeactivateTraceIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.DeactivateTraceIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.DeactivateTraceIEsTypeValuePresentNGRANTraceID:
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
	case ngapType.CellTrafficTraceIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentNGRANCGI:
		if value.NGRANCGI == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: NGRANCGI: NIL")
		}
		binData, bitEnd, err = NGRANCGI(*value.NGRANCGI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsTypeValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentTraceCollectionEntityIPAddress:
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
	case ngapType.LocationReportingControlIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingControlIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingControlIEsTypeValuePresentLocationReportingRequestType:
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
	case ngapType.LocationReportingFailureIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingFailureIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportingFailureIndicationIEsTypeValuePresentCause:
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
	case ngapType.LocationReportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentUEPresenceInAreaOfInterestList:
		if value.UEPresenceInAreaOfInterestList == nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: UEPresenceInAreaOfInterestList: NIL")
		}
		binData, bitEnd, err = UEPresenceInAreaOfInterestList(*value.UEPresenceInAreaOfInterestList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsTypeValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentLocationReportingRequestType:
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
	case ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentRANUENGAPID:
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
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapabilityForPaging:
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
	case ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentUERadioCapability:
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
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentIMSVoiceSupportIndicator:
		if value.IMSVoiceSupportIndicator == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: IMSVoiceSupportIndicator: NIL")
		}
		binData, bitEnd, err = IMSVoiceSupportIndicator(*value.IMSVoiceSupportIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsTypeValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentCriticalityDiagnostics:
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
	case ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentPDUSessionResourceSecondaryRATUsageList:
		if value.PDUSessionResourceSecondaryRATUsageList == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: PDUSessionResourceSecondaryRATUsageList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSecondaryRATUsageList(*value.PDUSessionResourceSecondaryRATUsageList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentHandoverFlag:
		if value.HandoverFlag == nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: HandoverFlag: NIL")
		}
		binData, bitEnd, err = HandoverFlag(*value.HandoverFlag, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEsTypeValue: " + err.Error())
		}
	case ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentUserLocationInformation:
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
	case ngapType.UplinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer:
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
	case ngapType.DownlinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer:
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
