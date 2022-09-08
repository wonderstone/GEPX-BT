package exporter

import (
	"testing"

	"github.com/wonderstone/QuantTools/account/virtualaccount"
)

// test ExportRealtimeYaml function
func TestExportRealtimeYaml(t *testing.T) {
	va := virtualaccount.NewVirtualAccount("2017/10/9 9:39", 100000, 100000)
	va.SAcct.ResetMVSlice()
	tmp := make(map[string]interface{})
	tmp["SIndiNmsAfter"] = []string{"S1"}
	tmp["FIndiNmsAfter"] = []string{"F1"}
	tmp["SDTfields"] = []string{"S1"}
	tmp["FDTfields"] = []string{"F1"}
	ExportRealtimeYaml("../config/Manual", "Default", va, tmp)
}

func TestExportSKE(t *testing.T) {
	ExportSKE("../config/Manual", "GEP", []string{"S1"})
}
