package tmp_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-tmp"
)

func TestTemporal_UnmarshalJSON_string(t *testing.T) {

	tests := []struct{
		JSON string
		Expected tmp.Temporal[string]
	}{
		{
			JSON: `""`,
			Expected: tmp.Permanent(""),
		},



		{
			JSON: `"apple"`,
			Expected: tmp.Permanent("apple"),
		},
		{
			JSON: `"banana"`,
			Expected: tmp.Permanent("banana"),
		},
		{
			JSON: `"cherry"`,
			Expected: tmp.Permanent("cherry"),
		},



		{
			JSON: `"ONCE"`,
			Expected: tmp.Permanent("ONCE"),
		},
		{
			JSON: `"TWICE"`,
			Expected: tmp.Permanent("TWICE"),
		},
		{
			JSON: `"THRICE"`,
			Expected: tmp.Permanent("THRICE"),
		},
		{
			JSON: `"FOURCE"`,
			Expected: tmp.Permanent("FOURCE"),
		},



		{
			JSON: `"🙂"`,
			Expected: tmp.Permanent("🙂"),
		},
		{
			JSON: `"😈"`,
			Expected: tmp.Permanent("😈"),
		},
		{
			JSON: `"❤️"`,
			Expected: tmp.Permanent("❤️"),
		},



		{
			JSON: `"٠١٢٣۴۵۶٧٨٩"`,
			Expected: tmp.Permanent("٠١٢٣۴۵۶٧٨٩"),
		},



		{
			JSON: `"𐏑𐏓𐏕"`,
			Expected: tmp.Permanent("𐏑𐏓𐏕"),
		},
	}

	for testNumber, test := range tests {

		var actual tmp.Temporal[string]

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value of the nullable optional type is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON: %q", test.JSON)
				continue
			}
		}
	}
}

func TestTemporal_UnmarshalJSON_string_fail(t *testing.T) {

	tests := []struct{
		JSON string
	}{
		{
			JSON: `null`,
		},
		{
			JSON: `12345`,
		},
		{
			JSON: `false`,
		},
		{
			JSON: `true`,
		},
	}

	for testNumber, test := range tests {

		var op tmp.Temporal[string]

		err := json.Unmarshal([]byte(test.JSON), &op)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: %q", test.JSON)
			continue
		}
	}
}
