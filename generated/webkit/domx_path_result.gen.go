// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMXPathResult */

/* debug [class_header]: Header for DOMXPathResult */
// The class instance for the [DOMXPathResult] class.
var (
	DOMXPathResultClass     _DOMXPathResultClass
	DOMXPathResultClassOnce sync.Once
)

func getDOMXPathResultClass() _DOMXPathResultClass {
	DOMXPathResultClassOnce.Do(func() {
		DOMXPathResultClass = _DOMXPathResultClass{objc.GetClass("DOMXPathResult")}
	})
	return DOMXPathResultClass
}

type _DOMXPathResultClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMXPathResult */
// An interface definition for the [DOMXPathResult] class.
type IDOMXPathResult interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMXPathResult */
	// properties:
	BooleanValue() bool
	InvalidIteratorState() bool
	NumberValue() float64
	ResultType() unsafe.Pointer
	SingleNodeValue() IDOMNode
	SnapshotLength() unsafe.Pointer
	StringValue() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMXPathResult */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMXPathResult */
// Alloc allocates a new instance without initialization.
func (dc _DOMXPathResultClass) Alloc() DOMXPathResult {
	rv := objc.Send[DOMXPathResult](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMXPathResultClass) New() DOMXPathResult {
	rv := objc.Send[DOMXPathResult](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMXPathResult) Init() DOMXPathResult {
	rv := objc.Send[DOMXPathResult](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMXPathResult) Autorelease() DOMXPathResult {
	rv := objc.Send[DOMXPathResult](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMXPathResult creates a new DOMXPathResult instance.
func NewDOMXPathResult() DOMXPathResult {
	return getDOMXPathResultClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMXPathResult */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult
type DOMXPathResult struct {
	DOMObject
}

// DOMXPathResultFrom constructs a [DOMXPathResult] from an unsafe.Pointer.
func DOMXPathResultFrom(ptr unsafe.Pointer) DOMXPathResult {
	return DOMXPathResult{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMXPathResult */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMXPathResult */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMXPathResult */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMXPathResult */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMXPathResult */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/booleanValue
func (d_ DOMXPathResult) BooleanValue() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("booleanValue"))
	return rv
} /* debug [instance_properties/getter]: booleanValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/invalidIteratorState
func (d_ DOMXPathResult) InvalidIteratorState() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("invalidIteratorState"))
	return rv
} /* debug [instance_properties/getter]: invalidIteratorState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/numberValue
func (d_ DOMXPathResult) NumberValue() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("numberValue"))
	return rv
} /* debug [instance_properties/getter]: numberValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/resultType
func (d_ DOMXPathResult) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("resultType"))
	return rv
} /* debug [instance_properties/getter]: resultType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/singleNodeValue
func (d_ DOMXPathResult) SingleNodeValue() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("singleNodeValue"))
	return rv
} /* debug [instance_properties/getter]: singleNodeValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/snapshotLength
func (d_ DOMXPathResult) SnapshotLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("snapshotLength"))
	return rv
} /* debug [instance_properties/getter]: snapshotLength */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathResult/stringValue
func (d_ DOMXPathResult) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("stringValue"))
	return rv
} /* debug [instance_properties/getter]: stringValue */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMXPathResult */
