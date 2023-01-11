// Code generated by xgen. DO NOT EDIT.

package cain_001_001_01

import (
	"encoding/xml"
)

// Document ...
type Document *Document

// AccountChoiceMethod1Code ...
type AccountChoiceMethod1Code string

// AccountIdentification30Choice ...
type AccountIdentification30Choice struct {
	XMLName *xml.Name `xml:"AccountIdentification30Choice"`
	Card    string    `xml:"Card"`
	MSISDN  string    `xml:"MSISDN"`
	EMail   string    `xml:"EMail"`
	IBAN    string    `xml:"IBAN"`
	BBAN    string    `xml:"BBAN"`
	UPIC    string    `xml:"UPIC"`
	Dmst    string    `xml:"Dmst"`
	Othr    string    `xml:"Othr"`
}

// Acquirer6 ...
type Acquirer6 struct {
	XMLName *xml.Name `xml:"Acquirer6"`
	Id      string    `xml:"Id"`
	Issr    string    `xml:"Issr"`
	CtryCd  string    `xml:"CtryCd"`
}

// AcquirerAuthorisationInitiation ...
type AcquirerAuthorisationInitiation struct {
	XMLName      *xml.Name                         `xml:"AcquirerAuthorisationInitiation"`
	Hdr          *Header17                         `xml:"Hdr"`
	AuthstnInitn *AcquirerAuthorisationInitiation1 `xml:"AuthstnInitn"`
	SctyTrlr     *ContentInformationType15         `xml:"SctyTrlr"`
}

// AcquirerAuthorisationInitiation1 ...
type AcquirerAuthorisationInitiation1 struct {
	XMLName *xml.Name                    `xml:"AcquirerAuthorisationInitiation1"`
	Envt    *CardTransactionEnvironment1 `xml:"Envt"`
	Cntxt   *CardTransactionContext1     `xml:"Cntxt"`
	Tx      *CardTransaction15           `xml:"Tx"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName *xml.Name `xml:"ActionMessage3"`
	Dstn    string    `xml:"Dstn"`
	Frmt    string    `xml:"Frmt"`
	Cntt    string    `xml:"Cntt"`
}

// ActionType4Code ...
type ActionType4Code string

// AddressType2Code ...
type AddressType2Code string

// AddressVerification1 ...
type AddressVerification1 struct {
	XMLName    *xml.Name `xml:"AddressVerification1"`
	AdrDgts    string    `xml:"AdrDgts"`
	PstlCdDgts string    `xml:"PstlCdDgts"`
}

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
	XMLName *xml.Name   `xml:"AlgorithmIdentification11"`
	Algo    string      `xml:"Algo"`
	Param   *Parameter4 `xml:"Param"`
}

// AlgorithmIdentification12 ...
type AlgorithmIdentification12 struct {
	XMLName *xml.Name   `xml:"AlgorithmIdentification12"`
	Algo    string      `xml:"Algo"`
	Param   *Parameter5 `xml:"Param"`
}

// AlgorithmIdentification13 ...
type AlgorithmIdentification13 struct {
	XMLName *xml.Name   `xml:"AlgorithmIdentification13"`
	Algo    string      `xml:"Algo"`
	Param   *Parameter6 `xml:"Param"`
}

// AlgorithmIdentification14 ...
type AlgorithmIdentification14 struct {
	XMLName *xml.Name   `xml:"AlgorithmIdentification14"`
	Algo    string      `xml:"Algo"`
	Param   *Parameter6 `xml:"Param"`
}

// AlgorithmIdentification15 ...
type AlgorithmIdentification15 struct {
	XMLName *xml.Name   `xml:"AlgorithmIdentification15"`
	Algo    string      `xml:"Algo"`
	Param   *Parameter7 `xml:"Param"`
}

// AntiMoneyLaundering1 ...
type AntiMoneyLaundering1 struct {
	XMLName             *xml.Name            `xml:"AntiMoneyLaundering1"`
	SndrNm              string               `xml:"SndrNm"`
	SndrAdr             *PostalAddress18     `xml:"SndrAdr"`
	SndrNtlIdr          string               `xml:"SndrNtlIdr"`
	NtlIdrCtry          string               `xml:"NtlIdrCtry"`
	SndrPsptNb          string               `xml:"SndrPsptNb"`
	PsptIssgCtry        string               `xml:"PsptIssgCtry"`
	SndrTaxIdr          string               `xml:"SndrTaxIdr"`
	TaxCtry             string               `xml:"TaxCtry"`
	SndrCstmrIdr        string               `xml:"SndrCstmrIdr"`
	SndrDtAndPlcOfBirth *DateAndPlaceOfBirth `xml:"SndrDtAndPlcOfBirth"`
	RcvrNm              string               `xml:"RcvrNm"`
	TxRef               string               `xml:"TxRef"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttendanceContext1Code ...
type AttendanceContext1Code string

// AttributeType1Code ...
type AttributeType1Code string

// AuthenticatedData4 ...
type AuthenticatedData4 struct {
	XMLName     *xml.Name                  `xml:"AuthenticatedData4"`
	Vrsn        float64                    `xml:"Vrsn"`
	Rcpt        []*Recipient4Choice        `xml:"Rcpt"`
	MACAlgo     *AlgorithmIdentification15 `xml:"MACAlgo"`
	NcpsltdCntt *EncapsulatedContent3      `xml:"NcpsltdCntt"`
	MAC         string                     `xml:"MAC"`
}

// AuthenticationEntity2Code ...
type AuthenticationEntity2Code string

// AuthenticationMethod5Code ...
type AuthenticationMethod5Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult7 ...
type AuthorisationResult7 struct {
	XMLName     *xml.Name                `xml:"AuthorisationResult7"`
	AuthstnNtty *GenericIdentification75 `xml:"AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"TxRspn"`
	AuthstnCd   string                   `xml:"AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAcceptorTerminal1 ...
type CardAcceptorTerminal1 struct {
	XMLName  *xml.Name                        `xml:"CardAcceptorTerminal1"`
	Id       *GenericIdentification32         `xml:"Id"`
	Lctn     *PostalAddress18                 `xml:"Lctn"`
	Cpblties *PointOfInteractionCapabilities4 `xml:"Cpblties"`
}

// CardAccount1 ...
type CardAccount1 struct {
	XMLName      *xml.Name                      `xml:"CardAccount1"`
	SelctnMtd    string                         `xml:"SelctnMtd"`
	SelctdAcctTp string                         `xml:"SelctdAcctTp"`
	AcctNm       string                         `xml:"AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"AcctOwnr"`
	AcctIdr      *AccountIdentification30Choice `xml:"AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"Svcr"`
}

// CardAccountType2Code ...
type CardAccountType2Code string

// CardDataReading2Code ...
type CardDataReading2Code string

// CardDataReading3Code ...
type CardDataReading3Code string

// CardFallback1Code ...
type CardFallback1Code string

// CardPaymentServiceType3Code ...
type CardPaymentServiceType3Code string

// CardPaymentServiceType7Code ...
type CardPaymentServiceType7Code string

// CardPaymentServiceType8Code ...
type CardPaymentServiceType8Code string

// CardPaymentToken4 ...
type CardPaymentToken4 struct {
	XMLName       *xml.Name                 `xml:"CardPaymentToken4"`
	Tkn           string                    `xml:"Tkn"`
	CardSeqNb     string                    `xml:"CardSeqNb"`
	TknXpryDt     string                    `xml:"TknXpryDt"`
	TknChrtc      []string                  `xml:"TknChrtc"`
	TknRqstr      *PaymentTokenIdentifiers1 `xml:"TknRqstr"`
	TknAssrncLvl  float64                   `xml:"TknAssrncLvl"`
	TknAssrncData string                    `xml:"TknAssrncData"`
}

// CardTransaction15 ...
type CardTransaction15 struct {
	XMLName           *xml.Name               `xml:"CardTransaction15"`
	TxTp              string                  `xml:"TxTp"`
	AddtlSvc          []string                `xml:"AddtlSvc"`
	SvcAttr           string                  `xml:"SvcAttr"`
	MrchntCtgyCd      string                  `xml:"MrchntCtgyCd"`
	Rcncltn           *TransactionIdentifier2 `xml:"Rcncltn"`
	AccptrTxDtTm      string                  `xml:"AccptrTxDtTm"`
	AccptrTxId        string                  `xml:"AccptrTxId"`
	InitrTxId         string                  `xml:"InitrTxId"`
	TxLifeCyclId      string                  `xml:"TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"TxLifeCyclSeqCntr"`
	OrgnlTx           *CardTransaction3       `xml:"OrgnlTx"`
	TxDtls            *CardTransactionDetail1 `xml:"TxDtls"`
	AuthstnRslt       *AuthorisationResult7   `xml:"AuthstnRslt"`
}

// CardTransaction3 ...
type CardTransaction3 struct {
	XMLName      *xml.Name `xml:"CardTransaction3"`
	AccptrTxDtTm string    `xml:"AccptrTxDtTm"`
	InitrTxId    string    `xml:"InitrTxId"`
	InitrId      string    `xml:"InitrId"`
}

// CardTransactionAmount1 ...
type CardTransactionAmount1 struct {
	XMLName          *xml.Name          `xml:"CardTransactionAmount1"`
	TtlAmt           *CurrencyAndAmount `xml:"TtlAmt"`
	AmtQlfr          string             `xml:"AmtQlfr"`
	CrdhldrBllgTxAmt *DetailedAmount8   `xml:"CrdhldrBllgTxAmt"`
	DtldAmt          []*DetailedAmount9 `xml:"DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	XMLName *xml.Name `xml:"CardTransactionCondition1"`
	Prgm    string    `xml:"Prgm"`
	Val     string    `xml:"Val"`
}

// CardTransactionContext1 ...
type CardTransactionContext1 struct {
	XMLName   *xml.Name                `xml:"CardTransactionContext1"`
	TxCntxt   *CardTransactionContext2 `xml:"TxCntxt"`
	SaleCntxt *SaleContext1            `xml:"SaleCntxt"`
}

// CardTransactionContext2 ...
type CardTransactionContext2 struct {
	XMLName        *xml.Name                        `xml:"CardTransactionContext2"`
	CardPres       bool                             `xml:"CardPres"`
	CrdhldrPres    bool                             `xml:"CrdhldrPres"`
	LctnCtgy       string                           `xml:"LctnCtgy"`
	AttndncCntxt   string                           `xml:"AttndncCntxt"`
	TxEnvt         string                           `xml:"TxEnvt"`
	HstgCtgy       string                           `xml:"HstgCtgy"`
	TxChanl        string                           `xml:"TxChanl"`
	CardDataNtryMd string                           `xml:"CardDataNtryMd"`
	FllbckInd      string                           `xml:"FllbckInd"`
	SpprtdOptn     []string                         `xml:"SpprtdOptn"`
	SpclConds      []*CardTransactionCondition1     `xml:"SpclConds"`
	RskInd         []*CardTransactionRiskIndicator1 `xml:"RskInd"`
}

// CardTransactionDetail1 ...
type CardTransactionDetail1 struct {
	XMLName        *xml.Name               `xml:"CardTransactionDetail1"`
	TxAmts         *CardTransactionAmount1 `xml:"TxAmts"`
	AddtlAmts      []*DetailedAmount10     `xml:"AddtlAmts"`
	MsgRsn         string                  `xml:"MsgRsn"`
	VldtyDt        string                  `xml:"VldtyDt"`
	UattnddLvlCtgy string                  `xml:"UattnddLvlCtgy"`
	AcctFr         *CardAccount1           `xml:"AcctFr"`
	AcctTo         *CardAccount1           `xml:"AcctTo"`
	Instlmt        *RecurringTransaction2  `xml:"Instlmt"`
	AML            *AntiMoneyLaundering1   `xml:"AML"`
	ICCRltdData    string                  `xml:"ICCRltdData"`
}

// CardTransactionEnvironment1 ...
type CardTransactionEnvironment1 struct {
	XMLName           *xml.Name                 `xml:"CardTransactionEnvironment1"`
	Acqrr             *Acquirer6                `xml:"Acqrr"`
	CardSchmeId       string                    `xml:"CardSchmeId"`
	Accptr            *Organisation18           `xml:"Accptr"`
	Termnl            *CardAcceptorTerminal1    `xml:"Termnl"`
	Card              *PaymentCard12            `xml:"Card"`
	CstmrDvc          *CustomerDevice1          `xml:"CstmrDvc"`
	Wllt              *CustomerDevice1          `xml:"Wllt"`
	PmtTkn            *CardPaymentToken4        `xml:"PmtTkn"`
	Crdhldr           *Cardholder9              `xml:"Crdhldr"`
	PrtctdCrdhldrData *ContentInformationType10 `xml:"PrtctdCrdhldrData"`
}

// CardTransactionRiskIndicator1 ...
type CardTransactionRiskIndicator1 struct {
	XMLName     *xml.Name `xml:"CardTransactionRiskIndicator1"`
	Rsn         []string  `xml:"Rsn"`
	Lvl         float64   `xml:"Lvl"`
	RcmmnddActn []string  `xml:"RcmmnddActn"`
}

// CardTransactionRiskReason1Code ...
type CardTransactionRiskReason1Code string

// Cardholder9 ...
type Cardholder9 struct {
	XMLName      *xml.Name                         `xml:"Cardholder9"`
	Id           *PersonIdentification7            `xml:"Id"`
	Nm           string                            `xml:"Nm"`
	Lang         string                            `xml:"Lang"`
	BllgAdr      *PostalAddress18                  `xml:"BllgAdr"`
	ShppgAdr     *PostalAddress18                  `xml:"ShppgAdr"`
	Authntcn     []*CardholderAuthentication7      `xml:"Authntcn"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"TxVrfctnRslt"`
	PrsnlData    string                            `xml:"PrsnlData"`
}

// CardholderAuthentication7 ...
type CardholderAuthentication7 struct {
	XMLName           *xml.Name                 `xml:"CardholderAuthentication7"`
	AuthntcnMtd       string                    `xml:"AuthntcnMtd"`
	AuthntcnVal       string                    `xml:"AuthntcnVal"`
	PrtctdAuthntcnVal *ContentInformationType10 `xml:"PrtctdAuthntcnVal"`
	CrdhldrOnLinePIN  *OnLinePIN4               `xml:"CrdhldrOnLinePIN"`
	CrdhldrId         *PersonIdentification7    `xml:"CrdhldrId"`
	AdrVrfctn         *AddressVerification1     `xml:"AdrVrfctn"`
}

// CardholderVerificationCapability2Code ...
type CardholderVerificationCapability2Code string

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        *xml.Name                     `xml:"CertificateIssuer1"`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"RltvDstngshdNm"`
}

// CommunicationAddress5 ...
type CommunicationAddress5 struct {
	XMLName      *xml.Name        `xml:"CommunicationAddress5"`
	PstlAdr      *PostalAddress18 `xml:"PstlAdr"`
	Email        string           `xml:"Email"`
	URLAdr       string           `xml:"URLAdr"`
	Phne         string           `xml:"Phne"`
	CstmrSvc     string           `xml:"CstmrSvc"`
	AddtlCtctInf string           `xml:"AddtlCtctInf"`
}

// ContentInformationType10 ...
type ContentInformationType10 struct {
	XMLName    *xml.Name       `xml:"ContentInformationType10"`
	CnttTp     string          `xml:"CnttTp"`
	EnvlpdData *EnvelopedData4 `xml:"EnvlpdData"`
}

// ContentInformationType15 ...
type ContentInformationType15 struct {
	XMLName      *xml.Name           `xml:"ContentInformationType15"`
	CnttTp       string              `xml:"CnttTp"`
	AuthntcdData *AuthenticatedData4 `xml:"AuthntcdData"`
}

// ContentType2Code ...
type ContentType2Code string

// CountryCode ...
type CountryCode string

// CurrencyAndAmountSimpleType ...
type CurrencyAndAmountSimpleType float64

// CurrencyAndAmount ...
type CurrencyAndAmount struct {
	XMLName *xml.Name `xml:"CurrencyAndAmount"`
	CcyAttr string    `json:"-,omitempty", bson:"-,omitempty", xml:"Ccy,attr"`
	Value   float64   `xml:",chardata"`
}

// CurrencyCode ...
type CurrencyCode string

// CustomerDevice1 ...
type CustomerDevice1 struct {
	XMLName *xml.Name `xml:"CustomerDevice1"`
	Id      string    `xml:"Id"`
	Tp      string    `xml:"Tp"`
	Prvdr   string    `xml:"Prvdr"`
}

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	XMLName     *xml.Name `xml:"DateAndPlaceOfBirth"`
	BirthDt     string    `xml:"BirthDt"`
	PrvcOfBirth string    `xml:"PrvcOfBirth"`
	CityOfBirth string    `xml:"CityOfBirth"`
	CtryOfBirth string    `xml:"CtryOfBirth"`
}

