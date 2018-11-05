package main

import (
	"fmt"
	"os"
)

func readMiGonFile(finePath string)[][] int{
	var row ,column int
	file, err := os.Open(finePath)

	if err != nil{
		panic(err)
	}
	fmt.Fscanf(file,"%d %d ",&row,&column)

	fileArray := make([][]int,row)
	//然后进行赋值
	//for index ,cs := range fileArray{
	//	fmt.Println(index,cs)
	//}
	for index  := range fileArray{
		fileArray[index] = make([]int,column)
		for j := range fileArray[index]{
			fmt.Fscanf(file,"%d",&fileArray[index][j])
		}
	}
	return fileArray
}



type point struct {
	i ,j  int
}
func main() {
	dataArray := readMiGonFile("simple/migong.in")
	for _ , columnArray := range dataArray{
		for _ ,val:= range columnArray{

			//fmt.Printf("fileArray[%d][%d] = %d",index,jndex,fileArray[index][jndex])
			//fmt.Printf("[%d][%d]=%d ",i,j, val)
			fmt.Printf("%d " ,val)
		}
		fmt.Println("")
	}
	//目的是：对这个数组进行操作，是的每个点进行移动，躲避障碍点，达到指定的终点
	start := point{
		0,0,
	}
	end := point{
		i : len(dataArray) - 1,
		j :len(dataArray[0]) - 1,//针对于切片，必须要减去1，
	}
	steps := wolrk(dataArray,start ,end)
	for _ ,row :=range steps{
		for val := range row{
			fmt.Printf("%d ",val)
		}
		fmt.Println()
	}
}
//需要定义四个方向  上左下右 i,j
var direct = [4] point{
	 {-1,0,},
	 {0,-1,},
	 {1,0,},
	 {0,1,},
}
func (p point) addDiri(cur point) (point){
	return point{p.i+cur.i,p.j+cur.j}
}
//返回当前的值
func  (p point) at(grid [][]int) (int,bool) {
	if p.i < 0 || p.i > len(grid) { //第一行往上往下 不越界
		return 0,false
	}
	if p.j < 0 || p.j >= len(grid[p.i]){ //第一行过完之后，
		return 0,false
	}
	return grid[p.i][p.j] ,true
}
func wolrk(dataArray [][]int, start , end point) [][] int {
	stepts := make([][] int,len(dataArray))
	for i :=range  stepts {
		stepts[i] = make([]int,len(dataArray[i]))
	}
	//搞一个起点,得到的就是开始的点
	Q :=  [] point{start}

	for len(Q)>0{
		cur := Q[0]
		Q = Q[1:]

		//发现终点
		if cur == end {
			break
		}

		//需要定义四个方向
		//遍历这四个方向
		for _,dir := range direct{
			next := cur.addDiri(dir)
			//判断限制 1.stepts  next为0 才能探索，dataArray next 才能探索 ，并且 next != start
			val ,ok   := next.at(dataArray)
			if !ok || val == 1{//1是撞墙啦
				continue
			}
			val ,ok   =next.at(stepts)
			if !ok || val  != 0{
				continue
			}
			if next == start {
				continue
			}
			//接下来，填充stepts
			//得到当前的步骤点数
			surrentStep ,_ := cur.at(stepts)
			stepts[next.i][next.j] = surrentStep
			Q = append(Q,next)
		}
	}
	return stepts
}
