// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterProgramGuideResponseParams */


/* debug [class_header]: Header for MTRChannelClusterProgramGuideResponseParams */
// The class instance for the [MTRChannelClusterProgramGuideResponseParams] class.
var (
	MTRChannelClusterProgramGuideResponseParamsClass     _MTRChannelClusterProgramGuideResponseParamsClass
	MTRChannelClusterProgramGuideResponseParamsClassOnce sync.Once
)

func getMTRChannelClusterProgramGuideResponseParamsClass() _MTRChannelClusterProgramGuideResponseParamsClass {
	MTRChannelClusterProgramGuideResponseParamsClassOnce.Do(func() {
		MTRChannelClusterProgramGuideResponseParamsClass = _MTRChannelClusterProgramGuideResponseParamsClass{objc.GetClass("MTRChannelClusterProgramGuideResponseParams")}
	})
	return MTRChannelClusterProgramGuideResponseParamsClass
}

type _MTRChannelClusterProgramGuideResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterProgramGuideResponseParams */
// An interface definition for the [MTRChannelClusterProgramGuideResponseParams] class.
type IMTRChannelClusterProgramGuideResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterProgramGuideResponseParams */
	// properties:
	Paging() IMTRChannelClusterChannelPagingStruct
	SetPaging(value IMTRChannelClusterChannelPagingStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterProgramGuideResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterProgramGuideResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramGuideResponseParamsClass) Alloc() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterProgramGuideResponseParamsClass) New() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramGuideResponseParams) Init() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramGuideResponseParams) Autorelease() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramGuideResponseParams creates a new MTRChannelClusterProgramGuideResponseParams instance.
func NewMTRChannelClusterProgramGuideResponseParams() MTRChannelClusterProgramGuideResponseParams {
	return getMTRChannelClusterProgramGuideResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterProgramGuideResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams
type MTRChannelClusterProgramGuideResponseParams struct {
	objectivec.Object
}

// MTRChannelClusterProgramGuideResponseParamsFrom constructs a [MTRChannelClusterProgramGuideResponseParams] from an unsafe.Pointer.
func MTRChannelClusterProgramGuideResponseParamsFrom(ptr unsafe.Pointer) MTRChannelClusterProgramGuideResponseParams {
	return MTRChannelClusterProgramGuideResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterProgramGuideResponseParams */

// Initialize an MTRChannelClusterProgramGuideResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams/init(responseValue:)
func NewMTRChannelClusterProgramGuideResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRChannelClusterProgramGuideResponseParams {
	instance := getMTRChannelClusterProgramGuideResponseParamsClass().Alloc()
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRChannelClusterProgramGuideResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterProgramGuideResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterProgramGuideResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterProgramGuideResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterProgramGuideResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramguideresponseparams/paging
func (m_ MTRChannelClusterProgramGuideResponseParams) Paging() IMTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](m_.ID, objc.Sel("paging"))
	return rv
}/* debug [instance_properties/getter]: paging */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramguideresponseparams/paging
func (m_ MTRChannelClusterProgramGuideResponseParams) SetPaging(value IMTRChannelClusterChannelPagingStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPaging:"), value)
}/* debug [instance_properties/setter]: paging */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterProgramGuideResponseParams */


