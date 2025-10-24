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
	// properties:
	BrandingInformation() IMTRContentLauncherClusterBrandingInformationStruct
	SetBrandingInformation(value IMTRContentLauncherClusterBrandingInformationStruct)
	ContentURL() objc.IObject /* cross-framework: NSString */
	SetContentURL(value objc.IObject /* cross-framework: NSString */)
	DisplayString() objc.IObject /* cross-framework: NSString */
	SetDisplayString(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/brandinginformation
func (m_ MTRContentLauncherClusterLaunchURLParams) BrandingInformation() IMTRContentLauncherClusterBrandingInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformationStruct](m_.ID, objc.Sel("brandingInformation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/brandinginformation
func (m_ MTRContentLauncherClusterLaunchURLParams) SetBrandingInformation(value IMTRContentLauncherClusterBrandingInformationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBrandingInformation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/contenturl
func (m_ MTRContentLauncherClusterLaunchURLParams) ContentURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("contentURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/contenturl
func (m_ MTRContentLauncherClusterLaunchURLParams) SetContentURL(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContentURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/displaystring
func (m_ MTRContentLauncherClusterLaunchURLParams) DisplayString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("displayString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/displaystring
func (m_ MTRContentLauncherClusterLaunchURLParams) SetDisplayString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchURLParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchURLParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchURLParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchurlparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchURLParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



