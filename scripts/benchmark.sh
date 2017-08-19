#!/usr/bin/env bash

go test -run=XXX -bench=. -benchmem > testdata/benchmark_results.txt
