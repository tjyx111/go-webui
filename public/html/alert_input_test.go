package html

import (
	"fmt"
	"lbswebui/public"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAlertSwitch(t *testing.T) {
	wd, err := public.StartChromDriver()
	assert.Equal(t, err, nil)

	currDir, err := os.Getwd()
	assert.Equal(t, err, nil)

	url := currDir + string(os.PathSeparator) + "alert_input_test.html"
	err = wd.Get(url)
	assert.Equal(t, err, nil)

	// 点击
	clickCss := "input#clicktest"
	err = public.EleClickByCssSelector(wd, clickCss)
	assert.Equal(t, err, nil)

	text, err := wd.AlertText()
	assert.Equal(t, err, nil)
	fmt.Println(text)

	// 进入弹窗
	err = wd.SetAlertText("12345678")
	assert.Equal(t, err, nil)

	// 点击弹窗
	err = wd.AcceptAlert()
	assert.Equal(t, err, nil)

	err = wd.Quit()
	assert.Equal(t, err, nil)
}
