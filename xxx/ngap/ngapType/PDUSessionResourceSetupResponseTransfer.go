// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceSetupResponseTransfer struct {
	QosFlowPerTNLInformation           QosFlowPerTNLInformation                                                 `vht:"valueExt"`
	AdditionalQosFlowPerTNLInformation *QosFlowPerTNLInformation                                                `vht:"valueExt,optional"`
	SecurityResult                     *SecurityResult                                                          `vht:"valueExt,optional"`
	QosFlowFailedToSetupList           *QosFlowList                                                             `vht:"optional"`
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs `vht:"optional"`
}
