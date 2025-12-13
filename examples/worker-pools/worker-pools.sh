# Запущенная программа показывает выполнение 5 задач
# разными воркерами. Программа занимает всего около
# 2 секунд, хотя общий объём работы составляет около
# 5 секунд, потому что 3 воркера работают конкурентно.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
