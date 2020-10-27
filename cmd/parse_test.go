package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Sean0628/ctap/test"
)

const cronExpression = "0 05 * * 1-5"

func TestGetParsedExpressionDisplay(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprint(cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := "0 05 * * 1-5: At 05:00 AM, Monday through Friday"
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplay_Locale(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-l@ja@%s", cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := "0 05 * * 1-5: 次において実施05:00 AM、月曜日 から 金曜日 まで"
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplay_LocaleLong(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("--locale@ja@%s", cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := "0 05 * * 1-5: 次において実施05:00 AM、月曜日 から 金曜日 まで"
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplay_24hrsFormat(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("--24-hour@%s", cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := "0 05 * * 1-5: At 05:00, Monday through Friday"
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplay_DowStartsAtOne(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-d@%s", cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := "0 05 * * 1-5: At 05:00 AM, Sunday through Thursday"
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplay_DowStartsAtOneLong(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("--dow-starts-at-one@%s", cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := "0 05 * * 1-5: At 05:00 AM, Sunday through Thursday"
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplay_Output(t *testing.T) {
	fileName := "../tmp/crontab.test.txt"
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-o@%s@%s", fileName, cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := fmt.Sprintf("ctap: the results are printed out to %s.", fileName)
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))

	expectedOutput = "0 05 * * 1-5: At 05:00 AM, Monday through Friday"
	data, _ := ioutil.ReadFile(fileName)
	test.AssertResult(t, expectedOutput, strings.Trim(string(data), "\n "))
}

func TestGetParsedExpressionDisplay_OutputLong(t *testing.T) {
	fileName := "../tmp/crontab.test.txt"
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("--output@%s@%s", fileName, cronExpression))
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := fmt.Sprintf("ctap: the results are printed out to %s.", fileName)
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))

	expectedOutput = "0 05 * * 1-5: At 05:00 AM, Monday through Friday"
	data, _ := ioutil.ReadFile(fileName)
	test.AssertResult(t, expectedOutput, strings.Trim(string(data), "\n "))
}
