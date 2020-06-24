// Created By HaoDHH-245789 VHT2020
package ngapType

type UserPlaneSecurityInformation struct {
	SecurityResult     SecurityResult                                                `vht:"valueExt"`
	SecurityIndication SecurityIndication                                            `vht:"valueExt"`
	IEExtensions       *ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs `vht:"optional"`
}
