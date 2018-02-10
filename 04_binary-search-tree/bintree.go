package main

import (
	"errors"
	"fmt"
	"log"
)

/*
## A Tree Node
Based on the above definition of a binary tree, a tree node consists of
* a value,
* a left subtree, and
* a right subtree.
By the way, this is a *recursive* data structure: Each subtree of a node is also a node containing subtrees.
In this minimal setup, the tree contains simple string data.
*/

// `Node` contains the search value, some data, a left child node, and a right child node.
type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

/* ## Node Operations
### Insert
To insert a value into a sorted tree, we need to find the correct place to add a new node with this value. Here is how:
1. Start at the root node.
2. Compare the new value with the current node's value.
	* If it is the same, stop. We have that value already.
	* If it is smaller, repeat 2. with the left child node. If there is no left child node, add a new one with the new value. Stop.
	* IF it is greater, repeat 2. with the right child node, or create a new one if none exists. Stop.
Sounds quite easy, doesn't it? Just keep in mind we do not take care of keeping the tree balanced. Doing so adds a bit of complexity but for now we don't care about this.
HYPE[Insert](TreeInsert.html)
The Insert method we define here works *recursively*. That is, it calls itself but with one of the child nodes as the new receiver. If you are unfamiliar with recursion, see the little example [here](https://en.wikipedia.org/wiki/Recursion#In_computer_science) or have a look at [this factorial function](https://play.golang.org/p/feMIAgYWg3).
*/

// `Insert` inserts new data into the tree, at the position determined by the search value.
// Return values:
//
// * `true` if the data was successfully inserted,
// * `false` if the data value already exists in the tree.
func (n *Node) Insert(value, data string) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	// If the data is already in the tree, return.
	case value == n.Value:
		return nil
		// If the data value is less than the current node's value, and if the left child node is `nil`, insert a new left child node. Else call `Insert` on the left subtree.
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
		// If the data value is greater than the current node's value, do the same but for the right subtree.
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

/*
### Find
Finding a value works as seen in the second animation of this article. (Hence, no animation here.) The Find method is also recursive.
It returns either the data of the found node and `true`, or "" and `false` if the node is not found.
*/

// `Find` searches for a string. It returns:
//
// * The data associated with the value and `true`, or
// * "" and `false` if the search string is not found in the tree.
func (n *Node) Find(s string) (string, bool) {

	if n == nil {
		return "", false
	}

	switch {
	// If the current node contains the value, return the node.
	case s == n.Value:
		return n.Data, true
		// If the data value is less than the current node's value, call `Find` for the left child node,
	case s < n.Value:
		return n.Left.Find(s)
		// else call `Find` for the right child node.
	default:
		return n.Right.Find(s)
	}
}

/*
### Delete
Deleting a value is a bit more complicated. Or rather, it is easy for two cases, and complicated for the third.
*The easy cases:*
**Delete a leaf node.**
This one is dead simple: Just set the parent's pointer to the node to nil.
**Delete a half-leaf node.**
Still easy: Replace the node by its child node. The tree's order remains intact.
*The complicated case:*
**Delete an inner node.**
This is the interesting part. The node to be deleted has two children, and we cannot assign both to the deleted node's parent node. Here is how this is solved. For simplicity, we assume that the node to be deleted is the *right* child of its parent node. The steps also apply if the node is the *left* child; you only need to swap "right" for "left", and "large" for "small".
1. In the node's left subtree, find the node with the largest value. Let's call this node "Node B".
2. Replace the node's value with B's value.
3. If B is a leaf node or a half-leaf node, delete it as described above for the leaf and half-leaf cases.
4. If B is an inner node, recursively call the Delete method on this node.
The animation shows how the root node is deleted. This is a simple case where "Node B" is a half-leaf node and hence does not require a recursive delete.
HYPE[Delete](TreeDelete.html)
To implement this, we first need two helper functions. The first one finds the maximum element in the subtree of the given node. The second one removes a node from its parent. To do so, it first determines if the node is the left child or the right child. Then it replaces the appropriate pointer with either nil (in the leaf case) or with the node's child node (in the half-leaf case).
*/

// `findMax` finds the maximum element in a (sub-)tree. Its value replaces the value of the
// to-be-deleted node.
// Return values: the node itself and its parent node.
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// `replaceNode` replaces the `parent`'s child pointer to `n` with a pointer to the `replacement` node.
// `parent` must not be `nil`.
func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

// `Delete` removes an element from the tree.
// It is an error to try deleting an element that does not exist.
// In order to remove an element properly, `Delete` needs to know the node's parent node.
// `parent` must not be `nil`.
func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	// Search the node to be deleted.
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		// We found the node to be deleted.
		// If the node has no children, simply remove it from its parent.
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		// If the node has one child: Replace the node with its child.
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		// If the node has two children:
		// Find the maximum element in the left subtree...
		replacement, replParent := n.Left.findMax(n)

		//...and replace the node's value and data with the replacement's value and data.
		n.Value = replacement.Value
		n.Data = replacement.Data

		// Then remove the replacement node.
		return replacement.Delete(replacement.Value, replParent)
	}
}

/*
## The Tree
One of a binary tree's nodes is the root node - the "entry point" of the tree.
The Tree data type wraps the root node and applies some special treatment. Especially, it handles the cases where the tree is completely empty or consists of a single node.
The Tree data type also provides an additional function for traversing the whole tree.
*/

// A `Tree` basically consists of a root node.
type Tree struct {
	Root *Node
}

// `Insert` calls `Node.Insert` unless the root node is `nil`
func (t *Tree) Insert(value, data string) error {
	// If the tree is empty, create a new node,...
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	// ...else call `Node.Insert`.
	return t.Root.Insert(value, data)
}

// `Find` calls `Node.Find` unless the root node is `nil`
func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

// `Delete` has one special case: the empty tree. (And deleting from an empty tree is an error.)
// In all other cases, it calls `Node.Delete`.
func (t *Tree) Delete(s string) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// Call`Node.Delete`. Passing a "fake" parent node here *almost* avoids
	// having to treat the root node as a special case, with one exception.
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	// If the root node is the only node in the tree, and if it is deleted,
	// then it *only* got removed from `fakeParent`. `t.Root` still points to the old node.
	// We rectify this by setting t.Root to nil.
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// `Traverse` is a simple method that traverses the tree in left-to-right order
// (which, *by pure incidence* ;-), is the same as traversing from smallest to
// largest value) and calls a custom function on each node.
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

/* ## A Couple Of Tree Operations
Our `main` function does a quick sort by filling a tree and reading it out again. Then it searches for a particular node. No fancy output to see here; this is just the proof that the whole code above works as it should.
*/

// `main`
func main() {

	// Set up a slice of strings.
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	// Create a tree and fill it from the values.
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	// Print the sorted values.
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Find values.
	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	// Delete a value.
	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Special case: A single-node tree. (See `Tree.Delete` about why this is a special case.)
	fmt.Println("Single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

}

