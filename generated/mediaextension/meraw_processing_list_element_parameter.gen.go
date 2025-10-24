// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MERAWProcessingListElementParameter */


/* debug [class_header]: Header for MERAWProcessingListElementParameter */
// The class instance for the [MERAWProcessingListElementParameter] class.
var (
	MERAWProcessingListElementParameterClass     _MERAWProcessingListElementParameterClass
	MERAWProcessingListElementParameterClassOnce sync.Once
)

func getMERAWProcessingListElementParameterClass() _MERAWProcessingListElementParameterClass {
	MERAWProcessingListElementParameterClassOnce.Do(func() {
		MERAWProcessingListElementParameterClass = _MERAWProcessingListElementParameterClass{objc.GetClass("MERAWProcessingListElementParameter")}
	})
	return MERAWProcessingListElementParameterClass
}

type _MERAWProcessingListElementParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingListElementParameter */
// An interface definition for the [MERAWProcessingListElementParameter] class.
type IMERAWProcessingListElementParameter interface {
	IMERAWProcessingParameter
	
/* debug [class_interface_properties]: Properties for MERAWProcessingListElementParameter */
	// properties:
	ListElementID() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingListElementParameter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingListElementParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingListElementParameterClass) Alloc() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingListElementParameterClass) New() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingListElementParameter) Init() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingListElementParameter) Autorelease() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingListElementParameter creates a new MERAWProcessingListElementParameter instance.
func NewMERAWProcessingListElementParameter() MERAWProcessingListElementParameter {
	return getMERAWProcessingListElementParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingListElementParameter */
// An object that describes a list element parameter of a RAW processor.
//
// The protocol provides an interface for to query descriptions of the different elements in a parameter list for a list element in a . A distinct is created for each list element.


// An object that describes a list element parameter of a RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement
type MERAWProcessingListElementParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingListElementParameterFrom constructs a [MERAWProcessingListElementParameter] from an unsafe.Pointer.
//
// An object that describes a list element parameter of a RAW processor.
func MERAWProcessingListElementParameterFrom(ptr unsafe.Pointer) MERAWProcessingListElementParameter {
	return MERAWProcessingListElementParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingListElementParameter */

// Creates a list element parameter object with the element id value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement/init(name:description:elementID:)
func NewMERAWProcessingListElementParameterWithNameDescriptionElementID(name objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, elementID int) MERAWProcessingListElementParameter {
	instance := getMERAWProcessingListElementParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingListElementParameter](instance.ID, objc.Sel("initWithName:description:elementID:"), name, description, elementID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingListElementParameterWithNameDescriptionElementID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingListElementParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingListElementParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingListElementParameter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingListElementParameter */

// A unique number in the list which represents this list option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement/listElementID
func (m_ MERAWProcessingListElementParameter) ListElementID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("listElementID"))
	return rv
}/* debug [instance_properties/getter]: listElementID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingListElementParameter */


