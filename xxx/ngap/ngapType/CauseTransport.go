// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CauseTransportPresentTransportResourceUnavailable Enumerated = 0
	CauseTransportPresentUnspecified                  Enumerated = 1
)

type CauseTransport struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
