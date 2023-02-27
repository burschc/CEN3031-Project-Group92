import argparse
import os
import subprocess
import platform

parser = argparse.ArgumentParser("run_venv")
parser.add_argument("filename", help="is the python file to run in the app's python virtual environment", type=str)
parser.add_argument("args", help="are the arguments you want to run the python script with.", nargs="+", default=[])
args = parser.parse_args()

curr_os = platform.system()
file = os.path.basename(__file__)

if curr_os == "Windows":
    print(f"({file}) Detected OS as Windows.")
    venv_python = "python/venv/Scripts/python"
elif curr_os == "Linux" or curr_os == "Darwin":
    print(f"({file}) Detected OS as Linux/Darwin.")
    venv_python = "python/venv/bin/python"
else:
    print(f"({file}) Current operating system's python location is not defined! Using linux default...")
    venv_python = "python/venv/bin/python"

subprocess.run([venv_python, "python/" + args.filename])
