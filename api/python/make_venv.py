import os
import subprocess

file = os.path.basename(__file__)

if not os.path.exists("python/venv/"):
    print(f"({file}) Python virtual environment does not exist!")
    print(f"({file}) Making virtual environment...")
    os.system("py -m venv python/venv/")
    os.system("py -m ensurepip")

venv_python = "python/venv/Scripts/python.exe"
args = [venv_python, 'python/get_dependencies.py']
subprocess.run(args)
