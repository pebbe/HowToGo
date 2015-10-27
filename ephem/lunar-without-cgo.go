package main

import (
	"fmt"
	"math"
	"time"
)

const (
	lunar_cycles_per_year = float64(365.25 / 29.530588853)
)

var (
	moonphase = [4]string{"nieuwemaan", "eerste kwartier", "vollemaan", "laatste kwartier"}
)

type Date struct {
	fract   float64
	D, M, Y int
}

func DateNew(year, month, day, hour, minute, second int) Date {
	return Date{
		D: day,
		M: month,
		Y: year,
		fract: float64(hour)/24 +
			float64(minute)/24/60 +
			float64(second)/24/60/60,
	}
}

func DateFromTime(t time.Time) Date {
	t0 := t.UTC()
	return DateNew(t0.Year(), int(t0.Month()), t0.Day(), t0.Hour(), t0.Minute(), t0.Second())
}

type Day float64

func (d Day) Local() time.Time {
	return d.UTC().Local()
}

func (d Day) UTC() time.Time {
	date := calendar_gregorian_from_absolute(float64(d))
	h := int(date.fract * 24)
	m := int(date.fract*24*60) % 60
	s := int(date.fract*24*60*60) % 60
	return time.Date(date.Y, time.Month(date.M), date.D, h, m, s, 0, time.UTC)
}

func degrees_to_radians(deg float64) float64 {
	return deg / 180 * math.Pi
}

func solar_sin_degrees(x float64) float64 {
	return math.Sin(degrees_to_radians(math.Mod(x, 360)))
}

func solar_cosine_degrees(x float64) float64 {
	return math.Cos(degrees_to_radians(math.Mod(x, 360)))
}

func calendar_astro_from_absolute(d Day) float64 {
	return float64(d) + 1721424.5
}

func calendar_absolute_from_gregorian(date Date) (day Day) {
	year := date.Y
	if year == 0 {
		panic("There was no year zero")
	}
	var dayI int
	if year > 0 {
		offset_years := year - 1
		dayI = calendar_day_number(date) +
			365*offset_years +
			offset_years/4 -
			offset_years/100 +
			offset_years/400
	} else {
		offset_years := -1 - year
		dayI = calendar_day_number(date) -
			365*offset_years -
			offset_years/4 +
			offset_years/100 -
			offset_years/400 -
			calendar_day_number(DateNew(-1, 12, 31, 0, 0, 0))
	}
	day = Day(float64(dayI) + date.fract)
	return
}

func calendar_last_day_of_month(month, year int) int {
	if month == 2 && calendar_leap_year_p(year) {
		return 29
	}
	return []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[month]
}

func calendar_gregorian_from_absolute(date float64) Date {
	dateI := int(math.Floor(date))
	fract := date - float64(dateI)
	d0 := dateI - 1
	n400 := d0 / 146097
	d1 := d0 % 146097
	n100 := d1 / 36524
	d2 := d1 % 36524
	n4 := d2 / 1461
	d3 := d2 % 1461
	n1 := d3 / 365
	day := (d3 % 365) + 1
	year := 400*n400 + 100*n100 + n4*4 + n1
	month := 1
	var mdays int
	if n100 == 4 || n1 == 4 {
		return DateNew(year, 12, 31, 0, 0, 0)
	}
	year += 1
	for {
		mdays = calendar_last_day_of_month(month, year)
		if mdays >= day {
			break
		}
		day -= mdays
		month++
	}
	return DateNew(year, month, day, int(fract * 24), int(fract * 24 * 60) % 60, int(fract * 24 * 60 * 60) % 60)
}

