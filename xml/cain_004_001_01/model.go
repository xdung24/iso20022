// Code generated by xgen. DO NOT EDIT.

package cain_004_001_01

import (
	"encoding/xml"
)

// Document ...
type Document *Document

// AccountIdentification30Choice ...
type AccountIdentification30Choice struct {
	XMLName *xml.Name `xml:"urn1:AccountIdentification30Choice"`
	Card    string    `xml:"urn1:Card"`
	MSISDN  string    `xml:"urn1:MSISDN"`
	EMail   string    `xml:"urn1:EMail"`
	IBAN    string    `xml:"urn1:IBAN"`
	BBAN    string    `xml:"urn1:BBAN"`
	UPIC    string    `xml:"urn1:UPIC"`
	Dmst    string    `xml:"urn1:Dmst"`
	Othr    string    `xml:"urn1:Othr"`
}

// AcquirerFinancialResponse ...
type AcquirerFinancialResponse struct {
	XMLName  *xml.Name                   `xml:"urn1:AcquirerFinancialResponse"`
	Hdr      *Header17                   `xml:"urn1:Hdr"`
	FinRspn  *AcquirerFinancialResponse1 `xml:"urn1:FinRspn"`
	SctyTrlr *ContentInformationType15   `xml:"urn1:SctyTrlr"`
}

// AcquirerFinancialResponse1 ...
type AcquirerFinancialResponse1 struct {
	XMLName *xml.Name                    `xml:"urn1:AcquirerFinancialResponse1"`
	Envt    *CardTransactionEnvironment2 `xml:"urn1:Envt"`
	Cntxt   *CardTransactionContext3     `xml:"urn1:Cntxt"`
	Tx      *CardTransaction6            `xml:"urn1:Tx"`
}

// Action4 ...
type Action4 struct {
	XMLName   *xml.Name       `xml:"urn1:Action4"`
	ActnTp    string          `xml:"urn1:ActnTp"`
	MsgToPres *ActionMessage2 `xml:"urn1:MsgToPres"`
}

// ActionMessage2 ...
type ActionMessage2 struct {
	XMLName      *xml.Name `xml:"urn1:ActionMessage2"`
	MsgDstn      string    `xml:"urn1:MsgDstn"`
	Frmt         string    `xml:"urn1:Frmt"`
	MsgCntt      string    `xml:"urn1:MsgCntt"`
	MsgCnttSgntr string    `xml:"urn1:MsgCnttSgntr"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName *xml.Name `xml:"urn1:ActionMessage3"`
	Dstn    string    `xml:"urn1:Dstn"`
	Frmt    string    `xml:"urn1:Frmt"`
	Cntt    string    `xml:"urn1:Cntt"`
}

// ActionType5Code ...
type ActionType5Code string

// ActiveCurrencyCode ...
type ActiveCurrencyCode string

// AddressType2Code ...
type AddressType2Code string

// Algorithm11Code ...
type Algorithm11Code string

// Algorithm12Code ...
type Algorithm12Code string

// Algorithm13Code ...
type Algorithm13Code string

// Algorithm15Code ...
type Algorithm15Code string

// Algorithm7Code ...
type Algorithm7Code string

// Algorithm8Code ...
type Algorithm8Code string

// AlgorithmIdentification11 ...
type AlgorithmIdentification11 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification11"`
	Algo    string      `xml:"urn1:Algo"`
	Param   *Parameter4 `xml:"urn1:Param"`
}

// AlgorithmIdentification12 ...
type AlgorithmIdentification12 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification12"`
	Algo    string      `xml:"urn1:Algo"`
	Param   *Parameter5 `xml:"urn1:Param"`
}

// AlgorithmIdentification13 ...
type AlgorithmIdentification13 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification13"`
	Algo    string      `xml:"urn1:Algo"`
	Param   *Parameter6 `xml:"urn1:Param"`
}

// AlgorithmIdentification14 ...
type AlgorithmIdentification14 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification14"`
	Algo    string      `xml:"urn1:Algo"`
	Param   *Parameter6 `xml:"urn1:Param"`
}

