## Two Sum

> Given an array of integers, return indices of the two numbers such that they add up to a specific target.
> You may assume that each input would have exactly one solution, and you may not use the same element twice.

> 给定一个整数数组`nums`和一个目标值`target`，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
> 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

**Example:**

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## Tasking

#### ？？

- [x] target 和 数组 负数是否考虑: yes
- [x] 最小和最大值的限制 -> int.min, int.max
- [x] 没有找到答案处理 -> return empty array
- [x] 错误处理 -> errors. yes
- [x] 数组是排序的吗？ no

### Tasking

- [x] happy pass for input array
- [x] 遍历查找
- [x] 没有答案时，返回空数组
- [x] 当输入不符合规则时提示错误
  - [x] 检查 nums 是否重复
  - [x] 检查是否多过一个，解决方案
  - [x] 检查最大值，最小值，和不能大于 int 最大值


## REF

- https://leetcode.com/problems/two-sum/
