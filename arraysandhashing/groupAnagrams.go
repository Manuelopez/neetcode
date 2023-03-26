package arraysandhashing


// my current own solution O(n^3) * too slow *
func groupAnagrams(strs []string) [][]string {
    a := make([][]int)

    m := make([len(strs)]map[int]interface{})
    currIndex = 0
    for i, str := range strs{
        tm := make(map[int]interface{})
        for _, c := range str{
            if _, ok:= tm[c]; !ok{
                tm[c]= 1
            }else{
                tm[c] = tm[c].(int) + 1
            }
        }
        tm[-1] = str
        
        same := true
        for _, cm := range m{
            if(len(cm) != len(tm)){
                continue
            }
            for key, val := tm{
                if(key.(int) == -1){
                    continue
                }
                if amount, ok := cm[key]; !ok || amount.(int) != val.(int){
                    same = false
                }
            }
            
        }
    }
    

