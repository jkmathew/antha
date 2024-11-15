package main

import (
	//"github.com/jkmathew/antha/antha/anthalib/wtype"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/jkmathew/antha/antha/AnthaStandardLibrary/Packages/Optimization"
)

func DefaultParameters() Optimization.AssemblyOptimizerParameters {
	prm := NewGAParameters()
	/*
		prm["max_iterations"] = 100
		prm["recom_p"] = 0.25
		prm["mut_p"] = 0.25
		prm["step_size"] = 3
		prm["pop_size"] = 100
	*/
	prm.Set("max_iterations", 200)
	prm.Set("recom_p", 0.5)
	prm.Set("mut_p", 0.5)
	prm.Set("step_size", 1)
	prm.Set("pop_size", 250)

	return prm
}

func DefaultConstraints() Constraints {
	return Constraints{
		MaxLen:       500,
		MinLen:       10,
		MaxSplits:    4,
		MinDistToMut: 2,
	}
}

func BasicProblem() AssemblyProblem {
	length := 1000
	mut := make(PointSet2D, 0, 3)

	mut = append(mut, Point2D{X: 100, Y: 4})
	mut = append(mut, Point2D{X: 200, Y: 3})
	mut = append(mut, Point2D{X: 600, Y: 8})

	return AssemblyProblem{Len: length, Mutations: mut}
}

// TODO --> end constraints for assembly
// 	    costs for ends ?

func main() {

	rand.Seed(time.Now().UnixNano())
	//func OptimizeAssembly(problem AssemblyProblem, constraints Constraints, parameters Optimization.AssemblyOptimizerParameters) {

	prm := DefaultParameters()
	cnstr := DefaultConstraints()
	problem := BasicProblem()

	OptimizeAssembly(problem, cnstr, prm)

}

type PointSet1D []int

func (ps1D PointSet1D) Dup() PointSet1D {
	r := make([]int, len(ps1D))

	for i := 0; i < len(ps1D); i++ {
		r[i] = ps1D[i]
	}

	return r
}

type PointSet2D []Point2D

type Point2D struct {
	X int
	Y int
}

func (ps2d PointSet2D) Less(i, j int) bool {
	return ps2d[i].X < ps2d[j].X
}

func (ps2d PointSet2D) Swap(i, j int) {
	t := ps2d[i]
	ps2d[i] = ps2d[j]
	ps2d[j] = t
}

func (ps2d PointSet2D) Len() int {
	return len(ps2d)
}

func abs(i int) int {
	if i < 0 {
		i = i * -1
	}

	return i
}

func (ps2d PointSet2D) MinDistTo(i int) int {
	m := ps2d[len(ps2d)-1].X

	for _, p := range ps2d {
		d := abs(p.X - i)

		if d < m {
			m = d
		}
	}

	return m
}

type Constraints struct {
	MaxSplits    int
	MinLen       int
	MaxLen       int
	MinDistToMut int
}

type AssemblyProblem struct {
	Mutations PointSet2D // set of mutations
	Len       int        // total sequence length
}

func OptimizeAssembly(problem AssemblyProblem, constraints Constraints, parameters Optimization.AssemblyOptimizerParameters) {
	// core of problem:
	// given N pairs of points (x_i,y_i)
	// choose up to K points such that
	// a) k_j =/= x_i for any i,j; also all k_j distinct
	// b) we minimize a cost function f which sums the products
	//    of all y_is corresponding to x_is which are between
	//    pairs of k_js
	rand.Seed(time.Now().UnixNano())

	pop := NewPop(problem, constraints, parameters)

	scores := pop.Assess()

	bestScore := scores.BestScore
	bestMember := scores.BestMember.Dup()

	for time := 1; time <= parameters.MaxIterations(); time++ {
		pop = pop.Regenerate(scores, parameters, constraints)
		scores = pop.Assess()

		if scores.BestScore < bestScore {
			bestScore = scores.BestScore
			bestMember = scores.BestMember.Dup()
		}
	}

	fmt.Println("PROB: ", problem)
	fmt.Println("BEST: ", bestScore)
	fmt.Println("MEM : ", bestMember)
}

type Population struct {
	Members     []PointSet1D
	Problem     AssemblyProblem
	FitnessTest func(f int, fs []int) bool
}

type FitnessValues struct {
	Fit        []int
	BestScore  int
	BestMember PointSet1D
}

