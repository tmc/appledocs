// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterUpdateNOCParams */


/* debug [class_header]: Header for MTROperationalCredentialsClusterUpdateNOCParams */
// The class instance for the [MTROperationalCredentialsClusterUpdateNOCParams] class.
var (
	MTROperationalCredentialsClusterUpdateNOCParamsClass     _MTROperationalCredentialsClusterUpdateNOCParamsClass
	MTROperationalCredentialsClusterUpdateNOCParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterUpdateNOCParamsClass() _MTROperationalCredentialsClusterUpdateNOCParamsClass {
	MTROperationalCredentialsClusterUpdateNOCParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterUpdateNOCParamsClass = _MTROperationalCredentialsClusterUpdateNOCParamsClass{objc.GetClass("MTROperationalCredentialsClusterUpdateNOCParams")}
	})
	return MTROperationalCredentialsClusterUpdateNOCParamsClass
}

type _MTROperationalCredentialsClusterUpdateNOCParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterUpdateNOCParams */
// An interface definition for the [MTROperationalCredentialsClusterUpdateNOCParams] class.
type IMTROperationalCredentialsClusterUpdateNOCParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterUpdateNOCParams */
	// properties:
	IcacValue() objc.IObject /* cross-framework: NSData */
	SetIcacValue(value objc.IObject /* cross-framework: NSData */)
	NocValue() objc.IObject /* cross-framework: NSData */
	SetNocValue(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterUpdateNOCParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterUpdateNOCParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterUpdateNOCParamsClass) Alloc() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterUpdateNOCParamsClass) New() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) Init() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) Autorelease() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterUpdateNOCParams creates a new MTROperationalCredentialsClusterUpdateNOCParams instance.
func NewMTROperationalCredentialsClusterUpdateNOCParams() MTROperationalCredentialsClusterUpdateNOCParams {
	return getMTROperationalCredentialsClusterUpdateNOCParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterUpdateNOCParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams
type MTROperationalCredentialsClusterUpdateNOCParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterUpdateNOCParamsFrom constructs a [MTROperationalCredentialsClusterUpdateNOCParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterUpdateNOCParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterUpdateNOCParams {
	return MTROperationalCredentialsClusterUpdateNOCParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterUpdateNOCParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterUpdateNOCParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterUpdateNOCParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterUpdateNOCParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterUpdateNOCParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/icacValue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) IcacValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("icacValue"))
	return rv
}/* debug [instance_properties/getter]: icacValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/icacValue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetIcacValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcacValue:"), value)
}/* debug [instance_properties/setter]: icacValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/nocValue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) NocValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("nocValue"))
	return rv
}/* debug [instance_properties/getter]: nocValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/nocValue
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetNocValue(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocValue:"), value)
}/* debug [instance_properties/setter]: nocValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/serverSideProcessingTimeout
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams/timedInvokeTimeoutMs
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterUpdateNOCParams */



