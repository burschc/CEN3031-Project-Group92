import argparse
import subprocess

parser = argparse.ArgumentParser("run_venv")
parser.add_argument("filename", help="is the python file to run in the app's python virtual environment", type=str)
parser.add_argument("args", help="are the arguments you want to run the python script with.", nargs="+", default=[])
args = parser.parse_args()

venv_python = "python/venv/Scripts/python"
subprocess.run([venv_python, "python/" + args.filename])
