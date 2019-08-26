package main

/**

go run main.go -DsrcList=/Users/i311688/gitrepo/go-repo/manhua-downloader/list.missing.txt -DbaseFolder=/Users/i311688/entertainment/manga/hunter/
go run main.go -DsrcUrl=https://manhua.fzdm.com/10/ -DbaseFolder=/Users/i311688/entertainment/manga/hunter/
**/
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"sync"

	// "io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

const ARG_BASEFOLER = "baseFolder"
const ARG_SRCURL = "srcUrl"
const ARG_SRCLIST = "srcList"

var picUrlKey = "var mhurl="
var picHost = "http://p1.manhuapan.com/"

// var baseFolder = "C:/Users/i311688/Desktop/MyTemp/manga/hzw/"
var baseFolder = "/Users/i311688/entertainment/manga/one_piece/"
var httpClient *http.Client

func init() {
	proxyStr := "http://proxy.pvgl.sap.corp:8080"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	cookieJar, _ := cookiejar.New(nil)
	httpClient = &http.Client{
		Transport: transport,
		Jar:       cookieJar,
	}
}

func main() {
	// checkFailedUrl()

	clean := setLogger()
	defer clean()

	args := parseArgs()

	if _, ok := args[ARG_SRCURL]; ok {
		downloadWithSrcUrl(args)
	} else if _, ok := args[ARG_SRCLIST]; ok {
		downloadWithList(args)
	} else {
		fmt.Println("source url or source list file path is required")
	}
}

func checkFailedUrl() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("can't get the log file")
		os.Exit(1)
	}
	logPath := dir + "/main.log"

	file, err := os.Open(logPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	flag := "check page error: Get"
	flagLen := len(flag)
	set := make(map[string]bool)
	for scanner.Scan() {
		url := scanner.Text()
		idx := strings.Index(url, flag)
		if idx >= 0 {
			idx2 := strings.LastIndex(url, ":")
			url = url[idx+flagLen : idx2]
			url = strings.ReplaceAll(url, ": read tcp 10.59.188.100:65222->172.20.5.58:8080: read", "")
			url = strings.ReplaceAll(url, ": read tcp 10.59.188.100:49547->172.20.5.58:8080: read", "")
			url = strings.Join(strings.Split(url, "/")[:5], "/") + "/"
			set[url] = true
		}
	}
	file.Close()
	for key, _ := range set {
		fmt.Println(key)
	}

	fmt.Println("done")
	os.Exit(0)
}

func setLogger() func() {
	dir, err := os.Getwd()
	if err != nil {
		dir = os.TempDir()
	}
	logPath := dir + "/main.log"
	fmt.Printf("log file path: %s\n", logPath)
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		os.Stdout = logFile
		return func() {
			logFile.Close()
		}
	} else {
		fmt.Printf("%v\n", err)
		return func() {
			// do nothing
		}
	}
}

func downloadWithSrcUrl(args map[string]string) {
	if bFolder, ok := args[ARG_BASEFOLER]; ok {
		baseFolder = bFolder
	}
	url := args[ARG_SRCURL]
	fmt.Printf("source URL is: %s \nbase folder is: %s\n", url, baseFolder)
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Printf("source url %s not avaiable: %v\n", url, err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("source url %s not avaiable: %v\n", url, err)
		resp.Body.Close()
		return
	}
	var urls []string
	doc.Find("#content li").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		urls = append(urls, fmt.Sprintf("%s%s", url, href))
	})
	downloadWithUrls(urls)
}

func downloadWithList(args map[string]string) {
	path := args[ARG_SRCLIST]
	fmt.Printf("file path is: %s\n", path)
	if p, ok := args[ARG_BASEFOLER]; ok {
		baseFolder = p
	}
	fmt.Printf("destination folder is: %s\n", baseFolder)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		url := scanner.Text()
		url = strings.Trim(url, " ")
		if url != "" && !strings.HasPrefix(url, "#") {
			urls = append(urls, url)
		}
	}
	file.Close()

	downloadWithUrls(urls)
}

