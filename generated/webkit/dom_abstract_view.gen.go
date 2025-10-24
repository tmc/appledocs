// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMAbstractView */


/* debug [class_header]: Header for DOMAbstractView */
// The class instance for the [DOMAbstractView] class.
var (
	DOMAbstractViewClass     _DOMAbstractViewClass
	DOMAbstractViewClassOnce sync.Once
)

func getDOMAbstractViewClass() _DOMAbstractViewClass {
	DOMAbstractViewClassOnce.Do(func() {
		DOMAbstractViewClass = _DOMAbstractViewClass{objc.GetClass("DOMAbstractView")}
	})
	return DOMAbstractViewClass
}

type _DOMAbstractViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMAbstractView */
// An interface definition for the [DOMAbstractView] class.
type IDOMAbstractView interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMAbstractView */
	// properties:
	Document() IDOMDocument
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMAbstractView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMAbstractView */
// Alloc allocates a new instance without initialization.
func (dc _DOMAbstractViewClass) Alloc() DOMAbstractView {
	rv := objc.Send[DOMAbstractView](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMAbstractViewClass) New() DOMAbstractView {
	rv := objc.Send[DOMAbstractView](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMAbstractView) Init() DOMAbstractView {
	rv := objc.Send[DOMAbstractView](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMAbstractView) Autorelease() DOMAbstractView {
	rv := objc.Send[DOMAbstractView](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMAbstractView creates a new DOMAbstractView instance.
func NewDOMAbstractView() DOMAbstractView {
	return getDOMAbstractViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMAbstractView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAbstractView
type DOMAbstractView struct {
	DOMObject
}

// DOMAbstractViewFrom constructs a [DOMAbstractView] from an unsafe.Pointer.
func DOMAbstractViewFrom(ptr unsafe.Pointer) DOMAbstractView {
	return DOMAbstractView{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMAbstractView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMAbstractView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMAbstractView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMAbstractView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMAbstractView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAbstractView/document
func (d_ DOMAbstractView) Document() IDOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMAbstractView */



