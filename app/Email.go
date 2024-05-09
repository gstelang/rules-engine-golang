package app

type Email struct {
	Subject        string `val:"subject"`
	Body           string `val:"body"`
	From_email     string `val:"from_email"`
	Reply_to       string `val:"reply_to"`
	Domain         string `val:"domain"`
	Company_domain string `val:"company_domain"`
}

type ScoredEmail struct {
	Email
	Score float64
}
