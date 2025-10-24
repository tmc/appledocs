// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROnOffClusterOffWithEffectParams */


/* debug [class_header]: Header for MTROnOffClusterOffWithEffectParams */
// The class instance for the [MTROnOffClusterOffWithEffectParams] class.
var (
	MTROnOffClusterOffWithEffectParamsClass     _MTROnOffClusterOffWithEffectParamsClass
	MTROnOffClusterOffWithEffectParamsClassOnce sync.Once
)

func getMTROnOffClusterOffWithEffectParamsClass() _MTROnOffClusterOffWithEffectParamsClass {
	MTROnOffClusterOffWithEffectParamsClassOnce.Do(func() {
		MTROnOffClusterOffWithEffectParamsClass = _MTROnOffClusterOffWithEffectParamsClass{objc.GetClass("MTROnOffClusterOffWithEffectParams")}
	})
	return MTROnOffClusterOffWithEffectParamsClass
}

type _MTROnOffClusterOffWithEffectParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROnOffClusterOffWithEffectParams */
// An interface definition for the [MTROnOffClusterOffWithEffectParams] class.
type IMTROnOffClusterOffWithEffectParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROnOffClusterOffWithEffectParams */
	// properties:
	EffectId() objc.IObject /* cross-framework: NSNumber */
	SetEffectId(value objc.IObject /* cross-framework: NSNumber */)
	EffectIdentifier() objc.IObject /* cross-framework: NSNumber */
	SetEffectIdentifier(value objc.IObject /* cross-framework: NSNumber */)
	EffectVariant() objc.IObject /* cross-framework: NSNumber */
	SetEffectVariant(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROnOffClusterOffWithEffectParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROnOffClusterOffWithEffectParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOffWithEffectParamsClass) Alloc() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROnOffClusterOffWithEffectParamsClass) New() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOffWithEffectParams) Init() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOffWithEffectParams) Autorelease() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOffWithEffectParams creates a new MTROnOffClusterOffWithEffectParams instance.
func NewMTROnOffClusterOffWithEffectParams() MTROnOffClusterOffWithEffectParams {
	return getMTROnOffClusterOffWithEffectParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROnOffClusterOffWithEffectParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams
type MTROnOffClusterOffWithEffectParams struct {
	objectivec.Object
}

// MTROnOffClusterOffWithEffectParamsFrom constructs a [MTROnOffClusterOffWithEffectParams] from an unsafe.Pointer.
func MTROnOffClusterOffWithEffectParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOffWithEffectParams {
	return MTROnOffClusterOffWithEffectParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROnOffClusterOffWithEffectParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROnOffClusterOffWithEffectParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROnOffClusterOffWithEffectParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROnOffClusterOffWithEffectParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROnOffClusterOffWithEffectParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/effectId
func (m_ MTROnOffClusterOffWithEffectParams) EffectId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("effectId"))
	return rv
}/* debug [instance_properties/getter]: effectId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/effectId
func (m_ MTROnOffClusterOffWithEffectParams) SetEffectId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectId:"), value)
}/* debug [instance_properties/setter]: effectId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/effectIdentifier
func (m_ MTROnOffClusterOffWithEffectParams) EffectIdentifier() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("effectIdentifier"))
	return rv
}/* debug [instance_properties/getter]: effectIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/effectIdentifier
func (m_ MTROnOffClusterOffWithEffectParams) SetEffectIdentifier(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectIdentifier:"), value)
}/* debug [instance_properties/setter]: effectIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/effectVariant
func (m_ MTROnOffClusterOffWithEffectParams) EffectVariant() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("effectVariant"))
	return rv
}/* debug [instance_properties/getter]: effectVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/effectVariant
func (m_ MTROnOffClusterOffWithEffectParams) SetEffectVariant(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectVariant:"), value)
}/* debug [instance_properties/setter]: effectVariant */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOffWithEffectParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/serverSideProcessingTimeout
func (m_ MTROnOffClusterOffWithEffectParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOffWithEffectParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams/timedInvokeTimeoutMs
func (m_ MTROnOffClusterOffWithEffectParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROnOffClusterOffWithEffectParams */



