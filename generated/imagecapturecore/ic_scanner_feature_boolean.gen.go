// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ICScannerFeatureBoolean] class.
var (
	ICScannerFeatureBooleanClass     _ICScannerFeatureBooleanClass
	ICScannerFeatureBooleanClassOnce sync.Once
)

func getICScannerFeatureBooleanClass() _ICScannerFeatureBooleanClass {
	ICScannerFeatureBooleanClassOnce.Do(func() {
		ICScannerFeatureBooleanClass = _ICScannerFeatureBooleanClass{objc.GetClass("ICScannerFeatureBoolean")}
	})
	return ICScannerFeatureBooleanClass
}

type _ICScannerFeatureBooleanClass struct {
	class objc.Class
}

// An interface definition for the [ICScannerFeatureBoolean] class.
type IICScannerFeatureBoolean interface {
	objectivec.IObject
	Value() bool
	SetValue(value bool)
}

// A feature with a value of or .


// A feature with a value of or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean
type ICScannerFeatureBoolean struct {
	objectivec.Object
}

// ICScannerFeatureBooleanFrom constructs a [ICScannerFeatureBoolean] from an unsafe.Pointer.
//
// A feature with a value of or .
func ICScannerFeatureBooleanFrom(ptr unsafe.Pointer) ICScannerFeatureBoolean {
	return ICScannerFeatureBoolean{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureBooleanClass) Alloc() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICScannerFeatureBooleanClass) New() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeatureBoolean) Init() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeatureBoolean) Autorelease() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureBoolean creates a new ICScannerFeatureBoolean instance.
func NewICScannerFeatureBoolean() ICScannerFeatureBoolean {
	return getICScannerFeatureBooleanClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean/value
func (i_ ICScannerFeatureBoolean) Value() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean/value
func (i_ ICScannerFeatureBoolean) SetValue(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setValue:"), value)
}




