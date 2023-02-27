import os
import platform
import sys
import subprocess

file = os.path.basename(__file__)
curr_os = platform.system()

if curr_os == "Windows":
    print(f"({file}) Detected OS as Windows.")
    pip = "python/venv/Scripts/pip"
    py = "python/venv/Scripts/python"
elif curr_os == "Linux" or curr_os == "Darwin":
    print(f"({file}) Detected OS as Linux/Darwin.")
    pip = "python/venv/bin/pip"
    py = "python/venv/bin/python"
else:
    print(f"({file}) Current operating system's python and pip location is not defined! Using linux default...")
    pip = "python/venv/bin/pip"
    py = "python/venv/bin/python"

print(f"({file}) Checking Python dependencies for UF Parking Map Plus...")

dependencies = ["gjf"]

subprocess.run(f"{py} -m pip install --upgrade pip")

for dep in dependencies:
    if dep in sys.modules:
        print(f"({file}) {dep} is already in the virtual environment. Checking for updates...")
        subprocess.run(f"{pip} --upgrade {dep}")
    else:
        subprocess.run(f"{pip} install {dep}")

print(f"({file}) All dependencies were verified or installed.\n")
