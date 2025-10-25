// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetResourceLoadingDataRequest */


/* debug [class_header]: Header for AVAssetResourceLoadingDataRequest */
// The class instance for the [AssetResourceLoadingDataRequest] class.
var (
	AssetResourceLoadingDataRequestClass     _AssetResourceLoadingDataRequestClass
	AssetResourceLoadingDataRequestClassOnce sync.Once
)

func getAssetResourceLoadingDataRequestClass() _AssetResourceLoadingDataRequestClass {
	AssetResourceLoadingDataRequestClassOnce.Do(func() {
		AssetResourceLoadingDataRequestClass = _AssetResourceLoadingDataRequestClass{objc.GetClass("AVAssetResourceLoadingDataRequest")}
	})
	return AssetResourceLoadingDataRequestClass
}

type _AssetResourceLoadingDataRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetResourceLoadingDataRequest */
// An interface definition for the [AssetResourceLoadingDataRequest] class.
type IAssetResourceLoadingDataRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetResourceLoadingDataRequest */
	// properties:
	CurrentOffset() objectivec.IObject
	RequestedLength() int
	RequestedOffset() objectivec.IObject
	RequestsAllDataToEndOfResource() bool
	DataRequest() IAVAssetResourceLoadingDataRequest
	SetDataRequest(value IAVAssetResourceLoadingDataRequest)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetResourceLoadingDataRequest */
	// methods:
	RespondWithData(data objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetResourceLoadingDataRequest */
// Alloc allocates a new instance without initialization.
func (ac _AssetResourceLoadingDataRequestClass) Alloc() AssetResourceLoadingDataRequest {
	rv := objc.Send[AssetResourceLoadingDataRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetResourceLoadingDataRequestClass) New() AssetResourceLoadingDataRequest {
	rv := objc.Send[AssetResourceLoadingDataRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetResourceLoadingDataRequest) Init() AssetResourceLoadingDataRequest {
	rv := objc.Send[AssetResourceLoadingDataRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetResourceLoadingDataRequest) Autorelease() AssetResourceLoadingDataRequest {
	rv := objc.Send[AssetResourceLoadingDataRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetResourceLoadingDataRequest creates a new AssetResourceLoadingDataRequest instance.
func NewAssetResourceLoadingDataRequest() AssetResourceLoadingDataRequest {
	return getAssetResourceLoadingDataRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetResourceLoadingDataRequest */
// An object for requesting data from a resource that an asset resource-loading request references.
//
// The uses the class to do the actual data reading, and its methods will be invoked, as necessary, to acquire data for the instance. When the resource loading delegate, which implements the  protocol, receives an instance of  as the second parameter of the delegate’s method, it has the option of accepting responsibility for loading the referenced resource. If it accepts that responsibility, by returning , it must check whether the  property of the  instance is not . If it is not , the resource loading delegate is informed of the range of bytes within the resource that are required by the underlying media system. In response, the data is provided by one or more invocations of as required to provide the requested data. The data can be provided in increments determined by the resource loading delegate according to convenience or efficiency. When the method is invoked, the data request is considered fully satisfied. If the entire range of bytes requested has not yet been provided, the underlying media system assumes that the resource’s length is limited to the provided content.


// An object for requesting data from a resource that an asset resource-loading request references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest
type AssetResourceLoadingDataRequest struct {
	objectivec.Object
}

// AssetResourceLoadingDataRequestFrom constructs a [AssetResourceLoadingDataRequest] from an unsafe.Pointer.
//
// An object for requesting data from a resource that an asset resource-loading request references.
func AssetResourceLoadingDataRequestFrom(ptr unsafe.Pointer) AssetResourceLoadingDataRequest {
	return AssetResourceLoadingDataRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetResourceLoadingDataRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetResourceLoadingDataRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetResourceLoadingDataRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetResourceLoadingDataRequest */

// Provides data to the loading request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/respond(with:)
func (a_ AssetResourceLoadingDataRequest) RespondWithData(data objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("respondWithData:"), data)
}/* debug [instance_methods/method]: RespondWithData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetResourceLoadingDataRequest */

// The position within the resource of the next byte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/currentOffset
func (a_ AssetResourceLoadingDataRequest) CurrentOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("currentOffset"))
	return rv
}/* debug [instance_properties/getter]: currentOffset */


// The length, in bytes, of the data requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/requestedLength
func (a_ AssetResourceLoadingDataRequest) RequestedLength() int {
	rv := objc.Send[int](a_.ID, objc.Sel("requestedLength"))
	return rv
}/* debug [instance_properties/getter]: requestedLength */


// The position within the resource of the first byte requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/requestedOffset
func (a_ AssetResourceLoadingDataRequest) RequestedOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("requestedOffset"))
	return rv
}/* debug [instance_properties/getter]: requestedOffset */


// A Boolean value that indicates the entire remaining length of the resource from the offest to the end of the resource is being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingDataRequest/requestsAllDataToEndOfResource
func (a_ AssetResourceLoadingDataRequest) RequestsAllDataToEndOfResource() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("requestsAllDataToEndOfResource"))
	return rv
}/* debug [instance_properties/getter]: requestsAllDataToEndOfResource */


// The range of requested resource data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/datarequest
func (a_ AssetResourceLoadingDataRequest) DataRequest() IAVAssetResourceLoadingDataRequest {
	rv := objc.Send[AssetResourceLoadingDataRequest](a_.ID, objc.Sel("dataRequest"))
	return rv
}/* debug [instance_properties/getter]: dataRequest */


// The range of requested resource data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingrequest/datarequest
func (a_ AssetResourceLoadingDataRequest) SetDataRequest(value IAVAssetResourceLoadingDataRequest) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataRequest:"), value)
}/* debug [instance_properties/setter]: dataRequest */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetResourceLoadingDataRequest */



