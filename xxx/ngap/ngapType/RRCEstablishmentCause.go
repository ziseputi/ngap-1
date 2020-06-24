// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	RRCEstablishmentCausePresentEmergency          Enumerated = 0
	RRCEstablishmentCausePresentHighPriorityAccess Enumerated = 1
	RRCEstablishmentCausePresentMtAccess           Enumerated = 2
	RRCEstablishmentCausePresentMoSignalling       Enumerated = 3
	RRCEstablishmentCausePresentMoData             Enumerated = 4
	RRCEstablishmentCausePresentMoVoiceCall        Enumerated = 5
	RRCEstablishmentCausePresentMoVideoCall        Enumerated = 6
	RRCEstablishmentCausePresentMoSMS              Enumerated = 7
	RRCEstablishmentCausePresentMpsPriorityAccess  Enumerated = 8
	RRCEstablishmentCausePresentMcsPriorityAccess  Enumerated = 9
)

type RRCEstablishmentCause struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:9"`
}
