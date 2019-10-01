pact-sample
==

Generate protocol buffers
--

### `golang`
Use `protoc` to output the Golang implementation of Protocol buffers. This requires `protoc-gen-go`.

```shell
$ protoc --go_out=./pb/addressbook ./pb/addressbook.proto
```

### `javascript`
Similarly, output the Javascript implementation.

```shell
protoc --js_out=import_style=commonjs,binary:./pb/addressbook ./pb/addressbook.proto
```
