// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathControl] class.
var (
	pathControlClass     _PathControlClass
	pathControlClassOnce sync.Once
)

func getPathControlClass() _PathControlClass {
	pathControlClassOnce.Do(func() {
		pathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}
	})
	return pathControlClass
}

type _PathControlClass struct {
	class objc.Class
}

// An interface definition for the [PathControl] class.
type IPathControl interface {
	IControl
}

// A display of a file system path or virtual path information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl
type PathControl struct {
	Control
}

// PathControlFrom constructs a [PathControl] from an unsafe.Pointer.
//
// A display of a file system path or virtual path information.
func PathControlFrom(ptr unsafe.Pointer) PathControl {
	return PathControl{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PathControlClass) Alloc() PathControl {
	rv := objc.Send[PathControl](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PathControlClass) New() PathControl {
	rv := objc.Send[PathControl](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathControl) Init() PathControl {
	rv := objc.Send[PathControl](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathControl) Autorelease() PathControl {
	rv := objc.Send[PathControl](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathControl creates a new PathControl instance.
func NewPathControl() PathControl {
	return getPathControlClass().New()
}




