/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking_detail.go (c) 2023
Desc: description
Created:  2023-05-22T08:51:37.100Z
Modified: !date!
*/

package entity

import "time"

type BookingDetail struct {
	Id        int       `json:"id"`
	BookingId int       `json:"booking_id"`
	StartTime time.Time `json:"start_time"`
	DateBook  time.Time `json:"date_book"`
	EndTime   time.Time `json:"end_time"`
}
