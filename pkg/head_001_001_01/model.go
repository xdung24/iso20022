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
	XMLName    xml.Name                             `json:"-,omitempty", bson:"-,omitempty", xml:""`
	FinInstnId *FinancialInstitutionIdentification8 `xml:"FinInstnId", json:"FinInstnId", bson:"FinInstnId"`
	BrnchId    *BranchData2                         `xml:"BrnchId", json:"BrnchId", bson:"BrnchId"`
}

// BranchData2 ...
type BranchData2 struct {
	XMLName xml.Name        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string          `xml:"Id", json:"Id", bson:"Id"`
	Nm      string          `xml:"Nm", json:"Nm", bson:"Nm"`
	PstlAdr *PostalAddress6 `xml:"PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
}

// BusinessApplicationHeader1 ...
type BusinessApplicationHeader1 struct {
	XMLName    xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CharSet    string             `xml:"CharSet", json:"CharSet", bson:"CharSet"`
	Fr         *Party9Choice      `xml:"Fr", json:"Fr", bson:"Fr"`
	To         *Party9Choice      `xml:"To", json:"To", bson:"To"`
	BizMsgIdr  string             `xml:"BizMsgIdr", json:"BizMsgIdr", bson:"BizMsgIdr"`
	MsgDefIdr  string             `xml:"MsgDefIdr", json:"MsgDefIdr", bson:"MsgDefIdr"`
	BizSvc     string             `xml:"BizSvc", json:"BizSvc", bson:"BizSvc"`
	CreDt      string             `xml:"CreDt", json:"CreDt", bson:"CreDt"`
	CpyDplct   string             `xml:"CpyDplct", json:"CpyDplct", bson:"CpyDplct"`
	PssblDplct bool               `xml:"PssblDplct", json:"PssblDplct", bson:"PssblDplct"`
	Prty       string             `xml:"Prty", json:"Prty", bson:"Prty"`
	Sgntr      *SignatureEnvelope `xml:"Sgntr", json:"Sgntr", bson:"Sgntr"`
}

// BusinessApplicationHeaderV01 ...
type BusinessApplicationHeaderV01 struct {
	XMLName    xml.Name                    `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CharSet    string                      `xml:"CharSet", json:"CharSet", bson:"CharSet"`
	Fr         *Party9Choice               `xml:"Fr", json:"Fr", bson:"Fr"`
	To         *Party9Choice               `xml:"To", json:"To", bson:"To"`
	BizMsgIdr  string                      `xml:"BizMsgIdr", json:"BizMsgIdr", bson:"BizMsgIdr"`
	MsgDefIdr  string                      `xml:"MsgDefIdr", json:"MsgDefIdr", bson:"MsgDefIdr"`
	BizSvc     string                      `xml:"BizSvc", json:"BizSvc", bson:"BizSvc"`
	CreDt      string                      `xml:"CreDt", json:"CreDt", bson:"CreDt"`
	CpyDplct   string                      `xml:"CpyDplct", json:"CpyDplct", bson:"CpyDplct"`
	PssblDplct bool                        `xml:"PssblDplct", json:"PssblDplct", bson:"PssblDplct"`
	Prty       string                      `xml:"Prty", json:"Prty", bson:"Prty"`
	Sgntr      *SignatureEnvelope          `xml:"Sgntr", json:"Sgntr", bson:"Sgntr"`
	Rltd       *BusinessApplicationHeader1 `xml:"Rltd", json:"Rltd", bson:"Rltd"`
}

// BusinessMessagePriorityCode ...
type BusinessMessagePriorityCode string

// ClearingSystemIdentification2Choice ...
type ClearingSystemIdentification2Choice struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Cd      string   `xml:"Cd", json:"Cd", bson:"Cd"`
	Prtry   string   `xml:"Prtry", json:"Prtry", bson:"Prtry"`
}

// ClearingSystemMemberIdentification2 ...
type ClearingSystemMemberIdentification2 struct {
	XMLName  xml.Name                             `json:"-,omitempty", bson:"-,omitempty", xml:""`
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId", json:"ClrSysId", bson:"ClrSysId"`
	MmbId    string                               `xml:"MmbId", json:"MmbId", bson:"MmbId"`
}

// ContactDetails2 ...
type ContactDetails2 struct {
	XMLName  xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	NmPrfx   string   `xml:"NmPrfx", json:"NmPrfx", bson:"NmPrfx"`
	Nm       string   `xml:"Nm", json:"Nm", bson:"Nm"`
	PhneNb   string   `xml:"PhneNb", json:"PhneNb", bson:"PhneNb"`
	MobNb    string   `xml:"MobNb", json:"MobNb", bson:"MobNb"`
	FaxNb    string   `xml:"FaxNb", json:"FaxNb", bson:"FaxNb"`
	EmailAdr string   `xml:"EmailAdr", json:"EmailAdr", bson:"EmailAdr"`
	Othr     string   `xml:"Othr", json:"Othr", bson:"Othr"`
}

// CopyDuplicate1Code ...
type CopyDuplicate1Code string

// CountryCode ...
type CountryCode string

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	XMLName     xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	BirthDt     string   `xml:"BirthDt", json:"BirthDt", bson:"BirthDt"`
	PrvcOfBirth string   `xml:"PrvcOfBirth", json:"PrvcOfBirth", bson:"PrvcOfBirth"`
	CityOfBirth string   `xml:"CityOfBirth", json:"CityOfBirth", bson:"CityOfBirth"`
	CtryOfBirth string   `xml:"CtryOfBirth", json:"CtryOfBirth", bson:"CtryOfBirth"`
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
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Cd      string   `xml:"Cd", json:"Cd", bson:"Cd"`
	Prtry   string   `xml:"Prtry", json:"Prtry", bson:"Prtry"`
}

