// Code generated automatically. DO NOT EDIT.

package cain_004_001_01

import (
	"encoding/xml"
)

// Document ...
type Document *Document

// AccountIdentification30Choice ...
type AccountIdentification30Choice struct {
	XMLName *xml.Name `xml:"urn1:AccountIdentification30Choice", json:"-,omitempty", bson:"-,omitempty"`
	Card    string    `xml:"urn1:Card", json:"Card", bson:"Card"`
	MSISDN  string    `xml:"urn1:MSISDN", json:"MSISDN", bson:"MSISDN"`
	EMail   string    `xml:"urn1:EMail", json:"EMail", bson:"EMail"`
	IBAN    string    `xml:"urn1:IBAN", json:"IBAN", bson:"IBAN"`
	BBAN    string    `xml:"urn1:BBAN", json:"BBAN", bson:"BBAN"`
	UPIC    string    `xml:"urn1:UPIC", json:"UPIC", bson:"UPIC"`
	Dmst    string    `xml:"urn1:Dmst", json:"Dmst", bson:"Dmst"`
	Othr    string    `xml:"urn1:Othr", json:"Othr", bson:"Othr"`
}

// AcquirerFinancialResponse ...
type AcquirerFinancialResponse struct {
	XMLName  *xml.Name                   `xml:"urn1:AcquirerFinancialResponse", json:"-,omitempty", bson:"-,omitempty"`
	Hdr      *Header17                   `xml:"urn1:Hdr", json:"Hdr", bson:"Hdr"`
	FinRspn  *AcquirerFinancialResponse1 `xml:"urn1:FinRspn", json:"FinRspn", bson:"FinRspn"`
	SctyTrlr *ContentInformationType15   `xml:"urn1:SctyTrlr", json:"SctyTrlr", bson:"SctyTrlr"`
}

// AcquirerFinancialResponse1 ...
type AcquirerFinancialResponse1 struct {
	XMLName *xml.Name                    `xml:"urn1:AcquirerFinancialResponse1", json:"-,omitempty", bson:"-,omitempty"`
	Envt    *CardTransactionEnvironment2 `xml:"urn1:Envt", json:"Envt", bson:"Envt"`
	Cntxt   *CardTransactionContext3     `xml:"urn1:Cntxt", json:"Cntxt", bson:"Cntxt"`
	Tx      *CardTransaction6            `xml:"urn1:Tx", json:"Tx", bson:"Tx"`
}

// Action4 ...
type Action4 struct {
	XMLName   *xml.Name       `xml:"urn1:Action4", json:"-,omitempty", bson:"-,omitempty"`
	ActnTp    string          `xml:"urn1:ActnTp", json:"ActnTp", bson:"ActnTp"`
	MsgToPres *ActionMessage2 `xml:"urn1:MsgToPres", json:"MsgToPres", bson:"MsgToPres"`
}

