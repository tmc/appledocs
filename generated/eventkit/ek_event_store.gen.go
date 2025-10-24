// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKEventStore */


/* debug [class_header]: Header for EKEventStore */
// The class instance for the [EKEventStore] class.
var (
	EKEventStoreClass     _EKEventStoreClass
	EKEventStoreClassOnce sync.Once
)

func getEKEventStoreClass() _EKEventStoreClass {
	EKEventStoreClassOnce.Do(func() {
		EKEventStoreClass = _EKEventStoreClass{objc.GetClass("EKEventStore")}
	})
	return EKEventStoreClass
}

type _EKEventStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKEventStore */
// An interface definition for the [EKEventStore] class.
type IEKEventStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EKEventStore */
	// properties:
	DefaultCalendarForNewEvents() IEKCalendar
	DelegateSources() []EKSource
	EventStoreIdentifier() objc.IObject /* cross-framework: NSString */
	Sources() []EKSource
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKEventStore */
	// methods:
	CalendarWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKCalendar
	CalendarItemWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKCalendarItem
	CalendarItemsWithExternalIdentifier(externalIdentifier objc.IObject /* cross-framework: NSString */) []EKCalendarItem
	CalendarsForEntityType(entityType EKEntityType) []EKCalendar
	CancelFetchRequest(fetchIdentifier objc.IObject)
	Commit(error_ objectivec.IObject) bool
	DefaultCalendarForNewReminders() IEKCalendar
	EnumerateEventsMatchingPredicateUsingBlock(predicate foundation.Predicate, block objectivec.IObject)
	EventWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKEvent
	EventsMatchingPredicate(predicate foundation.Predicate) []EKEvent
	FetchRemindersMatchingPredicateCompletion(predicate foundation.Predicate, completion unsafe.Pointer) objc.ID
	PredicateForCompletedRemindersWithCompletionDateStartingEndingCalendars(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, calendars []EKCalendar) foundation.Predicate
	PredicateForEventsWithStartDateEndDateCalendars(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, calendars []EKCalendar) foundation.Predicate
	PredicateForIncompleteRemindersWithDueDateStartingEndingCalendars(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, calendars []EKCalendar) foundation.Predicate
	PredicateForRemindersInCalendars(calendars []EKCalendar) foundation.Predicate
	RefreshSourcesIfNecessary()
	RemoveReminderCommitError(reminder IEKReminder, commit bool, error_ objectivec.IObject) bool
	RemoveEventSpanError(event IEKEvent, span EKSpan, error_ objectivec.IObject) bool
	RemoveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool, error_ objectivec.IObject) bool
	RemoveCalendarCommitError(calendar IEKCalendar, commit bool, error_ objectivec.IObject) bool
	RequestFullAccessToEventsWithCompletion(completion objectivec.IObject)
	RequestFullAccessToRemindersWithCompletion(completion objectivec.IObject)
	RequestWriteOnlyAccessToEventsWithCompletion(completion objectivec.IObject)
	Reset()
	SaveReminderCommitError(reminder IEKReminder, commit bool, error_ objectivec.IObject) bool
	SaveEventSpanError(event IEKEvent, span EKSpan, error_ objectivec.IObject) bool
	SaveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool, error_ objectivec.IObject) bool
	SaveCalendarCommitError(calendar IEKCalendar, commit bool, error_ objectivec.IObject) bool
	SourceWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKSource
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKEventStore */
// Alloc allocates a new instance without initialization.
func (ec _EKEventStoreClass) Alloc() EKEventStore {
	rv := objc.Send[EKEventStore](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKEventStoreClass) New() EKEventStore {
	rv := objc.Send[EKEventStore](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKEventStore) Init() EKEventStore {
	rv := objc.Send[EKEventStore](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKEventStore) Autorelease() EKEventStore {
	rv := objc.Send[EKEventStore](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKEventStore creates a new EKEventStore instance.
func NewEKEventStore() EKEventStore {
	return getEKEventStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKEventStore */
// An object that accesses a person’s calendar events and reminders and supports the scheduling of new events.
//
// The class is an app’s point of contact for accessing calendar and reminder data. After initializing the event store, you must request access to events or reminders before attempting to fetch or create data. To request access to reminders, call . To request access to events, call or . A typical workflow for using an event store is: Create a predicate, or a search query for events, with . Fetch and process events that match the predicate with the and methods. Save and delete events from the event store with the and methods. Use similar methods to access and manipulate reminders. After receiving an object from an event store, don’t use that object with a different event store. This restriction applies to subclasses such as , , , and , as well as predicates that the event store creates. For example, don’t fetch an event from one event store, modify the event, and then pass it to in a different store.


// An object that accesses a person’s calendar events and reminders and supports the scheduling of new events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore
type EKEventStore struct {
	objectivec.Object
}

// EKEventStoreFrom constructs a [EKEventStore] from an unsafe.Pointer.
//
// An object that accesses a person’s calendar events and reminders and supports the scheduling of new events.
func EKEventStoreFrom(ptr unsafe.Pointer) EKEventStore {
	return EKEventStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKEventStore */

// Initializes access to the event store with support for the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/initWithAccessToEntityTypes:
func NewEKEventStoreWithAccessToEntityTypes(entityTypes EKEntityMask) EKEventStore {
	instance := getEKEventStoreClass().Alloc()
	rv := objc.Send[EKEventStore](instance.ID, objc.Sel("initWithAccessToEntityTypes:"), entityTypes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKEventStoreWithAccessToEntityTypes */


// Creates an event store that contains data for the specified sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/init(sources:)
func NewEKEventStoreWithSources(sources []EKSource) EKEventStore {
	instance := getEKEventStoreClass().Alloc()
	rv := objc.Send[EKEventStore](instance.ID, objc.Sel("initWithSources:"), sources)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKEventStoreWithSources */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKEventStore */

// Determines the authorization status for the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/authorizationStatus(for:)
func (ec _EKEventStoreClass) AuthorizationStatusForEntityType(entityType EKEntityType) EKAuthorizationStatus {
	rv := objc.Send[EKAuthorizationStatus](objc.ID(ec.class), objc.Sel("authorizationStatusForEntityType:"), entityType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatusForEntityType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKEventStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKEventStore */

// Locates a calendar with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendar(withIdentifier:)
func (e_ EKEventStore) CalendarWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("calendarWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: CalendarWithIdentifier */


// Locates a reminder or the first occurrence of an event with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendarItem(withIdentifier:)
func (e_ EKEventStore) CalendarItemWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKCalendarItem {
	rv := objc.Send[EKCalendarItem](e_.ID, objc.Sel("calendarItemWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: CalendarItemWithIdentifier */


// Locates all reminders or the first occurrences of all events with the specified external identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendarItems(withExternalIdentifier:)
func (e_ EKEventStore) CalendarItemsWithExternalIdentifier(externalIdentifier objc.IObject /* cross-framework: NSString */) []EKCalendarItem {
	rv := objc.Send[[]EKCalendarItem](e_.ID, objc.Sel("calendarItemsWithExternalIdentifier:"), externalIdentifier)
	return rv
}/* debug [instance_methods/method]: CalendarItemsWithExternalIdentifier */


// Identifies the calendars that support a given entity type, such as reminders or events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendars(for:)
func (e_ EKEventStore) CalendarsForEntityType(entityType EKEntityType) []EKCalendar {
	rv := objc.Send[[]EKCalendar](e_.ID, objc.Sel("calendarsForEntityType:"), entityType)
	return rv
}/* debug [instance_methods/method]: CalendarsForEntityType */


// Cancels the request to fetch reminders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/cancelFetchRequest(_:)
func (e_ EKEventStore) CancelFetchRequest(fetchIdentifier objc.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("cancelFetchRequest:"), fetchIdentifier)
}/* debug [instance_methods/method]: CancelFetchRequest */


// Commits all unsaved changes to the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/commit()
func (e_ EKEventStore) Commit(error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("commit:"), error_)
	return rv
}/* debug [instance_methods/method]: Commit */


// Identifies the default calendar for adding reminders to, as specified by user settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/defaultCalendarForNewReminders()
func (e_ EKEventStore) DefaultCalendarForNewReminders() IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("defaultCalendarForNewReminders"))
	return rv
}/* debug [instance_methods/method]: DefaultCalendarForNewReminders */


// Finds all events that match a given predicate and calls a given callback for each event found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/enumerateEvents(matching:using:)
func (e_ EKEventStore) EnumerateEventsMatchingPredicateUsingBlock(predicate foundation.Predicate, block objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateEventsMatchingPredicate:usingBlock:"), predicate, block)
}/* debug [instance_methods/method]: EnumerateEventsMatchingPredicateUsingBlock */


// Locates the first occurrence of an event with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/event(withIdentifier:)
func (e_ EKEventStore) EventWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKEvent {
	rv := objc.Send[EKEvent](e_.ID, objc.Sel("eventWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: EventWithIdentifier */


// Finds all events that match a given predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/events(matching:)
func (e_ EKEventStore) EventsMatchingPredicate(predicate foundation.Predicate) []EKEvent {
	rv := objc.Send[[]EKEvent](e_.ID, objc.Sel("eventsMatchingPredicate:"), predicate)
	return rv
}/* debug [instance_methods/method]: EventsMatchingPredicate */


// Fetches reminders that match a given predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/fetchReminders(matching:completion:)
func (e_ EKEventStore) FetchRemindersMatchingPredicateCompletion(predicate foundation.Predicate, completion unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("fetchRemindersMatchingPredicate:completion:"), predicate, completion)
	return rv
}/* debug [instance_methods/method]: FetchRemindersMatchingPredicateCompletion */


// Creates a predicate to identify all completed reminders that occur within a given date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForCompletedReminders(withCompletionDateStarting:ending:calendars:)
func (e_ EKEventStore) PredicateForCompletedRemindersWithCompletionDateStartingEndingCalendars(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, calendars []EKCalendar) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](e_.ID, objc.Sel("predicateForCompletedRemindersWithCompletionDateStarting:ending:calendars:"), startDate, endDate, calendars)
	return rv
}/* debug [instance_methods/method]: PredicateForCompletedRemindersWithCompletionDateStartingEndingCalendars */


// Creates a predicate to identify events that occur within a given date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForEvents(withStart:end:calendars:)
func (e_ EKEventStore) PredicateForEventsWithStartDateEndDateCalendars(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, calendars []EKCalendar) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](e_.ID, objc.Sel("predicateForEventsWithStartDate:endDate:calendars:"), startDate, endDate, calendars)
	return rv
}/* debug [instance_methods/method]: PredicateForEventsWithStartDateEndDateCalendars */


// Creates a predicate to identify all incomplete reminders that occur within a given date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForIncompleteReminders(withDueDateStarting:ending:calendars:)
func (e_ EKEventStore) PredicateForIncompleteRemindersWithDueDateStartingEndingCalendars(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, calendars []EKCalendar) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](e_.ID, objc.Sel("predicateForIncompleteRemindersWithDueDateStarting:ending:calendars:"), startDate, endDate, calendars)
	return rv
}/* debug [instance_methods/method]: PredicateForIncompleteRemindersWithDueDateStartingEndingCalendars */


// Creates a predicate to identify all reminders in a collection of calendars.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForReminders(in:)
func (e_ EKEventStore) PredicateForRemindersInCalendars(calendars []EKCalendar) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](e_.ID, objc.Sel("predicateForRemindersInCalendars:"), calendars)
	return rv
}/* debug [instance_methods/method]: PredicateForRemindersInCalendars */


// Pulls new data from remote sources, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/refreshSourcesIfNecessary()
func (e_ EKEventStore) RefreshSourcesIfNecessary() {
	objc.Send[objc.ID](e_.ID, objc.Sel("refreshSourcesIfNecessary"))
}/* debug [instance_methods/method]: RefreshSourcesIfNecessary */


// Removes a reminder from the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/remove(_:commit:)
func (e_ EKEventStore) RemoveReminderCommitError(reminder IEKReminder, commit bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeReminder:commit:error:"), reminder, commit, error_)
	return rv
}/* debug [instance_methods/method]: RemoveReminderCommitError */


// Removes an event from the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/remove(_:span:)
func (e_ EKEventStore) RemoveEventSpanError(event IEKEvent, span EKSpan, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeEvent:span:error:"), event, span, error_)
	return rv
}/* debug [instance_methods/method]: RemoveEventSpanError */


// Removes an event or recurring events from the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/remove(_:span:commit:)
func (e_ EKEventStore) RemoveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeEvent:span:commit:error:"), event, span, commit, error_)
	return rv
}/* debug [instance_methods/method]: RemoveEventSpanCommitError */


// Removes a calendar from the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/removeCalendar(_:commit:)
func (e_ EKEventStore) RemoveCalendarCommitError(calendar IEKCalendar, commit bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeCalendar:commit:error:"), calendar, commit, error_)
	return rv
}/* debug [instance_methods/method]: RemoveCalendarCommitError */


// Prompts people to grant or deny read and write access to event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/requestFullAccessToEvents(completion:)
func (e_ EKEventStore) RequestFullAccessToEventsWithCompletion(completion objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestFullAccessToEventsWithCompletion:"), completion)
}/* debug [instance_methods/method]: RequestFullAccessToEventsWithCompletion */


// Prompts people to grant or deny read and write access to reminders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/requestFullAccessToReminders(completion:)
func (e_ EKEventStore) RequestFullAccessToRemindersWithCompletion(completion objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestFullAccessToRemindersWithCompletion:"), completion)
}/* debug [instance_methods/method]: RequestFullAccessToRemindersWithCompletion */


// Prompts the person using your app to grant or deny write access to event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/requestWriteOnlyAccessToEvents(completion:)
func (e_ EKEventStore) RequestWriteOnlyAccessToEventsWithCompletion(completion objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestWriteOnlyAccessToEventsWithCompletion:"), completion)
}/* debug [instance_methods/method]: RequestWriteOnlyAccessToEventsWithCompletion */


// Reverts the event store to its saved state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/reset()
func (e_ EKEventStore) Reset() {
	objc.Send[objc.ID](e_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Saves changes to a reminder by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/save(_:commit:)
func (e_ EKEventStore) SaveReminderCommitError(reminder IEKReminder, commit bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveReminder:commit:error:"), reminder, commit, error_)
	return rv
}/* debug [instance_methods/method]: SaveReminderCommitError */


// Saves changes to an event permanently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/save(_:span:)
func (e_ EKEventStore) SaveEventSpanError(event IEKEvent, span EKSpan, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveEvent:span:error:"), event, span, error_)
	return rv
}/* debug [instance_methods/method]: SaveEventSpanError */


// Saves an event or recurring events to the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/save(_:span:commit:)
func (e_ EKEventStore) SaveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveEvent:span:commit:error:"), event, span, commit, error_)
	return rv
}/* debug [instance_methods/method]: SaveEventSpanCommitError */


// Saves a calendar to the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/saveCalendar(_:commit:)
func (e_ EKEventStore) SaveCalendarCommitError(calendar IEKCalendar, commit bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveCalendar:commit:error:"), calendar, commit, error_)
	return rv
}/* debug [instance_methods/method]: SaveCalendarCommitError */


// Locates an event source with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/source(withIdentifier:)
func (e_ EKEventStore) SourceWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) IEKSource {
	rv := objc.Send[EKSource](e_.ID, objc.Sel("sourceWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: SourceWithIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKEventStore */

// The calendar that events are added to by default, as specified by user settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/defaultCalendarForNewEvents
func (e_ EKEventStore) DefaultCalendarForNewEvents() IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("defaultCalendarForNewEvents"))
	return rv
}/* debug [instance_properties/getter]: defaultCalendarForNewEvents */


// The event sources delegated to the person using your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/delegateSources
func (e_ EKEventStore) DelegateSources() []EKSource {
	rv := objc.Send[[]EKSource](e_.ID, objc.Sel("delegateSources"))
	return rv
}/* debug [instance_properties/getter]: delegateSources */


// The unique identifier for the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/eventStoreIdentifier
func (e_ EKEventStore) EventStoreIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("eventStoreIdentifier"))
	return rv
}/* debug [instance_properties/getter]: eventStoreIdentifier */


// An unordered array of objects that represent accounts that contain calendars.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/sources
func (e_ EKEventStore) Sources() []EKSource {
	rv := objc.Send[[]EKSource](e_.ID, objc.Sel("sources"))
	return rv
}/* debug [instance_properties/getter]: sources */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKEventStore */


