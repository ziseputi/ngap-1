// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	TimeToWaitPresentV1s  Enumerated = 0
	TimeToWaitPresentV2s  Enumerated = 1
	TimeToWaitPresentV5s  Enumerated = 2
	TimeToWaitPresentV10s Enumerated = 3
	TimeToWaitPresentV20s Enumerated = 4
	TimeToWaitPresentV60s Enumerated = 5
)

type TimeToWait struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:5"`
}
