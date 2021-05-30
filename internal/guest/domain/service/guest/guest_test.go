package guest

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/angryronald/guestlist/config"
	constant "github.com/angryronald/guestlist/internal/guest"
	"github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	"github.com/angryronald/guestlist/internal/guest/infrastructure/repository/database"
	"github.com/angryronald/guestlist/internal/guest/public"
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

func TestGetGuest(t *testing.T) {
	guestRepo := database.NewGuestPostgres()
	guestService := NewService(guestRepo)
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
		r        ServiceInterface
		input    uuid.UUID
		expected error
	}{
		"Object is found": {
			r:        guestService,
			input:    input.ID,
			expected: nil,
		},
		"Object is not found": {
			r:        guestService,
			input:    input2,
			expected: gorm.ErrRecordNotFound,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		_, err := c.r.GetGuest(context.Background(), c.input)
		if err != c.expected {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected, err)
		}
	}

	defer garbageCollector(input)
}

func TestGetGuestByName(t *testing.T) {
	guestRepo := database.NewGuestPostgres()
	guestService := NewService(guestRepo)
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
		r        ServiceInterface
		input    string
		expected error
	}{
		"Object is found": {
			r:        guestService,
			input:    key,
			expected: nil,
		},
		"Object is not found": {
			r:        guestService,
			input:    "def",
			expected: constant.ErrNotFound,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		_, err := c.r.GetGuestByName(context.Background(), c.input)
		if c.expected != err {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected, err)
		}
	}

	defer garbageCollector(input)
}

func TestFindAllGuest(t *testing.T) {
	guestRepo := database.NewGuestPostgres()
	guestService := NewService(guestRepo)
	input, err := guestRepo.Insert(context.Background(), &repository.Guest{
		Name:               "abc",
		AccompanyingGuests: 10,
		Table:              1,
	})
	if err != nil {
		t.Errorf("Runtime error: %v", err)
	}

	testCases := map[string]struct {
		r        ServiceInterface
		input    bool
		expected int
	}{
		"There is a guest": {
			r:        guestService,
			input:    false,
			expected: 1,
		},
		"There is no arrived guest": {
			r:        guestService,
			input:    true,
			expected: 0,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		res, err := c.r.ListGuests(context.Background(), c.input)
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
	guestRepo := database.NewGuestPostgres()
	guestService := NewService(guestRepo)

	testCases := map[string]struct {
		r               ServiceInterface
		input           *public.Guest
		expected        *public.Guest
		fieldsToCompare []string
	}{
		"Object is successfully inserted": {
			r: guestService,
			input: &public.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			expected: &public.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			fieldsToCompare: []string{"Name", "Table", "AccompanyingGuests"},
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.CreateGuest(context.Background(), c.input)
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
	guestRepo := database.NewGuestPostgres()
	guestService := NewService(guestRepo)

	testCases := map[string]struct {
		r               ServiceInterface
		initData        *public.Guest
		updated         *public.Guest
		expected        *public.Guest
		fieldsToCompare []string
	}{
		"Object is successfully updated": {
			r: guestService,
			initData: &public.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			updated: &public.Guest{
				Name:               "def",
				AccompanyingGuests: 5,
				Table:              2,
			},
			expected: &public.Guest{
				Name:               "def",
				AccompanyingGuests: 5,
				Table:              2,
			},
			fieldsToCompare: []string{"Name", "Table", "AccompanyingGuests"},
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.CreateGuest(context.Background(), c.initData)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		output.Name = c.updated.Name
		output.Table = c.updated.Table
		output.AccompanyingGuests = c.updated.AccompanyingGuests

		output, err = c.r.UpdateGuest(context.Background(), output)
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
	guestRepo := database.NewGuestPostgres()
	guestService := NewService(guestRepo)

	testCases := map[string]struct {
		r        ServiceInterface
		initData *public.Guest
		expected error
	}{
		"Object is successfully deleted": {
			r: guestService,
			initData: &public.Guest{
				Name:               "abc",
				AccompanyingGuests: 10,
				Table:              1,
			},
			expected: gorm.ErrRecordNotFound,
		},
	}

	// cannot used deep equal since the uuid is unique
	for name, c := range testCases {
		output, err := c.r.CreateGuest(context.Background(), c.initData)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		err = c.r.DeleteGuest(context.Background(), output.ID)
		if err != nil {
			t.Errorf("case [%s]: Runtime error: %v", name, err)
		}

		_, err = c.r.GetGuest(context.Background(), output.ID)
		if err != c.expected {
			t.Errorf("case [%s]: failed: expected: %v, got: %v", name, c.expected.Error(), err.Error())
		}

		garbageCollector(output)
	}
}
