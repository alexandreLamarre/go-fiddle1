### Go Channel implementation

Internally, channels are represented by the
hchan struct:

```
type hchan struct{
    qcount uint //total data in the queue
    dataqsiz uint //size of the circular queue
    buf unsafe.Pointer //points to an array of dataqsiz elements
    elemsize uint16
    closed uint32
    elemtype *_type //element type
    sendx uint //send index
    recvx uint // recieve index
    recvq waitq //list of receive waiters
    sendq waitq // list of send receivers

    lock mutex // lock protects all fields in hchan

}

type waitq struct{
    first *sudog
    last *sudog
}

// sudog represents a g in a wait list such as sending/receiving
// on a channel

type sudog struct{
    g *g

    next *sudog
    prev *sudog
    elem unsafe.Pointer //data element (may point to stack)
    ...
    c *hchan //channel
}
```
So the operation 
```
chmake := make(chan int, 3)
```

translates to :
- allocating hchan struct on the heap
- return a pointer to it

Now since `ch` is a pointer it can be between functions for send and receive

### Send and Receive Channel implementation

Consider:
```
ch := make(chan int, 3)

//G1-goroutine
func G1(ch chan<- int>) {
    for _, v := range []int{1,2,3,4} {
        ch <- v
    }
}

//G2-goroutine
func G2(ch <-chan int>) {
    for v := range ch {
        fmt.Println(v)
    }
}
```

When the above executes, or in general:
- There is no memory share betwen goroutines
- Goroutines copy elements
- hchan is protected by mutex lock

The paradigm behind this behaviour is :
"Do not communicate by sharing memory; Instead share memory by communications"

### Case Study : Buffer full
In the above scenarion, imagine the channel buffer is full and a goroutine tries to send a value:
- Sender goroutine gets blocked, it is parked on sendQ
- Data will be saved in the elem filed of the sudog structure
- When receiver comes along, it deques the value from the buffer
- Enqueues the data from elem field to the buffer
- Pops the goroutine in sendq and puts it into runnable state


### Case Study : Buffer empty
In the above scenario, image the channel buffer is empty and a goroutine tries to read from the empty channel :

- Goroutine is blocked, it is parked into recvq.
- elem field of the sudog structure holds the reference to the stack variable of the receiver goroutine
- When sender comes along, Sender finds the goroutine in recvq
- Sender copies the data, into the stack variable on the receiver goroutine directly
- Pops the goroutine in recvq, and puts it into runnable state


## Send and Receive on an Unbuffered Channel

When sender goroutines wants to send values
- If there is corresponding receiver waiting in recvq
    - Sender will write the value directly into the receiver goroutine stack variable
    - Sender goroutine puts the receiver goroutine back to runnable state

- If there is no receiver goroutine in recvq
    - Sender gets parked into sendq
    - Data is saved in elem field in sudog struct
    - Receiver comes and copies the data
    - Puts the sender to runnable state again

When receiver goroutine wants to receive value
- If it finds a gouroutine in waiting in sendq
    - Receiver copies the value in elem field to its variable
    - Puts the sender goroutine to runnable state

- If there was no sender goroutine in sendq
    - Receiver gets parked into recvq
    - Reference to variable is saved in the elem field in sudog struct
    - Sender comes and copies the data directly to receiver stack variable
    - Puts receiver back to runnable state


### Summary
- hchan struct represents channel
- It contains a circular ring buffer and mutex lock
- Goroutines that gets blocked on send or recv are parked in sendq or recvq
- Go scheduler moves the blocked goroutines, out of OS thread
- Once channel operation is complete, goroutine is moved back to local run queue
