// Created By HaoDHH-245789 VHT2020
package ngapType

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        UPTransportLayerInformation   `vht:"valueLB:0,valueUB:1"`
	DLNGUTNLInformationReused    *DLNGUTNLInformationReused    `vht:"optional"`
	UserPlaneSecurityInformation *UserPlaneSecurityInformation `vht:"valueExt,optional"`
	QosFlowAcceptedList          QosFlowAcceptedList
	IEExtensions                 *ProtocolExtensionContainerPathSwitchRequestTransferExtIEs `vht:"optional"`
}
