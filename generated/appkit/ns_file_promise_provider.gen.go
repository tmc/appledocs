// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFilePromiseProvider */


/* debug [class_header]: Header for NSFilePromiseProvider */
// The class instance for the [FilePromiseProvider] class.
var (
	FilePromiseProviderClass     _FilePromiseProviderClass
	FilePromiseProviderClassOnce sync.Once
)

func getFilePromiseProviderClass() _FilePromiseProviderClass {
	FilePromiseProviderClassOnce.Do(func() {
		FilePromiseProviderClass = _FilePromiseProviderClass{objc.GetClass("NSFilePromiseProvider")}
	})
	return FilePromiseProviderClass
}

type _FilePromiseProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FilePromiseProvider */
// An interface definition for the [FilePromiseProvider] class.
type IFilePromiseProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FilePromiseProvider */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FileType() objc.IObject /* cross-framework: NSString */
	SetFileType(value objc.IObject /* cross-framework: NSString */)
	UserInfo() objc.ID
	SetUserInfo(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FilePromiseProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FilePromiseProvider */
// Alloc allocates a new instance without initialization.
func (fc _FilePromiseProviderClass) Alloc() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FilePromiseProviderClass) New() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilePromiseProvider) Init() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilePromiseProvider) Autorelease() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilePromiseProvider creates a new FilePromiseProvider instance.
func NewFilePromiseProvider() FilePromiseProvider {
	return getFilePromiseProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FilePromiseProvider */
// An object that provides a promise for the pasteboard.
//
// A file promise is a possible future file of a specified type. When you’re working with drag and drop, use promises to indicate intent for future action. Avoid loading or performing any actions on the file until the promise completes. Use the class when creating file promises. Instantiate one for each file promised. Set the and properties before writing any to the pasteboard. The file type must be a Uniform Type Identifier (UTI) that ultimately conforms to or . The will write the promised file to the destination directory. Optionally, you may attach a object of your choosing to the to determine which promise is being referenced when promising multiple files under the same instance.


// An object that provides a promise for the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider
type FilePromiseProvider struct {
	objectivec.Object
}

// FilePromiseProviderFrom constructs a [FilePromiseProvider] from an unsafe.Pointer.
//
// An object that provides a promise for the pasteboard.
func FilePromiseProviderFrom(ptr unsafe.Pointer) FilePromiseProvider {
	return FilePromiseProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FilePromiseProvider */

// Initializes a file promise provider for a certain file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/init(fileType:delegate:)
func NewFilePromiseProviderWithFileTypeDelegate(fileType objc.IObject /* cross-framework: NSString */, delegate unsafe.Pointer) FilePromiseProvider {
	instance := getFilePromiseProviderClass().Alloc()
	rv := objc.Send[FilePromiseProvider](instance.ID, objc.Sel("initWithFileType:delegate:"), fileType, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFilePromiseProviderWithFileTypeDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FilePromiseProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FilePromiseProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FilePromiseProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FilePromiseProvider */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/delegate
func (f_ FilePromiseProvider) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/delegate
func (f_ FilePromiseProvider) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The file type of the file promise provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/fileType
func (f_ FilePromiseProvider) FileType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("fileType"))
	return rv
}/* debug [instance_properties/getter]: fileType */


// The file type of the file promise provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/fileType
func (f_ FilePromiseProvider) SetFileType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileType:"), value)
}/* debug [instance_properties/setter]: fileType */


// Optional user information to pass to the file promise provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/userInfo
func (f_ FilePromiseProvider) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// Optional user information to pass to the file promise provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/userInfo
func (f_ FilePromiseProvider) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFilePromiseProvider */



