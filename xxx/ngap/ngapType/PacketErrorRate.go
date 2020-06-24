// Created By HaoDHH-245789 VHT2020
package ngapType

type PacketErrorRate struct {
	PERScalar    int64                                            `vht:"valueExt,valueLB:0,valueUB:9"`
	PERExponent  int64                                            `vht:"valueExt,valueLB:0,valueUB:9"`
	IEExtensions *ProtocolExtensionContainerPacketErrorRateExtIEs `vht:"optional"`
}
