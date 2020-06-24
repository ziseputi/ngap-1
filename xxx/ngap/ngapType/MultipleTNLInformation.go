// Created By HaoDHH-245789 VHT2020
package ngapType

type MultipleTNLInformation struct {
	TNLInformationList TNLInformationList
	IEExtensions       *ProtocolExtensionContainerMultipleTNLInformationExtIEs `vht:"optional"`
}
