// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PresencePresentOptional    Enumerated = 0
	PresencePresentConditional Enumerated = 1
	PresencePresentMandatory   Enumerated = 2
)

type Presence struct {
	Value Enumerated `vht:"valueLB:0,valueUB:2"`
}
