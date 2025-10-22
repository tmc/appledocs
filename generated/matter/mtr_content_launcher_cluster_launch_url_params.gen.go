// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterLaunchURLParams] class.
var (
	MTRContentLauncherClusterLaunchURLParamsClass     _MTRContentLauncherClusterLaunchURLParamsClass
	MTRContentLauncherClusterLaunchURLParamsClassOnce sync.Once
)

func getMTRContentLauncherClusterLaunchURLParamsClass() _MTRContentLauncherClusterLaunchURLParamsClass {
	MTRContentLauncherClusterLaunchURLParamsClassOnce.Do(func() {
		MTRContentLauncherClusterLaunchURLParamsClass = _MTRContentLauncherClusterLaunchURLParamsClass{objc.GetClass("MTRContentLauncherClusterLaunchURLParams")}
	})
	return MTRContentLauncherClusterLaunchURLParamsClass
}

type _MTRContentLauncherClusterLaunchURLParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterLaunchURLParams] class.
type IMTRContentLauncherClusterLaunchURLParams interface {
	objectivec.IObject
	BrandingInformation() MTRContentLauncherClusterBrandingInformationStruct
	SetBrandingInformation(value IMTRContentLauncherClusterBrandingInformationStruct)
	ContentURL() string
	SetContentURL(value string)
	DisplayString() string
	SetDisplayString(value string)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterLaunchURLParams
type MTRContentLauncherClusterLaunchURLParams struct {
	objectivec.Object
}

// MTRContentLauncherClusterLaunchURLParamsFrom constructs a [MTRContentLauncherClusterLaunchURLParams] from an unsafe.Pointer.
func MTRContentLauncherClusterLaunchURLParamsFrom(ptr unsafe.Pointer) MTRContentLauncherClusterLaunchURLParams {
	return MTRContentLauncherClusterLaunchURLParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterLaunchURLParamsClass) Alloc() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterLaunchURLParamsClass) New() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterLaunchURLParams) Init() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterLaunchURLParams) Autorelease() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterLaunchURLParams creates a new MTRContentLauncherClusterLaunchURLParams instance.
func NewMTRContentLauncherClusterLaunchURLParams() MTRContentLauncherClusterLaunchURLParams {
	return getMTRContentLauncherClusterLaunchURLParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/brandinginformation
func (m_ MTRContentLauncherClusterLaunchURLParams) BrandingInformation() MTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](m_.ID, objc.Sel("brandingInformation"))
	return rv
}


// SetBrandingInformation sets the value of the brandingInformation property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/brandinginformation
func (m_ MTRContentLauncherClusterLaunchURLParams) SetBrandingInformation(value IMTRContentLauncherClusterBrandingInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBrandingInformation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/contenturl
func (m_ MTRContentLauncherClusterLaunchURLParams) ContentURL() string {
	rv := objc.Send[string](m_.ID, objc.Sel("contentURL"))
	return rv
}


// SetContentURL sets the value of the contentURL property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/contenturl
func (m_ MTRContentLauncherClusterLaunchURLParams) SetContentURL(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContentURL:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/displaystring
func (m_ MTRContentLauncherClusterLaunchURLParams) DisplayString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("displayString"))
	return rv
}


// SetDisplayString sets the value of the displayString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/displaystring
func (m_ MTRContentLauncherClusterLaunchURLParams) SetDisplayString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchURLParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchURLParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchURLParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchURLParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



