package encode

type PriorityQueue []Three

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	less := pq[i].Freq() < pq[j].Freq()
	if pq[i].Freq() == pq[j].Freq() {
		return pq[i].Root.Char() > pq[j].Root.Char()
	}
	return less
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(Three)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = Three{} // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
