// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DirectForwardingPathAvailabilityPresentDirectPathAvailable Enumerated = 0
)

type DirectForwardingPathAvailability struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
