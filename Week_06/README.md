学习笔记

64. 最小路径和

走到当前单元格(i,j)(i,j)的最小路径和为
“从左方单元格(i-1,j)(i−1,j)与从上方单元格 (i,j-1)(i,j−1)走来的两个最小路径和中较小的”+当前单元格值grid[i][j]。
具体分为以下4种情况：
1) 当左边和上边都不是矩阵边界时： 即当i!=0,j!=0时，
dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j]
2) 当只有左边是矩阵边界时： 只能从上面来，即当i=0,j!=0时，
dp[i][j] = dp[i][j-1]+grid[i][j]
3) 当只有上边是矩阵边界时： 只能从左面来，即当i!=0,j=0时，
dp[i][j] = dp[i-1][j]+grid[i][j]
4) 当左边和上边都是矩阵边界时： 即当i=0,j=0i时，其实就是起点，
dp[i][j] = grid[i][j]dp[i][j]=grid[i][j]

如果grid允许被修改，直接使用grid替代dp能减少空间复杂度


91. 解码方法

由于有效的编码数字范围是1-26,因此针对字符串中的某一位，其编码方式只有单独编码或与前一位合并编码两种方式，不会超过两位。
设置动态数组dp，dp[i]表示从起始到第i个字符，有多少编码方式。
那么第i个字符的编码方式，为字符i单独编码、或与前一个字符合并编码这两种方式之和，即
dp[i] = dp[i-1]) + dp[i-2]，
同时字符i存在编码无效的情况，一位编码的有效范围1-9，两位的有效范围10-26，加上一个系数count，编码有效为1，无效为0，即
dp[i] = dp[i-1]*count(i) + dp[i-2]*count(i-1,i)

dp数组的最后一个项即是最终结果
注意golang中string类型的坑，可先转为[]rune，再遍历。