package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Total struct {
	TotalTime         int `json:"totalTime"`
	TotalBillableTime int `json:"totalBillableTime"`
}

// Generated go struct
type Response struct {
	Total   []Total            `json:"totals"`
	Clients []ClientGroupEntry `json:"groupOne"`
}

type ClientGroupEntry struct {
	Duration   int64               `json:"duration"`
	ClientName string              `json:"name"`
	Projects   []ProjectGroupEntry `json:"children"`
}

type ProjectGroupEntry struct {
	Duration    int64        `json:"duration"`
	ProjectName string       `json:"name"`
	People      []GroupEntry `json:"children"`
}

type GroupEntry struct {
	Duration int64  `json:"duration"`
	Name     string `json:"name"`
}

type Rate struct {
	resource string  ""
	rate     float64 ""
}

type Client struct {
	name          string
	projectIds    string
	id            string
	rateMap       map[string]float64
	equityRateMap map[string]float64
	equityEnabled bool
	grouped       bool
}

func main() {
	int18RateMap := make(map[string]float64)

	clientRateMap := make(map[string]float64)

	wxRateMap := make(map[string]float64)

	clientPokercowsCashRateMap := make(map[string]float64)

	clientPokercowsEquityRateMap := make(map[string]float64)

	wb2ProjectIds := `["64026dae21b9ba7d2f298f4b"]`
	clientInt18 := Client{name: "int18", id: "64bf322e40486b3fa56d19fe", rateMap: int18RateMap, grouped: true, equityEnabled: false}
	clientElationWb2 := Client{name: "Elation", id: "64026d86264092281bfbcaa6", projectIds: wb2ProjectIds, rateMap: clientRateMap, grouped: false, equityEnabled: false}
	clientWx := Client{name: "WeatherSTEM", id: "640a8179ef9f495fb7aa90b9", rateMap: wxRateMap, grouped: true, equityEnabled: false}
	clientPokercows := Client{name: "Poker Cows", id: "64026d8b2b547d4bb3880da7", rateMap: clientPokercowsCashRateMap, grouped: true, equityEnabled: true, equityRateMap: clientPokercowsEquityRateMap}

	clients := make(map[string]Client)
	clients[clientInt18.name] = clientInt18
	clients[clientElationWb2.name] = clientElationWb2
	clients[clientWx.name] = clientWx
	clients[clientPokercows.name] = clientPokercows

	// workspaceId := ""

	// url := "https://reports.api.clockify.me/v1/workspaces/"+workspaceId+"/reports/summary"
	// fmt.Println("URL:>", url)

	// startDate1 := `"2024-04-07T00:00:00.000Z"`
	// endDate1 := `"2024-04-13T00:00:00.000Z"`
	// reportType := "month"
	now := time.Now().UTC()
	year, month, day := now.Date()
	endDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	fmt.Println("year:>", year, ", month=", month, ", day1=", day)
	weekDay := endDay.Weekday()
	fmt.Println("now:>", endDay, ", day=", weekDay)
	if weekDay == time.Sunday {
		endDay = endDay.AddDate(0, 0, -1)
	} else if weekDay == time.Monday {
		endDay = endDay.AddDate(0, 0, -2)
	} else if weekDay == time.Tuesday {
		endDay = endDay.AddDate(0, 0, -3)
	} else if weekDay == time.Wednesday {
		endDay = endDay.AddDate(0, 0, -4)
	} else if weekDay == time.Thursday {
		endDay = endDay.AddDate(0, 0, -5)
	} else if weekDay == time.Friday {
		endDay = endDay.AddDate(0, 0, -6)
	}
	startDay := endDay.AddDate(0, 0, -6)
	fmt.Println("endDay:>", endDay, ", day=", weekDay, ", startDay=", startDay)
	layout := `"2006-01-02T15:04:05.000Z"`
	// fmt.Println("existing endDate:>", endDate1, ", startDate=", startDate1, ", reportType=", reportType)
	startDate := startDay.Format(layout)
	endDate := endDay.Format(layout)
	reportType := "month"
	/*if endDate1 == endDate {
		fmt.Println("endDate match")
	} else {
		fmt.Println("endDate dont match", endDate, endDate1)
	}
	if startDate1 == startDate {
		fmt.Println("startDate match")
	} else {
		fmt.Println("startDate dont match", startDate, startDate1)
	}*/

	fmt.Println("endDate:>", endDate, ", startDate=", startDate, ", reportType=", reportType)
	run := true
	if run == true {
		//         for _, client := range clients {
		getSummary(clients, startDate, endDate, reportType)
		//             time.Sleep(1 * time.Second)
		//         }
	}
	// clients = []Client{clientInt18}
}

