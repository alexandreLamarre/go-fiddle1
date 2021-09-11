Go provides race detector tool for finding race conditions in Go code

- Binary needs to be race enabled
- When racy behaviour is detected a waning is printed
- Race enabled binary will be 10 times slower and consume 10 times more moemory 
- Integration tests and load tests are good candidates to test with binary race enabled