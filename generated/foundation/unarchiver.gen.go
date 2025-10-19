// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Unarchiver] class.
var (
	unarchiverClass     _UnarchiverClass
	unarchiverClassOnce sync.Once
)

func getUnarchiverClass() _UnarchiverClass {
	unarchiverClassOnce.Do(func() {
		unarchiverClass = _UnarchiverClass{objc.GetClass("NSUnarchiver")}
	})
	return unarchiverClass
}

type _UnarchiverClass struct {
	class objc.Class
}

// An interface definition for the [Unarchiver] class.
type IUnarchiver interface {
	ICoder
}

// A decoder that restores data from an archive.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver
type Unarchiver struct {
	Coder
}

// UnarchiverFrom constructs a [Unarchiver] from an unsafe.Pointer.
//
// A decoder that restores data from an archive.
func UnarchiverFrom(ptr unsafe.Pointer) Unarchiver {
	return Unarchiver{
		Coder: CoderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnarchiverClass) Alloc() Unarchiver {
	rv := objc.Send[Unarchiver](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnarchiverClass) New() Unarchiver {
	rv := objc.Send[Unarchiver](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ Unarchiver) Init() Unarchiver {
	rv := objc.Send[Unarchiver](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ Unarchiver) Autorelease() Unarchiver {
	rv := objc.Send[Unarchiver](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnarchiver creates a new Unarchiver instance.
func NewUnarchiver() Unarchiver {
	return getUnarchiverClass().New()
}




