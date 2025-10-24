//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for TimedMetadata


// iOS-only properties

// A dictionary containing all the metadata in the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata/allMetadata
func (t_ TimedMetadata) AllMetadata() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](t_.ID, objc.Sel("allMetadata"))
	return rv
}

// A key that identifies a piece of timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata/key
func (t_ TimedMetadata) Key() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("key"))
	return rv
}

// The namespace of the identifying key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata/keyspace
func (t_ TimedMetadata) Keyspace() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("keyspace"))
	return rv
}

// The timestamp of the metadata, in the timebase of the media stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata/timestamp
func (t_ TimedMetadata) Timestamp() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("timestamp"))
	return rv
}

// The timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata/value
func (t_ TimedMetadata) Value() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("value"))
	return rv
}







