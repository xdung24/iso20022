// Code generated by xgen. DO NOT EDIT.

package cain_002_001_01

// Document ...
type Document *Document

// AccountIdentification30Choice ...
type AccountIdentification30Choice struct {
	Card   string `json:"Card"`
	MSISDN string `json:"MSISDN"`
	EMail  string `json:"EMail"`
	IBAN   string `json:"IBAN"`
	BBAN   string `json:"BBAN"`
	UPIC   string `json:"UPIC"`
	Dmst   string `json:"Dmst"`
	Othr   string `json:"Othr"`
}

// AcquirerAuthorisationResponse ...
type AcquirerAuthorisationResponse struct {
	Hdr         *Header17                       `json:"Hdr"`
	AuthstnRspn *AcquirerAuthorisationResponse1 `json:"AuthstnRspn"`
	SctyTrlr    *ContentInformationType15       `json:"SctyTrlr"`
}

// AcquirerAuthorisationResponse1 ...
type AcquirerAuthorisationResponse1 struct {
	Envt  *CardTransactionEnvironment2 `json:"Envt"`
	Cntxt *CardTransactionContext3     `json:"Cntxt"`
	Tx    *CardTransaction4            `json:"Tx"`
}

// Action4 ...
type Action4 struct {
	ActnTp    string          `json:"ActnTp"`
	MsgToPres *ActionMessage2 `json:"MsgToPres"`
}

// ActionMessage2 ...
type ActionMessage2 struct {
	MsgDstn      string `json:"MsgDstn"`
	Frmt         string `json:"Frmt"`
	MsgCntt      string `json:"MsgCntt"`
	MsgCnttSgntr string `json:"MsgCnttSgntr"`
}

// ActionMessage3 ...
type ActionMessage3 struct {
	Dstn string `json:"Dstn"`
	Frmt string `json:"Frmt"`
	Cntt string `json:"Cntt"`
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
	Algo  string      `json:"Algo"`
	Param *Parameter4 `json:"Param"`
}

// AlgorithmIdentification12 ...
type AlgorithmIdentification12 struct {
	Algo  string      `json:"Algo"`
	Param *Parameter5 `json:"Param"`
}

// AlgorithmIdentification13 ...
type AlgorithmIdentification13 struct {
	Algo  string      `json:"Algo"`
	Param *Parameter6 `json:"Param"`
}

// AlgorithmIdentification14 ...
type AlgorithmIdentification14 struct {
	Algo  string      `json:"Algo"`
	Param *Parameter6 `json:"Param"`
}

// AlgorithmIdentification15 ...
type AlgorithmIdentification15 struct {
	Algo  string      `json:"Algo"`
	Param *Parameter7 `json:"Param"`
}

