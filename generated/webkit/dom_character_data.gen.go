// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCharacterData */

/* debug [class_header]: Header for DOMCharacterData */
// The class instance for the [DOMCharacterData] class.
var (
	DOMCharacterDataClass     _DOMCharacterDataClass
	DOMCharacterDataClassOnce sync.Once
)

func getDOMCharacterDataClass() _DOMCharacterDataClass {
	DOMCharacterDataClassOnce.Do(func() {
		DOMCharacterDataClass = _DOMCharacterDataClass{objc.GetClass("DOMCharacterData")}
	})
	return DOMCharacterDataClass
}

type _DOMCharacterDataClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCharacterData */
// An interface definition for the [DOMCharacterData] class.
type IDOMCharacterData interface {
	IDOMNode

	/* debug [class_interface_properties]: Properties for DOMCharacterData */
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	Length() unsafe.Pointer
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCharacterData */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCharacterData */
// Alloc allocates a new instance without initialization.
func (dc _DOMCharacterDataClass) Alloc() DOMCharacterData {
	rv := objc.Send[DOMCharacterData](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCharacterDataClass) New() DOMCharacterData {
	rv := objc.Send[DOMCharacterData](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCharacterData) Init() DOMCharacterData {
	rv := objc.Send[DOMCharacterData](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCharacterData) Autorelease() DOMCharacterData {
	rv := objc.Send[DOMCharacterData](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCharacterData creates a new DOMCharacterData instance.
func NewDOMCharacterData() DOMCharacterData {
	return getDOMCharacterDataClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCharacterData */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCharacterData
type DOMCharacterData struct {
	DOMNode
}

// DOMCharacterDataFrom constructs a [DOMCharacterData] from an unsafe.Pointer.
func DOMCharacterDataFrom(ptr unsafe.Pointer) DOMCharacterData {
	return DOMCharacterData{
		DOMNode: DOMNodeFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCharacterData */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCharacterData */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCharacterData */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCharacterData */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCharacterData */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCharacterData/data
func (d_ DOMCharacterData) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("data"))
	return rv
} /* debug [instance_properties/getter]: data */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCharacterData/data
func (d_ DOMCharacterData) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setData:"), value)
} /* debug [instance_properties/setter]: data */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCharacterData/length
func (d_ DOMCharacterData) Length() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("length"))
	return rv
} /* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCharacterData */
