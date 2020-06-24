// Created By HaoDHH-245789 VHT2020
package ngapType

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         UPTransportLayerInformation  `vht:"valueLB:0,valueUB:1"`
	DLForwardingUPTNLInformation  *UPTransportLayerInformation `vht:"valueLB:0,valueUB:1,optional"`
	SecurityResult                *SecurityResult              `vht:"valueExt,optional"`
	QosFlowSetupResponseList      QosFlowSetupResponseListHOReqAck
	QosFlowFailedToSetupList      *QosFlowList                                                        `vht:"optional"`
	DataForwardingResponseDRBList *DataForwardingResponseDRBList                                      `vht:"optional"`
	IEExtensions                  *ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs `vht:"optional"`
}