// ActionMessage2 ...
type ActionMessage2 struct {
	XMLName      *xml.Name `xml:"urn1:ActionMessage2", json:"-,omitempty", bson:"-,omitempty"`
	MsgDstn      string    `xml:"urn1:MsgDstn", json:"MsgDstn", bson:"MsgDstn"`
	Frmt         string    `xml:"urn1:Frmt", json:"Frmt", bson:"Frmt"`
	MsgCntt      string    `xml:"urn1:MsgCntt", json:"MsgCntt", bson:"MsgCntt"`
	MsgCnttSgntr []byte    `xml:"urn1:MsgCnttSgntr", json:"MsgCnttSgntr", bson:"MsgCnttSgntr"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName *xml.Name `xml:"urn1:ActionMessage3", json:"-,omitempty", bson:"-,omitempty"`
	Dstn    string    `xml:"urn1:Dstn", json:"Dstn", bson:"Dstn"`
	Frmt    string    `xml:"urn1:Frmt", json:"Frmt", bson:"Frmt"`
	Cntt    string    `xml:"urn1:Cntt", json:"Cntt", bson:"Cntt"`
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
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification11", json:"-,omitempty", bson:"-,omitempty"`
	Algo    string      `xml:"urn1:Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter4 `xml:"urn1:Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification12 ...
type AlgorithmIdentification12 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification12", json:"-,omitempty", bson:"-,omitempty"`
	Algo    string      `xml:"urn1:Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter5 `xml:"urn1:Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification13 ...
type AlgorithmIdentification13 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification13", json:"-,omitempty", bson:"-,omitempty"`
	Algo    string      `xml:"urn1:Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter6 `xml:"urn1:Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification14 ...
type AlgorithmIdentification14 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification14", json:"-,omitempty", bson:"-,omitempty"`
	Algo    string      `xml:"urn1:Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter6 `xml:"urn1:Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification15 ...
type AlgorithmIdentification15 struct {
	XMLName *xml.Name   `xml:"urn1:AlgorithmIdentification15", json:"-,omitempty", bson:"-,omitempty"`
	Algo    string      `xml:"urn1:Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter7 `xml:"urn1:Param", json:"Param", bson:"Param"`
}

// AmountAndDirection41 ...
type AmountAndDirection41 struct {
	XMLName *xml.Name          `xml:"urn1:AmountAndDirection41", json:"-,omitempty", bson:"-,omitempty"`
	Amt     *CurrencyAndAmount `xml:"urn1:Amt", json:"Amt", bson:"Amt"`
	Sgn     bool               `xml:"urn1:Sgn", json:"Sgn", bson:"Sgn"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttributeType1Code ...
type AttributeType1Code string

// AuthenticatedData4 ...
type AuthenticatedData4 struct {
	XMLName     *xml.Name                  `xml:"urn1:AuthenticatedData4", json:"-,omitempty", bson:"-,omitempty"`
	Vrsn        float64                    `xml:"urn1:Vrsn", json:"Vrsn", bson:"Vrsn"`
	Rcpt        []*Recipient4Choice        `xml:"urn1:Rcpt", json:"Rcpt", bson:"Rcpt"`
	MACAlgo     *AlgorithmIdentification15 `xml:"urn1:MACAlgo", json:"MACAlgo", bson:"MACAlgo"`
	NcpsltdCntt *EncapsulatedContent3      `xml:"urn1:NcpsltdCntt", json:"NcpsltdCntt", bson:"NcpsltdCntt"`
	MAC         []byte                     `xml:"urn1:MAC", json:"MAC", bson:"MAC"`
}

// AuthenticationEntity2Code ...
type AuthenticationEntity2Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult8 ...
type AuthorisationResult8 struct {
	XMLName     *xml.Name                `xml:"urn1:AuthorisationResult8", json:"-,omitempty", bson:"-,omitempty"`
	AuthstnNtty *GenericIdentification75 `xml:"urn1:AuthstnNtty", json:"AuthstnNtty", bson:"AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"urn1:TxRspn", json:"TxRspn", bson:"TxRspn"`
	Actn        []*Action4               `xml:"urn1:Actn", json:"Actn", bson:"Actn"`
	AuthstnCd   string                   `xml:"urn1:AuthstnCd", json:"AuthstnCd", bson:"AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"urn1:AddtlInf", json:"AddtlInf", bson:"AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAccount2 ...
type CardAccount2 struct {
	XMLName      *xml.Name                      `xml:"urn1:CardAccount2", json:"-,omitempty", bson:"-,omitempty"`
	SelctdAcctTp string                         `xml:"urn1:SelctdAcctTp", json:"SelctdAcctTp", bson:"SelctdAcctTp"`
	AcctNm       string                         `xml:"urn1:AcctNm", json:"AcctNm", bson:"AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"urn1:AcctOwnr", json:"AcctOwnr", bson:"AcctOwnr"`
	Ccy          string                         `xml:"urn1:Ccy", json:"Ccy", bson:"Ccy"`
	AcctIdr      *AccountIdentification30Choice `xml:"urn1:AcctIdr", json:"AcctIdr", bson:"AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"urn1:Svcr", json:"Svcr", bson:"Svcr"`
	Bal          *AmountAndDirection41          `xml:"urn1:Bal", json:"Bal", bson:"Bal"`
	BalDt        string                         `xml:"urn1:BalDt", json:"BalDt", bson:"BalDt"`
}

// CardAccountType2Code ...
type CardAccountType2Code string

// CardPaymentServiceType7Code ...
type CardPaymentServiceType7Code string

// CardPaymentToken2 ...
type CardPaymentToken2 struct {
	XMLName      *xml.Name `xml:"urn1:CardPaymentToken2", json:"-,omitempty", bson:"-,omitempty"`
	TknChrtc     []string  `xml:"urn1:TknChrtc", json:"TknChrtc", bson:"TknChrtc"`
	TknAssrncLvl float64   `xml:"urn1:TknAssrncLvl", json:"TknAssrncLvl", bson:"TknAssrncLvl"`
}

// CardProductType1Code ...
type CardProductType1Code string

// CardTransaction6 ...
type CardTransaction6 struct {
	XMLName           *xml.Name               `xml:"urn1:CardTransaction6", json:"-,omitempty", bson:"-,omitempty"`
	TxTp              string                  `xml:"urn1:TxTp", json:"TxTp", bson:"TxTp"`
	Rcncltn           *TransactionIdentifier2 `xml:"urn1:Rcncltn", json:"Rcncltn", bson:"Rcncltn"`
	AccptrTxDtTm      string                  `xml:"urn1:AccptrTxDtTm", json:"AccptrTxDtTm", bson:"AccptrTxDtTm"`
	InitrTxId         string                  `xml:"urn1:InitrTxId", json:"InitrTxId", bson:"InitrTxId"`
	TxLifeCyclId      string                  `xml:"urn1:TxLifeCyclId", json:"TxLifeCyclId", bson:"TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"urn1:TxLifeCyclSeqNb", json:"TxLifeCyclSeqNb", bson:"TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"urn1:TxLifeCyclSeqCntr", json:"TxLifeCyclSeqCntr", bson:"TxLifeCyclSeqCntr"`
	CardIssrRefData   string                  `xml:"urn1:CardIssrRefData", json:"CardIssrRefData", bson:"CardIssrRefData"`
	TxDtls            *CardTransactionDetail4 `xml:"urn1:TxDtls", json:"TxDtls", bson:"TxDtls"`
	AuthstnRslt       *AuthorisationResult8   `xml:"urn1:AuthstnRslt", json:"AuthstnRslt", bson:"AuthstnRslt"`
}

// CardTransactionAmount4 ...
type CardTransactionAmount4 struct {
	XMLName          *xml.Name          `xml:"urn1:CardTransactionAmount4", json:"-,omitempty", bson:"-,omitempty"`
	TtlAmt           *CurrencyAndAmount `xml:"urn1:TtlAmt", json:"TtlAmt", bson:"TtlAmt"`
	CrdhldrBllgTxAmt *DetailedAmount8   `xml:"urn1:CrdhldrBllgTxAmt", json:"CrdhldrBllgTxAmt", bson:"CrdhldrBllgTxAmt"`
	RcncltnTxAmt     *DetailedAmount8   `xml:"urn1:RcncltnTxAmt", json:"RcncltnTxAmt", bson:"RcncltnTxAmt"`
	DtldAmt          []*DetailedAmount9 `xml:"urn1:DtldAmt", json:"DtldAmt", bson:"DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	XMLName *xml.Name `xml:"urn1:CardTransactionCondition1", json:"-,omitempty", bson:"-,omitempty"`
	Prgm    string    `xml:"urn1:Prgm", json:"Prgm", bson:"Prgm"`
	Val     string    `xml:"urn1:Val", json:"Val", bson:"Val"`
}

// CardTransactionContext3 ...
type CardTransactionContext3 struct {
	XMLName *xml.Name                `xml:"urn1:CardTransactionContext3", json:"-,omitempty", bson:"-,omitempty"`
	TxCntxt *CardTransactionContext4 `xml:"urn1:TxCntxt", json:"TxCntxt", bson:"TxCntxt"`
}

// CardTransactionContext4 ...
type CardTransactionContext4 struct {
	XMLName   *xml.Name                    `xml:"urn1:CardTransactionContext4", json:"-,omitempty", bson:"-,omitempty"`
	SpclConds []*CardTransactionCondition1 `xml:"urn1:SpclConds", json:"SpclConds", bson:"SpclConds"`
}

// CardTransactionDetail4 ...
type CardTransactionDetail4 struct {
	XMLName      *xml.Name                         `xml:"urn1:CardTransactionDetail4", json:"-,omitempty", bson:"-,omitempty"`
	TxAmts       *CardTransactionAmount4           `xml:"urn1:TxAmts", json:"TxAmts", bson:"TxAmts"`
	TxFees       []*DetailedAmount11               `xml:"urn1:TxFees", json:"TxFees", bson:"TxFees"`
	AddtlAmts    []*DetailedAmount10               `xml:"urn1:AddtlAmts", json:"AddtlAmts", bson:"AddtlAmts"`
	AcctAndBal   []*CardAccount2                   `xml:"urn1:AcctAndBal", json:"AcctAndBal", bson:"AcctAndBal"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"urn1:TxVrfctnRslt", json:"TxVrfctnRslt", bson:"TxVrfctnRslt"`
	VldtyDt      string                            `xml:"urn1:VldtyDt", json:"VldtyDt", bson:"VldtyDt"`
	ICCRltdData  []byte                            `xml:"urn1:ICCRltdData", json:"ICCRltdData", bson:"ICCRltdData"`
}

// CardTransactionEnvironment2 ...
type CardTransactionEnvironment2 struct {
	XMLName     *xml.Name          `xml:"urn1:CardTransactionEnvironment2", json:"-,omitempty", bson:"-,omitempty"`
	AcqrrId     string             `xml:"urn1:AcqrrId", json:"AcqrrId", bson:"AcqrrId"`
	CardSchmeId string             `xml:"urn1:CardSchmeId", json:"CardSchmeId", bson:"CardSchmeId"`
	AccptrId    string             `xml:"urn1:AccptrId", json:"AccptrId", bson:"AccptrId"`
	TermnlId    string             `xml:"urn1:TermnlId", json:"TermnlId", bson:"TermnlId"`
	Card        *PaymentCard13     `xml:"urn1:Card", json:"Card", bson:"Card"`
	PmtTkn      *CardPaymentToken2 `xml:"urn1:PmtTkn", json:"PmtTkn", bson:"PmtTkn"`
	ShppgAdr    *PostalAddress18   `xml:"urn1:ShppgAdr", json:"ShppgAdr", bson:"ShppgAdr"`
}

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        *xml.Name                     `xml:"urn1:CertificateIssuer1", json:"-,omitempty", bson:"-,omitempty"`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"urn1:RltvDstngshdNm", json:"RltvDstngshdNm", bson:"RltvDstngshdNm"`
}

// ContentInformationType10 ...
type ContentInformationType10 struct {
	XMLName    *xml.Name       `xml:"urn1:ContentInformationType10", json:"-,omitempty", bson:"-,omitempty"`
	CnttTp     string          `xml:"urn1:CnttTp", json:"CnttTp", bson:"CnttTp"`
	EnvlpdData *EnvelopedData4 `xml:"urn1:EnvlpdData", json:"EnvlpdData", bson:"EnvlpdData"`
}

// ContentInformationType15 ...
type ContentInformationType15 struct {
	XMLName      *xml.Name           `xml:"urn1:ContentInformationType15", json:"-,omitempty", bson:"-,omitempty"`
	CnttTp       string              `xml:"urn1:CnttTp", json:"CnttTp", bson:"CnttTp"`
	AuthntcdData *AuthenticatedData4 `xml:"urn1:AuthntcdData", json:"AuthntcdData", bson:"AuthntcdData"`
}

// ContentType2Code ...
type ContentType2Code string

// CountryCode ...
type CountryCode string

// CurrencyAndAmountSimpleType ...
type CurrencyAndAmountSimpleType float64

// CurrencyAndAmount ...
type CurrencyAndAmount struct {
	XMLName *xml.Name `xml:"urn1:CurrencyAndAmount", json:"-,omitempty", bson:"-,omitempty"`
	CcyAttr string    `json:"-,omitempty", bson:"-,omitempty", xml:"Ccy,attr"`
	Value   float64   `xml:",chardata", json:"float64", bson:"float64"`
}

// CurrencyCode ...
type CurrencyCode string

// DetailedAmount10 ...
type DetailedAmount10 struct {
	XMLName *xml.Name          `xml:"urn1:DetailedAmount10", json:"-,omitempty", bson:"-,omitempty"`
	Tp      string             `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	AddtlTp string             `xml:"urn1:AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt     *CurrencyAndAmount `xml:"urn1:Amt", json:"Amt", bson:"Amt"`
	Labl    string             `xml:"urn1:Labl", json:"Labl", bson:"Labl"`
}

// DetailedAmount11 ...
type DetailedAmount11 struct {
	XMLName  *xml.Name             `xml:"urn1:DetailedAmount11", json:"-,omitempty", bson:"-,omitempty"`
	Tp       string                `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	AddtlTp  string                `xml:"urn1:AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt      *AmountAndDirection41 `xml:"urn1:Amt", json:"Amt", bson:"Amt"`
	OrgnlAmt *AmountAndDirection41 `xml:"urn1:OrgnlAmt", json:"OrgnlAmt", bson:"OrgnlAmt"`
}

// DetailedAmount8 ...
type DetailedAmount8 struct {
	XMLName  *xml.Name `xml:"urn1:DetailedAmount8", json:"-,omitempty", bson:"-,omitempty"`
	Amt      float64   `xml:"urn1:Amt", json:"Amt", bson:"Amt"`
	XchgRate float64   `xml:"urn1:XchgRate", json:"XchgRate", bson:"XchgRate"`
	QtnDt    string    `xml:"urn1:QtnDt", json:"QtnDt", bson:"QtnDt"`
	Labl     string    `xml:"urn1:Labl", json:"Labl", bson:"Labl"`
}

// DetailedAmount9 ...
type DetailedAmount9 struct {
	XMLName *xml.Name `xml:"urn1:DetailedAmount9", json:"-,omitempty", bson:"-,omitempty"`
	Tp      string    `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	AddtlTp string    `xml:"urn1:AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt     float64   `xml:"urn1:Amt", json:"Amt", bson:"Amt"`
	Labl    string    `xml:"urn1:Labl", json:"Labl", bson:"Labl"`
}

// EncapsulatedContent3 ...
type EncapsulatedContent3 struct {
	XMLName *xml.Name `xml:"urn1:EncapsulatedContent3", json:"-,omitempty", bson:"-,omitempty"`
	CnttTp  string    `xml:"urn1:CnttTp", json:"CnttTp", bson:"CnttTp"`
	Cntt    []byte    `xml:"urn1:Cntt", json:"Cntt", bson:"Cntt"`
}

// EncryptedContent3 ...
type EncryptedContent3 struct {
	XMLName        *xml.Name                  `xml:"urn1:EncryptedContent3", json:"-,omitempty", bson:"-,omitempty"`
	CnttTp         string                     `xml:"urn1:CnttTp", json:"CnttTp", bson:"CnttTp"`
	CnttNcrptnAlgo *AlgorithmIdentification14 `xml:"urn1:CnttNcrptnAlgo", json:"CnttNcrptnAlgo", bson:"CnttNcrptnAlgo"`
	NcrptdData     []byte                     `xml:"urn1:NcrptdData", json:"NcrptdData", bson:"NcrptdData"`
}

// EncryptionFormat1Code ...
type EncryptionFormat1Code string

// EnvelopedData4 ...
type EnvelopedData4 struct {
	XMLName    *xml.Name           `xml:"urn1:EnvelopedData4", json:"-,omitempty", bson:"-,omitempty"`
	Vrsn       float64             `xml:"urn1:Vrsn", json:"Vrsn", bson:"Vrsn"`
	Rcpt       []*Recipient4Choice `xml:"urn1:Rcpt", json:"Rcpt", bson:"Rcpt"`
	NcrptdCntt *EncryptedContent3  `xml:"urn1:NcrptdCntt", json:"NcrptdCntt", bson:"NcrptdCntt"`
}

// Exact1NumericText ...
type Exact1NumericText string

// Exact3NumericText ...
type Exact3NumericText string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification1", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	SchmeNm string    `xml:"urn1:SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
}

// GenericIdentification73 ...
type GenericIdentification73 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification73", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	Tp      string    `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
	Ctry    string    `xml:"urn1:Ctry", json:"Ctry", bson:"Ctry"`
	ShrtNm  string    `xml:"urn1:ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// GenericIdentification74 ...
type GenericIdentification74 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification74", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	Tp      string    `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
	Ctry    string    `xml:"urn1:Ctry", json:"Ctry", bson:"Ctry"`
	ShrtNm  string    `xml:"urn1:ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// GenericIdentification75 ...
type GenericIdentification75 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification75", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	Tp      string    `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
	Ctry    string    `xml:"urn1:Ctry", json:"Ctry", bson:"Ctry"`
	ShrtNm  string    `xml:"urn1:ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// Header17 ...
type Header17 struct {
	XMLName        *xml.Name                `xml:"urn1:Header17", json:"-,omitempty", bson:"-,omitempty"`
	MsgFctn        string                   `xml:"urn1:MsgFctn", json:"MsgFctn", bson:"MsgFctn"`
	PrtcolVrsn     string                   `xml:"urn1:PrtcolVrsn", json:"PrtcolVrsn", bson:"PrtcolVrsn"`
	XchgId         string                   `xml:"urn1:XchgId", json:"XchgId", bson:"XchgId"`
	ReTrnsmssnCntr string                   `xml:"urn1:ReTrnsmssnCntr", json:"ReTrnsmssnCntr", bson:"ReTrnsmssnCntr"`
	CreDtTm        string                   `xml:"urn1:CreDtTm", json:"CreDtTm", bson:"CreDtTm"`
	InitgPty       *GenericIdentification73 `xml:"urn1:InitgPty", json:"InitgPty", bson:"InitgPty"`
	RcptPty        *GenericIdentification73 `xml:"urn1:RcptPty", json:"RcptPty", bson:"RcptPty"`
	Tracblt        []*Traceability3         `xml:"urn1:Tracblt", json:"Tracblt", bson:"Tracblt"`
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
	XMLName *xml.Name           `xml:"urn1:IssuerAndSerialNumber1", json:"-,omitempty", bson:"-,omitempty"`
	Issr    *CertificateIssuer1 `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
	SrlNb   []byte              `xml:"urn1:SrlNb", json:"SrlNb", bson:"SrlNb"`
}

