package main

func main() {
	newNode := Node{nil, 10, nil}
	newNode.AddNode(20)
	newNode.AddNode(30)
	newNode.AddNode(40)
	newNode.PrintNode()
	newNode.DeleteLast()
	newNode.PrintNode()
	newNode.AddNode(50)
	newNode.AddNode(60)
	newNode.AddNode(70)
	newNode.PrintNode()
	newNode.DeleteLast()
	newNode.PrintNode()
}
