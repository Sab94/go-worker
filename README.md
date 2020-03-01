# go-worker

go-worker is a work queue to run tasks in go routines

# Install

`go get -u github.com/Sab94/go-worker`

# Usage

Once the package is imported under the name `gw`, an instance of Manager can be created like so:

```
  manager := gw.NewManager(2)
```

It can then be used to dispatch works in a buffered workChannle

```
  work := Work{
    Name: "Example Work"
  }

  workChannle := manager.NewBufferedManager(5)
  workChannle <- work
```
or `GoWork` can be called to dispatch works

```
  work := Work{
    Name: "Example Work"
  }
  gw.GoWork(work)
```

# Examples

Check out the example

# Project52

It is one of my [project 52](https://github.com/Sab94/project52).

Note: This project is inspired by [https://github.com/aloknerurkar/task-runner](https://github.com/aloknerurkar/task-runner)
