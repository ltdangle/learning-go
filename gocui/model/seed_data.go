package model

import (
	"math/rand"
)

func SeedData() []EmailAccount {
	emails1 := []Email{
		{
			Id:          rand.Intn(100),
			Path:        "/1/email1",
			From:        "john@example.com",
			To:          "jane@example.com",
			DeliveredTo: "jane@example.com",
			Subject:     "Meeting Tomorrow",
			Text:        "Hi Jane, I just wanted to remind you about our meeting tomorrow at 2pm.",
			HTML:        "",
			Date:        "2022-03-02T10:00:00Z",
			IsSeen:      false,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/1/email2",
			From:        "jane@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "RE: Meeting Tomorrow",
			Text:        "Thanks for the reminder, John. I'll be there at 2pm tomorrow.",
			HTML:        "",
			Date:        "2022-03-02T11:00:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  true,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/1/email3",
			From:        "bob@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Project Update",
			Text:        "Hi John, just wanted to give you a quick update on the project. We're making good progress and should be able to hit our deadline.",
			HTML:        "",
			Date:        "2022-03-01T15:30:00Z",
			IsSeen:      true,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/1/email4",
			From:        "alice@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Important Announcement",
			Text:        "Hello, we have an important company-wide announcement. Please see the attached document for details.",
			HTML:        "",
			Date:        "2022-02-28T09:15:00Z",
			IsSeen:      true,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/1/email5",
			From:        "sales@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Special Offer",
			Text:        "Don't miss out on our special offer! Use code SPECIAL10 for 10% off your next purchase.",
			HTML:        "",
			Date:        "2022-02-25T14:00:00Z",
			IsSeen:      true,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		},
	}
	emails2 := []Email{
		{
			Id:          rand.Intn(100),
			Path:        "/2/email1",
			From:        "mike@example.com",
			To:          "jane@example.com",
			DeliveredTo: "jane@example.com",
			Subject:     "New Product Launch",
			Text:        "Hi Jane, I wanted to let you know that we just launched our new product line. Check it out on our website!",
			HTML:        "",
			Date:        "2023-03-10T13:00:00Z",
			IsSeen:      false,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/2/email2",
			From:        "jane@example.com",
			To:          "mike@example.com",
			DeliveredTo: "mike@example.com",
			Subject:     "RE: New Product Launch",
			Text:        "Thanks for letting me know, Mike. I'll check it out!",
			HTML:        "",
			Date:        "2023-03-11T09:30:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  true,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/2/email3",
			From:        "bob@example.com",
			To:          "mike@example.com",
			DeliveredTo: "mike@example.com",
			Subject:     "Meeting Request",
			Text:        "Hi Mike, I'd like to schedule a meeting with you next week to discuss our project. Can you let me know your availability?",
			HTML:        "",
			Date:        "2023-03-08T10:15:00Z",
			IsSeen:      true,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/2/email4",
			From:        "alice@example.com",
			To:          "mike@example.com",
			DeliveredTo: "mike@example.com",
			Subject:     "Reminder: Team Meeting Today",
			Text:        "Hello, this is a reminder that we have our team meeting today at 3pm in the conference room.",
			HTML:        "",
			Date:        "2023-03-07T11:30:00Z",
			IsSeen:      true,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
	}
	emails3 := []Email{
		{
			Id:          rand.Intn(100),
			Path:        "/2/email1",
			From:        "susan@example.com",
			To:          "jane@example.com",
			DeliveredTo: "jane@example.com",
			Subject:     "Volunteer Opportunity",
			Text:        "Hi Jane, I hope this email finds you well. I wanted to let you know about a volunteer opportunity happening next week. Let me know if you're interested!",
			HTML:        "",
			Date:        "2023-03-10T09:00:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/2/email2",
			From:        "mark@example.com",
			To:          "jane@example.com",
			DeliveredTo: "jane@example.com",
			Subject:     "New Product Launch",
			Text:        "Hey Jane, just wanted to give you a heads up that we're launching a new product next week. I think you'll be really excited about it!",
			HTML:        "",
			Date:        "2023-03-09T14:30:00Z",
			IsSeen:      true,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/2/email3",
			From:        "jim@example.com",
			To:          "susan@example.com",
			DeliveredTo: "susan@example.com",
			Subject:     "Conference Call",
			Text:        "Hi Susan, can we schedule a conference call for next week? I have a few things I'd like to discuss with you.",
			HTML:        "",
			Date:        "2023-03-08T11:45:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/2/email4",
			From:        "jane@example.com",
			To:          "mark@example.com",
			DeliveredTo: "mark@example.com",
			Subject:     "RE: New Product Launch",
			Text:        "Thanks for letting me know, Mark! Can't wait to see what you guys have been working on.",
			HTML:        "",
			Date:        "2023-03-09T16:15:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  true,
			IsSelected:  false,
		},
	}
	emails4 := []Email{
		{
			Id:          rand.Intn(100),
			Path:        "/4/email1",
			From:        "marketing@example.com",
			To:          "jane@example.com",
			DeliveredTo: "jane@example.com",
			Subject:     "New Product Launch",
			Text:        "Hi Jane, we're excited to announce the launch of our new product! Check it out at our website.",
			HTML:        "",
			Date:        "2023-03-10T09:30:00Z",
			IsSeen:      false,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/4/email2",
			From:        "jim@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Meeting Rescheduled",
			Text:        "Hi John, I need to reschedule our meeting to next Thursday at 2pm. Can you make it then?",
			HTML:        "",
			Date:        "2023-03-11T11:00:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		}}
	emails5 := []Email{
		{
			Id:          rand.Intn(100),
			Path:        "/5/email1",
			From:        "john@example.com",
			To:          "jane@example.com",
			DeliveredTo: "jane@example.com",
			Subject:     "Meeting Tomorrow",
			Text:        "Hi Jane, I just wanted to remind you about our meeting tomorrow at 2pm.",
			HTML:        "",
			Date:        "2022-03-02T10:00:00Z",
			IsSeen:      false,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/5/email2",
			From:        "jane@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "RE: Meeting Tomorrow",
			Text:        "Thanks for the reminder, John. I'll be there at 2pm tomorrow.",
			HTML:        "",
			Date:        "2022-03-02T11:00:00Z",
			IsSeen:      false,
			IsImportant: false,
			IsAnswered:  true,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/5/email3",
			From:        "bob@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Project Update",
			Text:        "Hi John, just wanted to give you a quick update on the project. We're making good progress and should be able to hit our deadline.",
			HTML:        "",
			Date:        "2022-03-01T15:30:00Z",
			IsSeen:      true,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/5/email4",
			From:        "alice@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Important Announcement",
			Text:        "Hello, we have an important company-wide announcement. Please see the attached document for details.",
			HTML:        "",
			Date:        "2022-02-28T09:15:00Z",
			IsSeen:      true,
			IsImportant: true,
			IsAnswered:  false,
			IsSelected:  false,
		},
		{
			Id:          rand.Intn(100),
			Path:        "/5/email5",
			From:        "sales@example.com",
			To:          "john@example.com",
			DeliveredTo: "john@example.com",
			Subject:     "Special Offer",
			Text:        "Don't miss out on our special offer! Use code SPECIAL10 for 10% off your next purchase.",
			HTML:        "",
			Date:        "2022-02-25T14:00:00Z",
			IsSeen:      true,
			IsImportant: false,
			IsAnswered:  false,
			IsSelected:  false,
		},
	}
	emailAccounts := []EmailAccount{
		{
			AccountType:       "regular",
			ShortName:         "Inbox",
			Email:             "john.smith@example.com",
			InboxShellCommand: "ls /path/to/inbox",
			TrashFolder:       "/path/to/trash",
			DeliveryTransport: "smtp",
			Emails:            emails1,
		},
		{
			AccountType:       "virtual",
			ShortName:         "Important",
			Email:             "jane.doe@example.com",
			InboxShellCommand: "notmuch search --output=files tag:inbox",
			TrashFolder:       "/path/to/trash",
			DeliveryTransport: "smtp",
			Emails:            emails2,
		},
		{
			AccountType:       "regular",
			ShortName:         "Work",
			Email:             "bob.johnson@example.com",
			InboxShellCommand: "ls /path/to/inbox",
			TrashFolder:       "/path/to/trash",
			DeliveryTransport: "smtp",
			Emails:            emails3,
		},
		{
			AccountType:       "virtual",
			ShortName:         "Home",
			Email:             "alice.williams@example.com",
			InboxShellCommand: "notmuch search --output=files tag:inbox",
			TrashFolder:       "/path/to/trash",
			DeliveryTransport: "smtp",
			Emails:            emails4,
		},
		{
			AccountType:       "regular",
			ShortName:         "Football",
			Email:             "charlie.brown@example.com",
			InboxShellCommand: "ls /path/to/inbox",
			TrashFolder:       "/path/to/trash",
			DeliveryTransport: "smtp",
			Emails:            emails5,
		},
	}
	return emailAccounts
}
