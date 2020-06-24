// Created By HaoDHH-245789 VHT2020
package ngapType

type SecurityContext struct {
	NextHopChainingCount NextHopChainingCount
	NextHopNH            SecurityKey
	IEExtensions         *ProtocolExtensionContainerSecurityContextExtIEs `vht:"optional"`
}
