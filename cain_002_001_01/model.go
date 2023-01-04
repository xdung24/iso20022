package cain_002_001_01

import (
	"encoding/xml"
)

// Document ...
type Document *Document

// AccountIdentification30Choice ...
type AccountIdentification30Choice struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Card    string   `xml:"Card", json:"Card", bson:"Card"`
	MSISDN  string   `xml:"MSISDN", json:"MSISDN", bson:"MSISDN"`
	EMail   string   `xml:"EMail", json:"EMail", bson:"EMail"`
	IBAN    string   `xml:"IBAN", json:"IBAN", bson:"IBAN"`
	BBAN    string   `xml:"BBAN", json:"BBAN", bson:"BBAN"`
	UPIC    string   `xml:"UPIC", json:"UPIC", bson:"UPIC"`
	Dmst    string   `xml:"Dmst", json:"Dmst", bson:"Dmst"`
	Othr    string   `xml:"Othr", json:"Othr", bson:"Othr"`
}

// AcquirerAuthorisationResponse ...
type AcquirerAuthorisationResponse struct {
	XMLName     xml.Name                        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Hdr         *Header17                       `xml:"Hdr", json:"Hdr", bson:"Hdr"`
	AuthstnRspn *AcquirerAuthorisationResponse1 `xml:"AuthstnRspn", json:"AuthstnRspn", bson:"AuthstnRspn"`
	SctyTrlr    *ContentInformationType15       `xml:"SctyTrlr", json:"SctyTrlr", bson:"SctyTrlr"`
}

