package tmp_test

import (
	"testing"

	"fmt"

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
		Until    int64
		Expected string
	}{
		{
			Value:                           "",
			Until:                               0,
			Expected: `tmp.Temporary[string]("", 0)`,
		},
		{
			Value:                           "once twice thrice fource",
			Until:                                                       1234,
			Expected: `tmp.Temporary[string]("once twice thrice fource", 1234)`,
		},
		{
			Value:                           "apple banana cherry",
			Until:                                                  979899,
			Expected: `tmp.Temporary[string]("apple banana cherry", 979899)`,
		},



		{
			Value:                   uint8 (0x0),
			Until:                               101,
			Expected: `tmp.Temporary[uint8](0x0, 101)`,
		},
		{
			Value:                   uint8 (0x1),
			Until:                               111,
			Expected: `tmp.Temporary[uint8](0x1, 111)`,
		},
		{
			Value:                   uint8 (0x2),
			Until:                               121,
			Expected: `tmp.Temporary[uint8](0x2, 121)`,
		},
		{
			Value:                   uint8 (0xfe),
			Until:                                989,
			Expected: `tmp.Temporary[uint8](0xfe, 989)`,
		},
		{
			Value:                   uint8 (0xff),
			Until:                                999,
			Expected: `tmp.Temporary[uint8](0xff, 999)`,
		},



		{
			Value:                   uint16 (0x0),
			Until:                                303,
			Expected: `tmp.Temporary[uint16](0x0, 303)`,
		},
		{
			Value:                   uint16 (0x1),
			Until:                                313,
			Expected: `tmp.Temporary[uint16](0x1, 313)`,
		},
		{
			Value:                   uint16 (0x2),
			Until:                                323,
			Expected: `tmp.Temporary[uint16](0x2, 323)`,
		},
		{
			Value:                   uint16 (0xfe),
			Until:                                 383,
			Expected: `tmp.Temporary[uint16](0xfe, 383)`,
		},
		{
			Value:                   uint16 (0xff),
			Until:                                 393,
			Expected: `tmp.Temporary[uint16](0xff, 393)`,
		},
		{
			Value:                   uint16 (0x100),
			Until:                                  3003,
			Expected: `tmp.Temporary[uint16](0x100, 3003)`,
		},
		{
			Value:                   uint16 (0x101),
			Until:                                  3113,
			Expected: `tmp.Temporary[uint16](0x101, 3113)`,
		},
		{
			Value:                   uint16 (0x102),
			Until:                                  3223,
			Expected: `tmp.Temporary[uint16](0x102, 3223)`,
		},
		{
			Value:                   uint16 (0xfffe),
			Until:                                   3883,
			Expected: `tmp.Temporary[uint16](0xfffe, 3883)`,
		},
		{
			Value:                   uint16 (0xffff),
			Until:                                   3993,
			Expected: `tmp.Temporary[uint16](0xffff, 3993)`,
		},



		{
			Value:                   struct { A string; B int }{A:"joeblow",B:7},
			Until:                                                                                             9876543210,
			Expected: `tmp.Temporary[struct { A string; B int }](struct { A string; B int }{A:"joeblow", B:7}, 9876543210)`,
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