func (p *Population) Regenerate(scores FitnessValues, prm Optimization.AssemblyOptimizerParameters, cnstr Constraints) *Population {
	newMembers := make([]PointSet1D, 0, len(p.Members))

	for i := 0; i < len(p.Members); i++ {
		// choose a member

		mem1 := p.Pick(scores, nil)

		// decide what to do
		r := rand.Float64()
		recomP, _ := prm.GetFloat("recom_p")
		if r < recomP {
			mem2 := p.Pick(scores, mem1)
			mem1 = p.Recombine(mem1, mem2, prm, cnstr)
		} else {
			for {
				mem1 = p.Mutate(mem1, prm, cnstr)
				mp, _ := prm.GetFloat("mut_p")
				f := rand.Float64()
				if f > mp {
					break
				}
			}
		}

		newMembers = append(newMembers, mem1)
	}

	ret := Population{Members: newMembers, Problem: p.Problem, FitnessTest: p.FitnessTest}

	return &ret
}

func (p *Population) Pick(fit FitnessValues, m PointSet1D) PointSet1D {
	var picked PointSet1D

	for tries := 0; tries < len(fit.Fit); tries++ {
		if picked != nil && !reflect.DeepEqual(picked, m) {
			break
		}

		for {
			i := rand.Intn(len(p.Members))

			picked = p.Members[i]

			if p.FitnessTest(fit.Fit[i], fit.Fit) {
				break
			}

		}
	}

	return picked
}

func (pop *Population) Assess() FitnessValues {

	fit := make([]int, len(pop.Members))
	best := -1
	bestAt := -1

	for i := 0; i < len(pop.Members); i++ {
		fit[i] = Cost(pop.Members[i], pop.Problem)

		if bestAt == -1 || best > fit[i] {
			best = fit[i]
			bestAt = i
		}
	}

	return FitnessValues{Fit: fit, BestScore: best, BestMember: pop.Members[bestAt].Dup()}
}

func dist(a, b int) int {
	return b - a + 1
}

func valid(m PointSet1D, p AssemblyProblem, cnstr Constraints) bool {
	if len(m) > cnstr.MaxSplits {
		return false
	}

	last := 0
	for i := 0; i < len(m); i++ {
		if m[i] < 0 || m[i] >= p.Len {
			return false
		}

		d := dist(last, m[i])

		if d < cnstr.MinLen || d > cnstr.MaxLen {
			return false
		}

		// check for distance to mutation

		dTM := p.Mutations.MinDistTo(m[i])

		if dTM < cnstr.MinDistToMut {
			return false
		}

		last = m[i]
	}

	d := dist(m[len(m)-1], p.Len)

	if d < cnstr.MinLen || d > cnstr.MaxLen {
		return false
	}

	return true
}

func (p *Population) Recombine(m1, m2 PointSet1D, prm Optimization.AssemblyOptimizerParameters, cnstr Constraints) PointSet1D {
	// we will keep retrying until it works (or we give up, which we may)
	var mem PointSet1D

	for {
		l := len(m1)

		if len(m2) > l {
			l = len(m2)
		}

		mem = make(PointSet1D, 0, l)

		for i := 0; i < l; i++ {
			if rand.Intn(100) > 49 {
				if i < len(m1) {
					mem = append(mem, m1[i])
				}
			} else {
				if i < len(m2) {
					mem = append(mem, m2[i])
				}
			}
		}

		if valid(mem, p.Problem, cnstr) {
			break
		}
	}

	// sort the member
	sort.Ints(mem)
	return mem
}

func (p *Population) Mutate(mem PointSet1D, prm Optimization.AssemblyOptimizerParameters, cnstr Constraints) PointSet1D {
	const (
		ADD = iota
		DELETE
	)

	for {
		stop := false
		// we might add, delete or move
		move := rand.Intn(3)
		switch move {
		case ADD:
			if len(mem) == cnstr.MaxSplits {
				// can't add more
				continue
			}
			ret := p.addMutation(mem, prm, cnstr)
			if ret != nil {
				mem = ret
				stop = true
			}
		case DELETE:
			if len(mem) == 0 {
				// can't delete
				continue
			}
			ret := p.delMutation(mem, prm, cnstr)

			if ret != nil {
				mem = ret
				stop = true
			}
		default: // i.e. MOVE
			if len(mem) == 0 {
				// can't move
				continue
			}
			ret := p.movMutation(mem, prm, cnstr)
			if ret != nil {
				stop = true
				mem = ret
			}
		}
		if stop {
			break
		}
	}

	return mem
}

func (p *Population) addMutation(mem PointSet1D, prm Optimization.AssemblyOptimizerParameters, cnstr Constraints) PointSet1D {
	for {
		m := mem.Dup()
		// choose a spot

		l := rand.Intn(p.Problem.Len)

		// append to m

		m = append(m, l)

		// check if it's OK

		sort.Ints(m)
		if valid(m, p.Problem, cnstr) {
			mem = m
			break
		}
	}

	// sort the member
	return mem
}

func (pop *Population) delMutation(mem PointSet1D, prm Optimization.AssemblyOptimizerParameters, cnstr Constraints) PointSet1D {
	// should be pretty sure we've tried everything
	for tries := 0; tries < len(mem)*2; tries++ {
		p := rand.Intn(len(mem))

		prev := 0

		if p > 0 {
			prev = mem[p-1]
		}

		next := pop.Problem.Len

		if p < len(mem)-1 {
			next = mem[p+1]
		}

		// omit the pth member... but
		// don't leave too big a gap
		if (prev - next + 1) < cnstr.MaxLen {
			// must already be > minlen

			ret := make(PointSet1D, len(mem)-1)
			ret = append(ret, mem[:p]...)
			ret = append(ret, mem[p+1:]...)

			if valid(ret, pop.Problem, cnstr) {
				return ret
			}
		}
	}
	return nil
}

func (pop *Population) movMutation(mem PointSet1D, prm Optimization.AssemblyOptimizerParameters, cnstr Constraints) PointSet1D {
	moved := false
	for tries := 0; tries < len(mem)*2; tries++ {
		m := mem.Dup()

		// choose a position
		p := rand.Intn(len(m))

		// move it

		stepSize, _ := prm.GetInt("step_size")
		s := rand.Intn(stepSize*2 + 1)
		s -= stepSize
		m[p] += s

		if valid(m, pop.Problem, cnstr) {
			moved = true
			mem = m
			break
		}
	}

	if !moved {
		return nil
	}

	// sort the member

	sort.Ints(mem)
	return mem
}

func scale(f int, fs []int) float64 {
	max := fs[0]
	min := fs[0]

	for _, v := range fs {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return float64(f-min) / float64(max-min)
}

func ScaledFitnessTest(f int, fs []int) bool {
	s := scale(f, fs)

	r := rand.Float64()

	if r > s || s >= 1.0 {
		return true
	}

	return false
}

func NewPop(problem AssemblyProblem, constraints Constraints, parameters Optimization.AssemblyOptimizerParameters) *Population {
	/*
		sft := func(f int, fs []int) bool {
			s := scale(f, fs)

			r := rand.Float64()

			if r > s {
				return true
			}

			return false
		}
		sft = sft
	*/
	popSize, _ := parameters.GetInt("pop_size")
	members := make([]PointSet1D, 0, popSize)

	for i := 0; i < popSize; i++ {
		members = append(members, NewMember(problem, constraints, parameters))
	}

	p := Population{Members: members, Problem: problem, FitnessTest: ScaledFitnessTest}

	return &p
}

func NewMember(problem AssemblyProblem, constraints Constraints, parameters Optimization.AssemblyOptimizerParameters) PointSet1D {

	// we just keep trying

	for {
		m := makeMember(problem, constraints)

		if m != nil {
			return m
		}
	}
}

func makeMember(problem AssemblyProblem, constraints Constraints) PointSet1D {
	ret := make(PointSet1D, 0, constraints.MaxSplits)

	// minimum number of splits is problem.Len / constraints.MaxLen (integer div)

	minSplit := problem.Len / constraints.MaxLen

	if constraints.MaxSplits < minSplit {
		panic("too long for this number of splits")
	}

	for {
		nSplit := rand.Intn(constraints.MaxSplits-minSplit) + minSplit
		// now place the n split points
		last := 0
		for i := 0; i < nSplit; i++ {
			p := rand.Intn(constraints.MaxLen-constraints.MinLen) + constraints.MinLen + last
			ret = append(ret, p)
			last = p
		}

		// the only remaining question is if the last position is invalid

		left := problem.Len - last

		// or if the mutation distance is too low

		if left < constraints.MinLen || left > constraints.MaxLen || !valid(ret, problem, constraints) {
			// start again!
			ret = make(PointSet1D, 0, constraints.MaxSplits)
			continue
		} else {
			break
		}
	}

	return ret
}

func Cost(k PointSet1D, problem AssemblyProblem) int {
	x := problem.Mutations
	last := 0

	sort.Sort(x)
	sort.Ints(k)

	kk := k.Dup()
	kk = append(kk, problem.Len)

	tot := 0
	for _, p := range kk {
		s := p - last + 1 // score for a segment
		m := 1            // mutation multiplier
		for _, p2 := range x {
			if p2.X >= last && p2.X < p {
				m *= p2.Y
			}
		}

		tot += s * m
		last = p
	}

	return tot
}
