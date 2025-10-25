// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMetadataGroup */


/* debug [class_header]: Header for AVMetadataGroup */
// The class instance for the [MetadataGroup] class.
var (
	MetadataGroupClass     _MetadataGroupClass
	MetadataGroupClassOnce sync.Once
)

func getMetadataGroupClass() _MetadataGroupClass {
	MetadataGroupClassOnce.Do(func() {
		MetadataGroupClass = _MetadataGroupClass{objc.GetClass("AVMetadataGroup")}
	})
	return MetadataGroupClass
}

type _MetadataGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataGroup */
// An interface definition for the [MetadataGroup] class.
type IMetadataGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MetadataGroup */
	// properties:
	ClassifyingLabel() objc.IObject /* cross-framework: NSString */
	Items() []MetadataItem
	UniqueID() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataGroup */
// Alloc allocates a new instance without initialization.
func (mc _MetadataGroupClass) Alloc() MetadataGroup {
	rv := objc.Send[MetadataGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataGroupClass) New() MetadataGroup {
	rv := objc.Send[MetadataGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataGroup) Init() MetadataGroup {
	rv := objc.Send[MetadataGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataGroup) Autorelease() MetadataGroup {
	rv := objc.Send[MetadataGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataGroup creates a new MetadataGroup instance.
func NewMetadataGroup() MetadataGroup {
	return getMetadataGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataGroup */
// A collection of metadata items associated with a timeline segment.


// A collection of metadata items associated with a timeline segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup
type MetadataGroup struct {
	objectivec.Object
}

// MetadataGroupFrom constructs a [MetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items associated with a timeline segment.
func MetadataGroupFrom(ptr unsafe.Pointer) MetadataGroup {
	return MetadataGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataGroup */

// The classifying label associated with the metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/classifyingLabel
func (m_ MetadataGroup) ClassifyingLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("classifyingLabel"))
	return rv
}/* debug [instance_properties/getter]: classifyingLabel */


// The array of metadata items associated with the metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/items
func (m_ MetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// The unique identifier for the metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/uniqueID
func (m_ MetadataGroup) UniqueID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("uniqueID"))
	return rv
}/* debug [instance_properties/getter]: uniqueID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataGroup */



