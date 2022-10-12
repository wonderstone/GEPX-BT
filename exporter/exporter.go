package exporter

import (
	// "fmt"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/wonderstone/QuantTools/account/virtualaccount"
	"gopkg.in/yaml.v3"
)

// export realtime yaml file in config dir
func ExportRealtimeYaml(configDir string, sec string, va virtualaccount.VAcct, AInfo interface{}) {
	// read BackTest configuration from file  viper is not thread safe
	viper.SetConfigName("BackTest")
	viper.AddConfigPath(configDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})

	// Add Virtual account fields
	m["VA"] = va
	// Add additional fields
	m["AFields"] = AInfo
	// Add Data fields
	tdm := make(map[string]interface{})
	// fmt.Println(viper.GetString("SMName"))
	tmpMap := viper.GetStringMap(sec)
	var sinstrnames []string
	for _, v := range tmpMap["sinstrnames"].([]interface{}) {
		sinstrnames = append(sinstrnames, v.(string))
	}
	tdm["sinstrnames"] = sinstrnames
	var sindinames []string
	for _, v := range tmpMap["sindinames"].([]interface{}) {
		sindinames = append(sindinames, v.(string))
	}
	tdm["sindinames"] = sindinames
	var scsvdatafields []string
	for _, v := range tmpMap["scsvdatafields"].([]interface{}) {
		scsvdatafields = append(scsvdatafields, v.(string))
	}
	tdm["scsvdatafields"] = scsvdatafields
	var finstrnames []string
	for _, v := range tmpMap["finstrnames"].([]interface{}) {
		finstrnames = append(finstrnames, v.(string))
	}
	tdm["finstrnames"] = finstrnames
	var findinames []string
	for _, v := range tmpMap["findinames"].([]interface{}) {
		findinames = append(findinames, v.(string))
	}
	tdm["findinames"] = findinames
	var fcsvdatafields []string
	for _, v := range tmpMap["fcsvdatafields"].([]interface{}) {
		fcsvdatafields = append(fcsvdatafields, v.(string))
	}
	tdm["fcsvdatafields"] = fcsvdatafields
	m["DataFields"] = tdm

	// #  Section for ContractProp
	tCPm := make(map[string]interface{})
	tCPm["ConfName"] = tmpMap["confname"]
	tCPm["CPDataDir"] = tmpMap["cpdatadir"]
	m["ContractProp"] = tCPm
	// #  Section for Matcher parameter
	tMPm := make(map[string]interface{})
	tMPm["MatcherSlippage4S"] = tmpMap["matcherslippage4s"]
	tMPm["MatcherSlippage4F"] = tmpMap["matcherslippage4f"]
	m["MatcherParam"] = tMPm
	// #  Section for Performance Analytics Parameter
	tPAm := make(map[string]interface{})
	tPAm["RiskFreeRate"] = tmpMap["riskfreerate"]
	tPAm["PAType"] = tmpMap["patype"]
	m["PA"] = tPAm
	// #  Section for Strategy Module Selection
	tSMm := make(map[string]interface{})
	tSMm["StrategyModule"] = tmpMap["strategymodule"]
	tSMm["SMGEPType"] = tmpMap["smgeptype"]
	tSMm["SMName"] = tmpMap["smname"]
	tSMm["SMDataDir"] = tmpMap["smdatadir"]
	m["StgModel"] = tSMm
	// export yaml file with yaml.v3
	data, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal().Msg(err.Error())

	}
	err2 := ioutil.WriteFile("./realtime.yaml", data, 0777)
	if err2 != nil {
		log.Fatal().Msg(err2.Error())
	}
	// fmt.Println("data written")
}

// export the simplified Karva expression to refactor the realtime expression trees(ETs)
func ExportSKE(configDir string, sec string, KES interface{}) {
	// read BackTest configuration from file  viper is not thread safe
	viper.SetConfigName("GEP")
	viper.AddConfigPath(configDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// fmt.Println(viper.GetString("SMName"))
	tmpMap := viper.GetStringMap(sec)
	m := make(map[string]interface{})

	// Add Data fields
	// make a slice to store all the function names
	var funcnames []string
	// fmt.Println(tmpMap["funcweight"])
	for _, v := range tmpMap["funcweight"].([]interface{}) {
		tmp := v.([]interface{})
		// fmt.Println(tmp[0])
		funcnames = append(funcnames, tmp[0].(string))
	}
	m["FuncNames"] = funcnames
	// data fields
	m["HeadSize"] = tmpMap["headsize"]
	m["numConstants"] = tmpMap["numconstants"]
	m["linkFunc"] = tmpMap["linkfunc"]
	m["Mode"] = tmpMap["mode"]
	//
	m["KES"] = KES
	// export yaml file with yaml.v3
	data, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	err2 := ioutil.WriteFile("./KarvaExp.yaml", data, 0777)
	if err2 != nil {
		log.Fatal().Msg(err2.Error())
	}
	// fmt.Println("data written")
}
