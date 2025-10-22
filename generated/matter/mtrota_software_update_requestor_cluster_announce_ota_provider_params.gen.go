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
	AnnouncementReason() foundation.Number
	SetAnnouncementReason(value foundation.INumber)
	Endpoint() foundation.Number
	SetEndpoint(value foundation.INumber)
	MetadataForNode() foundation.Data
	SetMetadataForNode(value foundation.IData)
	ProviderNodeID() foundation.Number
	SetProviderNodeID(value foundation.INumber)
	ProviderNodeId() foundation.Number
	SetProviderNodeId(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	VendorID() foundation.Number
	SetVendorID(value foundation.INumber)
	VendorId() foundation.Number
	SetVendorId(value foundation.INumber)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/announcementreason
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) AnnouncementReason() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("announcementReason"))
	return rv
}


// SetAnnouncementReason sets the value of the announcementReason property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/announcementreason
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetAnnouncementReason(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnouncementReason:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/metadatafornode
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) MetadataForNode() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForNode"))
	return rv
}


// SetMetadataForNode sets the value of the metadataForNode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/metadatafornode
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetMetadataForNode(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForNode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l401
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ProviderNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("providerNodeID"))
	return rv
}


// SetProviderNodeID sets the value of the providerNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l401
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetProviderNodeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l40x
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ProviderNodeId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("providerNodeId"))
	return rv
}


// SetProviderNodeId sets the value of the providerNodeId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/providernodeid-7l40x
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetProviderNodeId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g13v
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g13v
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetVendorID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g14r
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) VendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorId"))
	return rv
}


// SetVendorId sets the value of the vendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterannounceotaproviderparams-8dobu/vendorid-4g14r
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetVendorId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



