// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  Cause                                                                        `vht:"valueLB:0,valueUB:5"`
	CriticalityDiagnostics *CriticalityDiagnostics                                                      `vht:"valueExt,optional"`
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs `vht:"optional"`
}
