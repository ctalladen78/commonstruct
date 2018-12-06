package edvmodel

import "time"

//EsSchedule holds the schedule in elasticsearch
type EsSchedule struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty" location:"body"`
	Code      *string `json:"code,omitempty" location:"body"`

	//EventCategory handles what type of event is this
	//It can be:
	//course + class schedule: COURSE_CLASS_SCHEDULE
	//class schedule: CLASS_SCHEDULE
	//course schedule: COURSE_SCHEDULE
	//field trip: FIELD_TRIP_SCHEDULE
	//billing schedule: BILLING_SCHEDULE
	//exam schedule: EXAM_SCHEDULE
	//public events: PUBLIC_SCHEDULE
	EventCategory *string `json:"event_category,omitempty" location:"body"`

	//Participants can contain 1 or more filters (who are this event intended for)
	//For course schedule, participants should contain 2 elements [course_code, class_code]
	//For single class event, participants should contain 1 element [class_code].
	//For a batch event, participants should contain 1 element [batch_code]
	//For public events, participants should contain 1 element [*]
	//use this field with combination with event_category to achieve unique participants
	Participants []string `json:"participants,omitempty" location:"body"`

	//POSSIBLE VALUES: DAILY, MONTHLY, WEEKLY, YEARLY, SINGLE (for one-time event)
	SchedulePeriod *string `json:"schedule_period,omitempty" location:"body"`

	//recurrenceValues is....
	//If daily: an array of time
	//If weekly: an array of dayOfWeek + time (e.g. 1/13:00:00 for monday, 1 PM) [0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday]
	//If Monthly: an array of dayOfMonth + time (e.g. 26/13:00:00 for every 26th, 1PM)
	//If yearly: an array of dayOfMonth + month + time (e.g. 26-0/13:00:00 for every 26th January, 1PM) [MONTH STARTS FROM 0(jan) AND ENDS AT 11(dec)]
	RecurrenceValues []string   `json:"recurrence_values,omitempty" location:"body"`
	StartTime        *time.Time `json:"start_time,omitempty" location:"body"`
	EndTime          *time.Time `json:"end_time,omitempty" location:"body"`

	//ScheduleTargets only for this platform
	//Possible values are student & parent
	//Schedule for admins (e.g. teachers) need to be in iteacloud
	//However, teacher can view student schedules
	ScheduleTargets []string `json:"schedule_targets,omitempty" location:"body"`

	Description *string `json:"description" location:"body"`
	Title       *string `json:"title,omitempty" location:"body"`

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
