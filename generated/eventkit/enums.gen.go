// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

// Enum types and constants
// EKAlarmProximity - A value indicating whether an alarm is triggered by entering or exiting a region.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmProximity
type EKAlarmProximity uint

// EKAlarmType - A value that specifies what type of action occurs when the alarm triggers.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmType
type EKAlarmType uint

// EKAuthorizationStatus - The current authorization status for a specific entity type.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus
type EKAuthorizationStatus uint

const (
	// EKAuthorizationStatusFullAccess - The app has both read and write access to the requested entity type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/fullAccess
	EKAuthorizationStatusFullAccess EKAuthorizationStatus = 0
	// EKAuthorizationStatusNotDetermined - The person hasn’t chosen whether the app may access the service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/notDetermined
	EKAuthorizationStatusNotDetermined EKAuthorizationStatus = 0
)

// EKCalendarEventAvailabilityMask - A bitmask indicating the event availability settings that the calendar can support.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask
type EKCalendarEventAvailabilityMask uint

// EKCalendarType - Possible calendar types.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType
type EKCalendarType uint

// EKEntityType - The type of entities allowed for a source.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityType
type EKEntityType uint

const (
	// EKEntityTypeReminder - Represents a reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityType/reminder
	EKEntityTypeReminder EKEntityType = 0
)

// EKErrorCode - Error codes for EventKit errors.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code
type EKErrorCode uint

const (
	// EKErrorAlarmGreaterThanRecurrence - The alarm interval is greater than the recurrence interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/alarmGreaterThanRecurrence
	EKErrorAlarmGreaterThanRecurrence EKErrorCode = 0
	// EKErrorAlarmProximityNotSupported - The source doesn’t allow geofences on alarms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/alarmProximityNotSupported
	EKErrorAlarmProximityNotSupported EKErrorCode = 0
	// EKErrorCalendarDoesNotAllowEvents - The calendar doesn’t allow you to add events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/calendarDoesNotAllowEvents
	EKErrorCalendarDoesNotAllowEvents EKErrorCode = 0
	// EKErrorCalendarDoesNotAllowReminders - The calendar doesn’t allow you to add reminders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/calendarDoesNotAllowReminders
	EKErrorCalendarDoesNotAllowReminders EKErrorCode = 0
	// EKErrorCalendarHasNoSource - You can’t save the calendar without setting a source first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/calendarHasNoSource
	EKErrorCalendarHasNoSource EKErrorCode = 0
	// EKErrorCalendarIsImmutable - The calendar is immutable and you can’t modify or delete it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/calendarIsImmutable
	EKErrorCalendarIsImmutable EKErrorCode = 0
	// EKErrorCalendarReadOnly - The calendar is read-only and you can’t add events to it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/calendarReadOnly
	EKErrorCalendarReadOnly EKErrorCode = 0
	// EKErrorCalendarSourceCannotBeModified - You can’t move the calendar to another source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/calendarSourceCannotBeModified
	EKErrorCalendarSourceCannotBeModified EKErrorCode = 0
	// EKErrorDatesInverted - The event’s end date occurs before its start date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/datesInverted
	EKErrorDatesInverted EKErrorCode = 0
	// EKErrorDurationGreaterThanRecurrence - The duration of an event is greater than its recurrence interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/durationGreaterThanRecurrence
	EKErrorDurationGreaterThanRecurrence EKErrorCode = 0
	// EKErrorEventNotMutable - The event isn’t mutable and you can’t save or delete it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/eventNotMutable
	EKErrorEventNotMutable EKErrorCode = 0
	// EKErrorEventStoreNotAuthorized - The user hasn’t authorized your app to access events or reminders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/eventStoreNotAuthorized
	EKErrorEventStoreNotAuthorized EKErrorCode = 0
	// EKErrorInternalFailure - An internal error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/internalFailure
	EKErrorInternalFailure EKErrorCode = 0
	// EKErrorInvalidEntityType - The entity type is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/invalidEntityType
	EKErrorInvalidEntityType EKErrorCode = 0
	// EKErrorInvalidInviteReplyCalendar - The calendar is invalid or nil.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/invalidInviteReplyCalendar
	EKErrorInvalidInviteReplyCalendar EKErrorCode = 0
	// EKErrorInvalidSpan - The system encountered an invalid span during a save or deletion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/invalidSpan
	EKErrorInvalidSpan EKErrorCode = 0
	// EKErrorInvitesCannotBeMoved - You can’t move the event because it’s an invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/invitesCannotBeMoved
	EKErrorInvitesCannotBeMoved EKErrorCode = 0
	// EKErrorLast - This error is for internal use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/last
	EKErrorLast EKErrorCode = 0
	// EKErrorNoCalendar - The event isn’t associated with a calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/noCalendar
	EKErrorNoCalendar EKErrorCode = 0
	// EKErrorNoEndDate - The event has no end date set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/noEndDate
	EKErrorNoEndDate EKErrorCode = 0
	// EKErrorNoStartDate - The event has no start date set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/noStartDate
	EKErrorNoStartDate EKErrorCode = 0
	// EKErrorNotificationCollectionMismatch - The notification collection that contains this notification doesn’t match the collection the system is trying to save.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/notificationCollectionMismatch
	EKErrorNotificationCollectionMismatch EKErrorCode = 0
	// EKErrorNotificationSavedWithoutCollection - The notification can’t save because you haven’t added it to a notification collection and saved the collection first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/notificationSavedWithoutCollection
	EKErrorNotificationSavedWithoutCollection EKErrorCode = 0
	// EKErrorNotificationsCollectionFlagNotSet - The notification collection doesn’t have the notifications collection flag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/notificationsCollectionFlagNotSet
	EKErrorNotificationsCollectionFlagNotSet EKErrorCode = 0
	// EKErrorObjectBelongsToDifferentStore - The object belongs to a different calendar store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/objectBelongsToDifferentStore
	EKErrorObjectBelongsToDifferentStore EKErrorCode = 0
	// EKErrorOSNotSupported - The action isn’t supported on the current operating system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/osNotSupported
	EKErrorOSNotSupported EKErrorCode = 0
	// EKErrorPriorityIsInvalid - The priority number for the reminder is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/priorityIsInvalid
	EKErrorPriorityIsInvalid EKErrorCode = 0
	// EKErrorProcedureAlarmsNotMutable - You can’t create or modify procedure alarms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/procedureAlarmsNotMutable
	EKErrorProcedureAlarmsNotMutable EKErrorCode = 0
	// EKErrorRecurringReminderRequiresDueDate - The recurring reminder requires a due date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/recurringReminderRequiresDueDate
	EKErrorRecurringReminderRequiresDueDate EKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/reminderAlarmContainsEmailOrUrl
	EKErrorReminderAlarmContainsEmailOrUrl EKErrorCode = 0
	// EKErrorReminderLocationsNotSupported - The source doesn’t support locations on reminders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/reminderLocationsNotSupported
	EKErrorReminderLocationsNotSupported EKErrorCode = 0
	// EKErrorSourceDoesNotAllowCalendarAddDelete - The source doesn’t allow you to add or delete calendars.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/sourceDoesNotAllowCalendarAddDelete
	EKErrorSourceDoesNotAllowCalendarAddDelete EKErrorCode = 0
	// EKErrorSourceDoesNotAllowEvents - The source doesn’t allow calendars supporting event entity types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/sourceDoesNotAllowEvents
	EKErrorSourceDoesNotAllowEvents EKErrorCode = 0
	// EKErrorSourceDoesNotAllowReminders - The source doesn’t allow calendars supporting reminder entity types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/sourceDoesNotAllowReminders
	EKErrorSourceDoesNotAllowReminders EKErrorCode = 0
	// EKErrorSourceMismatch - The object’s source doesn’t match its container’s source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/sourceMismatch
	EKErrorSourceMismatch EKErrorCode = 0
	// EKErrorStartDateCollidesWithOtherOccurrence - The event’s start date collides with another occurrence of the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/startDateCollidesWithOtherOccurrence
	EKErrorStartDateCollidesWithOtherOccurrence EKErrorCode = 0
	// EKErrorStartDateTooFarInFuture - The start date is further into the future than the calendar can display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/startDateTooFarInFuture
	EKErrorStartDateTooFarInFuture EKErrorCode = 0
	// EKErrorStructuredLocationsNotSupported - The source to which this calendar belongs doesn’t support structured locations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKError/Code/structuredLocationsNotSupported
	EKErrorStructuredLocationsNotSupported EKErrorCode = 0
)

