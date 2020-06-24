// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	NotificationCausePresentFulfilled    Enumerated = 0
	NotificationCausePresentNotFulfilled Enumerated = 1
)

type NotificationCause struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
