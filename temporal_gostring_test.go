package tmp_test

import (
	"testing"

	"fmt"
	"time"

	"sourcecode.social/reiver/go-tmp"
)

func TestTemporal_GoString_permanent(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value:                           "",
			Expected: `tmp.Permanent[string]("")`,
		},
		{
			Value:                           "once twice thrice fource",
			Expected: `tmp.Permanent[string]("once twice thrice fource")`,
		},
		{
			Value:                           "apple banana cherry",
			Expected: `tmp.Permanent[string]("apple banana cherry")`,
		},



		{
			Value:                   uint8 (0x0),
			Expected: `tmp.Permanent[uint8](0x0)`,
		},
		{
			Value:                   uint8 (0x1),
			Expected: `tmp.Permanent[uint8](0x1)`,
		},
		{
			Value:                   uint8 (0x2),
			Expected: `tmp.Permanent[uint8](0x2)`,
		},
		{
			Value:                   uint8 (0xfe),
			Expected: `tmp.Permanent[uint8](0xfe)`,
		},
		{
			Value:                   uint8 (0xff),
			Expected: `tmp.Permanent[uint8](0xff)`,
		},



		{
			Value:                   uint16 (0x0),
			Expected: `tmp.Permanent[uint16](0x0)`,
		},
		{
			Value:                   uint16 (0x1),
			Expected: `tmp.Permanent[uint16](0x1)`,
		},
		{
			Value:                   uint16 (0x2),
			Expected: `tmp.Permanent[uint16](0x2)`,
		},
		{
			Value:                   uint16 (0xfe),
			Expected: `tmp.Permanent[uint16](0xfe)`,
		},
		{
			Value:                   uint16 (0xff),
			Expected: `tmp.Permanent[uint16](0xff)`,
		},
		{
			Value:                   uint16 (0x100),
			Expected: `tmp.Permanent[uint16](0x100)`,
		},
		{
			Value:                   uint16 (0x101),
			Expected: `tmp.Permanent[uint16](0x101)`,
		},
		{
			Value:                   uint16 (0x102),
			Expected: `tmp.Permanent[uint16](0x102)`,
		},
		{
			Value:                   uint16 (0xfffe),
			Expected: `tmp.Permanent[uint16](0xfffe)`,
		},
		{
			Value:                   uint16 (0xffff),
			Expected: `tmp.Permanent[uint16](0xffff)`,
		},



		{
			Value:                   struct { A string; B int }{A:"joeblow",B:7},
			Expected: `tmp.Permanent[struct { A string; B int }](struct { A string; B int }{A:"joeblow", B:7})`,
		},
	}

	for testNumber, test := range tests {

		op := tmp.Permanent(test.Value)
		gostring := op.GoString()

		{
			expected := test.Expected
			actual   := gostring

			if expected != actual {
				t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE-TYPE: %T", test.Value)
				t.Logf("VALUE: %#v", test.Value)
	//////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestTemporal_GoString_nothing(t *testing.T) {

	tests := []struct{
		Temporal fmt.GoStringer
		Expected string
	}{
		{
			Temporal:  tmp.Nothing[string](),
			Expected: `tmp.Nothing[string]()`,
		},



		{
			Temporal:  tmp.Nothing[int8](),
			Expected: `tmp.Nothing[int8]()`,
		},
		{
			Temporal:  tmp.Nothing[int16](),
			Expected: `tmp.Nothing[int16]()`,
		},
		{
			Temporal:  tmp.Nothing[int32](),
			Expected: `tmp.Nothing[int32]()`,
		},
		{
			Temporal:  tmp.Nothing[int64](),
			Expected: `tmp.Nothing[int64]()`,
		},



		{
			Temporal:  tmp.Nothing[uint8](),
			Expected: `tmp.Nothing[uint8]()`,
		},
		{
			Temporal:  tmp.Nothing[uint16](),
			Expected: `tmp.Nothing[uint16]()`,
		},
		{
			Temporal:  tmp.Nothing[uint32](),
			Expected: `tmp.Nothing[uint32]()`,
		},
		{
			Temporal:  tmp.Nothing[uint64](),
			Expected: `tmp.Nothing[uint64]()`,
		},
	}

	for testNumber, test := range tests {

		expected := test.Expected
		actual   := test.Temporal.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual go-string value is not what was expected.", testNumber)
			t.Logf("EXPECTED GO-STRING: %q", expected)
			t.Logf("ACTUAL   GO-STRING: %q", actual)
			t.Logf("TYPE: %T", test.Temporal)
	/////////////// CONTINUE
			continue
		}
	}
}

func TestTemporal_GoString_temporary(t *testing.T) {

	tests := []struct{
		Value any
		Until    time.Time
		Expected string
	}{
		{
			Value:                           "",
			Until:                               time.Unix(0, 0),
			Expected: `tmp.Temporary[string]("", time.Unix(0, 0))`,
		},
		{
			Value:                           "once twice thrice fource",
			Until:                                                       time.Unix(1234, 0),
			Expected: `tmp.Temporary[string]("once twice thrice fource", time.Unix(1234, 0))`,
		},
		{
			Value:                           "apple banana cherry",
			Until:                                                  time.Unix(979899, 0),
			Expected: `tmp.Temporary[string]("apple banana cherry", time.Unix(979899, 0))`,
		},



		{
			Value:                   uint8 (0x0),
			Until:                               time.Unix(101, 0),
			Expected: `tmp.Temporary[uint8](0x0, time.Unix(101, 0))`,
		},
		{
			Value:                   uint8 (0x1),
			Until:                               time.Unix(111, 0),
			Expected: `tmp.Temporary[uint8](0x1, time.Unix(111, 0))`,
		},
		{
			Value:                   uint8 (0x2),
			Until:                               time.Unix(121, 0),
			Expected: `tmp.Temporary[uint8](0x2, time.Unix(121, 0))`,
		},
		{
			Value:                   uint8 (0xfe),
			Until:                                time.Unix(989, 0),
			Expected: `tmp.Temporary[uint8](0xfe, time.Unix(989, 0))`,
		},
		{
			Value:                   uint8 (0xff),
			Until:                                time.Unix(999, 0),
			Expected: `tmp.Temporary[uint8](0xff, time.Unix(999, 0))`,
		},



		{
			Value:                   uint16 (0x0),
			Until:                                time.Unix(303, 0),
			Expected: `tmp.Temporary[uint16](0x0, time.Unix(303, 0))`,
		},
		{
			Value:                   uint16 (0x1),
			Until:                                time.Unix(313, 0),
			Expected: `tmp.Temporary[uint16](0x1, time.Unix(313, 0))`,
		},
		{
			Value:                   uint16 (0x2),
			Until:                                time.Unix(323, 0),
			Expected: `tmp.Temporary[uint16](0x2, time.Unix(323, 0))`,
		},
		{
			Value:                   uint16 (0xfe),
			Until:                                 time.Unix(383, 0),
			Expected: `tmp.Temporary[uint16](0xfe, time.Unix(383, 0))`,
		},
		{
			Value:                   uint16 (0xff),
			Until:                                 time.Unix(393, 0),
			Expected: `tmp.Temporary[uint16](0xff, time.Unix(393, 0))`,
		},
		{
			Value:                   uint16 (0x100),
			Until:                                  time.Unix(3003, 0),
			Expected: `tmp.Temporary[uint16](0x100, time.Unix(3003, 0))`,
		},
		{
			Value:                   uint16 (0x101),
			Until:                                  time.Unix(3113, 0),
			Expected: `tmp.Temporary[uint16](0x101, time.Unix(3113, 0))`,
		},
		{
			Value:                   uint16 (0x102),
			Until:                                  time.Unix(3223, 0),
			Expected: `tmp.Temporary[uint16](0x102, time.Unix(3223, 0))`,
		},
		{
			Value:                   uint16 (0xfffe),
			Until:                                   time.Unix(3883, 0),
			Expected: `tmp.Temporary[uint16](0xfffe, time.Unix(3883, 0))`,
		},
		{
			Value:                   uint16 (0xffff),
			Until:                                   time.Unix(3993, 0),
			Expected: `tmp.Temporary[uint16](0xffff, time.Unix(3993, 0))`,
		},



		{
			Value:                   struct { A string; B int }{A:"joeblow",B:7},
			Until:                                                                                             time.Unix(9876543210, 0),
			Expected: `tmp.Temporary[struct { A string; B int }](struct { A string; B int }{A:"joeblow", B:7}, time.Unix(9876543210, 0))`,
		},
	}

	for testNumber, test := range tests {

		op := tmp.Temporary(test.Value, test.Until)
		gostring := op.GoString()

		{
			expected := test.Expected
			actual   := gostring

			if expected != actual {
				t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE-TYPE: %T", test.Value)
				t.Logf("VALUE: %#v", test.Value)
	//////////////////////// CONTINUE
				continue
			}
		}
	}
}
