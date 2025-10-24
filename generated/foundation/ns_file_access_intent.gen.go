// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileAccessIntent */


/* debug [class_header]: Header for NSFileAccessIntent */
// The class instance for the [FileAccessIntent] class.
var (
	FileAccessIntentClass     _FileAccessIntentClass
	FileAccessIntentClassOnce sync.Once
)

func getFileAccessIntentClass() _FileAccessIntentClass {
	FileAccessIntentClassOnce.Do(func() {
		FileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}
	})
	return FileAccessIntentClass
}

type _FileAccessIntentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileAccessIntent */
// An interface definition for the [FileAccessIntent] class.
type IFileAccessIntent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileAccessIntent */
	// properties:
	Url() IURL
	SetUrl(value IURL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileAccessIntent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileAccessIntent */
// Alloc allocates a new instance without initialization.
func (fc _FileAccessIntentClass) Alloc() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileAccessIntentClass) New() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileAccessIntent) Init() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileAccessIntent) Autorelease() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileAccessIntent creates a new FileAccessIntent instance.
func NewFileAccessIntent() FileAccessIntent {
	return getFileAccessIntentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileAccessIntent */
// The details of a coordinated-read or coordinated-write operation.
//
// Use this class when performing asynchronous operations with a file coordinator using the coordinator’s method.


// The details of a coordinated-read or coordinated-write operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent
type FileAccessIntent struct {
	objectivec.Object
}

// FileAccessIntentFrom constructs a [FileAccessIntent] from an unsafe.Pointer.
//
// The details of a coordinated-read or coordinated-write operation.
func FileAccessIntentFrom(ptr unsafe.Pointer) FileAccessIntent {
	return FileAccessIntent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileAccessIntent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileAccessIntent */

// Returns a file access intent object for writing to the given URL with the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent/writingIntent(with:options:)
func (fc _FileAccessIntentClass) WritingIntentWithURLOptions(url IURL, options FileCoordinatorWritingOptions) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("writingIntentWithURL:options:"), url, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WritingIntentWithURLOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileAccessIntent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileAccessIntent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileAccessIntent */

// The current URL for the item managed by the file access intent instance. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileaccessintent/url
func (f_ FileAccessIntent) Url() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The current URL for the item managed by the file access intent instance. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileaccessintent/url
func (f_ FileAccessIntent) SetUrl(value IURL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileAccessIntent */



