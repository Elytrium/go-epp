package requests

import "encoding/xml"

type EPPLogoutCommand struct {
	XMLName xml.Name `xml:"logout"`
}

func NewEPPLogoutRequest() EPPCommandRequest {
	return WrapEPPCommand(EPPLogoutCommand{})
}
