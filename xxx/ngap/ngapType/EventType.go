// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	EventTypePresentDirect                          Enumerated = 0
	EventTypePresentChangeOfServeCell               Enumerated = 1
	EventTypePresentUePresenceInAreaOfInterest      Enumerated = 2
	EventTypePresentStopChangeOfServeCell           Enumerated = 3
	EventTypePresentStopUePresenceInAreaOfInterest  Enumerated = 4
	EventTypePresentCancelLocationReportingForTheUe Enumerated = 5
)

type EventType struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:5"`
}