// KEK4 ...
type KEK4 struct {
	XMLName       *xml.Name                  `xml:"urn1:KEK4", json:"-,omitempty", bson:"-,omitempty"`
	Vrsn          float64                    `xml:"urn1:Vrsn", json:"Vrsn", bson:"Vrsn"`
	KEKId         *KEKIdentifier2            `xml:"urn1:KEKId", json:"KEKId", bson:"KEKId"`
	KeyNcrptnAlgo *AlgorithmIdentification13 `xml:"urn1:KeyNcrptnAlgo", json:"KeyNcrptnAlgo", bson:"KeyNcrptnAlgo"`
	NcrptdKey     []byte                     `xml:"urn1:NcrptdKey", json:"NcrptdKey", bson:"NcrptdKey"`
}

// KEKIdentifier2 ...
type KEKIdentifier2 struct {
	XMLName   *xml.Name `xml:"urn1:KEKIdentifier2", json:"-,omitempty", bson:"-,omitempty"`
	KeyId     string    `xml:"urn1:KeyId", json:"KeyId", bson:"KeyId"`
	KeyVrsn   string    `xml:"urn1:KeyVrsn", json:"KeyVrsn", bson:"KeyVrsn"`
	SeqNb     float64   `xml:"urn1:SeqNb", json:"SeqNb", bson:"SeqNb"`
	DerivtnId []byte    `xml:"urn1:DerivtnId", json:"DerivtnId", bson:"DerivtnId"`
}

// KeyTransport4 ...
type KeyTransport4 struct {
	XMLName       *xml.Name                  `xml:"urn1:KeyTransport4", json:"-,omitempty", bson:"-,omitempty"`
	Vrsn          float64                    `xml:"urn1:Vrsn", json:"Vrsn", bson:"Vrsn"`
	RcptId        *Recipient5Choice          `xml:"urn1:RcptId", json:"RcptId", bson:"RcptId"`
	KeyNcrptnAlgo *AlgorithmIdentification11 `xml:"urn1:KeyNcrptnAlgo", json:"KeyNcrptnAlgo", bson:"KeyNcrptnAlgo"`
	NcrptdKey     []byte                     `xml:"urn1:NcrptdKey", json:"NcrptdKey", bson:"NcrptdKey"`
}

// Max10000Binary ...
type Max10000Binary []byte

// Max100KBinary ...
type Max100KBinary []byte

// Max10Text ...
type Max10Text string

// Max140Binary ...
type Max140Binary []byte

// Max140Text ...
type Max140Text string

// Max16Text ...
type Max16Text string

// Max20000Text ...
type Max20000Text string

// Max30Text ...
type Max30Text string

// Max35Binary ...
type Max35Binary []byte

// Max35Text ...
type Max35Text string

// Max3NumericText ...
type Max3NumericText string

