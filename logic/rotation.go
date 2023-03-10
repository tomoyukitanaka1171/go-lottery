
// holdValue must be declare on memory
func Rotation() (holdValue int, target string) {
	var targetIndex int
	
	for i, s := range shikais {
		
		if  i < 7 && holdValue < i {
			targetIndex = i
			target = s
			break
		}

		if holdValue == 6 {
			targetIndex = 0
			target = shikais[0]
		}

	}
	
	holdValue = targetIndex
	return
}