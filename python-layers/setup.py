from setuptools import setup

setup(
    name="docklingcake",
    packages=["docklingcake"],
    entry_points={
        'console_scripts': [
            'example = docklingcake.example:get_example',
        ],
    },
)
