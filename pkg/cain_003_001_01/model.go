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

// Acquirer6 ...
type Acquirer6 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
	CtryCd  string   `xml:"CtryCd", json:"CtryCd", bson:"CtryCd"`
}

// AcquirerFinancialInitiation ...
type AcquirerFinancialInitiation struct {
	XMLName  xml.Name                      `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Hdr      *Header17                     `xml:"Hdr", json:"Hdr", bson:"Hdr"`
	FinInitn *AcquirerFinancialInitiation1 `xml:"FinInitn", json:"FinInitn", bson:"FinInitn"`
	SctyTrlr *ContentInformationType15     `xml:"SctyTrlr", json:"SctyTrlr", bson:"SctyTrlr"`
}

// AcquirerFinancialInitiation1 ...
type AcquirerFinancialInitiation1 struct {
	XMLName xml.Name                     `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Envt    *CardTransactionEnvironment1 `xml:"Envt", json:"Envt", bson:"Envt"`
	Cntxt   *CardTransactionContext1     `xml:"Cntxt", json:"Cntxt", bson:"Cntxt"`
	Tx      *CardTransaction5            `xml:"Tx", json:"Tx", bson:"Tx"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Dstn    string   `xml:"Dstn", json:"Dstn", bson:"Dstn"`
	Frmt    string   `xml:"Frmt", json:"Frmt", bson:"Frmt"`
	Cntt    string   `xml:"Cntt", json:"Cntt", bson:"Cntt"`
}

// ActionType4Code ...
type ActionType4Code string

// AddressType2Code ...
type AddressType2Code string

// AddressVerification1 ...
type AddressVerification1 struct {
	XMLName    xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AdrDgts    string   `xml:"AdrDgts", json:"AdrDgts", bson:"AdrDgts"`
	PstlCdDgts string   `xml:"PstlCdDgts", json:"PstlCdDgts", bson:"PstlCdDgts"`
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

// AntiMoneyLaundering1 ...
type AntiMoneyLaundering1 struct {
	XMLName             xml.Name             `json:"-,omitempty", bson:"-,omitempty", xml:""`
	SndrNm              string               `xml:"SndrNm", json:"SndrNm", bson:"SndrNm"`
	SndrAdr             *PostalAddress18     `xml:"SndrAdr", json:"SndrAdr", bson:"SndrAdr"`
	SndrNtlIdr          string               `xml:"SndrNtlIdr", json:"SndrNtlIdr", bson:"SndrNtlIdr"`
	NtlIdrCtry          string               `xml:"NtlIdrCtry", json:"NtlIdrCtry", bson:"NtlIdrCtry"`
	SndrPsptNb          string               `xml:"SndrPsptNb", json:"SndrPsptNb", bson:"SndrPsptNb"`
	PsptIssgCtry        string               `xml:"PsptIssgCtry", json:"PsptIssgCtry", bson:"PsptIssgCtry"`
	SndrTaxIdr          string               `xml:"SndrTaxIdr", json:"SndrTaxIdr", bson:"SndrTaxIdr"`
	TaxCtry             string               `xml:"TaxCtry", json:"TaxCtry", bson:"TaxCtry"`
	SndrCstmrIdr        string               `xml:"SndrCstmrIdr", json:"SndrCstmrIdr", bson:"SndrCstmrIdr"`
	SndrDtAndPlcOfBirth *DateAndPlaceOfBirth `xml:"SndrDtAndPlcOfBirth", json:"SndrDtAndPlcOfBirth", bson:"SndrDtAndPlcOfBirth"`
	RcvrNm              string               `xml:"RcvrNm", json:"RcvrNm", bson:"RcvrNm"`
	TxRef               string               `xml:"TxRef", json:"TxRef", bson:"TxRef"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttendanceContext1Code ...
type AttendanceContext1Code string

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

// AuthenticationMethod5Code ...
type AuthenticationMethod5Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult7 ...
type AuthorisationResult7 struct {
	XMLName     xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AuthstnNtty *GenericIdentification75 `xml:"AuthstnNtty", json:"AuthstnNtty", bson:"AuthstnNtty"`
	TxRspn      *ResponseType2           `xml:"TxRspn", json:"TxRspn", bson:"TxRspn"`
	AuthstnCd   string                   `xml:"AuthstnCd", json:"AuthstnCd", bson:"AuthstnCd"`
	AddtlInf    []*ActionMessage3        `xml:"AddtlInf", json:"AddtlInf", bson:"AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAcceptorTerminal1 ...
type CardAcceptorTerminal1 struct {
	XMLName  xml.Name                         `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id       *GenericIdentification32         `xml:"Id", json:"Id", bson:"Id"`
	Lctn     *PostalAddress18                 `xml:"Lctn", json:"Lctn", bson:"Lctn"`
	Cpblties *PointOfInteractionCapabilities4 `xml:"Cpblties", json:"Cpblties", bson:"Cpblties"`
}

// CardAccount1 ...
type CardAccount1 struct {
	XMLName      xml.Name                       `json:"-,omitempty", bson:"-,omitempty", xml:""`
	SelctnMtd    string                         `xml:"SelctnMtd", json:"SelctnMtd", bson:"SelctnMtd"`
	SelctdAcctTp string                         `xml:"SelctdAcctTp", json:"SelctdAcctTp", bson:"SelctdAcctTp"`
	AcctNm       string                         `xml:"AcctNm", json:"AcctNm", bson:"AcctNm"`
	AcctOwnr     *NameAndAddress3               `xml:"AcctOwnr", json:"AcctOwnr", bson:"AcctOwnr"`
	AcctIdr      *AccountIdentification30Choice `xml:"AcctIdr", json:"AcctIdr", bson:"AcctIdr"`
	Svcr         *PartyIdentification72Choice   `xml:"Svcr", json:"Svcr", bson:"Svcr"`
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
	XMLName       xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Tkn           string                    `xml:"Tkn", json:"Tkn", bson:"Tkn"`
	CardSeqNb     string                    `xml:"CardSeqNb", json:"CardSeqNb", bson:"CardSeqNb"`
	TknXpryDt     string                    `xml:"TknXpryDt", json:"TknXpryDt", bson:"TknXpryDt"`
	TknChrtc      []string                  `xml:"TknChrtc", json:"TknChrtc", bson:"TknChrtc"`
	TknRqstr      *PaymentTokenIdentifiers1 `xml:"TknRqstr", json:"TknRqstr", bson:"TknRqstr"`
	TknAssrncLvl  float64                   `xml:"TknAssrncLvl", json:"TknAssrncLvl", bson:"TknAssrncLvl"`
	TknAssrncData []byte                    `xml:"TknAssrncData", json:"TknAssrncData", bson:"TknAssrncData"`
}

// CardTransaction3 ...
type CardTransaction3 struct {
	XMLName      xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AccptrTxDtTm string   `xml:"AccptrTxDtTm", json:"AccptrTxDtTm", bson:"AccptrTxDtTm"`
	InitrTxId    string   `xml:"InitrTxId", json:"InitrTxId", bson:"InitrTxId"`
	InitrId      string   `xml:"InitrId", json:"InitrId", bson:"InitrId"`
}

// CardTransaction5 ...
type CardTransaction5 struct {
	XMLName           xml.Name                `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TxTp              string                  `xml:"TxTp", json:"TxTp", bson:"TxTp"`
	AddtlSvc          []string                `xml:"AddtlSvc", json:"AddtlSvc", bson:"AddtlSvc"`
	SvcAttr           string                  `xml:"SvcAttr", json:"SvcAttr", bson:"SvcAttr"`
	MrchntCtgyCd      string                  `xml:"MrchntCtgyCd", json:"MrchntCtgyCd", bson:"MrchntCtgyCd"`
	Rcncltn           *TransactionIdentifier2 `xml:"Rcncltn", json:"Rcncltn", bson:"Rcncltn"`
	AccptrTxDtTm      string                  `xml:"AccptrTxDtTm", json:"AccptrTxDtTm", bson:"AccptrTxDtTm"`
	AccptrTxId        string                  `xml:"AccptrTxId", json:"AccptrTxId", bson:"AccptrTxId"`
	InitrTxId         string                  `xml:"InitrTxId", json:"InitrTxId", bson:"InitrTxId"`
	TxLifeCyclId      string                  `xml:"TxLifeCyclId", json:"TxLifeCyclId", bson:"TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `xml:"TxLifeCyclSeqNb", json:"TxLifeCyclSeqNb", bson:"TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `xml:"TxLifeCyclSeqCntr", json:"TxLifeCyclSeqCntr", bson:"TxLifeCyclSeqCntr"`
	AcqrrTxRef        string                  `xml:"AcqrrTxRef", json:"AcqrrTxRef", bson:"AcqrrTxRef"`
	CardIssrRefData   string                  `xml:"CardIssrRefData", json:"CardIssrRefData", bson:"CardIssrRefData"`
	OrgnlTx           *CardTransaction3       `xml:"OrgnlTx", json:"OrgnlTx", bson:"OrgnlTx"`
	TxDtls            *CardTransactionDetail3 `xml:"TxDtls", json:"TxDtls", bson:"TxDtls"`
	AuthstnRslt       *AuthorisationResult7   `xml:"AuthstnRslt", json:"AuthstnRslt", bson:"AuthstnRslt"`
}

// CardTransactionAmount3 ...
type CardTransactionAmount3 struct {
	XMLName          xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TtlAmt           *CurrencyAndAmount `xml:"TtlAmt", json:"TtlAmt", bson:"TtlAmt"`
	AmtQlfr          string             `xml:"AmtQlfr", json:"AmtQlfr", bson:"AmtQlfr"`
	CrdhldrBllgTxAmt *DetailedAmount8   `xml:"CrdhldrBllgTxAmt", json:"CrdhldrBllgTxAmt", bson:"CrdhldrBllgTxAmt"`
	RcncltnTxAmt     *DetailedAmount8   `xml:"RcncltnTxAmt", json:"RcncltnTxAmt", bson:"RcncltnTxAmt"`
	DtldAmt          []*DetailedAmount9 `xml:"DtldAmt", json:"DtldAmt", bson:"DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Prgm    string   `xml:"Prgm", json:"Prgm", bson:"Prgm"`
	Val     string   `xml:"Val", json:"Val", bson:"Val"`
}

// CardTransactionContext1 ...
type CardTransactionContext1 struct {
	XMLName   xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TxCntxt   *CardTransactionContext2 `xml:"TxCntxt", json:"TxCntxt", bson:"TxCntxt"`
	SaleCntxt *SaleContext1            `xml:"SaleCntxt", json:"SaleCntxt", bson:"SaleCntxt"`
}

// CardTransactionContext2 ...
type CardTransactionContext2 struct {
	XMLName        xml.Name                         `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CardPres       bool                             `xml:"CardPres", json:"CardPres", bson:"CardPres"`
	CrdhldrPres    bool                             `xml:"CrdhldrPres", json:"CrdhldrPres", bson:"CrdhldrPres"`
	LctnCtgy       string                           `xml:"LctnCtgy", json:"LctnCtgy", bson:"LctnCtgy"`
	AttndncCntxt   string                           `xml:"AttndncCntxt", json:"AttndncCntxt", bson:"AttndncCntxt"`
	TxEnvt         string                           `xml:"TxEnvt", json:"TxEnvt", bson:"TxEnvt"`
	HstgCtgy       string                           `xml:"HstgCtgy", json:"HstgCtgy", bson:"HstgCtgy"`
	TxChanl        string                           `xml:"TxChanl", json:"TxChanl", bson:"TxChanl"`
	CardDataNtryMd string                           `xml:"CardDataNtryMd", json:"CardDataNtryMd", bson:"CardDataNtryMd"`
	FllbckInd      string                           `xml:"FllbckInd", json:"FllbckInd", bson:"FllbckInd"`
	SpprtdOptn     []string                         `xml:"SpprtdOptn", json:"SpprtdOptn", bson:"SpprtdOptn"`
	SpclConds      []*CardTransactionCondition1     `xml:"SpclConds", json:"SpclConds", bson:"SpclConds"`
	RskInd         []*CardTransactionRiskIndicator1 `xml:"RskInd", json:"RskInd", bson:"RskInd"`
}

// CardTransactionDetail3 ...
type CardTransactionDetail3 struct {
	XMLName        xml.Name                `json:"-,omitempty", bson:"-,omitempty", xml:""`
	TxAmts         *CardTransactionAmount3 `xml:"TxAmts", json:"TxAmts", bson:"TxAmts"`
	TxFees         []*DetailedAmount11     `xml:"TxFees", json:"TxFees", bson:"TxFees"`
	AddtlAmts      []*DetailedAmount10     `xml:"AddtlAmts", json:"AddtlAmts", bson:"AddtlAmts"`
	MsgRsn         string                  `xml:"MsgRsn", json:"MsgRsn", bson:"MsgRsn"`
	VldtyDt        string                  `xml:"VldtyDt", json:"VldtyDt", bson:"VldtyDt"`
	UattnddLvlCtgy string                  `xml:"UattnddLvlCtgy", json:"UattnddLvlCtgy", bson:"UattnddLvlCtgy"`
	AcctFr         *CardAccount1           `xml:"AcctFr", json:"AcctFr", bson:"AcctFr"`
	AcctTo         *CardAccount1           `xml:"AcctTo", json:"AcctTo", bson:"AcctTo"`
	Instlmt        *RecurringTransaction2  `xml:"Instlmt", json:"Instlmt", bson:"Instlmt"`
	AML            *AntiMoneyLaundering1   `xml:"AML", json:"AML", bson:"AML"`
	ICCRltdData    []byte                  `xml:"ICCRltdData", json:"ICCRltdData", bson:"ICCRltdData"`
}

// CardTransactionEnvironment1 ...
type CardTransactionEnvironment1 struct {
	XMLName           xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Acqrr             *Acquirer6                `xml:"Acqrr", json:"Acqrr", bson:"Acqrr"`
	CardSchmeId       string                    `xml:"CardSchmeId", json:"CardSchmeId", bson:"CardSchmeId"`
	Accptr            *Organisation18           `xml:"Accptr", json:"Accptr", bson:"Accptr"`
	Termnl            *CardAcceptorTerminal1    `xml:"Termnl", json:"Termnl", bson:"Termnl"`
	Card              *PaymentCard12            `xml:"Card", json:"Card", bson:"Card"`
	CstmrDvc          *CustomerDevice1          `xml:"CstmrDvc", json:"CstmrDvc", bson:"CstmrDvc"`
	Wllt              *CustomerDevice1          `xml:"Wllt", json:"Wllt", bson:"Wllt"`
	PmtTkn            *CardPaymentToken4        `xml:"PmtTkn", json:"PmtTkn", bson:"PmtTkn"`
	Crdhldr           *Cardholder9              `xml:"Crdhldr", json:"Crdhldr", bson:"Crdhldr"`
	PrtctdCrdhldrData *ContentInformationType10 `xml:"PrtctdCrdhldrData", json:"PrtctdCrdhldrData", bson:"PrtctdCrdhldrData"`
}

// CardTransactionRiskIndicator1 ...
type CardTransactionRiskIndicator1 struct {
	XMLName     xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Rsn         []string `xml:"Rsn", json:"Rsn", bson:"Rsn"`
	Lvl         float64  `xml:"Lvl", json:"Lvl", bson:"Lvl"`
	RcmmnddActn []string `xml:"RcmmnddActn", json:"RcmmnddActn", bson:"RcmmnddActn"`
}

// CardTransactionRiskReason1Code ...
type CardTransactionRiskReason1Code string

// Cardholder9 ...
type Cardholder9 struct {
	XMLName      xml.Name                          `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id           *PersonIdentification7            `xml:"Id", json:"Id", bson:"Id"`
	Nm           string                            `xml:"Nm", json:"Nm", bson:"Nm"`
	Lang         string                            `xml:"Lang", json:"Lang", bson:"Lang"`
	BllgAdr      *PostalAddress18                  `xml:"BllgAdr", json:"BllgAdr", bson:"BllgAdr"`
	ShppgAdr     *PostalAddress18                  `xml:"ShppgAdr", json:"ShppgAdr", bson:"ShppgAdr"`
	Authntcn     []*CardholderAuthentication7      `xml:"Authntcn", json:"Authntcn", bson:"Authntcn"`
	TxVrfctnRslt []*TransactionVerificationResult4 `xml:"TxVrfctnRslt", json:"TxVrfctnRslt", bson:"TxVrfctnRslt"`
	PrsnlData    string                            `xml:"PrsnlData", json:"PrsnlData", bson:"PrsnlData"`
}

// CardholderAuthentication7 ...
type CardholderAuthentication7 struct {
	XMLName           xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	AuthntcnMtd       string                    `xml:"AuthntcnMtd", json:"AuthntcnMtd", bson:"AuthntcnMtd"`
	AuthntcnVal       []byte                    `xml:"AuthntcnVal", json:"AuthntcnVal", bson:"AuthntcnVal"`
	PrtctdAuthntcnVal *ContentInformationType10 `xml:"PrtctdAuthntcnVal", json:"PrtctdAuthntcnVal", bson:"PrtctdAuthntcnVal"`
	CrdhldrOnLinePIN  *OnLinePIN4               `xml:"CrdhldrOnLinePIN", json:"CrdhldrOnLinePIN", bson:"CrdhldrOnLinePIN"`
	CrdhldrId         *PersonIdentification7    `xml:"CrdhldrId", json:"CrdhldrId", bson:"CrdhldrId"`
	AdrVrfctn         *AddressVerification1     `xml:"AdrVrfctn", json:"AdrVrfctn", bson:"AdrVrfctn"`
}

// CardholderVerificationCapability2Code ...
type CardholderVerificationCapability2Code string

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	XMLName        xml.Name                      `json:"-,omitempty", bson:"-,omitempty", xml:""`
	RltvDstngshdNm []*RelativeDistinguishedName1 `xml:"RltvDstngshdNm", json:"RltvDstngshdNm", bson:"RltvDstngshdNm"`
}

// CommunicationAddress5 ...
type CommunicationAddress5 struct {
	XMLName      xml.Name         `json:"-,omitempty", bson:"-,omitempty", xml:""`
	PstlAdr      *PostalAddress18 `xml:"PstlAdr", json:"PstlAdr", bson:"PstlAdr"`
	Email        string           `xml:"Email", json:"Email", bson:"Email"`
	URLAdr       string           `xml:"URLAdr", json:"URLAdr", bson:"URLAdr"`
	Phne         string           `xml:"Phne", json:"Phne", bson:"Phne"`
	CstmrSvc     string           `xml:"CstmrSvc", json:"CstmrSvc", bson:"CstmrSvc"`
	AddtlCtctInf string           `xml:"AddtlCtctInf", json:"AddtlCtctInf", bson:"AddtlCtctInf"`
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

// CustomerDevice1 ...
type CustomerDevice1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	Tp      string   `xml:"Tp", json:"Tp", bson:"Tp"`
	Prvdr   string   `xml:"Prvdr", json:"Prvdr", bson:"Prvdr"`
}

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	XMLName     xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	BirthDt     string   `xml:"BirthDt", json:"BirthDt", bson:"BirthDt"`
	PrvcOfBirth string   `xml:"PrvcOfBirth", json:"PrvcOfBirth", bson:"PrvcOfBirth"`
	CityOfBirth string   `xml:"CityOfBirth", json:"CityOfBirth", bson:"CityOfBirth"`
	CtryOfBirth string   `xml:"CtryOfBirth", json:"CtryOfBirth", bson:"CtryOfBirth"`
}

// DetailedAmount10 ...
type DetailedAmount10 struct {
	XMLName xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Tp      string             `xml:"Tp", json:"Tp", bson:"Tp"`
	AddtlTp string             `xml:"AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt     *CurrencyAndAmount `xml:"Amt", json:"Amt", bson:"Amt"`
	Labl    string             `xml:"Labl", json:"Labl", bson:"Labl"`
}

// DetailedAmount11 ...
type DetailedAmount11 struct {
	XMLName  xml.Name              `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Tp       string                `xml:"Tp", json:"Tp", bson:"Tp"`
	AddtlTp  string                `xml:"AddtlTp", json:"AddtlTp", bson:"AddtlTp"`
	Amt      *AmountAndDirection41 `xml:"Amt", json:"Amt", bson:"Amt"`
	OrgnlAmt *AmountAndDirection41 `xml:"OrgnlAmt", json:"OrgnlAmt", bson:"OrgnlAmt"`
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

// DisplayCapabilities3 ...
type DisplayCapabilities3 struct {
	XMLName   xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Dstn      string   `xml:"Dstn", json:"Dstn", bson:"Dstn"`
	AvlblFrmt []string `xml:"AvlblFrmt", json:"AvlblFrmt", bson:"AvlblFrmt"`
	NbOfLines float64  `xml:"NbOfLines", json:"NbOfLines", bson:"NbOfLines"`
	LineWidth float64  `xml:"LineWidth", json:"LineWidth", bson:"LineWidth"`
	AvlblLang []string `xml:"AvlblLang", json:"AvlblLang", bson:"AvlblLang"`
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

// Exact3AlphaNumericText ...
type Exact3AlphaNumericText string

// Exact3NumericText ...
type Exact3NumericText string

// Frequency3Code ...
type Frequency3Code string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	SchmeNm string   `xml:"SchmeNm", json:"SchmeNm", bson:"SchmeNm"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
}

// GenericIdentification32 ...
type GenericIdentification32 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	Tp      string   `xml:"Tp", json:"Tp", bson:"Tp"`
	Issr    string   `xml:"Issr", json:"Issr", bson:"Issr"`
	ShrtNm  string   `xml:"ShrtNm", json:"ShrtNm", bson:"ShrtNm"`
}

