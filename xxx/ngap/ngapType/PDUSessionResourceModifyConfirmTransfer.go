// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList  QosFlowModifyConfirmList
	TNLMappingList            *TNLMappingList                                                          `vht:"optional"`
	QosFlowFailedToModifyList *QosFlowList                                                             `vht:"optional"`
	IEExtensions              *ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs `vht:"optional"`
}
