// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceModifyUnsuccessfulTransfer struct {
	Cause                  Cause                                                                         `vht:"valueLB:0,valueUB:5"`
	CriticalityDiagnostics *CriticalityDiagnostics                                                       `vht:"valueExt,optional"`
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs `vht:"optional"`
}