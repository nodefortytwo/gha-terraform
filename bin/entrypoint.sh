#!/bin/sh -l
pwd
tree
cd $INPUT_ROOT_DIR
tree
terraform init
terraform $INPUT_ACTION