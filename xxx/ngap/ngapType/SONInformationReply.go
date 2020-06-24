// Created By HaoDHH-245789 VHT2020
package ngapType

type SONInformationReply struct {
	XnTNLConfigurationInfo *XnTNLConfigurationInfo                              `vht:"valueExt,optional"`
	IEExtensions           *ProtocolExtensionContainerSONInformationReplyExtIEs `vht:"optional"`
}
