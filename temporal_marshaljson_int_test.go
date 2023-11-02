package tmp_test

import (
	"testing"

	"time"

	"sourcecode.social/reiver/go-tmp"
)

func TestTemporal_MarshalJSON_int(t *testing.T) {

	tests := []struct{
		Value tmp.Temporal[int]
	}{
		{
			Value: tmp.Permanent(-65536),
		},
		{
			Value: tmp.Permanent(-65535),
		},



		{
			Value: tmp.Permanent(-256),
		},
		{
			Value: tmp.Permanent(-255),
		},



		{
			Value: tmp.Permanent(-5),
		},
		{
			Value: tmp.Permanent(-4),
		},
		{
			Value: tmp.Permanent(-3),
		},
		{
			Value: tmp.Permanent(-2),
		},
		{
			Value: tmp.Permanent(-1),
		},
		{
			Value: tmp.Permanent(0),
		},
		{
			Value: tmp.Permanent(1),
		},
		{
			Value: tmp.Permanent(2),
		},
		{
			Value: tmp.Permanent(3),
		},
		{
			Value: tmp.Permanent(4),
		},
		{
			Value: tmp.Permanent(5),
		},



		{
			Value: tmp.Permanent(255),
		},
		{
			Value: tmp.Permanent(256),
		},



		{
			Value: tmp.Permanent(65535),
		},
		{
			Value: tmp.Permanent(65536),
		},









		{
			Value: tmp.Temporary(-65536, time.Unix(9_223_372_036_854_775_807,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(-65535, time.Unix(9_223_372_036_854_775_806,0)), // supposed to be a time far in the future
		},



		{
			Value: tmp.Temporary(-256, time.Unix(9_223_372_036_854_775_805,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(-255, time.Unix(9_223_372_036_854_775_804,0)), // supposed to be a time far in the future
		},



		{
			Value: tmp.Temporary(-5, time.Unix(9_223_372_036_854_775_803,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(-4, time.Unix(9_223_372_036_854_775_802,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(-3, time.Unix(9_223_372_036_854_775_802,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(-2, time.Unix(9_223_372_036_854_775_801,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(-1, time.Unix(9_223_372_036_854_775_800,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(0, time.Unix(9_223_372_036_854_775_799,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(1, time.Unix(9_223_372_036_854_775_798,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(2, time.Unix(9_223_372_036_854_775_797,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(3, time.Unix(9_223_372_036_854_775_796,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(4, time.Unix(9_223_372_036_854_775_795,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(5, time.Unix(9_223_372_036_854_775_794,0)), // supposed to be a time far in the future
		},



		{
			Value: tmp.Temporary(255, time.Unix(9_223_372_036_854_775_793,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(256, time.Unix(9_223_372_036_854_775_792,0)), // supposed to be a time far in the future
		},



		{
			Value: tmp.Temporary(65535, time.Unix(9_223_372_036_854_775_791,0)), // supposed to be a time far in the future
		},
		{
			Value: tmp.Temporary(65536, time.Unix(9_223_372_036_854_775_790,0)), // supposed to be a time far in the future
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one." , testNumber)
			t.Logf("ACTUAL BYTES: %q", actualBytes)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}

func TestTemporal_MarshalJSON_int_fail(t *testing.T) {

	var nothing tmp.Temporal[int]

	actualBytes, err := nothing.MarshalJSON()
	if nil == err {
		t.Error("Expected an error but did not actually get one.")
		t.Logf("ACTUAL: %q", actualBytes)
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
	if nil != actualBytes {
		t.Error("Expected not bytes but actually get some.")
		t.Logf("ACTUAL: %q", actualBytes)
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
}
