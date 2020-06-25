package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func PDUSessionResourceSetupRequestIEs(value ngapType.PDUSessionResourceSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsPresentRANPagingPriority {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsPresentNASPDU {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsPresentPDUSessionResourceSetupListSUReq {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupRequestIEsPresentUEAggregateMaximumBitRate {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListSURes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceSetupListSURes {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListSURes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceFailedToSetupListSURes {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceSetupResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsPresentRANPagingPriority {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsPresentNASPDU {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceToReleaseListRelCmd:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseCommandIEsPresentPDUSessionResourceToReleaseListRelCmd {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListRelRes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsPresentPDUSessionResourceReleasedListRelRes {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceReleaseResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsPresentRANPagingPriority {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModReq:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyRequestIEsPresentPDUSessionResourceModifyListModReq {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModRes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceModifyListModRes {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModRes:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceFailedToModifyListModRes {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceNotifyList:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsPresentPDUSessionResourceNotifyList {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListNot:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsPresentPDUSessionResourceReleasedListNot {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceNotifyIEsPresentUserLocationInformation {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsPresentPDUSessionResourceModifyListModInd {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyIndicationIEsPresentUserLocationInformation {
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
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceModifyListModCfm:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceModifyListModCfm {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceFailedToModifyListModCfm {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PDUSessionResourceModifyConfirmIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDOldAMF:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentOldAMF {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGUAMI:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentGUAMI {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtReq:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentPDUSessionResourceSetupListCxtReq {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentAllowedNSSAI {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityKey:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentSecurityKey {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceActivation:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentTraceActivation {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMobilityRestrictionList:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentMobilityRestrictionList {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapability:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentUERadioCapability {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIndexToRFSP:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentIndexToRFSP {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMaskedIMEISV:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentMaskedIMEISV {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentNASPDU {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDEmergencyFallbackIndicator:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentEmergencyFallbackIndicator {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentUERadioCapabilityForPaging {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRedirectionVoiceFallback:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentRedirectionVoiceFallback {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentLocationReportingRequestType {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsPresentSRVCCOperationPossible {
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
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtRes:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsPresentPDUSessionResourceSetupListCxtRes {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsPresentPDUSessionResourceFailedToSetupListCxtRes {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.InitialContextSetupResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsPresentPDUSessionResourceFailedToSetupListCxtFail {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsPresentCause {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.InitialContextSetupFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelReq:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsPresentPDUSessionResourceListCxtRelReq {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.UEContextReleaseRequestIEsPresentCause {
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
		if value.TypeValue.Present != ngapType.UEContextReleaseCommandIEsPresentUENGAPIDs {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.UEContextReleaseCommandIEsPresentCause {
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
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsPresentInfoOnRecommendedCellsAndRANNodesForPaging {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelCpl:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsPresentPDUSessionResourceListCxtRelCpl {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UEContextReleaseCompleteIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentRANPagingPriority {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityKey:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentSecurityKey {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIndexToRFSP:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentIndexToRFSP {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDEmergencyFallbackIndicator:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentEmergencyFallbackIndicator {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewAMFUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentNewAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewGUAMI:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentNewGUAMI {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsPresentSRVCCOperationPossible {
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
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCState:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsPresentRRCState {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UEContextModificationResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsPresentCause {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UEContextModificationFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCState:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsPresentRRCState {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.RRCInactiveTransitionReportIEsPresentUserLocationInformation {
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
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverType:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentHandoverType {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentCause {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTargetID:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentTargetID {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDirectForwardingPathAvailability:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentDirectForwardingPathAvailability {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceListHORqd:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentPDUSessionResourceListHORqd {
			return binData, bitEnd, errors.New("HandoverRequiredIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverRequiredIEsPresentSourceToTargetTransparentContainer {
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
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverType:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentHandoverType {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASSecurityParametersFromNGRAN:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentNASSecurityParametersFromNGRAN {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceHandoverList:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentPDUSessionResourceHandoverList {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentPDUSessionResourceToReleaseListHOCmd {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTargetToSourceTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentTargetToSourceTransparentContainer {
			return binData, bitEnd, errors.New("HandoverCommandIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverCommandIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsPresentCause {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverPreparationFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverType:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentHandoverType {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentCause {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityContext:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentSecurityContext {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewSecurityContextInd:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentNewSecurityContextInd {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASC:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentNASC {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSetupListHOReq:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentPDUSessionResourceSetupListHOReq {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentAllowedNSSAI {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceActivation:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentTraceActivation {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMaskedIMEISV:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentMaskedIMEISV {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentSourceToTargetTransparentContainer {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMobilityRestrictionList:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentMobilityRestrictionList {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentLocationReportingRequestType {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGUAMI:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentGUAMI {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRedirectionVoiceFallback:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentRedirectionVoiceFallback {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsPresentSRVCCOperationPossible {
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
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceAdmittedList:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsPresentPDUSessionResourceAdmittedList {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsPresentPDUSessionResourceFailedToSetupListHOAck {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTargetToSourceTransparentContainer:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsPresentTargetToSourceTransparentContainer {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverRequestAcknowledgeIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.HandoverFailureIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverFailureIEsPresentCause {
			return binData, bitEnd, errors.New("HandoverFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.HandoverNotifyIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverNotifyIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverNotifyIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.HandoverNotifyIEsPresentUserLocationInformation {
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
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceAMFUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsPresentSourceAMFUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsPresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsPresentPDUSessionResourceToBeSwitchedDLList {
			return binData, bitEnd, errors.New("PathSwitchRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq:
		if value.TypeValue.Present != ngapType.PathSwitchRequestIEsPresentPDUSessionResourceFailedToSetupListPSReq {
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
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUESecurityCapabilities:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentUESecurityCapabilities {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSecurityContext:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentSecurityContext {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNewSecurityContextInd:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentNewSecurityContextInd {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSwitchedList:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceSwitchedList {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSAck:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceReleasedListPSAck {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentAllowedNSSAI {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentCoreNetworkAssistanceInformationForInactive {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentRRCInactiveTransitionReportRequest {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRedirectionVoiceFallback:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentRedirectionVoiceFallback {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCNAssistedRANTuning:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentCNAssistedRANTuning {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsPresentSRVCCOperationPossible {
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
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSFail:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsPresentPDUSessionResourceReleasedListPSFail {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PathSwitchRequestFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.HandoverCancelIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCancelIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.HandoverCancelIEsPresentCause {
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
		if value.TypeValue.Present != ngapType.HandoverCancelAcknowledgeIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.HandoverCancelAcknowledgeIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.UplinkRANStatusTransferIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkRANStatusTransferIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANStatusTransferTransparentContainer:
		if value.TypeValue.Present != ngapType.UplinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer {
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
		if value.TypeValue.Present != ngapType.DownlinkRANStatusTransferIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkRANStatusTransferIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANStatusTransferTransparentContainer:
		if value.TypeValue.Present != ngapType.DownlinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer {
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
		if value.TypeValue.Present != ngapType.PagingIEsPresentUEPagingIdentity {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPagingDRX:
		if value.TypeValue.Present != ngapType.PagingIEsPresentPagingDRX {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTAIListForPaging:
		if value.TypeValue.Present != ngapType.PagingIEsPresentTAIListForPaging {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPagingPriority:
		if value.TypeValue.Present != ngapType.PagingIEsPresentPagingPriority {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
		if value.TypeValue.Present != ngapType.PagingIEsPresentUERadioCapabilityForPaging {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPagingOrigin:
		if value.TypeValue.Present != ngapType.PagingIEsPresentPagingOrigin {
			return binData, bitEnd, errors.New("PagingIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAssistanceDataForPaging:
		if value.TypeValue.Present != ngapType.PagingIEsPresentAssistanceDataForPaging {
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
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentNASPDU {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRRCEstablishmentCause:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentRRCEstablishmentCause {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDFiveGSTMSI:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentFiveGSTMSI {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFSetID:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentAMFSetID {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEContextRequest:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentUEContextRequest {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentAllowedNSSAI {
			return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetAMFInformationReroute:
		if value.TypeValue.Present != ngapType.InitialUEMessageIEsPresentSourceToTargetAMFInformationReroute {
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
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDOldAMF:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentOldAMF {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANPagingPriority:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentRANPagingPriority {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentNASPDU {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDMobilityRestrictionList:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentMobilityRestrictionList {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIndexToRFSP:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentIndexToRFSP {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentUEAggregateMaximumBitRate {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentAllowedNSSAI {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSRVCCOperationPossible:
		if value.TypeValue.Present != ngapType.DownlinkNASTransportIEsPresentSRVCCOperationPossible {
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
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsPresentNASPDU {
			return binData, bitEnd, errors.New("UplinkNASTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.UplinkNASTransportIEsPresentUserLocationInformation {
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
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNASPDU:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsPresentNASPDU {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.NASNonDeliveryIndicationIEsPresentCause {
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
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFUENGAPID:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGAPMessage:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsPresentNGAPMessage {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFSetID:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsPresentAMFSetID {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAllowedNSSAI:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsPresentAllowedNSSAI {
			return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSourceToTargetAMFInformationReroute:
		if value.TypeValue.Present != ngapType.RerouteNASRequestIEsPresentSourceToTargetAMFInformationReroute {
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
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsPresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANNodeName:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsPresentRANNodeName {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSupportedTAList:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsPresentSupportedTAList {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDefaultPagingDRX:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsPresentDefaultPagingDRX {
			return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERetentionInformation:
		if value.TypeValue.Present != ngapType.NGSetupRequestIEsPresentUERetentionInformation {
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
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsPresentAMFName {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDServedGUAMIList:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsPresentServedGUAMIList {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRelativeAMFCapacity:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsPresentRelativeAMFCapacity {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPLMNSupportList:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsPresentPLMNSupportList {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsPresentCriticalityDiagnostics {
			return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERetentionInformation:
		if value.TypeValue.Present != ngapType.NGSetupResponseIEsPresentUERetentionInformation {
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
		if value.TypeValue.Present != ngapType.NGSetupFailureIEsPresentCause {
			return binData, bitEnd, errors.New("NGSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTimeToWait:
		if value.TypeValue.Present != ngapType.NGSetupFailureIEsPresentTimeToWait {
			return binData, bitEnd, errors.New("NGSetupFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.NGSetupFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsPresentRANNodeName {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSupportedTAList:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsPresentSupportedTAList {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDefaultPagingDRX:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsPresentDefaultPagingDRX {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsPresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTNLAssociationToRemoveList:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateIEsPresentNGRANTNLAssociationToRemoveList {
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
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateFailureIEsPresentCause {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTimeToWait:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateFailureIEsPresentTimeToWait {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.RANConfigurationUpdateFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentAMFName {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDServedGUAMIList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentServedGUAMIList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRelativeAMFCapacity:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentRelativeAMFCapacity {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPLMNSupportList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentPLMNSupportList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationToAddList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToAddList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationToRemoveList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToRemoveList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationToUpdateList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToUpdateList {
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
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationSetupList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTNLAssociationFailedToSetupList:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationFailedToSetupList {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateFailureIEsPresentCause {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTimeToWait:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateFailureIEsPresentTimeToWait {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.AMFConfigurationUpdateFailureIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.AMFStatusIndicationIEsPresentUnavailableGUAMIList {
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
		if value.TypeValue.Present != ngapType.NGResetIEsPresentCause {
			return binData, bitEnd, errors.New("NGResetIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDResetType:
		if value.TypeValue.Present != ngapType.NGResetIEsPresentResetType {
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
		if value.TypeValue.Present != ngapType.NGResetAcknowledgeIEsPresentUEAssociatedLogicalNGConnectionList {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.NGResetAcknowledgeIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsPresentCause {
			return binData, bitEnd, errors.New("ErrorIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.ErrorIndicationIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.OverloadStartIEsPresentAMFOverloadResponse {
			return binData, bitEnd, errors.New("OverloadStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDAMFTrafficLoadReductionIndication:
		if value.TypeValue.Present != ngapType.OverloadStartIEsPresentAMFTrafficLoadReductionIndication {
			return binData, bitEnd, errors.New("OverloadStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDOverloadStartNSSAIList:
		if value.TypeValue.Present != ngapType.OverloadStartIEsPresentOverloadStartNSSAIList {
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
		if value.TypeValue.Present != ngapType.UplinkRANConfigurationTransferIEsPresentSONConfigurationTransferUL {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDENDCSONConfigurationTransferUL:
		if value.TypeValue.Present != ngapType.UplinkRANConfigurationTransferIEsPresentENDCSONConfigurationTransferUL {
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
		if value.TypeValue.Present != ngapType.DownlinkRANConfigurationTransferIEsPresentSONConfigurationTransferDL {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDENDCSONConfigurationTransferDL:
		if value.TypeValue.Present != ngapType.DownlinkRANConfigurationTransferIEsPresentENDCSONConfigurationTransferDL {
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
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentMessageIdentifier {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentSerialNumber {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningAreaList:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentWarningAreaList {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRepetitionPeriod:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentRepetitionPeriod {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNumberOfBroadcastsRequested:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentNumberOfBroadcastsRequested {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningType:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentWarningType {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningSecurityInfo:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentWarningSecurityInfo {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDDataCodingScheme:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentDataCodingScheme {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningMessageContents:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentWarningMessageContents {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDConcurrentWarningMessageInd:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentConcurrentWarningMessageInd {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningAreaCoordinates:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningRequestIEsPresentWarningAreaCoordinates {
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
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsPresentMessageIdentifier {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsPresentSerialNumber {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDBroadcastCompletedAreaList:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsPresentBroadcastCompletedAreaList {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.WriteReplaceWarningResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsPresentMessageIdentifier {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsPresentSerialNumber {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDWarningAreaList:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsPresentWarningAreaList {
			return binData, bitEnd, errors.New("PWSCancelRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCancelAllWarningMessages:
		if value.TypeValue.Present != ngapType.PWSCancelRequestIEsPresentCancelAllWarningMessages {
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
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsPresentMessageIdentifier {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDSerialNumber:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsPresentSerialNumber {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDBroadcastCancelledAreaList:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsPresentBroadcastCancelledAreaList {
			return binData, bitEnd, errors.New("PWSCancelResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.PWSCancelResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsPresentCellIDListForRestart {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsPresentGlobalRANNodeID {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTAIListForRestart:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsPresentTAIListForRestart {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDEmergencyAreaIDListForRestart:
		if value.TypeValue.Present != ngapType.PWSRestartIndicationIEsPresentEmergencyAreaIDListForRestart {
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
		if value.TypeValue.Present != ngapType.PWSFailureIndicationIEsPresentPWSFailedCellIDList {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDGlobalRANNodeID:
		if value.TypeValue.Present != ngapType.PWSFailureIndicationIEsPresentGlobalRANNodeID {
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
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRoutingID:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRoutingID {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU {
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
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRoutingID:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentRoutingID {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU {
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
		if value.TypeValue.Present != ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU {
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
		if value.TypeValue.Present != ngapType.UplinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNRPPaPDU:
		if value.TypeValue.Present != ngapType.UplinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU {
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
		if value.TypeValue.Present != ngapType.TraceStartIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("TraceStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.TraceStartIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("TraceStartIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceActivation:
		if value.TypeValue.Present != ngapType.TraceStartIEsPresentTraceActivation {
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
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTraceID:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsPresentNGRANTraceID {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.TraceFailureIndicationIEsPresentCause {
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
		if value.TypeValue.Present != ngapType.DeactivateTraceIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("DeactivateTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.DeactivateTraceIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("DeactivateTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTraceID:
		if value.TypeValue.Present != ngapType.DeactivateTraceIEsPresentNGRANTraceID {
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
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANTraceID:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsPresentNGRANTraceID {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDNGRANCGI:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsPresentNGRANCGI {
			return binData, bitEnd, errors.New("CellTrafficTraceIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDTraceCollectionEntityIPAddress:
		if value.TypeValue.Present != ngapType.CellTrafficTraceIEsPresentTraceCollectionEntityIPAddress {
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
		if value.TypeValue.Present != ngapType.LocationReportingControlIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingControlIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportingControlIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingControlIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.LocationReportingControlIEsPresentLocationReportingRequestType {
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
		if value.TypeValue.Present != ngapType.LocationReportingFailureIndicationIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportingFailureIndicationIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCause:
		if value.TypeValue.Present != ngapType.LocationReportingFailureIndicationIEsPresentCause {
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
		if value.TypeValue.Present != ngapType.LocationReportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.LocationReportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.LocationReportIEsPresentUserLocationInformation {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUEPresenceInAreaOfInterestList:
		if value.TypeValue.Present != ngapType.LocationReportIEsPresentUEPresenceInAreaOfInterestList {
			return binData, bitEnd, errors.New("LocationReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDLocationReportingRequestType:
		if value.TypeValue.Present != ngapType.LocationReportIEsPresentLocationReportingRequestType {
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
		if value.TypeValue.Present != ngapType.UETNLABindingReleaseRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UETNLABindingReleaseRequestIEsPresentRANUENGAPID {
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
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapability:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsPresentUERadioCapability {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
		if value.TypeValue.Present != ngapType.UERadioCapabilityInfoIndicationIEsPresentUERadioCapabilityForPaging {
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
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckRequestIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckRequestIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUERadioCapability:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckRequestIEsPresentUERadioCapability {
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
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDIMSVoiceSupportIndicator:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsPresentIMSVoiceSupportIndicator {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDCriticalityDiagnostics:
		if value.TypeValue.Present != ngapType.UERadioCapabilityCheckResponseIEsPresentCriticalityDiagnostics {
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
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsPresentAMFUENGAPID {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDRANUENGAPID:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsPresentRANUENGAPID {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDPDUSessionResourceSecondaryRATUsageList:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsPresentPDUSessionResourceSecondaryRATUsageList {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDHandoverFlag:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsPresentHandoverFlag {
			return binData, bitEnd, errors.New("SecondaryRATDataUsageReportIEs: ProtocolIEID: INVALID")
		}
	case ngapType.ProtocolIEIDUserLocationInformation:
		if value.TypeValue.Present != ngapType.SecondaryRATDataUsageReportIEsPresentUserLocationInformation {
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
		if value.TypeValue.Present != ngapType.UplinkRIMInformationTransferIEsPresentRIMInformationTransfer {
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
		if value.TypeValue.Present != ngapType.DownlinkRIMInformationTransferIEsPresentRIMInformationTransfer {
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
