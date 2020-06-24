// Created By HaoDHH-245789 VHT2020
package ngapType

type AMFTNLAssociationToUpdateItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation                                    `vht:"valueLB:0,valueUB:1"`
	TNLAssociationUsage      *TNLAssociationUsage                                           `vht:"optional"`
	TNLAddressWeightFactor   *TNLAddressWeightFactor                                        `vht:"optional"`
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs `vht:"optional"`
}
