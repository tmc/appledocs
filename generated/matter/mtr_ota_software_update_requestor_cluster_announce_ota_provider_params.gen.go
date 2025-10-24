// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */


/* debug [class_header]: Header for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams] class.
type IMTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams interface {
	IMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
	// properties:
	AnnouncementReason() objc.IObject /* cross-framework: NSNumber */
	SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	MetadataForNode() objc.IObject /* cross-framework: NSData */
	SetMetadataForNode(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParamsClass) Alloc() MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/announcementReason
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) AnnouncementReason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("announcementReason"))
	return rv
}/* debug [instance_properties/getter]: announcementReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/announcementReason
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetAnnouncementReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnouncementReason:"), value)
}/* debug [instance_properties/setter]: announcementReason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/endpoint
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}/* debug [instance_properties/setter]: endpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/metadataForNode
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) MetadataForNode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("metadataForNode"))
	return rv
}/* debug [instance_properties/getter]: metadataForNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/metadataForNode
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetMetadataForNode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForNode:"), value)
}/* debug [instance_properties/setter]: metadataForNode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams-1ucwe/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateRequestorClusterAnnounceOtaProviderParams */