// AcquirerAuthorisationResponse1 ...
type AcquirerAuthorisationResponse1 struct {
	XMLName xml.Name                     `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Envt    *CardTransactionEnvironment2 `xml:"Envt", json:"Envt", bson:"Envt"`
	Cntxt   *CardTransactionContext3     `xml:"Cntxt", json:"Cntxt", bson:"Cntxt"`
	Tx      *CardTransaction4            `xml:"Tx", json:"Tx", bson:"Tx"`
}

// Action4 ...
type Action4 struct {
	XMLName   xml.Name        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	ActnTp    string          `xml:"ActnTp", json:"ActnTp", bson:"ActnTp"`
	MsgToPres *ActionMessage2 `xml:"MsgToPres", json:"MsgToPres", bson:"MsgToPres"`
}

// ActionMessage2 ...
type ActionMessage2 struct {
	XMLName      xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	MsgDstn      string   `xml:"MsgDstn", json:"MsgDstn", bson:"MsgDstn"`
	Frmt         string   `xml:"Frmt", json:"Frmt", bson:"Frmt"`
	MsgCntt      string   `xml:"MsgCntt", json:"MsgCntt", bson:"MsgCntt"`
	MsgCnttSgntr []byte   `xml:"MsgCnttSgntr", json:"MsgCnttSgntr", bson:"MsgCnttSgntr"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Dstn    string   `xml:"Dstn", json:"Dstn", bson:"Dstn"`
	Frmt    string   `xml:"Frmt", json:"Frmt", bson:"Frmt"`
	Cntt    string   `xml:"Cntt", json:"Cntt", bson:"Cntt"`
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
	XMLName xml.Name    `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Algo    string      `xml:"Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter4 `xml:"Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification12 ...
type AlgorithmIdentification12 struct {
	XMLName xml.Name    `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Algo    string      `xml:"Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter5 `xml:"Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification13 ...
type AlgorithmIdentification13 struct {
	XMLName xml.Name    `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Algo    string      `xml:"Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter6 `xml:"Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification14 ...
type AlgorithmIdentification14 struct {
	XMLName xml.Name    `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Algo    string      `xml:"Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter6 `xml:"Param", json:"Param", bson:"Param"`
}

// AlgorithmIdentification15 ...
type AlgorithmIdentification15 struct {
	XMLName xml.Name    `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Algo    string      `xml:"Algo", json:"Algo", bson:"Algo"`
	Param   *Parameter7 `xml:"Param", json:"Param", bson:"Param"`
}

// AmountAndDirection41 ...
type AmountAndDirection41 struct {
	XMLName xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Amt     *CurrencyAndAmount `xml:"Amt", json:"Amt", bson:"Amt"`
	Sgn     bool               `xml:"Sgn", json:"Sgn", bson:"Sgn"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttributeType1Code ...
type AttributeType1Code string

// AuthenticatedData4 ...
type AuthenticatedData4 struct {
	XMLName     xml.Name                   `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Vrsn        float64                    `xml:"Vrsn", json:"Vrsn", bson:"Vrsn"`
	Rcpt        []*Recipient4Choice        `xml:"Rcpt", json:"Rcpt", bson:"Rcpt"`
	MACAlgo     *AlgorithmIdentification15 `xml:"MACAlgo", json:"MACAlgo", bson:"MACAlgo"`
	NcpsltdCntt *EncapsulatedContent3      `xml:"NcpsltdCntt", json:"NcpsltdCntt", bson:"NcpsltdCntt"`
	MAC         []byte                     `xml:"MAC", json:"MAC", bson:"MAC"`
}

// AuthenticationEntity2Code ...
type AuthenticationEntity2Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult8 ...
type AuthorisationResult8 struct {
	XMLName     xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AuthstnNtty *GenericIdentification75 `xml:"AuthstnNtty", json:"AuthstnNtty", bson:"AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"TxRspn", json:"TxRspn", bson:"TxRspn"`
	Actn        []*Action4               `xml:"Actn", json:"Actn", bson:"Actn"`
	AuthstnCd   string                   `xml:"AuthstnCd", json:"AuthstnCd", bson:"AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"AddtlInf", json:"AddtlInf", bson:"AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAccount2 ...
type CardAccount2 struct {
	XMLName      xml.Name                       `json:"-,omitempty", bson:"-,omitempty", xml:""`
	SelctdAcctTp string                         `xml:"SelctdAcctTp", json:"SelctdAcctTp", bson:"SelctdAcctTp"`
	AcctNm       string                         `xml:"AcctNm", json:"AcctNm", bson:"AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"AcctOwnr", json:"AcctOwnr", bson:"AcctOwnr"`
	Ccy          string                         `xml:"Ccy", json:"Ccy", bson:"Ccy"`
	AcctIdr      *AccountIdentification30Choice `xml:"AcctIdr", json:"AcctIdr", bson:"AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"Svcr", json:"Svcr", bson:"Svcr"`
	Bal          *AmountAndDirection41          `xml:"Bal", json:"Bal", bson:"Bal"`
	BalDt        string                         `xml:"BalDt", json:"BalDt", bson:"BalDt"`
}

// CardAccountType2Code ...
type CardAccountType2Code string

// CardPaymentServiceType7Code ...
type CardPaymentServiceType7Code string

// CardPaymentToken2 ...
type CardPaymentToken2 struct {
	XMLName      xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TknChrtc     []string `xml:"TknChrtc", json:"TknChrtc", bson:"TknChrtc"`
	TknAssrncLvl float64  `xml:"TknAssrncLvl", json:"TknAssrncLvl", bson:"TknAssrncLvl"`
}

// CardProductType1Code ...
type CardProductType1Code string

// CardTransaction4 ...
type CardTransaction4 struct {
	XMLName           xml.Name                `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TxTp              string                  `xml:"TxTp", json:"TxTp", bson:"TxTp"`
	Rcncltn           *TransactionIdentifier2 `xml:"Rcncltn", json:"Rcncltn", bson:"Rcncltn"`
	AccptrTxDtTm      string                  `xml:"AccptrTxDtTm", json:"AccptrTxDtTm", bson:"AccptrTxDtTm"`
	InitrTxId         string                  `xml:"InitrTxId", json:"InitrTxId", bson:"InitrTxId"`
	TxLifeCyclId      string                  `xml:"TxLifeCyclId", json:"TxLifeCyclId", bson:"TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"TxLifeCyclSeqNb", json:"TxLifeCyclSeqNb", bson:"TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"TxLifeCyclSeqCntr", json:"TxLifeCyclSeqCntr", bson:"TxLifeCyclSeqCntr"`
	CardIssrRefData   string                  `xml:"CardIssrRefData", json:"CardIssrRefData", bson:"CardIssrRefData"`
	TxDtls            *CardTransactionDetail2 `xml:"TxDtls", json:"TxDtls", bson:"TxDtls"`
	AuthstnRslt       *AuthorisationResult8   `xml:"AuthstnRslt", json:"AuthstnRslt", bson:"AuthstnRslt"`
}

// CardTransactionAmount2 ...
type CardTransactionAmount2 struct {
	XMLName          xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TtlAmt           *CurrencyAndAmount `xml:"TtlAmt", json:"TtlAmt", bson:"TtlAmt"`
	CrdhldrBllgTxAmt *DetailedAmount8   `xml:"CrdhldrBllgTxAmt", json:"CrdhldrBllgTxAmt", bson:"CrdhldrBllgTxAmt"`
	DtldAmt          []*DetailedAmount9 `xml:"DtldAmt", json:"DtldAmt", bson:"DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Prgm    string   `xml:"Prgm", json:"Prgm", bson:"Prgm"`
	Val     string   `xml:"Val", json:"Val", bson:"Val"`
}

// CardTransactionContext3 ...
type CardTransactionContext3 struct {
	XMLName xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TxCntxt *CardTransactionContext4 `xml:"TxCntxt", json:"TxCntxt", bson:"TxCntxt"`
}

// CardTransactionContext4 ...
type CardTransactionContext4 struct {
	XMLName   xml.Name                     `json:"-,omitempty", bson:"-,omitempty", xml:""`
	SpclConds []*CardTransactionCondition1 `xml:"SpclConds", json:"SpclConds", bson:"SpclConds"`
}

// CardTransactionDetail2 ...
type CardTransactionDetail2 struct {
	XMLName      xml.Name                          `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TxAmts       *CardTransactionAmount2           `xml:"TxAmts", json:"TxAmts", bson:"TxAmts"`
	AddtlAmts    []*DetailedAmount10               `xml:"AddtlAmts", json:"AddtlAmts", bson:"AddtlAmts"`
	AcctAndBal   []*CardAccount2                   `xml:"AcctAndBal", json:"AcctAndBal", bson:"AcctAndBal"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"TxVrfctnRslt", json:"TxVrfctnRslt", bson:"TxVrfctnRslt"`
	VldtyDt      string                            `xml:"VldtyDt", json:"VldtyDt", bson:"VldtyDt"`
	ICCRltdData  []byte                            `xml:"ICCRltdData", json:"ICCRltdData", bson:"ICCRltdData"`
}

// CardTransactionEnvironment2 ...
type CardTransactionEnvironment2 struct {
	XMLName     xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AcqrrId     string             `xml:"AcqrrId", json:"AcqrrId", bson:"AcqrrId"`
	CardSchmeId string             `xml:"CardSchmeId", json:"CardSchmeId", bson:"CardSchmeId"`
	AccptrId    string             `xml:"AccptrId", json:"AccptrId", bson:"AccptrId"`
	TermnlId    string             `xml:"TermnlId", json:"TermnlId", bson:"TermnlId"`
	Card        *PaymentCard13     `xml:"Card", json:"Card", bson:"Card"`
	PmtTkn      *CardPaymentToken2 `xml:"PmtTkn", json:"PmtTkn", bson:"PmtTkn"`
	ShppgAdr    *PostalAddress18   `xml:"ShppgAdr", json:"ShppgAdr", bson:"ShppgAdr"`
}

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        xml.Name                      `json:"-,omitempty", bson:"-,omitempty", xml:""`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"RltvDstngshdNm", json:"RltvDstngshdNm", bson:"RltvDstngshdNm"`
}

// ContentInformationType10 ...
type ContentInformationType10 struct {
	XMLName    xml.Name        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CnttTp     string          `xml:"CnttTp", json:"CnttTp", bson:"CnttTp"`
	EnvlpdData *EnvelopedData4 `xml:"EnvlpdData", json:"EnvlpdData", bson:"EnvlpdData"`
}

// ContentInformationType15 ...
type ContentInformationType15 struct {
	XMLName      xml.Name            `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CnttTp       string              `xml:"CnttTp", json:"CnttTp", bson:"CnttTp"`
	AuthntcdData *AuthenticatedData4 `xml:"AuthntcdData", json:"AuthntcdData", bson:"AuthntcdData"`
}

// ContentType2Code ...
type ContentType2Code string

// CountryCode ...
type CountryCode string

// CurrencyAndAmountSimpleType ...
type CurrencyAndAmountSimpleType float64

// CurrencyAndAmount ...
type CurrencyAndAmount struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CcyAttr string   `json:"-,omitempty", bson:"-,omitempty", xml:"Ccy,attr"`
	Value   float64  `xml:",chardata", json:"float64", bson:"float64"`
}

// CurrencyCode ...
type CurrencyCode string

// DetailedAmount10 ...
type DetailedAmount10 struct {
	XMLName xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Tp      string             `xml:"Tp", json:"Tp", bson:"Tp"`
	AddtlTp string             `xml:"AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt     *CurrencyAndAmount `xml:"Amt", json:"Amt", bson:"Amt"`
	Labl    string             `xml:"Labl", json:"Labl", bson:"Labl"`
}

// DetailedAmount8 ...
type DetailedAmount8 struct {
	XMLName  xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Amt      float64  `xml:"Amt", json:"Amt", bson:"Amt"`
	XchgRate float64  `xml:"XchgRate", json:"XchgRate", bson:"XchgRate"`
	QtnDt    string   `xml:"QtnDt", json:"QtnDt", bson:"QtnDt"`
	Labl     string   `xml:"Labl", json:"Labl", bson:"Labl"`
}

// DetailedAmount9 ...
type DetailedAmount9 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Tp      string   `xml:"Tp", json:"Tp", bson:"Tp"`
	AddtlTp string   `xml:"AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt     float64  `xml:"Amt", json:"Amt", bson:"Amt"`
	Labl    string   `xml:"Labl", json:"Labl", bson:"Labl"`
}

// EncapsulatedContent3 ...
type EncapsulatedContent3 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CnttTp  string   `xml:"CnttTp", json:"CnttTp", bson:"CnttTp"`
	Cntt    []byte   `xml:"Cntt", json:"Cntt", bson:"Cntt"`
}

// EncryptedContent3 ...
type EncryptedContent3 struct {
	XMLName        xml.Name                   `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CnttTp         string                     `xml:"CnttTp", json:"CnttTp", bson:"CnttTp"`
	CnttNcrptnAlgo *AlgorithmIdentification14 `xml:"CnttNcrptnAlgo", json:"CnttNcrptnAlgo", bson:"CnttNcrptnAlgo"`
	NcrptdData     []byte                     `xml:"NcrptdData", json:"NcrptdData", bson:"NcrptdData"`
}

// EncryptionFormat1Code ...
type EncryptionFormat1Code string

// EnvelopedData4 ...
type EnvelopedData4 struct {
	XMLName    xml.Name            `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Vrsn       float64             `xml:"Vrsn", json:"Vrsn", bson:"Vrsn"`
	Rcpt       []*Recipient4Choice `xml:"Rcpt", json:"Rcpt", bson:"Rcpt"`
	NcrptdCntt *EncryptedContent3  `xml:"NcrptdCntt", json:"NcrptdCntt", bson:"NcrptdCntt"`
}

// Exact1NumericText ...
type Exact1NumericText string

// Exact3NumericText ...
type Exact3NumericText string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	SchmeNm string   `xml:"SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
}

// GenericIdentification73 ...
type GenericIdentification73 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	Tp      string   `xml:"Tp", json:"Tp", bson:"Tp"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
	Ctry    string   `xml:"Ctry", json:"Ctry", bson:"Ctry"`
	ShrtNm  string   `xml:"ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// GenericIdentification74 ...
type GenericIdentification74 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	Tp      string   `xml:"Tp", json:"Tp", bson:"Tp"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
	Ctry    string   `xml:"Ctry", json:"Ctry", bson:"Ctry"`
	ShrtNm  string   `xml:"ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// GenericIdentification75 ...
type GenericIdentification75 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	Tp      string   `xml:"Tp", json:"Tp", bson:"Tp"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
	Ctry    string   `xml:"Ctry", json:"Ctry", bson:"Ctry"`
	ShrtNm  string   `xml:"ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// Header17 ...
type Header17 struct {
	XMLName        xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	MsgFctn        string                   `xml:"MsgFctn", json:"MsgFctn", bson:"MsgFctn"`
	PrtcolVrsn     string                   `xml:"PrtcolVrsn", json:"PrtcolVrsn", bson:"PrtcolVrsn"`
	XchgId         string                   `xml:"XchgId", json:"XchgId", bson:"XchgId"`
	ReTrnsmssnCntr string                   `xml:"ReTrnsmssnCntr", json:"ReTrnsmssnCntr", bson:"ReTrnsmssnCntr"`
	CreDtTm        string                   `xml:"CreDtTm", json:"CreDtTm", bson:"CreDtTm"`
	InitgPty       *GenericIdentification73 `xml:"InitgPty", json:"InitgPty", bson:"InitgPty"`
	RcptPty        *GenericIdentification73 `xml:"RcptPty", json:"RcptPty", bson:"RcptPty"`
	Tracblt        []*Traceability3         `xml:"Tracblt", json:"Tracblt", bson:"Tracblt"`
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
	XMLName xml.Name            `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Issr    *CertificateIssuer1 `xml:"Issr", json:"Issr", bson:"Issr"`
	SrlNb   []byte              `xml:"SrlNb", json:"SrlNb", bson:"SrlNb"`
}

// KEK4 ...
type KEK4 struct {
	XMLName       xml.Name                   `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Vrsn          float64                    `xml:"Vrsn", json:"Vrsn", bson:"Vrsn"`
	KEKId         *KEKIdentifier2            `xml:"KEKId", json:"KEKId", bson:"KEKId"`
	KeyNcrptnAlgo *AlgorithmIdentification13 `xml:"KeyNcrptnAlgo", json:"KeyNcrptnAlgo", bson:"KeyNcrptnAlgo"`
	NcrptdKey     []byte                     `xml:"NcrptdKey", json:"NcrptdKey", bson:"NcrptdKey"`
}

// KEKIdentifier2 ...
type KEKIdentifier2 struct {
	XMLName   xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	KeyId     string   `xml:"KeyId", json:"KeyId", bson:"KeyId"`
	KeyVrsn   string   `xml:"KeyVrsn", json:"KeyVrsn", bson:"KeyVrsn"`
	SeqNb     float64  `xml:"SeqNb", json:"SeqNb", bson:"SeqNb"`
	DerivtnId []byte   `xml:"DerivtnId", json:"DerivtnId", bson:"DerivtnId"`
}

// KeyTransport4 ...
type KeyTransport4 struct {
	XMLName       xml.Name                   `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Vrsn          float64                    `xml:"Vrsn", json:"Vrsn", bson:"Vrsn"`
	RcptId        *Recipient5Choice          `xml:"RcptId", json:"RcptId", bson:"RcptId"`
	KeyNcrptnAlgo *AlgorithmIdentification11 `xml:"KeyNcrptnAlgo", json:"KeyNcrptnAlgo", bson:"KeyNcrptnAlgo"`
	NcrptdKey     []byte                     `xml:"NcrptdKey", json:"NcrptdKey", bson:"NcrptdKey"`
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
	XMLName xml.Name        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Nm      string          `xml:"Nm", json:"Nm", bson:"Nm"`
	Adr     *PostalAddress1 `xml:"Adr", json:"Adr", bson:"Adr"`
}

// Number ...
type Number float64

// OutputFormat1Code ...
type OutputFormat1Code string

// Parameter4 ...
type Parameter4 struct {
	XMLName      xml.Name                   `json:"-,omitempty", bson:"-,omitempty", xml:""`
	NcrptnFrmt   string                     `xml:"NcrptnFrmt", json:"NcrptnFrmt", bson:"NcrptnFrmt"`
	DgstAlgo     string                     `xml:"DgstAlgo", json:"DgstAlgo", bson:"DgstAlgo"`
	MskGnrtrAlgo *AlgorithmIdentification12 `xml:"MskGnrtrAlgo", json:"MskGnrtrAlgo", bson:"MskGnrtrAlgo"`
}

// Parameter5 ...
type Parameter5 struct {
	XMLName  xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	DgstAlgo string   `xml:"DgstAlgo", json:"DgstAlgo", bson:"DgstAlgo"`
}

// Parameter6 ...
type Parameter6 struct {
	XMLName      xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	NcrptnFrmt   string   `xml:"NcrptnFrmt", json:"NcrptnFrmt", bson:"NcrptnFrmt"`
	InitlstnVctr []byte   `xml:"InitlstnVctr", json:"InitlstnVctr", bson:"InitlstnVctr"`
	BPddg        string   `xml:"BPddg", json:"BPddg", bson:"BPddg"`
}

// Parameter7 ...
type Parameter7 struct {
	XMLName      xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	InitlstnVctr []byte   `xml:"InitlstnVctr", json:"InitlstnVctr", bson:"InitlstnVctr"`
	BPddg        string   `xml:"BPddg", json:"BPddg", bson:"BPddg"`
}

// PartyIdentification72Choice ...
type PartyIdentification72Choice struct {
	XMLName xml.Name                `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AnyBIC  string                  `xml:"AnyBIC", json:"AnyBIC", bson:"AnyBIC"`
	PrtryId *GenericIdentification1 `xml:"PrtryId", json:"PrtryId", bson:"PrtryId"`
}

// PartyType10Code ...
type PartyType10Code string

// PartyType11Code ...
type PartyType11Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard13 ...
type PaymentCard13 struct {
	XMLName        xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	PrtctdCardData *ContentInformationType10 `xml:"PrtctdCardData", json:"PrtctdCardData", bson:"PrtctdCardData"`
	PlainCardData  *PlainCardData9           `xml:"PlainCardData", json:"PlainCardData", bson:"PlainCardData"`
	MskdPAN        string                    `xml:"MskdPAN", json:"MskdPAN", bson:"MskdPAN"`
	CardPdctTp     string                    `xml:"CardPdctTp", json:"CardPdctTp", bson:"CardPdctTp"`
	CardPdctNm     string                    `xml:"CardPdctNm", json:"CardPdctNm", bson:"CardPdctNm"`
}

// PlainCardData9 ...
type PlainCardData9 struct {
	XMLName   xml.Name      `json:"-,omitempty", bson:"-,omitempty", xml:""`
	PAN       string        `xml:"PAN", json:"PAN", bson:"PAN"`
	CardSeqNb string        `xml:"CardSeqNb", json:"CardSeqNb", bson:"CardSeqNb"`
	FctvDt    string        `xml:"FctvDt", json:"FctvDt", bson:"FctvDt"`
	XpryDt    string        `xml:"XpryDt", json:"XpryDt", bson:"XpryDt"`
	SvcCd     string        `xml:"SvcCd", json:"SvcCd", bson:"SvcCd"`
	TrckData  []*TrackData1 `xml:"TrckData", json:"TrckData", bson:"TrckData"`
}

// PlusOrMinusIndicator ...
type PlusOrMinusIndicator bool

// PostalAddress1 ...
type PostalAddress1 struct {
	XMLName     xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AdrTp       string   `xml:"AdrTp", json:"AdrTp", bson:"AdrTp"`
	AdrLine     []string `xml:"AdrLine", json:"AdrLine", bson:"AdrLine"`
	StrtNm      string   `xml:"StrtNm", json:"StrtNm", bson:"StrtNm"`
	BldgNb      string   `xml:"BldgNb", json:"BldgNb", bson:"BldgNb"`
	PstCd       string   `xml:"PstCd", json:"PstCd", bson:"PstCd"`
	TwnNm       string   `xml:"TwnNm", json:"TwnNm", bson:"TwnNm"`
	CtrySubDvsn string   `xml:"CtrySubDvsn", json:"CtrySubDvsn", bson:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry", json:"Ctry", bson:"Ctry"`
}

// PostalAddress18 ...
type PostalAddress18 struct {
	XMLName     xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AdrLine     []string `xml:"AdrLine", json:"AdrLine", bson:"AdrLine"`
	StrtNm      string   `xml:"StrtNm", json:"StrtNm", bson:"StrtNm"`
	BldgNb      string   `xml:"BldgNb", json:"BldgNb", bson:"BldgNb"`
	PstCd       string   `xml:"PstCd", json:"PstCd", bson:"PstCd"`
	TwnNm       string   `xml:"TwnNm", json:"TwnNm", bson:"TwnNm"`
	CtrySubDvsn []string `xml:"CtrySubDvsn", json:"CtrySubDvsn", bson:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry", json:"Ctry", bson:"Ctry"`
}

// Recipient4Choice ...
type Recipient4Choice struct {
	XMLName    xml.Name        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	KeyTrnsprt *KeyTransport4  `xml:"KeyTrnsprt", json:"KeyTrnsprt", bson:"KeyTrnsprt"`
	KEK        *KEK4           `xml:"KEK", json:"KEK", bson:"KEK"`
	KeyIdr     *KEKIdentifier2 `xml:"KeyIdr", json:"KeyIdr", bson:"KeyIdr"`
}

// Recipient5Choice ...
type Recipient5Choice struct {
	XMLName      xml.Name                `json:"-,omitempty", bson:"-,omitempty", xml:""`
	IssrAndSrlNb *IssuerAndSerialNumber1 `xml:"IssrAndSrlNb", json:"IssrAndSrlNb", bson:"IssrAndSrlNb"`
	KeyIdr       *KEKIdentifier2         `xml:"KeyIdr", json:"KeyIdr", bson:"KeyIdr"`
}

// RelativeDistinguishedName1 ...
type RelativeDistinguishedName1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AttrTp  string   `xml:"AttrTp", json:"AttrTp", bson:"AttrTp"`
	AttrVal string   `xml:"AttrVal", json:"AttrVal", bson:"AttrVal"`
}

// Response3Code ...
type Response3Code string

// ResponseType2 ...
type ResponseType2 struct {
	XMLName      xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Rslt         string   `xml:"Rslt", json:"Rslt", bson:"Rslt"`
	RsltDtls     string   `xml:"RsltDtls", json:"RsltDtls", bson:"RsltDtls"`
	AddtlRsltInf string   `xml:"AddtlRsltInf", json:"AddtlRsltInf", bson:"AddtlRsltInf"`
}

// ResultDetail1Code ...
type ResultDetail1Code string

// Traceability3 ...
type Traceability3 struct {
	XMLName     xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	RlayId      *GenericIdentification74 `xml:"RlayId", json:"RlayId", bson:"RlayId"`
	TracDtTmIn  string                   `xml:"TracDtTmIn", json:"TracDtTmIn", bson:"TracDtTmIn"`
	TracDtTmOut string                   `xml:"TracDtTmOut", json:"TracDtTmOut", bson:"TracDtTmOut"`
}

// TrackData1 ...
type TrackData1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TrckNb  string   `xml:"TrckNb", json:"TrckNb", bson:"TrckNb"`
	TrckVal string   `xml:"TrckVal", json:"TrckVal", bson:"TrckVal"`
}

// TransactionIdentifier2 ...
type TransactionIdentifier2 struct {
	XMLName   xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	RcncltnDt string   `xml:"RcncltnDt", json:"RcncltnDt", bson:"RcncltnDt"`
	RcncltnId string   `xml:"RcncltnId", json:"RcncltnId", bson:"RcncltnId"`
}

// TransactionVerificationResult4 ...
type TransactionVerificationResult4 struct {
	XMLName    xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Mtd        string   `xml:"Mtd", json:"Mtd", bson:"Mtd"`
	VrfctnNtty string   `xml:"VrfctnNtty", json:"VrfctnNtty", bson:"VrfctnNtty"`
	Rslt       string   `xml:"Rslt", json:"Rslt", bson:"Rslt"`
	AddtlRslt  string   `xml:"AddtlRslt", json:"AddtlRslt", bson:"AddtlRslt"`
}

// TypeOfAmount5Code ...
type TypeOfAmount5Code string

// TypeOfAmount6Code ...
type TypeOfAmount6Code string

// UPICIdentifier ...
type UPICIdentifier string

// UserInterface3Code ...
type UserInterface3Code string

// UserInterface4Code ...
type UserInterface4Code string

// Verification1Code ...
type Verification1Code string
