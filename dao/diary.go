package dao

import (
	"encoding/json"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"log"
	"strings"
)

const (
	ALLKEY = "ALL_KEYS"
)

type (
	UI struct {
		Content     *widget.Entry
		Title       *widget.Entry
		DropDown    *widget.Select
		ThemeButton *widget.Button
		BarLayout   *widget.SplitContainer
		DateLabel   *widget.Label
		BtnList     *widget.Box
		AppDate     *CurrentDate
		Current     *Diary
		DiaryList   *DiaryList
	}

	Diary struct {
		Title   string
		Date    string
		Content string
	}

	DiaryList struct {
		List    []*Diary
		Allkeys []string
		Prefs   fyne.Preferences
	}
)

// Load the diary contents ...
func (dl *DiaryList) LoadDiary(selectedMonth string) {
	dl.List = nil
	dl.Allkeys = dl.LoadSelectedKeys(selectedMonth)
	if len(dl.Allkeys) != 0 {
		for _, key := range dl.Allkeys {
			diary := &Diary{}
			content := dl.Prefs.String(key)
			byteContent := []byte(content)
			fromJSON(byteContent, diary)
			dl.List = append(dl.List, diary)
		}
	}
}

// Save the diary contents (title, content) in app preferences ...
func (dl *DiaryList) SaveDiary(d *Diary) {
	var doesKeyExist bool = false
	keyDetails := dl.Prefs.String(ALLKEY)
	dl.Allkeys = *LoadStoredKeys(keyDetails)
	for _, key := range dl.Allkeys {
		if key == d.Date {
			doesKeyExist = true
		}
	}

	if doesKeyExist != true {
		dl.Allkeys = append(dl.Allkeys, d.Date)
		stringValue := toJSON(dl.Allkeys)
		dl.Prefs.SetString(ALLKEY, stringValue)
	}
	contentValue := d.ConvertDiaryToString()
	//log.Println("Content to save :: ", d.Date, contentValue)
	dl.Prefs.SetString(d.Date, contentValue)
	//log.Println("Content Saved ...")
}

// Convert diary struct to JSON string ...
func (d *Diary) ConvertDiaryToString() string {
	stringValue := toJSON(d)
	return stringValue
}

// Fetch the content for a given date ...
func (dl *DiaryList) GetSpecificDateContent(dt string) *Diary {
	d := &Diary{}
	content := dl.Prefs.String(dt)
	if content != "" {
		//log.Println("Specified Date Content :: ", content)
		fromJSON([]byte(content), d)
		return d
	}
	return nil
}

// Fetch the stored keys for selected month ...
func (dl *DiaryList) LoadSelectedKeys(givenMonth string) []string {
	var selectedKeys []string
	keyDetails := dl.Prefs.String(ALLKEY)
	keys := *LoadStoredKeys(keyDetails)
	for _, key := range keys {
		keyMonth := strings.Split(key, "-")[1]
		if keyMonth == givenMonth {
			selectedKeys = append(selectedKeys, key)
		}
	}
	return selectedKeys
}

// Fetch all the stored keys ...
func LoadStoredKeys(keyDetails string) *[]string {
	storedKeys := &[]string{}
	if keyDetails != "" {
		byteData := []byte(keyDetails)
		fromJSON(byteData, storedKeys)
	}
	//log.Println("Stored Keys :: ", storedKeys)
	return storedKeys
}

// convert given type to JSON string ...
func toJSON(content interface{}) string {
	byteData, err := json.Marshal(content)
	if err != nil {
		log.Fatal("An error occurred :: toJSON :: ", err)
	}
	return string(byteData)
}

// Unmarshal the given JSON ...
func fromJSON(byteData []byte, structType interface{}) interface{} {
	//log.Println("fromJson :: byteData :: ", string(byteData))
	//log.Println("fromJSON :: structType :: ", structType)
	err := json.Unmarshal(byteData, structType)
	if err != nil {
		log.Fatal("An error occurred :: fromJSON :: ", err)
	}
	return structType
}
