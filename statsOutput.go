package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type output struct {
	Totaltime     string `json:"total_time"`
	New_records   int    `json:"new_records"`
	Duplicates    int    `json:"dupes"`
	Total_records int    `json:"records"`
}

func newOutput(totalTime time.Duration, newRecs, dupes, totalRecs int) (stats *output) {
	stats = &output{
		Totaltime:     fmt.Sprintf("%v", totalTime),
		New_records:   newRecs,
		Duplicates:    dupes,
		Total_records: recs,
	}

	return
}

func (stats *output) printJson() {
	stats_bytes, _ := json.Marshal(stats)
	fmt.Println(string(stats_bytes))
}
