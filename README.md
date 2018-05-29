# Golang combinations

[![GoDoc](https://godoc.org/github.com/mxschmitt/golang-combinations?status.svg)](https://godoc.org/github.com/mxschmitt/golang-combinations)
[![Build Status](https://travis-ci.com/mxschmitt/golang-combinations.svg?branch=master)](https://travis-ci.com/mxschmitt/golang-combinations)
[![Coverage Status](https://coveralls.io/repos/github/mxschmitt/golang-combinations/badge.svg?branch=master)](https://coveralls.io/github/mxschmitt/golang-combinations?branch=master)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

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