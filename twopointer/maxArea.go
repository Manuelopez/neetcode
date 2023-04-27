package twopointer

func MaxArea(height []int) int {
   l := 0 
   r := len(height) -1
   maxA := 0
   for l < r{
       h := 0
       if(height[l] < height[r]){
           h = height[l]
       }else{
           h = height[r]
        }

       area := h * (r - l)
       if area > maxA{
           maxA = area
       }
       if(height[l] < height[r]){
           l++
       }else{
           r--
        }
   }

   return maxA
}
