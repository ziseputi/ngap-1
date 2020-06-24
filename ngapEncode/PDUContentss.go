package ngapEncode

import (
	"errors"
	"strconv"

	"ngap/ngapType"
)

func ProtocolIEContainerPDUSessionResourceSetupRequestIEs(value ngapType.ProtocolIEContainerPDUSessionResourceSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceSetupRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceSetupResponseIEs(value ngapType.ProtocolIEContainerPDUSessionResourceSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceSetupResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceSetupResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceReleaseCommandIEs(value ngapType.ProtocolIEContainerPDUSessionResourceReleaseCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleaseCommandIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceReleaseCommandIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceReleaseResponseIEs(value ngapType.ProtocolIEContainerPDUSessionResourceReleaseResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceReleaseResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceReleaseResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyRequestIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyResponseIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceNotifyIEs(value ngapType.ProtocolIEContainerPDUSessionResourceNotifyIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceNotifyIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceNotifyIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceNotifyIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyIndicationIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyConfirmIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyConfirmIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PDUSessionResourceModifyConfirmIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPDUSessionResourceModifyConfirmIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialContextSetupRequestIEs(value ngapType.ProtocolIEContainerInitialContextSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerInitialContextSetupRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = InitialContextSetupRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerInitialContextSetupRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialContextSetupResponseIEs(value ngapType.ProtocolIEContainerInitialContextSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerInitialContextSetupResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = InitialContextSetupResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerInitialContextSetupResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialContextSetupFailureIEs(value ngapType.ProtocolIEContainerInitialContextSetupFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerInitialContextSetupFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = InitialContextSetupFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerInitialContextSetupFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextReleaseRequestIEs(value ngapType.ProtocolIEContainerUEContextReleaseRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUEContextReleaseRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEContextReleaseRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUEContextReleaseRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextReleaseCommandIEs(value ngapType.ProtocolIEContainerUEContextReleaseCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUEContextReleaseCommandIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEContextReleaseCommandIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUEContextReleaseCommandIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextReleaseCompleteIEs(value ngapType.ProtocolIEContainerUEContextReleaseCompleteIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUEContextReleaseCompleteIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEContextReleaseCompleteIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUEContextReleaseCompleteIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextModificationRequestIEs(value ngapType.ProtocolIEContainerUEContextModificationRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUEContextModificationRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEContextModificationRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUEContextModificationRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextModificationResponseIEs(value ngapType.ProtocolIEContainerUEContextModificationResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUEContextModificationResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEContextModificationResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUEContextModificationResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextModificationFailureIEs(value ngapType.ProtocolIEContainerUEContextModificationFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUEContextModificationFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UEContextModificationFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUEContextModificationFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRRCInactiveTransitionReportIEs(value ngapType.ProtocolIEContainerRRCInactiveTransitionReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerRRCInactiveTransitionReportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RRCInactiveTransitionReportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerRRCInactiveTransitionReportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverRequiredIEs(value ngapType.ProtocolIEContainerHandoverRequiredIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverRequiredIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverRequiredIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverRequiredIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverCommandIEs(value ngapType.ProtocolIEContainerHandoverCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverCommandIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverCommandIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverCommandIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverPreparationFailureIEs(value ngapType.ProtocolIEContainerHandoverPreparationFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverPreparationFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverPreparationFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverPreparationFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverRequestIEs(value ngapType.ProtocolIEContainerHandoverRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverRequestAcknowledgeIEs(value ngapType.ProtocolIEContainerHandoverRequestAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverRequestAcknowledgeIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverRequestAcknowledgeIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverRequestAcknowledgeIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverFailureIEs(value ngapType.ProtocolIEContainerHandoverFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverNotifyIEs(value ngapType.ProtocolIEContainerHandoverNotifyIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverNotifyIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverNotifyIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverNotifyIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPathSwitchRequestIEs(value ngapType.ProtocolIEContainerPathSwitchRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPathSwitchRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPathSwitchRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPathSwitchRequestAcknowledgeIEs(value ngapType.ProtocolIEContainerPathSwitchRequestAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestAcknowledgeIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPathSwitchRequestAcknowledgeIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPathSwitchRequestFailureIEs(value ngapType.ProtocolIEContainerPathSwitchRequestFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPathSwitchRequestFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PathSwitchRequestFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPathSwitchRequestFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverCancelIEs(value ngapType.ProtocolIEContainerHandoverCancelIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverCancelIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverCancelIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverCancelIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverCancelAcknowledgeIEs(value ngapType.ProtocolIEContainerHandoverCancelAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerHandoverCancelAcknowledgeIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = HandoverCancelAcknowledgeIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerHandoverCancelAcknowledgeIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkRANStatusTransferIEs(value ngapType.ProtocolIEContainerUplinkRANStatusTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUplinkRANStatusTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UplinkRANStatusTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUplinkRANStatusTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkRANStatusTransferIEs(value ngapType.ProtocolIEContainerDownlinkRANStatusTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkRANStatusTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DownlinkRANStatusTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkRANStatusTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPagingIEs(value ngapType.ProtocolIEContainerPagingIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPagingIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PagingIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPagingIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialUEMessageIEs(value ngapType.ProtocolIEContainerInitialUEMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerInitialUEMessageIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = InitialUEMessageIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerInitialUEMessageIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkNASTransportIEs(value ngapType.ProtocolIEContainerDownlinkNASTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkNASTransportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DownlinkNASTransportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkNASTransportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkNASTransportIEs(value ngapType.ProtocolIEContainerUplinkNASTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUplinkNASTransportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UplinkNASTransportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUplinkNASTransportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNASNonDeliveryIndicationIEs(value ngapType.ProtocolIEContainerNASNonDeliveryIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerNASNonDeliveryIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NASNonDeliveryIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerNASNonDeliveryIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRerouteNASRequestIEs(value ngapType.ProtocolIEContainerRerouteNASRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerRerouteNASRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RerouteNASRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerRerouteNASRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGSetupRequestIEs(value ngapType.ProtocolIEContainerNGSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerNGSetupRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGSetupRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerNGSetupRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGSetupResponseIEs(value ngapType.ProtocolIEContainerNGSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerNGSetupResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGSetupResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerNGSetupResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGSetupFailureIEs(value ngapType.ProtocolIEContainerNGSetupFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerNGSetupFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGSetupFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerNGSetupFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRANConfigurationUpdateIEs(value ngapType.ProtocolIEContainerRANConfigurationUpdateIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerRANConfigurationUpdateIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RANConfigurationUpdateIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerRANConfigurationUpdateIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs(value ngapType.ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RANConfigurationUpdateAcknowledgeIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRANConfigurationUpdateFailureIEs(value ngapType.ProtocolIEContainerRANConfigurationUpdateFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerRANConfigurationUpdateFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = RANConfigurationUpdateFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerRANConfigurationUpdateFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFConfigurationUpdateIEs(value ngapType.ProtocolIEContainerAMFConfigurationUpdateIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerAMFConfigurationUpdateIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFConfigurationUpdateIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerAMFConfigurationUpdateIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs(value ngapType.ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFConfigurationUpdateAcknowledgeIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFConfigurationUpdateFailureIEs(value ngapType.ProtocolIEContainerAMFConfigurationUpdateFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerAMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFConfigurationUpdateFailureIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerAMFConfigurationUpdateFailureIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFStatusIndicationIEs(value ngapType.ProtocolIEContainerAMFStatusIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerAMFStatusIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = AMFStatusIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerAMFStatusIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGResetIEs(value ngapType.ProtocolIEContainerNGResetIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerNGResetIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGResetIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerNGResetIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGResetAcknowledgeIEs(value ngapType.ProtocolIEContainerNGResetAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerNGResetAcknowledgeIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = NGResetAcknowledgeIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerNGResetAcknowledgeIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerErrorIndicationIEs(value ngapType.ProtocolIEContainerErrorIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerErrorIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = ErrorIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerErrorIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerOverloadStartIEs(value ngapType.ProtocolIEContainerOverloadStartIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerOverloadStartIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = OverloadStartIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerOverloadStartIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerOverloadStopIEs(value ngapType.ProtocolIEContainerOverloadStopIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerOverloadStopIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = OverloadStopIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerOverloadStopIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkRANConfigurationTransferIEs(value ngapType.ProtocolIEContainerUplinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUplinkRANConfigurationTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UplinkRANConfigurationTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUplinkRANConfigurationTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkRANConfigurationTransferIEs(value ngapType.ProtocolIEContainerDownlinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DownlinkRANConfigurationTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkRANConfigurationTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerWriteReplaceWarningRequestIEs(value ngapType.ProtocolIEContainerWriteReplaceWarningRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerWriteReplaceWarningRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = WriteReplaceWarningRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerWriteReplaceWarningRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerWriteReplaceWarningResponseIEs(value ngapType.ProtocolIEContainerWriteReplaceWarningResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerWriteReplaceWarningResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = WriteReplaceWarningResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerWriteReplaceWarningResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSCancelRequestIEs(value ngapType.ProtocolIEContainerPWSCancelRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPWSCancelRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PWSCancelRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPWSCancelRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSCancelResponseIEs(value ngapType.ProtocolIEContainerPWSCancelResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPWSCancelResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PWSCancelResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPWSCancelResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSRestartIndicationIEs(value ngapType.ProtocolIEContainerPWSRestartIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPWSRestartIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PWSRestartIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPWSRestartIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSFailureIndicationIEs(value ngapType.ProtocolIEContainerPWSFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerPWSFailureIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PWSFailureIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerPWSFailureIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DownlinkUEAssociatedNRPPaTransportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UplinkUEAssociatedNRPPaTransportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DownlinkNonUEAssociatedNRPPaTransportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UplinkNonUEAssociatedNRPPaTransportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerTraceStartIEs(value ngapType.ProtocolIEContainerTraceStartIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerTraceStartIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TraceStartIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerTraceStartIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerTraceFailureIndicationIEs(value ngapType.ProtocolIEContainerTraceFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerTraceFailureIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = TraceFailureIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerTraceFailureIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDeactivateTraceIEs(value ngapType.ProtocolIEContainerDeactivateTraceIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDeactivateTraceIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DeactivateTraceIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDeactivateTraceIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerCellTrafficTraceIEs(value ngapType.ProtocolIEContainerCellTrafficTraceIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerCellTrafficTraceIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = CellTrafficTraceIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerCellTrafficTraceIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerLocationReportingControlIEs(value ngapType.ProtocolIEContainerLocationReportingControlIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerLocationReportingControlIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LocationReportingControlIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerLocationReportingControlIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerLocationReportingFailureIndicationIEs(value ngapType.ProtocolIEContainerLocationReportingFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerLocationReportingFailureIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LocationReportingFailureIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerLocationReportingFailureIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerLocationReportIEs(value ngapType.ProtocolIEContainerLocationReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerLocationReportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = LocationReportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerLocationReportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUETNLABindingReleaseRequestIEs(value ngapType.ProtocolIEContainerUETNLABindingReleaseRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUETNLABindingReleaseRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UETNLABindingReleaseRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUETNLABindingReleaseRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUERadioCapabilityInfoIndicationIEs(value ngapType.ProtocolIEContainerUERadioCapabilityInfoIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UERadioCapabilityInfoIndicationIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUERadioCapabilityInfoIndicationIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUERadioCapabilityCheckRequestIEs(value ngapType.ProtocolIEContainerUERadioCapabilityCheckRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UERadioCapabilityCheckRequestIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUERadioCapabilityCheckRequestIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUERadioCapabilityCheckResponseIEs(value ngapType.ProtocolIEContainerUERadioCapabilityCheckResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UERadioCapabilityCheckResponseIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUERadioCapabilityCheckResponseIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func PrivateIEContainerPrivateMessageIEs(value ngapType.PrivateIEContainerPrivateMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfPrivateIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PrivateIEContainerPrivateMessageIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = PrivateMessageIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PrivateIEContainerPrivateMessageIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerSecondaryRATDataUsageReportIEs(value ngapType.ProtocolIEContainerSecondaryRATDataUsageReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerSecondaryRATDataUsageReportIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = SecondaryRATDataUsageReportIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerSecondaryRATDataUsageReportIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkRIMInformationTransferIEs(value ngapType.ProtocolIEContainerUplinkRIMInformationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerUplinkRIMInformationTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = UplinkRIMInformationTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerUplinkRIMInformationTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkRIMInformationTransferIEs(value ngapType.ProtocolIEContainerDownlinkRIMInformationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkRIMInformationTransferIEs: " + err.Error())
	}
	for i := 0; i < numIEs; i++ {
		var binIEs []byte
		binIEs, bitEnd, err = DownlinkRIMInformationTransferIEs(value.List[i], binIEs, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ProtocolIEContainerDownlinkRIMInformationTransferIEs: [" + strconv.Itoa(i) + "]: " + err.Error())
		}
		binData = append(binData, binIEs...)
	}
	return binData, bitEnd, err
}
