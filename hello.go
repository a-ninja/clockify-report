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

type ReportType struct {
	ReportName string
	StartDate  string
	EndDate    string
}

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

	clients := getRates()

	now := time.Now().UTC()
	year, month, day := now.Date()
	endDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	origEndDay := endDay
	fmt.Println("year:>", year, ", month=", month, ", day=", day)
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

	layout := `"2006-01-02T15:04:05.000Z"`
	endDate := endDay.Format(layout)

	startDayForWeeklyReport := endDay.AddDate(0, 0, -6).Format(layout)
	// fmt.Println("endDay:>", endDay, ", day=", weekDay, ", startDay=", startDayForWeeklyReport)
	startDayForMonthlyReport := origEndDay.AddDate(0, 0, 0-day+1).Format(layout)
	// startDayForMonthlyReport = "\"2024-04-01T00:00:00.000Z\""
	// fmt.Println("endDay:>", endDay, ", day=", weekDay, ", startDay=", startDayForMonthlyReport)
	weeklyReport := ReportType{ReportName: "last-week-from-Sun-Sat", StartDate: startDayForWeeklyReport, EndDate: endDate}
	monthlyReport := ReportType{ReportName: "monthly", StartDate: startDayForMonthlyReport, EndDate: endDate}

	run := true
	if run == true {
		var reportTypes [2]ReportType
		reportTypes[0] = weeklyReport
		reportTypes[1] = monthlyReport
		for _, client := range clients {
			fmt.Println("Client:", client.name, client.id)
		}

		for _, reportType := range reportTypes {
			fmt.Println("\nendDate:>", reportType.EndDate, ", startDate", reportType.StartDate, ", reportType", reportType.ReportName)
			getSummary(clients, reportType)
		}
	}
}

// func getSummary(clients map[string]Client, startDate string, endDate string, reportType string) {
func getSummary(clients map[string]Client, reportType ReportType) {
	startDate := reportType.StartDate
	endDate := reportType.EndDate
	reportName := reportType.ReportName

	workspaceId := getWorkspaceId()
	//TODO: there should be a better way to create array of strings
	var clientIds string = `[`
	for _, ptClient := range clients {
		if clientIds != `[` {
			clientIds = clientIds + `,`
		}
		clientIds = clientIds + `"` + ptClient.id + `"`
	}
	clientIds = clientIds + `]`
	// clientIds := []string{"64026d86264092281bfbcaa6", "64bf322e40486b3fa56d19fe", "64026d8b2b547d4bb3880da7", "640a8179ef9f495fb7aa90b9", "659383608f580f174cafe8fa"}
	// clientIdStr, err := json.Marshal(clientIds)
	// fmt.Println("clientIds:", clientIds)
	// clientIds1 := ["64bf322e40486b3fa56d19fe","64026d86264092281bfbcaa6","640a8179ef9f495fb7aa90b9","64026d8b2b547d4bb3880da7","659383608f580f174cafe8fa"]

	url := "https://reports.api.clockify.me/v1/workspaces/" + workspaceId + "/reports/summary"
	// var str = `{"dateRangeStart":"2024-03-31T00:00:00.000Z","dateRangeEnd":"2024-04-06T23:59:59.999Z","sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":["64026d86264092281bfbcaa6","64bf322e40486b3fa56d19fe"],"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["PROJECT","USER"],"summaryChartType":"BILLABILITY"}}`
	// str = `{"dateRangeStart":"2024-03-31T00:00:00.000Z","dateRangeEnd":"2024-04-06T23:59:59.999Z","sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":["64026d86264092281bfbcaa6","64bf322e40486b3fa56d19fe","64026d8b2b547d4bb3880da7","640a8179ef9f495fb7aa90b9"],"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["PROJECT","USER"],"summaryChartType":"BILLABILITY"}}`
	var str = `{"dateRangeStart":` + startDate + `,"dateRangeEnd":` + endDate + `,"sortOrder":"ASCENDING","description":"","rounding":false,"withoutDescription":false,"amounts":["EARNED"],"amountShown":"EARNED","zoomLevel":"WEEK","userLocale":"en-US","customFields":null,"userCustomFields":null,"kioskIds":[],"clients":{"contains":"CONTAINS","ids":` + clientIds + `,"status":"ACTIVE","numberOfDeleted":0},"summaryFilter":{"sortColumn":"GROUP","groups":["CLIENT","PROJECT","USER"],"summaryChartType":"BILLABILITY"}}`

	var jsonStr = []byte(str)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	apiKey := getApiKey()
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//     fmt.Println("result:", result)
	var total = 0.0

	file, err := os.Create("report-" + reportName + ".csv")
	fmt.Println("filename", file.Name(), reportName)
	if err != nil {
		fmt.Println("failed creating file: %s", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	var data [][]string

	totalDuration := 0.0
	total = 0.0
	totalEquity := 0.0
	projectTotalDuration := 0.0
	projectTotal := 0.0
	projectTotalEquity := 0.0

	row := []string{"Duration", startDate + "-" + endDate}
	data = append(data, row)
	row = []string{"Client", "Project", "Resource", "Duration", "Cash Rate", "Total Cash", "Equity Rate", "Total Equity"}
	data = append(data, row)

	for _, clientEntry := range result.Clients {
		clientName := clientEntry.ClientName
		totalDuration = 0.0
		total = 0.0
		totalEquity = 0.0
		ptClient := clients[clientName]
		fmt.Println("\n" + clientName + ">>")

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
			row = []string{"Project TOTAL - " + clientName, projectName, "", fmt.Sprintf("%f", projectTotalDuration), "-", fmt.Sprintf("$%f", projectTotal), "-", fmt.Sprintf("$%f", projectTotalEquity)}
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
	fmt.Println("\n ================Report has been exported to", file.Name())

}
