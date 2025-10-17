// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserActivity] class.
var UserActivityClass objc.Class

func init() {
	UserActivityClass = objc.GetClass("NSUserActivity")
}

type UserActivity struct {
	objc.ID
}

func UserActivityFrom(ptr unsafe.Pointer) UserActivity {
	return UserActivity{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc UserActivity) Alloc() UserActivity {
	ret := objc.ID(UserActivityClass).Send(objc.RegisterName("alloc"))
	return UserActivity{ret}
}

// Init initializes the instance.
func (u_ UserActivity) Init() UserActivity {
	ret := u_.ID.Send(objc.RegisterName("init"))
	return UserActivity{ret}
}
// Creates a user activity object using the first activity type declared in the app’s information property list file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/init()
func NewUserActivity() UserActivity {
	instance := UserActivity{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a user activity object with the specified type. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/init(activityType:)
func NewUserActivityWithActivityType(activityType string) UserActivity {
	instance := UserActivity{}.Alloc()
	sel := objc.RegisterName("initWithActivityType:")
	ret := instance.ID.Send(sel, activityType)
	instance = UserActivity{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Deletes all user activities created by your app. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/deleteAllSavedUserActivities(completionHandler:)
func (uc UserActivity) DeleteAllSavedUserActivitiesWithCompletionHandler(handler unsafe.Pointer) {
	sel := objc.RegisterName("deleteAllSavedUserActivitiesWithCompletionHandler:")
	objc.ID(UserActivityClass).Send(sel, handler)
}
// Deletes user activities created by your app that have the specified persistent identifiers. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/deleteSavedUserActivities(withPersistentIdentifiers:completionHandler:)
func (uc UserActivity) DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers unsafe.Pointer, handler unsafe.Pointer) {
	sel := objc.RegisterName("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:")
	objc.ID(UserActivityClass).Send(sel, persistentIdentifiers, handler)
}
// Adds the contents of the specified dictionary to the user info dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/addUserInfoEntries(from:)
func (u_ UserActivity) AddUserInfoEntriesFromDictionary(otherDictionary unsafe.Pointer) {
	sel := objc.RegisterName("addUserInfoEntriesFromDictionary:")
	u_.ID.Send(sel, otherDictionary)
}
// Marks the activity as currently in use by the user. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/becomeCurrent()
func (u_ UserActivity) BecomeCurrent() {
	sel := objc.RegisterName("becomeCurrent")
	u_.ID.Send(sel)
}
// Requests streams back to the originating app. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/getContinuationStreams(completionHandler:)
func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("getContinuationStreamsWithCompletionHandler:")
	u_.ID.Send(sel, completionHandler)
}
// Invalidates an activity and marks it as no longer eligible for continuation. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/invalidate()
func (u_ UserActivity) Invalidate() {
	sel := objc.RegisterName("invalidate")
	u_.ID.Send(sel)
}
// Marks this activity object as inactive without invalidating it. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSUserActivity/resignCurrent()
func (u_ UserActivity) ResignCurrent() {
	sel := objc.RegisterName("resignCurrent")
	u_.ID.Send(sel)
}

