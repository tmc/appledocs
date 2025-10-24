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

// iOS-only methods for MovieErrorLogEvent


// iOS-only properties

// The date and time when the error occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/date
func (m_ MovieErrorLogEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("date"))
	return rv
}

// A description of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/errorComment
func (m_ MovieErrorLogEvent) ErrorComment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("errorComment"))
	return rv
}

// The network domain of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/errorDomain
func (m_ MovieErrorLogEvent) ErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("errorDomain"))
	return rv
}

// A unique error code identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/errorStatusCode
func (m_ MovieErrorLogEvent) ErrorStatusCode() int {
	rv := objc.Send[int](m_.ID, objc.Sel("errorStatusCode"))
	return rv
}

// A globally unique identifier (GUID) for the playback session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/playbackSessionID
func (m_ MovieErrorLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("playbackSessionID"))
	return rv
}

// The IP address of the web server that was the source of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/serverAddress
func (m_ MovieErrorLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serverAddress"))
	return rv
}

// The URI of the item playing when the error occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent/uri
func (m_ MovieErrorLogEvent) URI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("URI"))
	return rv
}