// AlgorithmIdentification15 ...
type AlgorithmIdentification15 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification15"`
	Algo    string      `xml:"urn1:Algo"`
	Param   *Parameter7 `xml:"urn1:Param"`
}

// AmountAndDirection41 ...
type AmountAndDirection41 struct {
	XMLName *xml.Name          `xml:"urn1:AmountAndDirection41"`
	Amt     *CurrencyAndAmount `xml:"urn1:Amt"`
	Sgn     bool               `xml:"urn1:Sgn"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttributeType1Code ...
type AttributeType1Code string

// AuthenticatedData4 ...
type AuthenticatedData4 struct {
	XMLName     *xml.Name                  `xml:"urn1:AuthenticatedData4"`
	Vrsn        float64                    `xml:"urn1:Vrsn"`
	Rcpt        []*Recipient4Choice        `xml:"urn1:Rcpt"`
	MACAlgo     *AlgorithmIdentification15 `xml:"urn1:MACAlgo"`
	NcpsltdCntt *EncapsulatedContent3      `xml:"urn1:NcpsltdCntt"`
	MAC         string                     `xml:"urn1:MAC"`
}

// AuthenticationEntity2Code ...
type AuthenticationEntity2Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult8 ...
type AuthorisationResult8 struct {
	XMLName     *xml.Name                `xml:"urn1:AuthorisationResult8"`
	AuthstnNtty *GenericIdentification75 `xml:"urn1:AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"urn1:TxRspn"`
	Actn        []*Action4               `xml:"urn1:Actn"`
	AuthstnCd   string                   `xml:"urn1:AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"urn1:AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAccount2 ...
type CardAccount2 struct {
	XMLName      *xml.Name                      `xml:"urn1:CardAccount2"`
	SelctdAcctTp string                         `xml:"urn1:SelctdAcctTp"`
	AcctNm       string                         `xml:"urn1:AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"urn1:AcctOwnr"`
	Ccy          string                         `xml:"urn1:Ccy"`
	AcctIdr      *AccountIdentification30Choice `xml:"urn1:AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"urn1:Svcr"`
	Bal          *AmountAndDirection41          `xml:"urn1:Bal"`
	BalDt        string                         `xml:"urn1:BalDt"`
}

// CardAccountType2Code ...
type CardAccountType2Code string

// CardPaymentServiceType7Code ...
type CardPaymentServiceType7Code string

// CardPaymentToken2 ...
type CardPaymentToken2 struct {
	XMLName      *xml.Name `xml:"urn1:CardPaymentToken2"`
	TknChrtc     []string  `xml:"urn1:TknChrtc"`
	TknAssrncLvl float64   `xml:"urn1:TknAssrncLvl"`
}

// CardProductType1Code ...
type CardProductType1Code string

// CardTransaction6 ...
type CardTransaction6 struct {
	XMLName           *xml.Name               `xml:"urn1:CardTransaction6"`
	TxTp              string                  `xml:"urn1:TxTp"`
	Rcncltn           *TransactionIdentifier2 `xml:"urn1:Rcncltn"`
	AccptrTxDtTm      string                  `xml:"urn1:AccptrTxDtTm"`
	InitrTxId         string                  `xml:"urn1:InitrTxId"`
	TxLifeCyclId      string                  `xml:"urn1:TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"urn1:TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"urn1:TxLifeCyclSeqCntr"`
	CardIssrRefData   string                  `xml:"urn1:CardIssrRefData"`
	TxDtls            *CardTransactionDetail4 `xml:"urn1:TxDtls"`
	AuthstnRslt       *AuthorisationResult8   `xml:"urn1:AuthstnRslt"`
}

