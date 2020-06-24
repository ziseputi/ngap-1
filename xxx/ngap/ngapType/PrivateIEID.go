// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PrivateIEIDPresentNothing int = iota /* No components present */
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

type PrivateIEID struct {
	Present int
	Local   *int64 `vht:"valueLB:0,valueUB:65535"`
	Global  *ObjectIdentifier
}
