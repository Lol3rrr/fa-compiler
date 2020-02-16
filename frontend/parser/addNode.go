package parser

func (node *PT_Node) addChild(childType, content string) {
  // Adds a new Child node to the given Node, with the provided parameters
  node.Children = append(node.Children, &PT_Node{
    Type: childType,
    Content: content,
    Parent: node,
    Children: make([]*PT_Node, 0),
  })
}
