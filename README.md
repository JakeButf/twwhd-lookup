## The Wind Waker HD Memory Address API

### CLI Usage

To generate usable YAML run the following command: 
```
go run . --version <twwhd version> --output <filename>
```

To generate addresses for aroma for example you can run: 
```
go run . --version aroma --output aroma.yaml
```

### Build

If you would rather build this as a binary:
```
go build -o twwhd-lookup
./twwhd-lookup --version <twwhd version> --output <filename>
```
