// Created By HaoDHH-245789 VHT2020
package ngapType

type UnavailableGUAMIItem struct {
	GUAMI                        GUAMI                                                 `vht:"valueExt"`
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval                         `vht:"optional"`
	BackupAMFName                *AMFName                                              `vht:"sizeExt,sizeLB:1,sizeUB:150,optional"`
	IEExtensions                 *ProtocolExtensionContainerUnavailableGUAMIItemExtIEs `vht:"optional"`
}
