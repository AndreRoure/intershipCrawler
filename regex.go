package main

import (
	"fmt"
	"regexp"
)

func next(body string) (link string) {
	//fmt.Println(body)
	next_regex, _ := regexp.Compile(`<a\shref='\/jobs\?q.*aria-label='Pr`)
	next_results := next_regex.FindAllString(body, -1)
	if len(next_results) == 0 {
		return ""
	}
	v := next_results[0][len(next_results[0])-80:]
	regex2, _ := regexp.Compile(`href='.*?'`)
	v = regex2.FindString(v)
	preprocess10, _ := regexp.Compile(`href='`)
	preprocess11, _ := regexp.Compile(`'`)
	v = preprocess10.ReplaceAllString(v, "")
	v = preprocess11.ReplaceAllString(v, "")
	v = "https://br.indeed.com" + v

	return v
}

func regex(body string, c chan map[string]info) {
	info_list := make([][]string, 0)
	//fmt.Println("ok")

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
	//fmt.Println(hash_results)

	/////////////////////////////// RECUPERANDO TITLE ///////////////////////////////
	title_regex, _ := regexp.Compile(`<span\stitle=".*?"`)
	title_results := (title_regex.FindAllString(body, -1))

	preprocess3, _ := regexp.Compile(`<span\stitle="`)

	for index, title := range title_results {
		title := (preprocess3.ReplaceAllString(title, ""))
		title = (preprocess2.ReplaceAllString(title, ""))
		title_results[index] = title
	}
	info_list = append(info_list, title_results)
	//fmt.Println(title_results)

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
	info_list = append(info_list, local_results)
	//fmt.Println(local_results)

	/////////////////////////////// RECUPERANDO EMPRESA ///////////////////////////////
	company_regex, _ := regexp.Compile(`<div class="heading6 company_location tapItem-gutter companyInfo">.*?<\/span>`)
    company_results := (company_regex.FindAllString(body, -1))
    // fmt.Println("------")
    // fmt.Println(company_results)
    // fmt.Println("------")
    preprocess6, _ := regexp.Compile(`.*">`)
    preprocess7, _ := regexp.Compile(`</span>`)
	preprocess8, _ := regexp.Compile(`</a>`)
    
    for index, company := range company_results {
        company := (preprocess6.ReplaceAllString(company, ""))
        company = (preprocess7.ReplaceAllString(company, ""))
		company = (preprocess8.ReplaceAllString(company, ""))
        company_results[index] = company
		fmt.Println(company)
    }
	info_list = append(info_list, company_results)



	/////////////////////////////// Criando Map ///////////////////////////////
	internships := make(map[string]info)
	n := len(hash_results)
	for _, i := range info_list {
		if n != len(i) {
			panic("Error in regex")
		}
	}
	for index, hash := range hash_results {
		i := info{info_list[0][index], info_list[1][index]}
		internships[hash] = i
	}
	c <- internships

}
