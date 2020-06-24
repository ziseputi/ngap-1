// Created By HaoDHH-245789 VHT2020
package ngapType

type TNLInformationItem struct {
	QosFlowPerTNLInformation QosFlowPerTNLInformation                            `vht:"valueExt"`
	IEExtensions             *ProtocolExtensionContainerTNLInformationItemExtIEs `vht:"optional"`
}
