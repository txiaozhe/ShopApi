package bs

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type bs struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
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

// New returns a new instance of translator for the 'bs' locale
func New() locales.Translator {
	return &bs{
		locale:                 "bs",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "KM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "kn", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "din.", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "maj", "jun", "jul", "avg", "sep", "okt", "nov", "dec"},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januar", "februar", "mart", "april", "maj", "juni", "juli", "avgust", "septembar", "oktobar", "novembar", "decembar"},
		daysAbbreviated:        []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysNarrow:             []string{"N", "P", "U", "S", "Č", "P", "S"},
		daysShort:              []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysWide:               []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
		periodsAbbreviated:     []string{"prijepodne", "popodne"},
		periodsNarrow:          []string{"prijepodne", "popodne"},
		periodsWide:            []string{"prijepodne", "popodne"},
		erasAbbreviated:        []string{"p. n. e.", "n. e."},
		erasNarrow:             []string{"pr.n.e.", "AD"},
		erasWide:               []string{"prije nove ere", "nove ere"},
		timezones:              map[string]string{"MDT": "Makao letnje računanje vremena", "PDT": "Sjevernoameričko pacifičko ljetno vrijeme", "OEZ": "Istočnoevropsko standardno vrijeme", "WEZ": "Zapadnoevropsko standardno vrijeme", "HKT": "Hongkonško standardno vrijeme", "COT": "Kolumbijsko standardno vrijeme", "CDT": "Sjevernoameričko centralno ljetno vrijeme", "JST": "Japansko standardno vrijeme", "AST": "Sjevernoameričko atlantsko standardno vrijeme", "MST": "Makao standardno vreme", "AEDT": "Istočnoaustralijsko ljetno vrijeme", "WIT": "Istočnoindonezijsko vrijeme", "ECT": "Ekvadorsko vrijeme", "HEEG": "Istočnogrenlandsko ljetno vrijeme", "CST": "Sjevernoameričko centralno standardno vrijeme", "TMT": "Turkmenistansko standardno vrijeme", "HKST": "Hongkonško ljetno vrijeme", "ACST": "Centralnoaustralijsko standardno vrijeme", "HEPMX": "Meksičko pacifičko ljetno vrijeme", "AWST": "Zapadnoaustralijsko standardno vrijeme", "CLST": "Čileansko ljetno vrijeme", "ARST": "Argentinsko ljetno vrijeme", "COST": "Kolumbijsko ljetno vrijeme", "HNPM": "Standardno vrijeme na Ostrvima Sen Pjer i Mikelon", "LHDT": "Ljetno vrijeme na Ostrvu Lord Hau", "SRT": "Surinamsko vrijeme", "BOT": "Bolivijsko vrijeme", "MESZ": "Centralnoevropsko ljetno vrijeme", "ART": "Argentinsko standardno vrijeme", "HNNOMX": "Sjeverozapadno meksičko standardno vrijeme", "∅∅∅": "Peruansko ljetno vrijeme", "WITA": "Centralnoindonezijsko vrijeme", "HNPMX": "Meksičko pacifičko standardno vrijeme", "VET": "Venecuelansko vrijeme", "ADT": "Sjevernoameričko atlantsko ljetno vrijeme", "EAT": "Istočnoafričko vrijeme", "CAT": "Centralnoafričko vrijeme", "CHAST": "Čatamsko standardno vrijeme", "CHADT": "Čatamsko ljetno vrijeme", "SGT": "Singapursko standardno vrijeme", "HENOMX": "Sjeverozapadno meksičko ljetno vrijeme", "HNEG": "Istočnogrenlandsko standardno vrijeme", "AKDT": "Aljaskansko ljetno vrijeme", "ChST": "Čamorsko standardno vrijeme", "HECU": "Kubansko ljetno vrijeme", "ACWDT": "Australijsko centralnozapadno ljetno vrijeme", "SAST": "Južnoafričko standardno vrijeme", "GMT": "Griničko vrijeme", "WAST": "Zapadnoafričko ljetno vrijeme", "EST": "Sjevernoameričko istočno standardno vrijeme", "AKST": "Aljaskansko standardno vrijeme", "HEPM": "Ljetno vrijeme na Ostrvima Sen Pjer i Mikelon", "JDT": "Japansko ljetno vrijeme", "HNOG": "Zapadnogrenlandsko standardno vrijeme", "MYT": "Malezijsko vrijeme", "BT": "Butansko vrijeme", "GYT": "Gvajansko vrijeme", "NZST": "Novozelandsko standardno vrijeme", "IST": "Indijsko standardno vrijeme", "WART": "Zapadnoargentinsko standardno vrijeme", "WESZ": "Zapadnoevropsko ljetno vrijeme", "LHST": "Standardno vrijeme na Ostrvu Lord Hau", "AWDT": "Zapadnoaustralijsko ljetno vrijeme", "HADT": "Havajsko-aleućansko ljetno vrijeme", "HNT": "Njufaundlendsko standardno vrijeme", "MEZ": "Centralnoevropsko standardno vrijeme", "CLT": "Čileansko standardno vrijeme", "ACDT": "Centralnoaustralijsko ljetno vrijeme", "HNCU": "Kubansko standardno vrijeme", "WIB": "Zapadnoindonezijsko vrijeme", "PST": "Sjevernoameričko pacifičko standardno vrijeme", "HAST": "Havajsko-aleućansko standardno vrijeme", "ACWST": "Australijsko centralnozapadno standardno vrijeme", "NZDT": "Novozelandsko ljetno vrijeme", "HEOG": "Zapadnogrenlandsko ljetno vrijeme", "OESZ": "Istočnoevropsko ljetno vrijeme", "WAT": "Zapadnoafričko standardno vrijeme", "HAT": "Njufaundlendsko ljetno vrijeme", "UYST": "Urugvajsko ljetno vrijeme", "UYT": "Urugvajsko standardno vrijeme", "WARST": "Zapadnoargentinsko ljetno vrijeme", "TMST": "Turkmenistansko ljetno vrijeme", "EDT": "Sjevernoameričko istočno ljetno vrijeme", "GFT": "Francuskogvajansko vrijeme", "AEST": "Istočnoaustralijsko standardno vrijeme"},
	}
}

