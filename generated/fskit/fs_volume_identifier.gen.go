// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FSVolumeIdentifier] class.
var (
	FSVolumeIdentifierClass     _FSVolumeIdentifierClass
	FSVolumeIdentifierClassOnce sync.Once
)

func getFSVolumeIdentifierClass() _FSVolumeIdentifierClass {
	FSVolumeIdentifierClassOnce.Do(func() {
		FSVolumeIdentifierClass = _FSVolumeIdentifierClass{objc.GetClass("FSVolumeIdentifier")}
	})
	return FSVolumeIdentifierClass
}

type _FSVolumeIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [FSVolumeIdentifier] class.
type IFSVolumeIdentifier interface {
	IFSEntityIdentifier
}

// A type that identifies a volume.
//
// For most volumes, the volume identifier is the UUID identifying the volume. Network file systems may access the same underlying volume using different authentication credentials. To handle this situation, add qualifying data to identify the specific container, as discussed in the superclass, .


// A type that identifies a volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/Identifier
type FSVolumeIdentifier struct {
	FSEntityIdentifier
}

// FSVolumeIdentifierFrom constructs a [FSVolumeIdentifier] from an unsafe.Pointer.
//
// A type that identifies a volume.
func FSVolumeIdentifierFrom(ptr unsafe.Pointer) FSVolumeIdentifier {
	return FSVolumeIdentifier{
		FSEntityIdentifier: FSEntityIdentifierFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FSVolumeIdentifierClass) Alloc() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSVolumeIdentifierClass) New() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSVolumeIdentifier) Init() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSVolumeIdentifier) Autorelease() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolumeIdentifier creates a new FSVolumeIdentifier instance.
func NewFSVolumeIdentifier() FSVolumeIdentifier {
	return getFSVolumeIdentifierClass().New()
}




