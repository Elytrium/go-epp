package types

import (
	"encoding/xml"
	"time"
)

type EPPHostAddrIPType string

const (
	EPPHostAddrIPTypeV4 EPPHostAddrIPType = "v4"
	EPPHostAddrIPTypeV6 EPPHostAddrIPType = "v6"
)

type EPPHostAddr struct {
	XMLName xml.Name          `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:addr"`
	IPType  EPPHostAddrIPType `xml:"ip,attr"`
	IP      string
}

type EPPHostStatusType string

const (
	EPPHostStatusOK                     EPPHostStatusType = "ok"
	EPPHostStatusLinked                 EPPHostStatusType = "linked"
	EPPHostStatusServerUpdateProhibited EPPHostStatusType = "serverUpdateProhibited"
	EPPHostStatusServerDeleteProhibited EPPHostStatusType = "serverDeleteProhibited"
	EPPHostStatusPendingCreate          EPPHostStatusType = "pendingCreate"
	EPPHostStatusPendingDelete          EPPHostStatusType = "pendingDelete"
	EPPHostStatusPendingUpdate          EPPHostStatusType = "pendingUpdate"
)

type EPPHostStatus struct {
	XMLName xml.Name          `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:status"`
	Type    EPPHostStatusType `xml:"s,attr"`
}

type EPPHostData struct {
	Addresses []EPPHostAddr
}

type EPPHostExtendedData struct {
	ClientID    string    `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:clID"`
	CreatorID   string    `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:crID"`
	CreatedDate time.Time `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:crDate"`
}
