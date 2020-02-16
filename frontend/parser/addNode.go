package parser

func (node *PT_Node) addChild(childType, content string) {
  node.Children = append(node.Children, &PT_Node{
    Type: childType,
    Content: content,
    Parent: node,
    Children: make([]*PT_Node, 0),
  })
}
