### Select 

- select statement is like a switch

```
select{
case <-ch1:
    // some block of statements
case <-ch2:
    // another block of statements
case ch3 <- struct{}{}:
    // yet another block of statements
}
```

- Each case specifies communication
- All channel operations are considered simultaneously

- Select waits until some case is ready to proceed
- When one of the channels is ready, that operation will proceed

Select is also very helpful in implementing:
- Timeouts
- Non-blocking communication

#### Timeout waiting on channel

```
select {
    case v:= <-ch:
        fmt.Println(v)
    case <- time.After(3*time.Second):
        fmt.Println("timeout")
}
```

- select waits until there is an event on channel ch or until timeout is reached

#### None blocking communication

```
select {
case m:= <-ch:
    fmt.Println("received message", m)
default:
    fmt.Println("no message received")
}
```

This code allows us to send or receive on a channel, but avoid blocking if the channel is not ready.

- default allows you to exit a select block without blocking

#### Empty Select

- Empty select state will <span style="color:red"> block forever </span> : 
```
Select {}
```

- Select on nil channel will <span style="color:red"> block forever </span>.
```
var ch chan string
select {
    case v := <-ch :

    case ch <-v :
}
```