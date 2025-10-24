// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMMediaList */

/* debug [class_header]: Header for DOMMediaList */
// The class instance for the [DOMMediaList] class.
var (
	DOMMediaListClass     _DOMMediaListClass
	DOMMediaListClassOnce sync.Once
)

func getDOMMediaListClass() _DOMMediaListClass {
	DOMMediaListClassOnce.Do(func() {
		DOMMediaListClass = _DOMMediaListClass{objc.GetClass("DOMMediaList")}
	})
	return DOMMediaListClass
}

type _DOMMediaListClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMMediaList */
// An interface definition for the [DOMMediaList] class.
type IDOMMediaList interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMMediaList */
	// properties:
	Length() unsafe.Pointer
	MediaText() objc.IObject /* cross-framework: NSString */
	SetMediaText(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMMediaList */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMMediaList */
// Alloc allocates a new instance without initialization.
func (dc _DOMMediaListClass) Alloc() DOMMediaList {
	rv := objc.Send[DOMMediaList](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMMediaListClass) New() DOMMediaList {
	rv := objc.Send[DOMMediaList](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMMediaList) Init() DOMMediaList {
	rv := objc.Send[DOMMediaList](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMMediaList) Autorelease() DOMMediaList {
	rv := objc.Send[DOMMediaList](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMMediaList creates a new DOMMediaList instance.
func NewDOMMediaList() DOMMediaList {
	return getDOMMediaListClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMMediaList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMediaList
type DOMMediaList struct {
	DOMObject
}

// DOMMediaListFrom constructs a [DOMMediaList] from an unsafe.Pointer.
func DOMMediaListFrom(ptr unsafe.Pointer) DOMMediaList {
	return DOMMediaList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMMediaList */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMMediaList */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMMediaList */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMMediaList */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMMediaList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMediaList/length
func (d_ DOMMediaList) Length() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("length"))
	return rv
} /* debug [instance_properties/getter]: length */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMediaList/mediaText
func (d_ DOMMediaList) MediaText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("mediaText"))
	return rv
} /* debug [instance_properties/getter]: mediaText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMediaList/mediaText
func (d_ DOMMediaList) SetMediaText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMediaText:"), value)
} /* debug [instance_properties/setter]: mediaText */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMMediaList */
