package tmp

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
)

var _ json.Unmarshaler = new(Temporal[bool])
var _ json.Unmarshaler = new(Temporal[string])

// UnmarshalJSON makes it so json.Unmarshaler is implemented.
func (receiver *Temporal[T]) UnmarshalJSON(data []byte) error {
	switch interface{}(receiver.value).(type) {
	case bool, string,json.Unmarshaler:
		if 4 == len(data) && 'n' == data[0] && 'u' == data[1] && 'l' == data[2] && 'l' == data[3] {
			return erorr.Errorf("tmp: cannot unmarshal a JSON `null` into something of type %T", receiver.value)
		}

		// these are OK.
	default:
		return erorr.Errorf("tmp: cannot unmarshal into something of type %T from JSON because parameterized type is ‘%T’ rather than ‘bool’, ‘string’, or ‘json.Unmarshaler’", receiver, receiver.value)
	}

	{
		var dst T

		err := json.Unmarshal(data, &dst)
		if nil != err {
			return err
		}

		*receiver = Permanent(dst)
	}

	return nil
}
