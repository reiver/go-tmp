package tmp_test

import (
	"testing"

	"sourcecode.social/reiver/go-tmp"
)

func TestTemporal_Filter_int(t *testing.T) {

	tests := []struct{
		Temporal tmp.Temporal[int]
		Expected tmp.Temporal[int]
	}{
		{
			Temporal: tmp.Nothing[int](),
			Expected: tmp.Nothing[int](),
		},



		{
			Temporal: tmp.Permanent[int](-2),
			Expected: tmp.Permanent[int](-2),
		},
		{
			Temporal: tmp.Permanent[int](-1),
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Permanent[int](0),
			Expected: tmp.Permanent[int](0),
		},
		{
			Temporal: tmp.Permanent[int](1),
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Permanent[int](2),
			Expected: tmp.Permanent[int](2),
		},



		{
			Temporal: tmp.Temporary[int](-2, 9_223_372_036_854_775_807), // int64 problem
			Expected: tmp.Temporary[int](-2, 9_223_372_036_854_775_807),
		},
		{
			Temporal: tmp.Temporary[int](-1, 9_223_372_036_854_775_806), // int64 problem
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Temporary[int](0, 9_223_372_036_854_775_805), // int64 problem
			Expected: tmp.Temporary[int](0, 9_223_372_036_854_775_805),
		},
		{
			Temporal: tmp.Temporary[int](1, 9_223_372_036_854_775_804), // int64 problem
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Temporary[int](2, 9_223_372_036_854_775_803),
			Expected: tmp.Temporary[int](2, 9_223_372_036_854_775_803), // int64 problem
		},



		{
			Temporal: tmp.Temporary[int](-2, 111), // supposed to be in the past
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Temporary[int](-1, 111), // supposed to be in the past
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Temporary[int](0, 111), // supposed to be in the past
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Temporary[int](1, 111), // supposed to be in the past
			Expected: tmp.Nothing[int](),
		},
		{
			Temporal: tmp.Temporary[int](2, 111), // supposed to be in the past
			Expected: tmp.Nothing[int](),
		},
	}

	for testNumber, test := range tests {

		fn := func(value int) bool {
			return 0 == (value % 2)
		}

		actual   := test.Temporal.Filter(fn)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
	/////////////// CONTINUE
			continue
		}
	}
}
