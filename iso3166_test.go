package iso3166

import "testing"

func TestDecode(t *testing.T) {

	tests := []struct {
		Country     string
		Region      string
		Expected    string
		ExpectError bool
	}{
		{"GB", "", "United Kingdom", false},
		{"GB", "CHS", "Cheshire", false},
		{"GB", "NTT", "Nottinghamshire", false},
		{"GB", "ESS", "Essex", false},
		{"GB", "ESX", "East Sussex", false},
		{"GB", "SEX", "", true},
		{"", "", "", true},
		{"", "CHS", "", true},
	}

	for _, test := range tests {
		if actual, err := Decode(test.Country, test.Region); actual != test.Expected || test.ExpectError != (err != nil) {
			t.Errorf("iso3166.Decode(%q, %q) gave %q but wanted %q: %s", test.Country, test.Region, actual, test.Expected, err)
		}
	}
}
