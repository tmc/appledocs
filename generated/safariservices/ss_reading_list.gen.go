// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SSReadingList] class.
var (
	SSReadingListClass     _SSReadingListClass
	SSReadingListClassOnce sync.Once
)

func getSSReadingListClass() _SSReadingListClass {
	SSReadingListClassOnce.Do(func() {
		SSReadingListClass = _SSReadingListClass{objc.GetClass("SSReadingList")}
	})
	return SSReadingListClass
}

type _SSReadingListClass struct {
	class objc.Class
}

// An interface definition for the [SSReadingList] class.
type ISSReadingList interface {
	objectivec.IObject
	// properties:
	SSReadingListErrorDomain() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An object for adding items to a user’s Safari Reading List.


// An object for adding items to a user’s Safari Reading List.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList
type SSReadingList struct {
	objectivec.Object
}

// SSReadingListFrom constructs a [SSReadingList] from an unsafe.Pointer.
//
// An object for adding items to a user’s Safari Reading List.
func SSReadingListFrom(ptr unsafe.Pointer) SSReadingList {
	return SSReadingList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SSReadingListClass) Alloc() SSReadingList {
	rv := objc.Send[SSReadingList](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SSReadingListClass) New() SSReadingList {
	rv := objc.Send[SSReadingList](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SSReadingList) Init() SSReadingList {
	rv := objc.Send[SSReadingList](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SSReadingList) Autorelease() SSReadingList {
	rv := objc.Send[SSReadingList](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSSReadingList creates a new SSReadingList instance.
func NewSSReadingList() SSReadingList {
	return getSSReadingListClass().New()
}



// Returns the Safari Reading List singleton object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList/default()
func (sc _SSReadingListClass) DefaultReadingList() SSReadingList {
	rv := objc.Send[SSReadingList](objc.ID(sc.class), objc.Sel("defaultReadingList"))
	return rv
}


// Determines whether a URL can be added to the Reading List.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList/supportsURL(_:)
func (sc _SSReadingListClass) SupportsURL(URL objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("supportsURL:"), URL)
	return rv
}


// The domain for Safari Reading List errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/ssreadinglisterrordomain
func (s_ SSReadingList) SSReadingListErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SSReadingListErrorDomain"))
	return rv
}


