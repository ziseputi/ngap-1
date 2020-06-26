package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func CheckMandatory(pdu ngapType.NGAPPDU) error {
	switch pdu.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		switch pdu.InitiatingMessage.Value.Present {
		case ngapType.InitiatingMessagePresentAMFConfigurationUpdate:
			return nil
		case ngapType.InitiatingMessagePresentAMFStatusIndication:
			return nil
		case ngapType.InitiatingMessagePresentCellTrafficTrace:
			return nil
		case ngapType.InitiatingMessagePresentDeactivateTrace:
			return nil
		case ngapType.InitiatingMessagePresentDownlinkNASTransport:
			return nil
		case ngapType.InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport:
			return nil
		case ngapType.InitiatingMessagePresentDownlinkRANConfigurationTransfer:
			return nil
		case ngapType.InitiatingMessagePresentDownlinkRANStatusTransfer:
			return nil
		case ngapType.InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport:
			return nil
		case ngapType.InitiatingMessagePresentErrorIndication:
			return nil
		case ngapType.InitiatingMessagePresentHandoverCancel:
			return nil
		case ngapType.InitiatingMessagePresentHandoverNotify:
			return nil
		case ngapType.InitiatingMessagePresentHandoverRequired:
			return nil
		case ngapType.InitiatingMessagePresentHandoverRequest:
			return nil
		case ngapType.InitiatingMessagePresentInitialContextSetupRequest:
			return nil
		case ngapType.InitiatingMessagePresentInitialUEMessage:
			return nil
		case ngapType.InitiatingMessagePresentLocationReport:
			return nil
		case ngapType.InitiatingMessagePresentLocationReportingControl:
			return nil
		case ngapType.InitiatingMessagePresentLocationReportingFailureIndication:
			return nil
		case ngapType.InitiatingMessagePresentNASNonDeliveryIndication:
			return nil
		case ngapType.InitiatingMessagePresentNGReset:
			return nil
		case ngapType.InitiatingMessagePresentNGSetupRequest:
			return nil
		case ngapType.InitiatingMessagePresentOverloadStart:
			return nil
		case ngapType.InitiatingMessagePresentOverloadStop:
			return nil
		case ngapType.InitiatingMessagePresentPaging:
			return nil
		case ngapType.InitiatingMessagePresentPathSwitchRequest:
			return nil
		case ngapType.InitiatingMessagePresentPDUSessionResourceModifyRequest:
			return nil
		case ngapType.InitiatingMessagePresentPDUSessionResourceModifyIndication:
			return nil
		case ngapType.InitiatingMessagePresentPDUSessionResourceNotify:
			return nil
		case ngapType.InitiatingMessagePresentPDUSessionResourceReleaseCommand:
			return nil
		case ngapType.InitiatingMessagePresentPDUSessionResourceSetupRequest:
			return nil
		case ngapType.InitiatingMessagePresentPrivateMessage:
			return nil
		case ngapType.InitiatingMessagePresentPWSCancelRequest:
			return nil
		case ngapType.InitiatingMessagePresentPWSFailureIndication:
			return nil
		case ngapType.InitiatingMessagePresentPWSRestartIndication:
			return nil
		case ngapType.InitiatingMessagePresentRANConfigurationUpdate:
			return nil
		case ngapType.InitiatingMessagePresentRerouteNASRequest:
			return nil
		case ngapType.InitiatingMessagePresentRRCInactiveTransitionReport:
			return nil
		case ngapType.InitiatingMessagePresentSecondaryRATDataUsageReport:
			return nil
		case ngapType.InitiatingMessagePresentTraceFailureIndication:
			return nil
		case ngapType.InitiatingMessagePresentTraceStart:
			return nil
		case ngapType.InitiatingMessagePresentUEContextModificationRequest:
			return nil
		case ngapType.InitiatingMessagePresentUEContextReleaseCommand:
			return nil
		case ngapType.InitiatingMessagePresentUEContextReleaseRequest:
			return nil
		case ngapType.InitiatingMessagePresentUERadioCapabilityCheckRequest:
			return nil
		case ngapType.InitiatingMessagePresentUERadioCapabilityInfoIndication:
			return nil
		case ngapType.InitiatingMessagePresentUETNLABindingReleaseRequest:
			return nil
		case ngapType.InitiatingMessagePresentUplinkNASTransport:
			return nil
		case ngapType.InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport:
			return nil
		case ngapType.InitiatingMessagePresentUplinkRANConfigurationTransfer:
			return nil
		case ngapType.InitiatingMessagePresentUplinkRANStatusTransfer:
			return nil
		case ngapType.InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport:
			return nil
		case ngapType.InitiatingMessagePresentWriteReplaceWarningRequest:
			return nil
		case ngapType.InitiatingMessagePresentUplinkRIMInformationTransfer:
			return nil
		case ngapType.InitiatingMessagePresentDownlinkRIMInformationTransfer:
			return nil
		default:
			return errors.New("NGAPPDU: InitiatingMessage: Present: INVALID")
		}
	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		switch pdu.SuccessfulOutcome.Value.Present {
		case ngapType.SuccessfulOutcomePresentAMFConfigurationUpdateAcknowledge:
			return nil
		case ngapType.SuccessfulOutcomePresentHandoverCancelAcknowledge:
			return nil
		case ngapType.SuccessfulOutcomePresentHandoverCommand:
			return nil
		case ngapType.SuccessfulOutcomePresentHandoverRequestAcknowledge:
			return nil
		case ngapType.SuccessfulOutcomePresentInitialContextSetupResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentNGResetAcknowledge:
			return nil
		case ngapType.SuccessfulOutcomePresentNGSetupResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentPathSwitchRequestAcknowledge:
			return nil
		case ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyConfirm:
			return nil
		case ngapType.SuccessfulOutcomePresentPDUSessionResourceReleaseResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentPDUSessionResourceSetupResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentPWSCancelResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge:
			return nil
		case ngapType.SuccessfulOutcomePresentUEContextModificationResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentUEContextReleaseComplete:
			return nil
		case ngapType.SuccessfulOutcomePresentUERadioCapabilityCheckResponse:
			return nil
		case ngapType.SuccessfulOutcomePresentWriteReplaceWarningResponse:
			return nil
		default:
			return errors.New("NGAPPDU: SuccessfulOutcome: Present: INVALID")
		}
	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		switch pdu.UnsuccessfulOutcome.Value.Present {
		case ngapType.UnsuccessfulOutcomePresentAMFConfigurationUpdateFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentHandoverPreparationFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentHandoverFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentInitialContextSetupFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentNGSetupFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentPathSwitchRequestFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentRANConfigurationUpdateFailure:
			return nil
		case ngapType.UnsuccessfulOutcomePresentUEContextModificationFailure:
			return nil
		default:
			return errors.New("NGAPPDU: UnsuccessfulOutcome: Present: INVALID")
		}
	default:
		return errors.New("NGAPPDU: Present: INVALID")
	}
}
