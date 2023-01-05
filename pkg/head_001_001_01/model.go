// Code generated automatically. DO NOT EDIT.

package head_001_001_01

import (
	"encoding/xml"
)

// AppHdr ...
type AppHdr *BusinessApplicationHeaderV01

// AddressType2Code ...
type AddressType2Code string

// AnyBICIdentifier ...
type AnyBICIdentifier string

// BICFIIdentifier ...
type BICFIIdentifier string

// BranchAndFinancialInstitutionIdentification5 ...
type BranchAndFinancialInstitutionIdentification5 struct {
	XMLName    *xml.Name                            `xml:"urn:BranchAndFinancialInstitutionIdentification5", json:"-,omitempty", bson:"-,omitempty"`
	FinInstnId *FinancialInstitutionIdentification8 `xml:"urn:FinInstnId", json:"FinInstnId", bson:"FinInstnId"`
	BrnchId    *BranchData2                         `xml:"urn:BrnchId", json:"BrnchId", bson:"BrnchId"`
}

// BranchData2 ...
type BranchData2 struct {
	XMLName *xml.Name       `xml:"urn:BranchData2", json:"-,omitempty", bson:"-,omitempty"`
	Id      string          `xml:"urn:Id", json:"Id", bson:"Id"`
	Nm      string          `xml:"urn:Nm", json:"Nm", bson:"Nm"`
	PstlAdr *PostalAddress6 `xml:"urn:PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
}

// BusinessApplicationHeader1 ...
type BusinessApplicationHeader1 struct {
	XMLName    *xml.Name          `xml:"urn:BusinessApplicationHeader1", json:"-,omitempty", bson:"-,omitempty"`
	CharSet    string             `xml:"urn:CharSet", json:"CharSet", bson:"CharSet"`
	Fr         *Party9Choice      `xml:"urn:Fr", json:"Fr", bson:"Fr"`
	To         *Party9Choice      `xml:"urn:To", json:"To", bson:"To"`
	BizMsgIdr  string             `xml:"urn:BizMsgIdr", json:"BizMsgIdr", bson:"BizMsgIdr"`
	MsgDefIdr  string             `xml:"urn:MsgDefIdr", json:"MsgDefIdr", bson:"MsgDefIdr"`
	BizSvc     string             `xml:"urn:BizSvc", json:"BizSvc", bson:"BizSvc"`
	CreDt      string             `xml:"urn:CreDt", json:"CreDt", bson:"CreDt"`
	CpyDplct   string             `xml:"urn:CpyDplct", json:"CpyDplct", bson:"CpyDplct"`
	PssblDplct bool               `xml:"urn:PssblDplct", json:"PssblDplct", bson:"PssblDplct"`
	Prty       string             `xml:"urn:Prty", json:"Prty", bson:"Prty"`
	Sgntr      *SignatureEnvelope `xml:"urn:Sgntr", json:"Sgntr", bson:"Sgntr"`
}

// BusinessApplicationHeaderV01 ...
type BusinessApplicationHeaderV01 struct {
	XMLName    *xml.Name                   `xml:"urn:BusinessApplicationHeaderV01", json:"-,omitempty", bson:"-,omitempty"`
	CharSet    string                      `xml:"urn:CharSet", json:"CharSet", bson:"CharSet"`
	Fr         *Party9Choice               `xml:"urn:Fr", json:"Fr", bson:"Fr"`
	To         *Party9Choice               `xml:"urn:To", json:"To", bson:"To"`
	BizMsgIdr  string                      `xml:"urn:BizMsgIdr", json:"BizMsgIdr", bson:"BizMsgIdr"`
	MsgDefIdr  string                      `xml:"urn:MsgDefIdr", json:"MsgDefIdr", bson:"MsgDefIdr"`
	BizSvc     string                      `xml:"urn:BizSvc", json:"BizSvc", bson:"BizSvc"`
	CreDt      string                      `xml:"urn:CreDt", json:"CreDt", bson:"CreDt"`
	CpyDplct   string                      `xml:"urn:CpyDplct", json:"CpyDplct", bson:"CpyDplct"`
	PssblDplct bool                        `xml:"urn:PssblDplct", json:"PssblDplct", bson:"PssblDplct"`
	Prty       string                      `xml:"urn:Prty", json:"Prty", bson:"Prty"`
	Sgntr      *SignatureEnvelope          `xml:"urn:Sgntr", json:"Sgntr", bson:"Sgntr"`
	Rltd       *BusinessApplicationHeader1 `xml:"urn:Rltd", json:"Rltd", bson:"Rltd"`
}

// BusinessMessagePriorityCode ...
type BusinessMessagePriorityCode string

// ClearingSystemIdentification2Choice ...
type ClearingSystemIdentification2Choice struct {
	XMLName *xml.Name `xml:"urn:ClearingSystemIdentification2Choice", json:"-,omitempty", bson:"-,omitempty"`
	Cd      string    `xml:"urn:Cd", json:"Cd", bson:"Cd"`
	Prtry   string    `xml:"urn:Prtry", json:"Prtry", bson:"Prtry"`
}

// ClearingSystemMemberIdentification2 ...
type ClearingSystemMemberIdentification2 struct {
	XMLName  *xml.Name                            `xml:"urn:ClearingSystemMemberIdentification2", json:"-,omitempty", bson:"-,omitempty"`
	ClrSysId *ClearingSystemIdentification2Choice `xml:"urn:ClrSysId", json:"ClrSysId", bson:"ClrSysId"`
	MmbId    string                               `xml:"urn:MmbId", json:"MmbId", bson:"MmbId"`
}

// ContactDetails2 ...
type ContactDetails2 struct {
	XMLName  *xml.Name `xml:"urn:ContactDetails2", json:"-,omitempty", bson:"-,omitempty"`
	NmPrfx   string    `xml:"urn:NmPrfx", json:"NmPrfx", bson:"NmPrfx"`
	Nm       string    `xml:"urn:Nm", json:"Nm", bson:"Nm"`
	PhneNb   string    `xml:"urn:PhneNb", json:"PhneNb", bson:"PhneNb"`
	MobNb    string    `xml:"urn:MobNb", json:"MobNb", bson:"MobNb"`
	FaxNb    string    `xml:"urn:FaxNb", json:"FaxNb", bson:"FaxNb"`
	EmailAdr string    `xml:"urn:EmailAdr", json:"EmailAdr", bson:"EmailAdr"`
	Othr     string    `xml:"urn:Othr", json:"Othr", bson:"Othr"`
}

// CopyDuplicate1Code ...
type CopyDuplicate1Code string

// CountryCode ...
type CountryCode string

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	XMLName     *xml.Name `xml:"urn:DateAndPlaceOfBirth", json:"-,omitempty", bson:"-,omitempty"`
	BirthDt     string    `xml:"urn:BirthDt", json:"BirthDt", bson:"BirthDt"`
	PrvcOfBirth string    `xml:"urn:PrvcOfBirth", json:"PrvcOfBirth", bson:"PrvcOfBirth"`
	CityOfBirth string    `xml:"urn:CityOfBirth", json:"CityOfBirth", bson:"CityOfBirth"`
	CtryOfBirth string    `xml:"urn:CtryOfBirth", json:"CtryOfBirth", bson:"CtryOfBirth"`
}

// ExternalClearingSystemIdentification1Code ...
type ExternalClearingSystemIdentification1Code string

// ExternalFinancialInstitutionIdentification1Code ...
type ExternalFinancialInstitutionIdentification1Code string

// ExternalOrganisationIdentification1Code ...
type ExternalOrganisationIdentification1Code string

// ExternalPersonIdentification1Code ...
type ExternalPersonIdentification1Code string

// FinancialIdentificationSchemeName1Choice ...
type FinancialIdentificationSchemeName1Choice struct {
	XMLName *xml.Name `xml:"urn:FinancialIdentificationSchemeName1Choice", json:"-,omitempty", bson:"-,omitempty"`
	Cd      string    `xml:"urn:Cd", json:"Cd", bson:"Cd"`
	Prtry   string    `xml:"urn:Prtry", json:"Prtry", bson:"Prtry"`
}

// FinancialInstitutionIdentification8 ...
type FinancialInstitutionIdentification8 struct {
	XMLName     *xml.Name                            `xml:"urn:FinancialInstitutionIdentification8", json:"-,omitempty", bson:"-,omitempty"`
	BICFI       string                               `xml:"urn:BICFI", json:"BICFI", bson:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"urn:ClrSysMmbId", json:"ClrSysMmbId", bson:"ClrSysMmbId"`
	Nm          string                               `xml:"urn:Nm", json:"Nm", bson:"Nm"`
	PstlAdr     *PostalAddress6                      `xml:"urn:PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"urn:Othr", json:"Othr", bson:"Othr"`
}

