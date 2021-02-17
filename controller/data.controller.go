package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/api/sheets/v4"

	"github.com/Goolgae/GSM_data/server/util"
)

//Data
type Data struct {
	Major        map[string]float64 `json:"major" xml:"major"`
	Reason       map[string]float64 `json:"reason" xml:"reason"`
	StudyMethod  map[string]float64 `json:"studyMethod" xml:"studyMethod"`
	StudyTime    map[string]float64 `json:"studyTime" xml:"studyTime"`
	LikeLanguage map[string]float64 `json:"likeLanguage" xml:"likeLanguage"`
}

func GetDataAboutTrandTaste(c echo.Context) error {
	spreadsheetId := "1gluPeTYankLaAElZfz4FxkpSjeBlhuSD5gnIOF_Lc20"
	readRange := "설문지 응답 시트1!C2:G"

	sheetData := util.GetSheet(spreadsheetId, readRange)
	totalLen := len(sheetData.Values)

	if len(sheetData.Values) == 0 {
		errorCode, _ := json.Marshal(map[string]string{"error": "have a no data"})
		return c.String(http.StatusNotFound, string(errorCode))
	}

	// major := getMajor(sheetData, totalLen)
	// reason := getReason(sheetData, totalLen)
	// studyMethod := getStudyMethod(sheetData, totalLen)
	// studyTime := getStudyTime(sheetData, totalLen)
	// likeLanguage := getLikeLanguage(sheetData, totalLen)

	d := &Data{
		Major:        getMajor(sheetData, totalLen),
		Reason:       getReason(sheetData, totalLen),
		StudyMethod:  getStudyMethod(sheetData, totalLen),
		StudyTime:    getStudyTime(sheetData, totalLen),
		LikeLanguage: getLikeLanguage(sheetData, totalLen),
	}
	return c.JSON(http.StatusOK, d)
}
func getStudyMethod(data *sheets.ValueRange, totalLen int) map[string]float64 {
	method := map[string]float64{
		"Google":            0,
		"QuestionToSenior":  0,
		"OfficialDocument":  0,
		"QuestionToFriends": 0,
		"Lecture":           0,
		"Book":              0,
		"QuestionToTeacher": 0,
		"Etc":               0,
	}
	for _, row := range data.Values {
		methodData := row[2]
		switch methodData {
		case "모르는게 있으면 구글링해봐요!":
			method["Google"] += 1
		case "선배들에게 궁금한걸 물어봐요":
			method["QuestionToSenior"] += 1
		case "공식문서를 보며 개념을 정리해요":
			method["OfficialDocument"] += 1
		case "친구들에게 모르는 부분을 물어봐요":
			method["QuestionToFriends"] += 1
		case "온라인 강의를 보며 공부해요":
			method["Lecture"] += 1
		case "책을 보고 공부해요":
			method["Book"] += 1
		case "선생님께 질문":
			method["QuestionToTeacher"] += 1
		default:
			method["Etc"] += 1
		}
	}

	for key, _ := range method {
		method[key] /= float64(totalLen)
	}

	return method
}
func getStudyTime(data *sheets.ValueRange, totalLen int) map[string]float64 {
	time := map[string]float64{
		"One":   0,
		"Two":   0,
		"Three": 0,
		"Four":  0,
		"Six":   0,
	}

	for _, row := range data.Values {
		timeData := row[3]
		switch timeData {
		case "1~2":
			time["One"] += 1
		case "2~3":
			time["Two"] += 1
		case "3~4":
			time["Three"] += 1
		case "4~6":
			time["Four"] += 1
		case "6시간 이상":
			time["Six"] += 1
		}
	}
	for key, _ := range time {
		time[key] /= float64(totalLen)
	}
	return time
}
func getLikeLanguage(data *sheets.ValueRange, totalLen int) map[string]float64 {
	language := map[string]float64{
		"C/Cpp":  0,
		"C#":     0,
		"Js":     0,
		"Ts":     0,
		"Go":     0,
		"Java":   0,
		"Python": 0,
		"Swift":  0,
		"Kotlin": 0,
		"Etc":    0,
	}

	for _, row := range data.Values {
		languageData := row[4]
		switch languageData {
		case "c/cpp":
			language["C/Cpp"] += 1
		case "c#":
			language["C#"] += 1
		case "js":
			language["Js"] += 1
		case "ts":
			language["Ts"] += 1
		case "go":
			language["Go"] += 1
		case "java":
			language["Java"] += 1
		case "python":
			language["Python"] += 1
		case "swift":
			language["Swift"] += 1
		case "kotlin":
			language["Kotlin"] += 1
		default:
			language["Etc"] += 1
		}
	}

	for key, _ := range language {
		language[key] /= float64(totalLen)
	}
	return language
}
func getReason(data *sheets.ValueRange, totalLen int) map[string]float64 {
	reason := map[string]float64{
		"Money":          0,
		"Easy":           0,
		"Cool":           0,
		"Want":           0,
		"Recommendation": 0,
		"Etc":            0,
	}
	for _, row := range data.Values {
		reasonData := row[1]
		switch reasonData {
		case "돈이 될듯 하여서":
			reason["Money"] += 1
		case "멋있어서":
			reason["Cool"] += 1
		case "쉬워 보여서":
			reason["Easy"] += 1
		case "자신이 원하는 결과물이 있어서":
			reason["Want"] += 1
		case "선배가 추천해서":
			reason["Recommendation"] += 1
		default:
			reason["Etc"] += 1
		}
	}
	for key, _ := range reason {
		reason[key] /= float64(totalLen)
	}
	return reason
}
func getMajor(data *sheets.ValueRange, totalLen int) map[string]float64 {
	major := map[string]float64{
		"BE":       0,
		"FE":       0,
		"AI":       0,
		"IOS":      0,
		"Android":  0,
		"Game":     0,
		"IOT":      0,
		"Security": 0,
		"Etc":      0,
	}

	for _, row := range data.Values {
		majorData := row[0]
		switch majorData {
		case "백엔드":
			major["BE"] += 1
		case "프론트엔드":
			major["FE"] += 1
		case "안드로이드앱":
			major["Android"] += 1
		case "ios앱":
			major["IOS"] += 1
		case "게임개발":
			major["Game"] += 1
		case "인공지능":
			major["AI"] += 1
		case "iot":
			major["IOT"] += 1
		case "정보 보안":
			major["Security"] += 1
		default:
			major["Etc"] += 1
		}
	}
	for key, _ := range major {
		major[key] /= float64(totalLen)
	}

	return major
}
