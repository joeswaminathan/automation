package hl7
import (
    "reflect"
    "log"
    "fmt"
)
var table_0018 []string = []string{
    "AA",
    "MD",
    "PhD",
}

var table_0201 []string = []string{
    "PRN",
    "WPN",
}

var table_0207 []string = []string{
    "a",
    "i",
    "r",
}

var table_0445 []string = []string{
    "ISS",
    "NS",
}

var table_0061 []string = []string{
    "BCV",
    "ISO",
}

var table_0203 []string = []string{
    "MR",
    "SS",
    "NI",
    "MC",
}

var table_0429 []string = []string{
    "Y",
    "N",
}

var table_0618 []string = []string{
    "P",
    "U",
}

var table_0104 []string = []string{
    "2",
    "2.0D",
    "2.1",
    "2.2",
    "2.3",
    "2.6",
    "2.4",
    "2.5",
    "2.5.1",
    "2.6",
}

var table_0204 []string = []string{
    "L",
    "D",
    "A",
}

var table_0300 []string = []string{
    "ISO",
}

var table_0212 []string = []string{
    "US",
    "CA",
    "MX",
}

var table_0399 []string = []string{
    "USA",
    "CAN",
    "MEX",
}

var table_0616 []string = []string{
    "EX",
    "ED",
}

var table_0617 []string = []string{
    "H",
    "W",
    "O",
}

var table_0002 []string = []string{
    "A",
    "D",
    "M",
    "S",
    "W",
    "C",
    "P",
    "U",
}

var table_0003 []string = []string{
    "A01",
    "A02",
    "A03",
    "A04",
    "A05",
    "A06",
    "A07",
    "A08",
    "A09",
    "A10",
    "A11",
    "A12",
    "A13",
    "A14",
    "A15",
    "A16",
    "A17",
    "A18",
    "A19",
    "A20",
    "A21",
    "A22",
    "A23",
    "A24",
    "A25",
    "A26",
    "A27",
    "A28",
    "A29",
    "A30",
    "A31",
    "A32",
    "A33",
    "A34",
    "A35",
    "A36",
    "A37",
    "A38",
    "A39",
    "A40",
    "A41",
    "A42",
    "A43",
    "A44",
    "A45",
    "A46",
    "A47",
    "A48",
    "A49",
    "A50",
    "A51",
    "C01",
    "C02",
    "C03",
    "C04",
    "C05",
    "C06",
    "C07",
    "C08",
    "C09",
    "C10",
    "C11",
    "C12",
    "CNQ",
    "I01",
    "I02",
    "I03",
    "I04",
    "I05",
    "I06",
    "I07",
    "I08",
    "I09",
    "I10",
    "I11",
    "I12",
    "I13",
    "I14",
    "I15",
    "M01",
    "M02",
    "M03",
    "M04",
    "M05",
    "M06",
    "M07",
    "M08",
    "M09",
    "M10",
    "M11",
    "O01",
    "O02",
    "P01",
    "P02",
    "P03",
    "P04",
    "P05",
    "P06",
    "P07",
    "P08",
    "P09",
    "PC1",
    "PC2",
    "PC3",
    "PC4",
    "PC5",
    "PC6",
    "PC7",
    "PC8",
    "PC9",
    "PCA",
    "PCB",
    "PCC",
    "PCD",
    "PCE",
    "PCF",
    "PCG",
    "PCH",
    "PCJ",
    "PCK",
    "PCL",
    "Q01",
    "Q02",
    "Q03",
    "Q04",
    "Q05",
    "Q06",
    "Q07",
    "Q08",
    "Q09",
    "R01",
    "R02",
    "R03",
    "R04",
    "R05",
    "R06",
    "R07",
    "R08",
    "R09",
    "R0R",
    "RAR",
    "RDR",
    "RER",
    "RGR",
    "S01",
    "S02",
    "S03",
    "S04",
    "S05",
    "S06",
    "S07",
    "S08",
    "S09",
    "S10",
    "S11",
    "S12",
    "S13",
    "S14",
    "S15",
    "S16",
    "S17",
    "S18",
    "S19",
    "S20",
    "S21",
    "S22",
    "S23",
    "S24",
    "S25",
    "S26",
    "T01",
    "T02",
    "T03",
    "T04",
    "T05",
    "T06",
    "T07",
    "T08",
    "T09",
    "T10",
    "T11",
    "T12",
    "V01",
    "V02",
    "V03",
    "V04",
    "W01",
    "W02",
}

var table_0306 []string = []string{
    "B",
    "R",
}

var table_0448 []string = []string{
    "IDE",
    "PRF",
}

var table_0001 []string = []string{
    "F",
    "M",
    "O",
    "U",
}

var table_0211 []string = []string{
    "8859/1",
    "UNICODE",
    "ASCII",
}

var table_0301 []string = []string{
    "DNS",
}

var table_0004 []string = []string{
    "E",
    "I",
    "O",
    "P",
    "R",
    "B",
    "C",
    "U",
}

var table_0076 []string = []string{
    "ACK",
    "ADR",
    "ADT",
    "ARD",
    "BAR",
    "CRM",
    "CSU",
    "DFT",
    "DOC",
    "DSR",
    "EDR",
    "EQQ",
    "ERP",
    "MCF",
    "MDM",
    "MFD",
    "MFK",
    "MFN",
    "MFQ",
    "MFR",
    "NMD",
    "NMQ",
    "NMR",
    "ORF",
    "ORM",
    "ORR",
    "ORU",
    "OSQ",
    "OSR",
    "PEX",
    "PGL",
    "PIN",
    "PPG",
    "PPP",
    "PPR",
    "PPT",
    "PPV",
    "PRR",
    "PBR",
    "PTR",
    "QCK",
    "QRY",
    "RAR",
    "RAS",
    "RCI",
    "RCL",
    "RDE",
    "RDR",
    "RDS",
    "REF",
    "RER",
    "RGR",
    "RGV",
    "ROR",
    "RPA",
    "RPI",
    "RPL",
    "RPR",
    "RQA",
    "RQC",
    "RQI",
    "RQP",
    "RQQ",
    "RRA",
    "RRD",
    "RRE",
    "RRG",
    "RRI",
    "SIU",
    "SPQ",
    "SQM",
    "SQR",
    "SRM",
    "SRR",
    "SUR",
    "TBR",
    "UDM",
    "VQQ",
    "VXQ",
    "VXR",
    "VXU",
    "VXX",
}

var table_0202 []string = []string{
    "PH",
    "FX",
}

var table_0447 []string = []string{
    "LAB",
    "GER",
}

var table_0103 []string = []string{
    "D",
    "P",
    "T",
}

var table_0125 []string = []string{
    "AD",
    "CF",
    "CK",
    "CN",
    "CP",
    "CWE",
    "CX",
    "DT",
    "DTM",
    "ED",
    "FT",
    "MO",
    "NM",
    "PN",
    "RP",
    "SN",
    "ST",
    "TM",
    "TN",
    "TX",
    "XAD",
    "XCN",
    "XON",
    "XPN",
    "XTN",
}

var table_0288 []string = []string{
    "001",
    "002",
}

var table_0289 []string = []string{
    "001",
    "002",
}

