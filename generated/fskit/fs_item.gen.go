// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/metal"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSItem */


/* debug [class_header]: Header for FSItem */
// The class instance for the [FSItem] class.
var (
	FSItemClass     _FSItemClass
	FSItemClassOnce sync.Once
)

func getFSItemClass() _FSItemClass {
	FSItemClassOnce.Do(func() {
		FSItemClass = _FSItemClass{objc.GetClass("FSItem")}
	})
	return FSItemClass
}

type _FSItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSItem */
// An interface definition for the [FSItem] class.
type IFSItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSItem */
	// properties:
	WantedAttributes() metal.Attribute
	SetWantedAttributes(value metal.Attribute)
	ConsumedAttributes() metal.Attribute
	SetConsumedAttributes(value metal.Attribute)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSItem */
// Alloc allocates a new instance without initialization.
func (fc _FSItemClass) Alloc() FSItem {
	rv := objc.Send[FSItem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSItemClass) New() FSItem {
	rv := objc.Send[FSItem](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSItem) Init() FSItem {
	rv := objc.Send[FSItem](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSItem) Autorelease() FSItem {
	rv := objc.Send[FSItem](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItem creates a new FSItem instance.
func NewFSItem() FSItem {
	return getFSItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSItem */
// A distinct object in a file hierarchy, such as a file, directory, symlink, socket, and more.
//
// An is a mostly opaque object, which your file system implementation defines as needed. The class defines nonatomic properties to support instances. An instance contains a snapshot of the attributes of an at one point in time. The properties have no explicit thread safety provisions, since the operations that either get or set these properties enforce thread safety. You test an attribute’s validity with the the method . If the value is (Swift) or (Objective-C), it’s safe to use the attribute. Methods that get or set an item’s attribute use or , respectively. Both are subclasses of . An contains a property to indicate the attributes a file system provides for the request. Similarly, uses the property for a file system to signal back which attributes it successfully used. is the FSKit equivelant of a vnode in the kernel. For every FSKit vnode in the kernel, the hosting the volume has an instantiated .


// A distinct object in a file hierarchy, such as a file, directory, symlink, socket, and more.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem
type FSItem struct {
	objectivec.Object
}

// FSItemFrom constructs a [FSItem] from an unsafe.Pointer.
//
// A distinct object in a file hierarchy, such as a file, directory, symlink, socket, and more.
func FSItemFrom(ptr unsafe.Pointer) FSItem {
	return FSItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSItem */

// The attributes requested by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/getattributesrequest/wantedattributes
func (f_ FSItem) WantedAttributes() metal.Attribute {
	rv := objc.Send[metal.Attribute](f_.ID, objc.Sel("wantedAttributes"))
	return rv
}/* debug [instance_properties/getter]: wantedAttributes */


// The attributes requested by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/getattributesrequest/wantedattributes
func (f_ FSItem) SetWantedAttributes(value metal.Attribute) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWantedAttributes:"), value)
}/* debug [instance_properties/setter]: wantedAttributes */


// The attributes successfully used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/setattributesrequest/consumedattributes
func (f_ FSItem) ConsumedAttributes() metal.Attribute {
	rv := objc.Send[metal.Attribute](f_.ID, objc.Sel("consumedAttributes"))
	return rv
}/* debug [instance_properties/getter]: consumedAttributes */


// The attributes successfully used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/setattributesrequest/consumedattributes
func (f_ FSItem) SetConsumedAttributes(value metal.Attribute) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConsumedAttributes:"), value)
}/* debug [instance_properties/setter]: consumedAttributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSItem */



