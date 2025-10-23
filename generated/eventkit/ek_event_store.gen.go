// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [EKEventStore] class.
type IEKEventStore interface {
	objectivec.IObject
	// properties:
	Calendars() []EKCalendar /* primitive/slice/pointer. */
	DefaultCalendarForNewEvents() IEKCalendar
	DelegateSources() []EKSource /* primitive/slice/pointer. */
	EventStoreIdentifier() string /* primitive/slice/pointer. */
	Sources() []EKSource /* primitive/slice/pointer. */
	// methods:
	CalendarWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKCalendar
	CalendarItemWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKCalendarItem
	CalendarItemsWithExternalIdentifier(externalIdentifier string /* primitive/slice/pointer. */) []EKCalendarItem /* primitive/slice/pointer. */
	CalendarsForEntityType(entityType EKEntityType) []EKCalendar /* primitive/slice/pointer. */
	CancelFetchRequest(fetchIdentifier objectivec.IObject)
	Commit(error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	DefaultCalendarForNewReminders() IEKCalendar
	EnumerateEventsMatchingPredicateUsingBlock(predicate objc.IObject /* cross-framework Predicate */, block unsafe.Pointer)
	EventWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKEvent
	EventsMatchingPredicate(predicate objc.IObject /* cross-framework Predicate */) []EKEvent /* primitive/slice/pointer. */
	FetchRemindersMatchingPredicateCompletion(predicate objc.IObject /* cross-framework Predicate */, completion unsafe.Pointer) objc.ID
	PredicateForCompletedRemindersWithCompletionDateStartingEndingCalendars(startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */
	PredicateForEventsWithStartDateEndDateCalendars(startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */
	PredicateForIncompleteRemindersWithDueDateStartingEndingCalendars(startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */
	PredicateForRemindersInCalendars(calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */
	RefreshSourcesIfNecessary()
	RemoveReminderCommitError(reminder IEKReminder, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	RemoveEventSpanError(event IEKEvent, span EKSpan, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	RemoveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	RemoveCalendarCommitError(calendar IEKCalendar, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	RequestFullAccessToEventsWithCompletion(completion unsafe.Pointer)
	RequestFullAccessToRemindersWithCompletion(completion unsafe.Pointer)
	RequestWriteOnlyAccessToEventsWithCompletion(completion unsafe.Pointer)
	Reset()
	SaveReminderCommitError(reminder IEKReminder, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	SaveEventSpanError(event IEKEvent, span EKSpan, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	SaveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	SaveCalendarCommitError(calendar IEKCalendar, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	SourceWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKSource
}

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

// Alloc allocates a new instance without initialization.
func (ec _EKEventStoreClass) Alloc() EKEventStore {
	rv := objc.Send[EKEventStore](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes access to the event store with support for the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/initWithAccessToEntityTypes:
func NewEKEventStoreWithAccessToEntityTypes(entityTypes unsafe.Pointer) EKEventStore {
	instance := getEKEventStoreClass().Alloc()
	rv := objc.Send[EKEventStore](instance.ID, objc.Sel("initWithAccessToEntityTypes:"), entityTypes)
	rv.Autorelease()
	return rv
}


// Creates an event store that contains data for the specified sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/init(sources:)
func NewEKEventStoreWithSources(sources []EKSource /* primitive/slice/pointer. */) EKEventStore {
	instance := getEKEventStoreClass().Alloc()
	rv := objc.Send[EKEventStore](instance.ID, objc.Sel("initWithSources:"), sources)
	rv.Autorelease()
	return rv
}



// Determines the authorization status for the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/authorizationStatus(for:)
func (ec _EKEventStoreClass) AuthorizationStatusForEntityType(entityType EKEntityType) EKAuthorizationStatus {
	rv := objc.Send[EKAuthorizationStatus](objc.ID(ec.class), objc.Sel("authorizationStatusForEntityType:"), entityType)
	return rv
}


// Locates a calendar with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendar(withIdentifier:)
func (e_ EKEventStore) CalendarWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("calendarWithIdentifier:"), objc.String(identifier))
	return rv
}


// Locates a reminder or the first occurrence of an event with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendarItem(withIdentifier:)
func (e_ EKEventStore) CalendarItemWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKCalendarItem {
	rv := objc.Send[EKCalendarItem](e_.ID, objc.Sel("calendarItemWithIdentifier:"), objc.String(identifier))
	return rv
}


// Locates all reminders or the first occurrences of all events with the specified external identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendarItems(withExternalIdentifier:)
func (e_ EKEventStore) CalendarItemsWithExternalIdentifier(externalIdentifier string /* primitive/slice/pointer. */) []EKCalendarItem /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKCalendarItem](e_.ID, objc.Sel("calendarItemsWithExternalIdentifier:"), objc.String(externalIdentifier))
	return rv
}


// Identifies the calendars that support a given entity type, such as reminders or events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendars(for:)
func (e_ EKEventStore) CalendarsForEntityType(entityType EKEntityType) []EKCalendar /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKCalendar](e_.ID, objc.Sel("calendarsForEntityType:"), entityType)
	return rv
}


// Cancels the request to fetch reminders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/cancelFetchRequest(_:)
func (e_ EKEventStore) CancelFetchRequest(fetchIdentifier objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("cancelFetchRequest:"), fetchIdentifier)
}


// Commits all unsaved changes to the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/commit()
func (e_ EKEventStore) Commit(error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("commit:"), error_)
	return rv
}


// Identifies the default calendar for adding reminders to, as specified by user settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/defaultCalendarForNewReminders()
func (e_ EKEventStore) DefaultCalendarForNewReminders() IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("defaultCalendarForNewReminders"))
	return rv
}


// Finds all events that match a given predicate and calls a given callback for each event found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/enumerateEvents(matching:using:)
func (e_ EKEventStore) EnumerateEventsMatchingPredicateUsingBlock(predicate objc.IObject /* cross-framework Predicate */, block unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateEventsMatchingPredicate:usingBlock:"), predicate, block)
}


// Locates the first occurrence of an event with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/event(withIdentifier:)
func (e_ EKEventStore) EventWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKEvent {
	rv := objc.Send[EKEvent](e_.ID, objc.Sel("eventWithIdentifier:"), objc.String(identifier))
	return rv
}


// Finds all events that match a given predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/events(matching:)
func (e_ EKEventStore) EventsMatchingPredicate(predicate objc.IObject /* cross-framework Predicate */) []EKEvent /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKEvent](e_.ID, objc.Sel("eventsMatchingPredicate:"), predicate)
	return rv
}


// Fetches reminders that match a given predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/fetchReminders(matching:completion:)
func (e_ EKEventStore) FetchRemindersMatchingPredicateCompletion(predicate objc.IObject /* cross-framework Predicate */, completion unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("fetchRemindersMatchingPredicate:completion:"), predicate, completion)
	return rv
}