func getSummary(clients map[string]Client, startDate string, endDate string, reportType string) {
	workspaceId := "64026d55264092281bfbc652"
	url := "https://reports.api.clockify.me/v1/workspaces/" + workspaceId + "/reports/summary"
	//     fmt.Println("URL:>", url)

	//     clientIds := `["`+ptClient.id+`"]`

	var str = `{"dateRangeStart":"2024-03-31T00:00:00.000Z","dateRangeEnd":"2024-04-06T23:59:59.999Z","sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":["64026d86264092281bfbcaa6","64bf322e40486b3fa56d19fe"],"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["PROJECT","USER"],"summaryChartType":"BILLABILITY"}}`
	str = `{"dateRangeStart":"2024-03-31T00:00:00.000Z","dateRangeEnd":"2024-04-06T23:59:59.999Z","sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":["64026d86264092281bfbcaa6","64bf322e40486b3fa56d19fe","64026d8b2b547d4bb3880da7","640a8179ef9f495fb7aa90b9"],"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["PROJECT","USER"],"summaryChartType":"BILLABILITY"}}`
	str = `{"dateRangeStart":` + startDate + `,"dateRangeEnd":` + endDate + `,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":["64026d86264092281bfbcaa6","64bf322e40486b3fa56d19fe","64026d8b2b547d4bb3880da7","640a8179ef9f495fb7aa90b9"],"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["CLIENT","PROJECT","USER"],"summaryChartType":"BILLABILITY"}}`

	/* if len(ptClient.projectIds) >0 {
	       str = `{"dateRangeStart":`+startDate+`,"dateRangeEnd":`+endDate+`,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":`+clientIds+`,"status":"ACTIVE","numberOfDeleted":0},"projects":{"contains":"CONTAINS","ids":`+ptClient.projectIds+`,"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["USER"],"summaryChartType":"BILLABILITY"}}`
	   } else {
	       str = `{"dateRangeStart":`+startDate+`,"dateRangeEnd":`+endDate+`,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":`+clientIds+`,"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["USER"],"summaryChartType":"BILLABILITY"}}`
	   } */

	/*
	   if ptClient.name == "wb2" {
	       str = `{"dateRangeStart":`+startDate+`,"dateRangeEnd":`+endDate+`,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":`+clientIds+`,"status":"ACTIVE","numberOfDeleted":0},"projects":{"contains":"CONTAINS","ids":`+ptClient.projectIds+`,"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["USER"],"summaryChartType":"BILLABILITY"}}`
	   } else {
	       str = `{"dateRangeStart":`+startDate+`,"dateRangeEnd":`+endDate+`,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":`+clientIds+`,"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["USER"],"summaryChartType":"BILLABILITY"}}`
	   }
	*/

	var jsonStr = []byte(str)

	//var jsonStr = []byte(`{"dateRangeStart":`+startDate+`,"dateRangeEnd":`+endDate+`,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":["64bf322e40486b3fa56d19fe"],"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["USER"],"summaryChartType":"BILLABILITY"}}`)
	//     fmt.Println(str)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	apiKey := ""
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//     fmt.Println("response Status:", resp.Status)
	//     fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	//     fmt.Println("response Body:", string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//     fmt.Println("result:", result)
	var total = 0.0

	file, err := os.Create("report-" + reportType + ".csv")
	if err != nil {
		fmt.Println("failed creating file: %s", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	var data [][]string

	//     var prevClient = ""
	totalDuration := 0.0
	total = 0.0
	totalEquity := 0.0
	projectTotalDuration := 0.0
	projectTotal := 0.0
	projectTotalEquity := 0.0

	row := []string{"Client", "Project", "Resource", "Duration", "Cash Rate", "Total Cash", "Equity Rate", "Total Equity"}
	data = append(data, row)

	for _, clientEntry := range result.Clients {
		clientName := clientEntry.ClientName
		totalDuration = 0.0
		total = 0.0
		totalEquity = 0.0
		ptClient := clients[clientName]
		for _, projectEntry := range clientEntry.Projects {
			projectName := projectEntry.ProjectName
			projectTotalDuration = 0.0
			projectTotal = 0.0
			projectTotalEquity = 0.0
			for _, element := range projectEntry.People {
				duration := float64(element.Duration) / 3600.0
				totalDuration = totalDuration + duration
				projectTotalDuration = projectTotalDuration + duration
				rate := clients[clientName].rateMap[element.Name]
				equityRate := clients[clientName].equityRateMap[element.Name]

				cost := duration * float64(rate)
				total = total + cost
				projectTotal = projectTotal + cost
				equityCost := 0.0
				if ptClient.equityEnabled {
					equityCost = duration * float64(equityRate)
					totalEquity = totalEquity + equityCost
					projectTotalEquity = projectTotalEquity + equityCost
				}
				row = []string{clientName, "", element.Name, fmt.Sprintf("%f", duration), fmt.Sprintf("$%f", rate), fmt.Sprintf("$%f", cost), fmt.Sprintf("$%f", equityRate), fmt.Sprintf("$%f", equityCost)}
				data = append(data, row)
			}
			row := []string{"Project TOTAL - " + clientName, projectName, "", fmt.Sprintf("%f", projectTotalDuration), "-", fmt.Sprintf("$%f", projectTotal), "-", fmt.Sprintf("$%f", projectTotalEquity)}
			data = append(data, row)
			fmt.Println(clientEntry.ClientName, projectEntry.ProjectName, projectTotalDuration, projectTotal, projectTotalEquity, ", Duration:", totalDuration, "total:", total, ", total Equity:", totalEquity)
		}
		row := []string{"TOTAL-" + clientName, "", "", fmt.Sprintf("%f", totalDuration), "-", fmt.Sprintf("$%f", total), "-", fmt.Sprintf("$%f", totalEquity)}
		data = append(data, row)
		row = []string{"", "", "", "", "", "", "", ""}
		data = append(data, row)
	}
	if err := w.WriteAll(data); err != nil {
		fmt.Println("error writing record to file", err)
	}

}