func calendar_leap_year_p(year int) bool {
	if year < 0 {
		year = -1 - year
	}
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func calendar_day_number(date Date) (day_of_year int) {
	day_of_year = int(date.D) + 31*(date.M-1)
	if date.M > 2 {
		day_of_year -= (4*date.M + 23) / 10
		if calendar_leap_year_p(date.Y) {
			day_of_year++
		}
	}
	return
}

// Returns index of last new moon before date, for date after 1-1-1900,
// or index of first new moon after date, for date before 1900
func lunar_index(date Date) int {
	// Years since 1900, as a float64
	years := float64(date.Y) + float64(calendar_day_number(date))/366 - 1900
	return 4 * int(lunar_cycles_per_year*years)
}

func lunar_phase(index int) (day Day, phase int) {
	phaseI := index % 4
	if phaseI < 0 {
		phaseI += 4
	}
	indexF := float64(index) / 4
	timeF := indexF / 1236.85
	timeF2 := timeF * timeF
	timeF3 := timeF2 * timeF
	dateF := float64(calendar_absolute_from_gregorian(Date{fract: .5, D: 0, M: 1, Y: 1900})) +
		0.75933 +
		29.53058868*indexF +
		0.0001178*timeF2 +
		-0.000000155*timeF3 +
		0.00033*solar_sin_degrees(166.56+132.87*timeF-0.009173*timeF2)
	sun_anomaly := math.Mod(359.2242+29.105356*indexF-0.0000333*timeF2-0.00000347*timeF3, 360)
	moon_anomaly := math.Mod(306.0253+385.81691806*indexF+0.0107306*timeF2+0.00001236*timeF3, 360)
	moon_lat := math.Mod(21.2964+390.67050646*indexF-0.0016528*timeF2-0.00000239*timeF3, 360)
	var adjustmentF float64
	if phaseI == 0 || phaseI == 2 {
		adjustmentF = (0.1734-0.000393*timeF)*solar_sin_degrees(sun_anomaly) +
			0.0021*solar_sin_degrees(2*sun_anomaly) +
			-0.4068*solar_sin_degrees(moon_anomaly) +
			0.0161*solar_sin_degrees(2*moon_anomaly) +
			-0.0004*solar_sin_degrees(3*moon_anomaly) +
			0.0104*solar_sin_degrees(2*moon_lat) +
			-0.0051*solar_sin_degrees(sun_anomaly+moon_anomaly) +
			-0.0074*solar_sin_degrees(sun_anomaly-moon_anomaly) +
			0.0004*solar_sin_degrees(2*moon_lat+sun_anomaly) +
			-0.0004*solar_sin_degrees(2*moon_lat-sun_anomaly) +
			-0.0006*solar_sin_degrees(2*moon_lat+moon_anomaly) +
			0.0010*solar_sin_degrees(2*moon_lat-moon_anomaly) +
			0.0005*solar_sin_degrees(2*moon_anomaly+sun_anomaly)
	} else {
		adjustmentF = (0.1721-0.0004*timeF)*solar_sin_degrees(sun_anomaly) +
			0.0021*solar_sin_degrees(2*sun_anomaly) +
			-0.6280*solar_sin_degrees(moon_anomaly) +
			0.0089*solar_sin_degrees(2*moon_anomaly) +
			-0.0004*solar_sin_degrees(3*moon_anomaly) +
			0.0079*solar_sin_degrees(2*moon_lat) +
			-0.0119*solar_sin_degrees(sun_anomaly+moon_anomaly) +
			-0.0047*solar_sin_degrees(sun_anomaly-moon_anomaly) +
			0.0003*solar_sin_degrees(2*moon_lat+sun_anomaly) +
			-0.0004*solar_sin_degrees(2*moon_lat-sun_anomaly) +
			-0.0006*solar_sin_degrees(2*moon_lat+moon_anomaly) +
			0.0021*solar_sin_degrees(2*moon_lat-moon_anomaly) +
			0.0003*solar_sin_degrees(2*moon_anomaly+sun_anomaly) +
			0.0004*solar_sin_degrees(sun_anomaly-2*moon_anomaly) +
			-0.0003*solar_sin_degrees(2*sun_anomaly+moon_anomaly)
	}
	adjF := 0.0028 +
		-0.0004*solar_cosine_degrees(sun_anomaly) +
		0.0003*solar_cosine_degrees(moon_anomaly)
	if phaseI == 1 {
		adjustmentF += adjF
	} else if phaseI == 2 {
		adjustmentF -= adjF
	}
	dateF += adjustmentF
	dateF -= solar_ephemeris_correction(
		calendar_gregorian_from_absolute(dateF).Y) / 60.0 / 24.0
	return Day(dateF), phaseI
}

func solar_ephemeris_correction(year int) float64 {
	if year >= 1988 && year < 2020 {
		return float64(year-2000+67) / (60 * 60 * 24)
	}
	if year >= 1900 && year < 1988 {
		theta := (calendar_astro_from_absolute(calendar_absolute_from_gregorian(DateNew(year, 7, 1, 0, 0, 0))) -
			calendar_astro_from_absolute(calendar_absolute_from_gregorian(DateNew(1900, 1, 1, 0, 0, 0)))) / 36525
		theta2 := theta * theta
		theta3 := theta2 * theta
		theta4 := theta2 * theta2
		theta5 := theta3 * theta2
		return -0.00002 +
			0.000297*theta +
			0.025184*theta2 +
			-0.181133*theta3 +
			0.553040*theta4 +
			-0.861938*theta5 +
			0.677066*theta3*theta3 +
			-0.212591*theta4*theta3
	}
	if year >= 1800 && year < 1900 {
		theta := (calendar_astro_from_absolute(calendar_absolute_from_gregorian(DateNew(year, 7, 1, 0, 0, 0))) -
			calendar_astro_from_absolute(calendar_absolute_from_gregorian(DateNew(1900, 1, 1, 0, 0, 0)))) / 36525
		theta2 := theta * theta
		theta3 := theta2 * theta
		theta4 := theta2 * theta2
		theta5 := theta3 * theta2
		return -0.000009 +
			0.003844*theta +
			0.083563*theta2 +
			0.865736*theta3 +
			4.867575*theta4 +
			15.845535*theta5 +
			31.332267*theta3*theta3 +
			38.291999*theta4*theta3 +
			28.316289*theta4*theta4 +
			11.636204*theta4*theta5 +
			2.043794*theta5*theta5
	}
	if year >= 1620 && year < 1800 {
		x := float64(year-1600) / 10
		return (2.19167*x*x - 40.675*x + 196.58333) / (60 * 60 * 24)
	}
	// if year < 1620 || year >= 2020
	tmp := calendar_astro_from_absolute(calendar_absolute_from_gregorian(DateNew(year, 1, 1, 0, 0, 0))) - 2382148
	second := tmp*tmp/41048480 - 15
	return second / (60 * 60 * 24)
}

func main() {

	date := DateFromTime(time.Now())
	//date = DateNew(1600, 1, 1, 0, 0, 0)
	day0 := calendar_absolute_from_gregorian(date)
	idx := lunar_index(date)
	var d Day
	var f int
	if idx < 0 {
		for {
			d1, f1 := lunar_phase(idx)
			if d1 < day0 {
				idx++
				break
			}
			d, f = d1, f1
			idx--
		}
	} else {
		for {
			idx += 1
			d, f = lunar_phase(idx)
			if d > day0 {
				break
			}
		}
	}
	fmt.Println(moonphase[f], d.Local())

	//fmt.Printf("%g\n%g\n\n", 0.0029898130127028167, solar_ephemeris_correction(2100))
	//fmt.Printf("%g\n%g\n\n", 0.000775462962962963, solar_ephemeris_correction(2000))
	//fmt.Printf("%g\n%g\n\n", 0.0003330603524900667, solar_ephemeris_correction(1950))
	//fmt.Printf("%g\n%g\n\n", 8.461254094840995e-05, solar_ephemeris_correction(1850))
	//fmt.Printf("%g\n%g\n\n", 0.00010417048611111067, solar_ephemeris_correction(1700))
	//fmt.Printf("%g\n%g\n\n", 0.0014851565792882565, solar_ephemeris_correction(1600))

}
