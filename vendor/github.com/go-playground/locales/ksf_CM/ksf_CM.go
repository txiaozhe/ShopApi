package ksf_CM

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ksf_CM struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'ksf_CM' locale
func New() locales.Translator {
	return &ksf_CM{
		locale:                 "ksf_CM",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "ŋ1", "ŋ2", "ŋ3", "ŋ4", "ŋ5", "ŋ6", "ŋ7", "ŋ8", "ŋ9", "ŋ10", "ŋ11", "ŋ12"},
		monthsWide:             []string{"", "ŋwíí a ntɔ́ntɔ", "ŋwíí akǝ bɛ́ɛ", "ŋwíí akǝ ráá", "ŋwíí akǝ nin", "ŋwíí akǝ táan", "ŋwíí akǝ táafɔk", "ŋwíí akǝ táabɛɛ", "ŋwíí akǝ táaraa", "ŋwíí akǝ táanin", "ŋwíí akǝ ntɛk", "ŋwíí akǝ ntɛk di bɔ́k", "ŋwíí akǝ ntɛk di bɛ́ɛ"},
		daysAbbreviated:        []string{"sɔ́n", "lǝn", "maa", "mɛk", "jǝǝ", "júm", "sam"},
		daysNarrow:             []string{"s", "l", "m", "m", "j", "j", "s"},
		daysWide:               []string{"sɔ́ndǝ", "lǝndí", "maadí", "mɛkrɛdí", "jǝǝdí", "júmbá", "samdí"},
		periodsAbbreviated:     []string{"sárúwá", "cɛɛ́nko"},
		periodsWide:            []string{"sárúwá", "cɛɛ́nko"},
		erasAbbreviated:        []string{"d.Y.", "k.Y."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"di Yɛ́sus aká yálɛ", "cámɛɛn kǝ kǝbɔpka Y"},
		timezones:              map[string]string{"HEPMX": "HEPMX", "IST": "IST", "ACST": "ACST", "∅∅∅": "∅∅∅", "WIB": "WIB", "NZDT": "NZDT", "MEZ": "MEZ", "ADT": "ADT", "HNCU": "HNCU", "MESZ": "MESZ", "AST": "AST", "WEZ": "WEZ", "EDT": "EDT", "BT": "BT", "LHST": "LHST", "PDT": "PDT", "HNOG": "HNOG", "GMT": "GMT", "TMT": "TMT", "HEPM": "HEPM", "OESZ": "OESZ", "MYT": "MYT", "TMST": "TMST", "HNNOMX": "HNNOMX", "WITA": "WITA", "ECT": "ECT", "ACWDT": "ACWDT", "NZST": "NZST", "MDT": "MDT", "HNEG": "HNEG", "UYST": "UYST", "VET": "VET", "OEZ": "OEZ", "EST": "EST", "HAT": "HAT", "ChST": "ChST", "HAST": "HAST", "HADT": "HADT", "WARST": "WARST", "ACDT": "ACDT", "AKST": "AKST", "LHDT": "LHDT", "WIT": "WIT", "CHAST": "CHAST", "WART": "WART", "CLST": "CLST", "WAT": "WAT", "COT": "COT", "EAT": "EAT", "HECU": "HECU", "CHADT": "CHADT", "CAT": "CAT", "HENOMX": "HENOMX", "SAST": "SAST", "AWST": "AWST", "JST": "JST", "COST": "COST", "AKDT": "AKDT", "HNPMX": "HNPMX", "BOT": "BOT", "CLT": "CLT", "HNT": "HNT", "AEST": "AEST", "HNPM": "HNPM", "ART": "ART", "GYT": "GYT", "CST": "CST", "CDT": "CDT", "PST": "PST", "ACWST": "ACWST", "JDT": "JDT", "MST": "MST", "WAST": "WAST", "UYT": "UYT", "AWDT": "AWDT", "SGT": "SGT", "ARST": "ARST", "HKT": "HKT", "AEDT": "AEDT", "GFT": "GFT", "SRT": "SRT", "HEOG": "HEOG", "WESZ": "WESZ", "HKST": "HKST", "HEEG": "HEEG"},
	}
}

// Locale returns the current translators string locale
func (ksf *ksf_CM) Locale() string {
	return ksf.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ksf_CM'
func (ksf *ksf_CM) PluralsCardinal() []locales.PluralRule {
	return ksf.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ksf_CM'
func (ksf *ksf_CM) PluralsOrdinal() []locales.PluralRule {
	return ksf.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ksf_CM'
func (ksf *ksf_CM) PluralsRange() []locales.PluralRule {
	return ksf.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ksf_CM'
func (ksf *ksf_CM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ksf_CM'
func (ksf *ksf_CM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ksf_CM'
func (ksf *ksf_CM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ksf *ksf_CM) MonthAbbreviated(month time.Month) string {
	return ksf.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ksf *ksf_CM) MonthsAbbreviated() []string {
	return ksf.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ksf *ksf_CM) MonthNarrow(month time.Month) string {
	return ksf.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ksf *ksf_CM) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (ksf *ksf_CM) MonthWide(month time.Month) string {
	return ksf.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ksf *ksf_CM) MonthsWide() []string {
	return ksf.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ksf *ksf_CM) WeekdayAbbreviated(weekday time.Weekday) string {
	return ksf.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ksf *ksf_CM) WeekdaysAbbreviated() []string {
	return ksf.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ksf *ksf_CM) WeekdayNarrow(weekday time.Weekday) string {
	return ksf.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ksf *ksf_CM) WeekdaysNarrow() []string {
	return ksf.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ksf *ksf_CM) WeekdayShort(weekday time.Weekday) string {
	return ksf.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ksf *ksf_CM) WeekdaysShort() []string {
	return ksf.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ksf *ksf_CM) WeekdayWide(weekday time.Weekday) string {
	return ksf.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ksf *ksf_CM) WeekdaysWide() []string {
	return ksf.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ksf_CM' and handles both Whole and Real numbers based on 'v'
func (ksf *ksf_CM) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 1 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ksf.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ksf.group) - 1; j >= 0; j-- {
					b = append(b, ksf.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ksf.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ksf_CM' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ksf *ksf_CM) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ksf_CM'
func (ksf *ksf_CM) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ksf.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ksf.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ksf.group) - 1; j >= 0; j-- {
					b = append(b, ksf.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ksf.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ksf.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ksf.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ksf_CM'
// in accounting notation.
func (ksf *ksf_CM) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ksf.currencies[currency]
	l := len(s) + len(symbol) + 3 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ksf.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ksf.group) - 1; j >= 0; j-- {
					b = append(b, ksf.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ksf.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ksf.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ksf.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ksf.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ksf.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ksf.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ksf.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ksf.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ksf_CM'
func (ksf *ksf_CM) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ksf.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ksf.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
