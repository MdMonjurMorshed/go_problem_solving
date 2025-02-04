func maxOperations(nums []int, k int) int {

    value_map := make(map[int]int)
    process := 0
    for _,num := range(nums){
        diff := k - num
        if value_map[diff]>0{
            value_map[diff]-=1
            process+=1
        }else{
            value_map[num]+=1
        }
        


    }
    return process
    
}