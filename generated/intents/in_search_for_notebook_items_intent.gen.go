// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForNotebookItemsIntent] class.
var (
	INSearchForNotebookItemsIntentClass     _INSearchForNotebookItemsIntentClass
	INSearchForNotebookItemsIntentClassOnce sync.Once
)

func getINSearchForNotebookItemsIntentClass() _INSearchForNotebookItemsIntentClass {
	INSearchForNotebookItemsIntentClassOnce.Do(func() {
		INSearchForNotebookItemsIntentClass = _INSearchForNotebookItemsIntentClass{objc.GetClass("INSearchForNotebookItemsIntent")}
	})
	return INSearchForNotebookItemsIntentClass
}

type _INSearchForNotebookItemsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForNotebookItemsIntent] class.
type IINSearchForNotebookItemsIntent interface {
	IINIntent
}

// A request to search for notes, tasks, and reminders.
//
// Siri creates an object when the user asks to search for existing notes, tasks, and reminders. The intent object contains search parameters such as the type of items to return, strings to match against the title or content, the completion status of tasks, or the trigger conditions used to generate reminders. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the search results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForNotebookItemsIntent
type INSearchForNotebookItemsIntent struct {
	INIntent
}

// INSearchForNotebookItemsIntentFrom constructs a [INSearchForNotebookItemsIntent] from an unsafe.Pointer.
//
// A request to search for notes, tasks, and reminders.
func INSearchForNotebookItemsIntentFrom(ptr unsafe.Pointer) INSearchForNotebookItemsIntent {
	return INSearchForNotebookItemsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForNotebookItemsIntentClass) Alloc() INSearchForNotebookItemsIntent {
	rv := objc.Send[INSearchForNotebookItemsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForNotebookItemsIntentClass) New() INSearchForNotebookItemsIntent {
	rv := objc.Send[INSearchForNotebookItemsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForNotebookItemsIntent) Init() INSearchForNotebookItemsIntent {
	rv := objc.Send[INSearchForNotebookItemsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForNotebookItemsIntent) Autorelease() INSearchForNotebookItemsIntent {
	rv := objc.Send[INSearchForNotebookItemsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForNotebookItemsIntent creates a new INSearchForNotebookItemsIntent instance.
func NewINSearchForNotebookItemsIntent() INSearchForNotebookItemsIntent {
	return getINSearchForNotebookItemsIntentClass().New()
}


// An indicator of how to apply location values to your search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/locationsearchtype
func (i_ INSearchForNotebookItemsIntent) LocationSearchType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("locationSearchType"))
	return rv
}


// SetLocationSearchType sets the value of the locationSearchType property.
// An indicator of how to apply location values to your search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/locationsearchtype
func (i_ INSearchForNotebookItemsIntent) SetLocationSearchType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocationSearchType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/temporaleventtriggertypes
func (i_ INSearchForNotebookItemsIntent) TemporalEventTriggerTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("temporalEventTriggerTypes"))
	return rv
}


// SetTemporalEventTriggerTypes sets the value of the temporalEventTriggerTypes property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/temporaleventtriggertypes
func (i_ INSearchForNotebookItemsIntent) SetTemporalEventTriggerTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemporalEventTriggerTypes:"), value)
}

// An indicator of how to apply date values to your search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/datesearchtype
func (i_ INSearchForNotebookItemsIntent) DateSearchType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dateSearchType"))
	return rv
}


// SetDateSearchType sets the value of the dateSearchType property.
// An indicator of how to apply date values to your search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/datesearchtype
func (i_ INSearchForNotebookItemsIntent) SetDateSearchType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateSearchType:"), value)
}

// The value to use when searching for location-triggered reminders.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/location
func (i_ INSearchForNotebookItemsIntent) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The value to use when searching for location-triggered reminders.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/location
func (i_ INSearchForNotebookItemsIntent) SetLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocation:"), value)
}

// The value to use when performing date-based searches.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/datetime
func (i_ INSearchForNotebookItemsIntent) DateTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dateTime"))
	return rv
}


// SetDateTime sets the value of the dateTime property.
// The value to use when performing date-based searches.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/datetime
func (i_ INSearchForNotebookItemsIntent) SetDateTime(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/taskpriority
func (i_ INSearchForNotebookItemsIntent) TaskPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("taskPriority"))
	return rv
}


// SetTaskPriority sets the value of the taskPriority property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/taskpriority
func (i_ INSearchForNotebookItemsIntent) SetTaskPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTaskPriority:"), value)
}

// The type of items to include in your search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/itemtype
func (i_ INSearchForNotebookItemsIntent) ItemType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("itemType"))
	return rv
}


// SetItemType sets the value of the itemType property.
// The type of items to include in your search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/itemtype
func (i_ INSearchForNotebookItemsIntent) SetItemType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setItemType:"), value)
}

// The unique identifier that your app assigned to the item.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/notebookitemidentifier
func (i_ INSearchForNotebookItemsIntent) NotebookItemIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("notebookItemIdentifier"))
	return rv
}


// SetNotebookItemIdentifier sets the value of the notebookItemIdentifier property.
// The unique identifier that your app assigned to the item.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/notebookitemidentifier
func (i_ INSearchForNotebookItemsIntent) SetNotebookItemIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNotebookItemIdentifier:"), objc.String(value))
}

// The title text to search for in a note, task, or task list.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/title
func (i_ INSearchForNotebookItemsIntent) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title text to search for in a note, task, or task list.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/title
func (i_ INSearchForNotebookItemsIntent) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTitle:"), value)
}

// The completion state to look for when searching for tasks.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/status
func (i_ INSearchForNotebookItemsIntent) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The completion state to look for when searching for tasks.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/status
func (i_ INSearchForNotebookItemsIntent) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStatus:"), value)
}

// The text to search for in the body of a note.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/content
func (i_ INSearchForNotebookItemsIntent) Content() string {
	rv := objc.Send[string](i_.ID, objc.Sel("content"))
	return rv
}


// SetContent sets the value of the content property.
// The text to search for in the body of a note.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchfornotebookitemsintent/content
func (i_ INSearchForNotebookItemsIntent) SetContent(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContent:"), objc.String(value))
}



