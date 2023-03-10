// Code generated by xgen. DO NOT EDIT.

package cain_003_001_01

import (
	"encoding/xml"
)

// Document ...
type Document *Document

// AccountChoiceMethod1Code ...
type AccountChoiceMethod1Code string

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

// Acquirer6 ...
type Acquirer6 struct {
	XMLName *xml.Name `xml:"urn1:Acquirer6"`
	Id      string    `xml:"urn1:Id"`
	Issr    string    `xml:"urn1:Issr"`
	CtryCd  string    `xml:"urn1:CtryCd"`
}

// AcquirerFinancialInitiation ...
type AcquirerFinancialInitiation struct {
	XMLName  *xml.Name                     `xml:"urn1:AcquirerFinancialInitiation"`
	Hdr      *Header17                     `xml:"urn1:Hdr"`
	FinInitn *AcquirerFinancialInitiation1 `xml:"urn1:FinInitn"`
	SctyTrlr *ContentInformationType15     `xml:"urn1:SctyTrlr"`
}

// AcquirerFinancialInitiation1 ...
type AcquirerFinancialInitiation1 struct {
	XMLName *xml.Name                    `xml:"urn1:AcquirerFinancialInitiation1"`
	Envt    *CardTransactionEnvironment1 `xml:"urn1:Envt"`
	Cntxt   *CardTransactionContext1     `xml:"urn1:Cntxt"`
	Tx      *CardTransaction5            `xml:"urn1:Tx"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName *xml.Name `xml:"urn1:ActionMessage3"`
	Dstn    string    `xml:"urn1:Dstn"`
	Frmt    string    `xml:"urn1:Frmt"`
	Cntt    string    `xml:"urn1:Cntt"`
}

// ActionType4Code ...
type ActionType4Code string

// AddressType2Code ...
type AddressType2Code string

// AddressVerification1 ...
type AddressVerification1 struct {
	XMLName    *xml.Name `xml:"urn1:AddressVerification1"`
	AdrDgts    string    `xml:"urn1:AdrDgts"`
	PstlCdDgts string    `xml:"urn1:PstlCdDgts"`
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

// AntiMoneyLaundering1 ...
type AntiMoneyLaundering1 struct {
	XMLName             *xml.Name            `xml:"urn1:AntiMoneyLaundering1"`
	SndrNm              string               `xml:"urn1:SndrNm"`
	SndrAdr             *PostalAddress18     `xml:"urn1:SndrAdr"`
	SndrNtlIdr          string               `xml:"urn1:SndrNtlIdr"`
	NtlIdrCtry          string               `xml:"urn1:NtlIdrCtry"`
	SndrPsptNb          string               `xml:"urn1:SndrPsptNb"`
	PsptIssgCtry        string               `xml:"urn1:PsptIssgCtry"`
	SndrTaxIdr          string               `xml:"urn1:SndrTaxIdr"`
	TaxCtry             string               `xml:"urn1:TaxCtry"`
	SndrCstmrIdr        string               `xml:"urn1:SndrCstmrIdr"`
	SndrDtAndPlcOfBirth *DateAndPlaceOfBirth `xml:"urn1:SndrDtAndPlcOfBirth"`
	RcvrNm              string               `xml:"urn1:RcvrNm"`
	TxRef               string               `xml:"urn1:TxRef"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttendanceContext1Code ...
type AttendanceContext1Code string

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

// AuthenticationMethod5Code ...
type AuthenticationMethod5Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult7 ...
type AuthorisationResult7 struct {
	XMLName     *xml.Name                `xml:"urn1:AuthorisationResult7"`
	AuthstnNtty *GenericIdentification75 `xml:"urn1:AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"urn1:TxRspn"`
	AuthstnCd   string                   `xml:"urn1:AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"urn1:AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAcceptorTerminal1 ...
type CardAcceptorTerminal1 struct {
	XMLName  *xml.Name                        `xml:"urn1:CardAcceptorTerminal1"`
	Id       *GenericIdentification32         `xml:"urn1:Id"`
	Lctn     *PostalAddress18                 `xml:"urn1:Lctn"`
	Cpblties *PointOfInteractionCapabilities4 `xml:"urn1:Cpblties"`
}

// CardAccount1 ...
type CardAccount1 struct {
	XMLName      *xml.Name                      `xml:"urn1:CardAccount1"`
	SelctnMtd    string                         `xml:"urn1:SelctnMtd"`
	SelctdAcctTp string                         `xml:"urn1:SelctdAcctTp"`
	AcctNm       string                         `xml:"urn1:AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"urn1:AcctOwnr"`
	AcctIdr      *AccountIdentification30Choice `xml:"urn1:AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"urn1:Svcr"`
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
	XMLName       *xml.Name                 `xml:"urn1:CardPaymentToken4"`
	Tkn           string                    `xml:"urn1:Tkn"`
	CardSeqNb     string                    `xml:"urn1:CardSeqNb"`
	TknXpryDt     string                    `xml:"urn1:TknXpryDt"`
	TknChrtc      []string                  `xml:"urn1:TknChrtc"`
	TknRqstr      *PaymentTokenIdentifiers1 `xml:"urn1:TknRqstr"`
	TknAssrncLvl  float64                   `xml:"urn1:TknAssrncLvl"`
	TknAssrncData string                    `xml:"urn1:TknAssrncData"`
}

// CardTransaction3 ...
type CardTransaction3 struct {
	XMLName      *xml.Name `xml:"urn1:CardTransaction3"`
	AccptrTxDtTm string    `xml:"urn1:AccptrTxDtTm"`
	InitrTxId    string    `xml:"urn1:InitrTxId"`
	InitrId      string    `xml:"urn1:InitrId"`
}

// CardTransaction5 ...
type CardTransaction5 struct {
	XMLName           *xml.Name               `xml:"urn1:CardTransaction5"`
	TxTp              string                  `xml:"urn1:TxTp"`
	AddtlSvc          []string                `xml:"urn1:AddtlSvc"`
	SvcAttr           string                  `xml:"urn1:SvcAttr"`
	MrchntCtgyCd      string                  `xml:"urn1:MrchntCtgyCd"`
	Rcncltn           *TransactionIdentifier2 `xml:"urn1:Rcncltn"`
	AccptrTxDtTm      string                  `xml:"urn1:AccptrTxDtTm"`
	AccptrTxId        string                  `xml:"urn1:AccptrTxId"`
	InitrTxId         string                  `xml:"urn1:InitrTxId"`
	TxLifeCyclId      string                  `xml:"urn1:TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"urn1:TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"urn1:TxLifeCyclSeqCntr"`
	AcqrrTxRef        string                  `xml:"urn1:AcqrrTxRef"`
	CardIssrRefData   string                  `xml:"urn1:CardIssrRefData"`
	OrgnlTx           *CardTransaction3       `xml:"urn1:OrgnlTx"`
	TxDtls            *CardTransactionDetail3 `xml:"urn1:TxDtls"`
	AuthstnRslt       *AuthorisationResult7   `xml:"urn1:AuthstnRslt"`
}

// CardTransactionAmount3 ...
type CardTransactionAmount3 struct {
	XMLName          *xml.Name          `xml:"urn1:CardTransactionAmount3"`
	TtlAmt           *CurrencyAndAmount `xml:"urn1:TtlAmt"`
	AmtQlfr          string             `xml:"urn1:AmtQlfr"`
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

// CardTransactionContext1 ...
type CardTransactionContext1 struct {
	XMLName   *xml.Name                `xml:"urn1:CardTransactionContext1"`
	TxCntxt   *CardTransactionContext2 `xml:"urn1:TxCntxt"`
	SaleCntxt *SaleContext1            `xml:"urn1:SaleCntxt"`
}

// CardTransactionContext2 ...
type CardTransactionContext2 struct {
	XMLName        *xml.Name                        `xml:"urn1:CardTransactionContext2"`
	CardPres       bool                             `xml:"urn1:CardPres"`
	CrdhldrPres    bool                             `xml:"urn1:CrdhldrPres"`
	LctnCtgy       string                           `xml:"urn1:LctnCtgy"`
	AttndncCntxt   string                           `xml:"urn1:AttndncCntxt"`
	TxEnvt         string                           `xml:"urn1:TxEnvt"`
	HstgCtgy       string                           `xml:"urn1:HstgCtgy"`
	TxChanl        string                           `xml:"urn1:TxChanl"`
	CardDataNtryMd string                           `xml:"urn1:CardDataNtryMd"`
	FllbckInd      string                           `xml:"urn1:FllbckInd"`
	SpprtdOptn     []string                         `xml:"urn1:SpprtdOptn"`
	SpclConds      []*CardTransactionCondition1     `xml:"urn1:SpclConds"`
	RskInd         []*CardTransactionRiskIndicator1 `xml:"urn1:RskInd"`
}

// CardTransactionDetail3 ...
type CardTransactionDetail3 struct {
	XMLName        *xml.Name               `xml:"urn1:CardTransactionDetail3"`
	TxAmts         *CardTransactionAmount3 `xml:"urn1:TxAmts"`
	TxFees         []*DetailedAmount11     `xml:"urn1:TxFees"`
	AddtlAmts      []*DetailedAmount10     `xml:"urn1:AddtlAmts"`
	MsgRsn         string                  `xml:"urn1:MsgRsn"`
	VldtyDt        string                  `xml:"urn1:VldtyDt"`
	UattnddLvlCtgy string                  `xml:"urn1:UattnddLvlCtgy"`
	AcctFr         *CardAccount1           `xml:"urn1:AcctFr"`
	AcctTo         *CardAccount1           `xml:"urn1:AcctTo"`
	Instlmt        *RecurringTransaction2  `xml:"urn1:Instlmt"`
	AML            *AntiMoneyLaundering1   `xml:"urn1:AML"`
	ICCRltdData    string                  `xml:"urn1:ICCRltdData"`
}

// CardTransactionEnvironment1 ...
type CardTransactionEnvironment1 struct {
	XMLName           *xml.Name                 `xml:"urn1:CardTransactionEnvironment1"`
	Acqrr             *Acquirer6                `xml:"urn1:Acqrr"`
	CardSchmeId       string                    `xml:"urn1:CardSchmeId"`
	Accptr            *Organisation18           `xml:"urn1:Accptr"`
	Termnl            *CardAcceptorTerminal1    `xml:"urn1:Termnl"`
	Card              *PaymentCard12            `xml:"urn1:Card"`
	CstmrDvc          *CustomerDevice1          `xml:"urn1:CstmrDvc"`
	Wllt              *CustomerDevice1          `xml:"urn1:Wllt"`
	PmtTkn            *CardPaymentToken4        `xml:"urn1:PmtTkn"`
	Crdhldr           *Cardholder9              `xml:"urn1:Crdhldr"`
	PrtctdCrdhldrData *ContentInformationType10 `xml:"urn1:PrtctdCrdhldrData"`
}

// CardTransactionRiskIndicator1 ...
type CardTransactionRiskIndicator1 struct {
	XMLName     *xml.Name `xml:"urn1:CardTransactionRiskIndicator1"`
	Rsn         []string  `xml:"urn1:Rsn"`
	Lvl         float64   `xml:"urn1:Lvl"`
	RcmmnddActn []string  `xml:"urn1:RcmmnddActn"`
}

// CardTransactionRiskReason1Code ...
type CardTransactionRiskReason1Code string

// Cardholder9 ...
type Cardholder9 struct {
	XMLName      *xml.Name                         `xml:"urn1:Cardholder9"`
	Id           *PersonIdentification7            `xml:"urn1:Id"`
	Nm           string                            `xml:"urn1:Nm"`
	Lang         string                            `xml:"urn1:Lang"`
	BllgAdr      *PostalAddress18                  `xml:"urn1:BllgAdr"`
	ShppgAdr     *PostalAddress18                  `xml:"urn1:ShppgAdr"`
	Authntcn     []*CardholderAuthentication7      `xml:"urn1:Authntcn"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"urn1:TxVrfctnRslt"`
	PrsnlData    string                            `xml:"urn1:PrsnlData"`
}

// CardholderAuthentication7 ...
type CardholderAuthentication7 struct {
	XMLName           *xml.Name                 `xml:"urn1:CardholderAuthentication7"`
	AuthntcnMtd       string                    `xml:"urn1:AuthntcnMtd"`
	AuthntcnVal       string                    `xml:"urn1:AuthntcnVal"`
	PrtctdAuthntcnVal *ContentInformationType10 `xml:"urn1:PrtctdAuthntcnVal"`
	CrdhldrOnLinePIN  *OnLinePIN4               `xml:"urn1:CrdhldrOnLinePIN"`
	CrdhldrId         *PersonIdentification7    `xml:"urn1:CrdhldrId"`
	AdrVrfctn         *AddressVerification1     `xml:"urn1:AdrVrfctn"`
}

// CardholderVerificationCapability2Code ...
type CardholderVerificationCapability2Code string

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        *xml.Name                     `xml:"urn1:CertificateIssuer1"`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"urn1:RltvDstngshdNm"`
}

// CommunicationAddress5 ...
type CommunicationAddress5 struct {
	XMLName      *xml.Name        `xml:"urn1:CommunicationAddress5"`
	PstlAdr      *PostalAddress18 `xml:"urn1:PstlAdr"`
	Email        string           `xml:"urn1:Email"`
	URLAdr       string           `xml:"urn1:URLAdr"`
	Phne         string           `xml:"urn1:Phne"`
	CstmrSvc     string           `xml:"urn1:CstmrSvc"`
	AddtlCtctInf string           `xml:"urn1:AddtlCtctInf"`
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

// CustomerDevice1 ...
type CustomerDevice1 struct {
	XMLName *xml.Name `xml:"urn1:CustomerDevice1"`
	Id      string    `xml:"urn1:Id"`
	Tp      string    `xml:"urn1:Tp"`
	Prvdr   string    `xml:"urn1:Prvdr"`
}

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	XMLName     *xml.Name `xml:"urn1:DateAndPlaceOfBirth"`
	BirthDt     string    `xml:"urn1:BirthDt"`
	PrvcOfBirth string    `xml:"urn1:PrvcOfBirth"`
	CityOfBirth string    `xml:"urn1:CityOfBirth"`
	CtryOfBirth string    `xml:"urn1:CtryOfBirth"`
}

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

// DisplayCapabilities3 ...
type DisplayCapabilities3 struct {
	XMLName   *xml.Name `xml:"urn1:DisplayCapabilities3"`
	Dstn      string    `xml:"urn1:Dstn"`
	AvlblFrmt []string  `xml:"urn1:AvlblFrmt"`
	NbOfLines float64   `xml:"urn1:NbOfLines"`
	LineWidth float64   `xml:"urn1:LineWidth"`
	AvlblLang []string  `xml:"urn1:AvlblLang"`
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

// Exact3AlphaNumericText ...
type Exact3AlphaNumericText string

// Exact3NumericText ...
type Exact3NumericText string

// Frequency3Code ...
type Frequency3Code string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification1"`
	Id      string    `xml:"urn1:Id"`
	SchmeNm string    `xml:"urn1:SchmeNm"`
	Issr    string    `xml:"urn1:Issr"`
}

// GenericIdentification32 ...
type GenericIdentification32 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification32"`
	Id      string    `xml:"urn1:Id"`
	Tp      string    `xml:"urn1:Tp"`
	Issr    string    `xml:"urn1:Issr"`
	ShrtNm  string    `xml:"urn1:ShrtNm"`
}

// GenericIdentification4 ...
type GenericIdentification4 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification4"`
	Id      string    `xml:"urn1:Id"`
	IdTp    string    `xml:"urn1:IdTp"`
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
	XMLName *xml.Name       `xml:"urn1:NameAndAddress3"`
	Nm      string          `xml:"urn1:Nm"`
	Adr     *PostalAddress1 `xml:"urn1:Adr"`
}

// Number ...
type Number float64

// OnLineCapability1Code ...
type OnLineCapability1Code string

// OnLinePIN4 ...
type OnLinePIN4 struct {
	XMLName       *xml.Name                 `xml:"urn1:OnLinePIN4"`
	NcrptdPINBlck *ContentInformationType10 `xml:"urn1:NcrptdPINBlck"`
	PINFrmt       string                    `xml:"urn1:PINFrmt"`
	AddtlInpt     string                    `xml:"urn1:AddtlInpt"`
}

// Organisation18 ...
type Organisation18 struct {
	XMLName    *xml.Name                `xml:"urn1:Organisation18"`
	Id         *GenericIdentification32 `xml:"urn1:Id"`
	CmonNm     string                   `xml:"urn1:CmonNm"`
	Lctn       *CommunicationAddress5   `xml:"urn1:Lctn"`
	SelctdLang string                   `xml:"urn1:SelctdLang"`
	SchmeData  string                   `xml:"urn1:SchmeData"`
}

