// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathControl] class.
var (
	PathControlClass     _PathControlClass
	PathControlClassOnce sync.Once
)

func getPathControlClass() _PathControlClass {
	PathControlClassOnce.Do(func() {
		PathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}
	})
	return PathControlClass
}

type _PathControlClass struct {
	class objc.Class
}

// An interface definition for the [PathControl] class.
type IPathControl interface {
	IControl
}

// A display of a file system path or virtual path information.
//
// The class uses to implement its user interface. provides cover methods for most methods—the cover method simply invokes the corresponding cell method. See also , which represents individual components of the path, and two associated protocols: and . has three styles represented by the enumeration constants , , and . The represented path can be a file system path or any other type of path leading through a sequence of nodes or components, as defined by the programmer. automatically supports drag and drop, which can be further customized via delegate methods. To accept drag and drop, calls with and . When the URL value in the object changes because of an automatic drag and drop operation or the user selecting a new path via the open panel, the action is sent. In OS X v10.5 the value returned by is , in macOS 10.6 and later, returns the clicked cell.
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
