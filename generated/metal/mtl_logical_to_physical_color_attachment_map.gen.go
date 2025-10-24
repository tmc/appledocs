// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLLogicalToPhysicalColorAttachmentMap */


/* debug [class_header]: Header for MTLLogicalToPhysicalColorAttachmentMap */
// The class instance for the [LogicalToPhysicalColorAttachmentMap] class.
var (
	LogicalToPhysicalColorAttachmentMapClass     _LogicalToPhysicalColorAttachmentMapClass
	LogicalToPhysicalColorAttachmentMapClassOnce sync.Once
)

func getLogicalToPhysicalColorAttachmentMapClass() _LogicalToPhysicalColorAttachmentMapClass {
	LogicalToPhysicalColorAttachmentMapClassOnce.Do(func() {
		LogicalToPhysicalColorAttachmentMapClass = _LogicalToPhysicalColorAttachmentMapClass{objc.GetClass("MTLLogicalToPhysicalColorAttachmentMap")}
	})
	return LogicalToPhysicalColorAttachmentMapClass
}

type _LogicalToPhysicalColorAttachmentMapClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LogicalToPhysicalColorAttachmentMap */
// An interface definition for the [LogicalToPhysicalColorAttachmentMap] class.
type ILogicalToPhysicalColorAttachmentMap interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LogicalToPhysicalColorAttachmentMap */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LogicalToPhysicalColorAttachmentMap */
	// methods:
	GetPhysicalIndexForLogicalIndex(logicalIndex uint) uint
	Reset()
	SetPhysicalIndexForLogicalIndex(physicalIndex uint, logicalIndex uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LogicalToPhysicalColorAttachmentMap */
// Alloc allocates a new instance without initialization.
func (lc _LogicalToPhysicalColorAttachmentMapClass) Alloc() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LogicalToPhysicalColorAttachmentMapClass) New() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogicalToPhysicalColorAttachmentMap) Init() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogicalToPhysicalColorAttachmentMap) Autorelease() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogicalToPhysicalColorAttachmentMap creates a new LogicalToPhysicalColorAttachmentMap instance.
func NewLogicalToPhysicalColorAttachmentMap() LogicalToPhysicalColorAttachmentMap {
	return getLogicalToPhysicalColorAttachmentMapClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LogicalToPhysicalColorAttachmentMap */
// Allows you to easily specify color attachment remapping from logical to physical indices.


// Allows you to easily specify color attachment remapping from logical to physical indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap
type LogicalToPhysicalColorAttachmentMap struct {
	objectivec.Object
}

// LogicalToPhysicalColorAttachmentMapFrom constructs a [LogicalToPhysicalColorAttachmentMap] from an unsafe.Pointer.
//
// Allows you to easily specify color attachment remapping from logical to physical indices.
func LogicalToPhysicalColorAttachmentMapFrom(ptr unsafe.Pointer) LogicalToPhysicalColorAttachmentMap {
	return LogicalToPhysicalColorAttachmentMap{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LogicalToPhysicalColorAttachmentMap *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LogicalToPhysicalColorAttachmentMap */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LogicalToPhysicalColorAttachmentMap */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LogicalToPhysicalColorAttachmentMap */

// Queries the physical color attachment index corresponding to a logical index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/getPhysicalIndexForLogicalIndex:
func (l_ LogicalToPhysicalColorAttachmentMap) GetPhysicalIndexForLogicalIndex(logicalIndex uint) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("getPhysicalIndexForLogicalIndex:"), logicalIndex)
	return rv
}/* debug [instance_methods/method]: GetPhysicalIndexForLogicalIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/reset()
func (l_ LogicalToPhysicalColorAttachmentMap) Reset() {
	objc.Send[objc.ID](l_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Maps a physical color attachment index to a logical index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/setPhysicalIndex:forLogicalIndex:
func (l_ LogicalToPhysicalColorAttachmentMap) SetPhysicalIndexForLogicalIndex(physicalIndex uint, logicalIndex uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPhysicalIndex:forLogicalIndex:"), physicalIndex, logicalIndex)
}/* debug [instance_methods/method]: SetPhysicalIndexForLogicalIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LogicalToPhysicalColorAttachmentMap */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLLogicalToPhysicalColorAttachmentMap */



