// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNSocialProfile] class.
var (
	CNSocialProfileClass     _CNSocialProfileClass
	CNSocialProfileClassOnce sync.Once
)

func getCNSocialProfileClass() _CNSocialProfileClass {
	CNSocialProfileClassOnce.Do(func() {
		CNSocialProfileClass = _CNSocialProfileClass{objc.GetClass("CNSocialProfile")}
	})
	return CNSocialProfileClass
}

type _CNSocialProfileClass struct {
	class objc.Class
}

// An interface definition for the [CNSocialProfile] class.
type ICNSocialProfile interface {
	objectivec.IObject
}

// An immutable object that represents one of the user’s social profiles.
//
// Some social profile services, such as Facebook and Twitter, are predefined in this class. You can also specify your own social profile service with the method. objects are thread-safe, and you may access their properties from any thread of your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile
type CNSocialProfile struct {
	objectivec.Object
}

// CNSocialProfileFrom constructs a [CNSocialProfile] from an unsafe.Pointer.
//
// An immutable object that represents one of the user’s social profiles.
func CNSocialProfileFrom(ptr unsafe.Pointer) CNSocialProfile {
	return CNSocialProfile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNSocialProfileClass) Alloc() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNSocialProfileClass) New() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNSocialProfile) Init() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNSocialProfile) Autorelease() CNSocialProfile {
	rv := objc.Send[CNSocialProfile](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNSocialProfile creates a new CNSocialProfile instance.
func NewCNSocialProfile() CNSocialProfile {
	return getCNSocialProfileClass().New()
}


// Returns the localized name of the property for the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSocialProfile/localizedString(forKey:)
func (cc _CNSocialProfileClass) LocalizedStringForKey(key string) string {
	rv := objc.Send[string](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), objc.String(key))
	return rv
}



