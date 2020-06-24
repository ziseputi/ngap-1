// Created By HaoDHH-245789 VHT2020
package ngapType

type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation `vht:"valueLB:0,valueUB:1"`
	TNLAssociationUsage      *TNLAssociationUsage        `vht:"optional"`
	TNLAddressWeightFactor   TNLAddressWeightFactor
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs `vht:"optional"`
}