// Locale returns the current translators string locale
func (bs *bs) Locale() string {
	return bs.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bs'
func (bs *bs) PluralsCardinal() []locales.PluralRule {
	return bs.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bs'
func (bs *bs) PluralsOrdinal() []locales.PluralRule {
	return bs.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bs'
func (bs *bs) PluralsRange() []locales.PluralRule {
	return bs.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (bs *bs) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14)) || (fMod10 >= 2 && fMod10 <= 4 && (fMod100 < 12 || fMod100 > 14)) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (bs *bs) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bs'
func (bs *bs) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := bs.CardinalPluralRule(num1, v1)
	end := bs.CardinalPluralRule(num2, v2)

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
func (bs *bs) MonthAbbreviated(month time.Month) string {
	return bs.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bs *bs) MonthsAbbreviated() []string {
	return bs.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bs *bs) MonthNarrow(month time.Month) string {
	return bs.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bs *bs) MonthsNarrow() []string {
	return bs.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bs *bs) MonthWide(month time.Month) string {
	return bs.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bs *bs) MonthsWide() []string {
	return bs.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bs *bs) WeekdayAbbreviated(weekday time.Weekday) string {
	return bs.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bs *bs) WeekdaysAbbreviated() []string {
	return bs.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bs *bs) WeekdayNarrow(weekday time.Weekday) string {
	return bs.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bs *bs) WeekdaysNarrow() []string {
	return bs.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bs *bs) WeekdayShort(weekday time.Weekday) string {
	return bs.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bs *bs) WeekdaysShort() []string {
	return bs.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bs *bs) WeekdayWide(weekday time.Weekday) string {
	return bs.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bs *bs) WeekdaysWide() []string {
	return bs.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bs' and handles both Whole and Real numbers based on 'v'
func (bs *bs) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bs' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bs *bs) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bs.percentSuffix...)

	b = append(b, bs.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bs'
func (bs *bs) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bs.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bs.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bs'
// in accounting notation.
func (bs *bs) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bs.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bs.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bs.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, bs.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bs.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bs.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bs.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bs'
func (bs *bs) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'bs'
func (bs *bs) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bs'
func (bs *bs) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bs'
func (bs *bs) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bs.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, bs.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(t.Year()*-1), 10)
	}

	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bs'
func (bs *bs) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bs'
func (bs *bs) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bs'
func (bs *bs) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bs'
func (bs *bs) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bs.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bs.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
