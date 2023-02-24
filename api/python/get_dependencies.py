import os
import sys
import subprocess

file = os.path.basename(__file__)
pip = "python/venv/Scripts/pip"
py = "python/venv/Scripts/python"

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
