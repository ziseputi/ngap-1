// Created By HaoDHH-245789 VHT2020
package ngapType

type DRBStatusUL18 struct {
	ULCOUNTValue              COUNTValueForPDCPSN18                          `vht:"valueExt"`
	ReceiveStatusOfULPDCPSDUs *BitString                                `vht:"sizeLB:1,sizeUB:131072,optional"`
	IEExtension               *ProtocolExtensionContainerDRBStatusUL18ExtIEs `vht:"optional"`
}
