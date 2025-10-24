// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [uuidString] class.
var (
	UuidStringClass     _uuidStringClass
	UuidStringClassOnce sync.Once
)

func getuuidStringClass() _uuidStringClass {
	UuidStringClassOnce.Do(func() {
		UuidStringClass = _uuidStringClass{objc.GetClass("uuidString")}
	})
	return UuidStringClass
}

type _uuidStringClass struct {
	class objc.Class
}

// An interface definition for the [uuidString] class.
type IuuidString interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-c.ivar
type uuidString struct {
	objectivec.Object
}

// uuidStringFrom constructs a [uuidString] from an unsafe.Pointer.
func uuidStringFrom(ptr unsafe.Pointer) uuidString {
	return uuidString{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _uuidStringClass) Alloc() uuidString {
	rv := objc.Send[uuidString](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _uuidStringClass) New() uuidString {
	rv := objc.Send[uuidString](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ uuidString) Init() uuidString {
	rv := objc.Send[uuidString](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ uuidString) Autorelease() uuidString {
	rv := objc.Send[uuidString](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewuuidString creates a new uuidString instance.
func NewuuidString() uuidString {
	return getuuidStringClass().New()
}




