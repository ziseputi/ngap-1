// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	SourceOfUEActivityBehaviourInformationPresentSubscriptionInformation Enumerated = 0
	SourceOfUEActivityBehaviourInformationPresentStatistics              Enumerated = 1
)

type SourceOfUEActivityBehaviourInformation struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