var table_0354 []string = []string{
    "ADT_A01",
    "ADT_A02",
    "ADT_A03",
    "ADT_A06",
    "ADT_A09",
    "ADT_A12",
    "ADT_A16",
    "ADT_A17",
    "ADT_A18",
    "ADT_A20",
    "ADT_A24",
    "ADT_A28",
    "ADT_A30",
    "ADT_A37",
    "ADT_A38",
    "ADT_A39",
    "ADT_A43",
    "ADT_A45",
    "ADT_A50",
    "ARD_A19",
    "BAR_P01",
    "BAR_P02",
    "BAR_P06",
    "CRM_C01",
    "CSU_C09",
    "DFT_P03",
    "DOC_T12",
    "DSR_Q01",
    "DSR_Q03",
    "EDR_R07",
    "EQQ_Q04",
    "ERP_R09",
    "MDM_T01",
    "MDM_T02",
    "MFD_P09",
    "MFK_M01",
    "MFN_M01",
    "MFN_M02",
    "MFN_M03",
    "MFN_M05",
    "MFN_M06",
    "MFN_M07",
    "MFN_M08",
    "MFN_M09",
    "MFN_M10",
    "MFN_M11",
    "ORF_R02",
    "ORM__O01",
    "ORM_Q06",
    "ORR_O02",
    "ORR_Q06",
    "ORU_R01",
    "ORU_W01",
    "OSQ_Q06",
    "OSR_Q06",
    "PEX_P07",
    "PGL_PC6",
    "PIN_107",
    "PPG_PCG",
    "PPP_PCB",
    "PPR_PC1",
    "PPT_PCL",
    "PPV_PCA",
    "PRR_PC5",
    "PTR_PCF",
    "QCK_Q02",
    "QRY_A19",
    "QRY_PC4",
    "QRY_Q01",
    "QRY_Q02",
    "QRY_R02",
    "QRY_T12",
    "RAR_RAR",
    "RAS_O01",
    "RAS_O02",
    "RCI_I05",
    "RCL_I06",
    "RDE_O01",
    "RDR_RDR",
    "RDS_O01",
    "REF_I12",
    "RER_RER",
    "RGR_RGR",
    "RGV_O01",
    "RPA_I08",
    "RPI_I0I",
    "RPL_I02",
    "RPR_I03",
    "RQA_I08",
    "RQC_I05",
    "RQC_I06",
    "RQI_I0I",
    "RQP_I04",
    "RQQ_Q09",
    "RRA_O02",
    "RRD_O02",
    "RRE_O01",
    "RRG_O02",
    "RRI_I12",
    "RROR_ROR",
    "SIIU_S12",
    "SPQ_Q08",
    "SQM_S25",
    "SQR_S25",
    "SRM_S01",
    "SRM_T12",
    "SRR_S01",
    "SRR_T12",
    "SUR_P09",
    "TBR_R09",
    "UDM_Q05",
    "VQQ_Q07",
    "VXQ_V01",
    "VXR_V03",
    "VXU_V04",
    "VXX_V02",
    "",
}

var table_0363 []string = []string{
    "ISO",
    "HL7",
}

var table_0444 []string = []string{
    "F",
    "G",
}

var table_0446 []string = []string{
    "CANINE",
    "FELINE",
}

var table_0465 []string = []string{
    "A",
    "N",
}

func findTable(tn string) []string {
    switch tn {
    case "0018":
        return table_0018
    case "0201":
        return table_0201
    case "0207":
        return table_0207
    case "0445":
        return table_0445
    case "0061":
        return table_0061
    case "0203":
        return table_0203
    case "0429":
        return table_0429
    case "0618":
        return table_0618
    case "0104":
        return table_0104
    case "0204":
        return table_0204
    case "0300":
        return table_0300
    case "0212":
        return table_0212
    case "0399":
        return table_0399
    case "0616":
        return table_0616
    case "0617":
        return table_0617
    case "0002":
        return table_0002
    case "0003":
        return table_0003
    case "0306":
        return table_0306
    case "0448":
        return table_0448
    case "0001":
        return table_0001
    case "0211":
        return table_0211
    case "0301":
        return table_0301
    case "0004":
        return table_0004
    case "0076":
        return table_0076
    case "0202":
        return table_0202
    case "0447":
        return table_0447
    case "0103":
        return table_0103
    case "0125":
        return table_0125
    case "0288":
        return table_0288
    case "0289":
        return table_0289
    case "0354":
        return table_0354
    case "0363":
        return table_0363
    case "0444":
        return table_0444
    case "0446":
        return table_0446
    case "0465":
        return table_0465
    default:
        return nil
    }
}

type NM  float64


type ID  string


type MO struct {
    Quantity NM `json:"quantity,omitempty" hl7:"part=1 length=16"`
    Denomination ID `json:"denomination,omitempty" hl7:"part=2 length=3 table=ISO4217"`
}

type ST  string


type SI  uint


type IS  string


type EI struct {
    EntityIdentifier ST `json:"entityIdentifier,omitempty" hl7:"part=1 length=199"`
    NamespaceId IS `json:"namespaceId,omitempty" hl7:"part=2 length=20"`
    UniversalId ST `json:"universalId,omitempty" hl7:"part=3 length=199"`
    UniversalIdType ID `json:"universalIdType,omitempty" hl7:"part=4 length=6"`
}

type CWE struct {
    Identifier ST `json:"identifier,omitempty" hl7:"part=1 length=-1"`
    Text ST `json:"text,omitempty" hl7:"part=2 length=-1"`
    NameOfCodingSystem ST `json:"nameOfCodingSystem,omitempty" hl7:"part=3 length=-1"`
    AlternateIdentifier ST `json:"alternateIdentifier,omitempty" hl7:"part=4 length=-1"`
    AlternateText ST `json:"alternateText,omitempty" hl7:"part=5 length=-1"`
    NameOfAlternateCodingSystem ST `json:"nameOfAlternateCodingSystem,omitempty" hl7:"part=6 length=-1"`
    CodingSystemVersionId ST `json:"codingSystemVersionId,omitempty" hl7:"part=7 length=-1"`
    AlternateCodingSystemVersionId ST `json:"alternateCodingSystemVersionId,omitempty" hl7:"part=8 length=-1"`
    OriginalText ST `json:"originalText,omitempty" hl7:"part=9 length=-1"`
}

type DTM  string


type CQ struct {
    Quantity NM `json:"quantity,omitempty" hl7:"part=1 length=-1"`
    Units *CWE `json:"units,omitempty" hl7:"part=2 length=705"`
}

type FN struct {
    Surname ST `json:"surname,omitempty" hl7:"part=1 length=50"`
    OwnSurnamePrefix ST `json:"ownSurnamePrefix,omitempty" hl7:"part=2 length=20"`
    OwnSurname ST `json:"ownSurname,omitempty" hl7:"part=3 length=50"`
    SurnamePrefixFromPartnerSpouse ST `json:"surnamePrefixFromPartnerSpouse,omitempty" hl7:"part=4 length=20"`
    SurnameFromPartnerSpouse ST `json:"surnameFromPartnerSpouse,omitempty" hl7:"part=5 length=50"`
}

type HD struct {
    NamespaceId IS `json:"namespaceId,omitempty" hl7:"part=1 length=-1 table=0300"`
    UniversalId ST `json:"universalId,omitempty" hl7:"part=2 length=-1"`
    UniversalIdType ID `json:"universalIdType,omitempty" hl7:"part=3 length=-1 table=0301"`
}

type XCN struct {
    IdNumber ST `json:"idNumber,omitempty" hl7:"part=1 length=-1"`
    FamilyName *FN `json:"familyName,omitempty" hl7:"part=2 length=194"`
    GivenName ST `json:"givenName,omitempty" hl7:"part=3 length=30"`
    SecondNames ST `json:"secondNames,omitempty" hl7:"part=4 length=30"`
    Suffix ST `json:"suffix,omitempty" hl7:"part=5 length=20"`
    Prefix ST `json:"prefix,omitempty" hl7:"part=6 length=20"`
    Degree IS `json:"degree,omitempty" hl7:"part=7 length=6 table=0360"`
    SourceTable ID `json:"sourceTable,omitempty" hl7:"part=8 length=-1"`
    AssigningAuthority *HD `json:"assigningAuthority,omitempty" hl7:"part=9 length=227"`
    NameTypeCode ID `json:"nameTypeCode,omitempty" hl7:"part=10 length=1 table=0200"`
    IdentifierCheckDigit ST `json:"identifierCheckDigit,omitempty" hl7:"part=11 length=-1"`
    CheckDigitScheme ID `json:"checkDigitScheme,omitempty" hl7:"part=12 length=3 table=0061"`
    AssigningFacility *HD `json:"assigningFacility,omitempty" hl7:"part=13 length=227"`
}

type TX  string


