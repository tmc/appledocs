// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FSItem] class.
type IFSItem interface {
	objectivec.IObject
	WantedAttributes() unsafe.Pointer
	SetWantedAttributes(value unsafe.Pointer)
	ConsumedAttributes() unsafe.Pointer
	SetConsumedAttributes(value unsafe.Pointer)
}

// A distinct object in a file hierarchy, such as a file, directory, symlink, socket, and more.
//
// An is a mostly opaque object, which your file system implementation defines as needed. The class defines nonatomic properties to support instances. An instance contains a snapshot of the attributes of an at one point in time. The properties have no explicit thread safety provisions, since the operations that either get or set these properties enforce thread safety. You test an attribute’s validity with the the method . If the value is (Swift) or (Objective-C), it’s safe to use the attribute. Methods that get or set an item’s attribute use or , respectively. Both are subclasses of . An contains a property to indicate the attributes a file system provides for the request. Similarly, uses the property for a file system to signal back which attributes it successfully used. is the FSKit equivelant of a vnode in the kernel. For every FSKit vnode in the kernel, the hosting the volume has an instantiated .
//
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

// Alloc allocates a new instance without initialization.
func (fc _FSItemClass) Alloc() FSItem {
	rv := objc.Send[FSItem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The attributes requested by the request.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/getattributesrequest/wantedattributes
func (f_ FSItem) WantedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("wantedAttributes"))
	return rv
}


// SetWantedAttributes sets the value of the wantedAttributes property.
// The attributes requested by the request.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/getattributesrequest/wantedattributes
func (f_ FSItem) SetWantedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWantedAttributes:"), value)
}

// The attributes successfully used by the file system.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/setattributesrequest/consumedattributes
func (f_ FSItem) ConsumedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("consumedAttributes"))
	return rv
}


// SetConsumedAttributes sets the value of the consumedAttributes property.
// The attributes successfully used by the file system.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsitem/setattributesrequest/consumedattributes
func (f_ FSItem) SetConsumedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConsumedAttributes:"), value)
}



