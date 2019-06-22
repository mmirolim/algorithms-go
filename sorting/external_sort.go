package sorting

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
)

func printMemUsage(header string) {
	var KB uint64 = 1024
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(header)
	fmt.Printf("Mem Alloc %+vKiB\n", m.Alloc/KB)           // output for debug
	fmt.Printf("Mem Sys %+vKiB\n", m.Sys/KB)               // output for debug
	fmt.Printf("Mem HeapAlloc %+vKiB\n", m.HeapAlloc/KB)   // output for debug
	fmt.Printf("Mem TotalAlloc %+vKiB\n", m.TotalAlloc/KB) // output for debug
	fmt.Printf("NumGC %+v\n", m.NumGC)                     // output for debug
	fmt.Printf("GCCPUFraction %+v\n", m.GCCPUFraction)     // output for debug
}

// GenerateFileWithRandomNumbers create file with random n numbers
func GenerateFileWithRandomNumbers(n int, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		num := strconv.Itoa(rand.Intn(math.MaxInt64))
		f.WriteString(num)
		f.WriteString("\n")
	}
	f.Close()
	return nil
}

// ExternalSort sorts numbers (int64) in file where each number on separate line,
// does not load the whole file into RAM, sortes in chunks (bytes) and using temp files
// returns name of the sorted file
func ExternalSort(fname string, chunk int) (fnameSorted string, err error) {
	tempDir, err := ioutil.TempDir("", fname)
	if err != nil {
		return
	}
	defer os.RemoveAll(tempDir)

	fnameSorted = fname + "_sorted"
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()
	pathTo := func(fname string) string {
		return tempDir + "/" + fname
	}

	readChunk := func(rd *bufio.Reader, buf []int64) (int, error) {
		readNumbers := 0
		for i := 0; i < len(buf); i++ {
			data, err := rd.ReadBytes('\n')
			if err != nil {
				return readNumbers, err
			}
			readNumbers++
			buf[i], err = strconv.ParseInt(string(data[:len(data)-1]), 10, 64)
			if err != nil {
				return readNumbers, err
			}
		}
		return readNumbers, nil
	}
	sort := func(chunk []int64) {
		sort.Slice(chunk, func(i, j int) bool {
			return chunk[i] < chunk[j]
		})
	}
	writeToTempFile := func(chunkNum int, chunk []int64, howManyNumToWrite int) error {
		name := pathTo(strconv.Itoa(chunkNum))
		f, err := os.Create(name)
		if err != nil {
			return err
		}
		for i := 0; i < howManyNumToWrite; i++ {
			f.WriteString(strconv.FormatInt(chunk[i], 10))
			f.WriteString("\n")
		}
		return f.Close()
	}
	chunkBuf := make([]int64, chunk)
	// read, sort and save chunks
	rd := bufio.NewReader(f)
	count := 0
	readNumbers := 0
	var errReadChunk error
	for errReadChunk == nil {
		readNumbers, errReadChunk = readChunk(rd, chunkBuf)
		if errReadChunk != nil {
			if errReadChunk != io.EOF {
				return
			} else if readNumbers == 0 {
				// EOF but no data to write
				break
			}
			// write remaining data on EOF
		}

		sort(chunkBuf)
		count++
		err = writeToTempFile(count, chunkBuf, readNumbers)
		if err != nil {
			return
		}
	}
	mergeFiles := func(suffix string, f1, f2 os.FileInfo) error {
		fd1, e1 := os.Open(pathTo(f1.Name()))
		if e1 != nil {
			return e1
		}
		defer fd1.Close()
		fd2, e2 := os.Open(pathTo(f2.Name()))
		if e2 != nil {
			return e2
		}
		defer fd2.Close()
		mfname := pathTo("merge_iter_" + suffix)
		f, err := os.Create(mfname)
		if err != nil {
			return err
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		var readers [2]*bufio.Reader
		var nums [2]int64
		var bufs [2][]byte
		for i, r := range []io.Reader{fd1, fd2} {
			readers[i] = bufio.NewReader(r)
			bufs[i], e1 = readers[i].ReadBytes('\n')
			if e1 != nil {
				return e1
			}
			nums[i], e1 = strconv.ParseInt(string(bufs[i][:len(bufs[i])-1]), 10, 64)
			if e1 != nil {
				return e1
			}
		}

		which := 0
		for {
			if nums[0] <= nums[1] {
				w.Write(bufs[0])
				which = 0
			} else {
				w.Write(bufs[1])
				which = 1
			}
			bufs[which], e1 = readers[which].ReadBytes('\n')
			if e1 != nil {
				if e1 == io.EOF {
					break
				}
				return e1
			}
			nums[which], e1 = strconv.ParseInt(string(bufs[which][:len(bufs[which])-1]), 10, 64)
			if e1 != nil {
				return e1
			}
		}
		// write what left in buf
		// which one returned error
		if which == 1 {
			w.Write(bufs[0])
			which = 0
		} else {
			w.Write(bufs[1])
			which = 1
		}
		// append remaining file
		for {
			bufs[which], e1 = readers[which].ReadBytes('\n')
			if e1 != nil {
				if e1 == io.EOF {
					break
				}
				return e1
			}
			w.Write(bufs[which])
		}

		w.Flush()
		return nil
	}
	// merge in order chunks
	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		return
	}
	count = 0
	for len(files) > 1 {
		// if there is odd numbers of files, merge on next iteration
		for i := 0; i < len(files)-1; i += 2 {
			count++
			err = mergeFiles(strconv.Itoa(count), files[i], files[i+1])
			if err != nil {
				return
			}
			err = os.Remove(pathTo(files[i].Name()))
			err = os.Remove(pathTo(files[i+1].Name()))
		}
		files, err = ioutil.ReadDir(tempDir)
		if err != nil {
			return
		}
	}
	w, err := os.Create(fnameSorted)
	if err != nil {
		return
	}
	defer w.Close()
	r, err := os.Open(pathTo(files[0].Name()))
	if err != nil {
		return
	}
	defer r.Close()
	io.Copy(w, r)
	return
}