type SPS struct {
    SpecimenSourceNameOrCode *CWE `json:"specimenSourceNameOrCode,omitempty" hl7:"part=1 length=705"`
    Additives *CWE `json:"additives,omitempty" hl7:"part=2 length=705 table=0371"`
    SpecimenCollectionMethod TX `json:"specimenCollectionMethod,omitempty" hl7:"part=3 length=200"`
    BodySite *CWE `json:"bodySite,omitempty" hl7:"part=4 length=705 table=0163"`
    SiteModifier *CWE `json:"siteModifier,omitempty" hl7:"part=5 length=705 table=0495"`
    CollectionMethodModifierCode *CWE `json:"collectionMethodModifierCode,omitempty" hl7:"part=6 length=705"`
    SpecimenRole *CWE `json:"specimenRole,omitempty" hl7:"part=7 length=705 table=0369"`
}

type TN  string


type XTN struct {
    TelephoneNumber TN `json:"telephoneNumber,omitempty" hl7:"part=1 length=-1"`
    TelecommunicationUseCode ID `json:"telecommunicationUseCode,omitempty" hl7:"part=2 length=-1 table=0201"`
    TelecommunicationEquipmentType ID `json:"telecommunicationEquipmentType,omitempty" hl7:"part=3 length=-1 table=0202"`
    EmailAddress ST `json:"emailAddress,omitempty" hl7:"part=4 length=-1"`
    CountryCode NM `json:"countryCode,omitempty" hl7:"part=5 length=-1"`
    AreaCode NM `json:"areaCode,omitempty" hl7:"part=6 length=-1"`
    PhoneNumber NM `json:"phoneNumber,omitempty" hl7:"part=7 length=-1"`
    Extension NM `json:"extension,omitempty" hl7:"part=8 length=-1"`
    AnyText ST `json:"anyText,omitempty" hl7:"part=9 length=-1"`
}

type MOC struct {
    MonetaryAmount *MO `json:"monetaryAmount,omitempty" hl7:"part=1 length=20"`
    ChargeCode *CWE `json:"chargeCode,omitempty" hl7:"part=2 length=705"`
}

type PRL struct {
    ParentObservationIdentifier *CWE `json:"parentObservationIdentifier,omitempty" hl7:"part=1 length=705"`
    ParentObservationSubIdentifier ST `json:"parentObservationSubIdentifier,omitempty" hl7:"part=2 length=20"`
    ParentObservationValueDescriptor TX `json:"parentObservationValueDescriptor,omitempty" hl7:"part=3 length=250"`
}

type RI struct {
    RepeatPattern IS `json:"repeatPattern,omitempty" hl7:"part=1 length=6 table=0335"`
    ExplicitTimeInterval ST `json:"explicitTimeInterval,omitempty" hl7:"part=2 length=199"`
}

type OSD struct {
    SequenceResultsFlag ID `json:"sequenceResultsFlag,omitempty" hl7:"part=1 length=1 table=0524"`
    PlacerOrderNumberEntityIdentifier ST `json:"placerOrderNumberEntityIdentifier,omitempty" hl7:"part=2 length=15"`
    PlacerOrderNumberNamespaceId ID `json:"placerOrderNumberNamespaceId,omitempty" hl7:"part=3 length=6 table=0363"`
    FillerOrderNumberEntityIdentifier ST `json:"fillerOrderNumberEntityIdentifier,omitempty" hl7:"part=4 length=15"`
    FillerOrderNumberNamespaceId ID `json:"fillerOrderNumberNamespaceId,omitempty" hl7:"part=5 length=6 table=0363"`
    SequenceConditionValue ST `json:"sequenceConditionValue,omitempty" hl7:"part=6 length=12"`
    MaximumNumberOfRepeats NM `json:"maximumNumberOfRepeats,omitempty" hl7:"part=7 length=3"`
    PlacerOrderNumberUniversalId ST `json:"placerOrderNumberUniversalId,omitempty" hl7:"part=8 length=15"`
    PlacerOrderNumberUniversalIdType ID `json:"placerOrderNumberUniversalIdType,omitempty" hl7:"part=9 length=6 table=0301"`
    FillerOrderNumberUniversalId ST `json:"fillerOrderNumberUniversalId,omitempty" hl7:"part=10 length=15"`
    FillerOrderNumberUniversalIdType ID `json:"fillerOrderNumberUniversalIdType,omitempty" hl7:"part=11 length=6 table=0301"`
}

type CE struct {
    Identifier ST `json:"identifier,omitempty" hl7:"part=1"`
    Text ST `json:"text,omitempty" hl7:"part=2"`
    NameOfCodingSystem ID `json:"nameOfCodingSystem,omitempty" hl7:"part=3 table=0396"`
    AlternateIdentifier ST `json:"alternateIdentifier,omitempty" hl7:"part=4"`
    AlternateText ST `json:"alternateText,omitempty" hl7:"part=5"`
    NameOfAlternateCodingSystem ID `json:"nameOfAlternateCodingSystem,omitempty" hl7:"part=6 table=0396"`
}

type TQ struct {
    Quantity *CQ `json:"quantity,omitempty" hl7:"part=1 length=267"`
    Interval *RI `json:"interval,omitempty" hl7:"part=2 length=206"`
    Duration ST `json:"duration,omitempty" hl7:"part=3 length=6"`
    StartDateTime DTM `json:"startDateTime,omitempty" hl7:"part=4 length=24"`
    EndDateTime DTM `json:"endDateTime,omitempty" hl7:"part=5 length=24"`
    Priority ST `json:"priority,omitempty" hl7:"part=6 length=6"`
    Condition ST `json:"condition,omitempty" hl7:"part=7 length=199"`
    Text TX `json:"text,omitempty" hl7:"part=8 length=200"`
    Conjunction ID `json:"conjunction,omitempty" hl7:"part=9 length=1 table=0472"`
    OrderSequencing *OSD `json:"orderSequencing,omitempty" hl7:"part=10 length=110"`
    OccurrenceDuration *CE `json:"occurrenceDuration,omitempty" hl7:"part=11 length=705"`
    TotalOccurrences NM `json:"totalOccurrences,omitempty" hl7:"part=12 length=4"`
}

type EIP struct {
    Placer *EI `json:"placer,omitempty" hl7:"part=1 length=427"`
    Filler *EI `json:"filler,omitempty" hl7:"part=2 length=427"`
}

type CNN struct {
    IdNumber ST `json:"idNumber,omitempty" hl7:"part=1 length=15"`
    FamilyName ST `json:"familyName,omitempty" hl7:"part=2 length=50"`
    GivenName ST `json:"givenName,omitempty" hl7:"part=3 length=30"`
    SecondAndFurtherGivenNamesOrInitialsThereof ST `json:"secondAndFurtherGivenNamesOrInitialsThereof,omitempty" hl7:"part=4 length=30"`
    Suffix ST `json:"suffix,omitempty" hl7:"part=5 length=20"`
    Prefix ST `json:"prefix,omitempty" hl7:"part=6 length=20"`
    Degree IS `json:"degree,omitempty" hl7:"part=7 length=5 table=0360"`
    SourceTable IS `json:"sourceTable,omitempty" hl7:"part=8 length=4 table=0297"`
    AssigningAuthorityNamespaceId IS `json:"assigningAuthorityNamespaceId,omitempty" hl7:"part=9 length=20 table=0363"`
    AssigningAuthorityUniversalId ST `json:"assigningAuthorityUniversalId,omitempty" hl7:"part=10 length=199"`
    AssigningAuthorityUniversalIdType ID `json:"assigningAuthorityUniversalIdType,omitempty" hl7:"part=11 length=6 table=0301"`
}

type NDL struct {
    Name *CNN `json:"name,omitempty" hl7:"part=1 length=406"`
    StartDateTime DTM `json:"startDateTime,omitempty" hl7:"part=2 length=24"`
    EndDateTime DTM `json:"endDateTime,omitempty" hl7:"part=3 length=24"`
    PointOfCare IS `json:"pointOfCare,omitempty" hl7:"part=4 length=20 table=0302"`
    Room IS `json:"room,omitempty" hl7:"part=5 length=20 table=0303"`
    Bed IS `json:"bed,omitempty" hl7:"part=6 length=20 table=0304"`
    Facility *HD `json:"facility,omitempty" hl7:"part=7 length=227"`
    LocationStatus IS `json:"locationStatus,omitempty" hl7:"part=8 length=20 table=0306"`
    PatientLocationType IS `json:"patientLocationType,omitempty" hl7:"part=9 length=20 table=0305"`
    Building IS `json:"building,omitempty" hl7:"part=10 length=20 table=0307"`
    Floor IS `json:"floor,omitempty" hl7:"part=11 length=20 table=0308"`
}

type CNE struct {
    Identifier ST `json:"identifier,omitempty" hl7:"part=1 length=20"`
    Text ST `json:"text,omitempty" hl7:"part=2 length=199"`
    NameOfCodingSystem ID `json:"nameOfCodingSystem,omitempty" hl7:"part=3 length=20 table=0396"`
    AlternateIdentifier ST `json:"alternateIdentifier,omitempty" hl7:"part=4 length=20"`
    AlternateText ST `json:"alternateText,omitempty" hl7:"part=5 length=199"`
    NameOfAlternateCodingSystem ID `json:"nameOfAlternateCodingSystem,omitempty" hl7:"part=6 length=20 table=0396"`
    CodingSystemVersionId ST `json:"codingSystemVersionId,omitempty" hl7:"part=7 length=10"`
    AlternateCodingSystemVersionId ST `json:"alternateCodingSystemVersionId,omitempty" hl7:"part=8 length=10"`
    OriginalText ST `json:"originalText,omitempty" hl7:"part=9 length=199"`
}

type OBR struct {
    Hl7SegmentName ST `json:"hl7SegmentName,omitempty" hl7:"part=1 length=3 table=0076"`
    SetId SI `json:"setId,omitempty" hl7:"part=2 length=4"`
    PlacerOrderNumber *EI `json:"placerOrderNumber,omitempty" hl7:"part=3 length=427"`
    FillerOrderNumber *EI `json:"fillerOrderNumber,omitempty" hl7:"part=4 length=427"`
    UniversalServiceIdentifier *CWE `json:"universalServiceIdentifier,omitempty" hl7:"part=5 length=705"`
    Priority ID `json:"priority,omitempty" hl7:"part=6 length=2"`
    RequestedDateTime DTM `json:"requestedDateTime,omitempty" hl7:"part=7 length=24"`
    ObservationDateTime DTM `json:"observationDateTime,omitempty" hl7:"part=8 length=24"`
    ObservationEndDateTime DTM `json:"observationEndDateTime,omitempty" hl7:"part=9 length=24"`
    CollectionVolume *CQ `json:"collectionVolume,omitempty" hl7:"part=10 length=722"`
    CollectorIdentifier []*XCN `json:"collectorIdentifier,omitempty" hl7:"part=11 length=3220"`
    SpecimenActionCode ID `json:"specimenActionCode,omitempty" hl7:"part=12 length=1 table=0065"`
    DangerCode *CWE `json:"dangerCode,omitempty" hl7:"part=13 length=705"`
    RelevantClinicalInformation ST `json:"relevantClinicalInformation,omitempty" hl7:"part=14 length=300"`
    SpecimenReceivedDateTime DTM `json:"specimenReceivedDateTime,omitempty" hl7:"part=15 length=24"`
    SpecimenSource *SPS `json:"specimenSource,omitempty" hl7:"part=16 length=300"`
    OrderingProvider []*XCN `json:"orderingProvider,omitempty" hl7:"part=17 length=3220"`
    OrderCallbackPhoneNumber []*XTN `json:"orderCallbackPhoneNumber,omitempty" hl7:"part=18 length=2743"`
    PlacerField1 ST `json:"placerField1,omitempty" hl7:"part=19 length=199"`
    PlacerField2 ST `json:"placerField2,omitempty" hl7:"part=20 length=199"`
    FillerField1 ST `json:"fillerField1,omitempty" hl7:"part=21 length=199"`
    FillerField2 ST `json:"fillerField2,omitempty" hl7:"part=22 length=199"`
    ResultsRptStatusChngDateTime DTM `json:"resultsRptStatusChngDateTime,omitempty" hl7:"part=23 length=24"`
    ChargeToPractice *MOC `json:"chargeToPractice,omitempty" hl7:"part=24 length=504"`
    DiagnosticServSectId ID `json:"diagnosticServSectId,omitempty" hl7:"part=25 length=10 table=0074"`
    ResultStatus ID `json:"resultStatus,omitempty" hl7:"part=26 length=1 table=0123"`
    ParentResult *PRL `json:"parentResult,omitempty" hl7:"part=27 length=977"`
    QuantityTiming []*TQ `json:"quantityTiming,omitempty" hl7:"part=28 length=705"`
    ResultCopiesTo []*XCN `json:"resultCopiesTo,omitempty" hl7:"part=29 length=3220"`
    Parent *EIP `json:"parent,omitempty" hl7:"part=30 length=855"`
    TransportationMode ID `json:"transportationMode,omitempty" hl7:"part=31 length=20 table=0124"`
    ReasonForStudy []*CWE `json:"reasonForStudy,omitempty" hl7:"part=32 length=705"`
    PrincipalResultInterpreter []*NDL `json:"principalResultInterpreter,omitempty" hl7:"part=33 length=831"`
    AssistantResultInterpreter []*NDL `json:"assistantResultInterpreter,omitempty" hl7:"part=34 length=831"`
    Technician []*NDL `json:"technician,omitempty" hl7:"part=35 length=831"`
    Transcriptionist []*NDL `json:"transcriptionist,omitempty" hl7:"part=36 length=831"`
    ScheduledDateTime DTM `json:"scheduledDateTime,omitempty" hl7:"part=37 length=24"`
    NumberOfSampleContainers NM `json:"numberOfSampleContainers,omitempty" hl7:"part=38 length=16"`
    TransportLogisticsOfCollectedSample []*CWE `json:"transportLogisticsOfCollectedSample,omitempty" hl7:"part=39 length=705"`
    CollectorsComment []*CWE `json:"collectorsComment,omitempty" hl7:"part=40 length=705"`
    TransportArrangementResponsibility *CWE `json:"transportArrangementResponsibility,omitempty" hl7:"part=41 length=705 table=0224"`
    TransportArranged ID `json:"transportArranged,omitempty" hl7:"part=42 length=30 table=0225"`
    EscortRequired ID `json:"escortRequired,omitempty" hl7:"part=43 length=1"`
    PlannedPatientTransportComment []*CWE `json:"plannedPatientTransportComment,omitempty" hl7:"part=44 length=705"`
    ProcedureCode *CNE `json:"procedureCode,omitempty" hl7:"part=45 length=705 table=0088"`
    ProcedureCodeModifier []*CNE `json:"procedureCodeModifier,omitempty" hl7:"part=46 length=705 table=0340"`
    PlacerSupplementalServiceInformation []*CWE `json:"placerSupplementalServiceInformation,omitempty" hl7:"part=47 length=705 table=0411"`
    FillerSupplementalServiceInformation []*CWE `json:"fillerSupplementalServiceInformation,omitempty" hl7:"part=48 length=705 table=0411"`
    MedicallyNecessaryDuplicateProcedureReason *CWE `json:"medicallyNecessaryDuplicateProcedureReason,omitempty" hl7:"part=49 length=705 table=0476"`
    ResultHandling IS `json:"resultHandling,omitempty" hl7:"part=50 length=2 table=0507"`
    ParentUniversalServiceIdentifier *CWE `json:"parentUniversalServiceIdentifier,omitempty" hl7:"part=51 length=705"`
}

