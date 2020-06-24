// Created By HaoDHH-245789 VHT2020
package ngapType

type UserLocationInformationEUTRA struct {
	EUTRACGI     EUTRACGI                                                      `vht:"valueExt"`
	TAI          TAI                                                           `vht:"valueExt"`
	TimeStamp    *TimeStamp                                                    `vht:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs `vht:"optional"`
}
