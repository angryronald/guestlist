package database

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/angryronald/guestlist/config"
	"github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func compareTwoStructs(expected interface{}, got interface{}, fieldsToCompare []string) error {
	expectedValue := reflect.Indirect(reflect.ValueOf(expected))
	gotValue := reflect.Indirect(reflect.ValueOf(got))

	for i, j := 0, 0; i < expectedValue.NumField() && j < len(fieldsToCompare); {
		if expectedValue.Type().Field(i).Name == fieldsToCompare[j] {
			if !reflect.DeepEqual(expectedValue.Field(i).Interface(), gotValue.Field(i).Interface()) {

				return fmt.Errorf("expected [%v]: %v, got: %v",
					fieldsToCompare[j],
					expectedValue.Field(i).Interface(),
					gotValue.Field(i).Interface())
			} else {
				i++
				j++
			}
		} else {
			i++
		}
	}
	return nil
}

func garbageCollector(object interface{}) {
	db := config.DB()
	db.Delete(object)
}

func TestFindByIDGuest(t *testing.T) {
	guestRepo := NewGuestPostgres()
	input, err := guestRepo.Insert(context.Background(), &repository.Guest{
		Name:               "abc",
		AccompanyingGuests: 10,
		Table:              1,
	})
	if err != nil {
		t.Errorf("Runtime error: %v", err)
	}

	input2, _ := uuid.NewRandom()

	testCases := map[string]struct {
		r        repository.GuestRepository
		input    uuid.UUID
		expected error
	}{
		"Object is found": {
			r:        guestRepo,
			input:    input.ID,
			expected: nil,
		},
		"Object is not found": {
			r:        guestRepo,
			input:    input2,
			expected: gorm.ErrRecordNotFound,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		_, err := c.r.FindByID(context.Background(), c.input)
		if err != c.expected {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected, err)
		}
	}

	defer garbageCollector(input)
}

func TestFindByNameGuest(t *testing.T) {
	guestRepo := NewGuestPostgres()
	key := "abc"
	input, err := guestRepo.Insert(context.Background(), &repository.Guest{
		Name:               key,
		AccompanyingGuests: 10,
		Table:              1,
	})
	if err != nil {
		t.Errorf("Runtime error: %v", err)
	}

	testCases := map[string]struct {
		r        repository.GuestRepository
		input    string
		expected *repository.Guest
	}{
		"Object is found": {
			r:        guestRepo,
			input:    key,
			expected: input,
		},
		"Object is not found": {
			r:        guestRepo,
			input:    "def",
			expected: nil,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.FindByName(context.Background(), c.input)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}
		// if !reflect.DeepEqual(output, c.expected) {
		if c.expected != nil && output != nil && c.expected.ID != output.ID {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected, output)
		}
	}

	defer garbageCollector(input)
}

func TestFindAllGuest(t *testing.T) {
	guestRepo := NewGuestPostgres()
	input, err := guestRepo.Insert(context.Background(), &repository.Guest{
		Name:               "abc",
		AccompanyingGuests: 10,
		Table:              1,
	})
	if err != nil {
		t.Errorf("Runtime error: %v", err)
	}

	testCases := map[string]struct {
		r        repository.GuestRepository
		input    bool
		expected int
	}{
		"There is a guest": {
			r:        guestRepo,
			input:    false,
			expected: 1,
		},
		"There is no arrived guest": {
			r:        guestRepo,
			input:    true,
			expected: 0,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		res, err := c.r.FindAll(context.Background(), c.input)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}
		if len(res) != c.expected {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected, len(res))
		}
	}

	defer garbageCollector(input)
}

func TestInsertGuest(t *testing.T) {
	guestRepo := NewGuestPostgres()

	testCases := map[string]struct {
		r               repository.GuestRepository
		input           *repository.Guest
		expected        *repository.Guest
		fieldsToCompare []string
	}{
		"Object is successfully inserted": {
			r: guestRepo,
			input: &repository.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			expected: &repository.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			fieldsToCompare: []string{"Name", "Table", "AccompanyingGuests"},
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.Insert(context.Background(), c.input)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}
		errCompare := compareTwoStructs(c.expected, output, c.fieldsToCompare)
		if errCompare != nil {
			t.Errorf("case [%s]: failed: %s", name, errCompare.Error())
		}

		garbageCollector(output)
	}
}

func TestUpdateGuest(t *testing.T) {
	guestRepo := NewGuestPostgres()

	testCases := map[string]struct {
		r               repository.GuestRepository
		initData        *repository.Guest
		updated         *repository.Guest
		expected        *repository.Guest
		fieldsToCompare []string
	}{
		"Object is successfully updated": {
			r: guestRepo,
			initData: &repository.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			updated: &repository.Guest{
				Name:               "def",
				AccompanyingGuests: 5,
				Table:              2,
			},
			expected: &repository.Guest{
				Name:               "def",
				AccompanyingGuests: 5,
				Table:              2,
			},
			fieldsToCompare: []string{"Name", "Table", "AccompanyingGuests"},
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.Insert(context.Background(), c.initData)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		output.Name = c.updated.Name
		output.Table = c.updated.Table
		output.AccompanyingGuests = c.updated.AccompanyingGuests

		output, err = c.r.Update(context.Background(), output)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		errCompare := compareTwoStructs(c.expected, output, c.fieldsToCompare)
		if errCompare != nil {
			t.Errorf("case [%s]: failed: %s", name, errCompare.Error())
		}

		garbageCollector(output)
	}
}

func TestDeleteGuest(t *testing.T) {
	guestRepo := NewGuestPostgres()

	testCases := map[string]struct {
		r        repository.GuestRepository
		initData *repository.Guest
		expected error
	}{
		"Object is successfully deleted": {
			r: guestRepo,
			initData: &repository.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			expected: gorm.ErrRecordNotFound,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.Insert(context.Background(), c.initData)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		err = c.r.Delete(context.Background(), output)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		_, err = c.r.FindByID(context.Background(), output.ID)
		if err != c.expected {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected.Error(), err.Error())
		}

		garbageCollector(output)
	}
}
