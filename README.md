## leetcode-51

本题来自于经典的八皇后问题，这里是N皇后，不仅仅需要计算个数，而且输出皇后摆放的结果。因此：需要解决以下几个问题：

- N皇后问题基本逻辑：回溯法。一旦发生冲突则回溯，从第0行到n-1行遍历，依次判断每行的第0列-N-1列。因此采用用了**递归**的方式来回溯
- 确定是否发生冲突，一是记录对角线特征（主对角线：行-列=定值，从对角线：行+列=定值），二是记录某列是否已被占用
- 记录在哪一行的哪一列存放数据

可能遇到的坑：
- 起初使用数组来记录对角线和列是否被占用，后来发现golang无法初始化非常量长度的数组，于是改用了map
- 最坑的是本地和线上测试都是正确的，但是提交结果不正确，很气，神坑，还没找到原因


## leetcode-52

同上题思路，leetcode直接提交仍显示错误，本地测试和线上测试均无误。

## leetcode-53
在这里使用了分治法，时间复杂度为O(NlogN)，其实有O(N)的解法，这里只是为了练习分治法。注意：

- go里面没有while，统一用for，go里面没有直接判断整数大小的，需要自己实现d
- 注意数组不要越界
- 代码尽量抽象简洁 

## leetcode-54
明确四个值，startX,startY代表每一圈左上角起始位置，endX,endY代表每一圈的右下角结束点。然后一圈一圈环绕，根据这两个点所在的坐标，确定环绕时的横纵坐标。

- 注意横纵坐标的起始位置，> <和>= <=的区别
- 别忘了循环的退出条件，注意每一圈循环起始位置横纵坐标+1，结束位置横纵坐标-1
- 只要startX<endX并且startY<endY，就可以继续环绕的第三步
- 注意输入矩阵为空的情况

## leetcode-55
始终比较两个值，当前可走的步数和之间的数组里可走的最大步程。如果这个值在未抵达终点前始终大于0，则可以继续。

- 中途一旦出现不能抵达的情况，需要直接返回false.

## leetcode-56
首先按照二维数组的第一位进行排序，然后始终维护一个数组变量temp去保存当前的最大范围，如果下一个数组和该范围没有重叠部分，就把temp加进结果数组里面并更新temp数组。如果有重叠部分，则要视重叠情况更新temp数组。

以下几个注意点：

- 重叠分为两种情况，a)完全覆盖；b)部分覆盖；这两种情况更新状况不同。
- 遍历到最后一个数组的时候，需要区分考虑。原因：每次将temp数组加入结果数组时都是因为下一个区间与temp不重合，而最后一个数组是因为没有下一个区间了。
- 需要考虑数组为空的情况，还需要考虑数组个数为1的情况。
- 注意golang里面只有值传递，因此append temp数组时传进去的是地址，必须复制temp数组的值再添进去。从注释里面可以看到，原本使用的是复制值的办法，后来发现有个copy函数，不用一个值一个值的复制，写起来更简洁。
- 整理golang里面的sort方法。


### golang里面sort方法 ###
  1.普通排序 

    对切片排序，普通类型如int,float,string
    
  2.自定义排序

   a)结构体排序
   实现以下三个函数

    func (a Type)Len() int {}
    func (a Type)Less(i,j int) bool {}
    func (a Type)Swap(i,j int){}

   b)golang提供以下几个方法对slice进行排序

    func Slice(slice interface{}, less func(i, j int) bool) 
    func SliceStable(slice interface{}, less func(i, j int) bool) 
    func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool 
    func Search(n int, f func(int) bool) int
  

## leetcode-57

这里直接append再sort，套用了56的方法。事实上如果先找到插入的位置，也要newInterval能不能和前面的数组合并（最多只可能合并一个），再考虑后面的能不能全部合并，且必须遍历后面的数组。找插入点最快的二分查找复杂度o(logn),再合并应该是o(n)。而直接套用56的算法时间复杂度应该是o(nlogn)。

## leetcode-59
沿用前面螺旋矩阵的方法，确定了每层的左上角起始点和右下角结束点。解决该问题很容易。

值得注意的点：

- golang里面不能初始化长度为变量的数组，但能够初始化切片。
- go里面++/--运算符仅能放在变量后面（后自增或后自减），且不能用于赋值语句，只能作表达式。