// DetailedAmount10 ...
type DetailedAmount10 struct {
	XMLName *xml.Name          `xml:"DetailedAmount10"`
	Tp      string             `xml:"Tp"`
	AddtlTp string             `xml:"AddtlTp"`
	Amt     *CurrencyAndAmount `xml:"Amt"`
	Labl    string             `xml:"Labl"`
}

// DetailedAmount8 ...
type DetailedAmount8 struct {
	XMLName  *xml.Name `xml:"DetailedAmount8"`
	Amt      float64   `xml:"Amt"`
	XchgRate float64   `xml:"XchgRate"`
	QtnDt    string    `xml:"QtnDt"`
	Labl     string    `xml:"Labl"`
}

// DetailedAmount9 ...
type DetailedAmount9 struct {
	XMLName *xml.Name `xml:"DetailedAmount9"`
	Tp      string    `xml:"Tp"`
	AddtlTp string    `xml:"AddtlTp"`
	Amt     float64   `xml:"Amt"`
	Labl    string    `xml:"Labl"`
}

// DisplayCapabilities3 ...
type DisplayCapabilities3 struct {
	XMLName   *xml.Name `xml:"DisplayCapabilities3"`
	Dstn      string    `xml:"Dstn"`
	AvlblFrmt []string  `xml:"AvlblFrmt"`
	NbOfLines float64   `xml:"NbOfLines"`
	LineWidth float64   `xml:"LineWidth"`
	AvlblLang []string  `xml:"AvlblLang"`
}

// EncapsulatedContent3 ...
type EncapsulatedContent3 struct {
	XMLName *xml.Name `xml:"EncapsulatedContent3"`
	CnttTp  string    `xml:"CnttTp"`
	Cntt    string    `xml:"Cntt"`
}

// EncryptedContent3 ...
type EncryptedContent3 struct {
	XMLName        *xml.Name                  `xml:"EncryptedContent3"`
	CnttTp         string                     `xml:"CnttTp"`
	CnttNcrptnAlgo *AlgorithmIdentification14 `xml:"CnttNcrptnAlgo"`
	NcrptdData     string                     `xml:"NcrptdData"`
}

// EncryptionFormat1Code ...
type EncryptionFormat1Code string

// EnvelopedData4 ...
type EnvelopedData4 struct {
	XMLName    *xml.Name           `xml:"EnvelopedData4"`
	Vrsn       float64             `xml:"Vrsn"`
	Rcpt       []*Recipient4Choice `xml:"Rcpt"`
	NcrptdCntt *EncryptedContent3  `xml:"NcrptdCntt"`
}

// Exact1NumericText ...
type Exact1NumericText string

// Exact3AlphaNumericText ...
type Exact3AlphaNumericText string

// Exact3NumericText ...
type Exact3NumericText string

// Frequency3Code ...
type Frequency3Code string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName *xml.Name `xml:"GenericIdentification1"`
	Id      string    `xml:"Id"`
	SchmeNm string    `xml:"SchmeNm"`
	Issr    string    `xml:"Issr"`
}

