package tmp

import (
	"fmt"
	"time"

	"github.com/reiver/go-opt"
)

type Temporal[T any] struct {
	value T
	until int64 // unix-time
	istemporary bool
	ispermanent bool
}

func Nothing[T any]() Temporal[T] {
	var nothing Temporal[T]

	return nothing
}

func Permanent[T any](value T) Temporal[T] {
	return Temporal[T]{
		ispermanent:true,
		value:value,
	}
}

func Temporary[T any](value T, until time.Time) Temporal[T] {
	return Temporal[T]{
		istemporary:true,
		value:value,
		until:until.Unix(),
	}
}

// Optional returns a optional-type (also called an option-type and maybe-type) based on the temporal-type.
// If the temporal-type or defunct, then it returns 'nothing'.
// Else it returns 'something'.
func (receiver Temporal[T]) Optional() opt.Optional[T] {
	switch {
	case receiver.isnothing():
		return opt.Nothing[T]()
	case receiver.IsDefunct():
		return opt.Nothing[T]()
	default:
		return opt.Something[T](receiver.value)
	}
}

func (receiver Temporal[T]) isnothing() bool {
	return !receiver.istemporary && !receiver.ispermanent
}

func (receiver Temporal[T]) Filter(fn func(T)bool) Temporal[T] {
	if receiver.isnothing() {
		return Nothing[T]()
	}
	if receiver.IsDefunct() {
		return Nothing[T]()
	}

	if !fn(receiver.value) {
		return Nothing[T]()
	}

	return receiver
}

func (receiver Temporal[T]) Get() (T, bool) {
	return receiver.value, receiver.IsExtant()
}

func (receiver Temporal[T]) GoString() string {
	if receiver.isnothing() {
		return fmt.Sprintf("tmp.Nothing[%T]()", receiver.value)
	}
	if receiver.ispermanent {
		return fmt.Sprintf("tmp.Permanent[%T](%#v)", receiver.value, receiver.value)
	}

	return fmt.Sprintf("tmp.Temporary[%T](%#v, time.Unix(%d, 0))", receiver.value, receiver.value, receiver.until)
}

func (receiver Temporal[T]) IsDefunct() bool {
	if receiver.ispermanent {
		return false
	}

	if receiver.istemporary {
		var now int64 = time.Now().Unix()

		return receiver.until < now
	}

	return true
}

func (receiver Temporal[T]) IsDefunctWhen(when time.Time) bool {
	if receiver.ispermanent {
		return false
	}

	if receiver.istemporary {
		var then int64 = when.Unix()

		return receiver.until < then
	}

	return true
}

func (receiver Temporal[T]) IsExtant() bool {
	if receiver.ispermanent {
		return true
	}

	if receiver.istemporary {
		var now int64 = time.Now().Unix()

		return now <= receiver.until
	}

	return false
}

func (receiver Temporal[T]) WhenNothing(fn func()) {
	if receiver.isnothing() {
		fn()
	}
}

func (receiver Temporal[T]) WhenDefunct(fn func()) {
	if receiver.IsDefunct() {
		fn()
	}
}

func (receiver Temporal[T]) WhenExtant(fn func(T)) {
	if receiver.IsExtant() {
		fn(receiver.value)
	}
}
