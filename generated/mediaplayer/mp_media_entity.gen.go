// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaEntity] class.
var (
	MediaEntityClass     _MediaEntityClass
	MediaEntityClassOnce sync.Once
)

func getMediaEntityClass() _MediaEntityClass {
	MediaEntityClassOnce.Do(func() {
		MediaEntityClass = _MediaEntityClass{objc.GetClass("MPMediaEntity")}
	})
	return MediaEntityClass
}

type _MediaEntityClass struct {
	class objc.Class
}

// An interface definition for the [MediaEntity] class.
type IMediaEntity interface {
	objectivec.IObject
	EnumerateValuesForPropertiesUsingBlock(properties unsafe.Pointer, block unsafe.Pointer)
	ObjectForKeyedSubscript(key objc.ID) objc.ID
	ValueForProperty(property string) objc.ID
}

// The abstract superclass for media items, media item collections, and media playlist instances.
//
// This is the superclass for and instances, and in turn for instances.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity
type MediaEntity struct {
	objectivec.Object
}

// MediaEntityFrom constructs a [MediaEntity] from an unsafe.Pointer.
//
// The abstract superclass for media items, media item collections, and media playlist instances.
func MediaEntityFrom(ptr unsafe.Pointer) MediaEntity {
	return MediaEntity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaEntityClass) Alloc() MediaEntity {
	rv := objc.Send[MediaEntity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaEntityClass) New() MediaEntity {
	rv := objc.Send[MediaEntity](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaEntity) Init() MediaEntity {
	rv := objc.Send[MediaEntity](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaEntity) Autorelease() MediaEntity {
	rv := objc.Send[MediaEntity](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaEntity creates a new MediaEntity instance.
func NewMediaEntity() MediaEntity {
	return getMediaEntityClass().New()
}


// Indicates whether you can use the media property key that you specify to construct a media property predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/canFilter(byProperty:)
func (mc _MediaEntityClass) CanFilterByProperty(property string) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("canFilterByProperty:"), objc.String(property))
	return rv
}

// Executes a provided block with the fetched values for the given item properties.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/enumerateValues(forProperties:using:)
func (m_ MediaEntity) EnumerateValuesForPropertiesUsingBlock(properties unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enumerateValuesForProperties:usingBlock:"), properties, block)
}

// Returns the object specified by the key.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/subscript(_:)
func (m_ MediaEntity) ObjectForKeyedSubscript(key objc.ID) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Retrieves the value for a specified media property key.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/value(forProperty:)
func (m_ MediaEntity) ValueForProperty(property string) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("valueForProperty:"), objc.String(property))
	return rv
}

// The persistent identifier for a media entity.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/persistentID
func (m_ MediaEntity) PersistentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("persistentID"))
	return rv
}



