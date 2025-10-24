// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ICScannerFeatureTemplate] class.
var (
	ICScannerFeatureTemplateClass     _ICScannerFeatureTemplateClass
	ICScannerFeatureTemplateClassOnce sync.Once
)

func getICScannerFeatureTemplateClass() _ICScannerFeatureTemplateClass {
	ICScannerFeatureTemplateClassOnce.Do(func() {
		ICScannerFeatureTemplateClass = _ICScannerFeatureTemplateClass{objc.GetClass("ICScannerFeatureTemplate")}
	})
	return ICScannerFeatureTemplateClass
}

type _ICScannerFeatureTemplateClass struct {
	class objc.Class
}

// An interface definition for the [ICScannerFeatureTemplate] class.
type IICScannerFeatureTemplate interface {
	IICScannerFeature
	// properties:
	Targets() []objc.IObject /* cross-framework: MutableArray */
	// methods:
}

// A group of one or more rectangular scan areas that can be used with a scanner functional unit.


// A group of one or more rectangular scan areas that can be used with a scanner functional unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureTemplate
type ICScannerFeatureTemplate struct {
	ICScannerFeature
}

// ICScannerFeatureTemplateFrom constructs a [ICScannerFeatureTemplate] from an unsafe.Pointer.
//
// A group of one or more rectangular scan areas that can be used with a scanner functional unit.
func ICScannerFeatureTemplateFrom(ptr unsafe.Pointer) ICScannerFeatureTemplate {
	return ICScannerFeatureTemplate{
		ICScannerFeature: ICScannerFeatureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ICScannerFeatureTemplateClass) Alloc() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ICScannerFeatureTemplateClass) New() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFeatureTemplate) Init() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFeatureTemplate) Autorelease() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureTemplate creates a new ICScannerFeatureTemplate instance.
func NewICScannerFeatureTemplate() ICScannerFeatureTemplate {
	return getICScannerFeatureTemplateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureTemplate/targets
func (i_ ICScannerFeatureTemplate) Targets() []objc.IObject /* cross-framework: MutableArray */ {
	rv := objc.Send[[]foundation.MutableArray](i_.ID, objc.Sel("targets"))
	return rv
}




