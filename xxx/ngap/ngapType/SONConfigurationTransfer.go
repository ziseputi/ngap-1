// Created By HaoDHH-245789 VHT2020
package ngapType

type SONConfigurationTransfer struct {
	TargetRANNodeID        TargetRANNodeID                                           `vht:"valueExt"`
	SourceRANNodeID        SourceRANNodeID                                           `vht:"valueExt"`
	SONInformation         SONInformation                                            `vht:"valueLB:0,valueUB:2"`
	XnTNLConfigurationInfo XnTNLConfigurationInfo                                    `vht:"valueExt"`
	IEExtensions           *ProtocolExtensionContainerSONConfigurationTransferExtIEs `vht:"optional"`
}
