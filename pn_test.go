package pn

import "testing"

func TestGetCarrieOperator(t *testing.T) {
	mobTests := []struct {
		mobileNumber string
		res          CarrieOperatorType
	}{
		{"17610968213", CarrieOperatorUnicom},
		{"13210968213", CarrieOperatorUnicom},

		{"13410968213", CarrieOperatorMobile},
		{"13510968213", CarrieOperatorMobile},
		{"13610968213", CarrieOperatorMobile},
		{"13710968213", CarrieOperatorMobile},
		{"13810968213", CarrieOperatorMobile},
		{"13910968213", CarrieOperatorMobile},
		{"14710968213", CarrieOperatorMobile},
		{"15010968213", CarrieOperatorMobile},
		{"15110968213", CarrieOperatorMobile},
		{"15210968213", CarrieOperatorMobile},
		{"15710968213", CarrieOperatorMobile},
		{"15810968213", CarrieOperatorMobile},
		{"15910968213", CarrieOperatorMobile},
		{"17810968213", CarrieOperatorMobile},
		{"18210968213", CarrieOperatorMobile},
		{"18310968213", CarrieOperatorMobile},
		{"18410968213", CarrieOperatorMobile},
		{"18710968213", CarrieOperatorMobile},
		{"18810968213", CarrieOperatorMobile},
	}

	for _, v := range mobTests {
		co := GetCarrieOperator(v.mobileNumber)
		if co != v.res {
			t.Errorf("mobilenumber:%v  get carrieOperatorType %v expected %v", v.mobileNumber, co, v.res)
		}
	}
}
