学习笔记

77 组合
全排列问题通常可利用回溯解决。
定义一个栈临时保存当前找到的一个合法组合。
递归查找过程可分为选择当前值，与不选择当前值两条路径。
对于回溯操作，如果选择当前值，在结束后需要将当前值弹栈。
当查找到后面的值，可以体现结束。判断剪枝条件加快算法效率。


105 从前序与中序遍历序列构造二叉树
由于可在前序遍历中确定根节点，那么在中序遍历序列中根据找到的根节点位置，可将序列分为左子树，右子树两个序列。
问题可以利用分治求解，左右子树的构建依然可以递归求解。

