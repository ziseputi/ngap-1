// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PDUSessionTypePresentIpv4         Enumerated = 0
	PDUSessionTypePresentIpv6         Enumerated = 1
	PDUSessionTypePresentIpv4v6       Enumerated = 2
	PDUSessionTypePresentEthernet     Enumerated = 3
	PDUSessionTypePresentUnstructured Enumerated = 4
)

type PDUSessionType struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:4"`
}
