package reviewlibrary

import (
	"testing"
)

func TestCheckinTimebtn(t *testing.T) {
	expectedText := `您好，請問今天大概會幾點抵達入住呢~
	我們的櫃台開放時間是08:00~20:30，期間可先辦理入住及寄放行李唷！
	(提早到達可先拿房卡，15:00可進房)
	若會超過櫃檯服務時間抵達，請務必提前與我們取得聯絡，我們會為您準備自助入住流程
	LINE: @loosha520`

	result := CheckinTimebtn()
	if result != expectedText {
		t.Errorf("CheckinTimebtn() returned unexpected text.\nExpected:\n%s\n\nGot:\n%s\n", expectedText, result)
	}
}

func TestCheckinTimebtn2(t *testing.T) {
	expectedText := `您好，請問今天大概會幾點抵達入住呢~
	我們的櫃台開放時間是08:00~21:30，期間可先辦理入住及寄放行李唷！
	(提早到達可先拿房卡，15:00可進房)
	若會超過櫃檯服務時間抵達，請務必提前與我們取得聯絡，我們會為您準備自助入住流程
	LINE: @loosha520`

	result := CheckinTimebtn2()
	if result != expectedText {
		t.Errorf("CheckinTimebtn2() returned unexpected text.\nExpected:\n%s\n\nGot:\n%s\n", expectedText, result)
	}
}

func TestUploadidbtn(t *testing.T) {
	expectedText := `您好，要麻煩您上傳證件的照片(身分證、駕照、健保卡擇一)，以利幫您登記入住，稍後會由平台為您請款，確認款項後即會發送自助入住流程給您喔~`

	result := Uploadidbtn()
	if result != expectedText {
		t.Errorf("Uploadidbtn() returned unexpected text.\nExpected:\n%s\n\nGot:\n%s\n", expectedText, result)
	}
}

func TestBreakfastbtn(t *testing.T) {
	expectedText := `早餐一份為$160
	供應時間為時間為每日08:00-10:00，可在辦理入住時告知櫃台人員加購早餐`

	result := Breakfastbtn()
	if result != expectedText {
		t.Errorf("Breakfastbtn() returned unexpected text.\nExpected:\n%s\n\nGot:\n%s\n", expectedText, result)
	}
}
