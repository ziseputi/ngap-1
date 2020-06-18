package ngapType

// BIT STRING

// BitString is for an ASN.1 BIT STRING type, BitLength means the effective bits.
type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength uint64 // length in bits.
}

// OCTET STRING

// OctetString is for an ASN.1 OCTET STRING type
type OctetString []byte

// OBJECT IDENTIFIER

// ObjectIdentifier is for an ASN.1 OBJECT IDENTIFIER type
type ObjectIdentifier []byte

// ENUMERATED

// An Enumerated is represented as a plain uint64.
type Enumerated uint64

// INTEGER

// An Integer is represented as a plain int64.
type Integer int64

// PrintableString

// A PrintableString is represented as a plain string.
type PrintableString string

// All types, universal class

// | Name                                           | Value encodings                | Tag number
// |                                                |                                | Decimal | Hexadecimal
// | End-of-Content (EOC)                           | Primitive                      | 0       | 0
// | BOOLEAN                                        | Primitive                      | 1       | 1
// | INTEGER                                        | Primitive                      | 2       | 2
// | BIT STRING                                     | Both                           | 3       | 3
// | OCTET STRING                                   | Both                           | 4       | 4
// | NULL                                           | Primitive                      | 5       | 5
// | OBJECT IDENTIFIER                              | Primitive                      | 6       | 6
// | Object Descriptor                              | Both                           | 7       | 7
// | EXTERNAL                                       | Constructed                    | 8       | 8
// | REAL (float)                                   | Primitive                      | 9       | 9
// | ENUMERATED                                     | Primitive                      | 10      | A
// | EMBEDDED PDV                                   | Constructed                    | 11      | B
// | UTF8String                                     | Both                           | 12      | C
// | RELATIVE-OID                                   | Primitive                      | 13      | D
// | TIME                                           | Primitive                      | 14      | E
// | Reserved                                       |                                | 15      | F
// | SEQUENCE and SEQUENCE OF                       | Constructed                    | 16      | 10
// | SET and SET OF                                 | Constructed                    | 17      | 11
// | NumericString                                  | Both                           | 18      | 12
// | PrintableString                                | Both                           | 19      | 13
// | T61String                                      | Both                           | 20      | 14
// | VideotexString                                 | Both                           | 21      | 15
// | IA5String                                      | Both                           | 22      | 16
// | UTCTime                                        | Both                           | 23      | 17
// | GeneralizedTime                                | Both                           | 24      | 18
// | GraphicString                                  | Both                           | 25      | 19
// | VisibleString                                  | Both                           | 26      | 1A
// | GeneralString                                  | Both                           | 27      | 1B
// | UniversalString                                | Both                           | 28      | 1C
// | CHARACTER STRING                               | Both                           | 29      | 1D
// | BMPString                                      | Both                           | 30      | 1E
// | DATE                                           | Primitive                      | 31      | 1F
// | TIME-OF-DAY                                    | Primitive                      | 32      | 20
// | DATE-TIME                                      | Primitive                      | 33      | 21
// | DURATION                                       | Primitive                      | 34      | 22
// | OID-IRI                                        | Primitive                      | 35      | 23
// | RELATIVE-OID-IRI                               | Primitive                      | 36      | 24
