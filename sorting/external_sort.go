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
	// tempDir, err := ioutil.TempDir("", fname)
	// if err != nil {
	// 	return
	// }
	tempDir := os.TempDir() + "/" + fname
	os.RemoveAll(tempDir)
	err = os.Mkdir(tempDir, 0777)
	if err != nil {
		return
	}
	fmt.Printf("TempDir %+v\n", tempDir) // output for debug
	fnameSorted = fname + "_sorted"
	f, err := os.Open(fname)
	if err != nil {
		return
	}
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
		fmt.Printf("writeToTempFile name %+v\n", name) // output for debug
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
	printMemUsage("Mem before read loop")
	readNumbers := 0
	for {
		readNumbers, err = readChunk(rd, chunkBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		sort(chunkBuf)
		count++
		err = writeToTempFile(count, chunkBuf, readNumbers)
		printMemUsage("Mem after read loop iteration")
		if err != nil {
			return
		}
	}
	mergeFiles := func(suffix string, f1, f2 os.FileInfo) error {
		fd1, e1 := os.Open(pathTo(f1.Name()))
		if e1 != nil {
			return e1
		}
		fd2, e2 := os.Open(pathTo(f2.Name()))
		if e2 != nil {
			return e2
		}
		mfname := pathTo("merge_iter_" + suffix)
		fmt.Printf("Merged file name %+v\n", mfname) // output for debug
		f, err := os.Create(mfname)
		if err != nil {
			return err
		}
		w := bufio.NewWriter(f)
		rd1 := bufio.NewReader(fd1)
		rd2 := bufio.NewReader(fd2)
		var n1, n2 int64
		var d1, d2 []byte
		n1Written := false
		// read both n1, n2 first
		d1, e1 = rd1.ReadBytes('\n')
		if e1 != nil {
			return e1
		}
		n1, e1 = strconv.ParseInt(string(d1[:len(d1)-1]), 10, 64)
		if e1 != nil {
			return e1
		}
		d2, e2 = rd2.ReadBytes('\n')
		if e2 != nil {
			return e2
		}
		n2, e2 = strconv.ParseInt(string(d2[:len(d2)-1]), 10, 64)
		if e2 != nil {
			return e2
		}
		for {
			if n1 <= n2 {
				n1Written = true
				w.WriteString(strconv.FormatInt(n1, 10))
			} else {
				n1Written = false
				w.WriteString(strconv.FormatInt(n2, 10))
			}
			w.WriteByte('\n')

			// read from where it was written
			if n1Written {
				d1, e1 = rd1.ReadBytes('\n')
				if e1 != nil {
					if e1 == io.EOF {
						break
					}
					return e1
				}
				n1, e1 = strconv.ParseInt(string(d1[:len(d1)-1]), 10, 64)
				if e1 != nil {
					return e1
				}
			} else {
				d2, e2 = rd2.ReadBytes('\n')
				if e2 != nil {
					if e2 == io.EOF {
						break
					}
					return e2
				}
				n2, e2 = strconv.ParseInt(string(d2[:len(d2)-1]), 10, 64)
				if e2 != nil {
					return e2
				}
			}
		}
		w.Flush()
		// write rest of f1 or f2 to merged file
		if e1 == nil {
			// rd2 EOF, we have unprocessed d1 bytes stored
			w.Write(d1)
			for {
				d1, e1 = rd1.ReadBytes('\n')
				if e1 != nil {
					if e1 == io.EOF {
						break
					}
					return e1
				}
				w.Write(d1)
			}
		}
		if e2 == nil {
			// we didn't write d1, d2 yet after break first loop
			w.Write(d1)
			w.Write(d2)
			for {
				d2, e2 = rd2.ReadBytes('\n')
				if e2 != nil {
					if e2 == io.EOF {
						break
					}
					return e2
				}
				w.Write(d2)
			}
		}
		w.Flush()
		return f.Close()
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
		printMemUsage("Mem after mergeFiles loop iteration")
		files, err = ioutil.ReadDir(tempDir)
		if err != nil {
			return
		}
	}
	w, err := os.Create(fnameSorted)
	if err != nil {
		return
	}
	r, err := os.Open(pathTo(files[0].Name()))
	if err != nil {
		return
	}
	io.Copy(w, r)
	printMemUsage("Mem before return")

	return
}
