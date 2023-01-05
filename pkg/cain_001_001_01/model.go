// Code generated automatically. DO NOT EDIT.

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

// Acquirer6 ...
type Acquirer6 struct {
	XMLName *xml.Name `xml:"urn1:Acquirer6", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
	CtryCd  string    `xml:"urn1:CtryCd", json:"CtryCd", bson:"CtryCd"`
}

// AcquirerAuthorisationInitiation ...
type AcquirerAuthorisationInitiation struct {
	XMLName      *xml.Name                         `xml:"urn1:AcquirerAuthorisationInitiation", json:"-,omitempty", bson:"-,omitempty"`
	Hdr          *Header17                         `xml:"urn1:Hdr", json:"Hdr", bson:"Hdr"`
	AuthstnInitn *AcquirerAuthorisationInitiation1 `xml:"urn1:AuthstnInitn", json:"AuthstnInitn", bson:"AuthstnInitn"`
	SctyTrlr     *ContentInformationType15         `xml:"urn1:SctyTrlr", json:"SctyTrlr", bson:"SctyTrlr"`
}

// AcquirerAuthorisationInitiation1 ...
type AcquirerAuthorisationInitiation1 struct {
	XMLName *xml.Name                    `xml:"urn1:AcquirerAuthorisationInitiation1", json:"-,omitempty", bson:"-,omitempty"`
	Envt    *CardTransactionEnvironment1 `xml:"urn1:Envt", json:"Envt", bson:"Envt"`
	Cntxt   *CardTransactionContext1     `xml:"urn1:Cntxt", json:"Cntxt", bson:"Cntxt"`
	Tx      *CardTransaction15           `xml:"urn1:Tx", json:"Tx", bson:"Tx"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName *xml.Name `xml:"urn1:ActionMessage3", json:"-,omitempty", bson:"-,omitempty"`
	Dstn    string    `xml:"urn1:Dstn", json:"Dstn", bson:"Dstn"`
	Frmt    string    `xml:"urn1:Frmt", json:"Frmt", bson:"Frmt"`
	Cntt    string    `xml:"urn1:Cntt", json:"Cntt", bson:"Cntt"`
}

// ActionType4Code ...
type ActionType4Code string

// AddressType2Code ...
type AddressType2Code string

