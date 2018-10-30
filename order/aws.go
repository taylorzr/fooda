package main

type AwsEmail struct {
	// Receipt Receipt `json:"receipt"`
	Mail    Mail   `json:"mail"`
	Content string `json:"content"`
}

type Mail struct {
	Source        string        `json:"source"`
	Headers       []Header      `json:"headers"`
	CommonHeaders CommonHeaders `json:"commonHeaders"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CommonHeaders struct {
	From        []string `json:"from"`
	To          []string `json:"to"`
	DeliveredTo []string `json:"Delivered-To"`
	Subject     string   `json:"subject"`
}
