package main

import (
	"fmt"
	"regexp"
	"strings"
)

func regex(body string) {
	//info := make([][]string, 2)
	fmt.Println("ok")

	/////////////////////////////// RECUPERANDO HASHs ///////////////////////////////
	hash_regex, _ := regexp.Compile(`data-jk="(.*?)"`)
	hash_results := (hash_regex.FindAllString(body, -1))

	preprocess1, _ := regexp.Compile(`data-jk="`)
	preprocess2, _ := regexp.Compile(`"`)

	for index, hash := range hash_results {
		hash := (preprocess1.ReplaceAllString(hash, ""))
		hash = (preprocess2.ReplaceAllString(hash, ""))
		hash_results[index] = hash
	}
	fmt.Println(hash_results)

	/////////////////////////////// RECUPERANDO TITLE ///////////////////////////////
	title_regex, _ := regexp.Compile(`<span\stitle=".*?"`)
	title_results := (title_regex.FindAllString(body, -1))

	preprocess3, _ := regexp.Compile(`<span\stitle="`)

	for index, title := range title_results {
		title := (preprocess3.ReplaceAllString(title, ""))
		title = (preprocess2.ReplaceAllString(title, ""))
		title_results[index] = title
	}
	//info = append(info, title_results)
	fmt.Println(title_results)

	/////////////////////////////// RECUPERANDO LOCAL ///////////////////////////////
	local_regex, _ := regexp.Compile(`class="companyLocation">.*?<`)
	local_results := (local_regex.FindAllString(body, -1))

	preprocess4, _ := regexp.Compile(`class="companyLocation">`)
	preprocess5, _ := regexp.Compile(`<`)

	for index, local := range local_results {
		local := (preprocess4.ReplaceAllString(local, ""))
		local = (preprocess5.ReplaceAllString(local, ""))
		local_results[index] = local
	}
	//info = append(info, local_results)
	fmt.Println(local_results)

	/////////////////////////////// RECUPERANDO EMPRESA ///////////////////////////////
	company_regex, _ := regexp.Compile(`<span class="companyName">.*</span>`)
	company_results := (company_regex.FindAllString(body, -1))

	preprocess6, _ := regexp.Compile(`<span class="companyName">`)
	preprocess7, _ := regexp.Compile(`</span>`)
	preprocess8, _ := regexp.Compile(`rel="noopener">`)
	preprocess9, _ := regexp.Compile(`</`)

	company_regex_href, _ := regexp.Compile(`rel="noopener">.*?</`)

	for index, company := range company_results {
		if strings.Contains(company, `rel="noopener"`) {
			company := company_regex_href.FindString(company)
			company = preprocess8.ReplaceAllString(company, "")
			company = preprocess9.ReplaceAllString(company, "")
			company_results[index] = company

		} else {
			company := (preprocess6.ReplaceAllString(company, ""))
			company = (preprocess7.ReplaceAllString(company, ""))
			company_results[index] = company
		}
	}
	fmt.Println(company_results)

	/////////////////////////////// Criando Map ///////////////////////////////
	//internships := make(map[string][]string)
	//n := len(hash_results)
	//for _, i := range info {
	//	if n != len(i) {
	//		panic("Error in regex")
	//	}
	//}
	//for index

}
