package sr_Cyrl_ME

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sr_Cyrl_ME struct {
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

// New returns a new instance of translator for the 'sr_Cyrl_ME' locale
func New() locales.Translator {
	return &sr_Cyrl_ME{
		locale:                 "sr_Cyrl_ME",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ".",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "јан.", "феб.", "март", "апр.", "мај", "јун", "јул", "авг.", "септ.", "окт.", "нов.", "дец."},
		monthsNarrow:           []string{"", "ј", "ф", "м", "а", "м", "ј", "ј", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
		daysAbbreviated:        []string{"нед.", "пон.", "ут.", "ср.", "чет.", "пет.", "суб."},
		daysNarrow:             []string{"н", "п", "у", "с", "ч", "п", "с"},
		daysShort:              []string{"не", "по", "ут", "ср", "че", "пе", "су"},
		daysWide:               []string{"недјеља", "понедељак", "уторак", "сриједа", "четвртак", "петак", "субота"},
		periodsAbbreviated:     []string{"прије подне", "по подне"},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"прије подне", "по подне"},
		erasAbbreviated:        []string{"", ""},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"HNPM": "Сен Пјер и Микелон стандардно време", "CST": "Северноамеричко централно стандардно време", "AWST": "Аустралијско западно стандардно време", "ART": "Аргентина стандардно време", "HKT": "Хонг Конг стандардно време", "COT": "Колумбија стандардно време", "∅∅∅": "Акре летње рачунање времена", "LHDT": "Лорд Хов летње рачунање времена", "HNPMX": "Мексички Пацифик стандардно време", "WIB": "Западно-индонезијско време", "SGT": "Сингапур стандардно време", "JST": "Јапанско стандардно време", "WART": "Западна Аргентина стандардно време", "MST": "Макао стандардно време", "HKST": "Хонг Конг летње рачунање времена", "HNCU": "Куба стандардно време", "CDT": "Северноамеричко централно летње време", "CHAST": "Чатам стандардно време", "CLT": "Чиле стандардно време", "CLST": "Чиле летње рачунање времена", "COST": "Колумбија летње рачунање времена", "EAT": "Источно-афричко време", "MESZ": "Средњеевропско летње рачунање времена", "GMT": "Средње време по Гриничу", "EDT": "Северноамеричко источно летње време", "GFT": "Француска Гвајана време", "WITA": "Централно-индонезијско време", "ECT": "Еквадор време", "IST": "Индијско стандардно време", "ARST": "Аргентина летње рачунање времена", "AEDT": "Аустралијско источно летње рачунање времена", "GYT": "Гвајана време", "MEZ": "Средњеевропско стандардно време", "AKST": "Аљаска, стандардно време", "AEST": "Аустралијско источно стандардно време", "OESZ": "Источноевропско летње рачунање времена", "MYT": "Малезија време", "TMST": "Туркменистан летње рачунање времена", "EST": "Северноамеричко источно стандардно време", "LHST": "Лорд Хов стандардно време", "HADT": "Хавајско-алеутско летње рачунање времена", "ACWST": "Аустралијско централно западно стандардно време", "MDT": "Макао летње рачунање времена", "SAST": "Јужно-афричко време", "HECU": "Куба летње рачунање времена", "NZST": "Нови Зеланд стандардно време", "WARST": "Западна Аргентина летње рачунање времена", "WAST": "Западно-афричко летње рачунање времена", "ACDT": "Аустралијско централно летње рачунање времена", "HAT": "Њуфаундленд летње рачунање времена", "PDT": "Северноамеричко пацифичко летње време", "VET": "Венецуела време", "AST": "Атлантско стандардно време", "HEOG": "Западни Гренланд летње рачунање времена", "WAT": "Западно-афричко стандардно време", "ACST": "Аустралијско централно стандардно време", "ADT": "Атлантско летње рачунање времена", "UYST": "Уругвај летње рачунање времена", "SRT": "Суринам време", "WIT": "Источно-индонезијско време", "PST": "Северноамеричко пацифичко стандардно време", "JDT": "Јапанско летње рачунање времена", "OEZ": "Источноевропско стандардно време", "HNOG": "Западни Гренланд стандардно време", "HENOMX": "Северозападни Мексико летње рачунање времена", "BT": "Бутан време", "HNT": "Њуфаундленд стандардно време", "CAT": "Централно-афричко време", "WEZ": "Западноевропско стандардно време", "TMT": "Туркменистан стандардно време", "HNNOMX": "Северозападни Мексико стандардно време", "HEEG": "Источни Гренланд летње рачунање времена", "AKDT": "Аљаска, летње рачунање времена", "ChST": "Чаморо време", "HEPM": "Сен Пјер и Микелон летње рачунање времена", "HEPMX": "Мексички Пацифик летње рачунање времена", "BOT": "Боливија време", "NZDT": "Нови Зеланд летње рачунање времена", "AWDT": "Аустралијско западно летње рачунање времена", "CHADT": "Чатам летње рачунање времена", "WESZ": "Западноевропско летње рачунање времена", "HNEG": "Источни Гренланд стандардно време", "UYT": "Уругвај стандардно време", "HAST": "Хавајско-алеутско стандардно време", "ACWDT": "Аустралијско централно западно летње рачунање времена"},
	}
}

// Locale returns the current translators string locale
func (sr *sr_Cyrl_ME) Locale() string {
	return sr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) PluralsCardinal() []locales.PluralRule {
	return sr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) PluralsOrdinal() []locales.PluralRule {
	return sr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) PluralsRange() []locales.PluralRule {
	return sr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod100 := f % 100
	fMod10 := f % 10

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14)) || (fMod10 >= 2 && fMod10 <= 4 && (fMod100 < 12 || fMod100 > 14)) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sr.CardinalPluralRule(num1, v1)
	end := sr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sr *sr_Cyrl_ME) MonthAbbreviated(month time.Month) string {
	return sr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sr *sr_Cyrl_ME) MonthsAbbreviated() []string {
	return sr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sr *sr_Cyrl_ME) MonthNarrow(month time.Month) string {
	return sr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sr *sr_Cyrl_ME) MonthsNarrow() []string {
	return sr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sr *sr_Cyrl_ME) MonthWide(month time.Month) string {
	return sr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sr *sr_Cyrl_ME) MonthsWide() []string {
	return sr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sr *sr_Cyrl_ME) WeekdayAbbreviated(weekday time.Weekday) string {
	return sr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sr *sr_Cyrl_ME) WeekdaysAbbreviated() []string {
	return sr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sr *sr_Cyrl_ME) WeekdayNarrow(weekday time.Weekday) string {
	return sr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sr *sr_Cyrl_ME) WeekdaysNarrow() []string {
	return sr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sr *sr_Cyrl_ME) WeekdayShort(weekday time.Weekday) string {
	return sr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sr *sr_Cyrl_ME) WeekdaysShort() []string {
	return sr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sr *sr_Cyrl_ME) WeekdayWide(weekday time.Weekday) string {
	return sr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sr *sr_Cyrl_ME) WeekdaysWide() []string {
	return sr.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sr_Cyrl_ME' and handles both Whole and Real numbers based on 'v'
func (sr *sr_Cyrl_ME) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sr_Cyrl_ME' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sr *sr_Cyrl_ME) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sr_Cyrl_ME'
// in accounting notation.
func (sr *sr_Cyrl_ME) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sr.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, sr.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sr_Cyrl_ME'
func (sr *sr_Cyrl_ME) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
