// Created By HaoDHH-245789 VHT2020
package ngapBuild

import "ngap/ngapType"

func AMFConfigurationUpdate() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentAMFConfigurationUpdate
	msg.Value.AMFConfigurationUpdate = new(ngapType.AMFConfigurationUpdate)
	msgProtocolIEs := &msg.Value.AMFConfigurationUpdate.ProtocolIEs
	protocolIEs := ngapType.AMFConfigurationUpdateIEs{}

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFName
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFName
	protocolIEs.TypeValue.AMFName = new(ngapType.AMFName)
	*protocolIEs.TypeValue.AMFName = AMFName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDServedGUAMIList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentServedGUAMIList
	protocolIEs.TypeValue.ServedGUAMIList = new(ngapType.ServedGUAMIList)
	*protocolIEs.TypeValue.ServedGUAMIList = ServedGUAMIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRelativeAMFCapacity
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentRelativeAMFCapacity
	protocolIEs.TypeValue.RelativeAMFCapacity = new(ngapType.RelativeAMFCapacity)
	*protocolIEs.TypeValue.RelativeAMFCapacity = RelativeAMFCapacity()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPLMNSupportList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentPLMNSupportList
	protocolIEs.TypeValue.PLMNSupportList = new(ngapType.PLMNSupportList)
	*protocolIEs.TypeValue.PLMNSupportList = PLMNSupportList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFTNLAssociationToAddList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToAddList
	protocolIEs.TypeValue.AMFTNLAssociationToAddList = new(ngapType.AMFTNLAssociationToAddList)
	*protocolIEs.TypeValue.AMFTNLAssociationToAddList = AMFTNLAssociationToAddList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFTNLAssociationToRemoveList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToRemoveList
	protocolIEs.TypeValue.AMFTNLAssociationToRemoveList = new(ngapType.AMFTNLAssociationToRemoveList)
	*protocolIEs.TypeValue.AMFTNLAssociationToRemoveList = AMFTNLAssociationToRemoveList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFTNLAssociationToUpdateList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateIEsTypeValuePresentAMFTNLAssociationToUpdateList
	protocolIEs.TypeValue.AMFTNLAssociationToUpdateList = new(ngapType.AMFTNLAssociationToUpdateList)
	*protocolIEs.TypeValue.AMFTNLAssociationToUpdateList = AMFTNLAssociationToUpdateList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func AMFConfigurationUpdateAcknowledge() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentAMFConfigurationUpdateAcknowledge
	msg.Value.AMFConfigurationUpdateAcknowledge = new(ngapType.AMFConfigurationUpdateAcknowledge)
	msgProtocolIEs := &msg.Value.AMFConfigurationUpdateAcknowledge.ProtocolIEs
	protocolIEs := ngapType.AMFConfigurationUpdateAcknowledgeIEs{}

	protocolIEs = ngapType.AMFConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFTNLAssociationSetupList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationSetupList
	protocolIEs.TypeValue.AMFTNLAssociationSetupList = new(ngapType.AMFTNLAssociationSetupList)
	*protocolIEs.TypeValue.AMFTNLAssociationSetupList = AMFTNLAssociationSetupList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFTNLAssociationFailedToSetupList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentAMFTNLAssociationFailedToSetupList
	protocolIEs.TypeValue.AMFTNLAssociationFailedToSetupList = new(ngapType.TNLAssociationList)
	*protocolIEs.TypeValue.AMFTNLAssociationFailedToSetupList = TNLAssociationList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func AMFConfigurationUpdateFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentAMFConfigurationUpdateFailure
	msg.Value.AMFConfigurationUpdateFailure = new(ngapType.AMFConfigurationUpdateFailure)
	msgProtocolIEs := &msg.Value.AMFConfigurationUpdateFailure.ProtocolIEs
	protocolIEs := ngapType.AMFConfigurationUpdateFailureIEs{}

	protocolIEs = ngapType.AMFConfigurationUpdateFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTimeToWait
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentTimeToWait
	protocolIEs.TypeValue.TimeToWait = new(ngapType.TimeToWait)
	*protocolIEs.TypeValue.TimeToWait = TimeToWait()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.AMFConfigurationUpdateFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.AMFConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func AMFStatusIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFStatusIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentAMFStatusIndication
	msg.Value.AMFStatusIndication = new(ngapType.AMFStatusIndication)
	msgProtocolIEs := &msg.Value.AMFStatusIndication.ProtocolIEs
	protocolIEs := ngapType.AMFStatusIndicationIEs{}

	protocolIEs = ngapType.AMFStatusIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUnavailableGUAMIList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.AMFStatusIndicationIEsTypeValuePresentUnavailableGUAMIList
	protocolIEs.TypeValue.UnavailableGUAMIList = new(ngapType.UnavailableGUAMIList)
	*protocolIEs.TypeValue.UnavailableGUAMIList = UnavailableGUAMIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func CellTrafficTrace() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeCellTrafficTrace
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentCellTrafficTrace
	msg.Value.CellTrafficTrace = new(ngapType.CellTrafficTrace)
	msgProtocolIEs := &msg.Value.CellTrafficTrace.ProtocolIEs
	protocolIEs := ngapType.CellTrafficTraceIEs{}

	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.CellTrafficTraceIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.CellTrafficTraceIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNGRANTraceID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.CellTrafficTraceIEsTypeValuePresentNGRANTraceID
	protocolIEs.TypeValue.NGRANTraceID = new(ngapType.NGRANTraceID)
	*protocolIEs.TypeValue.NGRANTraceID = NGRANTraceID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNGRANCGI
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.CellTrafficTraceIEsTypeValuePresentNGRANCGI
	protocolIEs.TypeValue.NGRANCGI = new(ngapType.NGRANCGI)
	*protocolIEs.TypeValue.NGRANCGI = NGRANCGI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTraceCollectionEntityIPAddress
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.CellTrafficTraceIEsTypeValuePresentTraceCollectionEntityIPAddress
	protocolIEs.TypeValue.TraceCollectionEntityIPAddress = new(ngapType.TransportLayerAddress)
	*protocolIEs.TypeValue.TraceCollectionEntityIPAddress = TransportLayerAddress()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DeactivateTrace() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDeactivateTrace
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDeactivateTrace
	msg.Value.DeactivateTrace = new(ngapType.DeactivateTrace)
	msgProtocolIEs := &msg.Value.DeactivateTrace.ProtocolIEs
	protocolIEs := ngapType.DeactivateTraceIEs{}

	protocolIEs = ngapType.DeactivateTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DeactivateTraceIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DeactivateTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DeactivateTraceIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DeactivateTraceIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNGRANTraceID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DeactivateTraceIEsTypeValuePresentNGRANTraceID
	protocolIEs.TypeValue.NGRANTraceID = new(ngapType.NGRANTraceID)
	*protocolIEs.TypeValue.NGRANTraceID = NGRANTraceID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkNASTransport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkNASTransport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDownlinkNASTransport
	msg.Value.DownlinkNASTransport = new(ngapType.DownlinkNASTransport)
	msgProtocolIEs := &msg.Value.DownlinkNASTransport.ProtocolIEs
	protocolIEs := ngapType.DownlinkNASTransportIEs{}

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDOldAMF
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentOldAMF
	protocolIEs.TypeValue.OldAMF = new(ngapType.AMFName)
	*protocolIEs.TypeValue.OldAMF = AMFName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentRANPagingPriority
	protocolIEs.TypeValue.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.TypeValue.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMobilityRestrictionList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentMobilityRestrictionList
	protocolIEs.TypeValue.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)
	*protocolIEs.TypeValue.MobilityRestrictionList = MobilityRestrictionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDIndexToRFSP
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentIndexToRFSP
	protocolIEs.TypeValue.IndexToRFSP = new(ngapType.IndexToRFSP)
	*protocolIEs.TypeValue.IndexToRFSP = IndexToRFSP()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentUEAggregateMaximumBitRate
	protocolIEs.TypeValue.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.TypeValue.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentAllowedNSSAI
	protocolIEs.TypeValue.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.TypeValue.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSRVCCOperationPossible
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkNASTransportIEsTypeValuePresentSRVCCOperationPossible
	protocolIEs.TypeValue.SRVCCOperationPossible = new(ngapType.SRVCCOperationPossible)
	*protocolIEs.TypeValue.SRVCCOperationPossible = SRVCCOperationPossible()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkNonUEAssociatedNRPPaTransport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDownlinkNonUEAssociatedNRPPaTransport
	msg.Value.DownlinkNonUEAssociatedNRPPaTransport = new(ngapType.DownlinkNonUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.DownlinkNonUEAssociatedNRPPaTransport.ProtocolIEs
	protocolIEs := ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}

	protocolIEs = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID
	protocolIEs.TypeValue.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.TypeValue.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU
	protocolIEs.TypeValue.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.TypeValue.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkRANConfigurationTransfer() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkRANConfigurationTransfer
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDownlinkRANConfigurationTransfer
	msg.Value.DownlinkRANConfigurationTransfer = new(ngapType.DownlinkRANConfigurationTransfer)
	msgProtocolIEs := &msg.Value.DownlinkRANConfigurationTransfer.ProtocolIEs
	protocolIEs := ngapType.DownlinkRANConfigurationTransferIEs{}

	protocolIEs = ngapType.DownlinkRANConfigurationTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSONConfigurationTransferDL
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferDL
	protocolIEs.TypeValue.SONConfigurationTransferDL = new(ngapType.SONConfigurationTransfer)
	*protocolIEs.TypeValue.SONConfigurationTransferDL = SONConfigurationTransfer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkRANConfigurationTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDENDCSONConfigurationTransferDL
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferDL
	protocolIEs.TypeValue.ENDCSONConfigurationTransferDL = new(ngapType.ENDCSONConfigurationTransfer)
	*protocolIEs.TypeValue.ENDCSONConfigurationTransferDL = ENDCSONConfigurationTransfer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkRANStatusTransfer() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkRANStatusTransfer
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDownlinkRANStatusTransfer
	msg.Value.DownlinkRANStatusTransfer = new(ngapType.DownlinkRANStatusTransfer)
	msgProtocolIEs := &msg.Value.DownlinkRANStatusTransfer.ProtocolIEs
	protocolIEs := ngapType.DownlinkRANStatusTransferIEs{}

	protocolIEs = ngapType.DownlinkRANStatusTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkRANStatusTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkRANStatusTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANStatusTransferTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer
	protocolIEs.TypeValue.RANStatusTransferTransparentContainer = new(ngapType.RANStatusTransferTransparentContainer)
	*protocolIEs.TypeValue.RANStatusTransferTransparentContainer = RANStatusTransferTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkUEAssociatedNRPPaTransport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDownlinkUEAssociatedNRPPaTransport
	msg.Value.DownlinkUEAssociatedNRPPaTransport = new(ngapType.DownlinkUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.DownlinkUEAssociatedNRPPaTransport.ProtocolIEs
	protocolIEs := ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}

	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID
	protocolIEs.TypeValue.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.TypeValue.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU
	protocolIEs.TypeValue.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.TypeValue.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func ErrorIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeErrorIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentErrorIndication
	msg.Value.ErrorIndication = new(ngapType.ErrorIndication)
	msgProtocolIEs := &msg.Value.ErrorIndication.ProtocolIEs
	protocolIEs := ngapType.ErrorIndicationIEs{}

	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.ErrorIndicationIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.ErrorIndicationIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.ErrorIndicationIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.ErrorIndicationIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverCancel() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverCancel
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentHandoverCancel
	msg.Value.HandoverCancel = new(ngapType.HandoverCancel)
	msgProtocolIEs := &msg.Value.HandoverCancel.ProtocolIEs
	protocolIEs := ngapType.HandoverCancelIEs{}

	protocolIEs = ngapType.HandoverCancelIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCancelIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCancelIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCancelIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCancelIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCancelIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverCancelAcknowledge() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverCancel
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentHandoverCancelAcknowledge
	msg.Value.HandoverCancelAcknowledge = new(ngapType.HandoverCancelAcknowledge)
	msgProtocolIEs := &msg.Value.HandoverCancelAcknowledge.ProtocolIEs
	protocolIEs := ngapType.HandoverCancelAcknowledgeIEs{}

	protocolIEs = ngapType.HandoverCancelAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCancelAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCancelAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCancelAcknowledgeIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverNotify() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverNotification
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentHandoverNotify
	msg.Value.HandoverNotify = new(ngapType.HandoverNotify)
	msgProtocolIEs := &msg.Value.HandoverNotify.ProtocolIEs
	protocolIEs := ngapType.HandoverNotifyIEs{}

	protocolIEs = ngapType.HandoverNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverNotifyIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverNotifyIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverNotifyIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverRequired() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentHandoverRequired
	msg.Value.HandoverRequired = new(ngapType.HandoverRequired)
	msgProtocolIEs := &msg.Value.HandoverRequired.ProtocolIEs
	protocolIEs := ngapType.HandoverRequiredIEs{}

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDHandoverType
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentHandoverType
	protocolIEs.TypeValue.HandoverType = new(ngapType.HandoverType)
	*protocolIEs.TypeValue.HandoverType = HandoverType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTargetID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentTargetID
	protocolIEs.TypeValue.TargetID = new(ngapType.TargetID)
	*protocolIEs.TypeValue.TargetID = TargetID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDDirectForwardingPathAvailability
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentDirectForwardingPathAvailability
	protocolIEs.TypeValue.DirectForwardingPathAvailability = new(ngapType.DirectForwardingPathAvailability)
	*protocolIEs.TypeValue.DirectForwardingPathAvailability = DirectForwardingPathAvailability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceListHORqd
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentPDUSessionResourceListHORqd
	protocolIEs.TypeValue.PDUSessionResourceListHORqd = new(ngapType.PDUSessionResourceListHORqd)
	*protocolIEs.TypeValue.PDUSessionResourceListHORqd = PDUSessionResourceListHORqd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSourceToTargetTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequiredIEsTypeValuePresentSourceToTargetTransparentContainer
	protocolIEs.TypeValue.SourceToTargetTransparentContainer = new(ngapType.SourceToTargetTransparentContainer)
	*protocolIEs.TypeValue.SourceToTargetTransparentContainer = SourceToTargetTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverCommand() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentHandoverCommand
	msg.Value.HandoverCommand = new(ngapType.HandoverCommand)
	msgProtocolIEs := &msg.Value.HandoverCommand.ProtocolIEs
	protocolIEs := ngapType.HandoverCommandIEs{}

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDHandoverType
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentHandoverType
	protocolIEs.TypeValue.HandoverType = new(ngapType.HandoverType)
	*protocolIEs.TypeValue.HandoverType = HandoverType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASSecurityParametersFromNGRAN
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentNASSecurityParametersFromNGRAN
	protocolIEs.TypeValue.NASSecurityParametersFromNGRAN = new(ngapType.NASSecurityParametersFromNGRAN)
	*protocolIEs.TypeValue.NASSecurityParametersFromNGRAN = NASSecurityParametersFromNGRAN()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceHandoverList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceHandoverList
	protocolIEs.TypeValue.PDUSessionResourceHandoverList = new(ngapType.PDUSessionResourceHandoverList)
	*protocolIEs.TypeValue.PDUSessionResourceHandoverList = PDUSessionResourceHandoverList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentPDUSessionResourceToReleaseListHOCmd
	protocolIEs.TypeValue.PDUSessionResourceToReleaseListHOCmd = new(ngapType.PDUSessionResourceToReleaseListHOCmd)
	*protocolIEs.TypeValue.PDUSessionResourceToReleaseListHOCmd = PDUSessionResourceToReleaseListHOCmd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTargetToSourceTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentTargetToSourceTransparentContainer
	protocolIEs.TypeValue.TargetToSourceTransparentContainer = new(ngapType.TargetToSourceTransparentContainer)
	*protocolIEs.TypeValue.TargetToSourceTransparentContainer = TargetToSourceTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverCommandIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverPreparationFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentHandoverPreparationFailure
	msg.Value.HandoverPreparationFailure = new(ngapType.HandoverPreparationFailure)
	msgProtocolIEs := &msg.Value.HandoverPreparationFailure.ProtocolIEs
	protocolIEs := ngapType.HandoverPreparationFailureIEs{}

	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverPreparationFailureIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverPreparationFailureIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverPreparationFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverPreparationFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentHandoverRequest
	msg.Value.HandoverRequest = new(ngapType.HandoverRequest)
	msgProtocolIEs := &msg.Value.HandoverRequest.ProtocolIEs
	protocolIEs := ngapType.HandoverRequestIEs{}

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDHandoverType
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentHandoverType
	protocolIEs.TypeValue.HandoverType = new(ngapType.HandoverType)
	*protocolIEs.TypeValue.HandoverType = HandoverType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentUEAggregateMaximumBitRate
	protocolIEs.TypeValue.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.TypeValue.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive
	protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = new(ngapType.CoreNetworkAssistanceInformationForInactive)
	*protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentUESecurityCapabilities
	protocolIEs.TypeValue.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.TypeValue.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSecurityContext
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentSecurityContext
	protocolIEs.TypeValue.SecurityContext = new(ngapType.SecurityContext)
	*protocolIEs.TypeValue.SecurityContext = SecurityContext()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNewSecurityContextInd
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentNewSecurityContextInd
	protocolIEs.TypeValue.NewSecurityContextInd = new(ngapType.NewSecurityContextInd)
	*protocolIEs.TypeValue.NewSecurityContextInd = NewSecurityContextInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASC
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentNASC
	protocolIEs.TypeValue.NASC = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASC = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListHOReq
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentPDUSessionResourceSetupListHOReq
	protocolIEs.TypeValue.PDUSessionResourceSetupListHOReq = new(ngapType.PDUSessionResourceSetupListHOReq)
	*protocolIEs.TypeValue.PDUSessionResourceSetupListHOReq = PDUSessionResourceSetupListHOReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentAllowedNSSAI
	protocolIEs.TypeValue.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.TypeValue.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTraceActivation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentTraceActivation
	protocolIEs.TypeValue.TraceActivation = new(ngapType.TraceActivation)
	*protocolIEs.TypeValue.TraceActivation = TraceActivation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMaskedIMEISV
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentMaskedIMEISV
	protocolIEs.TypeValue.MaskedIMEISV = new(ngapType.MaskedIMEISV)
	*protocolIEs.TypeValue.MaskedIMEISV = MaskedIMEISV()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSourceToTargetTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentSourceToTargetTransparentContainer
	protocolIEs.TypeValue.SourceToTargetTransparentContainer = new(ngapType.SourceToTargetTransparentContainer)
	*protocolIEs.TypeValue.SourceToTargetTransparentContainer = SourceToTargetTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMobilityRestrictionList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentMobilityRestrictionList
	protocolIEs.TypeValue.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)
	*protocolIEs.TypeValue.MobilityRestrictionList = MobilityRestrictionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentLocationReportingRequestType
	protocolIEs.TypeValue.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.TypeValue.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest
	protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDGUAMI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentGUAMI
	protocolIEs.TypeValue.GUAMI = new(ngapType.GUAMI)
	*protocolIEs.TypeValue.GUAMI = GUAMI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRedirectionVoiceFallback
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentRedirectionVoiceFallback
	protocolIEs.TypeValue.RedirectionVoiceFallback = new(ngapType.RedirectionVoiceFallback)
	*protocolIEs.TypeValue.RedirectionVoiceFallback = RedirectionVoiceFallback()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCNAssistedRANTuning
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentCNAssistedRANTuning
	protocolIEs.TypeValue.CNAssistedRANTuning = new(ngapType.CNAssistedRANTuning)
	*protocolIEs.TypeValue.CNAssistedRANTuning = CNAssistedRANTuning()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSRVCCOperationPossible
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestIEsTypeValuePresentSRVCCOperationPossible
	protocolIEs.TypeValue.SRVCCOperationPossible = new(ngapType.SRVCCOperationPossible)
	*protocolIEs.TypeValue.SRVCCOperationPossible = SRVCCOperationPossible()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverRequestAcknowledge() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentHandoverRequestAcknowledge
	msg.Value.HandoverRequestAcknowledge = new(ngapType.HandoverRequestAcknowledge)
	msgProtocolIEs := &msg.Value.HandoverRequestAcknowledge.ProtocolIEs
	protocolIEs := ngapType.HandoverRequestAcknowledgeIEs{}

	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceAdmittedList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceAdmittedList
	protocolIEs.TypeValue.PDUSessionResourceAdmittedList = new(ngapType.PDUSessionResourceAdmittedList)
	*protocolIEs.TypeValue.PDUSessionResourceAdmittedList = PDUSessionResourceAdmittedList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceFailedToSetupListHOAck
	protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListHOAck = new(ngapType.PDUSessionResourceFailedToSetupListHOAck)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListHOAck = PDUSessionResourceFailedToSetupListHOAck()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTargetToSourceTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentTargetToSourceTransparentContainer
	protocolIEs.TypeValue.TargetToSourceTransparentContainer = new(ngapType.TargetToSourceTransparentContainer)
	*protocolIEs.TypeValue.TargetToSourceTransparentContainer = TargetToSourceTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentHandoverFailure
	msg.Value.HandoverFailure = new(ngapType.HandoverFailure)
	msgProtocolIEs := &msg.Value.HandoverFailure.ProtocolIEs
	protocolIEs := ngapType.HandoverFailureIEs{}

	protocolIEs = ngapType.HandoverFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverFailureIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.HandoverFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.HandoverFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialContextSetupRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentInitialContextSetupRequest
	msg.Value.InitialContextSetupRequest = new(ngapType.InitialContextSetupRequest)
	msgProtocolIEs := &msg.Value.InitialContextSetupRequest.ProtocolIEs
	protocolIEs := ngapType.InitialContextSetupRequestIEs{}

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDOldAMF
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentOldAMF
	protocolIEs.TypeValue.OldAMF = new(ngapType.AMFName)
	*protocolIEs.TypeValue.OldAMF = AMFName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate
	protocolIEs.TypeValue.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.TypeValue.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive
	protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = new(ngapType.CoreNetworkAssistanceInformationForInactive)
	*protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDGUAMI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentGUAMI
	protocolIEs.TypeValue.GUAMI = new(ngapType.GUAMI)
	*protocolIEs.TypeValue.GUAMI = GUAMI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtReq
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListCxtReq
	protocolIEs.TypeValue.PDUSessionResourceSetupListCxtReq = new(ngapType.PDUSessionResourceSetupListCxtReq)
	*protocolIEs.TypeValue.PDUSessionResourceSetupListCxtReq = PDUSessionResourceSetupListCxtReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentAllowedNSSAI
	protocolIEs.TypeValue.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.TypeValue.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentUESecurityCapabilities
	protocolIEs.TypeValue.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.TypeValue.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSecurityKey
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentSecurityKey
	protocolIEs.TypeValue.SecurityKey = new(ngapType.SecurityKey)
	*protocolIEs.TypeValue.SecurityKey = SecurityKey()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTraceActivation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentTraceActivation
	protocolIEs.TypeValue.TraceActivation = new(ngapType.TraceActivation)
	*protocolIEs.TypeValue.TraceActivation = TraceActivation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMobilityRestrictionList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentMobilityRestrictionList
	protocolIEs.TypeValue.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)
	*protocolIEs.TypeValue.MobilityRestrictionList = MobilityRestrictionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERadioCapability
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapability
	protocolIEs.TypeValue.UERadioCapability = new(ngapType.UERadioCapability)
	*protocolIEs.TypeValue.UERadioCapability = UERadioCapability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDIndexToRFSP
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentIndexToRFSP
	protocolIEs.TypeValue.IndexToRFSP = new(ngapType.IndexToRFSP)
	*protocolIEs.TypeValue.IndexToRFSP = IndexToRFSP()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMaskedIMEISV
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentMaskedIMEISV
	protocolIEs.TypeValue.MaskedIMEISV = new(ngapType.MaskedIMEISV)
	*protocolIEs.TypeValue.MaskedIMEISV = MaskedIMEISV()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDEmergencyFallbackIndicator
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentEmergencyFallbackIndicator
	protocolIEs.TypeValue.EmergencyFallbackIndicator = new(ngapType.EmergencyFallbackIndicator)
	*protocolIEs.TypeValue.EmergencyFallbackIndicator = EmergencyFallbackIndicator()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest
	protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentUERadioCapabilityForPaging
	protocolIEs.TypeValue.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
	*protocolIEs.TypeValue.UERadioCapabilityForPaging = UERadioCapabilityForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRedirectionVoiceFallback
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentRedirectionVoiceFallback
	protocolIEs.TypeValue.RedirectionVoiceFallback = new(ngapType.RedirectionVoiceFallback)
	*protocolIEs.TypeValue.RedirectionVoiceFallback = RedirectionVoiceFallback()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentLocationReportingRequestType
	protocolIEs.TypeValue.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.TypeValue.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCNAssistedRANTuning
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentCNAssistedRANTuning
	protocolIEs.TypeValue.CNAssistedRANTuning = new(ngapType.CNAssistedRANTuning)
	*protocolIEs.TypeValue.CNAssistedRANTuning = CNAssistedRANTuning()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSRVCCOperationPossible
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupRequestIEsTypeValuePresentSRVCCOperationPossible
	protocolIEs.TypeValue.SRVCCOperationPossible = new(ngapType.SRVCCOperationPossible)
	*protocolIEs.TypeValue.SRVCCOperationPossible = SRVCCOperationPossible()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialContextSetupResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentInitialContextSetupResponse
	msg.Value.InitialContextSetupResponse = new(ngapType.InitialContextSetupResponse)
	msgProtocolIEs := &msg.Value.InitialContextSetupResponse.ProtocolIEs
	protocolIEs := ngapType.InitialContextSetupResponseIEs{}

	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupResponseIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupResponseIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtRes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListCxtRes
	protocolIEs.TypeValue.PDUSessionResourceSetupListCxtRes = new(ngapType.PDUSessionResourceSetupListCxtRes)
	*protocolIEs.TypeValue.PDUSessionResourceSetupListCxtRes = PDUSessionResourceSetupListCxtRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtRes
	protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListCxtRes = new(ngapType.PDUSessionResourceFailedToSetupListCxtRes)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListCxtRes = PDUSessionResourceFailedToSetupListCxtRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialContextSetupFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentInitialContextSetupFailure
	msg.Value.InitialContextSetupFailure = new(ngapType.InitialContextSetupFailure)
	msgProtocolIEs := &msg.Value.InitialContextSetupFailure.ProtocolIEs
	protocolIEs := ngapType.InitialContextSetupFailureIEs{}

	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupFailureIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupFailureIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupFailureIEsTypeValuePresentPDUSessionResourceFailedToSetupListCxtFail
	protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListCxtFail = new(ngapType.PDUSessionResourceFailedToSetupListCxtFail)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListCxtFail = PDUSessionResourceFailedToSetupListCxtFail()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialContextSetupFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialUEMessage() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialUEMessage
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentInitialUEMessage
	msg.Value.InitialUEMessage = new(ngapType.InitialUEMessage)
	msgProtocolIEs := &msg.Value.InitialUEMessage.ProtocolIEs
	protocolIEs := ngapType.InitialUEMessageIEs{}

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCEstablishmentCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentRRCEstablishmentCause
	protocolIEs.TypeValue.RRCEstablishmentCause = new(ngapType.RRCEstablishmentCause)
	*protocolIEs.TypeValue.RRCEstablishmentCause = RRCEstablishmentCause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDFiveGSTMSI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentFiveGSTMSI
	protocolIEs.TypeValue.FiveGSTMSI = new(ngapType.FiveGSTMSI)
	*protocolIEs.TypeValue.FiveGSTMSI = FiveGSTMSI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFSetID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentAMFSetID
	protocolIEs.TypeValue.AMFSetID = new(ngapType.AMFSetID)
	*protocolIEs.TypeValue.AMFSetID = AMFSetID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEContextRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentUEContextRequest
	protocolIEs.TypeValue.UEContextRequest = new(ngapType.UEContextRequest)
	*protocolIEs.TypeValue.UEContextRequest = UEContextRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentAllowedNSSAI
	protocolIEs.TypeValue.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.TypeValue.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSourceToTargetAMFInformationReroute
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.InitialUEMessageIEsTypeValuePresentSourceToTargetAMFInformationReroute
	protocolIEs.TypeValue.SourceToTargetAMFInformationReroute = new(ngapType.SourceToTargetAMFInformationReroute)
	*protocolIEs.TypeValue.SourceToTargetAMFInformationReroute = SourceToTargetAMFInformationReroute()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func LocationReport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeLocationReport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentLocationReport
	msg.Value.LocationReport = new(ngapType.LocationReport)
	msgProtocolIEs := &msg.Value.LocationReport.ProtocolIEs
	protocolIEs := ngapType.LocationReportIEs{}

	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.LocationReportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.LocationReportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.LocationReportIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEPresenceInAreaOfInterestList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.LocationReportIEsTypeValuePresentUEPresenceInAreaOfInterestList
	protocolIEs.TypeValue.UEPresenceInAreaOfInterestList = new(ngapType.UEPresenceInAreaOfInterestList)
	*protocolIEs.TypeValue.UEPresenceInAreaOfInterestList = UEPresenceInAreaOfInterestList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.LocationReportIEsTypeValuePresentLocationReportingRequestType
	protocolIEs.TypeValue.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.TypeValue.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func LocationReportingControl() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeLocationReportingControl
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentLocationReportingControl
	msg.Value.LocationReportingControl = new(ngapType.LocationReportingControl)
	msgProtocolIEs := &msg.Value.LocationReportingControl.ProtocolIEs
	protocolIEs := ngapType.LocationReportingControlIEs{}

	protocolIEs = ngapType.LocationReportingControlIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.LocationReportingControlIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportingControlIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.LocationReportingControlIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportingControlIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.LocationReportingControlIEsTypeValuePresentLocationReportingRequestType
	protocolIEs.TypeValue.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.TypeValue.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func LocationReportingFailureIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeLocationReportingFailureIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentLocationReportingFailureIndication
	msg.Value.LocationReportingFailureIndication = new(ngapType.LocationReportingFailureIndication)
	msgProtocolIEs := &msg.Value.LocationReportingFailureIndication.ProtocolIEs
	protocolIEs := ngapType.LocationReportingFailureIndicationIEs{}

	protocolIEs = ngapType.LocationReportingFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.LocationReportingFailureIndicationIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportingFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.LocationReportingFailureIndicationIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.LocationReportingFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.LocationReportingFailureIndicationIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NASNonDeliveryIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNASNonDeliveryIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentNASNonDeliveryIndication
	msg.Value.NASNonDeliveryIndication = new(ngapType.NASNonDeliveryIndication)
	msgProtocolIEs := &msg.Value.NASNonDeliveryIndication.ProtocolIEs
	protocolIEs := ngapType.NASNonDeliveryIndicationIEs{}

	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NASNonDeliveryIndicationIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NASNonDeliveryIndicationIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NASNonDeliveryIndicationIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NASNonDeliveryIndicationIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGReset() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGReset
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentNGReset
	msg.Value.NGReset = new(ngapType.NGReset)
	msgProtocolIEs := &msg.Value.NGReset.ProtocolIEs
	protocolIEs := ngapType.NGResetIEs{}

	protocolIEs = ngapType.NGResetIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGResetIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGResetIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDResetType
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NGResetIEsTypeValuePresentResetType
	protocolIEs.TypeValue.ResetType = new(ngapType.ResetType)
	*protocolIEs.TypeValue.ResetType = ResetType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGResetAcknowledge() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGReset
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentNGResetAcknowledge
	msg.Value.NGResetAcknowledge = new(ngapType.NGResetAcknowledge)
	msgProtocolIEs := &msg.Value.NGResetAcknowledge.ProtocolIEs
	protocolIEs := ngapType.NGResetAcknowledgeIEs{}

	protocolIEs = ngapType.NGResetAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEAssociatedLogicalNGConnectionList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGResetAcknowledgeIEsTypeValuePresentUEAssociatedLogicalNGConnectionList
	protocolIEs.TypeValue.UEAssociatedLogicalNGConnectionList = new(ngapType.UEAssociatedLogicalNGConnectionList)
	*protocolIEs.TypeValue.UEAssociatedLogicalNGConnectionList = UEAssociatedLogicalNGConnectionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGResetAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGResetAcknowledgeIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGSetupRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentNGSetupRequest
	msg.Value.NGSetupRequest = new(ngapType.NGSetupRequest)
	msgProtocolIEs := &msg.Value.NGSetupRequest.ProtocolIEs
	protocolIEs := ngapType.NGSetupRequestIEs{}

	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NGSetupRequestIEsTypeValuePresentGlobalRANNodeID
	protocolIEs.TypeValue.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.TypeValue.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANNodeName
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupRequestIEsTypeValuePresentRANNodeName
	protocolIEs.TypeValue.RANNodeName = new(ngapType.RANNodeName)
	*protocolIEs.TypeValue.RANNodeName = RANNodeName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSupportedTAList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NGSetupRequestIEsTypeValuePresentSupportedTAList
	protocolIEs.TypeValue.SupportedTAList = new(ngapType.SupportedTAList)
	*protocolIEs.TypeValue.SupportedTAList = SupportedTAList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDDefaultPagingDRX
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupRequestIEsTypeValuePresentDefaultPagingDRX
	protocolIEs.TypeValue.DefaultPagingDRX = new(ngapType.PagingDRX)
	*protocolIEs.TypeValue.DefaultPagingDRX = PagingDRX()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERetentionInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupRequestIEsTypeValuePresentUERetentionInformation
	protocolIEs.TypeValue.UERetentionInformation = new(ngapType.UERetentionInformation)
	*protocolIEs.TypeValue.UERetentionInformation = UERetentionInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGSetupResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentNGSetupResponse
	msg.Value.NGSetupResponse = new(ngapType.NGSetupResponse)
	msgProtocolIEs := &msg.Value.NGSetupResponse.ProtocolIEs
	protocolIEs := ngapType.NGSetupResponseIEs{}

	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFName
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NGSetupResponseIEsTypeValuePresentAMFName
	protocolIEs.TypeValue.AMFName = new(ngapType.AMFName)
	*protocolIEs.TypeValue.AMFName = AMFName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDServedGUAMIList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NGSetupResponseIEsTypeValuePresentServedGUAMIList
	protocolIEs.TypeValue.ServedGUAMIList = new(ngapType.ServedGUAMIList)
	*protocolIEs.TypeValue.ServedGUAMIList = ServedGUAMIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRelativeAMFCapacity
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupResponseIEsTypeValuePresentRelativeAMFCapacity
	protocolIEs.TypeValue.RelativeAMFCapacity = new(ngapType.RelativeAMFCapacity)
	*protocolIEs.TypeValue.RelativeAMFCapacity = RelativeAMFCapacity()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPLMNSupportList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.NGSetupResponseIEsTypeValuePresentPLMNSupportList
	protocolIEs.TypeValue.PLMNSupportList = new(ngapType.PLMNSupportList)
	*protocolIEs.TypeValue.PLMNSupportList = PLMNSupportList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERetentionInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupResponseIEsTypeValuePresentUERetentionInformation
	protocolIEs.TypeValue.UERetentionInformation = new(ngapType.UERetentionInformation)
	*protocolIEs.TypeValue.UERetentionInformation = UERetentionInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGSetupFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentNGSetupFailure
	msg.Value.NGSetupFailure = new(ngapType.NGSetupFailure)
	msgProtocolIEs := &msg.Value.NGSetupFailure.ProtocolIEs
	protocolIEs := ngapType.NGSetupFailureIEs{}

	protocolIEs = ngapType.NGSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTimeToWait
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupFailureIEsTypeValuePresentTimeToWait
	protocolIEs.TypeValue.TimeToWait = new(ngapType.TimeToWait)
	*protocolIEs.TypeValue.TimeToWait = TimeToWait()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.NGSetupFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.NGSetupFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func OverloadStart() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeOverloadStart
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentOverloadStart
	msg.Value.OverloadStart = new(ngapType.OverloadStart)
	msgProtocolIEs := &msg.Value.OverloadStart.ProtocolIEs
	protocolIEs := ngapType.OverloadStartIEs{}

	protocolIEs = ngapType.OverloadStartIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFOverloadResponse
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.OverloadStartIEsTypeValuePresentAMFOverloadResponse
	protocolIEs.TypeValue.AMFOverloadResponse = new(ngapType.OverloadResponse)
	*protocolIEs.TypeValue.AMFOverloadResponse = OverloadResponse()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.OverloadStartIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFTrafficLoadReductionIndication
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.OverloadStartIEsTypeValuePresentAMFTrafficLoadReductionIndication
	protocolIEs.TypeValue.AMFTrafficLoadReductionIndication = new(ngapType.TrafficLoadReductionIndication)
	*protocolIEs.TypeValue.AMFTrafficLoadReductionIndication = TrafficLoadReductionIndication()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.OverloadStartIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDOverloadStartNSSAIList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.OverloadStartIEsTypeValuePresentOverloadStartNSSAIList
	protocolIEs.TypeValue.OverloadStartNSSAIList = new(ngapType.OverloadStartNSSAIList)
	*protocolIEs.TypeValue.OverloadStartNSSAIList = OverloadStartNSSAIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func OverloadStop() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeOverloadStop
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentOverloadStop
	// msg.Value.OverloadStop = new(ngapType.OverloadStop)
	// msgProtocolIEs := &msg.Value.OverloadStop.ProtocolIEs
	// protocolIEs := ngapType.OverloadStopIEs{}

	return
}

