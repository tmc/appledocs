// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRContentAppObserverClusterContentAppMessageResponseParams */


/* debug [class_header]: Header for MTRContentAppObserverClusterContentAppMessageResponseParams */
// The class instance for the [MTRContentAppObserverClusterContentAppMessageResponseParams] class.
var (
	MTRContentAppObserverClusterContentAppMessageResponseParamsClass     _MTRContentAppObserverClusterContentAppMessageResponseParamsClass
	MTRContentAppObserverClusterContentAppMessageResponseParamsClassOnce sync.Once
)

func getMTRContentAppObserverClusterContentAppMessageResponseParamsClass() _MTRContentAppObserverClusterContentAppMessageResponseParamsClass {
	MTRContentAppObserverClusterContentAppMessageResponseParamsClassOnce.Do(func() {
		MTRContentAppObserverClusterContentAppMessageResponseParamsClass = _MTRContentAppObserverClusterContentAppMessageResponseParamsClass{objc.GetClass("MTRContentAppObserverClusterContentAppMessageResponseParams")}
	})
	return MTRContentAppObserverClusterContentAppMessageResponseParamsClass
}

type _MTRContentAppObserverClusterContentAppMessageResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentAppObserverClusterContentAppMessageResponseParams */
// An interface definition for the [MTRContentAppObserverClusterContentAppMessageResponseParams] class.
type IMTRContentAppObserverClusterContentAppMessageResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRContentAppObserverClusterContentAppMessageResponseParams */
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	EncodingHint() objc.IObject /* cross-framework: NSString */
	SetEncodingHint(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentAppObserverClusterContentAppMessageResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentAppObserverClusterContentAppMessageResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentAppObserverClusterContentAppMessageResponseParamsClass) Alloc() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentAppObserverClusterContentAppMessageResponseParamsClass) New() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Init() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Autorelease() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentAppObserverClusterContentAppMessageResponseParams creates a new MTRContentAppObserverClusterContentAppMessageResponseParams instance.
func NewMTRContentAppObserverClusterContentAppMessageResponseParams() MTRContentAppObserverClusterContentAppMessageResponseParams {
	return getMTRContentAppObserverClusterContentAppMessageResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentAppObserverClusterContentAppMessageResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams
type MTRContentAppObserverClusterContentAppMessageResponseParams struct {
	objectivec.Object
}

// MTRContentAppObserverClusterContentAppMessageResponseParamsFrom constructs a [MTRContentAppObserverClusterContentAppMessageResponseParams] from an unsafe.Pointer.
func MTRContentAppObserverClusterContentAppMessageResponseParamsFrom(ptr unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageResponseParams {
	return MTRContentAppObserverClusterContentAppMessageResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentAppObserverClusterContentAppMessageResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentAppObserverClusterContentAppMessageResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentAppObserverClusterContentAppMessageResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentAppObserverClusterContentAppMessageResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentAppObserverClusterContentAppMessageResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageresponseparams/encodinghint
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) EncodingHint() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("encodingHint"))
	return rv
}/* debug [instance_properties/getter]: encodingHint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageresponseparams/encodinghint
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetEncodingHint(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodingHint:"), value)
}/* debug [instance_properties/setter]: encodingHint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageresponseparams/status
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentappobserverclustercontentappmessageresponseparams/status
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentAppObserverClusterContentAppMessageResponseParams */



