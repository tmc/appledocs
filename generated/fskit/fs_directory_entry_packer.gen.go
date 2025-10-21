// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSDirectoryEntryPacker] class.
var (
	FSDirectoryEntryPackerClass     _FSDirectoryEntryPackerClass
	FSDirectoryEntryPackerClassOnce sync.Once
)

func getFSDirectoryEntryPackerClass() _FSDirectoryEntryPackerClass {
	FSDirectoryEntryPackerClassOnce.Do(func() {
		FSDirectoryEntryPackerClass = _FSDirectoryEntryPackerClass{objc.GetClass("FSDirectoryEntryPacker")}
	})
	return FSDirectoryEntryPackerClass
}

type _FSDirectoryEntryPackerClass struct {
	class objc.Class
}

// An interface definition for the [FSDirectoryEntryPacker] class.
type IFSDirectoryEntryPacker interface {
	objectivec.IObject
}

// An object used to provide items during a directory enumeration.
//
// You use this type in your implementation of . Packing allows your implementation to provide information FSKit needs, including each item’s name, type, and identifier (such as an inode number). Some directory enumerations require other attributes, as indicated by the sent to the enumerate method.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSDirectoryEntryPacker
type FSDirectoryEntryPacker struct {
	objectivec.Object
}

// FSDirectoryEntryPackerFrom constructs a [FSDirectoryEntryPacker] from an unsafe.Pointer.
//
// An object used to provide items during a directory enumeration.
func FSDirectoryEntryPackerFrom(ptr unsafe.Pointer) FSDirectoryEntryPacker {
	return FSDirectoryEntryPacker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSDirectoryEntryPackerClass) Alloc() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSDirectoryEntryPackerClass) New() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSDirectoryEntryPacker) Init() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSDirectoryEntryPacker) Autorelease() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSDirectoryEntryPacker creates a new FSDirectoryEntryPacker instance.
func NewFSDirectoryEntryPacker() FSDirectoryEntryPacker {
	return getFSDirectoryEntryPackerClass().New()
}




