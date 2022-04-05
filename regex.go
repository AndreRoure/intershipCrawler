package main

import (
    "fmt"
    "regexp"
	"strings"
)

func regex() {

	/////////////////////////////// RECUPERANDO HASHs ///////////////////////////////
	hash_regex, _ := regexp.Compile(`data-jk="(.*?)"`)
	hash_results := (hash_regex.FindAllString(`data-mobtk="1fvtngi77pv1h804" data-jk="12824b18a96400bc" data-hiring-event="false" target="_blank" data-hide-spinner="true" rel="nofollow" class="tapItem fs-unmask result job_12824b18a96400bc resultWithShelf sponTapItem desktop vjs-highlight" href="/rc/clk?jk=12824b18a96400bc&from=mobhp_jobfeed&tk=1fvtngi77pv1h804"> <a id="job_57809beabc77b6f7" data-mobtk="1fvtngi77pv1h804" data-jk="57809beabc77b6f7" data-hiring-event="false" target="_blank" data-hide-spinner="true" rel="nofollow" class="tapItem fs-unmask result job_57809beabc77b6f7 resultWithShelf sponTapItem desktop" href="/rc/clk?jk=57809beabc77b6f7&from=mobhp_jobfeed&tk=1fvtngi77pv1h804">…</a> <a id="job_1e5a582ab71d94f5" data-mobtk="1fvtngi77pv1h804" data-jk="1e5a582ab71d94f5" data-hiring-event="false" target="_blank" data-hide-spinner="true" rel="nofollow" class="tapItem fs-unmask result job_1e5a582ab71d94f5 resultWithShelf sponTapItem desktop" h`, -1)) 
	
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
	title_results := (title_regex.FindAllString(`<span title="Estagiário Superior">Estagiário Superior</span>
	<span title="Estagiária em Pedagogia">Estagiária em Pedagogia</span>
	<span title="Estagiário na área de Enfermagem - Taguatinga">Estagiário na área de Enfermagem - Taguatinga</span>`, -1)) 
	
	preprocess3, _ := regexp.Compile(`<span\stitle="`)
 	
	for index, title := range title_results {
		title := (preprocess3.ReplaceAllString(title, ""))
		title = (preprocess2.ReplaceAllString(title, ""))
		title_results[index] = title
 	}
	fmt.Println(title_results)

	/////////////////////////////// RECUPERANDO LOCAL ///////////////////////////////
	local_regex, _ := regexp.Compile(`class="companyLocation">.*?<`)
	local_results := (local_regex.FindAllString(`<div class="companyLocation">Home office in Brasília, DF</div>flex
	<div class="companyLocation">Brasília, DF</div>flex
	<div class="companyLocation">Taguatinga, DF</div>flex
	`, -1)) 
	
	preprocess4, _ := regexp.Compile(`class="companyLocation">`)
	preprocess5, _ := regexp.Compile(`<`)
 	
	for index, local := range local_results {
		local := (preprocess4.ReplaceAllString(local, ""))
		local = (preprocess5.ReplaceAllString(local, ""))
		local_results[index] = local
 	}
	fmt.Println(local_results)

	/////////////////////////////// RECUPERANDO EMPRESA ///////////////////////////////
	company_regex, _ := regexp.Compile(`<span class="companyName">.*</span>`)
	company_results := (company_regex.FindAllString(`
	<span class="companyName"><a data-tn-element="companyName" class="turnstileLink companyOverviewLink" target="_blank" href="/cmp/Captamed Cuidados-Continuados" rel="noopener">Captamed Cuidados Continuados</a></span>
	<span class="companyName">Jobbol</span>
	`, -1)) 
	
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

		}else{
			company := (preprocess6.ReplaceAllString(company, ""))
			company = (preprocess7.ReplaceAllString(company, ""))
			company_results[index] = company
		}
	}
	fmt.Println(company_results)








}