type XON struct {
    OrganizationName ST `json:"organizationName,omitempty" hl7:"part=1 length=50"`
    OrganizationNameTypeCode IS `json:"organizationNameTypeCode,omitempty" hl7:"part=2 length=20 table=0204"`
    IdNumber ST `json:"idNumber,omitempty" hl7:"part=3 length=-1"`
    CheckDigitScheme ID `json:"checkDigitScheme,omitempty" hl7:"part=4 length=3 table=0061"`
    AssigningAuthority *HD `json:"assigningAuthority,omitempty" hl7:"part=5 length=227 table=0363"`
    IdentifierTypeCode ID `json:"identifierTypeCode,omitempty" hl7:"part=6 length=5"`
    AssigningFacility *HD `json:"assigningFacility,omitempty" hl7:"part=7 length=227"`
    NameRepresentationCode ID `json:"nameRepresentationCode,omitempty" hl7:"part=8 length=1 table=0465"`
    OrganizationIdentifier ST `json:"organizationIdentifier,omitempty" hl7:"part=9 length=20"`
}

type PL struct {
    PointOfCare ST `json:"pointOfCare,omitempty" hl7:"part=1 length=10"`
    Room ST `json:"room,omitempty" hl7:"part=2 length=10"`
    Bed ST `json:"bed,omitempty" hl7:"part=3 length=10"`
    Facility *HD `json:"facility,omitempty" hl7:"part=4 length=227"`
    LocationStatus IS `json:"locationStatus,omitempty" hl7:"part=5 length=2"`
    PatientLocationType IS `json:"patientLocationType,omitempty" hl7:"part=6 length=2"`
    Building ST `json:"building,omitempty" hl7:"part=7 length=20"`
    Floor ST `json:"floor,omitempty" hl7:"part=8 length=10"`
    LocationDescription ST `json:"locationDescription,omitempty" hl7:"part=9 length=20"`
}

type DT  string


type SN struct {
    Comparator ST `json:"comparator,omitempty" hl7:"part=1 length=2"`
    Num1 NM `json:"num1,omitempty" hl7:"part=2 length=15"`
    SeparatorSuffix ST `json:"separatorSuffix,omitempty" hl7:"part=3 length=2"`
    Num2 NM `json:"num2,omitempty" hl7:"part=4 length=15"`
}

type MSG struct {
    MessageType ID `json:"messageType,omitempty" hl7:"part=1 length=-1 table=0076"`
    TriggerEvent ID `json:"triggerEvent,omitempty" hl7:"part=2 length=-1 table=0003"`
    MessageStructure ID `json:"messageStructure,omitempty" hl7:"part=3 length=-1 table=0354"`
}

type ProcessingType struct {
    ProcessingId ID `json:"processingId,omitempty" hl7:"part=1 length=-1 table=0103"`
    ProcessingMode ID `json:"processingMode,omitempty" hl7:"part=2 length=-1 table=0207"`
}

type VID struct {
    VersionId ID `json:"versionId,omitempty" hl7:"part=1 length=-1 table=0104"`
    InternationalizationCode *CWE `json:"internationalizationCode,omitempty" hl7:"part=2 length=-1"`
    InternationalVersionId *CWE `json:"internationalVersionId,omitempty" hl7:"part=3 length=-1"`
}

type MSH struct {
    Hl7SegmentName ST `json:"hl7SegmentName,omitempty" hl7:"part=1 length=3 table=0076"`
    EncodingCharacters ST `json:"encodingCharacters,omitempty" hl7:"part=2 length=4"`
    SendingApplication *HD `json:"sendingApplication,omitempty" hl7:"part=3 length=180"`
    SendingFacility *HD `json:"sendingFacility,omitempty" hl7:"part=4 length=180"`
    ReceivingApplication *HD `json:"receivingApplication,omitempty" hl7:"part=5 length=180"`
    ReceivingFacility *HD `json:"receivingFacility,omitempty" hl7:"part=6 length=180"`
    DateOrTimeOfMessage []DTM `json:"dateOrTimeOfMessage,omitempty" hl7:"part=7 length=4"`
    Security ST `json:"security,omitempty" hl7:"part=8 length=40"`
    MessageType *MSG `json:"messageType,omitempty" hl7:"part=9 length=7"`
    MessageControlId ST `json:"messageControlId,omitempty" hl7:"part=10 length=20"`
    ProcessingId *ProcessingType `json:"processingId,omitempty" hl7:"part=11 length=3"`
    VersionId *VID `json:"versionId,omitempty" hl7:"part=12 length=60"`
    SequenceNumber NM `json:"sequenceNumber,omitempty" hl7:"part=13 length=15"`
    ContinuationPointer ST `json:"continuationPointer,omitempty" hl7:"part=14 length=180"`
    AcceptAcknowledgementType ID `json:"acceptAcknowledgementType,omitempty" hl7:"part=15 length=2 table=0155"`
    CountryCode ID `json:"countryCode,omitempty" hl7:"part=16 length=3 table=399"`
    CharacterSet []ID `json:"characterSet,omitempty" hl7:"part=17 length=16 table=0211"`
    PrincipalLanguageOfMessage *CWE `json:"principalLanguageOfMessage,omitempty" hl7:"part=18 length=250 table=0155"`
    AlternateCharacterSetHandling ID `json:"alternateCharacterSetHandling,omitempty" hl7:"part=19 length=20 table=0356"`
    MessageProfileIdentifier []*EI `json:"messageProfileIdentifier,omitempty" hl7:"part=20 length=427"`
    SendingResponseOrganization *XON `json:"sendingResponseOrganization,omitempty" hl7:"part=21 length=567"`
    ReceivingResponseOrganization *XON `json:"receivingResponseOrganization,omitempty" hl7:"part=22 length=567"`
    SendingNetworkAddress *HD `json:"sendingNetworkAddress,omitempty" hl7:"part=23 length=227"`
    ReceivingNetworkAddress *HD `json:"receivingNetworkAddress,omitempty" hl7:"part=24 length=227"`
}

type CX struct {
    IdNumber ST `json:"idNumber,omitempty" hl7:"part=1 length=15"`
    IdentifierCheckDigit ST `json:"identifierCheckDigit,omitempty" hl7:"part=2 length=4"`
    CheckDigitScheme ID `json:"checkDigitScheme,omitempty" hl7:"part=3 length=3 table=0061"`
    AssigningAuthority *HD `json:"assigningAuthority,omitempty" hl7:"part=4 length=227 table=0363"`
    IdentifierTypeCode ID `json:"identifierTypeCode,omitempty" hl7:"part=5 length=5 table=0203"`
    AssigningFacility *HD `json:"assigningFacility,omitempty" hl7:"part=6 length=227"`
    EffectiveDT DT `json:"effectiveDT,omitempty" hl7:"part=7 length=8"`
    ExpirationDT DT `json:"expirationDT,omitempty" hl7:"part=8 length=8"`
    AssigningJurisdiction *CWE `json:"assigningJurisdiction,omitempty" hl7:"part=9 length=705"`
    AssigningAgencyOrDepartment *CWE `json:"assigningAgencyOrDepartment,omitempty" hl7:"part=10 length=705"`
}

type DR struct {
    RangeStart DTM `json:"rangeStart,omitempty" hl7:"part=1 length=24"`
    RangeEnd DTM `json:"rangeEnd,omitempty" hl7:"part=2 length=24"`
}

type XPN struct {
    FamilyName *FN `json:"familyName,omitempty" hl7:"part=1 length=194"`
    GivenName ST `json:"givenName,omitempty" hl7:"part=2 length=30"`
    SecondName ST `json:"secondName,omitempty" hl7:"part=3 length=30"`
    Suffix ST `json:"suffix,omitempty" hl7:"part=4 length=20"`
    Prefix ST `json:"prefix,omitempty" hl7:"part=5 length=20"`
    Degree IS `json:"degree,omitempty" hl7:"part=6 length=6 table=0360"`
    NameTypeCode ID `json:"nameTypeCode,omitempty" hl7:"part=7 length=1 table=0200"`
    NameRepresentationCode ID `json:"nameRepresentationCode,omitempty" hl7:"part=8 length=1 table=0465"`
    NameContext *CWE `json:"nameContext,omitempty" hl7:"part=9 length=3 table=0448"`
    NameValidityRange *DR `json:"nameValidityRange,omitempty" hl7:"part=10 length=227"`
    NameAssembleOrder ID `json:"nameAssembleOrder,omitempty" hl7:"part=11 length=-1 table=0444"`
    EffectiveDT DTM `json:"effectiveDT,omitempty" hl7:"part=12 length=24"`
    ExpirationDT DTM `json:"expirationDT,omitempty" hl7:"part=13 length=24"`
    ProfessionalSuffix ST `json:"professionalSuffix,omitempty" hl7:"part=14 length=199"`
}

