// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ParameterKey] class.
var (
	ParameterKeyClass     _ParameterKeyClass
	ParameterKeyClassOnce sync.Once
)

func getParameterKeyClass() _ParameterKeyClass {
	ParameterKeyClassOnce.Do(func() {
		ParameterKeyClass = _ParameterKeyClass{objc.GetClass("MLParameterKey")}
	})
	return ParameterKeyClass
}

type _ParameterKeyClass struct {
	class objc.Class
}

// An interface definition for the [ParameterKey] class.
type IParameterKey interface {
	IKey
}

// The keys for the parameter dictionary in a model configuration or a model update context.
//
// Use an to retrieve a model’s parameter value using: The model’s method The dictionary of an The dictionary of an
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterKey
type ParameterKey struct {
	Key
}

// ParameterKeyFrom constructs a [ParameterKey] from an unsafe.Pointer.
//
// The keys for the parameter dictionary in a model configuration or a model update context.
func ParameterKeyFrom(ptr unsafe.Pointer) ParameterKey {
	return ParameterKey{
		Key: KeyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _ParameterKeyClass) Alloc() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ParameterKeyClass) New() ParameterKey {
	rv := objc.Send[ParameterKey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterKey) Init() ParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterKey) Autorelease() ParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterKey creates a new ParameterKey instance.
func NewParameterKey() ParameterKey {
	return getParameterKeyClass().New()
}




