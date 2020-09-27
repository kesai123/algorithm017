func maxArea(height []int) int {
    max := 0
    for l, r, area := 0, len(height)-1, 0; l<r; {
        if height[l]<height[r]{
            area = height[l]*(r-l)
            l++
        } else {
            area = height[r]*(r-l)
            r--
        }
        if area > max {
            max = area 
        }
    }
    return max
}