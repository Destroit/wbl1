### 1. Какой самый эффективный способ конкатенации строк?
Тип данных strings.Builder позволяет осуществить конкатенацию строк при помощи методов Builder.Write, сводящих к минимуму операции копирования
```go
var sb strings.Builder
sb.WriteString("123")
sb.WriteString("456")
fmt.Println(sb.String())
```
### 2. Что такое интерфейсы, как они применяются в Go?
Интерфейсы представляют собой абстракцию поведения других типов. В Go, интерфейс содержит в себе набор методов. Если у типа данных есть все методы из набора методов интерфейса, и их описания совпадают, то тогда можно сказать, что он реализует интерфейс. Это означает что значения такого типа данных могут быть присвоены переменной, типом данных которой является интерфейс.
### 3.Чем отличаются RWMutex от Mutex?
у Mutex два метода для блокировки доступа к общим данным - Lock() и Unlock(). У RWMutex в свою очередь есть дополнительные методы RLock() и RUnlock(). У нас может быть несколько параллельных процессов чтения данных (RLock()) или один процесс записи (Lock()), но не одновременно. Если использован Lock(), то другие вызовы Lock() и RLock() блокируются до Unlock(). А если исопльзован RLock(), то блокируются только вызовы Lock(), а вызовы RLock() - нет.
### 4. Чем отличаются буферизированные и не буферизированные каналы?
Небуферризированный канал блокируется для записи пока записанное в него значение не будет прочитано. Буфферизированный блокируется для записи когда буфер полон. Буфер освобождается по мере прочтения данных из канала. 
### 5. Какой размер у структуры struct{}{}?
0 байт.
### 6. Есть ли в Go перегрузка методов или операторов?
В Go нет перегрузки методов и операторов.
### 7. В какой последовательности будут выведены элементы map[int]int?
В случайной последовательности.
### 8. В чем разница make и new?
При помощи make можно создать только каналы, мапы и слайсы и возвращает значение. new выделяет память для любых типов и возвращает указатель на неё.
### 9. Сколько существует способов задать переменную типа slice или map?
slice:
```go
var sl []int
sl := []int{1, 2, 3}
sl := make([]int, 0, 5)
sl := new([]string)
```
map:
```go
var m map[string]int 
m := map[string]int {
    "one": 1,
    "two": 2,
}
m := make(map[string]int, 5)
m := new(map[string]int)
```
### 10. Что выведет данная программа и почему?
```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
Программа выведет\
1\
1\
Так как изменяется копия указателя p
### 11. Что выведет данная программа и почему?
```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
4\
2\
3\
0\
1\
Затем произойдет Deadlock потому что sync.WaitGroup в горутине является копией, из-за чего счётчик оригинала никогда не обнулится.
### 12. Что выведет данная программа и почему?
```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
0, в теле if в новой области видимости создана n, которая перекрывает n, созданную вне if.
### 13. Что выведет данная программа и почему?
```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
[100 2 3 4 5]\
Через копию среза можно изменить значение оригинального массива, но изменения ёмкости и длины копии среза не влияют на оригинальный срез
### 14. Что выведет данная программа и почему?
```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
[b b a][a a]\
После append в функции слайс перестал указывать на оригинальный срез, так как изменилась его ёмкость, из-за чего последующие операции не влияют на оригинальный массив