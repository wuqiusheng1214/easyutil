/*
@Time       : 2022/1/2
@Author     : wuqiusheng
@File       : math.go
@Description: 常用数据库
*/
package easyutil

import (
	"math"
	"math/rand"
)

//x^y
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

//返回x的二次方根
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

//取[0,number)区间的随机值
func RandNumber(number int) int {
	if number == 0 {
		return 0
	}
	return rand.Intn(number)
}

//[min,max)之间的随机数
func RandNumBetween(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min) + min
}

//正态分布:根据标准差和期望生成不同的正态分布值,sd标准差,mean期望
func RandNorm64(sd, mean int32) float64 {
	return rand.NormFloat64()*float64(sd) + float64(mean)
}

//正态分布,在一定范围内
func RandNormInt32(min, max, sd, mean int32) int32 {
	result := int32(Atoi(RoundStr("%0.0f", RandNorm64(sd, mean))))
	if result < min {
		return min
	}
	if result > max {
		return max
	}
	return result
}

//四舍五入保留n位小数,如保留整数"%.0f",保留3位小数"%.3f"...保留n位小数"%.nf"
func RoundStr(format string, decimal float64) string {
	return Sprintf(format, decimal)
}

//四舍五入保留n位小数
func Round(decimal float64, w int) float32 {
	format := "%." + Itoa(w) + "f"
	return Atof(Sprintf(format, decimal))
}

//四舍五入保留n位小数
func Round64(decimal float64, w int) float64 {
	format := "%." + Itoa(w) + "f"
	return Atof64(Sprintf(format, decimal))
}

//Log10为底x的对数
func Log10(x float64) float64 {
	return math.Log10(x)
}

//abs
func Abs(x float64) float64 {
	return math.Abs(x)
}

//int32
func MaxInt32() int32 {
	return math.MaxInt32
}

// 两个数的较大值
func Max(x, y int32) int32 {
	if x > y {
		return x
	}

	return y
}

// 两个数的较小值
func Min(x, y int32) int32 {
	if x > y {
		return y
	}

	return x
}

func SafeSubInt32(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return 0
}

func SafeSubUint32(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return 0
}
func SafeSubUint64(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return 0
}

func SafeSubInt64(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return 0
}

//Perm返回整数[0，n)的伪随机置换，返回n个整数的切片[0,n)
func RandPerm(n int) []int {
	return rand.Perm(n)
}

func RandString(count int) string {
	var randomstr string
	for r := 0; r < count; r++ {
		i := RandNumBetween(65, 90)
		a := rune(i)
		randomstr += string(a)
	}
	return randomstr
}

type valueWeightItem struct {
	weight uint32
	value  uint64
}

// 权值对，根据权重随机一个值出来
type GBValueWeightPair struct {
	allweight uint32
	valuelist []*valueWeightItem
}

func NewValueWeightPair() *GBValueWeightPair {
	vwp := new(GBValueWeightPair)
	vwp.valuelist = make([]*valueWeightItem, 0, 0)
	return vwp
}

func (r *GBValueWeightPair) Add(weight uint32, value uint64) {
	if weight == 0 {
		return
	}
	valueinfo := &valueWeightItem{weight, value}
	r.valuelist = append(r.valuelist, valueinfo)
	r.allweight += weight
}

func (r *GBValueWeightPair) Random() uint64 {
	if r.allweight > 0 {
		randvalue := uint32(rand.Intn(int(r.allweight))) + 1 //[1,allweight]
		addweight := uint32(0)
		for i := 0; i < len(r.valuelist); i++ {
			addweight += r.valuelist[i].weight
			if randvalue <= addweight {
				return r.valuelist[i].value
			}
		}
	}
	return 0
}

// 这个功能前提是value不重复
// 不重复随机
func (r *GBValueWeightPair) RandomUniqueList(count int, exceptSlice []uint64) []uint64 {
	if count <= 0 {
		return nil
	}
	retSlice := make([]uint64, 0, count)
	if count >= len(r.valuelist) {
		for _, v := range r.valuelist {
			retSlice = append(retSlice, v.value)
		}
		return retSlice
	}

	for i, j := 0, 0; i < count && j < 10*count; j++ {
		v := r.Random()
		found := false
		for _, vv := range retSlice {
			if v == vv {
				found = true
				break
			}
		}
		if !found {
			for _, vv := range exceptSlice {
				if v == vv {
					found = true
					break
				}
			}
		}
		if !found {
			i++
			retSlice = append(retSlice, v)
		}
	}
	if len(retSlice) < count {
		for _, v := range r.valuelist {
			found := false
			for _, vv := range retSlice {
				if vv == v.value {
					found = true
					break
				}
			}
			if !found {
				for _, vv := range exceptSlice {
					if v.value == vv {
						found = true
						break
					}
				}
			}
			if !found {
				retSlice = append(retSlice, v.value)
			}
		}
	}
	return retSlice
}

func Ternary(v bool, a, b interface{}) interface{} {
	if v {
		return a
	}
	return b
}
