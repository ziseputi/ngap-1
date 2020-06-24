// Created By HaoDHH-245789 VHT2020
package ngapEncode

import (
	"errors"
	"strconv"

	"ngap/ngapType"
)

func AMFConfigurationUpdate(value ngapType.AMFConfigurationUpdate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerAMFConfigurationUpdateIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdate: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFConfigurationUpdateIEs(value ngapType.ProtocolIEContainerAMFConfigurationUpdateIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFConfigurationUpdateIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func AMFConfigurationUpdateIEsValue(value ngapType.AMFConfigurationUpdateIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFName:
		if value.AMFName == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: AMFName: NIL")
		}
		binData, bitEnd, err = AMFName(*value.AMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentServedGUAMIList:
		if value.ServedGUAMIList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: ServedGUAMIList: NIL")
		}
		binData, bitEnd, err = ServedGUAMIList(*value.ServedGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentRelativeAMFCapacity:
		if value.RelativeAMFCapacity == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: RelativeAMFCapacity: NIL")
		}
		binData, bitEnd, err = RelativeAMFCapacity(*value.RelativeAMFCapacity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentPLMNSupportList:
		if value.PLMNSupportList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: PLMNSupportList: NIL")
		}
		binData, bitEnd, err = PLMNSupportList(*value.PLMNSupportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToAddList:
		if value.AMFTNLAssociationToAddList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: AMFTNLAssociationToAddList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToAddList(*value.AMFTNLAssociationToAddList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToRemoveList:
		if value.AMFTNLAssociationToRemoveList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: AMFTNLAssociationToRemoveList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToRemoveList(*value.AMFTNLAssociationToRemoveList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToUpdateList:
		if value.AMFTNLAssociationToUpdateList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: AMFTNLAssociationToUpdateList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationToUpdateList(*value.AMFTNLAssociationToUpdateList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCancel(value ngapType.HandoverCancel, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverCancelIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancel: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverCancelIEs(value ngapType.ProtocolIEContainerHandoverCancelIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCancelIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverCancelIEsValue(value ngapType.HandoverCancelIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCancelIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsValue: " + err.Error())
		}
	case ngapType.HandoverCancelIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsValue: " + err.Error())
		}
	case ngapType.HandoverCancelIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCancelIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequired(value ngapType.HandoverRequired, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverRequiredIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequired: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverRequiredIEs(value ngapType.ProtocolIEContainerHandoverRequiredIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequiredIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequiredIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverRequiredIEsValue(value ngapType.HandoverRequiredIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequiredIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentTargetID:
		if value.TargetID == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: TargetID: NIL")
		}
		binData, bitEnd, err = TargetID(*value.TargetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentDirectForwardingPathAvailability:
		if value.DirectForwardingPathAvailability == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: DirectForwardingPathAvailability: NIL")
		}
		binData, bitEnd, err = DirectForwardingPathAvailability(*value.DirectForwardingPathAvailability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentPDUSessionResourceListHORqd:
		if value.PDUSessionResourceListHORqd == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: PDUSessionResourceListHORqd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListHORqd(*value.PDUSessionResourceListHORqd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequiredIEsTypeValuePresentSourceToTargetTransparentContainer:
		if value.SourceToTargetTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: SourceToTargetTransparentContainer: NIL")
		}
		binData, bitEnd, err = SourceToTargetTransparentContainer(*value.SourceToTargetTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequiredIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequiredIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequest(value ngapType.HandoverRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverRequestIEs(value ngapType.ProtocolIEContainerHandoverRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformation:
		if value.TypeValue.Present != ngapType.HandoverRequestIEsTypeValuePresentCoreNetworkAssistanceInformation {
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
	default:
		return binData, bitEnd, errors.New("HandoverRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverRequestIEsValue(value ngapType.HandoverRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentCoreNetworkAssistanceInformation:
		if value.CoreNetworkAssistanceInformation == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: CoreNetworkAssistanceInformation: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformation(*value.CoreNetworkAssistanceInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentSecurityContext:
		if value.SecurityContext == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: SecurityContext: NIL")
		}
		binData, bitEnd, err = SecurityContext(*value.SecurityContext, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentNewSecurityContextInd:
		if value.NewSecurityContextInd == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: NewSecurityContextInd: NIL")
		}
		binData, bitEnd, err = NewSecurityContextInd(*value.NewSecurityContextInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentNASC:
		if value.NASC == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: NASC: NIL")
		}
		binData, bitEnd, err = NASC(*value.NASC, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentPDUSessionResourceSetupListHOReq:
		if value.PDUSessionResourceSetupListHOReq == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: PDUSessionResourceSetupListHOReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListHOReq(*value.PDUSessionResourceSetupListHOReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentMaskedIMEISV:
		if value.MaskedIMEISV == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: MaskedIMEISV: NIL")
		}
		binData, bitEnd, err = MaskedIMEISV(*value.MaskedIMEISV, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentSourceToTargetTransparentContainer:
		if value.SourceToTargetTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: SourceToTargetTransparentContainer: NIL")
		}
		binData, bitEnd, err = SourceToTargetTransparentContainer(*value.SourceToTargetTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestIEsTypeValuePresentGUAMI:
		if value.GUAMI == nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: GUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.GUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialContextSetupRequest(value ngapType.InitialContextSetupRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerInitialContextSetupRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialContextSetupRequestIEs(value ngapType.ProtocolIEContainerInitialContextSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformation:
		if value.TypeValue.Present != ngapType.InitialContextSetupRequestIEsTypeValuePresentCoreNetworkAssistanceInformation {
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
	default:
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialContextSetupRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func InitialContextSetupRequestIEsValue(value ngapType.InitialContextSetupRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentOldAMF:
		if value.OldAMF == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: OldAMF: NIL")
		}
		binData, bitEnd, err = OldAMF(*value.OldAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentCoreNetworkAssistanceInformation:
		if value.CoreNetworkAssistanceInformation == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: CoreNetworkAssistanceInformation: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformation(*value.CoreNetworkAssistanceInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentGUAMI:
		if value.GUAMI == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: GUAMI: NIL")
		}
		binData, bitEnd, err = GUAMI(*value.GUAMI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListCxtReq:
		if value.PDUSessionResourceSetupListCxtReq == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: PDUSessionResourceSetupListCxtReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListCxtReq(*value.PDUSessionResourceSetupListCxtReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentSecurityKey:
		if value.SecurityKey == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: SecurityKey: NIL")
		}
		binData, bitEnd, err = SecurityKey(*value.SecurityKey, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentMaskedIMEISV:
		if value.MaskedIMEISV == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: MaskedIMEISV: NIL")
		}
		binData, bitEnd, err = MaskedIMEISV(*value.MaskedIMEISV, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentEmergencyFallbackIndicator:
		if value.EmergencyFallbackIndicator == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: EmergencyFallbackIndicator: NIL")
		}
		binData, bitEnd, err = EmergencyFallbackIndicator(*value.EmergencyFallbackIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGReset(value ngapType.NGReset, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGResetIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGReset: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGResetIEs(value ngapType.ProtocolIEContainerNGResetIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGResetIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func NGResetIEsValue(value ngapType.NGResetIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGResetIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NGResetIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetIEsValue: " + err.Error())
		}
	case ngapType.NGResetIEsTypeValuePresentResetType:
		if value.ResetType == nil {
			return binData, bitEnd, errors.New("NGResetIEsValue: ResetType: NIL")
		}
		binData, bitEnd, err = ResetType(*value.ResetType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGResetIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGSetupRequest(value ngapType.NGSetupRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGSetupRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGSetupRequestIEs(value ngapType.ProtocolIEContainerNGSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("NGSetupRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGSetupRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func NGSetupRequestIEsValue(value ngapType.NGSetupRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGSetupRequestIEsTypeValuePresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentRANNodeName:
		if value.RANNodeName == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: RANNodeName: NIL")
		}
		binData, bitEnd, err = RANNodeName(*value.RANNodeName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentSupportedTAList:
		if value.SupportedTAList == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: SupportedTAList: NIL")
		}
		binData, bitEnd, err = SupportedTAList(*value.SupportedTAList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.NGSetupRequestIEsTypeValuePresentDefaultPagingDRX:
		if value.DefaultPagingDRX == nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: DefaultPagingDRX: NIL")
		}
		binData, bitEnd, err = DefaultPagingDRX(*value.DefaultPagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGSetupRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequest(value ngapType.PathSwitchRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPathSwitchRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPathSwitchRequestIEs(value ngapType.ProtocolIEContainerPathSwitchRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PathSwitchRequestIEsValue(value ngapType.PathSwitchRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentSourceAMFUENGAPID:
		if value.SourceAMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: SourceAMFUENGAPID: NIL")
		}
		binData, bitEnd, err = SourceAMFUENGAPID(*value.SourceAMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceToBeSwitchedDLList:
		if value.PDUSessionResourceToBeSwitchedDLList == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: PDUSessionResourceToBeSwitchedDLList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToBeSwitchedDLList(*value.PDUSessionResourceToBeSwitchedDLList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceFailedToSetupListPSReq:
		if value.PDUSessionResourceFailedToSetupListPSReq == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: PDUSessionResourceFailedToSetupListPSReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListPSReq(*value.PDUSessionResourceFailedToSetupListPSReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyRequest(value ngapType.PDUSessionResourceModifyRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyRequestIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceModifyRequestIEsValue(value ngapType.PDUSessionResourceModifyRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentPDUSessionResourceModifyListModReq:
		if value.PDUSessionResourceModifyListModReq == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: PDUSessionResourceModifyListModReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModReq(*value.PDUSessionResourceModifyListModReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndication(value ngapType.PDUSessionResourceModifyIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyIndicationIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceModifyIndicationIEsValue(value ngapType.PDUSessionResourceModifyIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentPDUSessionResourceModifyListModInd:
		if value.PDUSessionResourceModifyListModInd == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: PDUSessionResourceModifyListModInd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModInd(*value.PDUSessionResourceModifyListModInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseCommand(value ngapType.PDUSessionResourceReleaseCommand, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceReleaseCommandIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommand: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceReleaseCommandIEs(value ngapType.ProtocolIEContainerPDUSessionResourceReleaseCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleaseCommandIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseCommandIEsValue(value ngapType.PDUSessionResourceReleaseCommandIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentPDUSessionResourceToReleaseListRelCmd:
		if value.PDUSessionResourceToReleaseListRelCmd == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: PDUSessionResourceToReleaseListRelCmd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToReleaseListRelCmd(*value.PDUSessionResourceToReleaseListRelCmd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommandIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupRequest(value ngapType.PDUSessionResourceSetupRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceSetupRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceSetupRequestIEs(value ngapType.ProtocolIEContainerPDUSessionResourceSetupRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceSetupRequestIEsValue(value ngapType.PDUSessionResourceSetupRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListSUReq:
		if value.PDUSessionResourceSetupListSUReq == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: PDUSessionResourceSetupListSUReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListSUReq(*value.PDUSessionResourceSetupListSUReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSCancelRequest(value ngapType.PWSCancelRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPWSCancelRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSCancelRequestIEs(value ngapType.ProtocolIEContainerPWSCancelRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSCancelRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PWSCancelRequestIEsValue(value ngapType.PWSCancelRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSCancelRequestIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsTypeValuePresentWarningAreaList:
		if value.WarningAreaList == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: WarningAreaList: NIL")
		}
		binData, bitEnd, err = WarningAreaList(*value.WarningAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: " + err.Error())
		}
	case ngapType.PWSCancelRequestIEsTypeValuePresentCancelAllWarningMessages:
		if value.CancelAllWarningMessages == nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: CancelAllWarningMessages: NIL")
		}
		binData, bitEnd, err = CancelAllWarningMessages(*value.CancelAllWarningMessages, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSCancelRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANConfigurationUpdate(value ngapType.RANConfigurationUpdate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRANConfigurationUpdateIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdate: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRANConfigurationUpdateIEs(value ngapType.ProtocolIEContainerRANConfigurationUpdateIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANConfigurationUpdateIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func RANConfigurationUpdateIEsValue(value ngapType.RANConfigurationUpdateIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentRANNodeName:
		if value.RANNodeName == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: RANNodeName: NIL")
		}
		binData, bitEnd, err = RANNodeName(*value.RANNodeName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentSupportedTAList:
		if value.SupportedTAList == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: SupportedTAList: NIL")
		}
		binData, bitEnd, err = SupportedTAList(*value.SupportedTAList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateIEsTypeValuePresentDefaultPagingDRX:
		if value.DefaultPagingDRX == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: DefaultPagingDRX: NIL")
		}
		binData, bitEnd, err = DefaultPagingDRX(*value.DefaultPagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextModificationRequest(value ngapType.UEContextModificationRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextModificationRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextModificationRequestIEs(value ngapType.ProtocolIEContainerUEContextModificationRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformation:
		if value.TypeValue.Present != ngapType.UEContextModificationRequestIEsTypeValuePresentCoreNetworkAssistanceInformation {
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
	default:
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextModificationRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UEContextModificationRequestIEsValue(value ngapType.UEContextModificationRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextModificationRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentSecurityKey:
		if value.SecurityKey == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: SecurityKey: NIL")
		}
		binData, bitEnd, err = SecurityKey(*value.SecurityKey, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentCoreNetworkAssistanceInformation:
		if value.CoreNetworkAssistanceInformation == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: CoreNetworkAssistanceInformation: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformation(*value.CoreNetworkAssistanceInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentEmergencyFallbackIndicator:
		if value.EmergencyFallbackIndicator == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: EmergencyFallbackIndicator: NIL")
		}
		binData, bitEnd, err = EmergencyFallbackIndicator(*value.EmergencyFallbackIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentNewAMFUENGAPID:
		if value.NewAMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: NewAMFUENGAPID: NIL")
		}
		binData, bitEnd, err = NewAMFUENGAPID(*value.NewAMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextReleaseCommand(value ngapType.UEContextReleaseCommand, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextReleaseCommandIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommand: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextReleaseCommandIEs(value ngapType.ProtocolIEContainerUEContextReleaseCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextReleaseCommandIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UEContextReleaseCommandIEsValue(value ngapType.UEContextReleaseCommandIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextReleaseCommandIEsTypeValuePresentUENGAPIDs:
		if value.UENGAPIDs == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsValue: UENGAPIDs: NIL")
		}
		binData, bitEnd, err = UENGAPIDs(*value.UENGAPIDs, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCommandIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCommandIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseCommandIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityCheckRequest(value ngapType.UERadioCapabilityCheckRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUERadioCapabilityCheckRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUERadioCapabilityCheckRequestIEs(value ngapType.ProtocolIEContainerUERadioCapabilityCheckRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityCheckRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UERadioCapabilityCheckRequestIEsValue(value ngapType.UERadioCapabilityCheckRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func WriteReplaceWarningRequest(value ngapType.WriteReplaceWarningRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerWriteReplaceWarningRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerWriteReplaceWarningRequestIEs(value ngapType.ProtocolIEContainerWriteReplaceWarningRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = WriteReplaceWarningRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func WriteReplaceWarningRequestIEsValue(value ngapType.WriteReplaceWarningRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaList:
		if value.WarningAreaList == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: WarningAreaList: NIL")
		}
		binData, bitEnd, err = WarningAreaList(*value.WarningAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentRepetitionPeriod:
		if value.RepetitionPeriod == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: RepetitionPeriod: NIL")
		}
		binData, bitEnd, err = RepetitionPeriod(*value.RepetitionPeriod, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentNumberOfBroadcastsRequested:
		if value.NumberOfBroadcastsRequested == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: NumberOfBroadcastsRequested: NIL")
		}
		binData, bitEnd, err = NumberOfBroadcastsRequested(*value.NumberOfBroadcastsRequested, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningType:
		if value.WarningType == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: WarningType: NIL")
		}
		binData, bitEnd, err = WarningType(*value.WarningType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningSecurityInfo:
		if value.WarningSecurityInfo == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: WarningSecurityInfo: NIL")
		}
		binData, bitEnd, err = WarningSecurityInfo(*value.WarningSecurityInfo, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentDataCodingScheme:
		if value.DataCodingScheme == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: DataCodingScheme: NIL")
		}
		binData, bitEnd, err = DataCodingScheme(*value.DataCodingScheme, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningMessageContents:
		if value.WarningMessageContents == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: WarningMessageContents: NIL")
		}
		binData, bitEnd, err = WarningMessageContents(*value.WarningMessageContents, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentConcurrentWarningMessageInd:
		if value.ConcurrentWarningMessageInd == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: ConcurrentWarningMessageInd: NIL")
		}
		binData, bitEnd, err = ConcurrentWarningMessageInd(*value.ConcurrentWarningMessageInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaCoordinates:
		if value.WarningAreaCoordinates == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: WarningAreaCoordinates: NIL")
		}
		binData, bitEnd, err = WarningAreaCoordinates(*value.WarningAreaCoordinates, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("WriteReplaceWarningRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFStatusIndication(value ngapType.AMFStatusIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerAMFStatusIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFStatusIndicationIEs(value ngapType.ProtocolIEContainerAMFStatusIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFStatusIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFStatusIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func AMFStatusIndicationIEsValue(value ngapType.AMFStatusIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFStatusIndicationIEsTypeValuePresentUnavailableGUAMIList:
		if value.UnavailableGUAMIList == nil {
			return binData, bitEnd, errors.New("AMFStatusIndicationIEsValue: UnavailableGUAMIList: NIL")
		}
		binData, bitEnd, err = UnavailableGUAMIList(*value.UnavailableGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFStatusIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFStatusIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func CellTrafficTrace(value ngapType.CellTrafficTrace, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerCellTrafficTraceIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTrace: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerCellTrafficTraceIEs(value ngapType.ProtocolIEContainerCellTrafficTraceIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = CellTrafficTraceIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTraceIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func CellTrafficTraceIEsValue(value ngapType.CellTrafficTraceIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.CellTrafficTraceIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentNGRANCGI:
		if value.NGRANCGI == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: NGRANCGI: NIL")
		}
		binData, bitEnd, err = NGRANCGI(*value.NGRANCGI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: " + err.Error())
		}
	case ngapType.CellTrafficTraceIEsTypeValuePresentTraceCollectionEntityIPAddress:
		if value.TraceCollectionEntityIPAddress == nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: TraceCollectionEntityIPAddress: NIL")
		}
		binData, bitEnd, err = TraceCollectionEntityIPAddress(*value.TraceCollectionEntityIPAddress, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("CellTrafficTraceIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DeactivateTrace(value ngapType.DeactivateTrace, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDeactivateTraceIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTrace: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDeactivateTraceIEs(value ngapType.ProtocolIEContainerDeactivateTraceIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTraceIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTraceIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DeactivateTraceIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DeactivateTraceIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func DeactivateTraceIEsValue(value ngapType.DeactivateTraceIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DeactivateTraceIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsValue: " + err.Error())
		}
	case ngapType.DeactivateTraceIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsValue: " + err.Error())
		}
	case ngapType.DeactivateTraceIEsTypeValuePresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DeactivateTraceIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DeactivateTraceIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkNASTransport(value ngapType.DownlinkNASTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkNASTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkNASTransportIEs(value ngapType.ProtocolIEContainerDownlinkNASTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkNASTransportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func DownlinkNASTransportIEsValue(value ngapType.DownlinkNASTransportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkNASTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentOldAMF:
		if value.OldAMF == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: OldAMF: NIL")
		}
		binData, bitEnd, err = OldAMF(*value.OldAMF, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentRANPagingPriority:
		if value.RANPagingPriority == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: RANPagingPriority: NIL")
		}
		binData, bitEnd, err = RANPagingPriority(*value.RANPagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentMobilityRestrictionList:
		if value.MobilityRestrictionList == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: MobilityRestrictionList: NIL")
		}
		binData, bitEnd, err = MobilityRestrictionList(*value.MobilityRestrictionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentIndexToRFSP:
		if value.IndexToRFSP == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: IndexToRFSP: NIL")
		}
		binData, bitEnd, err = IndexToRFSP(*value.IndexToRFSP, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentUEAggregateMaximumBitRate:
		if value.UEAggregateMaximumBitRate == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: UEAggregateMaximumBitRate: NIL")
		}
		binData, bitEnd, err = UEAggregateMaximumBitRate(*value.UEAggregateMaximumBitRate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNASTransportIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkNASTransportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkNonUEAssociatedNRPPaTransport(value ngapType.DownlinkNonUEAssociatedNRPPaTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerDownlinkNonUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkNonUEAssociatedNRPPaTransportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func DownlinkNonUEAssociatedNRPPaTransportIEsValue(value ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkNonUEAssociatedNRPPaTransportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkRANConfigurationTransfer(value ngapType.DownlinkRANConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkRANConfigurationTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkRANConfigurationTransferIEs(value ngapType.ProtocolIEContainerDownlinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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

func DownlinkRANConfigurationTransferIEs(value ngapType.DownlinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDSONConfigurationTransferDL:
		if value.TypeValue.Present != ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferDL {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkRANConfigurationTransferIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func DownlinkRANConfigurationTransferIEsValue(value ngapType.DownlinkRANConfigurationTransferIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferDL:
		if value.SONConfigurationTransferDL == nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsValue: SONConfigurationTransferDL: NIL")
		}
		binData, bitEnd, err = SONConfigurationTransferDL(*value.SONConfigurationTransferDL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransferIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkRANStatusTransfer(value ngapType.DownlinkRANStatusTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkRANStatusTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkRANStatusTransferIEs(value ngapType.ProtocolIEContainerDownlinkRANStatusTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkRANStatusTransferIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func DownlinkRANStatusTransferIEsValue(value ngapType.DownlinkRANStatusTransferIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: " + err.Error())
		}
	case ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: " + err.Error())
		}
	case ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer:
		if value.RANStatusTransferTransparentContainer == nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: RANStatusTransferTransparentContainer: NIL")
		}
		binData, bitEnd, err = RANStatusTransferTransparentContainer(*value.RANStatusTransferTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkRANStatusTransferIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func DownlinkUEAssociatedNRPPaTransport(value ngapType.DownlinkUEAssociatedNRPPaTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerDownlinkUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = DownlinkUEAssociatedNRPPaTransportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func DownlinkUEAssociatedNRPPaTransportIEsValue(value ngapType.DownlinkUEAssociatedNRPPaTransportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("DownlinkUEAssociatedNRPPaTransportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func ErrorIndication(value ngapType.ErrorIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerErrorIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerErrorIndicationIEs(value ngapType.ProtocolIEContainerErrorIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = ErrorIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func ErrorIndicationIEsValue(value ngapType.ErrorIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.ErrorIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: " + err.Error())
		}
	case ngapType.ErrorIndicationIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("ErrorIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("ErrorIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverNotify(value ngapType.HandoverNotify, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverNotifyIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotify: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverNotifyIEs(value ngapType.ProtocolIEContainerHandoverNotifyIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotifyIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotifyIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverNotifyIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverNotifyIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverNotifyIEsValue(value ngapType.HandoverNotifyIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverNotifyIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsValue: " + err.Error())
		}
	case ngapType.HandoverNotifyIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsValue: " + err.Error())
		}
	case ngapType.HandoverNotifyIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverNotifyIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverNotifyIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialUEMessage(value ngapType.InitialUEMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerInitialUEMessageIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessage: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialUEMessageIEs(value ngapType.ProtocolIEContainerInitialUEMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("InitialUEMessageIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessageIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessageIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialUEMessageIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialUEMessageIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func InitialUEMessageIEsValue(value ngapType.InitialUEMessageIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialUEMessageIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentRRCEstablishmentCause:
		if value.RRCEstablishmentCause == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: RRCEstablishmentCause: NIL")
		}
		binData, bitEnd, err = RRCEstablishmentCause(*value.RRCEstablishmentCause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentFiveGSTMSI:
		if value.FiveGSTMSI == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: FiveGSTMSI: NIL")
		}
		binData, bitEnd, err = FiveGSTMSI(*value.FiveGSTMSI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentAMFSetID:
		if value.AMFSetID == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: AMFSetID: NIL")
		}
		binData, bitEnd, err = AMFSetID(*value.AMFSetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentUEContextRequest:
		if value.UEContextRequest == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: UEContextRequest: NIL")
		}
		binData, bitEnd, err = UEContextRequest(*value.UEContextRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	case ngapType.InitialUEMessageIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialUEMessageIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialUEMessageIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReport(value ngapType.LocationReport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerLocationReportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerLocationReportIEs(value ngapType.ProtocolIEContainerLocationReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func LocationReportIEsValue(value ngapType.LocationReportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentUEPresenceInAreaOfInterestList:
		if value.UEPresenceInAreaOfInterestList == nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: UEPresenceInAreaOfInterestList: NIL")
		}
		binData, bitEnd, err = UEPresenceInAreaOfInterestList(*value.UEPresenceInAreaOfInterestList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: " + err.Error())
		}
	case ngapType.LocationReportIEsTypeValuePresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReportingControl(value ngapType.LocationReportingControl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerLocationReportingControlIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControl: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerLocationReportingControlIEs(value ngapType.ProtocolIEContainerLocationReportingControlIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControlIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControlIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportingControlIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingControlIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func LocationReportingControlIEsValue(value ngapType.LocationReportingControlIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportingControlIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsValue: " + err.Error())
		}
	case ngapType.LocationReportingControlIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsValue: " + err.Error())
		}
	case ngapType.LocationReportingControlIEsTypeValuePresentLocationReportingRequestType:
		if value.LocationReportingRequestType == nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsValue: LocationReportingRequestType: NIL")
		}
		binData, bitEnd, err = LocationReportingRequestType(*value.LocationReportingRequestType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingControlIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingControlIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func LocationReportingFailureIndication(value ngapType.LocationReportingFailureIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerLocationReportingFailureIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerLocationReportingFailureIndicationIEs(value ngapType.ProtocolIEContainerLocationReportingFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = LocationReportingFailureIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func LocationReportingFailureIndicationIEsValue(value ngapType.LocationReportingFailureIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.LocationReportingFailureIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: " + err.Error())
		}
	case ngapType.LocationReportingFailureIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: " + err.Error())
		}
	case ngapType.LocationReportingFailureIndicationIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("LocationReportingFailureIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NASNonDeliveryIndication(value ngapType.NASNonDeliveryIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNASNonDeliveryIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNASNonDeliveryIndicationIEs(value ngapType.ProtocolIEContainerNASNonDeliveryIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NASNonDeliveryIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func NASNonDeliveryIndicationIEsValue(value ngapType.NASNonDeliveryIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: " + err.Error())
		}
	case ngapType.NASNonDeliveryIndicationIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NASNonDeliveryIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func OverloadStart(value ngapType.OverloadStart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerOverloadStartIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStart: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerOverloadStartIEs(value ngapType.ProtocolIEContainerOverloadStartIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = OverloadStartIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStartIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func OverloadStartIEsValue(value ngapType.OverloadStartIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.OverloadStartIEsTypeValuePresentAMFOverloadResponse:
		if value.AMFOverloadResponse == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsValue: AMFOverloadResponse: NIL")
		}
		binData, bitEnd, err = AMFOverloadResponse(*value.AMFOverloadResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsValue: " + err.Error())
		}
	case ngapType.OverloadStartIEsTypeValuePresentAMFTrafficLoadReductionIndication:
		if value.AMFTrafficLoadReductionIndication == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsValue: AMFTrafficLoadReductionIndication: NIL")
		}
		binData, bitEnd, err = AMFTrafficLoadReductionIndication(*value.AMFTrafficLoadReductionIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsValue: " + err.Error())
		}
	case ngapType.OverloadStartIEsTypeValuePresentOverloadStartNSSAIList:
		if value.OverloadStartNSSAIList == nil {
			return binData, bitEnd, errors.New("OverloadStartIEsValue: OverloadStartNSSAIList: NIL")
		}
		binData, bitEnd, err = OverloadStartNSSAIList(*value.OverloadStartNSSAIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("OverloadStartIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("OverloadStartIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func OverloadStop(value ngapType.OverloadStop, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	// binData, bitEnd, err = ProtocolIEContainerOverloadStopIEs(value.ProtocolIEs, binData, bitEnd)
	// if err != nil {
	//	return binData, bitEnd, errors.New("OverloadStop: " + err.Error())
	// }
	return binData, bitEnd, err
}

// func ProtocolIEContainerOverloadStopIEs (value ngapType.ProtocolIEContainerOverloadStopIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
//	var err error
//	numIEs := len(value.List)
//	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
//	for i:=0; i<numIEs; i++ {
//		var binIEs []byte
//		binIEs, bitEnd, err = OverloadStopIEs(value.List[i], binIEs, bitEnd)
//		if err != nil {
//			return binData, bitEnd, errors.New("ProtocolIEContainerOverloadStopIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
//		}
//		binData = append(binData, binIEs...)
//	}
//	return binData, bitEnd, err
// }

// func OverloadStopIEs (value ngapType.OverloadStopIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
//	var err error
//	switch value.ProtocolIEID.Value {
//	case ngapType.ProtocolIEIDNothing:
//		if value.TypeValue.Present != ngapType.OverloadStopIEsTypeValuePresentNothing {
//			return binData, bitEnd, errors.New("OverloadStopIEs: ProtocolIEID: INVALID")
//		}
//	default:
//		return binData, bitEnd, errors.New("OverloadStopIEs: ProtocolIEID: INVALID")
//	}
//	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
//	if err != nil {
//		return binData, bitEnd, errors.New("OverloadStopIEs: " + err.Error())
//	}
//	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
//	if err != nil {
//		return binData, bitEnd, errors.New("OverloadStopIEs: " + err.Error())
//	}
//	var binValue []byte
//	binValue, bitEnd, err = OverloadStopIEsValue(value.Value, binValue, 8)
//	if err != nil {
//		return binData, bitEnd, errors.New("OverloadStopIEs: " + err.Error())
//	}
//	binValue = EncodeLengthValue(binValue)
//	binData = append(binData, binValue...)
//	return binData, bitEnd, err
// }

// func OverloadStopIEsValue (value ngapType.OverloadStopIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
//	var err error
//	switch value.Present {
//	case ngapType.OverloadStopIEsTypeValuePresentNothing:
//		if value.Nothing == nil {
//			return binData, bitEnd, errors.New("OverloadStopIEsValue: Nothing: NIL")
//		}
//		binData, bitEnd, err = Nothing(*value.Nothing, binData, bitEnd)
//		if err != nil {
//			return binData, bitEnd, errors.New("OverloadStopIEsValue: " + err.Error())
//		}
//	default:
//		return binData, bitEnd, errors.New("OverloadStopIEsValue: Present: INVALID")
//	}
//	return binData, bitEnd, err
// }

func Paging(value ngapType.Paging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPagingIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Paging: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPagingIEs(value ngapType.ProtocolIEContainerPagingIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PagingIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PagingIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PagingIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PagingIEsValue(value ngapType.PagingIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PagingIEsTypeValuePresentUEPagingIdentity:
		if value.UEPagingIdentity == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: UEPagingIdentity: NIL")
		}
		binData, bitEnd, err = UEPagingIdentity(*value.UEPagingIdentity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentPagingDRX:
		if value.PagingDRX == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: PagingDRX: NIL")
		}
		binData, bitEnd, err = PagingDRX(*value.PagingDRX, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentTAIListForPaging:
		if value.TAIListForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: TAIListForPaging: NIL")
		}
		binData, bitEnd, err = TAIListForPaging(*value.TAIListForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentPagingPriority:
		if value.PagingPriority == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: PagingPriority: NIL")
		}
		binData, bitEnd, err = PagingPriority(*value.PagingPriority, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentPagingOrigin:
		if value.PagingOrigin == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: PagingOrigin: NIL")
		}
		binData, bitEnd, err = PagingOrigin(*value.PagingOrigin, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	case ngapType.PagingIEsTypeValuePresentAssistanceDataForPaging:
		if value.AssistanceDataForPaging == nil {
			return binData, bitEnd, errors.New("PagingIEsValue: AssistanceDataForPaging: NIL")
		}
		binData, bitEnd, err = AssistanceDataForPaging(*value.AssistanceDataForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PagingIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PagingIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceNotify(value ngapType.PDUSessionResourceNotify, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceNotifyIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotify: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceNotifyIEs(value ngapType.ProtocolIEContainerPDUSessionResourceNotifyIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceNotifyIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceNotifyIEsValue(value ngapType.PDUSessionResourceNotifyIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceNotifyList:
		if value.PDUSessionResourceNotifyList == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: PDUSessionResourceNotifyList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceNotifyList(*value.PDUSessionResourceNotifyList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceReleasedListNot:
		if value.PDUSessionResourceReleasedListNot == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: PDUSessionResourceReleasedListNot: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListNot(*value.PDUSessionResourceReleasedListNot, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceNotifyIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceNotifyIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PrivateMessage(value ngapType.PrivateMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	// binData, bitEnd, err = ProtocolIEContainerPrivateMessageIEs(value.ProtocolIEs, binData, bitEnd)
	// if err != nil {
	//	return binData, bitEnd, errors.New("PrivateMessage: " + err.Error())
	// }
	return binData, bitEnd, err
}

// func ProtocolIEContainerPrivateMessageIEs (value ngapType.ProtocolIEContainerPrivateMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
//	var err error
//	numIEs := len(value.List)
//	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
//	for i:=0; i<numIEs; i++ {
//		var binIEs []byte
//		binIEs, bitEnd, err = PrivateMessageIEs(value.List[i], binIEs, bitEnd)
//		if err != nil {
//			return binData, bitEnd, errors.New("ProtocolIEContainerPrivateMessageIEs: ["+strconv.Itoa(i)+"]: " + err.Error())
//		}
//		binData = append(binData, binIEs...)
//	}
//	return binData, bitEnd, err
// }

// func PrivateMessageIEs (value ngapType.PrivateMessageIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
//	var err error
//	switch value.ProtocolIEID.Value {
//	case ngapType.ProtocolIEIDNothing:
//		if value.TypeValue.Present != ngapType.PrivateMessageIEsTypeValuePresentNothing {
//			return binData, bitEnd, errors.New("PrivateMessageIEs: ProtocolIEID: INVALID")
//		}
//	default:
//		return binData, bitEnd, errors.New("PrivateMessageIEs: ProtocolIEID: INVALID")
//	}
//	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
//	if err != nil {
//		return binData, bitEnd, errors.New("PrivateMessageIEs: " + err.Error())
//	}
//	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
//	if err != nil {
//		return binData, bitEnd, errors.New("PrivateMessageIEs: " + err.Error())
//	}
//	var binValue []byte
//	binValue, bitEnd, err = PrivateMessageIEsValue(value.Value, binValue, 8)
//	if err != nil {
//		return binData, bitEnd, errors.New("PrivateMessageIEs: " + err.Error())
//	}
//	binValue = EncodeLengthValue(binValue)
//	binData = append(binData, binValue...)
//	return binData, bitEnd, err
// }

// func PrivateMessageIEsValue (value ngapType.PrivateMessageIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
//	var err error
//	switch value.Present {
//	case ngapType.PrivateMessageIEsTypeValuePresentNothing:
//		if value.Nothing == nil {
//			return binData, bitEnd, errors.New("PrivateMessageIEsValue: Nothing: NIL")
//		}
//		binData, bitEnd, err = Nothing(*value.Nothing, binData, bitEnd)
//		if err != nil {
//			return binData, bitEnd, errors.New("PrivateMessageIEsValue: " + err.Error())
//		}
//	default:
//		return binData, bitEnd, errors.New("PrivateMessageIEsValue: Present: INVALID")
//	}
//	return binData, bitEnd, err
// }

func PWSFailureIndication(value ngapType.PWSFailureIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPWSFailureIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSFailureIndicationIEs(value ngapType.ProtocolIEContainerPWSFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSFailureIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PWSFailureIndicationIEsValue(value ngapType.PWSFailureIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSFailureIndicationIEsTypeValuePresentPWSFailedCellIDList:
		if value.PWSFailedCellIDList == nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsValue: PWSFailedCellIDList: NIL")
		}
		binData, bitEnd, err = PWSFailedCellIDList(*value.PWSFailedCellIDList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsValue: " + err.Error())
		}
	case ngapType.PWSFailureIndicationIEsTypeValuePresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSFailureIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSFailureIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSRestartIndication(value ngapType.PWSRestartIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPWSRestartIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSRestartIndicationIEs(value ngapType.ProtocolIEContainerPWSRestartIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSRestartIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSRestartIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PWSRestartIndicationIEsValue(value ngapType.PWSRestartIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSRestartIndicationIEsTypeValuePresentCellIDListForRestart:
		if value.CellIDListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: CellIDListForRestart: NIL")
		}
		binData, bitEnd, err = CellIDListForRestart(*value.CellIDListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsTypeValuePresentGlobalRANNodeID:
		if value.GlobalRANNodeID == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: GlobalRANNodeID: NIL")
		}
		binData, bitEnd, err = GlobalRANNodeID(*value.GlobalRANNodeID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsTypeValuePresentTAIListForRestart:
		if value.TAIListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: TAIListForRestart: NIL")
		}
		binData, bitEnd, err = TAIListForRestart(*value.TAIListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: " + err.Error())
		}
	case ngapType.PWSRestartIndicationIEsTypeValuePresentEmergencyAreaIDListForRestart:
		if value.EmergencyAreaIDListForRestart == nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: EmergencyAreaIDListForRestart: NIL")
		}
		binData, bitEnd, err = EmergencyAreaIDListForRestart(*value.EmergencyAreaIDListForRestart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSRestartIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RerouteNASRequest(value ngapType.RerouteNASRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRerouteNASRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRerouteNASRequestIEs(value ngapType.ProtocolIEContainerRerouteNASRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RerouteNASRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RerouteNASRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func RerouteNASRequestIEsValue(value ngapType.RerouteNASRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RerouteNASRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentNGAPMessage:
		if value.NGAPMessage == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: NGAPMessage: NIL")
		}
		binData, bitEnd, err = NGAPMessage(*value.NGAPMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentAMFSetID:
		if value.AMFSetID == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: AMFSetID: NIL")
		}
		binData, bitEnd, err = AMFSetID(*value.AMFSetID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: " + err.Error())
		}
	case ngapType.RerouteNASRequestIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RerouteNASRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RRCInactiveTransitionReport(value ngapType.RRCInactiveTransitionReport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRRCInactiveTransitionReportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRRCInactiveTransitionReportIEs(value ngapType.ProtocolIEContainerRRCInactiveTransitionReportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RRCInactiveTransitionReportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func RRCInactiveTransitionReportIEsValue(value ngapType.RRCInactiveTransitionReportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRRCState:
		if value.RRCState == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: RRCState: NIL")
		}
		binData, bitEnd, err = RRCState(*value.RRCState, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: " + err.Error())
		}
	case ngapType.RRCInactiveTransitionReportIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TraceFailureIndication(value ngapType.TraceFailureIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerTraceFailureIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerTraceFailureIndicationIEs(value ngapType.ProtocolIEContainerTraceFailureIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TraceFailureIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func TraceFailureIndicationIEsValue(value ngapType.TraceFailureIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.TraceFailureIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsTypeValuePresentNGRANTraceID:
		if value.NGRANTraceID == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: NGRANTraceID: NIL")
		}
		binData, bitEnd, err = NGRANTraceID(*value.NGRANTraceID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: " + err.Error())
		}
	case ngapType.TraceFailureIndicationIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("TraceFailureIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func TraceStart(value ngapType.TraceStart, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerTraceStartIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStart: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerTraceStartIEs(value ngapType.ProtocolIEContainerTraceStartIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStartIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStartIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = TraceStartIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("TraceStartIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func TraceStartIEsValue(value ngapType.TraceStartIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.TraceStartIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceStartIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsValue: " + err.Error())
		}
	case ngapType.TraceStartIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("TraceStartIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsValue: " + err.Error())
		}
	case ngapType.TraceStartIEsTypeValuePresentTraceActivation:
		if value.TraceActivation == nil {
			return binData, bitEnd, errors.New("TraceStartIEsValue: TraceActivation: NIL")
		}
		binData, bitEnd, err = TraceActivation(*value.TraceActivation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("TraceStartIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("TraceStartIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextReleaseRequest(value ngapType.UEContextReleaseRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextReleaseRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextReleaseRequestIEs(value ngapType.ProtocolIEContainerUEContextReleaseRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextReleaseRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UEContextReleaseRequestIEsValue(value ngapType.UEContextReleaseRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentPDUSessionResourceListCxtRelReq:
		if value.PDUSessionResourceListCxtRelReq == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: PDUSessionResourceListCxtRelReq: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListCxtRelReq(*value.PDUSessionResourceListCxtRelReq, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseRequestIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityInfoIndication(value ngapType.UERadioCapabilityInfoIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUERadioCapabilityInfoIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndication: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUERadioCapabilityInfoIndicationIEs(value ngapType.ProtocolIEContainerUERadioCapabilityInfoIndicationIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityInfoIndicationIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UERadioCapabilityInfoIndicationIEsValue(value ngapType.UERadioCapabilityInfoIndicationIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapability:
		if value.UERadioCapability == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: UERadioCapability: NIL")
		}
		binData, bitEnd, err = UERadioCapability(*value.UERadioCapability, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapabilityForPaging:
		if value.UERadioCapabilityForPaging == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: UERadioCapabilityForPaging: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityForPaging(*value.UERadioCapabilityForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityInfoIndicationIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UETNLABindingReleaseRequest(value ngapType.UETNLABindingReleaseRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUETNLABindingReleaseRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequest: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUETNLABindingReleaseRequestIEs(value ngapType.ProtocolIEContainerUETNLABindingReleaseRequestIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UETNLABindingReleaseRequestIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UETNLABindingReleaseRequestIEsValue(value ngapType.UETNLABindingReleaseRequestIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsValue: " + err.Error())
		}
	case ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequestIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkNASTransport(value ngapType.UplinkNASTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkNASTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkNASTransportIEs(value ngapType.ProtocolIEContainerUplinkNASTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkNASTransportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNASTransportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UplinkNASTransportIEsValue(value ngapType.UplinkNASTransportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkNASTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsTypeValuePresentNASPDU:
		if value.NASPDU == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: NASPDU: NIL")
		}
		binData, bitEnd, err = NASPDU(*value.NASPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkNASTransportIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkNASTransportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkNonUEAssociatedNRPPaTransport(value ngapType.UplinkNonUEAssociatedNRPPaTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkNonUEAssociatedNRPPaTransportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UplinkNonUEAssociatedNRPPaTransportIEsValue(value ngapType.UplinkNonUEAssociatedNRPPaTransportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkRANConfigurationTransfer(value ngapType.UplinkRANConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkRANConfigurationTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkRANConfigurationTransferIEs(value ngapType.ProtocolIEContainerUplinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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

func UplinkRANConfigurationTransferIEs(value ngapType.UplinkRANConfigurationTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProtocolIEID.Value {
	case ngapType.ProtocolIEIDSONConfigurationTransferUL:
		if value.TypeValue.Present != ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferUL {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkRANConfigurationTransferIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UplinkRANConfigurationTransferIEsValue(value ngapType.UplinkRANConfigurationTransferIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferUL:
		if value.SONConfigurationTransferUL == nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsValue: SONConfigurationTransferUL: NIL")
		}
		binData, bitEnd, err = SONConfigurationTransferUL(*value.SONConfigurationTransferUL, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANConfigurationTransferIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkRANStatusTransfer(value ngapType.UplinkRANStatusTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkRANStatusTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkRANStatusTransferIEs(value ngapType.ProtocolIEContainerUplinkRANStatusTransferIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkRANStatusTransferIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UplinkRANStatusTransferIEsValue(value ngapType.UplinkRANStatusTransferIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: " + err.Error())
		}
	case ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: " + err.Error())
		}
	case ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer:
		if value.RANStatusTransferTransparentContainer == nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: RANStatusTransferTransparentContainer: NIL")
		}
		binData, bitEnd, err = RANStatusTransferTransparentContainer(*value.RANStatusTransferTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkRANStatusTransferIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UplinkUEAssociatedNRPPaTransport(value ngapType.UplinkUEAssociatedNRPPaTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransport: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs(value ngapType.ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UplinkUEAssociatedNRPPaTransportIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UplinkUEAssociatedNRPPaTransportIEsValue(value ngapType.UplinkUEAssociatedNRPPaTransportIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID:
		if value.RoutingID == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: RoutingID: NIL")
		}
		binData, bitEnd, err = RoutingID(*value.RoutingID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	case ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU:
		if value.NRPPaPDU == nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: NRPPaPDU: NIL")
		}
		binData, bitEnd, err = NRPPaPDU(*value.NRPPaPDU, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransportIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFConfigurationUpdateAcknowledge(value ngapType.AMFConfigurationUpdateAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledge: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs(value ngapType.ProtocolIEContainerAMFConfigurationUpdateAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFConfigurationUpdateAcknowledgeIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func AMFConfigurationUpdateAcknowledgeIEsValue(value ngapType.AMFConfigurationUpdateAcknowledgeIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationSetupList:
		if value.AMFTNLAssociationSetupList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: AMFTNLAssociationSetupList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationSetupList(*value.AMFTNLAssociationSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationFailedToSetupList:
		if value.AMFTNLAssociationFailedToSetupList == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: AMFTNLAssociationFailedToSetupList: NIL")
		}
		binData, bitEnd, err = AMFTNLAssociationFailedToSetupList(*value.AMFTNLAssociationFailedToSetupList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateAcknowledgeIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCancelAcknowledge(value ngapType.HandoverCancelAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverCancelAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledge: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverCancelAcknowledgeIEs(value ngapType.ProtocolIEContainerHandoverCancelAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCancelAcknowledgeIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverCancelAcknowledgeIEsValue(value ngapType.HandoverCancelAcknowledgeIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCancelAcknowledgeIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverCommand(value ngapType.HandoverCommand, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverCommandIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommand: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverCommandIEs(value ngapType.ProtocolIEContainerHandoverCommandIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverCommandIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommandIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverCommandIEsValue(value ngapType.HandoverCommandIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverCommandIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentHandoverType:
		if value.HandoverType == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: HandoverType: NIL")
		}
		binData, bitEnd, err = HandoverType(*value.HandoverType, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentNASSecurityParametersFromNGRAN:
		if value.NASSecurityParametersFromNGRAN == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: NASSecurityParametersFromNGRAN: NIL")
		}
		binData, bitEnd, err = NASSecurityParametersFromNGRAN(*value.NASSecurityParametersFromNGRAN, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceHandoverList:
		if value.PDUSessionResourceHandoverList == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: PDUSessionResourceHandoverList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceHandoverList(*value.PDUSessionResourceHandoverList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceToReleaseListHOCmd:
		if value.PDUSessionResourceToReleaseListHOCmd == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: PDUSessionResourceToReleaseListHOCmd: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceToReleaseListHOCmd(*value.PDUSessionResourceToReleaseListHOCmd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentTargetToSourceTransparentContainer:
		if value.TargetToSourceTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: TargetToSourceTransparentContainer: NIL")
		}
		binData, bitEnd, err = TargetToSourceTransparentContainer(*value.TargetToSourceTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	case ngapType.HandoverCommandIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverCommandIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverCommandIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverRequestAcknowledge(value ngapType.HandoverRequestAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverRequestAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledge: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverRequestAcknowledgeIEs(value ngapType.ProtocolIEContainerHandoverRequestAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverRequestAcknowledgeIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverRequestAcknowledgeIEsValue(value ngapType.HandoverRequestAcknowledgeIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceAdmittedList:
		if value.PDUSessionResourceAdmittedList == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: PDUSessionResourceAdmittedList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceAdmittedList(*value.PDUSessionResourceAdmittedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceFailedToSetupListHOAck:
		if value.PDUSessionResourceFailedToSetupListHOAck == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: PDUSessionResourceFailedToSetupListHOAck: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListHOAck(*value.PDUSessionResourceFailedToSetupListHOAck, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentTargetToSourceTransparentContainer:
		if value.TargetToSourceTransparentContainer == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: TargetToSourceTransparentContainer: NIL")
		}
		binData, bitEnd, err = TargetToSourceTransparentContainer(*value.TargetToSourceTransparentContainer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverRequestAcknowledgeIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialContextSetupResponse(value ngapType.InitialContextSetupResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerInitialContextSetupResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialContextSetupResponseIEs(value ngapType.ProtocolIEContainerInitialContextSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialContextSetupResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func InitialContextSetupResponseIEsValue(value ngapType.InitialContextSetupResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListCxtRes:
		if value.PDUSessionResourceSetupListCxtRes == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: PDUSessionResourceSetupListCxtRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListCxtRes(*value.PDUSessionResourceSetupListCxtRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtRes:
		if value.PDUSessionResourceFailedToSetupListCxtRes == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: PDUSessionResourceFailedToSetupListCxtRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListCxtRes(*value.PDUSessionResourceFailedToSetupListCxtRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGResetAcknowledge(value ngapType.NGResetAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGResetAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledge: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGResetAcknowledgeIEs(value ngapType.ProtocolIEContainerNGResetAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGResetAcknowledgeIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func NGResetAcknowledgeIEsValue(value ngapType.NGResetAcknowledgeIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGResetAcknowledgeIEsTypeValuePresentUEAssociatedLogicalNGConnectionList:
		if value.UEAssociatedLogicalNGConnectionList == nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsValue: UEAssociatedLogicalNGConnectionList: NIL")
		}
		binData, bitEnd, err = UEAssociatedLogicalNGConnectionList(*value.UEAssociatedLogicalNGConnectionList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.NGResetAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGResetAcknowledgeIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGResetAcknowledgeIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGSetupResponse(value ngapType.NGSetupResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGSetupResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGSetupResponseIEs(value ngapType.ProtocolIEContainerNGSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	default:
		return binData, bitEnd, errors.New("NGSetupResponseIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGSetupResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func NGSetupResponseIEsValue(value ngapType.NGSetupResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGSetupResponseIEsTypeValuePresentAMFName:
		if value.AMFName == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: AMFName: NIL")
		}
		binData, bitEnd, err = AMFName(*value.AMFName, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentServedGUAMIList:
		if value.ServedGUAMIList == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: ServedGUAMIList: NIL")
		}
		binData, bitEnd, err = ServedGUAMIList(*value.ServedGUAMIList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentRelativeAMFCapacity:
		if value.RelativeAMFCapacity == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: RelativeAMFCapacity: NIL")
		}
		binData, bitEnd, err = RelativeAMFCapacity(*value.RelativeAMFCapacity, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentPLMNSupportList:
		if value.PLMNSupportList == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: PLMNSupportList: NIL")
		}
		binData, bitEnd, err = PLMNSupportList(*value.PLMNSupportList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.NGSetupResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGSetupResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestAcknowledge(value ngapType.PathSwitchRequestAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPathSwitchRequestAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledge: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPathSwitchRequestAcknowledgeIEs(value ngapType.ProtocolIEContainerPathSwitchRequestAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	case ngapType.ProtocolIEIDCoreNetworkAssistanceInformation:
		if value.TypeValue.Present != ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCoreNetworkAssistanceInformation {
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
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: ProtocolIEID: INVALID")
	}
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestAcknowledgeIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PathSwitchRequestAcknowledgeIEsValue(value ngapType.PathSwitchRequestAcknowledgeIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentUESecurityCapabilities:
		if value.UESecurityCapabilities == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: UESecurityCapabilities: NIL")
		}
		binData, bitEnd, err = UESecurityCapabilities(*value.UESecurityCapabilities, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSecurityContext:
		if value.SecurityContext == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: SecurityContext: NIL")
		}
		binData, bitEnd, err = SecurityContext(*value.SecurityContext, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentNewSecurityContextInd:
		if value.NewSecurityContextInd == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: NewSecurityContextInd: NIL")
		}
		binData, bitEnd, err = NewSecurityContextInd(*value.NewSecurityContextInd, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceSwitchedList:
		if value.PDUSessionResourceSwitchedList == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: PDUSessionResourceSwitchedList: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSwitchedList(*value.PDUSessionResourceSwitchedList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceReleasedListPSAck:
		if value.PDUSessionResourceReleasedListPSAck == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: PDUSessionResourceReleasedListPSAck: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListPSAck(*value.PDUSessionResourceReleasedListPSAck, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAllowedNSSAI:
		if value.AllowedNSSAI == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: AllowedNSSAI: NIL")
		}
		binData, bitEnd, err = AllowedNSSAI(*value.AllowedNSSAI, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCoreNetworkAssistanceInformation:
		if value.CoreNetworkAssistanceInformation == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: CoreNetworkAssistanceInformation: NIL")
		}
		binData, bitEnd, err = CoreNetworkAssistanceInformation(*value.CoreNetworkAssistanceInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRRCInactiveTransitionReportRequest:
		if value.RRCInactiveTransitionReportRequest == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: RRCInactiveTransitionReportRequest: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReportRequest(*value.RRCInactiveTransitionReportRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestAcknowledgeIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyResponse(value ngapType.PDUSessionResourceModifyResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyResponseIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceModifyResponseIEsValue(value ngapType.PDUSessionResourceModifyResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceModifyListModRes:
		if value.PDUSessionResourceModifyListModRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: PDUSessionResourceModifyListModRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModRes(*value.PDUSessionResourceModifyListModRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceFailedToModifyListModRes:
		if value.PDUSessionResourceFailedToModifyListModRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: PDUSessionResourceFailedToModifyListModRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToModifyListModRes(*value.PDUSessionResourceFailedToModifyListModRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceModifyConfirm(value ngapType.PDUSessionResourceModifyConfirm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyConfirmIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirm: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceModifyConfirmIEs(value ngapType.ProtocolIEContainerPDUSessionResourceModifyConfirmIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceModifyConfirmIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceModifyConfirmIEsValue(value ngapType.PDUSessionResourceModifyConfirmIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceModifyListModCfm:
		if value.PDUSessionResourceModifyListModCfm == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: PDUSessionResourceModifyListModCfm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyListModCfm(*value.PDUSessionResourceModifyListModCfm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceFailedToModifyListModCfm:
		if value.PDUSessionResourceFailedToModifyListModCfm == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: PDUSessionResourceFailedToModifyListModCfm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToModifyListModCfm(*value.PDUSessionResourceFailedToModifyListModCfm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirmIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseResponse(value ngapType.PDUSessionResourceReleaseResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceReleaseResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceReleaseResponseIEs(value ngapType.ProtocolIEContainerPDUSessionResourceReleaseResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceReleaseResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceReleaseResponseIEsValue(value ngapType.PDUSessionResourceReleaseResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentPDUSessionResourceReleasedListRelRes:
		if value.PDUSessionResourceReleasedListRelRes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: PDUSessionResourceReleasedListRelRes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListRelRes(*value.PDUSessionResourceReleasedListRelRes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PDUSessionResourceSetupResponse(value ngapType.PDUSessionResourceSetupResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceSetupResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPDUSessionResourceSetupResponseIEs(value ngapType.ProtocolIEContainerPDUSessionResourceSetupResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PDUSessionResourceSetupResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PDUSessionResourceSetupResponseIEsValue(value ngapType.PDUSessionResourceSetupResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListSURes:
		if value.PDUSessionResourceSetupListSURes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: PDUSessionResourceSetupListSURes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupListSURes(*value.PDUSessionResourceSetupListSURes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListSURes:
		if value.PDUSessionResourceFailedToSetupListSURes == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: PDUSessionResourceFailedToSetupListSURes: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListSURes(*value.PDUSessionResourceFailedToSetupListSURes, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: " + err.Error())
		}
	case ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PDUSessionResourceSetupResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PWSCancelResponse(value ngapType.PWSCancelResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPWSCancelResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPWSCancelResponseIEs(value ngapType.ProtocolIEContainerPWSCancelResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PWSCancelResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PWSCancelResponseIEsValue(value ngapType.PWSCancelResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PWSCancelResponseIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsTypeValuePresentBroadcastCancelledAreaList:
		if value.BroadcastCancelledAreaList == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: BroadcastCancelledAreaList: NIL")
		}
		binData, bitEnd, err = BroadcastCancelledAreaList(*value.BroadcastCancelledAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: " + err.Error())
		}
	case ngapType.PWSCancelResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PWSCancelResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANConfigurationUpdateAcknowledge(value ngapType.RANConfigurationUpdateAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledge: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs(value ngapType.ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANConfigurationUpdateAcknowledgeIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func RANConfigurationUpdateAcknowledgeIEsValue(value ngapType.RANConfigurationUpdateAcknowledgeIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RANConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledgeIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextModificationResponse(value ngapType.UEContextModificationResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextModificationResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextModificationResponseIEs(value ngapType.ProtocolIEContainerUEContextModificationResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextModificationResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UEContextModificationResponseIEsValue(value ngapType.UEContextModificationResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextModificationResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentRRCState:
		if value.RRCState == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: RRCState: NIL")
		}
		binData, bitEnd, err = RRCState(*value.RRCState, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextReleaseComplete(value ngapType.UEContextReleaseComplete, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextReleaseCompleteIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseComplete: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextReleaseCompleteIEs(value ngapType.ProtocolIEContainerUEContextReleaseCompleteIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextReleaseCompleteIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UEContextReleaseCompleteIEsValue(value ngapType.UEContextReleaseCompleteIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentUserLocationInformation:
		if value.UserLocationInformation == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: UserLocationInformation: NIL")
		}
		binData, bitEnd, err = UserLocationInformation(*value.UserLocationInformation, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentInfoOnRecommendedCellsAndRANNodesForPaging:
		if value.InfoOnRecommendedCellsAndRANNodesForPaging == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: InfoOnRecommendedCellsAndRANNodesForPaging: NIL")
		}
		binData, bitEnd, err = InfoOnRecommendedCellsAndRANNodesForPaging(*value.InfoOnRecommendedCellsAndRANNodesForPaging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentPDUSessionResourceListCxtRelCpl:
		if value.PDUSessionResourceListCxtRelCpl == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: PDUSessionResourceListCxtRelCpl: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceListCxtRelCpl(*value.PDUSessionResourceListCxtRelCpl, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: " + err.Error())
		}
	case ngapType.UEContextReleaseCompleteIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextReleaseCompleteIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UERadioCapabilityCheckResponse(value ngapType.UERadioCapabilityCheckResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUERadioCapabilityCheckResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUERadioCapabilityCheckResponseIEs(value ngapType.ProtocolIEContainerUERadioCapabilityCheckResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UERadioCapabilityCheckResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UERadioCapabilityCheckResponseIEsValue(value ngapType.UERadioCapabilityCheckResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentIMSVoiceSupportIndicator:
		if value.IMSVoiceSupportIndicator == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: IMSVoiceSupportIndicator: NIL")
		}
		binData, bitEnd, err = IMSVoiceSupportIndicator(*value.IMSVoiceSupportIndicator, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: " + err.Error())
		}
	case ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UERadioCapabilityCheckResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func WriteReplaceWarningResponse(value ngapType.WriteReplaceWarningResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerWriteReplaceWarningResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponse: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerWriteReplaceWarningResponseIEs(value ngapType.ProtocolIEContainerWriteReplaceWarningResponseIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = WriteReplaceWarningResponseIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func WriteReplaceWarningResponseIEsValue(value ngapType.WriteReplaceWarningResponseIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentMessageIdentifier:
		if value.MessageIdentifier == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: MessageIdentifier: NIL")
		}
		binData, bitEnd, err = MessageIdentifier(*value.MessageIdentifier, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentSerialNumber:
		if value.SerialNumber == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: SerialNumber: NIL")
		}
		binData, bitEnd, err = SerialNumber(*value.SerialNumber, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentBroadcastCompletedAreaList:
		if value.BroadcastCompletedAreaList == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: BroadcastCompletedAreaList: NIL")
		}
		binData, bitEnd, err = BroadcastCompletedAreaList(*value.BroadcastCompletedAreaList, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: " + err.Error())
		}
	case ngapType.WriteReplaceWarningResponseIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("WriteReplaceWarningResponseIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func AMFConfigurationUpdateFailure(value ngapType.AMFConfigurationUpdateFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerAMFConfigurationUpdateFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerAMFConfigurationUpdateFailureIEs(value ngapType.ProtocolIEContainerAMFConfigurationUpdateFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = AMFConfigurationUpdateFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func AMFConfigurationUpdateFailureIEsValue(value ngapType.AMFConfigurationUpdateFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: " + err.Error())
		}
	case ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverPreparationFailure(value ngapType.HandoverPreparationFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverPreparationFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverPreparationFailureIEs(value ngapType.ProtocolIEContainerHandoverPreparationFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverPreparationFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverPreparationFailureIEsValue(value ngapType.HandoverPreparationFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: " + err.Error())
		}
	case ngapType.HandoverPreparationFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverPreparationFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func HandoverFailure(value ngapType.HandoverFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerHandoverFailureIEs(value ngapType.ProtocolIEContainerHandoverFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = HandoverFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func HandoverFailureIEsValue(value ngapType.HandoverFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.HandoverFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsValue: " + err.Error())
		}
	case ngapType.HandoverFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsValue: " + err.Error())
		}
	case ngapType.HandoverFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("HandoverFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("HandoverFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitialContextSetupFailure(value ngapType.InitialContextSetupFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerInitialContextSetupFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerInitialContextSetupFailureIEs(value ngapType.ProtocolIEContainerInitialContextSetupFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitialContextSetupFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func InitialContextSetupFailureIEsValue(value ngapType.InitialContextSetupFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtFail:
		if value.PDUSessionResourceFailedToSetupListCxtFail == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: PDUSessionResourceFailedToSetupListCxtFail: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceFailedToSetupListCxtFail(*value.PDUSessionResourceFailedToSetupListCxtFail, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: " + err.Error())
		}
	case ngapType.InitialContextSetupFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitialContextSetupFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func NGSetupFailure(value ngapType.NGSetupFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGSetupFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerNGSetupFailureIEs(value ngapType.ProtocolIEContainerNGSetupFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = NGSetupFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func NGSetupFailureIEsValue(value ngapType.NGSetupFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.NGSetupFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsValue: " + err.Error())
		}
	case ngapType.NGSetupFailureIEsTypeValuePresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsValue: " + err.Error())
		}
	case ngapType.NGSetupFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGSetupFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGSetupFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func PathSwitchRequestFailure(value ngapType.PathSwitchRequestFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPathSwitchRequestFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerPathSwitchRequestFailureIEs(value ngapType.ProtocolIEContainerPathSwitchRequestFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = PathSwitchRequestFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func PathSwitchRequestFailureIEsValue(value ngapType.PathSwitchRequestFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentPDUSessionResourceReleasedListPSFail:
		if value.PDUSessionResourceReleasedListPSFail == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: PDUSessionResourceReleasedListPSFail: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleasedListPSFail(*value.PDUSessionResourceReleasedListPSFail, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: " + err.Error())
		}
	case ngapType.PathSwitchRequestFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("PathSwitchRequestFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func RANConfigurationUpdateFailure(value ngapType.RANConfigurationUpdateFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRANConfigurationUpdateFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerRANConfigurationUpdateFailureIEs(value ngapType.ProtocolIEContainerRANConfigurationUpdateFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = RANConfigurationUpdateFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func RANConfigurationUpdateFailureIEsValue(value ngapType.RANConfigurationUpdateFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentTimeToWait:
		if value.TimeToWait == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: TimeToWait: NIL")
		}
		binData, bitEnd, err = TimeToWait(*value.TimeToWait, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: " + err.Error())
		}
	case ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("RANConfigurationUpdateFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UEContextModificationFailure(value ngapType.UEContextModificationFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextModificationFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailure: " + err.Error())
	}
	return binData, bitEnd, err
}

func ProtocolIEContainerUEContextModificationFailureIEs(value ngapType.ProtocolIEContainerUEContextModificationFailureIEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	numIEs := len(value.List)
	binData, bitEnd = EncodeNumberProtocolIE(numIEs, binData, bitEnd)
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
	binData, bitEnd, err = ProtocolIEID(value.Id, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UEContextModificationFailureIEsValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationFailureIEs: " + err.Error())
	}
	binValue = EncodeLengthValue(binValue)
	binData = append(binData, binValue...)
	return binData, bitEnd, err
}

func UEContextModificationFailureIEsValue(value ngapType.UEContextModificationFailureIEsValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UEContextModificationFailureIEsTypeValuePresentAMFUENGAPID:
		if value.AMFUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: AMFUENGAPID: NIL")
		}
		binData, bitEnd, err = AMFUENGAPID(*value.AMFUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsTypeValuePresentRANUENGAPID:
		if value.RANUENGAPID == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: RANUENGAPID: NIL")
		}
		binData, bitEnd, err = RANUENGAPID(*value.RANUENGAPID, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsTypeValuePresentCause:
		if value.Cause == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: Cause: NIL")
		}
		binData, bitEnd, err = Cause(*value.Cause, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: " + err.Error())
		}
	case ngapType.UEContextModificationFailureIEsTypeValuePresentCriticalityDiagnostics:
		if value.CriticalityDiagnostics == nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: CriticalityDiagnostics: NIL")
		}
		binData, bitEnd, err = CriticalityDiagnostics(*value.CriticalityDiagnostics, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UEContextModificationFailureIEsValue: Present: INVALID")
	}
	return binData, bitEnd, err
}
