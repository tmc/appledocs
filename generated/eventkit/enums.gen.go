// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

/* debug [enums.gen.go]: Generating 19 enums for EventKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum EKErrorCode (38 cases) */
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

/* debug [enums.gen.go]: Processing enum EKAlarmProximity (3 cases) */
// EKAlarmProximity - A value indicating whether an alarm is triggered by entering or exiting a region.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmProximity
type EKAlarmProximity uint

const (
	// EKAlarmProximityEnter - The alarm is set to fire when entering a region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmProximity/enter
	EKAlarmProximityEnter EKAlarmProximity = 0
	// EKAlarmProximityLeave - The alarm is set to fire when leaving a region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmProximity/leave
	EKAlarmProximityLeave EKAlarmProximity = 0
	// EKAlarmProximityNone - The alarm has no proximity trigger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmProximity/none
	EKAlarmProximityNone EKAlarmProximity = 0
)

/* debug [enums.gen.go]: Processing enum EKAlarmType (4 cases) */
// EKAlarmType - A value that specifies what type of action occurs when the alarm triggers.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmType
type EKAlarmType uint

const (
	// EKAlarmTypeAudio - The alarm plays a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmType/audio
	EKAlarmTypeAudio EKAlarmType = 0
	// EKAlarmTypeDisplay - The alarm displays a message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmType/display
	EKAlarmTypeDisplay EKAlarmType = 0
	// EKAlarmTypeEmail - The alarm sends an email.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmType/email
	EKAlarmTypeEmail EKAlarmType = 0
	// EKAlarmTypeProcedure - The alarm opens a URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarmType/procedure
	EKAlarmTypeProcedure EKAlarmType = 0
)

/* debug [enums.gen.go]: Processing enum EKAuthorizationStatus (6 cases) */
// EKAuthorizationStatus - The current authorization status for a specific entity type.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus
type EKAuthorizationStatus uint

const (
	// EKAuthorizationStatusAuthorized - The app can access the service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/authorized
	EKAuthorizationStatusAuthorized EKAuthorizationStatus = 0
	// EKAuthorizationStatusDenied - The person explicitly denied access to the service for the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/denied
	EKAuthorizationStatusDenied EKAuthorizationStatus = 0
	// EKAuthorizationStatusFullAccess - The app has both read and write access to the requested entity type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/fullAccess
	EKAuthorizationStatusFullAccess EKAuthorizationStatus = 0
	// EKAuthorizationStatusNotDetermined - The person hasn’t chosen whether the app may access the service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/notDetermined
	EKAuthorizationStatusNotDetermined EKAuthorizationStatus = 0
	// EKAuthorizationStatusRestricted - The app isn’t authorized to access the service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/restricted
	EKAuthorizationStatusRestricted EKAuthorizationStatus = 0
	// EKAuthorizationStatusWriteOnly - The app has write-only access to the requested entity type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAuthorizationStatus/writeOnly
	EKAuthorizationStatusWriteOnly EKAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum EKCalendarEventAvailabilityMask (5 cases) */
// EKCalendarEventAvailabilityMask - A bitmask indicating the event availability settings that the calendar can support.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask
type EKCalendarEventAvailabilityMask uint

const (
	// EKCalendarEventAvailabilityBusy - The calendar supports the busy event availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask/busy
	EKCalendarEventAvailabilityBusy EKCalendarEventAvailabilityMask = 0
	// EKCalendarEventAvailabilityNone - The calendar does not support event availability settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask/EKCalendarEventAvailabilityNone
	EKCalendarEventAvailabilityNone EKCalendarEventAvailabilityMask = 0
	// EKCalendarEventAvailabilityFree - The calendar supports the free event availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask/free
	EKCalendarEventAvailabilityFree EKCalendarEventAvailabilityMask = 0
	// EKCalendarEventAvailabilityTentative - The calendar supports the tentative event availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask/tentative
	EKCalendarEventAvailabilityTentative EKCalendarEventAvailabilityMask = 0
	// EKCalendarEventAvailabilityUnavailable - The calendar supports the unavailable event availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarEventAvailabilityMask/unavailable
	EKCalendarEventAvailabilityUnavailable EKCalendarEventAvailabilityMask = 0
)

/* debug [enums.gen.go]: Processing enum EKCalendarType (5 cases) */
// EKCalendarType - Possible calendar types.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType
type EKCalendarType uint

const (
	// EKCalendarTypeBirthday - A birthday calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType/birthday
	EKCalendarTypeBirthday EKCalendarType = 0
	// EKCalendarTypeCalDAV - A CalDAV or iCloud calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType/calDAV
	EKCalendarTypeCalDAV EKCalendarType = 0
	// EKCalendarTypeExchange - An Exchange calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType/exchange
	EKCalendarTypeExchange EKCalendarType = 0
	// EKCalendarTypeLocal - A local calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType/local
	EKCalendarTypeLocal EKCalendarType = 0
	// EKCalendarTypeSubscription - A locally subscribed calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarType/subscription
	EKCalendarTypeSubscription EKCalendarType = 0
)

/* debug [enums.gen.go]: Processing enum EKEntityMask (2 cases) */
// EKEntityMask - A bitmask of 
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityMask
type EKEntityMask uint

const (
	// EKEntityMaskEvent - Represents an event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityMask/event
	EKEntityMaskEvent EKEntityMask = 0
	// EKEntityMaskReminder - Represents a reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityMask/reminder
	EKEntityMaskReminder EKEntityMask = 0
)

/* debug [enums.gen.go]: Processing enum EKEntityType (2 cases) */
// EKEntityType - The type of entities allowed for a source.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityType
type EKEntityType uint

const (
	// EKEntityTypeEvent - Represents an event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityType/event
	EKEntityTypeEvent EKEntityType = 0
	// EKEntityTypeReminder - Represents a reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEntityType/reminder
	EKEntityTypeReminder EKEntityType = 0
)

/* debug [enums.gen.go]: Processing enum EKEventAvailability (5 cases) */
// EKEventAvailability - The event’s availability setting for scheduling purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability
type EKEventAvailability uint

const (
	// EKEventAvailabilityBusy - The event has a busy availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability/busy
	EKEventAvailabilityBusy EKEventAvailability = 0
	// EKEventAvailabilityFree - The event has a free availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability/free
	EKEventAvailabilityFree EKEventAvailability = 0
	// EKEventAvailabilityNotSupported - Availability settings are not supported by the event’s calendar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability/notSupported
	EKEventAvailabilityNotSupported EKEventAvailability = 0
	// EKEventAvailabilityTentative - The event has a tentative availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability/tentative
	EKEventAvailabilityTentative EKEventAvailability = 0
	// EKEventAvailabilityUnavailable - The event has an unavailable availability setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventAvailability/unavailable
	EKEventAvailabilityUnavailable EKEventAvailability = 0
)

/* debug [enums.gen.go]: Processing enum EKEventStatus (4 cases) */
// EKEventStatus - The event’s status.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStatus
type EKEventStatus uint

const (
	// EKEventStatusCanceled - The event is canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStatus/canceled
	EKEventStatusCanceled EKEventStatus = 0
	// EKEventStatusConfirmed - The event is confirmed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStatus/confirmed
	EKEventStatusConfirmed EKEventStatus = 0
	// EKEventStatusNone - The event has no status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStatus/none
	EKEventStatusNone EKEventStatus = 0
	// EKEventStatusTentative - The event is tentative.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStatus/tentative
	EKEventStatusTentative EKEventStatus = 0
)

/* debug [enums.gen.go]: Processing enum EKParticipantRole (5 cases) */
// EKParticipantRole - The participant’s role for an event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole
type EKParticipantRole uint

const (
	// EKParticipantRoleChair - The participant is the chair of the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole/chair
	EKParticipantRoleChair EKParticipantRole = 0
	// EKParticipantRoleNonParticipant - The participant does not have an active role in the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole/nonParticipant
	EKParticipantRoleNonParticipant EKParticipantRole = 0
	// EKParticipantRoleOptional - The participant’s attendance is optional.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole/optional
	EKParticipantRoleOptional EKParticipantRole = 0
	// EKParticipantRoleRequired - The participant’s attendance is required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole/required
	EKParticipantRoleRequired EKParticipantRole = 0
	// EKParticipantRoleUnknown - The participant’s role is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantRole/unknown
	EKParticipantRoleUnknown EKParticipantRole = 0
)

/* debug [enums.gen.go]: Processing enum EKParticipantScheduleStatus (9 cases) */
// EKParticipantScheduleStatus - The participant’s scheduled status.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus
type EKParticipantScheduleStatus uint

const (
	// EKParticipantScheduleStatusCannotDeliver - The invitation wasn’t delivered because the system is unsure of how to deliver it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/cannotDeliver
	EKParticipantScheduleStatusCannotDeliver EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusDelivered - The invitation has been sent and successfully delivered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/delivered
	EKParticipantScheduleStatusDelivered EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusDeliveryFailed - The invitation wasn’t delivered due to a temporary failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/deliveryFailed
	EKParticipantScheduleStatusDeliveryFailed EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusNone - The invitation hasn’t been sent yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/none
	EKParticipantScheduleStatusNone EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusNoPrivileges - The invitation wasn’t delivered because of insufficient privileges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/noPrivileges
	EKParticipantScheduleStatusNoPrivileges EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusPending - The invitation is in the process of being sent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/pending
	EKParticipantScheduleStatusPending EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusRecipientNotAllowed - The invitation wasn’t delivered because scheduling with the participant isn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/recipientNotAllowed
	EKParticipantScheduleStatusRecipientNotAllowed EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusRecipientNotRecognized - The invitation wasn’t delivered because the source doesn’t recognize the recipient.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/recipientNotRecognized
	EKParticipantScheduleStatusRecipientNotRecognized EKParticipantScheduleStatus = 0
	// EKParticipantScheduleStatusSent - The invitation has been sent, but it’s unclear if it was successfully delivered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantScheduleStatus/sent
	EKParticipantScheduleStatusSent EKParticipantScheduleStatus = 0
)

/* debug [enums.gen.go]: Processing enum EKParticipantStatus (8 cases) */
// EKParticipantStatus - The participant’s attendance status for an event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus
type EKParticipantStatus uint

const (
	// EKParticipantStatusAccepted - The participant has accepted the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/accepted
	EKParticipantStatusAccepted EKParticipantStatus = 0
	// EKParticipantStatusCompleted - The participant’s event has completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/completed
	EKParticipantStatusCompleted EKParticipantStatus = 0
	// EKParticipantStatusDeclined - The participant has declined the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/declined
	EKParticipantStatusDeclined EKParticipantStatus = 0
	// EKParticipantStatusDelegated - The participant has delegated attendance to another participant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/delegated
	EKParticipantStatusDelegated EKParticipantStatus = 0
	// EKParticipantStatusInProcess - The participant’s event is currently in process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/inProcess
	EKParticipantStatusInProcess EKParticipantStatus = 0
	// EKParticipantStatusPending - The participant has yet to respond to the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/pending
	EKParticipantStatusPending EKParticipantStatus = 0
	// EKParticipantStatusTentative - The participant’s attendance status is tentative.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/tentative
	EKParticipantStatusTentative EKParticipantStatus = 0
	// EKParticipantStatusUnknown - The participant’s attendance status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantStatus/unknown
	EKParticipantStatusUnknown EKParticipantStatus = 0
)

/* debug [enums.gen.go]: Processing enum EKParticipantType (5 cases) */
// EKParticipantType - The type of participant.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType
type EKParticipantType uint

const (
	// EKParticipantTypeGroup - The participant is a group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType/group
	EKParticipantTypeGroup EKParticipantType = 0
	// EKParticipantTypePerson - The participant is a person.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType/person
	EKParticipantTypePerson EKParticipantType = 0
	// EKParticipantTypeResource - The participant is a resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType/resource
	EKParticipantTypeResource EKParticipantType = 0
	// EKParticipantTypeRoom - The participant is a room.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType/room
	EKParticipantTypeRoom EKParticipantType = 0
	// EKParticipantTypeUnknown - The participant’s type is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipantType/unknown
	EKParticipantTypeUnknown EKParticipantType = 0
)

/* debug [enums.gen.go]: Processing enum EKRecurrenceFrequency (4 cases) */
// EKRecurrenceFrequency - The frequency for recurrence rules.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceFrequency
type EKRecurrenceFrequency uint

const (
	// EKRecurrenceFrequencyDaily - Indicates a daily recurrence rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceFrequency/daily
	EKRecurrenceFrequencyDaily EKRecurrenceFrequency = 0
	// EKRecurrenceFrequencyMonthly - Indicates a monthly recurrence rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceFrequency/monthly
	EKRecurrenceFrequencyMonthly EKRecurrenceFrequency = 0
	// EKRecurrenceFrequencyWeekly - Indicates a weekly recurrence rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceFrequency/weekly
	EKRecurrenceFrequencyWeekly EKRecurrenceFrequency = 0
	// EKRecurrenceFrequencyYearly - Indicates a yearly recurrence rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceFrequency/yearly
	EKRecurrenceFrequencyYearly EKRecurrenceFrequency = 0
)

/* debug [enums.gen.go]: Processing enum EKReminderPriority (4 cases) */
// EKReminderPriority - The priority of the reminder.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminderPriority
type EKReminderPriority uint

const (
	// EKReminderPriorityHigh - The reminder is high priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminderPriority/high
	EKReminderPriorityHigh EKReminderPriority = 0
	// EKReminderPriorityLow - The reminder is low priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminderPriority/low
	EKReminderPriorityLow EKReminderPriority = 0
	// EKReminderPriorityMedium - The reminder is medium priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminderPriority/medium
	EKReminderPriorityMedium EKReminderPriority = 0
	// EKReminderPriorityNone - The reminder has no priority set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminderPriority/none
	EKReminderPriorityNone EKReminderPriority = 0
)

/* debug [enums.gen.go]: Processing enum EKSourceType (6 cases) */
// EKSourceType - The type of source object.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType
type EKSourceType uint

const (
	// EKSourceTypeBirthdays - Represents a birthday source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType/birthdays
	EKSourceTypeBirthdays EKSourceType = 0
	// EKSourceTypeCalDAV - Represents a CalDAV or iCloud source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType/calDAV
	EKSourceTypeCalDAV EKSourceType = 0
	// EKSourceTypeExchange - Represents an Exchange source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType/exchange
	EKSourceTypeExchange EKSourceType = 0
	// EKSourceTypeLocal - Represents a local source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType/local
	EKSourceTypeLocal EKSourceType = 0
	// EKSourceTypeMobileMe - Represents a MobileMe source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType/mobileMe
	EKSourceTypeMobileMe EKSourceType = 0
	// EKSourceTypeSubscribed - Represents a subscribed source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSourceType/subscribed
	EKSourceTypeSubscribed EKSourceType = 0
)

/* debug [enums.gen.go]: Processing enum EKSpan (2 cases) */
// EKSpan - An object that indicates whether modifications should apply to a single event or all future events of a recurring event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSpan
type EKSpan uint

const (
	// EKSpanFutureEvents - Modifications to this event instance should also affect future instances of this event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSpan/futureEvents
	EKSpanFutureEvents EKSpan = 0
	// EKSpanThisEvent - Modifications to this event instance should affect only this instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSpan/thisEvent
	EKSpanThisEvent EKSpan = 0
)

/* debug [enums.gen.go]: Processing enum EKWeekday (14 cases) */
// EKWeekday - The day of the week.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday
type EKWeekday uint

const (
	// EKFriday - The value for Friday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKFriday
	EKFriday EKWeekday = 0
	// EKMonday - The value for Monday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKMonday
	EKMonday EKWeekday = 0
	// EKSaturday - The value for Saturday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKSaturday
	EKSaturday EKWeekday = 0
	// EKSunday - The value for Sunday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKSunday
	EKSunday EKWeekday = 0
	// EKThursday - The value for Thursday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKThursday
	EKThursday EKWeekday = 0
	// EKTuesday - The value for Tuesday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKTuesday
	EKTuesday EKWeekday = 0
	// EKWednesday - The value for Wednesday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/EKWednesday
	EKWednesday EKWeekday = 0
	// EKWeekdayFriday - The value for Friday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/friday
	EKWeekdayFriday EKWeekday = 0
	// EKWeekdayMonday - The value for Monday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/monday
	EKWeekdayMonday EKWeekday = 0
	// EKWeekdaySaturday - The value for Saturday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/saturday
	EKWeekdaySaturday EKWeekday = 0
	// EKWeekdaySunday - The value for Sunday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/sunday
	EKWeekdaySunday EKWeekday = 0
	// EKWeekdayThursday - The value for Thursday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/thursday
	EKWeekdayThursday EKWeekday = 0
	// EKWeekdayTuesday - The value for Tuesday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/tuesday
	EKWeekdayTuesday EKWeekday = 0
	// EKWeekdayWednesday - The value for Wednesday.
	//
	// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKWeekday/wednesday
	EKWeekdayWednesday EKWeekday = 0
)


