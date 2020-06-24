// Created By HaoDHH-245789 VHT2020
package ngapType

type AMFTNLAssociationSetupItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation                                 `vht:"valueLB:0,valueUB:1"`
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs `vht:"optional"`
}