func downloadWithUrls(urls []string) {
	// urls = urls[:20]
	var wg sync.WaitGroup
	wg.Add(len(urls))
	maxThread := 5

	queue := make(chan string, len(urls))
	fmt.Printf("url amount: %d\n", len(urls))

	for idx, url := range urls {
		queue <- url
		if idx < maxThread {
			fmt.Printf("create download worker: %d\n", idx)
			go startWorker(queue, &wg)
		}
		// go download(url, &wg)
	}
	wg.Wait()
	close(queue)
}

func startWorker(queue <-chan string, wg *sync.WaitGroup) {
	for {
		url, ok := <-queue
		if ok {
			fmt.Printf("start %s\n", url)
			download(url, wg)
		} else {
			fmt.Println("url queue is closed")
			break
		}
	}
}

func download(url string, wg *sync.WaitGroup) {
	baseUrl := url
	defer func() {
		wg.Done()
		fmt.Printf("download %s done\n", baseUrl)
	}()
	initFolder := false
	folder := ""
	picIdx := 0
	title := ""
	var jobs [](func())
	picWg := sync.WaitGroup{}
	for {
		hasNext := false
		// fmt.Printf("download %s\n", url)
		resp, err := httpClient.Get(url)
		if err != nil {
			fmt.Printf("check page error: %v\n", err)
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Printf("download %s failed: %v\n", url, err)
			resp.Body.Close()
			return
		}

		if initFolder == false {
			doc.Find("title").Each(func(i int, s *goquery.Selection) {
				title = strings.Trim(s.Text(), " ")
				folder = baseFolder + title + "/"
				os.MkdirAll(folder, 0777)
				initFolder = true
			})
		}

		doc.Find("a.pure-button.pure-button-primary").Each(func(i int, s *goquery.Selection) {
			if s.Text() == "下一页" {
				href, _ := s.Attr("href")
				url = baseUrl + href
				// fmt.Printf("next is %s\n", url)
				hasNext = true
			}
		})
		// find the pic url
		keyLen := len(picUrlKey)
		doc.Find("body script").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			idx := strings.Index(text, picUrlKey)
			if idx >= 0 {
				text = text[idx:]
				idx2 := strings.Index(text, ";")
				// fmt.Printf("script is: %d, %d, %s\n", keyLen, idx2, text)
				picUrl := text[keyLen+1 : idx2-1]
				jobs = append(jobs, createJob(fmt.Sprintf("%s%s", picHost, picUrl), folder, picIdx, &picWg))
				// downloadPic(fmt.Sprintf("%s%s", picHost, picUrl), folder, picIdx)
			}
		})

		resp.Body.Close()

		picIdx++
		if hasNext == false {
			break
		}
	}
	picWg.Add(picIdx)
	for _, job := range jobs {
		go job()
	}
	picWg.Wait()
	// fmt.Printf("%s finished: %d\n", title, picIdx)
}

func createJob(url string, folder string, picIdx int, wg *sync.WaitGroup) func() {
	return func() {
		downloadPic(url, folder, picIdx)
		wg.Done()
	}
}

func downloadPic(picUrl string, folder string, picIdx int) {
	// fmt.Printf("download %s\n", picUrl)
	filePath := folder + fmt.Sprintf("%04d", picIdx) + ".jpg"
	if _, e := os.Stat(filePath); os.IsNotExist(e) == false {
		// fmt.Printf("%s is already downloaded\n", picUrl)
		return
	}

	resp, err := httpClient.Get(picUrl)
	if err != nil {
		fmt.Printf("download pic %s failed: %v\n", picUrl, err)
		return
	}
	defer resp.Body.Close()
	img, _ := os.Create(filePath)
	defer img.Close()
	_, err = io.Copy(img, resp.Body)
	if err == nil {
		// fmt.Println("done")
	} else {
		fmt.Println("error")
	}
}

func parseArgs() map[string]string {
	args := make(map[string]string)

	for _, arg := range os.Args {
		idx := strings.Index(arg, "-D")
		if idx == 0 {
			idx2 := strings.Index(arg, "=")
			if idx2 >= 0 && idx2 < (len(arg)-1) {
				args[arg[idx+2:idx2]] = arg[idx2+1:]
			}
		}
	}
	fmt.Printf("args: %v\n", args)
	return args
}
