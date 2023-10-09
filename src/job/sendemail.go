package job

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"searchRecommend/schema"
	util "searchRecommend/utils"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type SendEmail struct {
	Db *util.Db
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	log.Println(authCode)

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (checkquery *SendEmail) CheckUsers() {

	//***********QUERY USERS WHO'S RENTDATE IS NEAR START***********
	db, err := checkquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	getUsers := `
	SELECT 
	users.userid,
	users.email,
	users.firstname,
	users.lastname,
	bookone.bookid,
	bookone.title,
	bookone.rented_date, 
	bookone.rentdate 
	FROM 
	    bookone 
	INNER JOIN 
	    users 
	ON 
	    bookone.userid = users.userid 
	AND 
	    bookone.rented = true;`

	query, err1 := db.Query(getUsers)
	if err1 != nil {
		panic(err1.Error())
	}

	var detailsArr []schema.RentDetails
	for query.Next() {
		var details schema.RentDetails
		var rented_date time.Time
		var rentdate time.Time
		data := query.Scan(
			&details.USER_ID,
			&details.EMAIL,
			&details.FIRST_NAME,
			&details.LAST_NAME,
			&details.BOOK_ID,
			&details.TITLE,
			&rented_date,
			&rentdate,
		)
		if data != nil {
			panic(data.Error())
		}
		details.RENTED_DATE = strings.Split(rented_date.Format("2006-01-02 15:04:05.999999"), " ")[0]

		var errr1 error
		var errr2 error
		var errr3 error
		details.RENTDATE.YEAR, errr1 = strconv.Atoi(strings.Split(rentdate.Format("2006-01-02"), "-")[0])
		if errr1 != nil {
			panic(errr1.Error())
		}
		details.RENTDATE.MONTH, errr2 = strconv.Atoi(strings.Split(rentdate.Format("2006-01-02"), "-")[1])
		if errr2 != nil {
			panic(errr2.Error())
		}
		details.RENTDATE.DAY, errr3 = strconv.Atoi(strings.Split(rentdate.Format("2006-01-02"), "-")[2])
		if errr3 != nil {
			panic(errr3.Error())
		}
		detailsArr = append(detailsArr, details)
	}
	log.Println(detailsArr)

	currentTime := time.Now()
	currYear, month, currDay := currentTime.Date()

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
	log.Println(currDay, currMonth, currYear)

	//***********QUERY USERS WHO'S RENTDATE IS NEAR END***********

	ctx := context.Background()
	b, err := os.ReadFile("/home/aamir/Desktop/My_code/searchRecommend/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailComposeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	// Get the first day of the next month
	//explanation: nextmont varibale goes to the 1st of nextmont then in lastDayOfMonth we subtract 1 from the nexmont and it lands on last day of current month
	nextMonth := time.Date(currYear, month+1, 1, 0, 0, 0, 0, time.UTC)
	// Subtract one day to get the last day of the current month
	lastDayOfMonth := nextMonth.AddDate(0, 0, -1).Day()

	for _, v := range detailsArr {

		if currMonth == v.RENTDATE.MONTH {
			if v.RENTDATE.DAY > currDay {
				if v.RENTDATE.DAY-currDay == 2 {
					log.Println("send Email to the users" + " " + v.EMAIL)

					rentDay := fmt.Sprintf("%d-%d-%d", v.RENTDATE.DAY, v.RENTDATE.MONTH, v.RENTDATE.YEAR)
					var submissionDay int
					var submissionMonth int
					var submissionYear int
					if v.RENTDATE.DAY == lastDayOfMonth {
						submissionDay = 1
						submissionMonth = v.RENTDATE.MONTH + 1

					} else if v.RENTDATE.MONTH == 12 && v.RENTDATE.DAY == lastDayOfMonth {
						submissionDay = 1
						submissionMonth = 1
						submissionYear = currYear + 1

					} else {
						submissionDay = v.RENTDATE.DAY + 1
						submissionMonth = v.RENTDATE.MONTH
						submissionYear = v.RENTDATE.YEAR
					}
					submissionDate := fmt.Sprintf("%d-%d-%d", submissionDay, submissionMonth, submissionYear)

					user := "me"

					subject := "BookApi rent remainder"
					body := fmt.Sprintf(
						"Hi %s %s Hope you are doing well and enjoying the rented books.\n This is a gentel remainder that the book '%s' that you rented on '%s' , is going to expire on '%s' 12:00 AM.\n Please submit the book on '%s' by 8:00 AM to the library authorities.\n\n P.S. This is a autogenerated Email do not send emails to this email id. Incase if you want to contact us send us email at  rubyparween13@gmail.com.",
						v.FIRST_NAME,
						v.LAST_NAME,
						v.TITLE,
						v.RENTED_DATE,
						rentDay,
						submissionDate,
					)
					message := fmt.Sprintf("From: %s\r\n", user)
					message += fmt.Sprintf("To: %s\r\n", v.EMAIL)
					message += fmt.Sprintf("Subject: %s\r\n", subject)
					message += "\r\n" + body

					rawMessage := base64.URLEncoding.EncodeToString([]byte(message))
					sendMessage := &gmail.Message{
						Raw: rawMessage,
					}

					_, err := srv.Users.Messages.Send(user, sendMessage).Do()
					if err != nil {
						panic(err.Error())
					}
				}

			}

		}

	}

}