// Max5000Binary ...
type Max5000Binary []byte

// Max500Binary ...
type Max500Binary []byte

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
type Min5Max16Binary []byte

// Min6Max8Text ...
type Min6Max8Text string

// Min8Max28NumericText ...
type Min8Max28NumericText string

// NameAndAddress3 ...
type NameAndAddress3 struct {
	XMLName *xml.Name       `xml:"urn1:NameAndAddress3", json:"-,omitempty", bson:"-,omitempty"`
	Nm      string          `xml:"urn1:Nm", json:"Nm", bson:"Nm"`
	Adr     *PostalAddress1 `xml:"urn1:Adr", json:"Adr", bson:"Adr"`
}

// Number ...
type Number float64

// OutputFormat1Code ...
type OutputFormat1Code string

// Parameter4 ...
type Parameter4 struct {
	XMLName      *xml.Name                  `xml:"urn1:Parameter4", json:"-,omitempty", bson:"-,omitempty"`
	NcrptnFrmt   string                     `xml:"urn1:NcrptnFrmt", json:"NcrptnFrmt", bson:"NcrptnFrmt"`
	DgstAlgo     string                     `xml:"urn1:DgstAlgo", json:"DgstAlgo", bson:"DgstAlgo"`
	MskGnrtrAlgo *AlgorithmIdentification12 `xml:"urn1:MskGnrtrAlgo", json:"MskGnrtrAlgo", bson:"MskGnrtrAlgo"`
}

