Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

в Go ключевое слово `defer` используется для отложенного выполнения функции до того момента, когда окружающая функция завершит свое выполнение
в методе test() x объявляется, как возвращаемое значение.
если мы поставим x в return, то defer выполнится после return, и мы вернём неактуальное значение, как 
в методе anotherTest()

fifo


```
