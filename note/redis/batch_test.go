package redis

/*
redis 批量插入
cat data.txt | redis-cli --pipe

data.txt为命令集文件，例如：
SET Key0 Value0
SET Key1 Value1
...
SET KeyN ValueN

输出结果：
All data transferred. Waiting for the last reply...
Last reply received from server.
errors: 0, replies: 1000000
*/
