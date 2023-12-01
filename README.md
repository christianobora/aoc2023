# 🎄 Advent of Code 2023 🎄

My solutions for Advent of Code 2023

## Setup

1. Install Go from [here](https://golang.org/doc/install).
2. Run the following script to install dependencies.
    ```shell
    go mod tidy
    ```

## Running

Run the following command to test a given day's problem.

```shell
go test ./puzzles/day01
```


Run the following to run the given puzzle input for the day.

```shell
go run ./puzzles/day01/main.go
```

## Template

To create a template of files needed for a new day's puzzle, run the following command.

```shell
./scripts/create-day 01
```

This will create a new folder named `day01` pre-created with files for the main code, test code, and input files.