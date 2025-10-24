// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCounter */

/* debug [class_header]: Header for DOMCounter */
// The class instance for the [DOMCounter] class.
var (
	DOMCounterClass     _DOMCounterClass
	DOMCounterClassOnce sync.Once
)

func getDOMCounterClass() _DOMCounterClass {
	DOMCounterClassOnce.Do(func() {
		DOMCounterClass = _DOMCounterClass{objc.GetClass("DOMCounter")}
	})
	return DOMCounterClass
}

type _DOMCounterClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCounter */
// An interface definition for the [DOMCounter] class.
type IDOMCounter interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMCounter */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	ListStyle() objc.IObject  /* cross-framework: NSString */
	Separator() objc.IObject  /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCounter */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCounter */
// Alloc allocates a new instance without initialization.
func (dc _DOMCounterClass) Alloc() DOMCounter {
	rv := objc.Send[DOMCounter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCounterClass) New() DOMCounter {
	rv := objc.Send[DOMCounter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCounter) Init() DOMCounter {
	rv := objc.Send[DOMCounter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCounter) Autorelease() DOMCounter {
	rv := objc.Send[DOMCounter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCounter creates a new DOMCounter instance.
func NewDOMCounter() DOMCounter {
	return getDOMCounterClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCounter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCounter
type DOMCounter struct {
	DOMObject
}

// DOMCounterFrom constructs a [DOMCounter] from an unsafe.Pointer.
func DOMCounterFrom(ptr unsafe.Pointer) DOMCounter {
	return DOMCounter{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCounter */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCounter */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCounter */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCounter */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCounter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCounter/identifier
func (d_ DOMCounter) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCounter/listStyle
func (d_ DOMCounter) ListStyle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("listStyle"))
	return rv
} /* debug [instance_properties/getter]: listStyle */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCounter/separator
func (d_ DOMCounter) Separator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("separator"))
	return rv
} /* debug [instance_properties/getter]: separator */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCounter */
