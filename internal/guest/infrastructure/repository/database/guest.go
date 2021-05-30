package database

import (
	"context"
	"errors"
	"time"

	"github.com/angryronald/guestlist/config"
	"github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	"github.com/angryronald/guestlist/lib/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// guestPostgres implements the guest repository service
type guestPostgres struct {
	db *gorm.DB
}

// FindByID get guest by guestId
func (s *guestPostgres) FindByID(ctx context.Context, guestID uuid.UUID) (*repository.Guest, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	guest := repository.Guest{}
	var err error
	if err = db.First(&guest, "`id` = ? AND `deleted_at` IS NULL", guestID).Error; err != nil {
		return nil, err
	}

	return &guest, nil
}

// FindByName get guest by name
func (s *guestPostgres) FindByName(ctx context.Context, name string) (*repository.Guest, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	guest := repository.Guest{}
	var err error
	err = db.First(&guest, "`name` = ? AND `deleted_at` IS NULL", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &guest, nil
}

// Insert inserts an guest
func (s *guestPostgres) Insert(ctx context.Context, singleGuest *repository.Guest) (*repository.Guest, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	singleGuest.ID, _ = uuid.NewRandom()
	singleGuest.CreatedAt = time.Now().UTC()
	singleGuest.UpdatedAt = time.Now().UTC()

	if err := db.Create(singleGuest).Error; err != nil {
		return nil, err
	}

	return singleGuest, nil
}

// Delete guest data
func (s *guestPostgres) Delete(ctx context.Context, guest *repository.Guest) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	currentGuest, err := s.FindByID(ctx, guest.ID)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	guest.CreatedAt = currentGuest.CreatedAt
	guest.UpdatedAt = currentGuest.UpdatedAt
	guest.DeletedAt = &now
	if err := db.Save(guest).Error; err != nil {
		return err
	}

	return nil
}

// Update guest data
func (s *guestPostgres) Update(ctx context.Context, updatedGuest *repository.Guest) (*repository.Guest, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	currentGuest, err := s.FindByID(ctx, updatedGuest.ID)
	if err != nil {
		return nil, err
	}

	updatedGuest.CreatedAt = currentGuest.CreatedAt
	updatedGuest.UpdatedAt = time.Now().UTC()
	if err := db.Save(updatedGuest).Error; err != nil {
		return nil, err
	}

	return updatedGuest, nil
}

//FindAll collect all guests data
func (s *guestPostgres) FindAll(ctx context.Context, isArrivedOnly bool) ([]*repository.Guest, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	guests := []repository.Guest{}
	args := []interface{}{}
	where := "`deleted_at` IS NULL"
	if isArrivedOnly {
		where += " AND `time_arrived` IS NOT NULL"
	}

	order := " `id` DESC"
	if err := db.Where(
		where,
		args...,
	).
		Order(order).
		Find(&guests).Error; err != nil {
		return nil, err
	}

	var result []*repository.Guest
	for i := 0; i < len(guests); i++ {
		result = append(result, &guests[i])
	}

	return result, nil
}

// NewGuestPostgres creates new guest repository service
func NewGuestPostgres() repository.GuestRepository {
	return &guestPostgres{
		db: config.DB(),
	}
}
