// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TimedMetadata] class.
var (
	TimedMetadataClass     _TimedMetadataClass
	TimedMetadataClassOnce sync.Once
)

func getTimedMetadataClass() _TimedMetadataClass {
	TimedMetadataClassOnce.Do(func() {
		TimedMetadataClass = _TimedMetadataClass{objc.GetClass("MPTimedMetadata")}
	})
	return TimedMetadataClass
}

type _TimedMetadataClass struct {
	class objc.Class
}

// An interface definition for the [TimedMetadata] class.
type ITimedMetadata interface {
	objectivec.IObject
	ImageCropRect() coregraphics.CGRect
	SetImageCropRect(value coregraphics.CGRect)
	MPMoviePlayerTimedMetadataUserInfoKey() string
	AllMetadata() unsafe.Pointer
	SetAllMetadata(value unsafe.Pointer)
	Key() string
	SetKey(value string)
	Keyspace() string
	SetKeyspace(value string)
	Timestamp() unsafe.Pointer
	SetTimestamp(value unsafe.Pointer)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
}

// A carries time-based information within HTTP streamed media.
//
// Content providers can embed these objects when creating a stream. The properties and constants in this class let you extract the metadata as you play the stream using an object. For example, the provider of a live sports video stream could use instances to embed game scores, with timestamps, in the stream. On the client side—that is, on the user’s device—their application could employ the properties of this class to update their app’s user interface in real time during the game. A Javascript implementation of this class is also available for use by web-based applications.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata
type TimedMetadata struct {
	objectivec.Object
}

// TimedMetadataFrom constructs a [TimedMetadata] from an unsafe.Pointer.
//
// A carries time-based information within HTTP streamed media.
func TimedMetadataFrom(ptr unsafe.Pointer) TimedMetadata {
	return TimedMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TimedMetadataClass) Alloc() TimedMetadata {
	rv := objc.Send[TimedMetadata](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TimedMetadataClass) New() TimedMetadata {
	rv := objc.Send[TimedMetadata](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TimedMetadata) Init() TimedMetadata {
	rv := objc.Send[TimedMetadata](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TimedMetadata) Autorelease() TimedMetadata {
	rv := objc.Send[TimedMetadata](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimedMetadata creates a new TimedMetadata instance.
func NewTimedMetadata() TimedMetadata {
	return getTimedMetadataClass().New()
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (t_ TimedMetadata) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (t_ TimedMetadata) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImageCropRect:"), value)
}

// An NSDictionary object containing the most recent
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayertimedmetadatauserinfokey
func (t_ TimedMetadata) MPMoviePlayerTimedMetadataUserInfoKey() string {
	rv := objc.Send[string](t_.ID, objc.Sel("MPMoviePlayerTimedMetadataUserInfoKey"))
	return rv
}

// A dictionary containing all the metadata in the object.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/allmetadata
func (t_ TimedMetadata) AllMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("allMetadata"))
	return rv
}


// SetAllMetadata sets the value of the allMetadata property.
// A dictionary containing all the metadata in the object.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/allmetadata
func (t_ TimedMetadata) SetAllMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllMetadata:"), value)
}

// A key that identifies a piece of timed metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/key
func (t_ TimedMetadata) Key() string {
	rv := objc.Send[string](t_.ID, objc.Sel("key"))
	return rv
}


// SetKey sets the value of the key property.
// A key that identifies a piece of timed metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/key
func (t_ TimedMetadata) SetKey(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKey:"), objc.String(value))
}

// The namespace of the identifying key.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/keyspace
func (t_ TimedMetadata) Keyspace() string {
	rv := objc.Send[string](t_.ID, objc.Sel("keyspace"))
	return rv
}


// SetKeyspace sets the value of the keyspace property.
// The namespace of the identifying key.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/keyspace
func (t_ TimedMetadata) SetKeyspace(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeyspace:"), objc.String(value))
}

// The timestamp of the metadata, in the timebase of the media stream.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/timestamp
func (t_ TimedMetadata) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
// The timestamp of the metadata, in the timebase of the media stream.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/timestamp
func (t_ TimedMetadata) SetTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimestamp:"), value)
}

// The timed metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/value
func (t_ TimedMetadata) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The timed metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mptimedmetadata/value
func (t_ TimedMetadata) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setValue:"), value)
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (t_ TimedMetadata) ShowsRouteButton() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (t_ TimedMetadata) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowsRouteButton:"), value)
}



