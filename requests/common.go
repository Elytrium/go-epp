package requests

import (
	"encoding/xml"
	"strconv"
	"time"
)

type EPPContainer struct {
	XMLName         xml.Name            `xml:"http://www.ripn.net/epp/ripn-epp-1.0 epp"`
	CommandRequest  *EPPCommandRequest  `xml:",omitempty"`
	CommandResponse *EPPCommandResponse `xml:",omitempty"`
	HelloRequest    *EPPHelloRequest    `xml:",omitempty"`
	HelloResponse   *EPPHelloResponse   `xml:",omitempty"`
}

type EPPCommandRequest struct {
	XMLName             xml.Name `xml:"command"`
	Body                interface{}
	ClientTransactionID string `xml:"clTRID"`
}

type EPPMessage struct {
	Text     string `xml:",chardata"`
	Language string `xml:"lang,attr"`
}

type EPPResultExtData struct {
	Reason string `xml:"reason"`
}

type EPPMessageQueue struct {
	Count   int        `xml:"count,attr"`
	ID      string     `xml:"id,attr"`
	QDate   time.Time  `xml:"qDate"`
	Message EPPMessage `xml:"msg"`
}

type EPPCommandResponse struct {
	XMLName xml.Name `xml:"response"`
	Result  struct {
		Code    int               `xml:"code,attr"`
		Message EPPMessage        `xml:"msg"`
		ExtData *EPPResultExtData `xml:"extData,omitempty"`
	} `xml:"result"`
	MessageQueue   *EPPMessageQueue `xml:"msgQ,omitempty"`
	Data           *EPPResponseData `xml:"resData,omitempty"`
	TransactionIDs struct {
		ClientTransactionID string `xml:"clTRID"`
		ServerTransactionID string `xml:"svTRID"`
	} `xml:"trID"`
}

type EPPResponseData struct {
	*EPPCreateResponse `xml:",omitempty"`
	*EPPInfoResponse   `xml:",omitempty"`
	TransferResponse   *EPPTransferResponse `xml:",omitempty"`
}

func WrapEPPCommand(body interface{}) EPPCommandRequest {
	return EPPCommandRequest{
		Body:                body,
		ClientTransactionID: strconv.FormatInt(time.Now().UnixNano(), 16),
	}
}
