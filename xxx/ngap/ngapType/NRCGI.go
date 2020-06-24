// Created By HaoDHH-245789 VHT2020
package ngapType

type NRCGI struct {
	PLMNIdentity   PLMNIdentity
	NRCellIdentity NRCellIdentity
	IEExtensions   *ProtocolExtensionContainerNRCGIExtIEs `vht:"optional"`
}
