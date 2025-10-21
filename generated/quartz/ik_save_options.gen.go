// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [IKSaveOptions] class.
var (
	IKSaveOptionsClass     _IKSaveOptionsClass
	IKSaveOptionsClassOnce sync.Once
)

func getIKSaveOptionsClass() _IKSaveOptionsClass {
	IKSaveOptionsClassOnce.Do(func() {
		IKSaveOptionsClass = _IKSaveOptionsClass{objc.GetClass("IKSaveOptions")}
	})
	return IKSaveOptionsClass
}

type _IKSaveOptionsClass struct {
	class objc.Class
}

// An interface definition for the [IKSaveOptions] class.
type IIKSaveOptions interface {
	objectivec.IObject
}

// The class initializes, adds, and manages user interface options for saving image data.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions
type IKSaveOptions struct {
	objectivec.Object
}

// IKSaveOptionsFrom constructs a [IKSaveOptions] from an unsafe.Pointer.
//
// The class initializes, adds, and manages user interface options for saving image data.
func IKSaveOptionsFrom(ptr unsafe.Pointer) IKSaveOptions {
	return IKSaveOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IKSaveOptionsClass) Alloc() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKSaveOptionsClass) New() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKSaveOptions) Init() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKSaveOptions) Autorelease() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKSaveOptions creates a new IKSaveOptions instance.
func NewIKSaveOptions() IKSaveOptions {
	return getIKSaveOptionsClass().New()
}




