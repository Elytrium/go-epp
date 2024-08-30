package types

import (
	"encoding/xml"
	"time"
)

type EPPRegistrarStatusType string

type EPPRegistrarStatus struct {
	XMLName xml.Name               `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:status"`
	Type    EPPRegistrarStatusType `xml:"s,attr"`
}

type EPPRegistrarOrganizationAddress struct {
	Organization string   `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:org"`
	Address      []string `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:address"`
}

type EPPRegistrarOrganizationLegalInfo struct {
	Address []string `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:address"`
}

type EPPRegistrarEmailType string

const (
	EPPRegistrarEmailTypeAuth       EPPRegistrarEmailType = "auth"
	EPPRegistrarEmailTypeAuthNotify EPPRegistrarEmailType = "authNotify"
	EPPRegistrarEmailTypeInfo       EPPRegistrarEmailType = "info"
	EPPRegistrarEmailTypeNotify     EPPRegistrarEmailType = "notify"
)

type EPPRegistrarEmail struct {
	XMLName xml.Name              `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:email"`
	Type    EPPRegistrarEmailType `xml:"type,attr"`
	Email   string                `xml:",chardata"`
}
type EPPRegistrarAddrIPType string

const (
	EPPRegistrarAddrIPTypeV4 EPPHostAddrIPType = "v4"
	EPPRegistrarAddrIPTypeV6 EPPHostAddrIPType = "v6"
)

type EPPRegistrarAddr struct {
	XMLName xml.Name          `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:addr"`
	IPType  EPPHostAddrIPType `xml:"ip,attr"`
	IP      string
}

type EPPRegistrarData struct {
	InternationalPostalInfo *EPPRegistrarOrganizationAddress   `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:intPostalInfo,omitempty"`
	LocalPostalInfo         *EPPRegistrarOrganizationAddress   `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:locPostalInfo,omitempty"`
	LegalInfo               *EPPRegistrarOrganizationLegalInfo `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:legalInfo,omitempty"`
	TaxpayerNumbers         string                             `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:taxpayerNumbers,omitempty"`
	Voice                   []string                           `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:voice,omitempty"`
	Fax                     []string                           `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:fax,omitempty"`
	Email                   []EPPRegistrarEmailType            `xml:",omitempty"`
	WWW                     []string                           `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:www,omitempty"`
	Whois                   []string                           `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:whois,omitempty"`
	AdminContact            []string                           `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:adminContact,omitempty"`
	Registrar               []EPPRegistrarAddr                 `xml:",omitempty"`
}

type EPPRegistrarExtendedData struct {
	CreatedDate  time.Time `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:crDate"`
	ModifiedDate time.Time `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:upDate"`
}
