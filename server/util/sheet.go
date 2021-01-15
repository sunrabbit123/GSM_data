package util

import (
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error)
	}
}
func GetSheet(spreadsheetId string, readRange string) *sheets.ValueRange {
	data, err := ioutil.ReadFile("secret.json")
	checkError(err)

	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	checkError(err)

	client := conf.Client(context.TODO())
	srv, err := sheets.New(client)
	checkError(err)

	// spreadsheetId := "1gluPeTYankLaAElZfz4FxkpSjeBlhuSD5gnIOF_Lc20"
	// readRange := "설문지 응답 시트1!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	return resp
}
