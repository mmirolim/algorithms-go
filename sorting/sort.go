package sorting

import (
	"math"
	"math/rand"
	"runtime"
	"sort"
	"sync"

	tree "github.com/mmirolim/algos/trees"
)

func merge(in1, in2 []int) []int {
	out := make([]int, len(in1)+len(in2))
	i, j, k := 0, 0, 0
	for ; j < len(in1) && k < len(in2); i++ {
		if in1[j] < in2[k] {
			out[i] = in1[j]
			j++
		} else {
			out[i] = in2[k]
			k++
		}
	}
	for ; j < len(in1); j++ {
		out[i] = in1[j]
		i++
	}
	for ; k < len(in2); k++ {
		out[i] = in2[k]
		i++
	}

	return out
}

func MergeSort(in []int) []int {
	if len(in) == 1 {
		return in
	}
	return merge(
		MergeSort(in[:len(in)/2]),
		MergeSort(in[len(in)/2:]))
}

func qspartition(in *[]int) int {
	sl := *in
	p := rand.Intn(len(sl))
	last := len(sl) - 1
	sl[last], sl[p] = sl[p], sl[last]
	count := 0
	for i := 0; i < len(sl); i++ {
		if sl[i] < sl[last] {
			sl[i], sl[count] = sl[count], sl[i]
			count++
		}
	}

	sl[last], sl[count] = sl[count], sl[last]
	// return border index
	return count
}

func QuickSort(in *[]int) {
	if len(*in) > 1 {
		p := qspartition(in)
		p1 := (*in)[:p]
		QuickSort(&p1)
		p2 := (*in)[p+1:]
		QuickSort(&p2)
	}
}

func QuickSortConcurrent(in []int) []int {
	var wg sync.WaitGroup
	numcpu := runtime.NumCPU()
	batch := len(in) / numcpu
	if len(in)%numcpu != 0 {
		batch++
	}
	out := make([][]int, numcpu, numcpu+1)
	for i := 0; i < numcpu; i++ {
		if i*batch > len(in)-1 {
			// skip redundant batches
			continue
		}
		wg.Add(1)
		go func(i int) {
			end := (i + 1) * batch
			if end > len(in) {
				end = len(in)
			}
			s := in[i*batch : end]
			QuickSort(&s)
			out[i] = s
			wg.Done()
		}(i)
	}
	wg.Wait()
	// merge
	for len(out) > 1 {
		o := merge(out[0], out[1])
		out = append(out, o)
		out = out[2:]
	}
	return out[0]
}

func QuickSortConcurrentHeapMerge(in []int) []int {
	var wg sync.WaitGroup
	numcpu := runtime.NumCPU()
	batch := len(in) / numcpu
	if len(in)%numcpu != 0 {
		batch++
	}
	out := make([][]int, numcpu, numcpu+1)
	count := 0
	for i := 0; i < numcpu; i++ {
		if i*batch > len(in)-1 {
			// skip redundant batches
			continue
		}
		wg.Add(1)
		count++
		go func(i int) {
			end := (i + 1) * batch
			if end > len(in) {
				end = len(in)
			}
			s := in[i*batch : end]
			QuickSort(&s)
			out[i] = s
			wg.Done()
		}(i)
	}
	wg.Wait()
	// merge
	res := make([]int, len(in))
	heap, err := tree.NewMinHeap(numcpu)
	if err != nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		err = heap.Insert(out[i][0], out[i])
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < len(res); i++ {
		v, data, err := heap.ExtractTop()
		if err != nil {
			panic(err)
		}
		res[i] = v
		arr := data.([]int)
		if len(arr) > 1 {
			arr = arr[1:]
			err = heap.Insert(arr[0], arr)
			if err != nil {
				panic(err)
			}
		}
	}
	return res
}

type Point struct {
	id   int
	x, y float64
}

func FindMaxWallIntersectAimPoint(points []Point, q Point) Point {
	var aim Point
	radialAngle := func(q, p Point) float64 {
		return math.Atan((p.y - q.y) / (p.x - q.x))
	}
	makeSegments := func(ps []Point) [][2]Point {
		out := make([][2]Point, 0, len(ps))
		for i := 0; i < len(ps); i++ {
			out = append(out, [2]Point{ps[i], ps[(i+1)%len(ps)]})
		}
		return out
	}
	segments := makeSegments(points)
	type taggedPoint struct {
		point Point
		start bool
		phi   float64
	}
	tagPointFromSegment := func(s [2]Point) (taggedPoint, taggedPoint) {
		var t1, t2 taggedPoint
		t1.phi = radialAngle(q, s[0])
		t1.point = s[0]
		t2.phi = radialAngle(q, s[1])
		t2.point = s[1]
		// counterclockwise orientation,
		// start from smaller to bigger angle in segment
		if t1.phi < t2.phi {
			t1.start = true
		} else {
			t2.start = true
		}
		return t1, t2
	}
	tagPoints := func(ss [][2]Point) []taggedPoint {
		out := make([]taggedPoint, 2*len(ss))
		for i := 0; i < len(ss); i++ {
			tp1, tp2 := tagPointFromSegment(ss[i])
			out[2*i], out[2*i+1] = tp1, tp2
		}
		return out
	}
	taggedPoints := tagPoints(segments)
	sortTaggedPoints := func(tps []taggedPoint) []taggedPoint {
		sort.Slice(tps, func(i, j int) bool {
			if tps[i].phi < tps[j].phi {
				return true
			} else if math.Abs(tps[i].phi-tps[j].phi) < 0.00001 && tps[i].start {
				// if equal place first start point
				return true
			}
			return false
		})
		return tps
	}
	sortedTaggedPoints := sortTaggedPoints(taggedPoints)
	count, lastMax := 0, 0
	for i := range sortedTaggedPoints {
		if sortedTaggedPoints[i].start {
			count++
			if count > lastMax {
				lastMax = count
				aim = sortedTaggedPoints[i].point
			}
		} else {
			count--
		}
	}
	return aim
}
