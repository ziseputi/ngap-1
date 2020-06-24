// Created By HaoDHH-245789 VHT2020
package ngapType

type COUNTValueForPDCPSN18 struct {
	PDCPSN18     int64                                                  `vht:"valueLB:0,valueUB:262143"`
	HFNPDCPSN18  int64                                                  `vht:"valueLB:0,valueUB:16383"`
	IEExtensions *ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs `vht:"optional"`
}
