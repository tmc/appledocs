// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams] class.
var (
	MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass     _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass
	MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass() _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass {
	MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass = _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams")}
	})
	return MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass
}

type _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams] class.
type IMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu
type MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsFrom constructs a [MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	return MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass) Alloc() MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass) New() MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) Init() MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) Autorelease() MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams creates a new MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams instance.
func NewMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams() MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	return getMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass().New()
}




