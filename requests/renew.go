package requests

import (
	"encoding/xml"

	"github.com/Elytrium/go-epp/types"
)

type EPPRenewRequest struct {
	XMLName xml.Name `xml:"renew"`
	Body    interface{}
}

type EPPRenewDomainRequest struct {
	XMLName           xml.Name `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:renew"`
	DomainName        string   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
	CurrentExpiryDate string   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:curExpDate"`
	Period            types.EPPDomainPeriod
}

func NewRIPNEPPRenewDomainRequest(domainName string, curExpDate string, period types.EPPDomainPeriod) EPPCommandRequest {
	return WrapEPPCommand(EPPRenewRequest{
		Body: EPPRenewDomainRequest{
			DomainName:        domainName,
			CurrentExpiryDate: curExpDate,
			Period:            period,
		},
	})
}
