package god

import (
	"log"
	"strings"
)

// How to optimize Skiping

func Filter_tag(html_ch chan HTMLcontent, content_ch chan DBobj, monitor_ch chan int) {
	for {
		count := <-monitor_ch
		if count >= 1000 {
			count++
			monitor_ch <- count
			break
		}
		monitor_ch <- count
		raw_data := <-html_ch
		site, temp_data := raw_data.Site, raw_data.Content

		// replace " ", "\n" to ""
		temp_data = strings.Replace(temp_data, "\n", "", -1)
		temp_data = strings.Replace(temp_data, " ", "", -1)
		// filter <body>
		start_body := strings.Index(temp_data, "<body")
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
			if body[0:4] == "<!--" {
				i = strings.Index(body, "-->")
				body = body[i+3:]
				continue
			}
			// skip script
			if body[0:8] == "<script>" {
				i = strings.Index(body, "</script>")
				body = body[i+8:]
				continue
			}
			// skip style
			if body[0:6] == "<style" {
				i = strings.Index(body, "</style")
				body = body[i+6:]
				i = strings.Index(body, ">")
				body = body[i:]
				continue
			}
			// skip reference ex. [1]
			if body[0:10] == "<spanclass" {
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
			if temp != "" {
				lst_content = append(lst_content, temp)
			}
			body = body[i:]
		}

		// pack data by get content between ">" and "<" than split by "\n"
		content := strings.Join(lst_content, "\n")
		var record DBobj
		record.site = site
		record.content = content
		log.Println("Filter success:", content)
		content_ch <- record

	}

}