type SAD struct {
    StreetOrMailingAddress ST `json:"streetOrMailingAddress,omitempty" hl7:"part=1 length=120"`
    StreetName ST `json:"streetName,omitempty" hl7:"part=2 length=50"`
    DwellingNumber ST `json:"dwellingNumber,omitempty" hl7:"part=3 length=12"`
}

type XAD struct {
    StreetAddress *SAD `json:"streetAddress,omitempty" hl7:"part=1 length=184"`
    OtherDestination ST `json:"otherDestination,omitempty" hl7:"part=2 length=120"`
    City ST `json:"city,omitempty" hl7:"part=3 length=50"`
    StateOrProvince ST `json:"stateOrProvince,omitempty" hl7:"part=4 length=50"`
    ZipOrPostalCode ST `json:"zipOrPostalCode,omitempty" hl7:"part=5 length=12"`
    Country ID `json:"country,omitempty" hl7:"part=6 length=3 table=0399"`
    AddressType ID `json:"addressType,omitempty" hl7:"part=7 length=3 table=0190"`
    OtherGeographicDesignation ST `json:"otherGeographicDesignation,omitempty" hl7:"part=8 length=50"`
    CountyOrParishCode IS `json:"countyOrParishCode,omitempty" hl7:"part=9 length=20 table=0289"`
    CensusTract IS `json:"censusTract,omitempty" hl7:"part=10 length=20 table=0288"`
    AddressRepresentationCode ID `json:"addressRepresentationCode,omitempty" hl7:"part=11 length=1 table=0465"`
    AddressValidtyRange *DR `json:"addressValidtyRange,omitempty" hl7:"part=12 length=49"`
    EffectiveDT DTM `json:"effectiveDT,omitempty" hl7:"part=13 length=24"`
    ExpirationDT DTM `json:"expirationDT,omitempty" hl7:"part=14 length=24"`
    ExpirationReason *CWE `json:"expirationReason,omitempty" hl7:"part=15 length=705 table=0616"`
    TemporaryIndicator ID `json:"temporaryIndicator,omitempty" hl7:"part=16 length=1 table=0136"`
    BusAddressIndicator ID `json:"busAddressIndicator,omitempty" hl7:"part=17 length=1 table=0136"`
    AddressUsage ID `json:"addressUsage,omitempty" hl7:"part=18 length=44 table=0617"`
    Addresse ST `json:"addresse,omitempty" hl7:"part=19 length=199"`
    Comment ST `json:"comment,omitempty" hl7:"part=20 length=199"`
    PreferenceOrder NM `json:"preferenceOrder,omitempty" hl7:"part=21 length=2"`
    ProtectionCode *CWE `json:"protectionCode,omitempty" hl7:"part=22 length=705 table=0618"`
    AddressIdentifier *EI `json:"addressIdentifier,omitempty" hl7:"part=23 length=427"`
}

type DLN struct {
    DriverLicenseNumber ST `json:"driverLicenseNumber,omitempty" hl7:"part=1 length=-1"`
    IssuingStateProvinceCountry ST `json:"issuingStateProvinceCountry,omitempty" hl7:"part=2 length=-1"`
    ExpirationDT DTM `json:"expirationDT,omitempty" hl7:"part=3 length=-1"`
}

type PID struct {
    Hl7SegmentName ST `json:"hl7SegmentName,omitempty" hl7:"part=1 length=3 table=0076"`
    SetId NM `json:"setId,omitempty" hl7:"part=2 length=4"`
    PatientId *CX `json:"patientId,omitempty" hl7:"part=3 length=20"`
    PatientIdentifierList []*CX `json:"patientIdentifierList,omitempty" hl7:"part=4 length=250"`
    AlternatePatientId []*CX `json:"alternatePatientId,omitempty" hl7:"part=5 length=250"`
    PatientName []*XPN `json:"patientName,omitempty" hl7:"part=6 length=180"`
    MothersMaidenName []*XPN `json:"mothersMaidenName,omitempty" hl7:"part=7 length=180"`
    DateTimeOfBirth DTM `json:"dateTimeOfBirth,omitempty" hl7:"part=8 length=24"`
    AdministrativeSex IS `json:"administrativeSex,omitempty" hl7:"part=9 length=1 table=0001"`
    PatientAlias []*XPN `json:"patientAlias,omitempty" hl7:"part=10 length=250"`
    Race []*CWE `json:"race,omitempty" hl7:"part=11 length=705 table=0005"`
    PatientAddress []*XAD `json:"patientAddress,omitempty" hl7:"part=12 length=250"`
    CountyCode IS `json:"countyCode,omitempty" hl7:"part=13 length=4 table=0289"`
    PhoneNumberHome []*XTN `json:"phoneNumberHome,omitempty" hl7:"part=14 length=250"`
    PhoneNumberBusiness []*XTN `json:"phoneNumberBusiness,omitempty" hl7:"part=15 length=250"`
    PrimaryLanguage *CWE `json:"primaryLanguage,omitempty" hl7:"part=16 length=705 table=0296"`
    MaritalStatus *CWE `json:"maritalStatus,omitempty" hl7:"part=17 length=705 table=0002"`
    Religion *CWE `json:"religion,omitempty" hl7:"part=18 length=705 table=0006"`
    PatientAccountNumber *CX `json:"patientAccountNumber,omitempty" hl7:"part=19 length=250"`
    SocialSecurityNumber ST `json:"socialSecurityNumber,omitempty" hl7:"part=20 length=16"`
    DriversLicenseNumber *DLN `json:"driversLicenseNumber,omitempty" hl7:"part=21 length=25"`
    MothersIdentifier []*CX `json:"mothersIdentifier,omitempty" hl7:"part=22 length=250"`
    EthnicGroup []*CX `json:"ethnicGroup,omitempty" hl7:"part=23 length=1 table=0189"`
    BirthPlace ST `json:"birthPlace,omitempty" hl7:"part=24 length=250"`
    MultipleBirthIndicator ID `json:"multipleBirthIndicator,omitempty" hl7:"part=25 length=1 table=0136"`
    BirthOrder uint16 `json:"birthOrder,omitempty" hl7:"part=26 length=2"`
    Citizenship []*CWE `json:"citizenship,omitempty" hl7:"part=27 length=705 table=0171"`
    VeteransMilitaryStatus *CWE `json:"veteransMilitaryStatus,omitempty" hl7:"part=28 length=705 table=0172"`
    Nationality *CWE `json:"nationality,omitempty" hl7:"part=29 length=705 table=0212"`
    PatientDeathDTAndTime DTM `json:"patientDeathDTAndTime,omitempty" hl7:"part=30 length=24"`
    PatientDeathIndicator ID `json:"patientDeathIndicator,omitempty" hl7:"part=31 length=1 table=0136"`
    IdentityUnknownIndicator ID `json:"identityUnknownIndicator,omitempty" hl7:"part=32 length=1 table=0136"`
    IdentityReliabilityCode []IS `json:"identityReliabilityCode,omitempty" hl7:"part=33 length=20 table=0445"`
    LastUpdateDTM DTM `json:"lastUpdateDTM,omitempty" hl7:"part=34 length=24"`
    LastUpdateFacility *HD `json:"lastUpdateFacility,omitempty" hl7:"part=35 length=241"`
    SpeciesCode *CWE `json:"speciesCode,omitempty" hl7:"part=36 length=705 table=0446"`
    BreedCode *CWE `json:"breedCode,omitempty" hl7:"part=37 length=705 table=0447"`
    Strain ST `json:"strain,omitempty" hl7:"part=38 length=705"`
    ProductionClassCode []*CWE `json:"productionClassCode,omitempty" hl7:"part=39 length=705 table=0429"`
    TribalCitizenship []*CWE `json:"tribalCitizenship,omitempty" hl7:"part=40 length=705 table=0171"`
}

type Varies  any


type OBX struct {
    Hl7SegmentName ST `json:"hl7SegmentName,omitempty" hl7:"part=1 length=3 table=0076"`
    SetId SI `json:"setId,omitempty" hl7:"part=2 length=4"`
    ValueType ID `json:"valueType,omitempty" hl7:"part=3 length=3 table=0125"`
    ObservationIdentifier *CWE `json:"observationIdentifier,omitempty" hl7:"part=4 length=705"`
    ObservationSubId ST `json:"observationSubId,omitempty" hl7:"part=5 length=20"`
    ObservationValue []Varies `json:"observationValue,omitempty" hl7:"part=6 length=180 variedTypeField=ValueType"`
    Units *CWE `json:"units,omitempty" hl7:"part=7 length=705"`
    ReferenceRange ST `json:"referenceRange,omitempty" hl7:"part=8 length=60"`
    AbnormalFlags []IS `json:"abnormalFlags,omitempty" hl7:"part=9 length=5 table=0078"`
    Probability []NM `json:"probability,omitempty" hl7:"part=10 length=250"`
    NatureOfAbnormalTest []ID `json:"natureOfAbnormalTest,omitempty" hl7:"part=11 length=2 table=0080"`
    ObservationResultStatus ID `json:"observationResultStatus,omitempty" hl7:"part=12 length=1 table=0085"`
    EffectiveDataOfReferenceRange DTM `json:"effectiveDataOfReferenceRange,omitempty" hl7:"part=13 length=24"`
    UserDefinedAccessChecks ST `json:"userDefinedAccessChecks,omitempty" hl7:"part=14 length=20"`
    DateAndTimeOfObservation DTM `json:"dateAndTimeOfObservation,omitempty" hl7:"part=15 length=24"`
    ProducersId *CWE `json:"producersId,omitempty" hl7:"part=16 length=705"`
    ReponsibleObserver []*XCN `json:"reponsibleObserver,omitempty" hl7:"part=17 length=3220"`
    ObservationMethod []*CWE `json:"observationMethod,omitempty" hl7:"part=18 length=705"`
    EquipmentInstanceIdentifier *EI `json:"equipmentInstanceIdentifier,omitempty" hl7:"part=19 length=427"`
    MoodCode *CNE `json:"moodCode,omitempty" hl7:"part=20 length=705 table=0725"`
    PerformingOrganizationName *XON `json:"performingOrganizationName,omitempty" hl7:"part=21 length=570"`
    PerformingOrganizationAddress *XAD `json:"performingOrganizationAddress,omitempty" hl7:"part=22 length=2915"`
    PerformingOrganizationMedicalDirector *XCN `json:"performingOrganizationMedicalDirector,omitempty" hl7:"part=23 length=3220"`
}

type FC struct {
    FinancialClassCode IS `json:"financialClassCode,omitempty" hl7:"part=1 length=20 table=0064"`
    EffectiveDate DTM `json:"effectiveDate,omitempty" hl7:"part=2 length=24"`
}

type DLD struct {
    DischargeToLocation *CWE `json:"dischargeToLocation,omitempty" hl7:"part=1 length=705 table=0113"`
    EffectiveDate DTM `json:"effectiveDate,omitempty" hl7:"part=2 length=24"`
}

type PV1 struct {
    Hl7SegmentName ST `json:"hl7SegmentName,omitempty" hl7:"part=1 length=3 table=0076"`
    SetId SI `json:"setId,omitempty" hl7:"part=2 length=4"`
    PatientClass IS `json:"patientClass,omitempty" hl7:"part=3 length=1 table=0004"`
    AssignedPatientLocation *PL `json:"assignedPatientLocation,omitempty" hl7:"part=4 length=80"`
    AdmissionType IS `json:"admissionType,omitempty" hl7:"part=5 length=2 table=0007"`
    PreadmitNumber *CX `json:"preadmitNumber,omitempty" hl7:"part=6 length=250"`
    PriorPatientLocation *PL `json:"priorPatientLocation,omitempty" hl7:"part=7 length=80"`
    AttendingDoctor []*XCN `json:"attendingDoctor,omitempty" hl7:"part=8 length=250 table=0010"`
    ReferringDoctor []*XCN `json:"referringDoctor,omitempty" hl7:"part=9 length=250 table=0010"`
    ConsultingDoctor []*XCN `json:"consultingDoctor,omitempty" hl7:"part=10 length=250 table=0010"`
    HospitalService IS `json:"hospitalService,omitempty" hl7:"part=11 length=3 table=0069"`
    TemporaryLocation *PL `json:"temporaryLocation,omitempty" hl7:"part=12 length=80"`
    PreadmitTestIndicator IS `json:"preadmitTestIndicator,omitempty" hl7:"part=13 length=2 table=0087"`
    ReAdmissionIndicator IS `json:"reAdmissionIndicator,omitempty" hl7:"part=14 length=2 table=0092"`
    AdmitSource IS `json:"admitSource,omitempty" hl7:"part=15 length=6 table=0023"`
    AmbulatoryStatus []IS `json:"ambulatoryStatus,omitempty" hl7:"part=16 length=2 table=0009"`
    VipIndicator IS `json:"vipIndicator,omitempty" hl7:"part=17 length=2 table=0099"`
    AdmittingDoctor []*XCN `json:"admittingDoctor,omitempty" hl7:"part=18 length=250 table=0010"`
    PatientType IS `json:"patientType,omitempty" hl7:"part=19 length=2 table=0018"`
    VisitNumber *CX `json:"visitNumber,omitempty" hl7:"part=20 length=250"`
    FinancialClass []*FC `json:"financialClass,omitempty" hl7:"part=21 length=50"`
    ChargePriceIndicator IS `json:"chargePriceIndicator,omitempty" hl7:"part=22 length=2 table=0032"`
    CourtesyCode IS `json:"courtesyCode,omitempty" hl7:"part=23 length=2 table=0045"`
    CreditRating IS `json:"creditRating,omitempty" hl7:"part=24 length=2 table=0046"`
    ContractCode []IS `json:"contractCode,omitempty" hl7:"part=25 length=2 table=0044"`
    ContractEffectiveDate DT `json:"contractEffectiveDate,omitempty" hl7:"part=26 length=8"`
    ContractAmount NM `json:"contractAmount,omitempty" hl7:"part=27 length=12"`
    ContractPeriod NM `json:"contractPeriod,omitempty" hl7:"part=28 length=3"`
    InterestCode IS `json:"interestCode,omitempty" hl7:"part=29 length=2 table=0073"`
    TransferToBadDebtCode IS `json:"transferToBadDebtCode,omitempty" hl7:"part=30 length=4 table=0110"`
    TransferToBadDebtDate DT `json:"transferToBadDebtDate,omitempty" hl7:"part=31 length=8"`
    BadDebtAgencyCode IS `json:"badDebtAgencyCode,omitempty" hl7:"part=32 length=10 table=0021"`
    BadDebtTransferAmount []NM `json:"badDebtTransferAmount,omitempty" hl7:"part=33 length=12"`
    BadDebtRecoveryAmount []NM `json:"badDebtRecoveryAmount,omitempty" hl7:"part=34 length=12"`
    DeleteAccountIndicator IS `json:"deleteAccountIndicator,omitempty" hl7:"part=35 length=1 table=0111"`
    DeleteAccountDate DT `json:"deleteAccountDate,omitempty" hl7:"part=36 length=8"`
    DischargeDisposition IS `json:"dischargeDisposition,omitempty" hl7:"part=37 length=3 table=0112"`
    DischargedToLocation *DLD `json:"dischargedToLocation,omitempty" hl7:"part=38 length=47 table=0114"`
    DietType *CWE `json:"dietType,omitempty" hl7:"part=39 length=705"`
    ServicingFacility IS `json:"servicingFacility,omitempty" hl7:"part=40 length=2 table=0115"`
    BedStatus IS `json:"bedStatus,omitempty" hl7:"part=41 length=1 table=0116"`
    AccountStatus IS `json:"accountStatus,omitempty" hl7:"part=42 length=2 table=0117"`
    PendingLocation *PL `json:"pendingLocation,omitempty" hl7:"part=43 length=80"`
    PriorTemporaryLocation *PL `json:"priorTemporaryLocation,omitempty" hl7:"part=44 length=80"`
    AdmitDateTime DTM `json:"admitDateTime,omitempty" hl7:"part=45 length=24"`
    DischargeDateTime DTM `json:"dischargeDateTime,omitempty" hl7:"part=46 length=24"`
    CurrentPatientBalance NM `json:"currentPatientBalance,omitempty" hl7:"part=47 length=12"`
    TotalCharges NM `json:"totalCharges,omitempty" hl7:"part=48 length=12"`
    TotalAdjustments NM `json:"totalAdjustments,omitempty" hl7:"part=49 length=12"`
    TotalPayments NM `json:"totalPayments,omitempty" hl7:"part=50 length=12"`
    AlternateVisitId *CX `json:"alternateVisitId,omitempty" hl7:"part=51 length=250 table=0326"`
    VisitIndicator IS `json:"visitIndicator,omitempty" hl7:"part=52 length=1"`
    OtherHealthcareProvider []*XCN `json:"otherHealthcareProvider,omitempty" hl7:"part=53 length=250 table=0010"`
}

