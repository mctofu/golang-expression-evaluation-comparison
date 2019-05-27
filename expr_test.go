package main

import (
	"testing"

	"github.com/antonmedv/expr"
)

func Benchmark_expr(b *testing.B) {
	params := createParams()

	program, err := expr.Compile(example, expr.Env(params))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