// GenericIdentification4 ...
type GenericIdentification4 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id      string   `xml:"Id", json:"Id", bson:"Id"`
	IdTp    string   `xml:"IdTp", json:"IdTp", bson:"IdTp"`
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
	XMLName xml.Name        `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Nm      string          `xml:"Nm", json:"Nm", bson:"Nm"`
	Adr     *PostalAddress1 `xml:"Adr", json:"Adr", bson:"Adr"`
}

// Number ...
type Number float64

// OnLineCapability1Code ...
type OnLineCapability1Code string

// OnLinePIN4 ...
type OnLinePIN4 struct {
	XMLName       xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	NcrptdPINBlck *ContentInformationType10 `xml:"NcrptdPINBlck", json:"NcrptdPINBlck", bson:"NcrptdPINBlck"`
	PINFrmt       string                    `xml:"PINFrmt", json:"PINFrmt", bson:"PINFrmt"`
	AddtlInpt     string                    `xml:"AddtlInpt", json:"AddtlInpt", bson:"AddtlInpt"`
}

// Organisation18 ...
type Organisation18 struct {
	XMLName    xml.Name                 `json:"-,omitempty", bson:"-,omitempty", xml:""`
	Id         *GenericIdentification32 `xml:"Id", json:"Id", bson:"Id"`
	CmonNm     string                   `xml:"CmonNm", json:"CmonNm", bson:"CmonNm"`
	Lctn       *CommunicationAddress5   `xml:"Lctn", json:"Lctn", bson:"Lctn"`
	SelctdLang string                   `xml:"SelctdLang", json:"SelctdLang", bson:"SelctdLang"`
	SchmeData  string                   `xml:"SchmeData", json:"SchmeData", bson:"SchmeData"`
}

// OutputFormat1Code ...
type OutputFormat1Code string

// PINFormat3Code ...
type PINFormat3Code string

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

// PartyType3Code ...
type PartyType3Code string

// PartyType4Code ...
type PartyType4Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard12 ...
type PaymentCard12 struct {
	XMLName        xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	PrtctdCardData *ContentInformationType10 `xml:"PrtctdCardData", json:"PrtctdCardData", bson:"PrtctdCardData"`
	PlainCardData  *PlainCardData10          `xml:"PlainCardData", json:"PlainCardData", bson:"PlainCardData"`
	IssrBIN        string                    `xml:"IssrBIN", json:"IssrBIN", bson:"IssrBIN"`
	CardCtryCd     string                    `xml:"CardCtryCd", json:"CardCtryCd", bson:"CardCtryCd"`
	CardCcyCd      string                    `xml:"CardCcyCd", json:"CardCcyCd", bson:"CardCcyCd"`
	AddtlCardData  string                    `xml:"AddtlCardData", json:"AddtlCardData", bson:"AddtlCardData"`
}

// PaymentTokenIdentifiers1 ...
type PaymentTokenIdentifiers1 struct {
	XMLName xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	PrvdrId string   `xml:"PrvdrId", json:"PrvdrId", bson:"PrvdrId"`
	RqstrId string   `xml:"RqstrId", json:"RqstrId", bson:"RqstrId"`
}

// PersonIdentification7 ...
type PersonIdentification7 struct {
	XMLName         xml.Name                  `json:"-,omitempty", bson:"-,omitempty", xml:""`
	DrvrsLicNb      string                    `xml:"DrvrsLicNb", json:"DrvrsLicNb", bson:"DrvrsLicNb"`
	CstmrNb         string                    `xml:"CstmrNb", json:"CstmrNb", bson:"CstmrNb"`
	SclSctyNb       string                    `xml:"SclSctyNb", json:"SclSctyNb", bson:"SclSctyNb"`
	AlnRegnNb       string                    `xml:"AlnRegnNb", json:"AlnRegnNb", bson:"AlnRegnNb"`
	PsptNb          string                    `xml:"PsptNb", json:"PsptNb", bson:"PsptNb"`
	TaxIdNb         string                    `xml:"TaxIdNb", json:"TaxIdNb", bson:"TaxIdNb"`
	IdntyCardNb     string                    `xml:"IdntyCardNb", json:"IdntyCardNb", bson:"IdntyCardNb"`
	MplyrIdNb       string                    `xml:"MplyrIdNb", json:"MplyrIdNb", bson:"MplyrIdNb"`
	MplyeeIdNb      string                    `xml:"MplyeeIdNb", json:"MplyeeIdNb", bson:"MplyeeIdNb"`
	EmailAdr        string                    `xml:"EmailAdr", json:"EmailAdr", bson:"EmailAdr"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth      `xml:"DtAndPlcOfBirth", json:"DtAndPlcOfBirth", bson:"DtAndPlcOfBirth"`
	Othr            []*GenericIdentification4 `xml:"Othr", json:"Othr", bson:"Othr"`
}

