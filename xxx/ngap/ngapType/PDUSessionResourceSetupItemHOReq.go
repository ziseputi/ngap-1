// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            PDUSessionID
	SNSSAI                  SNSSAI `vht:"valueExt"`
	HandoverRequestTransfer OctetString
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs `vht:"optional"`
}
