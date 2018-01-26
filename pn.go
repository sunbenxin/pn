// Package pn provides some func on cell-phone number
package pn

var mobilePre = []string{
	"134",
	"135",
	"136",
	"137",
	"138",
	"139",
	"147",
	"150",
	"151",
	"152",
	"157",
	"158",
	"159",
	"178",
	"182",
	"183",
	"184",
	"187",
	"188",
}

var unionPre = []string{
	"130",
	"131",
	"132",
	"145",
	"155",
	"156",
	"171",
	"175",
	"176",
	"185",
	"186",
}
var carrieOperatorDict = map[string]CarrieOperatorType{}

func init() {
	for _, pre := range unionPre {
		carrieOperatorDict[pre] = CarrieOperatorUnicom
	}
	for _, pre := range mobilePre {
		carrieOperatorDict[pre] = CarrieOperatorMobile
	}
}

type CarrieOperatorType int // carrier operator

const (
	CarrieOperatorTelcom CarrieOperatorType = iota
	CarrieOperatorMobile
	CarrieOperatorUnicom
	CarrieOperatorVirtual
)

func GetCarrieOperator(mb string) CarrieOperatorType {
	if len(mb) < 8 {
		return CarrieOperatorTelcom
	}
	mobilePre := mb[0:3]
	if t, ok := carrieOperatorDict[mobilePre]; ok {
		return t
	}
	return CarrieOperatorTelcom
}
