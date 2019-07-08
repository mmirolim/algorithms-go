package hk

import (
	"math"
	"sort"
)

func moveOriginPoint(x, y, xo, yo float64) (float64, float64) {
	return x - xo, y - yo
}

func polarAngle(x, y float64) float64 {
	return math.Atan2(float64(y), float64(x))
}

// https://www.hackerrank.com/challenges/queens-attack-2/problem
func QueensAttack2(size int, qpos [2]int, obs [][2]int) int {
	// convert to polar coordinates, move origin point to queen position
	rmax, rmin := float64(size-qpos[0])+0.5, float64(0-qpos[0])+0.5
	cmax, cmin := float64(size-qpos[1])+0.5, float64(0-qpos[1])+0.5
	phiph := polarAngle(1, 0)
	phipd := polarAngle(1, 1)
	phipv := polarAngle(0, 1)
	phinxd := polarAngle(-1, 1)
	phinh := polarAngle(-1, 0)
	phind := polarAngle(-1, -1)
	phinv := polarAngle(0, -1)
	phinyd := polarAngle(1, -1)
	qAttackAngles := map[float64][]float64{
		phiph: []float64{cmax}, phipd: []float64{min(cmax, rmax)},
		phipv: []float64{rmax}, phinxd: []float64{min(abs(cmin), rmax)},
		phinh: []float64{abs(cmin)}, phind: []float64{min(abs(cmin), abs(rmin))},
		phinv: []float64{abs(rmin)}, phinyd: []float64{min(cmax, abs(rmin))},
	}
	for _, v := range obs {
		x, y := moveOriginPoint(float64(v[1])-0.5, float64(v[0])-0.5, float64(qpos[1])-0.5, float64(qpos[0])-0.5)
		phi := polarAngle(x, y)
		// filter out which does not cross attack angles
		if list, ok := qAttackAngles[phi]; ok {
			// compute free distance till obstacle
			list = append(list, max(abs(x), abs(y))-1)
			qAttackAngles[phi] = list
		}
	}
	// sort obs by free distance
	for k, list := range qAttackAngles {
		sort.Float64s(list)
		qAttackAngles[k] = list
	}

	pos := 0
	for _, list := range qAttackAngles {
		// take closest one
		pos += int(list[0])
	}

	if pos < 0 {
		return 0
	}

	return pos
}
func abs(x float64) float64 {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(n1, n2 float64) float64 {
	if n1 > n2 {
		return n2
	}
	return n1
}

func max(n1, n2 float64) float64 {
	if n1 > n2 {
		return n1
	}
	return n2
}
