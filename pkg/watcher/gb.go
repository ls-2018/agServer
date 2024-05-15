package watcher

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	apsv1 "my.domain/guestbook/apis/apps/v1"
	"sync"
	"time"
)

type Watch struct {
	one  sync.Once
	time time.Time
}

func (w *Watch) Stop() {
	fmt.Println("stop")
}

func (w *Watch) ResultChan() <-chan watch.Event {
	res := make(chan watch.Event)
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Second)
			res <- watch.Event{
				Type: watch.Added,
				Object: &apsv1.GuestBook{
					ObjectMeta: metav1.ObjectMeta{Name: time.Now().String()},
				},
			}
		}
		close(res)
	}()
	return res
}
