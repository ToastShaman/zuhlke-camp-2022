#!/usr/bin/env python3

import glob
import os
import shutil
import subprocess
import jinja2
from tqdm import tqdm
from jinja2 import Template

def split(path_to_file):
    file_name = os.path.basename(path_to_file)
    file_name_without_extension = os.path.splitext(file_name)[0]
    return file_name, file_name_without_extension


if __name__ == "__main__":
    output_dir = os.path.join("build")
    shutil.rmtree(output_dir, ignore_errors=True)
    os.mkdir(output_dir)

    # Render the complete guide
    with open('./src/pair-programming-guide.md', 'w') as f:
        template_env = jinja2.Environment(loader=jinja2.FileSystemLoader(searchpath="./src/"))
        template = template_env.get_template("pair-programming-guide.jinja")
        f.write(template.render())

    # Convert markdown to PDF
    files = glob.glob("./src/**/*.md", recursive=True)
    for file in tqdm(files):
        cwd = os.path.join(os.getcwd(), "src")
        file_name, file_name_without_extension = split(file)
        
        if file.endswith(".md"):
            output_md = os.path.join(output_dir, f"{file_name_without_extension}.md")
            subprocess.check_output(f"docker run -it --rm -v {cwd}/../:/data minlag/mermaid-cli -i /data/src/{file_name} -o /data/{output_md} -t dark -b transparent -s 3", shell=True)        
            subprocess.check_output(f"npx md-to-pdf {output_md}", shell=True)
    
    # Cleanup temporary files
    for ext in ["*.md", "*.png", "*.svg"]:
        for f in glob.glob(f"{output_dir}/**/{ext}", recursive=True):
            os.remove(f)