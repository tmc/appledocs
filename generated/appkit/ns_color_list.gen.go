// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSColorList */


/* debug [class_header]: Header for NSColorList */
// The class instance for the [ColorList] class.
var (
	ColorListClass     _ColorListClass
	ColorListClassOnce sync.Once
)

func getColorListClass() _ColorListClass {
	ColorListClassOnce.Do(func() {
		ColorListClass = _ColorListClass{objc.GetClass("NSColorList")}
	})
	return ColorListClass
}

type _ColorListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ColorList */
// An interface definition for the [ColorList] class.
type IColorList interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ColorList */
	// properties:
	AllKeys() objectivec.IObject
	SetAllKeys(value objectivec.IObject)
	IsEditable() bool
	SetIsEditable(value bool)
	Name() objectivec.IObject
	SetName(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ColorList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ColorList */
// Alloc allocates a new instance without initialization.
func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ColorListClass) New() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorList) Init() ColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorList) Autorelease() ColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorList creates a new ColorList instance.
func NewColorList() ColorList {
	return getColorListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ColorList */
// An ordered list of color objects, identified by keys.
//
// A color list manages a list of objects, each of which has an associated name. The list mode color picker uses instances of to represent any lists of colors that come with the system, as well as any lists the user creates. An app can use a color list to manage document-specific color lists.


// An ordered list of color objects, identified by keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList
type ColorList struct {
	objectivec.Object
}

// ColorListFrom constructs a [ColorList] from an unsafe.Pointer.
//
// An ordered list of color objects, identified by keys.
func ColorListFrom(ptr unsafe.Pointer) ColorList {
	return ColorList{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ColorList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ColorList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ColorList */

// Returns an array of all color lists found in the standard color list directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList/availableColorLists
func (cc _ColorListClass) AvailableColorLists() []ColorList {
	rv := objc.Send[[]ColorList](objc.ID(cc.class), objc.Sel("availableColorLists"))
	return rv
}/* debug [class_properties_class/property]: availableColorLists */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ColorList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ColorList */

// Returns an array of all color lists found in the standard color list directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList/availableColorLists
func (c_ ColorList) AvailableColorLists() []ColorList {
	rv := objc.Send[[]ColorList](c_.ID, objc.Sel("availableColorLists"))
	return rv
}/* debug [instance_properties/getter]: availableColorLists */


// An array of the keys by which the color objects are stored in the color list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/allkeys
func (c_ ColorList) AllKeys() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("allKeys"))
	return rv
}/* debug [instance_properties/getter]: allKeys */


// An array of the keys by which the color objects are stored in the color list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/allkeys
func (c_ ColorList) SetAllKeys(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllKeys:"), value)
}/* debug [instance_properties/setter]: allKeys */


// A Boolean value that indicates whether the color list can be modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/iseditable
func (c_ ColorList) IsEditable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean value that indicates whether the color list can be modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/iseditable
func (c_ ColorList) SetIsEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// The name of the color list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/name-swift.property
func (c_ ColorList) Name() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the color list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/name-swift.property
func (c_ ColorList) SetName(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSColorList */



