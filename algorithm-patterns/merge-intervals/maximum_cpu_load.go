package merge_intervals

import "sort"

type Job struct {
	start int
	end   int
	load  int
}

type Jobs []Job

func (jo *Jobs) Len() int {
	return len(*jo)
}

func (jo *Jobs) Less(i, j int) bool {
	return (*jo)[i].end < (*jo)[j].end
}

func (jo *Jobs) Swap(i, j int) {
	(*jo)[i], (*jo)[j] = (*jo)[j], (*jo)[i]
}

type JobsMinHeap []Job

func (m *JobsMinHeap) Len() int {
	return len(*m)
}

func (m *JobsMinHeap) Less(i, j int) bool {
	return (*m)[i].end < (*m)[j].end
}

func (m *JobsMinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *JobsMinHeap) Push(x Job) {
	*m = append(*m, x)
}

func (m *JobsMinHeap) Pop() Job {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *JobsMinHeap) Peek() Job {
	return (*m)[0]
}

func MaximumCPULoad(jobs Jobs) int {
	sort.Sort(&jobs)

	maxCPULoad, currentCPULoad := 0, 0
	minHeap := JobsMinHeap{}
	for _, job := range jobs {
		for minHeap.Len() > 0 && job.start >= minHeap.Peek().end {
			currentCPULoad -= minHeap.Peek().load
			minHeap.Pop()
		}

		minHeap.Push(job)
		currentCPULoad += job.load

		maxCPULoad = max(maxCPULoad, currentCPULoad)
	}

	return maxCPULoad
}
