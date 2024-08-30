package requests

import "encoding/xml"

type EPPHelloRequest struct {
	XMLName xml.Name `xml:"hello"`
}

type EPPHelloResponse struct {
	XMLName     xml.Name `xml:"greeting"`
	ServerID    string   `xml:"svID"`
	ServerDate  string   `xml:"svDate"`
	ServiceMenu struct {
		Version    string   `xml:"version"`
		Language   []string `xml:"lang"`
		ObjectURIs []struct {
			URI            string `xml:",chardata"`
			XSD            string `xml:",attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
		} `xml:"objURI"`
	} `xml:"svcMenu"`
	DataCollectionPolicy struct {
		Access struct {
			All              *struct{} `xml:"all"`
			None             *struct{} `xml:"none"`
			Null             *struct{} `xml:"null"`
			Personal         *struct{} `xml:"personal"`
			PersonalAndOther *struct{} `xml:"personalAndOther"`
		} `xml:"access"`
		Statement struct {
			Purpose struct {
				Admin        *struct{} `xml:"admin"`
				Provisioning *struct{} `xml:"prov"`
				Contact      *struct{} `xml:"contact"`
				Other        *struct{} `xml:"other"`
			} `xml:"purpose"`
			Recipient struct {
				Other     *struct{} `xml:"other"`
				Ours      *struct{} `xml:"ours"`
				Public    *struct{} `xml:"public"`
				Same      *struct{} `xml:"same"`
				Unrelated *struct{} `xml:"unrelated"`
			} `xml:"recipient"`
			Retention struct {
				Business   *struct{} `xml:"business"`
				Indefinite *struct{} `xml:"indefinite"`
				Legal      *struct{} `xml:"legal"`
				None       *struct{} `xml:"none"`
				Stated     *struct{} `xml:"stated"`
			} `xml:"retention"`
		} `xml:"statement"`
	} `xml:"dcp"`
}

func NewEPPHelloRequest() EPPHelloRequest {
	return EPPHelloRequest{}
}
