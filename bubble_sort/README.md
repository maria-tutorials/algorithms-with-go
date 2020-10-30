## 03 - Bubble sort

One of the slowest sorting algorithms we will look at.

`O(N^2)`

At a high level, we are "bubbling" large values to the top of a list.



Works by walking through a list and comparing consecutive pairs.

If next item goes before the current in the list we swap.

*Show example*

One pass will always move the largest unsorted item to the correct spot
So at most N passes to sort
Each pass takes N checks (roughly)
N^2 (more on big O later)
Or in code...

```go
for i := 0; i < len(list); i++ {
  for j := 0; j < len(list); j++ {
    if itemAtJ < itemAtJ+1 {
      swap
    }
  }
}
```

### Optimizations

There are two major optimizations you can make with bubble sort in your code.

1. Stop looking at sorted numbers

Each pass we sort an item
So on the Nth pass, N items should be sorted at the end.
We can code to skip these checks and it cuts our checks roughly in half.
Super minor optimization, but worth looking at.

2. Stop once our list is sorted

If we never make a swap, that means our list is sorted
No need to keep checking things if we have a sorted list!
This is part of why this algorithm can perform faster than others - it is better at recognizing an already sorted list.
We will learn more about where this algo is used later.
