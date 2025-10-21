// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams] class.
var (
	MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass     _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass
	MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass() _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass {
	MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass = _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams")}
	})
	return MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass
}

type _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams] class.
type IMTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams interface {
	IMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe
type MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams struct {
	MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams
}

// MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsFrom constructs a [MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	return MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams{
		MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams: MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass) Alloc() MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass) New() MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) Init() MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) Autorelease() MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams creates a new MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams instance.
func NewMTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams() MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	return getMTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass().New()
}




