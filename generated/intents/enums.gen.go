// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

// Enum types and constants
// INCallAudioRoute - Constants that describe the audio route for the call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallAudioRoute
type INCallAudioRoute uint

const (
	// INCallAudioRouteBluetoothAudioRoute - A connected Bluetooth device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallAudioRoute/bluetoothAudioRoute
	INCallAudioRouteBluetoothAudioRoute INCallAudioRoute = 0
	// INCallAudioRouteSpeakerphoneAudioRoute - The device’s speakerphone mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallAudioRoute/speakerphoneAudioRoute
	INCallAudioRouteSpeakerphoneAudioRoute INCallAudioRoute = 0
	// INCallAudioRouteUnknown - An unknown audio route.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallAudioRoute/unknown
	INCallAudioRouteUnknown INCallAudioRoute = 0
)

// INCallCapability - Constants indicating the capabilities of the call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallCapability
type INCallCapability uint

const (
	// INCallCapabilityUnknown - An unknown type of call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallCapability/unknown
	INCallCapabilityUnknown INCallCapability = 0
)

// INCallDestinationType - Constants describing the destination of a call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallDestinationType
type INCallDestinationType uint

const (
	// INCallDestinationTypeVoicemailDestination - A call routed to the user’s voicemail.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallDestinationType/voicemailDestination
	INCallDestinationTypeVoicemailDestination INCallDestinationType = 0
)

// INCallRecordType - Constants describing the type of the call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecordType
type INCallRecordType uint

const (
	// INCallRecordTypeMissed - A received call that the user did not answer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecordType/missed
	INCallRecordTypeMissed INCallRecordType = 0
)

// INConditionalOperator - Constants indicating how search attributes are interpreted.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INConditionalOperator
type INConditionalOperator uint

// INFocusStatusAuthorizationStatus - A constant that indicates whether your app has authorization to access the user’s focus status.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusAuthorizationStatus
type INFocusStatusAuthorizationStatus uint

const (
	// INFocusStatusAuthorizationStatusAuthorized - Your app has authorization to access the user’s focus status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusAuthorizationStatus/authorized
	INFocusStatusAuthorizationStatusAuthorized INFocusStatusAuthorizationStatus = 0
	// INFocusStatusAuthorizationStatusDenied - The user has denied your app access to their focus status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusAuthorizationStatus/denied
	INFocusStatusAuthorizationStatusDenied INFocusStatusAuthorizationStatus = 0
	// INFocusStatusAuthorizationStatusNotDetermined - The user hasn’t chosen whether to grant your app access to their focus status yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusAuthorizationStatus/notDetermined
	INFocusStatusAuthorizationStatusNotDetermined INFocusStatusAuthorizationStatus = 0
	// INFocusStatusAuthorizationStatusRestricted - A restriction prevents your app from accessing the user’s focus status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusAuthorizationStatus/restricted
	INFocusStatusAuthorizationStatusRestricted INFocusStatusAuthorizationStatus = 0
)

// INMessageAttributeOptions - Constants that indicate a message search filter.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMessageAttributeOptions
type INMessageAttributeOptions uint

// INOutgoingMessageType - The format of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INOutgoingMessageType
type INOutgoingMessageType uint

const (
	// INOutgoingMessageTypeOutgoingMessageAudio - An audio recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INOutgoingMessageType/outgoingMessageAudio
	INOutgoingMessageTypeOutgoingMessageAudio INOutgoingMessageType = 0
)

// INPersonSuggestionType - Constants indicating how to display the person’s identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPersonSuggestionType
type INPersonSuggestionType uint

const (
	// INPersonSuggestionTypeNone - No contact information to donate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INPersonSuggestionType/none
	INPersonSuggestionTypeNone INPersonSuggestionType = 0
)

// INPlaybackRepeatMode - The possible repeat modes at the time the user plays the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlaybackRepeatMode
type INPlaybackRepeatMode uint

const (
	// INPlaybackRepeatModeNone - A mode that doesn’t repeat media items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlaybackRepeatMode/none
	INPlaybackRepeatModeNone INPlaybackRepeatMode = 0
)

// INRelevantShortcutRole - Roles for a relevant shortcut.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRelevantShortcutRole
type INRelevantShortcutRole uint

// INReservationStatus - Constants that describe the current status of the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservationStatus
type INReservationStatus uint

// INRidePhase - Constants indicating the current ride status.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase
type INRidePhase uint

const (
	// INRidePhaseApproachingPickup - The driver is approaching the pickup location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/approachingPickup
	INRidePhaseApproachingPickup INRidePhase = 0
	// INRidePhaseCompleted - The driver has transited through all of the waypoints and is now at the ride’s destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/completed
	INRidePhaseCompleted INRidePhase = 0
	// INRidePhaseConfirmed - You have booked the ride and communicated the pickup information to the driver of the vehicle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/confirmed
	INRidePhaseConfirmed INRidePhase = 0
	// INRidePhaseOngoing - The driver has picked up the user’s party and is en route to the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/ongoing
	INRidePhaseOngoing INRidePhase = 0
	// INRidePhasePickup - The driver is at the pickup location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/pickup
	INRidePhasePickup INRidePhase = 0
	// INRidePhaseReceived - You have received the booking request and are processing it, but have not yet confirmed the request or communicated the pickup information to the driver of the vehicle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/received
	INRidePhaseReceived INRidePhase = 0
	// INRidePhaseUnknown - The state of the ride is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INRidePhase/unknown
	INRidePhaseUnknown INRidePhase = 0
)

