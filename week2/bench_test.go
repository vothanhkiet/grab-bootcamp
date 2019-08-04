package main

import "testing"

func BenchmarkConcurrent(b *testing.B) {
	b.ReportAllocs()
	files := getFiles("/Users/kiettv/workplaces/personal/labs/k8s") // start up worker pool
	for n := 0; n < b.N; n++ {
		doInBackground(files)
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	b.ReportAllocs()
	files := getFiles("/Users/kiettv/workplaces/personal/labs/k8s")
	for n := 0; n < b.N; n++ {
		for index, file := range files {
			job := WordCount{File: file}
			job.DoWork(index)
		}
	}
}
