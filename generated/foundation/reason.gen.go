// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [reason] class.
var (
	ReasonClass     _reasonClass
	ReasonClassOnce sync.Once
)

func getreasonClass() _reasonClass {
	ReasonClassOnce.Do(func() {
		ReasonClass = _reasonClass{objc.GetClass("reason")}
	})
	return ReasonClass
}

type _reasonClass struct {
	class objc.Class
}

// An interface definition for the [reason] class.
type Ireason interface {
	objectivec.IObject
}

//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/reason-c.ivar
type reason struct {
	objectivec.Object
}

// reasonFrom constructs a [reason] from an unsafe.Pointer.
func reasonFrom(ptr unsafe.Pointer) reason {
	return reason{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _reasonClass) Alloc() reason {
	rv := objc.Send[reason](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _reasonClass) New() reason {
	rv := objc.Send[reason](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ reason) Init() reason {
	rv := objc.Send[reason](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ reason) Autorelease() reason {
	rv := objc.Send[reason](r_.ID, objc.Sel("autorelease"))
	return rv
}

// Newreason creates a new reason instance.
func Newreason() reason {
	return getreasonClass().New()
}




