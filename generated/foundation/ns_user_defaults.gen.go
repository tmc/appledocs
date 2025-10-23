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
	UserDefaultsClass     _UserDefaultsClass
	UserDefaultsClassOnce sync.Once
)

func getUserDefaultsClass() _UserDefaultsClass {
	UserDefaultsClassOnce.Do(func() {
		UserDefaultsClass = _UserDefaultsClass{objc.GetClass("NSUserDefaults")}
	})
	return UserDefaultsClass
}

type _UserDefaultsClass struct {
	class objc.Class
}

// An interface definition for the [UserDefaults] class.
type IUserDefaults interface {
	objectivec.IObject
	// properties:
	VolatileDomainNames() string /* primitive/slice/pointer */
	SetVolatileDomainNames(value string /* primitive/slice/pointer */)
	// methods:
}

// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app.
//
// The class provides a programmatic interface for interacting with the defaults system. The defaults system allows an app to customize its behavior to match a user’s preferences. For example, you can allow users to specify their preferred units of measurement or media playback speed. Apps store these preferences by assigning values to a set of parameters in a user’s defaults database. The parameters are referred to as because they’re commonly used to determine an app’s default state at startup or the way it acts by default. At runtime, you use objects to read the defaults that your app uses from a user’s defaults database. caches the information to avoid having to open the user’s defaults database each time you need a default value. When you set a default value, it’s changed synchronously within your process, and asynchronously to persistent storage and other processes. With the exception of managed devices in educational institutions, a user’s defaults are stored locally on a single device, and persisted for backup and restore. To synchronize preferences and other data across a user’s connected devices, use instead.


// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app.
//
// [Full Topic]
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



// The current volatile domain names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/userdefaults/volatiledomainnames
func (u_ UserDefaults) VolatileDomainNames() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("volatileDomainNames"))
	return rv
}


// The current volatile domain names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/userdefaults/volatiledomainnames
func (u_ UserDefaults) SetVolatileDomainNames(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setVolatileDomainNames:"), objc.String(value))
}



