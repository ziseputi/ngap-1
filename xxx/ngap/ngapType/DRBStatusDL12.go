// Created By HaoDHH-245789 VHT2020
package ngapType

type DRBStatusDL12 struct {
	DLCOUNTValue COUNTValueForPDCPSN12                          `vht:"valueExt"`
	IEExtension  *ProtocolExtensionContainerDRBStatusDL12ExtIEs `vht:"optional"`
}
