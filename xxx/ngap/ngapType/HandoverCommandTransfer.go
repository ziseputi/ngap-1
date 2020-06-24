// Created By HaoDHH-245789 VHT2020
package ngapType

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  *UPTransportLayerInformation                             `vht:"valueLB:0,valueUB:1,optional"`
	QosFlowToBeForwardedList      *QosFlowToBeForwardedList                                `vht:"optional"`
	DataForwardingResponseDRBList *DataForwardingResponseDRBList                           `vht:"optional"`
	IEExtensions                  *ProtocolExtensionContainerHandoverCommandTransferExtIEs `vht:"optional"`
}
