# Go Crontab
### Introduction
GoCrontab is like linux crontab, and it is for doing repeating jobs for every certain amount of time.

It is easy to use, because the timer rule is same to linux crontab.
#
### How to use?
````
// Install the package
go get github.com/ChangsongLiQD/gocrontab
````

````go
package main

import (
	"fmt"
	"github.com/ChangsongLiQD/gocrontab"
	"time"
)

func main(){
	t := gocrontab.New()
	t.AddJob("0 * * * * *", func() {
		fmt.Println(0)
	})
	t.AddJob("30 * * * * *", func() {
		fmt.Println(30)
	})

	t.AddJob("45 * * * * *", func() {
		fmt.Println(45)
	})

	t.Start()
	time.Sleep(10000 * time.Second)
}
````

### Timer rule?

Example of job definition:
````
30 */2 * * * -> 30 minutes past the hour every 2 hours
15,45 23 * * * -> 11:15PM and 11:45PM every day
0 1 ? * SUN -> 1AM every Sunday
0 1 * * SUN -> 1AM every Sunday (same as above)
0 0 1 jan/2 * 2011-2013 ->
    midnight on January 1, 2011 and the first of every odd month until
    the end of 2013
24 7 L * * -> 7:24 AM on the last day of every month
24 7 * * L5 -> 7:24 AM on the last friday of every month
24 7 * * Lwed-fri ->
    7:24 AM on the last wednesday, thursday, and friday of every month
````
