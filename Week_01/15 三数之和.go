func threeSum(nums []int) [][]int {
 	var ret [][]int
	length := len(nums)
	
	if length < 3 {
		return ret
	}
	//sort first
	sort.Ints(nums)
	//fix 1 member one by one, then find the others 2
	for i:=0;i<length-2;i++ {
		if nums[i]>0 {
			break
		}
		// compare left and right ptr in sorted sequence
		for l,r := i+1,length-1; l < r; {
			sum := nums[i]+nums[l]+nums[r]
            
			if sum < 0 { //move left
				// skip the same left member
				for(l<r && nums[l]==nums[l+1]){
					l++
				}
				l++
			} else if sum > 0 { //move right
				// skip the same right member
				for(l<r && nums[r]==nums[r-1]){
					r--
				}
				r--
			} else { //got one, move both left and right
				m := []int {nums[i], nums[l], nums[r]}
				ret = append(ret, m)
				// skip the same left member
				for(l<r && nums[l]==nums[l+1]){
					l++
				}
				l++
				// skip the same right member
				for(l<r && nums[r]==nums[r-1]){
					r--
				}
				r--
			}
		}
		//skip the same fix member
		for i<length-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}