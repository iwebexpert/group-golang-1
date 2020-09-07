package main

import (
	"fmt"
	"service"
	"strings"
)

func printResult(v []string) {
	if len(v) < 1 {
		fmt.Println("Match isnt found")
		return
	}

	fmt.Println(fmt.Sprintf("Result is %s", strings.Join(v, ", ")))
}

func main() {
	v := service.SearchV1("yandex", []string{"https://yandex.ru", "http://ya.ru"})
	printResult(v)
	v = service.SearchV1("Google", []string{"https://mathiasbynens.be", "https://cnn.com", "https://www.tutorialspoint.com/mvc_framework/mvc_framework_introduction.htm"})
	printResult(v)
}
