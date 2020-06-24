// Created By HaoDHH-245789 VHT2020
package ngapType

type InitiatingMessage struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality
	Value         InitiatingMessageValue `vht:"openType,referenceFieldName:ProcedureCode"`
}

const (
	InitiatingMessagePresentNothing int = iota /* No components present */
	InitiatingMessagePresentAMFConfigurationUpdate
	InitiatingMessagePresentHandoverCancel
	InitiatingMessagePresentHandoverRequired
	InitiatingMessagePresentHandoverRequest
	InitiatingMessagePresentInitialContextSetupRequest
	InitiatingMessagePresentNGReset
	InitiatingMessagePresentNGSetupRequest
	InitiatingMessagePresentPathSwitchRequest
	InitiatingMessagePresentPDUSessionResourceModifyRequest
	InitiatingMessagePresentPDUSessionResourceModifyIndication
	InitiatingMessagePresentPDUSessionResourceReleaseCommand
	InitiatingMessagePresentPDUSessionResourceSetupRequest
	InitiatingMessagePresentPWSCancelRequest
	InitiatingMessagePresentRANConfigurationUpdate
	InitiatingMessagePresentUEContextModificationRequest
	InitiatingMessagePresentUEContextReleaseCommand
	InitiatingMessagePresentUERadioCapabilityCheckRequest
	InitiatingMessagePresentWriteReplaceWarningRequest
	InitiatingMessagePresentAMFStatusIndication
	InitiatingMessagePresentCellTrafficTrace
	InitiatingMessagePresentDeactivateTrace
	InitiatingMessagePresentDownlinkNASTransport
	InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport
	InitiatingMessagePresentDownlinkRANConfigurationTransfer
	InitiatingMessagePresentDownlinkRANStatusTransfer
	InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport
	InitiatingMessagePresentErrorIndication
	InitiatingMessagePresentHandoverNotify
	InitiatingMessagePresentInitialUEMessage
	InitiatingMessagePresentLocationReport
	InitiatingMessagePresentLocationReportingControl
	InitiatingMessagePresentLocationReportingFailureIndication
	InitiatingMessagePresentNASNonDeliveryIndication
	InitiatingMessagePresentOverloadStart
	InitiatingMessagePresentOverloadStop
	InitiatingMessagePresentPaging
	InitiatingMessagePresentPDUSessionResourceNotify
	InitiatingMessagePresentPrivateMessage
	InitiatingMessagePresentPWSFailureIndication
	InitiatingMessagePresentPWSRestartIndication
	InitiatingMessagePresentRerouteNASRequest
	InitiatingMessagePresentRRCInactiveTransitionReport
	InitiatingMessagePresentTraceFailureIndication
	InitiatingMessagePresentTraceStart
	InitiatingMessagePresentUEContextReleaseRequest
	InitiatingMessagePresentUERadioCapabilityInfoIndication
	InitiatingMessagePresentUETNLABindingReleaseRequest
	InitiatingMessagePresentUplinkNASTransport
	InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport
	InitiatingMessagePresentUplinkRANConfigurationTransfer
	InitiatingMessagePresentUplinkRANStatusTransfer
	InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport
)

