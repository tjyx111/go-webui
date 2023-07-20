package server

import (
	"lbswebui/public"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/tebeka/selenium"
)

const (
	SYS_LOG_URL     = public.SERVER + "/sys/logger/systemlog"
	SYS_TIME_URL    = public.SERVER + "/sys/maintain/time"
	SYS_STORAGE_URL = public.SERVER + "/sys/maintain/storage#"
	SYS_BACKUP_URL  = public.SERVER + "/sys/maintain/backup_manage"
	DOWN_LOAD_PATH  = "C:\\Users\\Lenovo\\Downloads"
	SEQ             = "44164D56-6B73-C47B-E5C1-32AA63FA900C"
)

func TestLocalTime(t *testing.T) {
	wd, err := public.SwitchToPage(SYS_TIME_URL)
	assert.Equal(t, err, nil)
	// 点击下拉框
	selectTime := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/fieldset[2]/table/tbody/tr[1]/td[2]/select"
	err = public.WaitAndDo(wd, public.EleClickByXpath, selectTime)
	assert.Equal(t, err, nil)

	// 选择本地时间
	localTime := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/fieldset[2]/table/tbody/tr[1]/td[2]/select/option[1]"
	err = public.WaitAndDo(wd, public.EleClickByXpath, localTime)
	assert.Equal(t, err, nil)

	// 点击应用
	apply := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/a/span"
	err = public.WaitAndDo(wd, public.EleClickByXpath, apply)
	assert.Equal(t, err, nil)

	// 检查是否设置成功
	timeSetSuccess := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/div/p"
	timeSetSuccessText := "时间设置成功！"
	err = public.IsDisplayedWithTextByXpath(wd, timeSetSuccess, timeSetSuccessText)
	assert.Equal(t, err, nil)
}

func TestNtpTime(t *testing.T) {
	wd, err := public.SwitchToPage(SYS_TIME_URL)
	assert.Equal(t, err, nil)
	// 点击下拉框
	selectTime := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/fieldset[2]/table/tbody/tr[1]/td[2]/select"
	err = public.WaitAndDo(wd, public.EleClickByXpath, selectTime)
	assert.Equal(t, err, nil)

	// 选择ntp时间
	localTime := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/fieldset[2]/table/tbody/tr[1]/td[2]/select/option[2]"
	err = public.WaitAndDo(wd, public.EleClickByXpath, localTime)
	assert.Equal(t, err, nil)

	// 输入ntp地址
	ntp := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/form/fieldset[2]/table/tbody/tr[3]/td[2]/input"
	ntpUrl := "ntp.ntsc.ac.cn"
	err = public.WaitAndDo(wd, public.EleSendKeysByXpath, ntp, ntpUrl)
	assert.Equal(t, err, nil)

	// 点击应用
	apply := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr/td/a/span"
	err = public.WaitAndDo(wd, public.EleClickByXpath, apply)
	assert.Equal(t, err, nil)

	// 检查是否设置成功
	timeSetSuccess := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/div/p"
	timeSetSuccessText := "时间设置成功！"
	err = public.IsDisplayedWithTextByXpath(wd, timeSetSuccess, timeSetSuccessText)
	assert.Equal(t, err, nil)
}

func TestStorage(t *testing.T) {
	wd, err := public.SwitchToPage(SYS_STORAGE_URL)
	assert.Equal(t, err, nil)
	// 点击清理日志
	clearLog := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/table/tbody/tr/td/table/tbody/tr[2]/td[5]/a"
	err = public.WaitAndDo(wd, public.EleClickByXpath, clearLog)
	assert.Equal(t, err, nil)

	// 点击确认清理
	err = public.WaitAlert(wd, public.ACCEPT, "")
	assert.Equal(t, err, nil)
}

func TestBackupNoNetCfg(t *testing.T) {
	wd, err := public.SwitchToPage(SYS_BACKUP_URL)
	assert.Equal(t, err, nil)

	// 点击备份
	backup := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/form/table/tbody/tr/td[3]/a[1]/span"
	err = public.WaitAndDo(wd, public.EleClickByXpath, backup)
	assert.Equal(t, err, nil)

	// 确认备份
	err = public.WaitAlert(wd, public.ACCEPT, "")
	assert.Equal(t, err, nil)

	// 等待备份完成
	backupSuccess := "/html/body/table[1]/tbody/tr/td[4]/table[2]/tbody/tr/td[2]/table/tbody/tr/td[2]/div[1]/span"
	backupSuccessText := "执行完毕！"
	err = public.WaitAndDo(wd, public.IsDisplayedWithTextByXpath, backupSuccess, backupSuccessText)
	assert.Equal(t, err, nil)

	// 下载备份文件
	lastBackup := "table#table_item>tbody>tr:nth-of-type(2)>td.tbl_action>a:nth-of-type(2)"
	err = public.EleClickByCssSelector(wd, lastBackup)
	assert.Equal(t, err, nil)

	// 不下载
	err = public.WaitAlert(wd, public.DISMISS, "")
	assert.Equal(t, err, nil)

	// 点击确定
	err = public.WaitAlert(wd, public.ACCEPT, "")
	assert.Equal(t, err, nil)
}

func TestBackupLast(t *testing.T) {
	wd, err := public.SwitchToPage(SYS_BACKUP_URL)
	assert.Equal(t, err, nil)
	lastBackup := "table#table_item>tbody>tr:nth-of-type(2)>td.tbl_action>a:nth-of-type(2)"
	ele, err := wd.FindElement(selenium.ByCSSSelector, lastBackup)
	assert.Equal(t, err, nil)
	text, err := ele.Text()
	assert.Equal(t, err, nil)
	assert.Equal(t, text, "下载")
}
