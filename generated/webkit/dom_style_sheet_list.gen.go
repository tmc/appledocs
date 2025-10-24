// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMStyleSheetList */


/* debug [class_header]: Header for DOMStyleSheetList */
// The class instance for the [DOMStyleSheetList] class.
var (
	DOMStyleSheetListClass     _DOMStyleSheetListClass
	DOMStyleSheetListClassOnce sync.Once
)

func getDOMStyleSheetListClass() _DOMStyleSheetListClass {
	DOMStyleSheetListClassOnce.Do(func() {
		DOMStyleSheetListClass = _DOMStyleSheetListClass{objc.GetClass("DOMStyleSheetList")}
	})
	return DOMStyleSheetListClass
}

type _DOMStyleSheetListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMStyleSheetList */
// An interface definition for the [DOMStyleSheetList] class.
type IDOMStyleSheetList interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMStyleSheetList */
	// properties:
	Length() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMStyleSheetList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMStyleSheetList */
// Alloc allocates a new instance without initialization.
func (dc _DOMStyleSheetListClass) Alloc() DOMStyleSheetList {
	rv := objc.Send[DOMStyleSheetList](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMStyleSheetListClass) New() DOMStyleSheetList {
	rv := objc.Send[DOMStyleSheetList](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMStyleSheetList) Init() DOMStyleSheetList {
	rv := objc.Send[DOMStyleSheetList](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMStyleSheetList) Autorelease() DOMStyleSheetList {
	rv := objc.Send[DOMStyleSheetList](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMStyleSheetList creates a new DOMStyleSheetList instance.
func NewDOMStyleSheetList() DOMStyleSheetList {
	return getDOMStyleSheetListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMStyleSheetList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheetList
type DOMStyleSheetList struct {
	DOMObject
}

// DOMStyleSheetListFrom constructs a [DOMStyleSheetList] from an unsafe.Pointer.
func DOMStyleSheetListFrom(ptr unsafe.Pointer) DOMStyleSheetList {
	return DOMStyleSheetList{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMStyleSheetList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMStyleSheetList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMStyleSheetList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMStyleSheetList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMStyleSheetList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheetList/length
func (d_ DOMStyleSheetList) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMStyleSheetList */



