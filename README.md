# Golang combinations

[![GoDoc](https://godoc.org/github.com/mxschmitt/golang-combinations?status.svg)](https://godoc.org/github.com/mxschmitt/golang-combinations)
[![Build Status](https://travis-ci.com/mxschmitt/golang-combinations.svg?branch=master)](https://travis-ci.com/mxschmitt/golang-combinations)
[![Coverage Status](https://coveralls.io/repos/github/mxschmitt/golang-combinations/badge.svg?branch=master)](https://coveralls.io/github/mxschmitt/golang-combinations?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mxschmitt/golang-combinations)](https://goreportcard.com/report/github.com/mxschmitt/golang-combinations)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This package provides a method to generate all ordered combinations out of a given string array.
This essentially creates the powerset of the given array except that the empty set is disregarded.

## Examples

Take a look at the [godoc](https://godoc.org/github.com/mxschmitt/golang-combinations/#pkg-examples) for examples.

In general when you have e.g. `[]string{"A", "B", "C"}` you will get:

```json
[
    ["A"],
    ["B"],
    ["A", "B"],
    ["C"],
    ["A", "C"],
    ["B", "C"],
    ["A", "B", "C"]
]
```

## Background

The algorithm iterates over each number from `1` to `2^length(input)`, separating it by binary components and utilizes the true/false interpretation of binary 1's and 0's to extract all unique ordered combinations of the input slice.

E.g. a binary number `0011` means selecting the first and second index from the slice and ignoring the third and fourth. For input `{"A", "B", "C", "D"}` this signifies the combination `{"A", "B"}`.

For input slice `{"A", "B", "C", "D"}` there are `2^4 - 1 = 15` binary combinations, so mapping each bit position to a slice index and selecting the entry for binary `1` and discarding for binary `0` gives the full subset as:

```txt
1	=	0001	=>	---A	=>	{"A"}
2	=	0010	=>	--B-	=>	{"B"}
3	=	0011	=>	--BA	=>	{"A", "B"}
4	=	0100	=>	-C--	=>	{"C"}
5	=	0101	=>	-C-A	=>	{"A", "C"}
6	=	0110	=>	-CB-	=>	{"B", "C"}
7	=	0111	=>	-CBA	=>	{"A", "B", "C"}
8	=	1000	=>	D---	=>	{"D"}
9	=	1001	=>	D--A	=>	{"A", "D"}
10	=	1010	=>	D-B-	=>	{"B", "D"}
11	=	1011	=>	D-BA	=>	{"A", "B", "D"}
12	=	1100	=>	DC--	=>	{"C", "D"}
13	=	1101	=>	DC-A	=>	{"A", "C", "D"}
14	=	1110	=>	DCB-	=>	{"B", "C", "D"}
15	=	1111	=>	DCBA	=>	{"A", "B", "C", "D"}
```
