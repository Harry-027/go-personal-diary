package utils

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"go_personal_diary/dao"
)

// Set the current date as default date ...
func setDefaultDate() {
	dt = &dao.CurrentDate{}
	dt.SetDate()
	u.AppDate = dt
}

// Set the default title ...
func setDefaultTitle() {
	title := widget.NewEntry()
	title.PlaceHolder = titlePlaceholder
	u.Title = title
}

// Set the default theme toggling button ...
func setDefaultThemeButton() {
	themeButton := widget.NewButton(themeButtonLabel, func() {
		if themeType == darkTheme {
			a.Settings().SetTheme(theme.LightTheme())
			themeType = lightTheme
		} else {
			a.Settings().SetTheme(theme.DarkTheme())
			themeType = darkTheme
		}
	})
	u.ThemeButton = themeButton
}

// Set the default content for app ...
func setDefaultContent() {
	content := widget.NewMultiLineEntry()
	content.Wrapping = fyne.TextWrapWord
	u.Content = content
}

// Set the default barlayout ...
func setDefaultBarLayout() {
	dateBar := widget.NewLabel(todayLabel + u.AppDate.FormattedDate)
	dateBar.Alignment = fyne.TextAlignTrailing
	u.DateLabel = dateBar
	barLayout := widget.NewVSplitContainer(dateBar, u.Title)
	u.BarLayout = barLayout
}

// Set the default read and write status for text content ...
func defaultReadWriteStatus(date string) {
	if u.AppDate.GetCurrentDate() != date {
		u.Content.Disable()
		u.Title.Disable()
	} else {
		u.Content.Enable()
		u.Title.Enable()
	}
}
