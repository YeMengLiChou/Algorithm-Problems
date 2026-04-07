package main

// 通过取模可以求出所在边，进而给出方向；

type Robot struct {
	w, h, step int
}

func Constructor(width int, height int) Robot {
	return Robot{w: width, h: height, step: 0}
}

func (this *Robot) Step(num int) {
	this.step = (this.step+num-1)%((this.w+this.h-2)*2) + 1
}

func (this *Robot) GetPos() []int {
	x, y, _ := this.getState()
	return []int{x, y}
}

func (this *Robot) GetDir() string {
	_, _, dir := this.getState()
	return dir
}

func (this *Robot) getState() (x, y int, dir string) {
	w, h, s := this.w, this.h, this.step
	switch {
	case s < w:
		return s, 0, "East"
	case s < w+h-1:
		return w - 1, s - w + 1, "North"
	case s < 2*w+h-2:
		return w - 1 - (s - w - h + 2), h - 1, "West"
	default:
		return 0, (w+h)*2 - s - 4, "South"
	}
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
