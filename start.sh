#! /bin/sh

old_module_name="github.com\/drewdavis21\/go-start"

echo "Please specify a module name: "

read -r new_module_name

find . -type f \( -name '*.go' -o -name '*.mod' \) -exec sed -i'*.bak' -e "s/$old_module_name/$new_module_name/g" {} \;
find . -type f -name '*.bak' -exec rm {} \;

rm LICENSE

printf "You're ready to start!\nThe package name for this repository is now %s\n" "$new_module_name"
