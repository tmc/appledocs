// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSFileSystem] class.
var (
	FSFileSystemClass     _FSFileSystemClass
	FSFileSystemClassOnce sync.Once
)

func getFSFileSystemClass() _FSFileSystemClass {
	FSFileSystemClassOnce.Do(func() {
		FSFileSystemClass = _FSFileSystemClass{objc.GetClass("FSFileSystem")}
	})
	return FSFileSystemClass
}

type _FSFileSystemClass struct {
	class objc.Class
}

// An interface definition for the [FSFileSystem] class.
type IFSFileSystem interface {
	objectivec.IObject
}

// An abstract base class for implementing a full-featured file system.
//
// is a full-featured file system, which works with one or more instances and presents one or more references to callers. Implement your app extension by providing a subclass of as a delegate object. Your delegate also needs to implement the protocol so that it can probe, load, and unload resources.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSFileSystem
type FSFileSystem struct {
	objectivec.Object
}

// FSFileSystemFrom constructs a [FSFileSystem] from an unsafe.Pointer.
//
// An abstract base class for implementing a full-featured file system.
func FSFileSystemFrom(ptr unsafe.Pointer) FSFileSystem {
	return FSFileSystem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSFileSystemClass) Alloc() FSFileSystem {
	rv := objc.Send[FSFileSystem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSFileSystemClass) New() FSFileSystem {
	rv := objc.Send[FSFileSystem](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSFileSystem) Init() FSFileSystem {
	rv := objc.Send[FSFileSystem](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSFileSystem) Autorelease() FSFileSystem {
	rv := objc.Send[FSFileSystem](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSFileSystem creates a new FSFileSystem instance.
func NewFSFileSystem() FSFileSystem {
	return getFSFileSystemClass().New()
}