// EKEventAvailability - The event’s availability setting for scheduling purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability
type EKEventAvailability uint

// EKEventStatus - The event’s status.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStatus
type EKEventStatus uint

// EKParticipantRole - The participant’s role for an event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole
type EKParticipantRole uint

// EKParticipantScheduleStatus - The participant’s scheduled status.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus
type EKParticipantScheduleStatus uint

// EKParticipantStatus - The participant’s attendance status for an event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus
type EKParticipantStatus uint

// EKParticipantType - The type of participant.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType
type EKParticipantType uint

// EKRecurrenceFrequency - The frequency for recurrence rules.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceFrequency
type EKRecurrenceFrequency uint

// EKReminderPriority - The priority of the reminder.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminderPriority
type EKReminderPriority uint

// EKSourceType - The type of source object.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType
type EKSourceType uint

// EKSpan - An object that indicates whether modifications should apply to a single event or all future events of a recurring event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSpan
type EKSpan uint

const (
	// EKSpanFutureEvents - Modifications to this event instance should also affect future instances of this event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSpan/futureEvents
	EKSpanFutureEvents EKSpan = 0
)

// EKWeekday - The day of the week.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday
type EKWeekday uint

const (
	// EKFriday - The value for Friday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKFriday
	EKFriday EKWeekday = 0
	// EKTuesday - The value for Tuesday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKTuesday
	EKTuesday EKWeekday = 0
)


