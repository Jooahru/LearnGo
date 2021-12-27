package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

// Scrape Indeed
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	c := make(chan []extractedJob)
	var jobs []extractedJob
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...) //... 배열합치기
	}
	writeJobs(jobs)
	fmt.Println("Jobs Extract Done")
}

func writeJobs(jobs []extractedJob) {
	c := make(chan []string)
	file, err := os.Create("jobs.csv") // 파일생성
	checkErr(err)
	utf8bom := []byte{0xEF, 0xBB, 0xBF} // 한글깨짐방지
	file.Write(utf8bom)                 // 한글깨짐방지

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		go goRoutineWrite(job, c)
	}

	for i := 0; i < len(jobs); i++ {
		job := <-c
		jwErr := w.Write(job)
		checkErr(jwErr)

	}
}

func goRoutineWrite(job extractedJob, c chan<- []string) {
	jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
	c <- jobSlice
}

func getPage(page int, baseURL string, cMain chan<- []extractedJob) {
	c := make(chan extractedJob)
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50) // strconv go Packages
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	cMain <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".singleLineTitle>h2").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	salary := CleanString(card.Find(".salary-snippet").Text())
	summary := CleanString(card.Find(".job-snippet").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func getPages(baseURL string) int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Requestfailed with Status:", res.StatusCode)
	}
}

// CleanSTring
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "")
}
