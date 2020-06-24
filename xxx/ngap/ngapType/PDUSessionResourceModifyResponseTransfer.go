// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation              *UPTransportLayerInformation                                              `vht:"valueLB:0,valueUB:1,optional"`
	ULNGUUPTNLInformation              *UPTransportLayerInformation                                              `vht:"valueLB:0,valueUB:1,optional"`
	QosFlowAddOrModifyResponseList     *QosFlowAddOrModifyResponseList                                           `vht:"optional"`
	AdditionalQosFlowPerTNLInformation *QosFlowPerTNLInformation                                                 `vht:"valueExt,optional"`
	QosFlowFailedToAddOrModifyList     *QosFlowList                                                              `vht:"optional"`
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs `vht:"optional"`
}