// AmountAndDirection41 ...
type AmountAndDirection41 struct {
	Amt *CurrencyAndAmount `json:"Amt"`
	Sgn bool               `json:"Sgn"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// AttributeType1Code ...
type AttributeType1Code string

// AuthenticatedData4 ...
type AuthenticatedData4 struct {
	Vrsn        float64                    `json:"Vrsn"`
	Rcpt        []*Recipient4Choice        `json:"Rcpt"`
	MACAlgo     *AlgorithmIdentification15 `json:"MACAlgo"`
	NcpsltdCntt *EncapsulatedContent3      `json:"NcpsltdCntt"`
	MAC         string                     `json:"MAC"`
}

// AuthenticationEntity2Code ...
type AuthenticationEntity2Code string

// AuthenticationMethod6Code ...
type AuthenticationMethod6Code string

// AuthorisationResult8 ...
type AuthorisationResult8 struct {
	AuthstnNtty *GenericIdentification75 `json:"AuthstnNtty"`
	TxRspn      *ResponseType2           `json:"TxRspn"`
	Actn        []*Action4               `json:"Actn"`
	AuthstnCd   string                   `json:"AuthstnCd"`
	AddtlInf    []*ActionMessage3        `json:"AddtlInf"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BytePadding1Code ...
type BytePadding1Code string

// CardAccount2 ...
type CardAccount2 struct {
	SelctdAcctTp string                         `json:"SelctdAcctTp"`
	AcctNm       string                         `json:"AcctNm"`
	AcctOwnr     *NameAndAddress3               `json:"AcctOwnr"`
	Ccy          string                         `json:"Ccy"`
	AcctIdr      *AccountIdentification30Choice `json:"AcctIdr"`
	Svcr         *PartyIdentification72Choice   `json:"Svcr"`
	Bal          *AmountAndDirection41          `json:"Bal"`
	BalDt        string                         `json:"BalDt"`
}

// CardAccountType2Code ...
type CardAccountType2Code string

// CardPaymentServiceType7Code ...
type CardPaymentServiceType7Code string

// CardPaymentToken2 ...
type CardPaymentToken2 struct {
	TknChrtc     []string `json:"TknChrtc"`
	TknAssrncLvl float64  `json:"TknAssrncLvl"`
}

// CardProductType1Code ...
type CardProductType1Code string

// CardTransaction4 ...
type CardTransaction4 struct {
	TxTp              string                  `json:"TxTp"`
	Rcncltn           *TransactionIdentifier2 `json:"Rcncltn"`
	AccptrTxDtTm      string                  `json:"AccptrTxDtTm"`
	InitrTxId         string                  `json:"InitrTxId"`
	TxLifeCyclId      string                  `json:"TxLifeCyclId"`
	TxLifeCyclSeqNb   float64                 `json:"TxLifeCyclSeqNb"`
	TxLifeCyclSeqCntr float64                 `json:"TxLifeCyclSeqCntr"`
	CardIssrRefData   string                  `json:"CardIssrRefData"`
	TxDtls            *CardTransactionDetail2 `json:"TxDtls"`
	AuthstnRslt       *AuthorisationResult8   `json:"AuthstnRslt"`
}

// CardTransactionAmount2 ...
type CardTransactionAmount2 struct {
	TtlAmt           *CurrencyAndAmount `json:"TtlAmt"`
	CrdhldrBllgTxAmt *DetailedAmount8   `json:"CrdhldrBllgTxAmt"`
	DtldAmt          []*DetailedAmount9 `json:"DtldAmt"`
}

// CardTransactionCondition1 ...
type CardTransactionCondition1 struct {
	Prgm string `json:"Prgm"`
	Val  string `json:"Val"`
}

// CardTransactionContext3 ...
type CardTransactionContext3 struct {
	TxCntxt *CardTransactionContext4 `json:"TxCntxt"`
}

// CardTransactionContext4 ...
type CardTransactionContext4 struct {
	SpclConds []*CardTransactionCondition1 `json:"SpclConds"`
}

// CardTransactionDetail2 ...
type CardTransactionDetail2 struct {
	TxAmts       *CardTransactionAmount2           `json:"TxAmts"`
	AddtlAmts    []*DetailedAmount10               `json:"AddtlAmts"`
	AcctAndBal   []*CardAccount2                   `json:"AcctAndBal"`
	TxVrfctnRslt []*TransactionVerificationResult4 `json:"TxVrfctnRslt"`
	VldtyDt      string                            `json:"VldtyDt"`
	ICCRltdData  string                            `json:"ICCRltdData"`
}

// CardTransactionEnvironment2 ...
type CardTransactionEnvironment2 struct {
	AcqrrId     string             `json:"AcqrrId"`
	CardSchmeId string             `json:"CardSchmeId"`
	AccptrId    string             `json:"AccptrId"`
	TermnlId    string             `json:"TermnlId"`
	Card        *PaymentCard13     `json:"Card"`
	PmtTkn      *CardPaymentToken2 `json:"PmtTkn"`
	ShppgAdr    *PostalAddress18   `json:"ShppgAdr"`
}

// CertificateIssuer1 ...
type CertificateIssuer1 struct {
	RltvDstngshdNm []*RelativeDistinguishedName1 `json:"RltvDstngshdNm"`
}

// ContentInformationType10 ...
type ContentInformationType10 struct {
	CnttTp     string          `json:"CnttTp"`
	EnvlpdData *EnvelopedData4 `json:"EnvlpdData"`
}

// ContentInformationType15 ...
type ContentInformationType15 struct {
	CnttTp       string              `json:"CnttTp"`
	AuthntcdData *AuthenticatedData4 `json:"AuthntcdData"`
}

// ContentType2Code ...
type ContentType2Code string

// CountryCode ...
type CountryCode string

// CurrencyAndAmountSimpleType ...
type CurrencyAndAmountSimpleType float64

// CurrencyAndAmount ...
type CurrencyAndAmount struct {
	CcyAttr string  `json:"-,omitempty", bson:"-,omitempty", xml:"Ccy,attr"`
	Value   float64 `json:"float64"`
}

// CurrencyCode ...
type CurrencyCode string

// DetailedAmount10 ...
type DetailedAmount10 struct {
	Tp      string             `json:"Tp"`
	AddtlTp string             `json:"AddtlTp"`
	Amt     *CurrencyAndAmount `json:"Amt"`
	Labl    string             `json:"Labl"`
}

// DetailedAmount8 ...
type DetailedAmount8 struct {
	Amt      float64 `json:"Amt"`
	XchgRate float64 `json:"XchgRate"`
	QtnDt    string  `json:"QtnDt"`
	Labl     string  `json:"Labl"`
}

// DetailedAmount9 ...
type DetailedAmount9 struct {
	Tp      string  `json:"Tp"`
	AddtlTp string  `json:"AddtlTp"`
	Amt     float64 `json:"Amt"`
	Labl    string  `json:"Labl"`
}

// EncapsulatedContent3 ...
type EncapsulatedContent3 struct {
	CnttTp string `json:"CnttTp"`
	Cntt   string `json:"Cntt"`
}

// EncryptedContent3 ...
type EncryptedContent3 struct {
	CnttTp         string                     `json:"CnttTp"`
	CnttNcrptnAlgo *AlgorithmIdentification14 `json:"CnttNcrptnAlgo"`
	NcrptdData     string                     `json:"NcrptdData"`
}

// EncryptionFormat1Code ...
type EncryptionFormat1Code string

// EnvelopedData4 ...
type EnvelopedData4 struct {
	Vrsn       float64             `json:"Vrsn"`
	Rcpt       []*Recipient4Choice `json:"Rcpt"`
	NcrptdCntt *EncryptedContent3  `json:"NcrptdCntt"`
}

// Exact1NumericText ...
type Exact1NumericText string

// Exact3NumericText ...
type Exact3NumericText string

// GenericIdentification1 ...
type GenericIdentification1 struct {
	Id      string `json:"Id"`
	SchmeNm string `json:"SchmeNm"`
	Issr    string `json:"Issr"`
}

// GenericIdentification73 ...
type GenericIdentification73 struct {
	Id     string `json:"Id"`
	Tp     string `json:"Tp"`
	Issr   string `json:"Issr"`
	Ctry   string `json:"Ctry"`
	ShrtNm string `json:"ShrtNm"`
}

// GenericIdentification74 ...
type GenericIdentification74 struct {
	Id     string `json:"Id"`
	Tp     string `json:"Tp"`
	Issr   string `json:"Issr"`
	Ctry   string `json:"Ctry"`
	ShrtNm string `json:"ShrtNm"`
}

// GenericIdentification75 ...
type GenericIdentification75 struct {
	Id     string `json:"Id"`
	Tp     string `json:"Tp"`
	Issr   string `json:"Issr"`
	Ctry   string `json:"Ctry"`
	ShrtNm string `json:"ShrtNm"`
}

// Header17 ...
type Header17 struct {
	MsgFctn        string                   `json:"MsgFctn"`
	PrtcolVrsn     string                   `json:"PrtcolVrsn"`
	XchgId         string                   `json:"XchgId"`
	ReTrnsmssnCntr string                   `json:"ReTrnsmssnCntr"`
	CreDtTm        string                   `json:"CreDtTm"`
	InitgPty       *GenericIdentification73 `json:"InitgPty"`
	RcptPty        *GenericIdentification73 `json:"RcptPty"`
	Tracblt        []*Traceability3         `json:"Tracblt"`
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
	Issr  *CertificateIssuer1 `json:"Issr"`
	SrlNb string              `json:"SrlNb"`
}

// KEK4 ...
type KEK4 struct {
	Vrsn          float64                    `json:"Vrsn"`
	KEKId         *KEKIdentifier2            `json:"KEKId"`
	KeyNcrptnAlgo *AlgorithmIdentification13 `json:"KeyNcrptnAlgo"`
	NcrptdKey     string                     `json:"NcrptdKey"`
}

// KEKIdentifier2 ...
type KEKIdentifier2 struct {
	KeyId     string  `json:"KeyId"`
	KeyVrsn   string  `json:"KeyVrsn"`
	SeqNb     float64 `json:"SeqNb"`
	DerivtnId string  `json:"DerivtnId"`
}

// KeyTransport4 ...
type KeyTransport4 struct {
	Vrsn          float64                    `json:"Vrsn"`
	RcptId        *Recipient5Choice          `json:"RcptId"`
	KeyNcrptnAlgo *AlgorithmIdentification11 `json:"KeyNcrptnAlgo"`
	NcrptdKey     string                     `json:"NcrptdKey"`
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
	Nm  string          `json:"Nm"`
	Adr *PostalAddress1 `json:"Adr"`
}

// Number ...
type Number float64

// OutputFormat1Code ...
type OutputFormat1Code string

// Parameter4 ...
type Parameter4 struct {
	NcrptnFrmt   string                     `json:"NcrptnFrmt"`
	DgstAlgo     string                     `json:"DgstAlgo"`
	MskGnrtrAlgo *AlgorithmIdentification12 `json:"MskGnrtrAlgo"`
}

// Parameter5 ...
type Parameter5 struct {
	DgstAlgo string `json:"DgstAlgo"`
}

// Parameter6 ...
type Parameter6 struct {
	NcrptnFrmt   string `json:"NcrptnFrmt"`
	InitlstnVctr string `json:"InitlstnVctr"`
	BPddg        string `json:"BPddg"`
}

// Parameter7 ...
type Parameter7 struct {
	InitlstnVctr string `json:"InitlstnVctr"`
	BPddg        string `json:"BPddg"`
}

// PartyIdentification72Choice ...
type PartyIdentification72Choice struct {
	AnyBIC  string                  `json:"AnyBIC"`
	PrtryId *GenericIdentification1 `json:"PrtryId"`
}

// PartyType10Code ...
type PartyType10Code string

// PartyType11Code ...
type PartyType11Code string

// PartyType9Code ...
type PartyType9Code string

// PaymentCard13 ...
type PaymentCard13 struct {
	PrtctdCardData *ContentInformationType10 `json:"PrtctdCardData"`
	PlainCardData  *PlainCardData9           `json:"PlainCardData"`
	MskdPAN        string                    `json:"MskdPAN"`
	CardPdctTp     string                    `json:"CardPdctTp"`
	CardPdctNm     string                    `json:"CardPdctNm"`
}

// PlainCardData9 ...
type PlainCardData9 struct {
	PAN       string        `json:"PAN"`
	CardSeqNb string        `json:"CardSeqNb"`
	FctvDt    string        `json:"FctvDt"`
	XpryDt    string        `json:"XpryDt"`
	SvcCd     string        `json:"SvcCd"`
	TrckData  []*TrackData1 `json:"TrckData"`
}

// PlusOrMinusIndicator ...
type PlusOrMinusIndicator bool

// PostalAddress1 ...
type PostalAddress1 struct {
	AdrTp       string   `json:"AdrTp"`
	AdrLine     []string `json:"AdrLine"`
	StrtNm      string   `json:"StrtNm"`
	BldgNb      string   `json:"BldgNb"`
	PstCd       string   `json:"PstCd"`
	TwnNm       string   `json:"TwnNm"`
	CtrySubDvsn string   `json:"CtrySubDvsn"`
	Ctry        string   `json:"Ctry"`
}

// PostalAddress18 ...
type PostalAddress18 struct {
	AdrLine     []string `json:"AdrLine"`
	StrtNm      string   `json:"StrtNm"`
	BldgNb      string   `json:"BldgNb"`
	PstCd       string   `json:"PstCd"`
	TwnNm       string   `json:"TwnNm"`
	CtrySubDvsn []string `json:"CtrySubDvsn"`
	Ctry        string   `json:"Ctry"`
}

// Recipient4Choice ...
type Recipient4Choice struct {
	KeyTrnsprt *KeyTransport4  `json:"KeyTrnsprt"`
	KEK        *KEK4           `json:"KEK"`
	KeyIdr     *KEKIdentifier2 `json:"KeyIdr"`
}

// Recipient5Choice ...
type Recipient5Choice struct {
	IssrAndSrlNb *IssuerAndSerialNumber1 `json:"IssrAndSrlNb"`
	KeyIdr       *KEKIdentifier2         `json:"KeyIdr"`
}

// RelativeDistinguishedName1 ...
type RelativeDistinguishedName1 struct {
	AttrTp  string `json:"AttrTp"`
	AttrVal string `json:"AttrVal"`
}

// Response3Code ...
type Response3Code string

// ResponseType2 ...
type ResponseType2 struct {
	Rslt         string `json:"Rslt"`
	RsltDtls     string `json:"RsltDtls"`
	AddtlRsltInf string `json:"AddtlRsltInf"`
}

// ResultDetail1Code ...
type ResultDetail1Code string

// Traceability3 ...
type Traceability3 struct {
	RlayId      *GenericIdentification74 `json:"RlayId"`
	TracDtTmIn  string                   `json:"TracDtTmIn"`
	TracDtTmOut string                   `json:"TracDtTmOut"`
}

// TrackData1 ...
type TrackData1 struct {
	TrckNb  string `json:"TrckNb"`
	TrckVal string `json:"TrckVal"`
}

// TransactionIdentifier2 ...
type TransactionIdentifier2 struct {
	RcncltnDt string `json:"RcncltnDt"`
	RcncltnId string `json:"RcncltnId"`
}

// TransactionVerificationResult4 ...
type TransactionVerificationResult4 struct {
	Mtd        string `json:"Mtd"`
	VrfctnNtty string `json:"VrfctnNtty"`
	Rslt       string `json:"Rslt"`
	AddtlRslt  string `json:"AddtlRslt"`
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
