// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ICScannerFeature] class.
var (
	ICScannerFeatureClass     _ICScannerFeatureClass
	ICScannerFeatureClassOnce sync.Once
)

func getICScannerFeatureClass() _ICScannerFeatureClass {
	ICScannerFeatureClassOnce.Do(func() {
		ICScannerFeatureClass = _ICScannerFeatureClass{objc.GetClass("ICScannerFeature")}
	})
	return ICScannerFeatureClass
}

type _ICScannerFeatureClass struct {
	class objc.Class
}

// An interface definition for the [ICScannerFeature] class.
type IICScannerFeature interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other ImageCaptureCore classes.


// A parent class referenced by other ImageCaptureCore classes. [Full Topic]
type ICScannerFeature struct {
	objectivec.Object
}

// ICScannerFeatureFrom constructs a [ICScannerFeature] from an unsafe.Pointer.
//
// A parent class referenced by other ImageCaptureCore classes.
func ICScannerFeatureFrom(ptr unsafe.Pointer) ICScannerFeature {
	return ICScannerFeature{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureClass) Alloc() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICScannerFeatureClass) New() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeature) Init() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeature) Autorelease() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeature creates a new ICScannerFeature instance.
func NewICScannerFeature() ICScannerFeature {
	return getICScannerFeatureClass().New()
}




