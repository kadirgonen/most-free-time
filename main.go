package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	strArr1 := []string{"10:00AM-12:30PM", "02:00PM-02:45PM", "09:10AM-09:50AM"}
	strArr2 := []string{"12:15PM-02:00PM", "09:00AM-10:00AM", "10:30AM-12:00PM"}
	strArr3 := []string{"12:15PM-02:00PM", "09:00AM-12:11PM", "02:02PM-04:00PM"}
	fmt.Println(MostFreeTime(strArr1))
	fmt.Println(MostFreeTime(strArr2))
	fmt.Println(MostFreeTime(strArr3))

}

type Event struct {
	start int
	end   int
}

func MostFreeTime(strArr []string) string {
	var events []Event
	var max = math.MinInt
	for _, v := range strArr {
		times := strings.Split(v, "-")
		event := &Event{start: minuteOfDay(times[0]),
			end: minuteOfDay(times[1])}
		events = append(events, *event)

	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].start < events[j].start
	})

	for j := 0; j < len(events)-1; j++ {
		if events[j+1].start-events[j].end > max {
			max = events[j+1].start - events[j].end
		}
	}
	return convertMinToStr(max)

}

func convertMinToStr(minute int) string {
	var hours, mins int
	mins = minute % 60
	if minute/60 > 0 {
		hours = int(math.Floor(float64(minute / 60)))
	}
	return fmt.Sprintf("%0.2d:%0.2d", hours, mins)
}

func minuteOfDay(time string) int {
	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:5])
	mid := time[5:]

	var result int
	if hours == 12 {
		hours = 0
	}
	if mid == "PM" {
		result = ((12 + hours) * 60) + minutes
	} else {
		result = (hours * 60) + minutes
	}
	return result
}