// INSearchForMessagesIntentResponseCode - Constants that indicate the response state.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMessagesIntentResponseCode
type INSearchForMessagesIntentResponseCode uint

// INSendMessageIntentResponseCode - Constants that indicate the response state.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendMessageIntentResponseCode
type INSendMessageIntentResponseCode uint

// INShareFocusStatusIntentResponseCode - A constant that indicates your app’s ability to handle an intent to share the user’s focus status.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode
type INShareFocusStatusIntentResponseCode uint

const (
	// INShareFocusStatusIntentResponseCodeFailure - Your intent handler is unable to handle the intent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode/failure
	INShareFocusStatusIntentResponseCodeFailure INShareFocusStatusIntentResponseCode = 0
	// INShareFocusStatusIntentResponseCodeFailureRequiringAppLaunch - The user needs to launch your app to update their focus status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode/failureRequiringAppLaunch
	INShareFocusStatusIntentResponseCodeFailureRequiringAppLaunch INShareFocusStatusIntentResponseCode = 0
	// INShareFocusStatusIntentResponseCodeInProgress - Your intent handler is handling the intent, but it may take some time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode/inProgress
	INShareFocusStatusIntentResponseCodeInProgress INShareFocusStatusIntentResponseCode = 0
	// INShareFocusStatusIntentResponseCodeReady - Your intent handler is ready to handle the intent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode/ready
	INShareFocusStatusIntentResponseCodeReady INShareFocusStatusIntentResponseCode = 0
	// INShareFocusStatusIntentResponseCodeSuccess - Your intent handler successfully updated the user’s communication status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode/success
	INShareFocusStatusIntentResponseCodeSuccess INShareFocusStatusIntentResponseCode = 0
	// INShareFocusStatusIntentResponseCodeUnspecified - An unspecified response code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntentResponseCode/unspecified
	INShareFocusStatusIntentResponseCodeUnspecified INShareFocusStatusIntentResponseCode = 0
)

// INSiriAuthorizationStatus - Constants indicating the authorization status of your Intents extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSiriAuthorizationStatus
type INSiriAuthorizationStatus uint

const (
	// INSiriAuthorizationStatusAuthorized - Authorized. Siri is enabled and your app is authorized to interact with it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INSiriAuthorizationStatus/authorized
	INSiriAuthorizationStatusAuthorized INSiriAuthorizationStatus = 0
	// INSiriAuthorizationStatusDenied - Not authorized. The user explicitly denied authorization for this app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INSiriAuthorizationStatus/denied
	INSiriAuthorizationStatusDenied INSiriAuthorizationStatus = 0
	// INSiriAuthorizationStatusNotDetermined - Not yet determined. An authorization request has not yet been made or the user has not yet made a choice regarding the status of the app. Call the   method to request authorization from the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INSiriAuthorizationStatus/notDetermined
	INSiriAuthorizationStatusNotDetermined INSiriAuthorizationStatus = 0
	// INSiriAuthorizationStatusRestricted - Restricted. The app is not authorized to use Siri services. This status could be the result of active restrictions on Siri services rather than on the user denying access.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INSiriAuthorizationStatus/restricted
	INSiriAuthorizationStatusRestricted INSiriAuthorizationStatus = 0
)

// INStartCallCallRecordToCallBackUnsupportedReason - A reason why your app can’t use a record to call a person back.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallCallRecordToCallBackUnsupportedReason
type INStartCallCallRecordToCallBackUnsupportedReason uint

const (
	// INStartCallCallRecordToCallBackUnsupportedReasonNoMatchingCall - A reason indicating that no call record matches the intent’s parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallCallRecordToCallBackUnsupportedReason/noMatchingCall
	INStartCallCallRecordToCallBackUnsupportedReasonNoMatchingCall INStartCallCallRecordToCallBackUnsupportedReason = 0
)

// INStartCallIntentResponseCode - Constants that indicate the response state.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntentResponseCode
type INStartCallIntentResponseCode uint

// INUpcomingMediaPredictionMode - Prediction modes for upcoming media intent shortcuts.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpcomingMediaPredictionMode
type INUpcomingMediaPredictionMode uint

// INVocabularyStringType - Possible usages for a custom vocabulary term.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType
type INVocabularyStringType uint

const (
	// INVocabularyStringTypeMediaAudiobookTitle - The name or title of an audiobook.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType/mediaAudiobookTitle
	INVocabularyStringTypeMediaAudiobookTitle INVocabularyStringType = 0
	// INVocabularyStringTypeMediaPlaylistTitle - The name or title of a playlist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType/mediaPlaylistTitle
	INVocabularyStringTypeMediaPlaylistTitle INVocabularyStringType = 0
	// INVocabularyStringTypeNotebookItemGroupName - The title of a group (or folder) containing the user’s notes or task lists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType/notebookItemGroupName
	INVocabularyStringTypeNotebookItemGroupName INVocabularyStringType = 0
	// INVocabularyStringTypePaymentsOrganizationName - The name of the bank or company that holds the user’s account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType/paymentsOrganizationName
	INVocabularyStringTypePaymentsOrganizationName INVocabularyStringType = 0
	// INVocabularyStringTypePhotoTag - A user-defined keyword associated with an image or images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType/photoTag
	INVocabularyStringTypePhotoTag INVocabularyStringType = 0
	// INVocabularyStringTypeWorkoutActivityName - The user-given name assigned to a workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabularyStringType/workoutActivityName
	INVocabularyStringTypeWorkoutActivityName INVocabularyStringType = 0
)


