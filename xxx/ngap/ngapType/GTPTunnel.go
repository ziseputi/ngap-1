// Created By HaoDHH-245789 VHT2020
package ngapType

type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress
	GTPTEID               GTPTEID
	IEExtensions          *ProtocolExtensionContainerGTPTunnelExtIEs `vht:"optional"`
}
