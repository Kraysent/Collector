package ticktick

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID              string   `json:"id"`
	ProjectID       string   `json:"projectId"`
	SortOrder       int64    `json:"sortOrder"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	Desc            string   `json:"desc"`
	StartDate       string   `json:"startDate"`
	DueDate         string   `json:"dueDate"`
	TimeZone        string   `json:"timeZone"`
	IsFloating      bool     `json:"isFloating"`
	IsAllDay        bool     `json:"isAllDay"`
	Reminders       []any    `json:"reminders"`
	ExDate          []any    `json:"exDate"`
	CompletedTime   string   `json:"completedTime"`
	CompletedUserID int      `json:"completedUserId"`
	RepeatTaskID    string   `json:"repeatTaskId"`
	Priority        int      `json:"priority"`
	Status          int      `json:"status"`
	Items           []any    `json:"items"`
	Progress        int      `json:"progress"`
	ModifiedTime    string   `json:"modifiedTime"`
	Etag            string   `json:"etag"`
	Deleted         int      `json:"deleted"`
	CreatedTime     string   `json:"createdTime"`
	Creator         int      `json:"creator"`
	RepeatFrom      string   `json:"repeatFrom"`
	Tags            []string `json:"tags"`
	Attachments     []any    `json:"attachments"`
	CommentCount    int      `json:"commentCount"`
	FocusSummaries  []any    `json:"focusSummaries"`
	ColumnID        string   `json:"columnId"`
	Kind            string   `json:"kind"`
}

type UpdateTaskRequest struct {
	Update []Task `json:"update"`
}
