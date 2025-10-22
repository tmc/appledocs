// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKFilterUIView] class.
var (
	IKFilterUIViewClass     _IKFilterUIViewClass
	IKFilterUIViewClassOnce sync.Once
)

func getIKFilterUIViewClass() _IKFilterUIViewClass {
	IKFilterUIViewClassOnce.Do(func() {
		IKFilterUIViewClass = _IKFilterUIViewClass{objc.GetClass("IKFilterUIView")}
	})
	return IKFilterUIViewClass
}

type _IKFilterUIViewClass struct {
	class objc.Class
}

// An interface definition for the [IKFilterUIView] class.
type IIKFilterUIView interface {
	appkit.IView
}

// Input parameters for filtering core image filters.
//
// The class provides a view that contains input parameter controls for a Core Image filter ( ). You need to use this class when providing a user interface for a custom filter. The class creates a view that has an object controller for the given filter. It also retains the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterUIView
type IKFilterUIView struct {
	appkit.View
}

// IKFilterUIViewFrom constructs a [IKFilterUIView] from an unsafe.Pointer.
//
// Input parameters for filtering core image filters.
func IKFilterUIViewFrom(ptr unsafe.Pointer) IKFilterUIView {
	return IKFilterUIView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKFilterUIViewClass) Alloc() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKFilterUIViewClass) New() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKFilterUIView) Init() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKFilterUIView) Autorelease() IKFilterUIView {
	rv := objc.Send[IKFilterUIView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKFilterUIView creates a new IKFilterUIView instance.
func NewIKFilterUIView() IKFilterUIView {
	return getIKFilterUIViewClass().New()
}




