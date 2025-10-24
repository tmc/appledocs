// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */


/* debug [class_header]: Header for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
// An interface definition for the [MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams] class.
type IMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
	// properties:
	AnnouncementReason() objc.IObject /* cross-framework: NSNumber */
	SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	MetadataForNode() objc.IObject /* cross-framework: NSData */
	SetMetadataForNode(value objc.IObject /* cross-framework: NSData */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsClass) Alloc() MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu
type MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsFrom constructs a [MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams {
	return MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/announcementReason
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) AnnouncementReason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("announcementReason"))
	return rv
}/* debug [instance_properties/getter]: announcementReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/announcementReason
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnouncementReason:"), value)
}/* debug [instance_properties/setter]: announcementReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/endpoint
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/metadataForNode
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) MetadataForNode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("metadataForNode"))
	return rv
}/* debug [instance_properties/getter]: metadataForNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/metadataForNode
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetMetadataForNode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForNode:"), value)
}/* debug [instance_properties/setter]: metadataForNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/providerNodeID-7l401
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ProviderNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("providerNodeID"))
	return rv
}/* debug [instance_properties/getter]: providerNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/providerNodeID-7l401
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetProviderNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeID:"), value)
}/* debug [instance_properties/setter]: providerNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/providerNodeId-7l40x
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ProviderNodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("providerNodeId"))
	return rv
}/* debug [instance_properties/getter]: providerNodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/providerNodeId-7l40x
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetProviderNodeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProviderNodeId:"), value)
}/* debug [instance_properties/setter]: providerNodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/vendorID-4g13v
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/vendorID-4g13v
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}/* debug [instance_properties/setter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/vendorId-4g14r
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}/* debug [instance_properties/getter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams-8dobu/vendorId-4g14r
func (m_ MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}/* debug [instance_properties/setter]: vendorId */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams */



