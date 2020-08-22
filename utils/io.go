package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

var scanner = bufio.NewScanner(os.Stdin)

// WriteLine write a message.
func WriteLine(a ...interface{}) {
	fmt.Println(a...)
}

// Writef write a formated message.
func Writef(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// ReadLine read a line of user input from console.
// If the input text is 'quit', exit the application.
func ReadLine() string {
	for scanner.Scan() {
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return ""
}

// PrintJSONTags prints all json tags of an object.
func PrintJSONTags(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if jsonTag := t.Field(i).Tag.Get("json"); jsonTag != "" {
			WriteLine("  ", jsonTag)
		}
	}
}

// PrintObject prints all data of an object.
func PrintObject(obj interface{}) {
	WriteLine()
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if jsonTag := t.Field(i).Tag.Get("json"); jsonTag != "" {
			Writef("  %-30v%v\n", jsonTag, v.Field(i).Interface())
		}
	}
}
