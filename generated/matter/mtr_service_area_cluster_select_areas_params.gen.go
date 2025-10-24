// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterSelectAreasParams */


/* debug [class_header]: Header for MTRServiceAreaClusterSelectAreasParams */
// The class instance for the [MTRServiceAreaClusterSelectAreasParams] class.
var (
	MTRServiceAreaClusterSelectAreasParamsClass     _MTRServiceAreaClusterSelectAreasParamsClass
	MTRServiceAreaClusterSelectAreasParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSelectAreasParamsClass() _MTRServiceAreaClusterSelectAreasParamsClass {
	MTRServiceAreaClusterSelectAreasParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSelectAreasParamsClass = _MTRServiceAreaClusterSelectAreasParamsClass{objc.GetClass("MTRServiceAreaClusterSelectAreasParams")}
	})
	return MTRServiceAreaClusterSelectAreasParamsClass
}

type _MTRServiceAreaClusterSelectAreasParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterSelectAreasParams */
// An interface definition for the [MTRServiceAreaClusterSelectAreasParams] class.
type IMTRServiceAreaClusterSelectAreasParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterSelectAreasParams */
	// properties:
	NewAreas() objc.IObject /* cross-framework: NSArray */
	SetNewAreas(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterSelectAreasParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterSelectAreasParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSelectAreasParamsClass) Alloc() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterSelectAreasParamsClass) New() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSelectAreasParams) Init() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSelectAreasParams) Autorelease() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSelectAreasParams creates a new MTRServiceAreaClusterSelectAreasParams instance.
func NewMTRServiceAreaClusterSelectAreasParams() MTRServiceAreaClusterSelectAreasParams {
	return getMTRServiceAreaClusterSelectAreasParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterSelectAreasParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams
type MTRServiceAreaClusterSelectAreasParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSelectAreasParamsFrom constructs a [MTRServiceAreaClusterSelectAreasParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSelectAreasParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSelectAreasParams {
	return MTRServiceAreaClusterSelectAreasParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterSelectAreasParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterSelectAreasParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterSelectAreasParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterSelectAreasParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterSelectAreasParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/newAreas
func (m_ MTRServiceAreaClusterSelectAreasParams) NewAreas() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("newAreas"))
	return rv
}/* debug [instance_properties/getter]: newAreas */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/newAreas
func (m_ MTRServiceAreaClusterSelectAreasParams) SetNewAreas(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewAreas:"), value)
}/* debug [instance_properties/setter]: newAreas */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterselectareasparams/serversideprocessingtimeout
func (m_ MTRServiceAreaClusterSelectAreasParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterselectareasparams/serversideprocessingtimeout
func (m_ MTRServiceAreaClusterSelectAreasParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterselectareasparams/timedinvoketimeoutms
func (m_ MTRServiceAreaClusterSelectAreasParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterselectareasparams/timedinvoketimeoutms
func (m_ MTRServiceAreaClusterSelectAreasParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterSelectAreasParams */



