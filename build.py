#!/usr/bin/env python3

import sys
import os
from os import path
import subprocess

def run_cmd(cmd, name):
	proc = subprocess.Popen(cmd, stdout = subprocess.PIPE, shell = True)
	output = proc.communicate()[0].decode("utf-8").strip()
	status = proc.wait()
	if output != "":
		print(output)
	if status != 0:
		print("Failed building", name)
		exit(status)

def build(which):
	exe = subprocess.check_output("go env GOEXE", shell = True).decode("utf-8").strip()
	server_cmd = "go build -o rpcserver" + exe + " ./cmd/rpcserver"
	client_cmd = "go build -o rpcclient" + exe + " ./cmd/rpcclient"
	if which == "server":
		run_cmd(server_cmd, which)
	elif which == "client":
		run_cmd(client_cmd, which)
	else:
		run_cmd(server_cmd, "server")
		run_cmd(client_cmd, "client")

	print("Building finished successfully, run ./rpcserver in one terminal window and ./rpcclient in another (in that order)")

if __name__ == "__main__":
	cmd = ""
	try:
		cmd = sys.argv.pop(1)
	except Exception:
		cmd = sys.argv.pop(0)

	if cmd != "server" and cmd != "client" and cmd != "both" and cmd != "clean":
		print("usage: python build.py","[server|client|both|clean]", file=sys.stderr)
		exit(code = 1)

	if os.environ.get("GOPATH") == None:
		print("GOPATH environment variable not set", file = sys.stderr)
		exit(code = 1)
	
	if cmd == "clean":
		print("Cleaning up")
		del_files = ["rpcclient", "rpcclient.exe", "rpcserver", "rpcserver.exe"]
		for del_file in del_files:
			if path.exists(del_file):
				os.remove(del_file)
	else:
		build(cmd)
