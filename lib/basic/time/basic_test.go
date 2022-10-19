package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeStop(t *testing.T) {

	// time After
	//var c chan int
	//select {
	//case m := <-c:
	//	fmt.Println("chan gain content ", m)
	//case <-time.After(10 * time.Second):
	//	fmt.Println("timed out")
	//}

	// ticker
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//done := make(chan bool)
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	done <- true
	//}()
	//for {
	//	select {
	//	case <-done:
	//		fmt.Println("Done!")
	//		return
	//	case t := <-ticker.C:
	//		fmt.Println("Current time: ", t)
	//	}
	//}

	//fmt.Println(time.Now())
	//time.Sleep(time.Second)
	//fmt.Println(time.Now())

	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, time.Now().String())
	}
}

func TestLocation(t *testing.T) {
	//loc := time.FixedZone("UTC-8", -8*60*60)
	//td := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	//fmt.Println("The time is:", td.Format(time.RFC822))

	//secondsEastOfUTC := int((8 * time.Hour).Seconds())
	//beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	//timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	//sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	//timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	//fmt.Println(timesAreEqual)
}

func TestTimeNowParse(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	fmt.Println(cur.Location())

	//start := time.Now()
	//time.Sleep(time.Second)
	//end := time.Now()
	//fmt.Println(end.Sub(start))

	//time := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	//fmt.Printf("Go launched at %s\n", time.Local())

	//h, _ := time.ParseDuration("4h30m")
	//fmt.Printf("I've got %.1f hours of work left.", h.Hours())

	//u, _ := time.ParseDuration("1s")
	//fmt.Printf("One second is %d microseconds.\n", u.Microseconds())

	//fmt.Println(1*time.Hour + 2*time.Minute + 300*time.Millisecond)
	//fmt.Println(3200 * time.Millisecond)

	//d, err := time.ParseDuration("1h15m30.918273645s")
	//if err != nil {
	//	panic(err)
	//}
	//
	//trunc := []time.Duration{
	//	time.Nanosecond,
	//	time.Microsecond,
	//	time.Millisecond,
	//	time.Second,
	//	2 * time.Second,
	//	time.Minute,
	//	10 * time.Minute,
	//	time.Hour,
	//}
	//
	//for _, t := range trunc {
	//	fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())
	//}

	//d, err := time.ParseDuration("1h15m30.918273645s")
	//if err != nil {
	//	panic(err)
	//}
	//
	//round := []time.Duration{
	//	time.Nanosecond,
	//	time.Microsecond,
	//	time.Millisecond,
	//	time.Second,
	//	2 * time.Second,
	//	time.Minute,
	//	10 * time.Minute,
	//	time.Hour,
	//}
	//
	//for _, r := range round {
	//	fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	//}

	td := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", td.Local())
}
