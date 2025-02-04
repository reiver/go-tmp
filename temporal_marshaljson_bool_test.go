package tmp_test

import (
	"testing"

	"time"

	"github.com/reiver/go-tmp"
)

func TestTemporal_MarshalJSON_bool(t *testing.T) {

	tests := []struct{
		Value tmp.Temporal[bool]
		Expected string
	}{
		{
			Value: tmp.Permanent(false),
			Expected: "false",
		},
		{
			Value: tmp.Permanent(true),
			Expected: "true",
		},



		{
			Value: tmp.Temporary(false, time.Unix(9_223_372_036_854_775_807,0)), // supposed to be a time far in the future
			Expected: "false",
		},
		{
			Value: tmp.Temporary(true, time.Unix(9_223_372_036_854_775_806,0)), // supposed to be a time far in the future
			Expected: "true",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for the JSON marshaling is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}

func TestTemporal_MarshalJSON_bool_fail(t *testing.T) {

	tests := []struct{
		Value tmp.Temporal[bool]
	}{
		{
			Value: tmp.Nothing[bool](),
		},



		{
			Value: tmp.Temporary(false, time.Unix(1234,0)), // supposed to be a time that is already expired
		},
		{
			Value: tmp.Temporary(true, time.Unix(5678,0)), // supposed to be a time that is already expired
		},
	}


	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ACTUAL: %q", actualBytes)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}
		if nil != actualBytes {
			t.Errorf("For test #%d, expected not bytes but actually get some.", testNumber)
			t.Logf("ACTUAL: %q", actualBytes)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}
	}
}

