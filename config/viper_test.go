package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_viper(t *testing.T) {
	fmt.Println("test viper")

	viper.SetConfigName("ContractProp")
	viper.AddConfigPath("./Manual")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	tmpMap1 := viper.GetStringMap("TARGETPROP.STOCK.ST")
	ContractSize := tmpMap1["contractsize"]
	fmt.Println(ContractSize)
	assert.Equal(t, float64(100), tmpMap1["contractsize"].(float64))
	// read all config file into map

	tmpMap := viper.GetStringMap("TARGETPROP")
	if len(tmpMap) == 0 {
		panic("check config file for instrIDs")
	}
	for k, m := range tmpMap {
		if k == "stock" {
			for ks, vs := range m.(map[string]interface{}) {
				fmt.Println(ks, vs)
				fmt.Println(vs.(map[string]interface{})["contractsize"])
			}
		}
		if k == "futures" {
			for kf, vf := range m.(map[string]interface{}) {
				fmt.Println(kf, vf)
				fmt.Println(vf.(map[string]interface{})["contractsize"])
			}

		}
	}

}
