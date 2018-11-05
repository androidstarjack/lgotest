package queue

type Node [] interface{}

func (node *Node ) Push(value int){
	*node = append(*node,value)
}

func (node *Node) Pop() (int ){
	head :=  (*node)[0]
	*node =  (*node)[1:]
	return  head.(int)
}

func (node *Node) IsEmpety() bool{
	return len(*node) == 0
}

