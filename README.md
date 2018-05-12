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

## Documentation

[GoDoc](http://godoc.org/github.com/jackwakefield/cgscheduler)

## Usage

Create the scheduler with `cgscheduler.New()`

```go
scheduler := cgscheduler.New()

// or, limit the number of concurrent tasks executed at once,
// by default this is set to runtime.NumCPU()
scheduler := cgscheduler.New(cgscheduler.ConcurrentTasks(2))
```

Tasks can be created with functions matching the signature `func(ctx context.Context) error` 

```go
// add a task which outputs "World!"
taskWorld := scheduler.AddTask(func(ctx context.Context) error {
  fmt.Print("World!")
  return nil
})

// add a task which outputs "Hello"
taskHello := scheduler.AddTask(func(ctx context.Context) error {
  fmt.Print("Hello")
  return nil
})

// add a task which outputs a space (" ")
taskSeparator := scheduler.AddTask(func(ctx context.Context) error {
  fmt.Print(" ")
  return nil
})
```

Dependencies can be created between tasks with `Task.DependsOn`, which uses `Scheduler.AddDependency` internally.

```go
// execute taskHello before taskSeparator
taskSeparator.DependsOn(taskHello)

// execute taskSeparator before taskWorld
taskWorld.DependsOn(taskSeparator)
```

Run the scheduler with `Scheduler.Run`, this accepts a context as a parameter and returns an error.

The scheduler returns when all tasks are complete, or a task has returned an error.

```go
if err := scheduler.Run(context.Background()); err != nil {
	log.Fatalln(err)
}

// Outputs:
// Hello World!
```

## Example

[example/main.go](https://github.com/jackwakefield/cgscheduler/blob/master/example/main.go)

## Internals

Internally the scheduler uses a [Directed Acyclic Graph](https://github.com/jackwakefield/graff) to represent the tasks as nodes and their dependencies as edges.

When the scheduler is ran, and when the graph state has changed since the last run, the tasks are [topologically ordered](https://en.wikipedia.org/wiki/Topological_order) into levels using the [Coffman-Graham algorithm](https://en.wikipedia.org/wiki/Coffman–Graham_algorithm).