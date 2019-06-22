package sorting

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestGenerateFileWithRandomNumbers(t *testing.T) {
	start := time.Now()
	fname := "1e4"
	var lines int = 1e4
	err := GenerateFileWithRandomNumbers(lines, fname)
	if err != nil {
		t.Errorf("unexpected error %+v", err)
		t.FailNow()
	}
	fmt.Printf("Time spent on GenerateFileWithRandomNumbers %+v\n", time.Since(start)) // output for debug

	f, err := os.Open(fname)
	if err != nil {
		t.Errorf("unexpected error %+v", err)
		t.FailNow()
	}
	buf := bufio.NewReader(f)
	count := 0
	var data []byte
	for {
		data, err = buf.ReadBytes('\n')
		if err != nil {
			break
		}
		_, err = strconv.Atoi(string(data[:len(data)-1]))
		if err != nil {
			t.Errorf("unexpected value %+v, expected int, err %v", string(data), err)
			t.FailNow()
		}
		count++
	}
	if count != lines {
		t.Errorf("expected lines number in %v, equal to %v, got %v", fname, lines, count)
	}
}

func TestExternalSort(t *testing.T) {
	data := []struct {
		fname    string
		lines    int
		chunkSze int
	}{
		{"1e4", 1e4, 3001},
		{"1e4", 1e4, 777},
		{"1e4", 1e4, 1e3},
		{"1e5", 1e5, 1e4},
	}
	for i, d := range data {
		fname := d.fname
		lines := d.lines
		start := time.Now()
		err := GenerateFileWithRandomNumbers(lines, fname)
		if err != nil {
			t.Errorf("case [%v] unexpected error %+v", i, err)
			t.FailNow()
		}
		fmt.Printf("Time spent on GenerateRandomNumbers %+v\n", time.Since(start)) // output for debug
		start = time.Now()
		fnameSorted, err := ExternalSort(fname, d.chunkSze)
		if err != nil {
			t.Errorf("case [%v] unexpected error %+v", i, err)
			t.FailNow()
		}
		fmt.Printf("Time spent on ExternalSort %+v\n", time.Since(start)) // output for debug

		f, err := os.Open(fnameSorted)
		buf := bufio.NewReader(f)
		prev := -math.MaxInt64
		count := 0
		for {
			data, err := buf.ReadBytes('\n')
			if err != nil {
				break
			}
			count++
			n, err := strconv.Atoi(string(data[:len(data)-1]))
			if err != nil {
				t.Errorf("case [%v] unexpected value %+v, expected int, err %v", i, string(data), err)
				t.FailNow()
			}
			if prev <= n {
				prev = n
			} else {
				t.Errorf("case [%v] expected sorted sequence got prev %v, next %v", i, prev, n)
				t.FailNow()
			}
		}
		if count != lines {
			t.Errorf("case [%v] expected lines number in %v, equal to %v, got %v, chunk size %v", i, fnameSorted, lines, count, d.chunkSze)
		}
	}
}

func testPrintMemUsage(t *testing.T) {
	fname := "10e3"
	printMemUsage("start of test")
	GenerateFileWithRandomNumbers(10e3, fname)
	printMemUsage("read file to memory")
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	printMemUsage("no reads of file")
	sl, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file size %+v KiB\n", len(sl)/1024) // output for debug
	printMemUsage("before f.Close()")
	f.Close()
	printMemUsage("after f.Close()")
	runtime.GC()
	printMemUsage("after runtime.GC")
	f, err = os.Open(fname)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	var lastBytes []byte
	for err == nil {
		lastBytes, err = buf.ReadBytes('\n')
		if err != nil {
			fmt.Printf("err %v, %T\n", err, err) // output for debug

		}
	}
	fmt.Printf("Last Err %+v\n", err)                  // output for debug
	fmt.Printf("last string %+v\n", string(lastBytes)) // output for debug
	printMemUsage("after read file with readBytes")
}
