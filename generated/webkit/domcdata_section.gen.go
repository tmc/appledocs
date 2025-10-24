// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCDATASection */

/* debug [class_header]: Header for DOMCDATASection */
// The class instance for the [DOMCDATASection] class.
var (
	DOMCDATASectionClass     _DOMCDATASectionClass
	DOMCDATASectionClassOnce sync.Once
)

func getDOMCDATASectionClass() _DOMCDATASectionClass {
	DOMCDATASectionClassOnce.Do(func() {
		DOMCDATASectionClass = _DOMCDATASectionClass{objc.GetClass("DOMCDATASection")}
	})
	return DOMCDATASectionClass
}

type _DOMCDATASectionClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCDATASection */
// An interface definition for the [DOMCDATASection] class.
type IDOMCDATASection interface {
	IDOMText

	/* debug [class_interface_properties]: Properties for DOMCDATASection */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCDATASection */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCDATASection */
// Alloc allocates a new instance without initialization.
func (dc _DOMCDATASectionClass) Alloc() DOMCDATASection {
	rv := objc.Send[DOMCDATASection](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCDATASectionClass) New() DOMCDATASection {
	rv := objc.Send[DOMCDATASection](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCDATASection) Init() DOMCDATASection {
	rv := objc.Send[DOMCDATASection](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCDATASection) Autorelease() DOMCDATASection {
	rv := objc.Send[DOMCDATASection](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCDATASection creates a new DOMCDATASection instance.
func NewDOMCDATASection() DOMCDATASection {
	return getDOMCDATASectionClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCDATASection */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCDATASection
type DOMCDATASection struct {
	DOMText
}

// DOMCDATASectionFrom constructs a [DOMCDATASection] from an unsafe.Pointer.
func DOMCDATASectionFrom(ptr unsafe.Pointer) DOMCDATASection {
	return DOMCDATASection{
		DOMText: DOMTextFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCDATASection */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCDATASection */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCDATASection */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCDATASection */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCDATASection */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCDATASection */
