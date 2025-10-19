// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UUID] class.
var (
	uUIDClass     _UUIDClass
	uUIDClassOnce sync.Once
)

func getUUIDClass() _UUIDClass {
	uUIDClassOnce.Do(func() {
		uUIDClass = _UUIDClass{objc.GetClass("NSUUID")}
	})
	return uUIDClass
}

type _UUIDClass struct {
	class objc.Class
}

// An interface definition for the [UUID] class.
type IUUID interface {
	objectivec.IObject
}

// A universally unique value that can be used to identify types, interfaces, and other items.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUUID
type UUID struct {
	objectivec.Object
}

// UUIDFrom constructs a [UUID] from an unsafe.Pointer.
//
// A universally unique value that can be used to identify types, interfaces, and other items.
func UUIDFrom(ptr unsafe.Pointer) UUID {
	return UUID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UUIDClass) Alloc() UUID {
	rv := objc.Send[UUID](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UUIDClass) New() UUID {
	rv := objc.Send[UUID](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UUID) Init() UUID {
	rv := objc.Send[UUID](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UUID) Autorelease() UUID {
	rv := objc.Send[UUID](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUUID creates a new UUID instance.
func NewUUID() UUID {
	return getUUIDClass().New()
}




