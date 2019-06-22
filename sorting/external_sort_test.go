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
)

func TestGenerateFileWithRandomNumbers(t *testing.T) {
	fname := "1e3"
	var lines int = 1e3
	err := GenerateFileWithRandomNumbers(lines, fname)
	if err != nil {
		t.Errorf("unexpected error %+v", err)
		t.FailNow()
	}

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
	fname := "1e5"
	var lines int = 1e5
	err := GenerateFileWithRandomNumbers(lines, fname)
	if err != nil {
		t.Errorf("unexpected error %+v", err)
		t.FailNow()
	}
	fnameSorted, err := ExternalSort(fname, 1e4)
	if err != nil {
		t.Errorf("unexpected error %+v", err)
		t.FailNow()
	}
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
			t.Errorf("unexpected value %+v, expected int, err %v", string(data), err)
			t.FailNow()
		}
		if prev <= n {
			prev = n
		} else {
			t.Errorf("expected sorted sequence got prev %v, next %v", prev, n)
			t.FailNow()
		}
	}
	if count != lines {
		t.Errorf("expected lines number in %v, equal to %v, got %v", fnameSorted, lines, count)
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