// Parameter5 ...
type Parameter5 struct {
	XMLName  *xml.Name `xml:"urn1:Parameter5", json:"-,omitempty", bson:"-,omitempty"`
	DgstAlgo string    `xml:"urn1:DgstAlgo", json:"DgstAlgo", bson:"DgstAlgo"`
}

// Parameter6 ...
type Parameter6 struct {
	XMLName      *xml.Name `xml:"urn1:Parameter6", json:"-,omitempty", bson:"-,omitempty"`
	NcrptnFrmt   string    `xml:"urn1:NcrptnFrmt", json:"NcrptnFrmt", bson:"NcrptnFrmt"`
	InitlstnVctr []byte    `xml:"urn1:InitlstnVctr", json:"InitlstnVctr", bson:"InitlstnVctr"`
	BPddg        string    `xml:"urn1:BPddg", json:"BPddg", bson:"BPddg"`
}

// Parameter7 ...
type Parameter7 struct {
	XMLName      *xml.Name `xml:"urn1:Parameter7", json:"-,omitempty", bson:"-,omitempty"`
	InitlstnVctr []byte    `xml:"urn1:InitlstnVctr", json:"InitlstnVctr", bson:"InitlstnVctr"`
	BPddg        string    `xml:"urn1:BPddg", json:"BPddg", bson:"BPddg"`
}

