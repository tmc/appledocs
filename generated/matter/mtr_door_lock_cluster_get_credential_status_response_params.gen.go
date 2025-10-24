// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetCredentialStatusResponseParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetCredentialStatusResponseParams */
// The class instance for the [MTRDoorLockClusterGetCredentialStatusResponseParams] class.
var (
	MTRDoorLockClusterGetCredentialStatusResponseParamsClass     _MTRDoorLockClusterGetCredentialStatusResponseParamsClass
	MTRDoorLockClusterGetCredentialStatusResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetCredentialStatusResponseParamsClass() _MTRDoorLockClusterGetCredentialStatusResponseParamsClass {
	MTRDoorLockClusterGetCredentialStatusResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetCredentialStatusResponseParamsClass = _MTRDoorLockClusterGetCredentialStatusResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetCredentialStatusResponseParams")}
	})
	return MTRDoorLockClusterGetCredentialStatusResponseParamsClass
}

type _MTRDoorLockClusterGetCredentialStatusResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetCredentialStatusResponseParams */
// An interface definition for the [MTRDoorLockClusterGetCredentialStatusResponseParams] class.
type IMTRDoorLockClusterGetCredentialStatusResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetCredentialStatusResponseParams */
	// properties:
	CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	CredentialData() objc.IObject /* cross-framework: NSData */
	SetCredentialData(value objc.IObject /* cross-framework: NSData */)
	CredentialExists() objc.IObject /* cross-framework: NSNumber */
	SetCredentialExists(value objc.IObject /* cross-framework: NSNumber */)
	LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */
	SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetCredentialStatusResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetCredentialStatusResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetCredentialStatusResponseParamsClass) Alloc() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetCredentialStatusResponseParamsClass) New() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) Init() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) Autorelease() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetCredentialStatusResponseParams creates a new MTRDoorLockClusterGetCredentialStatusResponseParams instance.
func NewMTRDoorLockClusterGetCredentialStatusResponseParams() MTRDoorLockClusterGetCredentialStatusResponseParams {
	return getMTRDoorLockClusterGetCredentialStatusResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetCredentialStatusResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams
type MTRDoorLockClusterGetCredentialStatusResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetCredentialStatusResponseParamsFrom constructs a [MTRDoorLockClusterGetCredentialStatusResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetCredentialStatusResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetCredentialStatusResponseParams {
	return MTRDoorLockClusterGetCredentialStatusResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetCredentialStatusResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/init(responseValue:)
func NewMTRDoorLockClusterGetCredentialStatusResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDoorLockClusterGetCredentialStatusResponseParams {
	instance := getMTRDoorLockClusterGetCredentialStatusResponseParamsClass().Alloc()
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDoorLockClusterGetCredentialStatusResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetCredentialStatusResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetCredentialStatusResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetCredentialStatusResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetCredentialStatusResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/creatorFabricIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("creatorFabricIndex"))
	return rv
}/* debug [instance_properties/getter]: creatorFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/creatorFabricIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreatorFabricIndex:"), value)
}/* debug [instance_properties/setter]: creatorFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/credentialData
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CredentialData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("credentialData"))
	return rv
}/* debug [instance_properties/getter]: credentialData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/credentialData
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCredentialData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialData:"), value)
}/* debug [instance_properties/setter]: credentialData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/credentialExists
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CredentialExists() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialExists"))
	return rv
}/* debug [instance_properties/getter]: credentialExists */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/credentialExists
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCredentialExists(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialExists:"), value)
}/* debug [instance_properties/setter]: credentialExists */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/lastModifiedFabricIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lastModifiedFabricIndex"))
	return rv
}/* debug [instance_properties/getter]: lastModifiedFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/lastModifiedFabricIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastModifiedFabricIndex:"), value)
}/* debug [instance_properties/setter]: lastModifiedFabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/nextCredentialIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextCredentialIndex"))
	return rv
}/* debug [instance_properties/getter]: nextCredentialIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/nextCredentialIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextCredentialIndex:"), value)
}/* debug [instance_properties/setter]: nextCredentialIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/userIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams/userIndex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetCredentialStatusResponseParams */


