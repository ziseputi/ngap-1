// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CriticalityPresentReject Enumerated = 0
	CriticalityPresentIgnore Enumerated = 1
	CriticalityPresentNotify Enumerated = 2
)

type Criticality struct {
	Value Enumerated `vht:"valueLB:0,valueUB:2"`
}
