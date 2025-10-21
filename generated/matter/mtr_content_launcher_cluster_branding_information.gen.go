// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRContentLauncherClusterBrandingInformation] class.
var (
	MTRContentLauncherClusterBrandingInformationClass     _MTRContentLauncherClusterBrandingInformationClass
	MTRContentLauncherClusterBrandingInformationClassOnce sync.Once
)

func getMTRContentLauncherClusterBrandingInformationClass() _MTRContentLauncherClusterBrandingInformationClass {
	MTRContentLauncherClusterBrandingInformationClassOnce.Do(func() {
		MTRContentLauncherClusterBrandingInformationClass = _MTRContentLauncherClusterBrandingInformationClass{objc.GetClass("MTRContentLauncherClusterBrandingInformation")}
	})
	return MTRContentLauncherClusterBrandingInformationClass
}

type _MTRContentLauncherClusterBrandingInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterBrandingInformation] class.
type IMTRContentLauncherClusterBrandingInformation interface {
	IMTRContentLauncherClusterBrandingInformationStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation
type MTRContentLauncherClusterBrandingInformation struct {
	MTRContentLauncherClusterBrandingInformationStruct
}

// MTRContentLauncherClusterBrandingInformationFrom constructs a [MTRContentLauncherClusterBrandingInformation] from an unsafe.Pointer.
func MTRContentLauncherClusterBrandingInformationFrom(ptr unsafe.Pointer) MTRContentLauncherClusterBrandingInformation {
	return MTRContentLauncherClusterBrandingInformation{
		MTRContentLauncherClusterBrandingInformationStruct: MTRContentLauncherClusterBrandingInformationStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterBrandingInformationClass) Alloc() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterBrandingInformationClass) New() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterBrandingInformation) Init() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterBrandingInformation) Autorelease() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterBrandingInformation creates a new MTRContentLauncherClusterBrandingInformation instance.
func NewMTRContentLauncherClusterBrandingInformation() MTRContentLauncherClusterBrandingInformation {
	return getMTRContentLauncherClusterBrandingInformationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/background
func (m_ MTRContentLauncherClusterBrandingInformation) Background() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("background"))
	return rv
}


// SetBackground sets the value of the background property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/background
func (m_ MTRContentLauncherClusterBrandingInformation) SetBackground(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackground:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/logo
func (m_ MTRContentLauncherClusterBrandingInformation) Logo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("logo"))
	return rv
}


// SetLogo sets the value of the logo property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/logo
func (m_ MTRContentLauncherClusterBrandingInformation) SetLogo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/progressbar
func (m_ MTRContentLauncherClusterBrandingInformation) ProgressBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("progressBar"))
	return rv
}


// SetProgressBar sets the value of the progressBar property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/progressbar
func (m_ MTRContentLauncherClusterBrandingInformation) SetProgressBar(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressBar:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/providername
func (m_ MTRContentLauncherClusterBrandingInformation) ProviderName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("providerName"))
	return rv
}


// SetProviderName sets the value of the providerName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/providername
func (m_ MTRContentLauncherClusterBrandingInformation) SetProviderName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/splash
func (m_ MTRContentLauncherClusterBrandingInformation) Splash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("splash"))
	return rv
}


// SetSplash sets the value of the splash property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/splash
func (m_ MTRContentLauncherClusterBrandingInformation) SetSplash(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSplash:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/watermark
func (m_ MTRContentLauncherClusterBrandingInformation) WaterMark() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("waterMark"))
	return rv
}


// SetWaterMark sets the value of the waterMark property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterbrandinginformation/watermark
func (m_ MTRContentLauncherClusterBrandingInformation) SetWaterMark(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWaterMark:"), value)
}



