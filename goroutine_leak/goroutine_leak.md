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