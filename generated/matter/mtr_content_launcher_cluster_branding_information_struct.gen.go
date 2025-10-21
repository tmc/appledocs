// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterBrandingInformationStruct] class.
var (
	MTRContentLauncherClusterBrandingInformationStructClass     _MTRContentLauncherClusterBrandingInformationStructClass
	MTRContentLauncherClusterBrandingInformationStructClassOnce sync.Once
)

func getMTRContentLauncherClusterBrandingInformationStructClass() _MTRContentLauncherClusterBrandingInformationStructClass {
	MTRContentLauncherClusterBrandingInformationStructClassOnce.Do(func() {
		MTRContentLauncherClusterBrandingInformationStructClass = _MTRContentLauncherClusterBrandingInformationStructClass{objc.GetClass("MTRContentLauncherClusterBrandingInformationStruct")}
	})
	return MTRContentLauncherClusterBrandingInformationStructClass
}

type _MTRContentLauncherClusterBrandingInformationStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterBrandingInformationStruct] class.
type IMTRContentLauncherClusterBrandingInformationStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformationStruct
type MTRContentLauncherClusterBrandingInformationStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterBrandingInformationStructFrom constructs a [MTRContentLauncherClusterBrandingInformationStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterBrandingInformationStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterBrandingInformationStruct {
	return MTRContentLauncherClusterBrandingInformationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterBrandingInformationStructClass) Alloc() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterBrandingInformationStructClass) New() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Init() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Autorelease() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterBrandingInformationStruct creates a new MTRContentLauncherClusterBrandingInformationStruct instance.
func NewMTRContentLauncherClusterBrandingInformationStruct() MTRContentLauncherClusterBrandingInformationStruct {
	return getMTRContentLauncherClusterBrandingInformationStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/progressbar
func (m_ MTRContentLauncherClusterBrandingInformationStruct) ProgressBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("progressBar"))
	return rv
}


// SetProgressBar sets the value of the progressBar property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/progressbar
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetProgressBar(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressBar:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/watermark
func (m_ MTRContentLauncherClusterBrandingInformationStruct) WaterMark() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("waterMark"))
	return rv
}


// SetWaterMark sets the value of the waterMark property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/watermark
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetWaterMark(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWaterMark:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/providername
func (m_ MTRContentLauncherClusterBrandingInformationStruct) ProviderName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("providerName"))
	return rv
}


// SetProviderName sets the value of the providerName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/providername
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetProviderName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/splash
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Splash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("splash"))
	return rv
}


// SetSplash sets the value of the splash property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/splash
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetSplash(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSplash:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/logo
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Logo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("logo"))
	return rv
}


// SetLogo sets the value of the logo property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/logo
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetLogo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/background
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Background() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("background"))
	return rv
}


// SetBackground sets the value of the background property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/background
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetBackground(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackground:"), value)
}