// AddressVerification1 ...
type AddressVerification1 struct {
	XMLName    *xml.Name `xml:"urn1:AddressVerification1", json:"-,omitempty", bson:"-,omitempty"`
	AdrDgts    string    `xml:"urn1:AdrDgts", json:"AdrDgts", bson:"AdrDgts"`
	PstlCdDgts string    `xml:"urn1:PstlCdDgts", json:"PstlCdDgts", bson:"PstlCdDgts"`
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

// AntiMoneyLaundering1 ...
type AntiMoneyLaundering1 struct {
	XMLName             *xml.Name            `xml:"urn1:AntiMoneyLaundering1", json:"-,omitempty", bson:"-,omitempty"`
	SndrNm              string               `xml:"urn1:SndrNm", json:"SndrNm", bson:"SndrNm"`
	SndrAdr             *PostalAddress18     `xml:"urn1:SndrAdr", json:"SndrAdr", bson:"SndrAdr"`
	SndrNtlIdr          string               `xml:"urn1:SndrNtlIdr", json:"SndrNtlIdr", bson:"SndrNtlIdr"`
	NtlIdrCtry          string               `xml:"urn1:NtlIdrCtry", json:"NtlIdrCtry", bson:"NtlIdrCtry"`
	SndrPsptNb          string               `xml:"urn1:SndrPsptNb", json:"SndrPsptNb", bson:"SndrPsptNb"`
	PsptIssgCtry        string               `xml:"urn1:PsptIssgCtry", json:"PsptIssgCtry", bson:"PsptIssgCtry"`
	SndrTaxIdr          string               `xml:"urn1:SndrTaxIdr", json:"SndrTaxIdr", bson:"SndrTaxIdr"`
	TaxCtry             string               `xml:"urn1:TaxCtry", json:"TaxCtry", bson:"TaxCtry"`
	SndrCstmrIdr        string               `xml:"urn1:SndrCstmrIdr", json:"SndrCstmrIdr", bson:"SndrCstmrIdr"`
	SndrDtAndPlcOfBirth *DateAndPlaceOfBirth `xml:"urn1:SndrDtAndPlcOfBirth", json:"SndrDtAndPlcOfBirth", bson:"SndrDtAndPlcOfBirth"`
	RcvrNm              string               `xml:"urn1:RcvrNm", json:"RcvrNm", bson:"RcvrNm"`
	TxRef               string               `xml:"urn1:TxRef", json:"TxRef", bson:"TxRef"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttendanceContext1Code ...
type AttendanceContext1Code string

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

// AuthenticationMethod5Code ...
type AuthenticationMethod5Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult7 ...
type AuthorisationResult7 struct {
	XMLName     *xml.Name                `xml:"urn1:AuthorisationResult7", json:"-,omitempty", bson:"-,omitempty"`
	AuthstnNtty *GenericIdentification75 `xml:"urn1:AuthstnNtty", json:"AuthstnNtty", bson:"AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"urn1:TxRspn", json:"TxRspn", bson:"TxRspn"`
	AuthstnCd   string                   `xml:"urn1:AuthstnCd", json:"AuthstnCd", bson:"AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"urn1:AddtlInf", json:"AddtlInf", bson:"AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAcceptorTerminal1 ...
type CardAcceptorTerminal1 struct {
	XMLName  *xml.Name                        `xml:"urn1:CardAcceptorTerminal1", json:"-,omitempty", bson:"-,omitempty"`
	Id       *GenericIdentification32         `xml:"urn1:Id", json:"Id", bson:"Id"`
	Lctn     *PostalAddress18                 `xml:"urn1:Lctn", json:"Lctn", bson:"Lctn"`
	Cpblties *PointOfInteractionCapabilities4 `xml:"urn1:Cpblties", json:"Cpblties", bson:"Cpblties"`
}

// CardAccount1 ...
type CardAccount1 struct {
	XMLName      *xml.Name                      `xml:"urn1:CardAccount1", json:"-,omitempty", bson:"-,omitempty"`
	SelctnMtd    string                         `xml:"urn1:SelctnMtd", json:"SelctnMtd", bson:"SelctnMtd"`
	SelctdAcctTp string                         `xml:"urn1:SelctdAcctTp", json:"SelctdAcctTp", bson:"SelctdAcctTp"`
	AcctNm       string                         `xml:"urn1:AcctNm", json:"AcctNm", bson:"AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"urn1:AcctOwnr", json:"AcctOwnr", bson:"AcctOwnr"`
	AcctIdr      *AccountIdentification30Choice `xml:"urn1:AcctIdr", json:"AcctIdr", bson:"AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"urn1:Svcr", json:"Svcr", bson:"Svcr"`
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
	XMLName       *xml.Name                 `xml:"urn1:CardPaymentToken4", json:"-,omitempty", bson:"-,omitempty"`
	Tkn           string                    `xml:"urn1:Tkn", json:"Tkn", bson:"Tkn"`
	CardSeqNb     string                    `xml:"urn1:CardSeqNb", json:"CardSeqNb", bson:"CardSeqNb"`
	TknXpryDt     string                    `xml:"urn1:TknXpryDt", json:"TknXpryDt", bson:"TknXpryDt"`
	TknChrtc      []string                  `xml:"urn1:TknChrtc", json:"TknChrtc", bson:"TknChrtc"`
	TknRqstr      *PaymentTokenIdentifiers1 `xml:"urn1:TknRqstr", json:"TknRqstr", bson:"TknRqstr"`
	TknAssrncLvl  float64                   `xml:"urn1:TknAssrncLvl", json:"TknAssrncLvl", bson:"TknAssrncLvl"`
	TknAssrncData []byte                    `xml:"urn1:TknAssrncData", json:"TknAssrncData", bson:"TknAssrncData"`
}

// CardTransaction15 ...
type CardTransaction15 struct {
	XMLName           *xml.Name               `xml:"urn1:CardTransaction15", json:"-,omitempty", bson:"-,omitempty"`
	TxTp              string                  `xml:"urn1:TxTp", json:"TxTp", bson:"TxTp"`
	AddtlSvc          []string                `xml:"urn1:AddtlSvc", json:"AddtlSvc", bson:"AddtlSvc"`
	SvcAttr           string                  `xml:"urn1:SvcAttr", json:"SvcAttr", bson:"SvcAttr"`
	MrchntCtgyCd      string                  `xml:"urn1:MrchntCtgyCd", json:"MrchntCtgyCd", bson:"MrchntCtgyCd"`
	Rcncltn           *TransactionIdentifier2 `xml:"urn1:Rcncltn", json:"Rcncltn", bson:"Rcncltn"`
	AccptrTxDtTm      string                  `xml:"urn1:AccptrTxDtTm", json:"AccptrTxDtTm", bson:"AccptrTxDtTm"`
	AccptrTxId        string                  `xml:"urn1:AccptrTxId", json:"AccptrTxId", bson:"AccptrTxId"`
	InitrTxId         string                  `xml:"urn1:InitrTxId", json:"InitrTxId", bson:"InitrTxId"`
	TxLifeCyclId      string                  `xml:"urn1:TxLifeCyclId", json:"TxLifeCyclId", bson:"TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"urn1:TxLifeCyclSeqNb", json:"TxLifeCyclSeqNb", bson:"TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"urn1:TxLifeCyclSeqCntr", json:"TxLifeCyclSeqCntr", bson:"TxLifeCyclSeqCntr"`
	OrgnlTx           *CardTransaction3       `xml:"urn1:OrgnlTx", json:"OrgnlTx", bson:"OrgnlTx"`
	TxDtls            *CardTransactionDetail1 `xml:"urn1:TxDtls", json:"TxDtls", bson:"TxDtls"`
	AuthstnRslt       *AuthorisationResult7   `xml:"urn1:AuthstnRslt", json:"AuthstnRslt", bson:"AuthstnRslt"`
}

// CardTransaction3 ...
type CardTransaction3 struct {
	XMLName      *xml.Name `xml:"urn1:CardTransaction3", json:"-,omitempty", bson:"-,omitempty"`
	AccptrTxDtTm string    `xml:"urn1:AccptrTxDtTm", json:"AccptrTxDtTm", bson:"AccptrTxDtTm"`
	InitrTxId    string    `xml:"urn1:InitrTxId", json:"InitrTxId", bson:"InitrTxId"`
	InitrId      string    `xml:"urn1:InitrId", json:"InitrId", bson:"InitrId"`
}

// CardTransactionAmount1 ...
type CardTransactionAmount1 struct {
	XMLName          *xml.Name          `xml:"urn1:CardTransactionAmount1", json:"-,omitempty", bson:"-,omitempty"`
	TtlAmt           *CurrencyAndAmount `xml:"urn1:TtlAmt", json:"TtlAmt", bson:"TtlAmt"`
	AmtQlfr          string             `xml:"urn1:AmtQlfr", json:"AmtQlfr", bson:"AmtQlfr"`
	CrdhldrBllgTxAmt *DetailedAmount8   `xml:"urn1:CrdhldrBllgTxAmt", json:"CrdhldrBllgTxAmt", bson:"CrdhldrBllgTxAmt"`
	DtldAmt          []*DetailedAmount9 `xml:"urn1:DtldAmt", json:"DtldAmt", bson:"DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	XMLName *xml.Name `xml:"urn1:CardTransactionCondition1", json:"-,omitempty", bson:"-,omitempty"`
	Prgm    string    `xml:"urn1:Prgm", json:"Prgm", bson:"Prgm"`
	Val     string    `xml:"urn1:Val", json:"Val", bson:"Val"`
}

// CardTransactionContext1 ...
type CardTransactionContext1 struct {
	XMLName   *xml.Name                `xml:"urn1:CardTransactionContext1", json:"-,omitempty", bson:"-,omitempty"`
	TxCntxt   *CardTransactionContext2 `xml:"urn1:TxCntxt", json:"TxCntxt", bson:"TxCntxt"`
	SaleCntxt *SaleContext1            `xml:"urn1:SaleCntxt", json:"SaleCntxt", bson:"SaleCntxt"`
}

// CardTransactionContext2 ...
type CardTransactionContext2 struct {
	XMLName        *xml.Name                        `xml:"urn1:CardTransactionContext2", json:"-,omitempty", bson:"-,omitempty"`
	CardPres       bool                             `xml:"urn1:CardPres", json:"CardPres", bson:"CardPres"`
	CrdhldrPres    bool                             `xml:"urn1:CrdhldrPres", json:"CrdhldrPres", bson:"CrdhldrPres"`
	LctnCtgy       string                           `xml:"urn1:LctnCtgy", json:"LctnCtgy", bson:"LctnCtgy"`
	AttndncCntxt   string                           `xml:"urn1:AttndncCntxt", json:"AttndncCntxt", bson:"AttndncCntxt"`
	TxEnvt         string                           `xml:"urn1:TxEnvt", json:"TxEnvt", bson:"TxEnvt"`
	HstgCtgy       string                           `xml:"urn1:HstgCtgy", json:"HstgCtgy", bson:"HstgCtgy"`
	TxChanl        string                           `xml:"urn1:TxChanl", json:"TxChanl", bson:"TxChanl"`
	CardDataNtryMd string                           `xml:"urn1:CardDataNtryMd", json:"CardDataNtryMd", bson:"CardDataNtryMd"`
	FllbckInd      string                           `xml:"urn1:FllbckInd", json:"FllbckInd", bson:"FllbckInd"`
	SpprtdOptn     []string                         `xml:"urn1:SpprtdOptn", json:"SpprtdOptn", bson:"SpprtdOptn"`
	SpclConds      []*CardTransactionCondition1     `xml:"urn1:SpclConds", json:"SpclConds", bson:"SpclConds"`
	RskInd         []*CardTransactionRiskIndicator1 `xml:"urn1:RskInd", json:"RskInd", bson:"RskInd"`
}

// CardTransactionDetail1 ...
type CardTransactionDetail1 struct {
	XMLName        *xml.Name               `xml:"urn1:CardTransactionDetail1", json:"-,omitempty", bson:"-,omitempty"`
	TxAmts         *CardTransactionAmount1 `xml:"urn1:TxAmts", json:"TxAmts", bson:"TxAmts"`
	AddtlAmts      []*DetailedAmount10     `xml:"urn1:AddtlAmts", json:"AddtlAmts", bson:"AddtlAmts"`
	MsgRsn         string                  `xml:"urn1:MsgRsn", json:"MsgRsn", bson:"MsgRsn"`
	VldtyDt        string                  `xml:"urn1:VldtyDt", json:"VldtyDt", bson:"VldtyDt"`
	UattnddLvlCtgy string                  `xml:"urn1:UattnddLvlCtgy", json:"UattnddLvlCtgy", bson:"UattnddLvlCtgy"`
	AcctFr         *CardAccount1           `xml:"urn1:AcctFr", json:"AcctFr", bson:"AcctFr"`
	AcctTo         *CardAccount1           `xml:"urn1:AcctTo", json:"AcctTo", bson:"AcctTo"`
	Instlmt        *RecurringTransaction2  `xml:"urn1:Instlmt", json:"Instlmt", bson:"Instlmt"`
	AML            *AntiMoneyLaundering1   `xml:"urn1:AML", json:"AML", bson:"AML"`
	ICCRltdData    []byte                  `xml:"urn1:ICCRltdData", json:"ICCRltdData", bson:"ICCRltdData"`
}

// CardTransactionEnvironment1 ...
type CardTransactionEnvironment1 struct {
	XMLName           *xml.Name                 `xml:"urn1:CardTransactionEnvironment1", json:"-,omitempty", bson:"-,omitempty"`
	Acqrr             *Acquirer6                `xml:"urn1:Acqrr", json:"Acqrr", bson:"Acqrr"`
	CardSchmeId       string                    `xml:"urn1:CardSchmeId", json:"CardSchmeId", bson:"CardSchmeId"`
	Accptr            *Organisation18           `xml:"urn1:Accptr", json:"Accptr", bson:"Accptr"`
	Termnl            *CardAcceptorTerminal1    `xml:"urn1:Termnl", json:"Termnl", bson:"Termnl"`
	Card              *PaymentCard12            `xml:"urn1:Card", json:"Card", bson:"Card"`
	CstmrDvc          *CustomerDevice1          `xml:"urn1:CstmrDvc", json:"CstmrDvc", bson:"CstmrDvc"`
	Wllt              *CustomerDevice1          `xml:"urn1:Wllt", json:"Wllt", bson:"Wllt"`
	PmtTkn            *CardPaymentToken4        `xml:"urn1:PmtTkn", json:"PmtTkn", bson:"PmtTkn"`
	Crdhldr           *Cardholder9              `xml:"urn1:Crdhldr", json:"Crdhldr", bson:"Crdhldr"`
	PrtctdCrdhldrData *ContentInformationType10 `xml:"urn1:PrtctdCrdhldrData", json:"PrtctdCrdhldrData", bson:"PrtctdCrdhldrData"`
}

// CardTransactionRiskIndicator1 ...
type CardTransactionRiskIndicator1 struct {
	XMLName     *xml.Name `xml:"urn1:CardTransactionRiskIndicator1", json:"-,omitempty", bson:"-,omitempty"`
	Rsn         []string  `xml:"urn1:Rsn", json:"Rsn", bson:"Rsn"`
	Lvl         float64   `xml:"urn1:Lvl", json:"Lvl", bson:"Lvl"`
	RcmmnddActn []string  `xml:"urn1:RcmmnddActn", json:"RcmmnddActn", bson:"RcmmnddActn"`
}

// CardTransactionRiskReason1Code ...
type CardTransactionRiskReason1Code string

// Cardholder9 ...
type Cardholder9 struct {
	XMLName      *xml.Name                         `xml:"urn1:Cardholder9", json:"-,omitempty", bson:"-,omitempty"`
	Id           *PersonIdentification7            `xml:"urn1:Id", json:"Id", bson:"Id"`
	Nm           string                            `xml:"urn1:Nm", json:"Nm", bson:"Nm"`
	Lang         string                            `xml:"urn1:Lang", json:"Lang", bson:"Lang"`
	BllgAdr      *PostalAddress18                  `xml:"urn1:BllgAdr", json:"BllgAdr", bson:"BllgAdr"`
	ShppgAdr     *PostalAddress18                  `xml:"urn1:ShppgAdr", json:"ShppgAdr", bson:"ShppgAdr"`
	Authntcn     []*CardholderAuthentication7      `xml:"urn1:Authntcn", json:"Authntcn", bson:"Authntcn"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"urn1:TxVrfctnRslt", json:"TxVrfctnRslt", bson:"TxVrfctnRslt"`
	PrsnlData    string                            `xml:"urn1:PrsnlData", json:"PrsnlData", bson:"PrsnlData"`
}

// CardholderAuthentication7 ...
type CardholderAuthentication7 struct {
	XMLName           *xml.Name                 `xml:"urn1:CardholderAuthentication7", json:"-,omitempty", bson:"-,omitempty"`
	AuthntcnMtd       string                    `xml:"urn1:AuthntcnMtd", json:"AuthntcnMtd", bson:"AuthntcnMtd"`
	AuthntcnVal       []byte                    `xml:"urn1:AuthntcnVal", json:"AuthntcnVal", bson:"AuthntcnVal"`
	PrtctdAuthntcnVal *ContentInformationType10 `xml:"urn1:PrtctdAuthntcnVal", json:"PrtctdAuthntcnVal", bson:"PrtctdAuthntcnVal"`
	CrdhldrOnLinePIN  *OnLinePIN4               `xml:"urn1:CrdhldrOnLinePIN", json:"CrdhldrOnLinePIN", bson:"CrdhldrOnLinePIN"`
	CrdhldrId         *PersonIdentification7    `xml:"urn1:CrdhldrId", json:"CrdhldrId", bson:"CrdhldrId"`
	AdrVrfctn         *AddressVerification1     `xml:"urn1:AdrVrfctn", json:"AdrVrfctn", bson:"AdrVrfctn"`
}

// CardholderVerificationCapability2Code ...
type CardholderVerificationCapability2Code string

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        *xml.Name                     `xml:"urn1:CertificateIssuer1", json:"-,omitempty", bson:"-,omitempty"`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"urn1:RltvDstngshdNm", json:"RltvDstngshdNm", bson:"RltvDstngshdNm"`
}

// CommunicationAddress5 ...
type CommunicationAddress5 struct {
	XMLName      *xml.Name        `xml:"urn1:CommunicationAddress5", json:"-,omitempty", bson:"-,omitempty"`
	PstlAdr      *PostalAddress18 `xml:"urn1:PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
	Email        string           `xml:"urn1:Email", json:"Email", bson:"Email"`
	URLAdr       string           `xml:"urn1:URLAdr", json:"URLAdr", bson:"URLAdr"`
	Phne         string           `xml:"urn1:Phne", json:"Phne", bson:"Phne"`
	CstmrSvc     string           `xml:"urn1:CstmrSvc", json:"CstmrSvc", bson:"CstmrSvc"`
	AddtlCtctInf string           `xml:"urn1:AddtlCtctInf", json:"AddtlCtctInf", bson:"AddtlCtctInf"`
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

// CustomerDevice1 ...
type CustomerDevice1 struct {
	XMLName *xml.Name `xml:"urn1:CustomerDevice1", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	Tp      string    `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	Prvdr   string    `xml:"urn1:Prvdr", json:"Prvdr", bson:"Prvdr"`
}

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	XMLName     *xml.Name `xml:"urn1:DateAndPlaceOfBirth", json:"-,omitempty", bson:"-,omitempty"`
	BirthDt     string    `xml:"urn1:BirthDt", json:"BirthDt", bson:"BirthDt"`
	PrvcOfBirth string    `xml:"urn1:PrvcOfBirth", json:"PrvcOfBirth", bson:"PrvcOfBirth"`
	CityOfBirth string    `xml:"urn1:CityOfBirth", json:"CityOfBirth", bson:"CityOfBirth"`
	CtryOfBirth string    `xml:"urn1:CtryOfBirth", json:"CtryOfBirth", bson:"CtryOfBirth"`
}

// DetailedAmount10 ...
type DetailedAmount10 struct {
	XMLName *xml.Name          `xml:"urn1:DetailedAmount10", json:"-,omitempty", bson:"-,omitempty"`
	Tp      string             `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	AddtlTp string             `xml:"urn1:AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt     *CurrencyAndAmount `xml:"urn1:Amt", json:"Amt", bson:"Amt"`
	Labl    string             `xml:"urn1:Labl", json:"Labl", bson:"Labl"`
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

// DisplayCapabilities3 ...
type DisplayCapabilities3 struct {
	XMLName   *xml.Name `xml:"urn1:DisplayCapabilities3", json:"-,omitempty", bson:"-,omitempty"`
	Dstn      string    `xml:"urn1:Dstn", json:"Dstn", bson:"Dstn"`
	AvlblFrmt []string  `xml:"urn1:AvlblFrmt", json:"AvlblFrmt", bson:"AvlblFrmt"`
	NbOfLines float64   `xml:"urn1:NbOfLines", json:"NbOfLines", bson:"NbOfLines"`
	LineWidth float64   `xml:"urn1:LineWidth", json:"LineWidth", bson:"LineWidth"`
	AvlblLang []string  `xml:"urn1:AvlblLang", json:"AvlblLang", bson:"AvlblLang"`
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

// Exact3AlphaNumericText ...
type Exact3AlphaNumericText string

// Exact3NumericText ...
type Exact3NumericText string

// Frequency3Code ...
type Frequency3Code string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification1", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	SchmeNm string    `xml:"urn1:SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
}

// GenericIdentification32 ...
type GenericIdentification32 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification32", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	Tp      string    `xml:"urn1:Tp", json:"Tp", bson:"Tp"`
	Issr    string    `xml:"urn1:Issr", json:"Issr", bson:"Issr"`
	ShrtNm  string    `xml:"urn1:ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// GenericIdentification4 ...
type GenericIdentification4 struct {
	XMLName *xml.Name `xml:"urn1:GenericIdentification4", json:"-,omitempty", bson:"-,omitempty"`
	Id      string    `xml:"urn1:Id", json:"Id", bson:"Id"`
	IdTp    string    `xml:"urn1:IdTp", json:"IdTp", bson:"IdTp"`
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

// LanguageCode ...
type LanguageCode string

// LocationCategory2Code ...
type LocationCategory2Code string

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
type Max35Binary []byte

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
type Max5000Binary []byte

// Max500Binary ...
type Max500Binary []byte

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

// OnLineCapability1Code ...
type OnLineCapability1Code string

// OnLinePIN4 ...
type OnLinePIN4 struct {
	XMLName       *xml.Name                 `xml:"urn1:OnLinePIN4", json:"-,omitempty", bson:"-,omitempty"`
	NcrptdPINBlck *ContentInformationType10 `xml:"urn1:NcrptdPINBlck", json:"NcrptdPINBlck", bson:"NcrptdPINBlck"`
	PINFrmt       string                    `xml:"urn1:PINFrmt", json:"PINFrmt", bson:"PINFrmt"`
	AddtlInpt     string                    `xml:"urn1:AddtlInpt", json:"AddtlInpt", bson:"AddtlInpt"`
}

// Organisation18 ...
type Organisation18 struct {
	XMLName    *xml.Name                `xml:"urn1:Organisation18", json:"-,omitempty", bson:"-,omitempty"`
	Id         *GenericIdentification32 `xml:"urn1:Id", json:"Id", bson:"Id"`
	CmonNm     string                   `xml:"urn1:CmonNm", json:"CmonNm", bson:"CmonNm"`
	Lctn       *CommunicationAddress5   `xml:"urn1:Lctn", json:"Lctn", bson:"Lctn"`
	SelctdLang string                   `xml:"urn1:SelctdLang", json:"SelctdLang", bson:"SelctdLang"`
	SchmeData  string                   `xml:"urn1:SchmeData", json:"SchmeData", bson:"SchmeData"`
}

// OutputFormat1Code ...
type OutputFormat1Code string

// PINFormat3Code ...
type PINFormat3Code string

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

// PartyType3Code ...
type PartyType3Code string

// PartyType4Code ...
type PartyType4Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard12 ...
type PaymentCard12 struct {
	XMLName        *xml.Name                 `xml:"urn1:PaymentCard12", json:"-,omitempty", bson:"-,omitempty"`
	PrtctdCardData *ContentInformationType10 `xml:"urn1:PrtctdCardData", json:"PrtctdCardData", bson:"PrtctdCardData"`
	PlainCardData  *PlainCardData10          `xml:"urn1:PlainCardData", json:"PlainCardData", bson:"PlainCardData"`
	IssrBIN        string                    `xml:"urn1:IssrBIN", json:"IssrBIN", bson:"IssrBIN"`
	CardCtryCd     string                    `xml:"urn1:CardCtryCd", json:"CardCtryCd", bson:"CardCtryCd"`
	CardCcyCd      string                    `xml:"urn1:CardCcyCd", json:"CardCcyCd", bson:"CardCcyCd"`
	AddtlCardData  string                    `xml:"urn1:AddtlCardData", json:"AddtlCardData", bson:"AddtlCardData"`
}

// PaymentTokenIdentifiers1 ...
type PaymentTokenIdentifiers1 struct {
	XMLName *xml.Name `xml:"urn1:PaymentTokenIdentifiers1", json:"-,omitempty", bson:"-,omitempty"`
	PrvdrId string    `xml:"urn1:PrvdrId", json:"PrvdrId", bson:"PrvdrId"`
	RqstrId string    `xml:"urn1:RqstrId", json:"RqstrId", bson:"RqstrId"`
}

// PersonIdentification7 ...
type PersonIdentification7 struct {
	XMLName         *xml.Name                 `xml:"urn1:PersonIdentification7", json:"-,omitempty", bson:"-,omitempty"`
	DrvrsLicNb      string                    `xml:"urn1:DrvrsLicNb", json:"DrvrsLicNb", bson:"DrvrsLicNb"`
	CstmrNb         string                    `xml:"urn1:CstmrNb", json:"CstmrNb", bson:"CstmrNb"`
	SclSctyNb       string                    `xml:"urn1:SclSctyNb", json:"SclSctyNb", bson:"SclSctyNb"`
	AlnRegnNb       string                    `xml:"urn1:AlnRegnNb", json:"AlnRegnNb", bson:"AlnRegnNb"`
	PsptNb          string                    `xml:"urn1:PsptNb", json:"PsptNb", bson:"PsptNb"`
	TaxIdNb         string                    `xml:"urn1:TaxIdNb", json:"TaxIdNb", bson:"TaxIdNb"`
	IdntyCardNb     string                    `xml:"urn1:IdntyCardNb", json:"IdntyCardNb", bson:"IdntyCardNb"`
	MplyrIdNb       string                    `xml:"urn1:MplyrIdNb", json:"MplyrIdNb", bson:"MplyrIdNb"`
	MplyeeIdNb      string                    `xml:"urn1:MplyeeIdNb", json:"MplyeeIdNb", bson:"MplyeeIdNb"`
	EmailAdr        string                    `xml:"urn1:EmailAdr", json:"EmailAdr", bson:"EmailAdr"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth      `xml:"urn1:DtAndPlcOfBirth", json:"DtAndPlcOfBirth", bson:"DtAndPlcOfBirth"`
	Othr            []*GenericIdentification4 `xml:"urn1:Othr", json:"Othr", bson:"Othr"`
}

