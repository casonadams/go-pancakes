# go-pancakes

## Tests

### Unit

```bash
go test -v github.com/casonadams/go-pancakes/waiter
```

### Performance

```bash
go test -benchmem -run=^$ github.com/casonadams/go-pancakes/waiter -bench .
```
