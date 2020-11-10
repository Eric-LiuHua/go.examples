//分金币
package fjb

import "fmt"

func Fjb(str []string) {
	fmt.Println("fjb,str:", str)

	pc := make(map[string]int)
	for _, v := range str {
		uc := userJb(v)
		_, s := pc[v]
		if s {
			pc[v] = pc[v] + uc
		} else {
			pc[v] = uc
		}
	}

	for n, v := range pc {
		fmt.Printf("Fjb name:%s ,vcount:%d \n", n, v)
	}
}

//单个用户的
func userJb(name string) int {
	var sum int = 0
	for _, v := range name {
		switch v {
		case 'a', 'A':
			sum += 1

		case 'e', 'E':
			sum += 1

		case 'I', 'i':
			sum += 2

		case 'o', 'O':
			sum += 3

		case 'U', 'u':
			sum += 5

		}
	}
	return sum
}