// FinancialInstitutionIdentification8 ...
type FinancialInstitutionIdentification8 struct {
	XMLName     xml.Name                             `json:"-,omitempty", bson:"-,omitempty", xml:""`
	BICFI       string                               `xml:"BICFI", json:"BICFI", bson:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId", json:"ClrSysMmbId", bson:"ClrSysMmbId"`
	Nm          string                               `xml:"Nm", json:"Nm", bson:"Nm"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr", json:"Othr", bson:"Othr"`
}

// GenericFinancialIdentification1 ...
type GenericFinancialIdentification1 struct {
	XMLName xml.Name                                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string                                    `xml:"Id", json:"Id", bson:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string                                    `xml:"Issr", json:"Issr", bson:"Issr"`
}

// GenericOrganisationIdentification1 ...
type GenericOrganisationIdentification1 struct {
	XMLName xml.Name                                     `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string                                       `xml:"Id", json:"Id", bson:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string                                       `xml:"Issr", json:"Issr", bson:"Issr"`
}

// GenericPersonIdentification1 ...
type GenericPersonIdentification1 struct {
	XMLName xml.Name                               `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string                                 `xml:"Id", json:"Id", bson:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string                                 `xml:"Issr", json:"Issr", bson:"Issr"`
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
	XMLName xml.Name                              `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AnyBIC  string                                `xml:"AnyBIC", json:"AnyBIC", bson:"AnyBIC"`
	Othr    []*GenericOrganisationIdentification1 `xml:"Othr", json:"Othr", bson:"Othr"`
}

// OrganisationIdentificationSchemeName1Choice ...
type OrganisationIdentificationSchemeName1Choice struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Cd      string   `xml:"Cd", json:"Cd", bson:"Cd"`
	Prtry   string   `xml:"Prtry", json:"Prtry", bson:"Prtry"`
}

// Party10Choice ...
type Party10Choice struct {
	XMLName xml.Name                     `json:"-,omitempty", bson:"-,omitempty", xml:""`
	OrgId   *OrganisationIdentification7 `xml:"OrgId", json:"OrgId", bson:"OrgId"`
	PrvtId  *PersonIdentification5       `xml:"PrvtId", json:"PrvtId", bson:"PrvtId"`
}

// Party9Choice ...
type Party9Choice struct {
	XMLName xml.Name                                      `json:"-,omitempty", bson:"-,omitempty", xml:""`
	OrgId   *PartyIdentification42                        `xml:"OrgId", json:"OrgId", bson:"OrgId"`
	FIId    *BranchAndFinancialInstitutionIdentification5 `xml:"FIId", json:"FIId", bson:"FIId"`
}

// PartyIdentification42 ...
type PartyIdentification42 struct {
	XMLName   xml.Name         `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Nm        string           `xml:"Nm", json:"Nm", bson:"Nm"`
	PstlAdr   *PostalAddress6  `xml:"PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
	Id        *Party10Choice   `xml:"Id", json:"Id", bson:"Id"`
	CtryOfRes string           `xml:"CtryOfRes", json:"CtryOfRes", bson:"CtryOfRes"`
	CtctDtls  *ContactDetails2 `xml:"CtctDtls", json:"CtctDtls", bson:"CtctDtls"`
}

// PersonIdentification5 ...
type PersonIdentification5 struct {
	XMLName         xml.Name                        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	DtAndPlcOfBirth *DateAndPlaceOfBirth            `xml:"DtAndPlcOfBirth", json:"DtAndPlcOfBirth", bson:"DtAndPlcOfBirth"`
	Othr            []*GenericPersonIdentification1 `xml:"Othr", json:"Othr", bson:"Othr"`
}

// PersonIdentificationSchemeName1Choice ...
type PersonIdentificationSchemeName1Choice struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Cd      string   `xml:"Cd", json:"Cd", bson:"Cd"`
	Prtry   string   `xml:"Prtry", json:"Prtry", bson:"Prtry"`
}

// PhoneNumber ...
type PhoneNumber string

// PostalAddress6 ...
type PostalAddress6 struct {
	XMLName     xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AdrTp       string   `xml:"AdrTp", json:"AdrTp", bson:"AdrTp"`
	Dept        string   `xml:"Dept", json:"Dept", bson:"Dept"`
	SubDept     string   `xml:"SubDept", json:"SubDept", bson:"SubDept"`
	StrtNm      string   `xml:"StrtNm", json:"StrtNm", bson:"StrtNm"`
	BldgNb      string   `xml:"BldgNb", json:"BldgNb", bson:"BldgNb"`
	PstCd       string   `xml:"PstCd", json:"PstCd", bson:"PstCd"`
	TwnNm       string   `xml:"TwnNm", json:"TwnNm", bson:"TwnNm"`
	CtrySubDvsn string   `xml:"CtrySubDvsn", json:"CtrySubDvsn", bson:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry", json:"Ctry", bson:"Ctry"`
	AdrLine     []string `xml:"AdrLine", json:"AdrLine", bson:"AdrLine"`
}

// SignatureEnvelope ...
type SignatureEnvelope struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
}

// UnicodeChartsCode ...
type UnicodeChartsCode string

// YesNoIndicator ...
type YesNoIndicator bool
