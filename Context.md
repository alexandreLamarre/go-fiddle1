### Context Package

- We need a way to propagate request-scoped data down the call-graph.

- We need a way to propagate cancellation signal down the call-graph.

Context package servers two primary purpose:

- Provides API's for cancelling branches of call-graph
- Provides a data-bag for transporting request-scoped data through the call graph

TODO()
```
func fun() {
    ctx := context.TODO()
}
```


- TODO() returns an Empty Context
- TODOs intended purpose is to server as a placeholder

Context package can be used to send,
- Request-scoped values
- Cancellation signals

Across API boundaries to all goroutines involved in handling a request

context.Background() -> returns an empty context, it is the root of any Context tree

#### Context pacakge ... cancellation

- Context is immutable
- Context package provides function to add new behaviour
- To add concellation behaviour we have a function like:
  - context.WithCancel()
  - context.WithTimeout()
  - context.WithDeadline()

- The derived context is passed to child goroutines to facilitate their cancellation

#### WithCancel()

```
ctx, cancel := context.withCancel(context.Background())
defer cancel()
```

- WithCancel returns a copy of parent with a new Done channel
- cancel() can be used to close context's done channel.
- Closing the done channel indicates to an operation to abandon its work and return : Note! can cause memory leaks if you cancel before workload is done
- Cancelling the context release the resources associated with it

- `cancel()` does not wait for the work to stop
- `cancel()` may be called by multiple gouroutines simulataneously. 

- After the first call, subsequent calls to a cancel() will do nothing.

Example:

Workflow (Parent goroutine)
```
ctx, cancel := context.WithCancel(context.Background())
ch := generator(ctx)
...
if n == 5{
    cancel()
}
```

Workflow(Child goroutine)
```
for{
    select{
    case <-ctx.Done():
        return ctx.Err()
    
    case dst <- n:
        n ++
    }
}
```

#### WithDeadline()

```
deadline := time.Now().Add(5*time.Millisecond)
ctx, cancel := context.WithoutDeadline(context.Background(), deadline)
defer cancel()
```

- WithDeadline() takes parent context and clock time as input
- WithDeadline returns a new Context that closes its done channel when the machine's clock advances past the given deadline

#### WithTimeout()

```
duration := 5 * time.Millisecond
ctx, cancel := context.WithTimeout(context.Background(), duration)
defer cancel()
```

- WithTimeout() takes a parent context and time duration as input
- WithTimeout() returns a new Context that closes its done channel after the given timeout duration
- WithTimeout() is useful for setting a deadline on the requests to backend servers
- WithTimeout() is actually a wrapper over WithDeadline()
- WithTimeout() timer countdown begins from the moment the context is created
- WithDeadline() sets the explicit time when timer will expire

#### Context Package as Data bag

- Context Package can be used to transport request scoped data down the call graph
- context.WithValue() provides a way to associate request-scoped values with a Context


Example:

Parent goroutine
```
type userIDType string
ctx := context.WithValue(context.Background(), userIDType("userIDKey", "jane"))
```

Child goroutine
```
userid := ctx.Value(userIDType("userIDKey)).(userIDType)
```

#### Context (Go idioms)

Incoming requests to a server should create a Context

- Create context early in processing task or request
- Create a top level context
```
func main(){
    ctx := context.Background()
}
```
- http.Request value already contains a Context
```
func hadleFunc(w http.ResponseWriter, req *http.Request) {
    ctx, cancel := context.WithCancel(req.Context())
}
```

Outgoing calls to servers should acceot a Context
- Higher level calls need to tell lower level calls how long they are willing to wait
```
//Create a context with a timeout of 100 milliseconds
ctx, cancel := context.WithTimeout(req.Context(), 100*time.Millisecond)
defer cancel()

// Bind the new context into the request
req = req.WithContext(ctx)

// Do will handle the context level timeout
resp, err := http.DefaultClient.Do(req)
```

- http.DefaultClient.Do() method to respect cancellation signal on timer expiry and return with an error message

##### Always pass a context to function performing I/O

- Any function that is performing I/O should accept a Context value as its first parameter and respect any timeout
or deadline configured by the caller
- Any API's that take a Context, the idiom is to have the first parameter accept the Context value
- Any change to a Context value creates a new Context value that is then propagated forward
- When a context is cancelled, all Contexts derived from it are also cancelled

**Use TODO context if you are unsure about which Context to use**

- If a function is not responsible for creating top level context
- We need a temporary top-level Context until we figured out where the actual Context will come from

#### Use context values for request scoped data
- Do not use the Context value to pass data into a function which becomes essential for its successful execution.
- A function should be able to execute its logic with an empty Context value

##### Summary
- Incoming request to a server should create a Context
- Outgoing calls to servers should accept a Context
- Any function that is performing I/O should accept a Context value.
- Any change to a Context vakue creates a new Context value that is then propagated forward
- If a parent Context is cancelled, all children derived from it are also cancelled