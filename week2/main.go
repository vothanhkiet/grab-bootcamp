package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WordCount WordCount
type WordCount struct {
	File string
}

// DoWork DoWork
func (i WordCount) DoWork(index int) map[string]int {
	wordCount := make(map[string]int)
	inFile, err := os.Open(i.File)
	if err != nil {
		fmt.Println(err.Error() + `: ` + i.File)
		return wordCount
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for i := range words {
			wordCount[words[i]]++
		}

	}

	return wordCount
}

// Job Job
type Job interface {
	DoWork(index int) map[string]int
}

// Worker Worker
type Worker struct {
	index int
	job   chan Job
	end   chan bool
	comm  chan map[string]int
	pool  chan chan Job
}

func (w Worker) start() {
	for {
		// Worker is available, can pick job now
		w.pool <- w.job
		select {
		// New Job is arrived
		case job := <-w.job:
			// Process it
			ret := job.DoWork(w.index)
			// Return result
			w.comm <- ret
		case <-w.end:
			return
		}
	}
}

func (w Worker) stop() {
	// fmt.Printf("worker [%d] is stopped\n", w.index)
	w.end <- true
	// Clean up resource
	close(w.job)
}

// Pool Pool
type Pool struct {
	pool    chan chan Job
	job     chan Job
	workers []Worker
	end     chan bool
	comm    chan map[string]int
	results []map[string]int
	total   int
	counter int
}

// NewWorkerPool NewWorkerPool
func NewWorkerPool(workerCount int, jobCount int) *Pool {
	p := Pool{
		pool:    make(chan chan Job, workerCount),
		job:     make(chan Job),
		comm:    make(chan map[string]int, workerCount*2),
		total:   jobCount,
		end:     make(chan bool),
		counter: 0,
	}
	for i := 0; i < workerCount; i++ {
		p.workers = append(p.workers, Worker{
			index: i + 1,
			pool:  p.pool,
			job:   make(chan Job),
			comm:  p.comm,
			end:   make(chan bool),
		})
	}
	for _, worker := range p.workers {
		go worker.start()
	}
	go p.start()
	return &p
}

func (p *Pool) dispatch(job Job) {
	p.job <- job
}

func (p *Pool) start() {
	// Loop until end
	for {
		select {
		// New job is arrived
		case job := <-p.job:
			// get free worker
			workerJobChannel := <-p.pool
			// Send job to worker
			workerJobChannel <- job
		case ret := <-p.comm:
			p.results = append(p.results, ret)
			// Received result from worker, increase counter
			p.counter++
			if p.counter == p.total {
				p.stop()
				return
			}
		}
	}
}

func (p *Pool) stop() {
	// All results have been returned, signal to stop the worker
	for _, worker := range p.workers {
		worker.stop()
	}
	// clean up resource
	close(p.job)
	close(p.pool)
	close(p.comm)
	p.end <- true
}

func getFiles(path string) []string {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// files = []string{"/Users/kiettv/workplaces/personal/workshops/grab/go.mod"}
	return files
}

func doInBackground(files []string) *Pool {
	pool := NewWorkerPool(8, len(files))
	for _, file := range files {
		pool.dispatch(WordCount{File: file})
	}

	<-pool.end
	return pool
}

func main() {
	// Prepare input for jobs
	files := getFiles("/Users/kiettv/workplaces/personal/labs/k8s")
	pool := doInBackground(files)

	finalResult := make(map[string]int)
	for _, results := range pool.results {
		for key, val := range results {
			if _, ok := finalResult[key]; ok {
				finalResult[key] = finalResult[key] + val
			} else {
				finalResult[key] = val
			}
		}
	}
	fmt.Println("Total words: ", len(finalResult))
	index := 1
	for key, val := range finalResult {
		fmt.Println(index, " - ", key, " occurrences ", val, " time(s)")
		index++
	}
}
