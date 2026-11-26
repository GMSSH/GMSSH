# -*- encoding: utf-8 -*-
from setuptools import find_packages, setup

setup(
    version="1.0.8",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown; charset=UTF-8",
    keywords=[
        "gm",
        "gm-scaffold",
        "simplejrpc",
        "gmssh",
    ],
    python_requires=">=3.10",
    packages=find_packages(),
    # exclude_package_date={"": [".gitignore"]},
    install_requires=[
        "pydispatch==1.1.0",
        "pyfiglet==1.0.2",
        "jinja2==3.1.4",
        "fire==0.7.0",
        "cleo==2.1.0",
    ],
)
