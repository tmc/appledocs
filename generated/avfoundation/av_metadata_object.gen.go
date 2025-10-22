// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataObject] class.
var (
	MetadataObjectClass     _MetadataObjectClass
	MetadataObjectClassOnce sync.Once
)

func getMetadataObjectClass() _MetadataObjectClass {
	MetadataObjectClassOnce.Do(func() {
		MetadataObjectClass = _MetadataObjectClass{objc.GetClass("AVMetadataObject")}
	})
	return MetadataObjectClass
}

type _MetadataObjectClass struct {
	class objc.Class
}

// An interface definition for the [MetadataObject] class.
type IMetadataObject interface {
	objectivec.IObject
	Bounds() coregraphics.CGRect
	CinematicVideoFocusMode() CaptureCinematicVideoFocusMode
	Duration() unsafe.Pointer
	GroupID() int
	FixedFocus() bool
	ObjectID() int
	Time() unsafe.Pointer
	Type() MetadataObjectType
	IsFixedFocus() bool
	SetIsFixedFocus(value bool)
}

// The abstract superclass for objects provided by a metadata capture output.
//
// The class is an abstract class that defines the basic properties associated with a piece of metadata. These attributes reflect information either about the metadata itself or the media from which the metadata originated. Subclasses are responsible for providing appropriate values for each of the relevant properties. You shouldn’t subclass directly. Instead, you use one of the defined subclasses provided by the AVFoundation framework. Similarly, you don’t create instances of this class yourself but use an object to retrieve them from the captured data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject
type MetadataObject struct {
	objectivec.Object
}

// MetadataObjectFrom constructs a [MetadataObject] from an unsafe.Pointer.
//
// The abstract superclass for objects provided by a metadata capture output.
func MetadataObjectFrom(ptr unsafe.Pointer) MetadataObject {
	return MetadataObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataObjectClass) Alloc() MetadataObject {
	rv := objc.Send[MetadataObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataObjectClass) New() MetadataObject {
	rv := objc.Send[MetadataObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataObject) Init() MetadataObject {
	rv := objc.Send[MetadataObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataObject) Autorelease() MetadataObject {
	rv := objc.Send[MetadataObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataObject creates a new MetadataObject instance.
func NewMetadataObject() MetadataObject {
	return getMetadataObjectClass().New()
}


// The bounding rectangle associated with the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/bounds
func (m_ MetadataObject) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("bounds"))
	return rv
}

// The current focus mode when an object is detected during a Cinematic Video recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/cinematicVideoFocusMode
func (m_ MetadataObject) CinematicVideoFocusMode() CaptureCinematicVideoFocusMode {
	rv := objc.Send[CaptureCinematicVideoFocusMode](m_.ID, objc.Sel("cinematicVideoFocusMode"))
	return rv
}

// The duration of the media associated with this metadata object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/duration
func (m_ MetadataObject) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}

// An identifier associated with a metadata object used to group it with other metadata objects belonging to a common parent.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/groupID
func (m_ MetadataObject) GroupID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("groupID"))
	return rv
}

// A BOOL indicating whether this metadata object represents a fixed focus.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/isFixedFocus
func (m_ MetadataObject) FixedFocus() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("fixedFocus"))
	return rv
}

// A unique identifier for each detected object type (face, body, hands, heads and salient objects) in a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/objectID
func (m_ MetadataObject) ObjectID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("objectID"))
	return rv
}

// The media time value associated with the metadata object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/time
func (m_ MetadataObject) Time() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("time"))
	return rv
}

// The type of metadata that this object provides.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/type
func (m_ MetadataObject) Type() MetadataObjectType {
	rv := objc.Send[MetadataObjectType](m_.ID, objc.Sel("type"))
	return rv
}

// A BOOL indicating whether this metadata object represents a fixed focus.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject/isfixedfocus
func (m_ MetadataObject) IsFixedFocus() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isFixedFocus"))
	return rv
}


// SetIsFixedFocus sets the value of the isFixedFocus property.
// A BOOL indicating whether this metadata object represents a fixed focus.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataobject/isfixedfocus
func (m_ MetadataObject) SetIsFixedFocus(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsFixedFocus:"), value)
}



