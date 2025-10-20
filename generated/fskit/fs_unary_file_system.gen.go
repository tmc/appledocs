// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSUnaryFileSystem] class.
var (
	FSUnaryFileSystemClass     _FSUnaryFileSystemClass
	FSUnaryFileSystemClassOnce sync.Once
)

func getFSUnaryFileSystemClass() _FSUnaryFileSystemClass {
	FSUnaryFileSystemClassOnce.Do(func() {
		FSUnaryFileSystemClass = _FSUnaryFileSystemClass{objc.GetClass("FSUnaryFileSystem")}
	})
	return FSUnaryFileSystemClass
}

type _FSUnaryFileSystemClass struct {
	class objc.Class
}

// An interface definition for the [FSUnaryFileSystem] class.
type IFSUnaryFileSystem interface {
	objectivec.IObject
}

// An abstract base class for implementing a minimal file system.
//
// is a simplified file system, which works with one and presents it as one . The one volume and its container have a shared state and lifetime, a more constrained life cycle than the design flow. Implement your app extension by providing a subclass of as a delegate object. Your delegate also needs to implement the protocol so that it can load resources.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystem
type FSUnaryFileSystem struct {
	objectivec.Object
}

// FSUnaryFileSystemFrom constructs a [FSUnaryFileSystem] from an unsafe.Pointer.
//
// An abstract base class for implementing a minimal file system.
func FSUnaryFileSystemFrom(ptr unsafe.Pointer) FSUnaryFileSystem {
	return FSUnaryFileSystem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSUnaryFileSystemClass) Alloc() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSUnaryFileSystemClass) New() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSUnaryFileSystem) Init() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSUnaryFileSystem) Autorelease() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSUnaryFileSystem creates a new FSUnaryFileSystem instance.
func NewFSUnaryFileSystem() FSUnaryFileSystem {
	return getFSUnaryFileSystemClass().New()
}




