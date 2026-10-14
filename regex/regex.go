package regex

import (
	"regexp"
)

// 这里写好正则 ，在utils.go中进行结合判断，然后在proxy.go中执行utils
func SqlInjection(req1, req2 string) (bool, string) {
	sqliPatterns := []string{
		`@rx (?i:\b(?:(?:m(?:s(?:ys(?:ac(?:cess(?:objects|storage|xml)|es)|(?:relationship|object|querie)s|modules2?)|db)|aster\.\.sysdatabases|ysql\.db)|pg_(?:catalog|toast)|information_schema|northwind|tempdb)\b|s(?:(?:ys(?:\.database_name|aux)|qlite(?:_temp)?_master)\b|chema(?:_name\b|\W*\())|d(?:atabas|b_nam)e\W*\())`,
		`@rx (?i:(?:select|;)\s+(?:benchmark|sleep|if)\s*?\(\s*?\(?\s*?\w+)`,
		// `@rx (?i:(?:[\"'`](?:;?\s*?(?:having|select|union)\b\s*?[^\s]|\s*?!\s*?[\"'`\w])|(?:c(?:onnection_id|urrent_user)|database)\s*?\([^\)]*?|u(?:nion(?:[\w(\s]*?select| select @)|ser\s*?\([^\)]*?)|s(?:chema\s*?\([^\)]*?|elect.*?\w?user\()|into[\s+]+(?:dump|out)file\s*?[\"'`]|\s*?exec(?:ute)?.*?\Wxp_cmdshell|from\W+information_schema\W|exec(?:ute)?\s+master\.|\wiif\s*?\())`,
	}
	combined := req1 + " " + req2
	for _, pattern := range sqliPatterns {
		matched, _ := regexp.MatchString(pattern, combined)
		if matched {
			return true, "SQL injection"
		}
	}
	return false, ""
}

func RCE(req1, req2 string) (bool, string) {
	// 定义一系列用于检测 RCE 攻击的正则表达式
	rcePatterns := []string{
		`(?i)\b(eval|assert|preg_replace|create_function|call_user_func|call_user_func_array|system|exec|shell_exec|passthru|popen|proc_open|pcntl_exec)\b`,
		`(?i)\b(wget|curl|python|perl|ruby|php|bash|sh|nc|netcat|telnet|ftp|tftp|base64|chmod|chown|rm|mv|cp|ls|cat|tac|more|less|head|tail|grep|find|awk|sed|whoami|uname|ifconfig|ipconfig|netstat|ping|tracert)\b`,
		`(?i)(;|\||&{1,2})\s*(?:\w|/|\.)*\b(wget|curl|python|perl|ruby|php|bash|sh|nc|netcat|telnet|ftp|tftp|base64|chmod|chown|rm|mv|cp|ls|cat|tac|more|less|head|tail|grep|find|awk|sed|id|whoami|uname|ifconfig|ipconfig|netstat|ping|tracert)\b`,
		`(?i)(\$IFS|/bin/(ba)?sh|/bin/(da)?sh|\$\(|` + "`" + `|<\(|>\()`,
		`(?i)(\becho\b\s+.*?>\s*\S*\.php|\bcat\b\s+\*|base64\s+.*?\|\s*base64)`,
	}
	combined := req1 + " " + req2
	for _, pattern := range rcePatterns {
		matched, _ := regexp.MatchString(pattern, combined)
		if matched {
			return true, "RCE"
		}
	}
	return false, ""
}
func XSS(req1, req2 string) (bool, string) {
	// 定义一系列用于检测 XSS 攻击的正则表达式
	xssPatterns := []string{
		`(?i)\bon\w+\s*=`,
		`(?i)<\s*/?\s*(script|iframe|object|embed|svg|applet|meta|link)\b`,
		`(?i)(javascript|vbscript)\s*:`,
		`(?i)\b(alert|confirm|prompt|eval|expression|document\.cookie|window\.location)\s*\(`,
		`(?i)(&#\d+;|&#x[0-9a-f]+;)`,
	}
	combined := req1 + " " + req2
	for _, pattern := range xssPatterns {
		matched, _ := regexp.MatchString(pattern, combined)
		if matched {
			return true, "XSS Attack"
		}
	}
	return false, ""
}
