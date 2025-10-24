// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMCSSValueList */


/* debug [class_header]: Header for DOMCSSValueList */
// The class instance for the [DOMCSSValueList] class.
var (
	DOMCSSValueListClass     _DOMCSSValueListClass
	DOMCSSValueListClassOnce sync.Once
)

func getDOMCSSValueListClass() _DOMCSSValueListClass {
	DOMCSSValueListClassOnce.Do(func() {
		DOMCSSValueListClass = _DOMCSSValueListClass{objc.GetClass("DOMCSSValueList")}
	})
	return DOMCSSValueListClass
}

type _DOMCSSValueListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMCSSValueList */
// An interface definition for the [DOMCSSValueList] class.
type IDOMCSSValueList interface {
	IDOMCSSValue
	
/* debug [class_interface_properties]: Properties for DOMCSSValueList */
	// properties:
	Length() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMCSSValueList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMCSSValueList */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSValueListClass) Alloc() DOMCSSValueList {
	rv := objc.Send[DOMCSSValueList](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSValueListClass) New() DOMCSSValueList {
	rv := objc.Send[DOMCSSValueList](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSValueList) Init() DOMCSSValueList {
	rv := objc.Send[DOMCSSValueList](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSValueList) Autorelease() DOMCSSValueList {
	rv := objc.Send[DOMCSSValueList](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSValueList creates a new DOMCSSValueList instance.
func NewDOMCSSValueList() DOMCSSValueList {
	return getDOMCSSValueListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMCSSValueList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSValueList
type DOMCSSValueList struct {
	DOMCSSValue
}

// DOMCSSValueListFrom constructs a [DOMCSSValueList] from an unsafe.Pointer.
func DOMCSSValueListFrom(ptr unsafe.Pointer) DOMCSSValueList {
	return DOMCSSValueList{
		DOMCSSValue: DOMCSSValueFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMCSSValueList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMCSSValueList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMCSSValueList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMCSSValueList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMCSSValueList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSValueList/length
func (d_ DOMCSSValueList) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMCSSValueList */



