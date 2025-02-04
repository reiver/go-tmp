package tmp_test

import (
	"testing"

	"time"

	"github.com/reiver/go-tmp"
)

func TestNullable_MarshalJSON_string(t *testing.T) {

	tests := []struct{
		Value tmp.Temporal[string]
		Expected string
	}{
		{
			Value: tmp.Permanent(""),
			Expected: `""`,
		},



		{
			Value: tmp.Temporary("", time.Unix(9_223_372_036_854_775_807,0)), // this is supposed to be some time far in the future
			Expected: `""`,
		},



		{
			Value: tmp.Permanent("apple"),
			Expected: `"apple"`,
		},
		{
			Value: tmp.Permanent("banana"),
			Expected: `"banana"`,
		},
		{
			Value: tmp.Permanent("cherry"),
			Expected: `"cherry"`,
		},



		{
			Value: tmp.Temporary("apple", time.Unix(9_223_372_036_854_775_806,0)), // this is supposed to be some time far in the future
			Expected: `"apple"`,
		},
		{
			Value: tmp.Temporary("banana", time.Unix(9_223_372_036_854_775_805,0)), // this is supposed to be some time far in the future
			Expected: `"banana"`,
		},
		{
			Value: tmp.Temporary("cherry", time.Unix(9_223_372_036_854_775_804,0)), // this is supposed to be some time far in the future
			Expected: `"cherry"`,
		},



		{
			Value: tmp.Permanent("ONCE"),
			Expected: `"ONCE"`,
		},
		{
			Value: tmp.Permanent("TWICE"),
			Expected: `"TWICE"`,
		},
		{
			Value: tmp.Permanent("THRICE"),
			Expected: `"THRICE"`,
		},
		{
			Value: tmp.Permanent("FOURCE"),
			Expected: `"FOURCE"`,
		},



		{
			Value: tmp.Temporary("ONCE", time.Unix(9_223_372_036_854_775_803,0)), // this is supposed to be some time far in the future
			Expected: `"ONCE"`,
		},
		{
			Value: tmp.Temporary("TWICE", time.Unix(9_223_372_036_854_775_802,0)), // this is supposed to be some time far in the future
			Expected: `"TWICE"`,
		},
		{
			Value: tmp.Temporary("THRICE", time.Unix(9_223_372_036_854_775_801,0)), // this is supposed to be some time far in the future
			Expected: `"THRICE"`,
		},
		{
			Value: tmp.Temporary("FOURCE", time.Unix(9_223_372_036_854_775_800,0)), // this is supposed to be some time far in the future
			Expected: `"FOURCE"`,
		},



		{
			Value: tmp.Permanent("🙂"),
			Expected: `"🙂"`,
		},
		{
			Value: tmp.Permanent("😈"),
			Expected: `"😈"`,
		},
		{
			Value: tmp.Permanent("❤️"),
			Expected: `"❤️"`,
		},



		{
			Value: tmp.Temporary("🙂", time.Unix(9_223_372_036_854_775_799,0)), // this is supposed to be some time far in the future
			Expected: `"🙂"`,
		},
		{
			Value: tmp.Temporary("😈", time.Unix(9_223_372_036_854_775_798,0)), // this is supposed to be some time far in the future
			Expected: `"😈"`,
		},
		{
			Value: tmp.Temporary("❤️", time.Unix(9_223_372_036_854_775_797,0)), // this is supposed to be some time far in the future
			Expected: `"❤️"`,
		},



		{
			Value: tmp.Permanent("٠١٢٣۴۵۶٧٨٩"),
			Expected: `"٠١٢٣۴۵۶٧٨٩"`,
		},



		{
			Value: tmp.Temporary("٠١٢٣۴۵۶٧٨٩", time.Unix(9_223_372_036_854_775_796,0)), // this is supposed to be some time far in the future
			Expected: `"٠١٢٣۴۵۶٧٨٩"`,
		},



		{
			Value: tmp.Permanent("𐏑𐏓𐏕"),
			Expected: `"𐏑𐏓𐏕"`,
		},



		{
			Value: tmp.Temporary("𐏑𐏓𐏕", time.Unix(9_223_372_036_854_775_795,0)), // this is supposed to be some time far in the future
			Expected: `"𐏑𐏓𐏕"`,
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

func TestNullable_MarshalJSON_string_fail(t *testing.T) {

	tests := []struct{
		Value tmp.Temporal[string]
	}{
		{
			Value: tmp.Nothing[string](),
		},
		{
			Value: tmp.Temporary("expired", time.Unix(11,0)),
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
