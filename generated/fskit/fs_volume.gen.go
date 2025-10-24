// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSVolume */


/* debug [class_header]: Header for FSVolume */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSVolume */
// An interface definition for the [FSVolume] class.
type IFSVolume interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSVolume */
	// properties:
	Name() IFSFileName
	SetName(value IFSFileName)
	VolumeID() IFSVolumeIdentifier
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSVolume */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSVolume */
// Alloc allocates a new instance without initialization.
func (fc _FSVolumeClass) Alloc() FSVolume {
	rv := objc.Send[FSVolume](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSVolume */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSVolume */

// Creates a volume with the given identifier and name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/init(volumeID:volumeName:)
func NewFSVolumeWithVolumeIDVolumeName(volumeID IFSVolumeIdentifier, volumeName IFSFileName) FSVolume {
	instance := getFSVolumeClass().Alloc()
	rv := objc.Send[FSVolume](instance.ID, objc.Sel("initWithVolumeID:volumeName:"), volumeID, volumeName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFSVolumeWithVolumeIDVolumeName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSVolume */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSVolume */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSVolume */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSVolume */

// The name of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/name
func (f_ FSVolume) Name() IFSFileName {
	rv := objc.Send[FSFileName](f_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/name
func (f_ FSVolume) SetName(value IFSFileName) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// An identifier that uniquely identifies the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/volumeID
func (f_ FSVolume) VolumeID() IFSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("volumeID"))
	return rv
}/* debug [instance_properties/getter]: volumeID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSVolume */


