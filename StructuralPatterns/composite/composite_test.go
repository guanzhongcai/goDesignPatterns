package composite

import "testing"

func TestNewComposite(t *testing.T) {

    root := NewComponent(CompositeNode, "root")

    root1 := NewComponent(CompositeNode, "root1")
    root2 := NewComponent(CompositeNode, "root2")
    root3 := NewComponent(CompositeNode, "root3")

    leaf1 := NewComponent(LeafNode, "leaf1")
    leaf2 := NewComponent(LeafNode, "leaf2")
    leaf3 := NewComponent(LeafNode, "leaf3")

    root.AddChild(root1)
    root.AddChild(root2)

    root1.AddChild(root3)
    root1.AddChild(leaf1)

    root2.AddChild(leaf2)
    root2.AddChild(leaf3)

    root.Print("")
    /*
     root
         root1
             root3
             leaf1
         root2
             leaf2
             leaf3
     */
}