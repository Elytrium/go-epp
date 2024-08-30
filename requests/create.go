package requests

import (
	"encoding/xml"
	"time"

	"github.com/Elytrium/go-epp/types"
)

type EPPCreateRequest struct {
	XMLName xml.Name `xml:"create"`
	Body    interface{}
}

type EPPCreateResponse struct {
	CreateContactResponse *EPPCreateContactResponse `xml:",omitempty"`
	CreateHostResponse    *EPPCreateHostResponse    `xml:",omitempty"`
	CreateDomainResponse  *EPPCreateDomainResponse  `xml:",omitempty"`
}

type EPPCreateContactRequest struct {
	XMLName   xml.Name `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:create"`
	ContactID string   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:id"`
	types.EPPContactData
}

type EPPCreateContactResponse struct {
	XMLName     xml.Name  `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:creData"`
	ContactID   string    `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:id"`
	CreatedDate time.Time `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:crDate"`
}

type EPPCreateHostRequest struct {
	XMLName  xml.Name `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:create"`
	HostName string   `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:name"`
	types.EPPHostData
}

type EPPCreateHostResponse struct {
	XMLName     xml.Name  `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:creData"`
	HostName    string    `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:name"`
	CreatedDate time.Time `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:crDate"`
}

type EPPCreateDomainRequest struct {
	XMLName    xml.Name `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:create"`
	DomainName string   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
	types.EPPDomainData
}

type EPPCreateDomainResponse struct {
	XMLName     xml.Name  `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:creData"`
	DomainName  string    `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
	CreatedDate time.Time `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:crDate"`
	ExpiryDate  time.Time `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:exDate"`
}

func NewRIPNEPPCreatePersonRequest(contactID string, person types.EPPContactPerson, verified bool, oidEsia string) EPPCommandRequest {
	if oidEsia == "" {
		oidEsia = "UNKNOWN"
	}

	unverifiedObj := &struct{}{}
	verifiedObj := &struct{}{}

	if verified {
		unverifiedObj = nil
	} else {
		verifiedObj = nil
	}

	return WrapEPPCommand(EPPCreateRequest{
		Body: EPPCreateContactRequest{
			ContactID: contactID,
			EPPContactData: types.EPPContactData{
				Person:     &person,
				Unverified: unverifiedObj,
				Verified:   verifiedObj,
				OidEsia:    oidEsia,
			},
		},
	})
}

func NewRIPNEPPCreateOrganizationRequest(contactID string, organization types.EPPContactOrganization, verified bool, oidEsia string) EPPCommandRequest {
	if oidEsia == "" {
		oidEsia = "UNKNOWN"
	}

	unverifiedObj := &struct{}{}
	verifiedObj := &struct{}{}

	if verified {
		unverifiedObj = nil
	} else {
		verifiedObj = nil
	}

	return WrapEPPCommand(EPPCreateRequest{
		Body: EPPCreateContactRequest{
			ContactID: contactID,
			EPPContactData: types.EPPContactData{
				Organization: &organization,
				Unverified:   unverifiedObj,
				Verified:     verifiedObj,
				OidEsia:      oidEsia,
			},
		},
	})
}

func NewRIPNEPPCreateHostRequest(hostName string, addresses []types.EPPHostAddr) EPPCommandRequest {
	return WrapEPPCommand(EPPCreateRequest{
		Body: EPPCreateHostRequest{
			HostName: hostName,
			EPPHostData: types.EPPHostData{
				Addresses: addresses,
			},
		},
	})
}

func NewRIPNEPPCreateDomainRequest(domainName string, period types.EPPDomainPeriod, registrant string,
	description []string, authInfo *types.EPPDomainAuthInfo) EPPCommandRequest {
	return WrapEPPCommand(EPPCreateRequest{
		Body: EPPCreateDomainRequest{
			DomainName: domainName,
			EPPDomainData: types.EPPDomainData{
				Period:      &period,
				Registrant:  registrant,
				Description: description,
				AuthInfo:    authInfo,
			},
		},
	})
}
