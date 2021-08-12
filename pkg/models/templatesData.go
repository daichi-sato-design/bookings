package models

// TemplateData は、ハンドラーからテンプレートに送信されるデータを保持する
type TemplateData struct{
	StringMap	map[string]string
	IntMap		map[string]int
	FloatMap	map[string]float32
	Data		map[string]interface{}
	CSRFToken	string
	Flash		string
	Warning		string
	Error		string
}