// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	TypeOfErrorPresentNotUnderstood Enumerated = 0
	TypeOfErrorPresentMissing       Enumerated = 1
)

type TypeOfError struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
