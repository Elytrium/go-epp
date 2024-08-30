package types

import (
	"encoding/xml"
	"time"
)

type EPPContactOrganizationAddress struct {
	Organization string   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:org"`
	Address      []string `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:address"`
}

type EPPContactOrganizationLegalInfo struct {
	Address []string `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:address"`
}

type EPPContactPersonAddress struct {
	Name    string   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:name"`
	Address []string `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:address"`
}

type EPPContactOrganization struct {
	InternationalPostalInfo *EPPContactOrganizationAddress   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:intPostalInfo,omitempty"`
	LocalPostalInfo         *EPPContactOrganizationAddress   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:locPostalInfo,omitempty"`
	LegalInfo               *EPPContactOrganizationLegalInfo `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:legalInfo,omitempty"`
	TaxpayerNumbers         string                           `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:taxpayerNumbers,omitempty"`
	Voice                   []string                         `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:voice,omitempty"`
	Fax                     []string                         `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:fax,omitempty"`
	Email                   []string                         `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:email,omitempty"`
}

type EPPContactDiscloseFlag int

const (
	EPPContactDiscloseFlagDisabled EPPContactDiscloseFlag = 0
	EPPContactDiscloseFlagEnabled  EPPContactDiscloseFlag = 1
)

type EPPContactDisclose struct {
	XMLName           xml.Name               `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:disclose"`
	Flag              EPPContactDiscloseFlag `xml:"flag,attr"`
	InternationalName *struct{}              `xml:"intName,omitempty"`
}

type EPPContactPerson struct {
	InternationalPostalInfo *EPPContactPersonAddress `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:intPostalInfo,omitempty"`
	LocalPostalInfo         *EPPContactPersonAddress `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:locPostalInfo,omitempty"`
	TaxpayerNumbers         string                   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:taxpayerNumbers,omitempty"`
	Birthday                string                   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:birthday,omitempty"`
	Passport                []string                 `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:passport,omitempty"`
	Voice                   []string                 `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:voice,omitempty"`
	Fax                     []string                 `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:fax,omitempty"`
	Email                   []string                 `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:email,omitempty"`
}

type EPPContactStatusType string

const (
	EPPContactStatusOK                     EPPContactStatusType = "ok"
	EPPContactStatusLinked                 EPPContactStatusType = "linked"
	EPPContactStatusServerUpdateProhibited EPPContactStatusType = "serverUpdateProhibited"
	EPPContactStatusServerDeleteProhibited EPPContactStatusType = "serverDeleteProhibited"
	EPPContactStatusPendingCreate          EPPContactStatusType = "pendingCreate"
	EPPContactStatusPendingDelete          EPPContactStatusType = "pendingDelete"
	EPPContactStatusPendingUpdate          EPPContactStatusType = "pendingUpdate"
)

type EPPContactStatus struct {
	XMLName xml.Name             `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:status"`
	Type    EPPContactStatusType `xml:"s,attr"`
}

type EPPContactData struct {
	Organization *EPPContactOrganization `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:organization,omitempty"`
	Person       *EPPContactPerson       `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:person,omitempty"`
	Verified     *struct{}               `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:verified,omitempty"`
	Unverified   *struct{}               `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:unverified,omitempty"`
	OidEsia      string                  `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:oidEsia,omitempty"`
}

type EPPContactExtendedData struct {
	ClientID     string    `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:clID"`
	CreatorID    string    `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:crID"`
	CreatedDate  time.Time `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:crDate"`
	ModifierID   string    `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:upID"`
	ModifiedDate time.Time `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:upDate"`
}