// GenericFinancialIdentification1 ...
type GenericFinancialIdentification1 struct {
	XMLName *xml.Name                                 `xml:"urn:GenericFinancialIdentification1", json:"-,omitempty", bson:"-,omitempty"`
	Id      string                                    `xml:"urn:Id", json:"Id", bson:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"urn:SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string                                    `xml:"urn:Issr", json:"Issr", bson:"Issr"`
}

// GenericOrganisationIdentification1 ...
type GenericOrganisationIdentification1 struct {
	XMLName *xml.Name                                    `xml:"urn:GenericOrganisationIdentification1", json:"-,omitempty", bson:"-,omitempty"`
	Id      string                                       `xml:"urn:Id", json:"Id", bson:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string                                       `xml:"urn:Issr", json:"Issr", bson:"Issr"`
}

// GenericPersonIdentification1 ...
type GenericPersonIdentification1 struct {
	XMLName *xml.Name                              `xml:"urn:GenericPersonIdentification1", json:"-,omitempty", bson:"-,omitempty"`
	Id      string                                 `xml:"urn:Id", json:"Id", bson:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string                                 `xml:"urn:Issr", json:"Issr", bson:"Issr"`
}

// ISODate ...
type ISODate string

// ISONormalisedDateTime ...
type ISONormalisedDateTime string

// Max140Text ...
type Max140Text string

// Max16Text ...
type Max16Text string

// Max2048Text ...
type Max2048Text string

// Max35Text ...
type Max35Text string

// Max70Text ...
type Max70Text string

// NamePrefix1Code ...
type NamePrefix1Code string

// OrganisationIdentification7 ...
type OrganisationIdentification7 struct {
	XMLName *xml.Name                             `xml:"urn:OrganisationIdentification7", json:"-,omitempty", bson:"-,omitempty"`
	AnyBIC  string                                `xml:"urn:AnyBIC", json:"AnyBIC", bson:"AnyBIC"`
	Othr    []*GenericOrganisationIdentification1 `xml:"urn:Othr", json:"Othr", bson:"Othr"`
}

// OrganisationIdentificationSchemeName1Choice ...
type OrganisationIdentificationSchemeName1Choice struct {
	XMLName *xml.Name `xml:"urn:OrganisationIdentificationSchemeName1Choice", json:"-,omitempty", bson:"-,omitempty"`
	Cd      string    `xml:"urn:Cd", json:"Cd", bson:"Cd"`
	Prtry   string    `xml:"urn:Prtry", json:"Prtry", bson:"Prtry"`
}

// Party10Choice ...
type Party10Choice struct {
	XMLName *xml.Name                    `xml:"urn:Party10Choice", json:"-,omitempty", bson:"-,omitempty"`
	OrgId   *OrganisationIdentification7 `xml:"urn:OrgId", json:"OrgId", bson:"OrgId"`
	PrvtId  *PersonIdentification5       `xml:"urn:PrvtId", json:"PrvtId", bson:"PrvtId"`
}

// Party9Choice ...
type Party9Choice struct {
	XMLName *xml.Name                                     `xml:"urn:Party9Choice", json:"-,omitempty", bson:"-,omitempty"`
	OrgId   *PartyIdentification42                        `xml:"urn:OrgId", json:"OrgId", bson:"OrgId"`
	FIId    *BranchAndFinancialInstitutionIdentification5 `xml:"urn:FIId", json:"FIId", bson:"FIId"`
}

// PartyIdentification42 ...
type PartyIdentification42 struct {
	XMLName   *xml.Name        `xml:"urn:PartyIdentification42", json:"-,omitempty", bson:"-,omitempty"`
	Nm        string           `xml:"urn:Nm", json:"Nm", bson:"Nm"`
	PstlAdr   *PostalAddress6  `xml:"urn:PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
	Id        *Party10Choice   `xml:"urn:Id", json:"Id", bson:"Id"`
	CtryOfRes string           `xml:"urn:CtryOfRes", json:"CtryOfRes", bson:"CtryOfRes"`
	CtctDtls  *ContactDetails2 `xml:"urn:CtctDtls", json:"CtctDtls", bson:"CtctDtls"`
}