// PhoneNumber ...
type PhoneNumber string

// PlainCardData10 ...
type PlainCardData10 struct {
	XMLName   xml.Name      `json:"-,omitempty", bson:"-,omitempty", xml:""`
	PAN       string        `xml:"PAN", json:"PAN", bson:"PAN"`
	CardSeqNb string        `xml:"CardSeqNb", json:"CardSeqNb", bson:"CardSeqNb"`
	FctvDt    string        `xml:"FctvDt", json:"FctvDt", bson:"FctvDt"`
	XpryDt    string        `xml:"XpryDt", json:"XpryDt", bson:"XpryDt"`
	SvcCd     string        `xml:"SvcCd", json:"SvcCd", bson:"SvcCd"`
	TrckData  []*TrackData1 `xml:"TrckData", json:"TrckData", bson:"TrckData"`
	CrdhldrNm string        `xml:"CrdhldrNm", json:"CrdhldrNm", bson:"CrdhldrNm"`
}

// PlusOrMinusIndicator ...
type PlusOrMinusIndicator bool

// PointOfInteractionCapabilities4 ...
type PointOfInteractionCapabilities4 struct {
	XMLName               xml.Name                `json:"-,omitempty", bson:"-,omitempty", xml:""`
	CardRdngCpblties      []string                `xml:"CardRdngCpblties", json:"CardRdngCpblties", bson:"CardRdngCpblties"`
	CardWrttgCpblties     []string                `xml:"CardWrttgCpblties", json:"CardWrttgCpblties", bson:"CardWrttgCpblties"`
	CrdhldrVrfctnCpblties []string                `xml:"CrdhldrVrfctnCpblties", json:"CrdhldrVrfctnCpblties", bson:"CrdhldrVrfctnCpblties"`
	PINLngthCpblties      float64                 `xml:"PINLngthCpblties", json:"PINLngthCpblties", bson:"PINLngthCpblties"`
	ApprvlCdLngth         float64                 `xml:"ApprvlCdLngth", json:"ApprvlCdLngth", bson:"ApprvlCdLngth"`
	MxScrptLngth          float64                 `xml:"MxScrptLngth", json:"MxScrptLngth", bson:"MxScrptLngth"`
	CardCaptrCpbl         bool                    `xml:"CardCaptrCpbl", json:"CardCaptrCpbl", bson:"CardCaptrCpbl"`
	OnLineCpblties        string                  `xml:"OnLineCpblties", json:"OnLineCpblties", bson:"OnLineCpblties"`
	MsgCpblties           []*DisplayCapabilities3 `xml:"MsgCpblties", json:"MsgCpblties", bson:"MsgCpblties"`
}

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