// PartyIdentification72Choice ...
type PartyIdentification72Choice struct {
	XMLName *xml.Name               `xml:"urn1:PartyIdentification72Choice", json:"-,omitempty", bson:"-,omitempty"`
	AnyBIC  string                  `xml:"urn1:AnyBIC", json:"AnyBIC", bson:"AnyBIC"`
	PrtryId *GenericIdentification1 `xml:"urn1:PrtryId", json:"PrtryId", bson:"PrtryId"`
}

// PartyType10Code ...
type PartyType10Code string

// PartyType11Code ...
type PartyType11Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard13 ...
type PaymentCard13 struct {
	XMLName        *xml.Name                 `xml:"urn1:PaymentCard13", json:"-,omitempty", bson:"-,omitempty"`
	PrtctdCardData *ContentInformationType10 `xml:"urn1:PrtctdCardData", json:"PrtctdCardData", bson:"PrtctdCardData"`
	PlainCardData  *PlainCardData9           `xml:"urn1:PlainCardData", json:"PlainCardData", bson:"PlainCardData"`
	MskdPAN        string                    `xml:"urn1:MskdPAN", json:"MskdPAN", bson:"MskdPAN"`
	CardPdctTp     string                    `xml:"urn1:CardPdctTp", json:"CardPdctTp", bson:"CardPdctTp"`
	CardPdctNm     string                    `xml:"urn1:CardPdctNm", json:"CardPdctNm", bson:"CardPdctNm"`
}

// PlainCardData9 ...
type PlainCardData9 struct {
	XMLName   *xml.Name     `xml:"urn1:PlainCardData9", json:"-,omitempty", bson:"-,omitempty"`
	PAN       string        `xml:"urn1:PAN", json:"PAN", bson:"PAN"`
	CardSeqNb string        `xml:"urn1:CardSeqNb", json:"CardSeqNb", bson:"CardSeqNb"`
	FctvDt    string        `xml:"urn1:FctvDt", json:"FctvDt", bson:"FctvDt"`
	XpryDt    string        `xml:"urn1:XpryDt", json:"XpryDt", bson:"XpryDt"`
	SvcCd     string        `xml:"urn1:SvcCd", json:"SvcCd", bson:"SvcCd"`
	TrckData  []*TrackData1 `xml:"urn1:TrckData", json:"TrckData", bson:"TrckData"`
}

// PlusOrMinusIndicator ...
type PlusOrMinusIndicator bool

// PostalAddress1 ...
type PostalAddress1 struct {
	XMLName     *xml.Name `xml:"urn1:PostalAddress1", json:"-,omitempty", bson:"-,omitempty"`
	AdrTp       string    `xml:"urn1:AdrTp", json:"AdrTp", bson:"AdrTp"`
	AdrLine     []string  `xml:"urn1:AdrLine", json:"AdrLine", bson:"AdrLine"`
	StrtNm      string    `xml:"urn1:StrtNm", json:"StrtNm", bson:"StrtNm"`
	BldgNb      string    `xml:"urn1:BldgNb", json:"BldgNb", bson:"BldgNb"`
	PstCd       string    `xml:"urn1:PstCd", json:"PstCd", bson:"PstCd"`
	TwnNm       string    `xml:"urn1:TwnNm", json:"TwnNm", bson:"TwnNm"`
	CtrySubDvsn string    `xml:"urn1:CtrySubDvsn", json:"CtrySubDvsn", bson:"CtrySubDvsn"`
	Ctry        string    `xml:"urn1:Ctry", json:"Ctry", bson:"Ctry"`
}

// PostalAddress18 ...
type PostalAddress18 struct {
	XMLName     *xml.Name `xml:"urn1:PostalAddress18", json:"-,omitempty", bson:"-,omitempty"`
	AdrLine     []string  `xml:"urn1:AdrLine", json:"AdrLine", bson:"AdrLine"`
	StrtNm      string    `xml:"urn1:StrtNm", json:"StrtNm", bson:"StrtNm"`
	BldgNb      string    `xml:"urn1:BldgNb", json:"BldgNb", bson:"BldgNb"`
	PstCd       string    `xml:"urn1:PstCd", json:"PstCd", bson:"PstCd"`
	TwnNm       string    `xml:"urn1:TwnNm", json:"TwnNm", bson:"TwnNm"`
	CtrySubDvsn []string  `xml:"urn1:CtrySubDvsn", json:"CtrySubDvsn", bson:"CtrySubDvsn"`
	Ctry        string    `xml:"urn1:Ctry", json:"Ctry", bson:"Ctry"`
}

