// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterSeriesInfoStruct */


/* debug [class_header]: Header for MTRChannelClusterSeriesInfoStruct */
// The class instance for the [MTRChannelClusterSeriesInfoStruct] class.
var (
	MTRChannelClusterSeriesInfoStructClass     _MTRChannelClusterSeriesInfoStructClass
	MTRChannelClusterSeriesInfoStructClassOnce sync.Once
)

func getMTRChannelClusterSeriesInfoStructClass() _MTRChannelClusterSeriesInfoStructClass {
	MTRChannelClusterSeriesInfoStructClassOnce.Do(func() {
		MTRChannelClusterSeriesInfoStructClass = _MTRChannelClusterSeriesInfoStructClass{objc.GetClass("MTRChannelClusterSeriesInfoStruct")}
	})
	return MTRChannelClusterSeriesInfoStructClass
}

type _MTRChannelClusterSeriesInfoStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterSeriesInfoStruct */
// An interface definition for the [MTRChannelClusterSeriesInfoStruct] class.
type IMTRChannelClusterSeriesInfoStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterSeriesInfoStruct */
	// properties:
	Episode() objc.IObject /* cross-framework: NSString */
	SetEpisode(value objc.IObject /* cross-framework: NSString */)
	Season() objc.IObject /* cross-framework: NSString */
	SetSeason(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterSeriesInfoStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterSeriesInfoStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterSeriesInfoStructClass) Alloc() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterSeriesInfoStructClass) New() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterSeriesInfoStruct) Init() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterSeriesInfoStruct) Autorelease() MTRChannelClusterSeriesInfoStruct {
	rv := objc.Send[MTRChannelClusterSeriesInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterSeriesInfoStruct creates a new MTRChannelClusterSeriesInfoStruct instance.
func NewMTRChannelClusterSeriesInfoStruct() MTRChannelClusterSeriesInfoStruct {
	return getMTRChannelClusterSeriesInfoStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterSeriesInfoStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct
type MTRChannelClusterSeriesInfoStruct struct {
	objectivec.Object
}

// MTRChannelClusterSeriesInfoStructFrom constructs a [MTRChannelClusterSeriesInfoStruct] from an unsafe.Pointer.
func MTRChannelClusterSeriesInfoStructFrom(ptr unsafe.Pointer) MTRChannelClusterSeriesInfoStruct {
	return MTRChannelClusterSeriesInfoStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterSeriesInfoStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterSeriesInfoStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterSeriesInfoStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterSeriesInfoStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterSeriesInfoStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/episode
func (m_ MTRChannelClusterSeriesInfoStruct) Episode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("episode"))
	return rv
}/* debug [instance_properties/getter]: episode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSeriesInfoStruct/episode
func (m_ MTRChannelClusterSeriesInfoStruct) SetEpisode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpisode:"), value)
}/* debug [instance_properties/setter]: episode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterseriesinfostruct/season
func (m_ MTRChannelClusterSeriesInfoStruct) Season() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("season"))
	return rv
}/* debug [instance_properties/getter]: season */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterseriesinfostruct/season
func (m_ MTRChannelClusterSeriesInfoStruct) SetSeason(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeason:"), value)
}/* debug [instance_properties/setter]: season */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterSeriesInfoStruct */



