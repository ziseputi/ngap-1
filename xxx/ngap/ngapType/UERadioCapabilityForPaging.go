// Created By HaoDHH-245789 VHT2020
package ngapType

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR                             `vht:"optional"`
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA                          `vht:"optional"`
	IEExtensions                      *ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs `vht:"optional"`
}
