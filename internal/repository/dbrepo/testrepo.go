package dbrepo

import (
	"errors"
	"time"

	"github.com/mikeruu/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

//InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

//InserRoomResrtiction inserts a room restriction into database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

//SearchAvailabilityByDatesByRoomID returns true if the room is avail and false if it has a restriction
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil
}

//SearchAvailabilityForAllRooms returns a slice of avail rooms if any for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

//GEtROOByID gets a room by id from database
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some Error")
	}
	
	return room, nil
}
