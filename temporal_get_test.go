package tmp_test

import (
	"testing"

	"time"

	"github.com/reiver/go-tmp"
)

func TestOptional_Get_string(t *testing.T) {

	tests := []struct{
		Temporal tmp.Temporal[string]
		ExpectedValue     string
		ExpectedSomething bool
	}{
		{
			Temporal: tmp.Nothing[string](),
			ExpectedValue: "",
			ExpectedSomething: false,
		},



		{
			Temporal: tmp.Permanent(""),
			ExpectedValue:          "",
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent("once twice thrice fource"),
			ExpectedValue:          "once twice thrice fource",
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent("😈"),
			ExpectedValue:          "😈",
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent("۰۱۲۳۴۵۶۷۸۹"),
			ExpectedValue:          "۰۱۲۳۴۵۶۷۸۹",
			ExpectedSomething: true,
		},



		{
			Temporal: tmp.Temporary("", time.Unix(9_223_372_036_854_775_807,0)), // this is supposed to be some time far in the future.
			ExpectedValue:          "",
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Temporary("once twice thrice fource", time.Unix(9_223_372_036_854_775_807,0)), // this is supposed to be some time far in the future.
			ExpectedValue:          "once twice thrice fource",
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Temporary("😈", time.Unix(9_223_372_036_854_775_807,0)), // this is supposed to be some time far in the future.
			ExpectedValue:          "😈",
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Temporary("۰۱۲۳۴۵۶۷۸۹", time.Unix(9_223_372_036_854_775_807,0)), // this is supposed to be some time far in the future.
			ExpectedValue:          "۰۱۲۳۴۵۶۷۸۹",
			ExpectedSomething: true,
		},



		{
			Temporal: tmp.Temporary("", time.Unix(10,0)), // this is supposed to be some time in the past.
			ExpectedValue:          "",
			ExpectedSomething: false,
		},
		{
			Temporal: tmp.Temporary("once twice thrice fource", time.Unix(11,0)), // this is supposed to be some time in the past.
			ExpectedValue:          "once twice thrice fource",
			ExpectedSomething: false,
		},
		{
			Temporal: tmp.Temporary("😈", time.Unix(12,0)), // this is supposed to be some time in the past.
			ExpectedValue:          "😈",
			ExpectedSomething: false,
		},
		{
			Temporal: tmp.Temporary("۰۱۲۳۴۵۶۷۸۹", time.Unix(13,0)), // this is supposed to be some time in the past.
			ExpectedValue:          "۰۱۲۳۴۵۶۷۸۹",
			ExpectedSomething: false,
		},
	}

	for testNumber, test := range tests {

		value, something := test.Temporal.Get()

		{
			expected := test.ExpectedSomething
			actual   := something

			if expected != actual {
				t.Errorf("For test #%d, the actual something-flag is not what was expected.", testNumber)
				t.Logf("EXPECTED FLAG: %t", expected)
				t.Logf("ACTUAL   FLAG: %t", actual)
				t.Logf("NULLABLE: %#v", test.Temporal)
	/////////////////////// CONTINUE
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := value

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("NULLABLE: %#v", test.Temporal)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_Get_int8(t *testing.T) {

	tests := []struct{
		Temporal tmp.Temporal[int8]
		ExpectedValue     int8
		ExpectedSomething bool
	}{
		{
			Temporal: tmp.Nothing[int8](),
			ExpectedValue: 0,
			ExpectedSomething: false,
		},



		{
			Temporal: tmp.Permanent(int8(-127)),
			ExpectedValue:          int8(-127),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-126)),
			ExpectedValue:          int8(-126),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-125)),
			ExpectedValue:          int8(-125),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-124)),
			ExpectedValue:          int8(-124),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-123)),
			ExpectedValue:          int8(-123),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-122)),
			ExpectedValue:          int8(-122),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-121)),
			ExpectedValue:          int8(-121),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-120)),
			ExpectedValue:          int8(-120),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-9)),
			ExpectedValue:          int8(-9),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-8)),
			ExpectedValue:          int8(-8),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-7)),
			ExpectedValue:          int8(-7),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-6)),
			ExpectedValue:          int8(-6),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-5)),
			ExpectedValue:          int8(-5),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-4)),
			ExpectedValue:          int8(-4),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-3)),
			ExpectedValue:          int8(-3),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-2)),
			ExpectedValue:          int8(-2),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(-1)),
			ExpectedValue:          int8(-1),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(0)),
			ExpectedValue:          int8(0),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(1)),
			ExpectedValue:          int8(1),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(2)),
			ExpectedValue:          int8(2),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(3)),
			ExpectedValue:          int8(3),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(4)),
			ExpectedValue:          int8(4),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(5)),
			ExpectedValue:          int8(5),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(6)),
			ExpectedValue:          int8(6),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(7)),
			ExpectedValue:          int8(7),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(8)),
			ExpectedValue:          int8(8),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(9)),
			ExpectedValue:          int8(9),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(120)),
			ExpectedValue:          int8(120),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(121)),
			ExpectedValue:          int8(121),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(122)),
			ExpectedValue:          int8(122),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(123)),
			ExpectedValue:          int8(123),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(124)),
			ExpectedValue:          int8(124),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(125)),
			ExpectedValue:          int8(125),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(126)),
			ExpectedValue:          int8(126),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(int8(127)),
			ExpectedValue:          int8(127),
			ExpectedSomething: true,
		},
	}

	for testNumber, test := range tests {

		value, something := test.Temporal.Get()

		{
			expected := test.ExpectedSomething
			actual   := something

			if expected != actual {
				t.Errorf("For test #%d, the actual something-flag is not what was expected.", testNumber)
				t.Logf("EXPECTED FLAG: %t", expected)
				t.Logf("ACTUAL   FLAG: %t", actual)
				t.Logf("NULLABLE: %#v", test.Temporal)
	/////////////////////// CONTINUE
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := value

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("NULLABLE: %#v", test.Temporal)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_Get_uint8(t *testing.T) {

	tests := []struct{
		Temporal tmp.Temporal[uint8]
		ExpectedValue     uint8
		ExpectedSomething bool
	}{
		{
			Temporal: tmp.Nothing[uint8](),
			ExpectedValue: 0,
			ExpectedSomething: false,
		},



		{
			Temporal: tmp.Permanent(uint8(0)),
			ExpectedValue:          uint8(0),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(uint8(1)),
			ExpectedValue:          uint8(1),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(uint8(2)),
			ExpectedValue:          uint8(2),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(uint8(254)),
			ExpectedValue:          uint8(254),
			ExpectedSomething: true,
		},
		{
			Temporal: tmp.Permanent(uint8(255)),
			ExpectedValue:          uint8(255),
			ExpectedSomething: true,
		},
	}

	for testNumber, test := range tests {

		value, something := test.Temporal.Get()

		{
			expected := test.ExpectedSomething
			actual   := something

			if expected != actual {
				t.Errorf("For test #%d, the actual something-flag is not what was expected.", testNumber)
				t.Logf("EXPECTED FLAG: %t", expected)
				t.Logf("ACTUAL   FLAG: %t", actual)
				t.Logf("NULLABLE: %#v", test.Temporal)
	/////////////////////// CONTINUE
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := value

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("NULLABLE: %#v", test.Temporal)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}
