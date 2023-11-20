// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/gi/v2/giv"
	"goki.dev/gi/v2/keyfun"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/icons"
	"goki.dev/mat32/v2"
)

func main() { gimain.Run(app) }

func app() {
	// turn these on to see a traces of various stages of processing..
	// gi.UpdateTrace = true
	// gi.RenderTrace = true
	// gi.LayoutTrace = true
	// gi.WinEventTrace = true
	// gi.WinRenderTrace = true
	// gi.EventTrace = true
	// gi.KeyEventTrace = true
	// events.TraceEventCompression = true
	// events.TraceWindowPaint = true

	gi.SetAppName("widgets")
	gi.SetAppAbout(`This is a demo of the main widgets and general functionality of the <b>GoGi</b> graphical interface system, within the <b>GoKi</b> tree framework.  See <a href="https://github.com/goki">GoKi on GitHub</a>.
<p>The <a href="https://goki.dev/gi/v2/blob/master/examples/widgets/README.md">README</a> page for this example app has lots of further info.</p>`)

	bd := gi.NewBody("widgets").SetTitle("GoGi Widgets Demo")

	// gi.DefaultTopAppBar = nil // turns it off

	sc := gi.NewScene(bd)

	sc.Header.Add(func(par gi.Widget) {
		tb := sc.TopAppBar(par)
		// if gi.DefaultTopAppBar != nil {
		// 	gi.DefaultTopAppBar(tb)
		// }
		gi.NewButton(tb).SetText("Button 1").SetData(1).
			OnClick(func(e events.Event) {
				fmt.Println("TopAppBar Button 1")
				gi.NewSnackbar(tb).Text("Something went wrong!").
					Button("Try again", func(e events.Event) {
						fmt.Println("got snackbar try again event")
					}).
					Icon(icons.Close, func(e events.Event) {
						fmt.Println("got snackbar close icon event")
					}).Run()
			})
		gi.NewButton(tb).SetText("Button 2").SetData(2).
			OnClick(func(e events.Event) {
				fmt.Println("TopAppBar Button 2")
			})
	})

	trow := gi.NewLayout(bd, "trow")
	trow.Style(func(s *styles.Style) {
		s.Align.X = styles.AlignCenter
	})

	giedsc := keyfun.ChordFor(keyfun.Inspector)
	prsc := keyfun.ChordFor(keyfun.Prefs)

	gi.NewLabel(trow, "title").SetText(
		`This is a <b>demonstration</b> of the
		<span style="color:red">various</span> <a href="https://goki.dev/gi/v2">GoGi</a> <i>Widgets</i><br>
		<small>Shortcuts: <kbd>` + string(prsc) + `</kbd> = Preferences,
		<kbd>` + string(giedsc) + `</kbd> = Editor, <kbd>Ctrl/Cmd +/-</kbd> = zoom</small><br>
		See <a href="https://goki.dev/gi/v2/blob/master/examples/widgets/README.md">README</a> for detailed info and things to try.`).
		SetType(gi.LabelHeadlineSmall).
		Style(func(s *styles.Style) {
			s.Align.X = styles.AlignCenter
			s.Text.Align = styles.AlignCenter
			// s.Text.AlignV = styles.AlignCenter
			s.Font.Family = "Times New Roman, serif"
		})

	//////////////////////////////////////////
	//      Buttons

	gi.NewSpace(bd)
	gi.NewLabel(bd).SetText("Buttons:")

	brow := gi.NewLayout(bd, "brow").
		Style(func(s *styles.Style) {
			s.Gap.X.Em(1)
		})

	b1 := gi.NewButton(brow).SetIcon(icons.OpenInNew).SetTooltip("press this <i>button</i> to pop up a dialog box").
		Style(func(s *styles.Style) {
			s.Min.X.Em(1.5)
			s.Min.Y.Em(1.5)
		})

	b1.OnClick(func(e events.Event) {
		fmt.Printf("Button1 clicked\n")

		// b := gi.NewBody()
		// gi.NewLabel(b).SetType(gi.LabelHeadlineLarge).SetText("Test Dialog")
		// gi.NewLabel(b).SetText("This is a prompt")
		// sc := gi.NewScene(b)
		// sc.Sides.Bottom = func(par Widget) {
		// 	brow := gi.NewLayout(par).Style(func(s *styles.Style) {
		// 		s.Align.X = styles.AlignEnd
		// 	})
		// 	gi.NewButton(brow).SetText("Cancel").OnClick(func(e events.Event) {
		// 		sc.Close()
		// 	})
		// 	gi.NewButton(brow).SetText("OK").OnClick(func(e events.Event) {
		// 		sc.Close()
		// 	})
		// }
		// gi.NewDialog(sc).SetModal(true).Run()

		d := gi.NewScene(gi.NewBody().AddTitle("Test Dialog").AddText("This is a prompt"))
		d.Footer.Add(func(par gi.Widget) {
			d.AddCancel(par).OnClick(func(e events.Event) {
			})
			d.AddOk(par).OnClick(func(e events.Event) {
			})
		})
		gi.NewDialog(d).SetContext(b1).Run() // note: NewDialog returns Stage

		// d := gi.NewDialog(b1).AddTitle("Test Dialog").AddText("This is a prompt").
		// 	Modal(true).Cancel().Ok().
		// OnAccept(func(e events.Event) {
		// 	fmt.Println("ok")
		// }).OnCancel(func(e events.Event) {
		// 	fmt.Println("cancel")
		// }).Run()

	})

	button2 := gi.NewButton(brow).SetText("Open Inspector").
		SetTooltip("This button will open the GoGi GUI editor where you can edit this very GUI and see it update dynamically as you change things")
	button2.OnClick(func(e events.Event) {
		txt := ""

		d := gi.NewBody().AddTitle("What is it?").AddText("Please enter your response:")
		giv.NewValue(d, &txt).AsWidget().(*gi.TextField).SetPlaceholder("Enter string here...")
		sc := gi.NewScene(d)
		sc.Footer.Add(func(par gi.Widget) {
			sc.AddCancel(par)
			sc.AddOk(par).OnClick(func(e events.Event) {
				fmt.Println("dialog accepted; string entered:", txt)
			})
		})
		gi.NewDialog(sc).SetContext(button2).Run()
	})

	toggle := gi.NewSwitch(brow).SetText("Toggle")
	toggle.OnChange(func(e events.Event) {
		fmt.Println("toggled", toggle.StateIs(states.Checked))
	})

	mb := gi.NewButton(brow).SetText("Menu Button")
	mb.SetTooltip("Press this button to pull up a nested menu of buttons")

	mb.Menu = func(m *gi.Scene) {
		m1 := gi.NewButton(m).SetText("Menu Item 1").SetIcon(icons.Save).SetShortcut("Shift+Control+1").SetData(1)
		m1.SetTooltip("A standard menu item with an icon").
			OnClick(func(e events.Event) {
				fmt.Println("Received menu action with data", m1.Data)
			})

		m2 := gi.NewButton(m).SetText("Menu Item 2").SetIcon(icons.Open).SetData(2)
		m2.SetTooltip("A menu item with an icon and a sub menu")

		m2.Menu = func(m *gi.Scene) {
			sm2 := gi.NewButton(m).SetText("Sub Menu Item 2").SetIcon(icons.InstallDesktop).SetData(2.1)
			sm2.SetTooltip("A sub menu item with an icon").
				OnClick(func(e events.Event) {
					fmt.Println("Received menu action with data", sm2.Data)
				})
		}

		gi.NewSeparator(m)

		m3 := gi.NewButton(m).SetText("Menu Item 3").SetIcon(icons.Favorite).SetShortcut("Control+3").SetData(3)
		m3.SetTooltip("A standard menu item with an icon, below a separator").
			OnClick(func(e events.Event) {
				fmt.Println("Received menu action with data", m3.Data)
			})
	}

	//////////////////////////////////////////
	//      Sliders

	gi.NewSpace(bd)
	gi.NewLabel(bd).SetText("Sliders:")

	srow := gi.NewLayout(bd).
		Style(func(s *styles.Style) {
			s.Align.Y = styles.AlignCenter
			s.Gap.X.Ex(2)
		})

	slider0 := gi.NewSlider(srow).SetDim(mat32.X).SetValue(0.5).
		SetSnap(true).SetTracking(false).SetIcon(icons.RadioButtonChecked)
	slider0.OnChange(func(e events.Event) {
		fmt.Println("slider0", slider0.Value)
	})
	slider0.Style(func(s *styles.Style) {
		s.Align.Y = styles.AlignCenter
	})

	slider1 := gi.NewSlider(srow).SetDim(mat32.Y).
		SetTracking(true).SetValue(0.5).SetThumbSize(mat32.NewVec2(1, 4))
	slider1.OnChange(func(e events.Event) {
		fmt.Println("slider1", slider1.Value)
	})

	scroll0 := gi.NewSlider(srow).SetType(gi.SliderScrollbar).SetDim(mat32.X).
		SetVisiblePct(0.25).SetValue(0.25).SetTracking(true).SetStep(0.05).SetSnap(true)
	scroll0.OnChange(func(e events.Event) {
		fmt.Println("scroll0", scroll0.Value)
	})
	scroll0.Style(func(s *styles.Style) {
		s.Align.Y = styles.AlignCenter
	})

	scroll1 := gi.NewSlider(srow).SetType(gi.SliderScrollbar).SetDim(mat32.Y).
		SetVisiblePct(.01).SetValue(0).SetMax(3000).
		SetTracking(true).SetStep(1).SetPageStep(10)
	scroll1.OnChange(func(e events.Event) {
		fmt.Println("scroll1", scroll1.Value)
	})

	//////////////////////////////////////////
	//      Text Widgets

	gi.NewLabel(bd).SetText("Text Widgets:")

	txrow := gi.NewLayout(bd).
		Style(func(s *styles.Style) {
			s.Gap.X.Ex(2)
		})

	edit1 := gi.NewTextField(txrow, "edit1").SetPlaceholder("Enter text here...").AddClearButton()
	edit1.OnChange(func(e events.Event) {
		fmt.Println("Text:", edit1.Text())
	})
	edit1.Style(func(s *styles.Style) {
		s.Grow.Set(1, 0)
	})

	sb := gi.NewSpinner(txrow).SetMax(1000).SetMin(-1000).SetStep(5)
	sb.OnChange(func(e events.Event) {
		fmt.Println("spinbox value changed to", sb.Value)
	})

	ch := gi.NewChooser(txrow).SetType(gi.ChooserOutlined).SetEditable(true).
		SetTypes(gti.AllEmbeddersOf(gi.WidgetBaseType), true, true, 50)
	// ItemsFromEnum(gi.ButtonTypesN, true, 50)
	ch.OnChange(func(e events.Event) {
		fmt.Printf("Chooser selected index: %d data: %v\n", ch.CurIndex, ch.CurVal)
	})

	gi.NewWindow(sc).Run().Wait()
}
