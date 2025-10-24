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

// iOS-only methods for MovieErrorLog


// iOS-only properties

// The events in the movie error log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLog/events
func (m_ MovieErrorLog) Events() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("events"))
	return rv
}

// A textual version of the web server error log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLog/extendedLogData
func (m_ MovieErrorLog) ExtendedLogData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("extendedLogData"))
	return rv
}

// The string encoding for the extended log data property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLog/extendedLogDataStringEncoding
func (m_ MovieErrorLog) ExtendedLogDataStringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](m_.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}