// CardTransactionAmount4 ...
type CardTransactionAmount4 struct {
	XMLName          *xml.Name          `xml:"urn1:CardTransactionAmount4"`
	TtlAmt           *CurrencyAndAmount `xml:"urn1:TtlAmt"`
	CrdhldrBllgTxAmt *DetailedAmount8   `xml:"urn1:CrdhldrBllgTxAmt"`
	RcncltnTxAmt     *DetailedAmount8   `xml:"urn1:RcncltnTxAmt"`
	DtldAmt          []*DetailedAmount9 `xml:"urn1:DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	XMLName *xml.Name `xml:"urn1:CardTransactionCondition1"`
	Prgm    string    `xml:"urn1:Prgm"`
	Val     string    `xml:"urn1:Val"`
}

// CardTransactionContext3 ...
type CardTransactionContext3 struct {
	XMLName *xml.Name                `xml:"urn1:CardTransactionContext3"`
	TxCntxt *CardTransactionContext4 `xml:"urn1:TxCntxt"`
}

// CardTransactionContext4 ...
type CardTransactionContext4 struct {
	XMLName   *xml.Name                    `xml:"urn1:CardTransactionContext4"`
	SpclConds []*CardTransactionCondition1 `xml:"urn1:SpclConds"`
}

// CardTransactionDetail4 ...
type CardTransactionDetail4 struct {
	XMLName      *xml.Name                         `xml:"urn1:CardTransactionDetail4"`
	TxAmts       *CardTransactionAmount4           `xml:"urn1:TxAmts"`
	TxFees       []*DetailedAmount11               `xml:"urn1:TxFees"`
	AddtlAmts    []*DetailedAmount10               `xml:"urn1:AddtlAmts"`
	AcctAndBal   []*CardAccount2                   `xml:"urn1:AcctAndBal"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"urn1:TxVrfctnRslt"`
	VldtyDt      string                            `xml:"urn1:VldtyDt"`
	ICCRltdData  string                            `xml:"urn1:ICCRltdData"`
}

// CardTransactionEnvironment2 ...
type CardTransactionEnvironment2 struct {
	XMLName     *xml.Name          `xml:"urn1:CardTransactionEnvironment2"`
	AcqrrId     string             `xml:"urn1:AcqrrId"`
	CardSchmeId string             `xml:"urn1:CardSchmeId"`
	AccptrId    string             `xml:"urn1:AccptrId"`
	TermnlId    string             `xml:"urn1:TermnlId"`
	Card        *PaymentCard13     `xml:"urn1:Card"`
	PmtTkn      *CardPaymentToken2 `xml:"urn1:PmtTkn"`
	ShppgAdr    *PostalAddress18   `xml:"urn1:ShppgAdr"`
}

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        *xml.Name                     `xml:"urn1:CertificateIssuer1"`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"urn1:RltvDstngshdNm"`
}

// ContentInformationType10 ...
type ContentInformationType10 struct {
	XMLName    *xml.Name       `xml:"urn1:ContentInformationType10"`
	CnttTp     string          `xml:"urn1:CnttTp"`
	EnvlpdData *EnvelopedData4 `xml:"urn1:EnvlpdData"`
}

// ContentInformationType15 ...
type ContentInformationType15 struct {
	XMLName      *xml.Name           `xml:"urn1:ContentInformationType15"`
	CnttTp       string              `xml:"urn1:CnttTp"`
	AuthntcdData *AuthenticatedData4 `xml:"urn1:AuthntcdData"`
}

// ContentType2Code ...
type ContentType2Code string

// CountryCode ...
type CountryCode string

// CurrencyAndAmountSimpleType ...
type CurrencyAndAmountSimpleType float64

// CurrencyAndAmount ...
type CurrencyAndAmount struct {
	XMLName *xml.Name `xml:"urn1:CurrencyAndAmount"`
	CcyAttr string    `json:"-,omitempty", bson:"-,omitempty", xml:"Ccy,attr"`
	Value   float64   `xml:",chardata"`
}

// CurrencyCode ...
type CurrencyCode string

