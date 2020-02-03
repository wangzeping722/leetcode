package _359_logger_rate_limiter


type Logger struct {
	timeMap map[string]int
}


/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{make(map[string]int)}
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := this.timeMap[message]; !ok{
		this.timeMap[message]= timestamp
		return true
	}

	if timestamp - this.timeMap[message] < 10 {
		return false
	}
	this.timeMap[message]= timestamp
	return true
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */