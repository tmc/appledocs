// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSVolume] class.
var (
	FSVolumeClass     _FSVolumeClass
	FSVolumeClassOnce sync.Once
)

func getFSVolumeClass() _FSVolumeClass {
	FSVolumeClassOnce.Do(func() {
		FSVolumeClass = _FSVolumeClass{objc.GetClass("FSVolume")}
	})
	return FSVolumeClass
}

type _FSVolumeClass struct {
	class objc.Class
}

// An interface definition for the [FSVolume] class.
type IFSVolume interface {
	objectivec.IObject
	Name() FSFileName
	SetName(value IFSFileName)
	VolumeID() FSVolumeIdentifier
}

// A directory structure for files and folders.
//
// A file system, depending on its type, provides one or more volumes to clients. The by definition provides only one volume, while an supports multiple volumes. You implement a volume for your file system type by subclassing this class, and also conforming to the and protocols. This protocol defines the minimum set of operations supported by a volume, such as mounting, activating, creating and removing items, and more. Your volume can provide additional functionality by conforming to other volume operations protocols. These protocols add support for operations like open and close, read and write, extended attribute (Xattr) manipulation, and more.


// A directory structure for files and folders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume

type FSVolume struct {
	objectivec.Object
}

// FSVolumeFrom constructs a [FSVolume] from an unsafe.Pointer.
//
// A directory structure for files and folders.
func FSVolumeFrom(ptr unsafe.Pointer) FSVolume {
	return FSVolume{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSVolumeClass) Alloc() FSVolume {
	rv := objc.Send[FSVolume](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSVolumeClass) New() FSVolume {
	rv := objc.Send[FSVolume](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSVolume) Init() FSVolume {
	rv := objc.Send[FSVolume](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSVolume) Autorelease() FSVolume {
	rv := objc.Send[FSVolume](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolume creates a new FSVolume instance.
func NewFSVolume() FSVolume {
	return getFSVolumeClass().New()
}




// Creates a volume with the given identifier and name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/init(volumeID:volumeName:)

func NewFSVolumeWithVolumeIDVolumeName(volumeID IFSVolumeIdentifier, volumeName IFSFileName) FSVolume {
	instance := getFSVolumeClass().Alloc()
	rv := objc.Send[FSVolume](instance.ID, objc.Sel("initWithVolumeID:volumeName:"), volumeID, volumeName)
	rv.Autorelease()
	return rv
}



// The name of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/name

func (f_ FSVolume) Name() FSFileName {
	rv := objc.Send[FSFileName](f_.ID, objc.Sel("name"))
	return rv
}


// The name of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/name

func (f_ FSVolume) SetName(value IFSFileName) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}


// An identifier that uniquely identifies the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/volumeID

func (f_ FSVolume) VolumeID() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("volumeID"))
	return rv
}


