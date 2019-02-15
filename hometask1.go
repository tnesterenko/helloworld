package main

import (
	"bufio"
	"encoding/csv"
	"fmt" 
	"io"
	"os"
	"strconv"
	"time"
)

const dateMask = "2006-01-02 15:04:05"

const startHour = 7 // 10 утра по Москве
const finHour = 0   // 3 часа утра по Москве

type trade struct {
	code  string
	price float64
	count string // не важно для задачи
	time  time.Time
}

type ohcl struct {
	code       string
	time       time.Time
	o, h, c, l float64
}

type date struct {
	year  int
	month time.Month
	day   int
}

func main() {

	var res5, res30, res240 []ohcl

	trades, dates := readerCsvFile("trades.csv")
	fmt.Printf("dates %v ; format %T \n", dates, dates)
	res5, res30, res240 = initRes(dates)
	fmt.Printf("res5 %v ; format %T \n", res5, res5)
	fmt.Printf("res30 %v ; format %T \n", res30, res30)
	fmt.Printf("res240 %v ; format %T \n", res240, res240)
	// fmt.Printf("%v \n", trades)
	for _, tr := range trades {
		// tr.time && tr.time.Hour < startHour
		// fmt.Printf("tr.time.Hour %v ; format %T \n", tr.time.Hour(), tr.time.Hour())

		if tr.time.Hour() >= finHour && tr.time.Hour() < startHour { // биржа закрыта
			continue
		}
		fmt.Printf("tr.time %v ; format %T \n", tr.time, tr.time)
		fmt.Printf("tr.price %v ; format %T \n", tr.price, tr.price)

		break
	}

}

func initRes (dates map[date]bool) ([]ohcl, []ohcl, []ohcl) {
	var res5, res30, res240 []ohcl
	var ohclTemp ohcl
	for d:= range dates{
		d1 := time.Date(d.year, d.month, d.day, startHour, 0, 0, 0, time.UTC)
		d2 := time.Date(d.year, d.month, d.day+1, finHour, 0, 0, 0, time.UTC)

		fmt.Printf("d1 %v ; format %T \n", d1, d1)
		fmt.Printf("d2 %v ; format %T \n", d2, d2)

		ohclTemp.time = d1
		for i:=0;;i++{
			// fmt.Printf("i %v ; format %T \n", i, i)
			if (i>10){
				break
			}
			ohclTemp.time = ohclTemp.time.Add(time.Minute * time.Duration(5*i))
			// fmt.Printf("ohclTemp.time %v ; format %T \n", ohclTemp.time, ohclTemp.time)
			if(!ohclTemp.time.Before(d2)){
				break
			}
			res5= append(res5,ohclTemp)
			res30= append(res5,ohclTemp)
			res240= append(res5,ohclTemp)
		}
	}
	return  res5, res30, res240
}

func isTimeInPeriod(t, tStart, tEnd time.Time) bool {
	if t.Before(tEnd) && t.After(tStart) {
		return true
	} 
	return false
}

// func startExchange(t time.Time) time.Time {
// 	year, month, day := d.Date()
// }
func readerCsvFile(filename string) ([]trade, map[date]bool) {
	var trades []trade
	var t time.Time
	var p float64
	var curDate date
	dates := make(map[date]bool)
	csvFile, _ := os.Open(filename)
	// TODO обработка ошибки
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		// fmt.Printf("line: %s  === type %T \n", line, line)
		if error == io.EOF {
			// TODO обработка ошибки
			break
		} else if error != nil {
			fmt.Println(error)
		}
		// fmt.Printf("line[3] %v %T\n", line[3], line[3])
		// t1, _ = time.Parse(time.RFC3339, line[3])
		// fmt.Printf("t1 %+v ; format %T \n", t1, t1)

		t, error = time.Parse(dateMask, line[3])
		if error != nil {
			fmt.Println(error)
		}

		curDate.year, curDate.month, curDate.day = t.Date()
		if !dates[curDate] {
			dates[curDate] = true
		}

		p, error = strconv.ParseFloat(line[1], 64)
		if error != nil {
			fmt.Println(error)
		}
		// t = stringToTime(line[3])
		// fmt.Printf("t %+v ; format %T \n", t, t)
		trades = append(trades, trade{
			code:  line[0],
			price: p,
			count: line[2],
			time:  t,
		})
		// break
	}
	return trades, dates
}

// func stringToTime(s string) time.Time {
// 	// var year int
// 	t, _ := strconv.ParseInt(s[5:7], 10, 0)
// 	fmt.Printf("value: %v , type: %T \n", t, t)
// 	year, _ := strconv.ParseInt(s[0:4], 10, 0)
// 	// month, _ := strconv.ParseInt(s[5:7], 10, 0)
// 	day, _ := strconv.ParseInt(s[8:10], 10, 0)
// 	hour, _ := strconv.ParseInt(s[11:13], 10, 0)
// 	min, _ := strconv.ParseInt(s[14:16], 10, 0)
// 	sec, _ := strconv.ParseInt(s[17:19], 10, 0)
// 	time.Date(int(year), 1, int(day), int(hour), int(min), int(sec), 0, time.UTC)
// 	return time.Now()
// }
