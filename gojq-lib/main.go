//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/itchyny/gojq"
	"github.com/mattn/go-runewidth"
)

var gojqVersion = "unknown"

type OutputMode int

const (
	OutputModeJSON OutputMode = iota
	OutputModeJSONPretty
	OutputModeRaw
)

func NewOutputMode(mode string) OutputMode {
	switch mode {
	case "json":
		return OutputModeJSON
	case "json-pretty":
		return OutputModeJSONPretty
	case "raw":
		return OutputModeRaw
	default:
		return OutputModeJSON
	}
}

func stringParseError(queryStr string, err *gojq.ParseError) string {
	line := 0
	lineBuf := ""
	for i := 0; i < min(len(queryStr), err.Offset); i++ {
		if queryStr[i] == '\n' {
			line++
			lineBuf = ""
		} else {
			lineBuf += string(queryStr[i])
		}
	}
	col := runewidth.StringWidth(lineBuf)
	prefix := fmt.Sprintf("%d: ", line)
	strings.Repeat(" ", col)
	return fmt.Sprintf("%s%s\n%s%s", prefix, lineBuf, strings.Repeat(" ", col+len(prefix)-len(err.Token))+strings.Repeat("^", len(err.Token))+"--- ", err.Error())
}

func execJq(queryStr string, data any, outputMode OutputMode, maxIter int) string {
	query, err := gojq.Parse(queryStr)
	if err != nil {
		if err, ok := err.(*gojq.ParseError); ok {
			return stringParseError(queryStr, err)
		}
		return err.Error()
	}
	iter := query.Run(data)
	curIter := -1
	out := ""
	for {
		curIter++
		if maxIter > 0 && curIter >= maxIter {
			break
		}
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			out += err.Error() + "\n"
			break
		}
		serialized := []byte{}
		switch outputMode {
		case OutputModeJSON:
			serialized, err = json.Marshal(v)
			if err != nil {
				return err.Error()
			}
			out += string(serialized) + "\n"
		case OutputModeJSONPretty:
			serialized, err = json.MarshalIndent(v, "", "\t")
			if err != nil {
				return err.Error()
			}
			out += string(serialized) + "\n"
		case OutputModeRaw:
			if str, ok := v.(string); ok {
				out += str + "\n"
			} else {
				out += fmt.Sprintf("%v\n", v)
			}
		}
	}

	return strings.TrimSpace(out)
}

func gojqWrapper() js.Value {

	wrapper := make(map[string]interface{})
	wrapper["exec"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 4 {
			return "Invalid arguments"
		}
		var data any
		json.Unmarshal(
			[]byte(args[1].String()),
			&data,
		)
		return execJq(args[0].String(), data, NewOutputMode(args[2].String()), args[3].Int())
	})

	wrapper["prepare"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return "Invalid arguments"
		}
		var data any
		json.Unmarshal(
			[]byte(args[0].String()),
			&data,
		)
		return js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) < 3 {
				return "Invalid arguments"
			}
			return execJq(args[0].String(), data, NewOutputMode(args[1].String()), args[2].Int())
		})
	})

	wrapper["gojqVersion"] = js.ValueOf(gojqVersion)

	return js.ValueOf(wrapper)
}

func main() {
	js.Global().Set("gojq", gojqWrapper())
	<-make(chan struct{})
}