// Creates a predicate to identify all completed reminders that occur within a given date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForCompletedReminders(withCompletionDateStarting:ending:calendars:)
func (e_ EKEventStore) PredicateForCompletedRemindersWithCompletionDateStartingEndingCalendars(startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicateForCompletedRemindersWithCompletionDateStarting:ending:calendars:"), startDate, endDate, calendars)
	return rv
}


// Creates a predicate to identify events that occur within a given date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForEvents(withStart:end:calendars:)
func (e_ EKEventStore) PredicateForEventsWithStartDateEndDateCalendars(startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicateForEventsWithStartDate:endDate:calendars:"), startDate, endDate, calendars)
	return rv
}


// Creates a predicate to identify all incomplete reminders that occur within a given date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForIncompleteReminders(withDueDateStarting:ending:calendars:)
func (e_ EKEventStore) PredicateForIncompleteRemindersWithDueDateStartingEndingCalendars(startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicateForIncompleteRemindersWithDueDateStarting:ending:calendars:"), startDate, endDate, calendars)
	return rv
}


// Creates a predicate to identify all reminders in a collection of calendars.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/predicateForReminders(in:)
func (e_ EKEventStore) PredicateForRemindersInCalendars(calendars []EKCalendar /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](e_.ID, objc.Sel("predicateForRemindersInCalendars:"), calendars)
	return rv
}


