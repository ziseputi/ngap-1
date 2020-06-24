// Created By HaoDHH-245789 VHT2020
package ngapType

type EUTRACGI struct {
	PLMNIdentity      PLMNIdentity
	EUTRACellIdentity EUTRACellIdentity
	IEExtensions      *ProtocolExtensionContainerEUTRACGIExtIEs `vht:"optional"`
}
