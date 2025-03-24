package entities

type EventData struct {
    Action      string `json:"action"`
    PullRequest *struct {
        Title string `json:"title,omitempty"`
        State string `json:"state,omitempty"`
        URL   string `json:"html_url,omitempty"`
    } `json:"pull_request,omitempty"`
    Workflow *struct {
        Status     string `json:"status,omitempty"`
        Conclusion string `json:"conclusion,omitempty"`
        URL        string `json:"html_url,omitempty"`
    } `json:"workflow_run,omitempty"`
}