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