// DetailedAmount10 ...
type DetailedAmount10 struct {
	XMLName *xml.Name          `xml:"urn1:DetailedAmount10"`
	Tp      string             `xml:"urn1:Tp"`
	AddtlTp string             `xml:"urn1:AddtlTp"`
	Amt     *CurrencyAndAmount `xml:"urn1:Amt"`
	Labl    string             `xml:"urn1:Labl"`
}

// DetailedAmount11 ...
type DetailedAmount11 struct {
	XMLName  *xml.Name             `xml:"urn1:DetailedAmount11"`
	Tp       string                `xml:"urn1:Tp"`
	AddtlTp  string                `xml:"urn1:AddtlTp"`
	Amt      *AmountAndDirection41 `xml:"urn1:Amt"`
	OrgnlAmt *AmountAndDirection41 `xml:"urn1:OrgnlAmt"`
}

// DetailedAmount8 ...
type DetailedAmount8 struct {
	XMLName  *xml.Name `xml:"urn1:DetailedAmount8"`
	Amt      float64   `xml:"urn1:Amt"`
	XchgRate float64   `xml:"urn1:XchgRate"`
	QtnDt    string    `xml:"urn1:QtnDt"`
	Labl     string    `xml:"urn1:Labl"`
}

// DetailedAmount9 ...
type DetailedAmount9 struct {
	XMLName *xml.Name `xml:"urn1:DetailedAmount9"`
	Tp      string    `xml:"urn1:Tp"`
	AddtlTp string    `xml:"urn1:AddtlTp"`
	Amt     float64   `xml:"urn1:Amt"`
	Labl    string    `xml:"urn1:Labl"`
}

// EncapsulatedContent3 ...
type EncapsulatedContent3 struct {
	XMLName *xml.Name `xml:"urn1:EncapsulatedContent3"`
	CnttTp  string    `xml:"urn1:CnttTp"`
	Cntt    string    `xml:"urn1:Cntt"`
}

// EncryptedContent3 ...
type EncryptedContent3 struct {
	XMLName        *xml.Name                  `xml:"urn1:EncryptedContent3"`
	CnttTp         string                     `xml:"urn1:CnttTp"`
	CnttNcrptnAlgo *AlgorithmIdentification14 `xml:"urn1:CnttNcrptnAlgo"`
	NcrptdData     string                     `xml:"urn1:NcrptdData"`
}

// EncryptionFormat1Code ...
type EncryptionFormat1Code string

// EnvelopedData4 ...
type EnvelopedData4 struct {
	XMLName    *xml.Name           `xml:"urn1:EnvelopedData4"`
	Vrsn       float64             `xml:"urn1:Vrsn"`
	Rcpt       []*Recipient4Choice `xml:"urn1:Rcpt"`
	NcrptdCntt *EncryptedContent3  `xml:"urn1:NcrptdCntt"`
}

// Exact1NumericText ...
type Exact1NumericText string

// Exact3NumericText ...
type Exact3NumericText string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification1"`
	Id      string    `xml:"urn1:Id"`
	SchmeNm string    `xml:"urn1:SchmeNm"`
	Issr    string    `xml:"urn1:Issr"`
}

// GenericIdentification73 ...
type GenericIdentification73 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification73"`
	Id      string    `xml:"urn1:Id"`
	Tp      string    `xml:"urn1:Tp"`
	Issr    string    `xml:"urn1:Issr"`
	Ctry    string    `xml:"urn1:Ctry"`
	ShrtNm  string    `xml:"urn1:ShrtNm"`
}

// GenericIdentification74 ...
type GenericIdentification74 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification74"`
	Id      string    `xml:"urn1:Id"`
	Tp      string    `xml:"urn1:Tp"`
	Issr    string    `xml:"urn1:Issr"`
	Ctry    string    `xml:"urn1:Ctry"`
	ShrtNm  string    `xml:"urn1:ShrtNm"`
}

// GenericIdentification75 ...
type GenericIdentification75 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification75"`
	Id      string    `xml:"urn1:Id"`
	Tp      string    `xml:"urn1:Tp"`
	Issr    string    `xml:"urn1:Issr"`
	Ctry    string    `xml:"urn1:Ctry"`
	ShrtNm  string    `xml:"urn1:ShrtNm"`
}

