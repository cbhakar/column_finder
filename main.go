package main

import "fmt"

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}


func main()  {
	ctr := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}

	value := "AZ"
	const column  = 3
	const row  = 2
	var data [row][column]string

	if len(value) == 1{
		idx := indexOf(value, ctr )
		for i:=0;i<row;i++{
			for j:=0;j<column;j++{
				/*if idx == 25{
					data[i][j] = ctr[idx]
					idx = 0
				}else*/
				{
					data[i][j] = ctr[idx]
					idx++
				}
			}
		}
	}
	if len(value) == 2{
		idx1 := indexOf(string(value[0]), ctr )
		idx := indexOf(string(value[1]),ctr)
		for i:=0;i<row;i++{
			for j:=0;j<column;j++{
				if idx == 25{
					val := ctr[idx1]+""+ctr[idx]
					data[i][j] = val
					idx = 0
					idx1 +=1
				}else {
					val := ctr[idx1]+""+ctr[idx]
					data[i][j] = val
					idx++
				}
			}
		}
	}
	fmt.Println(data)

}
