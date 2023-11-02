package tmp_test

import (
	"testing"

	"time"

	"sourcecode.social/reiver/go-tmp"
)

func TestTemporal_IsDefunct(t *testing.T) {

	tests := []struct{
		Temporal tmp.Temporal[string]
		Expected bool
	}{
		{
			Temporal: tmp.Nothing[string](),
			Expected: true,
		},
		{
			Temporal: tmp.Permanent[string]("forever"),
			Expected: false,
		},
		{
			Temporal: tmp.Temporary[string]("expired", time.Now().Unix()-999), // supposed to be in the pastasx
			Expected: true,
		},
		{
			Temporal: tmp.Temporary[string]("not-expired", time.Now().Unix()+99999),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Temporal.IsDefunct()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for is-defunct is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("TEMPORAL: %#v", test.Temporal)
			continue
		}
	}
}
