package logic

var memory int

// holdValue must be declare on memory
func Rotation(f []string) (t string) {
	var targetIndex int
	
	for i, s := range f {
		
		if  i < len(f) && memory < i {
			targetIndex = i
			t = s
			break
		}

		if memory == 6 {
			targetIndex = 0
			t = f[0]
		}

	}
	
	memory = targetIndex
	return
}

func GetMemory() (m int) {
	return memory
}