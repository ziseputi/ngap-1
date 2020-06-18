package ngapType

const (
	CriticalityReject Enumerated = 0
	CriticalityIgnore Enumerated = 1
	CriticalityNotify Enumerated = 2
)

type Criticality struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	PresenceOptional    Enumerated = 0
	PresenceConditional Enumerated = 1
	PresenceMandatory   Enumerated = 2
)

type Presence struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	PrivateIEIDChoiceLocal  int = 0
	PrivateIEIDChoiceGlobal int = 1
)

type PrivateIEID struct {
	Choice int
	Local  *Integer `vht:"valueMin:0,valueMax:65535"`
	Global *ObjectIdentifier
}

type ProcedureCode struct {
	Value Integer `vht:"valueMin:0,valueMax:255"`
}

type ProtocolExtensionID struct {
	Value Integer `vht:"valueMin:0,valueMax:65535"`
}

type ProtocolIEID struct {
	Value Integer `vht:"valueMin:0,valueMax:65535"`
}

const (
	TriggeringMessageInitiatingMessage   Enumerated = 0
	TriggeringMessageSuccessfulOutcome   Enumerated = 1
	TriggeringMessageUnsuccessfulOutcome Enumerated = 2
)

type TriggeringMessage struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}
