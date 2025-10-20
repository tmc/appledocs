package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/mediaplayer"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("MediaPlayer Framework Examples")
	fmt.Println("==============================")

	// Example 1: Create MPMusicPlayerController
	fmt.Println("\n1. Creating MPMusicPlayerController:")

	musicPlayer := mediaplayer.NewMusicPlayerController()
	fmt.Printf("   Music player created: %v\n", musicPlayer)

	// Example 2: MediaPlayer components
	fmt.Println("\n2. MediaPlayer Components:")

	components := map[string]string{
		"MPMusicPlayerController":      "Playback controller for music",
		"MPMediaQuery":                 "Query music library",
		"MPMediaItem":                  "Individual music item",
		"MPMediaItemCollection":        "Collection of items (album, playlist)",
		"MPMediaPickerController":      "UI for selecting music",
		"MPNowPlayingInfoCenter":       "Update Now Playing screen",
		"MPRemoteCommandCenter":        "Handle remote control events",
		"MPVolumeView":                 "Volume and AirPlay controls",
		"MPPlayableContentManager":     "CarPlay integration",
	}

	for component, desc := range components {
		fmt.Printf("   %-30s: %s\n", component, desc)
	}

	// Example 3: Player types
	fmt.Println("\n3. Music Player Types:")

	playerTypes := map[string]string{
		"systemMusicPlayer":       "System-wide music app player",
		"applicationMusicPlayer":  "In-app music player",
		"applicationQueuePlayer":  "Queue-based player with playback queue",
	}

	for playerType, desc := range playerTypes {
		fmt.Printf("   %-25s: %s\n", playerType, desc)
	}

	// Example 4: Playback states
	fmt.Println("\n4. Playback States:")

	states := []string{
		"stopped - Not playing",
		"playing - Currently playing",
		"paused - Paused",
		"interrupted - Interrupted by system",
		"seekingForward - Fast forward",
		"seekingBackward - Rewind",
	}

	for i, state := range states {
		fmt.Printf("   %d. %s\n", i+1, state)
	}

	// Example 5: Media query workflow
	fmt.Println("\n5. Media Query Workflow:")

	queryWorkflow := []string{
		"1. Create MPMediaQuery with filter",
		"2. Add predicates to filter items",
		"3. Execute query to get items",
		"4. Create collection from items",
		"5. Set collection on player",
		"6. Start playback",
	}

	for _, step := range queryWorkflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 6: Media item properties
	fmt.Println("\n6. Media Item Properties:")

	properties := map[string]string{
		"title":                "Song title",
		"albumTitle":           "Album name",
		"artist":               "Artist name",
		"albumArtist":          "Album artist",
		"genre":                "Music genre",
		"composer":             "Composer name",
		"playbackDuration":     "Song length in seconds",
		"albumTrackNumber":     "Track number",
		"albumTrackCount":      "Total tracks in album",
		"discNumber":           "Disc number",
		"artwork":              "Album artwork image",
		"lyrics":               "Song lyrics",
		"isCloudItem":          "From iCloud Music Library",
		"releaseDate":          "Release date",
		"beatsPerMinute":       "Tempo (BPM)",
		"comments":             "User comments",
		"assetURL":             "File URL (local items)",
	}

	for property, desc := range properties {
		fmt.Printf("   %-20s: %s\n", property, desc)
	}

	// Example 7: Query predicates
	fmt.Println("\n7. Common Media Query Predicates:")

	predicates := []string{
		"Artist - Filter by artist name",
		"Album Artist - Filter by album artist",
		"Album - Filter by album title",
		"Genre - Filter by genre",
		"Composer - Filter by composer",
		"Playlist - Filter by playlist",
		"Media Type - Filter by type (music, podcast, audiobook)",
		"Is Cloud Item - Filter cloud vs local",
		"Has Artwork - Filter items with artwork",
	}

	for i, predicate := range predicates {
		fmt.Printf("   %2d. %s\n", i+1, predicate)
	}

	// Example 8: Playback control
	fmt.Println("\n8. Playback Controls:")

	controls := map[string]string{
		"play":                    "Start playback",
		"pause":                   "Pause playback",
		"stop":                    "Stop playback",
		"skipToNextItem":          "Next track",
		"skipToPreviousItem":      "Previous track",
		"skipToBeginning":         "Restart current track",
		"beginSeekingForward":     "Start fast forward",
		"beginSeekingBackward":    "Start rewind",
		"endSeeking":              "Stop seeking",
		"setQueueWithQuery":       "Set playback queue from query",
		"setQueueWithItemCollection": "Set queue from collection",
		"setQueueWithStoreIDs":    "Set queue from store IDs",
	}

	for control, desc := range controls {
		fmt.Printf("   %-30s: %s\n", control, desc)
	}

	// Example 9: Shuffle and repeat modes
	fmt.Println("\n9. Shuffle and Repeat Modes:")

	modes := map[string][]string{
		"Shuffle Modes": {
			"off - No shuffle",
			"songs - Shuffle songs",
			"albums - Shuffle albums",
		},
		"Repeat Modes": {
			"none - No repeat",
			"one - Repeat current item",
			"all - Repeat all items",
		},
	}

	for category, items := range modes {
		fmt.Printf("\n   %s:\n", category)
		for _, item := range items {
			fmt.Printf("     • %s\n", item)
		}
	}

	// Example 10: Notifications
	fmt.Println("\n10. Player Notifications:")

	notifications := []string{
		"playbackStateDidChange - State changed",
		"nowPlayingItemDidChange - Track changed",
		"volumeDidChange - Volume changed",
		"queueDidChange - Queue modified",
	}

	for i, notification := range notifications {
		fmt.Printf("   %2d. %s\n", i+1, notification)
	}

	// Example 11: Remote command events
	fmt.Println("\n11. Remote Command Events:")

	commands := []string{
		"Play command - Play button",
		"Pause command - Pause button",
		"Stop command - Stop button",
		"Toggle play/pause - Play/pause toggle",
		"Next track - Next button",
		"Previous track - Previous button",
		"Seek forward - Fast forward",
		"Seek backward - Rewind",
		"Skip forward - Skip 15/30 seconds",
		"Skip backward - Jump back 15/30 seconds",
		"Change playback rate - Speed control",
		"Like command - Heart/like button",
		"Dislike command - Dislike button",
		"Bookmark command - Add bookmark",
		"Change shuffle mode - Shuffle toggle",
		"Change repeat mode - Repeat toggle",
	}

	for i, command := range commands {
		fmt.Printf("   %2d. %s\n", i+1, command)
	}

	// Example 12: Now Playing info keys
	fmt.Println("\n12. Now Playing Info Keys:")

	infoKeys := []string{
		"Title - Song title",
		"Artist - Artist name",
		"Album Title - Album name",
		"Artwork - Album art image",
		"Duration - Total length",
		"Elapsed Time - Current position",
		"Playback Rate - Speed (1.0 = normal)",
		"Track Number - Track number",
		"Track Count - Total tracks",
		"Album Track Number - Album track",
		"Album Track Count - Album total",
		"Disc Number - Disc number",
		"Disc Count - Total discs",
		"Genre - Music genre",
		"Composer - Composer name",
		"Release Date - Release date",
	}

	for i, key := range infoKeys {
		fmt.Printf("   %2d. %s\n", i+1, key)
	}

	// Example 13: Media picker
	fmt.Println("\n13. Media Picker Features:")

	pickerFeatures := []string{
		"Browse music library",
		"Search for songs, albums, artists",
		"Select single or multiple items",
		"Show all media or filter by type",
		"Allow cloud items or local only",
		"Show item properties",
		"Prompt for library access permission",
	}

	for i, feature := range pickerFeatures {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	// Example 14: Volume view features
	fmt.Println("\n14. Volume View Features:")

	volumeFeatures := []string{
		"System volume slider",
		"AirPlay device picker",
		"AirPlay route button",
		"Wireless route detection",
		"Customizable appearance",
		"Volume change notifications",
	}

	for i, feature := range volumeFeatures {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	// Example 15: Common use cases
	fmt.Println("\n15. Common Use Cases:")

	useCases := map[string]string{
		"Music Player App":     "Full-featured music player",
		"Workout App":          "Music playback during exercise",
		"Party Mode":           "DJ app, playlist management",
		"Music Discovery":      "Browse and sample music",
		"Alarm Clock":          "Wake up to music",
		"Meditation App":       "Guided meditation with music",
		"Podcast Player":       "Audio podcast playback",
		"Audiobook Reader":     "Audiobook player",
		"DJ App":               "Mix and crossfade tracks",
		"Music Library Browser": "Explore music collection",
	}

	for useCase, desc := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, desc)
	}

	// Example 16: Authorization
	fmt.Println("\n16. Library Access Authorization:")

	authSteps := []string{
		"1. Check authorization status",
		"2. Request permission if needed",
		"3. User approves or denies",
		"4. Handle authorization response",
		"5. Query library if authorized",
		"6. Show permission UI if denied",
	}

	for _, step := range authSteps {
		fmt.Printf("   %s\n", step)
	}

	// Example 17: CarPlay integration
	fmt.Println("\n17. CarPlay Integration:")

	carPlayFeatures := []string{
		"MPPlayableContent protocol",
		"Browse music library in CarPlay",
		"Play music in car",
		"Now Playing screen",
		"Voice control with Siri",
		"Safe browsing interface",
		"Automatic UI adaptation",
	}

	for i, feature := range carPlayFeatures {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	fmt.Println("\n✓ MediaPlayer framework examples completed!")
	fmt.Println("\nNote: MediaPlayer provides access to the user's music library:")
	fmt.Println("  - Query and filter music items")
	fmt.Println("  - Control music playback")
	fmt.Println("  - Integrate with Now Playing")
	fmt.Println("  - Handle remote controls")
	fmt.Println("  - Display media picker UI")
	fmt.Println("  - Support CarPlay")
	fmt.Println("\nReal applications would:")
	fmt.Println("  - Request library access permission")
	fmt.Println("  - Query music library with predicates")
	fmt.Println("  - Create playlists and queues")
	fmt.Println("  - Control playback state")
	fmt.Println("  - Update Now Playing info")
	fmt.Println("  - Handle remote command events")
	fmt.Println("  - Display album artwork")
}
