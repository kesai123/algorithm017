学习笔记

127 单词接龙
分别以初始单词，目标单词为起点，执行双向BFS。
双向BFS高效执行的关键：始终对较小队列执行BFS，控制BFS扩散的规模。
在编码层面，分别以初始词目标词为起点定义两个队列，每次BTS扩散一层，将较小队列定义为新的操作队列，
只对操作队列执行BFS，可简化编码。当操作队列中的词在另一个队列中，双向BFS相遇，接龙完成。
为简化单次在队列中查找，使用map来表示队列，执行一层BTS，进行队列替换。
对某一个单次的BTS执行，穷举其各个位置26个字母的变化，并且变化后的单次在合法单词集合里。

200 岛屿数量
依次以矩阵的每个陆地点为起点执行BFS或DFS，每执行依次BFS或DFS，将所有陆地点变为海洋点，并将岛屿数量加1.
实际执行过程中，由于BFS需要队列出入，实际执行效率低于DFS。
