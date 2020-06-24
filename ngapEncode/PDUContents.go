package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func PDUSessionResourceSetupRequest(value ngapType.PDUSessionResourceSetupRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceSetupRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceSetupRequest: " + err.Error())
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

func PDUSessionResourceReleaseCommand(value ngapType.PDUSessionResourceReleaseCommand, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceReleaseCommandIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceReleaseCommand: " + err.Error())
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

func PDUSessionResourceModifyRequest(value ngapType.PDUSessionResourceModifyRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyRequest: " + err.Error())
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

func PDUSessionResourceNotify(value ngapType.PDUSessionResourceNotify, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceNotifyIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceNotify: " + err.Error())
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

func PDUSessionResourceModifyConfirm(value ngapType.PDUSessionResourceModifyConfirm, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPDUSessionResourceModifyConfirmIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PDUSessionResourceModifyConfirm: " + err.Error())
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

func InitialContextSetupResponse(value ngapType.InitialContextSetupResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerInitialContextSetupResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitialContextSetupResponse: " + err.Error())
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

func UEContextReleaseRequest(value ngapType.UEContextReleaseRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextReleaseRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseRequest: " + err.Error())
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

func UEContextReleaseComplete(value ngapType.UEContextReleaseComplete, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextReleaseCompleteIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextReleaseComplete: " + err.Error())
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

func UEContextModificationResponse(value ngapType.UEContextModificationResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUEContextModificationResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UEContextModificationResponse: " + err.Error())
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

func RRCInactiveTransitionReport(value ngapType.RRCInactiveTransitionReport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRRCInactiveTransitionReportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReport: " + err.Error())
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

func HandoverCommand(value ngapType.HandoverCommand, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverCommandIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCommand: " + err.Error())
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

func HandoverRequest(value ngapType.HandoverRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverRequest: " + err.Error())
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

func HandoverFailure(value ngapType.HandoverFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverFailure: " + err.Error())
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

func PathSwitchRequest(value ngapType.PathSwitchRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPathSwitchRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequest: " + err.Error())
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

func PathSwitchRequestFailure(value ngapType.PathSwitchRequestFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPathSwitchRequestFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PathSwitchRequestFailure: " + err.Error())
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

func HandoverCancelAcknowledge(value ngapType.HandoverCancelAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerHandoverCancelAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("HandoverCancelAcknowledge: " + err.Error())
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

func DownlinkRANStatusTransfer(value ngapType.DownlinkRANStatusTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkRANStatusTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANStatusTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func Paging(value ngapType.Paging, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPagingIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("Paging: " + err.Error())
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

func DownlinkNASTransport(value ngapType.DownlinkNASTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkNASTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkNASTransport: " + err.Error())
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

func NASNonDeliveryIndication(value ngapType.NASNonDeliveryIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNASNonDeliveryIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASNonDeliveryIndication: " + err.Error())
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

func NGSetupRequest(value ngapType.NGSetupRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGSetupRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupRequest: " + err.Error())
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

func NGSetupFailure(value ngapType.NGSetupFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGSetupFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGSetupFailure: " + err.Error())
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

func RANConfigurationUpdateAcknowledge(value ngapType.RANConfigurationUpdateAcknowledge, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerRANConfigurationUpdateAcknowledgeIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RANConfigurationUpdateAcknowledge: " + err.Error())
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

func AMFConfigurationUpdate(value ngapType.AMFConfigurationUpdate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerAMFConfigurationUpdateIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdate: " + err.Error())
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

func AMFConfigurationUpdateFailure(value ngapType.AMFConfigurationUpdateFailure, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerAMFConfigurationUpdateFailureIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("AMFConfigurationUpdateFailure: " + err.Error())
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

func NGReset(value ngapType.NGReset, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerNGResetIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NGReset: " + err.Error())
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

func ErrorIndication(value ngapType.ErrorIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerErrorIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ErrorIndication: " + err.Error())
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

func OverloadStop(value ngapType.OverloadStop, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerOverloadStopIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("OverloadStop: " + err.Error())
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

func DownlinkRANConfigurationTransfer(value ngapType.DownlinkRANConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkRANConfigurationTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRANConfigurationTransfer: " + err.Error())
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

func WriteReplaceWarningResponse(value ngapType.WriteReplaceWarningResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerWriteReplaceWarningResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("WriteReplaceWarningResponse: " + err.Error())
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

func PWSCancelResponse(value ngapType.PWSCancelResponse, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPWSCancelResponseIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSCancelResponse: " + err.Error())
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

func PWSFailureIndication(value ngapType.PWSFailureIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerPWSFailureIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PWSFailureIndication: " + err.Error())
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

func UplinkUEAssociatedNRPPaTransport(value ngapType.UplinkUEAssociatedNRPPaTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkUEAssociatedNRPPaTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkUEAssociatedNRPPaTransport: " + err.Error())
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

func UplinkNonUEAssociatedNRPPaTransport(value ngapType.UplinkNonUEAssociatedNRPPaTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkNonUEAssociatedNRPPaTransportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkNonUEAssociatedNRPPaTransport: " + err.Error())
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

func TraceFailureIndication(value ngapType.TraceFailureIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerTraceFailureIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TraceFailureIndication: " + err.Error())
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

func CellTrafficTrace(value ngapType.CellTrafficTrace, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerCellTrafficTraceIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CellTrafficTrace: " + err.Error())
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

func LocationReportingFailureIndication(value ngapType.LocationReportingFailureIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerLocationReportingFailureIndicationIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LocationReportingFailureIndication: " + err.Error())
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

func UETNLABindingReleaseRequest(value ngapType.UETNLABindingReleaseRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUETNLABindingReleaseRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UETNLABindingReleaseRequest: " + err.Error())
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

func UERadioCapabilityCheckRequest(value ngapType.UERadioCapabilityCheckRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUERadioCapabilityCheckRequestIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityCheckRequest: " + err.Error())
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

func PrivateMessage(value ngapType.PrivateMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = PrivateIEContainerPrivateMessageIEs(value.PrivateIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("PrivateMessage: " + err.Error())
	}
	return binData, bitEnd, err
}

func SecondaryRATDataUsageReport(value ngapType.SecondaryRATDataUsageReport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerSecondaryRATDataUsageReportIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SecondaryRATDataUsageReport: " + err.Error())
	}
	return binData, bitEnd, err
}

func UplinkRIMInformationTransfer(value ngapType.UplinkRIMInformationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerUplinkRIMInformationTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UplinkRIMInformationTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}

func DownlinkRIMInformationTransfer(value ngapType.DownlinkRIMInformationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = ProtocolIEContainerDownlinkRIMInformationTransferIEs(value.ProtocolIEs, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("DownlinkRIMInformationTransfer: " + err.Error())
	}
	return binData, bitEnd, err
}
