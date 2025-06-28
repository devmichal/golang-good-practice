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

### Deadlock

### Panic w gorutynie

# Narzędzia pomocnicze


### Wykrywa race condition

```go
go test -race ./...
```

### Analiza gorutyn i CPU
```go
go vet ./...
pprof, trace
```