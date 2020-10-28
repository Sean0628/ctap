package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Sean0628/ctap/test"
)

const crontabContent = `# run the drupal cron process every hour of every day
0 * * * * /usr/bin/wget -O - -q -t 1 http://localhost/cron.php

# reset the contact form just after midnight
5 0 * * * /var/www/devdaily.com/bin/resetContactForm.sh

# db backup script
0 05 * * 1-5 root /var/www/db-backup.sh `

func TestGetParsedExpressionFromFileDisplay_Input(t *testing.T) {
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("-i@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := `0 * * * *: Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: At 12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: At 05:00 AM, Monday through Friday | root /var/www/db-backup.sh`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_InputLong(t *testing.T) {
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("--input@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := `0 * * * *: Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: At 12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: At 05:00 AM, Monday through Friday | root /var/www/db-backup.sh`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_noExistingFile(t *testing.T) {
	fileName := "dummy.txt"

	cmd := getRootCommand()
	opts := fmt.Sprintf("--input@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := fmt.Sprintf("ctap: failed to open %s: no such file or directory.", fileName)
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_emptyFile(t *testing.T) {
	fileName := test.WriteTmpCrontabFile("")
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("--input@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := "ctap: noting to be printed out. possibly input file does not have any contents."
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_Locale(t *testing.T) {
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-i@%s@-l@ja", fileName))
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := `0 * * * *: 毎時 | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: 次において実施12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: 次において実施05:00 AM、月曜日 から 金曜日 まで | root /var/www/db-backup.sh`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_24hrsFormat(t *testing.T) {
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-i@%s@--24-hour", fileName))
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := `0 * * * *: Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: At 00:05 | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: At 05:00, Monday through Friday | root /var/www/db-backup.sh`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_DowStartsAtOne(t *testing.T) {
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-i@%s@-d", fileName))
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := `0 * * * *: Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: At 12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: At 05:00 AM, Sunday through Thursday | root /var/www/db-backup.sh`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_Output(t *testing.T) {
	inputFileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(inputFileName)

	outputFileName := "../tmp/crontab.test.txt"
	defer test.RemoveTmpCrontabFile(outputFileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-i@%s@-o@%s", inputFileName, outputFileName))
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := fmt.Sprintf("ctap: the results are printed out to %s.", outputFileName)
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))

	expectedOutput = `0 * * * *: Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: At 12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: At 05:00 AM, Monday through Friday | root /var/www/db-backup.sh`

	data, _ := ioutil.ReadFile(outputFileName)
	test.AssertResult(t, expectedOutput, strings.Trim(string(data), "\n "))
}

func TestGetParsedExpressionFromFileDisplay_OutputLong(t *testing.T) {
	inputFileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(inputFileName)

	outputFileName := "../tmp/crontab.test.txt"
	defer test.RemoveTmpCrontabFile(outputFileName)

	cmd := getRootCommand()
	result := test.RunCmd(cmd, fmt.Sprintf("-i@%s@--output@%s", inputFileName, outputFileName))
	if result.Error != nil {
		t.Error(result.Error)
	}

	expectedOutput := fmt.Sprintf("ctap: the results are printed out to %s.", outputFileName)
	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))

	expectedOutput = `0 * * * *: Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *: At 12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5: At 05:00 AM, Monday through Friday | root /var/www/db-backup.sh`

	data, _ := ioutil.ReadFile(outputFileName)
	test.AssertResult(t, expectedOutput, strings.Trim(string(data), "\n "))
}

func TestGetParsedExpressionFromFileDisplay_FormatMarkdown(t *testing.T) {
	formatType := "markdown"
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("-i@%s@-f@%s", fileName, formatType)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := `| Original | Translated | Command |
|---|---|---|
| 0 * * * * | Every hour | /usr/bin/wget -O - -q -t 1 http://localhost/cron.php |
| 5 0 * * * | At 12:05 AM | /var/www/devdaily.com/bin/resetContactForm.sh |
| 0 05 * * 1-5 | At 05:00 AM, Monday through Friday | root /var/www/db-backup.sh |`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionFromFileDisplay_FormatCsv(t *testing.T) {
	formatType := "csv"
	fileName := test.WriteTmpCrontabFile(crontabContent)
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("-i@%s@--format@%s", fileName, formatType)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}
	expectedOutput := `Original, Translated, Command
0 * * * *, Every hour, /usr/bin/wget -O - -q -t 1 http://localhost/cron.php
5 0 * * *, At 12:05 AM, /var/www/devdaily.com/bin/resetContactForm.sh
0 05 * * 1-5, At 05:00 AM; Monday through Friday, root /var/www/db-backup.sh`

	test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}
