// Created By HaoDHH-245789 VHT2020
package ngapBuild

import "vht5gc/lib/ngap/ngapType"

func AMFConfigurationUpdate () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentAMFConfigurationUpdate
	msg.Value.AMFConfigurationUpdate = new(ngapType.AMFConfigurationUpdate)
	msgProtocolIEs := &msg.Value.AMFConfigurationUpdate.ProtocolIEs

	protocolIEs := ngapType.AMFConfigurationUpdateIEs{}

	// AMFName - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFName
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFName
	protocolIEs.Value.AMFName = new(ngapType.AMFName)
	*protocolIEs.Value.AMFName = AMFName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// ServedGUAMIList - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDServedGUAMIList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentServedGUAMIList
	protocolIEs.Value.ServedGUAMIList = new(ngapType.ServedGUAMIList)
	*protocolIEs.Value.ServedGUAMIList = ServedGUAMIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RelativeAMFCapacity - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRelativeAMFCapacity
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentRelativeAMFCapacity
	protocolIEs.Value.RelativeAMFCapacity = new(ngapType.RelativeAMFCapacity)
	*protocolIEs.Value.RelativeAMFCapacity = RelativeAMFCapacity()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PLMNSupportList - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPLMNSupportList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentPLMNSupportList
	protocolIEs.Value.PLMNSupportList = new(ngapType.PLMNSupportList)
	*protocolIEs.Value.PLMNSupportList = PLMNSupportList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFTNLAssociationToAddList - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationToAddList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToAddList
	protocolIEs.Value.AMFTNLAssociationToAddList = new(ngapType.AMFTNLAssociationToAddList)
	*protocolIEs.Value.AMFTNLAssociationToAddList = AMFTNLAssociationToAddList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFTNLAssociationToRemoveList - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationToRemoveList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToRemoveList
	protocolIEs.Value.AMFTNLAssociationToRemoveList = new(ngapType.AMFTNLAssociationToRemoveList)
	*protocolIEs.Value.AMFTNLAssociationToRemoveList = AMFTNLAssociationToRemoveList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFTNLAssociationToUpdateList - OK
	protocolIEs = ngapType.AMFConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationToUpdateList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateIEsPresentAMFTNLAssociationToUpdateList
	protocolIEs.Value.AMFTNLAssociationToUpdateList = new(ngapType.AMFTNLAssociationToUpdateList)
	*protocolIEs.Value.AMFTNLAssociationToUpdateList = AMFTNLAssociationToUpdateList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverCancel () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverCancel
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentHandoverCancel
	msg.Value.HandoverCancel = new(ngapType.HandoverCancel)
	msgProtocolIEs := &msg.Value.HandoverCancel.ProtocolIEs

	protocolIEs := ngapType.HandoverCancelIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverCancelIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCancelIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverCancelIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCancelIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.HandoverCancelIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCancelIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverRequired () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentHandoverRequired
	msg.Value.HandoverRequired = new(ngapType.HandoverRequired)
	msgProtocolIEs := &msg.Value.HandoverRequired.ProtocolIEs

	protocolIEs := ngapType.HandoverRequiredIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// HandoverType - OK
	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDHandoverType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentHandoverType
	protocolIEs.Value.HandoverType = new(ngapType.HandoverType)
	*protocolIEs.Value.HandoverType = HandoverType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TargetID - OK
	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTargetID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentTargetID
	protocolIEs.Value.TargetID = new(ngapType.TargetID)
	*protocolIEs.Value.TargetID = TargetID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// DirectForwardingPathAvailability - OK
	protocolIEs = ngapType.HandoverRequiredIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDDirectForwardingPathAvailability
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentDirectForwardingPathAvailability
	protocolIEs.Value.DirectForwardingPathAvailability = new(ngapType.DirectForwardingPathAvailability)
	*protocolIEs.Value.DirectForwardingPathAvailability = DirectForwardingPathAvailability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	//// TODO: PDUSessionResourceListHORqd
	//protocolIEs = ngapType.HandoverRequiredIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceListHORqd
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentPDUSessionResourceListHORqd
	//protocolIEs.Value.PDUSessionResourceListHORqd = new(ngapType.PDUSessionResourceListHORqd)
	//*protocolIEs.Value.PDUSessionResourceListHORqd = PDUSessionResourceListHORqd()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	//// TODO: SourceToTargetTransparentContainer
	//protocolIEs = ngapType.HandoverRequiredIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDSourceToTargetTransparentContainer
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.HandoverRequiredIEsPresentSourceToTargetTransparentContainer
	//protocolIEs.Value.SourceToTargetTransparentContainer = new(ngapType.SourceToTargetTransparentContainer)
	//*protocolIEs.Value.SourceToTargetTransparentContainer = SourceToTargetTransparentContainer()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentHandoverRequest
	msg.Value.HandoverRequest = new(ngapType.HandoverRequest)
	msgProtocolIEs := &msg.Value.HandoverRequest.ProtocolIEs

	protocolIEs := ngapType.HandoverRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// HandoverType - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDHandoverType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentHandoverType
	protocolIEs.Value.HandoverType = new(ngapType.HandoverType)
	*protocolIEs.Value.HandoverType = HandoverType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UEAggregateMaximumBitRate - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentUEAggregateMaximumBitRate
	protocolIEs.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.Value.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CoreNetworkAssistanceInformation - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentCoreNetworkAssistanceInformation
	protocolIEs.Value.CoreNetworkAssistanceInformation = new(ngapType.CoreNetworkAssistanceInformation)
	*protocolIEs.Value.CoreNetworkAssistanceInformation = CoreNetworkAssistanceInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UESecurityCapabilities - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentUESecurityCapabilities
	protocolIEs.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.Value.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SecurityContext - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSecurityContext
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentSecurityContext
	protocolIEs.Value.SecurityContext = new(ngapType.SecurityContext)
	*protocolIEs.Value.SecurityContext = SecurityContext()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NewSecurityContextInd - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNewSecurityContextInd
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentNewSecurityContextInd
	protocolIEs.Value.NewSecurityContextInd = new(ngapType.NewSecurityContextInd)
	*protocolIEs.Value.NewSecurityContextInd = NewSecurityContextInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASC - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASC
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentNASC
	protocolIEs.Value.NASC = new(ngapType.NASPDU)
	*protocolIEs.Value.NASC = NASC()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	//// TODO: PDUSessionResourceSetupListHOReq
	//protocolIEs = ngapType.HandoverRequestIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListHOReq
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentPDUSessionResourceSetupListHOReq
	//protocolIEs.Value.PDUSessionResourceSetupListHOReq = new(ngapType.PDUSessionResourceSetupListHOReq)
	//*protocolIEs.Value.PDUSessionResourceSetupListHOReq = PDUSessionResourceSetupListHOReq()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AllowedNSSAI - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentAllowedNSSAI
	protocolIEs.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.Value.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TraceActivation - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTraceActivation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentTraceActivation
	protocolIEs.Value.TraceActivation = new(ngapType.TraceActivation)
	*protocolIEs.Value.TraceActivation = TraceActivation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// MaskedIMEISV - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMaskedIMEISV
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentMaskedIMEISV
	protocolIEs.Value.MaskedIMEISV = new(ngapType.MaskedIMEISV)
	*protocolIEs.Value.MaskedIMEISV = MaskedIMEISV()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	//// TODO: SourceToTargetTransparentContainer
	//protocolIEs = ngapType.HandoverRequestIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDSourceToTargetTransparentContainer
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentSourceToTargetTransparentContainer
	//protocolIEs.Value.SourceToTargetTransparentContainer = new(ngapType.SourceToTargetTransparentContainer)
	//*protocolIEs.Value.SourceToTargetTransparentContainer = SourceToTargetTransparentContainer()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// MobilityRestrictionList - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMobilityRestrictionList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentMobilityRestrictionList
	protocolIEs.Value.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)
	*protocolIEs.Value.MobilityRestrictionList = MobilityRestrictionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// LocationReportingRequestType - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentLocationReportingRequestType
	protocolIEs.Value.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.Value.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCInactiveTransitionReportRequest - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentRRCInactiveTransitionReportRequest
	protocolIEs.Value.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.Value.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// GUAMI - OK
	protocolIEs = ngapType.HandoverRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDGUAMI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestIEsPresentGUAMI
	protocolIEs.Value.GUAMI = new(ngapType.GUAMI)
	*protocolIEs.Value.GUAMI = GUAMI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialContextSetupRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentInitialContextSetupRequest
	msg.Value.InitialContextSetupRequest = new(ngapType.InitialContextSetupRequest)
	msgProtocolIEs := &msg.Value.InitialContextSetupRequest.ProtocolIEs

	protocolIEs := ngapType.InitialContextSetupRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// OldAMF - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDOldAMF
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentOldAMF
	protocolIEs.Value.OldAMF = new(ngapType.AMFName)
	*protocolIEs.Value.OldAMF = OldAMF()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UEAggregateMaximumBitRate - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUEAggregateMaximumBitRate
	protocolIEs.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.Value.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CoreNetworkAssistanceInformation - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentCoreNetworkAssistanceInformation
	protocolIEs.Value.CoreNetworkAssistanceInformation = new(ngapType.CoreNetworkAssistanceInformation)
	*protocolIEs.Value.CoreNetworkAssistanceInformation = CoreNetworkAssistanceInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// GUAMI - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDGUAMI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentGUAMI
	protocolIEs.Value.GUAMI = new(ngapType.GUAMI)
	*protocolIEs.Value.GUAMI = GUAMI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	//// PDUSessionResourceSetupListCxtReq
	//protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtReq
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentPDUSessionResourceSetupListCxtReq
	//protocolIEs.Value.PDUSessionResourceSetupListCxtReq = new(ngapType.PDUSessionResourceSetupListCxtReq)
	//*protocolIEs.Value.PDUSessionResourceSetupListCxtReq = PDUSessionResourceSetupListCxtReq()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AllowedNSSAI - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentAllowedNSSAI
	protocolIEs.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.Value.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UESecurityCapabilities - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUESecurityCapabilities
	protocolIEs.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.Value.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SecurityKey - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSecurityKey
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentSecurityKey
	protocolIEs.Value.SecurityKey = new(ngapType.SecurityKey)
	*protocolIEs.Value.SecurityKey = SecurityKey()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TraceActivation - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTraceActivation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentTraceActivation
	protocolIEs.Value.TraceActivation = new(ngapType.TraceActivation)
	*protocolIEs.Value.TraceActivation = TraceActivation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// MobilityRestrictionList - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMobilityRestrictionList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentMobilityRestrictionList
	protocolIEs.Value.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)
	*protocolIEs.Value.MobilityRestrictionList = MobilityRestrictionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UERadioCapability - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUERadioCapability
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUERadioCapability
	protocolIEs.Value.UERadioCapability = new(ngapType.UERadioCapability)
	*protocolIEs.Value.UERadioCapability = UERadioCapability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// IndexToRFSP - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDIndexToRFSP
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentIndexToRFSP
	protocolIEs.Value.IndexToRFSP = new(ngapType.IndexToRFSP)
	*protocolIEs.Value.IndexToRFSP = IndexToRFSP()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// MaskedIMEISV - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMaskedIMEISV
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentMaskedIMEISV
	protocolIEs.Value.MaskedIMEISV = new(ngapType.MaskedIMEISV)
	*protocolIEs.Value.MaskedIMEISV = MaskedIMEISV()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// EmergencyFallbackIndicator - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDEmergencyFallbackIndicator
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentEmergencyFallbackIndicator
	protocolIEs.Value.EmergencyFallbackIndicator = new(ngapType.EmergencyFallbackIndicator)
	*protocolIEs.Value.EmergencyFallbackIndicator = EmergencyFallbackIndicator()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCInactiveTransitionReportRequest - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentRRCInactiveTransitionReportRequest
	protocolIEs.Value.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.Value.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UERadioCapabilityForPaging - OK
	protocolIEs = ngapType.InitialContextSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupRequestIEsPresentUERadioCapabilityForPaging
	protocolIEs.Value.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
	*protocolIEs.Value.UERadioCapabilityForPaging = UERadioCapabilityForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGReset () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGReset
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentNGReset
	msg.Value.NGReset = new(ngapType.NGReset)
	msgProtocolIEs := &msg.Value.NGReset.ProtocolIEs

	protocolIEs := ngapType.NGResetIEs{}

	// Cause - OK
	protocolIEs = ngapType.NGResetIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGResetIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// ResetType - OK
	protocolIEs = ngapType.NGResetIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDResetType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGResetIEsPresentResetType
	protocolIEs.Value.ResetType = new(ngapType.ResetType)
	*protocolIEs.Value.ResetType = ResetType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGSetupRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentNGSetupRequest
	msg.Value.NGSetupRequest = new(ngapType.NGSetupRequest)
	msgProtocolIEs := &msg.Value.NGSetupRequest.ProtocolIEs

	protocolIEs := ngapType.NGSetupRequestIEs{}

	// GlobalRANNodeID
	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupRequestIEsPresentGlobalRANNodeID
	protocolIEs.Value.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.Value.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANNodeName
	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANNodeName
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupRequestIEsPresentRANNodeName
	protocolIEs.Value.RANNodeName = new(ngapType.RANNodeName)
	*protocolIEs.Value.RANNodeName = RANNodeName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SupportedTAList
	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSupportedTAList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupRequestIEsPresentSupportedTAList
	protocolIEs.Value.SupportedTAList = new(ngapType.SupportedTAList)
	*protocolIEs.Value.SupportedTAList = SupportedTAList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// DefaultPagingDRX
	protocolIEs = ngapType.NGSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDDefaultPagingDRX
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupRequestIEsPresentDefaultPagingDRX
	protocolIEs.Value.DefaultPagingDRX = new(ngapType.PagingDRX)
	*protocolIEs.Value.DefaultPagingDRX = DefaultPagingDRX()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PathSwitchRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPathSwitchRequest
	msg.Value.PathSwitchRequest = new(ngapType.PathSwitchRequest)
	msgProtocolIEs := &msg.Value.PathSwitchRequest.ProtocolIEs

	protocolIEs := ngapType.PathSwitchRequestIEs{}

	// RANUENGAPID - OK
	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SourceAMFUENGAPID
	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSourceAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestIEsPresentSourceAMFUENGAPID
	protocolIEs.Value.SourceAMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.SourceAMFUENGAPID = SourceAMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UESecurityCapabilities - OK
	protocolIEs = ngapType.PathSwitchRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestIEsPresentUESecurityCapabilities
	protocolIEs.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.Value.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	//// PDUSessionResourceToBeSwitchedDLList
	//protocolIEs = ngapType.PathSwitchRequestIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.PathSwitchRequestIEsPresentPDUSessionResourceToBeSwitchedDLList
	//protocolIEs.Value.PDUSessionResourceToBeSwitchedDLList = new(ngapType.PDUSessionResourceToBeSwitchedDLList)
	//*protocolIEs.Value.PDUSessionResourceToBeSwitchedDLList = PDUSessionResourceToBeSwitchedDLList()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)
	//
	//// PDUSessionResourceFailedToSetupListPSReq
	//protocolIEs = ngapType.PathSwitchRequestIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.PathSwitchRequestIEsPresentPDUSessionResourceFailedToSetupListPSReq
	//protocolIEs.Value.PDUSessionResourceFailedToSetupListPSReq = new(ngapType.PDUSessionResourceFailedToSetupListPSReq)
	//*protocolIEs.Value.PDUSessionResourceFailedToSetupListPSReq = PDUSessionResourceFailedToSetupListPSReq()
	//msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModify
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceModifyRequest
	msg.Value.PDUSessionResourceModifyRequest = new(ngapType.PDUSessionResourceModifyRequest)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyRequest.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceModifyRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANPagingPriority
	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentRANPagingPriority
	protocolIEs.Value.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.Value.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceModifyListModReq
	protocolIEs = ngapType.PDUSessionResourceModifyRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModReq
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyRequestIEsPresentPDUSessionResourceModifyListModReq
	protocolIEs.Value.PDUSessionResourceModifyListModReq = new(ngapType.PDUSessionResourceModifyListModReq)
	*protocolIEs.Value.PDUSessionResourceModifyListModReq = PDUSessionResourceModifyListModReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModifyIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceModifyIndication
	msg.Value.PDUSessionResourceModifyIndication = new(ngapType.PDUSessionResourceModifyIndication)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyIndication.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceModifyIndicationIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyIndicationIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyIndicationIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceModifyListModInd
	protocolIEs = ngapType.PDUSessionResourceModifyIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyIndicationIEsPresentPDUSessionResourceModifyListModInd
	protocolIEs.Value.PDUSessionResourceModifyListModInd = new(ngapType.PDUSessionResourceModifyListModInd)
	*protocolIEs.Value.PDUSessionResourceModifyListModInd = PDUSessionResourceModifyListModInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceReleaseCommand () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceRelease
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceReleaseCommand
	msg.Value.PDUSessionResourceReleaseCommand = new(ngapType.PDUSessionResourceReleaseCommand)
	msgProtocolIEs := &msg.Value.PDUSessionResourceReleaseCommand.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceReleaseCommandIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANPagingPriority
	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentRANPagingPriority
	protocolIEs.Value.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.Value.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceToReleaseListRelCmd
	protocolIEs = ngapType.PDUSessionResourceReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceToReleaseListRelCmd
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseCommandIEsPresentPDUSessionResourceToReleaseListRelCmd
	protocolIEs.Value.PDUSessionResourceToReleaseListRelCmd = new(ngapType.PDUSessionResourceToReleaseListRelCmd)
	*protocolIEs.Value.PDUSessionResourceToReleaseListRelCmd = PDUSessionResourceToReleaseListRelCmd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceSetupRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceSetupRequest
	msg.Value.PDUSessionResourceSetupRequest = new(ngapType.PDUSessionResourceSetupRequest)
	msgProtocolIEs := &msg.Value.PDUSessionResourceSetupRequest.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceSetupRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANPagingPriority
	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentRANPagingPriority
	protocolIEs.Value.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.Value.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceSetupListSUReq
	protocolIEs = ngapType.PDUSessionResourceSetupRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupRequestIEsPresentPDUSessionResourceSetupListSUReq
	protocolIEs.Value.PDUSessionResourceSetupListSUReq = new(ngapType.PDUSessionResourceSetupListSUReq)
	*protocolIEs.Value.PDUSessionResourceSetupListSUReq = PDUSessionResourceSetupListSUReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSCancelRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSCancel
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPWSCancelRequest
	msg.Value.PWSCancelRequest = new(ngapType.PWSCancelRequest)
	msgProtocolIEs := &msg.Value.PWSCancelRequest.ProtocolIEs

	protocolIEs := ngapType.PWSCancelRequestIEs{}

	// MessageIdentifier
	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelRequestIEsPresentMessageIdentifier
	protocolIEs.Value.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.Value.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SerialNumber
	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelRequestIEsPresentSerialNumber
	protocolIEs.Value.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.Value.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// WarningAreaList
	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDWarningAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelRequestIEsPresentWarningAreaList
	protocolIEs.Value.WarningAreaList = new(ngapType.WarningAreaList)
	*protocolIEs.Value.WarningAreaList = WarningAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CancelAllWarningMessages
	protocolIEs = ngapType.PWSCancelRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCancelAllWarningMessages
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelRequestIEsPresentCancelAllWarningMessages
	protocolIEs.Value.CancelAllWarningMessages = new(ngapType.CancelAllWarningMessages)
	*protocolIEs.Value.CancelAllWarningMessages = CancelAllWarningMessages()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RANConfigurationUpdate () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentRANConfigurationUpdate
	msg.Value.RANConfigurationUpdate = new(ngapType.RANConfigurationUpdate)
	msgProtocolIEs := &msg.Value.RANConfigurationUpdate.ProtocolIEs

	protocolIEs := ngapType.RANConfigurationUpdateIEs{}

	// RANNodeName
	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANNodeName
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateIEsPresentRANNodeName
	protocolIEs.Value.RANNodeName = new(ngapType.RANNodeName)
	*protocolIEs.Value.RANNodeName = RANNodeName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SupportedTAList
	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSupportedTAList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateIEsPresentSupportedTAList
	protocolIEs.Value.SupportedTAList = new(ngapType.SupportedTAList)
	*protocolIEs.Value.SupportedTAList = SupportedTAList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// DefaultPagingDRX
	protocolIEs = ngapType.RANConfigurationUpdateIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDDefaultPagingDRX
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateIEsPresentDefaultPagingDRX
	protocolIEs.Value.DefaultPagingDRX = new(ngapType.PagingDRX)
	*protocolIEs.Value.DefaultPagingDRX = DefaultPagingDRX()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextModificationRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUEContextModificationRequest
	msg.Value.UEContextModificationRequest = new(ngapType.UEContextModificationRequest)
	msgProtocolIEs := &msg.Value.UEContextModificationRequest.ProtocolIEs

	protocolIEs := ngapType.UEContextModificationRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANPagingPriority
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentRANPagingPriority
	protocolIEs.Value.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.Value.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SecurityKey - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSecurityKey
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentSecurityKey
	protocolIEs.Value.SecurityKey = new(ngapType.SecurityKey)
	*protocolIEs.Value.SecurityKey = SecurityKey()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// IndexToRFSP - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDIndexToRFSP
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentIndexToRFSP
	protocolIEs.Value.IndexToRFSP = new(ngapType.IndexToRFSP)
	*protocolIEs.Value.IndexToRFSP = IndexToRFSP()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UEAggregateMaximumBitRate - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentUEAggregateMaximumBitRate
	protocolIEs.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.Value.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UESecurityCapabilities - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentUESecurityCapabilities
	protocolIEs.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.Value.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CoreNetworkAssistanceInformation - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentCoreNetworkAssistanceInformation
	protocolIEs.Value.CoreNetworkAssistanceInformation = new(ngapType.CoreNetworkAssistanceInformation)
	*protocolIEs.Value.CoreNetworkAssistanceInformation = CoreNetworkAssistanceInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// EmergencyFallbackIndicator - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDEmergencyFallbackIndicator
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentEmergencyFallbackIndicator
	protocolIEs.Value.EmergencyFallbackIndicator = new(ngapType.EmergencyFallbackIndicator)
	*protocolIEs.Value.EmergencyFallbackIndicator = EmergencyFallbackIndicator()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NewAMFUENGAPID
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNewAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentNewAMFUENGAPID
	protocolIEs.Value.NewAMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.NewAMFUENGAPID = NewAMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCInactiveTransitionReportRequest - OK
	protocolIEs = ngapType.UEContextModificationRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationRequestIEsPresentRRCInactiveTransitionReportRequest
	protocolIEs.Value.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.Value.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextReleaseCommand () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextRelease
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUEContextReleaseCommand
	msg.Value.UEContextReleaseCommand = new(ngapType.UEContextReleaseCommand)
	msgProtocolIEs := &msg.Value.UEContextReleaseCommand.ProtocolIEs

	protocolIEs := ngapType.UEContextReleaseCommandIEs{}

	// UENGAPIDs
	protocolIEs = ngapType.UEContextReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUENGAPIDs
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCommandIEsPresentUENGAPIDs
	protocolIEs.Value.UENGAPIDs = new(ngapType.UENGAPIDs)
	*protocolIEs.Value.UENGAPIDs = UENGAPIDs()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.UEContextReleaseCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCommandIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UERadioCapabilityCheckRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityCheck
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUERadioCapabilityCheckRequest
	msg.Value.UERadioCapabilityCheckRequest = new(ngapType.UERadioCapabilityCheckRequest)
	msgProtocolIEs := &msg.Value.UERadioCapabilityCheckRequest.ProtocolIEs

	protocolIEs := ngapType.UERadioCapabilityCheckRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UERadioCapabilityCheckRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UERadioCapabilityCheckRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UERadioCapability - OK
	protocolIEs = ngapType.UERadioCapabilityCheckRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUERadioCapability
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckRequestIEsPresentUERadioCapability
	protocolIEs.Value.UERadioCapability = new(ngapType.UERadioCapability)
	*protocolIEs.Value.UERadioCapability = UERadioCapability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func WriteReplaceWarningRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeWriteReplaceWarning
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentWriteReplaceWarningRequest
	msg.Value.WriteReplaceWarningRequest = new(ngapType.WriteReplaceWarningRequest)
	msgProtocolIEs := &msg.Value.WriteReplaceWarningRequest.ProtocolIEs

	protocolIEs := ngapType.WriteReplaceWarningRequestIEs{}

	// MessageIdentifier
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentMessageIdentifier
	protocolIEs.Value.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.Value.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SerialNumber
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentSerialNumber
	protocolIEs.Value.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.Value.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// WarningAreaList
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDWarningAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentWarningAreaList
	protocolIEs.Value.WarningAreaList = new(ngapType.WarningAreaList)
	*protocolIEs.Value.WarningAreaList = WarningAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RepetitionPeriod
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRepetitionPeriod
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentRepetitionPeriod
	protocolIEs.Value.RepetitionPeriod = new(ngapType.RepetitionPeriod)
	*protocolIEs.Value.RepetitionPeriod = RepetitionPeriod()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NumberOfBroadcastsRequested
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNumberOfBroadcastsRequested
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentNumberOfBroadcastsRequested
	protocolIEs.Value.NumberOfBroadcastsRequested = new(ngapType.NumberOfBroadcastsRequested)
	*protocolIEs.Value.NumberOfBroadcastsRequested = NumberOfBroadcastsRequested()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// WarningType
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDWarningType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentWarningType
	protocolIEs.Value.WarningType = new(ngapType.WarningType)
	*protocolIEs.Value.WarningType = WarningType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// WarningSecurityInfo
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDWarningSecurityInfo
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentWarningSecurityInfo
	protocolIEs.Value.WarningSecurityInfo = new(ngapType.WarningSecurityInfo)
	*protocolIEs.Value.WarningSecurityInfo = WarningSecurityInfo()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// DataCodingScheme
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDDataCodingScheme
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentDataCodingScheme
	protocolIEs.Value.DataCodingScheme = new(ngapType.DataCodingScheme)
	*protocolIEs.Value.DataCodingScheme = DataCodingScheme()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// WarningMessageContents
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDWarningMessageContents
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentWarningMessageContents
	protocolIEs.Value.WarningMessageContents = new(ngapType.WarningMessageContents)
	*protocolIEs.Value.WarningMessageContents = WarningMessageContents()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// ConcurrentWarningMessageInd
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDConcurrentWarningMessageInd
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentConcurrentWarningMessageInd
	protocolIEs.Value.ConcurrentWarningMessageInd = new(ngapType.ConcurrentWarningMessageInd)
	*protocolIEs.Value.ConcurrentWarningMessageInd = ConcurrentWarningMessageInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// WarningAreaCoordinates
	protocolIEs = ngapType.WriteReplaceWarningRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDWarningAreaCoordinates
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningRequestIEsPresentWarningAreaCoordinates
	protocolIEs.Value.WarningAreaCoordinates = new(ngapType.WarningAreaCoordinates)
	*protocolIEs.Value.WarningAreaCoordinates = WarningAreaCoordinates()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func AMFStatusIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFStatusIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentAMFStatusIndication
	msg.Value.AMFStatusIndication = new(ngapType.AMFStatusIndication)
	msgProtocolIEs := &msg.Value.AMFStatusIndication.ProtocolIEs

	protocolIEs := ngapType.AMFStatusIndicationIEs{}

	// UnavailableGUAMIList
	protocolIEs = ngapType.AMFStatusIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUnavailableGUAMIList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFStatusIndicationIEsPresentUnavailableGUAMIList
	protocolIEs.Value.UnavailableGUAMIList = new(ngapType.UnavailableGUAMIList)
	*protocolIEs.Value.UnavailableGUAMIList = UnavailableGUAMIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func CellTrafficTrace () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeCellTrafficTrace
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentCellTrafficTrace
	msg.Value.CellTrafficTrace = new(ngapType.CellTrafficTrace)
	msgProtocolIEs := &msg.Value.CellTrafficTrace.ProtocolIEs

	protocolIEs := ngapType.CellTrafficTraceIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.CellTrafficTraceIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.CellTrafficTraceIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NGRANTraceID
	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNGRANTraceID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.CellTrafficTraceIEsPresentNGRANTraceID
	protocolIEs.Value.NGRANTraceID = new(ngapType.NGRANTraceID)
	*protocolIEs.Value.NGRANTraceID = NGRANTraceID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NGRANCGI
	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNGRANCGI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.CellTrafficTraceIEsPresentNGRANCGI
	protocolIEs.Value.NGRANCGI = new(ngapType.NGRANCGI)
	*protocolIEs.Value.NGRANCGI = NGRANCGI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TraceCollectionEntityIPAddress
	protocolIEs = ngapType.CellTrafficTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTraceCollectionEntityIPAddress
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.CellTrafficTraceIEsPresentTraceCollectionEntityIPAddress
	protocolIEs.Value.TraceCollectionEntityIPAddress = new(ngapType.TransportLayerAddress)
	*protocolIEs.Value.TraceCollectionEntityIPAddress = TraceCollectionEntityIPAddress()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DeactivateTrace () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDeactivateTrace
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentDeactivateTrace
	msg.Value.DeactivateTrace = new(ngapType.DeactivateTrace)
	msgProtocolIEs := &msg.Value.DeactivateTrace.ProtocolIEs

	protocolIEs := ngapType.DeactivateTraceIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.DeactivateTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DeactivateTraceIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.DeactivateTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DeactivateTraceIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NGRANTraceID
	protocolIEs = ngapType.DeactivateTraceIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNGRANTraceID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DeactivateTraceIEsPresentNGRANTraceID
	protocolIEs.Value.NGRANTraceID = new(ngapType.NGRANTraceID)
	*protocolIEs.Value.NGRANTraceID = NGRANTraceID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkNASTransport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkNASTransport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentDownlinkNASTransport
	msg.Value.DownlinkNASTransport = new(ngapType.DownlinkNASTransport)
	msgProtocolIEs := &msg.Value.DownlinkNASTransport.ProtocolIEs

	protocolIEs := ngapType.DownlinkNASTransportIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// OldAMF - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDOldAMF
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentOldAMF
	protocolIEs.Value.OldAMF = new(ngapType.AMFName)
	*protocolIEs.Value.OldAMF = OldAMF()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANPagingPriority
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentRANPagingPriority
	protocolIEs.Value.RANPagingPriority = new(ngapType.RANPagingPriority)
	*protocolIEs.Value.RANPagingPriority = RANPagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// MobilityRestrictionList - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMobilityRestrictionList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentMobilityRestrictionList
	protocolIEs.Value.MobilityRestrictionList = new(ngapType.MobilityRestrictionList)
	*protocolIEs.Value.MobilityRestrictionList = MobilityRestrictionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// IndexToRFSP - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDIndexToRFSP
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentIndexToRFSP
	protocolIEs.Value.IndexToRFSP = new(ngapType.IndexToRFSP)
	*protocolIEs.Value.IndexToRFSP = IndexToRFSP()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UEAggregateMaximumBitRate - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEAggregateMaximumBitRate
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentUEAggregateMaximumBitRate
	protocolIEs.Value.UEAggregateMaximumBitRate = new(ngapType.UEAggregateMaximumBitRate)
	*protocolIEs.Value.UEAggregateMaximumBitRate = UEAggregateMaximumBitRate()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AllowedNSSAI - OK
	protocolIEs = ngapType.DownlinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNASTransportIEsPresentAllowedNSSAI
	protocolIEs.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.Value.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkNonUEAssociatedNRPPaTransport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport
	msg.Value.DownlinkNonUEAssociatedNRPPaTransport = new(ngapType.DownlinkNonUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.DownlinkNonUEAssociatedNRPPaTransport.ProtocolIEs

	protocolIEs := ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}

	// RoutingID
	protocolIEs = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID
	protocolIEs.Value.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.Value.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NRPPaPDU
	protocolIEs = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU
	protocolIEs.Value.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.Value.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkRANConfigurationTransfer () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkRANConfigurationTransfer
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentDownlinkRANConfigurationTransfer
	msg.Value.DownlinkRANConfigurationTransfer = new(ngapType.DownlinkRANConfigurationTransfer)
	msgProtocolIEs := &msg.Value.DownlinkRANConfigurationTransfer.ProtocolIEs

	protocolIEs := ngapType.DownlinkRANConfigurationTransferIEs{}

	// SONConfigurationTransferDL
	protocolIEs = ngapType.DownlinkRANConfigurationTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSONConfigurationTransferDL
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkRANConfigurationTransferIEsPresentSONConfigurationTransferDL
	protocolIEs.Value.SONConfigurationTransferDL = new(ngapType.SONConfigurationTransfer)
	*protocolIEs.Value.SONConfigurationTransferDL = SONConfigurationTransferDL()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkRANStatusTransfer () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkRANStatusTransfer
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentDownlinkRANStatusTransfer
	msg.Value.DownlinkRANStatusTransfer = new(ngapType.DownlinkRANStatusTransfer)
	msgProtocolIEs := &msg.Value.DownlinkRANStatusTransfer.ProtocolIEs

	protocolIEs := ngapType.DownlinkRANStatusTransferIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.DownlinkRANStatusTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkRANStatusTransferIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.DownlinkRANStatusTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkRANStatusTransferIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANStatusTransferTransparentContainer
	protocolIEs = ngapType.DownlinkRANStatusTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANStatusTransferTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer
	protocolIEs.Value.RANStatusTransferTransparentContainer = new(ngapType.RANStatusTransferTransparentContainer)
	*protocolIEs.Value.RANStatusTransferTransparentContainer = RANStatusTransferTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func DownlinkUEAssociatedNRPPaTransport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeDownlinkUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport
	msg.Value.DownlinkUEAssociatedNRPPaTransport = new(ngapType.DownlinkUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.DownlinkUEAssociatedNRPPaTransport.ProtocolIEs

	protocolIEs := ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RoutingID
	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentRoutingID
	protocolIEs.Value.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.Value.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NRPPaPDU
	protocolIEs = ngapType.DownlinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.DownlinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU
	protocolIEs.Value.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.Value.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func ErrorIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeErrorIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentErrorIndication
	msg.Value.ErrorIndication = new(ngapType.ErrorIndication)
	msgProtocolIEs := &msg.Value.ErrorIndication.ProtocolIEs

	protocolIEs := ngapType.ErrorIndicationIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.ErrorIndicationIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.ErrorIndicationIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.ErrorIndicationIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.ErrorIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.ErrorIndicationIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverNotify () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverNotification
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentHandoverNotify
	msg.Value.HandoverNotify = new(ngapType.HandoverNotify)
	msgProtocolIEs := &msg.Value.HandoverNotify.ProtocolIEs

	protocolIEs := ngapType.HandoverNotifyIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverNotifyIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverNotifyIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.HandoverNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverNotifyIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialUEMessage () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialUEMessage
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentInitialUEMessage
	msg.Value.InitialUEMessage = new(ngapType.InitialUEMessage)
	msgProtocolIEs := &msg.Value.InitialUEMessage.ProtocolIEs

	protocolIEs := ngapType.InitialUEMessageIEs{}

	// RANUENGAPID - OK
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCEstablishmentCause
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCEstablishmentCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentRRCEstablishmentCause
	protocolIEs.Value.RRCEstablishmentCause = new(ngapType.RRCEstablishmentCause)
	*protocolIEs.Value.RRCEstablishmentCause = RRCEstablishmentCause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// FiveGSTMSI
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDFiveGSTMSI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentFiveGSTMSI
	protocolIEs.Value.FiveGSTMSI = new(ngapType.FiveGSTMSI)
	*protocolIEs.Value.FiveGSTMSI = FiveGSTMSI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFSetID
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFSetID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentAMFSetID
	protocolIEs.Value.AMFSetID = new(ngapType.AMFSetID)
	*protocolIEs.Value.AMFSetID = AMFSetID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UEContextRequest
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEContextRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentUEContextRequest
	protocolIEs.Value.UEContextRequest = new(ngapType.UEContextRequest)
	*protocolIEs.Value.UEContextRequest = UEContextRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AllowedNSSAI - OK
	protocolIEs = ngapType.InitialUEMessageIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialUEMessageIEsPresentAllowedNSSAI
	protocolIEs.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.Value.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func LocationReport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeLocationReport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentLocationReport
	msg.Value.LocationReport = new(ngapType.LocationReport)
	msgProtocolIEs := &msg.Value.LocationReport.ProtocolIEs

	protocolIEs := ngapType.LocationReportIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UEPresenceInAreaOfInterestList
	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEPresenceInAreaOfInterestList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportIEsPresentUEPresenceInAreaOfInterestList
	protocolIEs.Value.UEPresenceInAreaOfInterestList = new(ngapType.UEPresenceInAreaOfInterestList)
	*protocolIEs.Value.UEPresenceInAreaOfInterestList = UEPresenceInAreaOfInterestList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// LocationReportingRequestType - OK
	protocolIEs = ngapType.LocationReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportIEsPresentLocationReportingRequestType
	protocolIEs.Value.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.Value.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func LocationReportingControl () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeLocationReportingControl
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentLocationReportingControl
	msg.Value.LocationReportingControl = new(ngapType.LocationReportingControl)
	msgProtocolIEs := &msg.Value.LocationReportingControl.ProtocolIEs

	protocolIEs := ngapType.LocationReportingControlIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.LocationReportingControlIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportingControlIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.LocationReportingControlIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportingControlIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// LocationReportingRequestType - OK
	protocolIEs = ngapType.LocationReportingControlIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDLocationReportingRequestType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportingControlIEsPresentLocationReportingRequestType
	protocolIEs.Value.LocationReportingRequestType = new(ngapType.LocationReportingRequestType)
	*protocolIEs.Value.LocationReportingRequestType = LocationReportingRequestType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func LocationReportingFailureIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeLocationReportingFailureIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentLocationReportingFailureIndication
	msg.Value.LocationReportingFailureIndication = new(ngapType.LocationReportingFailureIndication)
	msgProtocolIEs := &msg.Value.LocationReportingFailureIndication.ProtocolIEs

	protocolIEs := ngapType.LocationReportingFailureIndicationIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.LocationReportingFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportingFailureIndicationIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.LocationReportingFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportingFailureIndicationIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.LocationReportingFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.LocationReportingFailureIndicationIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NASNonDeliveryIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNASNonDeliveryIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentNASNonDeliveryIndication
	msg.Value.NASNonDeliveryIndication = new(ngapType.NASNonDeliveryIndication)
	msgProtocolIEs := &msg.Value.NASNonDeliveryIndication.ProtocolIEs

	protocolIEs := ngapType.NASNonDeliveryIndicationIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NASNonDeliveryIndicationIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NASNonDeliveryIndicationIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NASNonDeliveryIndicationIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.NASNonDeliveryIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NASNonDeliveryIndicationIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func OverloadStart () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeOverloadStart
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentOverloadStart
	msg.Value.OverloadStart = new(ngapType.OverloadStart)
	msgProtocolIEs := &msg.Value.OverloadStart.ProtocolIEs

	protocolIEs := ngapType.OverloadStartIEs{}

	// AMFOverloadResponse
	protocolIEs = ngapType.OverloadStartIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFOverloadResponse
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.OverloadStartIEsPresentAMFOverloadResponse
	protocolIEs.Value.AMFOverloadResponse = new(ngapType.OverloadResponse)
	*protocolIEs.Value.AMFOverloadResponse = AMFOverloadResponse()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFTrafficLoadReductionIndication
	protocolIEs = ngapType.OverloadStartIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFTrafficLoadReductionIndication
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.OverloadStartIEsPresentAMFTrafficLoadReductionIndication
	protocolIEs.Value.AMFTrafficLoadReductionIndication = new(ngapType.TrafficLoadReductionIndication)
	*protocolIEs.Value.AMFTrafficLoadReductionIndication = AMFTrafficLoadReductionIndication()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// OverloadStartNSSAIList
	protocolIEs = ngapType.OverloadStartIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDOverloadStartNSSAIList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.OverloadStartIEsPresentOverloadStartNSSAIList
	protocolIEs.Value.OverloadStartNSSAIList = new(ngapType.OverloadStartNSSAIList)
	*protocolIEs.Value.OverloadStartNSSAIList = OverloadStartNSSAIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func OverloadStop () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeOverloadStop
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentOverloadStop
	msg.Value.OverloadStop = new(ngapType.OverloadStop)
	msgProtocolIEs := &msg.Value.OverloadStop.ProtocolIEs

	protocolIEs := ngapType.OverloadStopIEs{}

	// Nothing
	protocolIEs = ngapType.OverloadStopIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDNothing
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.OverloadStopIEsPresentNothing
	//protocolIEs.Value.Nothing = new(ngapType.Nothing)
	//*protocolIEs.Value.Nothing = Nothing()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func Paging () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePaging
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPaging
	msg.Value.Paging = new(ngapType.Paging)
	msgProtocolIEs := &msg.Value.Paging.ProtocolIEs

	protocolIEs := ngapType.PagingIEs{}

	// UEPagingIdentity
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEPagingIdentity
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentUEPagingIdentity
	protocolIEs.Value.UEPagingIdentity = new(ngapType.UEPagingIdentity)
	*protocolIEs.Value.UEPagingIdentity = UEPagingIdentity()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PagingDRX
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPagingDRX
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentPagingDRX
	protocolIEs.Value.PagingDRX = new(ngapType.PagingDRX)
	*protocolIEs.Value.PagingDRX = PagingDRX()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TAIListForPaging
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTAIListForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentTAIListForPaging
	protocolIEs.Value.TAIListForPaging = new(ngapType.TAIListForPaging)
	*protocolIEs.Value.TAIListForPaging = TAIListForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PagingPriority
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPagingPriority
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentPagingPriority
	protocolIEs.Value.PagingPriority = new(ngapType.PagingPriority)
	*protocolIEs.Value.PagingPriority = PagingPriority()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UERadioCapability - OKForPaging - OK
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentUERadioCapabilityForPaging
	protocolIEs.Value.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
	*protocolIEs.Value.UERadioCapabilityForPaging = UERadioCapabilityForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PagingOrigin
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPagingOrigin
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentPagingOrigin
	protocolIEs.Value.PagingOrigin = new(ngapType.PagingOrigin)
	*protocolIEs.Value.PagingOrigin = PagingOrigin()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AssistanceDataForPaging
	protocolIEs = ngapType.PagingIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAssistanceDataForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PagingIEsPresentAssistanceDataForPaging
	protocolIEs.Value.AssistanceDataForPaging = new(ngapType.AssistanceDataForPaging)
	*protocolIEs.Value.AssistanceDataForPaging = AssistanceDataForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceNotify () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceNotify
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPDUSessionResourceNotify
	msg.Value.PDUSessionResourceNotify = new(ngapType.PDUSessionResourceNotify)
	msgProtocolIEs := &msg.Value.PDUSessionResourceNotify.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceNotifyIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceNotifyIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceNotifyIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceNotifyList
	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceNotifyList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceNotifyIEsPresentPDUSessionResourceNotifyList
	protocolIEs.Value.PDUSessionResourceNotifyList = new(ngapType.PDUSessionResourceNotifyList)
	*protocolIEs.Value.PDUSessionResourceNotifyList = PDUSessionResourceNotifyList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceReleasedListNot
	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListNot
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceNotifyIEsPresentPDUSessionResourceReleasedListNot
	protocolIEs.Value.PDUSessionResourceReleasedListNot = new(ngapType.PDUSessionResourceReleasedListNot)
	*protocolIEs.Value.PDUSessionResourceReleasedListNot = PDUSessionResourceReleasedListNot()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.PDUSessionResourceNotifyIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceNotifyIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PrivateMessage () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePrivateMessage
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPrivateMessage
	msg.Value.PrivateMessage = new(ngapType.PrivateMessage)
	msgProtocolIEs := &msg.Value.PrivateMessage.PrivateIEs

	protocolIEs := ngapType.PrivateMessageIEs{}

	// Nothing
	protocolIEs = ngapType.PrivateMessageIEs{}
	//protocolIEs.Id.Value = ngapType.ProtocolIEIDNothing
	//protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	//protocolIEs.Value.Present = ngapType.PrivateMessageIEsPresentNothing
	//protocolIEs.Value.Nothing = new(ngapType.Nothing)
	//*protocolIEs.Value.Nothing = Nothing()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSFailureIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSFailureIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPWSFailureIndication
	msg.Value.PWSFailureIndication = new(ngapType.PWSFailureIndication)
	msgProtocolIEs := &msg.Value.PWSFailureIndication.ProtocolIEs

	protocolIEs := ngapType.PWSFailureIndicationIEs{}

	// PWSFailedCellIDList
	protocolIEs = ngapType.PWSFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPWSFailedCellIDList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSFailureIndicationIEsPresentPWSFailedCellIDList
	protocolIEs.Value.PWSFailedCellIDList = new(ngapType.PWSFailedCellIDList)
	*protocolIEs.Value.PWSFailedCellIDList = PWSFailedCellIDList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// GlobalRANNodeID
	protocolIEs = ngapType.PWSFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSFailureIndicationIEsPresentGlobalRANNodeID
	protocolIEs.Value.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.Value.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSRestartIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSRestartIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentPWSRestartIndication
	msg.Value.PWSRestartIndication = new(ngapType.PWSRestartIndication)
	msgProtocolIEs := &msg.Value.PWSRestartIndication.ProtocolIEs

	protocolIEs := ngapType.PWSRestartIndicationIEs{}

	// CellIDListForRestart
	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCellIDListForRestart
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSRestartIndicationIEsPresentCellIDListForRestart
	protocolIEs.Value.CellIDListForRestart = new(ngapType.CellIDListForRestart)
	*protocolIEs.Value.CellIDListForRestart = CellIDListForRestart()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// GlobalRANNodeID
	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDGlobalRANNodeID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSRestartIndicationIEsPresentGlobalRANNodeID
	protocolIEs.Value.GlobalRANNodeID = new(ngapType.GlobalRANNodeID)
	*protocolIEs.Value.GlobalRANNodeID = GlobalRANNodeID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TAIListForRestart
	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTAIListForRestart
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSRestartIndicationIEsPresentTAIListForRestart
	protocolIEs.Value.TAIListForRestart = new(ngapType.TAIListForRestart)
	*protocolIEs.Value.TAIListForRestart = TAIListForRestart()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// EmergencyAreaIDListForRestart
	protocolIEs = ngapType.PWSRestartIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDEmergencyAreaIDListForRestart
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSRestartIndicationIEsPresentEmergencyAreaIDListForRestart
	protocolIEs.Value.EmergencyAreaIDListForRestart = new(ngapType.EmergencyAreaIDListForRestart)
	*protocolIEs.Value.EmergencyAreaIDListForRestart = EmergencyAreaIDListForRestart()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RerouteNASRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRerouteNASRequest
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentRerouteNASRequest
	msg.Value.RerouteNASRequest = new(ngapType.RerouteNASRequest)
	msgProtocolIEs := &msg.Value.RerouteNASRequest.ProtocolIEs

	protocolIEs := ngapType.RerouteNASRequestIEs{}

	// RANUENGAPID - OK
	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RerouteNASRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFUENGAPID - OK
	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RerouteNASRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NGAPMessage
	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNGAPMessage
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RerouteNASRequestIEsPresentNGAPMessage
	protocolIEs.Value.NGAPMessage = new(ngapType.OctetString)
	*protocolIEs.Value.NGAPMessage = NGAPMessage()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFSetID
	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFSetID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RerouteNASRequestIEsPresentAMFSetID
	protocolIEs.Value.AMFSetID = new(ngapType.AMFSetID)
	*protocolIEs.Value.AMFSetID = AMFSetID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AllowedNSSAI - OK
	protocolIEs = ngapType.RerouteNASRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RerouteNASRequestIEsPresentAllowedNSSAI
	protocolIEs.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.Value.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RRCInactiveTransitionReport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRRCInactiveTransitionReport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentRRCInactiveTransitionReport
	msg.Value.RRCInactiveTransitionReport = new(ngapType.RRCInactiveTransitionReport)
	msgProtocolIEs := &msg.Value.RRCInactiveTransitionReport.ProtocolIEs

	protocolIEs := ngapType.RRCInactiveTransitionReportIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RRCInactiveTransitionReportIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RRCInactiveTransitionReportIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCState
	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCState
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RRCInactiveTransitionReportIEsPresentRRCState
	protocolIEs.Value.RRCState = new(ngapType.RRCState)
	*protocolIEs.Value.RRCState = RRCState()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.RRCInactiveTransitionReportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RRCInactiveTransitionReportIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func TraceFailureIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeTraceFailureIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentTraceFailureIndication
	msg.Value.TraceFailureIndication = new(ngapType.TraceFailureIndication)
	msgProtocolIEs := &msg.Value.TraceFailureIndication.ProtocolIEs

	protocolIEs := ngapType.TraceFailureIndicationIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceFailureIndicationIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceFailureIndicationIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NGRANTraceID
	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNGRANTraceID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceFailureIndicationIEsPresentNGRANTraceID
	protocolIEs.Value.NGRANTraceID = new(ngapType.NGRANTraceID)
	*protocolIEs.Value.NGRANTraceID = NGRANTraceID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.TraceFailureIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceFailureIndicationIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func TraceStart () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeTraceStart
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentTraceStart
	msg.Value.TraceStart = new(ngapType.TraceStart)
	msgProtocolIEs := &msg.Value.TraceStart.ProtocolIEs

	protocolIEs := ngapType.TraceStartIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.TraceStartIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceStartIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.TraceStartIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceStartIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TraceActivation - OK
	protocolIEs = ngapType.TraceStartIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTraceActivation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.TraceStartIEsPresentTraceActivation
	protocolIEs.Value.TraceActivation = new(ngapType.TraceActivation)
	*protocolIEs.Value.TraceActivation = TraceActivation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextReleaseRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextReleaseRequest
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUEContextReleaseRequest
	msg.Value.UEContextReleaseRequest = new(ngapType.UEContextReleaseRequest)
	msgProtocolIEs := &msg.Value.UEContextReleaseRequest.ProtocolIEs

	protocolIEs := ngapType.UEContextReleaseRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceListCxtRelReq
	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceListCxtRelReq
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseRequestIEsPresentPDUSessionResourceListCxtRelReq
	protocolIEs.Value.PDUSessionResourceListCxtRelReq = new(ngapType.PDUSessionResourceListCxtRelReq)
	*protocolIEs.Value.PDUSessionResourceListCxtRelReq = PDUSessionResourceListCxtRelReq()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.UEContextReleaseRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseRequestIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UERadioCapabilityInfoIndication () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityInfoIndication
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUERadioCapabilityInfoIndication
	msg.Value.UERadioCapabilityInfoIndication = new(ngapType.UERadioCapabilityInfoIndication)
	msgProtocolIEs := &msg.Value.UERadioCapabilityInfoIndication.ProtocolIEs

	protocolIEs := ngapType.UERadioCapabilityInfoIndicationIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityInfoIndicationIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityInfoIndicationIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UERadioCapability - OK
	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUERadioCapability
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityInfoIndicationIEsPresentUERadioCapability
	protocolIEs.Value.UERadioCapability = new(ngapType.UERadioCapability)
	*protocolIEs.Value.UERadioCapability = UERadioCapability()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UERadioCapability - OKForPaging - OK
	protocolIEs = ngapType.UERadioCapabilityInfoIndicationIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUERadioCapabilityForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityInfoIndicationIEsPresentUERadioCapabilityForPaging
	protocolIEs.Value.UERadioCapabilityForPaging = new(ngapType.UERadioCapabilityForPaging)
	*protocolIEs.Value.UERadioCapabilityForPaging = UERadioCapabilityForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UETNLABindingReleaseRequest () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUETNLABindingRelease
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUETNLABindingReleaseRequest
	msg.Value.UETNLABindingReleaseRequest = new(ngapType.UETNLABindingReleaseRequest)
	msgProtocolIEs := &msg.Value.UETNLABindingReleaseRequest.ProtocolIEs

	protocolIEs := ngapType.UETNLABindingReleaseRequestIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UETNLABindingReleaseRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UETNLABindingReleaseRequestIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UETNLABindingReleaseRequestIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UETNLABindingReleaseRequestIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkNASTransport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkNASTransport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUplinkNASTransport
	msg.Value.UplinkNASTransport = new(ngapType.UplinkNASTransport)
	msgProtocolIEs := &msg.Value.UplinkNASTransport.ProtocolIEs

	protocolIEs := ngapType.UplinkNASTransportIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkNASTransportIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkNASTransportIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASPDU - OK
	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkNASTransportIEsPresentNASPDU
	protocolIEs.Value.NASPDU = new(ngapType.NASPDU)
	*protocolIEs.Value.NASPDU = NASPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.UplinkNASTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkNASTransportIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkNonUEAssociatedNRPPaTransport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport
	msg.Value.UplinkNonUEAssociatedNRPPaTransport = new(ngapType.UplinkNonUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.UplinkNonUEAssociatedNRPPaTransport.ProtocolIEs

	protocolIEs := ngapType.UplinkNonUEAssociatedNRPPaTransportIEs{}

	// RoutingID
	protocolIEs = ngapType.UplinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkNonUEAssociatedNRPPaTransportIEsPresentRoutingID
	protocolIEs.Value.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.Value.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NRPPaPDU
	protocolIEs = ngapType.UplinkNonUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkNonUEAssociatedNRPPaTransportIEsPresentNRPPaPDU
	protocolIEs.Value.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.Value.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkRANConfigurationTransfer () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkRANConfigurationTransfer
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUplinkRANConfigurationTransfer
	msg.Value.UplinkRANConfigurationTransfer = new(ngapType.UplinkRANConfigurationTransfer)
	msgProtocolIEs := &msg.Value.UplinkRANConfigurationTransfer.ProtocolIEs

	protocolIEs := ngapType.UplinkRANConfigurationTransferIEs{}

	// SONConfigurationTransferUL
	protocolIEs = ngapType.UplinkRANConfigurationTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSONConfigurationTransferUL
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkRANConfigurationTransferIEsPresentSONConfigurationTransferUL
	protocolIEs.Value.SONConfigurationTransferUL = new(ngapType.SONConfigurationTransfer)
	*protocolIEs.Value.SONConfigurationTransferUL = SONConfigurationTransferUL()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkRANStatusTransfer () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkRANStatusTransfer
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUplinkRANStatusTransfer
	msg.Value.UplinkRANStatusTransfer = new(ngapType.UplinkRANStatusTransfer)
	msgProtocolIEs := &msg.Value.UplinkRANStatusTransfer.ProtocolIEs

	protocolIEs := ngapType.UplinkRANStatusTransferIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UplinkRANStatusTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkRANStatusTransferIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UplinkRANStatusTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkRANStatusTransferIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANStatusTransferTransparentContainer
	protocolIEs = ngapType.UplinkRANStatusTransferIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANStatusTransferTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkRANStatusTransferIEsPresentRANStatusTransferTransparentContainer
	protocolIEs.Value.RANStatusTransferTransparentContainer = new(ngapType.RANStatusTransferTransparentContainer)
	*protocolIEs.Value.RANStatusTransferTransparentContainer = RANStatusTransferTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UplinkUEAssociatedNRPPaTransport () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentInitiatingMessage
	pdu.InitiatingMessage = new(ngapType.InitiatingMessage)
	msg := pdu.InitiatingMessage
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport
	msg.Value.UplinkUEAssociatedNRPPaTransport = new(ngapType.UplinkUEAssociatedNRPPaTransport)
	msgProtocolIEs := &msg.Value.UplinkUEAssociatedNRPPaTransport.ProtocolIEs

	protocolIEs := ngapType.UplinkUEAssociatedNRPPaTransportIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RoutingID
	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRoutingID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentRoutingID
	protocolIEs.Value.RoutingID = new(ngapType.RoutingID)
	*protocolIEs.Value.RoutingID = RoutingID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NRPPaPDU
	protocolIEs = ngapType.UplinkUEAssociatedNRPPaTransportIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNRPPaPDU
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UplinkUEAssociatedNRPPaTransportIEsPresentNRPPaPDU
	protocolIEs.Value.NRPPaPDU = new(ngapType.NRPPaPDU)
	*protocolIEs.Value.NRPPaPDU = NRPPaPDU()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func AMFConfigurationUpdateAcknowledge () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentAMFConfigurationUpdateAcknowledge
	msg.Value.AMFConfigurationUpdateAcknowledge = new(ngapType.AMFConfigurationUpdateAcknowledge)
	msgProtocolIEs := &msg.Value.AMFConfigurationUpdateAcknowledge.ProtocolIEs

	protocolIEs := ngapType.AMFConfigurationUpdateAcknowledgeIEs{}

	// AMFTNLAssociationSetupList
	protocolIEs = ngapType.AMFConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationSetupList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationSetupList
	protocolIEs.Value.AMFTNLAssociationSetupList = new(ngapType.AMFTNLAssociationSetupList)
	*protocolIEs.Value.AMFTNLAssociationSetupList = AMFTNLAssociationSetupList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AMFTNLAssociationFailedToSetupList
	protocolIEs = ngapType.AMFConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFTNLAssociationFailedToSetupList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentAMFTNLAssociationFailedToSetupList
	protocolIEs.Value.AMFTNLAssociationFailedToSetupList = new(ngapType.TNLAssociationList)
	*protocolIEs.Value.AMFTNLAssociationFailedToSetupList = AMFTNLAssociationFailedToSetupList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.AMFConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverCancelAcknowledge () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverCancel
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentHandoverCancelAcknowledge
	msg.Value.HandoverCancelAcknowledge = new(ngapType.HandoverCancelAcknowledge)
	msgProtocolIEs := &msg.Value.HandoverCancelAcknowledge.ProtocolIEs

	protocolIEs := ngapType.HandoverCancelAcknowledgeIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverCancelAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverCancelAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.HandoverCancelAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCancelAcknowledgeIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverCommand () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentHandoverCommand
	msg.Value.HandoverCommand = new(ngapType.HandoverCommand)
	msgProtocolIEs := &msg.Value.HandoverCommand.ProtocolIEs

	protocolIEs := ngapType.HandoverCommandIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// HandoverType - OK
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDHandoverType
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentHandoverType
	protocolIEs.Value.HandoverType = new(ngapType.HandoverType)
	*protocolIEs.Value.HandoverType = HandoverType()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NASSecurityParametersFromNGRAN
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNASSecurityParametersFromNGRAN
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentNASSecurityParametersFromNGRAN
	protocolIEs.Value.NASSecurityParametersFromNGRAN = new(ngapType.NASSecurityParametersFromNGRAN)
	*protocolIEs.Value.NASSecurityParametersFromNGRAN = NASSecurityParametersFromNGRAN()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceHandoverList
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceHandoverList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentPDUSessionResourceHandoverList
	protocolIEs.Value.PDUSessionResourceHandoverList = new(ngapType.PDUSessionResourceHandoverList)
	*protocolIEs.Value.PDUSessionResourceHandoverList = PDUSessionResourceHandoverList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceToReleaseListHOCmd
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentPDUSessionResourceToReleaseListHOCmd
	protocolIEs.Value.PDUSessionResourceToReleaseListHOCmd = new(ngapType.PDUSessionResourceToReleaseListHOCmd)
	*protocolIEs.Value.PDUSessionResourceToReleaseListHOCmd = PDUSessionResourceToReleaseListHOCmd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TargetToSourceTransparentContainer
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTargetToSourceTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentTargetToSourceTransparentContainer
	protocolIEs.Value.TargetToSourceTransparentContainer = new(ngapType.TargetToSourceTransparentContainer)
	*protocolIEs.Value.TargetToSourceTransparentContainer = TargetToSourceTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.HandoverCommandIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverCommandIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverRequestAcknowledge () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentHandoverRequestAcknowledge
	msg.Value.HandoverRequestAcknowledge = new(ngapType.HandoverRequestAcknowledge)
	msgProtocolIEs := &msg.Value.HandoverRequestAcknowledge.ProtocolIEs

	protocolIEs := ngapType.HandoverRequestAcknowledgeIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestAcknowledgeIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestAcknowledgeIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceAdmittedList
	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceAdmittedList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestAcknowledgeIEsPresentPDUSessionResourceAdmittedList
	protocolIEs.Value.PDUSessionResourceAdmittedList = new(ngapType.PDUSessionResourceAdmittedList)
	*protocolIEs.Value.PDUSessionResourceAdmittedList = PDUSessionResourceAdmittedList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceFailedToSetupListHOAck
	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestAcknowledgeIEsPresentPDUSessionResourceFailedToSetupListHOAck
	protocolIEs.Value.PDUSessionResourceFailedToSetupListHOAck = new(ngapType.PDUSessionResourceFailedToSetupListHOAck)
	*protocolIEs.Value.PDUSessionResourceFailedToSetupListHOAck = PDUSessionResourceFailedToSetupListHOAck()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TargetToSourceTransparentContainer
	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTargetToSourceTransparentContainer
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestAcknowledgeIEsPresentTargetToSourceTransparentContainer
	protocolIEs.Value.TargetToSourceTransparentContainer = new(ngapType.TargetToSourceTransparentContainer)
	*protocolIEs.Value.TargetToSourceTransparentContainer = TargetToSourceTransparentContainer()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.HandoverRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverRequestAcknowledgeIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialContextSetupResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentInitialContextSetupResponse
	msg.Value.InitialContextSetupResponse = new(ngapType.InitialContextSetupResponse)
	msgProtocolIEs := &msg.Value.InitialContextSetupResponse.ProtocolIEs

	protocolIEs := ngapType.InitialContextSetupResponseIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupResponseIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupResponseIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceSetupListCxtRes
	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtRes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupResponseIEsPresentPDUSessionResourceSetupListCxtRes
	protocolIEs.Value.PDUSessionResourceSetupListCxtRes = new(ngapType.PDUSessionResourceSetupListCxtRes)
	*protocolIEs.Value.PDUSessionResourceSetupListCxtRes = PDUSessionResourceSetupListCxtRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceFailedToSetupListCxtRes
	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupResponseIEsPresentPDUSessionResourceFailedToSetupListCxtRes
	protocolIEs.Value.PDUSessionResourceFailedToSetupListCxtRes = new(ngapType.PDUSessionResourceFailedToSetupListCxtRes)
	*protocolIEs.Value.PDUSessionResourceFailedToSetupListCxtRes = PDUSessionResourceFailedToSetupListCxtRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.InitialContextSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGResetAcknowledge () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGReset
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentNGResetAcknowledge
	msg.Value.NGResetAcknowledge = new(ngapType.NGResetAcknowledge)
	msgProtocolIEs := &msg.Value.NGResetAcknowledge.ProtocolIEs

	protocolIEs := ngapType.NGResetAcknowledgeIEs{}

	// UEAssociatedLogicalNGConnectionList
	protocolIEs = ngapType.NGResetAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUEAssociatedLogicalNGConnectionList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGResetAcknowledgeIEsPresentUEAssociatedLogicalNGConnectionList
	protocolIEs.Value.UEAssociatedLogicalNGConnectionList = new(ngapType.UEAssociatedLogicalNGConnectionList)
	*protocolIEs.Value.UEAssociatedLogicalNGConnectionList = UEAssociatedLogicalNGConnectionList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.NGResetAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGResetAcknowledgeIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGSetupResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentNGSetupResponse
	msg.Value.NGSetupResponse = new(ngapType.NGSetupResponse)
	msgProtocolIEs := &msg.Value.NGSetupResponse.ProtocolIEs

	protocolIEs := ngapType.NGSetupResponseIEs{}

	// AMFName - OK
	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFName
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupResponseIEsPresentAMFName
	protocolIEs.Value.AMFName = new(ngapType.AMFName)
	*protocolIEs.Value.AMFName = AMFName()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// ServedGUAMIList - OK
	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDServedGUAMIList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupResponseIEsPresentServedGUAMIList
	protocolIEs.Value.ServedGUAMIList = new(ngapType.ServedGUAMIList)
	*protocolIEs.Value.ServedGUAMIList = ServedGUAMIList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RelativeAMFCapacity - OK
	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRelativeAMFCapacity
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupResponseIEsPresentRelativeAMFCapacity
	protocolIEs.Value.RelativeAMFCapacity = new(ngapType.RelativeAMFCapacity)
	*protocolIEs.Value.RelativeAMFCapacity = RelativeAMFCapacity()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PLMNSupportList - OK
	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPLMNSupportList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupResponseIEsPresentPLMNSupportList
	protocolIEs.Value.PLMNSupportList = new(ngapType.PLMNSupportList)
	*protocolIEs.Value.PLMNSupportList = PLMNSupportList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.NGSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PathSwitchRequestAcknowledge () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentPathSwitchRequestAcknowledge
	msg.Value.PathSwitchRequestAcknowledge = new(ngapType.PathSwitchRequestAcknowledge)
	msgProtocolIEs := &msg.Value.PathSwitchRequestAcknowledge.ProtocolIEs

	protocolIEs := ngapType.PathSwitchRequestAcknowledgeIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UESecurityCapabilities - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUESecurityCapabilities
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentUESecurityCapabilities
	protocolIEs.Value.UESecurityCapabilities = new(ngapType.UESecurityCapabilities)
	*protocolIEs.Value.UESecurityCapabilities = UESecurityCapabilities()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SecurityContext - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSecurityContext
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentSecurityContext
	protocolIEs.Value.SecurityContext = new(ngapType.SecurityContext)
	*protocolIEs.Value.SecurityContext = SecurityContext()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// NewSecurityContextInd - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDNewSecurityContextInd
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentNewSecurityContextInd
	protocolIEs.Value.NewSecurityContextInd = new(ngapType.NewSecurityContextInd)
	*protocolIEs.Value.NewSecurityContextInd = NewSecurityContextInd()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceSwitchedList
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSwitchedList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceSwitchedList
	protocolIEs.Value.PDUSessionResourceSwitchedList = new(ngapType.PDUSessionResourceSwitchedList)
	*protocolIEs.Value.PDUSessionResourceSwitchedList = PDUSessionResourceSwitchedList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceReleasedListPSAck
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSAck
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentPDUSessionResourceReleasedListPSAck
	protocolIEs.Value.PDUSessionResourceReleasedListPSAck = new(ngapType.PDUSessionResourceReleasedListPSAck)
	*protocolIEs.Value.PDUSessionResourceReleasedListPSAck = PDUSessionResourceReleasedListPSAck()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// AllowedNSSAI - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAllowedNSSAI
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentAllowedNSSAI
	protocolIEs.Value.AllowedNSSAI = new(ngapType.AllowedNSSAI)
	*protocolIEs.Value.AllowedNSSAI = AllowedNSSAI()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CoreNetworkAssistanceInformation - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCoreNetworkAssistanceInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentCoreNetworkAssistanceInformation
	protocolIEs.Value.CoreNetworkAssistanceInformation = new(ngapType.CoreNetworkAssistanceInformation)
	*protocolIEs.Value.CoreNetworkAssistanceInformation = CoreNetworkAssistanceInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCInactiveTransitionReportRequest - OK
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentRRCInactiveTransitionReportRequest
	protocolIEs.Value.RRCInactiveTransitionReportRequest = new(ngapType.RRCInactiveTransitionReportRequest)
	*protocolIEs.Value.RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PathSwitchRequestAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestAcknowledgeIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModify
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyResponse
	msg.Value.PDUSessionResourceModifyResponse = new(ngapType.PDUSessionResourceModifyResponse)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyResponse.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceModifyResponseIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyResponseIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyResponseIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceModifyListModRes
	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModRes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceModifyListModRes
	protocolIEs.Value.PDUSessionResourceModifyListModRes = new(ngapType.PDUSessionResourceModifyListModRes)
	*protocolIEs.Value.PDUSessionResourceModifyListModRes = PDUSessionResourceModifyListModRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceFailedToModifyListModRes
	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModRes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyResponseIEsPresentPDUSessionResourceFailedToModifyListModRes
	protocolIEs.Value.PDUSessionResourceFailedToModifyListModRes = new(ngapType.PDUSessionResourceFailedToModifyListModRes)
	*protocolIEs.Value.PDUSessionResourceFailedToModifyListModRes = PDUSessionResourceFailedToModifyListModRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyResponseIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PDUSessionResourceModifyResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceModifyConfirm () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceModify
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentPDUSessionResourceModifyConfirm
	msg.Value.PDUSessionResourceModifyConfirm = new(ngapType.PDUSessionResourceModifyConfirm)
	msgProtocolIEs := &msg.Value.PDUSessionResourceModifyConfirm.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceModifyConfirmIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceModifyListModCfm
	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceModifyListModCfm
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceModifyListModCfm
	protocolIEs.Value.PDUSessionResourceModifyListModCfm = new(ngapType.PDUSessionResourceModifyListModCfm)
	*protocolIEs.Value.PDUSessionResourceModifyListModCfm = PDUSessionResourceModifyListModCfm()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceFailedToModifyListModCfm
	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentPDUSessionResourceFailedToModifyListModCfm
	protocolIEs.Value.PDUSessionResourceFailedToModifyListModCfm = new(ngapType.PDUSessionResourceFailedToModifyListModCfm)
	*protocolIEs.Value.PDUSessionResourceFailedToModifyListModCfm = PDUSessionResourceFailedToModifyListModCfm()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PDUSessionResourceModifyConfirmIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceModifyConfirmIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceReleaseResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceRelease
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentPDUSessionResourceReleaseResponse
	msg.Value.PDUSessionResourceReleaseResponse = new(ngapType.PDUSessionResourceReleaseResponse)
	msgProtocolIEs := &msg.Value.PDUSessionResourceReleaseResponse.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceReleaseResponseIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseResponseIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseResponseIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceReleasedListRelRes
	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListRelRes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseResponseIEsPresentPDUSessionResourceReleasedListRelRes
	protocolIEs.Value.PDUSessionResourceReleasedListRelRes = new(ngapType.PDUSessionResourceReleasedListRelRes)
	*protocolIEs.Value.PDUSessionResourceReleasedListRelRes = PDUSessionResourceReleasedListRelRes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseResponseIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PDUSessionResourceReleaseResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceReleaseResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PDUSessionResourceSetupResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePDUSessionResourceSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentPDUSessionResourceSetupResponse
	msg.Value.PDUSessionResourceSetupResponse = new(ngapType.PDUSessionResourceSetupResponse)
	msgProtocolIEs := &msg.Value.PDUSessionResourceSetupResponse.ProtocolIEs

	protocolIEs := ngapType.PDUSessionResourceSetupResponseIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupResponseIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupResponseIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceSetupListSURes
	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceSetupListSURes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceSetupListSURes
	protocolIEs.Value.PDUSessionResourceSetupListSURes = new(ngapType.PDUSessionResourceSetupListSURes)
	*protocolIEs.Value.PDUSessionResourceSetupListSURes = PDUSessionResourceSetupListSURes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceFailedToSetupListSURes
	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListSURes
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupResponseIEsPresentPDUSessionResourceFailedToSetupListSURes
	protocolIEs.Value.PDUSessionResourceFailedToSetupListSURes = new(ngapType.PDUSessionResourceFailedToSetupListSURes)
	*protocolIEs.Value.PDUSessionResourceFailedToSetupListSURes = PDUSessionResourceFailedToSetupListSURes()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PDUSessionResourceSetupResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PDUSessionResourceSetupResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PWSCancelResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePWSCancel
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentPWSCancelResponse
	msg.Value.PWSCancelResponse = new(ngapType.PWSCancelResponse)
	msgProtocolIEs := &msg.Value.PWSCancelResponse.ProtocolIEs

	protocolIEs := ngapType.PWSCancelResponseIEs{}

	// MessageIdentifier
	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelResponseIEsPresentMessageIdentifier
	protocolIEs.Value.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.Value.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SerialNumber
	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelResponseIEsPresentSerialNumber
	protocolIEs.Value.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.Value.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// BroadcastCancelledAreaList
	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDBroadcastCancelledAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelResponseIEsPresentBroadcastCancelledAreaList
	protocolIEs.Value.BroadcastCancelledAreaList = new(ngapType.BroadcastCancelledAreaList)
	*protocolIEs.Value.BroadcastCancelledAreaList = BroadcastCancelledAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PWSCancelResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PWSCancelResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RANConfigurationUpdateAcknowledge () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentRANConfigurationUpdateAcknowledge
	msg.Value.RANConfigurationUpdateAcknowledge = new(ngapType.RANConfigurationUpdateAcknowledge)
	msgProtocolIEs := &msg.Value.RANConfigurationUpdateAcknowledge.ProtocolIEs

	protocolIEs := ngapType.RANConfigurationUpdateAcknowledgeIEs{}

	// CriticalityDiagnostics
	protocolIEs = ngapType.RANConfigurationUpdateAcknowledgeIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateAcknowledgeIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextModificationResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentUEContextModificationResponse
	msg.Value.UEContextModificationResponse = new(ngapType.UEContextModificationResponse)
	msgProtocolIEs := &msg.Value.UEContextModificationResponse.ProtocolIEs

	protocolIEs := ngapType.UEContextModificationResponseIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationResponseIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationResponseIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RRCState
	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRRCState
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationResponseIEsPresentRRCState
	protocolIEs.Value.RRCState = new(ngapType.RRCState)
	*protocolIEs.Value.RRCState = RRCState()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationResponseIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.UEContextModificationResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextReleaseComplete () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextRelease
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentUEContextReleaseComplete
	msg.Value.UEContextReleaseComplete = new(ngapType.UEContextReleaseComplete)
	msgProtocolIEs := &msg.Value.UEContextReleaseComplete.ProtocolIEs

	protocolIEs := ngapType.UEContextReleaseCompleteIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCompleteIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCompleteIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// UserLocationInformation
	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDUserLocationInformation
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCompleteIEsPresentUserLocationInformation
	protocolIEs.Value.UserLocationInformation = new(ngapType.UserLocationInformation)
	*protocolIEs.Value.UserLocationInformation = UserLocationInformation()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// InfoOnRecommendedCellsAndRANNodesForPaging
	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCompleteIEsPresentInfoOnRecommendedCellsAndRANNodesForPaging
	protocolIEs.Value.InfoOnRecommendedCellsAndRANNodesForPaging = new(ngapType.InfoOnRecommendedCellsAndRANNodesForPaging)
	*protocolIEs.Value.InfoOnRecommendedCellsAndRANNodesForPaging = InfoOnRecommendedCellsAndRANNodesForPaging()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceListCxtRelCpl
	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceListCxtRelCpl
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCompleteIEsPresentPDUSessionResourceListCxtRelCpl
	protocolIEs.Value.PDUSessionResourceListCxtRelCpl = new(ngapType.PDUSessionResourceListCxtRelCpl)
	*protocolIEs.Value.PDUSessionResourceListCxtRelCpl = PDUSessionResourceListCxtRelCpl()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.UEContextReleaseCompleteIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextReleaseCompleteIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UERadioCapabilityCheckResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUERadioCapabilityCheck
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentUERadioCapabilityCheckResponse
	msg.Value.UERadioCapabilityCheckResponse = new(ngapType.UERadioCapabilityCheckResponse)
	msgProtocolIEs := &msg.Value.UERadioCapabilityCheckResponse.ProtocolIEs

	protocolIEs := ngapType.UERadioCapabilityCheckResponseIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckResponseIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckResponseIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// IMSVoiceSupportIndicator
	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDIMSVoiceSupportIndicator
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckResponseIEsPresentIMSVoiceSupportIndicator
	protocolIEs.Value.IMSVoiceSupportIndicator = new(ngapType.IMSVoiceSupportIndicator)
	*protocolIEs.Value.IMSVoiceSupportIndicator = IMSVoiceSupportIndicator()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.UERadioCapabilityCheckResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UERadioCapabilityCheckResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func WriteReplaceWarningResponse () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentSuccessfulOutcome
	pdu.SuccessfulOutcome = new(ngapType.SuccessfulOutcome)
	msg := pdu.SuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeWriteReplaceWarning
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.SuccessfulOutcomePresentWriteReplaceWarningResponse
	msg.Value.WriteReplaceWarningResponse = new(ngapType.WriteReplaceWarningResponse)
	msgProtocolIEs := &msg.Value.WriteReplaceWarningResponse.ProtocolIEs

	protocolIEs := ngapType.WriteReplaceWarningResponseIEs{}

	// MessageIdentifier
	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDMessageIdentifier
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningResponseIEsPresentMessageIdentifier
	protocolIEs.Value.MessageIdentifier = new(ngapType.MessageIdentifier)
	*protocolIEs.Value.MessageIdentifier = MessageIdentifier()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// SerialNumber
	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDSerialNumber
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningResponseIEsPresentSerialNumber
	protocolIEs.Value.SerialNumber = new(ngapType.SerialNumber)
	*protocolIEs.Value.SerialNumber = SerialNumber()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// BroadcastCompletedAreaList
	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDBroadcastCompletedAreaList
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningResponseIEsPresentBroadcastCompletedAreaList
	protocolIEs.Value.BroadcastCompletedAreaList = new(ngapType.BroadcastCompletedAreaList)
	*protocolIEs.Value.BroadcastCompletedAreaList = BroadcastCompletedAreaList()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.WriteReplaceWarningResponseIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.WriteReplaceWarningResponseIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func AMFConfigurationUpdateFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeAMFConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentNotify
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentAMFConfigurationUpdateFailure
	msg.Value.AMFConfigurationUpdateFailure = new(ngapType.AMFConfigurationUpdateFailure)
	msgProtocolIEs := &msg.Value.AMFConfigurationUpdateFailure.ProtocolIEs

	protocolIEs := ngapType.AMFConfigurationUpdateFailureIEs{}

	// Cause - OK
	protocolIEs = ngapType.AMFConfigurationUpdateFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TimeToWait
	protocolIEs = ngapType.AMFConfigurationUpdateFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTimeToWait
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateFailureIEsPresentTimeToWait
	protocolIEs.Value.TimeToWait = new(ngapType.TimeToWait)
	*protocolIEs.Value.TimeToWait = TimeToWait()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.AMFConfigurationUpdateFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.AMFConfigurationUpdateFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverPreparationFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverPreparation
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentHandoverPreparationFailure
	msg.Value.HandoverPreparationFailure = new(ngapType.HandoverPreparationFailure)
	msgProtocolIEs := &msg.Value.HandoverPreparationFailure.ProtocolIEs

	protocolIEs := ngapType.HandoverPreparationFailureIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverPreparationFailureIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverPreparationFailureIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverPreparationFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.HandoverPreparationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverPreparationFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func HandoverFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeHandoverResourceAllocation
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentHandoverFailure
	msg.Value.HandoverFailure = new(ngapType.HandoverFailure)
	msgProtocolIEs := &msg.Value.HandoverFailure.ProtocolIEs

	protocolIEs := ngapType.HandoverFailureIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.HandoverFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverFailureIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.HandoverFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.HandoverFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.HandoverFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func InitialContextSetupFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeInitialContextSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentInitialContextSetupFailure
	msg.Value.InitialContextSetupFailure = new(ngapType.InitialContextSetupFailure)
	msgProtocolIEs := &msg.Value.InitialContextSetupFailure.ProtocolIEs

	protocolIEs := ngapType.InitialContextSetupFailureIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupFailureIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupFailureIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceFailedToSetupListCxtFail
	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupFailureIEsPresentPDUSessionResourceFailedToSetupListCxtFail
	protocolIEs.Value.PDUSessionResourceFailedToSetupListCxtFail = new(ngapType.PDUSessionResourceFailedToSetupListCxtFail)
	*protocolIEs.Value.PDUSessionResourceFailedToSetupListCxtFail = PDUSessionResourceFailedToSetupListCxtFail()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.InitialContextSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.InitialContextSetupFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func NGSetupFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeNGSetup
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentNGSetupFailure
	msg.Value.NGSetupFailure = new(ngapType.NGSetupFailure)
	msgProtocolIEs := &msg.Value.NGSetupFailure.ProtocolIEs

	protocolIEs := ngapType.NGSetupFailureIEs{}

	// Cause - OK
	protocolIEs = ngapType.NGSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TimeToWait
	protocolIEs = ngapType.NGSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTimeToWait
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupFailureIEsPresentTimeToWait
	protocolIEs.Value.TimeToWait = new(ngapType.TimeToWait)
	*protocolIEs.Value.TimeToWait = TimeToWait()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.NGSetupFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.NGSetupFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func PathSwitchRequestFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodePathSwitchRequest
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentPathSwitchRequestFailure
	msg.Value.PathSwitchRequestFailure = new(ngapType.PathSwitchRequestFailure)
	msgProtocolIEs := &msg.Value.PathSwitchRequestFailure.ProtocolIEs

	protocolIEs := ngapType.PathSwitchRequestFailureIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// PDUSessionResourceReleasedListPSFail
	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDPDUSessionResourceReleasedListPSFail
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentPDUSessionResourceReleasedListPSFail
	protocolIEs.Value.PDUSessionResourceReleasedListPSFail = new(ngapType.PDUSessionResourceReleasedListPSFail)
	*protocolIEs.Value.PDUSessionResourceReleasedListPSFail = PDUSessionResourceReleasedListPSFail()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.PathSwitchRequestFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.PathSwitchRequestFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func RANConfigurationUpdateFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeRANConfigurationUpdate
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentRANConfigurationUpdateFailure
	msg.Value.RANConfigurationUpdateFailure = new(ngapType.RANConfigurationUpdateFailure)
	msgProtocolIEs := &msg.Value.RANConfigurationUpdateFailure.ProtocolIEs

	protocolIEs := ngapType.RANConfigurationUpdateFailureIEs{}

	// Cause - OK
	protocolIEs = ngapType.RANConfigurationUpdateFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// TimeToWait
	protocolIEs = ngapType.RANConfigurationUpdateFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDTimeToWait
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateFailureIEsPresentTimeToWait
	protocolIEs.Value.TimeToWait = new(ngapType.TimeToWait)
	*protocolIEs.Value.TimeToWait = TimeToWait()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.RANConfigurationUpdateFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.RANConfigurationUpdateFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}

func UEContextModificationFailure () (pdu ngapType.NGAPPDU) {
	pdu.Present = ngapType.NGAPPDUPresentUnsuccessfulOutcome
	pdu.UnsuccessfulOutcome = new(ngapType.UnsuccessfulOutcome)
	msg := pdu.UnsuccessfulOutcome
	msg.ProcedureCode.Value = ngapType.ProcedureCodeUEContextModification
	msg.Criticality.Value = ngapType.CriticalityPresentIgnore
	msg.Value.Present = ngapType.UnsuccessfulOutcomePresentUEContextModificationFailure
	msg.Value.UEContextModificationFailure = new(ngapType.UEContextModificationFailure)
	msgProtocolIEs := &msg.Value.UEContextModificationFailure.ProtocolIEs

	protocolIEs := ngapType.UEContextModificationFailureIEs{}

	// AMFUENGAPID - OK
	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDAMFUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationFailureIEsPresentAMFUENGAPID
	protocolIEs.Value.AMFUENGAPID = new(ngapType.AMFUENGAPID)
	*protocolIEs.Value.AMFUENGAPID = AMFUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// RANUENGAPID - OK
	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDRANUENGAPID
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationFailureIEsPresentRANUENGAPID
	protocolIEs.Value.RANUENGAPID = new(ngapType.RANUENGAPID)
	*protocolIEs.Value.RANUENGAPID = RANUENGAPID()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// Cause - OK
	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCause
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationFailureIEsPresentCause
	protocolIEs.Value.Cause = new(ngapType.Cause)
	*protocolIEs.Value.Cause = Cause()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	// CriticalityDiagnostics
	protocolIEs = ngapType.UEContextModificationFailureIEs{}
	protocolIEs.Id.Value = ngapType.ProtocolIEIDCriticalityDiagnostics
	protocolIEs.Criticality.Value = ngapType.CriticalityPresentIgnore
	protocolIEs.Value.Present = ngapType.UEContextModificationFailureIEsPresentCriticalityDiagnostics
	protocolIEs.Value.CriticalityDiagnostics = new(ngapType.CriticalityDiagnostics)
	*protocolIEs.Value.CriticalityDiagnostics = CriticalityDiagnostics()
	msgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)

	return
}