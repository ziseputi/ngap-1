package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func PDUSessionResourceSetupRequestIEs(value ngapType.PDUSessionResourceSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANPagingPriority {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListSUReq {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceSetupResponseIEs(value ngapType.PDUSessionResourceSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListSURes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListSURes {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListSURes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListSURes {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleaseCommandIEs(value ngapType.PDUSessionResourceReleaseCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANPagingPriority {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceToReleaseListRelCmd:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentPDUSessionResourceToReleaseListRelCmd {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleaseCommandIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceReleaseResponseIEs(value ngapType.PDUSessionResourceReleaseResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListRelRes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentPDUSessionResourceReleasedListRelRes {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleaseResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyRequestIEs(value ngapType.PDUSessionResourceModifyRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANPagingPriority {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModReq:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentPDUSessionResourceModifyListModReq {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyResponseIEs(value ngapType.PDUSessionResourceModifyResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModRes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceModifyListModRes {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModRes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceFailedToModifyListModRes {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceNotifyIEs(value ngapType.PDUSessionResourceNotifyIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceNotifyList:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceNotifyList {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListNot:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceReleasedListNot {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceNotifyIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyIndicationIEs(value ngapType.PDUSessionResourceModifyIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentPDUSessionResourceModifyListModInd {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PDUSessionResourceModifyConfirmIEs(value ngapType.PDUSessionResourceModifyConfirmIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModCfm:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceModifyListModCfm {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceFailedToModifyListModCfm {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyConfirmIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func InitialContextSetupRequestIEs(value ngapType.InitialContextSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDOldAMF:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentOldAMF {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGUAMI:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentGUAMI {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtReq:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListCxtReq {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentAllowedNSSAI {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityKey:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentSecurityKey {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceActivation:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentTraceActivation {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMobilityRestrictionList:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentMobilityRestrictionList {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapability:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapability {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIndexToRFSP:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentIndexToRFSP {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMaskedIMEISV:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentMaskedIMEISV {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDEmergencyFallbackIndicator:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentEmergencyFallbackIndicator {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapabilityForPaging {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRedirectionVoiceFallback:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentRedirectionVoiceFallback {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentLocationReportingRequestType {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentSRVCCOperationPossible {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialContextSetupRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func InitialContextSetupResponseIEs(value ngapType.InitialContextSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtRes:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListCxtRes {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtRes {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialContextSetupResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func InitialContextSetupFailureIEs(value ngapType.InitialContextSetupFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtFail {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialContextSetupFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEContextReleaseRequestIEs(value ngapType.UEContextReleaseRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelReq:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsTypeValuePresentPDUSessionResourceListCxtRelReq {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextReleaseRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEContextReleaseCommandIEs(value ngapType.UEContextReleaseCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDUENGAPIDs:
		if value.TypeValue.Present != ngapType.UEContextReleaseCommandIEsTypeValuePresentUENGAPIDs {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.UEContextReleaseCommandIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextReleaseCommandIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEContextReleaseCompleteIEs(value ngapType.UEContextReleaseCompleteIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsTypeValuePresentInfoOnRecommendedCellsAndRANNodesForPaging {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelCpl:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsTypeValuePresentPDUSessionResourceListCxtRelCpl {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextReleaseCompleteIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEContextModificationRequestIEs(value ngapType.UEContextModificationRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentRANPagingPriority {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityKey:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentSecurityKey {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIndexToRFSP:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentIndexToRFSP {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDEmergencyFallbackIndicator:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentEmergencyFallbackIndicator {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentNewAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewGUAMI:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentNewGUAMI {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentSRVCCOperationPossible {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextModificationRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEContextModificationResponseIEs(value ngapType.UEContextModificationResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCState:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsTypeValuePresentRRCState {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextModificationResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UEContextModificationFailureIEs(value ngapType.UEContextModificationFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextModificationFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RRCInactiveTransitionReportIEs(value ngapType.RRCInactiveTransitionReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCState:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRRCState {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RRCInactiveTransitionReportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverRequiredIEs(value ngapType.HandoverRequiredIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverType:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentHandoverType {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTargetID:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentTargetID {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDirectForwardingPathAvailability:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentDirectForwardingPathAvailability {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceListHORqd:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentPDUSessionResourceListHORqd {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsTypeValuePresentSourceToTargetTransparentContainer {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequiredIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverCommandIEs(value ngapType.HandoverCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverType:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentHandoverType {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASSecurityParametersFromNGRAN:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentNASSecurityParametersFromNGRAN {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceHandoverList:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceHandoverList {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceToReleaseListHOCmd {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTargetToSourceTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentTargetToSourceTransparentContainer {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCommandIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverPreparationFailureIEs(value ngapType.HandoverPreparationFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverPreparationFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverRequestIEs(value ngapType.HandoverRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverType:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentHandoverType {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityContext:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentSecurityContext {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewSecurityContextInd:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentNewSecurityContextInd {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASC:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentNASC {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListHOReq:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentPDUSessionResourceSetupListHOReq {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentAllowedNSSAI {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceActivation:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentTraceActivation {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMaskedIMEISV:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentMaskedIMEISV {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentSourceToTargetTransparentContainer {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMobilityRestrictionList:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentMobilityRestrictionList {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentLocationReportingRequestType {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGUAMI:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentGUAMI {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRedirectionVoiceFallback:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentRedirectionVoiceFallback {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentSRVCCOperationPossible {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverRequestAcknowledgeIEs(value ngapType.HandoverRequestAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceAdmittedList:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceAdmittedList {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceFailedToSetupListHOAck {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTargetToSourceTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentTargetToSourceTransparentContainer {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequestAcknowledgeIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverFailureIEs(value ngapType.HandoverFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverFailureIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("HandoverFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("HandoverFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverNotifyIEs(value ngapType.HandoverNotifyIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverNotifyIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverNotifyIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.HandoverNotifyIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("HandoverNotifyIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverNotifyIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotifyIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotifyIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverNotifyIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotifyIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestIEs(value ngapType.PathSwitchRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsTypeValuePresentSourceAMFUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsTypeValuePresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceToBeSwitchedDLList {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceFailedToSetupListPSReq {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestAcknowledgeIEs(value ngapType.PathSwitchRequestAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityContext:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSecurityContext {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewSecurityContextInd:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentNewSecurityContextInd {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSwitchedList:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceSwitchedList {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSAck:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceReleasedListPSAck {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAllowedNSSAI {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRedirectionVoiceFallback:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRedirectionVoiceFallback {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSRVCCOperationPossible {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestAcknowledgeIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PathSwitchRequestFailureIEs(value ngapType.PathSwitchRequestFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSFail:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsTypeValuePresentPDUSessionResourceReleasedListPSFail {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverCancelIEs(value ngapType.HandoverCancelIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCancelIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCancelIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverCancelIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("HandoverCancelIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverCancelIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCancelIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func HandoverCancelAcknowledgeIEs(value ngapType.HandoverCancelAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCancelAcknowledgeIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UplinkRANStatusTransferIEs(value ngapType.UplinkRANStatusTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANStatusTransferTransparentContainer:
		if value.TypeValue.Present != ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkRANStatusTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DownlinkRANStatusTransferIEs(value ngapType.DownlinkRANStatusTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANStatusTransferTransparentContainer:
		if value.TypeValue.Present != ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkRANStatusTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PagingIEs(value ngapType.PagingIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDUEPagingIdentity:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentUEPagingIdentity {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPagingDRX:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentPagingDRX {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTAIListForPaging:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentTAIListForPaging {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPagingPriority:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentPagingPriority {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentUERadioCapabilityForPaging {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPagingOrigin:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentPagingOrigin {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAssistanceDataForPaging:
		if value.TypeValue.Present != ngapType.PagingIEsTypeValuePresentAssistanceDataForPaging {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PagingIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PagingIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func InitialUEMessageIEs(value ngapType.InitialUEMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCEstablishmentCause:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentRRCEstablishmentCause {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDFiveGSTMSI:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentFiveGSTMSI {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFSetID:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentAMFSetID {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEContextRequest:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentUEContextRequest {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentAllowedNSSAI {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetAMFInformationReroute:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsTypeValuePresentSourceToTargetAMFInformationReroute {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessageIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessageIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialUEMessageIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessageIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DownlinkNASTransportIEs(value ngapType.DownlinkNASTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDOldAMF:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentOldAMF {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentRANPagingPriority {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMobilityRestrictionList:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentMobilityRestrictionList {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIndexToRFSP:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentIndexToRFSP {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentAllowedNSSAI {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsTypeValuePresentSRVCCOperationPossible {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkNASTransportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UplinkNASTransportIEs(value ngapType.UplinkNASTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkNASTransportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NASNonDeliveryIndicationIEs(value ngapType.NASNonDeliveryIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsTypeValuePresentNASPDU {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NASNonDeliveryIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RerouteNASRequestIEs(value ngapType.RerouteNASRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGAPMessage:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsTypeValuePresentNGAPMessage {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFSetID:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsTypeValuePresentAMFSetID {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsTypeValuePresentAllowedNSSAI {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetAMFInformationReroute:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsTypeValuePresentSourceToTargetAMFInformationReroute {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RerouteNASRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGSetupRequestIEs(value ngapType.NGSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsTypeValuePresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANNodeName:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsTypeValuePresentRANNodeName {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSupportedTAList:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsTypeValuePresentSupportedTAList {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDefaultPagingDRX:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsTypeValuePresentDefaultPagingDRX {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERetentionInformation:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsTypeValuePresentUERetentionInformation {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGSetupRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGSetupResponseIEs(value ngapType.NGSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFName:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsTypeValuePresentAMFName {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDServedGUAMIList:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsTypeValuePresentServedGUAMIList {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRelativeAMFCapacity:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsTypeValuePresentRelativeAMFCapacity {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPLMNSupportList:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsTypeValuePresentPLMNSupportList {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERetentionInformation:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsTypeValuePresentUERetentionInformation {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGSetupResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGSetupFailureIEs(value ngapType.NGSetupFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.NGSetupFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("NGSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTimeToWait:
		if value.TypeValue.Present != ngapType.NGSetupFailureIEsTypeValuePresentTimeToWait {
			return binData, bitEnd, errors.New("NGSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.NGSetupFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("NGSetupFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("NGSetupFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGSetupFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RANConfigurationUpdateIEs(value ngapType.RANConfigurationUpdateIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRANNodeName:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsTypeValuePresentRANNodeName {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSupportedTAList:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsTypeValuePresentSupportedTAList {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDefaultPagingDRX:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsTypeValuePresentDefaultPagingDRX {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsTypeValuePresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTNLAssociationToRemoveList:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsTypeValuePresentNGRANTNLAssociationToRemoveList {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANConfigurationUpdateIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RANConfigurationUpdateAcknowledgeIEs(value ngapType.RANConfigurationUpdateAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANConfigurationUpdateAcknowledgeIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func RANConfigurationUpdateFailureIEs(value ngapType.RANConfigurationUpdateFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTimeToWait:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentTimeToWait {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANConfigurationUpdateFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFConfigurationUpdateIEs(value ngapType.AMFConfigurationUpdateIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFName:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFName {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDServedGUAMIList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentServedGUAMIList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRelativeAMFCapacity:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentRelativeAMFCapacity {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPLMNSupportList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentPLMNSupportList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationToAddList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToAddList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationToRemoveList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToRemoveList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationToUpdateList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToUpdateList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFConfigurationUpdateIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFConfigurationUpdateAcknowledgeIEs(value ngapType.AMFConfigurationUpdateAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFTNLAssociationSetupList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationSetupList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationFailedToSetupList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationFailedToSetupList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFConfigurationUpdateAcknowledgeIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFConfigurationUpdateFailureIEs(value ngapType.AMFConfigurationUpdateFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTimeToWait:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentTimeToWait {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFConfigurationUpdateFailureIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func AMFStatusIndicationIEs(value ngapType.AMFStatusIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDUnavailableGUAMIList:
		if value.TypeValue.Present != ngapType.AMFStatusIndicationIEsTypeValuePresentUnavailableGUAMIList {
			return binData, bitEnd, errors.New("AMFStatusIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFStatusIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGResetIEs(value ngapType.NGResetIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.NGResetIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("NGResetIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDResetType:
		if value.TypeValue.Present != ngapType.NGResetIEsTypeValuePresentResetType {
			return binData, bitEnd, errors.New("NGResetIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("NGResetIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGResetIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func NGResetAcknowledgeIEs(value ngapType.NGResetAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDUEAssociatedLogicalNGConnectionList:
		if value.TypeValue.Present != ngapType.NGResetAcknowledgeIEsTypeValuePresentUEAssociatedLogicalNGConnectionList {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.NGResetAcknowledgeIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGResetAcknowledgeIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func ErrorIndicationIEs(value ngapType.ErrorIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ErrorIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func OverloadStartIEs(value ngapType.OverloadStartIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFOverloadResponse:
		if value.TypeValue.Present != ngapType.OverloadStartIEsTypeValuePresentAMFOverloadResponse {
			return binData, bitEnd, errors.New("OverloadStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTrafficLoadReductionIndication:
		if value.TypeValue.Present != ngapType.OverloadStartIEsTypeValuePresentAMFTrafficLoadReductionIndication {
			return binData, bitEnd, errors.New("OverloadStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDOverloadStartNSSAIList:
		if value.TypeValue.Present != ngapType.OverloadStartIEsTypeValuePresentOverloadStartNSSAIList {
			return binData, bitEnd, errors.New("OverloadStartIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("OverloadStartIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = OverloadStartIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func OverloadStopIEs(value ngapType.OverloadStopIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	default:
		return binData, bitEnd, errors.New("OverloadStopIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStopIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStopIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = OverloadStopIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStopIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UplinkRANConfigurationTransferIEs(value ngapType.UplinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDSONConfigurationTransferUL:
		if value.TypeValue.Present != ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferUL {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDENDCSONConfigurationTransferUL:
		if value.TypeValue.Present != ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferUL {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkRANConfigurationTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DownlinkRANConfigurationTransferIEs(value ngapType.DownlinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDSONConfigurationTransferDL:
		if value.TypeValue.Present != ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferDL {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDENDCSONConfigurationTransferDL:
		if value.TypeValue.Present != ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferDL {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkRANConfigurationTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func WriteReplaceWarningRequestIEs(value ngapType.WriteReplaceWarningRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDMessageIdentifier:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentMessageIdentifier {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentSerialNumber {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningAreaList:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaList {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRepetitionPeriod:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentRepetitionPeriod {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNumberOfBroadcastsRequested:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentNumberOfBroadcastsRequested {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningType:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningType {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningSecurityInfo:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningSecurityInfo {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDataCodingScheme:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentDataCodingScheme {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningMessageContents:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningMessageContents {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDConcurrentWarningMessageInd:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentConcurrentWarningMessageInd {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningAreaCoordinates:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaCoordinates {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = WriteReplaceWarningRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func WriteReplaceWarningResponseIEs(value ngapType.WriteReplaceWarningResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDMessageIdentifier:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsTypeValuePresentMessageIdentifier {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsTypeValuePresentSerialNumber {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDBroadcastCompletedAreaList:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsTypeValuePresentBroadcastCompletedAreaList {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = WriteReplaceWarningResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PWSCancelRequestIEs(value ngapType.PWSCancelRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDMessageIdentifier:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsTypeValuePresentMessageIdentifier {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsTypeValuePresentSerialNumber {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningAreaList:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsTypeValuePresentWarningAreaList {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCancelAllWarningMessages:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsTypeValuePresentCancelAllWarningMessages {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSCancelRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PWSCancelResponseIEs(value ngapType.PWSCancelResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDMessageIdentifier:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsTypeValuePresentMessageIdentifier {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsTypeValuePresentSerialNumber {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDBroadcastCancelledAreaList:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsTypeValuePresentBroadcastCancelledAreaList {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSCancelResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PWSRestartIndicationIEs(value ngapType.PWSRestartIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDCellIDListForRestart:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsTypeValuePresentCellIDListForRestart {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsTypeValuePresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTAIListForRestart:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsTypeValuePresentTAIListForRestart {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDEmergencyAreaIDListForRestart:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsTypeValuePresentEmergencyAreaIDListForRestart {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSRestartIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PWSFailureIndicationIEs(value ngapType.PWSFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDPWSFailedCellIDList:
		if value.TypeValue.Present != ngapType.PWSFailureIndicationIEsTypeValuePresentPWSFailedCellIDList {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.PWSFailureIndicationIEsTypeValuePresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSFailureIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DownlinkUEAssociatedNRPPaTransportIEs(value ngapType.DownlinkUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRoutingID:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkUEAssociatedNRPPaTransportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UplinkUEAssociatedNRPPaTransportIEs(value ngapType.UplinkUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRoutingID:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkUEAssociatedNRPPaTransportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DownlinkNonUEAssociatedNRPPaTransportIEs(value ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRoutingID:
		if value.TypeValue.Present != ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkNonUEAssociatedNRPPaTransportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UplinkNonUEAssociatedNRPPaTransportIEs(value ngapType.UplinkNonUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRoutingID:
		if value.TypeValue.Present != ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkNonUEAssociatedNRPPaTransportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TraceStartIEs(value ngapType.TraceStartIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.TraceStartIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("TraceStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.TraceStartIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("TraceStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceActivation:
		if value.TypeValue.Present != ngapType.TraceStartIEsTypeValuePresentTraceActivation {
			return binData, bitEnd, errors.New("TraceStartIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("TraceStartIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStartIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStartIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TraceStartIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStartIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func TraceFailureIndicationIEs(value ngapType.TraceFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTraceID:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsTypeValuePresentNGRANTraceID {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TraceFailureIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DeactivateTraceIEs(value ngapType.DeactivateTraceIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.DeactivateTraceIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DeactivateTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DeactivateTraceIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("DeactivateTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTraceID:
		if value.TypeValue.Present != ngapType.DeactivateTraceIEsTypeValuePresentNGRANTraceID {
			return binData, bitEnd, errors.New("DeactivateTraceIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DeactivateTraceIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTraceIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTraceIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DeactivateTraceIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTraceIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func CellTrafficTraceIEs(value ngapType.CellTrafficTraceIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTraceID:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsTypeValuePresentNGRANTraceID {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANCGI:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsTypeValuePresentNGRANCGI {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceCollectionEntityIPAddress:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsTypeValuePresentTraceCollectionEntityIPAddress {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellTrafficTraceIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LocationReportingControlIEs(value ngapType.LocationReportingControlIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportingControlIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingControlIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportingControlIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingControlIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.LocationReportingControlIEsTypeValuePresentLocationReportingRequestType {
			return binData, bitEnd, errors.New("LocationReportingControlIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingControlIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControlIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControlIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportingControlIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControlIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LocationReportingFailureIndicationIEs(value ngapType.LocationReportingFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportingFailureIndicationIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportingFailureIndicationIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.LocationReportingFailureIndicationIEsTypeValuePresentCause {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportingFailureIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func LocationReportIEs(value ngapType.LocationReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.LocationReportIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEPresenceInAreaOfInterestList:
		if value.TypeValue.Present != ngapType.LocationReportIEsTypeValuePresentUEPresenceInAreaOfInterestList {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.LocationReportIEsTypeValuePresentLocationReportingRequestType {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UETNLABindingReleaseRequestIEs(value ngapType.UETNLABindingReleaseRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UETNLABindingReleaseRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UERadioCapabilityInfoIndicationIEs(value ngapType.UERadioCapabilityInfoIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapability:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapability {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapabilityForPaging {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityInfoIndicationIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UERadioCapabilityCheckRequestIEs(value ngapType.UERadioCapabilityCheckRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapability:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentUERadioCapability {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityCheckRequestIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UERadioCapabilityCheckResponseIEs(value ngapType.UERadioCapabilityCheckResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIMSVoiceSupportIndicator:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentIMSVoiceSupportIndicator {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityCheckResponseIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func PrivateMessageIEs(value ngapType.PrivateMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	// switch value.PrivateIEID.Value {
	// default:
		// return binData, bitEnd, errors.New("PrivateMessageIEs: PrivateIEID: INVALID")
	// }
	binData, bitEnd, err = PrivateIEID(value.PrivateIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PrivateMessageIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PrivateMessageIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PrivateMessageIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PrivateMessageIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SecondaryRATDataUsageReportIEs(value ngapType.SecondaryRATDataUsageReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentAMFUENGAPID {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentRANUENGAPID {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSecondaryRATUsageList:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentPDUSessionResourceSecondaryRATUsageList {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverFlag:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentHandoverFlag {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentUserLocationInformation {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SecondaryRATDataUsageReportIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UplinkRIMInformationTransferIEs(value ngapType.UplinkRIMInformationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRIMInformationTransfer:
		if value.TypeValue.Present != ngapType.UplinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer {
			return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkRIMInformationTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRIMInformationTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func DownlinkRIMInformationTransferIEs(value ngapType.DownlinkRIMInformationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDRIMInformationTransfer:
		if value.TypeValue.Present != ngapType.DownlinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer {
			return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkRIMInformationTransferIEsTypeValue(value.TypeValue, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRIMInformationTransferIEs: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}
