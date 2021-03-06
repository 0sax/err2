package err2

//v1.1.2

import (
	"fmt"
	"strings"
)

type ErrList struct {
	errs []string
}



func (e *ErrList) Add(s string) {
	e.errs = append(e.errs, s)
}

func (e *ErrList) AddF(s string, v... interface{}) {
	e.errs = append(e.errs, fmt.Sprintf(s, v...))
}

func (e *ErrList) Len() int {
	return len(e.errs)
}

func (e *ErrList) HasErrs() bool {
	return e.Len() != 0
}

func (e *ErrList) Errs() []string {
	return e.errs
}

func (e *ErrList) ErrsAsError() (err error) {

	if e.HasErrs() {
		ee := strings.Join(e.Errs(), ", ")
		err = fmt.Errorf(fmt.Sprintf("%v", ee))
	}

	return

}
