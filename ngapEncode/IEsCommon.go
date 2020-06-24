package ngapEncode

import (
	"errors"

	"ngap/ngapType"
)

func AdditionalQosFlowInformation (value ngapType.AdditionalQosFlowInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("AdditionalQosFlowInformation: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func AMFName (value ngapType.AMFName, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 150 {
		return binData, bitEnd, errors.New("AMFName: PrintableString: sizeMin:1,sizeMax:150")
	}
	binData, bitEnd = EncodePrintableString(value.Value,1, 150, binData, bitEnd)
	return binData, bitEnd, nil
}

func AMFPointer (value ngapType.AMFPointer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 6 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("AMFPointer: BitString: sizeMin:6,sizeMax:6")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func AMFRegionID (value ngapType.AMFRegionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("AMFRegionID: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func AMFSetID (value ngapType.AMFSetID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 10 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("AMFSetID: BitString: sizeMin:10,sizeMax:10")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func AMFUENGAPID (value ngapType.AMFUENGAPID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1099511627775 {
		return binData, bitEnd, errors.New("AMFUENGAPID: Integer: valueMin:0,valueMax:1099511627775")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func AveragingWindow (value ngapType.AveragingWindow, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4095 {
		return binData, bitEnd, errors.New("AveragingWindow: Integer: valueMin:0,valueMax:4095")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func BitRate (value ngapType.BitRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4000000000000 {
		return binData, bitEnd, errors.New("BitRate: Integer: valueMin:0,valueMax:4000000000000")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func CancelAllWarningMessages (value ngapType.CancelAllWarningMessages, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("CancelAllWarningMessages: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseMisc (value ngapType.CauseMisc, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 6 {
		return binData, bitEnd, errors.New("CauseMisc: Enumerated: valueMin:0,valueMax:6")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseNas (value ngapType.CauseNas, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4 {
		return binData, bitEnd, errors.New("CauseNas: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseProtocol (value ngapType.CauseProtocol, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 7 {
		return binData, bitEnd, errors.New("CauseProtocol: Enumerated: valueMin:0,valueMax:7")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseRadioNetwork (value ngapType.CauseRadioNetwork, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 48 {
		return binData, bitEnd, errors.New("CauseRadioNetwork: Enumerated: valueMin:0,valueMax:48")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 6, binData, bitEnd)
	return binData, bitEnd, nil
}

func CauseTransport (value ngapType.CauseTransport, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("CauseTransport: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func CellSize (value ngapType.CellSize, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4 {
		return binData, bitEnd, errors.New("CellSize: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func CNTypeRestrictionsForServing (value ngapType.CNTypeRestrictionsForServing, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("CNTypeRestrictionsForServing: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func CommonNetworkInstance(value ngapType.CommonNetworkInstance, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("CommonNetworkInstance: "+ err.Error())
	}
	return binData, bitEnd, err
}

func ConcurrentWarningMessageInd (value ngapType.ConcurrentWarningMessageInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("ConcurrentWarningMessageInd: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ConfidentialityProtectionIndication (value ngapType.ConfidentialityProtectionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("ConfidentialityProtectionIndication: Enumerated: valueMin:0,valueMax:3")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func ConfidentialityProtectionResult (value ngapType.ConfidentialityProtectionResult, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("ConfidentialityProtectionResult: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DataCodingScheme (value ngapType.DataCodingScheme, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("DataCodingScheme: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func DataForwardingAccepted (value ngapType.DataForwardingAccepted, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("DataForwardingAccepted: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DataForwardingNotPossible (value ngapType.DataForwardingNotPossible, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("DataForwardingNotPossible: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DelayCritical (value ngapType.DelayCritical, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("DelayCritical: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DLForwarding (value ngapType.DLForwarding, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("DLForwarding: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DLNGUTNLInformationReused (value ngapType.DLNGUTNLInformationReused, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("DLNGUTNLInformationReused: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DirectForwardingPathAvailability (value ngapType.DirectForwardingPathAvailability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("DirectForwardingPathAvailability: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func DRBID (value ngapType.DRBID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 32 {
		return binData, bitEnd, errors.New("DRBID: Integer: valueMin:1,valueMax:32")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func EmergencyAreaID (value ngapType.EmergencyAreaID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("EmergencyAreaID: OctetString: sizeMin:3,sizeMax:3")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func EmergencyFallbackRequestIndicator (value ngapType.EmergencyFallbackRequestIndicator, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("EmergencyFallbackRequestIndicator: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func EmergencyServiceTargetCN (value ngapType.EmergencyServiceTargetCN, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("EmergencyServiceTargetCN: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ENDCSONConfigurationTransfer(value ngapType.ENDCSONConfigurationTransfer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("ENDCSONConfigurationTransfer: "+ err.Error())
	}
	return binData, bitEnd, err
}

func EPSTAC (value ngapType.EPSTAC, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("EPSTAC: OctetString: sizeMin:2,sizeMax:2")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func ERABID (value ngapType.ERABID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 15 {
		return binData, bitEnd, errors.New("ERABID: Integer: valueMin:0,valueMax:15")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func EUTRACellIdentity (value ngapType.EUTRACellIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 28 || len(value.Value.Bytes) != 4 {
		return binData, bitEnd, errors.New("EUTRACellIdentity: BitString: sizeMin:28,sizeMax:28")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func EUTRAencryptionAlgorithms (value ngapType.EUTRAencryptionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("EUTRAencryptionAlgorithms: BitString: sizeMin:16,sizeMax:16")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func EUTRAintegrityProtectionAlgorithms (value ngapType.EUTRAintegrityProtectionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("EUTRAintegrityProtectionAlgorithms: BitString: sizeMin:16,sizeMax:16")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func EventType (value ngapType.EventType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 6 {
		return binData, bitEnd, errors.New("EventType: Enumerated: valueMin:0,valueMax:6")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedActivityPeriod (value ngapType.ExpectedActivityPeriod, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 30|40|50|60|80|100|120|150|180|181 {
		return binData, bitEnd, errors.New("ExpectedActivityPeriod: Integer: valueMin:1,valueMax:30|40|50|60|80|100|120|150|180|181")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedHOInterval (value ngapType.ExpectedHOInterval, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 7 {
		return binData, bitEnd, errors.New("ExpectedHOInterval: Enumerated: valueMin:0,valueMax:7")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedIdlePeriod (value ngapType.ExpectedIdlePeriod, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 30|40|50|60|80|100|120|150|180|181 {
		return binData, bitEnd, errors.New("ExpectedIdlePeriod: Integer: valueMin:1,valueMax:30|40|50|60|80|100|120|150|180|181")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExpectedUEMobility (value ngapType.ExpectedUEMobility, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("ExpectedUEMobility: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ExtendedRNCID (value ngapType.ExtendedRNCID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 4096 || value.Value > 65535 {
		return binData, bitEnd, errors.New("ExtendedRNCID: Integer: valueMin:4096,valueMax:65535")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func FiveGTMSI (value ngapType.FiveGTMSI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 4 {
		return binData, bitEnd, errors.New("FiveGTMSI: OctetString: sizeMin:4,sizeMax:4")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func FiveQI (value ngapType.FiveQI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("FiveQI: Integer: valueMin:0,valueMax:255")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func GTPTEID (value ngapType.GTPTEID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 4 {
		return binData, bitEnd, errors.New("GTPTEID: OctetString: sizeMin:4,sizeMax:4")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func GUAMIType (value ngapType.GUAMIType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("GUAMIType: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func HandoverFlag (value ngapType.HandoverFlag, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("HandoverFlag: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func HandoverType (value ngapType.HandoverType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4 {
		return binData, bitEnd, errors.New("HandoverType: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func IMSVoiceSupportIndicator (value ngapType.IMSVoiceSupportIndicator, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("IMSVoiceSupportIndicator: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func IndexToRFSP (value ngapType.IndexToRFSP, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 256 {
		return binData, bitEnd, errors.New("IndexToRFSP: Integer: valueMin:1,valueMax:256")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func IntegrityProtectionIndication (value ngapType.IntegrityProtectionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("IntegrityProtectionIndication: Enumerated: valueMin:0,valueMax:3")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func IntegrityProtectionResult (value ngapType.IntegrityProtectionResult, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("IntegrityProtectionResult: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func IntendedNumberOfPagingAttempts (value ngapType.IntendedNumberOfPagingAttempts, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 16 {
		return binData, bitEnd, errors.New("IntendedNumberOfPagingAttempts: Integer: valueMin:1,valueMax:16")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func InterfacesToTrace (value ngapType.InterfacesToTrace, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("InterfacesToTrace: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func LAC (value ngapType.LAC, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("LAC: OctetString: sizeMin:2,sizeMax:2")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func LastVisitedEUTRANCellInformation(value ngapType.LastVisitedEUTRANCellInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedEUTRANCellInformation: "+ err.Error())
	}
	return binData, bitEnd, err
}

func LastVisitedGERANCellInformation(value ngapType.LastVisitedGERANCellInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedGERANCellInformation: "+ err.Error())
	}
	return binData, bitEnd, err
}

func LastVisitedUTRANCellInformation(value ngapType.LastVisitedUTRANCellInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("LastVisitedUTRANCellInformation: "+ err.Error())
	}
	return binData, bitEnd, err
}

func LocationReportingAdditionalInfo (value ngapType.LocationReportingAdditionalInfo, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("LocationReportingAdditionalInfo: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func LocationReportingReferenceID (value ngapType.LocationReportingReferenceID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 64 {
		return binData, bitEnd, errors.New("LocationReportingReferenceID: Integer: valueMin:1,valueMax:64")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func MaskedIMEISV (value ngapType.MaskedIMEISV, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 64 || len(value.Value.Bytes) != 8 {
		return binData, bitEnd, errors.New("MaskedIMEISV: BitString: sizeMin:64,sizeMax:64")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func MaximumDataBurstVolume (value ngapType.MaximumDataBurstVolume, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2000000 {
		return binData, bitEnd, errors.New("MaximumDataBurstVolume: Integer: valueMin:0,valueMax:2000000")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func MessageIdentifier (value ngapType.MessageIdentifier, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("MessageIdentifier: BitString: sizeMin:16,sizeMax:16")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func MaximumIntegrityProtectedDataRate (value ngapType.MaximumIntegrityProtectedDataRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("MaximumIntegrityProtectedDataRate: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func MICOModeIndication (value ngapType.MICOModeIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("MICOModeIndication: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func NASPDU(value ngapType.NASPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASPDU: "+ err.Error())
	}
	return binData, bitEnd, err
}

func NASSecurityParametersFromNGRAN(value ngapType.NASSecurityParametersFromNGRAN, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NASSecurityParametersFromNGRAN: "+ err.Error())
	}
	return binData, bitEnd, err
}

func NetworkInstance (value ngapType.NetworkInstance, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 256 {
		return binData, bitEnd, errors.New("NetworkInstance: Integer: valueMin:1,valueMax:256")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func NewSecurityContextInd (value ngapType.NewSecurityContextInd, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("NewSecurityContextInd: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func NextHopChainingCount (value ngapType.NextHopChainingCount, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 7 {
		return binData, bitEnd, errors.New("NextHopChainingCount: Integer: valueMin:0,valueMax:7")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func NextPagingAreaScope (value ngapType.NextPagingAreaScope, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("NextPagingAreaScope: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func NGRANTraceID (value ngapType.NGRANTraceID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 8 {
		return binData, bitEnd, errors.New("NGRANTraceID: OctetString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func NotificationCause (value ngapType.NotificationCause, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("NotificationCause: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func NotificationControl (value ngapType.NotificationControl, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("NotificationControl: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func NRCellIdentity (value ngapType.NRCellIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 36 || len(value.Value.Bytes) != 5 {
		return binData, bitEnd, errors.New("NRCellIdentity: BitString: sizeMin:36,sizeMax:36")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func NRencryptionAlgorithms (value ngapType.NRencryptionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("NRencryptionAlgorithms: BitString: sizeMin:16,sizeMax:16")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func NRintegrityProtectionAlgorithms (value ngapType.NRintegrityProtectionAlgorithms, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("NRintegrityProtectionAlgorithms: BitString: sizeMin:16,sizeMax:16")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func NRPPaPDU(value ngapType.NRPPaPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("NRPPaPDU: "+ err.Error())
	}
	return binData, bitEnd, err
}

func NumberOfBroadcasts (value ngapType.NumberOfBroadcasts, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 65535 {
		return binData, bitEnd, errors.New("NumberOfBroadcasts: Integer: valueMin:0,valueMax:65535")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func NumberOfBroadcastsRequested (value ngapType.NumberOfBroadcastsRequested, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 65535 {
		return binData, bitEnd, errors.New("NumberOfBroadcastsRequested: Integer: valueMin:0,valueMax:65535")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func OverloadAction (value ngapType.OverloadAction, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4 {
		return binData, bitEnd, errors.New("OverloadAction: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func PacketDelayBudget (value ngapType.PacketDelayBudget, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1023 {
		return binData, bitEnd, errors.New("PacketDelayBudget: Integer: valueMin:0,valueMax:1023")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PacketLossRate (value ngapType.PacketLossRate, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1000 {
		return binData, bitEnd, errors.New("PacketLossRate: Integer: valueMin:0,valueMax:1000")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PagingAttemptCount (value ngapType.PagingAttemptCount, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 16 {
		return binData, bitEnd, errors.New("PagingAttemptCount: Integer: valueMin:1,valueMax:16")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PagingDRX (value ngapType.PagingDRX, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4 {
		return binData, bitEnd, errors.New("PagingDRX: Enumerated: valueMin:0,valueMax:4")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func PagingOrigin (value ngapType.PagingOrigin, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("PagingOrigin: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func PagingPriority (value ngapType.PagingPriority, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 8 {
		return binData, bitEnd, errors.New("PagingPriority: Enumerated: valueMin:0,valueMax:8")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func PDUSessionID (value ngapType.PDUSessionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("PDUSessionID: Integer: valueMin:0,valueMax:255")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PDUSessionType (value ngapType.PDUSessionType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 5 {
		return binData, bitEnd, errors.New("PDUSessionType: Enumerated: valueMin:0,valueMax:5")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func PeriodicRegistrationUpdateTimer (value ngapType.PeriodicRegistrationUpdateTimer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("PeriodicRegistrationUpdateTimer: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func PLMNIdentity (value ngapType.PLMNIdentity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("PLMNIdentity: OctetString: sizeMin:3,sizeMax:3")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func PortNumber (value ngapType.PortNumber, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("PortNumber: OctetString: sizeMin:2,sizeMax:2")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func PreEmptionCapability (value ngapType.PreEmptionCapability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("PreEmptionCapability: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func PreEmptionVulnerability (value ngapType.PreEmptionVulnerability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("PreEmptionVulnerability: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func PriorityLevelARP (value ngapType.PriorityLevelARP, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 15 {
		return binData, bitEnd, errors.New("PriorityLevelARP: Integer: valueMin:1,valueMax:15")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func PriorityLevelQos (value ngapType.PriorityLevelQos, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 127 {
		return binData, bitEnd, errors.New("PriorityLevelQos: Integer: valueMin:1,valueMax:127")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func QosFlowIdentifier (value ngapType.QosFlowIdentifier, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 63 {
		return binData, bitEnd, errors.New("QosFlowIdentifier: Integer: valueMin:0,valueMax:63")
	}
	binData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func QosMonitoringRequest (value ngapType.QosMonitoringRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("QosMonitoringRequest: Enumerated: valueMin:0,valueMax:3")
	}
	binData, bitEnd = EncodeEnumerated(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func RANNodeName (value ngapType.RANNodeName, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 150 {
		return binData, bitEnd, errors.New("RANNodeName: PrintableString: sizeMin:1,sizeMax:150")
	}
	binData, bitEnd = EncodePrintableString(value.Value,1, 150, binData, bitEnd)
	return binData, bitEnd, nil
}

func RANPagingPriority (value ngapType.RANPagingPriority, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 256 {
		return binData, bitEnd, errors.New("RANPagingPriority: Integer: valueMin:1,valueMax:256")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func RANUENGAPID (value ngapType.RANUENGAPID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4294967295 {
		return binData, bitEnd, errors.New("RANUENGAPID: Integer: valueMin:0,valueMax:4294967295")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func RATInformation (value ngapType.RATInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("RATInformation: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func RATRestrictionInformation (value ngapType.RATRestrictionInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 8 || len(value.Value.Bytes) != 1 {
		return binData, bitEnd, errors.New("RATRestrictionInformation: BitString: sizeMin:8,sizeMax:8")
	}
	binData, bitEnd = EncodeBitStringFixExx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func RedirectionVoiceFallback (value ngapType.RedirectionVoiceFallback, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("RedirectionVoiceFallback: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ReflectiveQosAttribute (value ngapType.ReflectiveQosAttribute, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("ReflectiveQosAttribute: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func RelativeAMFCapacity (value ngapType.RelativeAMFCapacity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("RelativeAMFCapacity: Integer: valueMin:0,valueMax:255")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ReportArea (value ngapType.ReportArea, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("ReportArea: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func RepetitionPeriod (value ngapType.RepetitionPeriod, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 131071 {
		return binData, bitEnd, errors.New("RepetitionPeriod: Integer: valueMin:0,valueMax:131071")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func ResetAll (value ngapType.ResetAll, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("ResetAll: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func RNCID (value ngapType.RNCID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4095 {
		return binData, bitEnd, errors.New("RNCID: Integer: valueMin:0,valueMax:4095")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func RoutingID(value ngapType.RoutingID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RoutingID: "+ err.Error())
	}
	return binData, bitEnd, err
}

func RRCContainer(value ngapType.RRCContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("RRCContainer: "+ err.Error())
	}
	return binData, bitEnd, err
}

func RRCEstablishmentCause (value ngapType.RRCEstablishmentCause, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 11 {
		return binData, bitEnd, errors.New("RRCEstablishmentCause: Enumerated: valueMin:0,valueMax:11")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 4, binData, bitEnd)
	return binData, bitEnd, nil
}

func RRCInactiveTransitionReportRequest (value ngapType.RRCInactiveTransitionReportRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("RRCInactiveTransitionReportRequest: Enumerated: valueMin:0,valueMax:3")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func RRCState (value ngapType.RRCState, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("RRCState: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func GNBSetID (value ngapType.GNBSetID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 22 || len(value.Value.Bytes) != 3 {
		return binData, bitEnd, errors.New("GNBSetID: BitString: sizeMin:22,sizeMax:22")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func SD (value ngapType.SD, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("SD: OctetString: sizeMin:3,sizeMax:3")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func SecurityKey (value ngapType.SecurityKey, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 256 || len(value.Value.Bytes) != 32 {
		return binData, bitEnd, errors.New("SecurityKey: BitString: sizeMin:256,sizeMax:256")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func SerialNumber (value ngapType.SerialNumber, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength != 16 || len(value.Value.Bytes) != 2 {
		return binData, bitEnd, errors.New("SerialNumber: BitString: sizeMin:16,sizeMax:16")
	}
	binData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func SgNBUEX2APID (value ngapType.SgNBUEX2APID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4294967295 {
		return binData, bitEnd, errors.New("SgNBUEX2APID: Integer: valueMin:0,valueMax:4294967295")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func SONInformationRequest (value ngapType.SONInformationRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("SONInformationRequest: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func SourceOfUEActivityBehaviourInformation (value ngapType.SourceOfUEActivityBehaviourInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("SourceOfUEActivityBehaviourInformation: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func SourceToTargetTransparentContainer(value ngapType.SourceToTargetTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("SourceToTargetTransparentContainer: "+ err.Error())
	}
	return binData, bitEnd, err
}

func SRVCCOperationPossible (value ngapType.SRVCCOperationPossible, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("SRVCCOperationPossible: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ConfiguredNSSAI (value ngapType.ConfiguredNSSAI, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 128 {
		return binData, bitEnd, errors.New("ConfiguredNSSAI: OctetString: sizeMin:128,sizeMax:128")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func RejectedNSSAIinPLMN (value ngapType.RejectedNSSAIinPLMN, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 32 {
		return binData, bitEnd, errors.New("RejectedNSSAIinPLMN: OctetString: sizeMin:32,sizeMax:32")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func RejectedNSSAIinTA (value ngapType.RejectedNSSAIinTA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 32 {
		return binData, bitEnd, errors.New("RejectedNSSAIinTA: OctetString: sizeMin:32,sizeMax:32")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func SST (value ngapType.SST, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 1 {
		return binData, bitEnd, errors.New("SST: OctetString: sizeMin:1,sizeMax:1")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func TAC (value ngapType.TAC, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 3 {
		return binData, bitEnd, errors.New("TAC: OctetString: sizeMin:3,sizeMax:3")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func TargetToSourceTransparentContainer(value ngapType.TargetToSourceTransparentContainer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("TargetToSourceTransparentContainer: "+ err.Error())
	}
	return binData, bitEnd, err
}

func TimerApproachForGUAMIRemoval (value ngapType.TimerApproachForGUAMIRemoval, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("TimerApproachForGUAMIRemoval: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func TimeStamp (value ngapType.TimeStamp, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 4 {
		return binData, bitEnd, errors.New("TimeStamp: OctetString: sizeMin:4,sizeMax:4")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func TimeToWait (value ngapType.TimeToWait, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 6 {
		return binData, bitEnd, errors.New("TimeToWait: Enumerated: valueMin:0,valueMax:6")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func TimeUEStayedInCell (value ngapType.TimeUEStayedInCell, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 4095 {
		return binData, bitEnd, errors.New("TimeUEStayedInCell: Integer: valueMin:0,valueMax:4095")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func TimeUEStayedInCellEnhancedGranularity (value ngapType.TimeUEStayedInCellEnhancedGranularity, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 40950 {
		return binData, bitEnd, errors.New("TimeUEStayedInCellEnhancedGranularity: Integer: valueMin:0,valueMax:40950")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func TNLAddressWeightFactor (value ngapType.TNLAddressWeightFactor, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 255 {
		return binData, bitEnd, errors.New("TNLAddressWeightFactor: Integer: valueMin:0,valueMax:255")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func TNLAssociationUsage (value ngapType.TNLAssociationUsage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("TNLAssociationUsage: Enumerated: valueMin:0,valueMax:3")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func TraceDepth (value ngapType.TraceDepth, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 6 {
		return binData, bitEnd, errors.New("TraceDepth: Enumerated: valueMin:0,valueMax:6")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func TrafficLoadReductionIndication (value ngapType.TrafficLoadReductionIndication, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 1 || value.Value > 99 {
		return binData, bitEnd, errors.New("TrafficLoadReductionIndication: Integer: valueMin:1,valueMax:99")
	}
	binData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)
	return binData, bitEnd, nil
}

func TransportLayerAddress (value ngapType.TransportLayerAddress, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value.BitLength < 1 || value.Value.BitLength > 160 || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {
		return binData, bitEnd, errors.New("TransportLayerAddress: BitString: sizeMin:1,sizeMax:160")
	}
	binData, bitEnd = EncodeBitStringExtExx(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func TypeOfError (value ngapType.TypeOfError, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 2 {
		return binData, bitEnd, errors.New("TypeOfError: Enumerated: valueMin:0,valueMax:2")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func UEContextRequest (value ngapType.UEContextRequest, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("UEContextRequest: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func UEPresence (value ngapType.UEPresence, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 3 {
		return binData, bitEnd, errors.New("UEPresence: Enumerated: valueMin:0,valueMax:3")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 3, binData, bitEnd)
	return binData, bitEnd, nil
}

func UERadioCapability(value ngapType.UERadioCapability, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapability: "+ err.Error())
	}
	return binData, bitEnd, err
}

func UERadioCapabilityForPagingOfNR(value ngapType.UERadioCapabilityForPagingOfNR, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingOfNR: "+ err.Error())
	}
	return binData, bitEnd, err
}

func UERadioCapabilityForPagingOfEUTRA(value ngapType.UERadioCapabilityForPagingOfEUTRA, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	var err error
	binData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)
	if err != nil {
		return binData, bitEnd, errors.New("UERadioCapabilityForPagingOfEUTRA: "+ err.Error())
	}
	return binData, bitEnd, err
}

func UERetentionInformation (value ngapType.UERetentionInformation, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("UERetentionInformation: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func ULForwarding (value ngapType.ULForwarding, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if value.Value < 0 || value.Value > 1 {
		return binData, bitEnd, errors.New("ULForwarding: Enumerated: valueMin:0,valueMax:1")
	}
	binData, bitEnd = EncodeEnumeratedExt(value.Value, 2, binData, bitEnd)
	return binData, bitEnd, nil
}

func WarningAreaCoordinates (value ngapType.WarningAreaCoordinates, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 1024 {
		return binData, bitEnd, errors.New("WarningAreaCoordinates: OctetString: sizeMin:1,sizeMax:1024")
	}
	binData, bitEnd = EncodeOctetStringExt(value.Value,1, 1024, binData, bitEnd)
	return binData, bitEnd, nil
}

func WarningMessageContents (value ngapType.WarningMessageContents, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) < 1 || len(value.Value) > 9600 {
		return binData, bitEnd, errors.New("WarningMessageContents: OctetString: sizeMin:1,sizeMax:9600")
	}
	binData, bitEnd = EncodeOctetStringExt(value.Value,1, 9600, binData, bitEnd)
	return binData, bitEnd, nil
}

func WarningSecurityInfo (value ngapType.WarningSecurityInfo, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 50 {
		return binData, bitEnd, errors.New("WarningSecurityInfo: OctetString: sizeMin:50,sizeMax:50")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}

func WarningType (value ngapType.WarningType, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	if len(value.Value) != 2 {
		return binData, bitEnd, errors.New("WarningType: OctetString: sizeMin:2,sizeMax:2")
	}
	binData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)
	return binData, bitEnd, nil
}
