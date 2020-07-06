package serializer

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"chaos-mesh/matrix/pkg/utils"
)

type StatementSerializer struct{}

func (s StatementSerializer) Dump(value interface{}, target string) error {
	var lines []string
	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	lines, err = valueToLines(value)
	if err != nil {
		return err
	}
	_, err = f.WriteString(strings.Join(lines, "\n"))
	return err
}

func itemToLine(value interface{}) (string, error) {
	if valueMap, ok := value.(map[string]interface{}); ok {
		templateValue, templateOk := valueMap["template"]
		valueValue, valueOk := valueMap["value"]
		if templateOk && valueOk && len(valueMap) == 2 {
			stringTemplateValue, templateOk := templateValue.(string)
			if templateOk {
				return fmt.Sprintf(stringTemplateValue, valueValue), nil
			} else {
				return "", errors.New(fmt.Sprintf("`template` not a string: %v", templateValue))
			}
		} else {
			return "", errors.New(fmt.Sprintf("keys should be `template` and `value`, got: %v", utils.Keys(valueMap)))
		}
	} else {
		return "", errors.New("")
	}
}

func valueToLines(value interface{}) ([]string, error) {
	var lines []string
	if valueList, ok := value.([]interface{}); ok {
		for _, v := range valueList {
			line, err := itemToLine(v)
			if err != nil {
				return nil, err
			}
			lines = append(lines, line)
		}
	} else {
		return nil, errors.New(fmt.Sprintf("`%v` is not a list", value))
	}

	if lines == nil {
		lines = append(lines, "")
	}
	return lines, nil
}