// PersonIdentification5 ...
type PersonIdentification5 struct {
	XMLName         *xml.Name                       `xml:"urn:PersonIdentification5", json:"-,omitempty", bson:"-,omitempty"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth            `xml:"urn:DtAndPlcOfBirth", json:"DtAndPlcOfBirth", bson:"DtAndPlcOfBirth"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:Othr", json:"Othr", bson:"Othr"`
}

// PersonIdentificationSchemeName1Choice ...
type PersonIdentificationSchemeName1Choice struct {
	XMLName *xml.Name `xml:"urn:PersonIdentificationSchemeName1Choice", json:"-,omitempty", bson:"-,omitempty"`
	Cd      string    `xml:"urn:Cd", json:"Cd", bson:"Cd"`
	Prtry   string    `xml:"urn:Prtry", json:"Prtry", bson:"Prtry"`
}

// PhoneNumber ...
type PhoneNumber string

// PostalAddress6 ...
type PostalAddress6 struct {
	XMLName     *xml.Name `xml:"urn:PostalAddress6", json:"-,omitempty", bson:"-,omitempty"`
	AdrTp       string    `xml:"urn:AdrTp", json:"AdrTp", bson:"AdrTp"`
	Dept        string    `xml:"urn:Dept", json:"Dept", bson:"Dept"`
	SubDept     string    `xml:"urn:SubDept", json:"SubDept", bson:"SubDept"`
	StrtNm      string    `xml:"urn:StrtNm", json:"StrtNm", bson:"StrtNm"`
	BldgNb      string    `xml:"urn:BldgNb", json:"BldgNb", bson:"BldgNb"`
	PstCd       string    `xml:"urn:PstCd", json:"PstCd", bson:"PstCd"`
	TwnNm       string    `xml:"urn:TwnNm", json:"TwnNm", bson:"TwnNm"`
	CtrySubDvsn string    `xml:"urn:CtrySubDvsn", json:"CtrySubDvsn", bson:"CtrySubDvsn"`
	Ctry        string    `xml:"urn:Ctry", json:"Ctry", bson:"Ctry"`
	AdrLine     []string  `xml:"urn:AdrLine", json:"AdrLine", bson:"AdrLine"`
}

// SignatureEnvelope ...
type SignatureEnvelope struct {
	XMLName *xml.Name `xml:"urn:SignatureEnvelope", json:"-,omitempty", bson:"-,omitempty"`
}

// UnicodeChartsCode ...
type UnicodeChartsCode string

// YesNoIndicator ...
type YesNoIndicator bool
