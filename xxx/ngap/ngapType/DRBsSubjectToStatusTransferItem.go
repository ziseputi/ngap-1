// Created By HaoDHH-245789 VHT2020
package ngapType

type DRBsSubjectToStatusTransferItem struct {
	DRBID       DRBID
	DRBStatusUL DRBStatusUL                                                      `vht:"valueLB:0,valueUB:2"`
	DRBStatusDL DRBStatusDL                                                      `vht:"valueLB:0,valueUB:2"`
	IEExtension *ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs `vht:"optional"`
}
