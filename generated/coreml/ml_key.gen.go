// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Key] class.
var (
	KeyClass     _KeyClass
	KeyClassOnce sync.Once
)

func getKeyClass() _KeyClass {
	KeyClassOnce.Do(func() {
		KeyClass = _KeyClass{objc.GetClass("MLKey")}
	})
	return KeyClass
}

type _KeyClass struct {
	class objc.Class
}

// An interface definition for the [Key] class.
type IKey interface {
	objectivec.IObject
}

// An abstract base class for machine learning key types.
//
// You don’t create use this class directly. Instead, use a class that inherits from this one, such as or .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLKey
type Key struct {
	objectivec.Object
}

// KeyFrom constructs a [Key] from an unsafe.Pointer.
//
// An abstract base class for machine learning key types.
func KeyFrom(ptr unsafe.Pointer) Key {
	return Key{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (kc _KeyClass) Alloc() Key {
	rv := objc.Send[Key](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KeyClass) New() Key {
	rv := objc.Send[Key](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ Key) Init() Key {
	rv := objc.Send[Key](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ Key) Autorelease() Key {
	rv := objc.Send[Key](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKey creates a new Key instance.
func NewKey() Key {
	return getKeyClass().New()
}




