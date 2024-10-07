package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Message    string
}
type Number interface {
	int | int16 | int32 | int64 | float32 | float64
}

func main() {
	//3
	response, err := GetRequest("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
	//2
	//output, err := createErrorMethod2(0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(output)
	//1
	//x := Sum(2, 7)
	//fmt.Printf("%d\n", x)
	//
	//y := Sum(2.5, 7.5)
	//fmt.Printf("%f\n", y)
}
func (error HttpError) Error() string {
	return fmt.Sprintf("Http error occurred: %d %s", error.StatusCode, error.Message)
}

func NewHttpError(statusCode int, message string) HttpError {
	return HttpError{StatusCode: statusCode, Message: message}
}

func GetRequest(url string) (string, error) {
	if len(url) == 0 {
		return "", NewHttpError(400, "Url is empty")
	}
	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "Error occurred")
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", NewHttpError(200, "Body is empty")
	}
	return string(responseBody), nil
}
func Sum[T Number](a, b T) T {
	return a + b
}
func createErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("Number is not valid")
	}
	return number * 5, nil
}

func createErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("Number is not valid: %d", number)
	}
	return number * 5, nil
}