// GenericIdentification32 ...
type GenericIdentification32 struct {
	XMLName *xml.Name `xml:"GenericIdentification32"`
	Id      string    `xml:"Id"`
	Tp      string    `xml:"Tp"`
	Issr    string    `xml:"Issr"`
	ShrtNm  string    `xml:"ShrtNm"`
}

// GenericIdentification4 ...
type GenericIdentification4 struct {
	XMLName *xml.Name `xml:"GenericIdentification4"`
	Id      string    `xml:"Id"`
	IdTp    string    `xml:"IdTp"`
}

// GenericIdentification73 ...
type GenericIdentification73 struct {
	XMLName *xml.Name `xml:"GenericIdentification73"`
	Id      string    `xml:"Id"`
	Tp      string    `xml:"Tp"`
	Issr    string    `xml:"Issr"`
	Ctry    string    `xml:"Ctry"`
	ShrtNm  string    `xml:"ShrtNm"`
}

// GenericIdentification74 ...
type GenericIdentification74 struct {
	XMLName *xml.Name `xml:"GenericIdentification74"`
	Id      string    `xml:"Id"`
	Tp      string    `xml:"Tp"`
	Issr    string    `xml:"Issr"`
	Ctry    string    `xml:"Ctry"`
	ShrtNm  string    `xml:"ShrtNm"`
}

// GenericIdentification75 ...
type GenericIdentification75 struct {
	XMLName *xml.Name `xml:"GenericIdentification75"`
	Id      string    `xml:"Id"`
	Tp      string    `xml:"Tp"`
	Issr    string    `xml:"Issr"`
	Ctry    string    `xml:"Ctry"`
	ShrtNm  string    `xml:"ShrtNm"`
}

// Header17 ...
type Header17 struct {
	XMLName        *xml.Name                `xml:"Header17"`
	MsgFctn        string                   `xml:"MsgFctn"`
	PrtcolVrsn     string                   `xml:"PrtcolVrsn"`
	XchgId         string                   `xml:"XchgId"`
	ReTrnsmssnCntr string                   `xml:"ReTrnsmssnCntr"`
	CreDtTm        string                   `xml:"CreDtTm"`
	InitgPty       *GenericIdentification73 `xml:"InitgPty"`
	RcptPty        *GenericIdentification73 `xml:"RcptPty"`
	Tracblt        []*Traceability3         `xml:"Tracblt"`
}

// IBANIdentifier ...
type IBANIdentifier string

// ISO3NumericCountryCode ...
type ISO3NumericCountryCode string

// ISODate ...
type ISODate string

// ISODateTime ...
type ISODateTime string

// ImpliedCurrencyAndAmount ...
type ImpliedCurrencyAndAmount float64

// InstalmentPlan1Code ...
type InstalmentPlan1Code string

// IssuerAndSerialNumber1 ...
type IssuerAndSerialNumber1 struct {
	XMLName *xml.Name           `xml:"IssuerAndSerialNumber1"`
	Issr    *CertificateIssuer1 `xml:"Issr"`
	SrlNb   string              `xml:"SrlNb"`
}

// KEK4 ...
type KEK4 struct {
	XMLName       *xml.Name                  `xml:"KEK4"`
	Vrsn          float64                    `xml:"Vrsn"`
	KEKId         *KEKIdentifier2            `xml:"KEKId"`
	KeyNcrptnAlgo *AlgorithmIdentification13 `xml:"KeyNcrptnAlgo"`
	NcrptdKey     string                     `xml:"NcrptdKey"`
}

