package five

// Chain represents a chain structure
type Chain struct {
	value int
	next  *Chain
	prev  *Chain
}

func (c Chain) isMonoChain() bool {
	return c.next == nil || c.value == (*c.next).value
}

// JosephusSurvivor get the last cyclic int in n modulo k
func JosephusSurvivor(n, k int) int {
	chain := buildChain(n)

	for !chain.isMonoChain() {
		for i := 0; i < k-1; i++ {
			chain = *chain.next
		}
		prev := chain.prev
		next := chain.next
		prev.next = next
		next.prev = prev
		chain = *chain.next
	}

	return chain.value
}

func buildChain(n int) Chain {
	chainArr := []Chain{}
	chainArr = append(chainArr, Chain{1, nil, nil})
	for i := 1; i < n; i++ {
		chainArr = append(chainArr, Chain{i + 1, nil, nil})
	}
	for i := 1; i < n; i++ {
		chainArr[i-1].next = &chainArr[i]
		chainArr[i].prev = &chainArr[i-1]
	}
	chainArr[n-1].next = &chainArr[0]
	chainArr[0].prev = &chainArr[n-1]

	return chainArr[0]
}
