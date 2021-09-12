### Pipeline

- Process streams or batches of data
- Stage - take data in, perform an operation on it, and send the data out

### Stages

- Separate the concerns of each stage
- Process individual stage concurrently
- A stage could consume and return the same state
```
square(in <-chan int) <- chan int{

}
```
- This enables composility of a pipeline

#### Pipelines summary
- Pipelines are used to process Streams or Batches of Data
- Pipelines enables us to make an efficient use of I/O and multiple CPU cores

### Real pipelines 
- Real pipelines: Receivetr stages may only need a subset of values to make progress
- A stage can exit early because an inbound value represents an error in an earlier stage
- Receiver should not have to wait for the remaining values to arrive
- We want earlier stages to stop producing values that later stages don't need. 


Example of faulty pipeline:
- Main goroutine just receives on value
- Abandons the inbound channel from merge
- Merge goroutines will be blocked on channel send operation
- Square and generator goroutines will also be blocked on send
===> GOROUTINE LEAK

### Cancellation of a goroutine

- Pass a read-only 'done' channel to goroutine
- Close the channel, to send broadcast signal to all goroutines
- On receiving the signal on done channel, Goroutines needs to abandon
their work and terminate
- We use `select` to send/receive operation on channel pre-emptible

```
select{
case out <- n:
case <- done:
    return
}
```