// Header17 ...
type Header17 struct {
	XMLName        *xml.Name                `xml:"urn1:Header17"`
	MsgFctn        string                   `xml:"urn1:MsgFctn"`
	PrtcolVrsn     string                   `xml:"urn1:PrtcolVrsn"`
	XchgId         string                   `xml:"urn1:XchgId"`
	ReTrnsmssnCntr string                   `xml:"urn1:ReTrnsmssnCntr"`
	CreDtTm        string                   `xml:"urn1:CreDtTm"`
	InitgPty       *GenericIdentification73 `xml:"urn1:InitgPty"`
	RcptPty        *GenericIdentification73 `xml:"urn1:RcptPty"`
	Tracblt        []*Traceability3         `xml:"urn1:Tracblt"`
}

// IBANIdentifier ...
type IBANIdentifier string

// ISODate ...
type ISODate string

// ISODateTime ...
type ISODateTime string

// ImpliedCurrencyAndAmount ...
type ImpliedCurrencyAndAmount float64

// IssuerAndSerialNumber1 ...
type IssuerAndSerialNumber1 struct {
	XMLName *xml.Name           `xml:"urn1:IssuerAndSerialNumber1"`
	Issr    *CertificateIssuer1 `xml:"urn1:Issr"`
	SrlNb   string              `xml:"urn1:SrlNb"`
}

// KEK4 ...
type KEK4 struct {
	XMLName       *xml.Name                  `xml:"urn1:KEK4"`
	Vrsn          float64                    `xml:"urn1:Vrsn"`
	KEKId         *KEKIdentifier2            `xml:"urn1:KEKId"`
	KeyNcrptnAlgo *AlgorithmIdentification13 `xml:"urn1:KeyNcrptnAlgo"`
	NcrptdKey     string                     `xml:"urn1:NcrptdKey"`
}

// KEKIdentifier2 ...
type KEKIdentifier2 struct {
	XMLName   *xml.Name `xml:"urn1:KEKIdentifier2"`
	KeyId     string    `xml:"urn1:KeyId"`
	KeyVrsn   string    `xml:"urn1:KeyVrsn"`
	SeqNb     float64   `xml:"urn1:SeqNb"`
	DerivtnId string    `xml:"urn1:DerivtnId"`
}

// KeyTransport4 ...
type KeyTransport4 struct {
	XMLName       *xml.Name                  `xml:"urn1:KeyTransport4"`
	Vrsn          float64                    `xml:"urn1:Vrsn"`
	RcptId        *Recipient5Choice          `xml:"urn1:RcptId"`
	KeyNcrptnAlgo *AlgorithmIdentification11 `xml:"urn1:KeyNcrptnAlgo"`
	NcrptdKey     string                     `xml:"urn1:NcrptdKey"`
}

// Max10000Binary ...
type Max10000Binary string

// Max100KBinary ...
type Max100KBinary string

// Max10Text ...
type Max10Text string

// Max140Binary ...
type Max140Binary string

// Max140Text ...
type Max140Text string

// Max16Text ...
type Max16Text string

// Max20000Text ...
type Max20000Text string

// Max30Text ...
type Max30Text string

// Max35Binary ...
type Max35Binary string

// Max35Text ...
type Max35Text string

// Max3NumericText ...
type Max3NumericText string

// Max5000Binary ...
type Max5000Binary string

// Max500Binary ...
type Max500Binary string

// Max500Text ...
type Max500Text string

// Max6Text ...
type Max6Text string

// Max70Text ...
type Max70Text string

// MessageFunction6Code ...
type MessageFunction6Code string

// Min2Max3AlphaText ...
type Min2Max3AlphaText string

// Min2Max3NumericText ...
type Min2Max3NumericText string

// Min5Max16Binary ...
type Min5Max16Binary string

// Min6Max8Text ...
type Min6Max8Text string

