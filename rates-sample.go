package main

func getWorkspaceIdSAMPLE() string {
	return ""
}

func getApiKeySAMPLE() string {

	return ""
}

func getRatesSAMPLE() map[string]Client {

	int18RateMap := make(map[string]float64)
	int18RateMap["Alex"] = 0.00
	int18RateMap["Andrew \"jamoozy\" Sabisch"] = 0.00
	int18RateMap["Evan Mallory"] = 0.00
	int18RateMap["Ian Ma"] = 0.00
	int18RateMap["James Dowdell"] = 0.00
	int18RateMap["Kalp Vyas"] = 0.00
	int18RateMap["Karan Panwar"] = 0.00
	int18RateMap["Mahesh Paliwal"] = 0.00
	int18RateMap["Matt Metzger"] = 0.00
	int18RateMap["Mike Smaili"] = 0.00
	int18RateMap["Ned Borisov"] = 0.00
	int18RateMap["Neeraj Jain"] = 0.00
	int18RateMap["Shannon"] = 0.00
	int18RateMap["sunim"] = 0.00
	int18RateMap["Susan"] = 0.00
	int18RateMap["Vishal Kodia"] = 0.00

	wb2ProjectIds := `["64026dae21b9ba7d2f298f4b"]`
	elationRateMap := make(map[string]float64)
	elationRateMap["Alex"] = 0.00
	elationRateMap["Antonio Trincao"] = 0.00
	elationRateMap["Cyrus"] = 0.00
	elationRateMap["Ian Ma"] = 0.00
	elationRateMap["James Dowdell"] = 0
	elationRateMap["Jeff White"] = 0.00
	elationRateMap["Jon Bach"] = 0.00
	elationRateMap["Kyle Holmberg"] = 0.00
	elationRateMap["Matt Metzger"] = 0.00
	elationRateMap["Mike Smaili"] = 0.00
	elationRateMap["Mukesh Shrimali"] = 0.00
	elationRateMap["Neeraj Jain"] = 0.00
	elationRateMap["Nihal Sharma"] = 0.00
	elationRateMap["Susan Somerset"] = 0.00
	elationRateMap["Thom Clark"] = 0.00
	elationRateMap["Vijay Prema"] = 0.00
	elationRateMap["Winston Denny"] = 0.00

	wxRateMap := make(map[string]float64)
	wxRateMap["Aditya Jain"] = 0.00
	wxRateMap["Alex"] = 0.00
	wxRateMap["Conrad Kreyling"] = 0.00
	wxRateMap["Cristopher Caceres Camargo"] = 0.00
	wxRateMap["Ian Ma"] = 0.00
	wxRateMap["Jake Tobin"] = 0.00
	wxRateMap["James Dowdell"] = 0.00
	wxRateMap["Jon Bach"] = 0.00
	wxRateMap["kalp vyas"] = 0.00
	wxRateMap["Ned Borisov"] = 0.00
	wxRateMap["Neeraj Jain"] = 0.00
	wxRateMap["sunim"] = 0.00
	wxRateMap["Vandana kumawat"] = 0.00

	pokercowsCashRateMap := make(map[string]float64)
	pokercowsCashRateMap["Alex"] = 0.00
	pokercowsCashRateMap["Antonio"] = 0.00
	pokercowsCashRateMap["Ian Ma"] = 0.00
	pokercowsCashRateMap["James Dowdell"] = 0.00
	pokercowsCashRateMap["Jon Bach"] = 0.00
	pokercowsCashRateMap["kalp vyas"] = 0.00
	pokercowsCashRateMap["Matt Metzger"] = 0.00
	pokercowsCashRateMap["Marissa Li"] = 0.00
	pokercowsCashRateMap["morris wu"] = 0.00
	pokercowsCashRateMap["Ned Borisov"] = 0.00
	pokercowsCashRateMap["Neeraj Jain"] = 0.00

	pokercowsEquityRateMap := make(map[string]float64)
	pokercowsEquityRateMap["Alex"] = 0.00
	pokercowsEquityRateMap["Antonio"] = 0.00
	pokercowsEquityRateMap["Ian Ma"] = 0.00
	pokercowsEquityRateMap["James Dowdell"] = 0.00
	pokercowsEquityRateMap["Jon Bach"] = 0.00
	pokercowsEquityRateMap["kalp vyas"] = 0.00
	pokercowsEquityRateMap["Matt Metzger"] = 0.00
	pokercowsEquityRateMap["Marissa Li"] = 0.00
	pokercowsEquityRateMap["morris wu"] = 0.00
	pokercowsEquityRateMap["Ned Borisov"] = 0.00
	pokercowsEquityRateMap["Neeraj Jain"] = 0.00

	clientInt18 := Client{name: "int18", id: "64bf322e40486b3fa56d19fe", rateMap: int18RateMap, grouped: true, equityEnabled: false}
	clientElationWb2 := Client{name: "Elation", id: "64026d86264092281bfbcaa6", projectIds: wb2ProjectIds, rateMap: elationRateMap, grouped: false, equityEnabled: false}
	clientWx := Client{name: "WeatherSTEM", id: "640a8179ef9f495fb7aa90b9", rateMap: wxRateMap, grouped: true, equityEnabled: false}
	clientPokercows := Client{name: "Poker Cows", id: "64026d8b2b547d4bb3880da7", rateMap: pokercowsCashRateMap, grouped: true, equityEnabled: true, equityRateMap: pokercowsEquityRateMap}

	clients := make(map[string]Client)
	clients[clientInt18.name] = clientInt18
	clients[clientElationWb2.name] = clientElationWb2
	clients[clientWx.name] = clientWx
	clients[clientPokercows.name] = clientPokercows

	return clients
}
