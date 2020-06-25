package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func NGAPPDU (value ngapType.NGAPPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)
	switch value.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		if value.InitiatingMessage == nil {
			return binData, bitEnd, errors.New("NGAPPDU: InitiatingMessage: NIL")
		}
		binData, bitEnd, err = InitiatingMessage(*value.InitiatingMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGAPPDU: " + err.Error())
		}
	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		if value.SuccessfulOutcome == nil {
			return binData, bitEnd, errors.New("NGAPPDU: SuccessfulOutcome: NIL")
		}
		binData, bitEnd, err = SuccessfulOutcome(*value.SuccessfulOutcome, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGAPPDU: " + err.Error())
		}
	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		if value.UnsuccessfulOutcome == nil {
			return binData, bitEnd, errors.New("NGAPPDU: UnsuccessfulOutcome: NIL")
		}
		binData, bitEnd, err = UnsuccessfulOutcome(*value.UnsuccessfulOutcome, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("NGAPPDU: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("NGAPPDU: Present: INVALID")
	}
	return binData, bitEnd, err
}

func InitiatingMessage (value ngapType.InitiatingMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProcedureCode.Value {
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		if value.Value.Present != ngapType.InitiatingMessagePresentAMFConfigurationUpdate {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeAMFStatusIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentAMFStatusIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeCellTrafficTrace:
		if value.Value.Present != ngapType.InitiatingMessagePresentCellTrafficTrace {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDeactivateTrace:
		if value.Value.Present != ngapType.InitiatingMessagePresentDeactivateTrace {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDownlinkNASTransport:
		if value.Value.Present != ngapType.InitiatingMessagePresentDownlinkNASTransport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport:
		if value.Value.Present != ngapType.InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDownlinkRANConfigurationTransfer:
		if value.Value.Present != ngapType.InitiatingMessagePresentDownlinkRANConfigurationTransfer {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDownlinkRANStatusTransfer:
		if value.Value.Present != ngapType.InitiatingMessagePresentDownlinkRANStatusTransfer {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDownlinkUEAssociatedNRPPaTransport:
		if value.Value.Present != ngapType.InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeErrorIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentErrorIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverCancel:
		if value.Value.Present != ngapType.InitiatingMessagePresentHandoverCancel {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverNotification:
		if value.Value.Present != ngapType.InitiatingMessagePresentHandoverNotify {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverPreparation:
		if value.Value.Present != ngapType.InitiatingMessagePresentHandoverRequired {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		if value.Value.Present != ngapType.InitiatingMessagePresentHandoverRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeInitialContextSetup:
		if value.Value.Present != ngapType.InitiatingMessagePresentInitialContextSetupRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeInitialUEMessage:
		if value.Value.Present != ngapType.InitiatingMessagePresentInitialUEMessage {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeLocationReport:
		if value.Value.Present != ngapType.InitiatingMessagePresentLocationReport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeLocationReportingControl:
		if value.Value.Present != ngapType.InitiatingMessagePresentLocationReportingControl {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeLocationReportingFailureIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentLocationReportingFailureIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeNASNonDeliveryIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentNASNonDeliveryIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeNGReset:
		if value.Value.Present != ngapType.InitiatingMessagePresentNGReset {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeNGSetup:
		if value.Value.Present != ngapType.InitiatingMessagePresentNGSetupRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeOverloadStart:
		if value.Value.Present != ngapType.InitiatingMessagePresentOverloadStart {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeOverloadStop:
		if value.Value.Present != ngapType.InitiatingMessagePresentOverloadStop {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePaging:
		if value.Value.Present != ngapType.InitiatingMessagePresentPaging {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePathSwitchRequest:
		if value.Value.Present != ngapType.InitiatingMessagePresentPathSwitchRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceModify:
		if value.Value.Present != ngapType.InitiatingMessagePresentPDUSessionResourceModifyRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentPDUSessionResourceModifyIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceNotify:
		if value.Value.Present != ngapType.InitiatingMessagePresentPDUSessionResourceNotify {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceRelease:
		if value.Value.Present != ngapType.InitiatingMessagePresentPDUSessionResourceReleaseCommand {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		if value.Value.Present != ngapType.InitiatingMessagePresentPDUSessionResourceSetupRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePrivateMessage:
		if value.Value.Present != ngapType.InitiatingMessagePresentPrivateMessage {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePWSCancel:
		if value.Value.Present != ngapType.InitiatingMessagePresentPWSCancelRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePWSFailureIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentPWSFailureIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePWSRestartIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentPWSRestartIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		if value.Value.Present != ngapType.InitiatingMessagePresentRANConfigurationUpdate {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeRerouteNASRequest:
		if value.Value.Present != ngapType.InitiatingMessagePresentRerouteNASRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeRRCInactiveTransitionReport:
		if value.Value.Present != ngapType.InitiatingMessagePresentRRCInactiveTransitionReport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeSecondaryRATDataUsageReport:
		if value.Value.Present != ngapType.InitiatingMessagePresentSecondaryRATDataUsageReport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeTraceFailureIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentTraceFailureIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeTraceStart:
		if value.Value.Present != ngapType.InitiatingMessagePresentTraceStart {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUEContextModification:
		if value.Value.Present != ngapType.InitiatingMessagePresentUEContextModificationRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUEContextRelease:
		if value.Value.Present != ngapType.InitiatingMessagePresentUEContextReleaseCommand {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUEContextReleaseRequest:
		if value.Value.Present != ngapType.InitiatingMessagePresentUEContextReleaseRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUERadioCapabilityCheck:
		if value.Value.Present != ngapType.InitiatingMessagePresentUERadioCapabilityCheckRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUERadioCapabilityInfoIndication:
		if value.Value.Present != ngapType.InitiatingMessagePresentUERadioCapabilityInfoIndication {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUETNLABindingRelease:
		if value.Value.Present != ngapType.InitiatingMessagePresentUETNLABindingReleaseRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUplinkNASTransport:
		if value.Value.Present != ngapType.InitiatingMessagePresentUplinkNASTransport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
		if value.Value.Present != ngapType.InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUplinkRANConfigurationTransfer:
		if value.Value.Present != ngapType.InitiatingMessagePresentUplinkRANConfigurationTransfer {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUplinkRANStatusTransfer:
		if value.Value.Present != ngapType.InitiatingMessagePresentUplinkRANStatusTransfer {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport:
		if value.Value.Present != ngapType.InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeWriteReplaceWarning:
		if value.Value.Present != ngapType.InitiatingMessagePresentWriteReplaceWarningRequest {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUplinkRIMInformationTransfer:
		if value.Value.Present != ngapType.InitiatingMessagePresentUplinkRIMInformationTransfer {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeDownlinkRIMInformationTransfer:
		if value.Value.Present != ngapType.InitiatingMessagePresentDownlinkRIMInformationTransfer {
			return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("InitiatingMessage: ProcedureCode: INVALID")
	}
	binData, bitEnd, err = ProcedureCode(value.ProcedureCode, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitiatingMessage: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("InitiatingMessage: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = InitiatingMessageValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("InitiatingMessage: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func InitiatingMessageValue (value ngapType.InitiatingMessageValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.InitiatingMessagePresentAMFConfigurationUpdate:
		if value.AMFConfigurationUpdate == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: AMFConfigurationUpdate: NIL")
		}
		binData, bitEnd, err = AMFConfigurationUpdate(*value.AMFConfigurationUpdate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentAMFStatusIndication:
		if value.AMFStatusIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: AMFStatusIndication: NIL")
		}
		binData, bitEnd, err = AMFStatusIndication(*value.AMFStatusIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentCellTrafficTrace:
		if value.CellTrafficTrace == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: CellTrafficTrace: NIL")
		}
		binData, bitEnd, err = CellTrafficTrace(*value.CellTrafficTrace, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDeactivateTrace:
		if value.DeactivateTrace == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DeactivateTrace: NIL")
		}
		binData, bitEnd, err = DeactivateTrace(*value.DeactivateTrace, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDownlinkNASTransport:
		if value.DownlinkNASTransport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DownlinkNASTransport: NIL")
		}
		binData, bitEnd, err = DownlinkNASTransport(*value.DownlinkNASTransport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport:
		if value.DownlinkNonUEAssociatedNRPPaTransport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DownlinkNonUEAssociatedNRPPaTransport: NIL")
		}
		binData, bitEnd, err = DownlinkNonUEAssociatedNRPPaTransport(*value.DownlinkNonUEAssociatedNRPPaTransport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDownlinkRANConfigurationTransfer:
		if value.DownlinkRANConfigurationTransfer == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DownlinkRANConfigurationTransfer: NIL")
		}
		binData, bitEnd, err = DownlinkRANConfigurationTransfer(*value.DownlinkRANConfigurationTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDownlinkRANStatusTransfer:
		if value.DownlinkRANStatusTransfer == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DownlinkRANStatusTransfer: NIL")
		}
		binData, bitEnd, err = DownlinkRANStatusTransfer(*value.DownlinkRANStatusTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport:
		if value.DownlinkUEAssociatedNRPPaTransport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DownlinkUEAssociatedNRPPaTransport: NIL")
		}
		binData, bitEnd, err = DownlinkUEAssociatedNRPPaTransport(*value.DownlinkUEAssociatedNRPPaTransport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentErrorIndication:
		if value.ErrorIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: ErrorIndication: NIL")
		}
		binData, bitEnd, err = ErrorIndication(*value.ErrorIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentHandoverCancel:
		if value.HandoverCancel == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: HandoverCancel: NIL")
		}
		binData, bitEnd, err = HandoverCancel(*value.HandoverCancel, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentHandoverNotify:
		if value.HandoverNotify == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: HandoverNotify: NIL")
		}
		binData, bitEnd, err = HandoverNotify(*value.HandoverNotify, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentHandoverRequired:
		if value.HandoverRequired == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: HandoverRequired: NIL")
		}
		binData, bitEnd, err = HandoverRequired(*value.HandoverRequired, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentHandoverRequest:
		if value.HandoverRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: HandoverRequest: NIL")
		}
		binData, bitEnd, err = HandoverRequest(*value.HandoverRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentInitialContextSetupRequest:
		if value.InitialContextSetupRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: InitialContextSetupRequest: NIL")
		}
		binData, bitEnd, err = InitialContextSetupRequest(*value.InitialContextSetupRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentInitialUEMessage:
		if value.InitialUEMessage == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: InitialUEMessage: NIL")
		}
		binData, bitEnd, err = InitialUEMessage(*value.InitialUEMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentLocationReport:
		if value.LocationReport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: LocationReport: NIL")
		}
		binData, bitEnd, err = LocationReport(*value.LocationReport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentLocationReportingControl:
		if value.LocationReportingControl == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: LocationReportingControl: NIL")
		}
		binData, bitEnd, err = LocationReportingControl(*value.LocationReportingControl, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentLocationReportingFailureIndication:
		if value.LocationReportingFailureIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: LocationReportingFailureIndication: NIL")
		}
		binData, bitEnd, err = LocationReportingFailureIndication(*value.LocationReportingFailureIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentNASNonDeliveryIndication:
		if value.NASNonDeliveryIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: NASNonDeliveryIndication: NIL")
		}
		binData, bitEnd, err = NASNonDeliveryIndication(*value.NASNonDeliveryIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentNGReset:
		if value.NGReset == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: NGReset: NIL")
		}
		binData, bitEnd, err = NGReset(*value.NGReset, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentNGSetupRequest:
		if value.NGSetupRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: NGSetupRequest: NIL")
		}
		binData, bitEnd, err = NGSetupRequest(*value.NGSetupRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentOverloadStart:
		if value.OverloadStart == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: OverloadStart: NIL")
		}
		binData, bitEnd, err = OverloadStart(*value.OverloadStart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentOverloadStop:
		if value.OverloadStop == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: OverloadStop: NIL")
		}
		binData, bitEnd, err = OverloadStop(*value.OverloadStop, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPaging:
		if value.Paging == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: Paging: NIL")
		}
		binData, bitEnd, err = Paging(*value.Paging, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPathSwitchRequest:
		if value.PathSwitchRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PathSwitchRequest: NIL")
		}
		binData, bitEnd, err = PathSwitchRequest(*value.PathSwitchRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPDUSessionResourceModifyRequest:
		if value.PDUSessionResourceModifyRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PDUSessionResourceModifyRequest: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyRequest(*value.PDUSessionResourceModifyRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPDUSessionResourceModifyIndication:
		if value.PDUSessionResourceModifyIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PDUSessionResourceModifyIndication: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyIndication(*value.PDUSessionResourceModifyIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPDUSessionResourceNotify:
		if value.PDUSessionResourceNotify == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PDUSessionResourceNotify: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceNotify(*value.PDUSessionResourceNotify, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPDUSessionResourceReleaseCommand:
		if value.PDUSessionResourceReleaseCommand == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PDUSessionResourceReleaseCommand: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleaseCommand(*value.PDUSessionResourceReleaseCommand, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPDUSessionResourceSetupRequest:
		if value.PDUSessionResourceSetupRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PDUSessionResourceSetupRequest: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupRequest(*value.PDUSessionResourceSetupRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPrivateMessage:
		if value.PrivateMessage == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PrivateMessage: NIL")
		}
		binData, bitEnd, err = PrivateMessage(*value.PrivateMessage, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPWSCancelRequest:
		if value.PWSCancelRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PWSCancelRequest: NIL")
		}
		binData, bitEnd, err = PWSCancelRequest(*value.PWSCancelRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPWSFailureIndication:
		if value.PWSFailureIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PWSFailureIndication: NIL")
		}
		binData, bitEnd, err = PWSFailureIndication(*value.PWSFailureIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentPWSRestartIndication:
		if value.PWSRestartIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: PWSRestartIndication: NIL")
		}
		binData, bitEnd, err = PWSRestartIndication(*value.PWSRestartIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentRANConfigurationUpdate:
		if value.RANConfigurationUpdate == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: RANConfigurationUpdate: NIL")
		}
		binData, bitEnd, err = RANConfigurationUpdate(*value.RANConfigurationUpdate, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentRerouteNASRequest:
		if value.RerouteNASRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: RerouteNASRequest: NIL")
		}
		binData, bitEnd, err = RerouteNASRequest(*value.RerouteNASRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentRRCInactiveTransitionReport:
		if value.RRCInactiveTransitionReport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: RRCInactiveTransitionReport: NIL")
		}
		binData, bitEnd, err = RRCInactiveTransitionReport(*value.RRCInactiveTransitionReport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentSecondaryRATDataUsageReport:
		if value.SecondaryRATDataUsageReport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: SecondaryRATDataUsageReport: NIL")
		}
		binData, bitEnd, err = SecondaryRATDataUsageReport(*value.SecondaryRATDataUsageReport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentTraceFailureIndication:
		if value.TraceFailureIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: TraceFailureIndication: NIL")
		}
		binData, bitEnd, err = TraceFailureIndication(*value.TraceFailureIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentTraceStart:
		if value.TraceStart == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: TraceStart: NIL")
		}
		binData, bitEnd, err = TraceStart(*value.TraceStart, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUEContextModificationRequest:
		if value.UEContextModificationRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UEContextModificationRequest: NIL")
		}
		binData, bitEnd, err = UEContextModificationRequest(*value.UEContextModificationRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUEContextReleaseCommand:
		if value.UEContextReleaseCommand == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UEContextReleaseCommand: NIL")
		}
		binData, bitEnd, err = UEContextReleaseCommand(*value.UEContextReleaseCommand, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUEContextReleaseRequest:
		if value.UEContextReleaseRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UEContextReleaseRequest: NIL")
		}
		binData, bitEnd, err = UEContextReleaseRequest(*value.UEContextReleaseRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUERadioCapabilityCheckRequest:
		if value.UERadioCapabilityCheckRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UERadioCapabilityCheckRequest: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityCheckRequest(*value.UERadioCapabilityCheckRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUERadioCapabilityInfoIndication:
		if value.UERadioCapabilityInfoIndication == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UERadioCapabilityInfoIndication: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityInfoIndication(*value.UERadioCapabilityInfoIndication, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUETNLABindingReleaseRequest:
		if value.UETNLABindingReleaseRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UETNLABindingReleaseRequest: NIL")
		}
		binData, bitEnd, err = UETNLABindingReleaseRequest(*value.UETNLABindingReleaseRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUplinkNASTransport:
		if value.UplinkNASTransport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UplinkNASTransport: NIL")
		}
		binData, bitEnd, err = UplinkNASTransport(*value.UplinkNASTransport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport:
		if value.UplinkNonUEAssociatedNRPPaTransport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UplinkNonUEAssociatedNRPPaTransport: NIL")
		}
		binData, bitEnd, err = UplinkNonUEAssociatedNRPPaTransport(*value.UplinkNonUEAssociatedNRPPaTransport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUplinkRANConfigurationTransfer:
		if value.UplinkRANConfigurationTransfer == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UplinkRANConfigurationTransfer: NIL")
		}
		binData, bitEnd, err = UplinkRANConfigurationTransfer(*value.UplinkRANConfigurationTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUplinkRANStatusTransfer:
		if value.UplinkRANStatusTransfer == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UplinkRANStatusTransfer: NIL")
		}
		binData, bitEnd, err = UplinkRANStatusTransfer(*value.UplinkRANStatusTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport:
		if value.UplinkUEAssociatedNRPPaTransport == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UplinkUEAssociatedNRPPaTransport: NIL")
		}
		binData, bitEnd, err = UplinkUEAssociatedNRPPaTransport(*value.UplinkUEAssociatedNRPPaTransport, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentWriteReplaceWarningRequest:
		if value.WriteReplaceWarningRequest == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: WriteReplaceWarningRequest: NIL")
		}
		binData, bitEnd, err = WriteReplaceWarningRequest(*value.WriteReplaceWarningRequest, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentUplinkRIMInformationTransfer:
		if value.UplinkRIMInformationTransfer == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: UplinkRIMInformationTransfer: NIL")
		}
		binData, bitEnd, err = UplinkRIMInformationTransfer(*value.UplinkRIMInformationTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	case ngapType.InitiatingMessagePresentDownlinkRIMInformationTransfer:
		if value.DownlinkRIMInformationTransfer == nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: DownlinkRIMInformationTransfer: NIL")
		}
		binData, bitEnd, err = DownlinkRIMInformationTransfer(*value.DownlinkRIMInformationTransfer, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("InitiatingMessageValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("InitiatingMessageValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func SuccessfulOutcome (value ngapType.SuccessfulOutcome, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProcedureCode.Value {
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentAMFConfigurationUpdateAcknowledge {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverCancel:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentHandoverCancelAcknowledge {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverPreparation:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentHandoverCommand {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentHandoverRequestAcknowledge {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeInitialContextSetup:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentInitialContextSetupResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeNGReset:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentNGResetAcknowledge {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeNGSetup:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentNGSetupResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePathSwitchRequest:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentPathSwitchRequestAcknowledge {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceModify:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyConfirm {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceRelease:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentPDUSessionResourceReleaseResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentPDUSessionResourceSetupResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePWSCancel:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentPWSCancelResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUEContextModification:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentUEContextModificationResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUEContextRelease:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentUEContextReleaseComplete {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUERadioCapabilityCheck:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentUERadioCapabilityCheckResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeWriteReplaceWarning:
		if value.Value.Present != ngapType.SuccessfulOutcomePresentWriteReplaceWarningResponse {
			return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("SuccessfulOutcome: ProcedureCode: INVALID")
	}
	binData, bitEnd, err = ProcedureCode(value.ProcedureCode, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SuccessfulOutcome: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SuccessfulOutcome: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = SuccessfulOutcomeValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("SuccessfulOutcome: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func SuccessfulOutcomeValue (value ngapType.SuccessfulOutcomeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.SuccessfulOutcomePresentAMFConfigurationUpdateAcknowledge:
		if value.AMFConfigurationUpdateAcknowledge == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: AMFConfigurationUpdateAcknowledge: NIL")
		}
		binData, bitEnd, err = AMFConfigurationUpdateAcknowledge(*value.AMFConfigurationUpdateAcknowledge, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentHandoverCancelAcknowledge:
		if value.HandoverCancelAcknowledge == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: HandoverCancelAcknowledge: NIL")
		}
		binData, bitEnd, err = HandoverCancelAcknowledge(*value.HandoverCancelAcknowledge, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentHandoverCommand:
		if value.HandoverCommand == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: HandoverCommand: NIL")
		}
		binData, bitEnd, err = HandoverCommand(*value.HandoverCommand, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentHandoverRequestAcknowledge:
		if value.HandoverRequestAcknowledge == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: HandoverRequestAcknowledge: NIL")
		}
		binData, bitEnd, err = HandoverRequestAcknowledge(*value.HandoverRequestAcknowledge, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentInitialContextSetupResponse:
		if value.InitialContextSetupResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: InitialContextSetupResponse: NIL")
		}
		binData, bitEnd, err = InitialContextSetupResponse(*value.InitialContextSetupResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentNGResetAcknowledge:
		if value.NGResetAcknowledge == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: NGResetAcknowledge: NIL")
		}
		binData, bitEnd, err = NGResetAcknowledge(*value.NGResetAcknowledge, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentNGSetupResponse:
		if value.NGSetupResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: NGSetupResponse: NIL")
		}
		binData, bitEnd, err = NGSetupResponse(*value.NGSetupResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentPathSwitchRequestAcknowledge:
		if value.PathSwitchRequestAcknowledge == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: PathSwitchRequestAcknowledge: NIL")
		}
		binData, bitEnd, err = PathSwitchRequestAcknowledge(*value.PathSwitchRequestAcknowledge, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyResponse:
		if value.PDUSessionResourceModifyResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: PDUSessionResourceModifyResponse: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyResponse(*value.PDUSessionResourceModifyResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyConfirm:
		if value.PDUSessionResourceModifyConfirm == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: PDUSessionResourceModifyConfirm: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceModifyConfirm(*value.PDUSessionResourceModifyConfirm, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentPDUSessionResourceReleaseResponse:
		if value.PDUSessionResourceReleaseResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: PDUSessionResourceReleaseResponse: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceReleaseResponse(*value.PDUSessionResourceReleaseResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentPDUSessionResourceSetupResponse:
		if value.PDUSessionResourceSetupResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: PDUSessionResourceSetupResponse: NIL")
		}
		binData, bitEnd, err = PDUSessionResourceSetupResponse(*value.PDUSessionResourceSetupResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentPWSCancelResponse:
		if value.PWSCancelResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: PWSCancelResponse: NIL")
		}
		binData, bitEnd, err = PWSCancelResponse(*value.PWSCancelResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge:
		if value.RANConfigurationUpdateAcknowledge == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: RANConfigurationUpdateAcknowledge: NIL")
		}
		binData, bitEnd, err = RANConfigurationUpdateAcknowledge(*value.RANConfigurationUpdateAcknowledge, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentUEContextModificationResponse:
		if value.UEContextModificationResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: UEContextModificationResponse: NIL")
		}
		binData, bitEnd, err = UEContextModificationResponse(*value.UEContextModificationResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentUEContextReleaseComplete:
		if value.UEContextReleaseComplete == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: UEContextReleaseComplete: NIL")
		}
		binData, bitEnd, err = UEContextReleaseComplete(*value.UEContextReleaseComplete, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentUERadioCapabilityCheckResponse:
		if value.UERadioCapabilityCheckResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: UERadioCapabilityCheckResponse: NIL")
		}
		binData, bitEnd, err = UERadioCapabilityCheckResponse(*value.UERadioCapabilityCheckResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.SuccessfulOutcomePresentWriteReplaceWarningResponse:
		if value.WriteReplaceWarningResponse == nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: WriteReplaceWarningResponse: NIL")
		}
		binData, bitEnd, err = WriteReplaceWarningResponse(*value.WriteReplaceWarningResponse, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("SuccessfulOutcomeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("SuccessfulOutcomeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}

func UnsuccessfulOutcome (value ngapType.UnsuccessfulOutcome, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.ProcedureCode.Value {
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentAMFConfigurationUpdateFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverPreparation:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentHandoverPreparationFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentHandoverFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeInitialContextSetup:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentInitialContextSetupFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeNGSetup:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentNGSetupFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodePathSwitchRequest:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentPathSwitchRequestFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentRANConfigurationUpdateFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	case ngapType.ProcedureCodeUEContextModification:
		if value.Value.Present != ngapType.UnsuccessfulOutcomePresentUEContextModificationFailure {
			return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
		}
	default:
		return binData, bitEnd, errors.New("UnsuccessfulOutcome: ProcedureCode: INVALID")
	}
	binData, bitEnd, err = ProcedureCode(value.ProcedureCode, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnsuccessfulOutcome: " + err.Error())
	}
	binData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UnsuccessfulOutcome: " + err.Error())
	}
	var binValue []byte
	binValue, bitEnd, err = UnsuccessfulOutcomeValue(value.Value, binValue, 8)
	if err != nil {
		return binData, bitEnd, errors.New("UnsuccessfulOutcome: " + err.Error())
	}
	return append(binData, EncodeLengthValue(binValue)...), bitEnd, err
}

func UnsuccessfulOutcomeValue (value ngapType.UnsuccessfulOutcomeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	switch value.Present {
	case ngapType.UnsuccessfulOutcomePresentAMFConfigurationUpdateFailure:
		if value.AMFConfigurationUpdateFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: AMFConfigurationUpdateFailure: NIL")
		}
		binData, bitEnd, err = AMFConfigurationUpdateFailure(*value.AMFConfigurationUpdateFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentHandoverPreparationFailure:
		if value.HandoverPreparationFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: HandoverPreparationFailure: NIL")
		}
		binData, bitEnd, err = HandoverPreparationFailure(*value.HandoverPreparationFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentHandoverFailure:
		if value.HandoverFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: HandoverFailure: NIL")
		}
		binData, bitEnd, err = HandoverFailure(*value.HandoverFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentInitialContextSetupFailure:
		if value.InitialContextSetupFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: InitialContextSetupFailure: NIL")
		}
		binData, bitEnd, err = InitialContextSetupFailure(*value.InitialContextSetupFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentNGSetupFailure:
		if value.NGSetupFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: NGSetupFailure: NIL")
		}
		binData, bitEnd, err = NGSetupFailure(*value.NGSetupFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentPathSwitchRequestFailure:
		if value.PathSwitchRequestFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: PathSwitchRequestFailure: NIL")
		}
		binData, bitEnd, err = PathSwitchRequestFailure(*value.PathSwitchRequestFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentRANConfigurationUpdateFailure:
		if value.RANConfigurationUpdateFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: RANConfigurationUpdateFailure: NIL")
		}
		binData, bitEnd, err = RANConfigurationUpdateFailure(*value.RANConfigurationUpdateFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	case ngapType.UnsuccessfulOutcomePresentUEContextModificationFailure:
		if value.UEContextModificationFailure == nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: UEContextModificationFailure: NIL")
		}
		binData, bitEnd, err = UEContextModificationFailure(*value.UEContextModificationFailure, binData, bitEnd)
		if err != nil {
			return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: " + err.Error())
		}
	default:
		return binData, bitEnd, errors.New("UnsuccessfulOutcomeValue: Present: INVALID")
	}
	return binData, bitEnd, err
}
