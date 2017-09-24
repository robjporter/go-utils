package count

import (
	"github.com/robjporter/go-utils/go/as"
)

type Sequence <-chan int
type Sequence2 <-chan string

func getSequence(start, end int) Sequence {
	ch := make(chan int)
	go func() {
		if start < end {
			for i := start; i <= end; i++ {
				ch <- i
			}
		} else {
			for i := start; i >= end; i-- {
				ch <- i
			}
		}
		close(ch)
	}()

	return Sequence((<-chan int)(ch))
}

func getStringSequence(a []string) Sequence2 {
	ch := make(chan string)
	go func() {
		for i := 0; i < len(a); i++ {
			ch <- a[i]
		}
		close(ch)
	}()

	return Sequence2((<-chan string)(ch))
}

func CountUp(number int) Sequence {
	if number > 0 {
		return getSequence(0, number)
	}
	return nil
}

func CountUpS(start, number int) Sequence {
	if number > 0 {
		return getSequence(start, number)
	}
	return nil
}

func CountDown(number int) Sequence {
	if number > 0 {
		return getSequence(number, 0)
	}
	return nil
}

func CountDownS(number, start int) Sequence {
	if number > 0 {
		return getSequence(number, start)
	}
	return nil
}

func Loop(a interface{}) Sequence2 {
	typee := as.OfType(a)
	typee2 := as.OfType([]string{})
	if typee == typee2 {
		tmp := as.ToStringSlice(a)
		tmp2 := getStringSequence(tmp)
		return tmp2
	}
	return nil
}
