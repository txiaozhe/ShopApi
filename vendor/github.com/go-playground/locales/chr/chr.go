package chr

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type chr struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'chr' locale
func New() locales.Translator {
	return &chr{
		locale:                 "chr",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "A$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "ᎤᏃ", "ᎧᎦ", "ᎠᏅ", "ᎧᏬ", "ᎠᏂ", "ᏕᎭ", "ᎫᏰ", "ᎦᎶ", "ᏚᎵ", "ᏚᏂ", "ᏅᏓ", "ᎥᏍ"},
		monthsNarrow:           []string{"", "Ꭴ", "Ꭷ", "Ꭰ", "Ꭷ", "Ꭰ", "Ꮥ", "Ꭻ", "Ꭶ", "Ꮪ", "Ꮪ", "Ꮕ", "Ꭵ"},
		monthsWide:             []string{"", "ᎤᏃᎸᏔᏅ", "ᎧᎦᎵ", "ᎠᏅᏱ", "ᎧᏬᏂ", "ᎠᏂᏍᎬᏘ", "ᏕᎭᎷᏱ", "ᎫᏰᏉᏂ", "ᎦᎶᏂ", "ᏚᎵᏍᏗ", "ᏚᏂᏅᏗ", "ᏅᏓᏕᏆ", "ᎥᏍᎩᏱ"},
		daysAbbreviated:        []string{"ᏆᏍᎬ", "ᏉᏅᎯ", "ᏔᎵᏁ", "ᏦᎢᏁ", "ᏅᎩᏁ", "ᏧᎾᎩ", "ᏈᏕᎾ"},
		daysNarrow:             []string{"Ꮖ", "Ꮙ", "Ꮤ", "Ꮶ", "Ꮕ", "Ꮷ", "Ꭴ"},
		daysShort:              []string{"ᏍᎬ", "ᏅᎯ", "ᏔᎵ", "ᏦᎢ", "ᏅᎩ", "ᏧᎾ", "ᏕᎾ"},
		daysWide:               []string{"ᎤᎾᏙᏓᏆᏍᎬ", "ᎤᎾᏙᏓᏉᏅᎯ", "ᏔᎵᏁᎢᎦ", "ᏦᎢᏁᎢᎦ", "ᏅᎩᏁᎢᎦ", "ᏧᎾᎩᎶᏍᏗ", "ᎤᎾᏙᏓᏈᏕᎾ"},
		periodsAbbreviated:     []string{"ᏌᎾᎴ", "ᏒᎯᏱᎢᏗᏢ"},
		periodsNarrow:          []string{"Ꮜ", "Ꮢ"},
		periodsWide:            []string{"ᏌᎾᎴ", "ᏒᎯᏱᎢᏗᏢ"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ᏧᏓᎷᎸ ᎤᎷᎯᏍᏗ ᎦᎶᏁᏛ", "ᎠᏃ ᏙᎻᏂ"},
		timezones:              map[string]string{"AWDT": "ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "CHADT": "ᏣᏝᎻ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "WEZ": "ᏭᏕᎵᎬ ᏗᏜ ᏳᎳᏈ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HNT": "ᎢᏤᎤᏂᏩᏛᏓᎦᏙᎯ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HADT": "ᎭᏩᏱ-ᎠᎵᏳᏏᎠᏂ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "MESZ": "ᎠᏰᏟ ᏳᎳᏈ ᎪᎩ ᎠᏟᎢᎵᏒ", "OEZ": "ᏗᎧᎸᎬ ᏗᏜ ᏳᎳᏈ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "TMT": "ᏛᎵᎩᎺᏂᏍᏔᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "EDT": "ᏗᎧᎸᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "HNCU": "ᎫᏆ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "NZDT": "ᎢᏤ ᏏᎢᎴᏂᏗ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "HEOG": "ᏭᏕᎵᎬ ᎢᏤᏍᏛᏱ ᎪᎩ ᎠᏟᎢᎵᏒ", "ACST": "ᎠᏰᏟ ᎡᎳᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "WITA": "ᎠᏰᏟ ᎢᏂᏙᏂᏍᏯ ᎠᏟᎢᎵᏒ", "PDT": "ᏭᏕᎵᎬ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "HNEG": "ᏗᎧᎸᎬ ᎢᏤᏍᏛᏱ ᎠᏟᎶᏍᏗ ᎠᎵᎢᎵᏒ", "UYST": "ᏳᎷᏇ ᎪᎩ ᎠᏟᎢᎵᏒ", "LHST": "ᎤᎬᏫᏳᎯ ᎭᏫ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "LHDT": "ᎤᎬᏫᏳᎯ ᎭᏫ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "SAST": "ᏧᎦᎾᏮ ᎬᎿᎨᏍᏛ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "GYT": "ᎦᏯᎾ ᎠᏟᎢᎵᏒ", "CST": "ᎠᏰᏟ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "JDT": "ᏣᏩᏂᏏ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "WARST": "ᏭᏕᎵᎬ ᏗᏜ ᎠᏥᏂᏘᏂᎠ ᎪᎩ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HECU": "ᎫᏆ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "EAT": "ᏗᎧᎸᎬ ᎬᎿᎨᏍᏛ ᎠᏟᎢᎵᏒ", "HEEG": "ᏗᎧᎸᎬ ᎢᏤᏍᏛᏱ ᎪᎩ ᎠᏟᎢᎵᏒ", "BT": "ᏊᏔᏂ ᎠᏟᎢᎵᏒ", "ACDT": "ᎠᏰᏟ ᎡᎳᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "HEPMX": "ᎠᏂᏍᏆᏂ ᏭᏕᎵᎬ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "CDT": "ᎠᏰᏟ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "BOT": "ᏉᎵᏫᎠ ᎠᏟᎢᎵᏒ", "NZST": "ᎢᏤ ᏏᎢᎴᏂᏗ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "JST": "ᏣᏩᏂᏏ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AWST": "ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "PST": "ᏭᏕᎵᎬ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HNOG": "ᏭᏕᎵᎬ ᎢᏤᏍᏛᏱ ᎠᏟᎶᏍᏗ ᎠᎵᎢᎵᏒ", "WAST": "ᏭᏕᎵᎬ ᎬᎿᎨᏍᏛ ᎪᎩ ᎠᏟᎢᎵᏒ", "ChST": "ᏣᎼᎶ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "CHAST": "ᏣᏝᎻ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "WESZ": "ᏭᏕᎵᎬ ᏗᏜ ᏳᎳᏈ ᎪᎩ ᎠᏟᎢᎵᏒ", "MYT": "ᎹᎴᏏᎢᎠ ᎠᏟᎢᎵᏒ", "HNNOMX": "ᏧᏴᏢ ᏭᏕᎵᎬ ᎠᏂᏍᏆᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "COT": "ᎪᎸᎻᏈᎢᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AEDT": "ᎡᎳᏗᏜ ᏗᎧᎸᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "HNPMX": "ᎠᏂᏍᏆᏂ ᏭᏕᎵᎬ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "OESZ": "ᏗᎧᎸᎬ ᏗᏜ ᏳᎳᏈ ᎪᎩ ᎠᏟᎢᎵᏒ", "CLST": "ᏥᎵ ᎪᎩ ᎠᏟᎢᎵᏒ", "ADT": "ᏗᎧᎸᎬ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "HKST": "ᎰᏂᎩ ᎪᏂᎩ ᎪᎩ ᎠᏟᎢᎵᏒ", "SRT": "ᏒᎵᎾᎻ ᎠᏟᎢᎵᏒ", "SGT": "ᏏᏂᎦᏉᎵ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "EST": "ᏗᎧᎸᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HAT": "ᎢᏤᎤᏂᏩᏛᏓᎦᏙᎯ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "AEST": "ᎡᎳᏗᏜ ᏗᎧᎸᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "WIT": "ᏗᎧᎸᎬ ᏗᏜ ᎢᏂᏙᏂᏍᏯ ᎠᏟᎢᎵᏒ", "IST": "ᎢᏂᏗᎢᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "GMT": "ᎢᏤ ᎢᏳᏍᏗ ᎠᏟᎢᎵᏒ", "WART": "ᏭᏕᎵᎬ ᏗᏜ ᎠᏥᏂᏘᏂᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "TMST": "ᏛᎵᎩᎺᏂᏍᏔᏂ ᎪᎩ ᎠᏟᎢᎵᏒ", "HEPM": "ᎤᏓᏅᏘ ᏈᏰ ᎠᎴ ᎻᏇᎶᏂ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "CLT": "ᏥᎵ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AST": "ᏗᎧᎸᎬ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "MST": "MST", "WAT": "ᏭᏕᎵᎬ ᎬᎿᎨᏍᏛ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AKDT": "ᎠᎳᏍᎦ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "ARST": "ᎠᏥᏂᏘᏂᎠ ᎪᎩ ᎠᏟᎢᎵᏒ", "COST": "ᎪᎸᎻᏈᎢᎠ ᎪᎩ ᎠᏟᎢᎵᏒ", "GFT": "ᎠᏂᎦᎸ ᏈᏯᎾ ᎠᏟᎢᎵᏒ", "HAST": "ᎭᏩᏱ-ᎠᎵᏳᏏᎠᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "ACWST": "ᎠᏰᏟ ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "VET": "ᏪᏁᏑᏪᎳ ᎠᏟᎢᎵᏒ", "MEZ": "ᎠᏰᏟ ᏳᎳᏈ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "ART": "ᎠᏥᏂᏘᏂᎠ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HNPM": "ᎤᏓᏅᏘ ᏈᏰ ᎠᎴ ᎻᏇᎶᏂ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "HKT": "ᎰᏂᎩ ᎪᏂᎩ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "AKST": "ᎠᎳᏍᎦ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ", "∅∅∅": "ᎠᏐᎴᏏ ᎪᎩ ᎠᏟᎢᎵᏒ", "WIB": "ᏭᏕᎵᎬ ᏗᏜ ᎢᏂᏙᏂᏍᏯ ᎠᏟᎢᎵᏒ", "ECT": "ᎡᏆᏙᎵ ᎠᏟᎢᎵᏒ", "ACWDT": "ᎠᏰᏟ ᎡᎳᏗᏜ ᏭᏕᎵᎬ ᏗᏜ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏒᎩ", "CAT": "ᎠᏰᏟ ᎬᎿᎨᏍᏛ ᎠᏟᎢᎵᏒ", "MDT": "MDT", "HENOMX": "ᏧᏴᏢ ᏭᏕᎵᎬ ᎠᏂᏍᏆᏂ ᎪᎯ ᎢᎦ ᎠᏟᎢᎵᏍᏒᎩ", "UYT": "ᏳᎷᏇ ᎠᏟᎶᏍᏗ ᎠᏟᎢᎵᏒ"},
	}
}

// Locale returns the current translators string locale
func (chr *chr) Locale() string {
	return chr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'chr'
func (chr *chr) PluralsCardinal() []locales.PluralRule {
	return chr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'chr'
func (chr *chr) PluralsOrdinal() []locales.PluralRule {
	return chr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'chr'
func (chr *chr) PluralsRange() []locales.PluralRule {
	return chr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'chr'
func (chr *chr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'chr'
func (chr *chr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'chr'
func (chr *chr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (chr *chr) MonthAbbreviated(month time.Month) string {
	return chr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (chr *chr) MonthsAbbreviated() []string {
	return chr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (chr *chr) MonthNarrow(month time.Month) string {
	return chr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (chr *chr) MonthsNarrow() []string {
	return chr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (chr *chr) MonthWide(month time.Month) string {
	return chr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (chr *chr) MonthsWide() []string {
	return chr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (chr *chr) WeekdayAbbreviated(weekday time.Weekday) string {
	return chr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (chr *chr) WeekdaysAbbreviated() []string {
	return chr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (chr *chr) WeekdayNarrow(weekday time.Weekday) string {
	return chr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (chr *chr) WeekdaysNarrow() []string {
	return chr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (chr *chr) WeekdayShort(weekday time.Weekday) string {
	return chr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (chr *chr) WeekdaysShort() []string {
	return chr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (chr *chr) WeekdayWide(weekday time.Weekday) string {
	return chr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (chr *chr) WeekdaysWide() []string {
	return chr.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'chr' and handles both Whole and Real numbers based on 'v'
func (chr *chr) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'chr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (chr *chr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, chr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'chr'
func (chr *chr) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := chr.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, chr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'chr'
// in accounting notation.
func (chr *chr) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := chr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, chr.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, chr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, chr.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'chr'
func (chr *chr) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'chr'
func (chr *chr) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, chr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'chr'
func (chr *chr) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, chr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'chr'
func (chr *chr) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, chr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, chr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'chr'
func (chr *chr) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'chr'
func (chr *chr) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'chr'
func (chr *chr) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'chr'
func (chr *chr) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := chr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
