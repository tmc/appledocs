// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ImageProperties() unsafe.Pointer
	SetImageProperties(value unsafe.Pointer)
	ImageUTType() objc.IObject /* cross-framework: NSString */
	SetImageUTType(value objc.IObject /* cross-framework: NSString */)
	RememberLastSetting() bool
	SetRememberLastSetting(value bool)
	UserSelection() unsafe.Pointer
	SetUserSelection(value unsafe.Pointer)
	// methods:
}

// The class initializes, adds, and manages user interface options for saving image data.


// The class initializes, adds, and manages user interface options for saving image data.
//
// [Full Topic]
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



// Specifies the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/delegate
func (i_ IKSaveOptions) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// Specifies the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/delegate
func (i_ IKSaveOptions) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}


// Returns a dictionary of updated image properties that reflects the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/imageproperties
func (i_ IKSaveOptions) ImageProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageProperties"))
	return rv
}


// Returns a dictionary of updated image properties that reflects the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/imageproperties
func (i_ IKSaveOptions) SetImageProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageProperties:"), value)
}


// Returns the uniform type identifier that reflects the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/imageuttype
func (i_ IKSaveOptions) ImageUTType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("imageUTType"))
	return rv
}


// Returns the uniform type identifier that reflects the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/imageuttype
func (i_ IKSaveOptions) SetImageUTType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageUTType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/rememberlastsetting
func (i_ IKSaveOptions) RememberLastSetting() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("rememberLastSetting"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/rememberlastsetting
func (i_ IKSaveOptions) SetRememberLastSetting(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRememberLastSetting:"), value)
}


// Returns a dictionary that contains the save options selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/userselection
func (i_ IKSaveOptions) UserSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("userSelection"))
	return rv
}


// Returns a dictionary that contains the save options selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/iksaveoptions/userselection
func (i_ IKSaveOptions) SetUserSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUserSelection:"), value)
}