// PhoneNumber ...
type PhoneNumber string

// PlainCardData10 ...
type PlainCardData10 struct {
	XMLName   *xml.Name     `xml:"urn1:PlainCardData10", json:"-,omitempty", bson:"-,omitempty"`
	PAN       string        `xml:"urn1:PAN", json:"PAN", bson:"PAN"`
	CardSeqNb string        `xml:"urn1:CardSeqNb", json:"CardSeqNb", bson:"CardSeqNb"`
	FctvDt    string        `xml:"urn1:FctvDt", json:"FctvDt", bson:"FctvDt"`
	XpryDt    string        `xml:"urn1:XpryDt", json:"XpryDt", bson:"XpryDt"`
	SvcCd     string        `xml:"urn1:SvcCd", json:"SvcCd", bson:"SvcCd"`
	TrckData  []*TrackData1 `xml:"urn1:TrckData", json:"TrckData", bson:"TrckData"`
	CrdhldrNm string        `xml:"urn1:CrdhldrNm", json:"CrdhldrNm", bson:"CrdhldrNm"`
}

// PointOfInteractionCapabilities4 ...
type PointOfInteractionCapabilities4 struct {
	XMLName               *xml.Name               `xml:"urn1:PointOfInteractionCapabilities4", json:"-,omitempty", bson:"-,omitempty"`
	CardRdngCpblties      []string                `xml:"urn1:CardRdngCpblties", json:"CardRdngCpblties", bson:"CardRdngCpblties"`
	CardWrttgCpblties     []string                `xml:"urn1:CardWrttgCpblties", json:"CardWrttgCpblties", bson:"CardWrttgCpblties"`
	CrdhldrVrfctnCpblties []string                `xml:"urn1:CrdhldrVrfctnCpblties", json:"CrdhldrVrfctnCpblties", bson:"CrdhldrVrfctnCpblties"`
	PINLngthCpblties      float64                 `xml:"urn1:PINLngthCpblties", json:"PINLngthCpblties", bson:"PINLngthCpblties"`
	ApprvlCdLngth         float64                 `xml:"urn1:ApprvlCdLngth", json:"ApprvlCdLngth", bson:"ApprvlCdLngth"`
	MxScrptLngth          float64                 `xml:"urn1:MxScrptLngth", json:"MxScrptLngth", bson:"MxScrptLngth"`
	CardCaptrCpbl         bool                    `xml:"urn1:CardCaptrCpbl", json:"CardCaptrCpbl", bson:"CardCaptrCpbl"`
	OnLineCpblties        string                  `xml:"urn1:OnLineCpblties", json:"OnLineCpblties", bson:"OnLineCpblties"`
	MsgCpblties           []*DisplayCapabilities3 `xml:"urn1:MsgCpblties", json:"MsgCpblties", bson:"MsgCpblties"`
}

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

