package main

import "fmt"

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

func (u *user) updateStatus(status string) error {
	if status != "email_bounced" && status != "email_failed" {
		return fmt.Errorf("invalid status: %s", status)
	}
	u.status = status
	return nil
}

type analytics struct {
	totalBounces int
}

func (a *analytics) track(event string) error {
	if event != "email_bounced" {
		return fmt.Errorf("invalid event: %s", event)
	}
	a.totalBounces++
	return nil
}

// Another possible solution that wont pass the testcases but will work

// func (u *user) updateStatus(status string) (*user, error) {
// 	if status != "email_bounced" && status != "email_failed" {
// 		return nil, fmt.Errorf("invalid status: %s", status)
// 	}
// 	u.status = status
// 	return u, nil
// }

// type analytics struct {
// 	totalBounces int
// }

// func (a *analytics) track(event string) (*analytics, error) {
// 	if event != "email_bounced" {
// 		return nil, fmt.Errorf("invalid event: %s", event)
// 	}
// 	a.totalBounces++
// 	return a, nil
// }
