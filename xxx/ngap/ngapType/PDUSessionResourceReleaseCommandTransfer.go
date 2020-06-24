// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceReleaseCommandTransfer struct {
	Cause        Cause                                                                     `vht:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs `vht:"optional"`
}
