To run
```bash
go run main.go <input_file> <output_file>
```

If you want to build the executable
```bash
go build main.go
```

and then run with
```bash
./main
```

The `check_result.sh` script will run `main.go` and check if `output.txt` is equal to `expected.txt`. If any files are missing then it will simply output "test failed".