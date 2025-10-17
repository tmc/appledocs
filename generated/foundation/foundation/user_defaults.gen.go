// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserDefaults] class.
var UserDefaultsClass objc.Class

func init() {
	UserDefaultsClass = objc.GetClass("NSUserDefaults")
}

type UserDefaults struct {
	objc.ID
}

func UserDefaultsFrom(ptr unsafe.Pointer) UserDefaults {
	return UserDefaults{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc UserDefaults) Alloc() UserDefaults {
	ret := objc.ID(UserDefaultsClass).Send(objc.RegisterName("alloc"))
	return UserDefaults{ret}
}

// Init initializes the instance.
func (u_ UserDefaults) Init() UserDefaults {
	ret := u_.ID.Send(objc.RegisterName("init"))
	return UserDefaults{ret}
}
// Creates a user defaults object initialized with the defaults for the specified database name. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UserDefaults/init(suiteName:)
func NewUserDefaultsWithSuiteName(suitename string) UserDefaults {
	instance := UserDefaults{}.Alloc()
	sel := objc.RegisterName("initWithSuiteName:")
	ret := instance.ID.Send(sel, suitename)
	instance = UserDefaults{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Sets the value of the specified default key to the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UserDefaults/set(_:forKey:)-2bqjt
func (u_ UserDefaults) SetURLForKey(url unsafe.Pointer, defaultName string) {
	sel := objc.RegisterName("setURL:forKey:")
	u_.ID.Send(sel, url, defaultName)
}
// Sets the value of the specified default key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UserDefaults/set(_:forKey:)-8ab6d
func (u_ UserDefaults) SetObjectForKey(value objc.ID, defaultName string) {
	sel := objc.RegisterName("setObject:forKey:")
	u_.ID.Send(sel, value, defaultName)
}
// Returns the string associated with the specified key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UserDefaults/string(forKey:)
func (u_ UserDefaults) StringForKey(defaultName string) unsafe.Pointer {
	sel := objc.RegisterName("stringForKey:")
	ret := u_.ID.Send(sel, defaultName)
	return unsafe.Pointer(ret)
}

