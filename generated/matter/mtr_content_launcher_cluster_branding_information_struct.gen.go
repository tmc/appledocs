// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Background() IMTRContentLauncherClusterStyleInformationStruct
	SetBackground(value IMTRContentLauncherClusterStyleInformationStruct)
	Logo() IMTRContentLauncherClusterStyleInformationStruct
	SetLogo(value IMTRContentLauncherClusterStyleInformationStruct)
	ProgressBar() IMTRContentLauncherClusterStyleInformationStruct
	SetProgressBar(value IMTRContentLauncherClusterStyleInformationStruct)
	ProviderName() objc.IObject /* cross-framework: NSString */
	SetProviderName(value objc.IObject /* cross-framework: NSString */)
	Splash() IMTRContentLauncherClusterStyleInformationStruct
	SetSplash(value IMTRContentLauncherClusterStyleInformationStruct)
	WaterMark() IMTRContentLauncherClusterStyleInformationStruct
	SetWaterMark(value IMTRContentLauncherClusterStyleInformationStruct)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/background
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Background() IMTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("background"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/background
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetBackground(value IMTRContentLauncherClusterStyleInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackground:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/logo
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Logo() IMTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("logo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/logo
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetLogo(value IMTRContentLauncherClusterStyleInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/progressbar
func (m_ MTRContentLauncherClusterBrandingInformationStruct) ProgressBar() IMTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("progressBar"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/progressbar
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetProgressBar(value IMTRContentLauncherClusterStyleInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressBar:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/providername
func (m_ MTRContentLauncherClusterBrandingInformationStruct) ProviderName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("providerName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/providername
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetProviderName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/splash
func (m_ MTRContentLauncherClusterBrandingInformationStruct) Splash() IMTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("splash"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/splash
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetSplash(value IMTRContentLauncherClusterStyleInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSplash:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/watermark
func (m_ MTRContentLauncherClusterBrandingInformationStruct) WaterMark() IMTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("waterMark"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformationstruct/watermark
func (m_ MTRContentLauncherClusterBrandingInformationStruct) SetWaterMark(value IMTRContentLauncherClusterStyleInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWaterMark:"), value)
}



