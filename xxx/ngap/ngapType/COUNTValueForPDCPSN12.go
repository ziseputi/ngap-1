// Created By HaoDHH-245789 VHT2020
package ngapType

type COUNTValueForPDCPSN12 struct {
	PDCPSN12     int64                                                  `vht:"valueLB:0,valueUB:4095"`
	HFNPDCPSN12  int64                                                  `vht:"valueLB:0,valueUB:1048575"`
	IEExtensions *ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs `vht:"optional"`
}
