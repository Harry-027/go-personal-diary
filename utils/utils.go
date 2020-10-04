package utils

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"go_personal_diary/dao"
)

var (
	a         fyne.App
	themeType = darkTheme
	dt        *dao.CurrentDate
	u         *dao.UI
)

// Set the Current as per given date ...
func setCurrent(date string) {
	currentDiary := u.DiaryList.GetSpecificDateContent(date)
	if currentDiary != nil {
		u.Current = currentDiary
		u.AppDate.SetGivenDate(currentDiary.Date)
		u.DateLabel.SetText(currentDiary.Date)
		u.Content.SetText(currentDiary.Content)
		u.Title.SetText(currentDiary.Title)
		defaultReadWriteStatus(currentDiary.Date)
	} else {
		setDefaultTitle()
		setDefaultContent()
		setDefaultDate()
	}
}

// Set the dropdown ...
func setDropDown() {
	dropdown := widget.NewSelect(months, func(str string) {
		u.DiaryList.LoadDiary(str)
		setBtnList()
	})
	dropdown.PlaceHolder = dropdownPlaceholder
	dropdown.Selected = u.AppDate.Month
	u.DropDown = dropdown
}

// Build the buttonList based on already saved contents ...
func setBtnList() {
	u.BtnList.Children = nil
	if len(u.DiaryList.List) != 0 {
		for _, newDiary := range u.DiaryList.List {
			diary := newDiary
			newBtn := widget.NewButton(diary.Date, func() {
				setCurrent(diary.Date)
				setBtnList()
			})
			if newDiary.Date == u.AppDate.FormattedDate {
				newBtn.Style = widget.PrimaryButton
			}
			u.BtnList.Append(newBtn)
		}
	} else {
		u.BtnList.Append(widget.NewLabel(noRecord))
	}
}

// Initialize various components and load the diary list ...
func initAndLoadDiaryList() {
	u.DiaryList = &dao.DiaryList{}
	u.DiaryList.Prefs = a.Preferences()
	u.DiaryList.LoadDiary(dt.Month)
	u.BtnList = widget.NewVBox()
	setBtnList()
}

// Save the newly inout content from user ...
func saveOnContentChange() {
	u.Content.OnChanged = func(content string) {
		diary := &dao.Diary{
			Title:   u.Title.Text,
			Date:    u.AppDate.FormattedDate,
			Content: content,
		}
		u.DiaryList.SaveDiary(diary)
	}
	u.Title.OnChanged = func(titleContent string) {
		diary := &dao.Diary{
			Title:   titleContent,
			Date:    u.AppDate.FormattedDate,
			Content: u.Content.Text,
		}
		u.DiaryList.SaveDiary(diary)
	}
}

// Initialize the UI for layout to render ...
func InitUI() {
	u = &dao.UI{}
	setDefaultDate()
	setDefaultTitle()
	setDefaultThemeButton()
	setDefaultBarLayout()
	setDefaultContent()
	initAndLoadDiaryList()
	setCurrent(u.AppDate.FormattedDate)
	setDropDown()
	saveOnContentChange()
}

// Render the layout ...
func LoadUI() fyne.CanvasObject {
	InitUI()
	leftSide := fyne.NewContainerWithLayout(layout.NewBorderLayout(u.DropDown, u.ThemeButton, nil, nil), u.DropDown, u.ThemeButton, u.BtnList)
	rightSide := fyne.NewContainerWithLayout(layout.NewBorderLayout(u.BarLayout, nil, nil, nil), u.BarLayout, u.Content)
	splitLayout := widget.NewHSplitContainer(leftSide, rightSide)
	splitLayout.Offset = layoutOffset
	return splitLayout
}

// Initialize & run the app & window ...
func InitAndRun() {
	a = app.NewWithID(appId)
	w := a.NewWindow(title)
	w.SetContent(LoadUI())
	w.Resize(fyne.NewSize(appWidth, appHeight))
	w.ShowAndRun()
}
