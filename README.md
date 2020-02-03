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

