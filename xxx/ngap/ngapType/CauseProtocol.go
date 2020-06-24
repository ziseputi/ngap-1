// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CauseProtocolPresentTransferSyntaxError                          Enumerated = 0
	CauseProtocolPresentAbstractSyntaxErrorReject                    Enumerated = 1
	CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify           Enumerated = 2
	CauseProtocolPresentMessageNotCompatibleWithReceiverState        Enumerated = 3
	CauseProtocolPresentSemanticError                                Enumerated = 4
	CauseProtocolPresentAbstractSyntaxErrorFalselyConstructedMessage Enumerated = 5
	CauseProtocolPresentUnspecified                                  Enumerated = 6
)

type CauseProtocol struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:6"`
}
