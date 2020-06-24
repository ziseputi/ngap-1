// Created By HaoDHH-245789 VHT2020
package ngapType

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation                                          `vht:"valueLB:0,valueUB:1,optional"`
	SecurityIndication    *SecurityIndication                                                   `vht:"valueExt,optional"`
	IEExtensions          *ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs `vht:"optional"`
}
