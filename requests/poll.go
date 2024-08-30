package requests

import "encoding/xml"

type EPPPollOperation string

const (
	EPPPollOperationRequest         = "req"
	EPPPollOperationAcknowledgement = "ack"
)

type EPPPollRequest struct {
	XMLName   xml.Name         `xml:"poll"`
	Operation EPPPollOperation `xml:"op,attr"`
	MessageID string           `xml:"msgID,attr,omitempty"`
}

// EPPPollResponse has any data only on transfer notifications
type EPPPollResponse EPPTransferResponse

func NewEPPPollRequest() EPPCommandRequest {
	return WrapEPPCommand(EPPPollRequest{
		Operation: EPPPollOperationRequest,
	})
}

func NewEPPPollAcknowledgement(id string) EPPCommandRequest {
	return WrapEPPCommand(EPPPollRequest{
		Operation: EPPPollOperationAcknowledgement,
		MessageID: id,
	})
}
