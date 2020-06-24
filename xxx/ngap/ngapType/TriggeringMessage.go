// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	TriggeringMessagePresentInitiatingMessage    Enumerated = 0
	TriggeringMessagePresentSuccessfulOutcome    Enumerated = 1
	TriggeringMessagePresentUnsuccessfullOutcome Enumerated = 2
)

type TriggeringMessage struct {
	Value Enumerated `vht:"valueLB:0,valueUB:2"`
}
