# MacOS_Useridle
A simple daemon that runs in the background to check users last activity. A MQTT packet is fired every X minutes if user activity was detected.

# Building the application
You'll need go installed to compile to code. Once complied you'll have a single binary called "user_idle".

# Config
This program uses a YAML config to store params. The config file default location is the launch directory of the binary.