// KEKIdentifier2 ...
type KEKIdentifier2 struct {
	XMLName   *xml.Name `xml:"KEKIdentifier2"`
	KeyId     string    `xml:"KeyId"`
	KeyVrsn   string    `xml:"KeyVrsn"`
	SeqNb     float64   `xml:"SeqNb"`
	DerivtnId string    `xml:"DerivtnId"`
}

// KeyTransport4 ...
type KeyTransport4 struct {
	XMLName       *xml.Name                  `xml:"KeyTransport4"`
	Vrsn          float64                    `xml:"Vrsn"`
	RcptId        *Recipient5Choice          `xml:"RcptId"`
	KeyNcrptnAlgo *AlgorithmIdentification11 `xml:"KeyNcrptnAlgo"`
	NcrptdKey     string                     `xml:"NcrptdKey"`
}

// LanguageCode ...
type LanguageCode string

// LocationCategory2Code ...
type LocationCategory2Code string

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

// Max15NumericText ...
type Max15NumericText string

// Max16Text ...
type Max16Text string

// Max20000Text ...
type Max20000Text string

// Max256Text ...
type Max256Text string

// Max2NumericText ...
type Max2NumericText string

// Max35Binary ...
type Max35Binary string

// Max35NumericText ...
type Max35NumericText string

// Max35Text ...
type Max35Text string

// Max3NumericText ...
type Max3NumericText string

// Max3Text ...
type Max3Text string

// Max45Text ...
type Max45Text string

// Max5000Binary ...
type Max5000Binary string

// Max500Binary ...
type Max500Binary string

// Max500Text ...
type Max500Text string

// Max5NumericText ...
type Max5NumericText string

// Max6Text ...
type Max6Text string

// Max70Text ...
type Max70Text string

// MessageFunction6Code ...
type MessageFunction6Code string

// MessageReason1Code ...
type MessageReason1Code string

// Min2Max3AlphaText ...
type Min2Max3AlphaText string

// Min2Max3NumericText ...
type Min2Max3NumericText string

// Min3Max4NumericText ...
type Min3Max4NumericText string

// Min5Max16Binary ...
type Min5Max16Binary string

// Min6Max8Text ...
type Min6Max8Text string

// Min8Max28NumericText ...
type Min8Max28NumericText string

// NameAndAddress3 ...
type NameAndAddress3 struct {
	XMLName *xml.Name       `xml:"NameAndAddress3"`
	Nm      string          `xml:"Nm"`
	Adr     *PostalAddress1 `xml:"Adr"`
}

// Number ...
type Number float64

// OnLineCapability1Code ...
type OnLineCapability1Code string

// OnLinePIN4 ...
type OnLinePIN4 struct {
	XMLName       *xml.Name                 `xml:"OnLinePIN4"`
	NcrptdPINBlck *ContentInformationType10 `xml:"NcrptdPINBlck"`
	PINFrmt       string                    `xml:"PINFrmt"`
	AddtlInpt     string                    `xml:"AddtlInpt"`
}

// Organisation18 ...
type Organisation18 struct {
	XMLName    *xml.Name                `xml:"Organisation18"`
	Id         *GenericIdentification32 `xml:"Id"`
	CmonNm     string                   `xml:"CmonNm"`
	Lctn       *CommunicationAddress5   `xml:"Lctn"`
	SelctdLang string                   `xml:"SelctdLang"`
	SchmeData  string                   `xml:"SchmeData"`
}

// OutputFormat1Code ...
type OutputFormat1Code string

// PINFormat3Code ...
type PINFormat3Code string

// Parameter4 ...
type Parameter4 struct {
	XMLName      *xml.Name                  `xml:"Parameter4"`
	NcrptnFrmt   string                     `xml:"NcrptnFrmt"`
	DgstAlgo     string                     `xml:"DgstAlgo"`
	MskGnrtrAlgo *AlgorithmIdentification12 `xml:"MskGnrtrAlgo"`
}

// Parameter5 ...
type Parameter5 struct {
	XMLName  *xml.Name `xml:"Parameter5"`
	DgstAlgo string    `xml:"DgstAlgo"`
}

// Parameter6 ...
type Parameter6 struct {
	XMLName      *xml.Name `xml:"Parameter6"`
	NcrptnFrmt   string    `xml:"NcrptnFrmt"`
	InitlstnVctr string    `xml:"InitlstnVctr"`
	BPddg        string    `xml:"BPddg"`
}

// Parameter7 ...
type Parameter7 struct {
	XMLName      *xml.Name `xml:"Parameter7"`
	InitlstnVctr string    `xml:"InitlstnVctr"`
	BPddg        string    `xml:"BPddg"`
}

// PartyIdentification72Choice ...
type PartyIdentification72Choice struct {
	XMLName *xml.Name               `xml:"PartyIdentification72Choice"`
	AnyBIC  string                  `xml:"AnyBIC"`
	PrtryId *GenericIdentification1 `xml:"PrtryId"`
}