// Pulls new data from remote sources, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/refreshSourcesIfNecessary()
func (e_ EKEventStore) RefreshSourcesIfNecessary() {
	objc.Send[objc.ID](e_.ID, objc.Sel("refreshSourcesIfNecessary"))
}


// Removes a reminder from the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/remove(_:commit:)
func (e_ EKEventStore) RemoveReminderCommitError(reminder IEKReminder, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeReminder:commit:error:"), reminder, commit, error_)
	return rv
}


// Removes an event from the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/remove(_:span:)
func (e_ EKEventStore) RemoveEventSpanError(event IEKEvent, span EKSpan, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeEvent:span:error:"), event, span, error_)
	return rv
}


// Removes an event or recurring events from the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/remove(_:span:commit:)
func (e_ EKEventStore) RemoveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeEvent:span:commit:error:"), event, span, commit, error_)
	return rv
}


// Removes a calendar from the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/removeCalendar(_:commit:)
func (e_ EKEventStore) RemoveCalendarCommitError(calendar IEKCalendar, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("removeCalendar:commit:error:"), calendar, commit, error_)
	return rv
}


// Prompts people to grant or deny read and write access to event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/requestFullAccessToEvents(completion:)
func (e_ EKEventStore) RequestFullAccessToEventsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestFullAccessToEventsWithCompletion:"), completion)
}


// Prompts people to grant or deny read and write access to reminders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/requestFullAccessToReminders(completion:)
func (e_ EKEventStore) RequestFullAccessToRemindersWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestFullAccessToRemindersWithCompletion:"), completion)
}


// Prompts the person using your app to grant or deny write access to event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/requestWriteOnlyAccessToEvents(completion:)
func (e_ EKEventStore) RequestWriteOnlyAccessToEventsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("requestWriteOnlyAccessToEventsWithCompletion:"), completion)
}


// Reverts the event store to its saved state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/reset()
func (e_ EKEventStore) Reset() {
	objc.Send[objc.ID](e_.ID, objc.Sel("reset"))
}


// Saves changes to a reminder by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/save(_:commit:)
func (e_ EKEventStore) SaveReminderCommitError(reminder IEKReminder, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveReminder:commit:error:"), reminder, commit, error_)
	return rv
}


// Saves changes to an event permanently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/save(_:span:)
func (e_ EKEventStore) SaveEventSpanError(event IEKEvent, span EKSpan, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveEvent:span:error:"), event, span, error_)
	return rv
}


// Saves an event or recurring events to the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/save(_:span:commit:)
func (e_ EKEventStore) SaveEventSpanCommitError(event IEKEvent, span EKSpan, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveEvent:span:commit:error:"), event, span, commit, error_)
	return rv
}


// Saves a calendar to the event store by either committing or batching the changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/saveCalendar(_:commit:)
func (e_ EKEventStore) SaveCalendarCommitError(calendar IEKCalendar, commit bool /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("saveCalendar:commit:error:"), calendar, commit, error_)
	return rv
}


// Locates an event source with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/source(withIdentifier:)
func (e_ EKEventStore) SourceWithIdentifier(identifier string /* primitive/slice/pointer. */) IEKSource {
	rv := objc.Send[EKSource](e_.ID, objc.Sel("sourceWithIdentifier:"), objc.String(identifier))
	return rv
}


// The calendars associated with the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendars
func (e_ EKEventStore) Calendars() []EKCalendar /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKCalendar](e_.ID, objc.Sel("calendars"))
	return rv
}


// The calendar that events are added to by default, as specified by user settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/defaultCalendarForNewEvents
func (e_ EKEventStore) DefaultCalendarForNewEvents() IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("defaultCalendarForNewEvents"))
	return rv
}


// The event sources delegated to the person using your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/delegateSources
func (e_ EKEventStore) DelegateSources() []EKSource /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKSource](e_.ID, objc.Sel("delegateSources"))
	return rv
}


// The unique identifier for the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/eventStoreIdentifier
func (e_ EKEventStore) EventStoreIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("eventStoreIdentifier"))
	return rv
}


// An unordered array of objects that represent accounts that contain calendars.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/sources
func (e_ EKEventStore) Sources() []EKSource /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKSource](e_.ID, objc.Sel("sources"))
	return rv
}


