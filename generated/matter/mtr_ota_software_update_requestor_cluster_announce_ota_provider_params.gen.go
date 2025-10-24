// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AnnouncementReason() objc.IObject /* cross-framework: NSNumber */
	SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	MetadataForNode() objc.IObject /* cross-framework: Data */
	SetMetadataForNode(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/announcementreason
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) AnnouncementReason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("announcementReason"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/announcementreason
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnouncementReason:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/metadatafornode
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) MetadataForNode() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForNode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/metadatafornode
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetMetadataForNode(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForNode:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-1ucwe/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



