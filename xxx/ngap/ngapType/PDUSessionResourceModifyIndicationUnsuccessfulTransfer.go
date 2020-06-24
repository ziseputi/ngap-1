// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceModifyIndicationUnsuccessfulTransfer struct {
	Cause        Cause                                                                                   `vht:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs `vht:"optional"`
}