// Min8Max28NumericText ...
type Min8Max28NumericText string

// NameAndAddress3 ...
type NameAndAddress3 struct {
	XMLName *xml.Name       `xml:"urn1:NameAndAddress3"`
	Nm      string          `xml:"urn1:Nm"`
	Adr     *PostalAddress1 `xml:"urn1:Adr"`
}

// Number ...
type Number float64

// OutputFormat1Code ...
type OutputFormat1Code string

// Parameter4 ...
type Parameter4 struct {
	XMLName      *xml.Name                  `xml:"urn1:Parameter4"`
	NcrptnFrmt   string                     `xml:"urn1:NcrptnFrmt"`
	DgstAlgo     string                     `xml:"urn1:DgstAlgo"`
	MskGnrtrAlgo *AlgorithmIdentification12 `xml:"urn1:MskGnrtrAlgo"`
}

// Parameter5 ...
type Parameter5 struct {
	XMLName  *xml.Name `xml:"urn1:Parameter5"`
	DgstAlgo string    `xml:"urn1:DgstAlgo"`
}

// Parameter6 ...
type Parameter6 struct {
	XMLName      *xml.Name `xml:"urn1:Parameter6"`
	NcrptnFrmt   string    `xml:"urn1:NcrptnFrmt"`
	InitlstnVctr string    `xml:"urn1:InitlstnVctr"`
	BPddg        string    `xml:"urn1:BPddg"`
}

// Parameter7 ...
type Parameter7 struct {
	XMLName      *xml.Name `xml:"urn1:Parameter7"`
	InitlstnVctr string    `xml:"urn1:InitlstnVctr"`
	BPddg        string    `xml:"urn1:BPddg"`
}

// PartyIdentification72Choice ...
type PartyIdentification72Choice struct {
	XMLName *xml.Name               `xml:"urn1:PartyIdentification72Choice"`
	AnyBIC  string                  `xml:"urn1:AnyBIC"`
	PrtryId *GenericIdentification1 `xml:"urn1:PrtryId"`
}

// PartyType10Code ...
type PartyType10Code string

// PartyType11Code ...
type PartyType11Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard13 ...
type PaymentCard13 struct {
	XMLName        *xml.Name                 `xml:"urn1:PaymentCard13"`
	PrtctdCardData *ContentInformationType10 `xml:"urn1:PrtctdCardData"`
	PlainCardData  *PlainCardData9           `xml:"urn1:PlainCardData"`
	MskdPAN        string                    `xml:"urn1:MskdPAN"`
	CardPdctTp     string                    `xml:"urn1:CardPdctTp"`
	CardPdctNm     string                    `xml:"urn1:CardPdctNm"`
}

// PlainCardData9 ...
type PlainCardData9 struct {
	XMLName   *xml.Name     `xml:"urn1:PlainCardData9"`
	PAN       string        `xml:"urn1:PAN"`
	CardSeqNb string        `xml:"urn1:CardSeqNb"`
	FctvDt    string        `xml:"urn1:FctvDt"`
	XpryDt    string        `xml:"urn1:XpryDt"`
	SvcCd     string        `xml:"urn1:SvcCd"`
	TrckData  []*TrackData1 `xml:"urn1:TrckData"`
}

// PlusOrMinusIndicator ...
type PlusOrMinusIndicator bool

// PostalAddress1 ...
type PostalAddress1 struct {
	XMLName     *xml.Name `xml:"urn1:PostalAddress1"`
	AdrTp       string    `xml:"urn1:AdrTp"`
	AdrLine     []string  `xml:"urn1:AdrLine"`
	StrtNm      string    `xml:"urn1:StrtNm"`
	BldgNb      string    `xml:"urn1:BldgNb"`
	PstCd       string    `xml:"urn1:PstCd"`
	TwnNm       string    `xml:"urn1:TwnNm"`
	CtrySubDvsn string    `xml:"urn1:CtrySubDvsn"`
	Ctry        string    `xml:"urn1:Ctry"`
}

