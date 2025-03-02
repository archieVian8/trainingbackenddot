package db

import (
	"time"
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	DB *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{DB: db}
}

// Create New Schedule
func (r *ScheduleRepository) CreateSchedule(schedule *domain.Schedule) error {
	return r.DB.Create(schedule).Error
}

// View All Schedules
func (r *ScheduleRepository) GetAllSchedules() ([]domain.Schedule, error) {
	var schedules []domain.Schedule
	err := r.DB.
		Preload("Studio").
		Preload("Film").
		Find(&schedules).Error
	return schedules, err
}

// Update Schedule
func (r *ScheduleRepository) UpdateSchedule(id uint, updatedSchedule *domain.Schedule) error {
	return r.DB.Model(&domain.Schedule{}).Where("id = ?", id).Updates(updatedSchedule).Error
}

// Delete Schedule
func (r *ScheduleRepository) DeleteSchedule(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&domain.Schedule{}).Error
}

// Create Promo
func (r *ScheduleRepository) SetPromo(id uint, promo int, promoTime, promoEnds time.Time) error {
	var schedule domain.Schedule
	if err := r.DB.First(&schedule, id).Error; err != nil {
		return err
	}

	promoPrice := schedule.Price - (schedule.Price * float64(promo) / 100)

	return r.DB.Model(&schedule).Updates(map[string]interface{}{
		"Promo":      promo,
		"PromoPrice": promoPrice,
		"PromoTime":  promoTime,
		"PromoEnds":  promoEnds,
	}).Error
}
