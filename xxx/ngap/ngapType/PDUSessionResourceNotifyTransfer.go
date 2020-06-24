// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList                                                `vht:"optional"`
	QosFlowReleasedList *QosFlowList                                                      `vht:"optional"`
	IEExtensions        *ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs `vht:"optional"`
}
