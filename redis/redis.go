package main

//import (
//	"fmt"
//	"github.com/go-redis/redis"
//	"math/rand"
//	"strconv"
//	"sync/atomic"
//	"time"
//)
//
//const (
//	letters     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//	randomLen   = 16
//	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
//	redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
//    return "OK"
//else
//    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
//end`
//	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
//    return redis.call("DEL", KEYS[1])
//else
//    return 0
//end`
//	tolerance       = 500
//	millisPerSecond = 1000
//)
//
//type redisLock struct {
//	store   *redis.Client
//	seconds uint32
//	key    string
//	value   string
//}
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
//
//func NewRedisClient(addr, password string, DB, PoolSize int) *redis.Client {
//	return redis.NewClient(&redis.Options{
//		Addr:     addr,
//		Password: password, // 密码
//		DB:       DB,       // 数据库
//		PoolSize: PoolSize, // 连接池大小
//	})
//}
//
//func NewRedisLock(store *redis.Client, key string) *redisLock {
//	return &redisLock{
//		store: store,
//		key:  key,
//		value: randomStr(randomLen),
//	}
//}
//
//func (rl *redisLock) Lock() (bool, error) {
//	seconds := atomic.LoadUint32(&rl.seconds)
//	resp := rl.store.Eval(lockCommand,[]string{rl.key},[]string{rl.value,strconv.Itoa(int(seconds)*millisPerSecond + tolerance)})
//
//}
//
//func randomStr(n int) string {
//	b := make([]byte, n)
//	for i := range b {
//		b[i] = letters[rand.Intn(len(letters))]
//	}
//	return string(b)
//}
//
//func main() {
//	fmt.Println(randomStr(16))
//}

//func modifyString(s string) int {
//	n := len(s)
//	cnt := 0
//	prev := byte(' ')
//	for i := 0; i < n; i++ {
//		if s[i] != '?' {
//			if s[i] == prev {
//				for j := byte('a'); j <= 'z'; j++ {
//					if j != prev {
//						s = s[:i] + string(j) + s[i+1:]
//						prev = j
//						cnt++
//						break
//					}
//				}
//			} else {
//				prev = s[i]
//			}
//		} else {
//			for j := byte('a'); j <= 'z'; j++ {
//				if j != prev {
//					if i == 0 || j != s[i-1] {
//						s = s[:i] + string(j) + s[i+1:]
//						prev = j
//						cnt++
//						break
//					}
//				}
//			}
//		}
//	}
//	return cnt
//}
//
//type Interval struct {
//	start int
//	end   int
//}
//
//type Intervals []Interval
//
//func (iv Intervals) Len() int           { return len(iv) }
//func (iv Intervals) Swap(i, j int)      { iv[i], iv[j] = iv[j], iv[i] }
//func (iv Intervals) Less(i, j int) bool { return iv[i].start < iv[j].start }
//
//func observeMostStars(intervals Intervals) (int, int) {
//	sort.Sort(intervals)
//	fmt.Println(intervals, 12313)
//	maxStars := 0
//	count := 0
//	currentStars := 0
//	for i := 0; i < len(intervals); {
//		j := i
//		for ; j < len(intervals) && intervals[j].start == intervals[i].start; j++ {
//			currentStars++
//		}
//		if currentStars > maxStars {
//			maxStars = currentStars
//			count = 1
//		} else if currentStars == maxStars {
//			count++
//		}
//		for k := i; k < j; k++ {
//			if intervals[k].end == intervals[i].start {
//				currentStars--
//			}
//		}
//		i = j
//	}
//	return maxStars, count
//}

//func main() {
//	// 读取输入
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	n, _ := strconv.Atoi(scanner.Text())
//
//	scanner.Scan()
//	sList := strings.Split(scanner.Text(), " ")
//	s := make([]int, n)
//	for i := 0; i < n; i++ {
//		s[i], _ = strconv.Atoi(sList[i])
//	}
//
//	scanner.Scan()
//	tList := strings.Split(scanner.Text(), " ")
//	t := make([]int, n)
//	for i := 0; i < n; i++ {
//		t[i], _ = strconv.Atoi(tList[i])
//	}
//
//	// 统计每个时刻流星的数量
//	cnt := make(map[int]int)
//	for i := 0; i < n; i++ {
//		for j := s[i]; j <= t[i]; j++ {
//			cnt[j]++
//		}
//	}
//
//	// 找出最多的流星数量
//	maxCnt := 0
//	for _, v := range cnt {
//		if v > maxCnt {
//			maxCnt = v
//		}
//	}
//
//	// 统计最多的流星数量出现的时刻数量
//	ans := 0
//	for _, v := range cnt {
//		if v == maxCnt {
//			ans++
//		}
//	}
//
//	// 输出结果
//	fmt.Println(maxCnt, ans)
//}
