package main

import (
	"flag"
	"fmt"

	"github.com/progrium/darwinkit/dispatch"
	"github.com/progrium/darwinkit/helper/action"
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

var (
	e2e          = flag.Bool("e2e", false, "Run end-to-end test mode")
	stepDelay    = flag.Duration("delay", 0, "Delay between training iterations (e.g. 100ms, 1s)")
	refreshEvery = flag.Int("refresh", 1, "Update UI every N iterations (default: 1 = every iteration)")
)

func main() {
	flag.Parse()

	if *e2e {
		fmt.Println("Running in e2e test mode...")
		runE2ETest()
		return
	}

	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)

	w := appkit.NewWindowWithSize(800, 700)
	objc.Retain(&w)
	w.SetTitle("MNIST Training with MPSGraph - Live Visualization")
	w.Center()

	// Container view
	containerView := appkit.NewView()
	containerView.SetFrame(foundation.Rect{
		Size: foundation.Size{Width: 800, Height: 700},
	})

	// Create visualization view
	vizView := NewVisualizationView(800, 500)
	vizView.imageView.SetFrameOrigin(foundation.Point{X: 0, Y: 150})

	// Status label
	statusLabel := appkit.NewTextFieldWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 110},
		Size:   foundation.Size{Width: 760, Height: 30},
	})
	statusLabel.SetStringValue("Ready to train - Click button to start")
	statusLabel.SetEditable(false)
	statusLabel.SetBezeled(false)
	statusLabel.SetDrawsBackground(false)
	statusLabel.SetFont(appkit.Font_SystemFontOfSize(14))

	// Train button
	trainButton := appkit.NewButtonWithFrame(foundation.Rect{
		Origin: foundation.Point{X: 20, Y: 50},
		Size:   foundation.Size{Width: 150, Height: 40},
	})
	trainButton.SetTitle("Start Training")
	trainButton.SetBezelStyle(appkit.BezelStyleRounded)

	action.Set(trainButton, func(sender objc.Object) {
		statusLabel.SetStringValue("Training in progress...")
		trainButton.SetEnabled(false)

		// Run training asynchronously with live visual updates
		go func() {
			trainer := NewMNISTTrainer(func(iter int, loss float32, samples []SampleData) {
				// Only update UI every N iterations
				if iter%*refreshEvery != 0 {
					return
				}

				// Calculate accuracy from samples
				correct := 0
				for _, s := range samples {
					if s.Predicted == s.Label {
						correct++
					}
				}
				accuracy := float32(correct) / float32(len(samples))

				// Update visualization on main thread
				dispatch.MainQueue().DispatchAsync(func() {
					vizView.Update(samples, accuracy, loss, iter)
					statusLabel.SetStringValue(fmt.Sprintf("Training - Iteration %d/%d - Test Acc: %.1f%% - Loss: %.4f", iter, numIterations, accuracy*100, loss))
				})
			})

			err := trainer.Train()

			// Update final status on main thread
			dispatch.MainQueue().DispatchAsync(func() {
				if err != nil {
					statusLabel.SetStringValue(fmt.Sprintf("Error: %v", err))
				} else {
					statusLabel.SetStringValue("Training complete! Final accuracy shown above.")
				}
				trainButton.SetEnabled(true)
			})
		}()
	})

	// Add all views
	containerView.AddSubview(vizView.View())
	containerView.AddSubview(statusLabel)
	containerView.AddSubview(trainButton)

	w.ContentView().AddSubview(containerView)
	w.MakeKeyAndOrderFront(nil)
	w.SetIsVisible(true)
	app.ActivateIgnoringOtherApps(true)

	if *e2e {
		// Auto-click train button in e2e mode
		trainButton.PerformClick(nil)
	}
}

func runE2ETest() {
	fmt.Println("E2E Test: Training MNIST neural network...")

	trainer := NewMNISTTrainer(func(iter int, loss float32, samples []SampleData) {
		if iter%10 == 0 {
			correct := 0
			for _, s := range samples {
				if s.Predicted == s.Label {
					correct++
				}
			}
			accuracy := float32(correct) / float32(len(samples))
			fmt.Printf("Iteration %3d - Loss: %.4f - Accuracy: %.1f%%\n", iter, loss, accuracy*100)
		}
	})

	err := trainer.Train()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("E2E Test: Training complete!")
}
