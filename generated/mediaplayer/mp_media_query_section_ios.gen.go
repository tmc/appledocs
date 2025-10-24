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

// iOS-only methods for MediaQuerySection


// iOS-only properties

// The range in the media query’s items or collections array that the media query section represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuerySection/range
func (m_ MediaQuerySection) Range() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](m_.ID, objc.Sel("range"))
	return rv
}

// The localized title of the media query section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuerySection/title
func (m_ MediaQuerySection) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}





