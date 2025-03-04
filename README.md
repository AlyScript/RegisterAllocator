### Usage
**To run**
```bash
go run main.go <input_file> <output_file>
```

**If you want to build the executable**
```bash
go build main.go
```

**and then run with**
```bash
./main
```

### Checking Results
The `check_result.sh` script will run `main.go` and check if `output.txt` is equal to `expected.txt`. If any files are missing then it will simply output "test failed".

**NOTE**: The program always terminates the output file with a blank line, so `expected.txt` **_MUST_** contain one otherwise there will be false negatives!