// Created By HaoDHH-245789 VHT2020
package ngapType

type TNLAssociationItem struct {
	TNLAssociationAddress CPTransportLayerInformation                         `vht:"valueLB:0,valueUB:1"`
	Cause                 Cause                                               `vht:"valueLB:0,valueUB:5"`
	IEExtensions          *ProtocolExtensionContainerTNLAssociationItemExtIEs `vht:"optional"`
}
