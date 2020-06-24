// Created By HaoDHH-245789 VHT2020
package ngapType

type TargeteNBID struct {
	GlobalENBID    GlobalNgENBID                                `vht:"valueExt"`
	SelectedEPSTAI EPSTAI                                       `vht:"valueExt"`
	IEExtensions   *ProtocolExtensionContainerTargeteNBIDExtIEs `vht:"optional"`
}
