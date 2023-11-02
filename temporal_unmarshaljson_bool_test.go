package tmp_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-tmp"
)

func TestTemporal_UnmarshalJSON_bool(t *testing.T) {

	tests := []struct{
		JSON string
		Expected tmp.Temporal[bool]
	}{
		{
			JSON: `false`,
			Expected: tmp.Permanent(false),
		},
		{
			JSON: `true`,
			Expected: tmp.Permanent(true),
		},
	}

	for testNumber, test := range tests {

		var actual tmp.Temporal[bool]

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
				t.Errorf("For test #%d, the actual value of the optional type is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON: %q", test.JSON)
				continue
			}
		}
	}
}

func TestTemporal_UnmarshalJSON_bool_fail(t *testing.T) {

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
			JSON: `"something"`,
		},
	}

	for testNumber, test := range tests {

		var op tmp.Temporal[bool]

		err := json.Unmarshal([]byte(test.JSON), &op)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("OP: %#v", op)
			continue
		}
	}
}