type HL7 struct {
    Msh *MSH `json:"msh,omitempty" hl7:"part=MSH length=-1"`
    Pid []*PID `json:"pid,omitempty" hl7:"part=PID length=-1"`
    Obx []*OBX `json:"obx,omitempty" hl7:"part=OBX length=-1"`
    Obr *OBR `json:"obr,omitempty" hl7:"part=OBR length=-1"`
    Pv1 *PV1 `json:"pv1,omitempty" hl7:"part=PV1 length=-1"`
}

type NA  []float64


func (ds *decodeState) instantiateAndParseHl7(dtn string) (parseStatus, error) {
    switch dtn {
    case "SI":
        rValue := reflect.New(reflect.TypeFor[SI]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "DR":
        rValue := reflect.New(reflect.TypeFor[DR]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "DLN":
        rValue := reflect.New(reflect.TypeFor[DLN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "XAD":
        rValue := reflect.New(reflect.TypeFor[XAD]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "HL7":
        rValue := reflect.New(reflect.TypeFor[HL7]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "CQ":
        rValue := reflect.New(reflect.TypeFor[CQ]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "PRL":
        rValue := reflect.New(reflect.TypeFor[PRL]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "RI":
        rValue := reflect.New(reflect.TypeFor[RI]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "CNE":
        rValue := reflect.New(reflect.TypeFor[CNE]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "MSH":
        rValue := reflect.New(reflect.TypeFor[MSH]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "CX":
        rValue := reflect.New(reflect.TypeFor[CX]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "DLD":
        rValue := reflect.New(reflect.TypeFor[DLD]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "NA":
        rValue := reflect.New(reflect.TypeFor[NA]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "IS":
        rValue := reflect.New(reflect.TypeFor[IS]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "EI":
        rValue := reflect.New(reflect.TypeFor[EI]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "MOC":
        rValue := reflect.New(reflect.TypeFor[MOC]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "DT":
        rValue := reflect.New(reflect.TypeFor[DT]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "SN":
        rValue := reflect.New(reflect.TypeFor[SN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "XPN":
        rValue := reflect.New(reflect.TypeFor[XPN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "PID":
        rValue := reflect.New(reflect.TypeFor[PID]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "OBX":
        rValue := reflect.New(reflect.TypeFor[OBX]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "HD":
        rValue := reflect.New(reflect.TypeFor[HD]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "TX":
        rValue := reflect.New(reflect.TypeFor[TX]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "XTN":
        rValue := reflect.New(reflect.TypeFor[XTN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "TQ":
        rValue := reflect.New(reflect.TypeFor[TQ]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "NDL":
        rValue := reflect.New(reflect.TypeFor[NDL]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "OBR":
        rValue := reflect.New(reflect.TypeFor[OBR]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "OSD":
        rValue := reflect.New(reflect.TypeFor[OSD]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "SAD":
        rValue := reflect.New(reflect.TypeFor[SAD]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "FC":
        rValue := reflect.New(reflect.TypeFor[FC]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "PV1":
        rValue := reflect.New(reflect.TypeFor[PV1]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "FN":
        rValue := reflect.New(reflect.TypeFor[FN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "SPS":
        rValue := reflect.New(reflect.TypeFor[SPS]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "NM":
        rValue := reflect.New(reflect.TypeFor[NM]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "ID":
        rValue := reflect.New(reflect.TypeFor[ID]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "MO":
        rValue := reflect.New(reflect.TypeFor[MO]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "ST":
        rValue := reflect.New(reflect.TypeFor[ST]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "CWE":
        rValue := reflect.New(reflect.TypeFor[CWE]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "DTM":
        rValue := reflect.New(reflect.TypeFor[DTM]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "CE":
        rValue := reflect.New(reflect.TypeFor[CE]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "EIP":
        rValue := reflect.New(reflect.TypeFor[EIP]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "CNN":
        rValue := reflect.New(reflect.TypeFor[CNN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "ProcessingType":
        rValue := reflect.New(reflect.TypeFor[ProcessingType]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "TN":
        rValue := reflect.New(reflect.TypeFor[TN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "XON":
        rValue := reflect.New(reflect.TypeFor[XON]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "PL":
        rValue := reflect.New(reflect.TypeFor[PL]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "XCN":
        rValue := reflect.New(reflect.TypeFor[XCN]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "MSG":
        rValue := reflect.New(reflect.TypeFor[MSG]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "VID":
        rValue := reflect.New(reflect.TypeFor[VID]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    case "Varies":
        rValue := reflect.New(reflect.TypeFor[Varies]())
        ds.value.Elem().Set(rValue)
        ds.value = rValue
    default:
        err := fmt.Errorf("Error setting any type field: type (%s) not found", dtn)
        log.Printf("\t%s\n", err.Error())
        return INVALID, err
    }
    return ds.parseHl7()
}

func (ds *decodeState) isTypeStruct(dtn string) bool {
    switch dtn {
    case "SI":
        return false
    case "DR":
        return true
    case "DLN":
        return true
    case "XAD":
        return true
    case "HL7":
        return true
    case "CQ":
        return true
    case "PRL":
        return true
    case "RI":
        return true
    case "CNE":
        return true
    case "MSH":
        return true
    case "CX":
        return true
    case "DLD":
        return true
    case "NA":
        return false
    case "IS":
        return false
    case "EI":
        return true
    case "MOC":
        return true
    case "DT":
        return false
    case "SN":
        return true
    case "XPN":
        return true
    case "PID":
        return true
    case "OBX":
        return true
    case "HD":
        return true
    case "TX":
        return false
    case "XTN":
        return true
    case "TQ":
        return true
    case "NDL":
        return true
    case "OBR":
        return true
    case "OSD":
        return true
    case "SAD":
        return true
    case "FC":
        return true
    case "PV1":
        return true
    case "FN":
        return true
    case "SPS":
        return true
    case "NM":
        return false
    case "ID":
        return false
    case "MO":
        return true
    case "ST":
        return false
    case "CWE":
        return true
    case "DTM":
        return false
    case "CE":
        return true
    case "EIP":
        return true
    case "CNN":
        return true
    case "ProcessingType":
        return true
    case "TN":
        return false
    case "XON":
        return true
    case "PL":
        return true
    case "XCN":
        return true
    case "MSG":
        return true
    case "VID":
        return true
    case "Varies":
        return false
    default:
        return false
    }
}

