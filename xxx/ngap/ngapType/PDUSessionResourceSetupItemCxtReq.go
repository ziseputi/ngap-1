// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceSetupItemCxtReq struct {
	PDUSessionID                           PDUSessionID
	NASPDU                                 *NASPDU `vht:"optional"`
	SNSSAI                                 SNSSAI  `vht:"valueExt"`
	PDUSessionResourceSetupRequestTransfer OctetString
	IEExtensions                           *ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs `vht:"optional"`
}
