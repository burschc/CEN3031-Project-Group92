import os
import subprocess
import platform

file = os.path.basename(__file__)
curr_os = platform.system()

if not os.path.exists("python/venv/"):
    print(f"({file}) Python virtual environment does not exist!")
    print(f"({file}) Making virtual environment...")
    try:
        subprocess.run(["python", "-m", "venv", "python/venv/"])
        subprocess.run(["python", "-m", "ensurepip"])
    except subprocess.CalledProcessError:
        print(f"({file}) Could not make python virtual environment!\n"
              f"Check to make sure python 3 is installed in your system or make sure it is updated.")

if curr_os == "Windows":
    print(f"({file}) Detected OS as Windows.")
    venv_python = "python/venv/Scripts/python"
elif curr_os == "Linux" or curr_os == "Darwin":
    print(f"({file}) Detected OS as Linux/Darwin.")
    venv_python = "python/venv/bin/python"
else:
    print(f"({file}) Current operating system's python location is not defined! Using linux default...")
    venv_python = "python/venv/bin/python"

args = [venv_python, 'python/get_dependencies.py']
subprocess.run(args)