type InitiatingMessageValue struct {
	Present                               int
	AMFConfigurationUpdate                *AMFConfigurationUpdate                `vht:"valueExt,referenceFieldValue:0"`
	HandoverCancel                        *HandoverCancel                        `vht:"valueExt,referenceFieldValue:10"`
	HandoverRequired                      *HandoverRequired                      `vht:"valueExt,referenceFieldValue:12"`
	HandoverRequest                       *HandoverRequest                       `vht:"valueExt,referenceFieldValue:13"`
	InitialContextSetupRequest            *InitialContextSetupRequest            `vht:"valueExt,referenceFieldValue:14"`
	NGReset                               *NGReset                               `vht:"valueExt,referenceFieldValue:20"`
	NGSetupRequest                        *NGSetupRequest                        `vht:"valueExt,referenceFieldValue:21"`
	PathSwitchRequest                     *PathSwitchRequest                     `vht:"valueExt,referenceFieldValue:25"`
	PDUSessionResourceModifyRequest       *PDUSessionResourceModifyRequest       `vht:"valueExt,referenceFieldValue:26"`
	PDUSessionResourceModifyIndication    *PDUSessionResourceModifyIndication    `vht:"valueExt,referenceFieldValue:27"`
	PDUSessionResourceReleaseCommand      *PDUSessionResourceReleaseCommand      `vht:"valueExt,referenceFieldValue:28"`
	PDUSessionResourceSetupRequest        *PDUSessionResourceSetupRequest        `vht:"valueExt,referenceFieldValue:29"`
	PWSCancelRequest                      *PWSCancelRequest                      `vht:"valueExt,referenceFieldValue:32"`
	RANConfigurationUpdate                *RANConfigurationUpdate                `vht:"valueExt,referenceFieldValue:35"`
	UEContextModificationRequest          *UEContextModificationRequest          `vht:"valueExt,referenceFieldValue:40"`
	UEContextReleaseCommand               *UEContextReleaseCommand               `vht:"valueExt,referenceFieldValue:41"`
	UERadioCapabilityCheckRequest         *UERadioCapabilityCheckRequest         `vht:"valueExt,referenceFieldValue:43"`
	WriteReplaceWarningRequest            *WriteReplaceWarningRequest            `vht:"valueExt,referenceFieldValue:51"`
	AMFStatusIndication                   *AMFStatusIndication                   `vht:"valueExt,referenceFieldValue:1"`
	CellTrafficTrace                      *CellTrafficTrace                      `vht:"valueExt,referenceFieldValue:2"`
	DeactivateTrace                       *DeactivateTrace                       `vht:"valueExt,referenceFieldValue:3"`
	DownlinkNASTransport                  *DownlinkNASTransport                  `vht:"valueExt,referenceFieldValue:4"`
	DownlinkNonUEAssociatedNRPPaTransport *DownlinkNonUEAssociatedNRPPaTransport `vht:"valueExt,referenceFieldValue:5"`
	DownlinkRANConfigurationTransfer      *DownlinkRANConfigurationTransfer      `vht:"valueExt,referenceFieldValue:6"`
	DownlinkRANStatusTransfer             *DownlinkRANStatusTransfer             `vht:"valueExt,referenceFieldValue:7"`
	DownlinkUEAssociatedNRPPaTransport    *DownlinkUEAssociatedNRPPaTransport    `vht:"valueExt,referenceFieldValue:8"`
	ErrorIndication                       *ErrorIndication                       `vht:"valueExt,referenceFieldValue:9"`
	HandoverNotify                        *HandoverNotify                        `vht:"valueExt,referenceFieldValue:11"`
	InitialUEMessage                      *InitialUEMessage                      `vht:"valueExt,referenceFieldValue:15"`
	LocationReport                        *LocationReport                        `vht:"valueExt,referenceFieldValue:18"`
	LocationReportingControl              *LocationReportingControl              `vht:"valueExt,referenceFieldValue:16"`
	LocationReportingFailureIndication    *LocationReportingFailureIndication    `vht:"valueExt,referenceFieldValue:17"`
	NASNonDeliveryIndication              *NASNonDeliveryIndication              `vht:"valueExt,referenceFieldValue:19"`
	OverloadStart                         *OverloadStart                         `vht:"valueExt,referenceFieldValue:22"`
	OverloadStop                          *OverloadStop                          `vht:"valueExt,referenceFieldValue:23"`
	Paging                                *Paging                                `vht:"valueExt,referenceFieldValue:24"`
	PDUSessionResourceNotify              *PDUSessionResourceNotify              `vht:"valueExt,referenceFieldValue:30"`
	PrivateMessage                        *PrivateMessage                        `vht:"valueExt,referenceFieldValue:31"`
	PWSFailureIndication                  *PWSFailureIndication                  `vht:"valueExt,referenceFieldValue:33"`
	PWSRestartIndication                  *PWSRestartIndication                  `vht:"valueExt,referenceFieldValue:34"`
	RerouteNASRequest                     *RerouteNASRequest                     `vht:"valueExt,referenceFieldValue:36"`
	RRCInactiveTransitionReport           *RRCInactiveTransitionReport           `vht:"valueExt,referenceFieldValue:37"`
	TraceFailureIndication                *TraceFailureIndication                `vht:"valueExt,referenceFieldValue:38"`
	TraceStart                            *TraceStart                            `vht:"valueExt,referenceFieldValue:39"`
	UEContextReleaseRequest               *UEContextReleaseRequest               `vht:"valueExt,referenceFieldValue:42"`
	UERadioCapabilityInfoIndication       *UERadioCapabilityInfoIndication       `vht:"valueExt,referenceFieldValue:44"`
	UETNLABindingReleaseRequest           *UETNLABindingReleaseRequest           `vht:"valueExt,referenceFieldValue:45"`
	UplinkNASTransport                    *UplinkNASTransport                    `vht:"valueExt,referenceFieldValue:46"`
	UplinkNonUEAssociatedNRPPaTransport   *UplinkNonUEAssociatedNRPPaTransport   `vht:"valueExt,referenceFieldValue:47"`
	UplinkRANConfigurationTransfer        *UplinkRANConfigurationTransfer        `vht:"valueExt,referenceFieldValue:48"`
	UplinkRANStatusTransfer               *UplinkRANStatusTransfer               `vht:"valueExt,referenceFieldValue:49"`
	UplinkUEAssociatedNRPPaTransport      *UplinkUEAssociatedNRPPaTransport      `vht:"valueExt,referenceFieldValue:50"`
}
