package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"searchRecommend/books/schema"
	service "searchRecommend/books/services"
	"searchRecommend/books/src/middlewares"
	"strconv"
	"strings"
	"time"
)

type RentBookHandler struct {
	Bookservice *service.RentBookService
}

// @Summary RentBook route
// @Description User can rent a book for a period of time
// @Tags books
// @Accept json
// @Produce json
// @Security bearerToken
// @Param id path int true "ID of the book to rent"
// @Param body body schema.RentBookDTO true "Request body in JSON format"
// @Success 200 {object} schema.RentBookSuccess
// @Failure 401 {object}  schema.Error
// @Router /rentBook/{id}  [post]
func (handler *RentBookHandler) RentbookHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	dynamicParam := parts[2]
	bookId, errr := strconv.Atoi(dynamicParam)
	if errr != nil {
		panic(errr.Error())
	}
	log.Println(bookId)

	userID, ok := r.Context().Value(middlewares.UserIDKey).(int64)

	if !ok {
		panic("UserID not found or has unexpected type")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	//here month is of type time.month which is a string type but we want int type

	//********CONVERTING MONTH STRING TO INT START*********
	var currMonth int
	switch month.String() {
	case "January":
		currMonth = 01
		log.Println(currMonth)
	case "February":
		currMonth = 02
		log.Println(currMonth)
	case "March":
		currMonth = 03
		log.Println(currMonth)
	case "April":
		currMonth = 04
		log.Println(currMonth)
	case "May":
		currMonth = 05
		log.Println(currMonth)
	case "June":
		currMonth = 06
		log.Println(currMonth)
	case "July":
		currMonth = 07
		log.Println(currMonth)
	case "August":
		currMonth = 8
		log.Println(currMonth)
	case "September":
		currMonth = 9
		log.Println(currMonth)
	case "October":
		currMonth = 10
		log.Println(currMonth)
	case "November":
		currMonth = 11
		log.Println(currMonth)
	case "December":
		currMonth = 12
		log.Println(currMonth)
	default:
		log.Println("Invalid month")

	}
	//*********CONVERTING MONTH STRING TO INT END**********

	//*********USING THE CONVERTED MONTH TO FIND LAST DAY OF CURRENT MONTH START************
	month = time.Month(currMonth)

	// Get the first day of the next month
	//explanation: nextmont varibale goes to the 1st of nextmont then in lastDayOfMonth we subtract 1 from the nexmont and it lands on last day of current month
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	// Subtract one day to get the last day of the current month
	lastDayOfMonth := nextMonth.AddDate(0, 0, -1).Day()

	//*********USING THE CONVERTED MONTH TO FIND LAST DAY OF CURRENT MONTH END************

	//*********CONVERTING USER'S RENTDURATION TO MODIFIED SCHEMA START***********
	var rentDuration schema.RentBookDTO
	if err1 := json.Unmarshal(body, &rentDuration); err1 != nil {
		panic(err1.Error())
	}
	//*********CONVERTING USER'S RENTDURATION TO MODIFIED SCHEMA END***********

	error := handler.Bookservice.RentbookService(rentDuration, lastDayOfMonth, currMonth, day, year, bookId, userID)
	if error.CODE != 0 || error.STATUSTEXT != "" || error.MESSAGE != "" {
		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(error.CODE)
		w.Write(json)

	} else {
		var success schema.RentBookSuccess
		message := fmt.Sprintf("Book successfully rented.Please take is bookId: [[%d]] with you at the time of returing the book submit the bookid along with it.", bookId)
		success.STATUS_CODE = 201
		success.STATUS_TEXT = http.StatusText(success.STATUS_CODE)
		success.MESSAGE = message

		json, err := json.Marshal(success)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(success.STATUS_CODE)
		w.Write(json)
	}

}

// @Summary Give Book back route
// @Description User can give the rented book back to the admin and admin can update the user rent details
// @Tags books
// @Accept json
// @Produce json
// @Security bearerToken
// @Param body body schema.GiveBookBackDTO true "Request body in JSON format"
// @Success 200 {object} schema.RentBookSuccess
// @Failure 401 {object}  schema.Error
// @Router /giveBookBack  [put]
func (handler *RentBookHandler) GiveBookBackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var req schema.GiveBookBackDTO
	if err := json.Unmarshal(body, &req); err != nil {
		panic(err.Error())
	}
	log.Println(req)

	error := handler.Bookservice.GiveBookBackService(req)
	if error.CODE != 0 || error.STATUSTEXT != "" || error.MESSAGE != "" {

		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(error.CODE)
		w.Write(json)

	} else {
		var success schema.RentBookSuccess
		success.STATUS_CODE = 201
		success.STATUS_TEXT = http.StatusText(success.STATUS_CODE)
		success.MESSAGE = "Updated User rent details successfully!"

		json, err := json.Marshal(success)
		if err != nil {
			panic(err.Error())
		}

		w.Write(json)
	}

}
