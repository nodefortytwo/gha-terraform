#!/bin/sh -l
echo $INPUT_LIST
echo $INPUT_MAP


cd $INPUT_ROOT_DIR
tree
terraform init
terraform $INPUT_ACTION