// OutputFormat1Code ...
type OutputFormat1Code string

// PINFormat3Code ...
type PINFormat3Code string

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

// PartyType3Code ...
type PartyType3Code string

// PartyType4Code ...
type PartyType4Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard12 ...
type PaymentCard12 struct {
	XMLName        *xml.Name                 `xml:"urn1:PaymentCard12"`
	PrtctdCardData *ContentInformationType10 `xml:"urn1:PrtctdCardData"`
	PlainCardData  *PlainCardData10          `xml:"urn1:PlainCardData"`
	IssrBIN        string                    `xml:"urn1:IssrBIN"`
	CardCtryCd     string                    `xml:"urn1:CardCtryCd"`
	CardCcyCd      string                    `xml:"urn1:CardCcyCd"`
	AddtlCardData  string                    `xml:"urn1:AddtlCardData"`
}

// PaymentTokenIdentifiers1 ...
type PaymentTokenIdentifiers1 struct {
	XMLName *xml.Name `xml:"urn1:PaymentTokenIdentifiers1"`
	PrvdrId string    `xml:"urn1:PrvdrId"`
	RqstrId string    `xml:"urn1:RqstrId"`
}

// PersonIdentification7 ...
type PersonIdentification7 struct {
	XMLName         *xml.Name                 `xml:"urn1:PersonIdentification7"`
	DrvrsLicNb      string                    `xml:"urn1:DrvrsLicNb"`
	CstmrNb         string                    `xml:"urn1:CstmrNb"`
	SclSctyNb       string                    `xml:"urn1:SclSctyNb"`
	AlnRegnNb       string                    `xml:"urn1:AlnRegnNb"`
	PsptNb          string                    `xml:"urn1:PsptNb"`
	TaxIdNb         string                    `xml:"urn1:TaxIdNb"`
	IdntyCardNb     string                    `xml:"urn1:IdntyCardNb"`
	MplyrIdNb       string                    `xml:"urn1:MplyrIdNb"`
	MplyeeIdNb      string                    `xml:"urn1:MplyeeIdNb"`
	EmailAdr        string                    `xml:"urn1:EmailAdr"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth      `xml:"urn1:DtAndPlcOfBirth"`
	Othr            []*GenericIdentification4 `xml:"urn1:Othr"`
}

// PhoneNumber ...
type PhoneNumber string

// PlainCardData10 ...
type PlainCardData10 struct {
	XMLName   *xml.Name     `xml:"urn1:PlainCardData10"`
	PAN       string        `xml:"urn1:PAN"`
	CardSeqNb string        `xml:"urn1:CardSeqNb"`
	FctvDt    string        `xml:"urn1:FctvDt"`
	XpryDt    string        `xml:"urn1:XpryDt"`
	SvcCd     string        `xml:"urn1:SvcCd"`
	TrckData  []*TrackData1 `xml:"urn1:TrckData"`
	CrdhldrNm string        `xml:"urn1:CrdhldrNm"`
}

// PlusOrMinusIndicator ...
type PlusOrMinusIndicator bool

// PointOfInteractionCapabilities4 ...
type PointOfInteractionCapabilities4 struct {
	XMLName               *xml.Name               `xml:"urn1:PointOfInteractionCapabilities4"`
	CardRdngCpblties      []string                `xml:"urn1:CardRdngCpblties"`
	CardWrttgCpblties     []string                `xml:"urn1:CardWrttgCpblties"`
	CrdhldrVrfctnCpblties []string                `xml:"urn1:CrdhldrVrfctnCpblties"`
	PINLngthCpblties      float64                 `xml:"urn1:PINLngthCpblties"`
	ApprvlCdLngth         float64                 `xml:"urn1:ApprvlCdLngth"`
	MxScrptLngth          float64                 `xml:"urn1:MxScrptLngth"`
	CardCaptrCpbl         bool                    `xml:"urn1:CardCaptrCpbl"`
	OnLineCpblties        string                  `xml:"urn1:OnLineCpblties"`
	MsgCpblties           []*DisplayCapabilities3 `xml:"urn1:MsgCpblties"`
}

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

// RecurringTransaction2 ...
type RecurringTransaction2 struct {
	XMLName     *xml.Name          `xml:"urn1:RecurringTransaction2"`
	InstlmtPlan []string           `xml:"urn1:InstlmtPlan"`
	PlanId      string             `xml:"urn1:PlanId"`
	SeqNb       float64            `xml:"urn1:SeqNb"`
	PrdUnit     string             `xml:"urn1:PrdUnit"`
	InstlmtPrd  float64            `xml:"urn1:InstlmtPrd"`
	TtlNbOfPmts float64            `xml:"urn1:TtlNbOfPmts"`
	FrstPmtDt   string             `xml:"urn1:FrstPmtDt"`
	TtlAmt      *CurrencyAndAmount `xml:"urn1:TtlAmt"`
	FrstAmt     float64            `xml:"urn1:FrstAmt"`
	Chrgs       float64            `xml:"urn1:Chrgs"`
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

// SaleContext1 ...
type SaleContext1 struct {
	XMLName       *xml.Name `xml:"urn1:SaleContext1"`
	SaleId        string    `xml:"urn1:SaleId"`
	SaleRefNb     string    `xml:"urn1:SaleRefNb"`
	SaleRcncltnId string    `xml:"urn1:SaleRcncltnId"`
	CshrId        string    `xml:"urn1:CshrId"`
	ShftNb        string    `xml:"urn1:ShftNb"`
	AddtlSaleData string    `xml:"urn1:AddtlSaleData"`
}

// SupportedPaymentOption1Code ...
type SupportedPaymentOption1Code string

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

// TransactionChannel3Code ...
type TransactionChannel3Code string

// TransactionEnvironment2Code ...
type TransactionEnvironment2Code string

// TransactionEnvironment3Code ...
type TransactionEnvironment3Code string

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

// TrueFalseIndicator ...
type TrueFalseIndicator bool

// TypeOfAmount1Code ...
type TypeOfAmount1Code string

// TypeOfAmount5Code ...
type TypeOfAmount5Code string

// TypeOfAmount6Code ...
type TypeOfAmount6Code string

// TypeOfAmount7Code ...
type TypeOfAmount7Code string

// UPICIdentifier ...
type UPICIdentifier string

// UserInterface1Code ...
type UserInterface1Code string

// UserInterface3Code ...
type UserInterface3Code string

// Verification1Code ...
type Verification1Code string
