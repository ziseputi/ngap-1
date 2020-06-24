// Created By HaoDHH-245789 VHT2020
package ngapType

type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication
	MaximumIntegrityProtectedDataRate   *MaximumIntegrityProtectedDataRate                  `vht:"optional"`
	IEExtensions                        *ProtocolExtensionContainerSecurityIndicationExtIEs `vht:"optional"`
}
