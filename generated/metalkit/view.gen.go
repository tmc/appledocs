// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [View] class.
var viewClass = _ViewClass{objc.GetClass("NSView")}

type _ViewClass struct {
	class objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	objectivec.IObject
}

// A parent class referenced by other MetalKit classes. [Full Topic]

type View struct {
	objectivec.Object
}

// ViewFrom constructs a [View] from an unsafe.Pointer.
//
// A parent class referenced by other MetalKit classes.
func ViewFrom(ptr unsafe.Pointer) View {
	return View{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ View) Autorelease() View {
	rv := objc.Send[View](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewView creates a new View instance.
func NewView() View {
	return viewClass.New()
}




