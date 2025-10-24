// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWriteParams */


/* debug [class_header]: Header for MTRWriteParams */
// The class instance for the [MTRWriteParams] class.
var (
	MTRWriteParamsClass     _MTRWriteParamsClass
	MTRWriteParamsClassOnce sync.Once
)

func getMTRWriteParamsClass() _MTRWriteParamsClass {
	MTRWriteParamsClassOnce.Do(func() {
		MTRWriteParamsClass = _MTRWriteParamsClass{objc.GetClass("MTRWriteParams")}
	})
	return MTRWriteParamsClass
}

type _MTRWriteParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWriteParams */
// An interface definition for the [MTRWriteParams] class.
type IMTRWriteParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWriteParams */
	// properties:
	DataVersion() objc.IObject /* cross-framework: NSNumber */
	SetDataVersion(value objc.IObject /* cross-framework: NSNumber */)
	TimedWriteTimeout() objc.IObject /* cross-framework: NSNumber */
	SetTimedWriteTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWriteParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWriteParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWriteParamsClass) Alloc() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWriteParamsClass) New() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWriteParams) Init() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWriteParams) Autorelease() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWriteParams creates a new MTRWriteParams instance.
func NewMTRWriteParams() MTRWriteParams {
	return getMTRWriteParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWriteParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWriteParams
type MTRWriteParams struct {
	objectivec.Object
}

// MTRWriteParamsFrom constructs a [MTRWriteParams] from an unsafe.Pointer.
func MTRWriteParamsFrom(ptr unsafe.Pointer) MTRWriteParams {
	return MTRWriteParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWriteParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWriteParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWriteParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWriteParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWriteParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWriteParams/dataVersion
func (m_ MTRWriteParams) DataVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dataVersion"))
	return rv
}/* debug [instance_properties/getter]: dataVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWriteParams/dataVersion
func (m_ MTRWriteParams) SetDataVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataVersion:"), value)
}/* debug [instance_properties/setter]: dataVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWriteParams/timedWriteTimeout
func (m_ MTRWriteParams) TimedWriteTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedWriteTimeout"))
	return rv
}/* debug [instance_properties/getter]: timedWriteTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWriteParams/timedWriteTimeout
func (m_ MTRWriteParams) SetTimedWriteTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedWriteTimeout:"), value)
}/* debug [instance_properties/setter]: timedWriteTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWriteParams */





