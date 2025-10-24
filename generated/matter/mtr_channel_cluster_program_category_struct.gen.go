// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterProgramCategoryStruct */


/* debug [class_header]: Header for MTRChannelClusterProgramCategoryStruct */
// The class instance for the [MTRChannelClusterProgramCategoryStruct] class.
var (
	MTRChannelClusterProgramCategoryStructClass     _MTRChannelClusterProgramCategoryStructClass
	MTRChannelClusterProgramCategoryStructClassOnce sync.Once
)

func getMTRChannelClusterProgramCategoryStructClass() _MTRChannelClusterProgramCategoryStructClass {
	MTRChannelClusterProgramCategoryStructClassOnce.Do(func() {
		MTRChannelClusterProgramCategoryStructClass = _MTRChannelClusterProgramCategoryStructClass{objc.GetClass("MTRChannelClusterProgramCategoryStruct")}
	})
	return MTRChannelClusterProgramCategoryStructClass
}

type _MTRChannelClusterProgramCategoryStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterProgramCategoryStruct */
// An interface definition for the [MTRChannelClusterProgramCategoryStruct] class.
type IMTRChannelClusterProgramCategoryStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterProgramCategoryStruct */
	// properties:
	SubCategory() objc.IObject /* cross-framework: NSString */
	SetSubCategory(value objc.IObject /* cross-framework: NSString */)
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterProgramCategoryStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterProgramCategoryStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramCategoryStructClass) Alloc() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterProgramCategoryStructClass) New() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramCategoryStruct) Init() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramCategoryStruct) Autorelease() MTRChannelClusterProgramCategoryStruct {
	rv := objc.Send[MTRChannelClusterProgramCategoryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramCategoryStruct creates a new MTRChannelClusterProgramCategoryStruct instance.
func NewMTRChannelClusterProgramCategoryStruct() MTRChannelClusterProgramCategoryStruct {
	return getMTRChannelClusterProgramCategoryStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterProgramCategoryStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct
type MTRChannelClusterProgramCategoryStruct struct {
	objectivec.Object
}

// MTRChannelClusterProgramCategoryStructFrom constructs a [MTRChannelClusterProgramCategoryStruct] from an unsafe.Pointer.
func MTRChannelClusterProgramCategoryStructFrom(ptr unsafe.Pointer) MTRChannelClusterProgramCategoryStruct {
	return MTRChannelClusterProgramCategoryStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterProgramCategoryStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterProgramCategoryStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterProgramCategoryStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterProgramCategoryStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterProgramCategoryStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/subCategory
func (m_ MTRChannelClusterProgramCategoryStruct) SubCategory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subCategory"))
	return rv
}/* debug [instance_properties/getter]: subCategory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCategoryStruct/subCategory
func (m_ MTRChannelClusterProgramCategoryStruct) SetSubCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubCategory:"), value)
}/* debug [instance_properties/setter]: subCategory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramcategorystruct/category
func (m_ MTRChannelClusterProgramCategoryStruct) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramcategorystruct/category
func (m_ MTRChannelClusterProgramCategoryStruct) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterProgramCategoryStruct */



