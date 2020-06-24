// Created By HaoDHH-245789 VHT2020
package ngapType

type UserLocationInformationNR struct {
	NRCGI        NRCGI                                                      `vht:"valueExt"`
	TAI          TAI                                                        `vht:"valueExt"`
	TimeStamp    *TimeStamp                                                 `vht:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationNRExtIEs `vht:"optional"`
}