// RecurringTransaction2 ...
type RecurringTransaction2 struct {
	XMLName     xml.Name           `json:"-,omitempty", bson:"-,omitempty", xml:""`
	InstlmtPlan []string           `xml:"InstlmtPlan", json:"InstlmtPlan", bson:"InstlmtPlan"`
	PlanId      string             `xml:"PlanId", json:"PlanId", bson:"PlanId"`
	SeqNb       float64            `xml:"SeqNb", json:"SeqNb", bson:"SeqNb"`
	PrdUnit     string             `xml:"PrdUnit", json:"PrdUnit", bson:"PrdUnit"`
	InstlmtPrd  float64            `xml:"InstlmtPrd", json:"InstlmtPrd", bson:"InstlmtPrd"`
	TtlNbOfPmts float64            `xml:"TtlNbOfPmts", json:"TtlNbOfPmts", bson:"TtlNbOfPmts"`
	FrstPmtDt   string             `xml:"FrstPmtDt", json:"FrstPmtDt", bson:"FrstPmtDt"`
	TtlAmt      *CurrencyAndAmount `xml:"TtlAmt", json:"TtlAmt", bson:"TtlAmt"`
	FrstAmt     float64            `xml:"FrstAmt", json:"FrstAmt", bson:"FrstAmt"`
	Chrgs       float64            `xml:"Chrgs", json:"Chrgs", bson:"Chrgs"`
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

// SaleContext1 ...
type SaleContext1 struct {
	XMLName       xml.Name `json:"-,omitempty", bson:"-,omitempty", xml:""`
	SaleId        string   `xml:"SaleId", json:"SaleId", bson:"SaleId"`
	SaleRefNb     string   `xml:"SaleRefNb", json:"SaleRefNb", bson:"SaleRefNb"`
	SaleRcncltnId string   `xml:"SaleRcncltnId", json:"SaleRcncltnId", bson:"SaleRcncltnId"`
	CshrId        string   `xml:"CshrId", json:"CshrId", bson:"CshrId"`
	ShftNb        string   `xml:"ShftNb", json:"ShftNb", bson:"ShftNb"`
	AddtlSaleData string   `xml:"AddtlSaleData", json:"AddtlSaleData", bson:"AddtlSaleData"`
}

// SupportedPaymentOption1Code ...
type SupportedPaymentOption1Code string

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

// TransactionChannel3Code ...
type TransactionChannel3Code string

// TransactionEnvironment2Code ...
type TransactionEnvironment2Code string

// TransactionEnvironment3Code ...
type TransactionEnvironment3Code string

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
