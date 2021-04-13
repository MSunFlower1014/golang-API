package runtime

/*
select :
1.无case与default - 直接阻塞
2.只有一个case，且case的管道为nil - 直接阻塞
*/
