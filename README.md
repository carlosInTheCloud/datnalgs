# datnalgs - Data Structures and Algorithms
The datnalgs (data structures and algorithms) repo's intend is for a developer to complete the Go functions on each one of the sections, and then test the results against the provided tests. Use this repo as a way to learn and/or practice common data structures and algorithms.
## Project Sstructure
### data/ Folder
This folder contains preset data that would be referenced by some tests. 
### Section folders
Section folders are those that contain the given exercises, as well as the tests and answers to the tests. e.g. bit/, search/, sort/.  
In each folder you will find the following files.
- reset.sh script. This script will delete a {section}.go file and recreate it with the function declarations to the exercises.
- {section}_solved.txt. The answer to the exercises. There is more than one way to peal a banana, so the answers represent the author's preferred approach.
- {section}t.txt. Template used by the reset.sh script to recreate the {section}.go file.
- {section}_test.go. Tests for each one of the functions in the section. 
Note that the {section}.go file is not included in the repo by design.
Folder structure example:
```
|
|- bit
    |- reset.sh
    |- bit_solved.txt
    |- bitt.txt
    |- bit_test.go
```
## Setup:
- After cloning the repo, execute the reset.sh file in the section you're working on. i.g. bit/reset.sh.
``` 
$ ./reset.sh
```
If the file doesn't have permissions to execute, then run:
```
$ chmod +x ./reset.sh
```
You should see a {section}.go file created. 
- Complete the body for the functions in the {section}.go file. Once you're ready to test, go to the command prompt and execute the tests.
```
$ go test
```
or for verbose outputs: 
```
$ go test -v
```
Alternatively, if you'd like for the testing to fail the moment one of the test fails, run:
```
$ go test --failfast
```
Refer to the go test package https://pkg.go.dev/cmd/go/internal/test for other test parameters. 

