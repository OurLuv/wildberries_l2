Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]

Слайс состоит из указателя на массив, длины сегмента(len) и его емкости(cap)

В этом примере i[0] = "3" изменяет значения в обоих слайсах(s, i), но так как они ссылаются на один базовый массив,
данные меняются в обоих местах.
После первого append произошло выделение нового массива, и теперь переменные s и i указывают на разные массивы, и i[1] = "5" уже не изменит s.

```