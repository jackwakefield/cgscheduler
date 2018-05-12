# cgscheduler

A concurrent scheduler for tasks with dependencies.

## Installation

```bash
go get -u github.com/jackwakefield/cgscheduler
```

or 

```bash
# go get -u github.com/golang/dep/cmd/dep
# dep init
dep ensure -add github.com/jackwakefield/cgscheduler
```

## Usage

[Documentation](http://godoc.org/github.com/jackwakefield/cgscheduler)

## Example

[example/main.go](https://github.com/jackwakefield/cgscheduler/blob/master/example/main.go).

## FAQ

### How do I limit the number of concurrent tasks executed at once?

This can be achieved with `ConcurrentTasks` when creating the scheduler.

```go
// limit the maximum number of tasks executed at once to 2
scheduler := cgscheduler.New(cgscheduler.ConcurrentTasks(2))
```

This is set to `runtime.NumCPU()` by default.