func Paging() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePaging
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPaging
	msg.Value.Paging = new(ngapType.Paging)
	msgProtocolIEs := &msg.Value.Paging.ProtocolIEs
	protocolIEs := ngapType.PagingIEs{}

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEPagingIdentity
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentUEPagingIdentity
	protocolIEs.TypeValue.UEPagingIdentity = new(ngapType.UEPagingIdentity)
	*protocolIEs.TypeValue.UEPagingIdentity = UEPagingIdentity()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPagingDRX
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentPagingDRX
	protocolIEs.TypeValue.PagingDRX = new(ngapType.PagingDRX)
	*protocolIEs.TypeValue.PagingDRX = PagingDRX()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTAIListForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentTAIListForPaging
	protocolIEs.TypeValue.TAIListForPaging = new(ngapType.TAIListForPaging)
	*protocolIEs.TypeValue.TAIListForPaging = TAIListForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentPagingPriority
	protocolIEs.TypeValue.PagingPriority = new(ngapType.PagingPriority)
	*protocolIEs.TypeValue.PagingPriority = PagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentUERadioCapabilityForPaging
	protocolIEs.TypeValue.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
	*protocolIEs.TypeValue.UERadioCapabilityForPaging = UERadioCapabilityForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPagingOrigin
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentPagingOrigin
	protocolIEs.TypeValue.PagingOrigin = new(ngapType.PagingOrigin)
	*protocolIEs.TypeValue.PagingOrigin = PagingOrigin()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAssistanceDataForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PagingIEsTypeValuePresentAssistanceDataForPaging
	protocolIEs.TypeValue.AssistanceDataForPaging = new(ngapType.AssistanceDataForPaging)
	*protocolIEs.TypeValue.AssistanceDataForPaging = AssistanceDataForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PathSwitchRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPathSwitchRequest
	msg.Value.PathSwitchRequest = new(ngapType.PathSwitchRequest)
	msgProtocolIEs := &msg.Value.PathSwitchRequest.ProtocolIEs
	protocolIEs := ngapType.PathSwitchRequestIEs{}

	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSourceAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestIEsTypeValuePresentSourceAMFUENGAPID
	protocolIEs.TypeValue.SourceAMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.SourceAMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestIEsTypeValuePresentUESecurityCapabilities
	protocolIEs.TypeValue.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.TypeValue.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceToBeSwitchedDLList
	protocolIEs.TypeValue.PDUSessionResourceToBeSwitchedDLList = new(ngapType.PDUSessionResourceToBeSwitchedDLList)
	*protocolIEs.TypeValue.PDUSessionResourceToBeSwitchedDLList = PDUSessionResourceToBeSwitchedDLList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestIEsTypeValuePresentPDUSessionResourceFailedToSetupListPSReq
	protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListPSReq = new(ngapType.PDUSessionResourceFailedToSetupListPSReq)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListPSReq = PDUSessionResourceFailedToSetupListPSReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PathSwitchRequestAcknowledge() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentPathSwitchRequestAcknowledge
	msg.Value.PathSwitchRequestAcknowledge = new(ngapType.PathSwitchRequestAcknowledge)
	msgProtocolIEs := &msg.Value.PathSwitchRequestAcknowledge.ProtocolIEs
	protocolIEs := ngapType.PathSwitchRequestAcknowledgeIEs{}

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentUESecurityCapabilities
	protocolIEs.TypeValue.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.TypeValue.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSecurityContext
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSecurityContext
	protocolIEs.TypeValue.SecurityContext = new(ngapType.SecurityContext)
	*protocolIEs.TypeValue.SecurityContext = SecurityContext()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNewSecurityContextInd
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentNewSecurityContextInd
	protocolIEs.TypeValue.NewSecurityContextInd = new(ngapType.NewSecurityContextInd)
	*protocolIEs.TypeValue.NewSecurityContextInd = NewSecurityContextInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSwitchedList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceSwitchedList
	protocolIEs.TypeValue.PDUSessionResourceSwitchedList = new(ngapType.PDUSessionResourceSwitchedList)
	*protocolIEs.TypeValue.PDUSessionResourceSwitchedList = PDUSessionResourceSwitchedList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSAck
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentPDUSessionResourceReleasedListPSAck
	protocolIEs.TypeValue.PDUSessionResourceReleasedListPSAck = new(ngapType.PDUSessionResourceReleasedListPSAck)
	*protocolIEs.TypeValue.PDUSessionResourceReleasedListPSAck = PDUSessionResourceReleasedListPSAck()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentAllowedNSSAI
	protocolIEs.TypeValue.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.TypeValue.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive
	protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = new(ngapType.CoreNetworkAssistanceInformationForInactive)
	*protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRRCInactiveTransitionReportRequest
	protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRedirectionVoiceFallback
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentRedirectionVoiceFallback
	protocolIEs.TypeValue.RedirectionVoiceFallback = new(ngapType.RedirectionVoiceFallback)
	*protocolIEs.TypeValue.RedirectionVoiceFallback = RedirectionVoiceFallback()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCNAssistedRANTuning
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentCNAssistedRANTuning
	protocolIEs.TypeValue.CNAssistedRANTuning = new(ngapType.CNAssistedRANTuning)
	*protocolIEs.TypeValue.CNAssistedRANTuning = CNAssistedRANTuning()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSRVCCOperationPossible
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestAcknowledgeIEsTypeValuePresentSRVCCOperationPossible
	protocolIEs.TypeValue.SRVCCOperationPossible = new(ngapType.SRVCCOperationPossible)
	*protocolIEs.TypeValue.SRVCCOperationPossible = SRVCCOperationPossible()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PathSwitchRequestFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentPathSwitchRequestFailure
	msg.Value.PathSwitchRequestFailure = new(ngapType.PathSwitchRequestFailure)
	msgProtocolIEs := &msg.Value.PathSwitchRequestFailure.ProtocolIEs
	protocolIEs := ngapType.PathSwitchRequestFailureIEs{}

	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestFailureIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestFailureIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSFail
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestFailureIEsTypeValuePresentPDUSessionResourceReleasedListPSFail
	protocolIEs.TypeValue.PDUSessionResourceReleasedListPSFail = new(ngapType.PDUSessionResourceReleasedListPSFail)
	*protocolIEs.TypeValue.PDUSessionResourceReleasedListPSFail = PDUSessionResourceReleasedListPSFail()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PathSwitchRequestFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModify
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPDUSessionResourceModifyRequest
	msg.Value.PDUSessionResourceModifyRequest = new(ngapType.PDUSessionResourceModifyRequest)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyRequest.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceModifyRequestIEs{}

	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentRANPagingPriority
	protocolIEs.TypeValue.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.TypeValue.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModReq
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyRequestIEsTypeValuePresentPDUSessionResourceModifyListModReq
	protocolIEs.TypeValue.PDUSessionResourceModifyListModReq = new(ngapType.PDUSessionResourceModifyListModReq)
	*protocolIEs.TypeValue.PDUSessionResourceModifyListModReq = PDUSessionResourceModifyListModReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModify
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentPDUSessionResourceModifyResponse
	msg.Value.PDUSessionResourceModifyResponse = new(ngapType.PDUSessionResourceModifyResponse)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyResponse.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceModifyResponseIEs{}

	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModRes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceModifyListModRes
	protocolIEs.TypeValue.PDUSessionResourceModifyListModRes = new(ngapType.PDUSessionResourceModifyListModRes)
	*protocolIEs.TypeValue.PDUSessionResourceModifyListModRes = PDUSessionResourceModifyListModRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModRes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentPDUSessionResourceFailedToModifyListModRes
	protocolIEs.TypeValue.PDUSessionResourceFailedToModifyListModRes = new(ngapType.PDUSessionResourceFailedToModifyListModRes)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToModifyListModRes = PDUSessionResourceFailedToModifyListModRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModifyIndication
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPDUSessionResourceModifyIndication
	msg.Value.PDUSessionResourceModifyIndication = new(ngapType.PDUSessionResourceModifyIndication)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyIndication.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceModifyIndicationIEs{}

	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentPDUSessionResourceModifyListModInd
	protocolIEs.TypeValue.PDUSessionResourceModifyListModInd = new(ngapType.PDUSessionResourceModifyListModInd)
	*protocolIEs.TypeValue.PDUSessionResourceModifyListModInd = PDUSessionResourceModifyListModInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyIndicationIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyConfirm() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModifyIndication
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentPDUSessionResourceModifyConfirm
	msg.Value.PDUSessionResourceModifyConfirm = new(ngapType.PDUSessionResourceModifyConfirm)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyConfirm.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceModifyConfirmIEs{}

	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModCfm
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceModifyListModCfm
	protocolIEs.TypeValue.PDUSessionResourceModifyListModCfm = new(ngapType.PDUSessionResourceModifyListModCfm)
	*protocolIEs.TypeValue.PDUSessionResourceModifyListModCfm = PDUSessionResourceModifyListModCfm()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentPDUSessionResourceFailedToModifyListModCfm
	protocolIEs.TypeValue.PDUSessionResourceFailedToModifyListModCfm = new(ngapType.PDUSessionResourceFailedToModifyListModCfm)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToModifyListModCfm = PDUSessionResourceFailedToModifyListModCfm()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceModifyConfirmIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceNotify() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceNotify
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPDUSessionResourceNotify
	msg.Value.PDUSessionResourceNotify = new(ngapType.PDUSessionResourceNotify)
	msgProtocolIEs := &msg.Value.PDUSessionResourceNotify.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceNotifyIEs{}

	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceNotifyIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceNotifyIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceNotifyList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceNotifyList
	protocolIEs.TypeValue.PDUSessionResourceNotifyList = new(ngapType.PDUSessionResourceNotifyList)
	*protocolIEs.TypeValue.PDUSessionResourceNotifyList = PDUSessionResourceNotifyList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListNot
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceNotifyIEsTypeValuePresentPDUSessionResourceReleasedListNot
	protocolIEs.TypeValue.PDUSessionResourceReleasedListNot = new(ngapType.PDUSessionResourceReleasedListNot)
	*protocolIEs.TypeValue.PDUSessionResourceReleasedListNot = PDUSessionResourceReleasedListNot()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceNotifyIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceReleaseCommand() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceRelease
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPDUSessionResourceReleaseCommand
	msg.Value.PDUSessionResourceReleaseCommand = new(ngapType.PDUSessionResourceReleaseCommand)
	msgProtocolIEs := &msg.Value.PDUSessionResourceReleaseCommand.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceReleaseCommandIEs{}

	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentRANPagingPriority
	protocolIEs.TypeValue.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.TypeValue.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceToReleaseListRelCmd
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseCommandIEsTypeValuePresentPDUSessionResourceToReleaseListRelCmd
	protocolIEs.TypeValue.PDUSessionResourceToReleaseListRelCmd = new(ngapType.PDUSessionResourceToReleaseListRelCmd)
	*protocolIEs.TypeValue.PDUSessionResourceToReleaseListRelCmd = PDUSessionResourceToReleaseListRelCmd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceReleaseResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceRelease
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentPDUSessionResourceReleaseResponse
	msg.Value.PDUSessionResourceReleaseResponse = new(ngapType.PDUSessionResourceReleaseResponse)
	msgProtocolIEs := &msg.Value.PDUSessionResourceReleaseResponse.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceReleaseResponseIEs{}

	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListRelRes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentPDUSessionResourceReleasedListRelRes
	protocolIEs.TypeValue.PDUSessionResourceReleasedListRelRes = new(ngapType.PDUSessionResourceReleasedListRelRes)
	*protocolIEs.TypeValue.PDUSessionResourceReleasedListRelRes = PDUSessionResourceReleasedListRelRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceReleaseResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceSetupRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPDUSessionResourceSetupRequest
	msg.Value.PDUSessionResourceSetupRequest = new(ngapType.PDUSessionResourceSetupRequest)
	msgProtocolIEs := &msg.Value.PDUSessionResourceSetupRequest.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceSetupRequestIEs{}

	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentRANPagingPriority
	protocolIEs.TypeValue.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.TypeValue.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentPDUSessionResourceSetupListSUReq
	protocolIEs.TypeValue.PDUSessionResourceSetupListSUReq = new(ngapType.PDUSessionResourceSetupListSUReq)
	*protocolIEs.TypeValue.PDUSessionResourceSetupListSUReq = PDUSessionResourceSetupListSUReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupRequestIEsTypeValuePresentUEAggregateMaximumBitRate
	protocolIEs.TypeValue.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.TypeValue.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceSetupResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceSetup
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentPDUSessionResourceSetupResponse
	msg.Value.PDUSessionResourceSetupResponse = new(ngapType.PDUSessionResourceSetupResponse)
	msgProtocolIEs := &msg.Value.PDUSessionResourceSetupResponse.ProtocolIEs
	protocolIEs := ngapType.PDUSessionResourceSetupResponseIEs{}

	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListSURes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceSetupListSURes
	protocolIEs.TypeValue.PDUSessionResourceSetupListSURes = new(ngapType.PDUSessionResourceSetupListSURes)
	*protocolIEs.TypeValue.PDUSessionResourceSetupListSURes = PDUSessionResourceSetupListSURes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListSURes
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentPDUSessionResourceFailedToSetupListSURes
	protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListSURes = new(ngapType.PDUSessionResourceFailedToSetupListSURes)
	*protocolIEs.TypeValue.PDUSessionResourceFailedToSetupListSURes = PDUSessionResourceFailedToSetupListSURes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PDUSessionResourceSetupResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PrivateMessage() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePrivateMessage
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPrivateMessage
	// msg.Value.PrivateMessage = new(ngapType.PrivateMessage)
	// msgProtocolIEs := &msg.Value.PrivateMessage.ProtocolIEs
	// protocolIEs := ngapType.PrivateMessageIEs{}

	return
}

func PWSCancelRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSCancel
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPWSCancelRequest
	msg.Value.PWSCancelRequest = new(ngapType.PWSCancelRequest)
	msgProtocolIEs := &msg.Value.PWSCancelRequest.ProtocolIEs
	protocolIEs := ngapType.PWSCancelRequestIEs{}

	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSCancelRequestIEsTypeValuePresentMessageIdentifier
	protocolIEs.TypeValue.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.TypeValue.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSCancelRequestIEsTypeValuePresentSerialNumber
	protocolIEs.TypeValue.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.TypeValue.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDWarningAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PWSCancelRequestIEsTypeValuePresentWarningAreaList
	protocolIEs.TypeValue.WarningAreaList = new(ngapType.WarningAreaList)
	*protocolIEs.TypeValue.WarningAreaList = WarningAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCancelAllWarningMessages
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSCancelRequestIEsTypeValuePresentCancelAllWarningMessages
	protocolIEs.TypeValue.CancelAllWarningMessages = new(ngapType.CancelAllWarningMessages)
	*protocolIEs.TypeValue.CancelAllWarningMessages = CancelAllWarningMessages()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSCancelResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSCancel
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentPWSCancelResponse
	msg.Value.PWSCancelResponse = new(ngapType.PWSCancelResponse)
	msgProtocolIEs := &msg.Value.PWSCancelResponse.ProtocolIEs
	protocolIEs := ngapType.PWSCancelResponseIEs{}

	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSCancelResponseIEsTypeValuePresentMessageIdentifier
	protocolIEs.TypeValue.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.TypeValue.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSCancelResponseIEsTypeValuePresentSerialNumber
	protocolIEs.TypeValue.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.TypeValue.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDBroadcastCancelledAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PWSCancelResponseIEsTypeValuePresentBroadcastCancelledAreaList
	protocolIEs.TypeValue.BroadcastCancelledAreaList = new(ngapType.BroadcastCancelledAreaList)
	*protocolIEs.TypeValue.BroadcastCancelledAreaList = BroadcastCancelledAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.PWSCancelResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSFailureIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSFailureIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPWSFailureIndication
	msg.Value.PWSFailureIndication = new(ngapType.PWSFailureIndication)
	msgProtocolIEs := &msg.Value.PWSFailureIndication.ProtocolIEs
	protocolIEs := ngapType.PWSFailureIndicationIEs{}

	protocolIEs = ngapType.PWSFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPWSFailedCellIDList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSFailureIndicationIEsTypeValuePresentPWSFailedCellIDList
	protocolIEs.TypeValue.PWSFailedCellIDList = new(ngapType.PWSFailedCellIDList)
	*protocolIEs.TypeValue.PWSFailedCellIDList = PWSFailedCellIDList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSFailureIndicationIEsTypeValuePresentGlobalRANNodeID
	protocolIEs.TypeValue.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.TypeValue.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSRestartIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSRestartIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentPWSRestartIndication
	msg.Value.PWSRestartIndication = new(ngapType.PWSRestartIndication)
	msgProtocolIEs := &msg.Value.PWSRestartIndication.ProtocolIEs
	protocolIEs := ngapType.PWSRestartIndicationIEs{}

	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCellIDListForRestart
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSRestartIndicationIEsTypeValuePresentCellIDListForRestart
	protocolIEs.TypeValue.CellIDListForRestart = new(ngapType.CellIDListForRestart)
	*protocolIEs.TypeValue.CellIDListForRestart = CellIDListForRestart()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSRestartIndicationIEsTypeValuePresentGlobalRANNodeID
	protocolIEs.TypeValue.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.TypeValue.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTAIListForRestart
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSRestartIndicationIEsTypeValuePresentTAIListForRestart
	protocolIEs.TypeValue.TAIListForRestart = new(ngapType.TAIListForRestart)
	*protocolIEs.TypeValue.TAIListForRestart = TAIListForRestart()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDEmergencyAreaIDListForRestart
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.PWSRestartIndicationIEsTypeValuePresentEmergencyAreaIDListForRestart
	protocolIEs.TypeValue.EmergencyAreaIDListForRestart = new(ngapType.EmergencyAreaIDListForRestart)
	*protocolIEs.TypeValue.EmergencyAreaIDListForRestart = EmergencyAreaIDListForRestart()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RANConfigurationUpdate() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentRANConfigurationUpdate
	msg.Value.RANConfigurationUpdate = new(ngapType.RANConfigurationUpdate)
	msgProtocolIEs := &msg.Value.RANConfigurationUpdate.ProtocolIEs
	protocolIEs := ngapType.RANConfigurationUpdateIEs{}

	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANNodeName
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateIEsTypeValuePresentRANNodeName
	protocolIEs.TypeValue.RANNodeName = new(ngapType.RANNodeName)
	*protocolIEs.TypeValue.RANNodeName = RANNodeName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSupportedTAList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateIEsTypeValuePresentSupportedTAList
	protocolIEs.TypeValue.SupportedTAList = new(ngapType.SupportedTAList)
	*protocolIEs.TypeValue.SupportedTAList = SupportedTAList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDDefaultPagingDRX
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateIEsTypeValuePresentDefaultPagingDRX
	protocolIEs.TypeValue.DefaultPagingDRX = new(ngapType.PagingDRX)
	*protocolIEs.TypeValue.DefaultPagingDRX = PagingDRX()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateIEsTypeValuePresentGlobalRANNodeID
	protocolIEs.TypeValue.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.TypeValue.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNGRANTNLAssociationToRemoveList
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateIEsTypeValuePresentNGRANTNLAssociationToRemoveList
	protocolIEs.TypeValue.NGRANTNLAssociationToRemoveList = new(ngapType.NGRANTNLAssociationToRemoveList)
	*protocolIEs.TypeValue.NGRANTNLAssociationToRemoveList = NGRANTNLAssociationToRemoveList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RANConfigurationUpdateAcknowledge() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentRANConfigurationUpdateAcknowledge
	msg.Value.RANConfigurationUpdateAcknowledge = new(ngapType.RANConfigurationUpdateAcknowledge)
	msgProtocolIEs := &msg.Value.RANConfigurationUpdateAcknowledge.ProtocolIEs
	protocolIEs := ngapType.RANConfigurationUpdateAcknowledgeIEs{}

	protocolIEs = ngapType.RANConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateAcknowledgeIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RANConfigurationUpdateFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentRANConfigurationUpdateFailure
	msg.Value.RANConfigurationUpdateFailure = new(ngapType.RANConfigurationUpdateFailure)
	msgProtocolIEs := &msg.Value.RANConfigurationUpdateFailure.ProtocolIEs
	protocolIEs := ngapType.RANConfigurationUpdateFailureIEs{}

	protocolIEs = ngapType.RANConfigurationUpdateFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RANConfigurationUpdateFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTimeToWait
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentTimeToWait
	protocolIEs.TypeValue.TimeToWait = new(ngapType.TimeToWait)
	*protocolIEs.TypeValue.TimeToWait = TimeToWait()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RANConfigurationUpdateFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RANConfigurationUpdateFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RerouteNASRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRerouteNASRequest
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentRerouteNASRequest
	msg.Value.RerouteNASRequest = new(ngapType.RerouteNASRequest)
	msgProtocolIEs := &msg.Value.RerouteNASRequest.ProtocolIEs
	protocolIEs := ngapType.RerouteNASRequestIEs{}

	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RerouteNASRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RerouteNASRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNGAPMessage
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RerouteNASRequestIEsTypeValuePresentNGAPMessage
	protocolIEs.TypeValue.NGAPMessage = new(ngapType.OctetString)
	*protocolIEs.TypeValue.NGAPMessage = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFSetID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RerouteNASRequestIEsTypeValuePresentAMFSetID
	protocolIEs.TypeValue.AMFSetID = new(ngapType.AMFSetID)
	*protocolIEs.TypeValue.AMFSetID = AMFSetID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RerouteNASRequestIEsTypeValuePresentAllowedNSSAI
	protocolIEs.TypeValue.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.TypeValue.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSourceToTargetAMFInformationReroute
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RerouteNASRequestIEsTypeValuePresentSourceToTargetAMFInformationReroute
	protocolIEs.TypeValue.SourceToTargetAMFInformationReroute = new(ngapType.SourceToTargetAMFInformationReroute)
	*protocolIEs.TypeValue.SourceToTargetAMFInformationReroute = SourceToTargetAMFInformationReroute()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RRCInactiveTransitionReport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRRCInactiveTransitionReport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentRRCInactiveTransitionReport
	msg.Value.RRCInactiveTransitionReport = new(ngapType.RRCInactiveTransitionReport)
	msgProtocolIEs := &msg.Value.RRCInactiveTransitionReport.ProtocolIEs
	protocolIEs := ngapType.RRCInactiveTransitionReportIEs{}

	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RRCInactiveTransitionReportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCState
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RRCInactiveTransitionReportIEsTypeValuePresentRRCState
	protocolIEs.TypeValue.RRCState = new(ngapType.RRCState)
	*protocolIEs.TypeValue.RRCState = RRCState()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.RRCInactiveTransitionReportIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func SecondaryRATDataUsageReport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeSecondaryRATDataUsageReport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentSecondaryRATDataUsageReport
	msg.Value.SecondaryRATDataUsageReport = new(ngapType.SecondaryRATDataUsageReport)
	msgProtocolIEs := &msg.Value.SecondaryRATDataUsageReport.ProtocolIEs
	protocolIEs := ngapType.SecondaryRATDataUsageReportIEs{}

	protocolIEs = ngapType.SecondaryRATDataUsageReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.SecondaryRATDataUsageReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.SecondaryRATDataUsageReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceSecondaryRATUsageList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentPDUSessionResourceSecondaryRATUsageList
	protocolIEs.TypeValue.PDUSessionResourceSecondaryRATUsageList = new(ngapType.PDUSessionResourceSecondaryRATUsageList)
	*protocolIEs.TypeValue.PDUSessionResourceSecondaryRATUsageList = PDUSessionResourceSecondaryRATUsageList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.SecondaryRATDataUsageReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDHandoverFlag
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentHandoverFlag
	protocolIEs.TypeValue.HandoverFlag = new(ngapType.HandoverFlag)
	*protocolIEs.TypeValue.HandoverFlag = HandoverFlag()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.SecondaryRATDataUsageReportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.SecondaryRATDataUsageReportIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func TraceFailureIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeTraceFailureIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentTraceFailureIndication
	msg.Value.TraceFailureIndication = new(ngapType.TraceFailureIndication)
	msgProtocolIEs := &msg.Value.TraceFailureIndication.ProtocolIEs
	protocolIEs := ngapType.TraceFailureIndicationIEs{}

	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.TraceFailureIndicationIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.TraceFailureIndicationIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNGRANTraceID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.TraceFailureIndicationIEsTypeValuePresentNGRANTraceID
	protocolIEs.TypeValue.NGRANTraceID = new(ngapType.NGRANTraceID)
	*protocolIEs.TypeValue.NGRANTraceID = NGRANTraceID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.TraceFailureIndicationIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func TraceStart() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeTraceStart
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentTraceStart
	msg.Value.TraceStart = new(ngapType.TraceStart)
	msgProtocolIEs := &msg.Value.TraceStart.ProtocolIEs
	protocolIEs := ngapType.TraceStartIEs{}

	protocolIEs = ngapType.TraceStartIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.TraceStartIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.TraceStartIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.TraceStartIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.TraceStartIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDTraceActivation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.TraceStartIEsTypeValuePresentTraceActivation
	protocolIEs.TypeValue.TraceActivation = new(ngapType.TraceActivation)
	*protocolIEs.TypeValue.TraceActivation = TraceActivation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextModificationRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUEContextModificationRequest
	msg.Value.UEContextModificationRequest = new(ngapType.UEContextModificationRequest)
	msgProtocolIEs := &msg.Value.UEContextModificationRequest.ProtocolIEs
	protocolIEs := ngapType.UEContextModificationRequestIEs{}

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentRANPagingPriority
	protocolIEs.TypeValue.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.TypeValue.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSecurityKey
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentSecurityKey
	protocolIEs.TypeValue.SecurityKey = new(ngapType.SecurityKey)
	*protocolIEs.TypeValue.SecurityKey = SecurityKey()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDIndexToRFSP
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentIndexToRFSP
	protocolIEs.TypeValue.IndexToRFSP = new(ngapType.IndexToRFSP)
	*protocolIEs.TypeValue.IndexToRFSP = IndexToRFSP()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentUEAggregateMaximumBitRate
	protocolIEs.TypeValue.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.TypeValue.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentUESecurityCapabilities
	protocolIEs.TypeValue.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.TypeValue.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformationForInactive
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentCoreNetworkAssistanceInformationForInactive
	protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = new(ngapType.CoreNetworkAssistanceInformationForInactive)
	*protocolIEs.TypeValue.CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDEmergencyFallbackIndicator
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentEmergencyFallbackIndicator
	protocolIEs.TypeValue.EmergencyFallbackIndicator = new(ngapType.EmergencyFallbackIndicator)
	*protocolIEs.TypeValue.EmergencyFallbackIndicator = EmergencyFallbackIndicator()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNewAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentNewAMFUENGAPID
	protocolIEs.TypeValue.NewAMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.NewAMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentRRCInactiveTransitionReportRequest
	protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.TypeValue.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNewGUAMI
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentNewGUAMI
	protocolIEs.TypeValue.NewGUAMI = new(ngapType.GUAMI)
	*protocolIEs.TypeValue.NewGUAMI = GUAMI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCNAssistedRANTuning
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentCNAssistedRANTuning
	protocolIEs.TypeValue.CNAssistedRANTuning = new(ngapType.CNAssistedRANTuning)
	*protocolIEs.TypeValue.CNAssistedRANTuning = CNAssistedRANTuning()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSRVCCOperationPossible
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationRequestIEsTypeValuePresentSRVCCOperationPossible
	protocolIEs.TypeValue.SRVCCOperationPossible = new(ngapType.SRVCCOperationPossible)
	*protocolIEs.TypeValue.SRVCCOperationPossible = SRVCCOperationPossible()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextModificationResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentUEContextModificationResponse
	msg.Value.UEContextModificationResponse = new(ngapType.UEContextModificationResponse)
	msgProtocolIEs := &msg.Value.UEContextModificationResponse.ProtocolIEs
	protocolIEs := ngapType.UEContextModificationResponseIEs{}

	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationResponseIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationResponseIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRRCState
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationResponseIEsTypeValuePresentRRCState
	protocolIEs.TypeValue.RRCState = new(ngapType.RRCState)
	*protocolIEs.TypeValue.RRCState = RRCState()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationResponseIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextModificationFailure() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.UnsuccessfulOutcomeValuePresentUEContextModificationFailure
	msg.Value.UEContextModificationFailure = new(ngapType.UEContextModificationFailure)
	msgProtocolIEs := &msg.Value.UEContextModificationFailure.ProtocolIEs
	protocolIEs := ngapType.UEContextModificationFailureIEs{}

	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationFailureIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationFailureIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationFailureIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextModificationFailureIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextReleaseCommand() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextRelease
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUEContextReleaseCommand
	msg.Value.UEContextReleaseCommand = new(ngapType.UEContextReleaseCommand)
	msgProtocolIEs := &msg.Value.UEContextReleaseCommand.ProtocolIEs
	protocolIEs := ngapType.UEContextReleaseCommandIEs{}

	protocolIEs = ngapType.UEContextReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUENGAPIDs
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCommandIEsTypeValuePresentUENGAPIDs
	protocolIEs.TypeValue.UENGAPIDs = new(ngapType.UENGAPIDs)
	*protocolIEs.TypeValue.UENGAPIDs = UENGAPIDs()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseCommandIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCommandIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextReleaseComplete() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextRelease
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentUEContextReleaseComplete
	msg.Value.UEContextReleaseComplete = new(ngapType.UEContextReleaseComplete)
	msgProtocolIEs := &msg.Value.UEContextReleaseComplete.ProtocolIEs
	protocolIEs := ngapType.UEContextReleaseCompleteIEs{}

	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCompleteIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCompleteIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCompleteIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCompleteIEsTypeValuePresentInfoOnRecommendedCellsAndRANNodesForPaging
	protocolIEs.TypeValue.InfoOnRecommendedCellsAndRANNodesForPaging = new(ngapType.InfoOnRecommendedCellsAndRANNodesForPaging)
	*protocolIEs.TypeValue.InfoOnRecommendedCellsAndRANNodesForPaging = InfoOnRecommendedCellsAndRANNodesForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceListCxtRelCpl
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCompleteIEsTypeValuePresentPDUSessionResourceListCxtRelCpl
	protocolIEs.TypeValue.PDUSessionResourceListCxtRelCpl = new(ngapType.PDUSessionResourceListCxtRelCpl)
	*protocolIEs.TypeValue.PDUSessionResourceListCxtRelCpl = PDUSessionResourceListCxtRelCpl()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseCompleteIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextReleaseRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextReleaseRequest
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUEContextReleaseRequest
	msg.Value.UEContextReleaseRequest = new(ngapType.UEContextReleaseRequest)
	msgProtocolIEs := &msg.Value.UEContextReleaseRequest.ProtocolIEs
	protocolIEs := ngapType.UEContextReleaseRequestIEs{}

	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDPDUSessionResourceListCxtRelReq
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseRequestIEsTypeValuePresentPDUSessionResourceListCxtRelReq
	protocolIEs.TypeValue.PDUSessionResourceListCxtRelReq = new(ngapType.PDUSessionResourceListCxtRelReq)
	*protocolIEs.TypeValue.PDUSessionResourceListCxtRelReq = PDUSessionResourceListCxtRelReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UEContextReleaseRequestIEsTypeValuePresentCause
	protocolIEs.TypeValue.Cause = new(ngapType.Cause)
	*protocolIEs.TypeValue.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UERadioCapabilityCheckRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityCheck
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUERadioCapabilityCheckRequest
	msg.Value.UERadioCapabilityCheckRequest = new(ngapType.UERadioCapabilityCheckRequest)
	msgProtocolIEs := &msg.Value.UERadioCapabilityCheckRequest.ProtocolIEs
	protocolIEs := ngapType.UERadioCapabilityCheckRequestIEs{}

	protocolIEs = ngapType.UERadioCapabilityCheckRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityCheckRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityCheckRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERadioCapability
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckRequestIEsTypeValuePresentUERadioCapability
	protocolIEs.TypeValue.UERadioCapability = new(ngapType.UERadioCapability)
	*protocolIEs.TypeValue.UERadioCapability = UERadioCapability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UERadioCapabilityCheckResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityCheck
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentUERadioCapabilityCheckResponse
	msg.Value.UERadioCapabilityCheckResponse = new(ngapType.UERadioCapabilityCheckResponse)
	msgProtocolIEs := &msg.Value.UERadioCapabilityCheckResponse.ProtocolIEs
	protocolIEs := ngapType.UERadioCapabilityCheckResponseIEs{}

	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDIMSVoiceSupportIndicator
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentIMSVoiceSupportIndicator
	protocolIEs.TypeValue.IMSVoiceSupportIndicator = new(ngapType.IMSVoiceSupportIndicator)
	*protocolIEs.TypeValue.IMSVoiceSupportIndicator = IMSVoiceSupportIndicator()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityCheckResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UERadioCapabilityInfoIndication() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityInfoIndication
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUERadioCapabilityInfoIndication
	msg.Value.UERadioCapabilityInfoIndication = new(ngapType.UERadioCapabilityInfoIndication)
	msgProtocolIEs := &msg.Value.UERadioCapabilityInfoIndication.ProtocolIEs
	protocolIEs := ngapType.UERadioCapabilityInfoIndicationIEs{}

	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERadioCapability
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapability
	protocolIEs.TypeValue.UERadioCapability = new(ngapType.UERadioCapability)
	*protocolIEs.TypeValue.UERadioCapability = UERadioCapability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UERadioCapabilityInfoIndicationIEsTypeValuePresentUERadioCapabilityForPaging
	protocolIEs.TypeValue.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
	*protocolIEs.TypeValue.UERadioCapabilityForPaging = UERadioCapabilityForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UETNLABindingReleaseRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUETNLABindingRelease
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUETNLABindingReleaseRequest
	msg.Value.UETNLABindingReleaseRequest = new(ngapType.UETNLABindingReleaseRequest)
	msgProtocolIEs := &msg.Value.UETNLABindingReleaseRequest.ProtocolIEs
	protocolIEs := ngapType.UETNLABindingReleaseRequestIEs{}

	protocolIEs = ngapType.UETNLABindingReleaseRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UETNLABindingReleaseRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UETNLABindingReleaseRequestIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkNASTransport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkNASTransport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUplinkNASTransport
	msg.Value.UplinkNASTransport = new(ngapType.UplinkNASTransport)
	msgProtocolIEs := &msg.Value.UplinkNASTransport.ProtocolIEs
	protocolIEs := ngapType.UplinkNASTransportIEs{}

	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkNASTransportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkNASTransportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkNASTransportIEsTypeValuePresentNASPDU
	protocolIEs.TypeValue.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.TypeValue.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UplinkNASTransportIEsTypeValuePresentUserLocationInformation
	protocolIEs.TypeValue.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.TypeValue.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkNonUEAssociatedNRPPaTransport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUplinkNonUEAssociatedNRPPaTransport
	msg.Value.UplinkNonUEAssociatedNRPPaTransport = new(ngapType.UplinkNonUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.UplinkNonUEAssociatedNRPPaTransport.ProtocolIEs
	protocolIEs := ngapType.UplinkNonUEAssociatedNRPPaTransportIEs{}

	protocolIEs = ngapType.UplinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID
	protocolIEs.TypeValue.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.TypeValue.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkNonUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU
	protocolIEs.TypeValue.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.TypeValue.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkRANConfigurationTransfer() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkRANConfigurationTransfer
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUplinkRANConfigurationTransfer
	msg.Value.UplinkRANConfigurationTransfer = new(ngapType.UplinkRANConfigurationTransfer)
	msgProtocolIEs := &msg.Value.UplinkRANConfigurationTransfer.ProtocolIEs
	protocolIEs := ngapType.UplinkRANConfigurationTransferIEs{}

	protocolIEs = ngapType.UplinkRANConfigurationTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSONConfigurationTransferUL
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentSONConfigurationTransferUL
	protocolIEs.TypeValue.SONConfigurationTransferUL = new(ngapType.SONConfigurationTransfer)
	*protocolIEs.TypeValue.SONConfigurationTransferUL = SONConfigurationTransfer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkRANConfigurationTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDENDCSONConfigurationTransferUL
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UplinkRANConfigurationTransferIEsTypeValuePresentENDCSONConfigurationTransferUL
	protocolIEs.TypeValue.ENDCSONConfigurationTransferUL = new(ngapType.ENDCSONConfigurationTransfer)
	*protocolIEs.TypeValue.ENDCSONConfigurationTransferUL = ENDCSONConfigurationTransfer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkRANStatusTransfer() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkRANStatusTransfer
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUplinkRANStatusTransfer
	msg.Value.UplinkRANStatusTransfer = new(ngapType.UplinkRANStatusTransfer)
	msgProtocolIEs := &msg.Value.UplinkRANStatusTransfer.ProtocolIEs
	protocolIEs := ngapType.UplinkRANStatusTransferIEs{}

	protocolIEs = ngapType.UplinkRANStatusTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkRANStatusTransferIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkRANStatusTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkRANStatusTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANStatusTransferTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkRANStatusTransferIEsTypeValuePresentRANStatusTransferTransparentContainer
	protocolIEs.TypeValue.RANStatusTransferTransparentContainer = new(ngapType.RANStatusTransferTransparentContainer)
	*protocolIEs.TypeValue.RANStatusTransferTransparentContainer = RANStatusTransferTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkUEAssociatedNRPPaTransport() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUplinkUEAssociatedNRPPaTransport
	msg.Value.UplinkUEAssociatedNRPPaTransport = new(ngapType.UplinkUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.UplinkUEAssociatedNRPPaTransport.ProtocolIEs
	protocolIEs := ngapType.UplinkUEAssociatedNRPPaTransportIEs{}

	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentAMFUENGAPID
	protocolIEs.TypeValue.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.TypeValue.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRANUENGAPID
	protocolIEs.TypeValue.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.TypeValue.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentRoutingID
	protocolIEs.TypeValue.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.TypeValue.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsTypeValuePresentNRPPaPDU
	protocolIEs.TypeValue.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.TypeValue.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func WriteReplaceWarningRequest() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeWriteReplaceWarning
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.InitiatingMessageValuePresentWriteReplaceWarningRequest
	msg.Value.WriteReplaceWarningRequest = new(ngapType.WriteReplaceWarningRequest)
	msgProtocolIEs := &msg.Value.WriteReplaceWarningRequest.ProtocolIEs
	protocolIEs := ngapType.WriteReplaceWarningRequestIEs{}

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentMessageIdentifier
	protocolIEs.TypeValue.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.TypeValue.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentSerialNumber
	protocolIEs.TypeValue.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.TypeValue.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDWarningAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaList
	protocolIEs.TypeValue.WarningAreaList = new(ngapType.WarningAreaList)
	*protocolIEs.TypeValue.WarningAreaList = WarningAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRepetitionPeriod
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentRepetitionPeriod
	protocolIEs.TypeValue.RepetitionPeriod = new(ngapType.RepetitionPeriod)
	*protocolIEs.TypeValue.RepetitionPeriod = RepetitionPeriod()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDNumberOfBroadcastsRequested
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentNumberOfBroadcastsRequested
	protocolIEs.TypeValue.NumberOfBroadcastsRequested = new(ngapType.NumberOfBroadcastsRequested)
	*protocolIEs.TypeValue.NumberOfBroadcastsRequested = NumberOfBroadcastsRequested()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDWarningType
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningType
	protocolIEs.TypeValue.WarningType = new(ngapType.WarningType)
	*protocolIEs.TypeValue.WarningType = WarningType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDWarningSecurityInfo
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningSecurityInfo
	protocolIEs.TypeValue.WarningSecurityInfo = new(ngapType.WarningSecurityInfo)
	*protocolIEs.TypeValue.WarningSecurityInfo = WarningSecurityInfo()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDDataCodingScheme
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentDataCodingScheme
	protocolIEs.TypeValue.DataCodingScheme = new(ngapType.DataCodingScheme)
	*protocolIEs.TypeValue.DataCodingScheme = DataCodingScheme()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDWarningMessageContents
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningMessageContents
	protocolIEs.TypeValue.WarningMessageContents = new(ngapType.WarningMessageContents)
	*protocolIEs.TypeValue.WarningMessageContents = WarningMessageContents()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDConcurrentWarningMessageInd
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentConcurrentWarningMessageInd
	protocolIEs.TypeValue.ConcurrentWarningMessageInd = new(ngapType.ConcurrentWarningMessageInd)
	*protocolIEs.TypeValue.ConcurrentWarningMessageInd = ConcurrentWarningMessageInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDWarningAreaCoordinates
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningRequestIEsTypeValuePresentWarningAreaCoordinates
	protocolIEs.TypeValue.WarningAreaCoordinates = new(ngapType.WarningAreaCoordinates)
	*protocolIEs.TypeValue.WarningAreaCoordinates = WarningAreaCoordinates()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func WriteReplaceWarningResponse() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeWriteReplaceWarning
	msg.Criticality.Value = ngapType.CriticalityReject
	msg.Value.Present = ngapType.SuccessfulOutcomeValuePresentWriteReplaceWarningResponse
	msg.Value.WriteReplaceWarningResponse = new(ngapType.WriteReplaceWarningResponse)
	msgProtocolIEs := &msg.Value.WriteReplaceWarningResponse.ProtocolIEs
	protocolIEs := ngapType.WriteReplaceWarningResponseIEs{}

	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningResponseIEsTypeValuePresentMessageIdentifier
	protocolIEs.TypeValue.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.TypeValue.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityReject
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningResponseIEsTypeValuePresentSerialNumber
	protocolIEs.TypeValue.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.TypeValue.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDBroadcastCompletedAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningResponseIEsTypeValuePresentBroadcastCompletedAreaList
	protocolIEs.TypeValue.BroadcastCompletedAreaList = new(ngapType.BroadcastCompletedAreaList)
	*protocolIEs.TypeValue.BroadcastCompletedAreaList = BroadcastCompletedAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.WriteReplaceWarningResponseIEsTypeValuePresentCriticalityDiagnostics
	protocolIEs.TypeValue.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.TypeValue.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkRIMInformationTransfer() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkRIMInformationTransfer
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentUplinkRIMInformationTransfer
	msg.Value.UplinkRIMInformationTransfer = new(ngapType.UplinkRIMInformationTransfer)
	msgProtocolIEs := &msg.Value.UplinkRIMInformationTransfer.ProtocolIEs
	protocolIEs := ngapType.UplinkRIMInformationTransferIEs{}

	protocolIEs = ngapType.UplinkRIMInformationTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRIMInformationTransfer
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.UplinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer
	protocolIEs.TypeValue.RIMInformationTransfer = new(ngapType.RIMInformationTransfer)
	*protocolIEs.TypeValue.RIMInformationTransfer = RIMInformationTransfer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkRIMInformationTransfer() (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkRIMInformationTransfer
	msg.Criticality.Value = ngapType.CriticalityIgnore
	msg.Value.Present = ngapType.InitiatingMessageValuePresentDownlinkRIMInformationTransfer
	msg.Value.DownlinkRIMInformationTransfer = new(ngapType.DownlinkRIMInformationTransfer)
	msgProtocolIEs := &msg.Value.DownlinkRIMInformationTransfer.ProtocolIEs
	protocolIEs := ngapType.DownlinkRIMInformationTransferIEs{}

	protocolIEs = ngapType.DownlinkRIMInformationTransferIEs{}
	protocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEIDRIMInformationTransfer
	protocolIEs.Criticality.Value = ngapType.CriticalityIgnore
	protocolIEs.TypeValue.Present = ngapType.DownlinkRIMInformationTransferIEsTypeValuePresentRIMInformationTransfer
	protocolIEs.TypeValue.RIMInformationTransfer = new(ngapType.RIMInformationTransfer)
	*protocolIEs.TypeValue.RIMInformationTransfer = RIMInformationTransfer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func ProcedureCode() (value ngapType.ProcedureCode) {
	value.Value = ngapType.ProcedureCodeDownlinkRIMInformationTransfer
	return
}

func TriggeringMessage() (value ngapType.TriggeringMessage) {
	value.Value = ngapType.TriggeringMessageInitiatingMessage
	return
}

func Criticality() (value ngapType.Criticality) {
	value.Value = ngapType.CriticalityNotify
	return
}

func ProtocolIEID() (value ngapType.ProtocolIEID) {
	value.Value = ngapType.ProtocolIEIDAdditionalDLForwardingUPTNLInformation
	return
}
