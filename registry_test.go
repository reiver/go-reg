package reg_test

import (
	"testing"

	"github.com/reiver/go-reg"
)

func TestRegistry_int(t *testing.T) {

	tests := []struct{
		Sets map[string]int
		GetsSuccess []string
		GetsFail []string
	}{
		{
			Sets: map[string]int{
				"ONCE":   1,
				"TWICE":  2,
				"THRICE": 3,
				"FOURCE": 4,
			},
			GetsSuccess: []string{
				"TWICE",
				"FOURCE",
			},
			GetsFail: []string{
				"apple",
				"banana",
				"cherry",
			},
		},
	}

	for testNumber, test := range tests {

		var registry reg.Registry[int]

		for  name, value := range test.Sets {
			registry.Set(name, value)
		}

		for _, name := range test.GetsSuccess {
			actual, found := registry.Get(name)
			if !found {
				t.Errorf("For test #%d, and name %q, expected to get but actually didn't.", testNumber, name)
				t.Logf("NAME: %q", name)
				continue
			}

			{
				expected, found := test.Sets[name]
				if !found {
					t.Errorf("For test #%d, and name %q, did not find the expected value.", testNumber, name)
					t.Logf("NAME: %q", name)
					t.Logf("FOUND: %t", found)
					continue
				}

				if expected != actual {
					t.Errorf("For test #%d, and name %q, the actual value is not what was expeceted.", testNumber, name)
					t.Logf("NAME: %q", name)
					continue
				}
			}
		}

		for _, name := range test.GetsFail {
			_, found := registry.Get(name)
			if found {
				t.Errorf("For test #%d, and name %q, expected to not get but actually did", testNumber, name)
				t.Logf("NAME: %q", name)
				continue
			}
		}
	}
}
