package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	1 * 2 |> _ * 3 |> _ * 4 |> fmt.Println(_)
	1 + 2 |> _ * 3 |> _ != 1 + 2 * 3 |> fmt.Println(_)
	// 1 + 2 |> (_) * 3 |> fmt.Println(_) not allowed but automatically done

	_, err := os.Getwd() |> fmt.Println(_)
	if err != nil {
		log.Fatalln(err)
	}

	[]int{1, 2, 3, 4, 5} |> filter(_, func(i int, x int) bool {
		return x % 2 != 0
	}) |> transform(_, func(i int, x int) string {
		return fmt.Sprint(x)
	}) |> reduce(_, func(i int, x string, prev string) string {
		return x + prev
	}, "") |> fmt.Println(_)

	"\tHello, World\t" |> NewStringWrapper(_) |> _.Self().Unwrap() |> strings.TrimSpace(_) |> fmt.Println(_)
	"\tHello, World\t" |> NewStringWrapper(_) |> _.Self().s |> strings.TrimSpace(_) |> fmt.Println(_)

	"1;2;3" |> strings.Split(_, ";")[0] |> fmt.Println(_)
	0 |> strings.Split("1;2;3", ";")[_] |> fmt.Println(_)

	"3;2;1" |> strings.Split(_, ";")[1:2] |> fmt.Println(_)
	1 |> strings.Split("1;2;3", ";")[_:2] |> fmt.Println(_)
	2 |> strings.Split("1;2;3", ";")[1:_] |> fmt.Println(_)

	"%d\n" |> any(_) |> _.(string) |> fmt.Printf(_, 10)

	(fmt.Sprint(1) == "1") |> fmt.Println(_)
	fmt.Sprint(1) == "1" |> fmt.Println(_)

	m := make(map[string]bool)
	"\tHello, World!\t" |> strings.TrimSpace(_) |> m[_] = true
	fmt.Println(m)

	strings.Split("2;3", ";") |> append([]string{"1"}, _...) |> fmt.Println(_)

	"1" + "2" |> fmt.Sprint(12) == _ |> fmt.Println(_)
	"1" |> fmt.Sprint(1) == _ |> fmt.Println(_)
	"1" |> _ + "0!" |> fmt.Println(_)

	`[
		"Ciao",
		"Bella",
		"Hello",
		"World"
	]` |> []byte(_) |> unmarshalJSON[[]string](_) |> fmt.Println(_)

	fmt.Println("----------------------")
	testAST()
}

func unmarshalJSON[T any](data []byte) (T, error) {
	var x T
	err := json.Unmarshal(data, &x)
	return x, err
}
