package main

import (
	//"fmt"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/micmonay/keybd_event"
)

// is bot running?
var running bool

func runRobotCore() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(180 * time.Second)
		if running {
			// Select keys to be pressed
			kb.SetKeys(keybd_event.VK_F5)
			//fmt.Println("Ligado")
			kb.Press()
			kb.Release()
		} else {
			break
		}
	}
}

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Persist Login"),
			app.Size(unit.Dp(400), unit.Dp(200)),
		)

		// ops are the operations from the UI
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defines the material design style
		th := material.NewTheme(gofont.Collection())

		// listen for events in the window.
		for e := range w.Events() {

			// detect what type of event
			switch e := e.(type) {

			// this is sent when the application should re-render.
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				// Let's try out the flexbox layout concept:
				if startButton.Clicked() {
					running = !running
					go runRobotCore()
				}
				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start, i.e. at the top
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					// We insert two rigid elements:
					// First a button ...
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							var text string
							if !running {
								text = "Start"
							} else {
								text = "Stop"
							}
							btn := material.Button(th, &startButton, text)
							return btn.Layout(gtx)
						},
					),
					// ... then an empty spacer
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				e.Frame(gtx.Ops)
			}
		}

		os.Exit(0)
	}()
	app.Main()
}
