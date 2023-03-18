#!/bin/bash

if [ -z "$1" ] && [ -z "$2" ] && [ -z "$3" ]; then
  echo "[Error] masukan nama parent folder, username facebook, username linkedin!"
  exit 1
fi

parrent_path="$1 $(date)"

mkdir "$parrent_path"

mkdir -p "$parrent_path/about_me/"{personal,professional}

mkdir -p "$parrent_path/"{my_friends,my_system_info}

echo "https://www.facebook.com/$2" >> "$parrent_path/about_me/personal/facebook.txt"

echo "https://www.linkedin.com/in/$3" >> "$parrent_path/about_me/professional/linkedin.txt"

curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt >> "$parrent_path/my_friends/list_of_my_friends.txt"

echo -e "My username: $(whoami)\nWith host: $(hostname) $(uname -a)" >> "$parrent_path/my_system_info/about_this_laptop.txt"

ping -c 3 google.com >> "$parrent_path/my_system_info/internet_connection.txt"
