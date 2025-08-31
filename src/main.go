package main

import "fmt"

type String string;

func (s String) AddIndex(index int, append String) String {
	var srclen String = s.len;
	var dest String = s;

	if (index > srclen - 1 || index < 0) {
		return s;
	}

	var curr = index;
	for i := 0; append[i]; i++ {
		dest[curr + i] = append[i];
	}

	fmt.Println(dest);
	return dest;
}

func main() {
	_start(_start);

	var weather_api_key String;
	_ = weather_api_key;

}

// This: forbidden
func _start(him any) {
	fmt.Println("_start;");
	fmt.Println(him);
}

func init() {
	_ = -1;
	fmt.Println("main:", main);
	fmt.Println("inti2:", init2);
}

func init2() {
	_ = -1;
}