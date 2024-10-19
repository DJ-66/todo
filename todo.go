package todo
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task		string
	Done		bool
	CreatedAt	time.Time
	CompletedAt time.Time
}
// list ToDo items
type List []item

// add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item {
		Task:		task,
		Done: 		false,
		CreatedAt:	time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}
 // complete method markes items as done = true at current time
 func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	// adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
 }
 // Delete method removes items from list
 func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	// adjusting index for 0 based index
	*l = append(ls[:i-1], ls [i:]...)
	return nil
}
// Save method encodes the list in JSON and saves it
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}
// Get method opens the profided file name, decodes the JSON data and parses it into the list
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
			return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}




