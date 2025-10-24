// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXCategoricalDataAxisDescriptor */


/* debug [class_header]: Header for AXCategoricalDataAxisDescriptor */
// The class instance for the [AXCategoricalDataAxisDescriptor] class.
var (
	AXCategoricalDataAxisDescriptorClass     _AXCategoricalDataAxisDescriptorClass
	AXCategoricalDataAxisDescriptorClassOnce sync.Once
)

func getAXCategoricalDataAxisDescriptorClass() _AXCategoricalDataAxisDescriptorClass {
	AXCategoricalDataAxisDescriptorClassOnce.Do(func() {
		AXCategoricalDataAxisDescriptorClass = _AXCategoricalDataAxisDescriptorClass{objc.GetClass("AXCategoricalDataAxisDescriptor")}
	})
	return AXCategoricalDataAxisDescriptorClass
}

type _AXCategoricalDataAxisDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXCategoricalDataAxisDescriptor */
// An interface definition for the [AXCategoricalDataAxisDescriptor] class.
type IAXCategoricalDataAxisDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXCategoricalDataAxisDescriptor */
	// properties:
	CategoryOrder() []string
	SetCategoryOrder(value []string)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXCategoricalDataAxisDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXCategoricalDataAxisDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AXCategoricalDataAxisDescriptorClass) Alloc() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXCategoricalDataAxisDescriptorClass) New() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXCategoricalDataAxisDescriptor) Init() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXCategoricalDataAxisDescriptor) Autorelease() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXCategoricalDataAxisDescriptor creates a new AXCategoricalDataAxisDescriptor instance.
func NewAXCategoricalDataAxisDescriptor() AXCategoricalDataAxisDescriptor {
	return getAXCategoricalDataAxisDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXCategoricalDataAxisDescriptor */
// An object that represents an axis of categorical data.
//
// A categorical data axis divides information into groups, or categories. For example, a categorical axis may represent blood type data divided into the possible categories , , , and .


// An object that represents an axis of categorical data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor
type AXCategoricalDataAxisDescriptor struct {
	objectivec.Object
}

// AXCategoricalDataAxisDescriptorFrom constructs a [AXCategoricalDataAxisDescriptor] from an unsafe.Pointer.
//
// An object that represents an axis of categorical data.
func AXCategoricalDataAxisDescriptorFrom(ptr unsafe.Pointer) AXCategoricalDataAxisDescriptor {
	return AXCategoricalDataAxisDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXCategoricalDataAxisDescriptor */

// Creates a categorical data axis with the specified attributed title and an array of categories in the specified order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/init(attributedTitle:categoryOrder:)
func NewAXCategoricalDataAxisDescriptorWithAttributedTitleCategoryOrder(attributedTitle foundation.AttributedString, categoryOrder []string) AXCategoricalDataAxisDescriptor {
	instance := getAXCategoricalDataAxisDescriptorClass().Alloc()
	rv := objc.Send[AXCategoricalDataAxisDescriptor](instance.ID, objc.Sel("initWithAttributedTitle:categoryOrder:"), attributedTitle, categoryOrder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXCategoricalDataAxisDescriptorWithAttributedTitleCategoryOrder */


// Creates a categorical data axis with the specified title and an array of categories in the specified order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/init(title:categoryOrder:)
func NewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder(title objc.IObject /* cross-framework: NSString */, categoryOrder []string) AXCategoricalDataAxisDescriptor {
	instance := getAXCategoricalDataAxisDescriptorClass().Alloc()
	rv := objc.Send[AXCategoricalDataAxisDescriptor](instance.ID, objc.Sel("initWithTitle:categoryOrder:"), title, categoryOrder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXCategoricalDataAxisDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXCategoricalDataAxisDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXCategoricalDataAxisDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXCategoricalDataAxisDescriptor */

// A list of every category value for the axis in the order they appear visually in the graph or legend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/categoryOrder
func (a_ AXCategoricalDataAxisDescriptor) CategoryOrder() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("categoryOrder"))
	return rv
}/* debug [instance_properties/getter]: categoryOrder */


// A list of every category value for the axis in the order they appear visually in the graph or legend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/categoryOrder
func (a_ AXCategoricalDataAxisDescriptor) SetCategoryOrder(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setCategoryOrder:"), nsArray)
}/* debug [instance_properties/setter]: categoryOrder */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXCategoricalDataAxisDescriptor */


