// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AnnouncementReason() objc.IObject /* cross-framework: NSNumber */
	SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	MetadataForNode() objc.IObject /* cross-framework: Data */
	SetMetadataForNode(value objc.IObject /* cross-framework: Data */)
	ProviderNodeID() objc.IObject /* cross-framework: NSNumber */
	SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */)
	ProviderNodeId() objc.IObject /* cross-framework: NSNumber */
	SetProviderNodeId(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/announcementreason
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) AnnouncementReason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("announcementReason"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/announcementreason
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnouncementReason:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/metadatafornode
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) MetadataForNode() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForNode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/metadatafornode
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetMetadataForNode(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForNode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l401
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ProviderNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("providerNodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l401
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l40x
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ProviderNodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("providerNodeId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l40x
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetProviderNodeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g13v
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g13v
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g14r
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g14r
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



