// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	NotificationControlPresentNotificationRequested Enumerated = 0
)

type NotificationControl struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