// PartyType10Code ...
type PartyType10Code string

// PartyType11Code ...
type PartyType11Code string

// PartyType3Code ...
type PartyType3Code string

// PartyType4Code ...
type PartyType4Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard12 ...
type PaymentCard12 struct {
	XMLName        *xml.Name                 `xml:"PaymentCard12"`
	PrtctdCardData *ContentInformationType10 `xml:"PrtctdCardData"`
	PlainCardData  *PlainCardData10          `xml:"PlainCardData"`
	IssrBIN        string                    `xml:"IssrBIN"`
	CardCtryCd     string                    `xml:"CardCtryCd"`
	CardCcyCd      string                    `xml:"CardCcyCd"`
	AddtlCardData  string                    `xml:"AddtlCardData"`
}

// PaymentTokenIdentifiers1 ...
type PaymentTokenIdentifiers1 struct {
	XMLName *xml.Name `xml:"PaymentTokenIdentifiers1"`
	PrvdrId string    `xml:"PrvdrId"`
	RqstrId string    `xml:"RqstrId"`
}

// PersonIdentification7 ...
type PersonIdentification7 struct {
	XMLName         *xml.Name                 `xml:"PersonIdentification7"`
	DrvrsLicNb      string                    `xml:"DrvrsLicNb"`
	CstmrNb         string                    `xml:"CstmrNb"`
	SclSctyNb       string                    `xml:"SclSctyNb"`
	AlnRegnNb       string                    `xml:"AlnRegnNb"`
	PsptNb          string                    `xml:"PsptNb"`
	TaxIdNb         string                    `xml:"TaxIdNb"`
	IdntyCardNb     string                    `xml:"IdntyCardNb"`
	MplyrIdNb       string                    `xml:"MplyrIdNb"`
	MplyeeIdNb      string                    `xml:"MplyeeIdNb"`
	EmailAdr        string                    `xml:"EmailAdr"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth      `xml:"DtAndPlcOfBirth"`
	Othr            []*GenericIdentification4 `xml:"Othr"`
}

// PhoneNumber ...
type PhoneNumber string

// PlainCardData10 ...
type PlainCardData10 struct {
	XMLName   *xml.Name     `xml:"PlainCardData10"`
	PAN       string        `xml:"PAN"`
	CardSeqNb string        `xml:"CardSeqNb"`
	FctvDt    string        `xml:"FctvDt"`
	XpryDt    string        `xml:"XpryDt"`
	SvcCd     string        `xml:"SvcCd"`
	TrckData  []*TrackData1 `xml:"TrckData"`
	CrdhldrNm string        `xml:"CrdhldrNm"`
}

// PointOfInteractionCapabilities4 ...
type PointOfInteractionCapabilities4 struct {
	XMLName               *xml.Name               `xml:"PointOfInteractionCapabilities4"`
	CardRdngCpblties      []string                `xml:"CardRdngCpblties"`
	CardWrttgCpblties     []string                `xml:"CardWrttgCpblties"`
	CrdhldrVrfctnCpblties []string                `xml:"CrdhldrVrfctnCpblties"`
	PINLngthCpblties      float64                 `xml:"PINLngthCpblties"`
	ApprvlCdLngth         float64                 `xml:"ApprvlCdLngth"`
	MxScrptLngth          float64                 `xml:"MxScrptLngth"`
	CardCaptrCpbl         bool                    `xml:"CardCaptrCpbl"`
	OnLineCpblties        string                  `xml:"OnLineCpblties"`
	MsgCpblties           []*DisplayCapabilities3 `xml:"MsgCpblties"`
}

// PostalAddress1 ...
type PostalAddress1 struct {
	XMLName     *xml.Name `xml:"PostalAddress1"`
	AdrTp       string    `xml:"AdrTp"`
	AdrLine     []string  `xml:"AdrLine"`
	StrtNm      string    `xml:"StrtNm"`
	BldgNb      string    `xml:"BldgNb"`
	PstCd       string    `xml:"PstCd"`
	TwnNm       string    `xml:"TwnNm"`
	CtrySubDvsn string    `xml:"CtrySubDvsn"`
	Ctry        string    `xml:"Ctry"`
}

// PostalAddress18 ...
type PostalAddress18 struct {
	XMLName     *xml.Name `xml:"PostalAddress18"`
	AdrLine     []string  `xml:"AdrLine"`
	StrtNm      string    `xml:"StrtNm"`
	BldgNb      string    `xml:"BldgNb"`
	PstCd       string    `xml:"PstCd"`
	TwnNm       string    `xml:"TwnNm"`
	CtrySubDvsn []string  `xml:"CtrySubDvsn"`
	Ctry        string    `xml:"Ctry"`
}

// Recipient4Choice ...
type Recipient4Choice struct {
	XMLName    *xml.Name       `xml:"Recipient4Choice"`
	KeyTrnsprt *KeyTransport4  `xml:"KeyTrnsprt"`
	KEK        *KEK4           `xml:"KEK"`
	KeyIdr     *KEKIdentifier2 `xml:"KeyIdr"`
}

// Recipient5Choice ...
type Recipient5Choice struct {
	XMLName      *xml.Name               `xml:"Recipient5Choice"`
	IssrAndSrlNb *IssuerAndSerialNumber1 `xml:"IssrAndSrlNb"`
	KeyIdr       *KEKIdentifier2         `xml:"KeyIdr"`
}

// RecurringTransaction2 ...
type RecurringTransaction2 struct {
	XMLName     *xml.Name          `xml:"RecurringTransaction2"`
	InstlmtPlan []string           `xml:"InstlmtPlan"`
	PlanId      string             `xml:"PlanId"`
	SeqNb       float64            `xml:"SeqNb"`
	PrdUnit     string             `xml:"PrdUnit"`
	InstlmtPrd  float64            `xml:"InstlmtPrd"`
	TtlNbOfPmts float64            `xml:"TtlNbOfPmts"`
	FrstPmtDt   string             `xml:"FrstPmtDt"`
	TtlAmt      *CurrencyAndAmount `xml:"TtlAmt"`
	FrstAmt     float64            `xml:"FrstAmt"`
	Chrgs       float64            `xml:"Chrgs"`
}

// RelativeDistinguishedName1 ...
type RelativeDistinguishedName1 struct {
	XMLName *xml.Name `xml:"RelativeDistinguishedName1"`
	AttrTp  string    `xml:"AttrTp"`
	AttrVal string    `xml:"AttrVal"`
}

// Response3Code ...
type Response3Code string

// ResponseType2 ...
type ResponseType2 struct {
	XMLName      *xml.Name `xml:"ResponseType2"`
	Rslt         string    `xml:"Rslt"`
	RsltDtls     string    `xml:"RsltDtls"`
	AddtlRsltInf string    `xml:"AddtlRsltInf"`
}

// ResultDetail1Code ...
type ResultDetail1Code string

// SaleContext1 ...
type SaleContext1 struct {
	XMLName       *xml.Name `xml:"SaleContext1"`
	SaleId        string    `xml:"SaleId"`
	SaleRefNb     string    `xml:"SaleRefNb"`
	SaleRcncltnId string    `xml:"SaleRcncltnId"`
	CshrId        string    `xml:"CshrId"`
	ShftNb        string    `xml:"ShftNb"`
	AddtlSaleData string    `xml:"AddtlSaleData"`
}

// SupportedPaymentOption1Code ...
type SupportedPaymentOption1Code string

// Traceability3 ...
type Traceability3 struct {
	XMLName     *xml.Name                `xml:"Traceability3"`
	RlayId      *GenericIdentification74 `xml:"RlayId"`
	TracDtTmIn  string                   `xml:"TracDtTmIn"`
	TracDtTmOut string                   `xml:"TracDtTmOut"`
}

// TrackData1 ...
type TrackData1 struct {
	XMLName *xml.Name `xml:"TrackData1"`
	TrckNb  string    `xml:"TrckNb"`
	TrckVal string    `xml:"TrckVal"`
}

// TransactionChannel3Code ...
type TransactionChannel3Code string

// TransactionEnvironment2Code ...
type TransactionEnvironment2Code string

// TransactionEnvironment3Code ...
type TransactionEnvironment3Code string

// TransactionIdentifier2 ...
type TransactionIdentifier2 struct {
	XMLName   *xml.Name `xml:"TransactionIdentifier2"`
	RcncltnDt string    `xml:"RcncltnDt"`
	RcncltnId string    `xml:"RcncltnId"`
}

// TransactionVerificationResult4 ...
type TransactionVerificationResult4 struct {
	XMLName    *xml.Name `xml:"TransactionVerificationResult4"`
	Mtd        string    `xml:"Mtd"`
	VrfctnNtty string    `xml:"VrfctnNtty"`
	Rslt       string    `xml:"Rslt"`
	AddtlRslt  string    `xml:"AddtlRslt"`
}

// TrueFalseIndicator ...
type TrueFalseIndicator bool

// TypeOfAmount1Code ...
type TypeOfAmount1Code string

// TypeOfAmount5Code ...
type TypeOfAmount5Code string

// TypeOfAmount6Code ...
type TypeOfAmount6Code string

// UPICIdentifier ...
type UPICIdentifier string

// UserInterface1Code ...
type UserInterface1Code string

// UserInterface3Code ...
type UserInterface3Code string

// Verification1Code ...
type Verification1Code string
