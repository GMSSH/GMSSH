# -*- encoding: utf-8 -*-
from setuptools import find_packages, setup

setup(
    version="2.1.7",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown; charset=UTF-8",
    keywords=[
        "gm",
        "gm-sdk",
        "simplejrpc",
        "gmssh",
        "jsonrpc",
        "jsonrpcserver",
        "jsonrpcclient",
    ],
    python_requires=">=3.10",
    packages=find_packages(),
    # exclude_package_date={"": [".gitignore"]},
    install_requires=[
        "jsonrpcclient==4.0.3",
        "jsonrpcserver==5.0.9",
        "loguru==0.7.3",
        "PyYAML==6.0.2",
        "WTForms==3.2.1",
    ],
)