// PostalAddress18 ...
type PostalAddress18 struct {
	XMLName     *xml.Name `xml:"urn1:PostalAddress18"`
	AdrLine     []string  `xml:"urn1:AdrLine"`
	StrtNm      string    `xml:"urn1:StrtNm"`
	BldgNb      string    `xml:"urn1:BldgNb"`
	PstCd       string    `xml:"urn1:PstCd"`
	TwnNm       string    `xml:"urn1:TwnNm"`
	CtrySubDvsn []string  `xml:"urn1:CtrySubDvsn"`
	Ctry        string    `xml:"urn1:Ctry"`
}

// Recipient4Choice ...
type Recipient4Choice struct {
	XMLName    *xml.Name       `xml:"urn1:Recipient4Choice"`
	KeyTrnsprt *KeyTransport4  `xml:"urn1:KeyTrnsprt"`
	KEK        *KEK4           `xml:"urn1:KEK"`
	KeyIdr     *KEKIdentifier2 `xml:"urn1:KeyIdr"`
}

// Recipient5Choice ...
type Recipient5Choice struct {
	XMLName      *xml.Name               `xml:"urn1:Recipient5Choice"`
	IssrAndSrlNb *IssuerAndSerialNumber1 `xml:"urn1:IssrAndSrlNb"`
	KeyIdr       *KEKIdentifier2         `xml:"urn1:KeyIdr"`
}

// RelativeDistinguishedName1 ...
type RelativeDistinguishedName1 struct {
	XMLName *xml.Name `xml:"urn1:RelativeDistinguishedName1"`
	AttrTp  string    `xml:"urn1:AttrTp"`
	AttrVal string    `xml:"urn1:AttrVal"`
}

// Response3Code ...
type Response3Code string

// ResponseType2 ...
type ResponseType2 struct {
	XMLName      *xml.Name `xml:"urn1:ResponseType2"`
	Rslt         string    `xml:"urn1:Rslt"`
	RsltDtls     string    `xml:"urn1:RsltDtls"`
	AddtlRsltInf string    `xml:"urn1:AddtlRsltInf"`
}

// ResultDetail1Code ...
type ResultDetail1Code string

// Traceability3 ...
type Traceability3 struct {
	XMLName     *xml.Name                `xml:"urn1:Traceability3"`
	RlayId      *GenericIdentification74 `xml:"urn1:RlayId"`
	TracDtTmIn  string                   `xml:"urn1:TracDtTmIn"`
	TracDtTmOut string                   `xml:"urn1:TracDtTmOut"`
}

// TrackData1 ...
type TrackData1 struct {
	XMLName *xml.Name `xml:"urn1:TrackData1"`
	TrckNb  string    `xml:"urn1:TrckNb"`
	TrckVal string    `xml:"urn1:TrckVal"`
}

// TransactionIdentifier2 ...
type TransactionIdentifier2 struct {
	XMLName   *xml.Name `xml:"urn1:TransactionIdentifier2"`
	RcncltnDt string    `xml:"urn1:RcncltnDt"`
	RcncltnId string    `xml:"urn1:RcncltnId"`
}

// TransactionVerificationResult4 ...
type TransactionVerificationResult4 struct {
	XMLName    *xml.Name `xml:"urn1:TransactionVerificationResult4"`
	Mtd        string    `xml:"urn1:Mtd"`
	VrfctnNtty string    `xml:"urn1:VrfctnNtty"`
	Rslt       string    `xml:"urn1:Rslt"`
	AddtlRslt  string    `xml:"urn1:AddtlRslt"`
}

// TypeOfAmount5Code ...
type TypeOfAmount5Code string

// TypeOfAmount6Code ...
type TypeOfAmount6Code string

// TypeOfAmount7Code ...
type TypeOfAmount7Code string

// UPICIdentifier ...
type UPICIdentifier string

// UserInterface3Code ...
type UserInterface3Code string

// UserInterface4Code ...
type UserInterface4Code string

// Verification1Code ...
type Verification1Code string
