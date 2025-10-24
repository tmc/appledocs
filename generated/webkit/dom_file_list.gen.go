// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMFileList */


/* debug [class_header]: Header for DOMFileList */
// The class instance for the [DOMFileList] class.
var (
	DOMFileListClass     _DOMFileListClass
	DOMFileListClassOnce sync.Once
)

func getDOMFileListClass() _DOMFileListClass {
	DOMFileListClassOnce.Do(func() {
		DOMFileListClass = _DOMFileListClass{objc.GetClass("DOMFileList")}
	})
	return DOMFileListClass
}

type _DOMFileListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMFileList */
// An interface definition for the [DOMFileList] class.
type IDOMFileList interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMFileList */
	// properties:
	Length() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMFileList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMFileList */
// Alloc allocates a new instance without initialization.
func (dc _DOMFileListClass) Alloc() DOMFileList {
	rv := objc.Send[DOMFileList](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMFileListClass) New() DOMFileList {
	rv := objc.Send[DOMFileList](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMFileList) Init() DOMFileList {
	rv := objc.Send[DOMFileList](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMFileList) Autorelease() DOMFileList {
	rv := objc.Send[DOMFileList](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMFileList creates a new DOMFileList instance.
func NewDOMFileList() DOMFileList {
	return getDOMFileListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMFileList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMFileList
type DOMFileList struct {
	DOMObject
}

// DOMFileListFrom constructs a [DOMFileList] from an unsafe.Pointer.
func DOMFileListFrom(ptr unsafe.Pointer) DOMFileList {
	return DOMFileList{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMFileList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMFileList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMFileList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMFileList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMFileList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMFileList/length
func (d_ DOMFileList) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMFileList */



