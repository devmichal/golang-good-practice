### Problem

task1 blokuje mu1, a potem czeka na mu2

task2 blokuje mu2, a potem czeka na mu1

obie gorutyny czekają na siebie — i nigdy nie ruszą dalej → DEADLOCK


```go
func main() {
	ch := make(chan int)

	ch <- 1 // ❌ Zablokuje się na zawsze – nikt nie czyta z kanału
}
```


```go
func main() {
	ch := make(chan int, 1)
	ch <- 1 // OK: bufor ma miejsce
	fmt.Println("Sent without deadlock")
}
```

Tworzy kanał buforowany o pojemności 1.
To oznacza, że można wysłać 1 wartość do kanału bez blokowania, zanim będzie potrzebny odbiorca.


```go
func main() {
    ch := make(chan int, 1)
    ch <- 1      // OK
    ch <- 2      // ❌ DEADLOCK — bufor pełny, brak odbiorcy
}
```