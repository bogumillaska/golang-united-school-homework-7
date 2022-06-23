package coverage

import (
	"errors"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeopleLen(t *testing.T){
	people := People{Person{firstName: "Bogumil", lastName: "Laska", birthDay: time.Now()},
			Person{firstName: "Jon", lastName: "Kowalski", birthDay: time.Now()},
			Person{firstName: "Somebody", lastName: "Else", birthDay: time.Now()}}
	
	if people.Len() != 3 {
		t.Errorf("Wrong lenght of the people slice, got %v", people.Len())
	}
}

func TestPeopleLess(t *testing.T){
	people := People{Person{firstName: "Bogumil", lastName: "Laska", birthDay: time.Now()},
			Person{firstName: "Jon", lastName: "Kowalski", birthDay: time.Now()},
			Person{firstName: "Jon", lastName: "John", birthDay: time.Now()},
			Person{firstName: "Somebody", lastName: "Else", birthDay: time.Now().AddDate(+1, +1, +1)}}
	
	sort.Sort(people)
	if (people[0].firstName != "Somebody" || people[1].firstName != "Bogumil" || people[2].lastName != "John" || people[3].lastName != "Kowalski") {
		t.Errorf("wrong order to people slice : %v", people)
	}
}

func TestPeopleSwap(t *testing.T){
	people := People{Person{firstName: "Bogumil", lastName: "Laska", birthDay: time.Now()},
			Person{firstName: "Jon", lastName: "Kowalski", birthDay: time.Now()},
			Person{firstName: "Somebody", lastName: "Else", birthDay: time.Now().AddDate(+1, +1, +1)}}
	
	people.Swap(0,1)
	if people[0].firstName != "Jon" || people[1].firstName != "Bogumil" || people[2].firstName != "Somebody" {
		t.Errorf("wrong order to people slice : %v", people)
	}
}

func TestNewMatrixOK(t *testing.T) {
	matrix, err := New("1 2 3 4 5 6\n 6 7 1 2 3 4")
	if err != nil {
		t.Errorf("should not create matrix with strings : %v", err.Error())
	}
	if matrix.rows != 2 || matrix.cols != 6 {
		t.Errorf("wrong number of rows (%v) or columns (%v) : should be %v,%v", matrix.rows, matrix.cols, 2, 6)
	}
}

func TestNewMatrix_notSameLenght(t *testing.T) {
	_, err := New("1 2 3 4 5 6\n 6 7 1 2 3")
	if err.Error() != "Rows need to be the same length" {
		t.Error("should return error 'Rows need to be the same length'")
	}
}

func TestNewMatrixColsOK(t *testing.T) {
	assert := assert.New(t)
	matrix, _ := New("1 2 3 4 5 6\n 6 7 1 2 3 4")
	cols := matrix.Cols()
	assert.Equal(cols[0][0], 1)
}

func TestNewMatrixCelsOK(t *testing.T) {
	assert := assert.New(t)
	matrix, _ := New("1 2\n 3 4\n 5 6\n 6 7\n 1 2\n 3 4")
	cols := matrix.Cols()
	assert.Equal(cols[0][0], 1)
}

func TestNewMatrixSetOK(t *testing.T) {
	assert := assert.New(t)
	matrix, _ := New("1 2\n 3 4\n 5 6\n 6 7\n 1 2\n 3 4")
	matrix.Set(3, 1, 10)
	assert.Equal(matrix.Rows()[3][1], 10)
}

func TestNewMatrixSet_outOfRange(t *testing.T) {
	matrix, _ := New("1 2\n 3 4\n 5 6\n 6 7\n 1 2\n 3 4")
	result := matrix.Set(99, 99, 10)
	if result {
		t.Errorf("set should return false for out of range")
	}
}

func TestNewMatrixStringError(t *testing.T) {
	_, err := New("this is a new matrix")
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("should not create matrix with strings : %v", err.Error())
	}
}