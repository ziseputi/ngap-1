// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	IMSVoiceSupportIndicatorPresentSupported    Enumerated = 0
	IMSVoiceSupportIndicatorPresentNotSupported Enumerated = 1
)

type IMSVoiceSupportIndicator struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
