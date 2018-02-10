Search trees are everywhere: In databases, in file systems, in board game algorithms,... This post explores the probably most basic form of a tree: a binary search tree.
<!--more-->
## The Problem
Searching a value in a linear data container (list, array) with *n* elements can take up to *n* steps. Click on "Play" in the animation below and see how many steps it takes to find the value "3" when this value is in the last element of a list container.
HYPE[Linear Search](LinearSearch.html)
Computer scientists say that this operation has an [*order of O(n)*](https://en.wikipedia.org/wiki/Big_O_notation). For very large values of *n*, this operation can become quite slow. Can we do better?
## The Solution
If the data is organized in a tree structure, access can be much faster. Rather than 6 steps in the above example, the search takes only two steps:
HYPE[Tree Search](TreeSearch.html)
The secret is that there is a sort order in this data structure. The search value is first compared against the value at the top of this structure. If the search value is smaller, the search continues with the next value to the left; else it continues with the next value to the right. Repeat until the value is found, or until there are no more values to the left or the right of the current value.
## Binary Search Tree Basics
But wait, what is this "tree structure" seen in the animation above? This structure is called a *binary search tree*. It has the following properties:
1. A tree consists of *nodes* that store unique values.
2. Each node has zero, one, or two child nodes.
3. One of the nodes is designated as the *root node* that is at the top of the tree structure. (This is the "entry point" where all operations start.)
3. Each node has exactly one parent node, except for the root node, which has no parent.
4. Each node's value is larger than the value of its left child but smaller than the value of its right child.
5. Each subtree to the left of a node only contains values that are smaller than the node's value, and each subtree to the right of a node contains only larger values.
Some quick definitions that help keeping the following text shorter:
1. A node with no children is called a *leaf node*.
2. A node with one child is called a *half leaf node*.
3. A node with two children is called an *inner node*.
4. The longest path from the root node to a leaf node is called the tree's *height*.
![Binary Tree Definitions](BinTreeDef.png)
In the best case, the height of a tree with *n* elements is log2(n+1) (where "log2" means the logarithm to base 2). All paths from the root node to a leaf node would have roughly the same length (plus/minus one). The tree is called a "balanced tree" in this case. Operations on this tree would have an *order* of *O(log(n)).
For large values of *n*, the logarithm of *n* is much smaller than *n* itself, which is why algorithms that need O(log(n)) time are much faster on average than algorithms that need O(n) time.
Take a calculator and find it out!
Let's say we have 1000 elements in our data store.
If the data store is a linear list, a search needs between 1 and 1000 steps to find a value. On average, this would be about 500 steps per search (if we assume the data is randomly distributed in this list).
If the data store is a balanced tree, a search needs at most log2(1001), or roughly 10 steps. What an improvement!
To visualize the difference, here is a diagram with a linear graph (in red) and a logarithmic graph (in blue). As the value on the x axis - the size of the data set - increases, the linear function keeps increasing with the same rate while the logarithmic function increases slower and slower the larger x gets.
![Logarithmic function versus linear function](logVsLinear.png)
Remember, O(log(n)) is only for the best case, where the tree is balanced, and no path to a leaf node is particularly longer than any other path.
In this post, however, we look at a very simple binary search tree. Especially, we do not care to minimize the height of the search tree. So in the worst case, a tree with *n* elements can have a height of *n*, which means it is not better than a linear list. In fact, in this case, the tree *is* effectively a linear list:
![Binary Tree Worst Case](BinTreeWorstCase.png)
So in the tree that we are going to implement, a search would take anywhere between O(log(n)) and O(n) time. In the next article, we'll see how to ensure that the tree is always balanced, so that a search always takes only O(log(n)) time.
## Today's code: A simple binary search tree
Let's go through implementing a very simple search tree. It has three operations: Insert, Delete, and Find. We also add a Traverse function for traversing the tree in sort order.
*/
