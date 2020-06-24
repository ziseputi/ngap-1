// Created By HaoDHH-245789 VHT2020
package ngapType

type GBRQosInformation struct {
	MaximumFlowBitRateDL    BitRate
	MaximumFlowBitRateUL    BitRate
	GuaranteedFlowBitRateDL BitRate
	GuaranteedFlowBitRateUL BitRate
	NotificationControl     *NotificationControl                               `vht:"optional"`
	MaximumPacketLossRateDL *PacketLossRate                                    `vht:"optional"`
	MaximumPacketLossRateUL *PacketLossRate                                    `vht:"optional"`
	IEExtensions            *ProtocolExtensionContainerGBRQosInformationExtIEs `vht:"optional"`
}
