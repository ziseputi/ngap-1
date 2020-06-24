// Created By HaoDHH-245789 VHT2020
package ngapType

type DRBStatusDL18 struct {
	DLCOUNTValue COUNTValueForPDCPSN18                          `vht:"valueExt"`
	IEExtension  *ProtocolExtensionContainerDRBStatusDL18ExtIEs `vht:"optional"`
}
