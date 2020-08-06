#coding=utf-8
import sys,re

# re map
reMap = {
	".":  "\\.",
	"+":  "\\+",
	"-":  "\\-",
	"[":  "\\[",
	"]":  "\\]",
	"?":  "\\?",
	"\"": "\\\"",
	"‘":  "\\'",
	"/":  "\\/",
	"(":  "\\(",
	")":  "\\)",
	"\\": "\\\\",
	"|":  "\\|",
	"^":  "\\^",
	"$":  "\\$",
}

# re -> string
def convertStr(args):
	for arg in args:
		d = ""
		for a in arg:
			if a in reMap:
				d += reMap[a]
			else:
				d += a
		print(d)

def convertRe(args):
	for arg in args:
		d = arg
		for k,v in reMap.items():
			if v in arg:
				d.replace(v,k)
		print(d)

if __name__ == '__main__':
	# -s re -> str
	# -r str -> re
	c = sys.argv[1] 
	args = sys.argv[2:]

	if c == "-r":
		convertRe(args)
	elif c == "-s":
		convertStr(args)
	else :
		print("Invalid args, -s or -r")
		sys.exit(1)