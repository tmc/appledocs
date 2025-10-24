// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSFileName */


/* debug [class_header]: Header for FSFileName */
// The class instance for the [FSFileName] class.
var (
	FSFileNameClass     _FSFileNameClass
	FSFileNameClassOnce sync.Once
)

func getFSFileNameClass() _FSFileNameClass {
	FSFileNameClassOnce.Do(func() {
		FSFileNameClass = _FSFileNameClass{objc.GetClass("FSFileName")}
	})
	return FSFileNameClass
}

type _FSFileNameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSFileName */
// An interface definition for the [FSFileName] class.
type IFSFileName interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSFileName */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	DebugDescription() objc.IObject /* cross-framework: NSString */
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSFileName */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSFileName */
// Alloc allocates a new instance without initialization.
func (fc _FSFileNameClass) Alloc() FSFileName {
	rv := objc.Send[FSFileName](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSFileNameClass) New() FSFileName {
	rv := objc.Send[FSFileName](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSFileName) Init() FSFileName {
	rv := objc.Send[FSFileName](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSFileName) Autorelease() FSFileName {
	rv := objc.Send[FSFileName](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSFileName creates a new FSFileName instance.
func NewFSFileName() FSFileName {
	return getFSFileNameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSFileName */
// The name of a file, expressed as a data buffer.
//
// is the class that carries filenames from the kernel to instances, and carries names back to the kernel as part of directory enumeration. A filename is usually a valid UTF-8 sequence, but can be an arbitrary byte sequence that doesn’t conform to that format. As a result, the property always contains a value, but the property may be empty. An can receive an that isn’t valid UTF-8 in two cases: A program passes erroneous data to a system call. The treats this situation as an error. An lacks the character encoding used for a file name. This situation occurs because some file system formats consider a filename to be an arbitrary “bag of bytes,” and leave character encoding up to the operating system. Without encoding information, the can only pass back the names it finds on disk. In this case, the behavior of upper layers such as is unspecified. However, the must support looking up such names and using them as the source name of rename operations. The must also be able to support filenames that are derivatives of filenames returned from directory enumeration. Derivative filenames include Apple Double filenames ( ), and editor backup filenames.


// The name of a file, expressed as a data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName
type FSFileName struct {
	objectivec.Object
}

// FSFileNameFrom constructs a [FSFileName] from an unsafe.Pointer.
//
// The name of a file, expressed as a data buffer.
func FSFileNameFrom(ptr unsafe.Pointer) FSFileName {
	return FSFileName{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSFileName */

// Initializes a file name by copying a character sequence from a byte array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/initWithBytes:length:
func NewFSFileNameWithBytesLength(bytes unsafe.Pointer, length uint) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[FSFileName](instance.ID, objc.Sel("initWithBytes:length:"), bytes, length)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSFileNameWithBytesLength */


// Initializes a filename from a null-terminated character sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/initWithCString:
func NewFSFileNameWithCString(name unsafe.Pointer) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[FSFileName](instance.ID, objc.Sel("initWithCString:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSFileNameWithCString */


// Creates a filename by copying a character sequence data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/init(data:)
func NewFSFileNameWithData(name objc.IObject /* cross-framework: NSData */) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[FSFileName](instance.ID, objc.Sel("initWithData:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSFileNameWithData */


// Creates a filename by copying a character sequence from a string instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/init(string:)
func NewFSFileNameWithString(name objc.IObject /* cross-framework: NSString */) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[FSFileName](instance.ID, objc.Sel("initWithString:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSFileNameWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSFileName */

// Creates a filename by copying a character sequence from a byte array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithBytes:length:
func (fc _FSFileNameClass) NameWithBytesLength(bytes unsafe.Pointer, length uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("nameWithBytes:length:"), bytes, length)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NameWithBytesLength) */


// Creates a filename from a null-terminated character sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithCString:
func (fc _FSFileNameClass) NameWithCString(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("nameWithCString:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NameWithCString) */


// Creates a filename by copying a character sequence data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithData:
func (fc _FSFileNameClass) NameWithData(name objc.IObject /* cross-framework: NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("nameWithData:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NameWithData) */


// Creates a filename by copying a character sequence from a string instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithString:
func (fc _FSFileNameClass) NameWithString(name objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("nameWithString:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NameWithString) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSFileName */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSFileName */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSFileName */

// The byte sequence of the filename, as a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/data
func (f_ FSFileName) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](f_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The filename, represented as a potentially lossy conversion to a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/debugDescription
func (f_ FSFileName) DebugDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("debugDescription"))
	return rv
}/* debug [instance_properties/getter]: debugDescription */


// The filename, represented as a Unicode string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileName/string
func (f_ FSFileName) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSFileName */


