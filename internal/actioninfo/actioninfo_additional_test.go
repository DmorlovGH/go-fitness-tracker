package actioninfo

import (
	"bytes"
	"errors"
	"log"
	"reflect"
	"strings"
	"testing"
)

// recoveryTestParser позволяет проверить порядок вызовов при ошибках.
type recoveryTestParser struct {
	current string
	calls   []string
}

func (p *recoveryTestParser) Parse(data string) error {
	p.calls = append(p.calls, "Parse:"+data)
	p.current = data

	if data == "bad-parse" {
		return errors.New("parse failure")
	}
	return nil
}

func (p *recoveryTestParser) ActionInfo() (string, error) {
	p.calls = append(p.calls, "ActionInfo:"+p.current)

	if p.current == "bad-report" {
		return "", errors.New("report failure")
	}
	return "готово", nil
}

func TestInfoContinuesAfterErrors(t *testing.T) {
	var logs bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&logs)
	defer log.SetOutput(originalOutput)

	parser := &recoveryTestParser{}
	Info([]string{"bad-parse", "bad-report", "valid"}, parser)

	wantCalls := []string{
		"Parse:bad-parse",
		"Parse:bad-report",
		"ActionInfo:bad-report",
		"Parse:valid",
		"ActionInfo:valid",
	}

	if !reflect.DeepEqual(parser.calls, wantCalls) {
		t.Errorf("порядок вызовов: получено %v, ожидалось %v",
			parser.calls, wantCalls)
	}

	for _, message := range []string{"parse failure", "report failure"} {
		if !strings.Contains(logs.String(), message) {
			t.Errorf("в логе отсутствует ошибка %q", message)
		}
	}
}
