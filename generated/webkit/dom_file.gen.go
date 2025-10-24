// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMFile */


/* debug [class_header]: Header for DOMFile */
// The class instance for the [DOMFile] class.
var (
	DOMFileClass     _DOMFileClass
	DOMFileClassOnce sync.Once
)

func getDOMFileClass() _DOMFileClass {
	DOMFileClassOnce.Do(func() {
		DOMFileClass = _DOMFileClass{objc.GetClass("DOMFile")}
	})
	return DOMFileClass
}

type _DOMFileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMFile */
// An interface definition for the [DOMFile] class.
type IDOMFile interface {
	IDOMBlob
	
/* debug [class_interface_properties]: Properties for DOMFile */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMFile */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMFile */
// Alloc allocates a new instance without initialization.
func (dc _DOMFileClass) Alloc() DOMFile {
	rv := objc.Send[DOMFile](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMFileClass) New() DOMFile {
	rv := objc.Send[DOMFile](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMFile) Init() DOMFile {
	rv := objc.Send[DOMFile](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMFile) Autorelease() DOMFile {
	rv := objc.Send[DOMFile](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMFile creates a new DOMFile instance.
func NewDOMFile() DOMFile {
	return getDOMFileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMFile */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMFile
type DOMFile struct {
	DOMBlob
}

// DOMFileFrom constructs a [DOMFile] from an unsafe.Pointer.
func DOMFileFrom(ptr unsafe.Pointer) DOMFile {
	return DOMFile{
		DOMBlob: DOMBlobFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMFile *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMFile */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMFile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMFile */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMFile */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMFile/name
func (d_ DOMFile) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMFile */



