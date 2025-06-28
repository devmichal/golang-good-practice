# Popularne przypadki dla concurrency issue (problem z współbieżnością)

### zamknięciem (closure) 

```go
for _, identity := range identities {
	index++ // zmieniasz index

	go func() {
		sqlIdentities[index] = ...
	}()
}
```

index jest zmienną z zewnętrznego zakresu (for-loop), i gorutyna "widzi" ją w momencie wykonania, a nie w momencie utworzenia gorutyny.

Efekt: wszystkie gorutyny mogą używać tej samej, ostatecznej wartości index – np. 1000 – zamiast tej, którą index miał w momencie, gdy gorutyna została zainicjowana.

#### Rozwiązanie

```go
go func(i ExtIdentities, index int) {
// używasz kopii indexa, zrobionej w momencie tworzenia gorutyny
}(identity, index)
```
Zapewniasz, że każda gorutyna ma własną lokalną kopię index

### Race Condition (warunek wyścigu)

wiele gorutyn jednocześnie odwołuje się do tej samej zmiennej

```go
count := 0

for i := 0; i < 1000; i++ {
	go func() {
		count++ // RACE: równoczesny zapis
	}()
}
```

aby uniknąć nadpisywanie przez wiele gorutyn:

```go
count := 0
var mu sync.Mutex
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
	go func() {
        defer wg.Done()
        mu.Lock()
		count++ // RACE: równoczesny zapis
		mu.Unlock()
	}()
}
wg.Wait()
```

### Goroutine Leak

Problem: gorutyna zostaje uruchomiona, ale nigdy się nie kończy, np. bo czeka na dane z kanału, których nikt nie wyśle.

```go
func listen(ch <-chan int) {
    for val := range ch {
        fmt.Println("Got:", val)
    }
}

func main() {
    ch := make(chan int)

    go listen(ch) // Gorutyna czeka na dane, ale ich nie dostanie

    // main kończy się bez zamknięcia kanału ani wysłania czegokolwiek
    time.Sleep(1 * time.Second)
}
```

#### Rozwiązanie 1) Zamknięcie kanału

```go
func listen(ch <-chan int) {
    for val := range ch {
        fmt.Println("Got:", val)
    }
    fmt.Println("Channel closed, exiting goroutine.")
}

func main() {
    ch := make(chan int)

    go listen(ch)

    ch <- 1
    ch <- 2
    close(ch) // Zamyka kanał – listen kończy pętlę
}
```

#### Rozwiązanie 2) Użycie context.Context do anulowania

```go
func listen(ctx context.Context, ch <-chan int) {
    for {
        select {
        case val := <-ch:
            fmt.Println("Got:", val)
        case <-ctx.Done():
            fmt.Println("Context cancelled, exiting goroutine.")
            return
        }
    }
}

func main() {
    ch := make(chan int)
    ctx, cancel := context.WithCancel(context.Background())

    go listen(ctx, ch)

    ch <- 42
    time.Sleep(500 * time.Millisecond)
    cancel() // kończy gorutynę przez context
}

```

### Deadlock

### Panic w gorutynie

# Narzędzia pomocnicze


### Wykrywa race condition

```go
go run -race
```

### Analiza gorutyn i CPU
```go
pprof, trace
```