// Code generated from Apple documentation for GameSave. DO NOT EDIT.

package gamesave

/* debug [enums.gen.go]: Generating 1 enums for GameSave */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum GSSyncState (7 cases) */
// GSSyncState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState
type GSSyncState uint

const (
	// GSSyncStateClosed - The directory is closed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/closed
	GSSyncStateClosed GSSyncState = 0
	// GSSyncStateConflicted - The directory has conflicts with the cloud, which the game needs to resolve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/conflicted
	GSSyncStateConflicted GSSyncState = 0
	// GSSyncStateError - The directory is in error state and can’t be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/error
	GSSyncStateError GSSyncState = 0
	// GSSyncStateLocal - The directory is local-only and not synced to iCloud.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/local
	GSSyncStateLocal GSSyncState = 0
	// GSSyncStateOffline - The directory is available locally, but not fully synced because the device is offline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/offline
	GSSyncStateOffline GSSyncState = 0
	// GSSyncStateReady - The directory is fully synced and ready to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/ready
	GSSyncStateReady GSSyncState = 0
	// GSSyncStateSyncing - The directory is currently syncing and is not ready yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncState/syncing
	GSSyncStateSyncing GSSyncState = 0
)


