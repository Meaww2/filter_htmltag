package god

import (
	"log"
	"strings"
)

// How to optimize Skiping

func Filter_tag(html_ch chan HTMLcontent, content_ch chan DBobj) {
	for {
		raw_data := <-html_ch
		var record DBobj
		record.IsDone = false
		if raw_data.IsDone {
			record.IsDone = true
			content_ch <- record
			return
		}

		site, temp_data := raw_data.Site, raw_data.Content

		if temp_data == "" {
			log.Printf("Skip %s cause not have data", site)
			continue
		}

		log.Println("Processing:", site)
		// replace " ", "\n" to ""
		temp_data = strings.Replace(temp_data, "\n", "", -1)
		temp_data = strings.Replace(temp_data, " ", "", -1)
		temp_data = strings.ToLower(temp_data)
		// filter <body>
		start_body := strings.Index(temp_data, "<body")

		if start_body < 0 {
			continue
		}

		end_body := strings.Index(temp_data[start_body:], "</body>") + start_body
		body := temp_data[start_body : end_body+7]

		// get content
		// find < and > to skip string
		// find > and < to add string to list
		var lst_content []string
		i := 0
		temp := ""
		for i < len(body) {
			// skip comment
			if len(body) >= 4 && body[0:4] == "<!--" {
				i = strings.Index(body, "-->")
				body = body[i+3:]
				continue
			}
			// skip script
			if len(body) >= 8 && body[0:8] == "<script>" {
				i = strings.Index(body, "</script>")
				body = body[i+8:]
				continue
			}
			// skip style
			if len(body) >= 6 && body[0:6] == "<style" {
				i = strings.Index(body, "</style")
				body = body[i+6:]
				i = strings.Index(body, ">")
				body = body[i:]
				continue
			}
			// skip reference ex. [1]
			if len(body) >= 10 && body[0:10] == "<spanclass" {
				i = strings.Index(body[10:], "<spanclass")
				body = body[i+20:]
				i = strings.Index(body, "</span>")
				body = body[i+6:]
				continue
			}

			// filter
			if body[0] == '<' {
				i = strings.Index(body, ">")
			} else if body[0] == '>' {
				i = strings.Index(body, "<")
				temp = body[1:i]
				temp = strings.Trim(temp, "\n")
				temp = strings.TrimSpace(temp)
			} else {
				i++
			}

			// skip blank string

			if (temp != "" && len(lst_content) == 0) || (temp != "" && temp != lst_content[len(lst_content)-1]) {
				lst_content = append(lst_content, temp)
			}
			body = body[i:]
		}
		// pack data by get content between ">" and "<" than split by "\n"
		content := strings.Join(lst_content, "\n")
		record.site = site
		record.content = content
		content_ch <- record
		log.Printf("Filter %s success!", site)

	}

}
