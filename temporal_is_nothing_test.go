package tmp

import (
	"testing"

	"time"
)

func TestTemporal_IsNothing(t *testing.T) {

	tests := []struct{
		Temporal Temporal[string]
		Expected bool
	}{
		{
			Temporal: Nothing[string](),
			Expected: true,
		},
		{
			Temporal: Permanent[string]("forever"),
			Expected: false,
		},
		{
			Temporal: Temporary[string]("expired", time.Unix(time.Now().Unix()-999,0)),
			Expected: false,
		},
		{
			Temporal: Temporary[string]("not-expired", time.Unix(time.Now().Unix()+99999,0)),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Temporal.isnothing()
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
