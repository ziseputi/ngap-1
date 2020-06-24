// Created By HaoDHH-245789 VHT2020
package ngapType

type DRBStatusUL12 struct {
	ULCOUNTValue              COUNTValueForPDCPSN12                          `vht:"valueExt"`
	ReceiveStatusOfULPDCPSDUs *BitString                                `vht:"sizeLB:1,sizeUB:2048,optional"`
	IEExtension               *ProtocolExtensionContainerDRBStatusUL12ExtIEs `vht:"optional"`
}
