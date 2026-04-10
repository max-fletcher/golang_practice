package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	if err := em.recipient.updateStatus(em.status); err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}

	if err := a.track(em.status); err != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)
	}

	return nil
}

// Another possible solution that wont pass the testcases but will work

// func (a *analytics) handleEmailBounce(em email) error {
// 	_, err1 := em.recipient.updateStatus(em.status)
// 	if err1 == nil {
// 		return fmt.Errorf("error updating user status: %w", err1)
// 	}

// 	_, err2 := a.track(em.status)
// 	if err2 == nil {
// 		return fmt.Errorf("error updating user status: %w", err2)
// 	}

// 	return nil
// }

func main() {
	type testCase struct {
		email           email
		expectedError   string
		expectedStatus  string
		expectedBounces int
	}

	runCases := []testCase{
		{
			email: email{
				status:    "email_bounced",
				recipient: &user{email: "bugs@acme.inc"},
			},
			expectedError:   "<nil>",
			expectedStatus:  "email_bounced",
			expectedBounces: 1,
		},
		{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "elmer@acme.inc"},
			},
			expectedError:   "error tracking user bounce: invalid event: email_failed",
			expectedStatus:  "email_failed",
			expectedBounces: 0,
		},
		{
			email: email{
				status:    "email_sent",
				recipient: &user{email: "daffy@acme.inc"},
			},
			expectedError:   "error updating user status: invalid status: email_sent",
			expectedStatus:  "",
			expectedBounces: 0,
		},
		{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "porky@acme.inc"},
			},
			expectedError:   "error tracking user bounce: invalid event: email_failed",
			expectedStatus:  "email_failed",
			expectedBounces: 0,
		},
	}

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		a := &analytics{}
		err := a.handleEmailBounce(test.email)
		actualError := fmt.Sprintf("%v", err)
		if actualError != test.expectedError {
			failCount++
			fmt.Printf(`---------------------------------
				Test Failed:
				status:    %v
				recipient: %v
				expected error:   %v
				actual error:     %v
				`, test.email.status, test.email.recipient.email, test.expectedError, actualError)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
				Test Passed:
				status:    %v
				recipient: %v
				expected error:   %v
				actual error:     %v
				`, test.email.status, test.email.recipient.email, test.expectedError, actualError)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