// RecurringTransaction2 ...
type RecurringTransaction2 struct {
	XMLName     *xml.Name          `xml:"urn1:RecurringTransaction2", json:"-,omitempty", bson:"-,omitempty"`
	InstlmtPlan []string           `xml:"urn1:InstlmtPlan", json:"InstlmtPlan", bson:"InstlmtPlan"`
	PlanId      string             `xml:"urn1:PlanId", json:"PlanId", bson:"PlanId"`
	SeqNb       float64            `xml:"urn1:SeqNb", json:"SeqNb", bson:"SeqNb"`
	PrdUnit     string             `xml:"urn1:PrdUnit", json:"PrdUnit", bson:"PrdUnit"`
	InstlmtPrd  float64            `xml:"urn1:InstlmtPrd", json:"InstlmtPrd", bson:"InstlmtPrd"`
	TtlNbOfPmts float64            `xml:"urn1:TtlNbOfPmts", json:"TtlNbOfPmts", bson:"TtlNbOfPmts"`
	FrstPmtDt   string             `xml:"urn1:FrstPmtDt", json:"FrstPmtDt", bson:"FrstPmtDt"`
	TtlAmt      *CurrencyAndAmount `xml:"urn1:TtlAmt", json:"TtlAmt", bson:"TtlAmt"`
	FrstAmt     float64            `xml:"urn1:FrstAmt", json:"FrstAmt", bson:"FrstAmt"`
	Chrgs       float64            `xml:"urn1:Chrgs", json:"Chrgs", bson:"Chrgs"`
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

// SaleContext1 ...
type SaleContext1 struct {
	XMLName       *xml.Name `xml:"urn1:SaleContext1", json:"-,omitempty", bson:"-,omitempty"`
	SaleId        string    `xml:"urn1:SaleId", json:"SaleId", bson:"SaleId"`
	SaleRefNb     string    `xml:"urn1:SaleRefNb", json:"SaleRefNb", bson:"SaleRefNb"`
	SaleRcncltnId string    `xml:"urn1:SaleRcncltnId", json:"SaleRcncltnId", bson:"SaleRcncltnId"`
	CshrId        string    `xml:"urn1:CshrId", json:"CshrId", bson:"CshrId"`
	ShftNb        string    `xml:"urn1:ShftNb", json:"ShftNb", bson:"ShftNb"`
	AddtlSaleData string    `xml:"urn1:AddtlSaleData", json:"AddtlSaleData", bson:"AddtlSaleData"`
}

// SupportedPaymentOption1Code ...
type SupportedPaymentOption1Code string

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

// TransactionChannel3Code ...
type TransactionChannel3Code string

// TransactionEnvironment2Code ...
type TransactionEnvironment2Code string

// TransactionEnvironment3Code ...
type TransactionEnvironment3Code string

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
