package usecase

import (
	"time"
	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type ScheduleUsecase struct {
	ScheduleRepo *db.ScheduleRepository
}

func NewScheduleUsecase(scheduleRepo *db.ScheduleRepository) *ScheduleUsecase {
	return &ScheduleUsecase{ScheduleRepo: scheduleRepo}
}

// Function for New Schedule
func (u *ScheduleUsecase) CreateSchedule(schedule *domain.Schedule) error {
	return u.ScheduleRepo.CreateSchedule(schedule)
}

// Function for View All Schedules
func (u *ScheduleUsecase) ViewAllSchedules() ([]domain.Schedule, error) {
	return u.ScheduleRepo.GetAllSchedules()
}

// Function for Update Schedule
func (u *ScheduleUsecase) UpdateSchedule(id uint, updatedSchedule *domain.Schedule) error {
	return u.ScheduleRepo.UpdateSchedule(id, updatedSchedule)
}

// Function for Delete Schedule
func (u *ScheduleUsecase) DeleteSchedule(id uint) error {
	return u.ScheduleRepo.DeleteSchedule(id)
}

// Function for Create Promo
func (u *ScheduleUsecase) ApplyPromo(id uint, promo int, promoTime, promoEnds time.Time) error {
	return u.ScheduleRepo.SetPromo(id, promo, promoTime, promoEnds)
}
