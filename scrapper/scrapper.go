package scrapper

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

type extractedJob struct {
	id        string
	company   string
	title     string
	date      string
	condition string
	link      string
}

var baseURL string

func Scrape(term string) {
	var allJobs []extractedJob

	var url = "https://www.saramin.co.kr/zf_user/search/recruit?search_area=main&search_done=y&search_optional_item=n&searchType=default_mysearch&searchword=" + term + "&loc_mcd=101000&cat_cd=404"

	baseURL = url
	ch := make(chan []extractedJob)

	totalPages := getPages()

	for i := 1; i <= totalPages; i++ {
		go getPage(i, ch)
	}

	for i := 1; i <= totalPages; i++ {
		jobs := <-ch
		allJobs = append(allJobs, jobs...)
	}

	writeAllJobs(allJobs)
}

func writeAllJobs(allJobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Company", "Title", "date", "Condition", "link"}

	wErr := w.Write(headers)

	checkErr(wErr)

	for _, job := range allJobs {

		// package install : go get -u golang.org/x/text/...
		// convert utf-8 to euc-kr
		got1, _, _ := transform.String(korean.EUCKR.NewEncoder(), job.id)
		got2, _, _ := transform.String(korean.EUCKR.NewEncoder(), job.company)
		got3, _, _ := transform.String(korean.EUCKR.NewEncoder(), job.title)
		got4, _, _ := transform.String(korean.EUCKR.NewEncoder(), job.date)
		got5, _, _ := transform.String(korean.EUCKR.NewEncoder(), job.condition)
		got6, _, _ := transform.String(korean.EUCKR.NewEncoder(), job.link)

		jobSlice := []string{got1, got2, got3, got4, got5, got6}
		// jobSlice := []string{job.id, job.company, job.title, job.date, job.condition, job.link}

		jwErr := w.Write(jobSlice)
		checkErr(jwErr)

	}

}

func getPage(page int, ch chan []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	cardList := doc.Find(".content").Find(".item_recruit")

	cardList.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < cardList.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	ch <- jobs
}

func extractJob(card *goquery.Selection, c chan extractedJob) {
	jobId, _ := card.Attr("value")
	corpName, _ := card.Find(".corp_name a").Attr("title")
	jobTitle := CleanString(card.Find(".job_tit").Text())
	jobDate := CleanString(card.Find(".job_date").Text())
	jobCondition := CleanString(card.Find(".job_condition").Text())
	jobLink, _ := card.Find(".job_tit a").Attr("href")
	jobLink = "https://www.saramin.co.kr" + jobLink

	c <- extractedJob{
		id:        jobId,
		company:   corpName,
		title:     jobTitle,
		date:      jobDate,
		condition: jobCondition,
		link:      jobLink,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	p := doc.Find(".pagination")
	pages = p.Find("a").Length()

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status :", res.Status)
	}
}

func CleanString(text string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(text)), " ")
}
