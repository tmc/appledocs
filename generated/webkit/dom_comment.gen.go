// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMComment */

/* debug [class_header]: Header for DOMComment */
// The class instance for the [DOMComment] class.
var (
	DOMCommentClass     _DOMCommentClass
	DOMCommentClassOnce sync.Once
)

func getDOMCommentClass() _DOMCommentClass {
	DOMCommentClassOnce.Do(func() {
		DOMCommentClass = _DOMCommentClass{objc.GetClass("DOMComment")}
	})
	return DOMCommentClass
}

type _DOMCommentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMComment */
// An interface definition for the [DOMComment] class.
type IDOMComment interface {
	IDOMCharacterData

	/* debug [class_interface_properties]: Properties for DOMComment */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMComment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMComment */
// Alloc allocates a new instance without initialization.
func (dc _DOMCommentClass) Alloc() DOMComment {
	rv := objc.Send[DOMComment](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCommentClass) New() DOMComment {
	rv := objc.Send[DOMComment](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMComment) Init() DOMComment {
	rv := objc.Send[DOMComment](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMComment) Autorelease() DOMComment {
	rv := objc.Send[DOMComment](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMComment creates a new DOMComment instance.
func NewDOMComment() DOMComment {
	return getDOMCommentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMComment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMComment
type DOMComment struct {
	DOMCharacterData
}

// DOMCommentFrom constructs a [DOMComment] from an unsafe.Pointer.
func DOMCommentFrom(ptr unsafe.Pointer) DOMComment {
	return DOMComment{
		DOMCharacterData: DOMCharacterDataFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMComment */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMComment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMComment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMComment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMComment */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMComment */
