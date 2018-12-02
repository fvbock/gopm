package gopm

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type TUI struct {
	app        *tview.Application
	inputField *tview.InputField
	textView   *tview.TextView
	layout     *tview.Flex
}

func (tui *TUI) checkAppLevelKeyFunctions(key tcell.Key) {
	if key == tcell.KeyEsc {
		tui.app.Stop()
	}
}

func NewTUI() (tui *TUI) {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	inputField := tview.NewInputField().
		SetLabel(" 	What are you searching for: ").
		SetPlaceholder("AWS Tokyo")
		// SetFieldWidth(10).
		// SetAcceptanceFunc(tview.InputFieldInteger).

	inputField.SetDoneFunc(func(key tcell.Key) {
		// checkAppLevelKeyFunctions(key)

		if key == tcell.KeyTab {
			app.SetFocus(textView)
		}
	})

	textView.SetDoneFunc(func(key tcell.Key) {
		// checkAppLevelKeyFunctions(key)
		// currentSelection := textView.GetHighlights()

		if key == tcell.KeyTab {
			app.SetFocus(inputField)
		}
		// } else if key == tcell.KeyEnter {
		// 	if len(currentSelection) > 0 {
		// 		textView.Highlight()
		// 	} else {
		// 		textView.Highlight("0").ScrollToHighlight()
		// 	}
		// } else if len(currentSelection) > 0 {
		// 	index, _ := strconv.Atoi(currentSelection[0])
		// 	if key == tcell.KeyTab {
		// 		index = (index + 1) % numSelections
		// 	} else if key == tcell.KeyBacktab {
		// 		index = (index - 1 + numSelections) % numSelections
		// 	} else {
		// 		return
		// 	}
		// 	textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
		// }
	})

	textView.SetBorder(true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(inputField, 1, 1, true).
		AddItem(textView, 0, 1, false)

	tui = &TUI{
		app:        app,
		inputField: inputField,
		textView:   textView,
		layout:     flex,
	}

	return
}
