package types

import (
	"encoding/xml"
	"time"
)

type EPPDomainPeriodUnit string

const (
	EPPDomainPeriodUnitYear EPPDomainPeriodUnit = "y"
)

type EPPDomainPeriod struct {
	XMLName xml.Name            `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:period"`
	Unit    EPPDomainPeriodUnit `xml:"unit,attr"`
	Period  uint                `xml:",chardata"`
}

type EPPDomainNS struct {
	XMLName     xml.Name `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:ns"`
	HostObjects []string `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:hostObj"`
}

type EPPDomainStatusType string

const (
	EPPDomainStatusOK                       EPPDomainStatusType = "ok"
	EPPDomainStatusInactive                 EPPDomainStatusType = "inactive"
	EPPDomainStatusServerHold               EPPDomainStatusType = "serverHold"
	EPPDomainStatusServerUpdateProhibited   EPPDomainStatusType = "serverUpdateProhibited"
	EPPDomainStatusServerDeleteProhibited   EPPDomainStatusType = "serverDeleteProhibited"
	EPPDomainStatusServerRenewProhibited    EPPDomainStatusType = "serverRenewProhibited"
	EPPDomainStatusServerTransferProhibited EPPDomainStatusType = "serverTransferProhibited"
	EPPDomainStatusPendingCreate            EPPDomainStatusType = "pendingCreate"
	EPPDomainStatusPendingDelete            EPPDomainStatusType = "pendingDelete"
	EPPDomainStatusPendingUpdate            EPPDomainStatusType = "pendingUpdate"
	EPPDomainStatusPendingRenew             EPPDomainStatusType = "pendingRenew"
	EPPDomainStatusPendingTransfer          EPPDomainStatusType = "pendingTransfer"
)

type EPPDomainStatus struct {
	XMLName xml.Name            `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:status"`
	Type    EPPDomainStatusType `xml:"s,attr"`
}

type EPPDomainData struct {
	Period      *EPPDomainPeriod   `xml:",omitempty"`
	Registrant  string             `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:registrant,omitempty"`
	Description []string           `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:description,omitempty"`
	NS          *EPPDomainNS       `xml:",omitempty"`
	AuthInfo    *EPPDomainAuthInfo `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:authInfo,omitempty"`
}

type EPPDomainExtendedData struct {
	ClientID    string    `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:clID"`
	CreatorID   string    `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:crID"`
	CreatedDate time.Time `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:crDate"`
	ExpiryDate  time.Time `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:exDate"`
}
