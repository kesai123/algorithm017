学习笔记

LC 8: 字符串转换整数 (atoi) 
应注意的点有：
跳过空格
提取正负号
每计算一位，就进行范围检查
一旦发现非法字符，提前退出。

LC 438：找到字符串中所有字母异位词
由于字符串只包含小写字母，因此创建数组统计p的26个字母个数
针对s创建包含左右指针的滑动窗口，并创建统计窗口内26个字母个数的数组
首先滑动右指针，并增加右指针指向字符的统计个数。
如果统计个数超过p中对应字符统计个数，持续滑动左指针，直到其不大于p中统计个数，左指针滑动过的字符计数都要减1.
当左指针停止后，如果此时滑动窗口长度与p长度相同：
由于窗口内字符个数均不大于p字符个数，此时窗口子串字符个数一定与p相同，就找到了一个左指针起始对应的异位串。

