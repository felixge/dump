package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkJSONUmarshal(b *testing.B) {
	data, err := os.ReadFile(filepath.Join("testdata", "example.json"))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	var keepAlive []interface{}
	for i := 0; i < b.N; i++ {
		var m interface{}
		err := json.Unmarshal(data, &m)
		if err != nil {
			b.Fatal(err)
		}
		keepAlive = append(keepAlive, m)
	}
}
