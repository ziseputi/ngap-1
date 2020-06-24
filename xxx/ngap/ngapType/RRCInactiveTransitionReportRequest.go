// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	RRCInactiveTransitionReportRequestPresentSubsequentStateTransitionReport Enumerated = 0
	RRCInactiveTransitionReportRequestPresentSingleRrcConnectedStateReport   Enumerated = 1
	RRCInactiveTransitionReportRequestPresentCancelReport                    Enumerated = 2
)

type RRCInactiveTransitionReportRequest struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:2"`
}
