package templates

// 优先队列、最小堆模板实现
type hq []int

func (hq *hq) Len() int               { return len(*hq) }
func (hq *hq) Less(i int, j int) bool { return (*hq)[i] < (*hq)[j] }
func (hq *hq) Swap(i int, j int)      { (*hq)[i], (*hq)[j] = (*hq)[j], (*hq)[i] }
func (hq *hq) Push(x any)             { (*hq) = append((*hq), x.(int)) }
func (hq *hq) Pop() any               { n := len(*hq); item := (*hq)[n-1]; (*hq) = (*hq)[:n-1]; return item }
func (hq *hq) Peek() *int {
	n := len(*hq)
	if n == 0 {
		return nil
	} else {
		return &(*hq)[n-1]
	}
}
