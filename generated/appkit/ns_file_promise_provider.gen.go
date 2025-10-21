// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FilePromiseProvider] class.
type IFilePromiseProvider interface {
	objectivec.IObject
}

// An object that provides a promise for the pasteboard.
//
// A file promise is a possible future file of a specified type. When you’re working with drag and drop, use promises to indicate intent for future action. Avoid loading or performing any actions on the file until the promise completes. Use the class when creating file promises. Instantiate one for each file promised. Set the and properties before writing any to the pasteboard. The file type must be a Uniform Type Identifier (UTI) that ultimately conforms to or . The will write the promised file to the destination directory. Optionally, you may attach a object of your choosing to the to determine which promise is being referenced when promising multiple files under the same instance.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FilePromiseProviderClass) Alloc() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a file promise provider for a certain file type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/init(fileType:delegate:)
func NewFilePromiseProviderWithFileTypeDelegate(fileType string, delegate objc.ID) FilePromiseProvider {
	instance := getFilePromiseProviderClass().Alloc()
	rv := objc.Send[FilePromiseProvider](instance.ID, objc.Sel("initWithFileType:delegate:"), objc.String(fileType), delegate)
	rv.Autorelease()
	return rv
}




