# Merge Sort

In computer science, merge sort (also commonly spelled 
mergesort) is an efficient, general-purpose, 
comparison-based sorting algorithm. Most implementations 
produce a stable sort, which means that the implementation 
preserves the input order of equal elements in the sorted 
output. Mergesort is a divide and conquer algorithm that 
was invented by John von Neumann in 1945.

An example of merge sort. First divide the list into 
the smallest unit (1 element), then compare each 
element with the adjacent list to sort and merge the 
two adjacent lists. Finally all the elements are sorted 
and merged.

![Merge Sort](https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif)

A recursive merge sort algorithm used to sort an array of 7 
integer values. These are the steps a human would take to 
emulate merge sort (top-down).

![Merge Sort](https://upload.wikimedia.org/wikipedia/commons/e/e6/Merge_sort_algorithm_diagram.svg)

## Complexity

| Name                  | Best            | Average             | Worst               | Memory    | Stable    | Comments  |
| --------------------- | :-------------: | :-----------------: | :-----------------: | :-------: | :-------: | :-------- |
| **Merge sort**        | n&nbsp;log(n)   | n&nbsp;log(n)       | n&nbsp;log(n)       | n         | Yes       |           |

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Merge_sort)
- [YouTube](https://www.youtube.com/watch?v=KF2j-9iSf4Q&index=27&list=PLLXdhg_r2hKA7DPDsunoDZ-Z769jWn4R8)

#Refs
![图解归并排序](https://images2015.cnblogs.com/blog/1024555/201612/1024555-20161218194508761-468169540.png)
- [图解归并排序](https://www.cnblogs.com/chengxiao/p/6194356.html#:~:text=%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F%EF%BC%88MERGE%2DSORT%EF%BC%89,%E5%9C%A8%E4%B8%80%E8%B5%B7%EF%BC%8C%E5%8D%B3%E5%88%86%E8%80%8C%E6%B2%BB%E4%B9%8B)%E3%80%82)