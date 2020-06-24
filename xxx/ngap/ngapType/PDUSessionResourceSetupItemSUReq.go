// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           PDUSessionID
	PDUSessionNASPDU                       *NASPDU `vht:"optional"`
	SNSSAI                                 SNSSAI  `vht:"valueExt"`
	PDUSessionResourceSetupRequestTransfer OctetString
	IEExtensions                           *ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs `vht:"optional"`
}
