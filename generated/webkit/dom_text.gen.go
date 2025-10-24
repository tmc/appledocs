// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMText */

/* debug [class_header]: Header for DOMText */
// The class instance for the [DOMText] class.
var (
	DOMTextClass     _DOMTextClass
	DOMTextClassOnce sync.Once
)

func getDOMTextClass() _DOMTextClass {
	DOMTextClassOnce.Do(func() {
		DOMTextClass = _DOMTextClass{objc.GetClass("DOMText")}
	})
	return DOMTextClass
}

type _DOMTextClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMText */
// An interface definition for the [DOMText] class.
type IDOMText interface {
	IDOMCharacterData

	/* debug [class_interface_properties]: Properties for DOMText */
	// properties:
	WholeText() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMText */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMText */
// Alloc allocates a new instance without initialization.
func (dc _DOMTextClass) Alloc() DOMText {
	rv := objc.Send[DOMText](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMTextClass) New() DOMText {
	rv := objc.Send[DOMText](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMText) Init() DOMText {
	rv := objc.Send[DOMText](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMText) Autorelease() DOMText {
	rv := objc.Send[DOMText](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMText creates a new DOMText instance.
func NewDOMText() DOMText {
	return getDOMTextClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMText
type DOMText struct {
	DOMCharacterData
}

// DOMTextFrom constructs a [DOMText] from an unsafe.Pointer.
func DOMTextFrom(ptr unsafe.Pointer) DOMText {
	return DOMText{
		DOMCharacterData: DOMCharacterDataFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMText */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMText */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMText */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMText */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMText/wholeText
func (d_ DOMText) WholeText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("wholeText"))
	return rv
} /* debug [instance_properties/getter]: wholeText */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMText */
