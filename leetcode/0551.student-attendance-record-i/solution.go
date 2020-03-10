package _551_student_attendance_record_i

import "strings"

func checkRecord(s string) bool {
	if strings.Count(s, "A") > 1 {
		return false
	}

	if strings.Contains(s, "LLL") {
		return false
	}
	return true
}
