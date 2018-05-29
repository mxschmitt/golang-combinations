# Golang combinations

[![Build Status](https://travis-ci.com/mxschmitt/golang-combinations.svg?token=zmdRXuHyrWEFbXwsbBEN&branch=master)](https://travis-ci.com/mxschmitt/golang-combinations)

This package provides a method to generate all combinations out of a given string array.

## Examples

For examples check out the godoc: [here](https://godoc.org/github.com/mxschmitt/golang-combinations/#pkg-examples)

In general when you have e.g. `[]string{"A", "B", "C"}` you will get:

```json
[
    [
        "A"
    ],
    [
        "B"
    ],
    [
        "A",
        "B"
    ],
    [
        "C"
    ],
    [
        "A",
        "C"
    ],
    [
        "B",
        "C"
    ],
    [
        "A",
        "B",
        "C"
    ]
]
```