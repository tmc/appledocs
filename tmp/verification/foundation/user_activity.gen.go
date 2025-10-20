// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var userActivityClass _UserActivityClass

func init() {
	userActivityClass = _UserActivityClass{objc.GetClass("NSUserActivity")}
}

type _UserActivityClass struct {
	class objc.Class
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
func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.Send[UserActivity](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _UserActivityClass) New() UserActivity {
	rv := objc.Send[UserActivity](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserActivity) Init() UserActivity {
	rv := objc.Send[UserActivity](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserActivity) Autorelease() UserActivity {
	rv := objc.Send[UserActivity](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserActivity creates a new UserActivity instance.
func NewUserActivity() UserActivity {
	return userActivityClass.New()
}
// Creates a user activity object with the specified type. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/init(activityType:)
func NewUserActivityWithActivityType(activityType string) UserActivity {
	instance := userActivityClass.Alloc()
	rv := objc.Send[UserActivity](instance.ID, objc.Sel("initWithActivityType:"), activityType)
	rv.Autorelease()
	return rv
}


// Deletes all user activities created by your app. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/deleteAllSavedUserActivities(completionHandler:)
func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("deleteAllSavedUserActivitiesWithCompletionHandler:"), handler)
}
// Deletes user activities created by your app that have the specified persistent identifiers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/deleteSavedUserActivities(withPersistentIdentifiers:completionHandler:)
func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:"), persistentIdentifiers, handler)
}
// Adds the contents of the specified dictionary to the user info dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/addUserInfoEntries(from:)
func (u_ UserActivity) AddUserInfoEntriesFromDictionary(otherDictionary unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addUserInfoEntriesFromDictionary:"), otherDictionary)
}
// Marks the activity as currently in use by the user. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/becomeCurrent()
func (u_ UserActivity) BecomeCurrent() {
	objc.Send[objc.ID](u_.ID, objc.Sel("becomeCurrent"))
}
// Requests streams back to the originating app. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/getContinuationStreams(completionHandler:)
func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getContinuationStreamsWithCompletionHandler:"), completionHandler)
}
// Invalidates an activity and marks it as no longer eligible for continuation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/invalidate()
func (u_ UserActivity) Invalidate() {
	objc.Send[objc.ID](u_.ID, objc.Sel("invalidate"))
}
// Marks this activity object as inactive without invalidating it. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/resignCurrent()
func (u_ UserActivity) ResignCurrent() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resignCurrent"))
}