// Recipient4Choice ...
type Recipient4Choice struct {
	XMLName    *xml.Name       `xml:"urn1:Recipient4Choice", json:"-,omitempty", bson:"-,omitempty"`
	KeyTrnsprt *KeyTransport4  `xml:"urn1:KeyTrnsprt", json:"KeyTrnsprt", bson:"KeyTrnsprt"`
	KEK        *KEK4           `xml:"urn1:KEK", json:"KEK", bson:"KEK"`
	KeyIdr     *KEKIdentifier2 `xml:"urn1:KeyIdr", json:"KeyIdr", bson:"KeyIdr"`
}

// Recipient5Choice ...
type Recipient5Choice struct {
	XMLName      *xml.Name               `xml:"urn1:Recipient5Choice", json:"-,omitempty", bson:"-,omitempty"`
	IssrAndSrlNb *IssuerAndSerialNumber1 `xml:"urn1:IssrAndSrlNb", json:"IssrAndSrlNb", bson:"IssrAndSrlNb"`
	KeyIdr       *KEKIdentifier2         `xml:"urn1:KeyIdr", json:"KeyIdr", bson:"KeyIdr"`
}

// RelativeDistinguishedName1 ...
type RelativeDistinguishedName1 struct {
	XMLName *xml.Name `xml:"urn1:RelativeDistinguishedName1", json:"-,omitempty", bson:"-,omitempty"`
	AttrTp  string    `xml:"urn1:AttrTp", json:"AttrTp", bson:"AttrTp"`
	AttrVal string    `xml:"urn1:AttrVal", json:"AttrVal", bson:"AttrVal"`
}

// Response3Code ...
type Response3Code string

// ResponseType2 ...
type ResponseType2 struct {
	XMLName      *xml.Name `xml:"urn1:ResponseType2", json:"-,omitempty", bson:"-,omitempty"`
	Rslt         string    `xml:"urn1:Rslt", json:"Rslt", bson:"Rslt"`
	RsltDtls     string    `xml:"urn1:RsltDtls", json:"RsltDtls", bson:"RsltDtls"`
	AddtlRsltInf string    `xml:"urn1:AddtlRsltInf", json:"AddtlRsltInf", bson:"AddtlRsltInf"`
}

// ResultDetail1Code ...
type ResultDetail1Code string

// Traceability3 ...
type Traceability3 struct {
	XMLName     *xml.Name                `xml:"urn1:Traceability3", json:"-,omitempty", bson:"-,omitempty"`
	RlayId      *GenericIdentification74 `xml:"urn1:RlayId", json:"RlayId", bson:"RlayId"`
	TracDtTmIn  string                   `xml:"urn1:TracDtTmIn", json:"TracDtTmIn", bson:"TracDtTmIn"`
	TracDtTmOut string                   `xml:"urn1:TracDtTmOut", json:"TracDtTmOut", bson:"TracDtTmOut"`
}

// TrackData1 ...
type TrackData1 struct {
	XMLName *xml.Name `xml:"urn1:TrackData1", json:"-,omitempty", bson:"-,omitempty"`
	TrckNb  string    `xml:"urn1:TrckNb", json:"TrckNb", bson:"TrckNb"`
	TrckVal string    `xml:"urn1:TrckVal", json:"TrckVal", bson:"TrckVal"`
}

// TransactionIdentifier2 ...
type TransactionIdentifier2 struct {
	XMLName   *xml.Name `xml:"urn1:TransactionIdentifier2", json:"-,omitempty", bson:"-,omitempty"`
	RcncltnDt string    `xml:"urn1:RcncltnDt", json:"RcncltnDt", bson:"RcncltnDt"`
	RcncltnId string    `xml:"urn1:RcncltnId", json:"RcncltnId", bson:"RcncltnId"`
}

// TransactionVerificationResult4 ...
type TransactionVerificationResult4 struct {
	XMLName    *xml.Name `xml:"urn1:TransactionVerificationResult4", json:"-,omitempty", bson:"-,omitempty"`
	Mtd        string    `xml:"urn1:Mtd", json:"Mtd", bson:"Mtd"`
	VrfctnNtty string    `xml:"urn1:VrfctnNtty", json:"VrfctnNtty", bson:"VrfctnNtty"`
	Rslt       string    `xml:"urn1:Rslt", json:"Rslt", bson:"Rslt"`
	AddtlRslt  string    `xml:"urn1:AddtlRslt", json:"AddtlRslt", bson:"AddtlRslt"`
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
