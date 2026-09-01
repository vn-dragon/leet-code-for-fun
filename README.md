# LeetCode for Fun

A personal collection of LeetCode solutions written in **Go** and **Python**, created just for the fun of it.

**153 problems**, each solved in both languages, grouped by the technique it teaches.

## Topics Covered

| Topic | Solved | What it covers |
|---|---|---|
| [Arrays](#arrays) | 33 | Hashing, prefix sums, two pointers, greedy scans and string manipulation. |
| [Backtracking](#backtracking) | 10 | Subsets, permutations, combinations, grid search and partitioning. |
| [Binary Search](#binary-search) | 32 | Classic search on sorted input plus binary search on the answer space. |
| [Binary Tree](#binary-tree) | 29 | Traversals, recursion on subtrees, BST properties and tree construction. |
| [Linked List](#linked-list) | 28 | Pointer surgery, fast/slow pointers, dummy heads and cache design. |
| [Sliding Window](#sliding-window) | 16 | Fixed and variable windows, at-most-K counting, monotonic deques. |
| [Stack](#stack) | 5 | Monotonic stacks, matching pairs and stack-backed data structures. |

## Structure

```
daily/
├── Arrays/
├── Backtracking/
├── Binary Search/
├── Binary Tree/
├── Linked List/
├── Sliding Window/
└── Stack/

# Each problem is a matching pair, named <number>-<Title>:
#   daily/Sliding Window/904.Fruit-Into-Baskets.go
#   daily/Sliding Window/904.Fruit-Into-Baskets.py
```

## Using Solutions

Files are LeetCode submission snippets, not standalone programs — Go files are plain
functions under `package daily`, Python files expose a `Solution` class. Paste the body
straight into the LeetCode editor, or import it into a scratch file to run it locally.

| Language | Version | Shape |
|---|---|---|
| Go | 1.21+ | `package daily`, bare function (uses builtin `min`/`max`) |
| Python | 3.10+ | `class Solution` with the typed method |

## Problems

### Arrays

<details>
<summary>33 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 1 | Two sum | [Go](daily/Arrays/1-Two-sum.go) · [Python](daily/Arrays/1-Two-sum.py) |
| 15 | 3sum | [Go](daily/Arrays/15-3sum.go) · [Python](daily/Arrays/15-3sum.py) |
| 121 | Best time to buy and sell stock | [Go](daily/Arrays/121-Best-time-to-buy-and-sell-stock.go) · [Python](daily/Arrays/121-Best-time-to-buy-and-sell-stock.py) |
| 128 | Longest consecutive sequence | [Go](daily/Arrays/128-Longest-consecutive-sequence.go) · [Python](daily/Arrays/128-Longest-consecutive-sequence.py) |
| 136 | Single number | [Go](daily/Arrays/136-Single-number.go) · [Python](daily/Arrays/136-Single-number.py) |
| 167 | Two sum ii input array is sorted | [Go](daily/Arrays/167-Two-sum-ii-input-array-is-sorted.go) · [Python](daily/Arrays/167-Two-sum-ii-input-array-is-sorted.py) |
| 189 | Rotate array | [Go](daily/Arrays/189-Rotate-array.go) · [Python](daily/Arrays/189-Rotate-array.py) |
| 202 | Happy Number | [Go](daily/Arrays/202-Happy-Number.go) · [Python](daily/Arrays/202-Happy-Number.py) |
| 205 | Isomorphic strings | [Go](daily/Arrays/205-Isomorphic-strings.go) · [Python](daily/Arrays/205-Isomorphic-strings.py) |
| 633 | Sum of square numbers | [Go](daily/Arrays/633.Sum-of-square-numbers.go) · [Python](daily/Arrays/633.Sum-of-square-numbers.py) |
| 680 | Valid Palindrome II | [Go](daily/Arrays/680-Valid-Palindrome-II.go) · [Python](daily/Arrays/680-Valid-Palindrome-II.py) |
| 696 | Count binary substrings | [Go](daily/Arrays/696.Count-binary-substrings.go) · [Python](daily/Arrays/696.Count-binary-substrings.py) |
| 763 | Partition labels | [Go](daily/Arrays/763.Partition-labels.go) · [Python](daily/Arrays/763.Partition-labels.py) |
| 777 | Swap adjacent in lr string | [Go](daily/Arrays/777.Swap-adjacent-in-lr-string.go) · [Python](daily/Arrays/777.Swap-adjacent-in-lr-string.py) |
| 795 | Number of subarrays with bounded maximum | [Go](daily/Arrays/795.Number-of-subarrays-with-bounded-maximum.go) · [Python](daily/Arrays/795.Number-of-subarrays-with-bounded-maximum.py) |
| 809 | Expressive words | [Go](daily/Arrays/809.Expressive-words.go) · [Python](daily/Arrays/809.Expressive-words.py) |
| 821 | Shortest distance to a character | [Go](daily/Arrays/821.Shortest-distance-to-a-character.go) · [Python](daily/Arrays/821.Shortest-distance-to-a-character.py) |
| 832 | Flipping an image | [Go](daily/Arrays/832.Flipping-an-image.go) · [Python](daily/Arrays/832.Flipping-an-image.py) |
| 845 | Longest mountain array | [Go](daily/Arrays/845.Longest-mountain-array.go) · [Python](daily/Arrays/845.Longest-mountain-array.py) |
| 862 | Shortest subarray with sum at least k | [Go](daily/Arrays/862-Shortest-subarray-with-sum-at-least-k.go) · [Python](daily/Arrays/862-Shortest-subarray-with-sum-at-least-k.py) |
| 881 | Boats to Save People | [Go](daily/Arrays/881-Boats-to-Save-People.go) · [Python](daily/Arrays/881-Boats-to-Save-People.py) |
| 905 | Sort Array By Parity | [Go](daily/Arrays/905.Sort-Array-By-Parity.go) · [Python](daily/Arrays/905.Sort-Array-By-Parity.py) |
| 917 | Reverse only letters | [Go](daily/Arrays/917.Reverse-only-letters.go) · [Python](daily/Arrays/917.Reverse-only-letters.py) |
| 922 | Sort array by parity II | [Go](daily/Arrays/922.Sort-array-by-parity-II.go) · [Python](daily/Arrays/922.Sort-array-by-parity-II.py) |
| 923 | 3sum with multiplicity | [Go](daily/Arrays/923.3sum-with-multiplicity.go) · [Python](daily/Arrays/923.3sum-with-multiplicity.py) |
| 925 | Long Pressed Name | [Go](daily/Arrays/925.Long-Pressed-Name.go) · [Python](daily/Arrays/925.Long-Pressed-Name.py) |
| 942 | DI String Match | [Go](daily/Arrays/942.DI-String-Match.go) · [Python](daily/Arrays/942.DI-String-Match.py) |
| 948 | Bag of tokens | [Go](daily/Arrays/948.Bag-of-tokens.go) · [Python](daily/Arrays/948.Bag-of-tokens.py) |
| 1089 | Duplicates zero | [Go](daily/Arrays/1089.Duplicates-zero.go) · [Python](daily/Arrays/1089.Duplicates-zero.py) |
| 1304 | Find n unique integers sum up to zero | [Go](daily/Arrays/1304.Find-n-unique-integers-sum-up-to-zero.go) · [Python](daily/Arrays/1304.Find-n-unique-integers-sum-up-to-zero.py) |
| 1317 | Convert integer to the sum of two no zero integers | [Go](daily/Arrays/1317.Convert-integer-to-the-sum-of-two-no-zero-integers.go) · [Python](daily/Arrays/1317.Convert-integer-to-the-sum-of-two-no-zero-integers.py) |
| 1935 | Maximum number of words you can type | [Go](daily/Arrays/1935.Maximum-number-of-words-you-can-type.go) · [Python](daily/Arrays/1935.Maximum-number-of-words-you-can-type.py) |
| 3541 | Find most frequent vowel and consonant | [Go](daily/Arrays/3541.Find-most-frequent-vowel-and-consonant.go) · [Python](daily/Arrays/3541.Find-most-frequent-vowel-and-consonant.py) |

</details>

### Backtracking

<details>
<summary>10 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 22 | Generate parentheses | [Go](daily/Backtracking/22-Generate-parentheses.go) · [Python](daily/Backtracking/22-Generate-parentheses.py) |
| 39 | Combination sum | [Go](daily/Backtracking/39-Combination-sum.go) · [Python](daily/Backtracking/39-Combination-sum.py) |
| 40 | Combination sum ii | [Go](daily/Backtracking/40-Combination-sum-ii.go) · [Python](daily/Backtracking/40-Combination-sum-ii.py) |
| 46 | Permutations | [Go](daily/Backtracking/46-Permutations.go) · [Python](daily/Backtracking/46-Permutations.py) |
| 47 | Permutations ii | [Go](daily/Backtracking/47-Permutations-ii.go) · [Python](daily/Backtracking/47-Permutations-ii.py) |
| 77 | Combinations | [Go](daily/Backtracking/77-Combinations.go) · [Python](daily/Backtracking/77-Combinations.py) |
| 78 | Subsets | [Go](daily/Backtracking/78-Subsets.go) · [Python](daily/Backtracking/78-Subsets.py) |
| 79 | Word search | [Go](daily/Backtracking/79-Word-search.go) · [Python](daily/Backtracking/79-Word-search.py) |
| 90 | Subsets II | [Go](daily/Backtracking/90-Subsets-II.go) · [Python](daily/Backtracking/90-Subsets-II.py) |
| 131 | Palindrome partitioning | [Go](daily/Backtracking/131-Palindrome-partitioning.go) · [Python](daily/Backtracking/131-Palindrome-partitioning.py) |

</details>

### Binary Search

<details>
<summary>32 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 4 | Median of Two Sorted Arrays | [Go](daily/Binary%20Search/4-Median-of-Two-Sorted-Arrays.go) · [Python](daily/Binary%20Search/4-Median-of-Two-Sorted-Arrays.py) |
| 33 | Search in Rotated Sorted Array | [Go](daily/Binary%20Search/33-Search-in-Rotated-Sorted-Array.go) · [Python](daily/Binary%20Search/33-Search-in-Rotated-Sorted-Array.py) |
| 35 | Search Insert Position | [Go](daily/Binary%20Search/35.Search-Insert-Position.go) · [Python](daily/Binary%20Search/35.Search-Insert-Position.py) |
| 69 | Sqrt(x) | [Go](daily/Binary%20Search/69-Sqrt%28x%29.go) · [Python](daily/Binary%20Search/69-Sqrt%28x%29.py) |
| 74 | Search a 2D Matrix | [Go](daily/Binary%20Search/74-Search-a-2D-Matrix.go) · [Python](daily/Binary%20Search/74-Search-a-2D-Matrix.py) |
| 81 | Search in Rotated Sorted Array II | [Go](daily/Binary%20Search/81-Search-in-Rotated-Sorted-Array-II.go) · [Python](daily/Binary%20Search/81-Search-in-Rotated-Sorted-Array-II.py) |
| 153 | Find Minimum in Rotated Sorted Array | [Go](daily/Binary%20Search/153-Find-Minimum-in-Rotated-Sorted-Array.go) · [Python](daily/Binary%20Search/153-Find-Minimum-in-Rotated-Sorted-Array.py) |
| 162 | Find Peak Element | [Go](daily/Binary%20Search/162-Find-Peak-Element.go) · [Python](daily/Binary%20Search/162-Find-Peak-Element.py) |
| 240 | Search a 2d matrix | [Go](daily/Binary%20Search/240.Search-a-2d-matrix.go) · [Python](daily/Binary%20Search/240.Search-a-2d-matrix.py) |
| 275 | H index ii | [Go](daily/Binary%20Search/275-H-index-ii.go) · [Python](daily/Binary%20Search/275-H-index-ii.py) |
| 278 | First bad version | [Go](daily/Binary%20Search/278.First-bad-version.go) · [Python](daily/Binary%20Search/278.First-bad-version.py) |
| 378 | kth smallest element in a sorted matrix | [Go](daily/Binary%20Search/378-kth-smallest-element-in-a-sorted-matrix.go) · [Python](daily/Binary%20Search/378-kth-smallest-element-in-a-sorted-matrix.py) |
| 410 | Split array largest sum | [Go](daily/Binary%20Search/410-Split-array-largest-sum.go) · [Python](daily/Binary%20Search/410-Split-array-largest-sum.py) |
| 436 | Find right interval | [Go](daily/Binary%20Search/436-Find-right-interval.go) · [Python](daily/Binary%20Search/436-Find-right-interval.py) |
| 528 | Random pick with weight | [Go](daily/Binary%20Search/528-Random-pick-with-weight.go) · [Python](daily/Binary%20Search/528-Random-pick-with-weight.py) |
| 658 | Find K Closest Elements | [Go](daily/Binary%20Search/658-Find-K-Closest-Elements.go) · [Python](daily/Binary%20Search/658-Find-K-Closest-Elements.py) |
| 704 | Binary Search | [Go](daily/Binary%20Search/704-Binary-Search.go) · [Python](daily/Binary%20Search/704-Binary-Search.py) |
| 744 | Find smallest letter greater than target | [Go](daily/Binary%20Search/744.Find-smallest-letter-greater-than-target.go) · [Python](daily/Binary%20Search/744.Find-smallest-letter-greater-than-target.py) |
| 826 | Most profit assigning work | [Go](daily/Binary%20Search/826.Most-profit-assigning-work.go) · [Python](daily/Binary%20Search/826.Most-profit-assigning-work.py) |
| 852 | Peak index in a mountain array | [Go](daily/Binary%20Search/852-Peak-index-in-a-mountain-array.go) · [Python](daily/Binary%20Search/852-Peak-index-in-a-mountain-array.py) |
| 875 | Koko Eating Bananas | [Go](daily/Binary%20Search/875-Koko-Eating-Bananas.go) · [Python](daily/Binary%20Search/875-Koko-Eating-Bananas.py) |
| 911 | Online election | [Go](daily/Binary%20Search/911-Online-election.go) · [Python](daily/Binary%20Search/911-Online-election.py) |
| 981 | Time Based Key Value Store | [Go](daily/Binary%20Search/981-Time-Based-Key-Value-Store.go) · [Python](daily/Binary%20Search/981-Time-Based-Key-Value-Store.py) |
| 1011 | Capacity To Ship Packages Within D Days | [Go](daily/Binary%20Search/1011-Capacity-To-Ship-Packages-Within-D-Days.go) · [Python](daily/Binary%20Search/1011-Capacity-To-Ship-Packages-Within-D-Days.py) |
| 1201 | Ugly number iii | [Go](daily/Binary%20Search/1201-Ugly-number-iii.go) · [Python](daily/Binary%20Search/1201-Ugly-number-iii.py) |
| 1283 | Find the smallest divisor given a threshold | [Go](daily/Binary%20Search/1283-Find-the-smallest-divisor-given-a-threshold.go) · [Python](daily/Binary%20Search/1283-Find-the-smallest-divisor-given-a-threshold.py) |
| 1292 | Maximum side length of a square with sum less than or equal to threshold | [Go](daily/Binary%20Search/1292-Maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold.go) · [Python](daily/Binary%20Search/1292-Maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold.py) |
| 1300 | Sum of mutated array closest to target | [Go](daily/Binary%20Search/1300-Sum-of-mutated-array-closest-to-target.go) · [Python](daily/Binary%20Search/1300-Sum-of-mutated-array-closest-to-target.py) |
| 1482 | Minimum number of days to make m bouquets | [Go](daily/Binary%20Search/1482-Minimum-number-of-days-to-make-m-bouquets.go) · [Python](daily/Binary%20Search/1482-Minimum-number-of-days-to-make-m-bouquets.py) |
| 1498 | Number of subsequences that satisfy the given sum condition | [Go](daily/Binary%20Search/1498-Number-of-subsequences-that-satisfy-the-given-sum-condition.go) · [Python](daily/Binary%20Search/1498-Number-of-subsequences-that-satisfy-the-given-sum-condition.py) |
| 1802 | Maximum value at a given index in a bounded array | [Go](daily/Binary%20Search/1802-Maximum-value-at-a-given-index-in-a-bounded-array.go) · [Python](daily/Binary%20Search/1802-Maximum-value-at-a-given-index-in-a-bounded-array.py) |
| 1901 | Find a peak element ii | [Go](daily/Binary%20Search/1901-Find-a-peak-element-ii.go) · [Python](daily/Binary%20Search/1901-Find-a-peak-element-ii.py) |

</details>

### Binary Tree

<details>
<summary>29 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 94 | Binary tree inorder traversal | [Go](daily/Binary%20Tree/94-Binary-tree-inorder-traversal.go) · [Python](daily/Binary%20Tree/94-Binary-tree-inorder-traversal.py) |
| 98 | Validate binary search tree | [Go](daily/Binary%20Tree/98-Validate-binary-search-tree.go) · [Python](daily/Binary%20Tree/98-Validate-binary-search-tree.py) |
| 100 | Same tree | [Go](daily/Binary%20Tree/100-Same-tree.go) · [Python](daily/Binary%20Tree/100-Same-tree.py) |
| 102 | Binary tree level order traversal | [Go](daily/Binary%20Tree/102.Binary-tree-level-order-traversal.go) · [Python](daily/Binary%20Tree/102.Binary-tree-level-order-traversal.py) |
| 103 | Binary tree zigzag level order traversal | [Go](daily/Binary%20Tree/103-Binary-tree-zigzag-level-order-traversal.go) · [Python](daily/Binary%20Tree/103-Binary-tree-zigzag-level-order-traversal.py) |
| 104 | Maximum depth of binary tree | [Go](daily/Binary%20Tree/104-Maximum-depth-of-binary-tree.go) · [Python](daily/Binary%20Tree/104-Maximum-depth-of-binary-tree.py) |
| 105 | Construct binary tree from preorder and inorder traversal | [Go](daily/Binary%20Tree/105-Construct-binary-tree-from-preorder-and-inorder-traversal.go) · [Python](daily/Binary%20Tree/105-Construct-binary-tree-from-preorder-and-inorder-traversal.py) |
| 110 | Balanced binary tree | [Go](daily/Binary%20Tree/110-Balanced-binary-tree.go) · [Python](daily/Binary%20Tree/110-Balanced-binary-tree.py) |
| 111 | Minimum depth of binary tree | [Go](daily/Binary%20Tree/111-Minimum-depth-of-binary-tree.go) · [Python](daily/Binary%20Tree/111-Minimum-depth-of-binary-tree.py) |
| 112 | Path sum | [Go](daily/Binary%20Tree/112-Path-sum.go) · [Python](daily/Binary%20Tree/112-Path-sum.py) |
| 116 | Populating next right pointers in each node | [Go](daily/Binary%20Tree/116-Populating-next-right-pointers-in-each-node.go) · [Python](daily/Binary%20Tree/116-Populating-next-right-pointers-in-each-node.py) |
| 124 | Binary tree maximum path sum | [Go](daily/Binary%20Tree/124-Binary-tree-maximum-path-sum.go) · [Python](daily/Binary%20Tree/124-Binary-tree-maximum-path-sum.py) |
| 144 | Binary tree preorder traversal | [Go](daily/Binary%20Tree/144-Binary-tree-preorder-traversal.go) · [Python](daily/Binary%20Tree/144-Binary-tree-preorder-traversal.py) |
| 145 | Binary tree postorder traversal | [Go](daily/Binary%20Tree/145-Binary-tree-postorder-traversal.go) · [Python](daily/Binary%20Tree/145-Binary-tree-postorder-traversal.py) |
| 199 | Binary tree right side view | [Go](daily/Binary%20Tree/199-Binary-tree-right-side-view.go) · [Python](daily/Binary%20Tree/199-Binary-tree-right-side-view.py) |
| 226 | Invert binary tree | [Go](daily/Binary%20Tree/226-Invert-binary-tree.go) · [Python](daily/Binary%20Tree/226-Invert-binary-tree.py) |
| 230 | kth smallest element in a bst | [Go](daily/Binary%20Tree/230-kth-smallest-element-in-a-bst.go) · [Python](daily/Binary%20Tree/230-kth-smallest-element-in-a-bst.py) |
| 235 | Lowest common ancestor of a binary search tree | [Go](daily/Binary%20Tree/235-Lowest-common-ancestor-of-a-binary-search-tree.go) · [Python](daily/Binary%20Tree/235-Lowest-common-ancestor-of-a-binary-search-tree.py) |
| 257 | Binary tree paths | [Go](daily/Binary%20Tree/257-Binary-tree-paths.go) · [Python](daily/Binary%20Tree/257-Binary-tree-paths.py) |
| 297 | Serialize and deserialize binary tree | [Go](daily/Binary%20Tree/297-Serialize-and-deserialize-binary-tree.go) · [Python](daily/Binary%20Tree/297-Serialize-and-deserialize-binary-tree.py) |
| 437 | Path sum iii | [Go](daily/Binary%20Tree/437-Path-sum-iii.go) · [Python](daily/Binary%20Tree/437-Path-sum-iii.py) |
| 450 | Delete node in a bst | [Go](daily/Binary%20Tree/450-Delete-node-in-a-bst.go) · [Python](daily/Binary%20Tree/450-Delete-node-in-a-bst.py) |
| 543 | Diameter of binary tree | [Go](daily/Binary%20Tree/543-Diameter-of-binary-tree.go) · [Python](daily/Binary%20Tree/543-Diameter-of-binary-tree.py) |
| 572 | Subtree of another tree | [Go](daily/Binary%20Tree/572-Subtree-of-another-tree.go) · [Python](daily/Binary%20Tree/572-Subtree-of-another-tree.py) |
| 617 | Merge two binary trees | [Go](daily/Binary%20Tree/617-Merge-two-binary-trees.go) · [Python](daily/Binary%20Tree/617-Merge-two-binary-trees.py) |
| 653 | Two sum iv input is a bst | [Go](daily/Binary%20Tree/653.Two-sum-iv-input-is-a-bst.go) · [Python](daily/Binary%20Tree/653.Two-sum-iv-input-is-a-bst.py) |
| 654 | Maximum binary tree | [Go](daily/Binary%20Tree/654-Maximum-binary-tree.go) · [Python](daily/Binary%20Tree/654-Maximum-binary-tree.py) |
| 701 | Insert into a binary search tree | [Go](daily/Binary%20Tree/701-Insert-into-a-binary-search-tree.go) · [Python](daily/Binary%20Tree/701-Insert-into-a-binary-search-tree.py) |
| 1448 | Count good nodes in binary tree | [Go](daily/Binary%20Tree/1448-Count-good-nodes-in-binary-tree.go) · [Python](daily/Binary%20Tree/1448-Count-good-nodes-in-binary-tree.py) |

</details>

### Linked List

<details>
<summary>28 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 2 | Add two numbers | [Go](daily/Linked%20List/2-Add-two-numbers.go) · [Python](daily/Linked%20List/2-Add-two-numbers.py) |
| 19 | Remove nth node from end of list | [Go](daily/Linked%20List/19-Remove-nth-node-from-end-of-list.go) · [Python](daily/Linked%20List/19-Remove-nth-node-from-end-of-list.py) |
| 21 | Merge two sorted lists | [Go](daily/Linked%20List/21-Merge-two-sorted-lists.go) · [Python](daily/Linked%20List/21-Merge-two-sorted-lists.py) |
| 23 | Merge k sorted lists | [Go](daily/Linked%20List/23-Merge-k-sorted-lists.go) · [Python](daily/Linked%20List/23-Merge-k-sorted-lists.py) |
| 24 | Swap nodes in pairs | [Go](daily/Linked%20List/24-Swap-nodes-in-pairs.go) · [Python](daily/Linked%20List/24-Swap-nodes-in-pairs.py) |
| 25 | Reverse nodes in k group | [Go](daily/Linked%20List/25-Reverse-nodes-in-k-group.go) · [Python](daily/Linked%20List/25-Reverse-nodes-in-k-group.py) |
| 86 | Partition list | [Go](daily/Linked%20List/86-Partition-list.go) · [Python](daily/Linked%20List/86-Partition-list.py) |
| 92 | Reverse linked list ii | [Go](daily/Linked%20List/92-Reverse-linked-list-ii.go) · [Python](daily/Linked%20List/92-Reverse-linked-list-ii.py) |
| 138 | Copy list with random pointer | [Go](daily/Linked%20List/138-Copy-list-with-random-pointer.go) · [Python](daily/Linked%20List/138-Copy-list-with-random-pointer.py) |
| 141 | Linked list cycle | [Go](daily/Linked%20List/141-Linked-list-cycle.go) · [Python](daily/Linked%20List/141-Linked-list-cycle.py) |
| 142 | Linked list cycle II | [Go](daily/Linked%20List/142-Linked-list-cycle-II.go) · [Python](daily/Linked%20List/142-Linked-list-cycle-II.py) |
| 143 | Reorder list | [Go](daily/Linked%20List/143-Reorder-list.go) · [Python](daily/Linked%20List/143-Reorder-list.py) |
| 146 | LRU cache | [Go](daily/Linked%20List/146-LRU-cache.go) · [Python](daily/Linked%20List/146-LRU-cache.py) |
| 147 | Insertion sort list | [Go](daily/Linked%20List/147-Insertion-sort-list.go) · [Python](daily/Linked%20List/147-Insertion-sort-list.py) |
| 160 | Intersection of two linked lists | [Go](daily/Linked%20List/160-Intersection-of-two-linked-lists.go) · [Python](daily/Linked%20List/160-Intersection-of-two-linked-lists.py) |
| 206 | Reverse linked list | [Go](daily/Linked%20List/206-Reverse-linked-list.go) · [Python](daily/Linked%20List/206-Reverse-linked-list.py) |
| 328 | Odd even linked list | [Go](daily/Linked%20List/328-Odd-even-linked-list.go) · [Python](daily/Linked%20List/328-Odd-even-linked-list.py) |
| 445 | Add two numbers ii | [Go](daily/Linked%20List/445-Add-two-numbers-ii.go) · [Python](daily/Linked%20List/445-Add-two-numbers-ii.py) |
| 460 | LFU Cache | [Go](daily/Linked%20List/460-LFU-Cache.go) · [Python](daily/Linked%20List/460-LFU-Cache.py) |
| 725 | Split linked list in parts | [Go](daily/Linked%20List/725-Split-linked-list-in-parts.go) · [Python](daily/Linked%20List/725-Split-linked-list-in-parts.py) |
| 1019 | Next greater node in linked list | [Go](daily/Linked%20List/1019-Next-greater-node-in-linked-list.go) · [Python](daily/Linked%20List/1019-Next-greater-node-in-linked-list.py) |
| 1669 | Merge in between linked lists | [Go](daily/Linked%20List/1669-Merge-in-between-linked-lists.go) · [Python](daily/Linked%20List/1669-Merge-in-between-linked-lists.py) |
| 1721 | Swapping nodes in a linked list | [Go](daily/Linked%20List/1721-Swapping-nodes-in-a-linked-list.go) · [Python](daily/Linked%20List/1721-Swapping-nodes-in-a-linked-list.py) |
| 2058 | Find the minimum and maximum number of nodes between critical points | [Go](daily/Linked%20List/2058-Find-the-minimum-and-maximum-number-of-nodes-between-critical-points.go) · [Python](daily/Linked%20List/2058-Find-the-minimum-and-maximum-number-of-nodes-between-critical-points.py) |
| 2130 | Maximum twin sum of a linked list | [Go](daily/Linked%20List/2130-Maximum-twin-sum-of-a-linked-list.go) · [Python](daily/Linked%20List/2130-Maximum-twin-sum-of-a-linked-list.py) |
| 2181 | Merge nodes in between zeros | [Go](daily/Linked%20List/2181-Merge-nodes-in-between-zeros.go) · [Python](daily/Linked%20List/2181-Merge-nodes-in-between-zeros.py) |
| 2487 | Remove nodes from linked list | [Go](daily/Linked%20List/2487-Remove-nodes-from-linked-list.go) · [Python](daily/Linked%20List/2487-Remove-nodes-from-linked-list.py) |
| 3217 | Delete nodes from linked list present in array | [Go](daily/Linked%20List/3217-Delete-nodes-from-linked-list-present-in-array.go) · [Python](daily/Linked%20List/3217-Delete-nodes-from-linked-list-present-in-array.py) |

</details>

### Sliding Window

<details>
<summary>16 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 3 | Longest substring without repeating characters | [Go](daily/Sliding%20Window/3-Longest-substring-without-repeating-characters.go) · [Python](daily/Sliding%20Window/3-Longest-substring-without-repeating-characters.py) |
| 76 | Minimum window substring | [Go](daily/Sliding%20Window/76-Minimum-window-substring.go) · [Python](daily/Sliding%20Window/76-Minimum-window-substring.py) |
| 209 | Minimun size subarray sum | [Go](daily/Sliding%20Window/209-Minimun-size-subarray-sum.go) · [Python](daily/Sliding%20Window/209-Minimun-size-subarray-sum.py) |
| 219 | Contains duplicate II | [Go](daily/Sliding%20Window/219-Contains-duplicate-II.go) · [Python](daily/Sliding%20Window/219-Contains-duplicate-II.py) |
| 239 | Sliding window maximum | [Go](daily/Sliding%20Window/239-Sliding-window-maximum.go) · [Python](daily/Sliding%20Window/239-Sliding-window-maximum.py) |
| 424 | Longest repeating character replacement | [Go](daily/Sliding%20Window/424-Longest-repeating-character-replacement.go) · [Python](daily/Sliding%20Window/424-Longest-repeating-character-replacement.py) |
| 567 | Permutation in string | [Go](daily/Sliding%20Window/567-Permutation-in-string.go) · [Python](daily/Sliding%20Window/567-Permutation-in-string.py) |
| 904 | Fruit Into Baskets | [Go](daily/Sliding%20Window/904.Fruit-Into-Baskets.go) · [Python](daily/Sliding%20Window/904.Fruit-Into-Baskets.py) |
| 930 | Binary subarrays with sum | [Go](daily/Sliding%20Window/930-Binary-subarrays-with-sum.go) · [Python](daily/Sliding%20Window/930-Binary-subarrays-with-sum.py) |
| 992 | Subarrays with k different integers | [Go](daily/Sliding%20Window/992-Subarrays-with-k-different-integers.go) · [Python](daily/Sliding%20Window/992-Subarrays-with-k-different-integers.py) |
| 1004 | Max consecutive ones iii | [Go](daily/Sliding%20Window/1004-Max-consecutive-ones-iii.go) · [Python](daily/Sliding%20Window/1004-Max-consecutive-ones-iii.py) |
| 1208 | Get equal substrings within budget | [Go](daily/Sliding%20Window/1208-Get-equal-substrings-within-budget.go) · [Python](daily/Sliding%20Window/1208-Get-equal-substrings-within-budget.py) |
| 1234 | Replace the substring for balanced string | [Go](daily/Sliding%20Window/1234-Replace-the-substring-for-balanced-string.go) · [Python](daily/Sliding%20Window/1234-Replace-the-substring-for-balanced-string.py) |
| 1248 | Count number of nice subarrays | [Go](daily/Sliding%20Window/1248-Count-number-of-nice-subarrays.go) · [Python](daily/Sliding%20Window/1248-Count-number-of-nice-subarrays.py) |
| 1358 | Number of substrings containing all three characters | [Go](daily/Sliding%20Window/1358-Number-of-substrings-containing-all-three-characters.go) · [Python](daily/Sliding%20Window/1358-Number-of-substrings-containing-all-three-characters.py) |
| 1438 | Longest continuous subarray with absolute diff less than or equal to limit | [Go](daily/Sliding%20Window/1438-Longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit.go) · [Python](daily/Sliding%20Window/1438-Longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit.py) |

</details>

### Stack

<details>
<summary>5 problems</summary>

| # | Problem | Solutions |
|---|---|---|
| 20 | Valid parentheses | [Go](daily/Stack/20-Valid-parentheses.go) · [Python](daily/Stack/20-Valid-parentheses.py) |
| 155 | Min stack | [Go](daily/Stack/155-Min-stack.go) · [Python](daily/Stack/155-Min-stack.py) |
| 503 | Next greater element ii | [Go](daily/Stack/503-Next-greater-element-ii.go) · [Python](daily/Stack/503-Next-greater-element-ii.py) |
| 739 | Daily temperatures | [Go](daily/Stack/739-Daily-temperatures.go) · [Python](daily/Stack/739-Daily-temperatures.py) |
| 853 | Car fleet | [Go](daily/Stack/853-Car-fleet.go) · [Python](daily/Stack/853-Car-fleet.py) |

</details>
