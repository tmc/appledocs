// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserDefaults] class.
var (
	userDefaultsClass     _UserDefaultsClass
	userDefaultsClassOnce sync.Once
)

func getUserDefaultsClass() _UserDefaultsClass {
	userDefaultsClassOnce.Do(func() {
		userDefaultsClass = _UserDefaultsClass{objc.GetClass("NSUserDefaults")}
	})
	return userDefaultsClass
}

type _UserDefaultsClass struct {
	class objc.Class
}

// An interface definition for the [UserDefaults] class.
type IUserDefaults interface {
	objectivec.IObject
	SetURLForKey(url unsafe.Pointer, defaultName string)
	SetObjectForKey(value objc.ID, defaultName string)
	StringForKey(defaultName string) unsafe.Pointer
}

// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults
type UserDefaults struct {
	objectivec.Object
}

// UserDefaultsFrom constructs a [UserDefaults] from an unsafe.Pointer.
//
// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app.
func UserDefaultsFrom(ptr unsafe.Pointer) UserDefaults {
	return UserDefaults{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserDefaultsClass) Alloc() UserDefaults {
	rv := objc.Send[UserDefaults](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserDefaultsClass) New() UserDefaults {
	rv := objc.Send[UserDefaults](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserDefaults) Init() UserDefaults {
	rv := objc.Send[UserDefaults](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserDefaults) Autorelease() UserDefaults {
	rv := objc.Send[UserDefaults](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserDefaults creates a new UserDefaults instance.
func NewUserDefaults() UserDefaults {
	return getUserDefaultsClass().New()
}


// Creates a user defaults object initialized with the defaults for the specified database name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/init(suiteName:)
func NewUserDefaultsWithSuiteName(suitename string) UserDefaults {
	instance := getUserDefaultsClass().Alloc()
	rv := objc.Send[UserDefaults](instance.ID, objc.Sel("initWithSuiteName:"), objc.String(suitename))
	rv.Autorelease()
	return rv
}


// Sets the value of the specified default key to the specified URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-2bqjt
func (u_ UserDefaults) SetURLForKey(url unsafe.Pointer, defaultName string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setURL:forKey:"), url, objc.String(defaultName))
}
// Sets the value of the specified default key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-8ab6d
func (u_ UserDefaults) SetObjectForKey(value objc.ID, defaultName string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setObject:forKey:"), value, objc.String(defaultName))
}
// Returns the string associated with the specified key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UserDefaults/string(forKey:)
func (u_ UserDefaults) StringForKey(defaultName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("stringForKey:"), objc.String(defaultName))
	return rv
}

