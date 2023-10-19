package service

import (
	"log"
	repository "searchRecommend/books/repositories"
	"searchRecommend/books/schema"
)

type RentBookService struct {
	RentBookRepo *repository.RentBookRepo
}

func (service *RentBookService) RentbookService(body schema.RentBookDTO, currMonthEnd, currMonth, currDay, currYear, bookId int) schema.Error {

	rentedDay := currDay
	rentedMonth := currMonth
	rentedYear := currYear

	count := 0
	var end_Day int
	if body.RENTDURATION.MONTHS == 0 {
		for i := 0; i < body.RENTDURATION.DAYS; i++ {
			if currDay != currMonthEnd {
				currDay++
				count++
				end_Day = currDay
			} else {
				if currMonth != 12 {
					end_Day = body.RENTDURATION.DAYS - count
					currMonth++
					break
				} else {
					end_Day = body.RENTDURATION.DAYS - count
					currMonth++
					currYear++
				}

			}

		}
	} else if body.RENTDURATION.DAYS == 0 {
		for i := 0; i < body.RENTDURATION.MONTHS; i++ {
			end_Day = currDay
			currMonth++

		}
	} else {
		log.Println(body)
		for i := 0; i < body.RENTDURATION.MONTHS; i++ {
			end_Day = currDay
			currMonth++

		}
		for i := 0; i < body.RENTDURATION.DAYS; i++ {
			if currDay != currMonthEnd {
				currDay++
				count++
				end_Day = currDay
			} else {
				if currMonth != 12 {
					end_Day = body.RENTDURATION.DAYS - count
					currMonth++
					break
				} else {
					end_Day = body.RENTDURATION.DAYS - count
					currMonth++
					currYear++
				}

			}

		}
	}
	var rent schema.RentBookSchema
	rent.USER_ID = body.USER_ID
	rent.RENTED_DATE.DAY = rentedDay
	rent.RENTED_DATE.MONTH = rentedMonth
	rent.RENTED_DATE.YEAR = rentedYear
	rent.RENT_DAY.DAYS = end_Day
	rent.RENT_DAY.MONTHS = currMonth
	rent.RENT_DAY.YEARS = currYear

	error := service.RentBookRepo.RentbookRepo(rent, bookId)
	// Log := fmt.Sprintf("%d-%d-%d", end_Day, currMonth, currYear)
	log.Println(rent)
	return error

}

func (service *RentBookService) GiveBookBackService(req schema.GiveBookBackDTO) schema.Error {

	error := service.RentBookRepo.GiveBookBackRepo(req)
	return error
}
