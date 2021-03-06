## leetcode-51

本题来自于经典的八皇后问题，这里是 N 皇后，不仅仅需要计算个数，而且输出皇后摆放的结果。因此：需要解决以下几个问题：

- N 皇后问题基本逻辑：回溯法。一旦发生冲突则回溯，从第 0 行到 n-1 行遍历，依次判断每行的第 0 列-N-1 列。因此采用用了**递归**的方式来回溯
- 确定是否发生冲突，一是记录对角线特征（主对角线：行-列=定值，从对角线：行+列=定值），二是记录某列是否已被占用
- 记录在哪一行的哪一列存放数据

可能遇到的坑：

- 起初使用数组来记录对角线和列是否被占用，后来发现 golang 无法初始化非常量长度的数组，于是改用了 map
- 最坑的是本地和线上测试都是正确的，但是提交结果不正确，很气，神坑，还没找到原因

## leetcode-52

同上题思路，leetcode 直接提交仍显示错误，本地测试和线上测试均无误。

## leetcode-53

在这里使用了分治法，时间复杂度为 O(NlogN)，其实有 O(N)的解法，这里只是为了练习分治法。注意：

- go 里面没有 while，统一用 for，go 里面没有直接判断整数大小的，需要自己实现 d
- 注意数组不要越界
- 代码尽量抽象简洁

## leetcode-54

明确四个值，startX,startY 代表每一圈左上角起始位置，endX,endY 代表每一圈的右下角结束点。然后一圈一圈环绕，根据这两个点所在的坐标，确定环绕时的横纵坐标。

- 注意横纵坐标的起始位置，> <和>= <=的区别
- 别忘了循环的退出条件，注意每一圈循环起始位置横纵坐标+1，结束位置横纵坐标-1
- 只要 startX<endX 并且 startY<endY，就可以继续环绕的第三步
- 注意输入矩阵为空的情况

## leetcode-55

始终比较两个值，当前可走的步数和之间的数组里可走的最大步程。如果这个值在未抵达终点前始终大于 0，则可以继续。

- 中途一旦出现不能抵达的情况，需要直接返回 false.

## leetcode-56

首先按照二维数组的第一位进行排序，然后始终维护一个数组变量 temp 去保存当前的最大范围，如果下一个数组和该范围没有重叠部分，就把 temp 加进结果数组里面并更新 temp 数组。如果有重叠部分，则要视重叠情况更新 temp 数组。

以下几个注意点：

- 重叠分为两种情况，a)完全覆盖；b)部分覆盖；这两种情况更新状况不同。
- 遍历到最后一个数组的时候，需要区分考虑。原因：每次将 temp 数组加入结果数组时都是因为下一个区间与 temp 不重合，而最后一个数组是因为没有下一个区间了。
- 需要考虑数组为空的情况，还需要考虑数组个数为 1 的情况。
- 注意 golang 里面只有值传递，因此 append temp 数组时传进去的是地址，必须复制 temp 数组的值再添进去。从注释里面可以看到，原本使用的是复制值的办法，后来发现有个 copy 函数，不用一个值一个值的复制，写起来更简洁。
- 整理 golang 里面的 sort 方法。

### golang 里面 sort 方法

1.普通排序

    对切片排序，普通类型如int,float,string

2.自定义排序

a)结构体排序
实现以下三个函数

    func (a Type)Len() int {}
    func (a Type)Less(i,j int) bool {}
    func (a Type)Swap(i,j int){}

b)golang 提供以下几个方法对 slice 进行排序

    func Slice(slice interface{}, less func(i, j int) bool)
    func SliceStable(slice interface{}, less func(i, j int) bool)
    func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
    func Search(n int, f func(int) bool) int

## leetcode-57

这里直接 append 再 sort，套用了 56 的方法。事实上如果先找到插入的位置，也要 newInterval 能不能和前面的数组合并（最多只可能合并一个），再考虑后面的能不能全部合并，且必须遍历后面的数组。找插入点最快的二分查找复杂度 o(logn),再合并应该是 o(n)。而直接套用 56 的算法时间复杂度应该是 o(nlogn)。

## leetcode-59

沿用前面螺旋矩阵的方法，确定了每层的左上角起始点和右下角结束点。解决该问题很容易。

值得注意的点：

- golang 里面不能初始化长度为变量的数组，但能够初始化切片。
- go 里面++/--运算符仅能放在变量后面（后自增或后自减），且不能用于赋值语句，只能作表达式。

## leetcode-60

解法 1 这题在本地和测试用例计算也是正确的，提交不正确，不知道原因.

解法 2 参考了评论区的解法，与解法 1 的思路其实是一样的，但是解法 2 写的简洁的多，而且通过了提交。

从解法 2 可以学习到以下三点：

- 在求 n!时，由于要从 n 的阶层求到 1 的阶层，可以直接用一个数组来保存结果，可以节省大量时间。也无需再另外开辟函数了。
- 解法 1 讨论了 k/factorial(n-1)能不能除尽的问题，因为如果刚好除尽，那么该数字属于上一个区间，如果除不尽，才属于下一个区间。事实上，直接将 k--,再进行计算，可以完美避开这个问题。
- 求剩下的数字中大小排号为 k 的数字，解法 1 中选择使用一个 map 记录哪些数字被使用了，然后再计数到 k 个数字，这样又消耗了 O(n)的时间，解法 2 直接使用了一个切片，用掉了就使用 append 方法切割掉这个数字，效率更高。

疑问：

解法 1 为什么不能通过测试，会不会是 leetcode 在提交时对全局变量 map 的处理不符合预期不得而知，不过不管怎么说，解法 2 要强很多。

## leetcode-61

先把单向链表构造成了循环链表，然后从表头处移动 len-k 步即可。

注意点：这里可能出现 k>len 的情况，注意先取余

## leetcode-62

略

## leetcode-63

沿用 62 题的做法。唯一区别在于，遇到为 1 的情况把路径数置为 0，需额外考虑矩阵第一位数为 1 的情况。

## leetcode-64

原本用递归写了，结果递归时间超了，递归会产生大量的重复计算。所以后来还是选择使用数组，和 62，63 一样的思路。

## leetcode-65

不管怎么说，用正则是最简单的方式。长时间不用会忘记正则的用法。注意：2. .2 这种情况也被标注为 true。用有限状态机的思路很牛。

[golang 正则](https://github.com/google/re2/wiki/Syntax)
