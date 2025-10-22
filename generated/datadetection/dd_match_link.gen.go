// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DDMatchLink] class.
var (
	DDMatchLinkClass     _DDMatchLinkClass
	DDMatchLinkClassOnce sync.Once
)

func getDDMatchLinkClass() _DDMatchLinkClass {
	DDMatchLinkClassOnce.Do(func() {
		DDMatchLinkClass = _DDMatchLinkClass{objc.GetClass("DDMatchLink")}
	})
	return DDMatchLinkClass
}

type _DDMatchLinkClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchLink] class.
type IDDMatchLink interface {
	IDDMatch
	URL() foundation.URL
}

// An object that contains a web link that the data detection system matches.
//
// The DataDetection framework returns a link match in a object, which contains a .


// An object that contains a web link that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchLink

type DDMatchLink struct {
	DDMatch
}

// DDMatchLinkFrom constructs a [DDMatchLink] from an unsafe.Pointer.
//
// An object that contains a web link that the data detection system matches.
func DDMatchLinkFrom(ptr unsafe.Pointer) DDMatchLink {
	return DDMatchLink{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchLinkClass) Alloc() DDMatchLink {
	rv := objc.Send[DDMatchLink](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchLinkClass) New() DDMatchLink {
	rv := objc.Send[DDMatchLink](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchLink) Init() DDMatchLink {
	rv := objc.Send[DDMatchLink](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchLink) Autorelease() DDMatchLink {
	rv := objc.Send[DDMatchLink](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchLink creates a new DDMatchLink instance.
func NewDDMatchLink() DDMatchLink {
	return getDDMatchLinkClass().New()
}



// An address for a web resource, such as a webpage or image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchLink/url

func (d_ DDMatchLink) URL() foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("URL"))
	return rv
}



