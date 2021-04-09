package redis

/*
1.ZADD KEY score member
将指定 member 和 score 添加到 key对应的有序集合中
参数:
XX: 仅仅更新存在的成员，不添加新成员。
NX: 不更新存在的成员。只添加新成员。
CH: 修改返回值为发生变化的成员总数，原始是返回新添加成员的总数 (CH 是 changed 的意思)。
更改的元素是新添加的成员，已经存在的成员更新分数。 所以在命令中指定的成员有相同的分数将不被计算在内。
注：在通常情况下，ZADD返回值只计算新添加成员的数量。
INCR: 当ZADD指定这个选项时，成员的操作就等同ZINCRBY命令，对成员的分数进行递增操作。

有序集合里面的成员是不能重复的都是唯一的，但是，不同成员间有可能有相同的分数。
当多个成员有相同的分数时，他们将是有序的字典（ordered lexicographically）
（仍由分数作为第一排序条件，然后，相同分数的成员按照字典规则相对排序）。

字典顺序排序用的是二进制，它比较的是字符串的字节数组。

redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 1 "uno"
(integer) 1
redis> ZADD myzset 2 "two" 3 "three"
(integer) 2
redis> ZRANGE myzset 0 -1 WITHSCORES
1) "one"
2) "1"
3) "uno"
4) "1"
5) "two"
6) "2"
7) "three"
8) "3"
redis>

2.ZCARD KEY   返回有序集合元素个数
3.ZCOUNT KEY MIN MAX   返回有序集key中，score值在min和max之间(默认包括score值等于min或max)的成员个数。
4.ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight] [SUM|MIN|MAX]
计算给定的numkeys个有序集合的交集，并且把结果放到destination中。
destination  -  交集结果集合key
numkeys  -   参与计算的key个数
key   -   子集key
weight  -   子集乘法因子，默认为1

redis> ZADD zset1 1 "one"
(integer) 1
redis> ZADD zset1 2 "two"
(integer) 1
redis> ZADD zset2 1 "one"
(integer) 1
redis> ZADD zset2 2 "two"
(integer) 1
redis> ZADD zset2 3 "three"
(integer) 1
redis> ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3
(integer) 2
redis> ZRANGE out 0 -1 WITHSCORES
1) "one"
2) "5"
3) "two"
4) "10"

将 zset1 zset2 的两个元素进行交集，乘法因子为 2，3
one = 1*2 + 1*3
two = 2*2 + 2*3

5.ZPOPMAX key [count]
删除并返回有序集合key中的最多count个具有最高得分的成员。count的默认值为1
6.ZPOPMIN key [count]
删除并返回有序集合key中的最多count个具有最低得分的成员。count的默认值为1
7.ZREVRANGE key start stop [WITHSCORES]
返回有序集key中，指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列。
WITHSCORES 表示返回成员以及分数信息

8.ZREVRANK key member
返回有序集key中成员member的排名，其中有序集成员按score值从大到小排列。排名以0为底
9.ZSCORE key member
返回有序集key中，成员member的score值
*/
