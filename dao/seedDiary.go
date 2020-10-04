package dao

// Seed the test data ...
func (dl *DiaryList) SeedDiary() {
	seedDiary := []Diary{
		Diary{
			Title:   "Hello",
			Date:    "1-October-2020",
			Content: "LMAO",
		},
		Diary{
			Title:   "Meow",
			Date:    "29-September-2020",
			Content: "Wow",
		},
	}

	for _, seedDetails := range seedDiary {
		var doesKeyExist bool = false
		keyDetails := dl.Prefs.String(ALLKEY)
		dl.Allkeys = *LoadStoredKeys(keyDetails)
		for _, key := range dl.Allkeys {
			if key == seedDetails.Date {
				doesKeyExist = true
			}
		}

		if doesKeyExist != true {
			dl.Allkeys = append(dl.Allkeys, seedDetails.Date)
			stringValue := toJSON(dl.Allkeys)
			dl.Prefs.SetString(ALLKEY, stringValue)
		}
		contentValue := seedDetails.ConvertDiaryToString()
		dl.Prefs.SetString(seedDetails.Date, contentValue)
	}

}
